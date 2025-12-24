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
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

// Mux is the primary interface for mapping HTTP requests to gRPC method calls.
type Mux interface {
	http.Handler
	AddService(client any, opt ...ServiceOption)
}

// grpcMux implements Mux.
type grpcMux struct {
	conn *grpc.ClientConn

	// customHandlers are tried in order, before matching against services
	customHandlers []func(w http.ResponseWriter, r *http.Request) bool

	// services are the gRPC services registered with the mux
	services []*grpcService
}

var _ Mux = (*grpcMux)(nil)

// NewGRPCMux creates a new grpcMux.
func NewGRPCMux(conn *grpc.ClientConn) (*grpcMux, error) {
	return &grpcMux{conn: conn}, nil
}

// grpcMethods holds state for a single gRPC method.
type grpcMethod struct {
	method       protoreflect.MethodDescriptor
	goMethod     reflect.Value
	goMethodType reflect.Method

	httpRule *annotations.HttpRule

	httpMethod string
	httpPath   string

	pathMatcher *pathMatcher
}

// Name returns the name of the gRPC method.
func (m *grpcMethod) Name() string {
	return string(m.method.FullName())
}

// ServeHTTP implements http.Handler.
// This is the primary entrypoint for HTTP requests,
// they are decoded and mapped to GRPC method calls.
func (m *grpcMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := klog.FromContext(ctx)

	url := r.URL.Path
	tokens := strings.Split(strings.TrimPrefix(url, "/"), "/")

	// Check for any custom handlers first
	for _, customHandler := range m.customHandlers {
		if handled := customHandler(w, r); handled {
			return
		}
	}

	// Otherwise look through all our registered services and methods
	for _, service := range m.services {
		for _, method := range service.methods {
			if method.httpMethod != r.Method {
				continue
			}

			matches, ok := method.pathMatcher.Match(tokens)
			if !ok {
				continue
			}

			m.serveHTTPMethod(w, r, method, matches)
			return
		}
	}

	log.Info("http request not matched", "method", r.Method, "url", r.URL.String())

	// For debugging it can be useful to dump all the possible paths
	dumpPaths := false
	if dumpPaths {
		for _, service := range m.services {
			for _, method := range service.methods {
				if method.httpMethod != r.Method {
					log.Info("skipping method due to http method mismatch", "method", method.Name(), "expected", method.httpMethod, "actual", r.Method)
					continue
				}
				matches, ok := method.pathMatcher.Match(tokens)
				if !ok {
					log.Info("skipping method due to path mismatch", "method", method.Name(), "path", method.httpPath, "tokens", tokens)
					continue
				}
				log.Info("found matching method", "method", method.Name(), "path", method.httpPath, "matches", matches)
			}
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

// serveHTTPMethod serves a single HTTP method mapped to a gRPC method.
func (m *grpcMux) serveHTTPMethod(w http.ResponseWriter, r *http.Request, method *grpcMethod, pathValues map[string]string) {
	ctx := r.Context()
	log := klog.FromContext(ctx)

	call := &httpMethodCall{
		parent: m,
		r:      r,
		w:      w,
	}

	var body []byte
	if r.Body != nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			klog.Errorf("failed to read body: %v", err)
			call.SendErrorResponse(status.Errorf(codes.InvalidArgument, "invalid body"))
			return
		}
		body = b
	} else {
		log.V(2).Info("no request body")
	}
	var responseOptions ResponseOptions

	// Build the input arguments for the gRPC method call
	var inArgs []reflect.Value
	for i := range method.goMethodType.Type.NumIn() {
		if i == 0 {
			// Skip receiver
			continue
		}
		inType := method.goMethodType.Type.In(i)

		prefix := ""
		if inType.Kind() == reflect.Ptr {
			inType = inType.Elem()
			prefix += "*"
		}
		if inType.Kind() == reflect.Slice {
			inType = inType.Elem()
			prefix += "[]"
		}

		inTypeName := prefix + inType.PkgPath() + "." + inType.Name()
		if inTypeName == "context.Context" {
			inArgs = append(inArgs, reflect.ValueOf(ctx))

		} else if inTypeName == "[]google.golang.org/grpc.CallOption" {
			// Do we need to pass any CallOptions?
			var callOptions []grpc.CallOption
			if method.goMethodType.Type.IsVariadic() {
				for _, callOption := range callOptions {
					inArgs = append(inArgs, reflect.ValueOf(callOption))
				}
			} else {
				inArgs = append(inArgs, reflect.ValueOf(callOptions))
			}
		} else {
			inArg := reflect.New(inType)

			protoMessage := inArg.Interface().(proto.Message)

			if len(body) != 0 {
				unmarshalOptions := protojson.UnmarshalOptions{
					DiscardUnknown: true,
				}

				dest := protoMessage

				if bodyFieldName := method.httpRule.GetBody(); bodyFieldName != "" && bodyFieldName != "*" {
					bodyFieldFD := protoMessage.ProtoReflect().Descriptor().Fields().ByTextName(bodyFieldName)
					if bodyFieldFD == nil {
						klog.Fatalf("body field %q not found in %v", bodyFieldName, protoMessage.ProtoReflect().Descriptor().FullName())
					}
					dest = protoMessage.ProtoReflect().Mutable(bodyFieldFD).Message().Interface()
				}

				if err := unmarshalOptions.Unmarshal(body, dest); err != nil {
					klog.Errorf("failed to unmarshal body: %v", err)
					call.SendErrorResponse(status.Errorf(codes.InvalidArgument, "invalid body"))
					return
				}
			}

			for k, v := range pathValues {
				if err := setProtoField(protoMessage, k, []string{v}); err != nil {
					klog.Errorf("failed to set field %q: %v", k, err)
					call.SendErrorResponse(status.Errorf(codes.InvalidArgument, "invalid value for %q", k))
					return
				}
			}

			q := r.URL.Query()
			for k, values := range q {
				if k == "alt" || k == "$alt" {
					responseOptions.Alt = values
					continue
				}
				// Convert camelCase to snake_case
				var protoKey []rune
				for _, c := range k {
					if c >= 'A' && c <= 'Z' {
						protoKey = append(protoKey, '_')
						protoKey = append(protoKey, c-'A'+'a')
					} else {
						protoKey = append(protoKey, c)
					}
				}
				k = string(protoKey)

				if err := setProtoField(protoMessage, k, values); err != nil {
					klog.Errorf("failed to set url-query field %q: %v", k, err)
					call.SendErrorResponse(status.Errorf(codes.InvalidArgument, "invalid value for query-param %q", k))
					return
				}
			}

			inArgs = append(inArgs, inArg)
		}
	}

	// Make the gRPC method call
	out := method.goMethod.Call(inArgs)

	if len(out) != 2 {
		klog.Fatalf("output format not handled, expected two output parameters")
	}

	// Check if the gRPC method returned an error
	if !out[1].IsNil() {
		err, ok := out[1].Interface().(error)
		if !ok {
			klog.Fatalf("expected second parameter to be error, was %T", out[1])
		}
		call.SendErrorResponse(err)
		return
	}

	// Otherwise send a normal response with the JSON-encoded proto response
	response, ok := out[0].Interface().(proto.Message)
	if !ok {
		klog.Fatalf("expected first parameter to be proto.Message, was %T", out[0])
	}
	call.SendResponse(response, responseOptions)
}

// setProtoField sets the specified proto field; this is normally used for request parameters
func setProtoField(protoMessage protoreflect.ProtoMessage, k string, values []string) error {
	tokens := strings.Split(k, ".")

	curr := protoMessage.ProtoReflect()
	for i := 0; i < len(tokens)-1; i++ {
		token := tokens[i]
		fd := protoMessage.ProtoReflect().Descriptor().Fields().ByTextName(token)
		if fd == nil {
			return fmt.Errorf("value field %q not found", k)
		}
		curr = curr.Mutable(fd).Message()
	}

	{
		tail := tokens[len(tokens)-1]
		fd := curr.Descriptor().Fields().ByTextName(tail)
		if fd == nil {
			return fmt.Errorf("value field %q not found in %v", k, curr.Descriptor().FullName())
		}
		switch fd.Kind() {
		case protoreflect.StringKind:
			if len(values) != 1 {
				return fmt.Errorf("expected one value for %q, got %v", k, values)
			}
			v := values[0]
			curr.Set(fd, protoreflect.ValueOfString(v))

		case protoreflect.BoolKind:
			if len(values) != 1 {
				return fmt.Errorf("expected one value for %q, got %v", k, values)
			}
			v := values[0]
			switch v {
			case "true", "True":
				curr.Set(fd, protoreflect.ValueOfBool(true))
			case "false", "False":
				curr.Set(fd, protoreflect.ValueOfBool(false))
			default:
				return fmt.Errorf("expected bool value for %q, got %v", k, v)
			}

		case protoreflect.EnumKind:
			if len(values) != 1 {
				return fmt.Errorf("expected one value for %q, got %v", k, values)
			}
			v := values[0]
			enumValueDescriptor := fd.Enum().Values().ByName(protoreflect.Name(v))
			if enumValueDescriptor == nil {
				return fmt.Errorf("invalid enum value %q for %q", v, k)
			}
			curr.Set(fd, protoreflect.ValueOfEnum(enumValueDescriptor.Number()))

		case protoreflect.MessageKind:
			messageFQN := fd.Message().FullName()
			switch messageFQN {
			case "google.protobuf.FieldMask":
				fieldMask := curr.Mutable(fd).Message()

				pathsFD := fieldMask.Descriptor().Fields().ByName("paths")
				paths := fieldMask.Mutable(pathsFD).List()
				for _, v := range values {
					for _, v2 := range strings.Split(v, ",") {
						paths.Append(protoreflect.ValueOfString(v2))
					}
				}

			default:
				return fmt.Errorf("unhandled message kind %v", messageFQN)
			}

		default:
			return fmt.Errorf("unhandled field kind %v", fd.Kind())
		}
	}
	return nil
}
