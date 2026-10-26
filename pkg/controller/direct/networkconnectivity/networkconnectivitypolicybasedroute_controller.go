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

// +generated:controller
// krm.group: networkconnectivity.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NetworkConnectivityPolicyBasedRoute

package networkconnectivity

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/networkconnectivity/apiv1"
	networkconnectivitypb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityPolicyBasedRouteGVK, NewPolicyBasedRouteModel)
}

func NewPolicyBasedRouteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPolicyBasedRoute{config: *config}, nil
}

var _ directbase.Model = &modelPolicyBasedRoute{}

type modelPolicyBasedRoute struct {
	config config.ControllerConfig
}

func (m *modelPolicyBasedRoute) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityPolicyBasedRoute{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetworkConnectivityPolicyBasedRouteIdentity)

	// Get networkconnectivity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PolicyBasedRouteAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPolicyBasedRoute) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//networkconnectivity.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetworkConnectivityPolicyBasedRouteIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PolicyBasedRouteAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPolicyBasedRoute) client(ctx context.Context) (*gcp.PolicyBasedRoutingClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewPolicyBasedRoutingClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PolicyBasedRoute client: %w", err)
	}
	return gcpClient, err
}

type PolicyBasedRouteAdapter struct {
	id        *krm.NetworkConnectivityPolicyBasedRouteIdentity
	gcpClient *gcp.PolicyBasedRoutingClient
	desired   *krm.NetworkConnectivityPolicyBasedRoute
	actual    *networkconnectivitypb.PolicyBasedRoute
}

var _ directbase.Adapter = &PolicyBasedRouteAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PolicyBasedRouteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PolicyBasedRoute", "name", a.id)

	req := &networkconnectivitypb.GetPolicyBasedRouteRequest{
		Name: a.id.String(),
	}
	policybasedroutepb, err := a.gcpClient.GetPolicyBasedRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PolicyBasedRoute %q: %w", a.id, err)
	}

	a.actual = policybasedroutepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PolicyBasedRouteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PolicyBasedRoute", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkConnectivityPolicyBasedRouteSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &networkconnectivitypb.CreatePolicyBasedRouteRequest{
		Parent:             a.id.ParentString(),
		PolicyBasedRouteId: a.id.Policybasedroutesid,
		PolicyBasedRoute:   resource,
	}
	op, err := a.gcpClient.CreatePolicyBasedRoute(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PolicyBasedRoute %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PolicyBasedRoute %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PolicyBasedRoute", "name", a.id)

	status := &krm.NetworkConnectivityPolicyBasedRouteStatus{}
	status.ObservedState = NetworkConnectivityPolicyBasedRouteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PolicyBasedRouteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PolicyBasedRoute", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NetworkConnectivityPolicyBasedRouteSpec_FromProto,
		NetworkConnectivityPolicyBasedRouteSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating NetworkConnectivityPolicyBasedRoute is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.NetworkConnectivityPolicyBasedRouteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkConnectivityPolicyBasedRouteObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PolicyBasedRouteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityPolicyBasedRoute{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityPolicyBasedRouteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetworkConnectivityPolicyBasedRouteGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PolicyBasedRouteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PolicyBasedRoute", "name", a.id)

	req := &networkconnectivitypb.DeletePolicyBasedRouteRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeletePolicyBasedRoute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PolicyBasedRoute, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PolicyBasedRoute %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PolicyBasedRoute", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete PolicyBasedRoute %s: %w", a.id, err)
	}
	return true, nil
}
