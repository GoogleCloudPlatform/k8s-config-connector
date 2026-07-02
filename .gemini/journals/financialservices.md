### 2026-06-15 Greenfield scaffolding of FinancialServicesInstance
- **Context**: Greenfield scaffolding of KRM types and IdentityV2 for `FinancialServicesInstance` under issue #10268.
- **Problem**: The GCP Instance protobuf definition for `google.cloud.financialservices.v1.Instance` specifies `kms_key` as a required CMEK field, but the default type generator scaffolds it as a plain `*string`.
- **Solution**: Replaced the scaffolded `KMSKey *string` field with `KMSKeyRef *refsv1beta1.KMSCryptoKeyRef` to properly implement a KCC-native reference to KMS crypto key, preventing exception file list additions and ensuring proper dependency mapping.
- **Impact**: The next agent implementing the actual controller and mapper can map this reference to GCP's KMS key URI cleanly.
