// Copyright 2022 Google LLC
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
// http.host: billingbudgets.googleapis.com
// proto.service: google.cloud.billing.budgets.v1.BudgetService

package mockbillingbudgets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httptogrpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked billingbudgets service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	budgets *BudgetServiceServer
}

type BudgetServiceServer struct {
	*MockService
	pb.UnimplementedBudgetServiceServer
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	s.budgets = &BudgetServiceServer{MockService: s}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"billingbudgets.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBudgetServiceServer(grpcServer, s.budgets)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	grpcMux, err := httptogrpc.NewGRPCMux(conn)
	if err != nil {
		return nil, fmt.Errorf("error building grpc service: %w", err)
	}

	grpcMux.AddService(pb.NewBudgetServiceClient(conn))

	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isV1Beta1 := strings.Contains(r.URL.Path, "/v1beta1/")
		if isV1Beta1 {
			if r.Body != nil && (r.Method == "POST" || r.Method == "PATCH" || r.Method == "PUT") {
				bodyBytes, err := io.ReadAll(r.Body)
				if err == nil {
					// v1beta1 uses allUpdatesRule, v1 uses notificationsRule
					bodyBytes = bytes.ReplaceAll(bodyBytes, []byte("\"allUpdatesRule\""), []byte("\"notificationsRule\""))
					bodyBytes = bytes.ReplaceAll(bodyBytes, []byte("\"all_updates_rule\""), []byte("\"notifications_rule\""))

					var raw map[string]json.RawMessage
					if json.Unmarshal(bodyBytes, &raw) == nil {
						if wrapped, ok := raw["budget"]; ok {
							bodyBytes = []byte(wrapped)
						}
					}

					r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
					r.ContentLength = int64(len(bodyBytes))
				}
			}
			r.URL.Path = strings.Replace(r.URL.Path, "/v1beta1/", "/v1/", 1)
			w = &v1beta1ResponseWriter{ResponseWriter: w}
		}
		grpcMux.ServeHTTP(w, r)
	})

	return mux, nil
}

type v1beta1ResponseWriter struct {
	http.ResponseWriter
}

func (w *v1beta1ResponseWriter) Write(b []byte) (int, error) {
	// v1 uses notificationsRule, v1beta1 uses allUpdatesRule
	b = bytes.ReplaceAll(b, []byte("\"notificationsRule\""), []byte("\"allUpdatesRule\""))
	b = bytes.ReplaceAll(b, []byte("\"notifications_rule\""), []byte("\"all_updates_rule\""))
	return w.ResponseWriter.Write(b)
}
