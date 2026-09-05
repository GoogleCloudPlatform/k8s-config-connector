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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/apigee/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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

func (m *modelApigeeAPIProduct) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.ApigeeAPIProduct{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeAPIProductIdentity)

	// Create a deep copy of obj to prevent mutating the original object,
	// keeping it clean for any future maintainers.
	desired := obj.DeepCopy()
	desired.Name = id.ID()

	return &ApigeeAPIProductAdapter{
		id:                id,
		desired:           desired,
		apiproductsClient: gcpClient.apiproductsClient(),
	}, nil
}

func (m *modelApigeeAPIProduct) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ApigeeAPIProductAdapter struct {
	id                *krm.ApigeeAPIProductIdentity
	desired           *krm.ApigeeAPIProduct
	actual            *api.GoogleCloudApigeeV1ApiProduct
	apiproductsClient *api.OrganizationsApiproductsService
}

var _ directbase.Adapter = &ApigeeAPIProductAdapter{}

func (a *ApigeeAPIProductAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeAPIProduct", "name", a.id)

	apiproduct, err := a.apiproductsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeAPIProduct %q: %w", a.id, err)
	}

	a.actual = apiproduct
	return true, nil
}

func (a *ApigeeAPIProductAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeAPIProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	req := ApigeeAPIProductSpec_ToApi(mapCtx, &a.desired.Spec, a.desired.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	created, err := a.apiproductsClient.Create(a.id.ParentString(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeAPIProduct %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ApigeeAPIProduct", "ApigeeAPIProduct", created)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ApigeeAPIProductAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeAPIProduct", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredAPI := ApigeeAPIProductSpec_ToApi(mapCtx, &a.desired.Spec, a.desired.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffResults, err := compareResource(ctx, a.actual, desiredAPI)
	if err != nil {
		return fmt.Errorf("comparing ApigeeAPIProduct: %w", err)
	}

	latest := a.actual
	if !diffResults.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id)

		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		updated, err := a.apiproductsClient.Update(a.id.String(), desiredAPI).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating ApigeeAPIProduct %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ApigeeAPIProduct", "ApigeeAPIProduct", updated)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ApigeeAPIProductAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeAPIProduct{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeAPIProductSpec_FromApi(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &apigeev1beta1.ApigeeOrganizationRef{External: a.id.ParentString()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ApigeeAPIProductGVK)

	u.Object = uObj
	return u, nil
}

func (a *ApigeeAPIProductAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeAPIProduct", "name", a.id)

	_, err := a.apiproductsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ApigeeAPIProduct, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeAPIProduct %s: %w", a.id, err)
	}

	return true, nil
}

func (a *ApigeeAPIProductAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.GoogleCloudApigeeV1ApiProduct) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ApigeeAPIProductStatus{}
	status.ObservedState = ApigeeAPIProductObservedState_FromApi(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareResource(ctx context.Context, actual, desired *api.GoogleCloudApigeeV1ApiProduct) (*structuredreporting.Diff, error) {
	mapCtx := &direct.MapContext{}
	maskedActualSpec := ApigeeAPIProductSpec_FromApi(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	maskedActual := ApigeeAPIProductSpec_ToApi(mapCtx, maskedActualSpec, desired.Name)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	diffs, _, err := diffs.GoogleAPI.Diff(ctx, maskedActual, desired)
	return diffs, err
}
