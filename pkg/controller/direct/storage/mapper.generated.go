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

package storage

import (
	pb "cloud.google.com/go/storage/internal/apiv2/storagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Bucket_FromProto(mapCtx *direct.MapContext, in *pb.Bucket) *krm.Bucket {
	if in == nil {
		return nil
	}
	out := &krm.Bucket{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: BucketID
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Project = direct.LazyPtr(in.GetProject())
	// MISSING: Metageneration
	out.Location = direct.LazyPtr(in.GetLocation())
	// MISSING: LocationType
	out.StorageClass = direct.LazyPtr(in.GetStorageClass())
	out.Rpo = direct.LazyPtr(in.GetRpo())
	out.Acl = direct.Slice_FromProto(mapCtx, in.Acl, BucketAccessControl_FromProto)
	out.DefaultObjectAcl = direct.Slice_FromProto(mapCtx, in.DefaultObjectAcl, ObjectAccessControl_FromProto)
	out.Lifecycle = Bucket_Lifecycle_FromProto(mapCtx, in.GetLifecycle())
	// MISSING: CreateTime
	out.Cors = direct.Slice_FromProto(mapCtx, in.Cors, Bucket_Cors_FromProto)
	// MISSING: UpdateTime
	out.DefaultEventBasedHold = direct.LazyPtr(in.GetDefaultEventBasedHold())
	out.Labels = in.Labels
	out.Website = Bucket_Website_FromProto(mapCtx, in.GetWebsite())
	out.Versioning = Bucket_Versioning_FromProto(mapCtx, in.GetVersioning())
	out.Logging = Bucket_Logging_FromProto(mapCtx, in.GetLogging())
	// MISSING: Owner
	out.Encryption = Bucket_Encryption_FromProto(mapCtx, in.GetEncryption())
	out.Billing = Bucket_Billing_FromProto(mapCtx, in.GetBilling())
	out.RetentionPolicy = Bucket_RetentionPolicy_FromProto(mapCtx, in.GetRetentionPolicy())
	out.IamConfig = Bucket_IamConfig_FromProto(mapCtx, in.GetIamConfig())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.CustomPlacementConfig = Bucket_CustomPlacementConfig_FromProto(mapCtx, in.GetCustomPlacementConfig())
	out.Autoclass = Bucket_Autoclass_FromProto(mapCtx, in.GetAutoclass())
	out.HierarchicalNamespace = Bucket_HierarchicalNamespace_FromProto(mapCtx, in.GetHierarchicalNamespace())
	out.SoftDeletePolicy = Bucket_SoftDeletePolicy_FromProto(mapCtx, in.GetSoftDeletePolicy())
	return out
}
func Bucket_ToProto(mapCtx *direct.MapContext, in *krm.Bucket) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: BucketID
	out.Etag = direct.ValueOf(in.Etag)
	out.Project = direct.ValueOf(in.Project)
	// MISSING: Metageneration
	out.Location = direct.ValueOf(in.Location)
	// MISSING: LocationType
	out.StorageClass = direct.ValueOf(in.StorageClass)
	out.Rpo = direct.ValueOf(in.Rpo)
	out.Acl = direct.Slice_ToProto(mapCtx, in.Acl, BucketAccessControl_ToProto)
	out.DefaultObjectAcl = direct.Slice_ToProto(mapCtx, in.DefaultObjectAcl, ObjectAccessControl_ToProto)
	out.Lifecycle = Bucket_Lifecycle_ToProto(mapCtx, in.Lifecycle)
	// MISSING: CreateTime
	out.Cors = direct.Slice_ToProto(mapCtx, in.Cors, Bucket_Cors_ToProto)
	// MISSING: UpdateTime
	out.DefaultEventBasedHold = direct.ValueOf(in.DefaultEventBasedHold)
	out.Labels = in.Labels
	out.Website = Bucket_Website_ToProto(mapCtx, in.Website)
	out.Versioning = Bucket_Versioning_ToProto(mapCtx, in.Versioning)
	out.Logging = Bucket_Logging_ToProto(mapCtx, in.Logging)
	// MISSING: Owner
	out.Encryption = Bucket_Encryption_ToProto(mapCtx, in.Encryption)
	out.Billing = Bucket_Billing_ToProto(mapCtx, in.Billing)
	out.RetentionPolicy = Bucket_RetentionPolicy_ToProto(mapCtx, in.RetentionPolicy)
	out.IamConfig = Bucket_IamConfig_ToProto(mapCtx, in.IamConfig)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.CustomPlacementConfig = Bucket_CustomPlacementConfig_ToProto(mapCtx, in.CustomPlacementConfig)
	out.Autoclass = Bucket_Autoclass_ToProto(mapCtx, in.Autoclass)
	out.HierarchicalNamespace = Bucket_HierarchicalNamespace_ToProto(mapCtx, in.HierarchicalNamespace)
	out.SoftDeletePolicy = Bucket_SoftDeletePolicy_ToProto(mapCtx, in.SoftDeletePolicy)
	return out
}
func BucketAccessControl_FromProto(mapCtx *direct.MapContext, in *pb.BucketAccessControl) *krm.BucketAccessControl {
	if in == nil {
		return nil
	}
	out := &krm.BucketAccessControl{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.ID = direct.LazyPtr(in.GetId())
	out.Entity = direct.LazyPtr(in.GetEntity())
	// MISSING: EntityAlt
	out.EntityID = direct.LazyPtr(in.GetEntityId())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Email = direct.LazyPtr(in.GetEmail())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.ProjectTeam = ProjectTeam_FromProto(mapCtx, in.GetProjectTeam())
	return out
}
func BucketAccessControl_ToProto(mapCtx *direct.MapContext, in *krm.BucketAccessControl) *pb.BucketAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.BucketAccessControl{}
	out.Role = direct.ValueOf(in.Role)
	out.Id = direct.ValueOf(in.ID)
	out.Entity = direct.ValueOf(in.Entity)
	// MISSING: EntityAlt
	out.EntityId = direct.ValueOf(in.EntityID)
	out.Etag = direct.ValueOf(in.Etag)
	out.Email = direct.ValueOf(in.Email)
	out.Domain = direct.ValueOf(in.Domain)
	out.ProjectTeam = ProjectTeam_ToProto(mapCtx, in.ProjectTeam)
	return out
}
func BucketAccessControlObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BucketAccessControl) *krm.BucketAccessControlObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BucketAccessControlObservedState{}
	// MISSING: Role
	// MISSING: ID
	// MISSING: Entity
	out.EntityAlt = direct.LazyPtr(in.GetEntityAlt())
	// MISSING: EntityID
	// MISSING: Etag
	// MISSING: Email
	// MISSING: Domain
	// MISSING: ProjectTeam
	return out
}
func BucketAccessControlObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BucketAccessControlObservedState) *pb.BucketAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.BucketAccessControl{}
	// MISSING: Role
	// MISSING: ID
	// MISSING: Entity
	out.EntityAlt = direct.ValueOf(in.EntityAlt)
	// MISSING: EntityID
	// MISSING: Etag
	// MISSING: Email
	// MISSING: Domain
	// MISSING: ProjectTeam
	return out
}
func BucketObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Bucket) *krm.BucketObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BucketObservedState{}
	// MISSING: Name
	out.BucketID = direct.LazyPtr(in.GetBucketId())
	// MISSING: Etag
	// MISSING: Project
	out.Metageneration = direct.LazyPtr(in.GetMetageneration())
	// MISSING: Location
	out.LocationType = direct.LazyPtr(in.GetLocationType())
	// MISSING: StorageClass
	// MISSING: Rpo
	out.Acl = direct.Slice_FromProto(mapCtx, in.Acl, BucketAccessControlObservedState_FromProto)
	out.DefaultObjectAcl = direct.Slice_FromProto(mapCtx, in.DefaultObjectAcl, ObjectAccessControlObservedState_FromProto)
	// MISSING: Lifecycle
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Cors
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	out.Autoclass = Bucket_AutoclassObservedState_FromProto(mapCtx, in.GetAutoclass())
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
func BucketObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BucketObservedState) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	// MISSING: Name
	out.BucketId = direct.ValueOf(in.BucketID)
	// MISSING: Etag
	// MISSING: Project
	out.Metageneration = direct.ValueOf(in.Metageneration)
	// MISSING: Location
	out.LocationType = direct.ValueOf(in.LocationType)
	// MISSING: StorageClass
	// MISSING: Rpo
	out.Acl = direct.Slice_ToProto(mapCtx, in.Acl, BucketAccessControlObservedState_ToProto)
	out.DefaultObjectAcl = direct.Slice_ToProto(mapCtx, in.DefaultObjectAcl, ObjectAccessControlObservedState_ToProto)
	// MISSING: Lifecycle
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Cors
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	out.Owner = Owner_ToProto(mapCtx, in.Owner)
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	out.Autoclass = Bucket_AutoclassObservedState_ToProto(mapCtx, in.Autoclass)
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
func Bucket_Autoclass_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Autoclass) *krm.Bucket_Autoclass {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Autoclass{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	// MISSING: ToggleTime
	out.TerminalStorageClass = in.TerminalStorageClass
	// MISSING: TerminalStorageClassUpdateTime
	return out
}
func Bucket_Autoclass_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Autoclass) *pb.Bucket_Autoclass {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Autoclass{}
	out.Enabled = direct.ValueOf(in.Enabled)
	// MISSING: ToggleTime
	out.TerminalStorageClass = in.TerminalStorageClass
	// MISSING: TerminalStorageClassUpdateTime
	return out
}
func Bucket_AutoclassObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Autoclass) *krm.Bucket_AutoclassObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_AutoclassObservedState{}
	// MISSING: Enabled
	out.ToggleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetToggleTime())
	// MISSING: TerminalStorageClass
	out.TerminalStorageClassUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetTerminalStorageClassUpdateTime())
	return out
}
func Bucket_AutoclassObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_AutoclassObservedState) *pb.Bucket_Autoclass {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Autoclass{}
	// MISSING: Enabled
	out.ToggleTime = direct.StringTimestamp_ToProto(mapCtx, in.ToggleTime)
	// MISSING: TerminalStorageClass
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.TerminalStorageClassUpdateTime); oneof != nil {
		out.TerminalStorageClassUpdateTime = &pb.Bucket_Autoclass_TerminalStorageClassUpdateTime{TerminalStorageClassUpdateTime: oneof}
	}
	return out
}
func Bucket_Billing_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Billing) *krm.Bucket_Billing {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Billing{}
	out.RequesterPays = direct.LazyPtr(in.GetRequesterPays())
	return out
}
func Bucket_Billing_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Billing) *pb.Bucket_Billing {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Billing{}
	out.RequesterPays = direct.ValueOf(in.RequesterPays)
	return out
}
func Bucket_Cors_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Cors) *krm.Bucket_Cors {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Cors{}
	out.Origin = in.Origin
	out.Method = in.Method
	out.ResponseHeader = in.ResponseHeader
	out.MaxAgeSeconds = direct.LazyPtr(in.GetMaxAgeSeconds())
	return out
}
func Bucket_Cors_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Cors) *pb.Bucket_Cors {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Cors{}
	out.Origin = in.Origin
	out.Method = in.Method
	out.ResponseHeader = in.ResponseHeader
	out.MaxAgeSeconds = direct.ValueOf(in.MaxAgeSeconds)
	return out
}
func Bucket_CustomPlacementConfig_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_CustomPlacementConfig) *krm.Bucket_CustomPlacementConfig {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_CustomPlacementConfig{}
	out.DataLocations = in.DataLocations
	return out
}
func Bucket_CustomPlacementConfig_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_CustomPlacementConfig) *pb.Bucket_CustomPlacementConfig {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_CustomPlacementConfig{}
	out.DataLocations = in.DataLocations
	return out
}
func Bucket_Encryption_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Encryption) *krm.Bucket_Encryption {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Encryption{}
	out.DefaultKMSKey = direct.LazyPtr(in.GetDefaultKmsKey())
	return out
}
func Bucket_Encryption_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Encryption) *pb.Bucket_Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Encryption{}
	out.DefaultKmsKey = direct.ValueOf(in.DefaultKMSKey)
	return out
}
func Bucket_HierarchicalNamespace_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_HierarchicalNamespace) *krm.Bucket_HierarchicalNamespace {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_HierarchicalNamespace{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func Bucket_HierarchicalNamespace_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_HierarchicalNamespace) *pb.Bucket_HierarchicalNamespace {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_HierarchicalNamespace{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func Bucket_IamConfig_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_IamConfig) *krm.Bucket_IamConfig {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_IamConfig{}
	out.UniformBucketLevelAccess = Bucket_IamConfig_UniformBucketLevelAccess_FromProto(mapCtx, in.GetUniformBucketLevelAccess())
	out.PublicAccessPrevention = direct.LazyPtr(in.GetPublicAccessPrevention())
	return out
}
func Bucket_IamConfig_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_IamConfig) *pb.Bucket_IamConfig {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_IamConfig{}
	out.UniformBucketLevelAccess = Bucket_IamConfig_UniformBucketLevelAccess_ToProto(mapCtx, in.UniformBucketLevelAccess)
	out.PublicAccessPrevention = direct.ValueOf(in.PublicAccessPrevention)
	return out
}
func Bucket_IamConfig_UniformBucketLevelAccess_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_IamConfig_UniformBucketLevelAccess) *krm.Bucket_IamConfig_UniformBucketLevelAccess {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_IamConfig_UniformBucketLevelAccess{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.LockTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLockTime())
	return out
}
func Bucket_IamConfig_UniformBucketLevelAccess_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_IamConfig_UniformBucketLevelAccess) *pb.Bucket_IamConfig_UniformBucketLevelAccess {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_IamConfig_UniformBucketLevelAccess{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.LockTime = direct.StringTimestamp_ToProto(mapCtx, in.LockTime)
	return out
}
func Bucket_Lifecycle_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle) *krm.Bucket_Lifecycle {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Lifecycle{}
	out.Rule = direct.Slice_FromProto(mapCtx, in.Rule, Bucket_Lifecycle_Rule_FromProto)
	return out
}
func Bucket_Lifecycle_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Lifecycle) *pb.Bucket_Lifecycle {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle{}
	out.Rule = direct.Slice_ToProto(mapCtx, in.Rule, Bucket_Lifecycle_Rule_ToProto)
	return out
}
func Bucket_Lifecycle_Rule_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule) *krm.Bucket_Lifecycle_Rule {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Lifecycle_Rule{}
	out.Action = Bucket_Lifecycle_Rule_Action_FromProto(mapCtx, in.GetAction())
	out.Condition = Bucket_Lifecycle_Rule_Condition_FromProto(mapCtx, in.GetCondition())
	return out
}
func Bucket_Lifecycle_Rule_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Lifecycle_Rule) *pb.Bucket_Lifecycle_Rule {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule{}
	out.Action = Bucket_Lifecycle_Rule_Action_ToProto(mapCtx, in.Action)
	out.Condition = Bucket_Lifecycle_Rule_Condition_ToProto(mapCtx, in.Condition)
	return out
}
func Bucket_Lifecycle_Rule_Action_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule_Action) *krm.Bucket_Lifecycle_Rule_Action {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Lifecycle_Rule_Action{}
	out.Type = direct.LazyPtr(in.GetType())
	out.StorageClass = direct.LazyPtr(in.GetStorageClass())
	return out
}
func Bucket_Lifecycle_Rule_Action_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Lifecycle_Rule_Action) *pb.Bucket_Lifecycle_Rule_Action {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule_Action{}
	out.Type = direct.ValueOf(in.Type)
	out.StorageClass = direct.ValueOf(in.StorageClass)
	return out
}
func Bucket_Lifecycle_Rule_Condition_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule_Condition) *krm.Bucket_Lifecycle_Rule_Condition {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Lifecycle_Rule_Condition{}
	out.AgeDays = in.AgeDays
	out.CreatedBefore = Date_FromProto(mapCtx, in.GetCreatedBefore())
	out.IsLive = in.IsLive
	out.NumNewerVersions = in.NumNewerVersions
	out.MatchesStorageClass = in.MatchesStorageClass
	out.DaysSinceCustomTime = in.DaysSinceCustomTime
	out.CustomTimeBefore = Date_FromProto(mapCtx, in.GetCustomTimeBefore())
	out.DaysSinceNoncurrentTime = in.DaysSinceNoncurrentTime
	out.NoncurrentTimeBefore = Date_FromProto(mapCtx, in.GetNoncurrentTimeBefore())
	out.MatchesPrefix = in.MatchesPrefix
	out.MatchesSuffix = in.MatchesSuffix
	return out
}
func Bucket_Lifecycle_Rule_Condition_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Lifecycle_Rule_Condition) *pb.Bucket_Lifecycle_Rule_Condition {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule_Condition{}
	out.AgeDays = in.AgeDays
	out.CreatedBefore = Date_ToProto(mapCtx, in.CreatedBefore)
	out.IsLive = in.IsLive
	out.NumNewerVersions = in.NumNewerVersions
	out.MatchesStorageClass = in.MatchesStorageClass
	out.DaysSinceCustomTime = in.DaysSinceCustomTime
	out.CustomTimeBefore = Date_ToProto(mapCtx, in.CustomTimeBefore)
	out.DaysSinceNoncurrentTime = in.DaysSinceNoncurrentTime
	out.NoncurrentTimeBefore = Date_ToProto(mapCtx, in.NoncurrentTimeBefore)
	out.MatchesPrefix = in.MatchesPrefix
	out.MatchesSuffix = in.MatchesSuffix
	return out
}
func Bucket_Logging_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Logging) *krm.Bucket_Logging {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Logging{}
	out.LogBucket = direct.LazyPtr(in.GetLogBucket())
	out.LogObjectPrefix = direct.LazyPtr(in.GetLogObjectPrefix())
	return out
}
func Bucket_Logging_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Logging) *pb.Bucket_Logging {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Logging{}
	out.LogBucket = direct.ValueOf(in.LogBucket)
	out.LogObjectPrefix = direct.ValueOf(in.LogObjectPrefix)
	return out
}
func Bucket_RetentionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_RetentionPolicy) *krm.Bucket_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_RetentionPolicy{}
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	out.IsLocked = direct.LazyPtr(in.GetIsLocked())
	out.RetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetRetentionDuration())
	return out
}
func Bucket_RetentionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_RetentionPolicy) *pb.Bucket_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_RetentionPolicy{}
	out.EffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime)
	out.IsLocked = direct.ValueOf(in.IsLocked)
	out.RetentionDuration = direct.StringDuration_ToProto(mapCtx, in.RetentionDuration)
	return out
}
func Bucket_SoftDeletePolicy_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_SoftDeletePolicy) *krm.Bucket_SoftDeletePolicy {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_SoftDeletePolicy{}
	out.RetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetRetentionDuration())
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	return out
}
func Bucket_SoftDeletePolicy_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_SoftDeletePolicy) *pb.Bucket_SoftDeletePolicy {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_SoftDeletePolicy{}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.RetentionDuration); oneof != nil {
		out.RetentionDuration = &pb.Bucket_SoftDeletePolicy_RetentionDuration{RetentionDuration: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime); oneof != nil {
		out.EffectiveTime = &pb.Bucket_SoftDeletePolicy_EffectiveTime{EffectiveTime: oneof}
	}
	return out
}
func Bucket_Versioning_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Versioning) *krm.Bucket_Versioning {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Versioning{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func Bucket_Versioning_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Versioning) *pb.Bucket_Versioning {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Versioning{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func Bucket_Website_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Website) *krm.Bucket_Website {
	if in == nil {
		return nil
	}
	out := &krm.Bucket_Website{}
	out.MainPageSuffix = direct.LazyPtr(in.GetMainPageSuffix())
	out.NotFoundPage = direct.LazyPtr(in.GetNotFoundPage())
	return out
}
func Bucket_Website_ToProto(mapCtx *direct.MapContext, in *krm.Bucket_Website) *pb.Bucket_Website {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Website{}
	out.MainPageSuffix = direct.ValueOf(in.MainPageSuffix)
	out.NotFoundPage = direct.ValueOf(in.NotFoundPage)
	return out
}
func ObjectAccessControl_FromProto(mapCtx *direct.MapContext, in *pb.ObjectAccessControl) *krm.ObjectAccessControl {
	if in == nil {
		return nil
	}
	out := &krm.ObjectAccessControl{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.ID = direct.LazyPtr(in.GetId())
	out.Entity = direct.LazyPtr(in.GetEntity())
	// MISSING: EntityAlt
	out.EntityID = direct.LazyPtr(in.GetEntityId())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Email = direct.LazyPtr(in.GetEmail())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.ProjectTeam = ProjectTeam_FromProto(mapCtx, in.GetProjectTeam())
	return out
}
func ObjectAccessControl_ToProto(mapCtx *direct.MapContext, in *krm.ObjectAccessControl) *pb.ObjectAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.ObjectAccessControl{}
	out.Role = direct.ValueOf(in.Role)
	out.Id = direct.ValueOf(in.ID)
	out.Entity = direct.ValueOf(in.Entity)
	// MISSING: EntityAlt
	out.EntityId = direct.ValueOf(in.EntityID)
	out.Etag = direct.ValueOf(in.Etag)
	out.Email = direct.ValueOf(in.Email)
	out.Domain = direct.ValueOf(in.Domain)
	out.ProjectTeam = ProjectTeam_ToProto(mapCtx, in.ProjectTeam)
	return out
}
func ObjectAccessControlObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ObjectAccessControl) *krm.ObjectAccessControlObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ObjectAccessControlObservedState{}
	// MISSING: Role
	// MISSING: ID
	// MISSING: Entity
	out.EntityAlt = direct.LazyPtr(in.GetEntityAlt())
	// MISSING: EntityID
	// MISSING: Etag
	// MISSING: Email
	// MISSING: Domain
	// MISSING: ProjectTeam
	return out
}
func ObjectAccessControlObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ObjectAccessControlObservedState) *pb.ObjectAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.ObjectAccessControl{}
	// MISSING: Role
	// MISSING: ID
	// MISSING: Entity
	out.EntityAlt = direct.ValueOf(in.EntityAlt)
	// MISSING: EntityID
	// MISSING: Etag
	// MISSING: Email
	// MISSING: Domain
	// MISSING: ProjectTeam
	return out
}
func Owner_FromProto(mapCtx *direct.MapContext, in *pb.Owner) *krm.Owner {
	if in == nil {
		return nil
	}
	out := &krm.Owner{}
	out.Entity = direct.LazyPtr(in.GetEntity())
	out.EntityID = direct.LazyPtr(in.GetEntityId())
	return out
}
func Owner_ToProto(mapCtx *direct.MapContext, in *krm.Owner) *pb.Owner {
	if in == nil {
		return nil
	}
	out := &pb.Owner{}
	out.Entity = direct.ValueOf(in.Entity)
	out.EntityId = direct.ValueOf(in.EntityID)
	return out
}
func ProjectTeam_FromProto(mapCtx *direct.MapContext, in *pb.ProjectTeam) *krm.ProjectTeam {
	if in == nil {
		return nil
	}
	out := &krm.ProjectTeam{}
	out.ProjectNumber = direct.LazyPtr(in.GetProjectNumber())
	out.Team = direct.LazyPtr(in.GetTeam())
	return out
}
func ProjectTeam_ToProto(mapCtx *direct.MapContext, in *krm.ProjectTeam) *pb.ProjectTeam {
	if in == nil {
		return nil
	}
	out := &pb.ProjectTeam{}
	out.ProjectNumber = direct.ValueOf(in.ProjectNumber)
	out.Team = direct.ValueOf(in.Team)
	return out
}
func StorageBucketObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Bucket) *krm.StorageBucketObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketObservedState{}
	// MISSING: Name
	// MISSING: BucketID
	// MISSING: Etag
	// MISSING: Project
	// MISSING: Metageneration
	// MISSING: Location
	// MISSING: LocationType
	// MISSING: StorageClass
	// MISSING: Rpo
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	// MISSING: Lifecycle
	// MISSING: CreateTime
	// MISSING: Cors
	// MISSING: UpdateTime
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	// MISSING: Owner
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	// MISSING: Autoclass
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
func StorageBucketObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketObservedState) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	// MISSING: Name
	// MISSING: BucketID
	// MISSING: Etag
	// MISSING: Project
	// MISSING: Metageneration
	// MISSING: Location
	// MISSING: LocationType
	// MISSING: StorageClass
	// MISSING: Rpo
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	// MISSING: Lifecycle
	// MISSING: CreateTime
	// MISSING: Cors
	// MISSING: UpdateTime
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	// MISSING: Owner
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	// MISSING: Autoclass
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
func StorageBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.Bucket) *krm.StorageBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketSpec{}
	// MISSING: Name
	// MISSING: BucketID
	// MISSING: Etag
	// MISSING: Project
	// MISSING: Metageneration
	// MISSING: Location
	// MISSING: LocationType
	// MISSING: StorageClass
	// MISSING: Rpo
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	// MISSING: Lifecycle
	// MISSING: CreateTime
	// MISSING: Cors
	// MISSING: UpdateTime
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	// MISSING: Owner
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	// MISSING: Autoclass
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
func StorageBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketSpec) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	// MISSING: Name
	// MISSING: BucketID
	// MISSING: Etag
	// MISSING: Project
	// MISSING: Metageneration
	// MISSING: Location
	// MISSING: LocationType
	// MISSING: StorageClass
	// MISSING: Rpo
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	// MISSING: Lifecycle
	// MISSING: CreateTime
	// MISSING: Cors
	// MISSING: UpdateTime
	// MISSING: DefaultEventBasedHold
	// MISSING: Labels
	// MISSING: Website
	// MISSING: Versioning
	// MISSING: Logging
	// MISSING: Owner
	// MISSING: Encryption
	// MISSING: Billing
	// MISSING: RetentionPolicy
	// MISSING: IamConfig
	// MISSING: SatisfiesPzs
	// MISSING: CustomPlacementConfig
	// MISSING: Autoclass
	// MISSING: HierarchicalNamespace
	// MISSING: SoftDeletePolicy
	return out
}
