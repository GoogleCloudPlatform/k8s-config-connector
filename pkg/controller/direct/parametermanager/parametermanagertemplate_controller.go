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
// krm.group: parametermanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ParameterManagerTemplate

package parametermanager

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/parametermanager/apiv1"
	parametermanagerpb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ParameterManagerTemplateGVK, NewTemplateModel)
}

func NewTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTemplate{config: *config}, nil
}

var _ directbase.Model = &modelTemplate{}

type modelTemplate struct {
	config config.ControllerConfig
}

func (m *modelTemplate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ParameterManagerTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ParameterManagerTemplateIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get parametermanager GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &TemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTemplate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//parametermanager.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ParameterManagerTemplateIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &TemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelTemplate) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("parametermanager.%s.rep.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Template client: %w", err)
	}
	return gcpClient, err
}

type TemplateAdapter struct {
	id        *krm.ParameterManagerTemplateIdentity
	gcpClient *gcp.Client
	desired   *krm.ParameterManagerTemplate
	actual    *parametermanagerpb.Template
}

var _ directbase.Adapter = &TemplateAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Template", "name", a.id)

	req := &parametermanagerpb.GetTemplateRequest{
		Name: a.id.String(),
	}
	templatepb, err := a.gcpClient.GetTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Template %q: %w", a.id, err)
	}

	a.actual = templatepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Template", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ParameterManagerTemplateSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)
	if resource.Format == parametermanagerpb.TemplateFormat_TEMPLATE_FORMAT_UNSPECIFIED {
		resource.Format = parametermanagerpb.TemplateFormat_TEMPLATE_FORMAT_YAML
	}

	req := &parametermanagerpb.CreateTemplateRequest{
		Parent:     a.id.ParentString(),
		TemplateId: a.id.Template,
		Template:   resource,
	}
	created, err := a.gcpClient.CreateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Template %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Template", "name", a.id)

	status := &krm.ParameterManagerTemplateStatus{}
	status.ObservedState = ParameterManagerTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Template", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		ParameterManagerTemplateSpec_FromProto,
		ParameterManagerTemplateSpec_ToProto,
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
		req := &parametermanagerpb.UpdateTemplateRequest{
			UpdateMask: updateMask,
			Template:   resource,
		}
		_, err = a.gcpClient.UpdateTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Template %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Template", "name", a.id)

		// Fetch fully-populated resource after update because UpdateTemplate response omits fields like labels
		getReq := &parametermanagerpb.GetTemplateRequest{
			Name: a.id.String(),
		}
		updated, err = a.gcpClient.GetTemplate(ctx, getReq)
		if err != nil {
			return fmt.Errorf("getting Template %s after update: %w", a.id, err)
		}
	}

	status := &krm.ParameterManagerTemplateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ParameterManagerTemplateObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ParameterManagerTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ParameterManagerTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ParameterManagerTemplateGVK)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Template", "name", a.id)

	req := &parametermanagerpb.DeleteTemplateRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Template, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Template %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Template", "name", a.id)
	return true, nil
}
