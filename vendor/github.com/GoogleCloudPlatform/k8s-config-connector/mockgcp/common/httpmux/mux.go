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

package httpmux

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// NewServeMux constructs an http server with our error handling etc
func NewServeMux(ctx context.Context, conn *grpc.ClientConn, handlers ...func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) (*runtime.ServeMux, error) {
	resolver := &protoResolver{}
	marshaler := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: false,
				Resolver:        resolver,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
				Resolver:       resolver,
			},
		},
	}

	outgoingHeaderMatcher := func(key string) (string, bool) {
		switch key {
		case "content-type":
			return "", false
		default:
			klog.Warningf("unknown grpc metadata header %q", key)
			return "", false
		}
	}

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(customErrorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler),
		runtime.WithOutgoingHeaderMatcher(outgoingHeaderMatcher),
		runtime.WithForwardResponseOption(addGCPHeaders),
	)

	for _, handler := range handlers {
		if err := handler(ctx, mux, conn); err != nil {
			return nil, err
		}
	}

	return mux, nil
}

func addGCPHeaders(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if w.Header().Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
	w.Header().Set("Cache-Control", "private")
	w.Header().Set("Server", "ESF")
	w.Header()["Vary"] = []string{"Origin", "X-Origin", "Referer"}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Xss-Protection", "0")

	return nil
}
