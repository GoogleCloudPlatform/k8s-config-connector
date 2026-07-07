### [2026-07-06] BigQueryDataTransferTransferConfig Greenfield Implementation
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for BigQueryDataTransferTransferConfig (v1alpha1) under Issue #9021.
- **Problem**: The proto definition of TransferConfig contains fields that can refer to other KCC resources (e.g., KMS key name, PubSub topic, PubSub subscription, and target dataset). Leaving these as plain strings in the Spec would break KCC reference standards.
- **Solution**: Manually configured EncryptionConfiguration (using `refsv1beta1.KMSCryptoKeyRef`), target dataset (using `bigqueryv1beta1.DatasetRef`), notification Pub/Sub topic (using `pubsubv1beta1.PubSubTopicRef`), and event-driven schedule Pub/Sub subscription (using `pubsubv1beta1.PubSubSubscriptionRef`). Because these are referenced in the spec, the generator correctly uncomments and reaches all related nested types in `types.generated.go`.
- **Impact**: All reference-like fields must be implemented as proper KCC reference fields rather than raw strings to ensure consistency and correct reconciliation.

### [2026-05-17] bigquerydatatransfer migration
- **Context**: Moving mockgcp bigquerydatatransfer to httptogrpc (PR 8302).
- **Problem**: Previously used RewriteError for 404s to clear error.Errors.
- **Solution**: Removed as httptogrpc does not support it.
- **Impact**: It is important to remember that RewriteError is not supported in httptogrpc.
