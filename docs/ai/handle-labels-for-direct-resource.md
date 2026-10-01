# How to Handle Labels for Direct Resources

This document describes how to handle GCP labels for direct resources in KCC.

Which field holds the labels depends on whether the resource is **greenfield**(implemented via direct controller) or **brownfield**(migrated from TF/DCL controller):

| Resource type | Source of truth for GCP labels |
|---|---|---|
| **greenfield** | `spec.labels` |
| **brownfield** | `metadata.labels` for backward compatibility and gradually migrate to `spec.labels` |

## Greenfield resources: use `spec.labels`

For greenfield direct resources, user-specified labels are modeled like any other GCP field.

> [!WARNING]
> Do **not** comment out the `Labels` field in the `Spec`, and do **not** mark `.labels` as unimplemented in the fuzzer (`Unimplemented_LabelsAnnotations`, `Unimplemented_NotYetTriaged`, or `UnimplementedFields.Insert`). If the GCP API supports labels, they must be exposed in `spec.labels`.
> If the GCP API supports labels, the direct controller **must** pass them through to GCP. Silently dropping labels is a bug.
> 
### 1. Keep `Labels` in the API Spec

**File to edit:** `apis/<service>/<version>/<kind>_types.go`

```go
// +kcc:spec:proto=google.cloud.<service>.v1.<Kind>
type <Kind>Spec struct {
    // ... other fields ...

    // Optional. Labels as key value pairs.
    // +kcc:proto:field=google.cloud.<service>.v1.<Kind>.labels
    // +optional
    Labels map[string]string `json:"labels,omitempty"`

    // ... other fields ...
}
```

If the field was previously commented out, uncomment it and regenerate the CRDs and mappers:

```bash
dev/tasks/generate-types-and-mappers
```

### 2. Map labels in the mapper

The generated `<Kind>Spec_ToProto` / `<Kind>Spec_FromProto` should now contain:

```go
out.Labels = in.Labels
```

If the mapper is hand-written, add this line to both directions. Make sure there is no leftover `// MISSING: Labels` comment.

### 3. Update the Fuzzer

**File to edit:** `pkg/controller/direct/<service>/<kind>_fuzzer.go`

Treat labels as a regular spec field:

```go
f.SpecField(".labels")
```

Remove any line that marks labels as unimplemented. Labels can be marked unimplemented in any of these forms, and all of them must be removed:

```go
f.Unimplemented_LabelsAnnotations(".labels")
f.Unimplemented_NotYetTriaged(".labels")
f.UnimplementedFields.Insert(".labels")
```

### 4. Update the Controller

The desired proto built from `<Kind>Spec_ToProto` already contains the labels. Do **not** overwrite `desired.Labels` with `metadata.labels`.

In `Update()`, make sure label changes are detected and that `labels` is included in the update mask. For example, if you compute the diff with `common.CompareProtoMessage`, `labels` is included automatically because it is a spec field.

### 5. Update Fixture Tests

**Path to edit:** `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/`

Set `spec.labels` in `create.yaml` and change it in `update.yaml` (add, modify, and/or remove a key) so that label updates are covered by the golden HTTP logs.

```yaml
# create.yaml
spec:
  labels:
    env: "test"
    team: "foo"
```

```yaml
# update.yaml
spec:
  labels:
    env: "prod"   # modified
    owner: "bar"  # added; "team" removed
```

---

## Brownfield resources: use `metadata.labels` for backward compatibility

For brownfield resources, the Terraform/DCL controllers map Kubernetes `metadata.labels` to GCP labels. To avoid breaking existing users, the direct controller **must** keep the same behavior.

See the [kcc-direct-brownfield-labels](../../.gemini/skills/kcc-direct-brownfield-labels/SKILL.md) skill has more detail.
