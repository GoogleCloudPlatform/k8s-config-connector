# CertificateManagerCertificate Export Journal

## Observations & Implementation Details

1. **Dual Reconciler Support**: `CertificateManagerCertificate` is supported by both the Terraform (default) and Direct controllers.
2. **E2E Testing Workflow**:
   - The unified test harness (`TestAllInSeries`) runs the test scenario twice:
     1. First, with the **Direct** reconciler. Since direct export support is now implemented via `AdapterForURL`, this run correctly exports the resource and compares/creates `_exported.yaml`.
     2. Second, with the fallback **Terraform** reconciler. The export execution is ignored with a warning because `DisableDirectExport` is true during the fallback run.
3. **CAIS Integration**: Since `CertificateManagerCertificateIdentity` implements `identity.IdentityV2` and `CertificateManagerCertificate` implements `identity.Resource`, `resolveCAISURI(h, obj)` successfully resolved the CAIS URL starting with `//certificatemanager.googleapis.com/...` out of the box without further changes.
