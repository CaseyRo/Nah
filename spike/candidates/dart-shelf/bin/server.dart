// Candidate D: Dart, shelf, sqlite3. The challenger, because the client is
// Flutter and one language across the stack is worth real money to a solo
// developer.
import 'dart:convert';
import 'dart:io';
import 'package:shelf/shelf.dart';
import 'package:shelf/shelf_io.dart' as io;
import 'package:sqlite3/sqlite3.dart';

late final Database db;

Response _json(Object? body, {int status = 200}) => Response(status,
    body: jsonEncode(body), headers: {'content-type': 'application/json'});

Future<Response> _handler(Request req) async {
  if (req.method == 'POST' && req.url.path == 'moments') {
    final d = jsonDecode(await req.readAsString()) as Map<String, dynamic>;
    final stmt = db.prepare(
        'INSERT INTO moments (circle_id, author_id, created_at, blob) VALUES (?,?,?,?)');
    stmt.execute([d['circle_id'], d['author_id'], d['created_at'], d['blob']]);
    stmt.dispose();
    return _json({'ok': true}, status: 201);
  }
  if (req.method == 'GET' && req.url.path == 'feed') {
    final p = req.url.queryParameters;
    var limit = int.tryParse(p['limit'] ?? '') ?? 30;
    if (limit <= 0 || limit > 200) limit = 30;
    final rows = db.select(
        'SELECT id, circle_id, author_id, created_at, blob FROM moments '
        'WHERE circle_id = ? ORDER BY created_at DESC LIMIT ?',
        [p['circle_id'], limit]);
    return _json(rows
        .map((r) => {
              'id': r['id'], 'circle_id': r['circle_id'], 'author_id': r['author_id'],
              'created_at': r['created_at'], 'blob': r['blob'],
            })
        .toList());
  }
  return Response.notFound('no');
}

void main() async {
  db = sqlite3.open(Platform.environment['DB'] ?? 'data.db');
  db.execute('PRAGMA journal_mode=WAL');
  db.execute('PRAGMA synchronous=NORMAL');
  db.execute('PRAGMA busy_timeout=5000');
  db.execute('''CREATE TABLE IF NOT EXISTS moments (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      circle_id TEXT NOT NULL, author_id TEXT NOT NULL,
      created_at INTEGER NOT NULL, blob TEXT NOT NULL)''');
  db.execute(
      'CREATE INDEX IF NOT EXISTS idx_moments_circle_created ON moments (circle_id, created_at DESC)');
  final addr = Platform.environment['ADDR'] ?? '127.0.0.1:8093';
  final parts = addr.split(':');
  await io.serve(_handler, parts[0], int.parse(parts[1]));
}
