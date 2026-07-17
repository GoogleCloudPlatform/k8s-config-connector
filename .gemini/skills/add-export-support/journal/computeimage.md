# ComputeImage Export Support Journal

## Context
Implemented export support for the `ComputeImage` resource using the Direct controller.

## Learnings & Observations
- **Critical Object Assignment Bug**: Discovered that the existing placeholder/stub `Export` method was assigning `u.Object = uObj` *after* calling `u.SetName(...)` and `u.SetGroupVersionKind(...)`. This was correctly identified and fixed, assigning `u.Object = uObj` first to prevent metadata truncation or erasure.
- **ResourceID Injection**: Injected `ResourceID` on the exported Spec manually via `obj.Spec.ResourceID = direct.LazyPtr(a.id.Image)`. This maps the identity cleanly into the exported yaml, matching the legacy exporter's format.
- **Project Annotation Assignment**: Since `ComputeImage` is not a reference-bound resource (it has no `spec.projectRef`), we correctly populated the `cnrm.cloud.google.com/project-id` annotation via `export.SetProjectID(u, a.id.Project)`.
- **Label Mapping**: Utilized the `export.SetLabels(u, a.actual.Labels)` helper to perfectly align metadata labels.
