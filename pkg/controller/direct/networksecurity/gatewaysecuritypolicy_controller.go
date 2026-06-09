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

package networksecurity

import (
	"context"
	"fmt"
	"strings"

	api "google.golang.org/api/networksecurity/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityGatewaySecurityPolicyGVK, NewGatewaySecurityPolicyModel)
}

func NewGatewaySecurityPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gatewaySecurityPolicyModel{config: *config}, nil
}

var _ directbase.Model = &gatewaySecurityPolicyModel{}

type gatewaySecurityPolicyModel struct {
	config config.ControllerConfig
}

func (m *gatewaySecurityPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityGatewaySecurityPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idIdentity

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newNetworkSecurityClient(ctx)
	if err != nil {
		return nil, err
	}
	return &gatewaySecurityPolicyAdapter{
		gcpClientWrapper: gcpClient,
		gcpClient:        client,
		id:               id,
		desired:          obj,
	}, nil
}

func (m *gatewaySecurityPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type gatewaySecurityPolicyAdapter struct {
	gcpClientWrapper *gcpClient
	gcpClient        *api.Service
	id               *krm.GatewaySecurityPolicyIdentity
	desired          *krm.NetworkSecurityGatewaySecurityPolicy
	actual           *api.GatewaySecurityPolicy
}

var _ directbase.Adapter = &gatewaySecurityPolicyAdapter{}

func (a *gatewaySecurityPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networksecurity gatewaysecuritypolicy", "name", a.id.String())

	actual, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity gatewaysecuritypolicy %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *gatewaySecurityPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networksecurity gatewaysecuritypolicy", "name", a.id.String())

	req := &api.GatewaySecurityPolicy{
		Name:                a.id.String(),
		Description:         direct.ValueOf(a.desired.Spec.Description),
		TlsInspectionPolicy: direct.ValueOf(a.desired.Spec.TlsInspectionPolicy),
	}

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Create(a.id.Parent, req).GatewaySecurityPolicyId(a.id.ID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating networksecurity gatewaysecuritypolicy %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for networksecurity gatewaysecuritypolicy %q creation: %w", a.id.String(), err)
	}
	return nil
}

func (a *gatewaySecurityPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networksecurity gatewaysecuritypolicy", "name", a.id.String())

	req := &api.GatewaySecurityPolicy{
		Name:                a.id.String(),
		Description:         direct.ValueOf(a.desired.Spec.Description),
		TlsInspectionPolicy: direct.ValueOf(a.desired.Spec.TlsInspectionPolicy),
	}

	updateMask := []string{"description", "tlsInspectionPolicy"}

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Patch(a.id.String(), req).UpdateMask(strings.Join(updateMask, ",")).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating networksecurity gatewaysecuritypolicy %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for networksecurity gatewaysecuritypolicy %q update: %w", a.id.String(), err)
	}

	return nil
}

func (a *gatewaySecurityPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityGatewaySecurityPolicy{}
	obj.SetGroupVersionKind(krm.NetworkSecurityGatewaySecurityPolicyGVK)
	obj.SetName(a.desired.GetName())

	obj.Spec.Description = direct.LazyPtr(a.actual.Description)
	obj.Spec.TlsInspectionPolicy = direct.LazyPtr(a.actual.TlsInspectionPolicy)

	obj.Status.CreateTime = direct.LazyPtr(a.actual.CreateTime)
	obj.Status.UpdateTime = direct.LazyPtr(a.actual.UpdateTime)
	obj.Status.ExternalRef = direct.LazyPtr(a.actual.Name)

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = m
	return u, nil
}

func (a *gatewaySecurityPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networksecurity gatewaysecuritypolicy", "name", a.id.String())

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity gatewaysecuritypolicy %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for networksecurity gatewaysecuritypolicy %q deletion: %w", a.id.String(), err)
	}

	return true, nil
}
