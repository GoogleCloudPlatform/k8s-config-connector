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
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestApplySQLInstanceGCPDefaults_RetainedBackups(t *testing.T) {
	// Case 1: User provides BackupConfiguration, but RetainedBackups is 0 (default/unset).
	t.Run("DefaultAppliedWhenZero", func(t *testing.T) {
		in := &krm.SQLInstance{
			Spec: krm.SQLInstanceSpec{
				Settings: krm.InstanceSettings{
					BackupConfiguration: &krm.InstanceBackupConfiguration{
						Enabled: boolPtr(true),
						BackupRetentionSettings: &krm.InstanceBackupRetentionSettings{
							RetainedBackups: 0,
						},
					},
					Edition: stringPtr("ENTERPRISE"),
				},
			},
		}
		out := &api.DatabaseInstance{
			Settings: &api.Settings{
				BackupConfiguration: &api.BackupConfiguration{
					Enabled: true,
					BackupRetentionSettings: &api.BackupRetentionSettings{
						RetainedBackups: 0,
					},
				},
				Edition: "ENTERPRISE",
			},
		}

		ApplySQLInstanceGCPDefaults(in, out, nil, nil)

		if out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups != 7 {
			t.Errorf("expected 7, got %d", out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups)
		}
		if out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit != "COUNT" {
			t.Errorf("expected COUNT, got %s", out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit)
		}
	})

	// Case 2: Enterprise Plus defaults to 15 when 0.
	t.Run("EnterprisePlusDefault", func(t *testing.T) {
		in := &krm.SQLInstance{
			Spec: krm.SQLInstanceSpec{
				Settings: krm.InstanceSettings{
					BackupConfiguration: &krm.InstanceBackupConfiguration{
						BackupRetentionSettings: &krm.InstanceBackupRetentionSettings{
							RetainedBackups: 0,
						},
					},
					Edition: stringPtr("ENTERPRISE_PLUS"),
				},
			},
		}
		out := &api.DatabaseInstance{
			Settings: &api.Settings{
				BackupConfiguration: &api.BackupConfiguration{
					BackupRetentionSettings: &api.BackupRetentionSettings{
						RetainedBackups: 0,
					},
				},
				Edition: "ENTERPRISE_PLUS",
			},
		}

		ApplySQLInstanceGCPDefaults(in, out, nil, nil)

		if out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups != 15 {
			t.Errorf("expected 15, got %d", out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups)
		}
	})

	// Case 3: No BackupRetentionSettings provided in KRM, but backups enabled.
	t.Run("NoRetentionSettingsButEnabled", func(t *testing.T) {
		in := &krm.SQLInstance{
			Spec: krm.SQLInstanceSpec{
				Settings: krm.InstanceSettings{
					BackupConfiguration: &krm.InstanceBackupConfiguration{
						Enabled: boolPtr(true),
					},
					Edition: stringPtr("ENTERPRISE"),
				},
			},
		}
		out := &api.DatabaseInstance{
			Settings: &api.Settings{
				BackupConfiguration: &api.BackupConfiguration{
					Enabled: true,
					// Simulate out having nil BackupRetentionSettings initially
					BackupRetentionSettings: nil,
				},
				Edition: "ENTERPRISE",
			},
		}

		ApplySQLInstanceGCPDefaults(in, out, nil, nil)

		if out.Settings.BackupConfiguration.BackupRetentionSettings == nil {
			t.Fatalf("expected BackupRetentionSettings to be created")
		}

		if out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups != 7 {
			t.Errorf("expected 7, got %d", out.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups)
		}
		if out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit != "COUNT" {
			t.Errorf("expected COUNT, got %s", out.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit)
		}
	})
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
