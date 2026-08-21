# MockGCP Journal: ComputeHTTPHealthCheck & ComputeHTTPSHealthCheck

## Context
Compute Engine has legacy HTTP and HTTPS Health Checks endpoints (`/projects/{project}/global/httpHealthChecks` and `/projects/{project}/global/httpsHealthChecks`). These are separate resource types from the modern, unified `/projects/{project}/global/healthChecks` endpoint.
However, in the GCP OpenAPI / gRPC proto schemas, only `HealthChecks` and its sub-types are defined in modern APIs, while legacy health check APIs might be omitted or marked as deprecated.

## Design Decision: The Facade Mapping Pattern
Instead of implementing full gRPC server methods and registering new services from scratch (which would require custom/legacy proto generation and extensive boilerplates), we implemented an **HTTP Middleware Facade Mapping** pattern inside `NewHTTPMux`:

1. **Request Interception & Rewrite**:
   - Detect incoming requests matching `/global/httpHealthChecks` or `/global/httpsHealthChecks`.
   - Read the payload body and map its top-level legacy HTTP/HTTPS fields (`host`, `port`, `requestPath`) into the modern nested `httpHealthCheck` or `httpsHealthCheck` format.
   - Insert `type: "HTTP"` or `type: "HTTPS"` appropriately.
   - Rewrite the request path to `/global/healthChecks`.

2. **Response Capture & Translation**:
   - Wrap the HTTP ResponseWriter to capture the response JSON.
   - Perform string replacement to translate `healthChecks/` back to the legacy `httpHealthChecks/` or `httpsHealthChecks/` resource URLs in selfLinks and Operations.
   - For `compute#healthCheck` kinds, restore the nested fields back to the top level, set the kind to `compute#httpHealthCheck` or `compute#httpsHealthCheck`, and delete the modern `type` field.

## Benefits
- No complex proto generation or registration needed.
- Leverages the robust existing `healthChecks` implementation.
- Extremely clean, simple, and completely transparent to the client (Terraform/KCC).
