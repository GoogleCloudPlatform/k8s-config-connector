---
name: move-away-from-grpc-gateway
description: Moves a mockgcp service from using locally generated grpc-gateway proto bindings to using reflection with the official Google Cloud Go client library (httptogrpc). Use this when tasked with stopping grpc-gateway generation for a GCP service in mockgcp.
---

# Move away from grpc-gateway

This skill outlines the steps to transition a mockgcp service from using `grpc-gateway` to using reflection via `httptogrpc`.

## Step 0: Verify gRPC client library existence

Before migrating, verify that an official gRPC Go client and protobuf module exists for the service. Some services (e.g., `servicenetworking`) only have REST-only client libraries (`google.golang.org/api`) and do not have official protobuf packages under `cloud.google.com/go/` (or `google.golang.org/genproto/googleapis/api/`).

Check if the client package can be imported:
```go
import pb "cloud.google.com/go/<service_name>/apiv1/<service_name>pb"
```
Or check if it exists in the standard Google Cloud Go module repository. If no such package exists, the migration is **blocked** and cannot be completed. In this case, please document this in a new journal file under `mockgcp/.gemini/skills/move-away-from-grpc-gateway/journal/<service_name>.md` and do not perform the migration.

## Step 1: Stop generating protos

In `mockgcp/Makefile`, remove the service from the `gen-proto-no-fixup` target. Find the line that looks like `./third_party/googleapis/google/cloud/<service_name>/v1/*.proto \` (or `./third_party/googleapis/mockgcp/cloud/...`) and delete it.

## Step 2: Delete generated files

Delete the generated code for the service located in:
`mockgcp/generated/google/cloud/<service_name>/` (or `mockgcp/generated/mockgcp/cloud/<service_name>/`)
Run `rm -rf` on the directory to remove the `.pb.go`, `.pb.gw.go`, and `_grpc.pb.go` files for all versions.

## Step 3: Update imports

Update the Go files in `mockgcp/mock<service_name>/` (typically `service.go`, `instance.go`, etc.):

- Remove the local generated import: `pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/<service_name>/<version>"`
- Add the official client library import: `pb "cloud.google.com/go/<service_name>/apiv1/<service_name>pb"` (or `apiv1beta` if the service is only in beta).
- Register the service with `mockgcpregistry` by adding an `init` function and updating the `New` function signature in `service.go`:
  ```go
  func init() {
      mockgcpregistry.Register(New)
  }
  func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
      // ...
  }
  ```
- Update `mockgcp/register.go` to include the new service package in the anonymous imports.
- Remove the manual registration of the service in `mockgcp/mock_http_roundtrip.go`.

## Step 4: Switch HTTP Multiplexer to httptogrpc

In `mockgcp/mock<service_name>/service.go`, update the `NewHTTPMux` method.

- Replace `"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"` with `"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"`
- Remove the `httpmux.NewServeMux` call and replace it with `httptogrpc.NewGRPCMux(conn)`.
- If the old code used `mux.RewriteError`, you should safely delete it. `httptogrpc` does not support it (and handles errors differently).
- If the old code used `mux.RewriteHeaders`, use `mux.OverrideHeaders(func(response http.ResponseWriter) { ... })`.

Example:

```go
func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	mux.AddService(pb.NewMemorystoreClient(conn)) // Replace MemorystoreClient with the correct client
	mux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	// Custom header handling
	mux.OverrideHeaders(func(response http.ResponseWriter) {
		response.Header().Del("Cache-Control")
	})

	return mux, nil
}
```

## Step 5: Fix type mismatches

The official client library proto types might differ slightly from the old grpc-gateway types:

- You may need to replace pointer field assignments with `new(bool)` or specific constant types.
- Replace type-casted `proto.Clone(x).(*pb.Message)` with `proto.CloneOf(x)`, which is available in recent versions of `google.golang.org/protobuf/proto`.
- Replace legacy protobuf types with their modern equivalents:
    - `github.com/golang/protobuf/ptypes/empty` -> `google.golang.org/protobuf/types/known/emptypb`
    - `github.com/golang/protobuf/ptypes/timestamp` -> `google.golang.org/protobuf/types/known/timestamppb`
- Check and fix compilation errors by running `go build` or `go test` in the service directory.
- Update `uuid` generation or default field behaviors to match the strict types in the official client.

- Mocks often handle FieldMask paths manually with a switch statement. The official client library via `httptogrpc` may send camelCase paths (matching JSON names) instead of snake_case (matching proto names). You should update your `switch` statements to support both formats.

## Journaling

If you discover any new patterns, edge cases, or workarounds during migration, you MUST document them in the `mockgcp/.gemini/skills/move-away-from-grpc-gateway/journal/` directory.

- **CRITICAL RULE:** Always create a new file with a descriptive, topic-based name (e.g., `netapp_leftover_generated_files.md` or `datastream_rewriteerror_not_needed.md`) to capture your findings, or **append** to an existing relevant file.
- **NEVER** overwrite, truncate, or delete any existing journal files or entries. All records must be preserved to prevent losing valuable historical migration context.
