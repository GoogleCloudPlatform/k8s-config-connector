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

package resourcemanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Organization_FromProto(mapCtx *direct.MapContext, in *pb.Organization) *krm.Organization {
	if in == nil {
		return nil
	}
	out := &krm.Organization{}
	// MISSING: Name
	// MISSING: DisplayName
	out.DirectoryCustomerID = direct.LazyPtr(in.GetDirectoryCustomerId())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func Organization_ToProto(mapCtx *direct.MapContext, in *krm.Organization) *pb.Organization {
	if in == nil {
		return nil
	}
	out := &pb.Organization{}
	// MISSING: Name
	// MISSING: DisplayName
	if oneof := Organization_DirectoryCustomerId_ToProto(mapCtx, in.DirectoryCustomerID); oneof != nil {
		out.Owner = oneof
	}
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func OrganizationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Organization) *krm.OrganizationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: DirectoryCustomerID
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func OrganizationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationObservedState) *pb.Organization {
	if in == nil {
		return nil
	}
	out := &pb.Organization{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: DirectoryCustomerID
	out.State = direct.Enum_ToProto[pb.Organization_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func ResourcemanagerOrganizationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Organization) *krm.ResourcemanagerOrganizationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerOrganizationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DirectoryCustomerID
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func ResourcemanagerOrganizationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerOrganizationObservedState) *pb.Organization {
	if in == nil {
		return nil
	}
	out := &pb.Organization{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DirectoryCustomerID
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func ResourcemanagerOrganizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Organization) *krm.ResourcemanagerOrganizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerOrganizationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DirectoryCustomerID
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func ResourcemanagerOrganizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerOrganizationSpec) *pb.Organization {
	if in == nil {
		return nil
	}
	out := &pb.Organization{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: DirectoryCustomerID
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
