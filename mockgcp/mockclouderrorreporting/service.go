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
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked clouderrorreporting service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	notificationSettings map[string]*NotificationSettings
}

type NotificationSettings struct {
	Name                 string   `json:"name,omitempty"`
	NotificationChannels []string `json:"notificationChannels,omitempty"`
	VersionSkewReportDelays []string `json:"versionSkewReportDelays,omitempty"`
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment:      env,
		storage:              storage,
		notificationSettings: make(map[string]*NotificationSettings),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"clouderrorreporting.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	// No gRPC services to register as we use HTTP directly
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	return &httpHandler{s}, nil
}

type httpHandler struct {
	s *MockService
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// projects/{project}/locations/global/notificationSettings
	if strings.HasSuffix(r.URL.Path, "/notificationSettings") {
		name := strings.TrimPrefix(r.URL.Path, "/v1beta1/")
		switch r.Method {
		case "GET":
			h.getNotificationSettings(w, r, name)
		case "PATCH":
			h.updateNotificationSettings(w, r, name)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
