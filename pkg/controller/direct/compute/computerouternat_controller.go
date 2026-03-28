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

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeRouterNATGVK, NewRouterNATModel)
}

func NewRouterNATModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &routerNATModel{config: config}, nil
}

type routerNATModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &routerNATModel{}

type routerNATAdapter struct {
	id            *krm.RouterNATIdentity
	routersClient *compute.RoutersClient
	desired       *krm.ComputeRouterNAT
	actual        *computepb.RouterNat
	reader        client.Reader
}

var _ directbase.Adapter = &routerNATAdapter{}

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
	routerNATID := id.(*krm.RouterNATIdentity)

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	routersClient, err := gcpClient.newRoutersClient(ctx)
	if err != nil {
		return nil, err
	}

	return &routerNATAdapter{
		id:            routerNATID,
		routersClient: routersClient,
		desired:       obj,
		reader:        reader,
	}, nil
}

func (m *routerNATModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *routerNATAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeRouterNAT", "name", a.id)

	router, err := a.getRouter(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Router %s: %w", a.id.RouterName, err)
	}

	for _, nat := range router.Nats {
		if nat.GetName() == a.id.ResourceID {
			a.actual = nat
			return true, nil
		}
	}

	return false, nil
}

func (a *routerNATAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeRouterNAT", "name", a.id)

	if err := resolveRouterNATRefs(ctx, a.reader, a.desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeRouterNATSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredProto.Name = direct.PtrTo(a.id.ResourceID)

	router, err := a.getRouter(ctx)
	if err != nil {
		return fmt.Errorf("getting Router %s: %w", a.id.RouterName, err)
	}

	// Check if already exists in list (shouldn't happen if Find returned false, but to be safe)
	found := false
	for i, nat := range router.Nats {
		if nat.GetName() == a.id.ResourceID {
			router.Nats[i] = desiredProto
			found = true
			break
		}
	}
	if !found {
		router.Nats = append(router.Nats, desiredProto)
	}

	if err := a.patchRouter(ctx, router.Nats); err != nil {
		return fmt.Errorf("patching Router %s for NAT %s: %w", a.id.RouterName, a.id.ResourceID, err)
	}

	log.V(2).Info("successfully created ComputeRouterNAT", "name", a.id)

	status := &krm.ComputeRouterNATStatus{}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *routerNATAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeRouterNAT", "name", a.id)

	if err := resolveRouterNATRefs(ctx, a.reader, a.desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeRouterNATSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredProto.Name = direct.PtrTo(a.id.ResourceID)

	// Check if update is needed
	paths, err := common.CompareProtoMessage(desiredProto, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	router, err := a.getRouter(ctx)
	if err != nil {
		return fmt.Errorf("getting Router %s: %w", a.id.RouterName, err)
	}

	// Find and replace in nats list
	found := false
	for i, nat := range router.Nats {
		if nat.GetName() == a.id.ResourceID {
			router.Nats[i] = desiredProto
			found = true
			break
		}
	}
	if !found {
		router.Nats = append(router.Nats, desiredProto)
	}

	if err := a.patchRouter(ctx, router.Nats); err != nil {
		return fmt.Errorf("patching Router %s for NAT %s: %w", a.id.RouterName, a.id.ResourceID, err)
	}

	log.V(2).Info("successfully updated ComputeRouterNAT", "name", a.id)

	status := &krm.ComputeRouterNATStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *routerNATAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeRouterNAT", "name", a.id)

	router, err := a.getRouter(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("getting Router %s: %w", a.id.RouterName, err)
	}

	// Remove from nats list
	newNats := make([]*computepb.RouterNat, 0, len(router.Nats))
	found := false
	for _, nat := range router.Nats {
		if nat.GetName() == a.id.ResourceID {
			found = true
			continue
		}
		newNats = append(newNats, nat)
	}

	if !found {
		log.V(2).Info("NAT not found in Router, considering deleted", "name", a.id)
		return true, nil
	}

	if err := a.patchRouter(ctx, newNats); err != nil {
		return false, fmt.Errorf("patching Router %s to remove NAT %s: %w", a.id.RouterName, a.id.ResourceID, err)
	}

	log.V(2).Info("successfully deleted ComputeRouterNAT", "name", a.id)
	return true, nil
}

func (a *routerNATAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("ComputeRouterNAT %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeRouterNATSpec_v1beta1_FromProto(mc, a.actual)
	if mc.Err() != nil {
		return nil, mc.Err()
	}

	// Add region and routerRef back as they are not in the proto RouterNat but required in KRM
	spec.Region = a.id.ParentID.Location
	spec.RouterRef = refs.ComputeRouterRef{
		External: fmt.Sprintf("projects/%s/regions/%s/routers/%s", a.id.ParentID.ProjectID, a.id.ParentID.Location, a.id.RouterName),
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting ComputeRouterNAT spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ComputeRouterNATGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *routerNATAdapter) getRouter(ctx context.Context) (*computepb.Router, error) {
	req := &computepb.GetRouterRequest{
		Project: a.id.ParentID.ProjectID,
		Region:  a.id.ParentID.Location,
		Router:  a.id.RouterName,
	}
	return a.routersClient.Get(ctx, req)
}

func (a *routerNATAdapter) patchRouter(ctx context.Context, nats []*computepb.RouterNat) error {
	// We only patch the nats field.
	req := &computepb.PatchRouterRequest{
		Project: a.id.ParentID.ProjectID,
		Region:  a.id.ParentID.Location,
		Router:  a.id.RouterName,
		RouterResource: &computepb.Router{
			Nats: nats,
		},
	}
	op, err := a.routersClient.Patch(ctx, req)
	if err != nil {
		return err
	}
	return op.Wait(ctx)
}
