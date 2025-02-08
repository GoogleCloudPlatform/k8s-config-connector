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

package contentwarehouse

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contentwarehouse/apiv1/contentwarehousepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
)
func AccessControlAction_FromProto(mapCtx *direct.MapContext, in *pb.AccessControlAction) *krm.AccessControlAction {
	if in == nil {
		return nil
	}
	out := &krm.AccessControlAction{}
	out.OperationType = direct.Enum_FromProto(mapCtx, in.GetOperationType())
	out.Policy = Policy_FromProto(mapCtx, in.GetPolicy())
	return out
}
func AccessControlAction_ToProto(mapCtx *direct.MapContext, in *krm.AccessControlAction) *pb.AccessControlAction {
	if in == nil {
		return nil
	}
	out := &pb.AccessControlAction{}
	out.OperationType = direct.Enum_ToProto[pb.AccessControlAction_OperationType](mapCtx, in.OperationType)
	out.Policy = Policy_ToProto(mapCtx, in.Policy)
	return out
}
func Action_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.Action {
	if in == nil {
		return nil
	}
	out := &krm.Action{}
	out.ActionID = direct.LazyPtr(in.GetActionId())
	out.AccessControl = AccessControlAction_FromProto(mapCtx, in.GetAccessControl())
	out.DataValidation = DataValidationAction_FromProto(mapCtx, in.GetDataValidation())
	out.DataUpdate = DataUpdateAction_FromProto(mapCtx, in.GetDataUpdate())
	out.AddToFolder = AddToFolderAction_FromProto(mapCtx, in.GetAddToFolder())
	out.PublishToPubSub = PublishAction_FromProto(mapCtx, in.GetPublishToPubSub())
	out.RemoveFromFolderAction = RemoveFromFolderAction_FromProto(mapCtx, in.GetRemoveFromFolderAction())
	out.DeleteDocumentAction = DeleteDocumentAction_FromProto(mapCtx, in.GetDeleteDocumentAction())
	return out
}
func Action_ToProto(mapCtx *direct.MapContext, in *krm.Action) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	out.ActionId = direct.ValueOf(in.ActionID)
	if oneof := AccessControlAction_ToProto(mapCtx, in.AccessControl); oneof != nil {
		out.Action = &pb.Action_AccessControl{AccessControl: oneof}
	}
	if oneof := DataValidationAction_ToProto(mapCtx, in.DataValidation); oneof != nil {
		out.Action = &pb.Action_DataValidation{DataValidation: oneof}
	}
	if oneof := DataUpdateAction_ToProto(mapCtx, in.DataUpdate); oneof != nil {
		out.Action = &pb.Action_DataUpdate{DataUpdate: oneof}
	}
	if oneof := AddToFolderAction_ToProto(mapCtx, in.AddToFolder); oneof != nil {
		out.Action = &pb.Action_AddToFolder{AddToFolder: oneof}
	}
	if oneof := PublishAction_ToProto(mapCtx, in.PublishToPubSub); oneof != nil {
		out.Action = &pb.Action_PublishToPubSub{PublishToPubSub: oneof}
	}
	if oneof := RemoveFromFolderAction_ToProto(mapCtx, in.RemoveFromFolderAction); oneof != nil {
		out.Action = &pb.Action_RemoveFromFolderAction{RemoveFromFolderAction: oneof}
	}
	if oneof := DeleteDocumentAction_ToProto(mapCtx, in.DeleteDocumentAction); oneof != nil {
		out.Action = &pb.Action_DeleteDocumentAction{DeleteDocumentAction: oneof}
	}
	return out
}
func AddToFolderAction_FromProto(mapCtx *direct.MapContext, in *pb.AddToFolderAction) *krm.AddToFolderAction {
	if in == nil {
		return nil
	}
	out := &krm.AddToFolderAction{}
	out.Folders = in.Folders
	return out
}
func AddToFolderAction_ToProto(mapCtx *direct.MapContext, in *krm.AddToFolderAction) *pb.AddToFolderAction {
	if in == nil {
		return nil
	}
	out := &pb.AddToFolderAction{}
	out.Folders = in.Folders
	return out
}
func ContentwarehouseRuleSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuleSet) *krm.ContentwarehouseRuleSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseRuleSetObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Source
	// MISSING: Rules
	return out
}
func ContentwarehouseRuleSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseRuleSetObservedState) *pb.RuleSet {
	if in == nil {
		return nil
	}
	out := &pb.RuleSet{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Source
	// MISSING: Rules
	return out
}
func ContentwarehouseRuleSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.RuleSet) *krm.ContentwarehouseRuleSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseRuleSetSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Source
	// MISSING: Rules
	return out
}
func ContentwarehouseRuleSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseRuleSetSpec) *pb.RuleSet {
	if in == nil {
		return nil
	}
	out := &pb.RuleSet{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Source
	// MISSING: Rules
	return out
}
func DataUpdateAction_FromProto(mapCtx *direct.MapContext, in *pb.DataUpdateAction) *krm.DataUpdateAction {
	if in == nil {
		return nil
	}
	out := &krm.DataUpdateAction{}
	out.Entries = in.Entries
	return out
}
func DataUpdateAction_ToProto(mapCtx *direct.MapContext, in *krm.DataUpdateAction) *pb.DataUpdateAction {
	if in == nil {
		return nil
	}
	out := &pb.DataUpdateAction{}
	out.Entries = in.Entries
	return out
}
func DataValidationAction_FromProto(mapCtx *direct.MapContext, in *pb.DataValidationAction) *krm.DataValidationAction {
	if in == nil {
		return nil
	}
	out := &krm.DataValidationAction{}
	out.Conditions = in.Conditions
	return out
}
func DataValidationAction_ToProto(mapCtx *direct.MapContext, in *krm.DataValidationAction) *pb.DataValidationAction {
	if in == nil {
		return nil
	}
	out := &pb.DataValidationAction{}
	out.Conditions = in.Conditions
	return out
}
func DeleteDocumentAction_FromProto(mapCtx *direct.MapContext, in *pb.DeleteDocumentAction) *krm.DeleteDocumentAction {
	if in == nil {
		return nil
	}
	out := &krm.DeleteDocumentAction{}
	out.EnableHardDelete = direct.LazyPtr(in.GetEnableHardDelete())
	return out
}
func DeleteDocumentAction_ToProto(mapCtx *direct.MapContext, in *krm.DeleteDocumentAction) *pb.DeleteDocumentAction {
	if in == nil {
		return nil
	}
	out := &pb.DeleteDocumentAction{}
	out.EnableHardDelete = direct.ValueOf(in.EnableHardDelete)
	return out
}
func PublishAction_FromProto(mapCtx *direct.MapContext, in *pb.PublishAction) *krm.PublishAction {
	if in == nil {
		return nil
	}
	out := &krm.PublishAction{}
	out.TopicID = direct.LazyPtr(in.GetTopicId())
	out.Messages = in.Messages
	return out
}
func PublishAction_ToProto(mapCtx *direct.MapContext, in *krm.PublishAction) *pb.PublishAction {
	if in == nil {
		return nil
	}
	out := &pb.PublishAction{}
	out.TopicId = direct.ValueOf(in.TopicID)
	out.Messages = in.Messages
	return out
}
func RemoveFromFolderAction_FromProto(mapCtx *direct.MapContext, in *pb.RemoveFromFolderAction) *krm.RemoveFromFolderAction {
	if in == nil {
		return nil
	}
	out := &krm.RemoveFromFolderAction{}
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Folder = direct.LazyPtr(in.GetFolder())
	return out
}
func RemoveFromFolderAction_ToProto(mapCtx *direct.MapContext, in *krm.RemoveFromFolderAction) *pb.RemoveFromFolderAction {
	if in == nil {
		return nil
	}
	out := &pb.RemoveFromFolderAction{}
	out.Condition = direct.ValueOf(in.Condition)
	out.Folder = direct.ValueOf(in.Folder)
	return out
}
func Rule_FromProto(mapCtx *direct.MapContext, in *pb.Rule) *krm.Rule {
	if in == nil {
		return nil
	}
	out := &krm.Rule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.RuleID = direct.LazyPtr(in.GetRuleId())
	out.TriggerType = direct.Enum_FromProto(mapCtx, in.GetTriggerType())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, Action_FromProto)
	return out
}
func Rule_ToProto(mapCtx *direct.MapContext, in *krm.Rule) *pb.Rule {
	if in == nil {
		return nil
	}
	out := &pb.Rule{}
	out.Description = direct.ValueOf(in.Description)
	out.RuleId = direct.ValueOf(in.RuleID)
	out.TriggerType = direct.Enum_ToProto[pb.Rule_TriggerType](mapCtx, in.TriggerType)
	out.Condition = direct.ValueOf(in.Condition)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, Action_ToProto)
	return out
}
func RuleSet_FromProto(mapCtx *direct.MapContext, in *pb.RuleSet) *krm.RuleSet {
	if in == nil {
		return nil
	}
	out := &krm.RuleSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Source = direct.LazyPtr(in.GetSource())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, Rule_FromProto)
	return out
}
func RuleSet_ToProto(mapCtx *direct.MapContext, in *krm.RuleSet) *pb.RuleSet {
	if in == nil {
		return nil
	}
	out := &pb.RuleSet{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Source = direct.ValueOf(in.Source)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, Rule_ToProto)
	return out
}
