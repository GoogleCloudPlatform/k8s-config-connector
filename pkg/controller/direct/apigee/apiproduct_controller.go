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

package apigee

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/apigee/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ApigeeAPIProductGVK, NewApigeeAPIProductModel)
}

func NewApigeeAPIProductModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeAPIProduct{config: config}, nil
}

var _ directbase.Model = &modelApigeeAPIProduct{}

type modelApigeeAPIProduct struct {
	config *config.ControllerConfig
}

func (m *modelApigeeAPIProduct) client(ctx context.Context) (*api.OrganizationsApiproductsService, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building gcp service client: %w", err)
	}
	return api.NewOrganizationsApiproductsService(gcpClient), nil
}

func (m *modelApigeeAPIProduct) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeAPIProduct{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeAPIProductIdentity)

	productsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApigeeAPIProductAdapter{
		id:             id,
		k8sClient:      reader,
		productsClient: productsClient,
		desired:        obj,
	}, nil
}

func (m *modelApigeeAPIProduct) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ApigeeAPIProductAdapter struct {
	id             *krm.ApigeeAPIProductIdentity
	k8sClient      client.Reader
	productsClient *api.OrganizationsApiproductsService
	desired        *krm.ApigeeAPIProduct
	actual         *api.GoogleCloudApigeeV1ApiProduct
}

var _ directbase.Adapter = &ApigeeAPIProductAdapter{}

func (a *ApigeeAPIProductAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeAPIProduct", "name", a.id)

	actual, err := a.productsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeAPIProduct %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ApigeeAPIProductAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeAPIProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	resource := ApigeeAPIProductSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.Apiproduct

	parent := fmt.Sprintf("organizations/%s", a.id.Organization)

	created, err := a.productsClient.Create(parent, resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeAPIProduct %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ApigeeAPIProduct", "name", a.id)

	status := &krm.ApigeeAPIProductStatus{}
	status.ObservedState = ApigeeAPIProductObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ApigeeAPIProductAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeAPIProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	resource := ApigeeAPIProductSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	hasChanges := false
	if resource.DisplayName != a.actual.DisplayName {
		hasChanges = true
	}
	if resource.Description != a.actual.Description {
		hasChanges = true
	}
	if resource.ApprovalType != a.actual.ApprovalType {
		hasChanges = true
	}
	if resource.Quota != a.actual.Quota {
		hasChanges = true
	}
	if resource.QuotaCounterScope != a.actual.QuotaCounterScope {
		hasChanges = true
	}
	if resource.QuotaInterval != a.actual.QuotaInterval {
		hasChanges = true
	}
	if resource.QuotaTimeUnit != a.actual.QuotaTimeUnit {
		hasChanges = true
	}
	if !reflect.DeepEqual(resource.ApiResources, a.actual.ApiResources) {
		hasChanges = true
	}
	if !reflect.DeepEqual(resource.Environments, a.actual.Environments) {
		hasChanges = true
	}
	if !reflect.DeepEqual(resource.Proxies, a.actual.Proxies) {
		hasChanges = true
	}
	if !reflect.DeepEqual(resource.Scopes, a.actual.Scopes) {
		hasChanges = true
	}
	if !reflect.DeepEqual(resource.Attributes, a.actual.Attributes) {
		hasChanges = true
	}

	if !hasChanges {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ApigeeAPIProductStatus{}
		status.ObservedState = ApigeeAPIProductObservedState_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	updated, err := a.productsClient.Update(a.id.String(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ApigeeAPIProduct %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ApigeeAPIProduct", "name", a.id)

	status := &krm.ApigeeAPIProductStatus{}
	status.ObservedState = ApigeeAPIProductObservedState_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ApigeeAPIProductAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeAPIProduct{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeAPIProductSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &apigeev1beta1.ApigeeOrganizationRef{External: fmt.Sprintf("organizations/%s", a.id.Organization)}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Apiproduct)
	u.SetGroupVersionKind(krm.ApigeeAPIProductGVK)

	u.Object = uObj
	return u, nil
}

func (a *ApigeeAPIProductAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeAPIProduct", "name", a.id)

	_, err := a.productsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ApigeeAPIProduct, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeAPIProduct %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeAPIProduct", "name", a.id)

	return true, nil
}
