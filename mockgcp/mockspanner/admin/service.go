// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockspanner

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	databasepb_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/database/v1"
	instancepb_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/instance/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	databaseV1 *SpannerDatabaseV1
	instanceV1 *SpannerInstanceV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.databaseV1 = &SpannerDatabaseV1{MockService: s}
	s.instanceV1 = &SpannerInstanceV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"spanner.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	databasepb_v1.RegisterDatabaseAdminServer(grpcServer, s.databaseV1)
	instancepb_v1.RegisterInstanceAdminServer(grpcServer, s.instanceV1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		databasepb_v1.RegisterDatabaseAdminHandler,
		instancepb_v1.RegisterInstanceAdminHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"))
	if err != nil {
		return nil, err
	}

	// Returns  non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == 404 {
			error.Errors = nil
		}
	}
	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("X-Content-Type-Options")
	}

	wrapped := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		brw := &bufferedResponseWriter{
			header: make(http.Header),
		}
		mux.ServeHTTP(brw, r)

		if brw.statusCode == 0 {
			brw.statusCode = 200
		}

		if brw.statusCode >= 200 && brw.statusCode < 300 && strings.Contains(brw.header.Get("Content-Type"), "application/json") {
			var decoded any
			if err := json.Unmarshal(brw.body.Bytes(), &decoded); err == nil {
				walkJSON(decoded, false)
				if reencoded, err := json.Marshal(decoded); err == nil {
					brw.body.Reset()
					brw.body.Write(reencoded)
				}
			}
		}

		brw.WriteTo(w)
	})

	return wrapped, nil
}

type bufferedResponseWriter struct {
	statusCode int
	body       bytes.Buffer
	header     http.Header
}

var _ http.ResponseWriter = &bufferedResponseWriter{}

func (w *bufferedResponseWriter) Header() http.Header {
	return w.header
}

func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.body.Write(b)
}

func (w *bufferedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

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

func walkJSON(v any, inMetadata bool) {
	switch val := v.(type) {
	case map[string]any:
		nextInMetadata := inMetadata
		for k, subVal := range val {
			childInMetadata := nextInMetadata
			if k == "metadata" {
				childInMetadata = true
			} else if k == "response" {
				childInMetadata = false
			}

			if k == "name" {
				if name, ok := subVal.(string); ok && strings.Contains(name, "/instances/") && !strings.Contains(name, "/operations/") && !strings.Contains(name, "/databases/") {
					if config, ok := val["config"].(string); ok {
						tokens := strings.Split(config, "/")
						if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instanceConfigs" {
							location := strings.TrimPrefix(tokens[3], "regional-")
							val["resourceLocation"] = location
						}
					}
				}
			}

			if k == "replicaSelection" {
				if replicaSelection, ok := subVal.(map[string]any); ok && childInMetadata {
					replicaSelection["replicaType"] = 1
				}
			}

			walkJSON(subVal, childInMetadata)
		}
	case []any:
		for _, subVal := range val {
			walkJSON(subVal, inMetadata)
		}
	}
}
