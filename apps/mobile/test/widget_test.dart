import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:nah/main.dart';
import 'package:nah/moment.dart';
import 'package:nah/server.dart';

/// A server that answers however the test needs, without a socket.
class _FakeServer implements Server {
  _FakeServer({this.moments = const []});

  List<Moment> moments;
  Completer<void>? hold;
  bool reachable = true;
  bool joined = true;
  String? refusal;
  final posted = <String>[];
  final connected = <String>[];

  @override
  String get me => 'me';

  @override
  Future<void> start() async {
    await hold?.future;
    if (!reachable) {
      throw Exception('SocketException: connection refused (127.0.0.1:8080)');
    }
    if (!joined) throw const NeedsInvitation();
  }

  @override
  Future<void> join(String invitation) async {
    if (refusal != null) throw Refused(refusal!);
    joined = true;
  }

  @override
  Future<List<Moment>> feed() async => moments;

  @override
  Future<void> post(String text) async {
    posted.add(text);
    moments = [
      Moment(author: me, id: posted.length, at: DateTime.now(), text: text),
      ...moments,
    ];
  }

  @override
  Future<String> invitation() async => 'me#an-invite';

  @override
  Future<void> connect(String invitation) async {
    if (refusal != null) throw Refused(refusal!);
    connected.add(invitation);
  }
}

void main() {
  testWidgets('says it is looking while it signs in', (tester) async {
    await tester.pumpWidget(NahApp(server: _FakeServer()..hold = Completer()));
    expect(find.text('Looking for your people…'), findsOneWidget);
  });

  testWidgets('an unreachable server is said in words, not in a stack trace', (
    tester,
  ) async {
    await tester.pumpWidget(NahApp(server: _FakeServer()..reachable = false));
    await tester.pumpAndSettle();
    expect(find.text('Cannot reach your people right now.'), findsOneWidget);
    // Nothing technical ever reaches a person's screen.
    expect(find.textContaining('127.0.0.1'), findsNothing);
    expect(find.textContaining('Exception'), findsNothing);
  });

  testWidgets('a phone that has not joined asks for an invitation, and joins '
      'with one', (tester) async {
    final server = _FakeServer()..joined = false;
    await tester.pumpWidget(NahApp(server: server));
    await tester.pumpAndSettle();
    expect(
      find.text('Nah? is by invitation. Paste the one someone sent you.'),
      findsOneWidget,
    );
    // There is nobody to connect with until this phone is someone.
    expect(find.byTooltip('Connect'), findsNothing);

    final join = find.widgetWithText(FilledButton, 'Join');
    server.refusal = 'This invitation is no longer valid.';
    await tester.enterText(find.byType(TextField), 'someone#their-invite');
    await tester.tap(join);
    await tester.pumpAndSettle();
    expect(find.text('This invitation is no longer valid.'), findsOneWidget);

    server.refusal = null;
    await tester.tap(join);
    await tester.pumpAndSettle();
    expect(server.joined, isTrue);
    expect(find.text('Say something to your people'), findsOneWidget);
    expect(find.byTooltip('Connect'), findsOneWidget);
  });

  testWidgets('the feed is shown in the order the server gives, yours marked', (
    tester,
  ) async {
    final now = DateTime.now();
    await tester.pumpWidget(
      NahApp(
        server: _FakeServer(
          moments: [
            Moment(author: 'them', id: 1, at: now, text: 'from them'),
            Moment(author: 'me', id: 1, at: now, text: 'from me'),
          ],
        ),
      ),
    );
    await tester.pumpAndSettle();
    expect(
      tester.getTopLeft(find.text('from them')).dy,
      lessThan(tester.getTopLeft(find.text('from me')).dy),
    );
    expect(find.textContaining('You ·'), findsOneWidget);
  });

  testWidgets('posting puts the moment at the top and empties the field', (
    tester,
  ) async {
    final server = _FakeServer(
      moments: [
        Moment(author: 'them', id: 1, at: DateTime.now(), text: 'earlier'),
      ],
    );
    await tester.pumpWidget(NahApp(server: server));
    await tester.pumpAndSettle();

    await tester.enterText(find.byType(TextField), 'The cake survived.');
    await tester.tap(find.byTooltip('Send'));
    await tester.pumpAndSettle();

    expect(server.posted, ['The cake survived.']);
    // Found once: in the feed, and no longer in the field.
    expect(find.text('The cake survived.'), findsOneWidget);
    expect(
      tester.getTopLeft(find.text('The cake survived.')).dy,
      lessThan(tester.getTopLeft(find.text('earlier')).dy),
    );
  });

  testWidgets('connecting: an invitation to copy, a refusal in the server\'s '
      'own words, then a connection', (tester) async {
    final server = _FakeServer();
    await tester.pumpWidget(NahApp(server: server));
    await tester.pumpAndSettle();

    await tester.tap(find.byTooltip('Connect'));
    await tester.pumpAndSettle();
    expect(find.text('me#an-invite'), findsOneWidget);

    final theirs = find.byWidgetPredicate(
      (w) => w is TextField && w.decoration?.hintText == 'Their invitation',
    );
    final connect = find.widgetWithText(FilledButton, 'Connect');

    server.refusal = 'There is no room for this connection right now.';
    await tester.enterText(theirs, 'them#their-invite');
    await tester.tap(connect);
    await tester.pumpAndSettle();
    expect(
      find.text('There is no room for this connection right now.'),
      findsOneWidget,
    );

    server.refusal = null;
    await tester.tap(connect);
    await tester.pumpAndSettle();
    expect(find.text('Connected.'), findsOneWidget);
    expect(server.connected, ['them#their-invite']);
  });
}
