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

package mockbigquery

import (
	"bytes"
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
)

// MockService represents a mocked bigquery service.
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
	return "bigquery.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDatasetsServer(grpcServer, &datasetsServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, pb.RegisterDatasetsHandler)
	if err != nil {
		return nil, err
	}

	// To be compatible with the real BigQuery API, we need to serve a 204 on DELETE
	// Add a filter to do that, it's hard to do with grpc-gateway
	filter := func(w http.ResponseWriter, r *http.Request) {
		brw := &bufferedResponseWriter{
			header: make(http.Header),
		}
		mux.ServeHTTP(brw, r)
		// Send a 204 on DELETE instead of a 200 with an empty response.
		if brw.statusCode == 200 && r.Method == http.MethodDelete && brw.body.String() == "{}" {
			brw.statusCode = 204
			brw.body.Reset()
			brw.header.Set("Content-Length", "0")
			brw.header.Del("Cache-Control")
		}
		brw.WriteTo(w)
	}

	return http.HandlerFunc(filter), nil
}

// bufferedResponseWriter implements http.ResponseWriter and stores the response.
type bufferedResponseWriter struct {
	statusCode int
	body       bytes.Buffer
	header     http.Header
}

var _ http.ResponseWriter = &bufferedResponseWriter{}

// Header implements http.ResponseWriter
func (w *bufferedResponseWriter) Header() http.Header {
	return w.header
}

// Write implements http.ResponseWriter
func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.body.Write(b)
}

// WriteHeader implements http.ResponseWriter
func (w *bufferedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// WriteTo writes the buffered response to a different http.ResponseWriter.
func (w *bufferedResponseWriter) WriteTo(out http.ResponseWriter) {
	for k, values := range w.header {
		out.Header()[k] = values
	}
	statusCode := w.statusCode
	if statusCode == 0 {
		statusCode = 200
	}
	out.WriteHeader(statusCode)
	out.Write(w.body.Bytes())
}
