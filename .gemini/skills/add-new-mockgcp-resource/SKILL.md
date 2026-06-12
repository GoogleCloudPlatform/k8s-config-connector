# Skill: Add New MockGCP Resource

This skill provides a structured workflow and expert guidance for adding a new mock service to `mockgcp`.

## Overview

MockGCP uses `grpc-gateway` to provide an HTTP interface for mocked GCP services. The services are implemented in Go using GCP's published proto definitions.

---

## Workflow

### 1. Identify the Service and Protos

*   Identify the GCP service name (e.g., `gkehub.googleapis.com`, `memcache.googleapis.com`).
*   Determine if the service is standard (in public `googleapis` repository) or custom (needs custom `.proto` files inside `apis/mockgcp/cloud/...`).
*   Identify the corresponding Go packages in `cloud.google.com/go`.

### 2. Update MockGCP Makefile & Protos

*   **Standard Protos (googleapis)**:
    *   If using a standard GCP service, add its proto wildcard path (e.g., `./third_party/googleapis/mockgcp/cloud/gkehub/v1/*.proto \`) under the protobuf compilation target in `mockgcp/Makefile`.
    *   Ensure any custom proto files needed are placed in `./apis/mockgcp/cloud/<service>/<version>/*.proto`.
*   **Running Code Generation**:
    *   Run `make gen-proto` inside the `mockgcp/` directory to generate the gRPC and gateway Go files.
    *   If you experience compilation issues like `"TestIamPermissions" is already defined`, check if `apply-proto-patches.sh` is trying to append methods that are already in the tracked `.proto` files in the repository. If so, modify `apply-proto-patches.sh` to remove the redundant appends.
    *   If you encounter `"No such file or directory"` errors during `make gen-proto`, verify if the `mockgcp/Makefile` contains stale references to a third-party directory that was moved or deleted (e.g., resources moved from `third_party/` to `apis/`).

### 3. Create the Mock Service Directory

Create a new directory `mockgcp/mock<servicename>` (e.g., `mockgcp/mockgkehub`).

### 4. Implement `service.go`

Create `mockgcp/mock<servicename>/service.go`. This file should:
*   Register the service with `mockgcpregistry`.
*   Define the `MockService` / `GkeHubV1` structs.
*   Implement `ExpectedHosts()`, `Register()`, and `NewHTTPMux()`.

Example `service.go` structure:

```go
package mock<servicename>

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/<service>/v1"
)

func init() {
	mockgcpregistry.Register(New)
}

type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations
}

func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"<service>.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.Register<Service>Server(grpcServer, &<service>Server{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.Register<Service>Handler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
	)
	if err != nil {
		return nil, err
	}
	return mux, nil
}
```

### 5. Implement the Resource Controller

Create `mockgcp/mock<servicename>/<resource>.go`. Implement the CRUD methods defined in the proto.
*   Use `s.storage` for persistence.
*   Use `s.operations` for long-running operations (LROs).
*   Parse resource names using a helper (e.g., `parseResourceName` in `names.go`).
*   Ensure the parsing helper supports the exact format (with or without parent scopes/locations) expected by KCC direct controller's Get/Update/Delete/Create operations. For example, GKEHub Namespace uses a parent scope on creation but is retrieved/referenced without a scope in its main resource name.
*   If KCC direct controller sends updates to specific fields, verify that `UpdateMask` paths (e.g. `role`, `labels`) are fully handled in your update method's switch block.

### 6. Implement `normalize.go`

Create `mockgcp/mock<servicename>/normalize.go` to handle stable golden logs.
**CRITICAL**: Always scope `Previsit` to your service host.

```go
func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
    if !strings.Contains(event.URL(), "<service>.googleapis.com") {
        return
    }
    // ... normalization logic ...
}
```

### 7. Register the Service

*   Add the new mock package to `mockgcp/register.go` as a side-effect import (`_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mock<servicename>"`).
*   **CRITICAL**: Avoid manual appends in `mockgcp/mock_http_roundtrip.go` as the registry handles instantiation automatically. If the mock was listed in `mock_http_roundtrip.go`, remove it to avoid `"duplicate service registration"` errors.

### 8. Update Known Services

Add the service to `mockgcp/mockserviceusage/knownservices.go`.

### 9. Enable Tests

If the resource has KRM tests, add its `GroupKind` to `config/tests/samples/create/harness.go` under `MaybeSkip` to ensure it is not skipped.

### 10. Verify and Align Logs

1.  Run `hack/compare-mock fixtures/<testname>` to verify and capture mock behavior.
2.  If the test passes but golden files differ, running with `WRITE_GOLDEN_OUTPUT=1` (which `hack/compare-mock` does by default) will update the `_http.log`.
3.  Check git diffs of `_http.log` files to make sure they are correct and clean.

---

## Troubleshooting & Common Pitfalls

### Version Path Routing Mismatches
If a KCC direct controller is written against a `v1beta` endpoint but the MockGCP protobuf service is registered in `v1`, gRPC-gateway won't route `/v1beta/...` HTTP requests to your mocked RPCs.
*   **Solution**: In `mockgcp/apply-proto-patches.sh`, add `additional_bindings` to the proto HTTP gateway routing option to map the `v1beta` path directly to the `v1` RPC method.
    ```protobuf
    rpc GetScope(GetScopeRequest) returns (Scope) {
      option (google.api.http) = {
        get: "/v1/{name=projects/*/locations/*/scopes/*}"
        additional_bindings {
          get: "/v1beta/{name=projects/*/locations/*/scopes/*}"
        }
      };
    }
    ```

### Duplicate Service Registration
If running a mock test results in:
`grpc: Server.RegisterService found duplicate service registration for "mockgcp..."`
*   **Solution**: Check `mock_http_roundtrip.go` and remove any legacy manual import or explicit `services = append(services, mock<service>.New(...))` additions. Ensure the mock is only registered via `mockgcpregistry` in `mockgcp/register.go`.
