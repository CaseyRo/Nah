import 'package:dio/dio.dart';

/// The circle server this build talks to.
///
/// A person never types this and never sees it. It is a build-time setting for
/// now; once the directory exists (CDI-1838) a circle's address is looked up
/// from its identifier, because joining must never involve a hostname.
const serverBaseUrl = String.fromEnvironment(
  'NAH_SERVER',
  defaultValue: 'http://127.0.0.1:8080',
);

/// Everything the app knows about talking to a circle server.
///
/// This is deliberately thin. The real client — sessions per server, the
/// per-circle cache, encrypting on the way out — is CDI-1861, and it belongs in
/// a package the app owns rather than in here.
class Server {
  Server({Dio? dio})
    : _dio =
          dio ??
          Dio(
            BaseOptions(
              baseUrl: serverBaseUrl,
              connectTimeout: const Duration(seconds: 10),
              receiveTimeout: const Duration(seconds: 20),
              // A moment is ciphertext, so nothing here negotiates content.
              responseType: ResponseType.json,
            ),
          );

  final Dio _dio;

  Future<bool> healthy() async {
    try {
      final res = await _dio.get<String>(
        '/healthz',
        options: Options(responseType: ResponseType.plain),
      );
      return res.statusCode == 200;
    } on DioException {
      return false;
    }
  }
}
