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

func (s *Operations) RegisterOperationsHandler(prefix string) func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
		forwardResponseOptions := mux.GetForwardResponseOptions()

		// GET /{prefix}/operations/{name}
		if err := mux.HandlePath("GET", "/"+prefix+"/operations/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			name := pathParams["name"]
			req := &longrunningpb.GetOperationRequest{Name: "operations/" + name}
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
