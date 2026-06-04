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
