// Copyright 2024 Google LLC
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
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.VmwareEngineNetwork
// crd.type: VMwareEngineNetwork
// crd.version: v1alpha1

package vmwareengine

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/vmwareengine/apiv1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.VMwareEngineNetworkGVK, NewNetworkModel)
}

func NewNetworkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkModel{config: *config}, nil
}

var _ directbase.Model = &networkModel{}

type networkModel struct {
	config config.ControllerConfig
}

func (m *networkModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VMwareEngineNetwork{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewVmwareEngineNetworkIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get VMwareEngine GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}
	return &networkAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *networkModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type networkAdapter struct {
	gcpClient *gcp.Client
	id        *krm.VmwareEngineNetworkIdentity
	desired   *krm.VMwareEngineNetwork
	actual    *pb.VmwareEngineNetwork
}

var _ directbase.Adapter = &networkAdapter{}

func (a *networkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine network", "name", a.id)

	req := &pb.GetVmwareEngineNetworkRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetVmwareEngineNetwork(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine network %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *networkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine network", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateVmwareEngineNetworkRequest{
		Parent:                a.id.Parent().String(),
		VmwareEngineNetworkId: a.id.ID(),
		VmwareEngineNetwork:   resource,
	}
	op, err := a.gcpClient.CreateVmwareEngineNetwork(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine network %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine network %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine network in gcp", "name", a.id)

	status := &krm.VMwareEngineNetworkStatus{}
	status.ObservedState = VMwareEngineNetworkObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *networkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine network", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Type != nil && !reflect.DeepEqual(resource.Type, a.actual.Type) {
		paths = append(paths, "type")
	}
	// TODO: etag

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
	req := &pb.UpdateVmwareEngineNetworkRequest{
		VmwareEngineNetwork: resource,
		UpdateMask:          &fieldmaskpb.FieldMask{Paths: paths},
	}
	op, err := a.gcpClient.UpdateVmwareEngineNetwork(ctx, req)
	if err != nil {
		return fmt.Errorf("updating vmwareengine network %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine network %s waiting for update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated vmwareengine network", "name", a.id)

	status := &krm.VMwareEngineNetworkStatus{}
	status.ObservedState = VMwareEngineNetworkObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *networkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEngineNetwork{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEngineNetworkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEngineNetworkGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *networkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine network", "name", a.id)

	req := &pb.DeleteVmwareEngineNetworkRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteVmwareEngineNetwork(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine network %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine network", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}
