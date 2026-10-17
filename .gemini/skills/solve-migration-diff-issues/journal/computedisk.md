# ComputeDisk Direct Takeover Migration Journal - 2026

## Takeover Diffs Diagnosed

During the migration validation for `ComputeDisk`, we identified two main takeover diffs when comparing the legacy Terraform controller state with the new Direct controller state:

### 1. `replica_zones` URI Canonicalization and Sorting (observed in `regionalcomputedisk`)
- **Symptom**: The legacy TF state returned `replica_zones` as fully-qualified URLs (e.g. `https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-a`), whereas KRM specified relative short names (e.g. `us-central1-a`).
- **Fix**: Implemented a `trimZonePrefix(z string) string` helper that resolves zone identifiers down to their short name, and integrated it into the comparison's `populateDefaults` step. To prevent diffs due to ordering differences, we also sort the resulting `ReplicaZones` list using `sort.Strings`.

### 2. Optional `SizeGb` Defaulting (observed in `regionalcomputedisk` and `computediskfromsourcedisk`)
- **Symptom**: `SizeGb` is an optional field in the KRM spec and was left omitted. However, GCP or the TF controller populated a default value of `500`. Since the desired state has `SizeGb` as `nil` but the actual GCP state has `500`, a false-positive takeover diff occurred.
- **Fix**: In the comparison logic (`compareComputeDisk`), if `clonedDesired.SizeGb` is `nil`, we explicitly set `maskedActual.SizeGb` to `nil` to ignore GCP's server-side default value and prevent a false diff.
