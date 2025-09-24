*   Special shout-outs to @justinsb and @xiaoweim for their contributions to this release.

## New features:

*   [#5231](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5231): Add more verbose logging during certificate validation to assist with debugging.

## Bug Fixes:

*   [#5230](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5230): Fixed an issue that could lead to premature certificate rotation by ensuring errors are not swallowed when reading a Secret.
