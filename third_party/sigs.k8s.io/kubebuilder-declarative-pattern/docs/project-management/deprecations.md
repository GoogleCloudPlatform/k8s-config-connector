# Deprecations and new functionality

As kubernetes and the controller ecosystem continues to evolve,
we will sometimes want to add new functionality or behaviour to kdp.

However, as a library, we do not want to break the functionality of controllers that upgrade their kdp version,
particularly as the coupling to kubernetes versions means that users must upgrade periodically.

We will aim to follow these rules therefore:

* Do not remove functionality or fundamentally change the behaviour in a way that will break users.

* Bug fixes are good, but we should think about existing users even for bug fixes.

* Consider introducing a new method or type instead of breaking existing functionality.

* Use the go `// Deprecated:` [convention](https://go.dev/wiki/Deprecated) to discourage usage of "old" methods, fields or types.

* Prefer errors at compilation time to errors at run time.  Reasonable code changes are acceptable (e.g. adding a context method).
  Crashing at run time because a field is not populated is not acceptable.
