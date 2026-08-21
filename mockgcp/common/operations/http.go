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

package operations

import (
	"context"
	"net/http"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

// RegisterOperationsPath will serve the operations REST API at `path`.
// This is currently a GET handler; `path` should include a parameter {name}
// which is the operation's name, and an optional parameter {prefix} which
// includes whatever else the service defines as the per-operation prefix.
// The value of path / prefix should match whatever the API defines as its
// operations endpoint, which is often most conveniently determined by looking
// at the API documentation, or by seeing what paths clients request.
func (s *Operations) RegisterOperationsPath(path string) func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
		forwardResponseOptions := mux.GetForwardResponseOptions()

		if err := mux.HandlePath("GET", path, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			name := pathParams["name"]
			prefix := pathParams["prefix"]
			req := &longrunningpb.GetOperationRequest{Name: "operations/" + name}
			if prefix != "" {
				req.Name = prefix + "/operations/" + name
			}
			op, err := s.GetOperation(ctx, req)
			if err != nil {
				if status.Code(err) == codes.NotFound {
					klog.Infof("operation not found %+v", req)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				klog.Warningf("error getting operation %T: %v", err, err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			for _, forwardResponseOption := range forwardResponseOptions {
				err := forwardResponseOption(ctx, w, op)
				if err != nil {
					klog.Warningf("error running forwardResponseOption %T: %v", forwardResponseOption, err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			b, err := httpmux.MarshalAsJSON(op)
			if err != nil {
				klog.Warningf("error converting to proto %T: %v", err, err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
		}); err != nil {
			return err
		}
		return nil
	}
}
