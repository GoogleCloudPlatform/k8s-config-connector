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
	"encoding/json"
	"net/http"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/klog/v2"
)

// httpMethodCall holds state for a single method call over HTTP.
type httpMethodCall struct {
	parent     *grpcMux
	grpcMethod *grpcMethod
	r          *http.Request
	w          http.ResponseWriter
}

// SendErrorResponse sends an error response for a GRPC method call over HTTP.
func (c *httpMethodCall) SendErrorResponse(err error) {
	ctx := c.r.Context()

	klog.Warningf("sending error response for %T %+v", err, err)

	// Try to map well known errors to HTTP status codes and structured responses
	statusErr, ok := status.FromError(err)
	if ok {
		response := statusErr.Proto()

		httpErrorResponse := &httpErrorResponse{
			Error: &httpError{
				Code:    http.StatusInternalServerError,
				Message: response.Message,
			},
		}

		switch statusErr.Code() {
		case codes.InvalidArgument:
			httpErrorResponse.Error.Code = http.StatusBadRequest
			httpErrorResponse.Error.Status = "INVALID_ARGUMENT"
		case codes.NotFound:
			httpErrorResponse.Error.Code = http.StatusNotFound
			httpErrorResponse.Error.Status = "NOT_FOUND"
		case codes.Internal:
			httpErrorResponse.Error.Code = http.StatusInternalServerError
			httpErrorResponse.Error.Status = "INTERNAL"
		}

		body, err := json.Marshal(httpErrorResponse)
		if err != nil {
			klog.Errorf("failed to marshal error: %v", err)
			http.Error(c.w, "internal error", http.StatusInternalServerError)
			return
		}

		c.w.Header().Set("Content-Type", "application/json")

		c.parent.addGCPHeaders(ctx, c.w, response)

		c.w.WriteHeader(httpErrorResponse.Error.Code)
		if _, err := c.w.Write(body); err != nil {
			klog.Errorf("failed to write error: %v", err)
		}
		klog.Infof("sent response %v with body %v", httpErrorResponse.Error.Code, string(body))
		return
	}
	klog.Warningf("stub-handling error %v", err)
	http.Error(c.w, err.Error(), http.StatusInternalServerError)
}

// ResponseOptions holds options for sending a response.
type ResponseOptions struct {
	Alt []string
}

// populateMarshalOptions populates marshal options based on the options set.
func (o *ResponseOptions) populateMarshalOptions(marshalOptions *protojson.MarshalOptions) {
	if o == nil {
		return
	}
	for _, alt := range o.Alt {
		switch alt {
		case "json;enum-encoding=int":
			marshalOptions.UseEnumNumbers = true
		case "json":
			// Default behavior
		default:
			klog.Fatalf("unhandled alt option: %q", alt)
		}
	}
}

// SendResponse sends a successful response for a GRPC method call over HTTP.
func (c *httpMethodCall) SendResponse(response proto.Message, responseOptions ResponseOptions) {
	ctx := c.r.Context()

	httpCode := http.StatusOK

	c.w.Header().Set("Content-Type", "application/json")

	c.parent.addGCPHeaders(ctx, c.w, response)

	marshalOptions := protojson.MarshalOptions{
		Resolver: &protoResolver{},
	}
	responseOptions.populateMarshalOptions(&marshalOptions)

	emitUnpopulated := false
	if c.grpcMethod != nil {
		if c.grpcMethod.parentService.options.EmitUnpopulated {
			emitUnpopulated = true
		}
	} else if c.parent != nil {
		if op, ok := response.(*longrunningpb.Operation); ok {
			var target string
			if op.Metadata != nil {
				if m, err := anypb.UnmarshalNew(op.Metadata, proto.UnmarshalOptions{Resolver: protoregistry.GlobalTypes}); err == nil {
					ref := m.ProtoReflect()
					fd := ref.Descriptor().Fields().ByName("target")
					if fd != nil {
						target = ref.Get(fd).String()
					}
				}
			}
			if target == "" && op.GetResponse() != nil {
				if m, err := anypb.UnmarshalNew(op.GetResponse(), proto.UnmarshalOptions{Resolver: protoregistry.GlobalTypes}); err == nil {
					ref := m.ProtoReflect()
					fd := ref.Descriptor().Fields().ByName("name")
					if fd != nil {
						target = ref.Get(fd).String()
					}
				}
			}

			if target != "" {
				targetTokens := tokenizePathString(target)
				for _, s := range c.parent.services {
					for _, m := range s.methods {
						if _, ok := m.pathMatcher.Match(targetTokens); ok {
							if s.options.EmitUnpopulated {
								emitUnpopulated = true
							}
							break
						}
					}
					if emitUnpopulated {
						break
					}
				}
			}
		}
	}

	if emitUnpopulated {
		marshalOptions.EmitUnpopulated = true
	}

	body, err := marshalOptions.Marshal(response)
	if err != nil {
		klog.Errorf("failed to marshal response: %v", err)
		http.Error(c.w, "internal error", http.StatusInternalServerError)
		return
	}

	c.w.WriteHeader(httpCode)
	if _, err := c.w.Write(body); err != nil {
		klog.Errorf("failed to write error: %v", err)
	}
	klog.Infof("sent response %v with body %v", httpCode, string(body))
}

func tokenizePathString(p string) []string {
	if idx := strings.Index(p, "://"); idx != -1 {
		p = p[idx+3:]
		if slashIdx := strings.Index(p, "/"); slashIdx != -1 {
			p = p[slashIdx+1:]
		}
	}
	suffix := ""
	if idx := strings.Index(p, ":"); idx != -1 {
		suffix = p[idx:]
		p = p[:idx]
	}
	tokens := strings.Split(strings.TrimPrefix(p, "/"), "/")
	if suffix != "" {
		tokens = append(tokens, suffix)
	}
	return tokens
}
