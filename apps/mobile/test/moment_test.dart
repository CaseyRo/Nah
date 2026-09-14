import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:nah/moment.dart';

List<int> unsealed(Object? envelope) => [
  0,
  ...utf8.encode(jsonEncode(envelope)),
];

Map<String, dynamic> envelopeOf(List<int> blob) =>
    jsonDecode(utf8.decode(blob.sublist(1))) as Map<String, dynamic>;

void main() {
  test('a text moment reads back as written', () {
    final blob = seal('The cake survived the drive.', DateTime(2026, 9, 14));
    expect(open(blob), 'The cake survived the drive.');
  });

  test('the envelope is versioned and carries its own fallback', () {
    final envelope = envelopeOf(seal('hello', DateTime(2026, 9, 14)));
    expect(envelope['v'], 1);
    expect(envelope['type'], 'text');
    expect(envelope['fallback'], 'hello');
    expect(envelope['body'], {'text': 'hello'});
  });

  test('a long moment still gets a short fallback', () {
    final envelope = envelopeOf(seal('word ' * 1000, DateTime(2026, 9, 14)));
    expect((envelope['fallback'] as String).runes.length, 500);
  });

  test('a type or version this app does not know shows the fallback', () {
    expect(
      open(
        unsealed({
          'v': 1,
          'type': 'hologram',
          'fallback': 'Anna sent a hologram.',
          'body': <String, dynamic>{},
        }),
      ),
      'Anna sent a hologram.',
    );
    expect(
      open(
        unsealed({
          'v': 2,
          'type': 'text',
          'fallback': 'Written by a newer app.',
          'body': {'text': 'hi'},
        }),
      ),
      'Written by a newer app.',
    );
  });

  test('every broken or hostile blob is unreadable, never an exception', () {
    final blobs = <String, List<int>>{
      'empty': [],
      'a seal byte and nothing': [0],
      'a seal this app cannot open yet': [1, ...utf8.encode('{"v":1}')],
      'not UTF-8': [0, 0xff, 0xfe],
      'not JSON': [0, ...utf8.encode('{"v":')],
      'JSON, but not an object': unsealed(['an', 'array']),
      'text that is not a string': unsealed({
        'v': 1,
        'type': 'text',
        'body': {'text': 42},
      }),
      'empty text': unsealed({
        'v': 1,
        'type': 'text',
        'body': {'text': ''},
      }),
      'text past the limit': unsealed({
        'v': 1,
        'type': 'text',
        'body': {'text': 'x' * (maxTextLength + 1)},
      }),
      'an unknown type with no fallback': unsealed({
        'v': 1,
        'type': 'hologram',
      }),
      'a fallback past the limit': unsealed({
        'v': 1,
        'type': 'hologram',
        'fallback': 'x' * 10000,
      }),
      'past the size limit': [0, ...List.filled(maxEnvelopeBytes + 1, 0x20)],
      'nested deep enough to hurt': [0, ...utf8.encode('[' * 60000)],
    };
    blobs.forEach((name, blob) {
      expect(open(blob), unreadableMoment, reason: name);
    });
  });
}
