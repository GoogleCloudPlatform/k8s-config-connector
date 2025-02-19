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

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ApihubAttributeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attribute) *krm.ApihubAttributeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubAttributeObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubAttributeObservedState) *pb.Attribute {
	if in == nil {
		return nil
	}
	out := &pb.Attribute{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Attribute) *krm.ApihubAttributeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubAttributeSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubAttributeSpec) *pb.Attribute {
	if in == nil {
		return nil
	}
	out := &pb.Attribute{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Attribute_AllowedValue_FromProto(mapCtx *direct.MapContext, in *pb.Attribute_AllowedValue) *krm.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &krm.Attribute_AllowedValue{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Immutable = direct.LazyPtr(in.GetImmutable())
	return out
}
func Attribute_AllowedValue_ToProto(mapCtx *direct.MapContext, in *krm.Attribute_AllowedValue) *pb.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &pb.Attribute_AllowedValue{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Immutable = direct.ValueOf(in.Immutable)
	return out
}
