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

package compute

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeAddressGVK, NewComputeAddressModel)
}

func NewComputeAddressModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeAddressModel{config: config}, nil
}

type computeAddressModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computeAddressModel{}

type computeAddressAdapter struct {
	id                    *krm.ComputeAddressIdentity
	addressesClient       *gcp.AddressesClient
	globalAddressesClient *gcp.GlobalAddressesClient
	desired               *krm.ComputeAddress
	actual                *computepb.Address
	reader                client.Reader
}

var _ directbase.Adapter = &computeAddressAdapter{}

func (m *computeAddressModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeAddress{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewComputeAddressIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	adapter := &computeAddressAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	if id.Location() == "global" {
		client, err := gcpClient.newGlobalAddressesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.globalAddressesClient = client
	} else {
		client, err := gcpClient.newAddressesClient(ctx)
		if err != nil {
			return nil, err
		}
		adapter.addressesClient = client
	}

	return adapter, nil
}

func (m *computeAddressModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *computeAddressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeAddress", "name", a.id)

	addr, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeAddress %q: %w", a.id, err)
	}
	a.actual = addr
	return true, nil
}

func (a *computeAddressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	var err error

	err = resolveComputeAddressRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeAddress", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	pbAddress := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	pbAddress.Name = direct.PtrTo(a.id.ID())

	if a.id.Location() == "global" {
		req := &computepb.InsertGlobalAddressRequest{
			Project:         a.id.Project(),
			AddressResource: pbAddress,
		}
		op, err := a.globalAddressesClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("inserting global ComputeAddress %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for global ComputeAddress %s creation: %w", a.id, err)
			}
		}
	} else {
		req := &computepb.InsertAddressRequest{
			Project:         a.id.Project(),
			Region:          a.id.Location(),
			AddressResource: pbAddress,
		}
		op, err := a.addressesClient.Insert(ctx, req)
		if err != nil {
			return fmt.Errorf("inserting regional ComputeAddress %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for regional ComputeAddress %s creation: %w", a.id, err)
			}
		}
	}

	// Get the created resource
	actual, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeAddress %s: %w", a.id, err)
	}

	status := &krm.ComputeAddressStatus{}
	ComputeAddressStatus_FromProto(mapCtx, actual, status)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ObservedGeneration = direct.PtrTo(a.desired.Generation)

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *computeAddressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	var err error

	err = resolveComputeAddressRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeAddress", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	pbAddress := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	pbAddress.Name = direct.PtrTo(a.id.ID())

	actualCopy := proto.Clone(a.actual).(*computepb.Address)
	// Normalization for comparison
	pbAddress.CreationTimestamp = actualCopy.CreationTimestamp
	pbAddress.Id = actualCopy.Id
	pbAddress.SelfLink = actualCopy.SelfLink
	pbAddress.Kind = actualCopy.Kind
	pbAddress.Status = actualCopy.Status
	pbAddress.Region = actualCopy.Region
	pbAddress.Users = actualCopy.Users
	pbAddress.LabelFingerprint = actualCopy.LabelFingerprint

	paths, err := common.CompareProtoMessage(pbAddress, actualCopy, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("ComputeAddress fields are immutable and cannot be updated")
	}

	// If no changes, just update status
	status := &krm.ComputeAddressStatus{}
	ComputeAddressStatus_FromProto(mapCtx, a.actual, status)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ObservedGeneration = direct.PtrTo(a.desired.Generation)

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *computeAddressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeAddress", "name", a.id)

	if a.id.Location() == "global" {
		req := &computepb.DeleteGlobalAddressRequest{
			Project: a.id.Project(),
			Address: a.id.ID(),
		}
		op, err := a.globalAddressesClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("deleting global ComputeAddress %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return false, fmt.Errorf("waiting for global ComputeAddress %s deletion: %w", a.id, err)
			}
		}
	} else {
		req := &computepb.DeleteAddressRequest{
			Project: a.id.Project(),
			Region:  a.id.Location(),
			Address: a.id.ID(),
		}
		op, err := a.addressesClient.Delete(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("deleting regional ComputeAddress %s: %w", a.id, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return false, fmt.Errorf("waiting for regional ComputeAddress %s deletion: %w", a.id, err)
			}
		}
	}

	return true, nil
}

func (a *computeAddressAdapter) get(ctx context.Context) (*computepb.Address, error) {
	if a.id.Location() == "global" {
		req := &computepb.GetGlobalAddressRequest{
			Project: a.id.Project(),
			Address: a.id.ID(),
		}
		return a.globalAddressesClient.Get(ctx, req)
	} else {
		req := &computepb.GetAddressRequest{
			Project: a.id.Project(),
			Region:  a.id.Location(),
			Address: a.id.ID(),
		}
		return a.addressesClient.Get(ctx, req)
	}
}

func (a *computeAddressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
