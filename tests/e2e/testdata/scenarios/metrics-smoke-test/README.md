# Metrics Smoke Test

This is a simple scenario test that verifies the metrics infrastructure is working correctly across different resource types.

## What it does

1. **Creates a GCP Project** - This triggers API calls to the Resource Manager API
2. **Creates a Logging Log Metric** - This triggers API calls to the Logging API (direct-based)
3. **Creates a Compute Network** - This triggers API calls to the Compute API (Terraform-based)
4. **Creates a Monitoring Alert Policy** - This triggers API calls to the Monitoring API (DCL-based)
5. **Waits 30 seconds** - This allows the system to run and collect HTTP logs
6. **Checks Transport Metrics** - Verifies that the transport layer is emitting the expected metrics


## Expected results

The HTTP logs should show:
- Successful API calls to GCP APIs
- Proper HTTP status codes (200, 201, etc.)
- Request/response bodies with GCP resource data

This verifies that:
1. The metrics transport is working
2. API calls are being made successfully
3. The HTTP logging infrastructure is capturing requests
4. The transport layer is properly instrumented with metrics
