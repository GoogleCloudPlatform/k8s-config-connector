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
// krm.group: backupdr.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: BackupDRRestoreTemplate

package backupdr

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/backupdr/apiv1"
	backupdrpb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BackupDRRestoreTemplateGVK, NewRestoreTemplateModel)
}

func NewRestoreTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelRestoreTemplate{config: *config}, nil
}

var _ directbase.Model = &modelRestoreTemplate{}

type modelRestoreTemplate struct {
	config config.ControllerConfig
}

func (m *modelRestoreTemplate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BackupDRRestoreTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.BackupDRRestoreTemplateIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get backupdr GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &RestoreTemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelRestoreTemplate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//backupdr.googleapis.com/") {
		return nil, nil
	}
	id := &krm.BackupDRRestoreTemplateIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &RestoreTemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelRestoreTemplate) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building RestoreTemplate client: %w", err)
	}
	return gcpClient, err
}

type RestoreTemplateAdapter struct {
	id        *krm.BackupDRRestoreTemplateIdentity
	gcpClient *gcp.Client
	desired   *krm.BackupDRRestoreTemplate
	actual    *backupdrpb.RestoreTemplate
}

var _ directbase.Adapter = &RestoreTemplateAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *RestoreTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting RestoreTemplate", "name", a.id)

	req := &backupdrpb.GetRestoreTemplateRequest{
		Name: a.id.String(),
	}
	restoretemplatepb, err := a.gcpClient.GetRestoreTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting RestoreTemplate %q: %w", a.id, err)
	}

	a.actual = restoretemplatepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RestoreTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating RestoreTemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BackupDRRestoreTemplateSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &backupdrpb.CreateRestoreTemplateRequest{
		Parent:            a.id.ParentString(),
		RestoreTemplateId: a.id.RestoreTemplate,
		RestoreTemplate:   resource,
	}
	op, err := a.gcpClient.CreateRestoreTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating RestoreTemplate %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("RestoreTemplate %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created RestoreTemplate", "name", a.id)

	status := &krm.BackupDRRestoreTemplateStatus{}
	status.ObservedState = BackupDRRestoreTemplateObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RestoreTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating RestoreTemplate", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		BackupDRRestoreTemplateSpec_v1alpha1_FromProto,
		BackupDRRestoreTemplateSpec_v1alpha1_ToProto,
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
		req := &backupdrpb.UpdateRestoreTemplateRequest{
			UpdateMask:      updateMask,
			RestoreTemplate: resource,
		}
		op, err := a.gcpClient.UpdateRestoreTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating RestoreTemplate %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("RestoreTemplate %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated RestoreTemplate", "name", a.id)
	}

	status := &krm.BackupDRRestoreTemplateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BackupDRRestoreTemplateObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *RestoreTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BackupDRRestoreTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupDRRestoreTemplateSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.BackupDRRestoreTemplateGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *RestoreTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting RestoreTemplate", "name", a.id)

	req := &backupdrpb.DeleteRestoreTemplateRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteRestoreTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent RestoreTemplate, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting RestoreTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted RestoreTemplate", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete RestoreTemplate %s: %w", a.id, err)
	}
	return true, nil
}
