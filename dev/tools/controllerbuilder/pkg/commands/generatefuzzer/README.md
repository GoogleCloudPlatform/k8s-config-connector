# Generate Fuzzer

This command generates and qualifies a fuzzer for a given resource.

## Example Usage

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.managedkafka.v1.Topic" \
  --api-version "managedkafka.cnrm.cloud.google.com/v1alpha1" \
  --max-attempts 3
```

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.bigquery.datatransfer.v1.TransferConfig" \
  --api-version "bigquerydatatransfer.cnrm.cloud.google.com/v1beta1" \
  --llm-model "gemini-2.0-flash-exp" \
  --max-attempts 3
```

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.redis.cluster.v1.Cluster" \
  --api-version "redis.cnrm.cloud.google.com/v1beta1" \
  --llm-model "gemini-2.0-flash-exp" \
  --max-attempts 5
```

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.securesourcemanager.v1.Instance" \
  --api-version "securesourcemanager.cnrm.cloud.google.com/v1alpha1" \
  --llm-model "gemini-2.0-flash-exp" \
  --max-attempts 5
```

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.securesourcemanager.v1.Repository" \
  --api-version "securesourcemanager.cnrm.cloud.google.com/v1alpha1" \
  --llm-model "gemini-2.0-flash-exp" \
  --max-attempts 5
```