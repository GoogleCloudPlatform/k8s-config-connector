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
// proto.service: google.cloud.compute.v1.Networks
// proto.message: google.cloud.compute.v1.Network
// crd.type: ComputeNetwork
// crd.version: v1beta1

package compute

import (
	"context"
	"errors"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkGVK, NewNetworkModel)
}

func NewNetworkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkModel{config: config}, nil
}

var _ directbase.Model = &networkModel{}

type networkModel struct {
	config *config.ControllerConfig
}

func (m *networkModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNetwork{}
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
	networksClient, err := gcpClient.newNetworksClient(ctx)
	if err != nil {
		return nil, err
	}
	routesClient, err := gcpClient.newRoutesClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeNetworkSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = direct.LazyPtr(id.(*krm.NetworkIdentity).Network)

	deleteDefaultRoutesOnCreate := false
	if obj.Spec.DeleteDefaultRoutesOnCreate != nil {
		deleteDefaultRoutesOnCreate = *obj.Spec.DeleteDefaultRoutesOnCreate
	}

	return &NetworkAdapter{
		gcpClient:                   networksClient,
		routesClient:                routesClient,
		id:                          id.(*krm.NetworkIdentity),
		desired:                     desired,
		deleteDefaultRoutesOnCreate: deleteDefaultRoutesOnCreate,
		reader:                      reader,
	}, nil
}

func (m *networkModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkAdapter struct {
	gcpClient                   *compute.NetworksClient
	routesClient                *compute.RoutesClient
	id                          *krm.NetworkIdentity
	desired                     *computepb.Network
	actual                      *computepb.Network
	deleteDefaultRoutesOnCreate bool
	reader                      client.Reader
}

var _ directbase.Adapter = &NetworkAdapter{}

func (a *NetworkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNetwork", "name", a.id)

	req := &computepb.GetNetworkRequest{
		Project: a.id.Project,
		Network: a.id.Network,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetwork %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *NetworkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNetwork", "name", a.id)

	req := &computepb.InsertNetworkRequest{
		Project:         a.id.Project,
		NetworkResource: a.desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetwork %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNetwork %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created ComputeNetwork in gcp", "name", a.id)

	if a.deleteDefaultRoutesOnCreate {
		if err := a.deleteDefaultRoutes(ctx); err != nil {
			return err
		}
	}

	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeNetwork after creation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *NetworkAdapter) deleteDefaultRoutes(ctx context.Context) error {
	log := klog.FromContext(ctx)
	log.Info("deleting default routes for ComputeNetwork", "network", a.id.Network)

	networkSelfLink := ""
	if a.actual != nil {
		networkSelfLink = *a.actual.SelfLink
	} else {
		network, err := a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeNetwork to resolve SelfLink: %w", err)
		}
		networkSelfLink = *network.SelfLink
	}

	filter := fmt.Sprintf("(network=\"%s\") AND (destRange=\"0.0.0.0/0\")", networkSelfLink)
	req := &computepb.ListRoutesRequest{
		Project: a.id.Project,
		Filter:  direct.LazyPtr(filter),
	}

	it := a.routesClient.List(ctx, req)
	for {
		route, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return fmt.Errorf("listing routes for ComputeNetwork %s: %w", a.id, err)
		}

		log.Info("deleting default route", "route", *route.Name)
		delReq := &computepb.DeleteRouteRequest{
			Project: a.id.Project,
			Route:   *route.Name,
		}
		op, err := a.routesClient.Delete(ctx, delReq)
		if err != nil {
			return fmt.Errorf("deleting route %s: %w", *route.Name, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for route deletion %s: %w", *route.Name, err)
		}
	}
	return nil
}

func (a *NetworkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNetwork", "name", a.id)

	diffs, err := compareNetwork(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for ComputeNetwork", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &computepb.PatchNetworkRequest{
		Project:         a.id.Project,
		Network:         a.id.Network,
		NetworkResource: a.desired,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeNetwork %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNetwork %s waiting for update: %w", a.id.String(), err)
	}

	updated, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeNetwork %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *NetworkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNetwork{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNetworkSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeNetworkGVK)

	u.Object = uObj
	return u, nil
}

func (a *NetworkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetwork", "name", a.id)

	req := &computepb.DeleteNetworkRequest{
		Project: a.id.Project,
		Network: a.id.Network,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeNetwork %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeNetwork", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeNetwork %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *NetworkAdapter) get(ctx context.Context) (*computepb.Network, error) {
	getReq := &computepb.GetNetworkRequest{
		Project: a.id.Project,
		Network: a.id.Network,
	}
	return a.gcpClient.Get(ctx, getReq)
}

func (a *NetworkAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *computepb.Network) error {
	mapCtx := &direct.MapContext{}
	status := ComputeNetworkStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareNetwork(ctx context.Context, actual, desired *computepb.Network) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeNetworkSpec_v1beta1_FromProto, ComputeNetworkSpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*computepb.Network)

	populateDefaults := func(obj *computepb.Network) {
		// No defaults need custom mapping, but we define pattern anyway
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
