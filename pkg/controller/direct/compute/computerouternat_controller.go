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
// proto.message: google.cloud.compute.v1.RouterNat
// crd.type: ComputeRouterNAT
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeRouterNATGVK, NewRouterNATModel)
}

func NewRouterNATModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &routerNATModel{config: config}, nil
}

var _ directbase.Model = &routerNATModel{}

type routerNATModel struct {
	config *config.ControllerConfig
}

func (m *routerNATModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeRouterNAT{}
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
	resource := ComputeRouterNATSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &RouterNATAdapter{
		gcpClient: routersClient,
		id:        id.(*krm.ComputeRouterNATIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *routerNATModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeRouterNATIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	routersClient, err := gcpClient.newRoutersClient(ctx)
	if err != nil {
		return nil, err
	}

	return &RouterNATAdapter{
		gcpClient: routersClient,
		id:        id,
	}, nil
}

type RouterNATAdapter struct {
	gcpClient *compute.RoutersClient
	id        *krm.ComputeRouterNATIdentity
	desired   *computepb.RouterNat
	actual    *computepb.RouterNat
	router    *computepb.Router
	reader    client.Reader
}

var _ directbase.Adapter = &RouterNATAdapter{}

func (a *RouterNATAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRouter for ComputeRouterNAT", "name", a.id)

	req := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	router, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeRouter %q: %w", a.id.Router, err)
	}

	a.router = router

	for _, nat := range router.Nats {
		if nat.GetName() == a.id.ComputeRouterNAT {
			a.actual = nat
			return true, nil
		}
	}

	return false, nil
}

func (a *RouterNATAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRouterNAT", "name", a.id)

	if a.router == nil {
		req := &computepb.GetRouterRequest{
			Project: a.id.Project,
			Region:  a.id.Region,
			Router:  a.id.Router,
		}
		router, err := a.gcpClient.Get(ctx, req)
		if err != nil {
			return fmt.Errorf("getting ComputeRouter %q: %w", a.id.Router, err)
		}
		a.router = router
	}

	a.desired.Name = proto.String(a.id.ComputeRouterNAT)

	// Check if already exists in a.router.Nats
	exists := false
	for _, nat := range a.router.Nats {
		if nat.GetName() == a.id.ComputeRouterNAT {
			exists = true
			break
		}
	}

	if !exists {
		a.router.Nats = append(a.router.Nats, a.desired)
	}

	patchRouter := &computepb.Router{
		Nats: a.router.Nats,
	}

	req := &computepb.PatchRouterRequest{
		Project:        a.id.Project,
		Region:         a.id.Region,
		Router:         a.id.Router,
		RouterResource: patchRouter,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeRouterNAT %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeRouterNAT %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeRouterNAT in gcp", "name", a.id)

	// Fetch latest router to get the created NAT
	latestRouter, err := a.getRouter(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeRouter %s: %w", a.id.Router, err)
	}
	a.router = latestRouter

	var created *computepb.RouterNat
	for _, nat := range latestRouter.Nats {
		if nat.GetName() == a.id.ComputeRouterNAT {
			created = nat
			break
		}
	}
	if created == nil {
		return fmt.Errorf("created NAT %s not found in latest router configuration", a.id.ComputeRouterNAT)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *RouterNATAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRouterNAT", "name", a.id)

	diffs, _, err := compareComputeRouterNAT(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.RouterNat
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)

		if a.router == nil {
			req := &computepb.GetRouterRequest{
				Project: a.id.Project,
				Region:  a.id.Region,
				Router:  a.id.Router,
			}
			router, err := a.gcpClient.Get(ctx, req)
			if err != nil {
				return fmt.Errorf("getting ComputeRouter %q: %w", a.id.Router, err)
			}
			a.router = router
		}

		a.desired.Name = proto.String(a.id.ComputeRouterNAT)

		// Update our NAT in a.router.Nats
		found := false
		for i, nat := range a.router.Nats {
			if nat.GetName() == a.id.ComputeRouterNAT {
				a.router.Nats[i] = a.desired
				found = true
				break
			}
		}
		if !found {
			a.router.Nats = append(a.router.Nats, a.desired)
		}

		patchRouter := &computepb.Router{
			Nats: a.router.Nats,
		}

		req := &computepb.PatchRouterRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			Router:         a.id.Router,
			RouterResource: patchRouter,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeRouterNAT %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeRouterNAT", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeRouterNAT %s waiting for update: %w", a.id.String(), err)
		}

		latestRouter, err := a.getRouter(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeRouter %s: %w", a.id.Router, err)
		}
		a.router = latestRouter

		for _, nat := range latestRouter.Nats {
			if nat.GetName() == a.id.ComputeRouterNAT {
				updated = nat
				break
			}
		}
		if updated == nil {
			return fmt.Errorf("updated NAT %s not found in latest router configuration", a.id.ComputeRouterNAT)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *RouterNATAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeRouterNAT{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeRouterNATSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Region = a.id.Region
	obj.Spec.RouterRef = krm.ComputeRouterRef{
		External: fmt.Sprintf("projects/%s/regions/%s/routers/%s", a.id.Project, a.id.Region, a.id.Router),
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ComputeRouterNAT)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ComputeRouterNAT)
	u.SetGroupVersionKind(krm.ComputeRouterNATGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *RouterNATAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRouterNAT", "name", a.id)

	req := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	router, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.Info("parent ComputeRouter not found, assuming NAT is deleted", "router", a.id.Router)
			return true, nil
		}
		return false, fmt.Errorf("getting ComputeRouter %q: %w", a.id.Router, err)
	}

	a.router = router

	var newNats []*computepb.RouterNat
	found := false
	for _, nat := range a.router.Nats {
		if nat.GetName() == a.id.ComputeRouterNAT {
			found = true
			continue
		} else {
			newNats = append(newNats, nat)
		}
	}

	if !found {
		log.Info("ComputeRouterNAT already deleted or not found on router", "name", a.id)
		return true, nil
	}

	patchRouter := &computepb.Router{
		Nats: newNats,
	}

	updateReq := &computepb.PatchRouterRequest{
		Project:        a.id.Project,
		Region:         a.id.Region,
		Router:         a.id.Router,
		RouterResource: patchRouter,
	}
	op, err := a.gcpClient.Patch(ctx, updateReq)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeRouterNAT %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of compute ComputeRouterNAT %s: %w", a.id.String(), err)
	}

	log.Info("successfully deleted compute ComputeRouterNAT", "name", a.id)
	return true, nil
}

func (a *RouterNATAdapter) getRouter(ctx context.Context) (*computepb.Router, error) {
	getReq := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeRouter %s: %w", a.id.Router, err)
	}
	return resource, nil
}

func (a *RouterNATAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.RouterNat) error {
	status := &krm.ComputeRouterNATStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeRouterNAT(ctx context.Context, actual, desired *computepb.RouterNat) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeRouterNATSpec_FromProto, ComputeRouterNATSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *computepb.RouterNat) {
		if obj.IcmpIdleTimeoutSec == nil {
			obj.IcmpIdleTimeoutSec = proto.Int32(30)
		}
		if obj.TcpEstablishedIdleTimeoutSec == nil {
			obj.TcpEstablishedIdleTimeoutSec = proto.Int32(1200)
		}
		if obj.TcpTimeWaitTimeoutSec == nil {
			obj.TcpTimeWaitTimeoutSec = proto.Int32(120)
		}
		if obj.TcpTransitoryIdleTimeoutSec == nil {
			obj.TcpTransitoryIdleTimeoutSec = proto.Int32(30)
		}
		if obj.UdpIdleTimeoutSec == nil {
			obj.UdpIdleTimeoutSec = proto.Int32(30)
		}
		if obj.EnableEndpointIndependentMapping == nil {
			obj.EnableEndpointIndependentMapping = proto.Bool(true)
		}
		if obj.EnableDynamicPortAllocation == nil {
			obj.EnableDynamicPortAllocation = proto.Bool(false)
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
