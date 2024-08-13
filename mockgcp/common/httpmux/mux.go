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

type Options struct {
	// If EmitUnpopulated is true, we will send empty proto fields (false / "" / 0 etc)
	// Some older APIs do this (e.g. cloudbilling)
	// While it likely doesn't matter, it makes golden testing easier to match.
	EmitUnpopulated bool
}

type ServeMux struct {
	ServeMux *runtime.ServeMux

	// RewriteError allows us to customize the error we return.
	// Error can be changed in-place.
	RewriteError func(ctx context.Context, error *ErrorResponse)

	// RewriteHeaders allows us to customize the headers we return.
	// Response is changed in-place.
	RewriteHeaders func(ctx context.Context, response http.ResponseWriter, payload proto.Message)
}

// NewServeMux constructs an http server with our error handling etc
func NewServeMux(ctx context.Context, conn *grpc.ClientConn, opt Options, handlers ...func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) (*ServeMux, error) {
	resolver := &protoResolver{}
	marshaler := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: opt.EmitUnpopulated,
				Resolver:        resolver,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
				Resolver:       resolver,
			},
		},
	}

	marshalerWithEnumNumbers := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: opt.EmitUnpopulated,
				UseEnumNumbers:  true,
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
		case MetadataKeyStatusCode:
			return "", false
		case MetadataKeyExpires:
			return "", false
		default:
			klog.Warningf("unknown grpc metadata header %q", key)
			return "", false
		}
	}

	m := &ServeMux{}

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(m.customErrorHandler),
		runtime.WithMarshalerOption("application/json;enum-encoding=int", marshalerWithEnumNumbers),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler),
		runtime.WithOutgoingHeaderMatcher(outgoingHeaderMatcher),
		runtime.WithForwardResponseOption(m.addGCPHeaders),
	)
	m.ServeMux = mux

	for _, handler := range handlers {
		if err := handler(ctx, mux, conn); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *ServeMux) addGCPHeaders(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if w.Header().Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
	w.Header().Set("Cache-Control", "private")
	w.Header().Set("Server", "ESF")
	w.Header()["Vary"] = []string{"Origin", "X-Origin", "Referer"}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Xss-Protection", "0")

	if m.RewriteHeaders != nil {
		// response is changed in place
		m.RewriteHeaders(ctx, w, resp)
	}

	return nil
}

func (m *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := m.ServeMux.ServeHTTP

	for k, values := range r.URL.Query() {
		if k == "$alt" {
			for _, v := range values {
				if v == "json;enum-encoding=int" {
					klog.Infof("found %q=%q, will convert to Accept header", k, v)
					r.Header.Set("Accept", "application/json;enum-encoding=int")
				}
			}
		}
	}
	handler(w, r)
}
