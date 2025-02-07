// Copyright 2025 Google LLC
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

package bigquery

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/datapolicies/apiv1beta1/datapoliciespb"
)
func DataMaskingPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DataMaskingPolicy) *krm.DataMaskingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.DataMaskingPolicy{}
	out.PredefinedExpression = direct.Enum_FromProto(mapCtx, in.GetPredefinedExpression())
	return out
}
func DataMaskingPolicy_ToProto(mapCtx *direct.MapContext, in *krm.DataMaskingPolicy) *pb.DataMaskingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataMaskingPolicy{}
	if oneof := DataMaskingPolicy_PredefinedExpression_ToProto(mapCtx, in.PredefinedExpression); oneof != nil {
		out.MaskingExpression = oneof
	}
	return out
}
func DataPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DataPolicy) *krm.DataPolicy {
	if in == nil {
		return nil
	}
	out := &krm.DataPolicy{}
	out.PolicyTag = direct.LazyPtr(in.GetPolicyTag())
	out.DataMaskingPolicy = DataMaskingPolicy_FromProto(mapCtx, in.GetDataMaskingPolicy())
	// MISSING: Name
	out.DataPolicyType = direct.Enum_FromProto(mapCtx, in.GetDataPolicyType())
	out.DataPolicyID = direct.LazyPtr(in.GetDataPolicyId())
	return out
}
func DataPolicy_ToProto(mapCtx *direct.MapContext, in *krm.DataPolicy) *pb.DataPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataPolicy{}
	if oneof := DataPolicy_PolicyTag_ToProto(mapCtx, in.PolicyTag); oneof != nil {
		out.MatchingLabel = oneof
	}
	if oneof := DataMaskingPolicy_ToProto(mapCtx, in.DataMaskingPolicy); oneof != nil {
		out.Policy = &pb.DataPolicy_DataMaskingPolicy{DataMaskingPolicy: oneof}
	}
	// MISSING: Name
	out.DataPolicyType = direct.Enum_ToProto[pb.DataPolicy_DataPolicyType](mapCtx, in.DataPolicyType)
	out.DataPolicyId = direct.ValueOf(in.DataPolicyID)
	return out
}
func DataPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataPolicy) *krm.DataPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataPolicyObservedState{}
	// MISSING: PolicyTag
	// MISSING: DataMaskingPolicy
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DataPolicyType
	// MISSING: DataPolicyID
	return out
}
func DataPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataPolicyObservedState) *pb.DataPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataPolicy{}
	// MISSING: PolicyTag
	// MISSING: DataMaskingPolicy
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DataPolicyType
	// MISSING: DataPolicyID
	return out
}
