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

package mockcontainer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked container service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"container.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterClusterManagerServer(grpcServer, &ClusterManagerV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterClusterManagerHandler)
	if err != nil {
		return nil, err
	}

	// Terraform uses the /v1beta1/ endpoints, but gcloud uses v1.
	// Rewrite for now (hoping they are compatible enough)
	rewriteV1ToBeta := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		if strings.HasPrefix(u.Path, "/v1/") {
			u2 := *u
			u2.Path = "/v1beta1/" + strings.TrimPrefix(u.Path, "/v1/")
			r = httpmux.RewriteRequest(r, &u2)
		}

		// Intercept Request Body: OS_2022 -> OS_VERSION_LTSC2022
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				// Replace short enum names with full proto enum names
				bodyBytes = bytes.ReplaceAll(bodyBytes, []byte(`"OS_2022"`), []byte(`"OS_VERSION_LTSC2022"`))
				bodyBytes = bytes.ReplaceAll(bodyBytes, []byte(`"OS_2019"`), []byte(`"OS_VERSION_LTSC2019"`))
				r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
				r.ContentLength = int64(len(bodyBytes))
			}
		}

		// Intercept Response Body: OS_VERSION_LTSC2022 -> OS_2022
		rec := &responseWrapper{ResponseWriter: w, body: &bytes.Buffer{}}
		mux.ServeHTTP(rec, r)

		respBytes := rec.body.Bytes()
		respBytes = bytes.ReplaceAll(respBytes, []byte(`"OS_VERSION_LTSC2022"`), []byte(`"OS_2022"`))
		respBytes = bytes.ReplaceAll(respBytes, []byte(`"OS_VERSION_LTSC2019"`), []byte(`"OS_2019"`))

		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(respBytes)))
		if rec.statusCode != 0 {
			w.WriteHeader(rec.statusCode)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Write(respBytes)
	}

	return http.HandlerFunc(rewriteV1ToBeta), nil
}

type responseWrapper struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (w *responseWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *responseWrapper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}
