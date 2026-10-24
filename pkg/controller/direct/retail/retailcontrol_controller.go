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
// krm.group: retail.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: RetailControl

package retail

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/retail/apiv2"
	retailpb "cloud.google.com/go/retail/apiv2/retailpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.RetailControlGVK, NewControlModel)
}

func NewControlModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelControl{config: *config}, nil
}

var _ directbase.Model = &modelControl{}

type modelControl struct {
	config config.ControllerConfig
}

func (m *modelControl) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RetailControl{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.RetailControlIdentity)

	// Get retail GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ControlAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelControl) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//retail.googleapis.com/") {
		return nil, nil
	}
	id := &krm.RetailControlIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ControlAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelControl) client(ctx context.Context) (*gcp.ControlClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewControlRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Control client: %w", err)
	}
	return gcpClient, err
}

type ControlAdapter struct {
	id        *krm.RetailControlIdentity
	gcpClient *gcp.ControlClient
	desired   *krm.RetailControl
	actual    *retailpb.Control
}

var _ directbase.Adapter = &ControlAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ControlAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Control", "name", a.id)

	req := &retailpb.GetControlRequest{
		Name: a.id.String(),
	}
	controlpb, err := a.gcpClient.GetControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Control %q: %w", a.id, err)
	}

	a.actual = controlpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ControlAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Control", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := RetailControlSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &retailpb.CreateControlRequest{
		Parent:    a.id.ParentString(),
		ControlId: a.id.Control,
		Control:   resource,
	}
	created, err := a.gcpClient.CreateControl(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Control %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Control", "name", a.id)

	status := &krm.RetailControlStatus{}
	status.ObservedState = RetailControlObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ControlAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Control", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		RetailControlSpec_FromProto,
		RetailControlSpec_ToProto,
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
		req := &retailpb.UpdateControlRequest{
			UpdateMask: updateMask,
			Control:    resource,
		}
		updated, err = a.gcpClient.UpdateControl(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Control %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Control", "name", a.id)
	}

	status := &krm.RetailControlStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = RetailControlObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ControlAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.RetailControl{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RetailControlSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.RetailControlGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ControlAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Control", "name", a.id)

	req := &retailpb.DeleteControlRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Control, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Control %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Control", "name", a.id)
	return true, nil
}
