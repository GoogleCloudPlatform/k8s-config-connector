### [2026-05-24] AIStreams Resource Initialization and Scaffolding
- **Context**: Implementing the initial direct types and IdentityV2 for `AIStreamsCluster` (GCP `google.partner.aistreams.v1alpha1.Cluster`).
- **Problem**: 
  1. The API protos reside under the `google/partner/aistreams/v1alpha1` directory in `googleapis` rather than the standard `google/cloud/aistreams/v1alpha1`. This causes compilation and message lookup failures if not explicitly configured in `generate-proto.sh`.
  2. The `.name` field on the GCP resource represents the resource identity/URL but is not a direct KRM field, which caused fuzz round-trip tests to fail.
- **Solution**:
  1. Updated `dev/tools/controllerbuilder/generate-proto.sh` to compile `${THIRD_PARTY}/googleapis/google/partner/aistreams/*/*.proto` and removed cached `.pb` files to force regeneration.
  2. Configured the fuzzer in `aistreamscluster_fuzzer.go` using `f.Unimplemented_Identity(".name")` to exclude the `.name` field from the spec/status round-trip testing.
- **Impact**: Enables standard KCC tooling (`generate-types`, `generate-mapper`) to discover the AIStreams protobuf package correctly, and ensures successful roundtrip fuzz testing validation.
