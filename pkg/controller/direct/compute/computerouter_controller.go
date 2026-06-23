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
// proto.service: google.cloud.compute.v1.Routers
// proto.message: google.cloud.compute.v1.Router
// crd.type: ComputeRouter
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeRouterGVK, NewRouterModel)
}

func NewRouterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &routerModel{config: config}, nil
}

var _ directbase.Model = &routerModel{}

type routerModel struct {
	config *config.ControllerConfig
}

func (m *routerModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeRouter{}
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
	routersClient, err := gcpClient.newRoutersClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeRouterSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &RouterAdapter{
		gcpClient: routersClient,
		id:        id.(*krm.ComputeRouterIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *routerModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type RouterAdapter struct {
	gcpClient *compute.RoutersClient
	id        *krm.ComputeRouterIdentity
	desired   *computepb.Router
	actual    *computepb.Router
	reader    client.Reader
}

var _ directbase.Adapter = &RouterAdapter{}

func (a *RouterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRouter", "name", a.id)

	req := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeRouter %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *RouterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRouter", "name", a.id)

	a.desired.Name = proto.String(a.id.Router)

	req := &computepb.InsertRouterRequest{
		Project:        a.id.Project,
		Region:         a.id.Region,
		RouterResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeRouter %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeRouter %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeRouter in gcp", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeRouter %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *RouterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRouter", "name", a.id)

	diffs, _, err := compareComputeRouter(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.Router
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &computepb.PatchRouterRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			Router:         a.id.Router,
			RouterResource: a.desired,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeRouter %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeRouter", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeRouter %s waiting for update: %w", a.id.String(), err)
		}

		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeRouter %s: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *RouterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeRouter{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeRouterSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeRouterGVK)

	u.Object = uObj
	return u, nil
}

func (a *RouterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRouter", "name", a.id)

	req := &computepb.DeleteRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeRouter %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeRouter", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute ComputeRouter %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *RouterAdapter) get(ctx context.Context) (*computepb.Router, error) {
	getReq := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeRouter %s: %w", a.id, err)
	}
	return resource, nil
}

func (a *RouterAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Router) error {
	mapCtx := &direct.MapContext{}
	status := ComputeRouterStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeRouter(ctx context.Context, actual, desired *computepb.Router) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeRouterSpec_v1beta1_FromProto, ComputeRouterSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*computepb.Router)

	populateDefaults := func(obj *computepb.Router) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
