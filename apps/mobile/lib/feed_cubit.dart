import 'package:flutter_bloc/flutter_bloc.dart';

import 'moment.dart';
import 'server.dart';

sealed class FeedState {
  const FeedState();
}

class Starting extends FeedState {
  const Starting();
}

class Unreachable extends FeedState {
  const Unreachable();
}

/// This phone has not joined yet, and needs an invitation to.
class Uninvited extends FeedState {
  const Uninvited();
}

/// This person has joined but has not said what they are called (CDI-1896).
class Unnamed extends FeedState {
  const Unnamed();
}

class Ready extends FeedState {
  const Ready(this.moments);
  final List<Moment> moments;
}

/// The one screen: joining, signing in, the feed, and the few things a person
/// can do from it. Anything that goes wrong comes back as a sentence.
class FeedCubit extends Cubit<FeedState> {
  FeedCubit(this._server) : super(const Starting());

  final Server _server;

  Future<void> start() async {
    emit(const Starting());
    try {
      await _server.start();
      await refresh();
    } on NeedsInvitation {
      emit(const Uninvited());
    } on Exception {
      emit(const Unreachable());
    }
  }

  /// Loads the page again. It never loads more (CDI-1833). A person with no
  /// name is asked for one first, because their circle has to know who posted.
  Future<void> refresh() async {
    if (_server.name == null) return emit(const Unnamed());
    try {
      emit(Ready(await _server.feed()));
    } on Exception {
      emit(const Unreachable());
    }
  }

  /// Joins with an invitation, then shows the feed. Returns a sentence to show
  /// if it did not go through.
  Future<String?> join(String invitation) =>
      _attempt(() => _server.join(invitation));

  /// Gives this person their name, then shows the feed. Returns a sentence to
  /// show if it did not go through.
  Future<String?> name(String name) => _attempt(() => _server.setName(name));

  /// Posts, then reloads, so the moment shows where everyone else will see it.
  /// Returns a sentence to show if it did not go through.
  Future<String?> post(String text) => _attempt(() => _server.post(text));

  Future<String?> connect(String invitation) =>
      _attempt(() => _server.connect(invitation));

  Future<String?> invitation() async {
    try {
      return await _server.invitation();
    } on Exception {
      return null;
    }
  }

  Future<String?> _attempt(Future<void> Function() action) async {
    try {
      await action();
    } on Refused catch (e) {
      return e.message;
    } on Exception {
      return 'That did not go through. Try again in a moment.';
    }
    await refresh();
    return null;
  }
}
