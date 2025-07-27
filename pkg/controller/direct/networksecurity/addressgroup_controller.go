// Copyright 2025 Google LLC
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

package networksecurity

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	api "google.golang.org/api/networksecurity/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityAddressGroupGVK, NewGroupModel)
}

func NewGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGroup{config: *config}, nil
}

var _ directbase.Model = &modelGroup{}

type modelGroup struct {
	config config.ControllerConfig
}

func (m *modelGroup) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkSecurityAddressGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAddressGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		if id.Parent().ProjectID != "" {
			config.BillingProject = id.Parent().ProjectID
		} else {
			config.BillingProject = "cnrm-yuhou"
		}
		// Folder and Organization parents are not billing projects
	}

	if obj.Spec.Purpose == nil {
		obj.Spec.Purpose = []string{"DEFAULT"}
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &AddressGroupAdapter{
		id:               id,
		gcpClient:        gcpClient,
		desired:          obj,
		operationsClient: api.NewOrganizationsLocationsOperationsService(gcpClient),
	}, nil
}

func (m *modelGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AddressGroupAdapter struct {
	id               *krm.AddressGroupIdentity
	gcpClient        *api.Service
	desired          *krm.NetworkSecurityAddressGroup
	actual           *api.AddressGroup
	operationsClient *api.OrganizationsLocationsOperationsService
}

var _ directbase.Adapter = &AddressGroupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AddressGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Group", "name", a.id)

	resource, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AddressGroup %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AddressGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AddressGroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityAddressGroupSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	var op *api.Operation
	var err error
	if desired.Spec.OrganizationRef != nil {
		op, err = a.gcpClient.Organizations.Locations.AddressGroups.Create(a.id.Parent().String(), resource).AddressGroupId(a.id.ID()).Context(ctx).Do()
	} else {
		op, err = a.gcpClient.Projects.Locations.AddressGroups.Create(a.id.Parent().String(), resource).AddressGroupId(a.id.ID()).Context(ctx).Do()
	}
	if err != nil {
		return fmt.Errorf("creating AddressGroup %s: %w", a.id, err)
	}

	if err := WaitForOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("error waiting AddressGroup %s creation: %w", a.id, err)
	}

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created AddressGroup %q: %w", a.id, err)
	}

	log.V(2).Info("successfully created AddressGroup", "name", a.id)

	status := &krm.NetworkSecurityAddressGroupStatus{}
	status.ObservedState = NetworkSecurityAddressGroupObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &created.Name
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AddressGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Group", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityAddressGroupSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string
	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if !reflect.DeepEqual(resource.Purpose, a.actual.Purpose) {
		paths = append(paths, "purpose")
	}

	if len(paths) == 0 || reflect.DeepEqual(resource.Items, a.actual.Items) {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.NetworkSecurityAddressGroupStatus{}
		status.ObservedState = NetworkSecurityAddressGroupObservedState_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	// Get update mask
	sort.Strings(paths)
	updateMask := strings.Join(paths, ",")

	var op *api.Operation
	var err error
	if desired.Spec.OrganizationRef != nil {
		op, err = a.gcpClient.Organizations.Locations.AddressGroups.Patch(a.id.String(), resource).UpdateMask(updateMask).Context(ctx).Do()
	} else {
		op, err = a.gcpClient.Projects.Locations.AddressGroups.Patch(a.id.String(), resource).UpdateMask(updateMask).Context(ctx).Do()
	}

	if err != nil {
		return fmt.Errorf("updating AddressGroup %s: %w", a.id, err)
	}

	if err := WaitForOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("error waiting AddressGroup %s update: %w", a.id, err)
	}

	// Get update items
	var addItems, removeItems []string
	beforeSet := make(map[string]bool)
	if a.actual.Items != nil {
		for _, item := range a.actual.Items {
			beforeSet[item] = true
		}
	}

	afterSet := make(map[string]bool)
	if resource.Items != nil {
		for _, item := range resource.Items {
			afterSet[item] = true
		}
	}

	for _, item := range resource.Items {
		if !beforeSet[item] {
			addItems = append(addItems, item)
		}
	}

	for _, item := range a.actual.Items {
		if !afterSet[item] {
			removeItems = append(removeItems, item)
		}
	}

	addItemsReq := &api.AddAddressGroupItemsRequest{
		Items: addItems,
	}
	removeItemsReq := &api.RemoveAddressGroupItemsRequest{
		Items: removeItems,
	}
	//todo: support clone items?

	if desired.Spec.OrganizationRef != nil {
		if len(addItems) != 0 {
			op, err = a.gcpClient.Organizations.Locations.AddressGroups.AddItems(a.id.String(), addItemsReq).Context(ctx).Do()
		}
		if len(removeItems) != 0 {
			op, err = a.gcpClient.Organizations.Locations.AddressGroups.RemoveItems(a.id.String(), removeItemsReq).Context(ctx).Do()
		}
	} else {
		if len(addItems) != 0 {
			op, err = a.gcpClient.Projects.Locations.AddressGroups.AddItems(a.id.String(), addItemsReq).Context(ctx).Do()
		}
		if len(removeItems) != 0 {
			op, err = a.gcpClient.Projects.Locations.AddressGroups.RemoveItems(a.id.String(), removeItemsReq).Context(ctx).Do()
		}
	}

	if err != nil {
		return fmt.Errorf("updating AddressGroup items %s: %w", a.id, err)
	}

	if err := WaitForOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("error waiting AddressGroup items %s update: %w", a.id, err)
	}

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting updated AddressGroup %q: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AddressGroup", "name", a.id)

	status := &krm.NetworkSecurityAddressGroupStatus{}
	status.ObservedState = NetworkSecurityAddressGroupObservedState_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AddressGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityAddressGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityAddressGroupSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NetworkSecurityAddressGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AddressGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AddressGroup", "name", a.id)

	var op *api.Operation
	var err error
	if a.desired.Spec.OrganizationRef != nil {
		op, err = a.gcpClient.Organizations.Locations.AddressGroups.Delete(a.id.String()).Context(ctx).Do()
	} else {
		op, err = a.gcpClient.Projects.Locations.AddressGroups.Delete(a.id.String()).Context(ctx).Do()
	}

	if err := WaitForOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("error waiting Group %s deletion: %w", a.id, err)
	}
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AddressGroup %q: %w", a.id, err)
	}

	//if err := WaitForOp(ctx, op); err != nil {
	//	return false, fmt.Errorf("error waiting Group %s deletion: %w", a.id, err)
	//}

	log.V(2).Info("successfully deleted AddressGroup", "name", a.id)
	return true, nil
}

func (a *AddressGroupAdapter) get(ctx context.Context) (resource *api.AddressGroup, err error) {
	if a.desired.Spec.OrganizationRef != nil {
		resource, err = a.gcpClient.Organizations.Locations.AddressGroups.Get(a.id.String()).Context(ctx).Do()
	} else {
		resource, err = a.gcpClient.Projects.Locations.AddressGroups.Get(a.id.String()).Context(ctx).Do()
	}
	return resource, err
}
