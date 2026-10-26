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
// krm.group: bigqueryanalyticshub.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: BigQueryAnalyticsHubQueryTemplate

package bigqueryanalyticshub

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigquery/analyticshub/apiv1"
	analyticshubpb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryAnalyticsHubQueryTemplateGVK, NewQueryTemplateModel)
}

func NewQueryTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelQueryTemplate{config: *config}, nil
}

var _ directbase.Model = &modelQueryTemplate{}

type modelQueryTemplate struct {
	config config.ControllerConfig
}

func (m *modelQueryTemplate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryAnalyticsHubQueryTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.BigQueryAnalyticsHubQueryTemplateIdentity)

	// Get bigqueryanalyticshub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QueryTemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelQueryTemplate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//analyticshub.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BigQueryAnalyticsHubQueryTemplateIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QueryTemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelQueryTemplate) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building QueryTemplate client: %w", err)
	}
	return gcpClient, err
}

type QueryTemplateAdapter struct {
	id        *krm.BigQueryAnalyticsHubQueryTemplateIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryAnalyticsHubQueryTemplate
	actual    *analyticshubpb.QueryTemplate
}

var _ directbase.Adapter = &QueryTemplateAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *QueryTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting QueryTemplate", "name", a.id)

	req := &analyticshubpb.GetQueryTemplateRequest{
		Name: a.id.String(),
	}
	querytemplatepb, err := a.gcpClient.GetQueryTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting QueryTemplate %q: %w", a.id, err)
	}

	a.actual = querytemplatepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QueryTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating QueryTemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryAnalyticsHubQueryTemplateSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &analyticshubpb.CreateQueryTemplateRequest{
		Parent:          a.id.ParentString(),
		QueryTemplateId: a.id.QueryTemplate,
		QueryTemplate:   resource,
	}
	created, err := a.gcpClient.CreateQueryTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating QueryTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created QueryTemplate", "name", a.id)

	status := &krm.BigQueryAnalyticsHubQueryTemplateStatus{}
	status.ObservedState = BigQueryAnalyticsHubQueryTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QueryTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating QueryTemplate", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		BigQueryAnalyticsHubQueryTemplateSpec_FromProto,
		BigQueryAnalyticsHubQueryTemplateSpec_ToProto,
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
		req := &analyticshubpb.UpdateQueryTemplateRequest{
			UpdateMask:    updateMask,
			QueryTemplate: resource,
		}
		updated, err = a.gcpClient.UpdateQueryTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating QueryTemplate %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated QueryTemplate", "name", a.id)
	}

	status := &krm.BigQueryAnalyticsHubQueryTemplateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BigQueryAnalyticsHubQueryTemplateObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *QueryTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryAnalyticsHubQueryTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryAnalyticsHubQueryTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.BigQueryAnalyticsHubQueryTemplateGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *QueryTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting QueryTemplate", "name", a.id)

	req := &analyticshubpb.DeleteQueryTemplateRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteQueryTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent QueryTemplate, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting QueryTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted QueryTemplate", "name", a.id)
	return true, nil
}
