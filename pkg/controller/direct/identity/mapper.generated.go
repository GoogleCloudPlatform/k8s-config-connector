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

package identity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/identity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
)
func AccessPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krm.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AccessPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Scopes = in.Scopes
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func AccessPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AccessPolicy) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.Title = direct.ValueOf(in.Title)
	out.Scopes = in.Scopes
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func IdentityAccessPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krm.IdentityAccessPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IdentityAccessPolicyObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Title
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func IdentityAccessPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IdentityAccessPolicyObservedState) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Title
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func IdentityAccessPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krm.IdentityAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.IdentityAccessPolicySpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Title
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func IdentityAccessPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.IdentityAccessPolicySpec) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Title
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
