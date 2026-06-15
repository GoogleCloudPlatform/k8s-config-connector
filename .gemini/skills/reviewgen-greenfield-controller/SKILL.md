---
name: reviewgen-greenfield-controller
description: Provides provides clear review criteria for reviewing PRs that add controllers for KCC Greenfield resources.
---

# Review guide for KCC Greenfield new types PRs
Please respect the following review criteria and invariants when reviewing.

## 1. Proto Diffs and Field Masks
*   Verify that `common.CompareProtoMessage` is used to calculate differences between the desired state (`krm` spec mapped to proto) and the actual state (actual proto returned from API).
*   Verify that the resulting diff paths are used to populate an `UpdateMask` (e.g. `&fieldmaskpb.FieldMask{Paths: sets.List(paths)}`) in the `Update` operation if required by the API.

## 2. Structured Reporting
*   Verify that `structuredreporting.ReportDiff(ctx, report)` is used in the `Update` method to report the exact diff to the user before submitting the update request to the GCP API.

## 3. Mandatory Status Updates
*   **Strict Rule:** The controller MUST always update the KRM status at the end of `Create` and `Update` operations, even if there are no fields that need updating in the GCP resource. (e.g., a `setStatus(u, status)` call must be present and invoked, updating `ObservedState`, `ExternalRef`, etc.).

## 4. Wait for LROs
*   **Strict Rule:** For operations that return a Long Running Operation (LRO) from the GCP client, the controller MUST wait for the LRO to finish before proceeding (e.g., calling `op.Wait(ctx)` after `Create` or `Update`).

## 5. Test Coverage (Minimal and Maximal, Edge Cases)
*   Ensure that the `pkg/test/resourcefixture/testdata/basic/` directory for this resource contains at minimum:
    *   A minimal test case (`create.yaml`, etc.)
    *   One or more maximal test cases containing all spec fields to ensure maximal branch coverage.
*   **Reference Example:** A good example of minimal and maximal test suites can be found here: [`gkehubscope`](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/d5ce0db71838fcb63ea99de8be8fd53fa90bf597/pkg/test/resourcefixture/testdata/basic/gkehub/v1alpha1/gkehubscope).
*   Check for appropriate edge-case testing if any complex logic exists in the controller (e.g., dependencies on other resources, complex identity mapping).

## 6. General Controller Structure
*   **Reference Implementation:** The canonical reference controller implementation is [`workerpool_controller.go`](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d5ce0db71838fcb63ea99de8be8fd53fa90bf597/pkg/controller/direct/cloudbuild/workerpool_controller.go) (Note: it predates `IdentityV2` and does not handle it correctly, so ensure new controllers implement `IdentityV2` properly).
*   The controller must implement the `directbase.Model` and `directbase.Adapter` interfaces.
*   It should contain appropriate dependency resolution (`resolveDependencies`).

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Proto Diffs & Update Mask**: [Pass/Fail] - (List any issues with diff calculation)
* **Structured Reporting**: [Pass/Fail] - (List if structured reporting is missing)
* **KRM Status Updates**: [Pass/Fail] - (List if status update is skipped on no-op updates)
* **LRO Wait**: [Pass/Fail] - (List if LRO waits are missing)
* **Test Coverage**: [Pass/Fail] - (List missing minimal/maximal test cases)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
