// Copyright 2026 Google LLC
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

package gkehub

import (
	gkehubapi "google.golang.org/api/gkehub/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/gkehub/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEHubScopeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Scope) *krm.GKEHubScopeSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeSpec{}
	out.NamespaceLabels = in.NamespaceLabels

	// Identity mapping
	id := &krm.GKEHubScopeIdentity{}
	if err := id.FromExternal(in.Name); err == nil {
		out.ProjectRef = refs.ProjectRef{External: "projects/" + id.ProjectID}
		out.Location = direct.LazyPtr(id.Location)
		out.ResourceID = direct.LazyPtr(id.ID())
	}

	return out
}

func GKEHubScopeSpec_ToProto(mapCtx *direct.MapContext, in *krm.GKEHubScopeSpec) *pb.Scope {
	if in == nil {
		return nil
	}
	out := &pb.Scope{}
	out.NamespaceLabels = in.NamespaceLabels
	return out
}

func GKEHubScopeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Scope) *krm.GKEHubScopeStatus {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeStatus{}
	out.Uid = direct.LazyPtr(in.Uid)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.DeleteTime)
	if in.State != nil {
		out.State = &krm.GKEHubScopeStateStatus{
			Code: direct.LazyPtr(in.State.Code.String()),
		}
	}
	return out
}

func GKEHubScopeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEHubScopeStatus) *pb.Scope {
	if in == nil {
		return nil
	}
	out := &pb.Scope{}
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	if in.State != nil {
		val := pb.ScopeLifecycleState_Code_value[direct.ValueOf(in.State.Code)]
		out.State = &pb.ScopeLifecycleState{
			Code: pb.ScopeLifecycleState_Code(val),
		}
	}
	return out
}

// --- Discovery API Mappings (Helper for Controller) ---

func GKEHubScopeSpec_FromAPI(mapCtx *direct.MapContext, in *gkehubapi.Scope) *krm.GKEHubScopeSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeSpec{}
	out.NamespaceLabels = in.NamespaceLabels
	return out
}

func GKEHubScopeObservedState_FromAPI(mapCtx *direct.MapContext, in *gkehubapi.Scope) *krm.GKEHubScopeStatus {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeStatus{}
	out.Uid = direct.LazyPtr(in.Uid)
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	out.DeleteTime = direct.LazyPtr(in.DeleteTime)
	if in.State != nil {
		out.State = &krm.GKEHubScopeStateStatus{
			Code: direct.LazyPtr(in.State.Code),
		}
	}
	return out
}

// --- Legacy Helpers ---

func GKEHubScopeSpecKRMtoAPI(mapCtx *direct.MapContext, r *krm.GKEHubScopeSpec) *gkehubapi.Scope {
	pbObj := GKEHubScopeSpec_ToProto(mapCtx, r)
	apiObj := &gkehubapi.Scope{}
	if err := Convert_v1_Scope_pb_to_api(pbObj, apiObj); err != nil {
		mapCtx.Errorf("error converting pb to api: %v", err)
	}
	return apiObj
}

func GKEHubScopeSpecAPIToKRM(mapCtx *direct.MapContext, r *gkehubapi.Scope, id *krm.GKEHubScopeIdentity) *krm.GKEHubScopeSpec {
	out := GKEHubScopeSpec_FromAPI(mapCtx, r)
	if out != nil {
		out.ProjectRef = refs.ProjectRef{External: "projects/" + id.ProjectID}
		out.Location = direct.LazyPtr(id.Location)
		out.ResourceID = direct.LazyPtr(id.ID())
	}
	return out
}

func GKEHubScopeStatusAPIToKRM(mapCtx *direct.MapContext, r *gkehubapi.Scope) *krm.GKEHubScopeStatus {
	return GKEHubScopeObservedState_FromAPI(mapCtx, r)
}
