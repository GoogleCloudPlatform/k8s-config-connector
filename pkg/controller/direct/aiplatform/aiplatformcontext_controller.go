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
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformContext

package aiplatform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformContextGVK, NewContextModel)
}

func NewContextModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelContext{config: *config}, nil
}

var _ directbase.Model = &modelContext{}

type modelContext struct {
	config config.ControllerConfig
}

func (m *modelContext) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformContext{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformContextIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &ContextAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelContext) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformContextIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &ContextAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelContext) client(ctx context.Context, location string) (*gcp.MetadataClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewMetadataClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Context client: %w", err)
	}
	return gcpClient, err
}

type ContextAdapter struct {
	id        *krm.AIPlatformContextIdentity
	gcpClient *gcp.MetadataClient
	desired   *krm.AIPlatformContext
	actual    *aiplatformpb.Context
}

var _ directbase.Adapter = &ContextAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ContextAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Context", "name", a.id)

	req := &aiplatformpb.GetContextRequest{
		Name: a.id.String(),
	}
	contextpb, err := a.gcpClient.GetContext(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Context %q: %w", a.id, err)
	}

	a.actual = contextpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ContextAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Context", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformContextSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)

	req := &aiplatformpb.CreateContextRequest{
		Parent:    a.id.ParentString(),
		ContextId: a.id.Context,
		Context:   resource,
	}
	created, err := a.gcpClient.CreateContext(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Context %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Context", "name", a.id)

	status := &krm.AIPlatformContextStatus{}
	status.ObservedState = AIPlatformContextObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ContextAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Context", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		AIPlatformContextSpec_FromProto,
		AIPlatformContextSpec_ToProto,
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
		req := &aiplatformpb.UpdateContextRequest{
			UpdateMask: updateMask,
			Context:    resource,
		}
		updated, err = a.gcpClient.UpdateContext(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Context %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Context", "name", a.id)
	}

	status := &krm.AIPlatformContextStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformContextObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ContextAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformContext{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformContextSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformContextGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ContextAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Context", "name", a.id)

	req := &aiplatformpb.DeleteContextRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteContext(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Context, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Context %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Context", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Context %s: %w", a.id, err)
	}
	return true, nil
}
