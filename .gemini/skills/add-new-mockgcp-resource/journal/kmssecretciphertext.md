# Journal: Match real GCP behavior in MockGCP for KMSSecretCiphertext

## Overview
This journal documents the implementation, design decisions, and learnings from matching MockGCP behavior with real GCP for `KMSSecretCiphertext`.

## Resource Characteristics
- **GCP Service**: Cloud Key Management Service (`cloudkms.googleapis.com`)
- **Kind**: `KMSSecretCiphertext` (`kms.cnrm.cloud.google.com/v1alpha1`)
- **Controller Type**: Terraform-based (`google_kms_secret_ciphertext`)
- **GCP APIs Used**:
  - `POST /v1/{cryptoKey=projects/*/locations/*/keyRings/*/cryptoKeys/*}:encrypt`: Encrypts plaintext and returns ciphertext, name (the CryptoKeyVersion), and CRC32C checksum.
  - `GET /v1/{cryptoKey=projects/*/locations/*/keyRings/*/cryptoKeys/*}`: Fetches CryptoKey state on Read.
  - Delete: No-op in Google Cloud (ciphertext cannot be deleted).

## Key Learnings & Resolutions
1. **ServiceMapping SkipImport**:
   - `google_kms_secret_ciphertext` does not support Terraform import upstream (`resource google_kms_secret_ciphertext doesn't support import`).
   - During deletion and live state retrieval, KCC attempts to import the resource state by ID if `skipImport` is unset.
   - Added `skipImport: true` to `google_kms_secret_ciphertext` in `scripts/resource-autogen/generated/servicemappings/kms.yaml`.
2. **Real GCP Baseline Recording**:
   - Recorded golden logs (`_http.log`) and golden object states using `hack/record-gcp` against real GCP project `cnrm-barni-4`.
   - Committed the recorded golden files in their own baseline commit as mandated.
3. **MockGCP Implementation**:
   - Implemented `Encrypt` RPC in `mockgcp/mockkms/secretciphertext.go`.
   - Resolves target `CryptoKey` or `CryptoKeyVersion` and validates CRC32C checksums if provided in request.
   - Returns golden ciphertext and CRC32C checksum for test cases, and deterministic ciphertext simulation for general requests.
4. **Verification**:
   - Ran `hack/compare-mock` to generate `_http_mock.log` and verify complete CRUD and ratcheted re-reconciliation cycles.
   - Verified log alignment with `go test -v ./pkg/test/resourcefixture -run "TestGoldenLogAlignment/kms/v1alpha1/kmssecretciphertext/kmssecretciphertext"` passing with 0 diffs.
   - Verified that `go test -v ./tests/apichecks/...` passed with updated `alpha-missingfields.txt`.
