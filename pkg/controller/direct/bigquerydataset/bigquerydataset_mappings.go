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
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/bigquery/v2"
)

func BigQueryDatasetSpec_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetSpec, name string) *api.Dataset {
	if in == nil {
		return nil
	}
	out := &api.Dataset{}
	acccessList := []*api.DatasetAccess{}
	for _, access := range in.Access {
		curAccess := Access_ToAPI(mapCtx, direct.LazyPtr(access))
		acccessList = append(acccessList, curAccess)
	}
	out.Access = acccessList
	out.DefaultCollation = direct.ValueOf(in.DefaultCollation)
	out.DefaultPartitionExpirationMs = direct.ValueOf(in.DefaultPartitionExpirationMs)
	out.DefaultTableExpirationMs = direct.ValueOf(in.DefaultTableExpirationMs)
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_ToAPI(mapCtx, in.DefaultEncryptionConfiguration)
	out.Description = direct.ValueOf(in.Description)
	out.FriendlyName = direct.ValueOf(in.FriendlyName)
	out.DatasetReference = DatasetReference_ToAPI(mapCtx, in, name)
	out.Location = direct.ValueOf(in.Location)
	out.IsCaseInsensitive = direct.ValueOf(in.IsCaseInsensitive)
	if in.MaxTimeTravelHours != nil {
		out.MaxTimeTravelHours, _ = strconv.ParseInt(direct.ValueOf(in.MaxTimeTravelHours), 10, 64)
	}
	out.StorageBillingModel = direct.ValueOf(in.StorageBillingModel)
	return out
}
func BigQueryDatasetSpec_FromAPI(mapCtx *direct.MapContext, in *api.Dataset) *krm.BigQueryDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetSpec{}
	accessList := []krm.Access{}
	for _, access := range in.Access {
		curAccess := Access_FromAPI(mapCtx, access)
		accessList = append(accessList, direct.ValueOf(curAccess))
	}
	out.Access = accessList
	out.DefaultCollation = direct.LazyPtr(in.DefaultCollation)
	out.DefaultPartitionExpirationMs = direct.LazyPtr(in.DefaultPartitionExpirationMs)
	out.DefaultTableExpirationMs = direct.LazyPtr(in.DefaultTableExpirationMs)
	out.DefaultEncryptionConfiguration = EncryptionConfiguration_FromAPI(mapCtx, in.DefaultEncryptionConfiguration)
	out.Description = direct.LazyPtr(in.Description)
	out.FriendlyName = direct.LazyPtr(in.FriendlyName)
	out.Location = direct.LazyPtr(in.Location)
	out.IsCaseInsensitive = direct.LazyPtr(in.IsCaseInsensitive)
	maxTime := strconv.FormatInt(in.MaxTimeTravelHours, 10)
	out.MaxTimeTravelHours = direct.LazyPtr(maxTime)
	out.StorageBillingModel = direct.LazyPtr(in.StorageBillingModel)
	return out
}
func BigQueryDatasetStatus_FromAPI(mapCtx *direct.MapContext, in *api.Dataset) *krm.BigQueryDatasetStatus {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetStatus{}
	out.Etag = direct.LazyPtr(in.Etag)
	out.CreationTime = direct.LazyPtr(in.CreationTime)
	out.LastModifiedTime = direct.LazyPtr(in.LastModifiedTime)
	out.SelfLink = direct.LazyPtr(in.SelfLink)
	return out
}
func BigQueryDatasetStatusObservedState_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetStatus) *api.Dataset {
	if in == nil {
		return nil
	}
	out := &api.Dataset{}
	out.Etag = direct.ValueOf(in.Etag)
	out.CreationTime = direct.ValueOf(in.CreationTime)
	out.LastModifiedTime = direct.ValueOf(in.LastModifiedTime)
	out.SelfLink = direct.ValueOf(in.SelfLink)
	return out
}
func Access_ToAPI(mapCtx *direct.MapContext, in *krm.Access) *api.DatasetAccess {
	if in == nil {
		return nil
	}
	out := &api.DatasetAccess{}
	out.Domain = direct.ValueOf(in.Domain)
	out.GroupByEmail = direct.ValueOf(in.GroupByEmail)
	out.IamMember = direct.ValueOf(in.IamMember)
	out.UserByEmail = direct.ValueOf(in.UserByEmail)
	out.SpecialGroup = direct.ValueOf(in.SpecialGroup)
	out.Role = direct.ValueOf(in.Role)
	out.Dataset = DatasetAccessEntry_ToAPI(mapCtx, in.Dataset)
	out.Routine = RoutineReference_ToAPI(mapCtx, in.Routine)
	out.View = TableReference_ToAPI(mapCtx, in.View)
	return out
}
func Access_FromAPI(mapCtx *direct.MapContext, in *api.DatasetAccess) *krm.Access {
	if in == nil {
		return nil
	}
	out := &krm.Access{}
	out.Domain = direct.LazyPtr(in.Domain)
	out.GroupByEmail = direct.LazyPtr(in.GroupByEmail)
	out.IamMember = direct.LazyPtr(in.IamMember)
	out.UserByEmail = direct.LazyPtr(in.UserByEmail)
	out.SpecialGroup = direct.LazyPtr(in.SpecialGroup)
	out.Role = direct.LazyPtr(in.Role)
	out.Dataset = DatasetAccessEntry_FromAPI(mapCtx, in.Dataset)
	out.Routine = RoutineReference_FromAPI(mapCtx, in.Routine)
	out.View = TableReference_FromAPI(mapCtx, in.View)
	return out
}
func DatasetAccessEntry_FromAPI(mapCtx *direct.MapContext, in *api.DatasetAccessEntry) *krm.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.DatasetAccessEntry{}
	out.Dataset = &krm.DatasetReference{
		DatasetId: direct.LazyPtr(in.Dataset.DatasetId),
		ProjectId: direct.LazyPtr(in.Dataset.ProjectId),
	}
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetAccessEntry_ToAPI(mapCtx *direct.MapContext, in *krm.DatasetAccessEntry) *api.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &api.DatasetAccessEntry{}
	out.Dataset = &api.DatasetReference{
		DatasetId: direct.ValueOf(in.Dataset.DatasetId),
		ProjectId: direct.ValueOf(in.Dataset.ProjectId),
	}
	out.TargetTypes = in.TargetTypes
	return out
}
func DatasetReference_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetSpec, name string) *api.DatasetReference {
	if in == nil {
		return nil
	}
	out := &api.DatasetReference{}
	out.DatasetId = name
	return out
}
func EncryptionConfiguration_ToAPI(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *api.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &api.EncryptionConfiguration{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}

func EncryptionConfiguration_FromAPI(mapCtx *direct.MapContext, in *api.EncryptionConfiguration) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	out.KmsKeyRef = &v1beta1.KMSCryptoKeyRef{
		External: in.KmsKeyName,
	}
	return out
}
func RoutineReference_FromAPI(mapCtx *direct.MapContext, in *api.RoutineReference) *krm.RoutineReference {
	if in == nil {
		return nil
	}
	out := &krm.RoutineReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetId)
	out.ProjectId = direct.LazyPtr(in.ProjectId)
	out.RoutineId = direct.LazyPtr(in.RoutineId)
	return out
}
func RoutineReference_ToAPI(mapCtx *direct.MapContext, in *krm.RoutineReference) *api.RoutineReference {
	if in == nil {
		return nil
	}
	out := &api.RoutineReference{}
	out.DatasetId = direct.ValueOf(in.DatasetId)
	out.ProjectId = direct.ValueOf(in.ProjectId)
	out.RoutineId = direct.ValueOf(in.RoutineId)
	return out
}
func TableReference_FromAPI(mapCtx *direct.MapContext, in *api.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.DatasetId = direct.LazyPtr(in.DatasetId)
	out.ProjectId = direct.LazyPtr(in.ProjectId)
	out.TableId = direct.LazyPtr(in.TableId)
	return out
}
func TableReference_ToAPI(mapCtx *direct.MapContext, in *krm.TableReference) *api.TableReference {
	if in == nil {
		return nil
	}
	out := &api.TableReference{}
	out.DatasetId = direct.ValueOf(in.DatasetId)
	out.ProjectId = direct.ValueOf(in.ProjectId)
	out.TableId = direct.ValueOf(in.TableId)
	return out
}
