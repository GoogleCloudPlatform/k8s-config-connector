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

package compute

import (
	"context"
	"sort"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestRouterNATAdapter_RequestBuilding(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(krm.GroupVersion, &krm.ComputeRouterNAT{}, &unstructured.Unstructured{})

	tests := []struct {
		name        string
		obj         *krm.ComputeRouterNAT
		objs        []runtime.Object
		wantRouter  string
		wantProject string
		wantRegion  string
	}{
		{
			name: "short name routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "proj-1",
					},
				},
				Spec: krm.ComputeRouterNATSpec{
					Region: "region-1",
					RouterRef: krm.ComputeRouterRef{
						External: "router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
		{
			name: "canonical path routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: krm.ComputeRouterNATSpec{
					RouterRef: krm.ComputeRouterRef{
						External: "projects/proj-1/regions/region-1/routers/router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
		{
			name: "full URI routerRef",
			obj: &krm.ComputeRouterNAT{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nat-1",
					Namespace: "ns-1",
				},
				Spec: krm.ComputeRouterNATSpec{
					RouterRef: krm.ComputeRouterRef{
						External: "https://www.googleapis.com/compute/v1/projects/proj-1/regions/region-1/routers/router-1",
					},
				},
			},
			wantRouter:  "router-1",
			wantProject: "proj-1",
			wantRegion:  "region-1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			fakeClient := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objs...).Build()

			ident, err := tc.obj.GetIdentity(ctx, fakeClient)
			if err != nil {
				t.Fatalf("unexpected error resolving identity: %v", err)
			}

			id := ident.(*krm.ComputeRouterNATIdentity)

			if id.Router != tc.wantRouter {
				t.Errorf("expected router %q, got %q", tc.wantRouter, id.Router)
			}
			if id.Project != tc.wantProject {
				t.Errorf("expected project %q, got %q", tc.wantProject, id.Project)
			}
			if id.Region != tc.wantRegion {
				t.Errorf("expected region %q, got %q", tc.wantRegion, id.Region)
			}
		})
	}
}

func TestCompareComputeRouterNAT(t *testing.T) {
	tests := []struct {
		name      string
		desired   *computepb.RouterNat
		actual    *computepb.RouterNat
		wantDiff  bool
		wantPaths []string
	}{
		{
			// 1. Reference URL Prefix Mismatch
			// Desired (KRM):
			// nat_ips: ["projects/p1/regions/r1/addresses/addr-1"]
			// subnetworks: [{"name": "projects/p1/regions/r1/subnetworks/sub-1", "source_ip_ranges_to_nat": ["ALL_IP_RANGES"]}]
			// Actual (GCP Proto):
			// nat_ips: ["https://www.googleapis.com/compute/v1/projects/p1/regions/r1/addresses/addr-1"]
			// subnetworks: [{"name": "https://www.googleapis.com/compute/v1/projects/p1/regions/r1/subnetworks/sub-1", "source_ip_ranges_to_nat": ["ALL_IP_RANGES"]}]
			name: "Reference URL Prefix Mismatch",
			desired: &computepb.RouterNat{
				NatIps: []string{"projects/p1/regions/r1/addresses/addr-1"},
				Subnetworks: []*computepb.RouterNatSubnetworkToNat{
					{
						Name:                proto.String("projects/p1/regions/r1/subnetworks/sub-1"),
						SourceIpRangesToNat: []string{"ALL_IP_RANGES"},
					},
				},
			},
			actual: &computepb.RouterNat{
				NatIps: []string{"https://www.googleapis.com/compute/v1/projects/p1/regions/r1/addresses/addr-1"},
				Subnetworks: []*computepb.RouterNatSubnetworkToNat{
					{
						Name:                proto.String("https://www.googleapis.com/compute/v1/projects/p1/regions/r1/subnetworks/sub-1"),
						SourceIpRangesToNat: []string{"ALL_IP_RANGES"},
					},
				},
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			// 2. Missing Server Defaults for Unset Fields
			// Desired (KRM):
			// type: nil
			// endpoint_types: nil
			// nat_ip_allocate_option: "AUTO_ONLY"
			// Actual (GCP Proto):
			// type: "PUBLIC"
			// endpoint_types: ["ENDPOINT_TYPE_VM"]
			// nat_ip_allocate_option: "AUTO_ONLY"
			name: "Missing Server Defaults for Unset Fields",
			desired: &computepb.RouterNat{
				NatIpAllocateOption: proto.String("AUTO_ONLY"),
			},
			actual: &computepb.RouterNat{
				Type:                proto.String("PUBLIC"),
				EndpointTypes:       []string{"ENDPOINT_TYPE_VM"},
				NatIpAllocateOption: proto.String("AUTO_ONLY"),
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			// 3. Wrong Default for enable_endpoint_independent_mapping
			// Desired (KRM):
			// enable_endpoint_independent_mapping: nil (omitted)
			// Actual (GCP Proto):
			// enable_endpoint_independent_mapping: false
			name:    "Wrong Default for enable_endpoint_independent_mapping",
			desired: &computepb.RouterNat{},
			actual: &computepb.RouterNat{
				EnableEndpointIndependentMapping: proto.Bool(false),
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			// 4. Wrong Default for enable_dynamic_port_allocation on Private NAT
			// Desired (KRM):
			// type: "PRIVATE"
			// enable_dynamic_port_allocation: nil (omitted)
			// Actual (GCP Proto):
			// type: "PRIVATE"
			// enable_dynamic_port_allocation: true
			name: "Wrong Default for enable_dynamic_port_allocation on Private NAT",
			desired: &computepb.RouterNat{
				Type: proto.String("PRIVATE"),
			},
			actual: &computepb.RouterNat{
				Type:                        proto.String("PRIVATE"),
				EnableDynamicPortAllocation: proto.Bool(true),
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			// 5. Terraform Set vs Direct Slice Ordering
			// Desired (KRM):
			// nat_ips: ["projects/.../addr-1", "projects/.../addr-2"]
			// Actual (GCP Proto):
			// nat_ips: ["projects/.../addr-2", "projects/.../addr-1"]
			name: "Terraform Set vs Direct Slice Ordering",
			desired: &computepb.RouterNat{
				NatIps: []string{
					"projects/p1/regions/r1/addresses/addr-1",
					"projects/p1/regions/r1/addresses/addr-2",
				},
			},
			actual: &computepb.RouterNat{
				NatIps: []string{
					"projects/p1/regions/r1/addresses/addr-2",
					"projects/p1/regions/r1/addresses/addr-1",
				},
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			// 6. Positive Control: Genuine Field Modification
			// Desired (KRM):
			// min_ports_per_vm: 128
			// Actual (GCP Proto):
			// min_ports_per_vm: 64
			name: "Positive Control: Genuine Field Modification",
			desired: &computepb.RouterNat{
				MinPortsPerVm: proto.Int32(128),
			},
			actual: &computepb.RouterNat{
				MinPortsPerVm: proto.Int32(64),
			},
			wantDiff:  true,
			wantPaths: []string{"min_ports_per_vm"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			diff, updateMask, err := compareComputeRouterNAT(ctx, tc.actual, tc.desired)
			if err != nil {
				t.Fatalf("unexpected error comparing: %v", err)
			}

			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("expected HasDiff() = %v, got %v", tc.wantDiff, diff.HasDiff())
			}

			gotPaths := updateMask.GetPaths()
			sort.Strings(gotPaths)
			sort.Strings(tc.wantPaths)

			if !cmp.Equal(gotPaths, tc.wantPaths) {
				t.Errorf("unexpected diff paths (-want +got):\n%s", cmp.Diff(tc.wantPaths, gotPaths))
			}
		})
	}
}
