import 'dart:async';

import 'package:flutter_test/flutter_test.dart';
import 'package:nah/main.dart';
import 'package:nah/server.dart';

/// A server that answers however the test needs, without a socket.
class _FakeServer implements Server {
  _FakeServer(this._answer);
  final Future<bool> _answer;
  @override
  Future<bool> healthy() => _answer;
}

void main() {
  testWidgets('says it is looking while the answer is outstanding', (
    tester,
  ) async {
    await tester.pumpWidget(
      NahApp(server: _FakeServer(Completer<bool>().future)),
    );
    expect(find.text('Looking for your circle…'), findsOneWidget);
  });

  testWidgets('a reachable server', (tester) async {
    await tester.pumpWidget(NahApp(server: _FakeServer(Future.value(true))));
    await tester.pumpAndSettle();
    expect(find.text('Your circle is there.'), findsOneWidget);
  });

  testWidgets('an unreachable server is said in words, not in a stack trace', (
    tester,
  ) async {
    await tester.pumpWidget(NahApp(server: _FakeServer(Future.value(false))));
    await tester.pumpAndSettle();
    expect(find.text('Cannot reach your circle right now.'), findsOneWidget);
    // Nothing technical ever reaches a person's screen.
    expect(find.textContaining('127.0.0.1'), findsNothing);
    expect(find.textContaining('Exception'), findsNothing);
  });
}
