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
	registry.RegisterModel(krm.NetworkSecurityGatewaySecurityPolicyRuleGVK, NewGatewaySecurityPolicyRuleModel)
}

func NewGatewaySecurityPolicyRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gatewaySecurityPolicyRuleModel{config: *config}, nil
}

var _ directbase.Model = &gatewaySecurityPolicyRuleModel{}

type gatewaySecurityPolicyRuleModel struct {
	config config.ControllerConfig
}

func (m *gatewaySecurityPolicyRuleModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityGatewaySecurityPolicyRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := obj.Spec.GatewaySecurityPolicyRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
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
	return &gatewaySecurityPolicyRuleAdapter{
		gcpClientWrapper: gcpClient,
		gcpClient:        client,
		id:               id,
		desired:          obj,
	}, nil
}

func (m *gatewaySecurityPolicyRuleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type gatewaySecurityPolicyRuleAdapter struct {
	gcpClientWrapper *gcpClient
	gcpClient        *api.Service
	id               *krm.GatewaySecurityPolicyRuleIdentity
	desired          *krm.NetworkSecurityGatewaySecurityPolicyRule
	actual           *api.GatewaySecurityPolicyRule
}

var _ directbase.Adapter = &gatewaySecurityPolicyRuleAdapter{}

func (a *gatewaySecurityPolicyRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networksecurity gatewaysecuritypolicyrule", "name", a.id.String())

	actual, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Rules.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity gatewaysecuritypolicyrule %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *gatewaySecurityPolicyRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networksecurity gatewaysecuritypolicyrule", "name", a.id.String())

	req := &api.GatewaySecurityPolicyRule{
		Name:                 a.id.String(),
		BasicProfile:         a.desired.Spec.BasicProfile,
		Enabled:              a.desired.Spec.Enabled,
		Priority:             a.desired.Spec.Priority,
		SessionMatcher:       a.desired.Spec.SessionMatcher,
		ApplicationMatcher:   direct.ValueOf(a.desired.Spec.ApplicationMatcher),
		Description:          direct.ValueOf(a.desired.Spec.Description),
		TlsInspectionEnabled: direct.ValueOf(a.desired.Spec.TlsInspectionEnabled),
	}

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Rules.Create(a.id.Parent, req).GatewaySecurityPolicyRuleId(a.id.ID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating networksecurity gatewaysecuritypolicyrule %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for networksecurity gatewaysecuritypolicyrule %q creation: %w", a.id.String(), err)
	}
	return nil
}

func (a *gatewaySecurityPolicyRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networksecurity gatewaysecuritypolicyrule", "name", a.id.String())

	req := &api.GatewaySecurityPolicyRule{
		Name:                 a.id.String(),
		BasicProfile:         a.desired.Spec.BasicProfile,
		Enabled:              a.desired.Spec.Enabled,
		Priority:             a.desired.Spec.Priority,
		SessionMatcher:       a.desired.Spec.SessionMatcher,
		ApplicationMatcher:   direct.ValueOf(a.desired.Spec.ApplicationMatcher),
		Description:          direct.ValueOf(a.desired.Spec.Description),
		TlsInspectionEnabled: direct.ValueOf(a.desired.Spec.TlsInspectionEnabled),
	}

	updateMask := []string{"basicProfile", "enabled", "priority", "sessionMatcher", "applicationMatcher", "description", "tlsInspectionEnabled"}

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Rules.Patch(a.id.String(), req).UpdateMask(strings.Join(updateMask, ",")).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating networksecurity gatewaysecuritypolicyrule %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for networksecurity gatewaysecuritypolicyrule %q update: %w", a.id.String(), err)
	}

	return nil
}

func (a *gatewaySecurityPolicyRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityGatewaySecurityPolicyRule{}
	obj.SetGroupVersionKind(krm.NetworkSecurityGatewaySecurityPolicyRuleGVK)
	obj.SetName(a.desired.GetName())

	obj.Spec.BasicProfile = a.actual.BasicProfile
	obj.Spec.Enabled = a.actual.Enabled
	obj.Spec.Priority = a.actual.Priority
	obj.Spec.SessionMatcher = a.actual.SessionMatcher
	obj.Spec.ApplicationMatcher = direct.LazyPtr(a.actual.ApplicationMatcher)
	obj.Spec.Description = direct.LazyPtr(a.actual.Description)
	obj.Spec.TlsInspectionEnabled = direct.LazyPtr(a.actual.TlsInspectionEnabled)

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

func (a *gatewaySecurityPolicyRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networksecurity gatewaysecuritypolicyrule", "name", a.id.String())

	op, err := a.gcpClient.Projects.Locations.GatewaySecurityPolicies.Rules.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity gatewaysecuritypolicyrule %q: %w", a.id.String(), err)
	}
	if err := a.gcpClientWrapper.waitOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for networksecurity gatewaysecuritypolicyrule %q deletion: %w", a.id.String(), err)
	}

	return true, nil
}
