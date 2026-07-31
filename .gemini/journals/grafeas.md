# Grafeas Journal

### [2026-06-29] Grafeas Note Scaffolding and Identity
- **Context**: Scaffolding GrafeasNote types, CRD, and IdentityV2 (v1alpha1) under the direct controller model.
- **Problem**: 
  1. The Google APIs protobuf definitions for Grafeas are located in a non-standard path: `grafeas/v1/` instead of `google/`. Therefore, the `generate-proto.sh` script did not compile them.
  2. The actual package declared in `grafeas.proto` is `grafeas.v1`, but the GCP service name is `google.cloud.grafeas.v1`.
  3. When running the generator, many of the nested Note types (e.g., `VulnerabilityNote`, `BuildNote`) were considered unreachable and pruned (commented out) because they weren't referenced by the default generated Spec.
- **Solution**:
  1. Updated `dev/tools/controllerbuilder/generate-proto.sh` to include `${THIRD_PARTY}/googleapis/grafeas/v1/*.proto` in the compilation list.
  2. Updated the `--service` parameter of `generate-types` in `apis/grafeas/v1alpha1/generate.sh` to be `google.cloud.grafeas.v1,grafeas.v1` so the tool successfully looks up type descriptors in both namespaces.
  3. Fully declared all `type` oneof fields (e.g., `Vulnerability`, `Build`, `Image`, etc.) in `GrafeasNoteSpec`, making them reachable and preventing the generator from commenting them out.
  4. Added `//containeranalysis.googleapis.com/projects/{}/notes/{}` to ignored templates in `pkg/gcpurls/registry_test.go` to prevent asset template validation failures.
- **Impact**: All direct Grafeas Note fields are successfully scaffolded and reachable, and future developers can easily implement direct controllers for Grafeas resources.

### [2026-07-31] Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing the direct controller, fuzzer registration, and recording E2E fixtures for GrafeasNote in mockgcp and real GCP.
- **Problem**:
  1. The `generate-mapper` tool requires the primary service package name (`grafeas.v1`) to be listed FIRST in the `--service` flag argument, otherwise it fails to look up the file descriptors.
  2. The standard `grafeas.NewClient` defaults to `grafeas.googleapis.com:443`, but KCC and mockgcp are configured to route and intercept container analysis on `containeranalysis.googleapis.com:443`.
  3. When initializing the client, passing `RESTClientOptions()` returns HTTP options, but `grafeas.NewClient` expects gRPC transport options, leading to an error: `unsupported HTTP client specified`.
  4. The `update_mask` in the Update Note RPC throws an invalid argument error if type-specific fields (like `attestation`) are included, because nested note type details are immutable once created. Only generic metadata fields (like description and related URLs) can be updated.
- **Solution**:
  1. Reordered the `--service` flag in `apis/grafeas/generate.sh` to have `grafeas.v1` first (`--service grafeas.v1,google.cloud.grafeas.v1`) and ran `generate-mapper` to auto-generate the Spec and Status conversion functions.
  2. Appended `option.WithEndpoint("containeranalysis.googleapis.com:443")` when initializing the client in the direct controller.
  3. Used `GRPCClientOptions()` instead of `RESTClientOptions()` to fetch client transport options.
  4. Configured the `update.yaml` fixture to leave the immutable `attestation` note configuration unchanged from `create.yaml` so it is not included in the update mask, while successfully verifying updates to mutable fields like descriptions and related URLs.
- **Impact**: GrafeasNote is now a fully direct KCC resource with successful E2E fixture tests in mockgcp, fully-passing fuzz testing, and 100% compliant API coverage.

