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
// krm.group: networksecurity.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NetworkSecurityGatewaySecurityPolicyRule

package networksecurity

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/networksecurity/apiv1"
	networksecuritypb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityGatewaySecurityPolicyRuleGVK, NewGatewaySecurityPolicyRuleModel)
}

func NewGatewaySecurityPolicyRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGatewaySecurityPolicyRule{config: *config}, nil
}

var _ directbase.Model = &modelGatewaySecurityPolicyRule{}

type modelGatewaySecurityPolicyRule struct {
	config config.ControllerConfig
}

func (m *modelGatewaySecurityPolicyRule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityGatewaySecurityPolicyRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetworkSecurityGatewaySecurityPolicyRuleIdentity)

	// Get networksecurity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GatewaySecurityPolicyRuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGatewaySecurityPolicyRule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//networksecurity.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetworkSecurityGatewaySecurityPolicyRuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GatewaySecurityPolicyRuleAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelGatewaySecurityPolicyRule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GatewaySecurityPolicyRule client: %w", err)
	}
	return gcpClient, err
}

type GatewaySecurityPolicyRuleAdapter struct {
	id        *krm.NetworkSecurityGatewaySecurityPolicyRuleIdentity
	gcpClient *gcp.Client
	desired   *krm.NetworkSecurityGatewaySecurityPolicyRule
	actual    *networksecuritypb.GatewaySecurityPolicyRule
}

var _ directbase.Adapter = &GatewaySecurityPolicyRuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GatewaySecurityPolicyRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GatewaySecurityPolicyRule", "name", a.id)

	req := &networksecuritypb.GetGatewaySecurityPolicyRuleRequest{
		Name: a.id.String(),
	}
	gatewaysecuritypolicyrulepb, err := a.gcpClient.GetGatewaySecurityPolicyRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GatewaySecurityPolicyRule %q: %w", a.id, err)
	}

	a.actual = gatewaysecuritypolicyrulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GatewaySecurityPolicyRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GatewaySecurityPolicyRule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityGatewaySecurityPolicyRuleSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &networksecuritypb.CreateGatewaySecurityPolicyRuleRequest{
		Parent:                      a.id.ParentString(),
		GatewaySecurityPolicyRuleId: a.id.Rule,
		GatewaySecurityPolicyRule:   resource,
	}
	op, err := a.gcpClient.CreateGatewaySecurityPolicyRule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GatewaySecurityPolicyRule %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("GatewaySecurityPolicyRule %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created GatewaySecurityPolicyRule", "name", a.id)

	status := &krm.NetworkSecurityGatewaySecurityPolicyRuleStatus{}
	status.ObservedState = NetworkSecurityGatewaySecurityPolicyRuleObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GatewaySecurityPolicyRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GatewaySecurityPolicyRule", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NetworkSecurityGatewaySecurityPolicyRuleSpec_v1alpha1_FromProto,
		NetworkSecurityGatewaySecurityPolicyRuleSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &networksecuritypb.UpdateGatewaySecurityPolicyRuleRequest{
			UpdateMask:                updateMask,
			GatewaySecurityPolicyRule: resource,
		}
		op, err := a.gcpClient.UpdateGatewaySecurityPolicyRule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating GatewaySecurityPolicyRule %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("GatewaySecurityPolicyRule %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated GatewaySecurityPolicyRule", "name", a.id)
	}

	status := &krm.NetworkSecurityGatewaySecurityPolicyRuleStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkSecurityGatewaySecurityPolicyRuleObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GatewaySecurityPolicyRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityGatewaySecurityPolicyRule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityGatewaySecurityPolicyRuleSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetworkSecurityGatewaySecurityPolicyRuleGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GatewaySecurityPolicyRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GatewaySecurityPolicyRule", "name", a.id)

	req := &networksecuritypb.DeleteGatewaySecurityPolicyRuleRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteGatewaySecurityPolicyRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent GatewaySecurityPolicyRule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting GatewaySecurityPolicyRule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted GatewaySecurityPolicyRule", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete GatewaySecurityPolicyRule %s: %w", a.id, err)
	}
	return true, nil
}
