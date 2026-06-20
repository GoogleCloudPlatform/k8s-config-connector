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
// proto.service: google.cloud.compute.v1.Addresses
// proto.message: google.cloud.compute.v1.Address
// crd.type: ComputeAddress
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeAddressGVK, NewComputeAddressModel)
}

func NewComputeAddressModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeAddressModel{config: config}, nil
}

var _ directbase.Model = &computeAddressModel{}

type computeAddressModel struct {
	config *config.ControllerConfig
}

func (m *computeAddressModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeAddress{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resolvedID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedID.(*krm.ComputeAddressIdentity)

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	var addressesClient *gcp.AddressesClient
	var globalAddressesClient *gcp.GlobalAddressesClient

	if id.IsGlobal() {
		globalAddressesClient, err = gcpClient.newGlobalAddressesClient(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		addressesClient, err = gcpClient.newAddressesClient(ctx)
		if err != nil {
			return nil, err
		}
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &computeAddressAdapter{
		addressesClient:       addressesClient,
		globalAddressesClient: globalAddressesClient,
		id:                    id,
		desired:               desired,
		reader:                reader,
	}, nil
}

func (m *computeAddressModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type computeAddressAdapter struct {
	addressesClient       *gcp.AddressesClient
	globalAddressesClient *gcp.GlobalAddressesClient
	id                    *krm.ComputeAddressIdentity
	desired               *computepb.Address
	actual                *computepb.Address
	reader                client.Reader
}

var _ directbase.Adapter = &computeAddressAdapter{}

func (a *computeAddressAdapter) get(ctx context.Context) (*computepb.Address, error) {
	if a.id.IsGlobal() {
		req := &computepb.GetGlobalAddressRequest{
			Project: a.id.Project,
			Address: a.id.Address,
		}
		return a.globalAddressesClient.Get(ctx, req)
	} else {
		req := &computepb.GetAddressRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Address: a.id.Address,
		}
		return a.addressesClient.Get(ctx, req)
	}
}

func (a *computeAddressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeAddress", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeAddress %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

func (a *computeAddressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeAddress", "name", a.id)

	a.desired.Name = direct.LazyPtr(a.id.Address)

	var op *gcp.Operation
	var err error

	if a.id.IsGlobal() {
		req := &computepb.InsertGlobalAddressRequest{
			Project:         a.id.Project,
			AddressResource: a.desired,
		}
		op, err = a.globalAddressesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertAddressRequest{
			Project:         a.id.Project,
			Region:          a.id.Region,
			AddressResource: a.desired,
		}
		op, err = a.addressesClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeAddress %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeAddress %s create failed: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeAddress", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting created ComputeAddress %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *computeAddressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeAddress", "name", a.id)

	diffs, err := compareComputeAddress(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		return fmt.Errorf("ComputeAddress is immutable and cannot be updated. Detected changes: %v", diffs.Fields)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func compareComputeAddress(ctx context.Context, actual, desired *computepb.Address) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeAddressSpec_v1beta1_FromProto, ComputeAddressSpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.Clone(desired).(*computepb.Address)

	populateDefaults := func(obj *computepb.Address) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *computeAddressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ComputeAddress{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeAddressSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = a.id.Region
	obj.Spec.ResourceID = direct.PtrTo(a.id.Address)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Address)
	u.SetGroupVersionKind(krm.ComputeAddressGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *computeAddressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeAddress", "name", a.id)

	var op *gcp.Operation
	var err error

	if a.id.IsGlobal() {
		req := &computepb.DeleteGlobalAddressRequest{
			Project: a.id.Project,
			Address: a.id.Address,
		}
		op, err = a.globalAddressesClient.Delete(ctx, req)
	} else {
		req := &computepb.DeleteAddressRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Address: a.id.Address,
		}
		op, err = a.addressesClient.Delete(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeAddress %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting ComputeAddress %s delete failed: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ComputeAddress", "name", a.id)
	return true, nil
}

func (a *computeAddressAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Address) error {
	mapCtx := &direct.MapContext{}
	status := ComputeAddressStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
