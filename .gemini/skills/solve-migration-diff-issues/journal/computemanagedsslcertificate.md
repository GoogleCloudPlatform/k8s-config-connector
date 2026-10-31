# ComputeManagedSSLCertificate Direct Migration Journal

During the direct migration validation for `ComputeManagedSSLCertificate`, we validated direct takeover and re-reconciliation against live GCP (`cnrm-barni-4`) and mock GCP.

## Observations & Root Cause

### Trailing Dot in Managed Domain Names
- **Issue:** During direct takeover (Phase 3), the direct controller detected a diff on `spec.managed.domains`:
  ```json
  {
    "id": "managed",
    "old": { "domains": ["sslcert.fixture.club"] },
    "new": { "domains": ["sslcert.fixture.club."] }
  }
  ```
  Because `ComputeManagedSSLCertificate` is an immutable resource, this diff caused the direct controller `Update` method to fail with `"ComputeManagedSSLCertificate is immutable and cannot be updated"`.
- **Root Cause:**
  - In user KRM manifests and test fixtures, domain names may be configured with a trailing dot (e.g. `sslcert.fixture.club.`, fully-qualified domain name format).
  - The GCP Compute API canonicalizes domain names by stripping the trailing dot upon creation and returning `sslcert.fixture.club`.
  - In the legacy Terraform controller, the Terraform provider configured `DiffSuppressFunc: tpgresource.AbsoluteDomainSuppress` on `managed.0.domains`, which treated domains with and without a trailing dot as equivalent.
  - The direct controller previously passed `nil` for the `normalize` parameter in `common.CompareBrownfieldSpec`, leading to a false diff between the desired spec and the actual GCP state.
- **Resolution:**
  - Implemented a `normalize` callback in `ComputeManagedSSLCertificateAdapter.Update` passed to `common.CompareBrownfieldSpec`. The callback trims any trailing dot (`strings.TrimSuffix(domain, ".")`) from all entries in `p.Managed.Domains`.
  - Re-ran `./hack/record-gcp TestMigrationToDirect/fixtures/computemanagedsslcertificate$` against live GCP project `cnrm-barni-4`. All phases (legacy create, legacy re-reconciliation, direct takeover, and direct re-reconciliation) completed with zero unexpected diffs.
  - Verified against mock GCP (`E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest RUN_E2E=1`) which also passed successfully.
