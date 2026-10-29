# ComputeManagedSSLCertificate MockGCP Alignment Journal

- **Date:** October 29, 2026
- **Resource Kind:** `ComputeManagedSSLCertificate`
- **GCP Resource:** `compute, sslCertificates`

## Overview

Existing mock implementation for SSL certificates exists under `mockgcp/mockcompute` in `globalsslcertificatesv1.go` and `regionalsslcertificatesv1.go`.
Previously, the mock only defaulted `Type` to `SELF_MANAGED` and always set `ExpireTime = CreationTimestamp`.
It did not handle `MANAGED` certificates or populate managed certificate fields.

## Discrepancies and Findings

When recording baseline HTTP logs against real GCP project `cnrm-barni-4` using `hack/record-gcp`:
1. **`expireTime`**: In real GCP, for managed certificates (`Type == "MANAGED"`) that are in `PROVISIONING` status, `expireTime` is not set. Previously, MockGCP unconditionally set `ExpireTime` on all certificates to `CreationTimestamp`.
2. **`managed.domains`**: Real GCP trims trailing dots (`.`) from domain names (e.g. `sslcert.fixture.club.` -> `sslcert.fixture.club`).
3. **`managed.status`**: Real GCP initializes `status` of managed certificate to `PROVISIONING`.
4. **`managed.domainStatus`**: Real GCP populates the `domainStatus` map with `PROVISIONING` for each configured domain.
5. **Ratcheting Exclusion**: Removed `ComputeManagedSSLCertificate` from `tests/e2e/ratcheting.go` to test re-reconciliation.

## Solution

Updated `Insert` in both `mockcompute/globalsslcertificatesv1.go` and `mockcompute/regionalsslcertificatesv1.go`:
- If `obj.Managed != nil` and `obj.Type == nil`, default `Type` to `MANAGED`.
- Only populate `ExpireTime = CreationTimestamp` when `Type == "SELF_MANAGED"`.
- When `Type == "MANAGED"`:
  - Default `Managed.Status` to `PROVISIONING`.
  - Normalize domains by trimming trailing dots.
  - Populate `Managed.DomainStatus` with `PROVISIONING` for each domain.

## Validation

- Recorded authentic baseline against real GCP (`cnrm-barni-4`) using `hack/record-gcp pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computemanagedsslcertificate`.
- Ran `hack/compare-mock pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computemanagedsslcertificate` successfully.
- Verified golden log alignment tests pass:
  - `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/compute/v1beta1/computemanagedsslcertificate"`
  - `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/compute/v1beta1/computesslcertificate"`
  - `go test -v ./pkg/test/resourcefixture -run "TestRealHTTPLogsDoNotContainMockGCP"`
