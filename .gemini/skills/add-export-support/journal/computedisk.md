# ComputeDisk Export Support Journal

## Context
Implemented export support for the `ComputeDisk` resource using the Direct controller.

## Learnings & Observations
- **URI Type Canonicalization**: GCP returns fully qualified URIs for fields like `spec.type` (e.g., `https://www.googleapis.com/compute/v1/projects/${projectId}/zones/us-central1-a/diskTypes/pd-standard`), but the legacy exporter expects the short name (e.g., `pd-standard`). Stripping fully qualified URIs to their last path segment in `Export` is critical to match the legacy exporter's format and avoid producing diffs.
- **Reference-Bound project ID**: Since `ComputeDisk` supports `spec.projectRef`, `export.SetProjectID` was correctly bypassed, and `obj.Spec.ProjectRef` was populated instead.
- **Location Mapping**: Handled location mapping during `Export` based on zonal vs regional identity types from `a.id`.
