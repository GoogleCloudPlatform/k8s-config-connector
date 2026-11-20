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
// krm.group: ces.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CESToolset

package ces

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/ces/apiv1"
	cespb "cloud.google.com/go/ces/apiv1/cespb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CESToolsetGVK, NewToolsetModel)
}

func NewToolsetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelToolset{config: *config}, nil
}

var _ directbase.Model = &modelToolset{}

type modelToolset struct {
	config config.ControllerConfig
}

func (m *modelToolset) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CESToolset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CESToolsetIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get ces GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ToolsetAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelToolset) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//ces.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CESToolsetIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ToolsetAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelToolset) client(ctx context.Context) (*gcp.AgentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAgentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Toolset client: %w", err)
	}
	return gcpClient, err
}

type ToolsetAdapter struct {
	id        *krm.CESToolsetIdentity
	gcpClient *gcp.AgentClient
	desired   *krm.CESToolset
	actual    *cespb.Toolset
}

var _ directbase.Adapter = &ToolsetAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ToolsetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Toolset", "name", a.id)

	req := &cespb.GetToolsetRequest{
		Name: a.id.String(),
	}
	toolsetpb, err := a.gcpClient.GetToolset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Toolset %q: %w", a.id, err)
	}

	a.actual = toolsetpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ToolsetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Toolset", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CESToolsetSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cespb.CreateToolsetRequest{
		Parent:    a.id.ParentString(),
		ToolsetId: a.id.Toolset,
		Toolset:   resource,
	}
	created, err := a.gcpClient.CreateToolset(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Toolset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Toolset", "name", a.id)

	status := &krm.CESToolsetStatus{}
	status.ObservedState = CESToolsetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ToolsetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Toolset", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CESToolsetSpec_FromProto,
		CESToolsetSpec_ToProto,
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
		req := &cespb.UpdateToolsetRequest{
			UpdateMask: updateMask,
			Toolset:    resource,
		}
		updated, err = a.gcpClient.UpdateToolset(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Toolset %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Toolset", "name", a.id)
	}

	status := &krm.CESToolsetStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CESToolsetObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ToolsetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CESToolset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CESToolsetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CESToolsetGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ToolsetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Toolset", "name", a.id)

	req := &cespb.DeleteToolsetRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteToolset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Toolset, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Toolset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Toolset", "name", a.id)
	return true, nil
}
