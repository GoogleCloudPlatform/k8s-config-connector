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

// +generated:controller
// krm.group: networkconnectivity.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NetworkConnectivityServiceConnectionToken

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
	registry.RegisterModel(krm.NetworkConnectivityServiceConnectionTokenGVK, NewServiceConnectionTokenModel)
}

func NewServiceConnectionTokenModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelServiceConnectionToken{config: *config}, nil
}

var _ directbase.Model = &modelServiceConnectionToken{}

type modelServiceConnectionToken struct {
	config config.ControllerConfig
}

func (m *modelServiceConnectionToken) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityServiceConnectionToken{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetworkConnectivityServiceConnectionTokenIdentity)

	// Get networkconnectivity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceConnectionTokenAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelServiceConnectionToken) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//networkconnectivity.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetworkConnectivityServiceConnectionTokenIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceConnectionTokenAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelServiceConnectionToken) client(ctx context.Context) (*gcp.CrossNetworkAutomationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCrossNetworkAutomationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServiceConnectionToken client: %w", err)
	}
	return gcpClient, err
}

type ServiceConnectionTokenAdapter struct {
	id        *krm.NetworkConnectivityServiceConnectionTokenIdentity
	gcpClient *gcp.CrossNetworkAutomationClient
	desired   *krm.NetworkConnectivityServiceConnectionToken
	actual    *networkconnectivitypb.ServiceConnectionToken
}

var _ directbase.Adapter = &ServiceConnectionTokenAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ServiceConnectionTokenAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ServiceConnectionToken", "name", a.id)

	req := &networkconnectivitypb.GetServiceConnectionTokenRequest{
		Name: a.id.String(),
	}
	serviceconnectiontokenpb, err := a.gcpClient.GetServiceConnectionToken(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServiceConnectionToken %q: %w", a.id, err)
	}

	a.actual = serviceconnectiontokenpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServiceConnectionTokenAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ServiceConnectionToken", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkConnectivityServiceConnectionTokenSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &networkconnectivitypb.CreateServiceConnectionTokenRequest{
		Parent:                   a.id.ParentString(),
		ServiceConnectionTokenId: a.id.ServiceConnectionToken,
		ServiceConnectionToken:   resource,
	}
	op, err := a.gcpClient.CreateServiceConnectionToken(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceConnectionToken %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ServiceConnectionToken %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ServiceConnectionToken", "name", a.id)

	status := &krm.NetworkConnectivityServiceConnectionTokenStatus{}
	status.ObservedState = NetworkConnectivityServiceConnectionTokenObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServiceConnectionTokenAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ServiceConnectionToken", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NetworkConnectivityServiceConnectionTokenSpec_FromProto,
		NetworkConnectivityServiceConnectionTokenSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating NetworkConnectivityServiceConnectionToken is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.NetworkConnectivityServiceConnectionTokenStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkConnectivityServiceConnectionTokenObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ServiceConnectionTokenAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityServiceConnectionToken{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityServiceConnectionTokenSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetworkConnectivityServiceConnectionTokenGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ServiceConnectionTokenAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ServiceConnectionToken", "name", a.id)

	req := &networkconnectivitypb.DeleteServiceConnectionTokenRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteServiceConnectionToken(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ServiceConnectionToken, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ServiceConnectionToken %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ServiceConnectionToken", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ServiceConnectionToken %s: %w", a.id, err)
	}
	return true, nil
}
