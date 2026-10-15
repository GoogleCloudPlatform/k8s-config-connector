# SQLSSLCert Migration Journal

## Overview
Migrated `SQLSSLCert` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Server-Generated ID & Identity Rules:**
   - `SQLSSLCert` uses a server-generated SHA-1 fingerprint as its resource ID (`spec.resourceID`).
   - Implements Rules 1-4 for service-generated resource IDs:
     - Rule 1: Identity uses `spec.resourceID` directly without falling back to `metadata.name`.
     - Rule 2: `Find()` pre-checks if `a.id.Sha1Fingerprint == ""` and returns `(false, nil)` before creation.
     - Rule 3: `Create()` does not specify the fingerprint in the request payload. Upon creation, it extracts the fingerprint from `resp.ClientCert.CertInfo.Sha1Fingerprint` and sets `spec.resourceID` via `createOp.SetSpecResourceID`.
     - Rule 4: `GetIdentity()` verifies identity transitions against `status.sha1Fingerprint`.

2. **Write-Only/Creation-Only Fields (`privateKey` and `serverCaCert`):**
   - The private key (`certPrivateKey`) and server CA certificate (`serverCaCert`) are only returned by GCP in the `SslCertsInsertResponse`. Subsequent `Get` calls do not return the private key.
   - The adapter captures them during `Create()` and preserves them in `updateStatus()` across re-reconciliations.

3. **Resource Immutability:**
   - GCP Cloud SQL SSL certificates do not support update or patch operations.
   - The `Update()` method validates that the spec (specifically `commonName`) has not changed. If a diff is detected, it returns an explicit error indicating the resource is immutable.

4. **Brownfield Dual-Testing Alignment:**
   - `pkg/controller/resourceconfig/static_config.go` was updated to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `SQLSSLCert` while keeping `DefaultController` as `k8s.ReconcilerTypeTerraform`.
   - Both direct and legacy controllers pass MockGCP tests and golden alignment.
