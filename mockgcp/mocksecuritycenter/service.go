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
// http.host: securitycenter.googleapis.com
// proto.service: google.cloud.securitycenter.v1.SecurityCenter

package mocksecuritycenter

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	grpcpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/securitycenter/v1"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked securitycenter service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

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
	return []string{"securitycenter.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterSecurityCenterServer(grpcServer, &SecurityCenterServer{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		grpcpb.RegisterSecurityCenterHandler,
		s.registerLocationBigQueryExportHandlers,
	)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func (s *MockService) registerLocationBigQueryExportHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := pb.NewSecurityCenterClient(conn)

	patternsCreate := []runtime.Pattern{
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 4, 4, 5, 3, 2, 4}, []string{"v1", "organizations", "locations", "parent", "bigQueryExports"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 4, 4, 5, 3, 2, 4}, []string{"v1", "folders", "locations", "parent", "bigQueryExports"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 4, 4, 5, 3, 2, 4}, []string{"v1", "projects", "locations", "parent", "bigQueryExports"}, "")),
	}
	for _, pattern := range patternsCreate {
		mux.Handle("POST", pattern, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
			var protoReq pb.CreateBigQueryExportRequest
			protoReq.BigQueryExport = &pb.BigQueryExport{}
			if err := inboundMarshaler.NewDecoder(req.Body).Decode(protoReq.BigQueryExport); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			protoReq.Parent = pathParams["parent"]
			protoReq.BigQueryExportId = req.URL.Query().Get("bigQueryExportId")

			resp, err := client.CreateBigQueryExport(ctx, &protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, req, resp)
		})
	}

	patternsGet := []runtime.Pattern{
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "organizations", "locations", "bigQueryExports", "name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "folders", "locations", "bigQueryExports", "name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "projects", "locations", "bigQueryExports", "name"}, "")),
	}
	for _, pattern := range patternsGet {
		mux.Handle("GET", pattern, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
			protoReq := &pb.GetBigQueryExportRequest{
				Name: pathParams["name"],
			}

			resp, err := client.GetBigQueryExport(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, req, resp)
		})
	}

	patternsDelete := []runtime.Pattern{
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "organizations", "locations", "bigQueryExports", "name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "folders", "locations", "bigQueryExports", "name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "projects", "locations", "bigQueryExports", "name"}, "")),
	}
	for _, pattern := range patternsDelete {
		mux.Handle("DELETE", pattern, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
			protoReq := &pb.DeleteBigQueryExportRequest{
				Name: pathParams["name"],
			}

			resp, err := client.DeleteBigQueryExport(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, req, resp)
		})
	}

	patternsUpdate := []runtime.Pattern{
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "organizations", "locations", "bigQueryExports", "big_query_export.name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "folders", "locations", "bigQueryExports", "big_query_export.name"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 1, 0, 4, 6, 5, 4}, []string{"v1", "projects", "locations", "bigQueryExports", "big_query_export.name"}, "")),
	}
	for _, pattern := range patternsUpdate {
		mux.Handle("PATCH", pattern, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
			var protoReq pb.UpdateBigQueryExportRequest
			protoReq.BigQueryExport = &pb.BigQueryExport{}
			if err := inboundMarshaler.NewDecoder(req.Body).Decode(protoReq.BigQueryExport); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			protoReq.BigQueryExport.Name = pathParams["big_query_export.name"]
			if updateMask := req.URL.Query().Get("updateMask"); updateMask != "" {
				protoReq.UpdateMask = &fieldmaskpb.FieldMask{Paths: strings.Split(updateMask, ",")}
			}

			resp, err := client.UpdateBigQueryExport(ctx, &protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, req, resp)
		})
	}

	patternsList := []runtime.Pattern{
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 4, 5, 5, 4}, []string{"v1", "organizations", "locations", "bigQueryExports", "parent"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 4, 5, 5, 4}, []string{"v1", "folders", "locations", "bigQueryExports", "parent"}, "")),
		runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 2, 2, 1, 0, 2, 3, 4, 5, 5, 4}, []string{"v1", "projects", "locations", "bigQueryExports", "parent"}, "")),
	}
	for _, pattern := range patternsList {
		mux.Handle("GET", pattern, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
			protoReq := &pb.ListBigQueryExportsRequest{
				Parent: pathParams["parent"],
			}

			resp, err := client.ListBigQueryExports(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
				return
			}
			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, req, resp)
		})
	}

	return nil
}
