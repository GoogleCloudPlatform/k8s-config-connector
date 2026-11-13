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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOracleDatabaseExadbVMClusterIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *OracleDatabaseExadbVMClusterIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/exadbVmClusters/my-cluster",
			want: &OracleDatabaseExadbVMClusterIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				ExadbVmCluster: "my-cluster",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://oracledatabase.googleapis.com/projects/my-project/locations/us-central1/exadbVmClusters/my-cluster",
			want: &OracleDatabaseExadbVMClusterIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				ExadbVmCluster: "my-cluster",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &OracleDatabaseExadbVMClusterIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestOracleDatabaseODBNetworkRef(t *testing.T) {
	tests := []struct {
		name    string
		ref     OracleDatabaseODBNetworkRef
		wantErr bool
	}{
		{
			name: "valid reference",
			ref: OracleDatabaseODBNetworkRef{
				External: "projects/my-project/locations/us-central1/odbNetworks/my-network",
			},
		},
		{
			name: "invalid reference format",
			ref: OracleDatabaseODBNetworkRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
		{
			name: "empty external",
			ref: OracleDatabaseODBNetworkRef{
				External: "",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, nil, "default")
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.ref.External != "" {
				err := tt.ref.ValidateExternal(tt.ref.External)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestOracleDatabaseODBSubnetRef(t *testing.T) {
	tests := []struct {
		name    string
		ref     OracleDatabaseODBSubnetRef
		wantErr bool
	}{
		{
			name: "valid reference",
			ref: OracleDatabaseODBSubnetRef{
				External: "projects/my-project/locations/us-central1/odbNetworks/my-network/odbSubnets/my-subnet",
			},
		},
		{
			name: "invalid reference format",
			ref: OracleDatabaseODBSubnetRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
		{
			name: "empty external",
			ref: OracleDatabaseODBSubnetRef{
				External: "",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, nil, "default")
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.ref.External != "" {
				err := tt.ref.ValidateExternal(tt.ref.External)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestOracleDatabaseExascaleDBStorageVaultRef(t *testing.T) {
	tests := []struct {
		name    string
		ref     OracleDatabaseExascaleDBStorageVaultRef
		wantErr bool
	}{
		{
			name: "valid reference",
			ref: OracleDatabaseExascaleDBStorageVaultRef{
				External: "projects/my-project/locations/us-central1/exascaleDbStorageVaults/my-vault",
			},
		},
		{
			name: "invalid reference format",
			ref: OracleDatabaseExascaleDBStorageVaultRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
		{
			name: "empty external",
			ref: OracleDatabaseExascaleDBStorageVaultRef{
				External: "",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, nil, "default")
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.ref.External != "" {
				err := tt.ref.ValidateExternal(tt.ref.External)
				if (err != nil) != tt.wantErr {
					t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
