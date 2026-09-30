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

package sql

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestSQLInstance_DataCacheConfigDefaults(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		in           *krm.SQLInstance
		actual       *api.DatabaseInstance
		wantEnabled  bool
		wantDiff     bool
		diffFieldIDs []string
	}{
		{
			name: "Enterprise Plus with unset dataCacheConfig preserves actual true and has no drift",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition: direct.PtrTo("ENTERPRISE_PLUS"),
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					Edition: "ENTERPRISE_PLUS",
					DataCacheConfig: &api.DataCacheConfig{
						DataCacheEnabled: true,
					},
				},
			},
			wantEnabled: true,
			wantDiff:    false,
		},
		{
			name: "Enterprise Plus with empty dataCacheConfig preserves actual true and has no drift",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition:         direct.PtrTo("ENTERPRISE_PLUS"),
						DataCacheConfig: &krm.InstanceDataCacheConfig{},
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					Edition: "ENTERPRISE_PLUS",
					DataCacheConfig: &api.DataCacheConfig{
						DataCacheEnabled: true,
					},
				},
			},
			wantEnabled: true,
			wantDiff:    false,
		},
		{
			name: "Enterprise Plus with explicit dataCacheEnabled: false detects drift when actual is true",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition: direct.PtrTo("ENTERPRISE_PLUS"),
						DataCacheConfig: &krm.InstanceDataCacheConfig{
							DataCacheEnabled: direct.PtrTo(false),
						},
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					Edition: "ENTERPRISE_PLUS",
					DataCacheConfig: &api.DataCacheConfig{
						DataCacheEnabled: true,
					},
				},
			},
			wantEnabled:  false,
			wantDiff:     true,
			diffFieldIDs: []string{".settings.dataCacheConfig.dataCacheEnabled"},
		},
		{
			name: "Enterprise Plus with explicit dataCacheEnabled: true detects drift when actual is false",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition: direct.PtrTo("ENTERPRISE_PLUS"),
						DataCacheConfig: &krm.InstanceDataCacheConfig{
							DataCacheEnabled: direct.PtrTo(true),
						},
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					Edition: "ENTERPRISE_PLUS",
					DataCacheConfig: &api.DataCacheConfig{
						DataCacheEnabled: false,
					},
				},
			},
			wantEnabled:  true,
			wantDiff:     true,
			diffFieldIDs: []string{".settings.dataCacheConfig.dataCacheEnabled"},
		},
		{
			name: "Enterprise edition with unset dataCacheConfig and nil actual has no drift",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition: direct.PtrTo("ENTERPRISE"),
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					Edition:         "ENTERPRISE",
					DataCacheConfig: nil,
				},
			},
			wantEnabled: false,
			wantDiff:    false,
		},
		{
			name: "Create phase with actual == nil and unset dataCacheConfig produces nil dataCacheConfig",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Edition: direct.PtrTo("ENTERPRISE_PLUS"),
					},
				},
			},
			actual:      nil,
			wantEnabled: false,
			wantDiff:    false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gcpObj, err := SQLInstanceKRMToGCP(tc.in, tc.actual, nil)
			if err != nil {
				t.Fatalf("SQLInstanceKRMToGCP() failed: %v", err)
			}

			if tc.actual == nil {
				if gcpObj.Settings.DataCacheConfig != nil {
					t.Errorf("expected nil DataCacheConfig on create, got: %v", gcpObj.Settings.DataCacheConfig)
				}
				return
			}

			gotEnabled := false
			if gcpObj.Settings.DataCacheConfig != nil {
				gotEnabled = gcpObj.Settings.DataCacheConfig.DataCacheEnabled
			}
			if gotEnabled != tc.wantEnabled {
				t.Errorf("expected DataCacheEnabled=%v, got=%v", tc.wantEnabled, gotEnabled)
			}

			diff := DiffDataCacheConfig(gcpObj.Settings.DataCacheConfig, tc.actual.Settings.DataCacheConfig)
			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("expected HasDiff()=%v, got=%v (diffs: %v)", tc.wantDiff, diff.HasDiff(), diff.Fields)
			}

			if tc.wantDiff {
				for _, expectedID := range tc.diffFieldIDs {
					found := false
					for _, field := range diff.Fields {
						if field.ID == expectedID {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("expected diff on %q, but diffs were: %v", expectedID, diff.Fields)
					}
				}
			}
		})
	}
}
