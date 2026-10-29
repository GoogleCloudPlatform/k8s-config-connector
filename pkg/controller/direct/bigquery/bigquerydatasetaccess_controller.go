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
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryDatasetAccessGVK, NewDatasetAccessModel)
}

func NewDatasetAccessModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &datasetAccessModel{config: config}, nil
}

type datasetAccessModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &datasetAccessModel{}

func (m *datasetAccessModel) datasetsService(ctx context.Context) (*api.DatasetsService, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpService, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigQuery client: %w", err)
	}
	return gcpService.Datasets, nil
}

func (m *datasetAccessModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryDatasetAccess{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	accessID, ok := id.(*krm.BigQueryDatasetAccessIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	datasetsService, err := m.datasetsService(ctx)
	if err != nil {
		return nil, err
	}

	return &datasetAccessAdapter{
		datasetsService: datasetsService,
		id:              accessID,
		desired:         obj,
	}, nil
}

func (m *datasetAccessModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type datasetAccessAdapter struct {
	datasetsService  *api.DatasetsService
	id               *krm.BigQueryDatasetAccessIdentity
	desired          *krm.BigQueryDatasetAccess
	actual           *api.DatasetAccess
	apiUpdatedMember bool
}

var _ directbase.Adapter = &datasetAccessAdapter{}

func canonicalRole(r string) string {
	switch r {
	case "roles/bigquery.dataOwner":
		return "OWNER"
	case "roles/bigquery.dataEditor":
		return "WRITER"
	case "roles/bigquery.dataViewer":
		return "READER"
	default:
		return r
	}
}

func matchesAccessEntry(entry *api.DatasetAccess, desired *api.DatasetAccess, desiredIamMember string) (bool, bool) {
	if entry == nil || desired == nil {
		return false, false
	}

	// 1. View
	if desired.View != nil {
		if entry.View == nil {
			return false, false
		}
		if entry.View.ProjectId == desired.View.ProjectId &&
			entry.View.DatasetId == desired.View.DatasetId &&
			entry.View.TableId == desired.View.TableId {
			return true, false
		}
		return false, false
	}

	// 2. Dataset
	if desired.Dataset != nil {
		if entry.Dataset == nil || entry.Dataset.Dataset == nil || desired.Dataset.Dataset == nil {
			return false, false
		}
		if entry.Dataset.Dataset.ProjectId == desired.Dataset.Dataset.ProjectId &&
			entry.Dataset.Dataset.DatasetId == desired.Dataset.Dataset.DatasetId &&
			reflect.DeepEqual(entry.Dataset.TargetTypes, desired.Dataset.TargetTypes) {
			return true, false
		}
		return false, false
	}

	// 3. For all other access entries, the role must match.
	if canonicalRole(entry.Role) != canonicalRole(desired.Role) {
		return false, false
	}

	if desired.Domain != "" {
		if strings.EqualFold(entry.Domain, desired.Domain) {
			return true, false
		}
	}
	if desired.GroupByEmail != "" {
		if strings.EqualFold(entry.GroupByEmail, desired.GroupByEmail) {
			return true, false
		}
	}
	if desired.UserByEmail != "" {
		if strings.EqualFold(entry.UserByEmail, desired.UserByEmail) {
			return true, false
		}
	}
	if desired.SpecialGroup != "" {
		if entry.SpecialGroup == desired.SpecialGroup {
			return true, false
		}
	}

	if desiredIamMember != "" {
		if entry.IamMember != "" && strings.EqualFold(entry.IamMember, desiredIamMember) {
			return true, false
		}
		parts := strings.SplitN(desiredIamMember, ":", 2)
		memberValue := desiredIamMember
		if len(parts) == 2 {
			memberValue = parts[1]
		}
		if entry.UserByEmail != "" && strings.EqualFold(entry.UserByEmail, memberValue) {
			return true, true
		}
		if entry.GroupByEmail != "" && strings.EqualFold(entry.GroupByEmail, memberValue) {
			return true, true
		}
		if entry.Domain != "" && strings.EqualFold(entry.Domain, memberValue) {
			return true, true
		}
		if entry.SpecialGroup != "" && strings.EqualFold(entry.SpecialGroup, memberValue) {
			return true, true
		}
	}

	return false, false
}

func (a *datasetAccessAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryDatasetAccess", "id", a.id.String())

	dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataset %s/%s: %w", a.id.Project, a.id.Dataset, err)
	}

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	for _, entry := range dataset.Access {
		if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
			a.actual = entry
			a.apiUpdatedMember = apiUpdated
			return true, nil
		}
	}

	return false, nil
}

func (a *datasetAccessAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigQueryDatasetAccess", "id", a.id.String())

	dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting dataset %s/%s for access creation: %w", a.id.Project, a.id.Dataset, err)
	}

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	for _, entry := range dataset.Access {
		if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
			a.actual = entry
			a.apiUpdatedMember = apiUpdated
			return a.updateStatus(ctx, op)
		}
	}

	newAccess := append(dataset.Access, desiredAccess)
	patchReq := &api.Dataset{
		Access: newAccess,
	}

	patchedDataset, err := a.datasetsService.Patch(a.id.Project, a.id.Dataset, patchReq).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("patching dataset %s/%s with new access entry: %w", a.id.Project, a.id.Dataset, err)
	}

	for _, entry := range patchedDataset.Access {
		if matched, apiUpdated := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
			a.actual = entry
			a.apiUpdatedMember = apiUpdated
			break
		}
	}
	if a.actual == nil {
		a.actual = desiredAccess
	}

	return a.updateStatus(ctx, op)
}

func (a *datasetAccessAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigQueryDatasetAccess", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	matched, apiUpdated := matchesAccessEntry(a.actual, desiredAccess, desiredIamMember)
	if !matched {
		return fmt.Errorf("BigQueryDatasetAccess is immutable and cannot be updated")
	}
	a.apiUpdatedMember = apiUpdated

	return a.updateStatus(ctx, op)
}

func (a *datasetAccessAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigQueryDatasetAccess", "id", a.id.String())

	dataset, err := a.datasetsService.Get(a.id.Project, a.id.Dataset).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("getting dataset %s/%s for access deletion: %w", a.id.Project, a.id.Dataset, err)
	}

	mapCtx := &direct.MapContext{}
	desiredAccess := BigQueryDatasetAccessSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	desiredIamMember := direct.ValueOf(a.desired.Spec.IamMember)

	foundIdx := -1
	for idx, entry := range dataset.Access {
		if matched, _ := matchesAccessEntry(entry, desiredAccess, desiredIamMember); matched {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return true, nil
	}

	updatedAccess := make([]*api.DatasetAccess, 0, len(dataset.Access)-1)
	for idx, entry := range dataset.Access {
		if idx != foundIdx {
			updatedAccess = append(updatedAccess, entry)
		}
	}

	patchReq := &api.Dataset{
		Access: updatedAccess,
	}

	_, err = a.datasetsService.Patch(a.id.Project, a.id.Dataset, patchReq).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("patching dataset %s/%s to delete access entry: %w", a.id.Project, a.id.Dataset, err)
	}

	return true, nil
}

func (a *datasetAccessAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryDatasetAccessStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.apiUpdatedMember {
		status.ApiUpdatedMember = direct.LazyPtr(true)
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *datasetAccessAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryDatasetAccess{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryDatasetAccessSpec_FromAPI(mapCtx, a.actual))
	obj.Spec.ProjectRef = parent.ProjectRef{External: a.id.Project}
	obj.Spec.DatasetId = a.id.Dataset
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}
