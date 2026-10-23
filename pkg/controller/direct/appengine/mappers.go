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
	"strconv"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krmappenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// AppEngineFirewallRuleSpec_v1alpha1_FromProto converts the protobuf FirewallRule to the KRM AppEngineFirewallRuleSpec.
// This is handcoded because Action is represented as an enum in protobuf but as a required non-pointer string in the existing CRD schema,
// SourceRange is a required non-pointer string in KRM, and Priority is an int32 in protobuf while represented by ResourceID (*string) in KRM.
func AppEngineFirewallRuleSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallRule) *krmappenginev1alpha1.AppEngineFirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineFirewallRuleSpec{}
	out.Action = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetAction()))
	out.SourceRange = in.GetSourceRange()
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetPriority() != 0 {
		priorityStr := strconv.FormatInt(int64(in.GetPriority()), 10)
		out.ResourceID = &priorityStr
	}
	return out
}

// AppEngineFirewallRuleSpec_v1alpha1_ToProto converts the KRM AppEngineFirewallRuleSpec to the protobuf FirewallRule.
// This is handcoded because Action is represented as an enum in protobuf but as a required non-pointer string in the existing CRD schema,
// SourceRange is a required non-pointer string in KRM, and Priority is an int32 in protobuf while represented by ResourceID (*string) in KRM.
func AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineFirewallRuleSpec) *pb.FirewallRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallRule{}
	out.Action = direct.Enum_ToProto[pb.FirewallRule_Action](mapCtx, &in.Action)
	out.SourceRange = in.SourceRange
	out.Description = direct.ValueOf(in.Description)
	if in.ResourceID != nil {
		priorityVal, err := strconv.ParseInt(*in.ResourceID, 10, 32)
		if err != nil {
			mapCtx.Errorf("cannot parse resourceID %q as int32 priority: %v", *in.ResourceID, err)
		} else {
			out.Priority = int32(priorityVal)
		}
	}
	return out
}

// DomainmappingSslSettings_v1alpha1_FromProto converts the protobuf SslSettings to the KRM DomainmappingSslSettings.
// This is handcoded because SSLManagementType is represented as an enum in protobuf but as a required non-pointer string in KRM.
func DomainmappingSslSettings_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SslSettings) *krmappenginev1alpha1.DomainmappingSslSettings {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.DomainmappingSslSettings{}
	out.CertificateID = direct.LazyPtr(in.GetCertificateId())
	out.SSLManagementType = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetSslManagementType()))
	out.PendingManagedCertificateID = direct.LazyPtr(in.GetPendingManagedCertificateId())
	return out
}

// DomainmappingSslSettings_v1alpha1_ToProto converts the KRM DomainmappingSslSettings to the protobuf SslSettings.
// This is handcoded because SSLManagementType is represented as an enum in protobuf but as a required non-pointer string in KRM.
func DomainmappingSslSettings_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.DomainmappingSslSettings) *pb.SslSettings {
	if in == nil {
		return nil
	}
	out := &pb.SslSettings{}
	out.CertificateId = direct.ValueOf(in.CertificateID)
	out.SslManagementType = direct.Enum_ToProto[pb.SslSettings_SslManagementType](mapCtx, &in.SSLManagementType)
	out.PendingManagedCertificateId = direct.ValueOf(in.PendingManagedCertificateID)
	return out
}
