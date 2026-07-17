# CertificateManagerCertificateMap Export Journal

## Observations & Implementation Details

1. **Standard Metadata & Annotations**:
   - `CertificateManagerCertificateMap` maps to `google.cloud.certificatemanager.v1.CertificateMap`.
   - Instead of setting the `cnrm.cloud.google.com/project-id` annotation (which uses `export.SetProjectID`), resources with `spec.projectRef` in their KRM schema should have `spec.projectRef.external` set to the GCP project ID, and no annotation set.
   
2. **Project Reference Behavior & `SetProjectRef`**:
   - Introduced a new helper `export.SetProjectRef(u, projectID)` under `pkg/export/helpers.go`.
   - This helper sets the `spec.projectRef.external` path directly on the unstructured representation of exported resources.
   - For `CertificateManagerCertificateMap`, `spec.projectRef` in `_exported.yaml` is now correctly exported as:
     ```yaml
     projectRef:
       external: ${projectId}
     ```
     and the `cnrm.cloud.google.com/project-id` annotation is omitted entirely.

3. **Short-Name and ResourceID Mapping**:
   - The GCP actual `Name` returns the full resource path (`projects/{project}/locations/global/certificateMaps/{certificatemap}`).
   - It is crucial to use the parsed short name `a.id.CertificateMap` when calling `u.SetName` to generate valid KRM metadata.
   - Set `obj.Spec.ResourceID = direct.LazyPtr(a.id.CertificateMap)` explicitly since identity fields are not automatically converted from proto.

4. **Isolated E2E Testing**:
   - When running the unified E2E test runner, specify the exact regex `^TestAllInSeries/fixtures/certificatemanagercertificatemap$` to prevent other tests (e.g., `certificatemanagercertificatemapentry`) from being executed and writing undesired golden output diffs.
