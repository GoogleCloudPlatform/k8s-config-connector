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

package appengine

import (
	"context"
	"math"
	"reflect"
	"strings"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
)

func TestAppEngineFirewallRuleAdapter_DesiredRule(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		obj          *krm.AppEngineFirewallRule
		wantPriority int32
		wantAction   pb.FirewallRule_Action
		wantSource   string
		wantDesc     string
		expectError  bool
	}{
		{
			name: "populates priority from metadata.name when spec.resourceID is unset",
			obj: &krm.AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{Name: "200"},
				Spec: krm.AppEngineFirewallRuleSpec{
					Project:     ptr.To("test-proj"),
					Action:      "ALLOW",
					SourceRange: "10.0.0.0/8",
					Description: ptr.To("corp range"),
				},
			},
			wantPriority: 200,
			wantAction:   pb.FirewallRule_ALLOW,
			wantSource:   "10.0.0.0/8",
			wantDesc:     "corp range",
		},
		{
			name: "populates default rule priority from metadata.name when spec.resourceID is unset",
			obj: &krm.AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{Name: "2147483647"},
				Spec: krm.AppEngineFirewallRuleSpec{
					Project:     ptr.To("test-proj"),
					Action:      "DENY",
					SourceRange: "*",
				},
			},
			wantPriority: math.MaxInt32,
			wantAction:   pb.FirewallRule_DENY,
			wantSource:   "*",
			wantDesc:     "",
		},
		{
			name: "populates priority from spec.resourceID with leading zeros canonicalized",
			obj: &krm.AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{Name: "allow-corp"},
				Spec: krm.AppEngineFirewallRuleSpec{
					Project:     ptr.To("test-proj"),
					ResourceID:  ptr.To("01000"),
					Action:      "ALLOW",
					SourceRange: "192.0.2.0/24",
				},
			},
			wantPriority: 1000,
			wantAction:   pb.FirewallRule_ALLOW,
			wantSource:   "192.0.2.0/24",
		},
		{
			name: "rejects priority 0 during Create and Update",
			obj: &krm.AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{Name: "zero-rule"},
				Spec: krm.AppEngineFirewallRuleSpec{
					Project:     ptr.To("test-proj"),
					ResourceID:  ptr.To("0"),
					Action:      "ALLOW",
					SourceRange: "192.0.2.0/24",
				},
			},
			expectError: true,
		},
		{
			name: "rejects negative priority during Create and Update",
			obj: &krm.AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{Name: "neg-rule"},
				Spec: krm.AppEngineFirewallRuleSpec{
					Project:     ptr.To("test-proj"),
					ResourceID:  ptr.To("-5"),
					Action:      "ALLOW",
					SourceRange: "192.0.2.0/24",
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rawID, idErr := tc.obj.GetIdentity(ctx, nil)
			if idErr != nil {
				t.Fatalf("unexpected GetIdentity error: %v", idErr)
			}
			adp := &adapter{
				id:      rawID.(*krm.AppEngineFirewallRuleIdentity),
				desired: tc.obj,
			}

			got, err := adp.desiredRule()

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error from desiredRule(), got %+v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error from desiredRule(): %v", err)
			}
			if got.GetPriority() != tc.wantPriority {
				t.Errorf("Priority = %d, want %d", got.GetPriority(), tc.wantPriority)
			}
			if got.GetAction() != tc.wantAction {
				t.Errorf("Action = %v, want %v", got.GetAction(), tc.wantAction)
			}
			if got.GetSourceRange() != tc.wantSource {
				t.Errorf("SourceRange = %q, want %q", got.GetSourceRange(), tc.wantSource)
			}
			if got.GetDescription() != tc.wantDesc {
				t.Errorf("Description = %q, want %q", got.GetDescription(), tc.wantDesc)
			}
		})
	}
}

func TestAppEngineFirewallRuleAdapter_IsDefaultRule(t *testing.T) {
	tests := []struct {
		name string
		id   *krm.AppEngineFirewallRuleIdentity
		want bool
	}{
		{
			name: "default rule 2147483647",
			id:   &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "2147483647"},
			want: true,
		},
		{
			name: "custom rule 200",
			id:   &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "200"},
			want: false,
		},
		{
			name: "zero priority",
			id:   &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "0"},
			want: false,
		},
		{
			name: "non-numeric priority",
			id:   &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "invalid"},
			want: false,
		},
		{
			name: "nil identity",
			id:   nil,
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adp := &adapter{id: tc.id}

			got := adp.isDefaultRule()

			if got != tc.want {
				t.Errorf("isDefaultRule() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAppEngineFirewallRuleAdapter_ValidateDefaultRuleUpdate_RejectsNonActionFields(t *testing.T) {
	adp := &adapter{
		id: &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "2147483647"},
	}
	tests := []struct {
		name      string
		paths     []string
		wantField string
	}{
		{
			name:      "source_range change rejected",
			paths:     []string{"source_range"},
			wantField: "[source_range]",
		},
		{
			name:      "description change rejected",
			paths:     []string{"description"},
			wantField: "[description]",
		},
		{
			name:      "mixed action and source_range change rejected",
			paths:     []string{"source_range", "action"},
			wantField: "[action source_range]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := adp.validateDefaultRuleUpdate(tc.paths)

			if err == nil {
				t.Fatalf("expected error for paths %v, got nil", tc.paths)
			}
			if !strings.Contains(err.Error(), "only the action of the default rule can be modified") {
				t.Errorf("unexpected error message: %v", err)
			}
			if !strings.Contains(err.Error(), tc.wantField) {
				t.Errorf("expected error %q to contain %q", err.Error(), tc.wantField)
			}
		})
	}
}

func TestAppEngineFirewallRuleAdapter_ValidateDefaultRuleUpdate_AllowsActionOnly(t *testing.T) {
	adp := &adapter{
		id: &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "2147483647"},
	}

	err := adp.validateDefaultRuleUpdate([]string{"action"})

	if err != nil {
		t.Fatalf("expected nil error for action-only default rule update, got: %v", err)
	}
}

func TestAppEngineFirewallRuleAdapter_ValidateDefaultRuleUpdate_AllowsCustomRuleFields(t *testing.T) {
	adp := &adapter{
		id: &krm.AppEngineFirewallRuleIdentity{App: "test-proj", IngressRule: "200"},
	}

	err := adp.validateDefaultRuleUpdate([]string{"action", "description", "source_range"})

	if err != nil {
		t.Fatalf("expected nil error for custom rule update, got: %v", err)
	}
}

func TestAppEngineFirewallRule_CompareBrownfieldSpec_ClearAndPreserveDescription(t *testing.T) {
	ctx := context.Background()
	actual := &pb.FirewallRule{
		Priority:    200,
		Action:      pb.FirewallRule_ALLOW,
		SourceRange: "10.0.0.0/8",
		Description: "existing description",
	}

	t.Run("explicit empty string produces description updateMask", func(t *testing.T) {
		desired := &krm.AppEngineFirewallRuleSpec{
			Action:      "ALLOW",
			SourceRange: "10.0.0.0/8",
			Description: ptr.To(""),
		}

		diffs, mask, err := common.CompareBrownfieldSpec(
			ctx,
			desired,
			actual,
			AppEngineFirewallRuleSpec_v1alpha1_FromProto,
			AppEngineFirewallRuleSpec_v1alpha1_ToProto,
			nil,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diffs.HasDiff() {
			t.Fatalf("expected diff when clearing description with empty string")
		}
		if !reflect.DeepEqual(mask.GetPaths(), []string{"description"}) {
			t.Errorf("updateMask paths = %v, want [description]", mask.GetPaths())
		}
	})

	t.Run("omitted description preserves live value without diff", func(t *testing.T) {
		desired := &krm.AppEngineFirewallRuleSpec{
			Action:      "ALLOW",
			SourceRange: "10.0.0.0/8",
		}

		diffs, _, err := common.CompareBrownfieldSpec(
			ctx,
			desired,
			actual,
			AppEngineFirewallRuleSpec_v1alpha1_FromProto,
			AppEngineFirewallRuleSpec_v1alpha1_ToProto,
			nil,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diffs.HasDiff() {
			t.Errorf("expected no diff when description is omitted, got %v", diffs)
		}
	})
}
