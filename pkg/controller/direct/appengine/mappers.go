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

// AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto converts protobuf Service to KRM AppEngineServiceSplitTrafficSpec.
// This is handcoded because Service.id maps to KRM ResourceID, and migrateTraffic / project are KRM-specific fields.
func AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{}
	out.Split = AppEngineServiceSplitTrafficSplit_v1alpha1_FromProto(mapCtx, in.GetSplit())
	if in.GetId() != "" {
		out.ResourceID = direct.LazyPtr(in.GetId())
	}
	return out
}

// AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto converts KRM AppEngineServiceSplitTrafficSpec to protobuf Service.
// This is handcoded because ResourceID maps to Service.id, and migrateTraffic / project are KRM-specific fields.
func AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	out.Split = AppEngineServiceSplitTrafficSplit_v1alpha1_ToProto(mapCtx, in.Split)
	if in.ResourceID != nil {
		out.Id = *in.ResourceID
	}
	return out
}

// Allocations_FromProto converts protobuf float64 traffic allocations to KRM string representations.
// This is handcoded because in protobuf allocations is map[string]double whereas in the existing CRD it is map[string]string.
func Allocations_FromProto(mapCtx *direct.MapContext, in map[string]float64) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return out
}

// Allocations_ToProto converts KRM string traffic allocations to protobuf float64 representations.
// This is handcoded because in protobuf allocations is map[string]double whereas in the existing CRD it is map[string]string.
func Allocations_ToProto(mapCtx *direct.MapContext, in map[string]string) map[string]float64 {
	if in == nil {
		return nil
	}
	out := make(map[string]float64, len(in))
	for k, v := range in {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse allocation %q for version %q as float64: %v", v, k, err)
			continue
		}
		out[k] = val
	}
	return out
}
