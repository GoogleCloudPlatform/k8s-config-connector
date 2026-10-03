---
name: reviewgen-brownfield-controller
description: Review guide and architectural invariants for reviewing PRs that migrate legacy (Terraform/DCL) resources to KCC Direct Controllers.
---

# Review Guide for KCC Brownfield Direct Controller PRs

This skill outlines the mandatory validation invariants, review checks, and architectural patterns required when reviewing Pull Requests migrating existing (brownfield) Config Connector resources from legacy Terraform/DCL controllers to the **Direct Controller** pattern.

---

## 1. Ratcheting Exclusion Check (`tests/e2e/ratcheting.go`) — **MANDATORY**

*   **Strict Invariant**: The target resource's `GroupKind` **must be removed** from the exemption list in `ShouldTestRereconiliation()` in `tests/e2e/ratcheting.go`.
*   **Why it's critical**: Inclusion in `tests/e2e/ratcheting.go` causes `ShouldTestRereconiliation()` to return `false`, completely bypassing Server-Side Apply (SSA) creation validation and 0-write re-reconciliation testing in CI. Without removing this exclusion, steady-state drift bugs and spurious GCP update calls cannot be detected.
*   **Review Action**: Fail the review if the resource's `GroupKind` case statement remains in `ShouldTestRereconiliation()` in `tests/e2e/ratcheting.go`.

---

## 2. Adapter Architecture & Reference Resolution Timing

*   **No Reference Normalization in `AdapterForObject()`**:
    *   Do **NOT** invoke `common.NormalizeReferences(ctx, reader, obj, ...)` or blocking reference lookups inside `AdapterForObject()`.
    *   `AdapterForObject()` is called on every reconciliation pass, including `Find()` and `Delete()`.
    *   If a referenced dependency is still provisioning or has already been deleted, resolving references in `AdapterForObject()` will fail prematurely before the controller can read or delete the target GCP resource.
*   **Resolve References Only in `Create()` and `Update()`**:
    *   Defer `common.NormalizeReferences` to `buildDesired` or the start of `Create()` and `Update()`.
*   **Export Handling**:
    *   In `Export()`, construct reference fields directly from proto strings or URLs without querying the Kubernetes API reader (`client.Reader`).

---

## 3. Short Name & URL Equivalence (Compute Resources Only)

*   **Scope**: Primarily applies to Compute Engine resources (and resources referencing compute networking, such as GKE clusters). Most modern GCP APIs do not exhibit this.
*   **Discrepancy**: Compute APIs accept short names (e.g., `default`) or relative paths, but return fully qualified URLs (`https://www.googleapis.com/compute/v1/...`), triggering false diffs during comparison.
*   **Normalization**: For compute network/subnetwork references, call `ref.CanonicalizeAndNormalize(...)` (or canonicalize URLs) directly after `common.NormalizeReferences()` to align URLs and prevent re-reconciliation drift.

---

## 4. `metadata.labels` Mapping & Handling

*   **Verify Mapping Definition**:
    *   Check `config/servicemappings/<service>.yaml` to confirm whether `metadataMapping.labels` is configured for the resource (e.g., mapping to `labels` or `user_labels`).
*   **Verify `Create()` Logic**:
    *   If mapped, verify the controller uses `label.GCPLabels(u)` (e.g. inside `buildDesired`) to compute and inject GCP labels into the creation request payload (`desired.Labels = label.GCPLabels(u)`).
*   **Verify `Update()` Comparison**:
    *   If mapped, verify desired labels are populated with `label.GCPLabels(u)`.
    *   Verify the controller uses `common.CompareBrownfieldSpecAndLabels` (or compares labels explicitly using `maps.Equal(desired, actual)` rather than `reflect.DeepEqual`) so that label changes populate the `UpdateMask` and do not trigger false drift.
*   **No Deprecated Helpers**: Reject usage of deprecated helpers like `label.NewGCPLabelsFromK8sLabels` or manual map copying.

---

## 5. Spec Field Comparison & Diff Output

*   **Verify Comparison Helper Usage**:
    *   In `Update()`, verify the controller uses `common.CompareBrownfieldSpec` / `common.CompareBrownfieldSpecAndLabels` (or a dedicated `CompareSpecFields(desired, actual)` helper) to compute diffs.
    *   Flag and reject ad-hoc struct equality checks or manual field-by-field comparisons that risk ignoring unmanaged or server-defaulted fields.
*   **Inspect Diff & Update Mask Generation**:
    *   Ensure update requests and `UpdateMask` paths are generated strictly from the diff produced by the comparison helper.
*   **Structured Reporting**:
    *   Verify that `structuredreporting.ReportDiff(ctx, diff)` is invoked in `Update()` before sending the update request to the GCP API.

---

## 6. Parity with Legacy Terraform Special Handling

*   **Inspect Legacy Provider Source**:
    *   Review the legacy TF code (`third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/...`) for custom diff suppressions (`DiffSuppressFunc`), separated PATCH requests, or specialized endpoints.
    *   Verify the direct controller replicates necessary diff suppressions or normalization behaviors to avoid regressions or spurious drift.
*   **Specialized Endpoints (Primarily Container / GKE Resources)**:
    *   Container resources frequently split modifications across multiple dedicated endpoints (e.g., `SetNodePoolSize`, `SetNodePoolManagement`) rather than a single generic `Patch`.
    *   Verify all applicable specialized endpoints are invoked so mutable fields are not silently dropped.
*   **Do NOT Preserve TF "Set to Zero/Empty on Unset" Behavior**:
    *   Unlike TF, which often sets fields to 0/empty when unset in configuration, KCC direct controllers treat unset fields as unmanaged.
    *   Verify the controller does NOT actively clear or zero out omitted fields.

---

## 7. Mappers & Backward Compatibility

*   **Handwritten Mappers (`*_mapper.go` / `mappers.go`)**:
    *   Ensure enum mappings support both short strings and proto enum constants.
    *   Ensure flattened structures (e.g., `maxPodsPerNode` <-> `maxPodsConstraint`) map bidirectionally without data loss.
*   **Completeness & Schema Parity**:
    *   Verify that all mutable and immutable fields present in the existing CRD are mapped in `Spec_ToProto` and `Spec_FromProto`.
    *   Verify that no fields are dropped or altered in type compared to the existing CRD.

---

## 8. Routing Configuration (`static_config.go`)

*   **Verify Static Controller Routing**:
    *   Check `pkg/controller/resourceconfig/static_config.go`.
    *   Verify `DefaultController` remains `ReconcilerTypeTerraform` (or `ReconcilerTypeDCL`) so existing production users are unaffected.
    *   Verify `ReconcilerTypeDirect` is added to `SupportedControllers` to allow opt-in execution and side-by-side golden testing.

---

## 9. Side-by-Side Golden & Diff Verification

*   **Inspect `_final_object.diff`**:
    *   Verify diffs between the legacy and direct controller outputs only contain expected schema improvements (e.g., `status.externalRef`, standard conditions).
    *   Ensure no existing spec or status fields are inadvertently dropped.
*   **Inspect `_http.diff` / `_http_mock.diff`**:
    *   Verify traffic alignment against the legacy controller.
    *   Confirm there are 0 unexpected PATCH or PUT calls during steady-state re-reconciliation.
*   **Inspect `_exported_object.diff`**:
    *   Confirm that exported KRM specifications match expected deployable formats without dropping spec fields.

---

## 10. General Controller Invariants

*   **Client Creation Preference**:
    *   Verify that the controller uses an official GAPIC Go client library REST constructor (e.g., `cloud.google.com/go/<service>/apivX` via `NewFooRESTClient`) instead of gRPC (`NewFooClient`) or raw `grpc.Dial` where REST is supported.
*   **Wait for LROs**:
    *   For operations that return a Long Running Operation (LRO) from the GCP client, the controller MUST wait for the LRO to complete before returning (e.g., calling `op.Wait(ctx)` after `Create` or `Update`).
*   **Status Updates**:
    *   The controller must update KRM status (`status.observedState`, `status.externalRef`, conditions) at the end of reconciliation.

---

# Review Comment Template

When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Ratcheting Exclusion Removed**: [Pass/Fail] - (Confirm removed from ShouldTestRereconiliation in tests/e2e/ratcheting.go)
* **Adapter Architecture & Reference Timing**: [Pass/Fail] - (No reference lookups in AdapterForObject; resolved only in Create/Update)
* **Short Name & URL Equivalence**: [Pass/Fail/N/A] - (Compute network/subnetwork URLs normalized if applicable)
* **Metadata Labels Handling**: [Pass/Fail/N/A] - (Verify label.GCPLabels(u) usage in Create/Update and maps.Equal comparison)
* **Spec Field Comparison & Structured Reporting**: [Pass/Fail] - (CompareBrownfieldSpec/CompareSpecFields and ReportDiff used in Update)
* **Legacy TF Parity & Specialized Endpoints**: [Pass/Fail] - (TF diff suppressions replicated; specialized endpoints called; omitted fields unmanaged)
* **Mappers & Backward Compatibility**: [Pass/Fail] - (Bidirectional mapping of all CRD fields; enums support strings/constants)
* **Static Controller Routing**: [Pass/Fail] - (DefaultController kept as legacy; Direct added to SupportedControllers in static_config.go)
* **Diff & Golden Log Verification**: [Pass/Fail] - (_http.diff, _final_object.diff, and _exported_object.diff verified with 0 steady-state writes)
* **LRO Wait & Client Creation**: [Pass/Fail] - (GAPIC REST client constructor used; op.Wait(ctx) invoked)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
