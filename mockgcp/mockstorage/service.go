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
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	// Note we use "real" protos (not mockgcp) ones as it's GRPC API.
	grpcpb "cloud.google.com/go/storage/control/apiv2/controlpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked storage service.
type MockService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations

	mutex      sync.Mutex
	objectData map[string][]byte
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
		objectData:      make(map[string][]byte),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	return []string{"storage.googleapis.com", "www.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterBucketsServerServer(grpcServer, &buckets{MockService: s})
	pb.RegisterObjectsServerServer(grpcServer, &objects{MockService: s})
	pb.RegisterFoldersServerServer(grpcServer, &folder{MockService: s})
	pb.RegisterNotificationsServerServer(grpcServer, &notifications{MockService: s})
	pb.RegisterManagedFoldersServerServer(grpcServer, &managedFolders{MockService: s})
	grpcpb.RegisterStorageControlServer(grpcServer, &StorageControlService{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{
		UnescapingMode: runtime.UnescapingModeAllExceptReserved,
	},
		pb.RegisterBucketsServerHandler,
		pb.RegisterObjectsServerHandler,
		pb.RegisterNotificationsServerHandler,
		pb.RegisterFoldersServerHandler,
		pb.RegisterManagedFoldersServerHandler,
		s.operations.RegisterOperationsPath("/v1/{prefix=**}/operations/{name}"),
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

	filterBodyHandler, err := httpmux.FilterBodyOn204(mux)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/upload/") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/upload")
			if r.Method == "POST" && strings.HasSuffix(r.URL.Path, "/o") {
				// This is an upload.
				bucket := strings.TrimPrefix(r.URL.Path, "/storage/v1/b/")
				bucket = strings.TrimSuffix(bucket, "/o")

				name := r.URL.Query().Get("name")

				// Read the body
				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				var data []byte
				contentType := r.Header.Get("Content-Type")
				if strings.HasPrefix(contentType, "multipart/related") {
					_, params, err := mime.ParseMediaType(contentType)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					boundary := params["boundary"]
					mr := multipart.NewReader(bytes.NewReader(bodyBytes), boundary)
					for {
						part, err := mr.NextPart()
						if err == io.EOF {
							break
						}
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						partData, err := io.ReadAll(part)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						// The first part is metadata (JSON), the second part is media.
						// We want the media.
						if part.Header.Get("Content-Type") != "application/json" {
							data = partData
						}
					}
				} else {
					data = bodyBytes
				}

				objectsClient := pb.NewObjectsServerClient(conn)
				req := &pb.InsertObjectRequest{
					Bucket: &bucket,
					Name:   &name,
				}
				obj, err := objectsClient.InsertObject(r.Context(), req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Store the data
				s.mutex.Lock()
				s.objectData[fmt.Sprintf("b/%s/o/%s", bucket, name)] = data
				s.mutex.Unlock()

				// Set some headers to match what ESF/GCS returns
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				b, _ := protojson.Marshal(obj)
				w.Write(b)
				return
			}
		}

		if r.Method == "GET" && r.URL.Query().Get("alt") == "media" {
			// This is a media download.
			bucketAndObject := strings.TrimPrefix(r.URL.Path, "/storage/v1/b/")
			parts := strings.SplitN(bucketAndObject, "/o/", 2)
			if len(parts) == 2 {
				bucket := parts[0]
				name := parts[1]

				s.mutex.Lock()
				data, found := s.objectData[fmt.Sprintf("b/%s/o/%s", bucket, name)]
				s.mutex.Unlock()

				if found {
					w.Header().Set("Content-Type", "application/octet-stream")
					w.WriteHeader(http.StatusOK)
					w.Write(data)
					return
				}
			}
		}

		// Handle simple media download: storage.googleapis.com/{bucket}/{name}
		if r.Method == "GET" && !strings.HasPrefix(r.URL.Path, "/storage/v1/") && !strings.HasPrefix(r.URL.Path, "/upload/") && !strings.HasPrefix(r.URL.Path, "/batch") {
			path := strings.TrimPrefix(r.URL.Path, "/")
			parts := strings.SplitN(path, "/", 2)
			if len(parts) == 2 {
				bucket := parts[0]
				name := parts[1]

				s.mutex.Lock()
				data, found := s.objectData[fmt.Sprintf("b/%s/o/%s", bucket, name)]
				s.mutex.Unlock()

				if found {
					w.Header().Set("Content-Type", "application/octet-stream")
					w.WriteHeader(http.StatusOK)
					w.Write(data)
					return
				}
			}
		}

		filterBodyHandler.ServeHTTP(w, r)
	}), nil
}
