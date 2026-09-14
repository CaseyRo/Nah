import 'dart:convert';
import 'dart:typed_data';

/// How a blob is sealed, in its first byte, so reading never depends on
/// guessing. M1 moments travel unsealed until each person has a content key
/// (CDI-1863). Sealed moments will take the next number, and everything posted
/// before then stays readable with no migration.
const _unsealed = 0;

/// What a reader will accept. A decrypted envelope is hostile input (ADR-0016),
/// and an unsealed one doubly so.
const maxEnvelopeBytes = 64 * 1024;
const maxTextLength = 4000;
const _maxFallbackLength = 500;

const unreadableMoment = 'This moment could not be read.';

/// A moment as the app shows it.
class Moment {
  const Moment({
    required this.author,
    required this.id,
    required this.at,
    required this.text,
  });

  /// One row of the server's feed.
  factory Moment.fromJson(Map<String, dynamic> json) {
    List<int> blob;
    try {
      blob = base64.decode(json['blob'] as String);
    } on FormatException {
      blob = const [];
    }
    return Moment(
      author: json['author_id'] as String,
      id: json['id'] as int,
      at: DateTime.fromMillisecondsSinceEpoch(json['created_at'] as int),
      text: open(blob),
    );
  }

  /// The author's person id.
  final String author;

  /// Unique among its author's moments, not across a feed.
  final int id;

  /// When the server took it, which is what the feed is ordered by.
  final DateTime at;

  /// The words to show: the moment itself, the fallback its author's app wrote,
  /// or [unreadableMoment].
  final String text;
}

/// Writes a text moment as a version 1 envelope, unsealed for now.
Uint8List seal(String text, DateTime now) {
  final runes = text.runes;
  final envelope = {
    'v': 1,
    'type': 'text',
    'fallback': runes.length <= _maxFallbackLength
        ? text
        : '${String.fromCharCodes(runes.take(_maxFallbackLength - 1))}…',
    'sent_at': now.millisecondsSinceEpoch,
    'body': {'text': text},
  };
  return Uint8List.fromList([_unsealed, ...utf8.encode(jsonEncode(envelope))]);
}

/// Reads a blob as hostile input. Anything that is not a well-formed, bounded
/// envelope comes back as [unreadableMoment], never as an exception.
String open(List<int> blob) {
  if (blob.isEmpty ||
      blob.length > maxEnvelopeBytes + 1 ||
      blob.first != _unsealed) {
    return unreadableMoment;
  }
  final Object? envelope;
  try {
    envelope = jsonDecode(utf8.decode(blob.sublist(1)));
  } on FormatException {
    return unreadableMoment;
  }
  if (envelope is! Map<String, dynamic>) return unreadableMoment;

  final body = envelope['body'];
  if (envelope['v'] == 1 &&
      envelope['type'] == 'text' &&
      body is Map<String, dynamic>) {
    final text = body['text'];
    if (text is String && text.isNotEmpty && text.length <= maxTextLength) {
      return text;
    }
  }

  // A version or a type this app does not know: show the sentence the author's
  // newer app wrote for exactly this case.
  final fallback = envelope['fallback'];
  if (fallback is String &&
      fallback.isNotEmpty &&
      fallback.runes.length <= _maxFallbackLength) {
    return fallback;
  }
  return unreadableMoment;
}
