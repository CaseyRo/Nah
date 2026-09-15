// CDI-1835 with the phones taken out, in the app's own code: two identities
// against a real apps/server, the second joining through the first's
// invitation, each posting, each reading both.
//
// The first person needs an operator's invitation, which is single-use, so make
// a fresh one for every run, from the same directory the server runs in:
//
//   (cd ../server && go run ./cmd/server) &
//   NAH_SERVER=http://127.0.0.1:8080 \
//   NAH_INVITE="$(cd ../server && go run ./cmd/server invite)" \
//   flutter test test/two_people_test.dart
import 'dart:io';

import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:nah/server.dart';

void main() {
  final url = Platform.environment['NAH_SERVER'];
  final operatorInvitation = Platform.environment['NAH_INVITE'];

  test(
    'one person joins through the other, and each reads both moments',
    () async {
      Future<Server> phone(Future<String> Function() invitation) async {
        // Each phone has its own keychain, but the mock is one map: empty it
        // before each identity is made. A joined Server keeps its keys and
        // session in memory, so the first phone is unaffected.
        FlutterSecureStorage.setMockInitialValues({});
        final server = Server(dio: Dio(BaseOptions(baseUrl: url!)));
        await expectLater(server.start(), throwsA(isA<NeedsInvitation>()));
        await server.join(await invitation());
        return server;
      }

      final a = await phone(() async => operatorInvitation!);
      final b = await phone(a.invitation);
      expect(a.me, isNot(b.me));

      await a.post('from A');
      await b.post('from B');

      for (final reader in [a, b]) {
        final feed = await reader.feed();
        expect(feed.map((m) => m.text), ['from B', 'from A']);
        expect(feed.map((m) => m.author), [b.me, a.me]);
      }
    },
    skip: url == null || operatorInvitation == null
        ? 'set NAH_SERVER and NAH_INVITE; see the comment at the top'
        : false,
  );
}
