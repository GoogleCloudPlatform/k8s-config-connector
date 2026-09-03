// Copyright 2026 Google LLC
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

// +tool:mockgcp-service
// http.host: saasservicemgmt.googleapis.com
// proto.service: google.cloud.saasplatform.saasservicemgmt.v1beta1.SaasDeployments

package mocksaasservicemgmt

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pbhttp "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/saasplatform/saasservicemgmt/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked saasservicemgmt service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	saasDeploymentsServer *SaasDeploymentsServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	s.saasDeploymentsServer = &SaasDeploymentsServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{
		"saasservicemgmt.googleapis.com",
		"saasservicemgmt.us-central1.googleapis.com",
	}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterSaasDeploymentsServer(grpcServer, s.saasDeploymentsServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pbhttp.RegisterSaasDeploymentsHandler)
	if err != nil {
		return nil, err
	}

	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	return mux, nil
}
