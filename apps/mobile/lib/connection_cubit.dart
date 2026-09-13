import 'package:flutter_bloc/flutter_bloc.dart';

import 'server.dart';

enum Connection { unknown, reachable, unreachable }

/// A Cubit rather than a full Bloc: there are no events worth naming yet, and
/// the stack note in the README says Bloc/Cubit, not Bloc.
class ConnectionCubit extends Cubit<Connection> {
  ConnectionCubit(this._server) : super(Connection.unknown);

  final Server _server;

  Future<void> check() async {
    emit(Connection.unknown);
    emit(
      await _server.healthy() ? Connection.reachable : Connection.unreachable,
    );
  }
}
