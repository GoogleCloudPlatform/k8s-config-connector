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

package v1beta1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRedisInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *RedisInstanceIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/instances/my-instance",
			want: &RedisInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://redis.googleapis.com/projects/my-project/locations/us-central1/instances/my-instance",
			want: &RedisInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &RedisInstanceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("RedisInstanceIdentity mismatch (-want +got):\n%s", diff)
				}
				expectedString := "projects/my-project/locations/us-central1/instances/my-instance"
				if i.String() != expectedString {
					t.Errorf("String() = %q, want %q", i.String(), expectedString)
				}
				expectedParent := "projects/my-project/locations/us-central1"
				if i.ParentString() != expectedParent {
					t.Errorf("ParentString() = %q, want %q", i.ParentString(), expectedParent)
				}
			}
		})
	}
}

func TestRedisInstanceRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/instances/my-instance",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/instances/my-instance",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/instances/my-instance",
			wantErr: true,
		},
		{
			name:    "missing instance",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisInstanceRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("RedisInstanceRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
