# KMSSecretCiphertext Direct Controller Journal

## Key Observations and Decisions

- **Direct Controller Implementation**: Created `KMSSecretCiphertext` direct controller in `pkg/controller/direct/kms/kmssecretciphertext_controller.go` using GAPIC REST `KeyManagementRESTClient`.
- **Identity and Server-Generated IDs**: Identity interface (`KMSSecretCiphertextIdentity`) was already defined. The resource identifier is the server-generated `ciphertext` string. `Find()` employs the service-generated ID guard (`if a.id.Ciphertext == "" { return false, nil }`), and `Create()` persists the generated ciphertext to `spec.resourceID` via `createOp.SetSpecResourceID`.
- **Immutable Resource**: All spec fields are immutable in GCP. `Update()` detects any attempted spec changes and returns an error, preventing drift.
- **Delete Behavior**: Deletions are a no-op since Cloud KMS does not support deleting ciphertexts, and the controller is registered with `registry.CannotBeDeleted()`.
- **Dual-Controller Testing**: Added `k8s.ReconcilerTypeDirect` to `SupportedControllers` in `static_config.go` while keeping `DefaultController` as Terraform for gradual opt-in. Recorded real GCP baseline traffic via `hack/record-gcp` and verified alignment with MockGCP via `hack/compare-mock`.
