// Copyright 2024 Google LLC
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

package mockredis

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/redis/apiv1/redispb"
	pbcluster "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked redis service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"redis.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterCloudRedisServer(grpcServer, &redisServer{MockService: s})
	pbcluster.RegisterCloudRedisClusterServer(grpcServer, &clusterServer{MockService: s})
}

type responseWrapper struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (w *responseWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *responseWrapper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, err
	}

	grpcMux.AddService(pb.NewCloudRedisClient(conn))
	grpcMux.AddService(pbcluster.NewCloudRedisClusterClient(conn))

	grpcMux.AddOperationsPath("/v1beta1/{prefix=**}/operations/{name}", conn)
	grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	// Terraform uses the /v1beta1/ endpoints, but we have protos only for v1.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		isV1 := false
		u := r.URL
		if strings.HasPrefix(u.Path, "/v1beta1/projects/") && strings.Contains(u.Path, "/instances") {
			isV1 = true
			u2 := *u
			u2.Path = "/v1" + strings.TrimPrefix(u.Path, "/v1beta1")
			r = httpmux.RewriteRequest(r, &u2)
		} else if strings.HasPrefix(u.Path, "/v1beta1/projects/") && strings.Contains(u.Path, "/operations") {
			isV1 = true
		}

		if isV1 {
			rec := &responseWrapper{ResponseWriter: w, body: &bytes.Buffer{}}
			grpcMux.ServeHTTP(rec, r)

			respBytes := rec.body.Bytes()
			respBytes = bytes.ReplaceAll(respBytes, []byte(`google.cloud.redis.v1.Instance`), []byte(`google.cloud.redis.v1beta1.Instance`))

			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(respBytes)))
			if rec.statusCode != 0 {
				w.WriteHeader(rec.statusCode)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			w.Write(respBytes)
		} else {
			grpcMux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}
