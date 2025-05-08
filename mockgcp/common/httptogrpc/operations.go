// Copyright 2025 Google LLC
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

package httptogrpc

import (
	"net/http"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

// RegisterOperationsPath will serve the operations REST API at `path`.
// This is currently a GET handler; `path` should include a parameter {name}
// which is the operation's name, and an optional parameter {prefix} which
// includes whatever else the service defines as the per-operation prefix.
// The value of path / prefix should match whatever the API defines as its
// operations endpoint, which is often most conveniently determined by looking
// at the API documentation, or by seeing what paths clients request.
func (m *grpcMux) RegisterWithHTTPMux(httpPath string, conn *grpc.ClientConn, inner http.Handler) http.Handler {
	client := longrunningpb.NewOperationsClient(conn)

	getMatcher, err := newPathMatcher(httpPath)
	if err != nil {
		klog.Fatalf("unable to parse path %q: %v", httpPath, err)
	}

	// klog.Infof("client is %T", client)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		urlPath := r.URL.Path
		tokens := strings.Split(strings.TrimPrefix(urlPath, "/"), "/")

		call := &httpMethodCall{
			parent: m,
			r:      r,
			w:      w,
		}

		if r.Method == "GET" {
			if vals, ok := getMatcher.Match(tokens); ok {
				klog.Infof("WILDCARD match for %v; matches get: %v", tokens, vals)
				name := vals["name"]
				prefix := vals["prefix"]
				req := &longrunningpb.GetOperationRequest{Name: "operations/" + name}
				if prefix != "" {
					req.Name = prefix + "/operations/" + name
				}

				op, err := client.GetOperation(ctx, req)
				if err != nil {
					call.SendErrorResponse(err)
				} else {
					call.SendResponse(op)
				}
				return
			}
		}

		inner.ServeHTTP(w, r)
	})
}
