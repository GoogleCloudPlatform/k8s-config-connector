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
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.bigtable.admin.v2

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmbigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	out.RowAffinity = AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx, in.GetRowAffinity())
	return out
}
func AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	if oneof := AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx, in.RowAffinity); oneof != nil {
		out.Affinity = &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_{RowAffinity: oneof}
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_SingleClusterRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krmbigtablev1beta1.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krmbigtablev1beta1.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
	return out
}
func BigtableAppProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmbigtablev1beta1.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.BigtableAppProfileObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAppProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAuthorizedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krm.BigtableAuthorizedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableAuthorizedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableAuthorizedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableAuthorizedViewObservedState) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableLogicalViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogicalView) *krm.BigtableLogicalViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableLogicalViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableLogicalViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableLogicalViewObservedState) *pb.LogicalView {
	if in == nil {
		return nil
	}
	out := &pb.LogicalView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableMaterializedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaterializedView) *krm.BigtableMaterializedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableMaterializedViewObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BigtableMaterializedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMaterializedViewObservedState) *pb.MaterializedView {
	if in == nil {
		return nil
	}
	out := &pb.MaterializedView{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmbigtablev1beta1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1beta1.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1beta1.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
