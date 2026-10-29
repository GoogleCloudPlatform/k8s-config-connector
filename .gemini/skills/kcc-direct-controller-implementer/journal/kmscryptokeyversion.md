# KMSCryptoKeyVersion Direct Controller Journal

## Key Observations

- **Pre-generated Mappers**: Conversions for `KMSCryptoKeyVersionSpec` and `KMSCryptoKeyVersionStatus` were already present and verified by fuzzer tests.
- **Service-Generated Resource IDs**: `KMSCryptoKeyVersion` uses service-generated integer IDs (e.g., `"1"`, `"2"`). Per `kcc-direct-service-generated-id`:
  - `Find()` guards on empty resource ID: `if !a.id.HasIdentitySpecified() { return false, nil }`.
  - `Create()` leaves `desiredPb.Name` empty so GCP assigns the next version ID.
- **KMS Delete Semantics**: Since KMS key versions cannot be deleted directly from GCP, `Delete()` issues `DestroyCryptoKeyVersionRequest` to transition the key version to `DESTROY_SCHEDULED`.
- **Normalization for E2E Tests**: In `mockkms` and `e2e` normalizers, timestamps such as `.status.generateTime`, `destroyTime`, and `cryptoKeyVersions[].destroyTime` were normalized with placeholders to ensure reproducible golden logs between test runs.
