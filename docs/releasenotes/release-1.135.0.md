# 1.135.0 - 2025-09-18

*   Special shout-outs to acpana, anhdle-sso, cheftako, fqbright, gemmahou, GoogleCloudPlatform, himanikh, jingyih, justinsb, kjasnoor0305, xiaoweim, yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   `APIQuotaPreference`
*   `PubSubSnapshot`
*   `VMwareEngineExternalAddress`

## New Alpha Resources (Direct Reconciler):

*   `BillingAccount`
*   `FirestoreIndex`
*   `IAMDenyPolicy`
*   `ServiceNetworkingPeeredDNSDomain`

## New Fields:

*   `AlloyDBCluster`
    *   Added `spec.databaseVersion` field to support specifying the database version for the cluster.

## Reconciliation Improvements

*   A new controller routing logic has been introduced to provide more control over which reconciler (direct, terraform, or dcl) is used for a given resource. This can be configured via the `ConfigConnectorContext` resource or a static map.

## New features:

*   The `AssetSavedQuery` resource now supports `{futureTimestamp}` in the query, which will be replaced with a timestamp 90 days in the future.
*   Added mockgcp support for label validation, DNS managed zones, and DNS resource record sets.
*   Introduced a client-side custom resourceLock for multi-cluster leader election.

## Bug Fixes:

*   Optimized `OrgPolicyPolicy` updates by comparing policies before applying changes.
*   Fixed a nil pointer dereference in the AlloyDB direct controller.
*   Fixed several fuzzers by adding missing fields or types for `NetworkConnectivity`, `ComputeForwardingRule`, `CertificateManager`, `Dataflow`, and `Monitoring`.
*   Fixed nil pointer panics in `ComputeFirewallPolicyRule` and `ComputeTargetTCPProxy`.
*   Moved `expireTime` from spec to the observed state for `PubSubSnapshot`.
*   Fixed an issue with unused imports in the mapper-generator.
*   Fixed a proto mismatch for PubSub in the API and mapper+fuzzer.
