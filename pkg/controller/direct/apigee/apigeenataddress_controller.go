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
	registry.RegisterModel(krm.ApigeeNATAddressGVK, NewApigeeNATAddressModel)
}

func NewApigeeNATAddressModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeNATAddress{config: config}, nil
}

var _ directbase.Model = &modelApigeeNATAddress{}

type modelApigeeNATAddress struct {
	config *config.ControllerConfig
}

func (m *modelApigeeNATAddress) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeNATAddress{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeNATAddressIdentity)

	// Get apigee GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeNATAddressAdapter{
		id:                 id,
		k8sClient:          reader,
		natAddressesClient: gcpClient.instancesNatAddressesClient(),
		operationsClient:   gcpClient.operationsClient(),
		desired:            obj,
	}, nil
}

func (m *modelApigeeNATAddress) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ApigeeNATAddressIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeNATAddressAdapter{
		id:                 id,
		natAddressesClient: gcpClient.instancesNatAddressesClient(),
		operationsClient:   gcpClient.operationsClient(),
	}, nil
}

type ApigeeNATAddressAdapter struct {
	id                 *krm.ApigeeNATAddressIdentity
	k8sClient          client.Reader
	natAddressesClient *api.OrganizationsInstancesNatAddressesService
	operationsClient   *api.OrganizationsOperationsService
	desired            *krm.ApigeeNATAddress
	actual             *api.GoogleCloudApigeeV1NatAddress
}

var _ directbase.Adapter = &ApigeeNATAddressAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeNATAddressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeNATAddress", "name", a.id)

	resource, err := a.natAddressesClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeNATAddress %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeNATAddressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeNATAddress", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ApigeeNATAddressSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.NatAddress

	op, err := a.natAddressesClient.Create(a.id.ParentString(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeNATAddress %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ApigeeNATAddress %s waiting creation: %w", a.id, err)
	}

	created, err := a.natAddressesClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeNATAddress %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ApigeeNATAddress", "name", a.id)

	status := ApigeeNATAddressStatus_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeNATAddressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeNATAddress", "id", a.id)
	mapCtx := &direct.MapContext{}

	// There are no mutable fields in the GCP ApigeeNATAddress API.
	// So, we will only update the KRM status.
	status := ApigeeNATAddressStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeNATAddressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeNATAddress{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeNATAddressSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceId = a.id.ParentString()
	obj.Spec.ResourceID = &a.actual.Name

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.NatAddress)
	u.SetGroupVersionKind(krm.ApigeeNATAddressGVK)

	u.Object = uObj
	return u, nil
}

// Delete deletes the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeNATAddressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeNATAddress", "id", a.id)

	op, err := a.natAddressesClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeNATAddress, assuming it was already deleted", "id", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeNATAddress %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeNATAddress", "id", a.id)

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete ApigeeNATAddress %s: %w", a.id, err)
	}
	return true, nil
}
