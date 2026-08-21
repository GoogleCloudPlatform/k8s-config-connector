# Skill: Opt-in to Strict Testing

This skill guides an automated agent through the process of opting a resource (or all resources in an API group) into strict testing. Strict testing enables server-side apply for creates and re-reconciliation tests to ensure the controllers are 100% correct in differencing detection.

## Steps

1. **Remove from Ratcheting Exclusions**
   - Open `tests/e2e/ratcheting.go`.
   - Locate the function `ShouldTestRereconiliation`.
   - Find the `switch` statement that checks `primaryResource.GroupVersionKind()`.
   - Remove the `case` line(s) corresponding to the `schema.GroupKind` of the resource(s) being opted in.

2. **Regenerate Golden Output**
   - Because strict testing enables server-side apply, the expected requests and generated objects will change in the tests. You must regenerate the golden test output (`_generated_object_*.golden.yaml` and `_http.log`).
   - Run the relevant presubmit script for the service with the `WRITE_GOLDEN_OUTPUT=1` environment variable.
     For example:
     `WRITE_GOLDEN_OUTPUT=1 dev/ci/presubmits/tests-e2e-fixtures-<servicename>`
   - Alternatively, you can use the `hack/compare-mock` or `run-e2e` scripts if you are targeting specific tests.

3. **Verify the Tests Pass**
   - After regenerating the golden files, verify that the tests actually pass.
   - Run the tests again without `WRITE_GOLDEN_OUTPUT=1`.
     For example:
     `dev/ci/presubmits/tests-e2e-fixtures-<servicename>`
   - If a test fails during the `TestRereconciliation` phase, the controller might have a bug in differencing detection.

4. **Investigate Re-reconciliation Bugs (If Any)**
   - If the re-reconciliation test fails, it means the controller thinks the resource needs an update even when its state matches the desired state, or the observed state doesn't match after creation.
   - You might need to add/fix logic in the controller or mappers to properly normalize the field, handle server-generated values, or ignore specific fields in diffs.

5. **Commit the Changes**
   - Create a commit that includes the change to `tests/e2e/ratcheting.go` and the updated golden files (`_generated_object_*.golden.yaml`, `_http.log`).