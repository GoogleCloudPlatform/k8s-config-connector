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

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AgentValidationResult_FromProto(mapCtx *direct.MapContext, in *pb.AgentValidationResult) *krm.AgentValidationResult {
	if in == nil {
		return nil
	}
	out := &krm.AgentValidationResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FlowValidationResults = direct.Slice_FromProto(mapCtx, in.FlowValidationResults, FlowValidationResult_FromProto)
	return out
}
func AgentValidationResult_ToProto(mapCtx *direct.MapContext, in *krm.AgentValidationResult) *pb.AgentValidationResult {
	if in == nil {
		return nil
	}
	out := &pb.AgentValidationResult{}
	out.Name = direct.ValueOf(in.Name)
	out.FlowValidationResults = direct.Slice_ToProto(mapCtx, in.FlowValidationResults, FlowValidationResult_ToProto)
	return out
}
func FlowValidationResult_FromProto(mapCtx *direct.MapContext, in *pb.FlowValidationResult) *krm.FlowValidationResult {
	if in == nil {
		return nil
	}
	out := &krm.FlowValidationResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ValidationMessages = direct.Slice_FromProto(mapCtx, in.ValidationMessages, ValidationMessage_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FlowValidationResult_ToProto(mapCtx *direct.MapContext, in *krm.FlowValidationResult) *pb.FlowValidationResult {
	if in == nil {
		return nil
	}
	out := &pb.FlowValidationResult{}
	out.Name = direct.ValueOf(in.Name)
	out.ValidationMessages = direct.Slice_ToProto(mapCtx, in.ValidationMessages, ValidationMessage_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ResourceName_FromProto(mapCtx *direct.MapContext, in *pb.ResourceName) *krm.ResourceName {
	if in == nil {
		return nil
	}
	out := &krm.ResourceName{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func ResourceName_ToProto(mapCtx *direct.MapContext, in *krm.ResourceName) *pb.ResourceName {
	if in == nil {
		return nil
	}
	out := &pb.ResourceName{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func ValidationMessage_FromProto(mapCtx *direct.MapContext, in *pb.ValidationMessage) *krm.ValidationMessage {
	if in == nil {
		return nil
	}
	out := &krm.ValidationMessage{}
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	out.Resources = in.Resources
	out.ResourceNames = direct.Slice_FromProto(mapCtx, in.ResourceNames, ResourceName_FromProto)
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Detail = direct.LazyPtr(in.GetDetail())
	return out
}
func ValidationMessage_ToProto(mapCtx *direct.MapContext, in *krm.ValidationMessage) *pb.ValidationMessage {
	if in == nil {
		return nil
	}
	out := &pb.ValidationMessage{}
	out.ResourceType = direct.Enum_ToProto[pb.ValidationMessage_ResourceType](mapCtx, in.ResourceType)
	out.Resources = in.Resources
	out.ResourceNames = direct.Slice_ToProto(mapCtx, in.ResourceNames, ResourceName_ToProto)
	out.Severity = direct.Enum_ToProto[pb.ValidationMessage_Severity](mapCtx, in.Severity)
	out.Detail = direct.ValueOf(in.Detail)
	return out
}
