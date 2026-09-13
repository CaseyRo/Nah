// Nah? — a private home for your closest people.
//
// This is the scaffold and nothing more (CDI-1832). The design system, the
// radial menu, the timeline clock, the spring physics and the onboarding ritual
// are all decided and none of them are on the path to two phones talking.
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'connection_cubit.dart';
import 'server.dart';

void main() => runApp(const NahApp());

class NahApp extends StatelessWidget {
  const NahApp({super.key, this.server});

  /// Injected by tests so a widget test never opens a socket.
  final Server? server;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => ConnectionCubit(server ?? Server())..check(),
      child: MaterialApp(
        title: 'Nah?',
        debugShowCheckedModeBanner: false,
        // No design tokens yet, on purpose. System defaults until M2.
        theme: ThemeData(useMaterial3: true),
        home: const ScaffoldCheck(),
      ),
    );
  }
}

/// The only screen: does this build reach a circle server. It exists to prove
/// the wiring and gets replaced by the feed in CDI-1833.
class ScaffoldCheck extends StatelessWidget {
  const ScaffoldCheck({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: BlocBuilder<ConnectionCubit, Connection>(
          builder: (context, state) => Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(switch (state) {
                Connection.unknown => 'Looking for your circle…',
                Connection.reachable => 'Your circle is there.',
                Connection.unreachable => 'Cannot reach your circle right now.',
              }, style: Theme.of(context).textTheme.titleMedium),
              const SizedBox(height: 16),
              TextButton(
                onPressed: () => context.read<ConnectionCubit>().check(),
                child: const Text('Try again'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
