// Copyright 2026 Google LLC
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

package bigquery

import (
	"errors"
	"testing"

	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAdapter_IsUnreadableButDeletable(t *testing.T) {
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
			name: "HTTP 400 Bad Request api error",
			err: &googleapi.Error{
				Code:    400,
				Message: "Bad Request",
			},
			want: true,
		},
		{
			name: "HTTP 404 Not Found api error",
			err: &googleapi.Error{
				Code:    404,
				Message: "Not Found",
			},
			want: false,
		},
		{
			name: "HTTP 403 Forbidden api error",
			err: &googleapi.Error{
				Code:    403,
				Message: "Forbidden",
			},
			want: false,
		},
		{
			name: "HTTP 500 Internal Server api error",
			err: &googleapi.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			want: false,
		},
		{
			name: "grpc InvalidArgument",
			err:  status.Error(codes.InvalidArgument, "invalid argument"),
			want: false,
		},
	}

	a := &Adapter{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.err
			var gErr *googleapi.Error
			if errors.As(err, &gErr) {
				if apiErr, wrapperOk := apierror.FromError(gErr); wrapperOk {
					err = apiErr
				}
			}

			got := a.IsUnreadableButDeletable(err)
			if got != tt.want {
				t.Errorf("IsUnreadableButDeletable() = %v, want %v for error: %v", got, tt.want, err)
			}
		})
	}
}
