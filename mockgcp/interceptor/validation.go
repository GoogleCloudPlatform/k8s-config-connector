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

package interceptor

import (
	"context"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/validation"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// getResourceFromRequest extracts a proto.Message from a request object.
// It looks for the first field that is a proto.Message.
// This is based on the convention that Create/Update requests have a field
// for the resource being modified (e.g., `Instance` in `CreateInstanceRequest`).
func getResourceFromRequest(req interface{}) proto.Message {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			if msg, ok := field.Interface().(proto.Message); ok {
				return msg
			}
		}
	}
	return nil
}

// LabelValidationInterceptor is a gRPC unary interceptor that validates labels on incoming requests.
func LabelValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Try to extract the resource from the request object
	resource := getResourceFromRequest(req)
	if resource == nil {
		// If we can't find a resource, just pass through
		return handler(ctx, req)
	}

	// Use reflection to check for a `GetLabels` method
	v := reflect.ValueOf(resource)
	m := v.MethodByName("GetLabels")
	if !m.IsValid() {
		// No GetLabels method, nothing to validate
		return handler(ctx, req)
	}

	// Call GetLabels()
	results := m.Call(nil)
	if len(results) != 1 {
		return handler(ctx, req)
	}
	labels, ok := results[0].Interface().(map[string]string)
	if !ok {
		return handler(ctx, req)
	}

	// Validate the labels
	if err := validation.ValidateLabels(labels); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
