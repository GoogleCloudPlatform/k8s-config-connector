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
// krm.group: financialservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: FinancialServicesEngineConfig

package financialservices

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/financialservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/financialservices/apiv1"
	financialservicespb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FinancialServicesEngineConfigGVK, NewEngineConfigModel)
}

func NewEngineConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEngineConfig{config: *config}, nil
}

var _ directbase.Model = &modelEngineConfig{}

type modelEngineConfig struct {
	config config.ControllerConfig
}

func (m *modelEngineConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FinancialServicesEngineConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.FinancialServicesEngineConfigIdentity)

	// Get financialservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EngineConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEngineConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//financialservices.googleapis.com/") {
		return nil, nil
	}
	id := &krm.FinancialServicesEngineConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EngineConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelEngineConfig) client(ctx context.Context) (*gcp.AMLClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAMLRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EngineConfig client: %w", err)
	}
	return gcpClient, err
}

type EngineConfigAdapter struct {
	id        *krm.FinancialServicesEngineConfigIdentity
	gcpClient *gcp.AMLClient
	desired   *krm.FinancialServicesEngineConfig
	actual    *financialservicespb.EngineConfig
}

var _ directbase.Adapter = &EngineConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EngineConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EngineConfig", "name", a.id)

	req := &financialservicespb.GetEngineConfigRequest{
		Name: a.id.String(),
	}
	engineconfigpb, err := a.gcpClient.GetEngineConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EngineConfig %q: %w", a.id, err)
	}

	a.actual = engineconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EngineConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EngineConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FinancialServicesEngineConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &financialservicespb.CreateEngineConfigRequest{
		Parent:         a.id.ParentString(),
		EngineConfigId: a.id.EngineConfig,
		EngineConfig:   resource,
	}
	op, err := a.gcpClient.CreateEngineConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating EngineConfig %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("EngineConfig %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created EngineConfig", "name", a.id)

	status := &krm.FinancialServicesEngineConfigStatus{}
	status.ObservedState = FinancialServicesEngineConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EngineConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EngineConfig", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		FinancialServicesEngineConfigSpec_FromProto,
		FinancialServicesEngineConfigSpec_ToProto,
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
		req := &financialservicespb.UpdateEngineConfigRequest{
			UpdateMask:   updateMask,
			EngineConfig: resource,
		}
		op, err := a.gcpClient.UpdateEngineConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating EngineConfig %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("EngineConfig %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated EngineConfig", "name", a.id)
	}

	status := &krm.FinancialServicesEngineConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FinancialServicesEngineConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EngineConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FinancialServicesEngineConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FinancialServicesEngineConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.FinancialServicesEngineConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EngineConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EngineConfig", "name", a.id)

	req := &financialservicespb.DeleteEngineConfigRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteEngineConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent EngineConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting EngineConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EngineConfig", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete EngineConfig %s: %w", a.id, err)
	}
	return true, nil
}
