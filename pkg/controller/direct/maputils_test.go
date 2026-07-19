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

package direct

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestStringDuration_FromProto(t *testing.T) {
	mapctx := &MapContext{}
	d := &durationpb.Duration{Seconds: 34312, Nanos: 20}
	krm := StringDuration_FromProto(mapctx, d)
	if *krm != "9h31m52.00000002s" {
		t.Fatalf("google.protobuf.Duration -> string, expect \"9h31m52.00000002s\", got %s", *krm)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> string error: %s", mapctx.Err())
	}
}

func TestStringDuration_ToProto(t *testing.T) {
	mapctx := &MapContext{}
	s := "1h1m"
	d := StringDuration_ToProto(mapctx, &s)
	if d.Seconds != 3660 || d.Nanos != 0 {
		t.Fatalf("string -> google.protobuf.Duration, expect \"seconds:3660 nanos:00\", got %s", d)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> String error: %s", mapctx.Err())
	}
}

func TestIsAlreadyExists(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "unrelated error",
			err:  fmt.Errorf("something went wrong"),
			want: false,
		},
		{
			name: "gRPC AlreadyExists",
			err:  status.Error(codes.AlreadyExists, "resource already exists"),
			want: true,
		},
		{
			name: "gRPC NotFound",
			err:  status.Error(codes.NotFound, "not found"),
			want: false,
		},
		{
			name: "gRPC AlreadyExists wrapped with apierror",
			err: func() error {
				grpcErr := status.Error(codes.AlreadyExists, "resource already exists")
				apiErr, _ := apierror.ParseError(grpcErr, true)
				return apiErr
			}(),
			want: true,
		},
		{
			name: "gRPC NotFound wrapped with apierror",
			err: func() error {
				grpcErr := status.Error(codes.NotFound, "not found")
				apiErr, _ := apierror.ParseError(grpcErr, true)
				return apiErr
			}(),
			want: false,
		},
		{
			name: "wrapped gRPC AlreadyExists",
			err:  fmt.Errorf("creating resource: %w", status.Error(codes.AlreadyExists, "already exists")),
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsAlreadyExists(tc.err)
			if got != tc.want {
				t.Errorf("IsAlreadyExists(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "unrelated error",
			err:  fmt.Errorf("something went wrong"),
			want: false,
		},
		{
			name: "gRPC NotFound",
			err:  status.Error(codes.NotFound, "not found"),
			want: true,
		},
		{
			name: "gRPC AlreadyExists",
			err:  status.Error(codes.AlreadyExists, "already exists"),
			want: false,
		},
		{
			name: "gRPC NotFound wrapped with apierror",
			err: func() error {
				grpcErr := status.Error(codes.NotFound, "not found")
				apiErr, _ := apierror.ParseError(grpcErr, true)
				return apiErr
			}(),
			want: true,
		},
		{
			name: "gRPC AlreadyExists wrapped with apierror",
			err: func() error {
				grpcErr := status.Error(codes.AlreadyExists, "already exists")
				apiErr, _ := apierror.ParseError(grpcErr, true)
				return apiErr
			}(),
			want: false,
		},
		{
			name: "wrapped gRPC NotFound",
			err:  fmt.Errorf("getting resource: %w", status.Error(codes.NotFound, "not found")),
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsNotFound(tc.err)
			if got != tc.want {
				t.Errorf("IsNotFound(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestJSON_FromProto(t *testing.T) {
	mapCtx := &MapContext{}

	// Test with a struct value
	m := map[string]interface{}{
		"a": 1.0,
		"b": "hello",
		"c": true,
	}
	v, err := structpb.NewValue(m)
	if err != nil {
		t.Fatalf("structpb.NewValue failed: %v", err)
	}

	jsonVal := JSON_FromProto(mapCtx, v)
	if mapCtx.Err() != nil {
		t.Fatalf("JSON_FromProto failed: %v", mapCtx.Err())
	}

	// Verify it's natural JSON (no "structValue" or "fields" at top level)
	var decoded map[string]interface{}
	if err := json.Unmarshal(jsonVal.Raw, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if decoded["a"] != 1.0 || decoded["b"] != "hello" || decoded["c"] != true {
		t.Errorf("decoded JSON mismatch: %v", decoded)
	}
}
