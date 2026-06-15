### 2026-06-05 Implement Direct Types, CRD, and IdentityV2 for DataLabelingDataset
- **Context**: Implementing the initial KRM types, CRD, and IdentityV2 for `DataLabelingDataset` (Issue #9268).
- **Problem**: The issue description mentioned using `google.cloud.datalabeling.v1` as the service name, but only `v1beta1` exists in the `googleapis` proto repository.
- **Solution**: Used `google.cloud.datalabeling.v1beta1` in the `generate.sh` configuration to successfully compile descriptor files and scaffold KRM Go types.
- **Impact**: The correct service version must be `v1beta1` for all subsequent reconciliation and mapper controller development.
- **Location**: Note that the `DataLabelingDataset` resource does not have a location/region field, as the Data Labeling service organizes datasets globally/directly under a project (`projects/{project}/datasets/{dataset}`).

