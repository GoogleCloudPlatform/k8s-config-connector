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
// krm.group: networkservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkservices.v1

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EndpointMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher) *krm.EndpointMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher{}
	out.MetadataLabelMatcher = EndpointMatcher_MetadataLabelMatcher_FromProto(mapCtx, in.GetMetadataLabelMatcher())
	return out
}
func EndpointMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher) *pb.EndpointMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher{}
	if oneof := EndpointMatcher_MetadataLabelMatcher_ToProto(mapCtx, in.MetadataLabelMatcher); oneof != nil {
		out.MatcherType = &pb.EndpointMatcher_MetadataLabelMatcher_{MetadataLabelMatcher: oneof}
	}
	return out
}
func EndpointMatcher_MetadataLabelMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher_MetadataLabelMatcher) *krm.EndpointMatcher_MetadataLabelMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher_MetadataLabelMatcher{}
	out.MetadataLabelMatchCriteria = direct.Enum_FromProto(mapCtx, in.GetMetadataLabelMatchCriteria())
	out.MetadataLabels = direct.Slice_FromProto(mapCtx, in.MetadataLabels, EndpointMatcher_MetadataLabelMatcher_MetadataLabels_FromProto)
	return out
}
func EndpointMatcher_MetadataLabelMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher_MetadataLabelMatcher) *pb.EndpointMatcher_MetadataLabelMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher_MetadataLabelMatcher{}
	out.MetadataLabelMatchCriteria = direct.Enum_ToProto[pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria](mapCtx, in.MetadataLabelMatchCriteria)
	out.MetadataLabels = direct.Slice_ToProto(mapCtx, in.MetadataLabels, EndpointMatcher_MetadataLabelMatcher_MetadataLabels_ToProto)
	return out
}
func EndpointMatcher_MetadataLabelMatcher_MetadataLabels_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels) *krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels{}
	out.LabelName = direct.LazyPtr(in.GetLabelName())
	out.LabelValue = direct.LazyPtr(in.GetLabelValue())
	return out
}
func EndpointMatcher_MetadataLabelMatcher_MetadataLabels_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels) *pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels{}
	out.LabelName = direct.ValueOf(in.LabelName)
	out.LabelValue = direct.ValueOf(in.LabelValue)
	return out
}
func TrafficPortSelector_FromProto(mapCtx *direct.MapContext, in *pb.TrafficPortSelector) *krm.TrafficPortSelector {
	if in == nil {
		return nil
	}
	out := &krm.TrafficPortSelector{}
	out.Ports = in.Ports
	return out
}
func TrafficPortSelector_ToProto(mapCtx *direct.MapContext, in *krm.TrafficPortSelector) *pb.TrafficPortSelector {
	if in == nil {
		return nil
	}
	out := &pb.TrafficPortSelector{}
	out.Ports = in.Ports
	return out
}
