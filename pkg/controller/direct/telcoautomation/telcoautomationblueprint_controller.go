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
// krm.group: telcoautomation.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: TelcoAutomationBlueprint

package telcoautomation

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/telcoautomation/apiv1"
	telcoautomationpb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.TelcoAutomationBlueprintGVK, NewBlueprintModel)
}

func NewBlueprintModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBlueprint{config: *config}, nil
}

var _ directbase.Model = &modelBlueprint{}

type modelBlueprint struct {
	config config.ControllerConfig
}

func (m *modelBlueprint) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TelcoAutomationBlueprint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.TelcoAutomationBlueprintIdentity)

	// Get telcoautomation GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BlueprintAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBlueprint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//telcoautomation.googleapis.com/") {
		return nil, nil
	}
	id := &krm.TelcoAutomationBlueprintIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BlueprintAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelBlueprint) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Blueprint client: %w", err)
	}
	return gcpClient, err
}

type BlueprintAdapter struct {
	id        *krm.TelcoAutomationBlueprintIdentity
	gcpClient *gcp.Client
	desired   *krm.TelcoAutomationBlueprint
	actual    *telcoautomationpb.Blueprint
}

var _ directbase.Adapter = &BlueprintAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BlueprintAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Blueprint", "name", a.id)

	req := &telcoautomationpb.GetBlueprintRequest{
		Name: a.id.String(),
	}
	blueprintpb, err := a.gcpClient.GetBlueprint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Blueprint %q: %w", a.id, err)
	}

	a.actual = blueprintpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BlueprintAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Blueprint", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := TelcoAutomationBlueprintSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &telcoautomationpb.CreateBlueprintRequest{
		Parent:      a.id.ParentString(),
		BlueprintId: a.id.Blueprint,
		Blueprint:   resource,
	}
	created, err := a.gcpClient.CreateBlueprint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Blueprint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Blueprint", "name", a.id)

	status := &krm.TelcoAutomationBlueprintStatus{}
	status.ObservedState = TelcoAutomationBlueprintObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BlueprintAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Blueprint", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		TelcoAutomationBlueprintSpec_FromProto,
		TelcoAutomationBlueprintSpec_ToProto,
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
		req := &telcoautomationpb.UpdateBlueprintRequest{
			UpdateMask: updateMask,
			Blueprint:  resource,
		}
		updated, err = a.gcpClient.UpdateBlueprint(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Blueprint %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Blueprint", "name", a.id)
	}

	status := &krm.TelcoAutomationBlueprintStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = TelcoAutomationBlueprintObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BlueprintAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TelcoAutomationBlueprint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TelcoAutomationBlueprintSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.TelcoAutomationBlueprintGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BlueprintAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Blueprint", "name", a.id)

	req := &telcoautomationpb.DeleteBlueprintRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteBlueprint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Blueprint, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Blueprint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Blueprint", "name", a.id)
	return true, nil
}
