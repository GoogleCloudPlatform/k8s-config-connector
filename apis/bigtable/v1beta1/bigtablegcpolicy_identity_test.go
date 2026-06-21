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
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestBigtableGCPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigtableGCPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/instances/my-instance/tables/my-table/columnFamilies/my-family",
			want: &BigtableGCPolicyIdentity{
				Project:      "my-project",
				Instance:     "my-instance",
				Table:        "my-table",
				ColumnFamily: "my-family",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigtable.googleapis.com/projects/my-project/instances/my-instance/tables/my-table/columnFamilies/my-family",
			want: &BigtableGCPolicyIdentity{
				Project:      "my-project",
				Instance:     "my-instance",
				Table:        "my-table",
				ColumnFamily: "my-family",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigtableGCPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				// Test String() method as well to ensure roundtrip works
				gotStr := i.String()
				wantStr := "projects/my-project/instances/my-instance/tables/my-table/columnFamilies/my-family"
				if gotStr != wantStr {
					t.Errorf("String() = %v, want %v", gotStr, wantStr)
				}
			}
		})
	}
}

type mockReader struct {
	client.Reader
	getFunc func(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error
}

func (m *mockReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if m.getFunc != nil {
		return m.getFunc(ctx, key, obj, opts...)
	}
	return nil
}

func TestGetIdentityFromBigtableGCPolicySpec(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *BigtableGCPolicy
		mockGetFunc func(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error
		wantErr     string
	}{
		{
			name: "empty TableRef and ColumnFamily",
			obj: &BigtableGCPolicy{
				Spec: BigtableGCPolicySpec{},
			},
			mockGetFunc: func(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
				return errors.New("not found")
			},
			wantErr: "reading referenced",
		},
		{
			name: "empty ColumnFamily but valid TableRef external",
			obj: &BigtableGCPolicy{
				Spec: BigtableGCPolicySpec{
					TableRef: TableRef{
						External: "projects/my-project/instances/my-instance/tables/my-table",
					},
				},
			},
			wantErr: "cannot resolve ColumnFamily: empty string",
		},
		{
			name: "both Name and External specified on TableRef",
			obj: &BigtableGCPolicy{
				Spec: BigtableGCPolicySpec{
					TableRef: TableRef{
						External: "projects/my-project/instances/my-instance/tables/my-table",
						Name:     "my-table",
					},
				},
			},
			wantErr: "cannot specify both name and external",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &mockReader{getFunc: tt.mockGetFunc}
			got, err := getIdentityFromBigtableGCPolicySpec(ctx, reader, tt.obj)
			if err == nil {
				t.Fatalf("getIdentityFromBigtableGCPolicySpec() expected error, got nil identity: %v", got)
			}
			if tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("getIdentityFromBigtableGCPolicySpec() error = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}
