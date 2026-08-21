# KMSKeyHandle Direct Controller Journal

## Key Observations & Design Decisions

1. **Immutability and CannotBeDeleted**:
   - `KMSKeyHandle` resources are completely immutable once created in Google Cloud KMS (Autokey API). They do not support updates or deletion.
   - We registered the model using `registry.CannotBeDeleted()` to signify that the resource cannot be deleted from GCP.
   - The deletion operation `Delete` behaves as a no-op / returns `true, nil` since deletions are not supported.

2. **Immutable Update Handling**:
   - Even though the resource is immutable, KCC direct controller standards require the `Update` reconciliation loop to perform diff comparison.
   - We implemented `compareKeyHandle` using `mappers.OnlySpecFields` and `tags.DiffForTopLevelFields`.
   - If any difference/change is detected between the actual and desired specifications, we report the diff using `structuredreporting.ReportDiff` and return a clear descriptive error so that it is surfaced on the resource status rather than silently doing nothing.

3. **Status Integration**:
   - Status updates are integrated via a unified `updateStatus` helper function which maps from GCP Autokey response to KRM Status format and populates `externalRef` with the fully qualified GCP resource path.
