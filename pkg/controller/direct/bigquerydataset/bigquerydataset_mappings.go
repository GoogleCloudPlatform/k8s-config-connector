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

// This file only contains the mapper for BigQueryDataset Resource.
// The mapper is created based on cloud.google.com/go/bigquery library.
// The old library google.golang.org/api/bigquery/v2 is depreacted and
// the recommended cloud.google.com/go/bigquery library has a different
// schema. Since this is a beta resource, and we want to be backward compatible,
// we have manually added this custom mapper to accommodate the schema difference.
package bigquerydataset

import (
	"fmt"
	"time"

	pb "cloud.google.com/go/bigquery"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Access_FromProto(mapCtx *direct.MapContext, in *pb.AccessEntry) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.Role = direct.LazyPtr(fmt.Sprintf("%s", in.Role))
	switch in.EntityType {
	case 1:
		out.Domain = direct.LazyPtr(in.Entity)
	case 2:
		out.GroupByEmail = direct.LazyPtr(in.Entity)
	case 3:
		out.UserByEmail = direct.LazyPtr(in.Entity)
	case 4:
		out.SpecialGroup = direct.LazyPtr(in.Entity)
	case 6:
		out.IamMember = direct.LazyPtr(in.Entity)
	}
	out.View = TableReference_FromProto(mapCtx, in.View)
	out.Routine = RoutineReference_FromProto(mapCtx, in.Routine)
	out.Dataset = DatasetAccessEntry_FromProto(mapCtx, in.Dataset)
	return out
}
func Access_ToProto(mapCtx *direct.MapContext, in *krm.Access) *pb.AccessEntry {
	if in == nil {
		return nil
	}
	out := &pb.AccessEntry{}
	out.Role = pb.AccessRole(direct.ValueOf(in.Role))
	if in.Domain != nil {
		out.EntityType = 1
		out.Entity = direct.ValueOf(in.Domain)
	}
	if in.GroupByEmail != nil {
		out.EntityType = 2
		out.Entity = direct.ValueOf(in.GroupByEmail)
	}
	if in.UserByEmail != nil {
		out.EntityType = 3
		out.Entity = direct.ValueOf(in.UserByEmail)
	}
	if in.SpecialGroup != nil {
		out.EntityType = 4
		out.Entity = direct.ValueOf(in.SpecialGroup)
	}
	if in.IamMember != nil {
		out.EntityType = 6
		out.Entity = direct.ValueOf(in.IamMember)
	}
	out.View = TableReference_ToProto(mapCtx, in.View)
	out.Routine = RoutineReference_ToProto(mapCtx, in.Routine)
	out.Dataset = DatasetAccessEntry_ToProto(mapCtx, in.Dataset)
	return out
}
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.DatasetMetadata) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.Kind = direct.LazyPtr("BigQueryDataset")
	out.Etag = direct.LazyPtr(in.ETag)
	out.ID = direct.LazyPtr(in.FullID)
	out.FriendlyName = direct.LazyPtr(in.Name)
	out.Description = direct.LazyPtr(in.Description)
	defaultTableExpirationMs := int64(in.DefaultTableExpiration / time.Millisecond)
	out.DefaultTableExpirationMs = &defaultTableExpirationMs
	defaultPartitionExpirationMs := int64(in.DefaultPartitionExpiration / time.Millisecond)
	out.DefaultPartitionExpirationMs = &defaultPartitionExpirationMs
	out.Labels = in.Labels
	out.Access = direct.Slice_FromProto(mapCtx, in.Access, Access_FromProto)
	//TODO: convert from time.Time to int64
	// out.CreationTime = in.CreationTime
	// out.LastModifiedTime = in.LastModifiedTime
	time.Now().UnixNano()
	out.Location = direct.LazyPtr(in.Location)
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.DefaultEncryptionConfig)
	out.ExternalDatasetReference = ExternalDatasetReference_FromProto(mapCtx, in.ExternalDatasetReference)
	out.DefaultCollation = direct.LazyPtr(in.DefaultCollation)
	maxTimeTravelHours := (int64)(in.MaxTimeTravel / time.Hour)
	out.MaxTimeTravelHours = &maxTimeTravelHours
	out.Tags = direct.Slice_FromProto(mapCtx, in.Tags, DatasetTag_FromProto)
	out.StorageBillingModel = direct.LazyPtr(in.StorageBillingModel)
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.DatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DatasetMetadata{}
	out.ETag = direct.ValueOf(in.Etag)
	out.FullID = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.FriendlyName)
	out.Description = direct.ValueOf(in.Description)
	out.DefaultTableExpiration = time.Duration(*in.DefaultTableExpirationMs) * time.Millisecond
	out.DefaultPartitionExpiration = time.Duration(*in.DefaultPartitionExpirationMs) * time.Millisecond
	out.Labels = in.Labels
	out.Access = direct.Slice_ToProto(mapCtx, in.Access, Access_ToProto)
	out.CreationTime = time.UnixMilli(*in.CreationTime)
	out.LastModifiedTime = time.UnixMilli(*in.LastModifiedTime)
	out.Location = direct.ValueOf(in.Location)
	out.DefaultEncryptionConfig = EncryptionConfiguration_ToProto(mapCtx, in.DefaultEncryptionConfiguration)
	out.ExternalDatasetReference = ExternalDatasetReference_ToProto(mapCtx, in.ExternalDatasetReference)
	out.DefaultCollation = *in.DefaultCollation
	out.MaxTimeTravel = time.Duration(*in.MaxTimeTravelHours) * time.Hour
	out.Tags = direct.Slice_ToProto(mapCtx, in.Tags, DatasetTag_ToProto)
	out.StorageBillingModel = direct.ValueOf(in.StorageBillingModel)
	return out
}
func DatasetAccessEntry_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccessEntry) *krm.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.DatasetAccessEntry{}
	out.Dataset = DatasetReference_FromProto(mapCtx, in.Dataset)
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
func DatasetList_FromProto(mapCtx *direct.MapContext, in *pb.DatasetIterator) *krm.DatasetList {
	if in == nil {
		return nil
	}
	out := &krm.DatasetList{}
	in.ListHidden = true
	out.Kind = direct.LazyPtr("BigQueryDataset")
	var datasets []krm.ListFormatDataset
	var next *pb.Dataset
	next, _ = in.Next()
	for next != nil {
		datasets = append(datasets, *ListFormatDataset_FromProto(mapCtx, next))
		next, _ = in.Next()
	}
	out.Datasets = datasets

	return out
}
func DatasetList_ToProto(mapCtx *direct.MapContext, in *krm.DatasetList) *pb.DatasetIterator {
	if in == nil {
		return nil
	}
	out := &pb.DatasetIterator{}
	// Missing
	return out
}
func DatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.DatasetReference{}
	out.DatasetId = &in.DatasetID
	out.ProjectId = &in.ProjectID
	return out
}
func DatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.DatasetReference) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out.DatasetID = *in.DatasetId
	out.ProjectID = *in.ProjectId
	return out
}
func DatasetTag_FromProto(mapCtx *direct.MapContext, in *pb.DatasetTag) *krm.GcpTag {
	if in == nil {
		return nil
	}
	out := &krm.GcpTag{}
	out.TagKey = direct.LazyPtr(in.TagKey)
	out.TagValue = direct.LazyPtr(in.TagValue)
	return out
}
func DatasetTag_ToProto(mapCtx *direct.MapContext, in *krm.GcpTag) *pb.DatasetTag {
	if in == nil {
		return nil
	}
	out := &pb.DatasetTag{}
	out.TagKey = direct.ValueOf(in.TagKey)
	out.TagValue = direct.ValueOf(in.TagValue)
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	out.KmsKeyRef = &v1beta1.KMSCryptoKeyRef{
		Name: in.KMSKeyName,
	}
	return out
}

func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KMSKeyName = in.KmsKeyRef.Name
	return out
}
func ErrorProto_FromProto(mapCtx *direct.MapContext, in *pb.Error) *krm.ErrorProto {
	if in == nil {
		return nil
	}
	out := &krm.ErrorProto{}
	out.Reason = direct.LazyPtr(in.Reason)
	out.Location = direct.LazyPtr(in.Location)
	out.Message = direct.LazyPtr(in.Message)
	return out
}
func ErrorProto_ToProto(mapCtx *direct.MapContext, in *krm.ErrorProto) *pb.Error {
	if in == nil {
		return nil
	}
	out := &pb.Error{}
	out.Reason = direct.ValueOf(in.Reason)
	out.Location = direct.ValueOf(in.Location)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func ExternalDatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.ExternalDatasetReference) *krm.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.ExternalDatasetReference{}
	out.ExternalSource = direct.LazyPtr(in.ExternalSource)
	out.Connection = direct.LazyPtr(in.Connection)
	return out
}
func ExternalDatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.ExternalDatasetReference) *pb.ExternalDatasetReference {
	if in == nil {
		return nil
	}
	out := &pb.ExternalDatasetReference{}
	out.ExternalSource = direct.ValueOf(in.ExternalSource)
	out.Connection = direct.ValueOf(in.Connection)
	return out
}
func ListFormatDataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.ListFormatDataset {
	if in == nil {
		return nil
	}
	out := &krm.ListFormatDataset{}
	out.Kind = direct.LazyPtr("BigQueryDataset")
	out.DatasetReference = DatasetReference_FromProto(mapCtx, in)
	return out
}
func ListFormatDataset_ToProto(mapCtx *direct.MapContext, in *krm.ListFormatDataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out = DatasetReference_ToProto(mapCtx, in.DatasetReference)
	return out
}
func RoutineReference_FromProto(mapCtx *direct.MapContext, in *pb.Routine) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.DatasetId = &in.DatasetID
	out.ProjectId = &in.ProjectID
	out.RoutineId = &in.RoutineID
	return out
}
func RoutineReference_ToProto(mapCtx *direct.MapContext, in *krm.RoutineReference) *pb.Routine {
	if in == nil {
		return nil
	}
	out := &pb.Routine{}
	out.DatasetID = *in.DatasetId
	out.ProjectID = *in.ProjectId
	out.RoutineID = *in.RoutineId
	return out
}
func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.DatasetId = &in.DatasetID
	out.ProjectId = &in.ProjectID
	out.TableId = &in.TableID
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.DatasetID = *in.DatasetId
	out.ProjectID = *in.ProjectId
	out.TableID = *in.TableId
	return out
}
