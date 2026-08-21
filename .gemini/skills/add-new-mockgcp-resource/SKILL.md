# Skill: Add New MockGCP Resource

This skill provides a structured workflow for adding a new mock service to `mockgcp`.

## Overview

MockGCP uses `grpc-gateway` to provide an HTTP interface for mocked GCP services. The services are implemented in Go using GCP's published proto definitions.

## Workflow

### 1. Identify the Service and Protos

*   Identify the GCP service name (e.g., `memcache.googleapis.com`).
*   Locate the relevant `.proto` files in the `googleapis` repository.
*   Identify the corresponding Go packages in `cloud.google.com/go`.

### 2. Update MockGCP Makefile (Only if needed)

**Note: We should only do this if we cannot use `cloud.google.com/go` / `httptogrpc`. Generating protos manually should be avoided whenever possible.**

*   If the proto is not already being generated, add it to the `Makefile` in the `mockgcp` directory.
*   Run `make gen-proto` in the `mockgcp` directory to generate the Go code.
*   Alternatively, use `httptogrpc` if you want to avoid compiling protos and use existing Go SDK types.

### 3. Create the Mock Service Directory

Create a new directory `mockgcp/mock<servicename>`.

### 4. Implement `service.go`

Create `mockgcp/mock<servicename>/service.go`. This file should:
*   Register the service with `mockgcpregistry`.
*   Define the `MockService` struct.
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

	pb "cloud.google.com/go/<service>/apiv1/<service>pb"
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
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.New<Service>Client(conn))
    // Add LRO path if applicable
	grpcMux.AddOperationsPath("/v1/{prefix=projects/*/locations/*}/operations/{name=**}", conn)

	return grpcMux, nil
}
```

### 5. Implement Mock for the resource

Create `mockgcp/mock<servicename>/<resource>.go`. Implement the CRUD methods defined in the proto.
*   Use `s.storage` for persistence.
*   Use `s.operations` for long-running operations (LROs).
*   Parse resource names using a helper (e.g., `parseResourceName`).
*   Populate default values in a `populateDefaultsFor<Resource>` function.

Example implementations to look at for these patterns:
*   `mockfirestore` (Recent and shows the patterns nicely)
*   `mockmemorystore` (Recent and shows the patterns nicely)
*   `mockmemcache` (Good example of httptogrpc usage)

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

Add the new mock package to `mockgcp/register.go`.

### 8. Update Known Services

Add the service to `mockgcp/mockserviceusage/knownservices.go`.

### 9. Enable Tests

If the resource has KRM tests, add it to `config/tests/samples/create/harness.go`.

### 10. Verify and Align Logs

Please use the skill .gemini/skills/match-mockgcp-with-realgcp/SKILL.md to align the mockgcp behavior with the real gcp behavior.

## Tips

*   Use `httptogrpc` to avoid compiling protos if possible.
*   Follow established patterns in `mockfirestore` or other existing mocks.
*   For Terraform-based resources, check the provider code to understand how fields are flattened/expanded.
