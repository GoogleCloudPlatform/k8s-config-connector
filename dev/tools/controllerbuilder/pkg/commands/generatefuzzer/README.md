# Generate Fuzzer

This command generates and qualifies a fuzzer for a given resource.

## Example Usage

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.managedkafka.v1.Topic" \
  --api-version "managedkafka.cnrm.cloud.google.com/v1alpha1" \
  --kind "ManagedKafkaTopic" \
  --max-attempts 3
```

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.bigquery.datatransfer.v1.TransferConfig" \
  --api-version "bigquerydatatransfer.cnrm.cloud.google.com/v1beta1" \
  --kind "BigQueryTransferConfig" \
  --llm-model "gemini-2.0-flash-exp" \
  --max-attempts 3
```
