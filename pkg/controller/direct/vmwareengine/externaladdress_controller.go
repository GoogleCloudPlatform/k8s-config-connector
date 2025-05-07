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
// proto.message: google.cloud.vmwareengine.v1.ExternalAddress
// crd.type: VMwareEngineExternalAddress
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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.VMwareEngineExternalAddressGVK, NewExternalAddressModel)
}

func NewExternalAddressModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &externalAddressModel{config: *config}, nil
}

var _ directbase.Model = &externalAddressModel{}

type externalAddressModel struct {
	config config.ControllerConfig
}

func (m *externalAddressModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VMwareEngineExternalAddress{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewExternalAddressIdentity(ctx, reader, obj)
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
	return &externalAddressAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader, // Added reader for reference normalization
	}, nil
}

func (m *externalAddressModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type externalAddressAdapter struct {
	gcpClient *gcp.Client
	id        *krm.ExternalAddressIdentity
	desired   *krm.VMwareEngineExternalAddress
	actual    *pb.ExternalAddress
	reader    client.Reader // Added reader for reference normalization
}

var _ directbase.Adapter = &externalAddressAdapter{}

func (a *externalAddressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine external address", "name", a.id)

	req := &pb.GetExternalAddressRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetExternalAddress(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine external address %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *externalAddressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine external address", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineExternalAddressSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateExternalAddressRequest{
		Parent:            a.id.Parent().String(),
		ExternalAddressId: a.id.ID(),
		ExternalAddress:   resource,
	}
	op, err := a.gcpClient.CreateExternalAddress(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine external address %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine external address %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine external address in gcp", "name", a.id)

	status := &krm.VMwareEngineExternalAddressStatus{}
	status.ObservedState = VMwareEngineExternalAddressObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *externalAddressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine external address", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineExternalAddressSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	// UpdateMask for ExternalAddress allows updating description and internal_ip
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.InternalIP != nil && !reflect.DeepEqual(resource.InternalIp, a.actual.InternalIp) {
		paths = append(paths, "internal_ip")
	}

	var updated *pb.ExternalAddress
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateExternalAddressRequest{
			ExternalAddress: resource,
			UpdateMask:      &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateExternalAddress(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine external address %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine external address %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine external address", "name", a.id)
	}

	status := &krm.VMwareEngineExternalAddressStatus{}
	status.ObservedState = VMwareEngineExternalAddressObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *externalAddressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEngineExternalAddress{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEngineExternalAddressSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Add parent references from ID
	obj.Spec.PrivateCloudRef = &krm.PrivateCloudRef{External: a.id.Parent().String()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// The name in KCC is set from the last segment of the GCP resource name
	u.SetName(a.id.ID()) // Use the ID() method to get the last segment
	u.SetGroupVersionKind(krm.VMwareEngineExternalAddressGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *externalAddressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine external address", "name", a.id)

	req := &pb.DeleteExternalAddressRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteExternalAddress(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ExternalAddress, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine external address %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion of vmwareengine external address", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ExternalAddress %s: %w", a.id, err)
	}
	return true, nil
}
