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

package mockdataform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	gw_v1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked dataform service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1beta1 *RepositoryV1Beta1
	v1      *DataformV1

	mutex           sync.Mutex
	queryPriorities map[string]any
}

type DataformV1 struct {
	*MockService
	gw_v1.UnimplementedDataformServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
		queryPriorities: make(map[string]any),
	}
	s.v1beta1 = &RepositoryV1Beta1{MockService: s}
	s.v1 = &DataformV1{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"dataform.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterDataformServer(grpcServer, s.v1beta1)
	gw_v1.RegisterDataformServer(grpcServer, s.v1)
}

type dataformResponseWrapper struct {
	http.ResponseWriter
	header     http.Header
	body       *bytes.Buffer
	statusCode int
}

func (w *dataformResponseWrapper) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *dataformResponseWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *dataformResponseWrapper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	v1beta1Mux := runtime.NewServeMux()
	if err := pb.RegisterDataformHandler(ctx, v1beta1Mux, conn); err != nil {
		return nil, err
	}

	v1Mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{}, gw_v1.RegisterDataformHandler)
	if err != nil {
		return nil, err
	}

	// Wrapper to handle queryPriority in WorkflowConfig invocationConfig.
	// queryPriority is accepted by the real GCP Dataform API in InvocationConfig,
	// but is not present in mockgcp's generated v1 proto schema for InvocationConfig.
	// 1. Stripping queryPriority from the request body before grpc-gateway parses it avoids
	//    400 Bad Request ("could not find field queryPriority") and prevents spurious diffs in updateMask.
	// 2. Re-injecting queryPriority into the JSON response ensures the KCC controller sees the
	//    configured queryPriority in observed GCP state, matching what the controller mapped into its desired state.
	v1Handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isWorkflowConfig := strings.Contains(r.URL.Path, "/workflowConfigs")
		if isWorkflowConfig {
			var reqQueryPriority any
			hasQueryPriority := false

			trimmedPath := strings.TrimPrefix(r.URL.Path, "/v1/")
			resourceName := trimmedPath
			if r.Method == http.MethodPost {
				workflowConfigID := r.URL.Query().Get("workflowConfigId")
				if workflowConfigID != "" {
					resourceName = trimmedPath + "/" + workflowConfigID
				}
			}

			if (r.Method == http.MethodPost || r.Method == http.MethodPatch) && r.Body != nil {
				bodyBytes, err := io.ReadAll(r.Body)
				if err == nil {
					r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
					var rawMap map[string]any
					if err := json.Unmarshal(bodyBytes, &rawMap); err == nil {
						if inv, ok := rawMap["invocationConfig"].(map[string]any); ok {
							if qp, found := inv["queryPriority"]; found {
								reqQueryPriority = qp
								hasQueryPriority = true
								delete(inv, "queryPriority")
							}
						}
						if strippedBytes, err := json.Marshal(rawMap); err == nil {
							r.Body = io.NopCloser(bytes.NewReader(strippedBytes))
							r.ContentLength = int64(len(strippedBytes))
						}
					}
				}
			}

			if hasQueryPriority {
				s.mutex.Lock()
				s.queryPriorities[resourceName] = reqQueryPriority
				s.mutex.Unlock()
			} else if r.Method == http.MethodDelete {
				s.mutex.Lock()
				delete(s.queryPriorities, resourceName)
				s.mutex.Unlock()
			}

			// Capture response to re-inject queryPriority if invocationConfig is in response
			rec := &dataformResponseWrapper{
				ResponseWriter: w,
				header:         w.Header(),
				body:           &bytes.Buffer{},
			}
			v1Mux.ServeHTTP(rec, r)

			respBytes := rec.body.Bytes()
			statusCode := rec.statusCode
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			if statusCode >= 200 && statusCode < 300 && len(respBytes) > 0 {
				var respMap map[string]any
				if err := json.Unmarshal(respBytes, &respMap); err == nil {
					name, _ := respMap["name"].(string)
					if name == "" {
						name = resourceName
					}
					s.mutex.Lock()
					qp, found := s.queryPriorities[name]
					s.mutex.Unlock()

					if inv, ok := respMap["invocationConfig"].(map[string]any); ok {
						if found {
							inv["queryPriority"] = qp
						} else if _, hasQP := inv["queryPriority"]; !hasQP {
							inv["queryPriority"] = float64(0)
						}
						if _, hasExecMode := inv["executionMode"]; !hasExecMode {
							inv["executionMode"] = float64(1)
						}
						if modifiedBytes, err := json.Marshal(respMap); err == nil {
							respBytes = modifiedBytes
						}
					}
				}
			}
			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(respBytes)))
			if rec.statusCode != 0 {
				w.WriteHeader(rec.statusCode)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			w.Write(respBytes)
			return
		}

		v1Mux.ServeHTTP(w, r)
	})

	mainMux := http.NewServeMux()
	mainMux.Handle("/v1beta1/", v1beta1Mux)
	mainMux.Handle("/v1/", v1Handler)

	return mainMux, nil
}
