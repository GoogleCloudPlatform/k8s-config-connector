### [2026-06-15] Grafeas Note Direct Types Scaffolding
- **Context**: Implementing initial KRM types, IdentityV2, and CRD for `GrafeasNote` resource under `grafeas.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: 
  1. The googleapis package for Grafeas is `grafeas.v1` instead of standard `google.cloud.grafeas.v1`, so using `--service google.cloud.grafeas.v1` caused protobuf package lookup failure in `generate-types`.
  2. The Google APIs .pb file compiled during `generate-proto.sh` did not originally include `${THIRD_PARTY}/googleapis/grafeas/v1/*.proto`, leading to compilation errors.
  3. The initial `generate-types` scaffolding left the `Spec` struct empty, which caused the generator to prune and comment out all of Grafeas `Note`'s sub-structures and fields under `types.generated.go` as "unreachable".
- **Solution**:
  1. Added `${THIRD_PARTY}/googleapis/grafeas/v1/*.proto` to the `protoc` command in `dev/tools/controllerbuilder/generate-proto.sh` and cleared the cached pb files before regenerating.
  2. Changed `--service` in `generate.sh` to `grafeas.v1`.
  3. Manually filled the `GrafeasNoteSpec` and `GrafeasNoteObservedState` fields inside `note_types.go` and mapped them to their respective proto fields via `+kcc:proto:field` annotations, making all referenced sub-types (e.g., `VulnerabilityNote`, `BuildNote`, etc.) fully reachable and active in `types.generated.go`.
  4. Mapped the `GrafeasNote` Identity V2 template to `projects/{project}/notes/{note}` using the Container Analysis API host `containeranalysis.googleapis.com` as Grafeas is served through Container Analysis.
- **Impact**: Ensures that any subsequent steps (like controller logic or mappers) will have access to fully active, unpruned Go types for `grafeas.v1.Note` and its nested structures.
