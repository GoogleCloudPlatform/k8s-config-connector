# PubSubSnapshot Direct Controller Implementation Journal

## Observations
1. **Standardized IdentityV2 Pattern**: Refactored the legacy `SnapshotIdentity` struct to `PubSubSnapshotIdentity` using the modern `gcpurls.Template` pattern, matching other Pub/Sub direct controllers.
2. **Standardized Reference V2 Pattern**: Refactored `SnapshotRef` to implement the `refs.Ref` interface, allowing unified normalization and validation across namespaces.
3. **KRM-only Spec Fields**: Successfully isolated the KRM-only `pubSubSubscriptionRef` field in `AdapterForObject` and handled it inside the controller adapter rather than blending it into the `pb.Snapshot` protobuf, keeping the protobuf representations clean and strictly identical to GCP state.
4. **Structured Diff & Immutability**: Utilized `common.CompareProtoMessageStructuredDiff` for calculating updates deterministically. Added checks on immutable fields like `topic` to prevent user updates and surface precise immutability errors.
