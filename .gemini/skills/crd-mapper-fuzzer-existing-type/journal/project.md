# Project CRD Mapper Journal

## Context
Transitioned the `Project` resource under `apis/resourcemanager` to support direct KRM types and mappers.

## Learnings & Observations
- **Custom Reference / oneOf Mapping**:
  - The Project resource spec contains fields `folderRef` and `organizationRef`, which both map to the single protobuf field `parent` (representing a GCP parent URI such as `folders/xxxx` or `organizations/yyyy`).
  - Standard automatic mapper generation fails or creates incomplete mappings for fields mapping to the same protobuf field.
  - To solve this, we handcoded `ProjectSpec_FromProto`, `ProjectSpec_ToProto`, `ProjectStatus_FromProto`, and `ProjectStatus_ToProto` in `pkg/controller/direct/resourcemanager/mapper.go`.
  - The controller builder tool `generate-mapper` automatically detected our handcoded mapping functions and skipped generating them in `mapper.generated.go`, commenting them out.
- **Status Project Number Extraction**:
  - The project status `number` field corresponds to the numeric ID of the project.
  - In protobuf, RM v3 returns the resource name as `projects/123456789`. We extracted the project number by parsing/splitting the name prefix (`projects/`).
- **Protobuf Generation Dependency**:
  - Local `generate.sh` runs require `.build/googleapis.pb` to exist. If it's missing, executing `dev/tools/controllerbuilder/generate-proto.sh` resolves the error by compiling the protos into the local build folder.
- **Canonical Identity and Reference Types**:
  - Implemented the canonical `_identity.go` (`project_identity.go`) and `_reference.go` (`project_reference.go`) under `apis/resourcemanager/v1beta1/` following the `identity.IdentityV2` and `refs.Ref` interface models.
  - Leveraged `common.ToStructuredType[*Project](u)` in the `Normalize` fallback callback function of the reference implementation to safely convert unstructured objects into fully typed ones.
  - Validated template registration with `go test ./pkg/gcpurls/...`, syntax correctness via `go vet ./apis/resourcemanager/v1beta1/...`, and project-wide formatting via `make fmt`.

## KRM Round-trip Fuzzer
- Implemented the `Project` round-trip KRM fuzzer under `pkg/controller/direct/resourcemanager/project_fuzzer.go` and registered it via `fuzztesting.RegisterKRMFuzzer`.
- Because the `parent` and `name` fields are expected to follow strict naming prefixes (e.g. `folders/`, `organizations/`, and `projects/`), random protobuf input generation could cause mapper validation errors. We addressed this by implementing `FilterSpec` and `FilterStatus` functions within the fuzzer:
  - `FilterSpec` normalizes `.parent` to start with `folders/` or `organizations/` prefix using sanitized alphanumeric IDs.
  - `FilterStatus` normalizes `.name` to start with `projects/` prefix using sanitized alphanumeric IDs.
- The fuzzer covers 3 spec fields (`.display_name`, `.parent`, `.project_id`), 1 status field (`.name`), and sets 6 unimplemented fields (`.state`, `.create_time`, `.update_time`, `.delete_time`, `.etag`, `.labels`).
