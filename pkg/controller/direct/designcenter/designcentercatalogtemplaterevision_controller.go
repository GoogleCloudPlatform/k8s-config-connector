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
// krm.group: designcenter.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DesignCenterCatalogTemplateRevision

package designcenter

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/designcenter/apiv1"
	designcenterpb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DesignCenterCatalogTemplateRevisionGVK, NewCatalogTemplateRevisionModel)
}

func NewCatalogTemplateRevisionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCatalogTemplateRevision{config: *config}, nil
}

var _ directbase.Model = &modelCatalogTemplateRevision{}

type modelCatalogTemplateRevision struct {
	config config.ControllerConfig
}

func (m *modelCatalogTemplateRevision) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DesignCenterCatalogTemplateRevision{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DesignCenterCatalogTemplateRevisionIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get designcenter GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CatalogTemplateRevisionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCatalogTemplateRevision) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//designcenter.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DesignCenterCatalogTemplateRevisionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CatalogTemplateRevisionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCatalogTemplateRevision) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CatalogTemplateRevision client: %w", err)
	}
	return gcpClient, err
}

type CatalogTemplateRevisionAdapter struct {
	id        *krm.DesignCenterCatalogTemplateRevisionIdentity
	gcpClient *gcp.Client
	desired   *krm.DesignCenterCatalogTemplateRevision
	actual    *designcenterpb.CatalogTemplateRevision
}

var _ directbase.Adapter = &CatalogTemplateRevisionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CatalogTemplateRevisionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CatalogTemplateRevision", "name", a.id)

	req := &designcenterpb.GetCatalogTemplateRevisionRequest{
		Name: a.id.String(),
	}
	catalogtemplaterevisionpb, err := a.gcpClient.GetCatalogTemplateRevision(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CatalogTemplateRevision %q: %w", a.id, err)
	}

	a.actual = catalogtemplaterevisionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CatalogTemplateRevisionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CatalogTemplateRevision", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DesignCenterCatalogTemplateRevisionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &designcenterpb.CreateCatalogTemplateRevisionRequest{
		Parent:                    a.id.ParentString(),
		CatalogTemplateRevisionId: a.id.Revision,
		CatalogTemplateRevision:   resource,
	}
	op, err := a.gcpClient.CreateCatalogTemplateRevision(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CatalogTemplateRevision %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("CatalogTemplateRevision %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created CatalogTemplateRevision", "name", a.id)

	status := &krm.DesignCenterCatalogTemplateRevisionStatus{}
	status.ObservedState = DesignCenterCatalogTemplateRevisionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CatalogTemplateRevisionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CatalogTemplateRevision", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DesignCenterCatalogTemplateRevisionSpec_FromProto,
		DesignCenterCatalogTemplateRevisionSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating DesignCenterCatalogTemplateRevision is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.DesignCenterCatalogTemplateRevisionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DesignCenterCatalogTemplateRevisionObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CatalogTemplateRevisionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DesignCenterCatalogTemplateRevision{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DesignCenterCatalogTemplateRevisionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DesignCenterCatalogTemplateRevisionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CatalogTemplateRevisionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CatalogTemplateRevision", "name", a.id)

	req := &designcenterpb.DeleteCatalogTemplateRevisionRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteCatalogTemplateRevision(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CatalogTemplateRevision, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CatalogTemplateRevision %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CatalogTemplateRevision", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete CatalogTemplateRevision %s: %w", a.id, err)
	}
	return true, nil
}
