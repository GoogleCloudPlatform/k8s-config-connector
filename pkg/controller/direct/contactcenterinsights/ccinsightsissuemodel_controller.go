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
// krm.group: contactcenterinsights.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CCInsightsIssueModel

package contactcenterinsights

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CCInsightsIssueModelGVK, NewIssueModelModel)
}

func NewIssueModelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIssueModel{config: *config}, nil
}

var _ directbase.Model = &modelIssueModel{}

type modelIssueModel struct {
	config config.ControllerConfig
}

func (m *modelIssueModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsIssueModel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CCInsightsIssueModelIdentity)

	// Get contactcenterinsights GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IssueModelAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelIssueModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//contactcenterinsights.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CCInsightsIssueModelIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &IssueModelAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelIssueModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building IssueModel client: %w", err)
	}
	return gcpClient, err
}

type IssueModelAdapter struct {
	id        *krm.CCInsightsIssueModelIdentity
	gcpClient *gcp.Client
	desired   *krm.CCInsightsIssueModel
	actual    *contactcenterinsightspb.IssueModel
}

var _ directbase.Adapter = &IssueModelAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *IssueModelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting IssueModel", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &contactcenterinsightspb.GetIssueModelRequest{
		Name: a.id.String(),
	}
	issuemodelpb, err := a.gcpClient.GetIssueModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IssueModel %q: %w", a.id, err)
	}

	a.actual = issuemodelpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IssueModelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating IssueModel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CCInsightsIssueModelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &contactcenterinsightspb.CreateIssueModelRequest{
		Parent:     a.id.ParentString(),
		IssueModel: resource,
	}
	op, err := a.gcpClient.CreateIssueModel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IssueModel %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("IssueModel %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created IssueModel", "name", a.id)

	status := &krm.CCInsightsIssueModelStatus{}
	status.ObservedState = CCInsightsIssueModelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.CCInsightsIssueModelIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.IssueModel = parsedID.IssueModel
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *IssueModelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating IssueModel", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CCInsightsIssueModelSpec_FromProto,
		CCInsightsIssueModelSpec_ToProto,
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
		req := &contactcenterinsightspb.UpdateIssueModelRequest{
			UpdateMask: updateMask,
			IssueModel: resource,
		}
		updated, err = a.gcpClient.UpdateIssueModel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating IssueModel %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated IssueModel", "name", a.id)
	}

	status := &krm.CCInsightsIssueModelStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CCInsightsIssueModelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *IssueModelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CCInsightsIssueModel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CCInsightsIssueModelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CCInsightsIssueModelGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *IssueModelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting IssueModel", "name", a.id)

	req := &contactcenterinsightspb.DeleteIssueModelRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteIssueModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent IssueModel, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting IssueModel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted IssueModel", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete IssueModel %s: %w", a.id, err)
	}
	return true, nil
}
