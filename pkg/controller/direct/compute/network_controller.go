// Copyright 2024 Google LLC
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

	api "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddNetworkController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddNetworkController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.ComputeNetworkGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &networkModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m)
}

type networkModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &networkModel{}

type networkAdapter struct {
	projectID string

	resourceID string

	desired *krm.ComputeNetwork
	actual  *pb.Network

	*gcpClient
	networks *api.NetworksClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &networkAdapter{}

// AdapterForObject implements the Model interface.
func (m *networkModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	networks, err := m.newNetworksClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.ComputeNetwork{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Annotations["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	return &networkAdapter{
		projectID:  projectID,
		resourceID: resourceID,
		desired:    obj,
		gcpClient:  m.gcpClient,
		networks:   networks,
	}, nil
}

// Find implements the Adapter interface.
func (a *networkAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetNetworkRequest{
		Project: a.projectID,
		Network: a.resourceID,
	}
	network, err := a.networks.Get(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("network was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	a.actual = network

	return true, nil
}

// Delete implements the Adapter interface.
func (a *networkAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteNetworkRequest{
		Project: a.projectID,
		Network: a.resourceID,
	}

	op, err := a.networks.Delete(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting network: %w", err)
	}

	completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
	if err != nil {
		return false, fmt.Errorf("waiting for network deletion: %w", err)
	}
	// TODO: Move this check to wait?
	if completed.GetStatus() != pb.Operation_DONE {
		return false, fmt.Errorf("network deletion failed: %q", completed.GetStatus())
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *networkAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	network := &pb.Network{
		// IPv4Range:                             a.desired.Spec.IPv4Range,
		AutoCreateSubnetworks: (a.desired.Spec.AutoCreateSubnetworks),
		Description:           (a.desired.Spec.Description),
		EnableUlaInternalIpv6: (a.desired.Spec.EnableUlaInternalIpv6),
		InternalIpv6Range:     (a.desired.Spec.InternalIpv6Range),
		// TODO: Should be int32
		Mtu:                                   PtrTo(int32(ValueOf(a.desired.Spec.Mtu))),
		Name:                                  &a.resourceID,
		NetworkFirewallPolicyEnforcementOrder: (a.desired.Spec.NetworkFirewallPolicyEnforcementOrder),
	}

	network.RoutingConfig = &pb.NetworkRoutingConfig{
		RoutingMode: a.desired.Spec.RoutingMode,
	}

	req := &pb.InsertNetworkRequest{
		NetworkResource: network,
		Project:         a.projectID,
	}
	op, err := a.networks.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating network: %w", err)
	}
	// TODO: Return created object status

	completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
	if err != nil {
		return fmt.Errorf("waiting for network creation: %w", err)
	}
	if completed.GetStatus() != pb.Operation_DONE {
		return fmt.Errorf("network creation failed: %q", completed.GetStatus())
	}

	// Can we get this from the operation?
	created, err := a.networks.Get(ctx, &pb.GetNetworkRequest{Project: a.projectID, Network: a.resourceID})
	if err != nil {
		return fmt.Errorf("getting network after creation: %w", err)
	}
	log.V(2).Info("created network", "network", created)

	status := &krm.ComputeNetworkStatus{}
	if err := networkStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func networkStatusToKRM(in *pb.Network, out *krm.ComputeNetworkStatus) error {
	out.SelfLink = (in.SelfLink)
	out.GatewayIpv4 = (in.GatewayIPv4)
	return nil
}

// Update implements the Adapter interface.
func (a *networkAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}
	update := &pb.Network{}
	update.RoutingConfig = &pb.NetworkRoutingConfig{}

	// routingConfig.routingMode is the only field that can be updated
	if ValueOf(a.desired.Spec.RoutingMode) != a.actual.GetRoutingConfig().GetRoutingMode() {
		updateMask.Paths = append(updateMask.Paths, "routingConfig.routingMode")
		update.RoutingConfig.RoutingMode = (a.desired.Spec.RoutingMode)
	}

	// TODO: Where/how do we want to enforce immutability?

	req := &pb.PatchNetworkRequest{
		Network:         a.resourceID,
		Project:         a.projectID,
		NetworkResource: update,
	}
	if len(updateMask.Paths) != 0 {
		op, err := a.networks.Patch(ctx, req)
		if err != nil {
			return err
		}

		completed, err := a.waitForGlobalOperation(ctx, a.projectID, op.Name())
		if err != nil {
			return fmt.Errorf("waiting for network updte: %w", err)
		}
		// TODO: Move this check to wait?
		if completed.GetStatus() != pb.Operation_DONE {
			return fmt.Errorf("network update failed: %q", completed.GetStatus())
		}

	}

	// TODO: Return updated object status
	return nil
}
