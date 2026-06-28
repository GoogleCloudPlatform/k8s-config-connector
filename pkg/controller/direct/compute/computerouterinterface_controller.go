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
// proto.message: google.cloud.compute.v1.RouterInterface
// crd.type: ComputeRouterInterface
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
	registry.RegisterModel(krm.ComputeRouterInterfaceGVK, NewRouterInterfaceModel)
}

func NewRouterInterfaceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &routerInterfaceModel{config: config}, nil
}

var _ directbase.Model = &routerInterfaceModel{}

type routerInterfaceModel struct {
	config *config.ControllerConfig
}

func (m *routerInterfaceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeRouterInterface{}
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
	resource := ComputeRouterInterfaceSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &RouterInterfaceAdapter{
		gcpClient: routersClient,
		id:        id.(*krm.ComputeRouterInterfaceIdentity),
		desired:   resource,
		reader:    reader,
	}, nil
}

func (m *routerInterfaceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeRouterInterfaceIdentity{}
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

	return &RouterInterfaceAdapter{
		gcpClient: routersClient,
		id:        id,
	}, nil
}

type RouterInterfaceAdapter struct {
	gcpClient *compute.RoutersClient
	id        *krm.ComputeRouterInterfaceIdentity
	desired   *computepb.RouterInterface
	actual    *computepb.RouterInterface
	router    *computepb.Router
	reader    client.Reader
}

var _ directbase.Adapter = &RouterInterfaceAdapter{}

func (a *RouterInterfaceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRouter for ComputeRouterInterface", "name", a.id)

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

	for _, iface := range router.Interfaces {
		if iface.GetName() == a.id.Interface {
			a.actual = iface
			return true, nil
		}
	}

	return false, nil
}

func (a *RouterInterfaceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRouterInterface", "name", a.id)

	if a.router == nil {
		// Just in case, fetch the router
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

	a.desired.Name = proto.String(a.id.Interface)

	// Check if already exists in a.router.Interfaces
	exists := false
	for _, iface := range a.router.Interfaces {
		if iface.GetName() == a.id.Interface {
			exists = true
			break
		}
	}

	if !exists {
		a.router.Interfaces = append(a.router.Interfaces, a.desired)
	}

	patchRouter := &computepb.Router{
		Interfaces: a.router.Interfaces,
	}

	req := &computepb.PatchRouterRequest{
		Project:        a.id.Project,
		Region:         a.id.Region,
		Router:         a.id.Router,
		RouterResource: patchRouter,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeRouterInterface %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeRouterInterface %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeRouterInterface in gcp", "name", a.id)

	// Fetch latest router to get the created interface
	latestRouter, err := a.getRouter(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeRouter %s: %w", a.id.Router, err)
	}
	a.router = latestRouter

	var created *computepb.RouterInterface
	for _, iface := range latestRouter.Interfaces {
		if iface.GetName() == a.id.Interface {
			created = iface
			break
		}
	}
	if created == nil {
		return fmt.Errorf("created interface %s not found in latest router configuration", a.id.Interface)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *RouterInterfaceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRouterInterface", "name", a.id)

	diffs, _, err := compareComputeRouterInterface(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *computepb.RouterInterface
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

		a.desired.Name = proto.String(a.id.Interface)

		// Update our interface in a.router.Interfaces
		found := false
		for i, iface := range a.router.Interfaces {
			if iface.GetName() == a.id.Interface {
				a.router.Interfaces[i] = a.desired
				found = true
				break
			}
		}
		if !found {
			a.router.Interfaces = append(a.router.Interfaces, a.desired)
		}

		patchRouter := &computepb.Router{
			Interfaces: a.router.Interfaces,
		}

		req := &computepb.PatchRouterRequest{
			Project:        a.id.Project,
			Region:         a.id.Region,
			Router:         a.id.Router,
			RouterResource: patchRouter,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute ComputeRouterInterface %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute ComputeRouterInterface", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute ComputeRouterInterface %s waiting for update: %w", a.id.String(), err)
		}

		latestRouter, err := a.getRouter(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeRouter %s: %w", a.id.Router, err)
		}
		a.router = latestRouter

		for _, iface := range latestRouter.Interfaces {
			if iface.GetName() == a.id.Interface {
				updated = iface
				break
			}
		}
		if updated == nil {
			return fmt.Errorf("updated interface %s not found in latest router configuration", a.id.Interface)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *RouterInterfaceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeRouterInterface{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeRouterInterfaceSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Region = a.id.Region
	obj.Spec.RouterRef = krm.ComputeRouterRef{
		External: fmt.Sprintf("projects/%s/regions/%s/routers/%s", a.id.Project, a.id.Region, a.id.Router),
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Interface)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Interface)
	u.SetGroupVersionKind(krm.ComputeRouterInterfaceGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *RouterInterfaceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRouterInterface", "name", a.id)

	// Fetch the router first to get the current list of interfaces
	req := &computepb.GetRouterRequest{
		Project: a.id.Project,
		Region:  a.id.Region,
		Router:  a.id.Router,
	}
	router, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.Info("parent ComputeRouter not found, assuming interface is deleted", "router", a.id.Router)
			return true, nil
		}
		return false, fmt.Errorf("getting ComputeRouter %q: %w", a.id.Router, err)
	}

	a.router = router

	var newIfaces []*computepb.RouterInterface
	found := false
	for _, iface := range a.router.Interfaces {
		if iface.GetName() == a.id.Interface {
			found = true
			continue
		} else {
			// If this is a redundant interface, remove its reference from other interfaces as well
			if iface.GetRedundantInterface() == a.id.Interface {
				iface.RedundantInterface = proto.String("")
			}
			newIfaces = append(newIfaces, iface)
		}
	}

	if !found {
		log.Info("ComputeRouterInterface already deleted or not found on router", "name", a.id)
		return true, nil
	}

	patchRouter := &computepb.Router{
		Interfaces: newIfaces,
	}

	// GCE API Patch
	patchReq := &computepb.PatchRouterRequest{
		Project:        a.id.Project,
		Region:         a.id.Region,
		Router:         a.id.Router,
		RouterResource: patchRouter,
	}
	op, err := a.gcpClient.Patch(ctx, patchReq)
	if err != nil {
		return false, fmt.Errorf("deleting compute ComputeRouterInterface %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of compute ComputeRouterInterface %s: %w", a.id.String(), err)
	}

	log.Info("successfully deleted compute ComputeRouterInterface", "name", a.id)
	return true, nil
}

func (a *RouterInterfaceAdapter) getRouter(ctx context.Context) (*computepb.Router, error) {
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

func (a *RouterInterfaceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.RouterInterface) error {
	status := &krm.ComputeRouterInterfaceStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeRouterInterface(ctx context.Context, actual, desired *computepb.RouterInterface) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeRouterInterfaceSpec_v1beta1_FromProto, ComputeRouterInterfaceSpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*computepb.RouterInterface)

	populateDefaults := func(obj *computepb.RouterInterface) {
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
