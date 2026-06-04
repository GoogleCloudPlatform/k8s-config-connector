# Service Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Strict Schema Compatibility for References**:
   - The baseline `Service` CRD defines a `projectRef` without a `kind` field.
   - To match the baseline CRD exactly and ensure strict compatibility, we used the clean, kindless `refs.ProjectRef` type from `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"`. This replaced the old `parent.ProjectRef` while maintaining identical OpenAPI schema properties (no `kind` field added).

2. **GCP Proto Mapping Annotation**:
   - We annotated the `ServiceSpec` struct in `service_types.go` with `// +kcc:spec:proto=google.api.serviceusage.v1beta1.Service` to explicitly link it to the corresponding GCP Proto Service definition.

3. **Skipping Generation of Duplicate Resource and Mapper Types**:
   - Because the KRM Resource Kind is named `Service` and the Proto Message name is also `Service`, `controllerbuilder` automatically detects that a struct named `Service` already exists and comments out the duplicate `Service` definition inside `types.generated.go`.
   - Since custom/handcoded `ServiceSpec_FromProto`, `ServiceSpec_ToProto`, `ServiceStatus_FromProto`, and `ServiceStatus_ToProto` mapper functions are already present in `pkg/controller/direct/serviceusage/mapper.go`, the mapper generator correctly identifies them and comments out duplicate mapped configurations inside `mapper.generated.go`.

## Validation Results
- Running `dev/tasks/diff-crds` produced absolutely empty/zero output, confirming that our types are 100% strictly schema-compatible with the baseline CRD schema.
- Successfully ran `make lint-custom`, `make manifests`, `make generate-go-client`, and `go vet ./...` with clean passes and no compiler or linter warnings.
