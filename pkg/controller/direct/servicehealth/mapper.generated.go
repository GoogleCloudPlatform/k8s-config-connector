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

package servicehealth

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicehealth/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/servicehealth/apiv1/servicehealthpb"
)
func EventImpact_FromProto(mapCtx *direct.MapContext, in *pb.EventImpact) *krm.EventImpact {
	if in == nil {
		return nil
	}
	out := &krm.EventImpact{}
	out.Product = Product_FromProto(mapCtx, in.GetProduct())
	out.Location = Location_FromProto(mapCtx, in.GetLocation())
	return out
}
func EventImpact_ToProto(mapCtx *direct.MapContext, in *krm.EventImpact) *pb.EventImpact {
	if in == nil {
		return nil
	}
	out := &pb.EventImpact{}
	out.Product = Product_ToProto(mapCtx, in.Product)
	out.Location = Location_ToProto(mapCtx, in.Location)
	return out
}
func EventUpdate_FromProto(mapCtx *direct.MapContext, in *pb.EventUpdate) *krm.EventUpdate {
	if in == nil {
		return nil
	}
	out := &krm.EventUpdate{}
	// MISSING: UpdateTime
	// MISSING: Title
	// MISSING: Description
	// MISSING: Symptom
	// MISSING: Workaround
	return out
}
func EventUpdate_ToProto(mapCtx *direct.MapContext, in *krm.EventUpdate) *pb.EventUpdate {
	if in == nil {
		return nil
	}
	out := &pb.EventUpdate{}
	// MISSING: UpdateTime
	// MISSING: Title
	// MISSING: Description
	// MISSING: Symptom
	// MISSING: Workaround
	return out
}
func EventUpdateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventUpdate) *krm.EventUpdateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventUpdateObservedState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Symptom = direct.LazyPtr(in.GetSymptom())
	out.Workaround = direct.LazyPtr(in.GetWorkaround())
	return out
}
func EventUpdateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventUpdateObservedState) *pb.EventUpdate {
	if in == nil {
		return nil
	}
	out := &pb.EventUpdate{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Symptom = direct.ValueOf(in.Symptom)
	out.Workaround = direct.ValueOf(in.Workaround)
	return out
}
func Location_FromProto(mapCtx *direct.MapContext, in *pb.Location) *krm.Location {
	if in == nil {
		return nil
	}
	out := &krm.Location{}
	out.LocationName = direct.LazyPtr(in.GetLocationName())
	return out
}
func Location_ToProto(mapCtx *direct.MapContext, in *krm.Location) *pb.Location {
	if in == nil {
		return nil
	}
	out := &pb.Location{}
	out.LocationName = direct.ValueOf(in.LocationName)
	return out
}
func OrganizationEvent_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationEvent) *krm.OrganizationEvent {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationEvent{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
func OrganizationEvent_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationEvent) *pb.OrganizationEvent {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationEvent{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
func OrganizationEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationEvent) *krm.OrganizationEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationEventObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Category = direct.Enum_FromProto(mapCtx, in.GetCategory())
	out.DetailedCategory = direct.Enum_FromProto(mapCtx, in.GetDetailedCategory())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DetailedState = direct.Enum_FromProto(mapCtx, in.GetDetailedState())
	out.EventImpacts = direct.Slice_FromProto(mapCtx, in.EventImpacts, EventImpact_FromProto)
	out.Updates = direct.Slice_FromProto(mapCtx, in.Updates, EventUpdate_FromProto)
	out.ParentEvent = direct.LazyPtr(in.GetParentEvent())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.NextUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextUpdateTime())
	return out
}
func OrganizationEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationEventObservedState) *pb.OrganizationEvent {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationEvent{}
	out.Name = direct.ValueOf(in.Name)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Category = direct.Enum_ToProto[pb.OrganizationEvent_EventCategory](mapCtx, in.Category)
	out.DetailedCategory = direct.Enum_ToProto[pb.OrganizationEvent_DetailedCategory](mapCtx, in.DetailedCategory)
	out.State = direct.Enum_ToProto[pb.OrganizationEvent_State](mapCtx, in.State)
	out.DetailedState = direct.Enum_ToProto[pb.OrganizationEvent_DetailedState](mapCtx, in.DetailedState)
	out.EventImpacts = direct.Slice_ToProto(mapCtx, in.EventImpacts, EventImpact_ToProto)
	out.Updates = direct.Slice_ToProto(mapCtx, in.Updates, EventUpdate_ToProto)
	out.ParentEvent = direct.ValueOf(in.ParentEvent)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.NextUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.NextUpdateTime)
	return out
}
func Product_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.Product {
	if in == nil {
		return nil
	}
	out := &krm.Product{}
	out.ProductName = direct.LazyPtr(in.GetProductName())
	out.ID = direct.LazyPtr(in.GetId())
	return out
}
func Product_ToProto(mapCtx *direct.MapContext, in *krm.Product) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	out.ProductName = direct.ValueOf(in.ProductName)
	out.Id = direct.ValueOf(in.ID)
	return out
}
func ServicehealthOrganizationEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationEvent) *krm.ServicehealthOrganizationEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServicehealthOrganizationEventObservedState{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
func ServicehealthOrganizationEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServicehealthOrganizationEventObservedState) *pb.OrganizationEvent {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationEvent{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
func ServicehealthOrganizationEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationEvent) *krm.ServicehealthOrganizationEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.ServicehealthOrganizationEventSpec{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
func ServicehealthOrganizationEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.ServicehealthOrganizationEventSpec) *pb.OrganizationEvent {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationEvent{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Category
	// MISSING: DetailedCategory
	// MISSING: State
	// MISSING: DetailedState
	// MISSING: EventImpacts
	// MISSING: Updates
	// MISSING: ParentEvent
	// MISSING: UpdateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: NextUpdateTime
	return out
}
