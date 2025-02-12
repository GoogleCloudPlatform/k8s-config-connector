# Generate Fuzzer

This command generates and qualifies a fuzzer for a given resource.

## Usage

```bash
go run main.go generate-fuzzer \
  --message "google.cloud.managedkafka.v1.Topic" \
  --api-version "managedkafka.cnrm.cloud.google.com/v1alpha1" \
  --max-attempts 3
```
