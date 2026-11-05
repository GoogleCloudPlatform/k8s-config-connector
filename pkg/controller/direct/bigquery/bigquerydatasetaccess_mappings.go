// Copyright 2026 Google LLC
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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/bigquery/v2"
)

func BigQueryDatasetAccessSpec_FromAPI(mapCtx *direct.MapContext, in *api.DatasetAccess) *krm.BigQueryDatasetAccessSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetAccessSpec{}
	out.Dataset = BigQueryDatasetAccessDatasetAccessEntry_FromAPI(mapCtx, in.Dataset)
	out.Domain = direct.LazyPtr(in.Domain)
	out.GroupByEmail = direct.LazyPtr(in.GroupByEmail)
	out.IamMember = direct.LazyPtr(in.IamMember)
	out.Role = direct.LazyPtr(in.Role)
	out.SpecialGroup = direct.LazyPtr(in.SpecialGroup)
	out.UserByEmail = direct.LazyPtr(in.UserByEmail)
	out.View = BigQueryDatasetAccessView_FromAPI(mapCtx, in.View)
	return out
}

func BigQueryDatasetAccessSpec_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetAccessSpec) *api.DatasetAccess {
	if in == nil {
		return nil
	}
	out := &api.DatasetAccess{}
	out.Dataset = BigQueryDatasetAccessDatasetAccessEntry_ToAPI(mapCtx, in.Dataset)
	out.Domain = direct.ValueOf(in.Domain)
	out.GroupByEmail = direct.ValueOf(in.GroupByEmail)
	out.IamMember = direct.ValueOf(in.IamMember)
	out.Role = direct.ValueOf(in.Role)
	out.SpecialGroup = direct.ValueOf(in.SpecialGroup)
	out.UserByEmail = direct.ValueOf(in.UserByEmail)
	out.View = BigQueryDatasetAccessView_ToAPI(mapCtx, in.View)
	return out
}

func BigQueryDatasetAccessStatus_FromAPI(mapCtx *direct.MapContext, in *api.DatasetAccess) *krm.BigQueryDatasetAccessStatus {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetAccessStatus{}
	return out
}

func BigQueryDatasetAccessStatus_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetAccessStatus) *api.DatasetAccess {
	if in == nil {
		return nil
	}
	out := &api.DatasetAccess{}
	return out
}

func BigQueryDatasetAccessDatasetAccessEntry_FromAPI(mapCtx *direct.MapContext, in *api.DatasetAccessEntry) *krm.BigQueryDatasetAccessDatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetAccessDatasetAccessEntry{}
	if in.Dataset != nil {
		if ds := BigQueryDatasetAccessDataset_FromAPI(mapCtx, in.Dataset); ds != nil {
			out.Dataset = *ds
		}
	}
	out.TargetTypes = in.TargetTypes
	return out
}

func BigQueryDatasetAccessDatasetAccessEntry_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetAccessDatasetAccessEntry) *api.DatasetAccessEntry {
	if in == nil {
		return nil
	}
	out := &api.DatasetAccessEntry{}
	out.Dataset = BigQueryDatasetAccessDataset_ToAPI(mapCtx, &in.Dataset)
	out.TargetTypes = in.TargetTypes
	return out
}

func BigQueryDatasetAccessView_FromAPI(mapCtx *direct.MapContext, in *api.TableReference) *krm.BigQueryDatasetAccessView {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetAccessView{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	out.TableId = in.TableId
	return out
}

func BigQueryDatasetAccessView_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetAccessView) *api.TableReference {
	if in == nil {
		return nil
	}
	out := &api.TableReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	out.TableId = in.TableId
	return out
}

func BigQueryDatasetAccessDataset_FromAPI(mapCtx *direct.MapContext, in *api.DatasetReference) *krm.BigQueryDatasetAccessDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDatasetAccessDataset{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}

func BigQueryDatasetAccessDataset_ToAPI(mapCtx *direct.MapContext, in *krm.BigQueryDatasetAccessDataset) *api.DatasetReference {
	if in == nil {
		return nil
	}
	out := &api.DatasetReference{}
	out.DatasetId = in.DatasetId
	out.ProjectId = in.ProjectId
	return out
}
