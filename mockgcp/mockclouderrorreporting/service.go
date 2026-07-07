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

package mockclouderrorreporting

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

type GCPNotificationSettings struct {
	Name                 string   `json:"name,omitempty"`
	NotificationChannels []string `json:"notificationChannels,omitempty"`
}

type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	mu       sync.Mutex
	settings map[string]*GCPNotificationSettings
}

func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	return &MockService{
		MockEnvironment: env,
		storage:         storage,
		settings:        make(map[string]*GCPNotificationSettings),
	}
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"clouderrorreporting.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	// REST-only mock, no GRPC registration needed
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	return s, nil
}

func (s *MockService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	tokens := strings.Split(strings.Trim(path, "/"), "/")
	if len(tokens) < 6 || tokens[0] != "v1beta1" || tokens[1] != "projects" || tokens[3] != "locations" || tokens[4] != "global" || tokens[5] != "notificationSettings" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	projectID := tokens[2]
	key := fmt.Sprintf("projects/%s/locations/global/notificationSettings", projectID)

	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case "GET":
		setting, ok := s.settings[key]
		if !ok {
			setting = &GCPNotificationSettings{
				Name:                 key,
				NotificationChannels: []string{},
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(setting)

	case "PATCH":
		var req GCPNotificationSettings
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.settings[key] = &GCPNotificationSettings{
			Name:                 key,
			NotificationChannels: req.NotificationChannels,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.settings[key])

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
