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

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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

	id, err := krm.NewNetworkIdentity(ctx, reader, obj)
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

	return &NetworkAdapter{
		networksClient: networksClient,
		routesClient:   routesClient,
		id:             id,
		desired:        obj,
		reader:         reader,
	}, nil
}

func (m *networkModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkAdapter struct {
	networksClient *compute.NetworksClient
	routesClient   *compute.RoutesClient
	id             *krm.NetworkIdentity
	desired        *krm.ComputeNetwork
	actual         *computepb.Network
	reader         client.Reader
}

var _ directbase.Adapter = &NetworkAdapter{}

func (a *NetworkAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ID() == "" {
		return false, nil
	}

	req := &computepb.GetNetworkRequest{
		Project: a.id.Parent().ProjectID,
		Network: a.id.ID(),
	}
	actual, err := a.networksClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetwork %q: %w", a.id, err)
	}
	a.actual = actual
	return true, nil
}

func (a *NetworkAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName("ComputeNetwork")
	log.V(2).Info("creating ComputeNetwork", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeNetworkSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if err := mapCtx.Err(); err != nil {
		return err
	}

	req := &computepb.InsertNetworkRequest{
		Project:         a.id.Parent().ProjectID,
		NetworkResource: desiredProto,
	}
	opWait, err := a.networksClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetwork %s: %w", a.id, err)
	}

	if err := opWait.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for creation of ComputeNetwork %s: %w", a.id, err)
	}

	// Fetch the created resource to get the SelfLink
	actual, err := a.networksClient.Get(ctx, &computepb.GetNetworkRequest{
		Project: a.id.Parent().ProjectID,
		Network: a.id.ID(),
	})
	if err != nil {
		return fmt.Errorf("fetching created ComputeNetwork %s: %w", a.id, err)
	}

	// Post-creation: delete default routes if requested
	if direct.ValueOf(a.desired.Spec.DeleteDefaultRoutesOnCreate) {
		if err := a.deleteDefaultRoutes(ctx, actual.GetSelfLink()); err != nil {
			return fmt.Errorf("deleting default routes for ComputeNetwork %s: %w", a.id, err)
		}
	}

	status := &krm.ComputeNetworkStatus{}
	status.ObservedState = ComputeNetworkObservedState_v1beta1_FromProto(mapCtx, actual)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.GatewayIpv4 = actual.GatewayIPv4
	status.SelfLink = actual.SelfLink
	return op.UpdateStatus(ctx, status, nil)
}

func (a *NetworkAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName("ComputeNetwork")
	log.V(2).Info("updating ComputeNetwork", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredProto := ComputeNetworkSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if err := mapCtx.Err(); err != nil {
		return err
	}

	// For ComputeNetwork, many fields are Immutable.
	// We use Patch to update RoutingConfig and other mutable fields.
	// Check if any field actually changed.
	// Note: We don't have a generic diff yet for ComputeNetwork, but we can rely on Patch being safe.

	report := &structuredreporting.Diff{Object: op.GetUnstructured()}
	// TODO: Add fields to report if they changed
	structuredreporting.ReportDiff(ctx, report)

	req := &computepb.PatchNetworkRequest{
		Project:         a.id.Parent().ProjectID,
		Network:         a.id.ID(),
		NetworkResource: desiredProto,
	}
	opWait, err := a.networksClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeNetwork %s: %w", a.id, err)
	}

	if err := opWait.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for update of ComputeNetwork %s: %w", a.id, err)
	}

	actual, err := a.networksClient.Get(ctx, &computepb.GetNetworkRequest{
		Project: a.id.Parent().ProjectID,
		Network: a.id.ID(),
	})
	if err != nil {
		return fmt.Errorf("fetching updated ComputeNetwork %s: %w", a.id, err)
	}

	status := &krm.ComputeNetworkStatus{}
	status.ObservedState = ComputeNetworkObservedState_v1beta1_FromProto(mapCtx, actual)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.GatewayIpv4 = actual.GatewayIPv4
	status.SelfLink = actual.SelfLink
	return op.UpdateStatus(ctx, status, nil)
}

func (a *NetworkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNetwork{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNetworkSpec_v1beta1_FromProto(mapCtx, a.actual))
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ComputeNetworkGVK)

	u.Object = uObj
	return u, nil
}

func (a *NetworkAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName("ComputeNetwork")
	log.V(2).Info("deleting ComputeNetwork", "name", a.id)

	req := &computepb.DeleteNetworkRequest{
		Project: a.id.Parent().ProjectID,
		Network: a.id.ID(),
	}
	opWait, err := a.networksClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeNetwork %s: %w", a.id, err)
	}

	if err := opWait.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for deletion of ComputeNetwork %s: %w", a.id, err)
	}

	return true, nil
}

func (a *NetworkAdapter) deleteDefaultRoutes(ctx context.Context, networkSelfLink string) error {
	projectID := a.id.Parent().ProjectID
	filter := fmt.Sprintf("(network=\"%s\") AND (destRange=\"0.0.0.0/0\")", networkSelfLink)
	
	req := &computepb.ListRoutesRequest{
		Project: projectID,
		Filter:  &filter,
	}
	
	it := a.routesClient.List(ctx, req)
	for {
		route, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("listing routes: %w", err)
		}
		
		deleteReq := &computepb.DeleteRouteRequest{
			Project: projectID,
			Route:   route.GetName(),
		}
		opWait, err := a.routesClient.Delete(ctx, deleteReq)
		if err != nil {
			return fmt.Errorf("deleting route %s: %w", route.GetName(), err)
		}
		if err := opWait.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for deletion of route %s: %w", route.GetName(), err)
		}
	}
	return nil
}
