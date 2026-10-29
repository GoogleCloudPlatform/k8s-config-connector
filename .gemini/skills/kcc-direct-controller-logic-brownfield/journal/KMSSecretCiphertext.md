# KMSSecretCiphertext Migration Journal

## Overview
Migrated `KMSSecretCiphertext` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Resource Behavior and Service-Generated ID:**
   - `KMSSecretCiphertext` is a client-side logical resource in Cloud KMS that performs an `Encrypt` RPC against an existing `CryptoKey`.
   - GCP Cloud KMS does not maintain a discrete `SecretCiphertext` entity that can be fetched via GET or deleted via DELETE.
   - The unique identifier of the resource is the server-generated `ciphertext` string returned by the `Encrypt` RPC.
   - For service-generated IDs, following the `kcc-direct-service-generated-id` rules, `Find()` checks if `a.id.Ciphertext == ""` and returns `(false, nil)` before creation.
   - During `Create()`, after calling `KeyManagementClient.Encrypt()`, the server-generated ciphertext is recorded into `spec.resourceID` via `createOp.SetSpecResourceID` (for backwards compatibility) and updated into `status.ciphertext`.
   - `Create()` also performs a `GetCryptoKey` check to verify the existence of the parent key and maintain alignment with legacy controller HTTP traffic.

2. **Re-reconciliation and Immutability:**
   - On re-reconciliation, `Find()` retrieves the parent `CryptoKey` using `GetCryptoKey(a.id.ParentString())`. If the key exists, `Find()` returns `(true, nil)`.
   - Spec fields (`cryptoKey`, `plaintext`, `additionalAuthenticatedData`, `resourceID`) are immutable in GCP. `Update()` ensures that immutable fields have not been modified and returns a descriptive error if attempted.

3. **Delete Semantics:**
   - Cloud KMS does not support deleting ciphertexts from GCP.
   - In accordance with legacy Terraform semantics, `Delete()` is a no-op that logs the deletion and returns `(true, nil)` to allow Kubernetes finalizers to clear. The model is registered with `registry.CannotBeDeleted()`.

4. **Golden Log Normalization:**
   - Cloud KMS symmetric encryption generates a random initialization vector (IV) for every `Encrypt` call. Consequently, the resulting `ciphertext` and `ciphertextCrc32c` are non-deterministic across live GCP test executions.
   - Normalized `ciphertext` and `ciphertextCrc32c` in `golden_alignment_test.go` to ensure deterministic comparison between recorded live GCP logs and mock logs.
