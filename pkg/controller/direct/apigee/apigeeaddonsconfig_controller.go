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

package apigee

import (
	"context"
	"fmt"
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/apigee/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ApigeeAddonsConfigGVK, NewApigeeAddonsConfigModel)
}

func NewApigeeAddonsConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeAddonsConfig{config: config}, nil
}

var _ directbase.Model = &modelApigeeAddonsConfig{}

type modelApigeeAddonsConfig struct {
	config *config.ControllerConfig
}

func (m *modelApigeeAddonsConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ApigeeAddonsConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeAddonsConfigIdentity)

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeAddonsConfigAdapter{
		id:                  id,
		k8sClient:           reader,
		organizationsClient: gcpClient.organizationsClient(),
		operationsClient:    gcpClient.operationsClient(),
		desired:             obj,
	}, nil
}

func (m *modelApigeeAddonsConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	format := krm.ApigeeAddonsConfigIdentityFormat
	id, match, err := format.Parse(url)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, nil
	}
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeAddonsConfigAdapter{
		id:                  id,
		organizationsClient: gcpClient.organizationsClient(),
		operationsClient:    gcpClient.operationsClient(),
	}, nil
}

type ApigeeAddonsConfigAdapter struct {
	id                  *krm.ApigeeAddonsConfigIdentity
	k8sClient           client.Reader
	organizationsClient *api.OrganizationsService
	operationsClient    *api.OrganizationsOperationsService
	desired             *krm.ApigeeAddonsConfig
	actual              *api.GoogleCloudApigeeV1AddonsConfig
}

var _ directbase.Adapter = &ApigeeAddonsConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeAddonsConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeAddonsConfig", "name", a.id)

	org, err := a.organizationsClient.Get(a.id.ParentString()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting organization %q for ApigeeAddonsConfig: %w", a.id.ParentString(), err)
	}

	a.actual = org.AddonsConfig
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeAddonsConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeAddonsConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	addonsConfig := GoogleCloudApigeeV1AddonsConfig_ToAPI(mapCtx, desired.Spec.AddonsConfig)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &api.GoogleCloudApigeeV1SetAddonsRequest{
		AddonsConfig: addonsConfig,
	}

	op, err := a.organizationsClient.SetAddons(a.id.ParentString(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeAddonsConfig %s: %w", a.id, err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for ApigeeAddonsConfig creation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ApigeeAddonsConfig", "name", a.id)

	status := &krm.ApigeeAddonsConfigStatus{}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ApigeeAddonsConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeAddonsConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	diff := a.diff(a.actual, a.desired.Spec.AddonsConfig, report)
	if !diff {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ApigeeAddonsConfigStatus{}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	structuredreporting.ReportDiff(ctx, report)

	desired := a.desired.DeepCopy()
	addonsConfig := GoogleCloudApigeeV1AddonsConfig_ToAPI(mapCtx, desired.Spec.AddonsConfig)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &api.GoogleCloudApigeeV1SetAddonsRequest{
		AddonsConfig: addonsConfig,
	}

	op, err := a.organizationsClient.SetAddons(a.id.ParentString(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ApigeeAddonsConfig %s: %w", a.id, err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for ApigeeAddonsConfig update %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ApigeeAddonsConfig", "name", a.id)

	status := &krm.ApigeeAddonsConfigStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeAddonsConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or resource not found")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeAddonsConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeAddonsConfigSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Org = a.id.Organization
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Organization)
	u.SetGroupVersionKind(krm.ApigeeAddonsConfigGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeAddonsConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeAddonsConfig", "name", a.id)

	req := &api.GoogleCloudApigeeV1SetAddonsRequest{}

	op, err := a.organizationsClient.SetAddons(a.id.ParentString(), req).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeAddonsConfig %s: %w", a.id, err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting for ApigeeAddonsConfig deletion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeAddonsConfig", "name", a.id)
	return true, nil
}

func (a *ApigeeAddonsConfigAdapter) diff(actual *api.GoogleCloudApigeeV1AddonsConfig, desired *krm.GoogleCloudApigeeV1AddonsConfig, report *structuredreporting.Diff) bool {
	hasDiff := false

	// AdvancedApiOpsConfig
	actualAdvancedApiOpsEnabled := actual != nil && actual.AdvancedApiOpsConfig != nil && actual.AdvancedApiOpsConfig.Enabled
	desiredAdvancedApiOpsEnabled := desired != nil && desired.AdvancedApiOpsConfig != nil && direct.ValueOf(desired.AdvancedApiOpsConfig.Enabled)
	if actualAdvancedApiOpsEnabled != desiredAdvancedApiOpsEnabled {
		report.AddField("addonsConfig.advancedApiOpsConfig.enabled", actualAdvancedApiOpsEnabled, desiredAdvancedApiOpsEnabled)
		hasDiff = true
	}

	// IntegrationConfig
	actualIntegrationEnabled := actual != nil && actual.IntegrationConfig != nil && actual.IntegrationConfig.Enabled
	desiredIntegrationEnabled := desired != nil && desired.IntegrationConfig != nil && direct.ValueOf(desired.IntegrationConfig.Enabled)
	if actualIntegrationEnabled != desiredIntegrationEnabled {
		report.AddField("addonsConfig.integrationConfig.enabled", actualIntegrationEnabled, desiredIntegrationEnabled)
		hasDiff = true
	}

	// MonetizationConfig
	actualMonetizationEnabled := actual != nil && actual.MonetizationConfig != nil && actual.MonetizationConfig.Enabled
	desiredMonetizationEnabled := desired != nil && desired.MonetizationConfig != nil && direct.ValueOf(desired.MonetizationConfig.Enabled)
	if actualMonetizationEnabled != desiredMonetizationEnabled {
		report.AddField("addonsConfig.monetizationConfig.enabled", actualMonetizationEnabled, desiredMonetizationEnabled)
		hasDiff = true
	}

	// ApiSecurityConfig
	actualApiSecurityEnabled := actual != nil && actual.ApiSecurityConfig != nil && actual.ApiSecurityConfig.Enabled
	desiredApiSecurityEnabled := desired != nil && desired.ApiSecurityConfig != nil && direct.ValueOf(desired.ApiSecurityConfig.Enabled)
	if actualApiSecurityEnabled != desiredApiSecurityEnabled {
		report.AddField("addonsConfig.apiSecurityConfig.enabled", actualApiSecurityEnabled, desiredApiSecurityEnabled)
		hasDiff = true
	}
	if desired != nil && desired.ApiSecurityConfig != nil && desired.ApiSecurityConfig.ExpiresAt != nil {
		actualExpiresAt := ""
		if actual != nil && actual.ApiSecurityConfig != nil && actual.ApiSecurityConfig.ExpiresAt != 0 {
			actualExpiresAt = strconv.FormatInt(actual.ApiSecurityConfig.ExpiresAt, 10)
		}
		if actualExpiresAt != *desired.ApiSecurityConfig.ExpiresAt {
			report.AddField("addonsConfig.apiSecurityConfig.expiresAt", actualExpiresAt, *desired.ApiSecurityConfig.ExpiresAt)
			hasDiff = true
		}
	}

	// ConnectorsPlatformConfig
	actualConnectorsPlatformEnabled := actual != nil && actual.ConnectorsPlatformConfig != nil && actual.ConnectorsPlatformConfig.Enabled
	desiredConnectorsPlatformEnabled := desired != nil && desired.ConnectorsPlatformConfig != nil && direct.ValueOf(desired.ConnectorsPlatformConfig.Enabled)
	if actualConnectorsPlatformEnabled != desiredConnectorsPlatformEnabled {
		report.AddField("addonsConfig.connectorsPlatformConfig.enabled", actualConnectorsPlatformEnabled, desiredConnectorsPlatformEnabled)
		hasDiff = true
	}
	if desired != nil && desired.ConnectorsPlatformConfig != nil && desired.ConnectorsPlatformConfig.ExpiresAt != nil {
		actualExpiresAt := ""
		if actual != nil && actual.ConnectorsPlatformConfig != nil && actual.ConnectorsPlatformConfig.ExpiresAt != 0 {
			actualExpiresAt = strconv.FormatInt(actual.ConnectorsPlatformConfig.ExpiresAt, 10)
		}
		if actualExpiresAt != *desired.ConnectorsPlatformConfig.ExpiresAt {
			report.AddField("addonsConfig.connectorsPlatformConfig.expiresAt", actualExpiresAt, *desired.ConnectorsPlatformConfig.ExpiresAt)
			hasDiff = true
		}
	}

	return hasDiff
}
