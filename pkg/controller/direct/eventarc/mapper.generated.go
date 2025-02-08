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

package eventarc

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EventType_FromProto(mapCtx *direct.MapContext, in *pb.EventType) *krm.EventType {
	if in == nil {
		return nil
	}
	out := &krm.EventType{}
	// MISSING: Type
	// MISSING: Description
	// MISSING: FilteringAttributes
	// MISSING: EventSchemaURI
	return out
}
func EventType_ToProto(mapCtx *direct.MapContext, in *krm.EventType) *pb.EventType {
	if in == nil {
		return nil
	}
	out := &pb.EventType{}
	// MISSING: Type
	// MISSING: Description
	// MISSING: FilteringAttributes
	// MISSING: EventSchemaURI
	return out
}
func EventTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventType) *krm.EventTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventTypeObservedState{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.FilteringAttributes = direct.Slice_FromProto(mapCtx, in.FilteringAttributes, FilteringAttribute_FromProto)
	out.EventSchemaURI = direct.LazyPtr(in.GetEventSchemaUri())
	return out
}
func EventTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventTypeObservedState) *pb.EventType {
	if in == nil {
		return nil
	}
	out := &pb.EventType{}
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.FilteringAttributes = direct.Slice_ToProto(mapCtx, in.FilteringAttributes, FilteringAttribute_ToProto)
	out.EventSchemaUri = direct.ValueOf(in.EventSchemaURI)
	return out
}
func EventarcProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.EventarcProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcProviderObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func EventarcProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcProviderObservedState) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func EventarcProviderSpec_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.EventarcProviderSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcProviderSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func EventarcProviderSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcProviderSpec) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func FilteringAttribute_FromProto(mapCtx *direct.MapContext, in *pb.FilteringAttribute) *krm.FilteringAttribute {
	if in == nil {
		return nil
	}
	out := &krm.FilteringAttribute{}
	// MISSING: Attribute
	// MISSING: Description
	// MISSING: Required
	// MISSING: PathPatternSupported
	return out
}
func FilteringAttribute_ToProto(mapCtx *direct.MapContext, in *krm.FilteringAttribute) *pb.FilteringAttribute {
	if in == nil {
		return nil
	}
	out := &pb.FilteringAttribute{}
	// MISSING: Attribute
	// MISSING: Description
	// MISSING: Required
	// MISSING: PathPatternSupported
	return out
}
func FilteringAttributeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FilteringAttribute) *krm.FilteringAttributeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FilteringAttributeObservedState{}
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Required = direct.LazyPtr(in.GetRequired())
	out.PathPatternSupported = direct.LazyPtr(in.GetPathPatternSupported())
	return out
}
func FilteringAttributeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FilteringAttributeObservedState) *pb.FilteringAttribute {
	if in == nil {
		return nil
	}
	out := &pb.FilteringAttribute{}
	out.Attribute = direct.ValueOf(in.Attribute)
	out.Description = direct.ValueOf(in.Description)
	out.Required = direct.ValueOf(in.Required)
	out.PathPatternSupported = direct.ValueOf(in.PathPatternSupported)
	return out
}
func Provider_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.Provider {
	if in == nil {
		return nil
	}
	out := &krm.Provider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func Provider_ToProto(mapCtx *direct.MapContext, in *krm.Provider) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: EventTypes
	return out
}
func ProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Provider) *krm.ProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProviderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EventTypes = direct.Slice_FromProto(mapCtx, in.EventTypes, EventType_FromProto)
	return out
}
func ProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProviderObservedState) *pb.Provider {
	if in == nil {
		return nil
	}
	out := &pb.Provider{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EventTypes = direct.Slice_ToProto(mapCtx, in.EventTypes, EventType_ToProto)
	return out
}
