// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcompute

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

// MockService represents a mocked compute service.
type MockService struct {
	kube    client.Client
	storage storage.Storage

	projects   projects.ProjectStore
	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		kube:       env.GetKubeClient(),
		storage:    storage,
		projects:   env.GetProjects(),
		operations: operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "compute.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterAddressesServer(grpcServer, &RegionalAddressesV1{MockService: s})
	pb.RegisterGlobalAddressesServer(grpcServer, &GlobalAddressesV1{MockService: s})

	pb.RegisterDisksServer(grpcServer, &DisksV1{MockService: s})
	pb.RegisterRegionDisksServer(grpcServer, &RegionalDisksV1{MockService: s})

	pb.RegisterHealthChecksServer(grpcServer, &GlobalHealthCheckV1{MockService: s})
	pb.RegisterRegionHealthChecksServer(grpcServer, &RegionalHealthCheckV1{MockService: s})

	pb.RegisterNetworksServer(grpcServer, &NetworksV1{MockService: s})

	pb.RegisterSubnetworksServer(grpcServer, &SubnetsV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	// TODO: Is all of beta a direct mapping to v1?
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		u := r.URL
		u.Path = "/compute/v1/" + strings.TrimPrefix(u.Path, "/compute/beta/")
		r.URL = u

		mux.ServeHTTP(w, r)
	}

	if err := mux.HandlePath("GET", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("POST", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("DELETE", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PATCH", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}
	if err := mux.HandlePath("PUT", "/compute/beta/{path=**}", rewriteBetaToV1); err != nil {
		return nil, err
	}

	if err := pb.RegisterAddressesHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalAddressesHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterDisksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionDisksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterHealthChecksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionHealthChecksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterSubnetworksHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
