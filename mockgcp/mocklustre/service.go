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
// http.host: lustre.googleapis.com
// proto.service: google.cloud.lustre.v1.Lustre

package mocklustre

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/lustre/apiv1/lustrepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked lustre service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	*lustreServer

	accessRulesMu sync.Mutex
	accessRules   map[string]map[string]interface{}
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
		accessRules:     make(map[string]map[string]interface{}),
	}
	s.lustreServer = &lustreServer{MockService: s}
	return s
}

func (s *MockService) setAccessRules(name string, rules map[string]interface{}) {
	s.accessRulesMu.Lock()
	defer s.accessRulesMu.Unlock()
	if s.accessRules == nil {
		s.accessRules = make(map[string]map[string]interface{})
	}
	s.accessRules[name] = rules
}

func (s *MockService) getAccessRules(name string) map[string]interface{} {
	s.accessRulesMu.Lock()
	defer s.accessRulesMu.Unlock()
	if s.accessRules == nil {
		return nil
	}
	return s.accessRules[name]
}

func (s *MockService) deleteAccessRules(name string) {
	s.accessRulesMu.Lock()
	defer s.accessRulesMu.Unlock()
	if s.accessRules != nil {
		delete(s.accessRules, name)
	}
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"lustre.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterLustreServer(grpcServer, s.lustreServer)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewLustreClient(conn))

	grpcMux.AddOperationsPath("/v1/{prefix=**}/operations/{name}", conn)

	return &lustreHTTPHandler{mux: grpcMux, s: s}, nil
}

type lustreHTTPHandler struct {
	mux http.Handler
	s   *MockService
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (r *responseRecorder) Header() http.Header {
	return r.ResponseWriter.Header()
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func (h *lustreHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasPrefix(path, "/v1/projects/") && strings.Contains(path, "/instances/") {
		instanceName := strings.TrimPrefix(path, "/v1/")
		if !strings.Contains(instanceName, "/operations/") && strings.Count(instanceName, "/") == 5 {
			switch r.Method {
			case http.MethodGet:
				rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
				h.mux.ServeHTTP(rec, r)
				if rec.statusCode == http.StatusOK {
					rules := h.s.getAccessRules(instanceName)
					if rules != nil {
						var data map[string]interface{}
						if err := json.Unmarshal(rec.body.Bytes(), &data); err == nil {
							data["accessRulesOptions"] = rules
							if modified, err := json.Marshal(data); err == nil {
								w.Header().Set("Content-Type", "application/json")
								w.Header().Set("Content-Length", strconv.Itoa(len(modified)))
								w.WriteHeader(rec.statusCode)
								w.Write(modified)
								return
							}
						}
					}
				}
				w.WriteHeader(rec.statusCode)
				w.Write(rec.body.Bytes())
				return

			case http.MethodPatch:
				bodyBytes, _ := io.ReadAll(r.Body)
				var bodyMap map[string]interface{}
				_ = json.Unmarshal(bodyBytes, &bodyMap)

				if rules, ok := bodyMap["accessRulesOptions"].(map[string]interface{}); ok {
					h.s.setAccessRules(instanceName, rules)
				}

				q := r.URL.Query()
				mask := q.Get("updateMask")
				if mask == "accessRulesOptions" {
					now := time.Now()
					idx := strings.LastIndex(instanceName, "/instances/")
					prefix := instanceName[:idx]
					metadata := &pb.OperationMetadata{
						ApiVersion: "v1",
						CreateTime: timestamppb.New(now),
						Target:     instanceName,
						Verb:       "update",
					}
					op, err := h.s.operations.StartLRO(r.Context(), prefix, metadata, func() (proto.Message, error) {
						metadata.EndTime = timestamppb.Now()
						obj := &pb.Instance{}
						_ = h.s.storage.Get(r.Context(), instanceName, obj)
						obj.UpdateTime = timestamppb.New(now)
						_ = h.s.storage.Update(r.Context(), instanceName, obj)
						return obj, nil
					})
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					opJSON, _ := protojson.Marshal(op)
					w.WriteHeader(http.StatusOK)
					w.Write(opJSON)
					return
				}

				if strings.Contains(mask, "accessRulesOptions") {
					parts := strings.Split(mask, ",")
					var newParts []string
					for _, p := range parts {
						if p != "accessRulesOptions" && !strings.HasPrefix(p, "accessRulesOptions.") {
							newParts = append(newParts, p)
						}
					}
					q.Set("updateMask", strings.Join(newParts, ","))
					r.URL.RawQuery = q.Encode()
				}
				r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

			case http.MethodDelete:
				h.s.deleteAccessRules(instanceName)
			}
		}
	} else if strings.HasPrefix(path, "/v1/projects/") && strings.HasSuffix(path, "/instances") && r.Method == http.MethodPost {
		bodyBytes, _ := io.ReadAll(r.Body)
		var bodyMap map[string]interface{}
		_ = json.Unmarshal(bodyBytes, &bodyMap)

		instanceId := r.URL.Query().Get("instanceId")
		if instanceId != "" {
			if rules, ok := bodyMap["accessRulesOptions"].(map[string]interface{}); ok {
				instanceName := fmt.Sprintf("%s/%s", strings.TrimPrefix(path, "/v1/"), instanceId)
				h.s.setAccessRules(instanceName, rules)
			}
		}
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	h.mux.ServeHTTP(w, r)
}
