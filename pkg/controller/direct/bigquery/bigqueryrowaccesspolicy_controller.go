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
	"sort"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryRowAccessPolicyGVK, NewBigQueryRowAccessPolicyModel)
}

func NewBigQueryRowAccessPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &bigQueryRowAccessPolicyModel{config: config}, nil
}

type bigQueryRowAccessPolicyModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &bigQueryRowAccessPolicyModel{}

func (m *bigQueryRowAccessPolicyModel) rowAccessPoliciesService(ctx context.Context) (*bigquery.RowAccessPoliciesService, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpService, err := bigquery.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigQuery client: %w", err)
	}
	return gcpService.RowAccessPolicies, nil
}

func (m *bigQueryRowAccessPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryRowAccessPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	rawID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := rawID.(*krm.BigQueryRowAccessPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", rawID)
	}

	service, err := m.rowAccessPoliciesService(ctx)
	if err != nil {
		return nil, err
	}

	return &bigQueryRowAccessPolicyAdapter{
		service: service,
		id:      id,
		desired: obj,
	}, nil
}

func (m *bigQueryRowAccessPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type bigQueryRowAccessPolicyAdapter struct {
	service *bigquery.RowAccessPoliciesService
	id      *krm.BigQueryRowAccessPolicyIdentity
	desired *krm.BigQueryRowAccessPolicy
	actual  *bigquery.RowAccessPolicy
}

var _ directbase.Adapter = &bigQueryRowAccessPolicyAdapter{}

func (a *bigQueryRowAccessPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting RowAccessPolicy", "name", a.id.String())

	actual, err := a.service.Get(a.id.Project, a.id.Dataset, a.id.Table, a.id.RowAccessPolicy).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting RowAccessPolicy %q: %w", a.id.String(), err)
	}

	// BigQuery RowAccessPolicies.Get API does not populate the grantees field in RowAccessPolicy response.
	// Instead, grantees are stored as IAM policy bindings with role "roles/bigquery.filteredDataViewer".
	// Query GetIamPolicy to populate actual.Grantees for drift detection.
	iamResource := fmt.Sprintf("projects/%s/datasets/%s/tables/%s/rowAccessPolicies/%s", a.id.Project, a.id.Dataset, a.id.Table, a.id.RowAccessPolicy)
	iamPolicy, err := a.service.GetIamPolicy(iamResource, &bigquery.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("getting IAM policy for RowAccessPolicy %q: %w", a.id.String(), err)
	}
	var grantees []string
	for _, b := range iamPolicy.Bindings {
		if b.Role == "roles/bigquery.filteredDataViewer" {
			grantees = append(grantees, b.Members...)
		}
	}
	sort.Strings(grantees)
	actual.Grantees = grantees

	a.actual = actual
	return true, nil
}

func (a *bigQueryRowAccessPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating RowAccessPolicy", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := BigQueryRowAccessPolicySpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if desired.RowAccessPolicyReference == nil {
		desired.RowAccessPolicyReference = &bigquery.RowAccessPolicyReference{}
	}
	desired.RowAccessPolicyReference.ProjectId = a.id.Project
	desired.RowAccessPolicyReference.DatasetId = a.id.Dataset
	desired.RowAccessPolicyReference.TableId = a.id.Table
	desired.RowAccessPolicyReference.PolicyId = a.id.RowAccessPolicy

	created, err := a.service.Insert(a.id.Project, a.id.Dataset, a.id.Table, desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating RowAccessPolicy %q: %w", a.id.String(), err)
	}

	status := BigQueryRowAccessPolicyStatus_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(createOp.GetUnstructured(), status)
}

func (a *bigQueryRowAccessPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating RowAccessPolicy", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := BigQueryRowAccessPolicySpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	diff := false

	if a.actual.FilterPredicate != desired.FilterPredicate {
		report.AddField("filterPredicate", a.actual.FilterPredicate, desired.FilterPredicate)
		diff = true
	}
	actualGranteesSet := sets.New(a.actual.Grantees...)
	desiredGranteesSet := sets.New(desired.Grantees...)
	if !actualGranteesSet.Equal(desiredGranteesSet) {
		report.AddField("grantees", a.actual.Grantees, desired.Grantees)
		diff = true
	}

	if !diff {
		log.V(2).Info("no diff detected for RowAccessPolicy", "name", a.id.String())
		status := BigQueryRowAccessPolicyStatus_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return setStatus(updateOp.GetUnstructured(), status)
	}

	structuredreporting.ReportDiff(ctx, report)

	if desired.RowAccessPolicyReference == nil {
		desired.RowAccessPolicyReference = &bigquery.RowAccessPolicyReference{}
	}
	desired.RowAccessPolicyReference.ProjectId = a.id.Project
	desired.RowAccessPolicyReference.DatasetId = a.id.Dataset
	desired.RowAccessPolicyReference.TableId = a.id.Table
	desired.RowAccessPolicyReference.PolicyId = a.id.RowAccessPolicy

	updated, err := a.service.Update(a.id.Project, a.id.Dataset, a.id.Table, a.id.RowAccessPolicy, desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating RowAccessPolicy %q: %w", a.id.String(), err)
	}

	status := BigQueryRowAccessPolicyStatus_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.LazyPtr(a.id.String())
	return setStatus(updateOp.GetUnstructured(), status)
}

func (a *bigQueryRowAccessPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.BigQueryRowAccessPolicy{}
	mapCtx := &direct.MapContext{}
	spec := BigQueryRowAccessPolicySpec_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	tableParent := fmt.Sprintf("projects/%s/datasets/%s/tables/%s", a.id.Project, a.id.Dataset, a.id.Table)
	spec.ParentRef = &bigqueryv1beta1.BigQueryTableRef{
		External: tableParent,
	}
	spec.ResourceID = direct.LazyPtr(a.id.RowAccessPolicy)
	obj.Spec = *spec
	obj.TypeMeta = metav1.TypeMeta{
		APIVersion: krm.GroupVersion.String(),
		Kind:       krm.BigQueryRowAccessPolicyGVK.Kind,
	}

	status := BigQueryRowAccessPolicyStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	obj.Status = *status

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting %T to unstructured: %w", obj, err)
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *bigQueryRowAccessPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting RowAccessPolicy", "name", a.id.String())

	err := a.service.Delete(a.id.Project, a.id.Dataset, a.id.Table, a.id.RowAccessPolicy).Force(true).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting RowAccessPolicy %q: %w", a.id.String(), err)
	}

	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	oldConditions, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	if oldConditions != nil {
		if err := unstructured.SetNestedSlice(status, oldConditions, "conditions"); err != nil {
			return fmt.Errorf("error setting conditions: %w", err)
		}
	}

	u.Object["status"] = status
	return nil
}
