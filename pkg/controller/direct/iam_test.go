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

package direct

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
)

func TestMemberMatches(t *testing.T) {
	tests := []struct {
		member     string
		wantMember string
		matches    bool
	}{
		{
			member:     "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "deleted:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?uid=12345",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "deleted:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "serviceAccount:other@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    false,
		},
		{
			member:     "deleted:serviceAccount:other@project.iam.gserviceaccount.com?uid=67890",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    false,
		},
		{
			member:     "serviceAccount:KCC-REPRO-SA@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "DELETED:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?UID=12345",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "DELETED:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?UID=12345",
			matches:    true,
		},
	}

	for _, tc := range tests {
		got := memberMatches(tc.member, tc.wantMember)
		if got != tc.matches {
			t.Errorf("memberMatches(%q, %q) = %v, want %v", tc.member, tc.wantMember, got, tc.matches)
		}
	}
}

func TestIAMBindingsRoundTrip(t *testing.T) {
	// Arrange
	input := []v1beta1.IAMPolicyBinding{
		{
			Role: "roles/viewer",
			Members: []v1beta1.Member{
				"user:alice@example.com",
				"group:devs@example.com",
			},
			Condition: &v1beta1.IAMCondition{
				Title:       "expires-soon",
				Description: "access expires after end of year",
				Expression:  "request.time < timestamp('2027-01-01T00:00:00Z')",
			},
		},
		{
			Role: "roles/editor",
			Members: []v1beta1.Member{
				"serviceAccount:sa@project.iam.gserviceaccount.com",
			},
		},
	}

	// Act
	protoBindings := IAMBindingsToProto(input)
	got := IAMBindingsFromProto(protoBindings)

	// Assert
	if len(got) != len(input) {
		t.Fatalf("expected %d bindings, got %d", len(input), len(got))
	}
	for i := range input {
		if got[i].Role != input[i].Role {
			t.Errorf("binding[%d].Role = %q, want %q", i, got[i].Role, input[i].Role)
		}
		if len(got[i].Members) != len(input[i].Members) {
			t.Errorf("binding[%d].Members length = %d, want %d", i, len(got[i].Members), len(input[i].Members))
		}
		for j := range input[i].Members {
			if got[i].Members[j] != input[i].Members[j] {
				t.Errorf("binding[%d].Members[%d] = %q, want %q", i, j, got[i].Members[j], input[i].Members[j])
			}
		}
		if input[i].Condition != nil {
			if got[i].Condition == nil {
				t.Errorf("binding[%d].Condition is nil, want non-nil", i)
			} else {
				if got[i].Condition.Title != input[i].Condition.Title {
					t.Errorf("binding[%d].Condition.Title = %q, want %q", i, got[i].Condition.Title, input[i].Condition.Title)
				}
				if got[i].Condition.Description != input[i].Condition.Description {
					t.Errorf("binding[%d].Condition.Description = %q, want %q", i, got[i].Condition.Description, input[i].Condition.Description)
				}
				if got[i].Condition.Expression != input[i].Condition.Expression {
					t.Errorf("binding[%d].Condition.Expression = %q, want %q", i, got[i].Condition.Expression, input[i].Condition.Expression)
				}
			}
		} else if got[i].Condition != nil {
			t.Errorf("binding[%d].Condition = %+v, want nil", i, got[i].Condition)
		}
	}
}

func TestIAMAuditConfigsRoundTrip(t *testing.T) {
	// Arrange
	input := []v1beta1.IAMPolicyAuditConfig{
		{
			Service: "allServices",
			AuditLogConfigs: []v1beta1.AuditLogConfig{
				{
					LogType: "DATA_READ",
					ExemptedMembers: []v1beta1.Member{
						"user:admin@example.com",
					},
				},
				{
					LogType: "DATA_WRITE",
				},
				{
					LogType: "ADMIN_READ",
				},
			},
		},
		{
			Service: "compute.googleapis.com",
			AuditLogConfigs: []v1beta1.AuditLogConfig{
				{
					LogType: "DATA_READ",
				},
			},
		},
	}

	// Act
	protoAuditConfigs := IAMAuditConfigsToProto(input)
	got := IAMAuditConfigsFromProto(protoAuditConfigs)

	// Assert
	if len(got) != len(input) {
		t.Fatalf("expected %d audit configs, got %d", len(input), len(got))
	}
	for i := range input {
		if got[i].Service != input[i].Service {
			t.Errorf("auditConfig[%d].Service = %q, want %q", i, got[i].Service, input[i].Service)
		}
		if len(got[i].AuditLogConfigs) != len(input[i].AuditLogConfigs) {
			t.Errorf("auditConfig[%d].AuditLogConfigs length = %d, want %d", i, len(got[i].AuditLogConfigs), len(input[i].AuditLogConfigs))
		}
		for j := range input[i].AuditLogConfigs {
			if got[i].AuditLogConfigs[j].LogType != input[i].AuditLogConfigs[j].LogType {
				t.Errorf("auditConfig[%d].AuditLogConfigs[%d].LogType = %q, want %q", i, j, got[i].AuditLogConfigs[j].LogType, input[i].AuditLogConfigs[j].LogType)
			}
			for k := range input[i].AuditLogConfigs[j].ExemptedMembers {
				if got[i].AuditLogConfigs[j].ExemptedMembers[k] != input[i].AuditLogConfigs[j].ExemptedMembers[k] {
					t.Errorf("auditConfig[%d].AuditLogConfigs[%d].ExemptedMembers[%d] = %q, want %q", i, j, k, got[i].AuditLogConfigs[j].ExemptedMembers[k], input[i].AuditLogConfigs[j].ExemptedMembers[k])
				}
			}
		}
	}
}

func TestMapLogType(t *testing.T) {
	// Arrange
	tests := []struct {
		krmLogType string
		wantKRM    string
	}{
		{krmLogType: "DATA_READ", wantKRM: "DATA_READ"},
		{krmLogType: "DATA_WRITE", wantKRM: "DATA_WRITE"},
		{krmLogType: "ADMIN_READ", wantKRM: "ADMIN_READ"},
		{krmLogType: "UNKNOWN", wantKRM: "LOG_TYPE_UNSPECIFIED"},
	}

	for _, tc := range tests {
		// Act
		pbType := mapV1Beta1LogTypeToProto(tc.krmLogType)
		got := mapProtoLogTypeToKRM(pbType)

		// Assert
		if got != tc.wantKRM {
			t.Errorf("mapProtoLogTypeToKRM(mapV1Beta1LogTypeToProto(%q)) = %q, want %q", tc.krmLogType, got, tc.wantKRM)
		}
	}
}

