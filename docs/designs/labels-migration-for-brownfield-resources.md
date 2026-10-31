# Technical Design: Labels Migration for Brownfield Resources

## 1. Overview

This document addresses the handling of GCP resource `labels` when migrating existing "brownfield" (Terraform/DCL-based) resources to the modern "direct" controller architecture in Config Connector (KCC).

### 1.1. Greenfield vs. Brownfield Label Handling

It is critical to distinguish between Greenfield and Brownfield resources regarding how GCP `labels` are handled in KCC:

1. **Greenfield (New) Resources:**
   In brand new direct controllers, GCP `labels` are represented directly within the resource's `spec.labels` field. They are **not** mapped to the Kubernetes `metadata.labels` field. They are treated like any other standard Proto/KRM map field, requiring no custom backward compatibility mapping.

2. **Brownfield (Migrated) Resources:**
   To maintain strict backward compatibility for resources that migrated from legacy controllers (where GCP `labels` were mapped to `metadata.labels`), direct controllers must follow the same pattern:
   - The Kubernetes `metadata.labels` is the single source of truth for GCP resource `labels`.
   - `spec.labels` is **not** present in the KRM schema.
   - The controller must manually map `metadata.labels` to/from the GCP resource's `labels` field.

To ensure seamless, non-breaking migrations of brownfield resources to direct controllers, we must guarantee that:
1. Direct controllers correctly read and propagate `metadata.labels` to GCP resource `labels`.
2. Any existing resource's labels are not cleared, corrupted, or altered unexpectedly during or after the migration.

We propose:
1. **The cheapest, most robust way** to confirm and verify that the labels field is properly handled during migration.
2. **A Standard Implementation Template** that can be integrated into the GEMINI skills to guide developers and agents on implementing backward-compatible brownfield labels consistently.

---

## 2. Part 1: The Cheapest Way to Confirm Proper Label Handling During Migration

To guarantee that label migration works without introducing manual testing overhead, we identify and propose two complementary "cheap" validation strategies:

### 2.1. Hermetic E2E Fixture Verification (Cheapest Integration Level)
KCC already possesses a comprehensive fixture testing framework (`pkg/test/resourcefixture/testdata/basic`).
The **cheapest way** to confirm proper label propagation is to enforce that every migrated brownfield resource that already supported `labels` has a specific labels test case under its test data:

1. **Test Fixture Structure:**
   For a migrated resource (e.g., `BackupDRBackupPlan`), a dedicated test directory named `backupdrbackupplan-labels` is created containing:
   - `create.yaml`: Configures a resource with a set of valid custom labels in `metadata.labels`.
   - `update.yaml`: Modifies, adds, and removes some labels in `metadata.labels`.

2. **E2E Test Runner Traffic Validation (`_http.log`):**
   The standard E2E fixture test runner (`TestAllInSeries` in `tests/e2e`) executes these YAML files against `mockgcp` (or real GCP when recording). It captures all API request payloads.
   - **Why this is the cheapest:** No new test runner, library, or complex framework is required. It relies completely on the existing hermetic, deterministic, fast-running test suite (`E2E_GCP_TARGET=mock`).
   - If the direct controller forgets to map labels, or handles them incorrectly, the generated HTTP requests captured in `_http.log` will either omit the labels or contain incorrect values, causing the golden file comparison check in the test runner to fail.

3. **Mock Alignment Test (`golden_alignment_test.go`):**
   To ensure that `mockgcp` handles and maps labels identically to real GCP, KCC's `golden_alignment_test.go` acts as a static validation check comparing the real recorded traffic (`_http.log`) against mock recorded traffic (`_http_mock.log` or similar). Any structural discrepancy in label representation between mock and real GCP is immediately flagged during the alignment test.

### 2.2. Automated Fuzzer Schema Validation
To catch issues even before writing fixture tests, we propose adding a **static schema sanity check** in the KRM fuzzer framework (`pkg/fuzztesting`):
- When a fuzzer is registered for a Proto type, the fuzzer framework can inspect the Proto message descriptor via reflection.
- **Trigger Condition:** The validation check only triggers if the brownfield resource supported `labels` in its legacy representation (i.e. they are mapped in the Terraform service mapping, or in the DCL schema), but the fuzzer configuration has **NOT** registered `Unimplemented_LabelsAnnotations(".labels")` or another annotation-skipping helper.
- **Assertion:** If the condition is met, the fuzzer should fail if the direct mapper does not roundtrip the labels properly.
- This forces the developer to explicitly acknowledge how labels are handled (either they are implemented, or they are ignored because they are mapped to metadata).

---

## 3. Part 2: Standard Implementation Template for Migrating Brownfield Labels

Because greenfield and brownfield resources handle labels differently, automated global base-controller-level injection is not recommended. Instead, we propose a **Standard Implementation Template** for brownfield resource migration. This template ensures consistency, and is added to the GEMINI skills to guide developers and agents during migrations.

### Step 1: Remove Labels from the API Specification (KRM Schema)
The `labels` field must not be part of the resource's `Spec` (as `metadata.labels` is the single source of truth). If it exists, comment it out or remove it.

**File to edit:** `apis/<service>/<version>/<resource>_types.go`
```go
type MyResourceSpec struct {
    // ... other fields ...

    // Commented out to ensure metadata.labels is the single source of truth.
    // Labels map[string]string `json:"labels,omitempty"`
}
```
After editing, regenerate the CRDs by running:
```bash
dev/tasks/generate-crds
```

### Step 2: Update the Fuzzer Test
Since labels are not mapped to any field in the KRM spec, configure the fuzzer to ignore the Proto's `.labels` field during testing.

**File to edit:** `pkg/controller/direct/<service>/<resource>_fuzzer.go`
```go
func MyResourceFuzzer() {
    // ...
    f.Unimplemented_LabelsAnnotations(".labels")
}
```

### Step 3: Implement Label Mapping in the Controller Adapter
In the direct controller, map Kubernetes metadata labels to/from the GCP resource labels within the `AdapterForObject` and `Find` methods.

**File to edit:** `pkg/controller/direct/<service>/<resource>_controller.go`
```go
import (
    "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
    u := op.GetUnstructured()
    // ... Convert unstructured to KRM struct and map to desiredProto ...

    // Map Kubernetes metadata.labels to GCP desiredProto labels
    desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

    // ...
}
```
If the GCP resource supports partial patch updates, ensure that label updates are correctly included in the patch request field masks.

### Step 4: Add E2E Test Fixtures for Labels
Create a dedicated test case under `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/` to verify that labels are propagated correctly.

**Directory name:** `<resource_lowercase>-labels`
- `create.yaml`: Contains custom labels under `metadata.labels` (such as `custom-label: "foo"`).
- `update.yaml`: Modifies or removes some labels under `metadata.labels` (such as changing `custom-label: "foo"` to `"bar"`).

This standard template is registered in `.gemini/skills/kcc-direct-controller-logic-brownfield/SKILL.md` to guarantee a consistent developer and agent workflow.
