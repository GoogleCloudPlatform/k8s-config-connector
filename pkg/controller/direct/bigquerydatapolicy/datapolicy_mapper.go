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

// +generated:mapper
// krm.group: bigquerydatapolicy.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.bigquery.datapolicies.v1beta1

package bigquerydatapolicy

import (
	pb "cloud.google.com/go/bigquery/datapolicies/apiv1beta1/datapoliciespb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatapolicy/v1alpha1"
	datacalogref "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryDataPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataPolicy) *krmv1alpha1.BigQueryDataPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDataPolicyObservedState{}
	return out
}
func BigQueryDataPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDataPolicyObservedState) *pb.DataPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataPolicy{}
	return out
}
func BigQueryDataPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.DataPolicy) *krmv1alpha1.BigQueryDataPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDataPolicySpec{}
	if in.GetPolicyTag() != "" {
		out.PolicyTagRef = &datacalogref.PolicyTagRef{External: in.GetPolicyTag()}
	}
	out.DataMaskingPolicy = DataMaskingPolicy_FromProto(mapCtx, in.GetDataMaskingPolicy())
	out.DataPolicyType = direct.Enum_FromProto(mapCtx, in.GetDataPolicyType())
	out.DataPolicyID = direct.LazyPtr(in.GetDataPolicyId())
	return out
}
func BigQueryDataPolicySpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDataPolicySpec) *pb.DataPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataPolicy{}
	if in.PolicyTagRef != nil {
		out.MatchingLabel = &pb.DataPolicy_PolicyTag{PolicyTag: in.PolicyTagRef.External}
	}
	if oneof := DataMaskingPolicy_ToProto(mapCtx, in.DataMaskingPolicy); oneof != nil {
		out.Policy = &pb.DataPolicy_DataMaskingPolicy{DataMaskingPolicy: oneof}
	}
	out.DataPolicyType = direct.Enum_ToProto[pb.DataPolicy_DataPolicyType](mapCtx, in.DataPolicyType)
	out.DataPolicyId = direct.ValueOf(in.DataPolicyID)
	return out
}
func DataMaskingPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DataMaskingPolicy) *krmv1alpha1.DataMaskingPolicy {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataMaskingPolicy{}
	out.PredefinedExpression = direct.Enum_FromProto(mapCtx, in.GetPredefinedExpression())
	return out
}
func DataMaskingPolicy_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataMaskingPolicy) *pb.DataMaskingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DataMaskingPolicy{}
	if oneof := DataMaskingPolicy_PredefinedExpression_ToProto(mapCtx, in.PredefinedExpression); oneof != nil {
		out.MaskingExpression = oneof
	}
	return out
}

func DataMaskingPolicy_PredefinedExpression_ToProto(mapCtx *direct.MapContext, in *string) *pb.DataMaskingPolicy_PredefinedExpression_ {
	if in == nil {
		return nil
	}
	out := &pb.DataMaskingPolicy_PredefinedExpression_{PredefinedExpression: direct.Enum_ToProto[pb.DataMaskingPolicy_PredefinedExpression](mapCtx, in)}
	return out
}
