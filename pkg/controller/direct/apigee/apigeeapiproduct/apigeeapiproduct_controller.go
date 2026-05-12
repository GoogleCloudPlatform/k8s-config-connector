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

package apigeeapiproduct

import (
	"context"
	"fmt"

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
	registry.RegisterModel(krm.ApigeeApiProductGVK, NewApigeeApiProductModel)
}

func NewApigeeApiProductModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeApiProduct{config: config}, nil
}

var _ directbase.Model = &modelApigeeApiProduct{}

type modelApigeeApiProduct struct {
	config *config.ControllerConfig
}

func (m *modelApigeeApiProduct) client(ctx context.Context) (*api.OrganizationsApiproductsService, error) {
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

func (m *modelApigeeApiProduct) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeApiProduct{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeApiProductIdentity)

	productsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApigeeApiProductAdapter{
		id:             id,
		k8sClient:      reader,
		productsClient: productsClient,
		desired:        obj,
	}, nil
}

func (m *modelApigeeApiProduct) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ApigeeApiProductAdapter struct {
	id             *krm.ApigeeApiProductIdentity
	k8sClient      client.Reader
	productsClient *api.OrganizationsApiproductsService
	desired        *krm.ApigeeApiProduct
	actual         *api.GoogleCloudApigeeV1ApiProduct
}

var _ directbase.Adapter = &ApigeeApiProductAdapter{}

func (a *ApigeeApiProductAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeApiProduct", "name", a.id)

	actual, err := a.productsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeApiProduct %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ApigeeApiProductAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeApiProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	resource := ApigeeApiProductSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.Apiproduct

	parent := fmt.Sprintf("organizations/%s", a.id.Organization)

	created, err := a.productsClient.Create(parent, resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeApiProduct %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ApigeeApiProduct", "name", a.id)

	status := &krm.ApigeeApiProductStatus{}
	status.ObservedState = ApigeeApiProductObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ApigeeApiProductAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeApiProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	resource := ApigeeApiProductSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updated, err := a.productsClient.Update(a.id.String(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ApigeeApiProduct %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ApigeeApiProduct", "name", a.id)

	status := &krm.ApigeeApiProductStatus{}
	status.ObservedState = ApigeeApiProductObservedState_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ApigeeApiProductAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeApiProduct{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeApiProductSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &apigeev1beta1.ApigeeOrganizationRef{External: fmt.Sprintf("organizations/%s", a.id.Organization)}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Apiproduct)
	u.SetGroupVersionKind(krm.ApigeeApiProductGVK)

	u.Object = uObj
	return u, nil
}

func (a *ApigeeApiProductAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeApiProduct", "name", a.id)

	_, err := a.productsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ApigeeApiProduct, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeApiProduct %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeApiProduct", "name", a.id)

	return true, nil
}
