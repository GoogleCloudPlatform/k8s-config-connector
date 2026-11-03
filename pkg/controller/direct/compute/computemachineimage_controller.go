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
// proto.service: google.cloud.compute.v1.MachineImages
// proto.message: google.cloud.compute.v1.MachineImage
// crd.type: ComputeMachineImage
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeMachineImageGVK, NewComputeMachineImageModel)
}

func NewComputeMachineImageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeMachineImageModel{config: config}, nil
}

type computeMachineImageModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &computeMachineImageModel{}

type ComputeMachineImageAdapter struct {
	gcpClient *compute.MachineImagesClient
	id        *krm.ComputeMachineImageIdentity
	desired   *krm.ComputeMachineImage
	actual    *pb.MachineImage
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeMachineImageAdapter{}

func (m *computeMachineImageModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeMachineImage{}
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

	client, err := gcpClient.newMachineImagesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	return &ComputeMachineImageAdapter{
		gcpClient: client,
		id:        id.(*krm.ComputeMachineImageIdentity),
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *computeMachineImageModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeMachineImageIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	client, err := gcpClient.newMachineImagesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeMachineImageAdapter{
		id:        id,
		gcpClient: client,
	}, nil
}

func (a *ComputeMachineImageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeMachineImage", "name", a.id)

	req := &pb.GetMachineImageRequest{
		Project:      a.id.Project,
		MachineImage: a.id.MachineImage,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeMachineImage %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeMachineImageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeMachineImage", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := ComputeMachineImageSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = direct.LazyPtr(a.id.MachineImage)

	req := &pb.InsertMachineImageRequest{
		Project:              a.id.Project,
		MachineImageResource: desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeMachineImage %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeMachineImage %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeMachineImage", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeMachineImage %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeMachineImageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeMachineImage", "name", a.id)

	normalize := func(ctx context.Context, p *pb.MachineImage) error {
		if p == nil {
			return nil
		}
		if p.SourceInstance != nil {
			*p.SourceInstance = refs.TrimComputeURIPrefix(*p.SourceInstance)
		}
		return nil
	}

	diffs, _, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComputeMachineImageSpec_v1alpha1_FromProto,
		ComputeMachineImageSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	// Surface exact diff back to the user
	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("ComputeMachineImage is immutable and cannot be updated")
}

func (a *ComputeMachineImageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeMachineImage{}
	mapCtx := &direct.MapContext{}
	spec := ComputeMachineImageSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec == nil {
		spec = &krm.ComputeMachineImageSpec{}
	}

	spec.ResourceID = direct.LazyPtr(a.id.MachineImage)
	spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	if spec.SourceInstanceRef != nil && spec.SourceInstanceRef.External != "" {
		spec.SourceInstanceRef.External = refs.TrimComputeURIPrefix(spec.SourceInstanceRef.External)
	}

	obj.Spec = *spec

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.MachineImage)
	u.SetGroupVersionKind(krm.ComputeMachineImageGVK)

	export.SetProjectID(u, a.id.Project)
	return u, nil
}

func (a *ComputeMachineImageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeMachineImage", "name", a.id)

	req := &pb.DeleteMachineImageRequest{
		Project:      a.id.Project,
		MachineImage: a.id.MachineImage,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeMachineImage %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully requested deletion of ComputeMachineImage", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("waiting for deletion of ComputeMachineImage %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeMachineImageAdapter) get(ctx context.Context) (*pb.MachineImage, error) {
	getReq := &pb.GetMachineImageRequest{
		Project:      a.id.Project,
		MachineImage: a.id.MachineImage,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeMachineImage %s: %w", a.id, err)
	}
	return resource, nil
}

func (a *ComputeMachineImageAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.MachineImage) error {
	mapCtx := &direct.MapContext{}
	status := ComputeMachineImageStatus_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
