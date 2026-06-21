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
// proto.service: google.cloud.compute.v1.GlobalAddresses
// proto.message: google.cloud.compute.v1.Address
// crd.type: ComputeAddress
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
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

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	var addressesClient *compute.AddressesClient
	var globalAddressesClient *compute.GlobalAddressesClient
	addressId := id.(*krm.ComputeAddressIdentity)

	if addressId.IsGlobal() {
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

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeAddressSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	resource.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &ComputeAddressAdapter{
		gcpClient:       addressesClient,
		gcpGlobalClient: globalAddressesClient,
		id:              addressId,
		desired:         resource,
		reader:          reader,
	}, nil
}

func (m *computeAddressModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeAddressIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	var addressesClient *compute.AddressesClient
	var globalAddressesClient *compute.GlobalAddressesClient

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

	return &ComputeAddressAdapter{
		id:              id,
		gcpClient:       addressesClient,
		gcpGlobalClient: globalAddressesClient,
	}, nil
}

type ComputeAddressAdapter struct {
	gcpClient       *compute.AddressesClient
	gcpGlobalClient *compute.GlobalAddressesClient
	id              *krm.ComputeAddressIdentity
	desired         *pb.Address
	actual          *pb.Address
	reader          client.Reader
}

var _ directbase.Adapter = &ComputeAddressAdapter{}

func (a *ComputeAddressAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeAddress", "name", a.id)

	var actual *pb.Address
	var err error

	if a.id.IsGlobal() {
		req := &pb.GetGlobalAddressRequest{
			Project: a.id.Project,
			Address: a.id.Address,
		}
		actual, err = a.gcpGlobalClient.Get(ctx, req)
	} else {
		req := &pb.GetAddressRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Address: a.id.Address,
		}
		actual, err = a.gcpClient.Get(ctx, req)
	}

	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeAddress %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeAddressAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeAddress", "name", a.id)

	a.desired.Name = direct.LazyPtr(a.id.Address)

	var err error
	if a.id.IsGlobal() {
		req := &pb.InsertGlobalAddressRequest{
			Project:         a.id.Project,
			AddressResource: a.desired,
		}
		op, errInsert := a.gcpGlobalClient.Insert(ctx, req)
		if errInsert != nil {
			return fmt.Errorf("creating ComputeGlobalAddress %s: %w", a.id, errInsert)
		}
		err = op.Wait(ctx)
	} else {
		req := &pb.InsertAddressRequest{
			Project:         a.id.Project,
			Region:          a.id.Region,
			AddressResource: a.desired,
		}
		op, errInsert := a.gcpClient.Insert(ctx, req)
		if errInsert != nil {
			return fmt.Errorf("creating ComputeRegionalAddress %s: %w", a.id, errInsert)
		}
		err = op.Wait(ctx)
	}

	if err != nil {
		return fmt.Errorf("waiting ComputeAddress %s create failed: %w", a.id, err)
	}
	log.V(2).Info("successfully created ComputeAddress", "name", a.id)

	// Get latest state
	latest, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeAddress %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *ComputeAddressAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeAddress", "name", a.id)

	diffs, _, err := compareAddress(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	// Check if only labels changed by comparing actual vs desired without labels
	actualNoLabels := proto.Clone(a.actual).(*pb.Address)
	desiredNoLabels := proto.Clone(a.desired).(*pb.Address)
	actualNoLabels.Labels = nil
	desiredNoLabels.Labels = nil

	specDiffs, _, err := compareAddress(ctx, actualNoLabels, desiredNoLabels)
	if err != nil {
		return err
	}

	if !specDiffs.HasDiff() {
		// Only labels changed, so we can update labels!
		log.V(2).Info("updating ComputeAddress labels", "name", a.id)
		if a.id.IsGlobal() {
			req := &pb.SetLabelsGlobalAddressRequest{
				Project:  a.id.Project,
				Resource: a.id.Address,
				GlobalSetLabelsRequestResource: &pb.GlobalSetLabelsRequest{
					Labels:           a.desired.Labels,
					LabelFingerprint: a.actual.LabelFingerprint,
				},
			}
			op, err := a.gcpGlobalClient.SetLabels(ctx, req)
			if err != nil {
				return fmt.Errorf("updating ComputeGlobalAddress labels %s: %w", a.id, err)
			}
			if err = op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting ComputeGlobalAddress %s labels update: %w", a.id, err)
			}
		} else {
			req := &pb.SetLabelsAddressRequest{
				Project:  a.id.Project,
				Region:   a.id.Region,
				Resource: a.id.Address,
				RegionSetLabelsRequestResource: &pb.RegionSetLabelsRequest{
					Labels:           a.desired.Labels,
					LabelFingerprint: a.actual.LabelFingerprint,
				},
			}
			op, err := a.gcpClient.SetLabels(ctx, req)
			if err != nil {
				return fmt.Errorf("updating ComputeRegionalAddress labels %s: %w", a.id, err)
			}
			if err = op.Wait(ctx); err != nil {
				return fmt.Errorf("waiting ComputeRegionalAddress %s labels update: %w", a.id, err)
			}
		}

		// Retrieve latest state after update
		latest, err := a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeAddress %s after labels update: %w", a.id, err)
		}
		return a.updateStatus(ctx, updateOp, latest)
	}

	// Surfacing exact diff back to the user
	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	// Since ComputeAddress spec is immutable, any change is not supported.
	return fmt.Errorf("ComputeAddress is immutable and cannot be updated")
}

func (a *ComputeAddressAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeAddress", "name", a.id)

	var err error
	if a.id.IsGlobal() {
		req := &pb.DeleteGlobalAddressRequest{
			Project: a.id.Project,
			Address: a.id.Address,
		}
		op, errDelete := a.gcpGlobalClient.Delete(ctx, req)
		if errDelete != nil {
			if direct.IsNotFound(errDelete) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeGlobalAddress %s: %w", a.id, errDelete)
		}
		err = op.Wait(ctx)
	} else {
		req := &pb.DeleteAddressRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Address: a.id.Address,
		}
		op, errDelete := a.gcpClient.Delete(ctx, req)
		if errDelete != nil {
			if direct.IsNotFound(errDelete) {
				return true, nil
			}
			return false, fmt.Errorf("deleting ComputeRegionalAddress %s: %w", a.id, errDelete)
		}
		err = op.Wait(ctx)
	}

	if err != nil {
		return false, fmt.Errorf("waiting ComputeAddress %s delete failed: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeAddress", "name", a.id)
	return true, nil
}

func (a *ComputeAddressAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeAddress{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeAddressSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// ComputeAddress requires location and resourceID
	obj.Spec.Location = a.id.Region
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Address)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Address)
	u.SetGroupVersionKind(krm.ComputeAddressGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *ComputeAddressAdapter) get(ctx context.Context) (*pb.Address, error) {
	if a.id.IsGlobal() {
		req := &pb.GetGlobalAddressRequest{
			Project: a.id.Project,
			Address: a.id.Address,
		}
		return a.gcpGlobalClient.Get(ctx, req)
	}

	req := &pb.GetAddressRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Address: a.id.Address,
	}
	return a.gcpClient.Get(ctx, req)
}

func (a *ComputeAddressAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Address) error {
	mapCtx := &direct.MapContext{}
	status := ComputeAddressStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if status != nil {
		status.ObservedState = ComputeAddressObservedState_v1beta1_FromProto(mapCtx, latest)
	}

	return op.UpdateStatus(ctx, status, nil)
}

func compareAddress(ctx context.Context, actual, desired *pb.Address) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeAddressSpec_v1beta1_FromProto, ComputeAddressSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.Clone(desired).(*pb.Address)

	populateDefaults := func(obj *pb.Address) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
		if obj.AddressType == nil {
			obj.AddressType = direct.PtrTo("EXTERNAL")
		}
		if obj.IpVersion == nil {
			obj.IpVersion = direct.PtrTo("IPV4")
		}
		if obj.NetworkTier == nil {
			obj.NetworkTier = direct.PtrTo("PREMIUM")
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
