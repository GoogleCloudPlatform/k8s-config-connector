# Journal: LoggingLogBucket Transition to Direct Types

## Learnings & Patterns

### 1. Spec-level oneOf validation for parent references
`LoggingLogBucket` can belong to any one of four parents: `projectRef`, `folderRef`, `organizationRef`, or `billingAccountRef`.
The Baseline CRD enforced this using a `oneOf` constraint at `spec`. When we generated the types, the `oneOf` was lost because `controller-gen` does not generate spec-level `oneOf` constraints out of the box for independent fields.
We learned that the `scripts/add-validation-to-crds/parse-crds.go` script hard-codes these validations based on kind (previously done for `LoggingLogView`). We added `LoggingLogBucket` there:
```go
		} else if (kind == "LoggingLogView" || kind == "LoggingLogBucket") && fieldPath == ".spec" {
			ruleYAML = `
oneOf:
- required: [billingAccountRef]
- required: [folderRef]
- required: [organizationRef]
- required: [projectRef]
`
```
This correctly restored the `oneOf` constraint to match the baseline schema perfectly.

### 2. Handling integer type mismatches between proto and KRM
In the baseline CRD, `retentionDays` is mapped as `int64` (integer). However, in GCP's `LogBucket` proto, `retention_days` is `int32`.
The `generate-mapper` tool generates direct assignment mappings which fail to compile due to the `int32` vs `int64` mismatch.
We resolved this by writing custom/hand-coded spec mapper functions in `pkg/controller/direct/logging/mapper.go`:
- `LoggingLogBucketSpec_FromProto`
- `LoggingLogBucketSpec_ToProto`
- `LoggingLogBucketStatus_FromProto`
Writing these hand-coded mapping functions causes the mapper generator to automatically discover them and comment out the conflicting generated versions in `mapper.generated.go`.

### 3. Clean References (Kindless)
The parent references `projectRef`, `folderRef`, `organizationRef`, and `billingAccountRef` in the baseline CRD did not contain a `kind` field.
Importing standard references from `apis/refs/v1beta1/` would introduce `kind` to the OpenAPI schema. We resolved this by importing and using clean kindless references from `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"`:
- `refs.ProjectRef`
- `refs.FolderRef`
- `refs.OrganizationRef`
- `refs.BillingAccountRef`
This achieved 100% schema compatibility without manual replication of standard types.
