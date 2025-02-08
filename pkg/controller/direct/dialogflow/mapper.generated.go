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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Context_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.Context {
	if in == nil {
		return nil
	}
	out := &krm.Context{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LifespanCount = direct.LazyPtr(in.GetLifespanCount())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func Context_ToProto(mapCtx *direct.MapContext, in *krm.Context) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	out.Name = direct.ValueOf(in.Name)
	out.LifespanCount = direct.ValueOf(in.LifespanCount)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	return out
}
func DialogflowContextObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.DialogflowContextObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowContextObservedState{}
	// MISSING: Name
	// MISSING: LifespanCount
	// MISSING: Parameters
	return out
}
func DialogflowContextObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowContextObservedState) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	// MISSING: Name
	// MISSING: LifespanCount
	// MISSING: Parameters
	return out
}
func DialogflowContextSpec_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.DialogflowContextSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowContextSpec{}
	// MISSING: Name
	// MISSING: LifespanCount
	// MISSING: Parameters
	return out
}
func DialogflowContextSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowContextSpec) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	// MISSING: Name
	// MISSING: LifespanCount
	// MISSING: Parameters
	return out
}
