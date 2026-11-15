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

package mockappengine

import (
	"context"
	"math"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func newTestFirewallServer() *firewallServer {
	return &firewallServer{
		MockService: &MockService{
			storage: storage.NewInMemoryStorage(),
		},
	}
}

func TestFirewallServer_CreateAutoAssignsSequentialPriorities(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()
	first, err := srv.CreateIngressRule(ctx, &pb.CreateIngressRuleRequest{
		Parent: "apps/test-app",
		Rule: &pb.FirewallRule{
			Action:      pb.FirewallRule_ALLOW,
			SourceRange: "10.0.0.0/8",
		},
	})
	if err != nil {
		t.Fatalf("first CreateIngressRule setup failed: %v", err)
	}

	second, err := srv.CreateIngressRule(ctx, &pb.CreateIngressRuleRequest{
		Parent: "apps/test-app",
		Rule: &pb.FirewallRule{
			Action:      pb.FirewallRule_DENY,
			SourceRange: "192.168.0.0/16",
		},
	})

	if err != nil {
		t.Fatalf("second CreateIngressRule failed: %v", err)
	}
	if first.GetPriority() != 1000 {
		t.Errorf("first auto-assigned priority = %d, want 1000", first.GetPriority())
	}
	if second.GetPriority() != 2000 {
		t.Errorf("second auto-assigned priority = %d, want 2000", second.GetPriority())
	}
}

func TestFirewallServer_CreateRejectsDefaultPriority(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()

	_, err := srv.CreateIngressRule(ctx, &pb.CreateIngressRuleRequest{
		Parent: "apps/test-app",
		Rule: &pb.FirewallRule{
			Priority:    math.MaxInt32,
			Action:      pb.FirewallRule_ALLOW,
			SourceRange: "*",
		},
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("CreateIngressRule(MaxInt32) code = %v, want InvalidArgument", status.Code(err))
	}
}

func TestFirewallServer_CreateRejectsDuplicatePriority(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()
	if _, err := srv.CreateIngressRule(ctx, &pb.CreateIngressRuleRequest{
		Parent: "apps/test-app",
		Rule: &pb.FirewallRule{
			Priority:    200,
			Action:      pb.FirewallRule_ALLOW,
			SourceRange: "10.0.0.0/8",
		},
	}); err != nil {
		t.Fatalf("initial CreateIngressRule(200) failed: %v", err)
	}

	_, dupErr := srv.CreateIngressRule(ctx, &pb.CreateIngressRuleRequest{
		Parent: "apps/test-app",
		Rule: &pb.FirewallRule{
			Priority:    200,
			Action:      pb.FirewallRule_DENY,
			SourceRange: "10.0.0.0/8",
		},
	})

	if status.Code(dupErr) != codes.AlreadyExists {
		t.Errorf("duplicate CreateIngressRule(200) code = %v, want AlreadyExists", status.Code(dupErr))
	}
}

func TestFirewallServer_GetSeedsDefaultRule(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()

	seeded, err := srv.GetIngressRule(ctx, &pb.GetIngressRuleRequest{
		Name: "apps/test-app/firewall/ingressRules/2147483647",
	})

	if err != nil {
		t.Fatalf("GetIngressRule(default) failed: %v", err)
	}
	if seeded.GetAction() != pb.FirewallRule_ALLOW || seeded.GetDescription() != "The default action." {
		t.Errorf("unexpected seeded default rule: %+v", seeded)
	}
}

func TestFirewallServer_UpdateDefaultRuleRejectsNonActionMask(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()

	_, err := srv.UpdateIngressRule(ctx, &pb.UpdateIngressRuleRequest{
		Name: "apps/test-app/firewall/ingressRules/2147483647",
		Rule: &pb.FirewallRule{
			Priority:    math.MaxInt32,
			Action:      pb.FirewallRule_DENY,
			SourceRange: "10.0.0.0/8",
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"action", "source_range"}},
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("UpdateIngressRule(default, action+source_range) code = %v, want InvalidArgument", status.Code(err))
	}
}

func TestFirewallServer_UpdateDefaultRuleAllowsActionOnly(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()

	updated, err := srv.UpdateIngressRule(ctx, &pb.UpdateIngressRuleRequest{
		Name: "apps/test-app/firewall/ingressRules/2147483647",
		Rule: &pb.FirewallRule{
			Priority: math.MaxInt32,
			Action:   pb.FirewallRule_DENY,
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"action"}},
	})

	if err != nil {
		t.Fatalf("UpdateIngressRule(default, action) failed: %v", err)
	}
	if updated.GetAction() != pb.FirewallRule_DENY || updated.GetDescription() != "The default action." {
		t.Errorf("unexpected updated default rule: %+v", updated)
	}
}

func TestFirewallServer_DeleteDefaultRuleRejected(t *testing.T) {
	ctx := context.Background()
	srv := newTestFirewallServer()

	_, err := srv.DeleteIngressRule(ctx, &pb.DeleteIngressRuleRequest{
		Name: "apps/test-app/firewall/ingressRules/2147483647",
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("DeleteIngressRule(default) code = %v, want InvalidArgument", status.Code(err))
	}
}
