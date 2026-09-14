import 'dart:convert';
import 'dart:math';

import 'package:cryptography/cryptography.dart';
import 'package:dio/dio.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'moment.dart';

/// The Nah? server this build talks to.
///
/// A person never types this and never sees it. It is a build-time setting for
/// now; once the directory exists (CDI-1838) the address is looked up, because
/// connecting must never involve a hostname.
const serverBaseUrl = String.fromEnvironment(
  'NAH_SERVER',
  defaultValue: 'http://127.0.0.1:8080',
);

/// A refusal the server explained, already in words a person can read.
class Refused implements Exception {
  const Refused(this.message);
  final String message;

  @override
  String toString() => message;
}

/// Everything the app knows about talking to the server: who this device is,
/// how it signs in, and the few things a person can do in M1.
///
/// Still thin on purpose. Sessions per server, the local cache and encrypting on
/// the way out are CDI-1861, in a package the app owns.
class Server {
  Server({Dio? dio, FlutterSecureStorage? storage})
    : _dio =
          dio ??
          Dio(
            BaseOptions(
              baseUrl: serverBaseUrl,
              connectTimeout: const Duration(seconds: 10),
              receiveTimeout: const Duration(seconds: 20),
              responseType: ResponseType.json,
            ),
          ),
      _storage =
          storage ??
          const FlutterSecureStorage(
            // The identity key belongs to this device and never leaves it
            // (ADR-0015), so it stays out of backups and off other devices.
            iOptions: IOSOptions(
              accessibility: KeychainAccessibility.first_unlock_this_device,
            ),
          );

  static const _seedKey = 'nah.identity.seed';
  static const _personKey = 'nah.identity.person';

  final Dio _dio;
  final FlutterSecureStorage _storage;
  final _ed25519 = Ed25519();

  late SimpleKeyPair _keys;
  late List<int> _publicKey;
  String? _person;
  String? _token;

  /// This device's person id, once [start] has run.
  String get me => _person!;

  /// Loads this device's identity, making one on first run, and signs in.
  Future<void> start() async {
    var seed = await _storage.read(key: _seedKey);
    if (seed == null) {
      final random = Random.secure();
      seed = base64.encode(List<int>.generate(32, (_) => random.nextInt(256)));
      await _storage.write(key: _seedKey, value: seed);
    }
    _keys = await _ed25519.newKeyPairFromSeed(base64.decode(seed));
    _publicKey = (await _keys.extractPublicKey()).bytes;

    _person = await _storage.read(key: _personKey);
    if (_person == null) await _register();
    try {
      await _signIn();
    } on DioException catch (e) {
      // The server no longer knows this person, which is what a wiped
      // development server looks like. Become a new person rather than stick.
      final status = e.response?.statusCode;
      if (status != 403 && status != 404) rethrow;
      await _register();
      await _signIn();
    }
  }

  Future<void> _register() async {
    final res = await _dio.post<Map<String, dynamic>>(
      '/v1/people',
      data: {'public_key': base64.encode(_publicKey)},
    );
    _person = res.data!['id'] as String;
    await _storage.write(key: _personKey, value: _person);
  }

  Future<void> _signIn() async {
    final key = base64.encode(_publicKey);
    final challenge =
        (await _dio.post<Map<String, dynamic>>(
              '/v1/people/$_person/challenge',
              data: {'public_key': key},
            )).data!['challenge']
            as String;
    // Exactly the string apps/server/README.md says a device signs.
    final signature = await _ed25519.sign(
      utf8.encode('nah-auth-v1:$_person:$challenge'),
      keyPair: _keys,
    );
    final session = await _dio.post<Map<String, dynamic>>(
      '/v1/people/$_person/session',
      data: {
        'public_key': key,
        'challenge': challenge,
        'signature': base64.encode(signature.bytes),
      },
    );
    _token = session.data!['token'] as String;
  }

  /// The newest moments from this person and their people, one page and no
  /// more (CDI-1833).
  Future<List<Moment>> feed() async {
    final res = await _authed<List<dynamic>>('GET', '/v1/people/$_person/feed');
    return [
      for (final row in res.data!) Moment.fromJson(row as Map<String, dynamic>),
    ];
  }

  /// Posts a text moment (CDI-1834). Moments are immutable, so this is the only
  /// write there is.
  Future<void> post(String text) => _refusable(
    () => _authed<Object?>(
      'POST',
      '/v1/people/$_person/moments',
      data: {'blob': base64.encode(seal(text, DateTime.now()))},
    ),
  );

  /// What someone needs to connect to this person: who, then the secret after a
  /// `#`, the shape CDI-1836 gives the link. The touch and the link themselves
  /// are M2.
  Future<String> invitation() async {
    final res = await _authed<Map<String, dynamic>>(
      'POST',
      '/v1/people/$_person/invites',
    );
    return '$_person#${res.data!['invite']}';
  }

  /// Redeems someone's invitation. Connection is mutual and immediate.
  Future<void> connect(String invitation) {
    final (person, invite) = switch (invitation.trim().split('#')) {
      [final p, final i] when p.isNotEmpty && i.isNotEmpty => (p, i),
      _ => throw const Refused('That does not look like an invitation.'),
    };
    return _refusable(
      () => _authed<Object?>(
        'POST',
        '/v1/people/$_person/connections',
        data: {'person': person, 'invite': invite},
      ),
    );
  }

  Future<Response<T>> _authed<T>(
    String method,
    String path, {
    Object? data,
  }) async {
    Future<Response<T>> send() => _dio.request<T>(
      path,
      data: data,
      options: Options(
        method: method,
        headers: {'Authorization': 'Bearer $_token'},
      ),
    );
    try {
      return await send();
    } on DioException catch (e) {
      // A session running out is not an event (ADR-0015): sign in again,
      // quietly, once.
      if (e.response?.statusCode != 401) rethrow;
      await _signIn();
      return send();
    }
  }

  /// Turns a refusal the server explained into a [Refused] with its sentence.
  Future<void> _refusable(Future<Object?> Function() call) async {
    try {
      await call();
    } on DioException catch (e) {
      final body = e.response?.data;
      if (body is Map && body['error'] is String) {
        throw Refused(body['error'] as String);
      }
      rethrow;
    }
  }
}
