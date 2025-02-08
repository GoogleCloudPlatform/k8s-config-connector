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

package speech

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/speech/apiv1/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CustomClass_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.CustomClass {
	if in == nil {
		return nil
	}
	out := &krm.CustomClass{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CustomClassID = direct.LazyPtr(in.GetCustomClassId())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, CustomClass_ClassItem_FromProto)
	return out
}
func CustomClass_ToProto(mapCtx *direct.MapContext, in *krm.CustomClass) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	out.Name = direct.ValueOf(in.Name)
	out.CustomClassId = direct.ValueOf(in.CustomClassID)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, CustomClass_ClassItem_ToProto)
	return out
}
func CustomClass_ClassItem_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass_ClassItem) *krm.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &krm.CustomClass_ClassItem{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func CustomClass_ClassItem_ToProto(mapCtx *direct.MapContext, in *krm.CustomClass_ClassItem) *pb.CustomClass_ClassItem {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass_ClassItem{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func SpeechCustomClassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.SpeechCustomClassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechCustomClassObservedState{}
	// MISSING: Name
	// MISSING: CustomClassID
	// MISSING: Items
	return out
}
func SpeechCustomClassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechCustomClassObservedState) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	// MISSING: CustomClassID
	// MISSING: Items
	return out
}
func SpeechCustomClassSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krm.SpeechCustomClassSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpeechCustomClassSpec{}
	// MISSING: Name
	// MISSING: CustomClassID
	// MISSING: Items
	return out
}
func SpeechCustomClassSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpeechCustomClassSpec) *pb.CustomClass {
	if in == nil {
		return nil
	}
	out := &pb.CustomClass{}
	// MISSING: Name
	// MISSING: CustomClassID
	// MISSING: Items
	return out
}
