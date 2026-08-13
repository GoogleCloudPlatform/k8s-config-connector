// Copyright 2025 Google LLC
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
// crd.type: ComputeGlobalNetworkEndpoint
// crd.version: v1alpha1

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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ComputeGlobalNetworkEndpointGVK, NewGlobalNetworkEndpointModel)
}

func NewGlobalNetworkEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &globalNetworkEndpointModel{config: config}, nil
}

var _ directbase.Model = &globalNetworkEndpointModel{}

type globalNetworkEndpointModel struct {
	config *config.ControllerConfig
}

func (m *globalNetworkEndpointModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeGlobalNetworkEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewGlobalNetworkEndpointIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	negClient, err := gcpClient.newGlobalNetworkEndpointGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &GlobalNetworkEndpointAdapter{
		gcpClient: negClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *globalNetworkEndpointModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type GlobalNetworkEndpointAdapter struct {
	gcpClient *compute.GlobalNetworkEndpointGroupsClient
	id        *krm.GlobalNetworkEndpointIdentity
	desired   *krm.ComputeGlobalNetworkEndpoint
	actual    *computepb.NetworkEndpoint
}

var _ directbase.Adapter = &GlobalNetworkEndpointAdapter{}

// Find retrieves the GCP resource.
// Returns true if found (triggers Update), false if not found (triggers Create).
func (a *GlobalNetworkEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding GlobalNetworkEndpoint", "id", a.id)

	endpoint, err := a.findEndpoint(ctx)
	if err != nil {
		return false, err
	}
	if endpoint == nil {
		return false, nil
	}
	a.actual = endpoint
	return true, nil
}

// findEndpoint searches the NEG's endpoints for the one matching this identity.
func (a *GlobalNetworkEndpointAdapter) findEndpoint(ctx context.Context) (*computepb.NetworkEndpoint, error) {
	req := &computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest{
		Project:              a.id.Parent().ProjectID,
		NetworkEndpointGroup: a.id.Parent().GlobalNetworkEndpointGroup,
	}

	it := a.gcpClient.ListNetworkEndpoints(ctx, req)
	for {
		item, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return nil, nil
			}
			return nil, fmt.Errorf("listing GlobalNetworkEndpoints for NEG %s: %w", a.id.Parent().GlobalNetworkEndpointGroup, err)
		}
		ep := item.GetNetworkEndpoint()
		if ep == nil {
			continue
		}
		if endpointMatchesIdentity(ep, a.id) {
			return ep, nil
		}
	}
	return nil, nil
}

// endpointMatchesIdentity returns true if the given NetworkEndpoint matches the identity.
func endpointMatchesIdentity(ep *computepb.NetworkEndpoint, id *krm.GlobalNetworkEndpointIdentity) bool {
	if ep.GetPort() != id.Port() {
		return false
	}
	if id.Fqdn() != "" {
		return ep.GetFqdn() == id.Fqdn()
	}
	return ep.GetIpAddress() == id.IPAddress()
}

// Create creates the resource in GCP based on `spec`.
func (a *GlobalNetworkEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GlobalNetworkEndpoint", "id", a.id)

	endpoint := buildNetworkEndpoint(a.id)

	req := &computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest{
		Project:              a.id.Parent().ProjectID,
		NetworkEndpointGroup: a.id.Parent().GlobalNetworkEndpointGroup,
		GlobalNetworkEndpointGroupsAttachEndpointsRequestResource: &computepb.GlobalNetworkEndpointGroupsAttachEndpointsRequest{
			NetworkEndpoints: []*computepb.NetworkEndpoint{endpoint},
		},
	}

	op, err := a.gcpClient.AttachNetworkEndpoints(ctx, req)
	if err != nil {
		return fmt.Errorf("attaching GlobalNetworkEndpoint %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for GlobalNetworkEndpoint %s attach: %w", a.id, err)
	}

	log.Info("successfully attached GlobalNetworkEndpoint in GCP", "id", a.id)

	status := &krm.ComputeGlobalNetworkEndpointStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update is a no-op since all fields are immutable.
func (a *GlobalNetworkEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("GlobalNetworkEndpoint fields are immutable; no update needed", "id", a.id)

	status := &krm.ComputeGlobalNetworkEndpointStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource spec.
func (a *GlobalNetworkEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ComputeGlobalNetworkEndpoint{}
	obj.Spec.ProjectRef = &refv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.GlobalNetworkEndpointGroup = a.id.Parent().GlobalNetworkEndpointGroup
	port := a.actual.GetPort()
	obj.Spec.Port = port
	if fqdn := a.actual.GetFqdn(); fqdn != "" {
		obj.Spec.Fqdn = &fqdn
	}
	if ip := a.actual.GetIpAddress(); ip != "" {
		obj.Spec.IPAddress = &ip
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeGlobalNetworkEndpointGVK)
	u.Object = uObj
	return u, nil
}

// Delete detaches the endpoint from the NEG.
func (a *GlobalNetworkEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GlobalNetworkEndpoint", "id", a.id)

	// First, confirm it still exists.
	endpoint, err := a.findEndpoint(ctx)
	if err != nil {
		return false, fmt.Errorf("finding GlobalNetworkEndpoint %s before delete: %w", a.id, err)
	}
	if endpoint == nil {
		log.V(2).Info("GlobalNetworkEndpoint not found, presuming already deleted", "id", a.id)
		return false, nil
	}

	ep := buildNetworkEndpoint(a.id)

	req := &computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest{
		Project:              a.id.Parent().ProjectID,
		NetworkEndpointGroup: a.id.Parent().GlobalNetworkEndpointGroup,
		GlobalNetworkEndpointGroupsDetachEndpointsRequestResource: &computepb.GlobalNetworkEndpointGroupsDetachEndpointsRequest{
			NetworkEndpoints: []*computepb.NetworkEndpoint{ep},
		},
	}

	op, err := a.gcpClient.DetachNetworkEndpoints(ctx, req)
	if err != nil {
		return false, fmt.Errorf("detaching GlobalNetworkEndpoint %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for GlobalNetworkEndpoint %s detach: %w", a.id, err)
	}

	log.Info("successfully detached GlobalNetworkEndpoint from GCP", "id", a.id)
	return true, nil
}

// buildNetworkEndpoint creates the computepb.NetworkEndpoint from the identity.
func buildNetworkEndpoint(id *krm.GlobalNetworkEndpointIdentity) *computepb.NetworkEndpoint {
	port := id.Port()
	ep := &computepb.NetworkEndpoint{
		Port: &port,
	}
	if id.Fqdn() != "" {
		fqdn := id.Fqdn()
		ep.Fqdn = &fqdn
	}
	if id.IPAddress() != "" {
		ip := id.IPAddress()
		ep.IpAddress = &ip
	}
	return ep
}
