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

// +tool:controller
// proto.service: google.cloud.edgecontainer.v1.EdgeContainer
// proto.message: google.cloud.edgecontainer.v1.Machine
// crd.type: EdgeContainerMachine
// crd.version: v1alpha1

package edgecontainer

import (
	"context"
	"fmt"
	"reflect"

	api "cloud.google.com/go/edgecontainer/apiv1"
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.EdgeContainerMachineGVK, NewMachineModel)
}

func NewMachineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &machineModel{config: *config}, nil
}

var _ directbase.Model = &machineModel{}

type machineModel struct {
	config config.ControllerConfig
}

func (m *machineModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EdgeContainerMachine{}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.EdgeContainerMachineIdentity)

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}
	return &machineAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *machineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type machineAdapter struct {
	gcpClient *api.Client
	id        *krm.EdgeContainerMachineIdentity
	desired   *krm.EdgeContainerMachine
	actual    *pb.Machine
}

var _ directbase.Adapter = &machineAdapter{}

func (a *machineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting edgecontainer machine", "name", a.id)

	req := &pb.GetMachineRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMachine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting edgecontainer machine %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *machineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating edgecontainer machine is not supported", "name", a.id)
	return fmt.Errorf("EdgeContainerMachine %q not found on GCP; physical machines cannot be created via the GCP API and must be registered out of band", a.id.String())
}

func (a *machineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating edgecontainer machine", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := EdgeContainerMachineSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		return fmt.Errorf("EdgeContainerMachine %q is read-only and labels cannot be updated on GCP", a.id.String())
	}
	if resource.Zone != a.actual.Zone {
		return fmt.Errorf("EdgeContainerMachine %q is read-only and zone cannot be updated on GCP", a.id.String())
	}

	status := &krm.EdgeContainerMachineStatus{}
	status.ObservedState = EdgeContainerMachineObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *machineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EdgeContainerMachine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EdgeContainerMachineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Machine)
	u.SetGroupVersionKind(krm.EdgeContainerMachineGVK)
	u.Object = uObj
	return u, nil
}

func (a *machineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting edgecontainer machine is a no-op", "name", a.id)
	return true, nil
}
