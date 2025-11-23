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

// AddOperationsPath will serve the operations REST API at `path`.
// This is currently a GET handler; `path` should include a parameter {name}
// which is the operation's name, and an optional parameter {prefix} which
// includes whatever else the service defines as the per-operation prefix.
// The value of path / prefix should match whatever the API defines as its
// operations endpoint, which is often most conveniently determined by looking
// at the API documentation, or by seeing what paths clients request.
func (m *grpcMux) AddOperationsPath(httpPath string, conn *grpc.ClientConn) {
	client := longrunningpb.NewOperationsClient(conn)

	getMatcher, err := newPathMatcher(httpPath)
	if err != nil {
		klog.Fatalf("unable to parse path %q: %v", httpPath, err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) bool {
		ctx := r.Context()

		if r.Method != "GET" {
			return false
		}

		urlPath := r.URL.Path
		tokens := strings.Split(strings.TrimPrefix(urlPath, "/"), "/")

		vals, ok := getMatcher.Match(tokens)
		if !ok {
			return false
		}

		call := &httpMethodCall{
			parent: m,
			r:      r,
			w:      w,
		}

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
			var responseOptions ResponseOptions
			call.SendResponse(op, responseOptions)
		}
		return true
	}

	m.customHandlers = append(m.customHandlers, handler)
}
