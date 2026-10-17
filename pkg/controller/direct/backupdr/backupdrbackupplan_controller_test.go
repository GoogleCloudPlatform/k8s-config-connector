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

package backupdr

import (
	"context"
	"testing"

	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
)

// backupPlanResponse is what the API returns for the BackupPlan created from the spec used in TestCompareBackupPlan:
// the backup vault uses the project number, and output-only fields are populated.
const backupPlanResponse = `{
  "name": "projects/my-project/locations/us-central1/backupPlans/my-plan",
  "description": "Daily backups",
  "labels": {"env": "test", "managed-by-cnrm": "true"},
  "createTime": "2024-04-01T12:34:56.123456Z",
  "updateTime": "2024-04-01T12:34:56.123456Z",
  "backupRules": [
    {
      "ruleId": "daily",
      "backupRetentionDays": 30,
      "standardSchedule": {
        "recurrenceType": 2,
        "backupWindow": {"startHourOfDay": 2, "endHourOfDay": 6},
        "timeZone": "Europe/Paris"
      }
    }
  ],
  "state": 2,
  "resourceType": "compute.googleapis.com/Instance",
  "etag": "abcdef0123A=",
  "backupVault": "projects/123456789/locations/us-central1/backupVaults/my-vault",
  "backupVaultServiceAccount": "vault-123456789-12345@gcp-sa-backupdr-pr.iam.gserviceaccount.com",
  "revisionId": "v0",
  "supportedResourceTypes": ["compute.googleapis.com/Instance"]
}`

func TestCompareBackupPlan(t *testing.T) {
	ctx := context.Background()

	spec := &krm.BackupDRBackupPlanSpec{
		Description:  direct.LazyPtr("Daily backups"),
		ResourceType: direct.LazyPtr("compute.googleapis.com/Instance"),
		BackupVaultRef: &krm.BackupVaultRef{
			External: "projects/my-project/locations/us-central1/backupVaults/my-vault",
		},
		BackupRules: []krm.BackupRule{
			{
				RuleID:              direct.LazyPtr("daily"),
				BackupRetentionDays: direct.LazyPtr(int32(30)),
				StandardSchedule: &krm.StandardSchedule{
					RecurrenceType: direct.LazyPtr("DAILY"),
					BackupWindow: &krm.BackupWindow{
						StartHourOfDay: direct.LazyPtr(int32(2)),
						EndHourOfDay:   direct.LazyPtr(int32(6)),
					},
					TimeZone: direct.LazyPtr("Europe/Paris"),
				},
			},
		},
	}

	tests := []struct {
		name      string
		mutate    func(actual *pb.BackupPlan)
		wantPaths []string
	}{
		{
			name:      "same backup rules as returned by the API",
			mutate:    func(actual *pb.BackupPlan) {},
			wantPaths: nil,
		},
		{
			name: "different backup retention",
			mutate: func(actual *pb.BackupPlan) {
				actual.BackupRules[0].BackupRetentionDays = 7
			},
			wantPaths: []string{"backup_rules"},
		},
		{
			name: "different backup window",
			mutate: func(actual *pb.BackupPlan) {
				actual.BackupRules[0].GetStandardSchedule().GetBackupWindow().EndHourOfDay = 8
			},
			wantPaths: []string{"backup_rules"},
		},
		{
			name: "additional backup rule",
			mutate: func(actual *pb.BackupPlan) {
				actual.BackupRules = append(actual.BackupRules, &pb.BackupRule{RuleId: "weekly", BackupRetentionDays: 90})
			},
			wantPaths: []string{"backup_rules"},
		},
		{
			name: "different description",
			mutate: func(actual *pb.BackupPlan) {
				actual.Description = "Weekly backups"
			},
			wantPaths: []string{"description"},
		},
		{
			name: "different labels",
			mutate: func(actual *pb.BackupPlan) {
				actual.Labels["env"] = "prod"
			},
			wantPaths: []string{"labels"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			desired := BackupDRBackupPlanSpec_v1beta1_ToProto(mapCtx, spec)
			if mapCtx.Err() != nil {
				t.Fatalf("BackupDRBackupPlanSpec_v1beta1_ToProto() unexpected error: %v", mapCtx.Err())
			}
			desired.Labels = label.NewGCPLabelsFromK8sLabels(map[string]string{"env": "test"})

			// Decode the response the same way as the REST client.
			actual := &pb.BackupPlan{}
			unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
			if err := unm.Unmarshal([]byte(backupPlanResponse), actual); err != nil {
				t.Fatalf("unmarshaling BackupPlan: %v", err)
			}
			tc.mutate(actual)

			diffs, updateMask, err := compareBackupPlan(ctx, actual, desired)
			if err != nil {
				t.Fatalf("compareBackupPlan() unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantPaths, updateMask.GetPaths()); diff != "" {
				t.Errorf("compareBackupPlan() paths mismatch (-want +got):\n%s", diff)
			}
			if got, want := diffs.HasDiff(), len(tc.wantPaths) != 0; got != want {
				t.Errorf("compareBackupPlan() HasDiff() = %v, want %v", got, want)
			}
		})
	}
}
