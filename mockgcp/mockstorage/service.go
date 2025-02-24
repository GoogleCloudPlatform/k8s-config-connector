// Copyright 2023 Google LLC
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

package mockstorage

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked storage service.
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
	return []string{"storage.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBucketsServerServer(grpcServer, &buckets{MockService: s})
	pb.RegisterObjectsServerServer(grpcServer, &objects{MockService: s})
	pb.RegisterNotificationsServerServer(grpcServer, &notifications{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterBucketsServerHandler,
		pb.RegisterObjectsServerHandler,
		pb.RegisterNotificationsServerHandler,
	)
	if err != nil {
		return nil, err
	}

	// GCS has a different set of headers from most other APIs
	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {

		expires, found := httpmux.GetExpiresHeader(ctx)
		if found {
			response.Header().Set("Cache-Control", "private, max-age=0, must-revalidate, no-transform")
			response.Header().Set("Expires", expires)
		} else {
			response.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
			response.Header().Set("Pragma", "no-cache")
			response.Header().Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
		}

		response.Header().Set("Vary", "Origin")
		response.Header().Add("Vary", "X-Origin")

		response.Header().Set("Server", "UploadServer")

		response.Header().Del("X-Content-Type")
		response.Header().Del("X-Content-Type-Options")
		response.Header().Del("X-Frame-Options")
		response.Header().Del("X-Xss-Protection")

		// set http status code
		if code, found := httpmux.GetStatusCode(ctx); found {
			delete(response.Header(), "Grpc-Metadata-X-Http-Code")
			response.WriteHeader(code)
			if code == 204 {
				// GCS sends different headers on a 204
				response.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")

				response.Header().Set("Content-Type", "application/json")
			}
		}
	}

	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		if error.Code == http.StatusNotFound {
			if strings.HasPrefix(error.Message, "bucket") {
				error.Status = ""
				error.Message = "The specified bucket does not exist."
				error.Errors = []httpmux.ErrorResponseDetails{
					{
						Domain:  "global",
						Reason:  "notFound",
						Message: "The specified bucket does not exist.",
					},
				}
			}
			return
		}
	}

	return httpmux.FilterBodyOn204(mux)
}
