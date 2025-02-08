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

package vmwareengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DnsBindPermission_FromProto(mapCtx *direct.MapContext, in *pb.DnsBindPermission) *krm.DnsBindPermission {
	if in == nil {
		return nil
	}
	out := &krm.DnsBindPermission{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
func DnsBindPermission_ToProto(mapCtx *direct.MapContext, in *krm.DnsBindPermission) *pb.DnsBindPermission {
	if in == nil {
		return nil
	}
	out := &pb.DnsBindPermission{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
func DnsBindPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsBindPermission) *krm.DnsBindPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DnsBindPermissionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Principals = direct.Slice_FromProto(mapCtx, in.Principals, Principal_FromProto)
	return out
}
func DnsBindPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DnsBindPermissionObservedState) *pb.DnsBindPermission {
	if in == nil {
		return nil
	}
	out := &pb.DnsBindPermission{}
	out.Name = direct.ValueOf(in.Name)
	out.Principals = direct.Slice_ToProto(mapCtx, in.Principals, Principal_ToProto)
	return out
}
func Principal_FromProto(mapCtx *direct.MapContext, in *pb.Principal) *krm.Principal {
	if in == nil {
		return nil
	}
	out := &krm.Principal{}
	out.User = direct.LazyPtr(in.GetUser())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func Principal_ToProto(mapCtx *direct.MapContext, in *krm.Principal) *pb.Principal {
	if in == nil {
		return nil
	}
	out := &pb.Principal{}
	if oneof := Principal_User_ToProto(mapCtx, in.User); oneof != nil {
		out.Principal = oneof
	}
	if oneof := Principal_ServiceAccount_ToProto(mapCtx, in.ServiceAccount); oneof != nil {
		out.Principal = oneof
	}
	return out
}
func VmwareengineDnsBindPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DnsBindPermission) *krm.VmwareengineDnsBindPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineDnsBindPermissionObservedState{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
func VmwareengineDnsBindPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineDnsBindPermissionObservedState) *pb.DnsBindPermission {
	if in == nil {
		return nil
	}
	out := &pb.DnsBindPermission{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
func VmwareengineDnsBindPermissionSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsBindPermission) *krm.VmwareengineDnsBindPermissionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineDnsBindPermissionSpec{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
func VmwareengineDnsBindPermissionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineDnsBindPermissionSpec) *pb.DnsBindPermission {
	if in == nil {
		return nil
	}
	out := &pb.DnsBindPermission{}
	// MISSING: Name
	// MISSING: Principals
	return out
}
