# StorageInsightsReportConfig Journal

## Observations
- `StorageInsightsReportConfig` maps 1:1 to `google.cloud.storageinsights.v1.ReportConfig`.
- Proto sub-messages `CloudStorageFilters` and `CloudStorageDestinationOptions` reference GCS buckets, which map to `storagev1beta1.StorageBucketRef`.
- `FrequencyOptions` references `google.type.Date`, which requires `Date_FromProto` and `Date_ToProto` helper mapping functions in `pkg/controller/direct/storageinsights/mapper.go`.
