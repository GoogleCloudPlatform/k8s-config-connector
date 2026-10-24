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

func TestAppEngineDomainMappingSpecMapping(t *testing.T) {
	tests := []struct {
		name        string
		protoDomain *pb.DomainMapping
		krmSpec     *krmappenginev1alpha1.AppEngineDomainMappingSpec
	}{
		{
			name:        "nil",
			protoDomain: nil,
			krmSpec:     nil,
		},
		{
			name: "ssl settings automatic",
			protoDomain: &pb.DomainMapping{
				SslSettings: &pb.SslSettings{
					SslManagementType:           pb.SslSettings_AUTOMATIC,
					PendingManagedCertificateId: "cert-pending-1",
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineDomainMappingSpec{
				SSLSettings: &krmappenginev1alpha1.DomainmappingSslSettings{
					SSLManagementType:           "AUTOMATIC",
					PendingManagedCertificateID: ptr.To("cert-pending-1"),
				},
			},
		},
		{
			name: "ssl settings manual",
			protoDomain: &pb.DomainMapping{
				SslSettings: &pb.SslSettings{
					CertificateId:     "cert-12345",
					SslManagementType: pb.SslSettings_MANUAL,
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineDomainMappingSpec{
				SSLSettings: &krmappenginev1alpha1.DomainmappingSslSettings{
					CertificateID:     ptr.To("cert-12345"),
					SSLManagementType: "MANUAL",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineDomainMappingSpec_v1alpha1_FromProto(mapCtx, tc.protoDomain)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineDomainMappingSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoDomain) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoDomain)
			}
		})
	}
}

func TestAppEngineDomainMappingStatusMapping(t *testing.T) {
	tests := []struct {
		name        string
		protoDomain *pb.DomainMapping
		krmStatus   *krmappenginev1alpha1.AppEngineDomainMappingStatus
	}{
		{
			name:        "nil",
			protoDomain: nil,
			krmStatus:   nil,
		},
		{
			name: "name and resource records",
			protoDomain: &pb.DomainMapping{
				Name: "apps/my-project/domainMappings/example.com",
				ResourceRecords: []*pb.ResourceRecord{
					{
						Name:   "www",
						Rrdata: "example.com",
						Type:   pb.ResourceRecord_CNAME,
					},
					{
						Name:   "@",
						Rrdata: "192.0.2.1",
						Type:   pb.ResourceRecord_A,
					},
				},
			},
			krmStatus: &krmappenginev1alpha1.AppEngineDomainMappingStatus{
				Name: ptr.To("apps/my-project/domainMappings/example.com"),
				ResourceRecords: []krmappenginev1alpha1.DomainmappingResourceRecords{
					{
						Name:   ptr.To("www"),
						Rrdata: ptr.To("example.com"),
						Type:   ptr.To("CNAME"),
					},
					{
						Name:   ptr.To("@"),
						Rrdata: ptr.To("192.0.2.1"),
						Type:   ptr.To("A"),
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineDomainMappingStatus_v1alpha1_FromProto(mapCtx, tc.protoDomain)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmStatus) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmStatus)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineDomainMappingStatus_v1alpha1_ToProto(mapCtx, tc.krmStatus)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoDomain) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoDomain)
			}
		})
	}
}
