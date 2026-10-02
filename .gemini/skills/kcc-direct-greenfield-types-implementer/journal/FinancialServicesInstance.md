# FinancialServicesInstance Journal

## Context
Implementing Greenfield direct KRM types, identity, reference, and generate.sh support for `FinancialServicesInstance` under `financialservices.cnrm.cloud.google.com/v1alpha1`.

## Resource Details
- Proto: `google.cloud.financialservices.v1.Instance`
- Service: `google.cloud.financialservices.v1`
- Resource Type: `financialservices.googleapis.com/Instance`
- URL Template: `projects/{project}/locations/{location}/instances/{instance}`
- Fields:
  - `kms_key` (REQUIRED) mapped to `kmsKeyRef` (*refsv1beta1.KMSCryptoKeyRef)
  - `labels` (map[string]string)
  - `state` (OUTPUT_ONLY) mapped to `observedState.state` with enum validation
  - `create_time` (OUTPUT_ONLY) mapped to `observedState.createTime`
  - `update_time` (OUTPUT_ONLY) mapped to `observedState.updateTime`
  - `name` (OUTPUT_ONLY) tracked via `status.externalRef`
