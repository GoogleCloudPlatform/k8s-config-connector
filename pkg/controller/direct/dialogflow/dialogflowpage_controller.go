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
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowPage

package dialogflow

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowPageGVK, NewPageModel)
}

func NewPageModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPage{config: *config}, nil
}

var _ directbase.Model = &modelPage{}

type modelPage struct {
	config config.ControllerConfig
}

func (m *modelPage) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowPage{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowPageIdentity)

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PageAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPage) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowPageIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PageAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPage) client(ctx context.Context) (*gcp.PagesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewPagesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Page client: %w", err)
	}
	return gcpClient, err
}

type PageAdapter struct {
	id        *krm.DialogflowPageIdentity
	gcpClient *gcp.PagesClient
	desired   *krm.DialogflowPage
	actual    *cxpb.Page
}

var _ directbase.Adapter = &PageAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PageAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Page", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cxpb.GetPageRequest{
		Name: a.id.String(),
	}
	pagepb, err := a.gcpClient.GetPage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Page %q: %w", a.id, err)
	}

	a.actual = pagepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PageAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Page", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowPageSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cxpb.CreatePageRequest{
		Parent: a.id.ParentString(),
		Page:   resource,
	}
	created, err := a.gcpClient.CreatePage(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Page %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Page", "name", a.id)

	status := &krm.DialogflowPageStatus{}
	status.ObservedState = DialogflowPageObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.DialogflowPageIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Page = parsedID.Page
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PageAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Page", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowPageSpec_FromProto,
		DialogflowPageSpec_ToProto,
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
		req := &cxpb.UpdatePageRequest{
			UpdateMask: updateMask,
			Page:       resource,
		}
		updated, err = a.gcpClient.UpdatePage(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Page %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Page", "name", a.id)
	}

	status := &krm.DialogflowPageStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DialogflowPageObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PageAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowPage{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowPageSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowPageGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PageAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Page", "name", a.id)

	req := &cxpb.DeletePageRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeletePage(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Page, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Page %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Page", "name", a.id)
	return true, nil
}
