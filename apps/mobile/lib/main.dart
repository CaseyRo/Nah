// Nah? — a private home for your closest people.
//
// M1, the walking skeleton (CDI-1833, CDI-1834): join by invitation, sign in,
// give your name (CDI-1896), read the feed, post a text moment, and connect by
// pasting an invitation. The
// design system, the radial menu, the timeline clock, the spring physics and the
// onboarding ritual are all decided and none of them are on the path to two
// phones talking.
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'feed_cubit.dart';
import 'moment.dart';
import 'server.dart';

void main() => runApp(const NahApp());

class NahApp extends StatelessWidget {
  const NahApp({super.key, this.server});

  /// Injected by tests so a widget test never opens a socket.
  final Server? server;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => FeedCubit(server ?? Server())..start(),
      child: MaterialApp(
        title: 'Nah?',
        debugShowCheckedModeBanner: false,
        // No design tokens yet, on purpose. System defaults until M2.
        theme: ThemeData(useMaterial3: true),
        home: const FeedScreen(),
      ),
    );
  }
}

class FeedScreen extends StatelessWidget {
  const FeedScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final ready = context.select((FeedCubit feed) => feed.state is Ready);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Nah?'),
        actions: [
          if (ready)
            IconButton(
              tooltip: 'Connect',
              icon: const Icon(Icons.person_add_alt),
              onPressed: () => showModalBottomSheet<void>(
                context: context,
                isScrollControlled: true,
                builder: (_) => BlocProvider.value(
                  value: context.read<FeedCubit>(),
                  child: const ConnectSheet(),
                ),
              ),
            ),
        ],
      ),
      body: BlocBuilder<FeedCubit, FeedState>(
        builder: (context, state) => switch (state) {
          Starting() => const Center(child: Text('Looking for your circle…')),
          Unreachable() => Center(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text('Cannot reach your circle right now.'),
                const SizedBox(height: 16),
                TextButton(
                  onPressed: () => context.read<FeedCubit>().start(),
                  child: const Text('Try again'),
                ),
              ],
            ),
          ),
          Uninvited() => const JoinScreen(),
          Unnamed() => const NameScreen(),
          Ready(:final moments) => Column(
            children: [
              const Composer(),
              const Divider(height: 1),
              Expanded(
                child: RefreshIndicator(
                  onRefresh: context.read<FeedCubit>().refresh,
                  child: ListView.separated(
                    // Always scrollable, so pulling down works on an empty feed.
                    physics: const AlwaysScrollableScrollPhysics(),
                    itemCount: moments.length,
                    separatorBuilder: (_, _) => const Divider(height: 1),
                    itemBuilder: (_, i) => MomentTile(moments[i]),
                  ),
                ),
              ),
            ],
          ),
        },
      ),
    );
  }
}

/// Nah? is by invitation, so a phone that has not joined asks for one. The first
/// person on a server pastes the one `server invite` printed; everyone after
/// that pastes one from someone already here.
class JoinScreen extends StatefulWidget {
  const JoinScreen({super.key});

  @override
  State<JoinScreen> createState() => _JoinScreenState();
}

class _JoinScreenState extends State<JoinScreen> {
  final _invitation = TextEditingController();
  String? _problem;
  bool _busy = false;

  @override
  void dispose() {
    _invitation.dispose();
    super.dispose();
  }

  Future<void> _join() async {
    final feed = context.read<FeedCubit>();
    setState(() {
      _busy = true;
      _problem = null;
    });
    final problem = await feed.join(_invitation.text);
    if (!mounted) return;
    setState(() {
      _busy = false;
      _problem = problem;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(24),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          const Text('Nah? is by invitation. Paste the one someone sent you.'),
          const SizedBox(height: 12),
          TextField(
            controller: _invitation,
            decoration: const InputDecoration(hintText: 'Your invitation'),
          ),
          const SizedBox(height: 12),
          FilledButton(
            onPressed: _busy ? null : _join,
            child: const Text('Join'),
          ),
          if (_problem != null)
            Padding(
              padding: const EdgeInsets.only(top: 12),
              child: Text(_problem!),
            ),
        ],
      ),
    );
  }
}

/// The first thing asked after joining: the name the people closest to you
/// already call you by (CDI-1896). It is required, because a circle has to know
/// who posted, and it has no counter (ADR-0004).
class NameScreen extends StatefulWidget {
  const NameScreen({super.key});

  @override
  State<NameScreen> createState() => _NameScreenState();
}

class _NameScreenState extends State<NameScreen> {
  final _name = TextEditingController();
  String? _problem;
  bool _busy = false;

  @override
  void dispose() {
    _name.dispose();
    super.dispose();
  }

  Future<void> _continue() async {
    final name = _name.text.trim();
    if (name.isEmpty) return;
    final feed = context.read<FeedCubit>();
    setState(() {
      _busy = true;
      _problem = null;
    });
    final problem = await feed.name(name);
    if (!mounted) return;
    setState(() {
      _busy = false;
      _problem = problem;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(24),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          Text(
            'What do the people closest to you call you?',
            style: Theme.of(context).textTheme.headlineSmall,
          ),
          const SizedBox(height: 12),
          TextField(
            controller: _name,
            autofocus: true,
            textCapitalization: TextCapitalization.words,
            inputFormatters: [LengthLimitingTextInputFormatter(maxNameLength)],
            onChanged: (_) => setState(() {}),
            onSubmitted: (_) => _continue(),
            decoration: const InputDecoration(
              hintText: 'Your name',
              helperText: 'Only your circle sees it. You can change it later.',
            ),
          ),
          const SizedBox(height: 12),
          FilledButton(
            onPressed: _busy || _name.text.trim().isEmpty ? null : _continue,
            child: const Text('Continue'),
          ),
          if (_problem != null)
            Padding(
              padding: const EdgeInsets.only(top: 12),
              child: Text(_problem!),
            ),
        ],
      ),
    );
  }
}

class MomentTile extends StatelessWidget {
  const MomentTile(this.moment, {super.key});

  final Moment moment;

  @override
  Widget build(BuildContext context) {
    final l10n = MaterialLocalizations.of(context);
    final at = moment.at.toLocal();
    final when = DateUtils.isSameDay(at, DateTime.now())
        ? l10n.formatTimeOfDay(TimeOfDay.fromDateTime(at))
        : l10n.formatShortMonthDay(at);
    return ListTile(
      title: Text(moment.text),
      subtitle: Text('${moment.authorName ?? 'Someone'} · $when'),
    );
  }
}

/// Compose or post, nothing in between: there is no draft (CDI-1834).
class Composer extends StatefulWidget {
  const Composer({super.key});

  @override
  State<Composer> createState() => _ComposerState();
}

class _ComposerState extends State<Composer> {
  final _text = TextEditingController();
  bool _sending = false;

  @override
  void dispose() {
    _text.dispose();
    super.dispose();
  }

  Future<void> _send() async {
    final text = _text.text.trim();
    if (text.isEmpty || _sending) return;
    final feed = context.read<FeedCubit>();
    setState(() => _sending = true);
    final problem = await feed.post(text);
    if (!mounted) return;
    setState(() => _sending = false);
    if (problem == null) {
      _text.clear();
    } else {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(problem)));
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.fromLTRB(16, 8, 8, 8),
      child: Row(
        children: [
          Expanded(
            child: TextField(
              controller: _text,
              minLines: 1,
              maxLines: 5,
              // A limit with no counter: a moment has an edge, not a number
              // (ADR-0004).
              inputFormatters: [
                LengthLimitingTextInputFormatter(maxTextLength),
              ],
              decoration: const InputDecoration(
                hintText: 'Say something to your circle',
              ),
            ),
          ),
          IconButton(
            tooltip: 'Send',
            onPressed: _sending ? null : _send,
            icon: const Icon(Icons.send),
          ),
        ],
      ),
    );
  }
}

/// Standing in for the touch (CDI-1840) and the link (CDI-1839), both M2: copy
/// your invitation to someone, or paste theirs.
class ConnectSheet extends StatefulWidget {
  const ConnectSheet({super.key});

  @override
  State<ConnectSheet> createState() => _ConnectSheetState();
}

class _ConnectSheetState extends State<ConnectSheet> {
  final _theirs = TextEditingController();
  String? _mine;
  String? _message;
  bool _busy = false;

  @override
  void initState() {
    super.initState();
    context.read<FeedCubit>().invitation().then((code) {
      if (!mounted) return;
      setState(() {
        if (code == null) {
          _message = 'Could not make an invitation right now.';
        } else {
          _mine = code;
        }
      });
    });
  }

  @override
  void dispose() {
    _theirs.dispose();
    super.dispose();
  }

  Future<void> _connect() async {
    final feed = context.read<FeedCubit>();
    setState(() {
      _busy = true;
      _message = null;
    });
    final problem = await feed.connect(_theirs.text);
    if (!mounted) return;
    setState(() {
      _busy = false;
      _message = problem ?? 'Connected.';
    });
    if (problem == null) _theirs.clear();
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(
        24,
        24,
        24,
        24 + MediaQuery.viewInsetsOf(context).bottom,
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          const Text('Send this to someone to connect with them.'),
          const SizedBox(height: 8),
          SelectableText(_mine ?? '…'),
          TextButton(
            onPressed: _mine == null
                ? null
                : () async {
                    await Clipboard.setData(ClipboardData(text: _mine!));
                    if (mounted) setState(() => _message = 'Copied.');
                  },
            child: const Text('Copy'),
          ),
          const Divider(height: 32),
          const Text('Or paste theirs.'),
          TextField(
            controller: _theirs,
            decoration: const InputDecoration(hintText: 'Their invitation'),
          ),
          const SizedBox(height: 12),
          FilledButton(
            onPressed: _busy ? null : _connect,
            child: const Text('Connect'),
          ),
          if (_message != null)
            Padding(
              padding: const EdgeInsets.only(top: 12),
              child: Text(_message!),
            ),
        ],
      ),
    );
  }
}
