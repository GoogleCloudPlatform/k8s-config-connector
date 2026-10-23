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
// krm.group: securitycenter.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SecuritycenterEventThreatDetectionCustomModule

package securitycenter

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/securitycenter/apiv1"
	securitycenterpb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecuritycenterEventThreatDetectionCustomModuleGVK, NewEventThreatDetectionCustomModuleModel)
}

func NewEventThreatDetectionCustomModuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEventThreatDetectionCustomModule{config: *config}, nil
}

var _ directbase.Model = &modelEventThreatDetectionCustomModule{}

type modelEventThreatDetectionCustomModule struct {
	config config.ControllerConfig
}

func (m *modelEventThreatDetectionCustomModule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecuritycenterEventThreatDetectionCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SecuritycenterEventThreatDetectionCustomModuleIdentity)

	// Get securitycenter GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EventThreatDetectionCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEventThreatDetectionCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//securitycenter.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SecuritycenterEventThreatDetectionCustomModuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EventThreatDetectionCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelEventThreatDetectionCustomModule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EventThreatDetectionCustomModule client: %w", err)
	}
	return gcpClient, err
}

type EventThreatDetectionCustomModuleAdapter struct {
	id        *krm.SecuritycenterEventThreatDetectionCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecuritycenterEventThreatDetectionCustomModule
	actual    *securitycenterpb.EventThreatDetectionCustomModule
}

var _ directbase.Adapter = &EventThreatDetectionCustomModuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EventThreatDetectionCustomModuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EventThreatDetectionCustomModule", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &securitycenterpb.GetEventThreatDetectionCustomModuleRequest{
		Name: a.id.String(),
	}
	eventthreatdetectioncustommodulepb, err := a.gcpClient.GetEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EventThreatDetectionCustomModule %q: %w", a.id, err)
	}

	a.actual = eventthreatdetectioncustommodulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EventThreatDetectionCustomModuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecuritycenterEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securitycenterpb.CreateEventThreatDetectionCustomModuleRequest{
		Parent:                           a.id.ParentString(),
		EventThreatDetectionCustomModule: resource,
	}
	created, err := a.gcpClient.CreateEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating EventThreatDetectionCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created EventThreatDetectionCustomModule", "name", a.id)

	status := &krm.SecuritycenterEventThreatDetectionCustomModuleStatus{}
	status.ObservedState = SecuritycenterEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.SecuritycenterEventThreatDetectionCustomModuleIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Module = parsedID.Module
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EventThreatDetectionCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EventThreatDetectionCustomModule", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SecuritycenterEventThreatDetectionCustomModuleSpec_FromProto,
		SecuritycenterEventThreatDetectionCustomModuleSpec_ToProto,
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
		req := &securitycenterpb.UpdateEventThreatDetectionCustomModuleRequest{
			UpdateMask:                       updateMask,
			EventThreatDetectionCustomModule: resource,
		}
		updated, err = a.gcpClient.UpdateEventThreatDetectionCustomModule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating EventThreatDetectionCustomModule %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated EventThreatDetectionCustomModule", "name", a.id)
	}

	status := &krm.SecuritycenterEventThreatDetectionCustomModuleStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SecuritycenterEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EventThreatDetectionCustomModuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecuritycenterEventThreatDetectionCustomModule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecuritycenterEventThreatDetectionCustomModuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SecuritycenterEventThreatDetectionCustomModuleGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EventThreatDetectionCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EventThreatDetectionCustomModule", "name", a.id)

	req := &securitycenterpb.DeleteEventThreatDetectionCustomModuleRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent EventThreatDetectionCustomModule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting EventThreatDetectionCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EventThreatDetectionCustomModule", "name", a.id)
	return true, nil
}
