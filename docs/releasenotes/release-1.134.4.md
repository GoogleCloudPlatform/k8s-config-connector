# Release 1.134.4

*   Special shout-outs to @cheftako, @maqiuyujoyce, and @xiaoweim for their contributions to this release.

## New features:

*   [#6065](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6065): Enabled Vertical Pod Autoscaler (VPA) support. Users can now enable VPA for Config Connector components via `ControllerResource` and `NamespacedControllerResource` to automatically adjust resource requests.

## Bug Fixes:

*   [#6035](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/6035): Fixed an issue where `managedFields` metadata could be incorrectly attributed to the `status` subresource during spec updates, causing "Location must be set" errors.