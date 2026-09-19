# Journal Entry: ComputeImage Direct Takeover & Migration

## Issue & Diagnosis
During migration testing for `ComputeImage` under `TestMigrationToDirect/fixtures/computeimage`, several issues were encountered:
1. **Immutable fields in Patch request (400 Bad Request):** The direct controller's `Update` logic previously sent the entire desired state (including `sourceDisk` which is immutable) to `PatchImageRequest`. GCE GAPI fails with 400 when an immutable field is present in the patch payload (even if unchanged).
2. **Formatting differences of immutable fields:** Comparison detected false diffs on `source_disk`/`source_image`/`source_snapshot` due to URI format mismatches (`https://www.googleapis.com/...` vs `projects/...`).
3. **Defaults/Nil alignment:** Unspecified optional immutable fields like `disk_size_gb`, `storage_locations`, `licenses`, `rawDisk`, and `guestOsFeatures` generated diffs because GCP defaulted them on the actual object while KRM had them as nil/zero.
4. **Label checking mismatch:** System KRM metadata labels (like `managed-by-cnrm`) are not automatically set on GCP resources by the TF controller, causing `labelsChanged` (evaluated via comparing unstructured K8s labels to GCP labels) to be true, which triggered an unexpected `SetLabelsImageRequest` during direct takeover.

## Resolution
1. **Only compare mutable/mutable-defaulted fields:** Modified `compareComputeImage` to normalize resource URIs to the relative path (`projects/...`) format using canonicalization functions from the `compute` package.
2. **Align immutable defaulted fields:** Aligned defaulted immutable fields in `clonedDesired` (setting them to `maskedActual` values if not specified in KRM) to prevent any false diffs.
3. **Compare GCP-formatted labels:** Changed the label change detection logic in `Update` to compare `a.desired.Labels` (the GCP-mapped labels) with `a.actual.Labels` (the live GCP labels) instead of unstructured K8s metadata labels. This guarantees that system-added KRM labels don't trigger unexpected `setLabels` calls during takeover.
4. **Surgical Patch updates:** Refactored `Update` to construct a minimal patched `computepb.Image` containing only mutable fields (`Family` and `Description`) for `PatchImageRequest`, completely avoiding GCE API 400 errors about immutable fields.
5. **Log Normalization:** Added regex normalization rules for `sourceDiskId` and `labelFingerprint` in `tests/e2e/normalize_legacy.go` to make recorded cassettes deterministic.
