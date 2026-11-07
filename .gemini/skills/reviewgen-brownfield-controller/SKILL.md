---
name: reviewgen-brownfield-controller
description: Provides provides clear review criteria for reviewing PRs that add controllers for KCC Brownfield resources.
---

# Review guide for KCC Brownfield controller
Please respect the following review criteria and invariants when reviewing PRs for Brownfield migrations (legacy to direct controller).

## 1. Diff Analysis & Behavioral Alignment
*   **Mandatory Diff Inspection:** PRs modifying brownfield controllers MUST generate and include `.diff` files if the legacy controller is still active. Reviewers must inspect `_http.diff`, `_http_mock.diff`, `_final_object.diff`, and `_exported_object.diff` for all fixtures.
*   **Acceptable Diffs:** Check that any diffs are justified as acceptable improvements (e.g., adopting Identity v2, improved conditions, better error messages).
*   **Regressions:** Reject PRs that introduce unacceptable diffs such as dropped fields in HTTP requests/responses, missing status attributes previously exposed by legacy, or spurious update requests during re-reconciliation.
*   **100% Field Parity**: Reject any PR where existing fields supported by the legacy controller are missing from the direct implementation. Dropping a field causes KCC to wipe that configuration from GCP on the first reconcile loop after upgrade.

## 2. Zero-Write Steady State (Ratcheting)
*   **Ratcheting Exclusion Removed:** Verify that the resource has been removed from the exclusion list in `tests/e2e/ratcheting.go` (specifically in `ShouldTestRereconiliation`).
*   **No Spurious Updates:** Verify that steady-state re-reconciliation does not produce any extra writes to the GCP API (no `PATCH`/`PUT` calls in `_http.log` during the `update` phase if nothing in the spec changed).

## 3. Proto Diffs and Field Masks
*   Verify that `common.CompareBrownfieldSpec` or appropriate diffing logic is used to correctly identify differences between the desired KRM state and the actual GCP state, particularly handling default values returned by the server.
*   Verify that the resulting diff paths are used to populate an `UpdateMask` (e.g. `&fieldmaskpb.FieldMask{Paths: sets.List(paths)}`) in the `Update` operation if required by the API.

## 4. Structured Reporting
*   Verify that `structuredreporting.ReportDiff(ctx, report)` is used in the `Update` method to report the exact diff to the user before submitting the update request to the GCP API.

## 5. Wait for LROs
*   **Strict Rule:** For operations that return a Long Running Operation (LRO) from the GCP client, the controller MUST wait for the LRO to finish before proceeding (e.g., calling `op.Wait(ctx)` after `Create` or `Update`).

## 6. General Controller Structure & Client Creation
*   **Client Creation Preference**: Verify that the controller uses an official GAPIC Go client library REST constructor (e.g., `cloud.google.com/go/<service>/apivX` via `NewFooRESTClient`) instead of gRPC where possible.
*   **Desired Object in Adapter**: Verify that the `Adapter` struct stores the KRM object (`desired *krm.<Kind>`) rather than the GCP API/Proto representation, and that premature mapping is not performed in `AdapterForObject`.
*   **Reference Resolution Timing (CRITICAL)**: Verify that reference resolution is strictly performed on **line 1 of `Create()` and `Update()`** (or inside `buildDesired()`), NEVER in `AdapterForObject()`. `Delete()` must never attempt to resolve external references; resolving references in `AdapterForObject()` causes deletion deadlocks when referenced objects are deleted first (per PR #13573).
*   **Do Not Change Default Controller**: Verify that the PR does NOT switch the default controller to `direct` in `pkg/controller/resourceconfig/static_config.go` (or via the `cnrm.cloud.google.com/default-controller=direct` label comment in KRM type definitions). The legacy controller must remain the `DefaultController` for now, and the direct controller should only be added to `SupportedControllers` to allow side-by-side testing.
*   The controller must implement the `directbase.Model` and `directbase.Adapter` interfaces.

## 7. SOUL.md Invariant 7: Internal Information Sanitization (Blocking Finding)
*   Treat any Buganizer ID (`b/<id>`), internal bug URL (`buganizer.corp.google.com`), or customer name anywhere on GitHub (PR title, description, commit messages, code comments, or release notes) as an absolute **BLOCKING FINDING**. Request changes until removed.
*   Reviewer comments must **NEVER** contain Buganizer IDs or customer names. Refer strictly to public GitHub issues or PRs.

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **No Internal Info (SOUL.md Invariant 7)**: [Pass/Fail] - (Verified no b/<id>, internal URLs, or customer names)
* **Diff Analysis**: [Pass/Fail] - (List any unjustified or regression diffs in _http.diff, _final_object.diff, etc.)
* **Zero-Write Steady State**: [Pass/Fail] - (List if ratcheting exclusion is not removed or spurious updates occur)
* **Reference Resolution Timing**: [Pass/Fail] - (Verified references resolved in Create/Update, never in AdapterForObject)
* **Client Creation**: [Pass/Fail] - (List if raw pb/grpc.Dial is used instead of GAPIC go-client)
* **Default Reconciler Kept**: [Pass/Fail] - (Verify that the default controller is NOT switched to direct in static_config.go)
* **Proto Diffs & Update Mask**: [Pass/Fail] - (List any issues with diff calculation)
* **Structured Reporting**: [Pass/Fail] - (List if structured reporting is missing)
* **LRO Wait**: [Pass/Fail] - (List if LRO waits are missing)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
