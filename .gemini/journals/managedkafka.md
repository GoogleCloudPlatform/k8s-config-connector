# ManagedKafka Journal

### [2026-06-15] Underscores in URL Template Variable Matching
- **Context**: Implementing IdentityV2 for `ManagedKafkaConnectCluster`.
- **Problem**: The GCP URL path contains the placeholder `{connect_cluster}`, which `gcpurls.Template` parses. It then attempts to find a matching field in the identity struct by converting both the template placeholder and Go struct fields to lower case and matching them. A Go field named `ConnectCluster` is lower-cased to `"connectcluster"`, which does not match the lower-cased template placeholder `"connect_cluster"`, leading to a runtime panic during initialization.
- **Solution**: Named the struct field exactly `Connect_cluster string`. The lower-cased field name is `"connect_cluster"`, which matches the placeholder `"connect_cluster"` perfectly, resolving the initialization panic.
- **Impact**: When implementing IdentityV2 for other resources with underscores in their URL template path parameters, ensure that the Go struct field retains the underscore (e.g. `Field_name`) to match the placeholder.

### [2026-10-29] ManagedKafkaACL ParentRef and Etag
- **Context**: Triaging judgement queue and finalizing types for `ManagedKafkaACL`.
- **Problem**: `ManagedKafkaACL` is a child resource of `ManagedKafkaCluster` (`.../clusters/{cluster}/acls/{acl}`). The generator guessed `.spec.clusterRef`, and placed `etag` in Spec because the proto lacked `OUTPUT_ONLY`.
- **Solution**: Replaced `projectRef`, `location`, and guessed `clusterRef` with `ParentRef *ClusterRef` per the single-parent rule (using the existing `ClusterRef` type from `v1alpha1`). Moved `etag` from `Spec` to `ObservedState`, regenerated types, identity, mapper, controller, and fuzzer.
- **Impact**: Child resources under `ManagedKafkaCluster` must use `spec.parentRef *ClusterRef`. All ancestor segments are resolved from `parentRef`.

### [2026-11-02] ManagedKafkaACL Direct Controller, Concurrency Control, and E2E Testing
- **Context**: Implementing Greenfield direct controller, fuzzer, and recording E2E fixtures for `ManagedKafkaACL`.
- **Problem**:
  1. `UpdateAcl` requires `etag` for concurrency control. Because `etag` is modelled in `ObservedState`, `CompareGreenfieldSpec` constructs `resource` from desired spec without `etag`.
  2. GCP Managed Kafka enforces specific ACL resource pattern IDs: `cluster`, `allTopics`, `allConsumerGroups`, `allTransactionalIds`, or path patterns (`topic/{name}`). Arbitrary resource IDs are rejected by GCP.
  3. E2E tests for `ManagedKafkaACL` require a `ManagedKafkaCluster` dependency. Cluster provisioning and deletion take ~5-10 minutes each.
- **Solution**:
  1. In `Update`, explicitly set `resource.Etag = a.actual.Etag` before calling `UpdateAcl`.
  2. Configured test fixtures using canonical valid ACL IDs: `resourceID: cluster` for minimal fixture and `resourceID: allTopics` for maximal fixture.
  3. Scoped HTTP report proto schema validation to primary entries and ensured token boundary checks in classifier so that parent `/clusters` endpoints and `clusterId` query parameters are not misclassified as primary ACL mutations.
- **Impact**: All direct controller reconciliation passes real GCP tests with clean 7/7 HTTP lifecycle reports and re-reconciliation verification.
