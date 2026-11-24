// Copyright 2025 Google LLC
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

package mockbigquerybiglake

// +tool:mockgcp-service
// http.host: biglake.googleapis.com
// proto.service: google.cloud.bigquery.biglake.v1.MetastoreService
// proto.service: google.cloud.biglake.v1.IcebergCatalogService
// proto.service: google.cloud.biglake.hive.v1.HiveMetastoreService

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	hivepb "cloud.google.com/go/biglake/hive/apiv1/hivepb"
	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	icebergpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/biglake/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/biglake/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked biglake service.
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
	return []string{"biglake.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterMetastoreServiceServer(grpcServer, &bigLakeService{MockService: s})
	biglakepb.RegisterIcebergCatalogServiceServer(grpcServer, &icebergCatalogService{MockService: s})
	hivepb.RegisterHiveMetastoreServiceServer(grpcServer, &hiveMetastoreService{MockService: s})
}

type icebergResponseWriter struct {
	http.ResponseWriter
	buf        bytes.Buffer
	statusCode int
}

func (w *icebergResponseWriter) Write(b []byte) (int, error) {
	return w.buf.Write(b)
}

func (w *icebergResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterMetastoreServiceHandler,
		icebergpb.RegisterIcebergCatalogServiceHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 && strings.Contains(error.Message, "Namespace") {
			error.Errors = nil
			error.Status = ""
			error.Type = "NoSuchNamespaceException"
		}
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPatch && strings.HasSuffix(r.URL.Path, "/properties") {
			rw := &icebergResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			mux.ServeHTTP(rw, r)
			if rw.statusCode == http.StatusOK {
				var data map[string]any
				if err := json.Unmarshal(rw.buf.Bytes(), &data); err == nil {
					if _, ok := data["missing"]; !ok {
						data["missing"] = []string{}
					}
					if _, ok := data["removed"]; !ok {
						data["removed"] = []string{}
					}
					if modified, err := json.Marshal(data); err == nil {
						w.WriteHeader(rw.statusCode)
						w.Write(modified)
						return
					}
				}
			}
			w.WriteHeader(rw.statusCode)
			w.Write(rw.buf.Bytes())
			return
		}
		mux.ServeHTTP(w, r)
	})

	return handler, nil
}
