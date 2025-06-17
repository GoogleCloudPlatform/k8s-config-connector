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

// +generated:mapper
// krm.group: accesscontextmanager.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.identity.accesscontextmanager.v1

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessContextManagerAccessPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krmv1beta1.AccessContextManagerAccessPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AccessContextManagerAccessPolicyObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AccessContextManagerAccessPolicyObservedState) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krmv1beta1.AccessContextManagerAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AccessContextManagerAccessPolicySpec{}
	// MISSING: Name
	// MISSING: Parent
	out.Title = direct.LazyPtr(in.GetTitle())
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicySpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AccessContextManagerAccessPolicySpec) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	out.Title = direct.ValueOf(in.Title)
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krmv1beta1.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AccessPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Scopes = in.Scopes
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func AccessPolicy_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AccessPolicy) *pb.AccessPolicy {
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
