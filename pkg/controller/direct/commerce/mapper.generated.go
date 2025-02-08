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

package commerce

import (
	pb "cloud.google.com/go/commerce/consumer/procurement/apiv1/procurementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerce/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func CommerceOrderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Order) *krm.CommerceOrderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CommerceOrderObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func CommerceOrderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CommerceOrderObservedState) *pb.Order {
	if in == nil {
		return nil
	}
	out := &pb.Order{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func CommerceOrderSpec_FromProto(mapCtx *direct.MapContext, in *pb.Order) *krm.CommerceOrderSpec {
	if in == nil {
		return nil
	}
	out := &krm.CommerceOrderSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func CommerceOrderSpec_ToProto(mapCtx *direct.MapContext, in *krm.CommerceOrderSpec) *pb.Order {
	if in == nil {
		return nil
	}
	out := &pb.Order{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func LineItem_FromProto(mapCtx *direct.MapContext, in *pb.LineItem) *krm.LineItem {
	if in == nil {
		return nil
	}
	out := &krm.LineItem{}
	// MISSING: LineItemID
	// MISSING: LineItemInfo
	// MISSING: PendingChange
	// MISSING: ChangeHistory
	return out
}
func LineItem_ToProto(mapCtx *direct.MapContext, in *krm.LineItem) *pb.LineItem {
	if in == nil {
		return nil
	}
	out := &pb.LineItem{}
	// MISSING: LineItemID
	// MISSING: LineItemInfo
	// MISSING: PendingChange
	// MISSING: ChangeHistory
	return out
}
func LineItemChange_FromProto(mapCtx *direct.MapContext, in *pb.LineItemChange) *krm.LineItemChange {
	if in == nil {
		return nil
	}
	out := &krm.LineItemChange{}
	// MISSING: ChangeID
	out.ChangeType = direct.Enum_FromProto(mapCtx, in.GetChangeType())
	// MISSING: OldLineItemInfo
	out.NewLineItemInfo = LineItemInfo_FromProto(mapCtx, in.GetNewLineItemInfo())
	// MISSING: ChangeState
	// MISSING: StateReason
	// MISSING: ChangeStateReasonType
	// MISSING: ChangeEffectiveTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LineItemChange_ToProto(mapCtx *direct.MapContext, in *krm.LineItemChange) *pb.LineItemChange {
	if in == nil {
		return nil
	}
	out := &pb.LineItemChange{}
	// MISSING: ChangeID
	out.ChangeType = direct.Enum_ToProto[pb.LineItemChangeType](mapCtx, in.ChangeType)
	// MISSING: OldLineItemInfo
	out.NewLineItemInfo = LineItemInfo_ToProto(mapCtx, in.NewLineItemInfo)
	// MISSING: ChangeState
	// MISSING: StateReason
	// MISSING: ChangeStateReasonType
	// MISSING: ChangeEffectiveTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LineItemChangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LineItemChange) *krm.LineItemChangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LineItemChangeObservedState{}
	out.ChangeID = direct.LazyPtr(in.GetChangeId())
	// MISSING: ChangeType
	out.OldLineItemInfo = LineItemInfo_FromProto(mapCtx, in.GetOldLineItemInfo())
	// MISSING: NewLineItemInfo
	out.ChangeState = direct.Enum_FromProto(mapCtx, in.GetChangeState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.ChangeStateReasonType = direct.Enum_FromProto(mapCtx, in.GetChangeStateReasonType())
	out.ChangeEffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetChangeEffectiveTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func LineItemChangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LineItemChangeObservedState) *pb.LineItemChange {
	if in == nil {
		return nil
	}
	out := &pb.LineItemChange{}
	out.ChangeId = direct.ValueOf(in.ChangeID)
	// MISSING: ChangeType
	out.OldLineItemInfo = LineItemInfo_ToProto(mapCtx, in.OldLineItemInfo)
	// MISSING: NewLineItemInfo
	out.ChangeState = direct.Enum_ToProto[pb.LineItemChangeState](mapCtx, in.ChangeState)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.ChangeStateReasonType = direct.Enum_ToProto[pb.LineItemChangeStateReasonType](mapCtx, in.ChangeStateReasonType)
	out.ChangeEffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.ChangeEffectiveTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func LineItemInfo_FromProto(mapCtx *direct.MapContext, in *pb.LineItemInfo) *krm.LineItemInfo {
	if in == nil {
		return nil
	}
	out := &krm.LineItemInfo{}
	out.Offer = direct.LazyPtr(in.GetOffer())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Parameter_FromProto)
	// MISSING: Subscription
	return out
}
func LineItemInfo_ToProto(mapCtx *direct.MapContext, in *krm.LineItemInfo) *pb.LineItemInfo {
	if in == nil {
		return nil
	}
	out := &pb.LineItemInfo{}
	out.Offer = direct.ValueOf(in.Offer)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Parameter_ToProto)
	// MISSING: Subscription
	return out
}
func LineItemInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LineItemInfo) *krm.LineItemInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LineItemInfoObservedState{}
	// MISSING: Offer
	// MISSING: Parameters
	out.Subscription = Subscription_FromProto(mapCtx, in.GetSubscription())
	return out
}
func LineItemInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LineItemInfoObservedState) *pb.LineItemInfo {
	if in == nil {
		return nil
	}
	out := &pb.LineItemInfo{}
	// MISSING: Offer
	// MISSING: Parameters
	out.Subscription = Subscription_ToProto(mapCtx, in.Subscription)
	return out
}
func LineItemObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LineItem) *krm.LineItemObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LineItemObservedState{}
	out.LineItemID = direct.LazyPtr(in.GetLineItemId())
	out.LineItemInfo = LineItemInfo_FromProto(mapCtx, in.GetLineItemInfo())
	out.PendingChange = LineItemChange_FromProto(mapCtx, in.GetPendingChange())
	out.ChangeHistory = direct.Slice_FromProto(mapCtx, in.ChangeHistory, LineItemChange_FromProto)
	return out
}
func LineItemObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LineItemObservedState) *pb.LineItem {
	if in == nil {
		return nil
	}
	out := &pb.LineItem{}
	out.LineItemId = direct.ValueOf(in.LineItemID)
	out.LineItemInfo = LineItemInfo_ToProto(mapCtx, in.LineItemInfo)
	out.PendingChange = LineItemChange_ToProto(mapCtx, in.PendingChange)
	out.ChangeHistory = direct.Slice_ToProto(mapCtx, in.ChangeHistory, LineItemChange_ToProto)
	return out
}
func Order_FromProto(mapCtx *direct.MapContext, in *pb.Order) *krm.Order {
	if in == nil {
		return nil
	}
	out := &krm.Order{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Order_ToProto(mapCtx *direct.MapContext, in *krm.Order) *pb.Order {
	if in == nil {
		return nil
	}
	out := &pb.Order{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: LineItems
	// MISSING: CancelledLineItems
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func OrderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Order) *krm.OrderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.LineItems = direct.Slice_FromProto(mapCtx, in.LineItems, LineItem_FromProto)
	out.CancelledLineItems = direct.Slice_FromProto(mapCtx, in.CancelledLineItems, LineItem_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	return out
}
func OrderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrderObservedState) *pb.Order {
	if in == nil {
		return nil
	}
	out := &pb.Order{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.LineItems = direct.Slice_ToProto(mapCtx, in.LineItems, LineItem_ToProto)
	out.CancelledLineItems = direct.Slice_ToProto(mapCtx, in.CancelledLineItems, LineItem_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	return out
}
func Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Parameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = Parameter_Value_FromProto(mapCtx, in.GetValue())
	return out
}
func Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Parameter) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = Parameter_Value_ToProto(mapCtx, in.Value)
	return out
}
func Parameter_Value_FromProto(mapCtx *direct.MapContext, in *pb.Parameter_Value) *krm.Parameter_Value {
	if in == nil {
		return nil
	}
	out := &krm.Parameter_Value{}
	out.Int64Value = direct.LazyPtr(in.GetInt64Value())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	return out
}
func Parameter_Value_ToProto(mapCtx *direct.MapContext, in *krm.Parameter_Value) *pb.Parameter_Value {
	if in == nil {
		return nil
	}
	out := &pb.Parameter_Value{}
	if oneof := Parameter_Value_Int64Value_ToProto(mapCtx, in.Int64Value); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Parameter_Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Parameter_Value_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Kind = oneof
	}
	return out
}
func Subscription_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.Subscription {
	if in == nil {
		return nil
	}
	out := &krm.Subscription{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.AutoRenewalEnabled = direct.LazyPtr(in.GetAutoRenewalEnabled())
	return out
}
func Subscription_ToProto(mapCtx *direct.MapContext, in *krm.Subscription) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.AutoRenewalEnabled = direct.ValueOf(in.AutoRenewalEnabled)
	return out
}
