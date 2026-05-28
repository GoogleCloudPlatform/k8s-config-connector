# 2026-05-17 - composer migration
- Moved mockgcp composer to httptogrpc.
- The import path for composer is `cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb`. This is an example of a service with a deeper path in the official Go client.
- The `go_package` option in the protos often provides a hint for the correct import path.
