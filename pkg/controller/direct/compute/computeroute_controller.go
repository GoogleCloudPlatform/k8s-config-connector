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
// proto.service: google.cloud.compute.v1.Routes
// proto.message: google.cloud.compute.v1.Route
// crd.type: ComputeRoute
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
	registry.RegisterModel(krm.ComputeRouteGVK, NewRouteModel)
}

func NewRouteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &routeModel{config: config}, nil
}

var _ directbase.Model = &routeModel{}

type routeModel struct {
	config *config.ControllerConfig
}

func (m *routeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeRoute{}
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
	routesClient, err := gcpClient.newRoutesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeRouteSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &RouteAdapter{
		gcpClient: routesClient,
		id:        id.(*krm.ComputeRouteIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *routeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeRouteIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	routesClient, err := gcpClient.newRoutesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &RouteAdapter{
		gcpClient: routesClient,
		id:        id,
	}, nil
}

type RouteAdapter struct {
	gcpClient *compute.RoutesClient
	id        *krm.ComputeRouteIdentity
	desired   *computepb.Route
	actual    *computepb.Route
	reader    client.Reader
}

var _ directbase.Adapter = &RouteAdapter{}

func (a *RouteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRoute", "name", a.id)

	req := &computepb.GetRouteRequest{
		Project: a.id.Project,
		Route:   a.id.Route,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeRoute %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *RouteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRoute", "name", a.id)

	a.desired.Name = proto.String(a.id.Route)

	req := &computepb.InsertRouteRequest{
		Project:       a.id.Project,
		RouteResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeRoute %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeRoute %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeRoute in gcp", "name", a.id)

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeRoute %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *RouteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRoute", "name", a.id)

	diffs, _, err := compareComputeRoute(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("ComputeRoute is immutable and cannot be updated")
}

func (a *RouteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRoute", "name", a.id)

	req := &computepb.DeleteRouteRequest{
		Project: a.id.Project,
		Route:   a.id.Route,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeRoute %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting ComputeRoute %s delete failed: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ComputeRoute", "name", a.id)
	return true, nil
}

func (a *RouteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeRoute{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeRouteSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Route)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeRouteGVK)
	return u, nil
}

func (a *RouteAdapter) get(ctx context.Context) (*computepb.Route, error) {
	req := &computepb.GetRouteRequest{
		Project: a.id.Project,
		Route:   a.id.Route,
	}
	return a.gcpClient.Get(ctx, req)
}

func (a *RouteAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Route) error {
	mapCtx := &direct.MapContext{}
	status := ComputeRouteStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeRoute(ctx context.Context, actual, desired *computepb.Route) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeRouteSpec_v1beta1_FromProto, ComputeRouteSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.Route) {
		if obj.Priority == nil {
			obj.Priority = direct.PtrTo[uint32](1000)
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
