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
// proto.service: google.cloud.compute.v1.Subnetworks
// proto.message: google.cloud.compute.v1.Subnetwork
// crd.type: ComputeSubnetwork
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeSubnetworkGVK, NewSubnetworkModel)
}

func NewSubnetworkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &subnetworkModel{config: config}, nil
}

var _ directbase.Model = &subnetworkModel{}

type subnetworkModel struct {
	config *config.ControllerConfig
}

func (m *subnetworkModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeSubnetwork{}
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
	subnetworksClient, err := gcpClient.newSubnetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desiredSpec := obj.DeepCopy()
	resource := ComputeSubnetworkSpec_v1beta1_ToProto(mapCtx, &desiredSpec.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &SubnetworkAdapter{
		gcpClient: subnetworksClient,
		id:        id.(*krm.ComputeSubnetworkIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *subnetworkModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeSubnetworkIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	subnetworksClient, err := gcpClient.newSubnetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	return &SubnetworkAdapter{
		gcpClient: subnetworksClient,
		id:        id,
	}, nil
}

type SubnetworkAdapter struct {
	gcpClient *compute.SubnetworksClient
	id        *krm.ComputeSubnetworkIdentity
	desired   *computepb.Subnetwork
	actual    *computepb.Subnetwork
	reader    client.Reader
}

var _ directbase.Adapter = &SubnetworkAdapter{}

func (a *SubnetworkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeSubnetwork", "name", a.id)

	req := &computepb.GetSubnetworkRequest{
		Project:    a.id.Project,
		Region:     a.id.Region,
		Subnetwork: a.id.Subnetwork,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeSubnetwork %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *SubnetworkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeSubnetwork", "name", a.id)

	a.desired.Name = proto.String(a.id.Subnetwork)
	if a.desired.LogConfig != nil {
		a.desired.LogConfig.Enable = proto.Bool(true)
	}

	req := &computepb.InsertSubnetworkRequest{
		Project:            a.id.Project,
		Region:             a.id.Region,
		SubnetworkResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeSubnetwork %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeSubnetwork %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeSubnetwork in gcp", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeSubnetwork %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *SubnetworkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeSubnetwork", "name", a.id)

	diffs, _, err := compareComputeSubnetwork(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.Subnetwork
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		if a.desired.LogConfig != nil {
			a.desired.LogConfig.Enable = proto.Bool(true)
		} else if a.actual.LogConfig != nil {
			a.desired.LogConfig = &computepb.SubnetworkLogConfig{
				Enable: proto.Bool(false),
			}
		}

		req := &computepb.PatchSubnetworkRequest{
			Project:            a.id.Project,
			Region:             a.id.Region,
			Subnetwork:         a.id.Subnetwork,
			SubnetworkResource: a.desired,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeSubnetwork %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeSubnetwork", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeSubnetwork %s waiting for update: %w", a.id.String(), err)
		}

		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeSubnetwork %s: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *SubnetworkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeSubnetwork{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeSubnetworkSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Region = &a.id.Region
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Subnetwork)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Subnetwork)
	u.SetGroupVersionKind(krm.ComputeSubnetworkGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *SubnetworkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeSubnetwork", "name", a.id)

	req := &computepb.DeleteSubnetworkRequest{
		Project:    a.id.Project,
		Region:     a.id.Region,
		Subnetwork: a.id.Subnetwork,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting compute ComputeSubnetwork %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeSubnetwork", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeSubnetwork %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *SubnetworkAdapter) get(ctx context.Context) (*computepb.Subnetwork, error) {
	getReq := &computepb.GetSubnetworkRequest{
		Project:    a.id.Project,
		Region:     a.id.Region,
		Subnetwork: a.id.Subnetwork,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeSubnetwork %s: %w", a.id, err)
	}
	return resource, nil
}

func (a *SubnetworkAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Subnetwork) error {
	mapCtx := &direct.MapContext{}
	status := ComputeSubnetworkStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeSubnetwork(ctx context.Context, actual, desired *computepb.Subnetwork) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeSubnetworkSpec_v1beta1_FromProto, ComputeSubnetworkSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.Subnetwork) {
		if obj.PrivateIpGoogleAccess == nil {
			obj.PrivateIpGoogleAccess = proto.Bool(false)
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
