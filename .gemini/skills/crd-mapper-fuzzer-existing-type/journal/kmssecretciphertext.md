# Journal: KMSSecretCiphertext Transition to Direct KRM Types

## Learnings & Observations

### 1. Resource Without Proto Message
- `KMSSecretCiphertext` is a client-side wrapper representing a GCP KMS ciphertext resource (originally a Terraform resource `google_kms_secret_ciphertext`).
- It does not correspond to a standard REST resource or a direct Proto message in `google.cloud.kms.v1` (since encrypting is an RPC call, not a standalone resource lifecycle).
- Thus, the types in `apis/kms/v1alpha1/kmssecretciphertext_types.go` were hand-coded to ensure exact structural alignment with the baseline CRD schema.

### 2. Validation Constraints (`value`/`valueFrom` signature)
- To maintain strict schema compatibility (preventing changes on `diff-crds`), we mapped the `legacyRefRule` (`oneOf` rule requiring exactly one of `value` or `valueFrom`) onto the fields of `KMSSecretCiphertext`.
- This was done by updating `scripts/add-validation-to-crds/parse-crds.go` to include `KMSSecretCiphertext` in the check block for `value,valueFrom` signatures:
  ```go
  } else if signature == "value,valueFrom" && (kind == "AlloyDBUser" || kind == "ComputeInstance" || kind == "ContainerCluster" || kind == "MonitoringUptimeCheckConfig" || kind == "KMSSecretCiphertext") {
      ruleYAML = legacyRefRule
  ```

### 3. Verification & Compatibility
- Running `dev/tasks/diff-crds` produces an empty output, verifying that zero structural changes or regressions were introduced.
- Updated the GVK and CRD reports (`docs/reports/crd_report.csv` and `docs/reports/crd_report.md`), which now successfully shows `KMSSecretCiphertext` has direct KRM types.
