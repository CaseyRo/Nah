// CDI-1835 with the phones taken out, in the app's own code: two identities
// against a real apps/server, connected, each posting, each reading both.
//
//   (cd ../server && go run ./cmd/server) &
//   NAH_SERVER=http://127.0.0.1:8080 flutter test test/two_people_test.dart
import 'dart:io';

import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:nah/server.dart';

void main() {
  final url = Platform.environment['NAH_SERVER'];

  test('two people connect, and each reads both moments', () async {
    Future<Server> phone() async {
      // Each phone has its own keychain, but the mock is one map: empty it
      // before each identity is made. A started Server keeps its keys in
      // memory, so the first phone is unaffected.
      FlutterSecureStorage.setMockInitialValues({});
      final server = Server(dio: Dio(BaseOptions(baseUrl: url!)));
      await server.start();
      return server;
    }

    final a = await phone();
    final b = await phone();
    expect(a.me, isNot(b.me));

    await b.connect(await a.invitation());
    await a.post('from A');
    await b.post('from B');

    for (final reader in [a, b]) {
      final feed = await reader.feed();
      expect(feed.map((m) => m.text), ['from B', 'from A']);
      expect(feed.map((m) => m.author), [b.me, a.me]);
    }
  }, skip: url == null ? 'set NAH_SERVER to a running apps/server' : false);
}
