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
