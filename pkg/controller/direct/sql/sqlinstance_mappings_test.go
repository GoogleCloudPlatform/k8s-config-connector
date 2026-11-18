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
	"encoding/json"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestSQLInstanceKRMToGCP_RetainBackupsOnDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		spec      *bool
		live      *bool
		wantValue bool
		wantKey   bool
	}{
		{
			name:      "create with unset omits key and defaults value to false",
			spec:      nil,
			live:      nil,
			wantValue: false,
			wantKey:   false,
		},
		{
			name:      "create with true sends true",
			spec:      direct.PtrTo(true),
			live:      nil,
			wantValue: true,
			wantKey:   true,
		},
		{
			name:      "create with false force-sends false",
			spec:      direct.PtrTo(false),
			live:      nil,
			wantValue: false,
			wantKey:   true,
		},
		{
			name:      "update with unset preserves live true",
			spec:      nil,
			live:      direct.PtrTo(true),
			wantValue: true,
			wantKey:   true,
		},
		{
			name:      "update with unset and live false omits key",
			spec:      nil,
			live:      direct.PtrTo(false),
			wantValue: false,
			wantKey:   false,
		},
		{
			name:      "update with explicit false overrides live true and force-sends false",
			spec:      direct.PtrTo(false),
			live:      direct.PtrTo(true),
			wantValue: false,
			wantKey:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Settings: krm.InstanceSettings{
						Tier:                  "db-custom-1-3840",
						RetainBackupsOnDelete: tc.spec,
					},
				},
			}

			var actual *api.DatabaseInstance
			if tc.live != nil {
				actual = &api.DatabaseInstance{
					Settings: &api.Settings{
						BackupConfiguration:   &api.BackupConfiguration{},
						RetainBackupsOnDelete: *tc.live,
					},
				}
			}

			out, err := SQLInstanceKRMToGCP(in, actual, nil)
			if err != nil {
				t.Fatalf("SQLInstanceKRMToGCP() unexpected error: %v", err)
			}
			if out == nil || out.Settings == nil {
				t.Fatalf("SQLInstanceKRMToGCP() returned nil Settings")
			}

			if got := out.Settings.RetainBackupsOnDelete; got != tc.wantValue {
				t.Errorf("out.Settings.RetainBackupsOnDelete = %v, want %v", got, tc.wantValue)
			}

			raw, err := json.Marshal(out.Settings)
			if err != nil {
				t.Fatalf("json.Marshal(out.Settings) failed: %v", err)
			}
			var m map[string]any
			if err := json.Unmarshal(raw, &m); err != nil {
				t.Fatalf("json.Unmarshal(out.Settings) failed: %v", err)
			}
			gotVal, gotKey := m["retainBackupsOnDelete"]
			if gotKey != tc.wantKey {
				t.Errorf("JSON key \"retainBackupsOnDelete\" present = %v, want %v (json: %s)", gotKey, tc.wantKey, string(raw))
			}
			if tc.wantKey {
				gotBool, ok := gotVal.(bool)
				if !ok || gotBool != tc.wantValue {
					t.Errorf("JSON \"retainBackupsOnDelete\" = %v (%T), want %v", gotVal, gotVal, tc.wantValue)
				}
			}
		})
	}
}

func TestInstanceSettingsGCPToKRM_RetainBackupsOnDelete(t *testing.T) {
	t.Parallel()

	for _, want := range []bool{true, false} {
		got := InstanceSettingsGCPToKRM(&api.Settings{
			RetainBackupsOnDelete: want,
		})
		if got.RetainBackupsOnDelete == nil {
			t.Fatalf("InstanceSettingsGCPToKRM(RetainBackupsOnDelete=%v).RetainBackupsOnDelete = nil, want %v", want, want)
		}
		if *got.RetainBackupsOnDelete != want {
			t.Errorf("*InstanceSettingsGCPToKRM(RetainBackupsOnDelete=%v).RetainBackupsOnDelete = %v, want %v", want, *got.RetainBackupsOnDelete, want)
		}
	}
}
