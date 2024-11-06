// Copyright 2024 Google LLC
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

package bigquery

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Access_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccess) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.UserByEmail = direct.LazyPtr(in.GetUserByEmail())
	out.GroupByEmail = direct.LazyPtr(in.GetGroupByEmail())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.SpecialGroup = direct.LazyPtr(in.GetSpecialGroup())
	out.IamMember = direct.LazyPtr(in.GetIamMember())
	out.View = ReferencedTable_FromProto(mapCtx, in.GetView())
	out.Routine = ReferencedRoutine_FromProto(mapCtx, in.GetRoutine())
	out.Dataset = DatasetAccessEntry_FromProto(mapCtx, in.GetDataset())
	return out
}
func Access_ToProto(mapCtx *direct.MapContext, in *krm.Access) *pb.DatasetAccess {
	if in == nil {
		return nil
	}
	out := &pb.DatasetAccess{}
	out.Role = in.Role
	out.UserByEmail = in.UserByEmail
	out.GroupByEmail = in.GroupByEmail
	out.Domain = in.Domain
	out.SpecialGroup = in.SpecialGroup
	out.IamMember = in.IamMember
	out.View = ReferencedTable_ToProto(mapCtx, in.View)
	out.Routine = ReferencedRoutine_ToProto(mapCtx, in.Routine)
	out.Dataset = DatasetAccessEntry_ToProto(mapCtx, in.Dataset)
	return out
}

func DatasetAccessEntry_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccessEntry) *krm.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.DatasetAccessEntry{}
	out.Dataset = DatasetReference_FromProto(mapCtx, in.GetDataset())
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetAccessEntry_ToProto(mapCtx *direct.MapContext, in *krm.DatasetAccessEntry) *pb.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &pb.DatasetAccessEntry{}
	out.Dataset = DatasetReference_ToProto(mapCtx, in.Dataset)
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.DatasetReference) *krm.DatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.DatasetReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}
func DatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.DatasetReference) *pb.DatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.DatasetReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	out.KmsKeyRef.Name = *in.KmsKeyName
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	out.KmsKeyName = &in.KmsKeyRef.Name
	return out
}

func ExternalCatalogDatasetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ExternalCatalogDatasetOptions) *krm.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &krm.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	out.DefaultStorageLocationUri = direct.LazyPtr(in.GetDefaultStorageLocationUri())
	return out
}
func ExternalCatalogDatasetOptions_ToProto(mapCtx *direct.MapContext, in *krm.ExternalCatalogDatasetOptions) *pb.ExternalCatalogDatasetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ExternalCatalogDatasetOptions{}
	out.Parameters = in.Parameters
	out.DefaultStorageLocationUri = in.DefaultStorageLocationUri
	return out
}
func ExternalDatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDatasetReference) *krm.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.ExternalDatasetReference{}
	out.ExternalSource = direct.LazyPtr(in.GetExternalSource())
	out.Connection = direct.LazyPtr(in.GetConnection())
	return out
}
func ExternalDatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.ExternalDatasetReference) *pb.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDatasetReference{}
	out.ExternalSource = in.ExternalSource
	out.Connection = in.Connection
	return out
}
func GcpTag_FromProto(mapCtx *direct.MapContext, in *pb.DatasetTags) *krm.GcpTag {
	if in == nil {
		return nil
	}
	out := &krm.GcpTag{}
	out.TagKey = direct.LazyPtr(in.GetTagKey())
	out.TagValue = direct.LazyPtr(in.GetTagValue())
	return out
}
func GcpTag_ToProto(mapCtx *direct.MapContext, in *krm.GcpTag) *pb.DatasetTags {
	if in == nil {
		return nil
	}
	out := &pb.DatasetTags{}
	out.TagKey = in.TagKey
	out.TagValue = in.TagValue
	return out
}
func LinkedDatasetMetadata_FromProto(mapCtx *direct.MapContext, in *pb.LinkedDatasetMetadata) *krm.LinkedDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &krm.LinkedDatasetMetadata{}
	return out
}
func LinkedDatasetMetadata_ToProto(mapCtx *direct.MapContext, in *krm.LinkedDatasetMetadata) *pb.LinkedDatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.LinkedDatasetMetadata{}
	return out
}
func LinkedDatasetSource_FromProto(mapCtx *direct.MapContext, in *pb.LinkedDatasetSource) *krm.LinkedDatasetSource {
	if in == nil {
		return nil
	}
	out := &krm.LinkedDatasetSource{}
	out.SourceDataset = DatasetReference_FromProto(mapCtx, in.GetSourceDataset())
	return out
}
func LinkedDatasetSource_ToProto(mapCtx *direct.MapContext, in *krm.LinkedDatasetSource) *pb.LinkedDatasetSource {
	if in == nil {
		return nil
	}
	out := &pb.LinkedDatasetSource{}
	out.SourceDataset = DatasetReference_ToProto(mapCtx, in.SourceDataset)
	return out
}
func ReferencedTable_FromProto(mapCtx *direct.MapContext, in *pb.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func ReferencedTable_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.TableId = in.TableId
	return out
}
func ReferencedRoutine_FromProto(mapCtx *direct.MapContext, in *pb.RoutineReference) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
	return out
}
func ReferencedRoutine_ToProto(mapCtx *direct.MapContext, in *krm.RoutineReference) *pb.RoutineReference {
	if in == nil {
		return nil
	}
	out := &pb.RoutineReference{}
	out.ProjectId = in.ProjectId
	out.DatasetId = in.DatasetId
	out.RoutineId = in.RoutineId
	return out
}
func RestrictionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RestrictionConfig) *krm.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RestrictionConfig{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func RestrictionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RestrictionConfig) *pb.RestrictionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RestrictionConfig{}
	out.Type = in.Type
	return out
}
