# Export Support for PrivateCACertificateAuthority

## Observations & Implementation Details

1. **Reference-Bound Project Binding**:
   - `PrivateCACertificateAuthority` uses reference-based project binding via `spec.projectRef`.
   - In the `Export` method, we mapped the project ID to the `External` field: `obj.Spec.ProjectRef = refsv1alpha1.ProjectRef{External: a.id.Project}`.
   - We did NOT call the `export.SetProjectID(...)` helper, which avoided adding the deprecated `cnrm.cloud.google.com/project-id` annotation, in strict adherence to the project guidelines.

2. **Standard Label and Metadata Mapping**:
   - We leveraged `export.SetLabels(u, a.actual.Labels)` to set standard GCP labels on the Kubernetes resource metadata.
   - We mapped the resource short name `a.id.CertificateAuthority` via `u.SetName(...)`, rather than the full GCP resource URL.

3. **E2E Export Testing Integration**:
   - We registered `PrivateCACertificateAuthority` in `tests/e2e/export.go` within `switch gvk.GroupKind()` to resolve its CAIS URI:
     ```go
     case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACertificateAuthority"}:
         exportURI = resolveCAISURI(h, obj)
     ```
   - Running E2E tests with `GOLDEN_OBJECT_CHECKS=1` successfully generated and validated the golden `_exported.yaml` file in `pkg/test/resourcefixture/testdata/basic/privateca/v1beta1/privatecacertificateauthority/`.
   - Since `PrivateCACertificateAuthority` is a direct-only controller (no legacy fallback TF/DCL exporter exists for it), the test harness correctly ignored the old controller export failure.
