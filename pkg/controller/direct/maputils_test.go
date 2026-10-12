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
	"fmt"
	"testing"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
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
		{
			name: "googleapi 409 Conflict",
			err:  &googleapi.Error{Code: 409, Message: "Conflict"},
			want: true,
		},
		{
			name: "googleapi 400 Bad Request",
			err:  &googleapi.Error{Code: 400, Message: "Bad Request"},
			want: false,
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
		{
			name: "googleapi.Error 404",
			err: &googleapi.Error{
				Code:    404,
				Message: "Not Found",
			},
			want: true,
		},
		{
			name: "googleapi.Error 400",
			err: &googleapi.Error{
				Code:    400,
				Message: "Bad Request",
			},
			want: false,
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

func TestIsBadRequest(t *testing.T) {
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
			name: "googleapi.Error 400",
			err: &googleapi.Error{
				Code:    400,
				Message: "Bad Request",
			},
			want: true,
		},
		{
			name: "googleapi.Error 404",
			err: &googleapi.Error{
				Code:    404,
				Message: "Not Found",
			},
			want: false,
		},
		{
			name: "wrapped googleapi.Error 400",
			err: fmt.Errorf("calling API: %w", &googleapi.Error{
				Code:    400,
				Message: "Bad Request",
			}),
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsBadRequest(tc.err)
			if got != tc.want {
				t.Errorf("IsBadRequest(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
