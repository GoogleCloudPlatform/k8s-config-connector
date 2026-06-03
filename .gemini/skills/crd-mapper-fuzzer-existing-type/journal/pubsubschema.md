# PubSubSchema Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Strict Schema Compatibility for References**:
   - The baseline `PubSubSchema` CRD defines a `projectRef` without a `kind` field.
   - Using the standard `refsv1beta1.ProjectRef` automatically generates and includes `kind` under `projectRef` properties.
   - To match the baseline CRD exactly and ensure strict compatibility, we used the clean, kindless `refs.ProjectRef` type from `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"`.

2. **Removal of Unused Status Fields**:
   - The baseline CRD status only includes `conditions` and `observedGeneration`.
   - The scaffolded KRM types include `externalRef` and `observedState`. To achieve 100% schema parity, we removed `externalRef` and `observedState` from `PubSubSchemaStatus` in `schema_types.go`.

3. **Global Scope vs. Location**:
   - `PubSubSchema` is a global resource mapped to GCP projects (`projects/{project}/schemas/{schema}`). It does not have a `location` (region) property in GCP or in the baseline CRD.
   - We removed the scaffolded `Location` property from `PubSubSchemaSpec` to maintain perfect compatibility.

## Validation Results
- Running `dev/tasks/diff-crds` produced absolutely empty/zero output, confirming that the new generated CRD is 100% strictly schema-compatible with the baseline CRD.
- Successfully ran `make generate-go-client`, `go build`, and `go vet` without errors.
