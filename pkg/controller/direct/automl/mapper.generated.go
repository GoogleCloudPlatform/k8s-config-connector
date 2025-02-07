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

package automl

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ExampleCount = direct.LazyPtr(in.GetExampleCount())
	return out
}
func AnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ExampleCount = direct.ValueOf(in.ExampleCount)
	return out
}
func AutomlAnnotationSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AutomlAnnotationSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlAnnotationSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlAnnotationSpecObservedState) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AutomlAnnotationSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlAnnotationSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
func AutomlAnnotationSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlAnnotationSpecSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ExampleCount
	return out
}
