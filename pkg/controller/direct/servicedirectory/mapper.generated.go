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

package servicedirectory

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Namespace_FromProto(mapCtx *direct.MapContext, in *pb.Namespace) *krm.Namespace {
	if in == nil {
		return nil
	}
	out := &krm.Namespace{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	return out
}
func Namespace_ToProto(mapCtx *direct.MapContext, in *krm.Namespace) *pb.Namespace {
	if in == nil {
		return nil
	}
	out := &pb.Namespace{}
	out.Name = direct.ValueOf(in.Name)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	return out
}
func NamespaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Namespace) *krm.NamespaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NamespaceObservedState{}
	// MISSING: Name
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func NamespaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NamespaceObservedState) *pb.Namespace {
	if in == nil {
		return nil
	}
	out := &pb.Namespace{}
	// MISSING: Name
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
