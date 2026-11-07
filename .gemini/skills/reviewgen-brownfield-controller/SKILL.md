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
*   **Do Not Change Default Controller**: Verify that the PR does NOT switch the default controller to `direct` in `pkg/controller/resourceconfig/static_config.go` (or via the `cnrm.cloud.google.com/default-controller=direct` label comment in KRM type definitions). The legacy controller must remain the `DefaultController` for now, and the direct controller should only be added to `SupportedControllers` to allow side-by-side testing.
*   The controller must implement the `directbase.Model` and `directbase.Adapter` interfaces.

## 7. Code Quality, Style & Design Practices
Pay attention to race conditions, edge cases, and input/output correctness. Ensure correctness first, then do a second pass to ensure the code follows these style and design practices:

1. **1 File = 1 Class / Struct (Single Responsibility per File)**:
   * Follow a 1-file, 1-class (1 primary struct + its methods) design with a single responsibility per file; avoid dumping multiple responsibilities or bag-of-fields helper structs into a controller file.
   * Each `<file>.go` should have a matching `<file>_test.go` with tests structured into 3 visually separated blocks (Arrange, Act, Assert) separated by blank lines.

2. **Treat Files as Private by Default (No Cross-File Package Globals)**:
   * Even though Go shares lowercase symbols across a package, treat each `.go` file as a strict encapsulation boundary.
   * Across files in the same package, only interact via a struct's constructor (`newX(...)`) and its methods (`x.doWork(...)`).
   * Do NOT reach into another struct's unexported fields from a different file.
   * Do NOT call free helper functions or reference const declarations defined in another file unless that file's explicit sole purpose is a shared domain model.

3. **No Feature Envy on Value/Data Structs**:
   * Value structs (e.g., parsed identifiers, options) must only contain data and validation for that data.
   * Do not attach downstream domain logic (e.g., formatting IAM emails or computing GCS regions) onto a shared data struct; keep those as private methods/helpers in the file that actually performs that domain work.

4. **Own Collaborators via Constructor, Not Method Threading**:
   * If struct A uses collaborator B, wire B inside `newA(...)` (and store B as a field on A) rather than having the caller construct B and pass it into multiple methods of A (`a.foo(b)`, `a.bar(b)`).
   * Initialize optional collaborators (e.g., default nil `io.Writer` to `io.Discard`) in the constructor (`newFoo`) rather than repeating `if x.field != nil` guards across methods.

5. **Depth-First Call-Tree Ordering Within a File (Stepdown Rule)**:
   * A reader must be able to read any file from top to bottom without jumping to other files—or bouncing to the bottom of the same file—to hunt down helpers or constants.
   * Do NOT group all receiver methods at the top of a file and dump private helper functions at the bottom.
   * Place every private helper function or sub-method immediately below the function that calls it, in the exact order they are called: struct -> constructor -> constructor's helpers -> primary method -> its helpers in call order -> trivial 1-line getters/delegators at the bottom.

6. **Colocate Constants with Their First Caller**:
   * Do not dump all const declarations at the top of the file if they are only used by a specific sub-tree of functions further down.
   * Declare const blocks immediately above the first function or method that uses them.

7. **Orchestrator vs. Leaf Method Granularity**:
   * Reference: [Function Length](https://martinfowler.com/bliki/FunctionLength.html).
   * Orchestrator methods (methods that coordinate multiple logical steps, e.g., provision, reconcile) must stay at a single high level of abstraction and should NOT contain inline for loops or anonymous struct tables for sub-steps. Extract each loop/sub-step into a well-named helper placed immediately below in call order.
   * Leaf methods (methods that perform one concrete task) should treat the 5–15 line length as a guideline, not a hard quota. Do NOT shred a cohesive 20-line leaf method into 3–4 single-use 5-line wrappers.

8. **No `(nil, nil)` Return Contracts**:
   * Never return `(nil, nil)` from a fetch/lookup function. Return either a valid value or a descriptive error (or `(T, bool)` for pure in-memory lookups) so callers never need defensive `if val == nil` checks after `if err != nil`.

## 8. PR Description & Metadata
* **Release Note Block**: Verify that the PR description includes a release note block set to `NONE`:
  ```release-note
  NONE
  ```
  Fail or flag if the release note block is missing or not set to `NONE`.

# Review Comment Template
When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **Release Note Check**: [Pass/Fail] - (Confirm release-note block is set to NONE)
* **Diff Analysis**: [Pass/Fail] - (List any unjustified or regression diffs in _http.diff, _final_object.diff, etc.)
* **Zero-Write Steady State**: [Pass/Fail] - (List if ratcheting exclusion is not removed or spurious updates occur)
* **Client Creation**: [Pass/Fail] - (List if raw pb/grpc.Dial is used instead of GAPIC go-client)
* **Default Reconciler Kept**: [Pass/Fail] - (Verify that the default controller is NOT switched to direct in static_config.go)
* **Proto Diffs & Update Mask**: [Pass/Fail] - (List any issues with diff calculation)
* **Structured Reporting**: [Pass/Fail] - (List if structured reporting is missing)
* **LRO Wait**: [Pass/Fail] - (List if LRO waits are missing)
* **Code Quality & Design**: [Pass/Fail] - (Check single responsibility, encapsulation, Stepdown ordering, collaborator ownership, no (nil, nil))

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
