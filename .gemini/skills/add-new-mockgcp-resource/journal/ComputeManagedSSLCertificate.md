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

Centralized SSL certificate defaulting in `mockcompute/sslcertificate.go` (`populateSSLCertificateDefaults`), used by both `GlobalSSLCertificatesV1.Insert` and `RegionalSSLCertificatesV1.Insert`:
- If `obj.Managed != nil` and `obj.Type == nil`, default `Type` to `MANAGED`.
- For `SELF_MANAGED` certificates, populate `ExpireTime = CreationTimestamp`.
- For `MANAGED` certificates:
  - Default `Managed.Status` to `PROVISIONING` if omitted.
  - Normalize domains in `Managed.Domains` by trimming trailing dots (`.`).
  - Normalize keys in `Managed.DomainStatus` by trimming trailing dots (`.`).
  - Handle field presence and domain status according to `Managed.Status`:
    - `PROVISIONING`: `certificate`, `expireTime`, and `subjectAlternativeNames` are omitted. Unset domain statuses default to `PROVISIONING`.
    - `ACTIVE`: `certificate` is populated with a dummy PEM cert, `expireTime` is set to ~90 days from creation, and `subjectAlternativeNames` is populated with the domains. Unset domain statuses default to `ACTIVE`.
    - `PROVISIONING_FAILED` / `PROVISIONING_FAILED_PERMANENTLY`: `certificate`, `expireTime`, and `subjectAlternativeNames` are omitted. Unset domain statuses default to `FAILED_NOT_VISIBLE`.
    - `RENEWAL_FAILED`: `certificate` is populated with an existing/dummy cert, `expireTime` is set to ~30 days from creation, and `subjectAlternativeNames` is populated with domains. Failing/unset domain statuses default to `FAILED_NOT_VISIBLE`.

## Validation

- Added unit tests in `mockgcp/mockcompute/sslcertificate_test.go` covering each status case (`SELF_MANAGED`, `PROVISIONING`, `ACTIVE`, `PROVISIONING_FAILED`, `PROVISIONING_FAILED_PERMANENTLY`, and `RENEWAL_FAILED`).
- Recorded authentic baseline against real GCP (`cnrm-barni-4`) using `hack/record-gcp pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computemanagedsslcertificate`.
- Ran `hack/compare-mock pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computemanagedsslcertificate` successfully.
- Verified golden log alignment tests pass:
  - `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/compute/v1beta1/computemanagedsslcertificate"`
  - `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/compute/v1beta1/computesslcertificate"`
  - `go test -v ./pkg/test/resourcefixture -run "TestRealHTTPLogsDoNotContainMockGCP"`
- Verified end-to-end fixture reconciliation and re-reconciliation pass against MockGCP:
  - `RUN_E2E=1 E2E_GCP_TARGET=mock E2E_KUBE_TARGET=envtest go test -v ./tests/e2e -run TestAllInSeries/fixtures/computemanagedsslcertificate`
