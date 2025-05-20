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
	"slices"
	"strconv"
	"strings"
	"time"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	pb "cloud.google.com/go/bigquery"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDatasetSpec) *pb.DatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DatasetMetadata{}
	acccessList := []*pb.AccessEntry{}
	for _, access := range in.Access {
		curAccess := AccessEntry_ToProto(mapCtx, direct.LazyPtr(access))
		acccessList = append(acccessList, curAccess)
	}
	out.Access = acccessList
	out.DefaultCollation = direct.ValueOf(in.DefaultCollation)
	out.DefaultPartitionExpiration = time.Duration(direct.ValueOf(in.DefaultPartitionExpirationMs)) * time.Millisecond
	out.DefaultTableExpiration = time.Duration(direct.ValueOf(in.DefaultTableExpirationMs)) * time.Millisecond
	out.DefaultEncryptionConfig = EncryptionConfiguration_ToProto(mapCtx, in.DefaultEncryptionConfiguration)
	out.Description = direct.ValueOf(in.Description)
	out.Location = direct.ValueOf(in.Location)
	if in.MaxTimeTravelHours != nil {
		maxHours, _ := strconv.ParseInt(direct.ValueOf(in.MaxTimeTravelHours), 10, 64)
		out.MaxTimeTravel = time.Duration(maxHours) * time.Hour
	}
	out.IsCaseInsensitive = direct.ValueOf(in.IsCaseInsensitive)
	out.Name = direct.ValueOf(in.FriendlyName)
	out.StorageBillingModel = direct.ValueOf(in.StorageBillingModel)
	return out
}
func BigQueryDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetMetadata) *krm.BigQueryDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetSpec{}
	accessList := []krm.Access{}
	for _, access := range in.Access {
		curAccess := AccessEntry_FromProto(mapCtx, access)
		accessList = append(accessList, direct.ValueOf(curAccess))
	}
	out.Access = accessList
	out.DefaultCollation = direct.LazyPtr(in.DefaultCollation)
	out.DefaultPartitionExpirationMs = direct.LazyPtr(in.DefaultPartitionExpiration.Milliseconds())
	out.DefaultTableExpirationMs = direct.LazyPtr(in.DefaultTableExpiration.Milliseconds())
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.DefaultEncryptionConfig)
	out.Description = direct.LazyPtr(in.Description)
	out.FriendlyName = direct.LazyPtr(in.Name)
	out.Location = direct.LazyPtr(in.Location)
	maxTimeInHours := fmt.Sprintf("%v", in.MaxTimeTravel.Hours())
	out.MaxTimeTravelHours = direct.LazyPtr(maxTimeInHours)
	out.IsCaseInsensitive = direct.LazyPtr(in.IsCaseInsensitive)
	out.StorageBillingModel = direct.LazyPtr(in.StorageBillingModel)
	tokens := strings.Split(in.FullID, ":")
	if len(tokens) == 2 {
		out.ResourceID = direct.LazyPtr(tokens[1])
	}
	return out
}
func BigQueryDatasetStatus_FromProto(mapCtx *direct.MapContext, in *pb.DatasetMetadata) *krm.BigQueryDatasetStatus {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetStatus{}
	out.Etag = direct.LazyPtr(in.ETag)
	out.CreationTime = direct.LazyPtr(in.CreationTime.UnixMilli())
	out.LastModifiedTime = direct.LazyPtr(in.LastModifiedTime.UnixMilli())
	// The full dataset ID in the form projectID:datasetID
	tokens := strings.Split(in.FullID, ":")
	if len(tokens) == 2 {
		out.SelfLink = direct.LazyPtr(fmt.Sprintf("https://bigquery.googleapis.com/bigquery/v2/projects/%s/datasets/%s", tokens[0], tokens[1]))
	}
	out.ObservedState = &krm.BigQueryDatasetObservedState{Location: direct.LazyPtr(in.Location)}
	return out
}
func BigQueryDatasetStatus_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDatasetStatus) *pb.DatasetMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DatasetMetadata{}
	out.ETag = direct.ValueOf(in.Etag)
	out.CreationTime = direct.UnixMillisToTime(direct.ValueOf(in.CreationTime))
	out.LastModifiedTime = direct.UnixMillisToTime(direct.ValueOf(in.LastModifiedTime))
	// The full dataset ID in the form projectID:datasetID
	if in.SelfLink != nil {
		selfLink := strings.Trim(direct.ValueOf(in.SelfLink), "https://bigquery.googleapis.com/bigquery/v2/")
		tokens := strings.Split(selfLink, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "datasets" {
			out.FullID = fmt.Sprintf("%s:%s", tokens[1], tokens[3])
		}
	}
	return out
}
func AccessEntry_ToProto(mapCtx *direct.MapContext, in *krm.Access) *pb.AccessEntry {
	if in == nil {
		return nil
	}
	out := &pb.AccessEntry{}
	out.Role = pb.AccessRole(direct.ValueOf(in.Role))
	switch {
	case in.Domain != nil:
		out.Entity = direct.ValueOf(in.Domain)
		out.EntityType = pb.DomainEntity
	case in.GroupByEmail != nil:
		out.Entity = direct.ValueOf(in.GroupByEmail)
		out.EntityType = pb.GroupEmailEntity
	case in.UserByEmail != nil:
		out.Entity = direct.ValueOf(in.UserByEmail)
		out.EntityType = pb.UserEmailEntity
	case in.SpecialGroup != nil:
		out.Entity = direct.ValueOf(in.SpecialGroup)
		out.EntityType = pb.SpecialGroupEntity
	case in.View != nil:
		out.View = TableReference_ToProto(mapCtx, in.View)
		out.EntityType = pb.ViewEntity
	case in.IamMember != nil:
		out.Entity = direct.ValueOf(in.IamMember)
		out.EntityType = pb.IAMMemberEntity
	case in.Routine != nil:
		out.Routine = RoutineReference_ToProto(mapCtx, in.Routine)
		out.EntityType = pb.ViewEntity
	case in.Dataset != nil:
		out.Dataset = DatasetAccessEntry_ToProto(mapCtx, in.Dataset)
		out.EntityType = pb.ViewEntity
	}
	return out
}
func AccessEntry_FromProto(mapCtx *direct.MapContext, in *pb.AccessEntry) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.Role = direct.LazyPtr(string(in.Role))
	switch in.EntityType {
	case pb.DomainEntity:
		out.Domain = direct.LazyPtr(in.Entity)
	case pb.GroupEmailEntity:
		out.GroupByEmail = direct.LazyPtr(in.Entity)
	case pb.UserEmailEntity:
		out.UserByEmail = direct.LazyPtr(in.Entity)
	case pb.SpecialGroupEntity:
		out.SpecialGroup = direct.LazyPtr(in.Entity)
	case pb.ViewEntity:
		out.View = TableReference_FromProto(mapCtx, in.View)
	case pb.IAMMemberEntity:
		out.IamMember = direct.LazyPtr(in.Entity)
	case pb.RoutineEntity:
		out.Routine = RoutineReference_FromProto(mapCtx, in.Routine)
	case pb.DatasetEntity:
		out.Dataset = DatasetAccessEntry_FromProto(mapCtx, in.Dataset)
	}
	return out
}
func BigQueryDataset_ToMetadataToUpdate(mapCtx *direct.MapContext, in *pb.DatasetMetadata, updatePaths []string) *pb.DatasetMetadataToUpdate {
	if in == nil {
		return nil
	}
	out := &pb.DatasetMetadataToUpdate{}
	acccessList := []*pb.AccessEntry{}
	for _, access := range in.Access {
		acccessList = append(acccessList, access)
	}
	out.Access = acccessList
	if in.DefaultEncryptionConfig != nil {
		out.DefaultEncryptionConfig = &pb.EncryptionConfig{
			KMSKeyName: in.DefaultEncryptionConfig.KMSKeyName,
		}
	}
	// if the value to explicitly set to empty in the update request, we set the value.
	// Otherwise, we drop the value.
	if in.DefaultCollation != "" || slices.Contains(updatePaths, "default_collation") {
		out.DefaultCollation = in.DefaultCollation
	}
	if in.DefaultPartitionExpiration != 0 || slices.Contains(updatePaths, "default_partition_expiration") {
		out.DefaultPartitionExpiration = in.DefaultPartitionExpiration
	}
	if in.DefaultTableExpiration != 0 || slices.Contains(updatePaths, "default_table_expiration") {
		out.DefaultTableExpiration = in.DefaultTableExpiration
	}
	if in.Description != "" || slices.Contains(updatePaths, "description") {
		out.Description = in.Description
	}
	if in.MaxTimeTravel != 0 || slices.Contains(updatePaths, "max_time_travel") {
		out.MaxTimeTravel = in.MaxTimeTravel
	}
	out.IsCaseInsensitive = in.IsCaseInsensitive
	if in.Name != "" || slices.Contains(updatePaths, "friendly_name") {
		out.Name = in.Name
	}
	if in.StorageBillingModel != "" || slices.Contains(updatePaths, "storage_billing_model") {
		out.StorageBillingModel = in.StorageBillingModel
	}
	return out
}
func DatasetAccessEntry_FromProto(mapCtx *direct.MapContext, in *pb.DatasetAccessEntry) *krm.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.DatasetAccessEntry{}
	out.Dataset = DatasetReference_FromProto(mapCtx, in.Dataset)
	for _, targetType := range in.TargetTypes {
		out.TargetTypes = append(out.TargetTypes, targetType)
	}
	return out
}
func DatasetAccessEntry_ToProto(mapCtx *direct.MapContext, in *krm.DatasetAccessEntry) *pb.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &pb.DatasetAccessEntry{}
	out.Dataset = DatasetReference_ToProto(mapCtx, in.Dataset)
	out.TargetTypes = make([]string, len(in.TargetTypes))
	for _, targetType := range in.TargetTypes {
		out.TargetTypes = append(out.TargetTypes, targetType)
	}
	return out
}
func DatasetReference_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatasetReference {
	if in == nil {
		return nil
	}
	out := &krm.DatasetReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetID)
	out.ProjectId = direct.LazyPtr(in.ProjectID)
	return out
}
func DatasetReference_ToProto(mapCtx *direct.MapContext, in *krm.DatasetReference) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	out.DatasetID = direct.ValueOf(in.DatasetId)
	out.ProjectID = direct.ValueOf(in.ProjectId)
	return out
}
func DatasetSpec_ToExternalDatasetReference(mapCtx *direct.MapContext, in *krm.BigQueryDatasetSpec) *pb.ExternalDatasetReference {
	// **NOTYET**
	// There are no matching fields in KRM for now.
	// out.Connection
	// out.ExternalSource
	return nil
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	if in.KMSKeyRef != nil {
		out.KMSKeyName = in.KMSKeyRef.External
	}
	return out
}

func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	out.KMSKeyRef = &kmsv1beta1.KMSKeyRef_OneOf{
		External: in.KMSKeyName,
	}
	return out
}
func RoutineReference_FromProto(mapCtx *direct.MapContext, in *pb.Routine) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetID)
	out.ProjectId = direct.LazyPtr(in.ProjectID)
	out.RoutineId = direct.LazyPtr(in.RoutineID)
	return out
}
func RoutineReference_ToProto(mapCtx *direct.MapContext, in *krm.RoutineReference) *pb.Routine {
	if in == nil {
		return nil
	}
	out := &pb.Routine{}
	out.DatasetID = direct.ValueOf(in.DatasetId)
	out.ProjectID = direct.ValueOf(in.ProjectId)
	out.RoutineID = direct.ValueOf(in.RoutineId)
	return out
}
func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetID)
	out.ProjectId = direct.LazyPtr(in.ProjectID)
	out.TableId = direct.LazyPtr(in.TableID)
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.DatasetID = direct.ValueOf(in.DatasetId)
	out.ProjectID = direct.ValueOf(in.ProjectId)
	out.TableID = direct.ValueOf(in.TableId)
	return out
}
