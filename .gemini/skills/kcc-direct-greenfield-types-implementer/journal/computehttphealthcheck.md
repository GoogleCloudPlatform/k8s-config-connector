# ComputeHTTPHealthCheck Implementation Journal

## Observations & Learnings

- **Legacy Resource Proto Exclusion**:
  - The legacy global `HttpHealthCheck` GCP resource is REST-only and completely excluded from the compiled GCP compute proto files. In `googleapis/google/cloud/compute/v1/BUILD.bazel`, the messages `HttpHealthCheck`, `HttpHealthCheckList`, `GetHttpHealthCheckRequest`, etc., are listed in the `_MESSAGE_IGNORE_LIST` and are omitted during Discovery-to-Proto compilation.
  - This means there is no canonical top-level message `HttpHealthCheck` in `google.cloud.compute.v1`'s protobuf.

- **Sub-Message Name Matching**:
  - The proto compiler defines a nested message `HTTPHealthCheck` (with uppercase `HTTP`) inside the newer style `HealthCheck` proto.
  - Running `generate-types` with `--resource ComputeHTTPHealthCheck:HTTPHealthCheck` successfully resolves to this nested message.
  - However, the nested message is structurally incomplete relative to the legacy `/global/httpHealthChecks` resource (it is missing `checkIntervalSec`, `healthyThreshold`, `timeoutSec`, `unhealthyThreshold`, etc.).

- **Manual Schema Hand-writing**:
  - To support the exact KRM schema expected for `ComputeHTTPHealthCheck` without breaking compatibility, we hand-wrote the full `ComputeHTTPHealthCheckSpec` in `apis/compute/v1beta1/httphealthcheck_types.go`, aligning with the existing Terraform-based schema.
  - We also included `ProjectRef` (matching modern direct resource conventions) and structured output-only fields like `creationTimestamp` and `selfLink` within `status.observedState`.
  - Because `HttpHealthCheck` is a global compute resource, we omitted the `Location` field.

- **Stability Level**:
  - Since `ComputeHTTPHealthCheck` is already a stable `v1beta1` resource, we set `cnrm.cloud.google.com/stability-level=stable`.
