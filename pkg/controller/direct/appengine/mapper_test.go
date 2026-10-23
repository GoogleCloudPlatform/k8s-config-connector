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
	"reflect"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krmappenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/utils/ptr"
)

func TestAppEngineFirewallRuleSpecMapping(t *testing.T) {
	tests := []struct {
		name      string
		protoRule *pb.FirewallRule
		krmSpec   *krmappenginev1alpha1.AppEngineFirewallRuleSpec
	}{
		{
			name:      "nil",
			protoRule: nil,
			krmSpec:   nil,
		},
		{
			name: "full fields allow",
			protoRule: &pb.FirewallRule{
				Priority:    1000,
				Action:      pb.FirewallRule_ALLOW,
				SourceRange: "192.168.1.0/24",
				Description: "Allow local subnet",
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFirewallRuleSpec{
				ResourceID:  ptr.To("1000"),
				Action:      "ALLOW",
				SourceRange: "192.168.1.0/24",
				Description: ptr.To("Allow local subnet"),
			},
		},
		{
			name: "full fields deny",
			protoRule: &pb.FirewallRule{
				Priority:    2147483647,
				Action:      pb.FirewallRule_DENY,
				SourceRange: "*",
				Description: "Default deny",
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFirewallRuleSpec{
				ResourceID:  ptr.To("2147483647"),
				Action:      "DENY",
				SourceRange: "*",
				Description: ptr.To("Default deny"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineFirewallRuleSpec_v1alpha1_FromProto(mapCtx, tc.protoRule)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoRule) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoRule)
			}
		})
	}
}

func TestAppEngineServiceSplitTrafficSpecMapping(t *testing.T) {
	tests := []struct {
		name         string
		protoService *pb.Service
		krmSpec      *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec
	}{
		{
			name:         "nil",
			protoService: nil,
			krmSpec:      nil,
		},
		{
			name: "full fields IP shard",
			protoService: &pb.Service{
				Id: "default",
				Split: &pb.TrafficSplit{
					ShardBy: pb.TrafficSplit_IP,
					Allocations: map[string]float64{
						"v1": 0.5,
						"v2": 0.5,
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{
				ResourceID: ptr.To("default"),
				Split: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSplit{
					ShardBy: ptr.To("IP"),
					Allocations: map[string]string{
						"v1": "0.5",
						"v2": "0.5",
					},
				},
			},
		},
		{
			name: "cookie shard with 3 decimal precision",
			protoService: &pb.Service{
				Id: "backend",
				Split: &pb.TrafficSplit{
					ShardBy: pb.TrafficSplit_COOKIE,
					Allocations: map[string]float64{
						"v1": 0.8,
						"v2": 0.2,
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{
				ResourceID: ptr.To("backend"),
				Split: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSplit{
					ShardBy: ptr.To("COOKIE"),
					Allocations: map[string]string{
						"v1": "0.8",
						"v2": "0.2",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto(mapCtx, tc.protoService)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoService) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoService)
			}
		})
	}
}
