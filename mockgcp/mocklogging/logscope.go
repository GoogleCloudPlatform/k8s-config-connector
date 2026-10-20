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

package mocklogging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type LogScopeServer struct {
	*MockService
	pb.UnimplementedConfigServiceV2Server
}

type logScopeJSON struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	ResourceNames []string `json:"resourceNames,omitempty"`
	CreateTime    string   `json:"createTime,omitempty"`
	UpdateTime    string   `json:"updateTime,omitempty"`
}

func (s *MockService) buildLogScopeHandler() func(w http.ResponseWriter, r *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		path := r.URL.Path
		// Expected URL structure: /v2/{parent=projects/*/locations/*}/logScopes[/{id}]
		// e.g. /v2/projects/mock-project/locations/global/logScopes
		// e.g. /v2/projects/mock-project/locations/global/logScopes/my-scope
		idx := strings.Index(path, "/logScopes")
		if idx == -1 {
			return false
		}

		parent := strings.TrimPrefix(path[:idx], "/v2/")
		rest := path[idx+len("/logScopes"):]
		logScopeID := strings.TrimPrefix(rest, "/")

		// Ensure parent follows projects/*/locations/* or organizations/*/locations/* or folders/*/locations/*
		if !strings.HasPrefix(parent, "projects/") && !strings.HasPrefix(parent, "organizations/") && !strings.HasPrefix(parent, "folders/") {
			return false
		}

		ctx := r.Context()
		switch r.Method {
		case http.MethodGet:
			if logScopeID == "" {
				// ListLogScopes
				return false
			}
			name := fmt.Sprintf("%s/logScopes/%s", parent, logScopeID)
			// To store nicely in s.storage, let's look up by FQN
			fqn := name
			existing := &pb.LogBucket{}
			if err := s.storage.Get(ctx, fqn, existing); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": map[string]any{
						"code":    http.StatusNotFound,
						"message": fmt.Sprintf("LogScope %s not found.", name),
						"status":  "NOT_FOUND",
					},
				})
				return true
			}
			// Unpack stored json
			var scope logScopeJSON
			if existing.Description != "" {
				_ = json.Unmarshal([]byte(existing.Description), &scope)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(scope)
			return true

		case http.MethodPost:
			if logScopeID != "" {
				return false
			}
			id := r.URL.Query().Get("logScopeId")
			if id == "" {
				w.WriteHeader(http.StatusBadRequest)
				return true
			}
			name := fmt.Sprintf("%s/logScopes/%s", parent, id)

			var input logScopeJSON
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return true
			}

			// Validate project/location exists
			projectObj, err := s.Projects.GetProjectByID(strings.Split(parent, "/")[1])
			if err != nil {
				_ = projectObj
			}

			now := time.Now().UTC().Format(time.RFC3339Nano)
			input.Name = name
			input.CreateTime = now
			input.UpdateTime = now

			// Store serialized scope
			bytes, _ := json.Marshal(input)
			stored := &pb.LogBucket{
				Name:        name,
				Description: string(bytes),
				CreateTime:  timestamppb.Now(),
			}
			if err := s.storage.Create(ctx, name, stored); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(map[string]any{
					"error": map[string]any{
						"code":    http.StatusConflict,
						"message": fmt.Sprintf("LogScope %s already exists.", name),
						"status":  "ALREADY_EXISTS",
					},
				})
				return true
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(input)
			return true

		case http.MethodPatch:
			name := fmt.Sprintf("%s/logScopes/%s", parent, logScopeID)
			existing := &pb.LogBucket{}
			if err := s.storage.Get(ctx, name, existing); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": map[string]any{
						"code":    http.StatusNotFound,
						"message": fmt.Sprintf("LogScope %s not found.", name),
						"status":  "NOT_FOUND",
					},
				})
				return true
			}

			var scope logScopeJSON
			if existing.Description != "" {
				_ = json.Unmarshal([]byte(existing.Description), &scope)
			}

			var patch logScopeJSON
			if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return true
			}

			updateMask := r.URL.Query().Get("updateMask")
			if updateMask == "" {
				scope.Description = patch.Description
				scope.ResourceNames = patch.ResourceNames
			} else {
				for _, mask := range strings.Split(updateMask, ",") {
					switch strings.TrimSpace(mask) {
					case "description":
						scope.Description = patch.Description
					case "resourceNames":
						scope.ResourceNames = patch.ResourceNames
					}
				}
			}
			scope.UpdateTime = time.Now().UTC().Format(time.RFC3339Nano)

			bytes, _ := json.Marshal(scope)
			stored := proto.Clone(existing).(*pb.LogBucket)
			stored.Description = string(bytes)
			stored.UpdateTime = timestamppb.Now()

			if err := s.storage.Update(ctx, name, stored); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return true
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(scope)
			return true

		case http.MethodDelete:
			name := fmt.Sprintf("%s/logScopes/%s", parent, logScopeID)
			existing := &pb.LogBucket{}
			if err := s.storage.Delete(ctx, name, existing); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]any{
					"error": map[string]any{
						"code":    http.StatusNotFound,
						"message": fmt.Sprintf("LogScope %s not found.", name),
						"status":  "NOT_FOUND",
					},
				})
				return true
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{}"))
			return true
		}

		return false
	}
}

func (s *MockService) parseLogScopeName(name string) (*projects.ProjectData, string, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "logScopes" {
		projectName := tokens[1]
		logScopeID := tokens[5]
		project, err := s.Projects.GetProjectByID(projectName)
		if err != nil {
			return nil, "", err
		}
		return project, logScopeID, nil
	}
	return nil, "", fmt.Errorf("name %q is not valid", name)
}
