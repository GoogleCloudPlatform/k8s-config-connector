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

package datafusion

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datafusion/apiv1beta1/datafusionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datafusion/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatafusionNamespaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Namespace) *krm.DatafusionNamespaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatafusionNamespaceObservedState{}
	// MISSING: Name
	// MISSING: IamPolicy
	return out
}
func DatafusionNamespaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatafusionNamespaceObservedState) *pb.Namespace {
	if in == nil {
		return nil
	}
	out := &pb.Namespace{}
	// MISSING: Name
	// MISSING: IamPolicy
	return out
}
func DatafusionNamespaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Namespace) *krm.DatafusionNamespaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatafusionNamespaceSpec{}
	// MISSING: Name
	// MISSING: IamPolicy
	return out
}
func DatafusionNamespaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatafusionNamespaceSpec) *pb.Namespace {
	if in == nil {
		return nil
	}
	out := &pb.Namespace{}
	// MISSING: Name
	// MISSING: IamPolicy
	return out
}
func IAMPolicy_FromProto(mapCtx *direct.MapContext, in *pb.IAMPolicy) *krm.IAMPolicy {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicy{}
	out.Policy = Policy_FromProto(mapCtx, in.GetPolicy())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	return out
}
func IAMPolicy_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicy) *pb.IAMPolicy {
	if in == nil {
		return nil
	}
	out := &pb.IAMPolicy{}
	out.Policy = Policy_ToProto(mapCtx, in.Policy)
	out.Status = Status_ToProto(mapCtx, in.Status)
	return out
}
func Namespace_FromProto(mapCtx *direct.MapContext, in *pb.Namespace) *krm.Namespace {
	if in == nil {
		return nil
	}
	out := &krm.Namespace{}
	out.Name = direct.LazyPtr(in.GetName())
	out.IamPolicy = IAMPolicy_FromProto(mapCtx, in.GetIamPolicy())
	return out
}
func Namespace_ToProto(mapCtx *direct.MapContext, in *krm.Namespace) *pb.Namespace {
	if in == nil {
		return nil
	}
	out := &pb.Namespace{}
	out.Name = direct.ValueOf(in.Name)
	out.IamPolicy = IAMPolicy_ToProto(mapCtx, in.IamPolicy)
	return out
}
