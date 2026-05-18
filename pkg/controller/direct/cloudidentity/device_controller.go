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

package cloudidentity

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudIdentityDeviceGVK, NewDeviceModel)
}

func NewDeviceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDevice{config: config}, nil
}

var _ directbase.Model = &modelDevice{}

type modelDevice struct {
	config *config.ControllerConfig
}

func (m *modelDevice) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudidentity client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDevice) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudIdentityDevice{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeviceAdapter{
		id:        id.(*krm.CloudIdentityDeviceIdentity),
		inner:     obj,
		gcpClient: gcpClient,
	}, nil
}

func (m *modelDevice) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil // not implemented yet
}

type DeviceAdapter struct {
	id        *krm.CloudIdentityDeviceIdentity
	inner     *krm.CloudIdentityDevice
	gcpClient *api.Service
	actual    *api.Device
}

var _ directbase.Adapter = &DeviceAdapter{}

func (a *DeviceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudIdentityDevice", "name", a.id.String())

	req := a.gcpClient.Devices.Get(a.id.String())
	device, err := req.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CloudIdentityDevice %q: %w", a.id.String(), err)
	}

	a.actual = device
	return true, nil
}

func (a *DeviceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudIdentityDevice", "name", a.id.String())

	desired := a.inner.DeepCopy()
	resource := CloudIdentityDeviceSpec_ToAPI(&desired.Spec)

	req := a.gcpClient.Devices.Create(&api.CreateDeviceRequest{
		Device:   resource,
		Customer: "customers/my_customer", // As per API docs, this can be used for your own org.
	})

	op, err := req.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating CloudIdentityDevice %q: %w", a.id.String(), err)
	}

	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for create operation for CloudIdentityDevice %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created CloudIdentityDevice", "name", a.id.String())

	var data map[string]interface{}
	if err := json.Unmarshal(op.Response, &data); err != nil {
		return fmt.Errorf("decoding operation response: %w", err)
	}
	generatedName, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("name not found in operation response")
	}

	created, err := a.gcpClient.Devices.Get(generatedName).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created CloudIdentityDevice %q: %w", generatedName, err)
	}

	status := &krm.CloudIdentityDeviceStatus{}
	status.ExternalRef = &generatedName
	status.ObservedState = CloudIdentityDeviceObservedState_FromAPI(created)

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *DeviceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudIdentityDevice", "name", a.id.String())

	// The API doesn't support updating a Device.
	// If the desired state differs from actual, we should probably return an error or ignore.
	desired := a.inner.DeepCopy()
	desiredAPI := CloudIdentityDeviceSpec_ToAPI(&desired.Spec)

	// Since we can't update, let's just check if it matches, or simply do nothing.
	// Let's log it and skip.
	if !reflect.DeepEqual(a.actual.DeviceType, desiredAPI.DeviceType) {
		log.Info("cannot update DeviceType, device is immutable")
	}

	status := &krm.CloudIdentityDeviceStatus{}
	status.ObservedState = CloudIdentityDeviceObservedState_FromAPI(a.actual)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *DeviceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	spec := CloudIdentityDeviceSpec_FromAPI(a.actual)
	obj := &krm.CloudIdentityDevice{}
	if spec != nil {
		obj.Spec = *spec
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *DeviceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudIdentityDevice", "name", a.id.String())

	req := a.gcpClient.Devices.Delete(a.id.String())
	req.Customer("customers/my_customer") // We need to specify customer

	op, err := req.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting CloudIdentityDevice %q: %w", a.id.String(), err)
	}

	if err := WaitForCloudIdentityOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for delete operation for CloudIdentityDevice %q: %w", a.id.String(), err)
	}

	return true, nil
}
