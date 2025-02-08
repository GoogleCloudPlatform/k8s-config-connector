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

package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AwsSourceDetails_FromProto(mapCtx *direct.MapContext, in *pb.AwsSourceDetails) *krm.AwsSourceDetails {
	if in == nil {
		return nil
	}
	out := &krm.AwsSourceDetails{}
	out.AccessKeyCreds = AwsSourceDetails_AccessKeyCredentials_FromProto(mapCtx, in.GetAccessKeyCreds())
	out.AwsRegion = direct.LazyPtr(in.GetAwsRegion())
	// MISSING: State
	// MISSING: Error
	out.InventoryTagList = direct.Slice_FromProto(mapCtx, in.InventoryTagList, AwsSourceDetails_Tag_FromProto)
	out.InventorySecurityGroupNames = in.InventorySecurityGroupNames
	out.MigrationResourcesUserTags = in.MigrationResourcesUserTags
	// MISSING: PublicIP
	return out
}
func AwsSourceDetails_ToProto(mapCtx *direct.MapContext, in *krm.AwsSourceDetails) *pb.AwsSourceDetails {
	if in == nil {
		return nil
	}
	out := &pb.AwsSourceDetails{}
	if oneof := AwsSourceDetails_AccessKeyCredentials_ToProto(mapCtx, in.AccessKeyCreds); oneof != nil {
		out.CredentialsType = &pb.AwsSourceDetails_AccessKeyCreds{AccessKeyCreds: oneof}
	}
	out.AwsRegion = direct.ValueOf(in.AwsRegion)
	// MISSING: State
	// MISSING: Error
	out.InventoryTagList = direct.Slice_ToProto(mapCtx, in.InventoryTagList, AwsSourceDetails_Tag_ToProto)
	out.InventorySecurityGroupNames = in.InventorySecurityGroupNames
	out.MigrationResourcesUserTags = in.MigrationResourcesUserTags
	// MISSING: PublicIP
	return out
}
func AwsSourceDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AwsSourceDetails) *krm.AwsSourceDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AwsSourceDetailsObservedState{}
	// MISSING: AccessKeyCreds
	// MISSING: AwsRegion
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: InventoryTagList
	// MISSING: InventorySecurityGroupNames
	// MISSING: MigrationResourcesUserTags
	out.PublicIP = direct.LazyPtr(in.GetPublicIp())
	return out
}
func AwsSourceDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AwsSourceDetailsObservedState) *pb.AwsSourceDetails {
	if in == nil {
		return nil
	}
	out := &pb.AwsSourceDetails{}
	// MISSING: AccessKeyCreds
	// MISSING: AwsRegion
	out.State = direct.Enum_ToProto[pb.AwsSourceDetails_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: InventoryTagList
	// MISSING: InventorySecurityGroupNames
	// MISSING: MigrationResourcesUserTags
	out.PublicIp = direct.ValueOf(in.PublicIP)
	return out
}
func AwsSourceDetails_AccessKeyCredentials_FromProto(mapCtx *direct.MapContext, in *pb.AwsSourceDetails_AccessKeyCredentials) *krm.AwsSourceDetails_AccessKeyCredentials {
	if in == nil {
		return nil
	}
	out := &krm.AwsSourceDetails_AccessKeyCredentials{}
	out.AccessKeyID = direct.LazyPtr(in.GetAccessKeyId())
	out.SecretAccessKey = direct.LazyPtr(in.GetSecretAccessKey())
	return out
}
func AwsSourceDetails_AccessKeyCredentials_ToProto(mapCtx *direct.MapContext, in *krm.AwsSourceDetails_AccessKeyCredentials) *pb.AwsSourceDetails_AccessKeyCredentials {
	if in == nil {
		return nil
	}
	out := &pb.AwsSourceDetails_AccessKeyCredentials{}
	out.AccessKeyId = direct.ValueOf(in.AccessKeyID)
	out.SecretAccessKey = direct.ValueOf(in.SecretAccessKey)
	return out
}
func AwsSourceDetails_Tag_FromProto(mapCtx *direct.MapContext, in *pb.AwsSourceDetails_Tag) *krm.AwsSourceDetails_Tag {
	if in == nil {
		return nil
	}
	out := &krm.AwsSourceDetails_Tag{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func AwsSourceDetails_Tag_ToProto(mapCtx *direct.MapContext, in *krm.AwsSourceDetails_Tag) *pb.AwsSourceDetails_Tag {
	if in == nil {
		return nil
	}
	out := &pb.AwsSourceDetails_Tag{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.Source {
	if in == nil {
		return nil
	}
	out := &krm.Source{}
	out.Vmware = VmwareSourceDetails_FromProto(mapCtx, in.GetVmware())
	out.Aws = AwsSourceDetails_FromProto(mapCtx, in.GetAws())
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Source_ToProto(mapCtx *direct.MapContext, in *krm.Source) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	if oneof := VmwareSourceDetails_ToProto(mapCtx, in.Vmware); oneof != nil {
		out.SourceDetails = &pb.Source_Vmware{Vmware: oneof}
	}
	if oneof := AwsSourceDetails_ToProto(mapCtx, in.Aws); oneof != nil {
		out.SourceDetails = &pb.Source_Aws{Aws: oneof}
	}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	return out
}
func SourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.SourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SourceObservedState{}
	// MISSING: Vmware
	out.Aws = AwsSourceDetailsObservedState_FromProto(mapCtx, in.GetAws())
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	return out
}
func SourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SourceObservedState) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Vmware
	if oneof := AwsSourceDetailsObservedState_ToProto(mapCtx, in.Aws); oneof != nil {
		out.SourceDetails = &pb.Source_Aws{Aws: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	return out
}
func VmmigrationSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.VmmigrationSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationSourceObservedState{}
	// MISSING: Vmware
	// MISSING: Aws
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	return out
}
func VmmigrationSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationSourceObservedState) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Vmware
	// MISSING: Aws
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	return out
}
func VmmigrationSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.VmmigrationSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationSourceSpec{}
	// MISSING: Vmware
	// MISSING: Aws
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	return out
}
func VmmigrationSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationSourceSpec) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Vmware
	// MISSING: Aws
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	return out
}
func VmwareSourceDetails_FromProto(mapCtx *direct.MapContext, in *pb.VmwareSourceDetails) *krm.VmwareSourceDetails {
	if in == nil {
		return nil
	}
	out := &krm.VmwareSourceDetails{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.VcenterIP = direct.LazyPtr(in.GetVcenterIp())
	out.Thumbprint = direct.LazyPtr(in.GetThumbprint())
	return out
}
func VmwareSourceDetails_ToProto(mapCtx *direct.MapContext, in *krm.VmwareSourceDetails) *pb.VmwareSourceDetails {
	if in == nil {
		return nil
	}
	out := &pb.VmwareSourceDetails{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.VcenterIp = direct.ValueOf(in.VcenterIP)
	out.Thumbprint = direct.ValueOf(in.Thumbprint)
	return out
}
