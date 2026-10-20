---
name: reviewgen-greenfield-controller-legacy
description: Review criteria for in-flight legacy Greenfield controller PRs written before deterministic controller/fuzzer generation (`--prepopulate-resource`).
---

# Review Guide for KCC Greenfield Controller (Legacy / In-Flight PRs)

Use this review guide **only** for in-flight PRs whose controller (`pkg/controller/direct/<service>/*_controller.go`) was written before deterministic controller generation (`// +generated:controller`).

## 1. Proto Diffs and Field Masks
*   Verify that `common.CompareGreenfieldSpec` (preferred), `tags.DiffForTopLevelFields`, or `common.CompareProtoMessage` is used to calculate differences between the desired state (`krm` spec mapped to proto) and the actual state (actual proto returned from API).
*   Verify that the resulting diff paths are used to populate an `UpdateMask` in the `Update` operation if required by the API.

## 2. Structured Reporting
*   Verify that `structuredreporting.ReportDiff(ctx, report)` is used in the `Update` method to report the exact diff to the user before submitting the update request to the GCP API.

## 3. Mandatory Status Updates
*   **Strict Rule:** The controller MUST always update the KRM status at the end of `Create` and `Update` operations, even if there are no fields that need updating in the GCP resource (updating `ObservedState`, `ExternalRef`, etc.).

## 4. Wait for LROs
*   **Strict Rule:** For operations that return a Long Running Operation (LRO) from the GCP client, the controller MUST wait for the LRO to finish before proceeding (e.g., calling `op.Wait(ctx)` after `Create` or `Update`).

## 5. Test Coverage (Minimal and Maximal, Edge Cases)
*   Ensure that the `pkg/test/resourcefixture/testdata/basic/` directory for this resource contains at minimum:
    *   A minimal test case (`create.yaml`, etc.)
    *   One or more maximal test cases containing all spec fields to ensure maximal branch coverage.
*   Check for appropriate edge-case testing if any complex logic exists in the controller (e.g., dependencies on other resources, complex identity mapping).

## 6. General Controller Structure & Client Creation
*   **Client Creation Preference**: Verify that the controller uses an official GAPIC Go client library REST constructor (`NewFooRESTClient`) instead of `NewFooClient` (gRPC) or manually calling `grpc.Dial` with raw protobuf client interfaces (`pb.New...Client`), unless REST is unavailable or fails.
*   **Reference Resolution Timing (CRITICAL)**: Verify that reference resolution is strictly performed on **line 1 of `Create()` and `Update()`**, NEVER in `AdapterForObject()`. `Delete()` must never attempt to resolve external references; resolving references in `AdapterForObject()` causes deletion deadlocks when referenced objects are deleted first (per PR #13573).
*   **Do Not Change Default Controller**: Verify that the PR does NOT switch the default controller to `direct` in `pkg/controller/resourceconfig/static_config.go` if the resource already has an existing Terraform or DCL controller.
*   The controller must implement the `directbase.Model` and `directbase.Adapter` interfaces.

## 7. SOUL.md Invariant 7: Internal Information Sanitization (Blocking Finding)
*   Treat any Buganizer ID (`b/<id>`), internal bug URL (`buganizer.corp.google.com`), or customer name anywhere on GitHub (PR title, description, commit messages, code comments, or release notes) as an absolute **BLOCKING FINDING**. Request changes until removed.
*   Reviewer comments must **NEVER** contain Buganizer IDs or customer names. Refer strictly to public GitHub issues or PRs.

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results (Legacy In-Flight Controller)
* **Trigger criteria matched**: [Yes/No]
* **No Internal Info (SOUL.md Invariant 7)**: [Pass/Fail] - (Verified no b/<id>, internal URLs, or customer names)
* **Client Creation**: [Pass/Fail] - (List if raw pb/grpc.Dial is used instead of GAPIC go-client)
* **Reference Resolution Timing**: [Pass/Fail] - (Verified references resolved in Create/Update, never in AdapterForObject)
* **Default Reconciler Kept**: [Pass/Fail] - (Verify that the default controller is NOT switched to direct in static_config.go if there was an existing legacy controller)
* **Proto Diffs & Update Mask**: [Pass/Fail] - (List any issues with diff calculation)
* **Structured Reporting**: [Pass/Fail] - (List if structured reporting is missing)
* **KRM Status Updates**: [Pass/Fail] - (List if status update is skipped on no-op updates)
* **LRO Wait**: [Pass/Fail] - (List if LRO waits are missing)
* **Test Coverage**: [Pass/Fail] - (List missing minimal/maximal test cases)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
