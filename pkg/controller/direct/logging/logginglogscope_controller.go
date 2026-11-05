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

package logging

import (
	"context"
	"fmt"
	"strings"

	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/logging/v2"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LoggingLogScopeGVK, NewLoggingLogScopeModel)
}

func NewLoggingLogScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogScope{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogScope{}

type modelLoggingLogScope struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogScope) client(ctx context.Context) (*api.Service, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newLoggingService(ctx)
}

func (m *modelLoggingLogScope) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogScope{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	service, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := LoggingLogScopeSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &LoggingLogScopeAdapter{
		id:      id.(*krm.LoggingLogScopeIdentity),
		service: service,
		desired: desired,
	}, nil
}

func (m *modelLoggingLogScope) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.LoggingLogScopeIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	service, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &LoggingLogScopeAdapter{
		id:      id,
		service: service,
	}, nil
}

type LoggingLogScopeAdapter struct {
	id      *krm.LoggingLogScopeIdentity
	service *api.Service
	desired *api.LogScope
	actual  *api.LogScope
}

var _ directbase.Adapter = &LoggingLogScopeAdapter{}

func (a *LoggingLogScopeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogScope", "name", a.id.String())

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogScope %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *LoggingLogScopeAdapter) get(ctx context.Context) (*api.LogScope, error) {
	name := a.id.String()
	if a.id.Project != "" {
		return a.service.Projects.Locations.LogScopes.Get(name).Context(ctx).Do()
	}
	if a.id.Folder != "" {
		return a.service.Folders.Locations.LogScopes.Get(name).Context(ctx).Do()
	}
	if a.id.Organization != "" {
		return a.service.Organizations.Locations.LogScopes.Get(name).Context(ctx).Do()
	}
	return nil, fmt.Errorf("unsupported parent for LoggingLogScope %s", name)
}

func (a *LoggingLogScopeAdapter) create(ctx context.Context, parent string, logScopeID string, obj *api.LogScope) (*api.LogScope, error) {
	if a.id.Project != "" {
		return a.service.Projects.Locations.LogScopes.Create(parent, obj).LogScopeId(logScopeID).Context(ctx).Do()
	}
	if a.id.Folder != "" {
		return a.service.Folders.Locations.LogScopes.Create(parent, obj).LogScopeId(logScopeID).Context(ctx).Do()
	}
	if a.id.Organization != "" {
		return a.service.Organizations.Locations.LogScopes.Create(parent, obj).LogScopeId(logScopeID).Context(ctx).Do()
	}
	return nil, fmt.Errorf("unsupported parent for LoggingLogScope %s", a.id.String())
}

func (a *LoggingLogScopeAdapter) patch(ctx context.Context, name string, obj *api.LogScope, updateMask string) (*api.LogScope, error) {
	if a.id.Project != "" {
		call := a.service.Projects.Locations.LogScopes.Patch(name, obj).Context(ctx)
		if updateMask != "" {
			call = call.UpdateMask(updateMask)
		}
		return call.Do()
	}
	if a.id.Folder != "" {
		call := a.service.Folders.Locations.LogScopes.Patch(name, obj).Context(ctx)
		if updateMask != "" {
			call = call.UpdateMask(updateMask)
		}
		return call.Do()
	}
	if a.id.Organization != "" {
		call := a.service.Organizations.Locations.LogScopes.Patch(name, obj).Context(ctx)
		if updateMask != "" {
			call = call.UpdateMask(updateMask)
		}
		return call.Do()
	}
	return nil, fmt.Errorf("unsupported parent for LoggingLogScope %s", a.id.String())
}

func (a *LoggingLogScopeAdapter) delete(ctx context.Context, name string) error {
	if a.id.Project != "" {
		_, err := a.service.Projects.Locations.LogScopes.Delete(name).Context(ctx).Do()
		return err
	}
	if a.id.Folder != "" {
		_, err := a.service.Folders.Locations.LogScopes.Delete(name).Context(ctx).Do()
		return err
	}
	if a.id.Organization != "" {
		_, err := a.service.Organizations.Locations.LogScopes.Delete(name).Context(ctx).Do()
		return err
	}
	return fmt.Errorf("unsupported parent for LoggingLogScope %s", a.id.String())
}

func (a *LoggingLogScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LoggingLogScope", "id", fqn)

	parent := a.id.ParentString()
	resourceID := a.id.LogScope

	created, err := a.create(ctx, parent, resourceID, a.desired)
	if err != nil {
		return fmt.Errorf("creating LoggingLogScope %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created LoggingLogScope", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *LoggingLogScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogScope", "name", a.id.String())

	diffResults, changedFields, err := compareScope(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffResults.HasDiff() {
		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		// Convert json field names (resourceNames, description) to snake_case for GCP updateMask
		var updateMaskParts []string
		for _, f := range changedFields {
			switch f {
			case "resourceNames", "ResourceNames":
				updateMaskParts = append(updateMaskParts, "resource_names")
			case "description", "Description":
				updateMaskParts = append(updateMaskParts, "description")
				if a.desired.Description == "" {
					a.desired.ForceSendFields = append(a.desired.ForceSendFields, "Description")
				}
			}
		}
		updateMask := strings.Join(updateMaskParts, ",")

		updated, err := a.patch(ctx, a.id.String(), a.desired, updateMask)
		if err != nil {
			return fmt.Errorf("updating LoggingLogScope %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LoggingLogScopeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.LogScope) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogScopeStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogScope{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogScopeSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.LogScope)

	if a.id.Project != "" {
		obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	} else if a.id.Folder != "" {
		obj.Spec.FolderRef = &refsv1beta1.FolderRef{External: a.id.Folder}
	} else if a.id.Organization != "" {
		obj.Spec.OrganizationRef = &refsv1beta1.OrganizationRef{External: a.id.Organization}
	} else if a.id.BillingAccount != "" {
		obj.Spec.BillingAccountRef = &billingv1alpha1.BillingAccountRef{External: a.id.BillingAccount}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.LogScope)
	u.SetGroupVersionKind(krm.LoggingLogScopeGVK)
	return u, nil
}

func (a *LoggingLogScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LoggingLogScope", "name", a.id.String())

	err := a.delete(ctx, a.id.String())
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LoggingLogScope, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LoggingLogScope %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted LoggingLogScope", "name", a.id.String())
	return true, nil
}

func compareScope(ctx context.Context, actual, desired *api.LogScope) (*structuredreporting.Diff, []string, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LoggingLogScopeSpec_FromAPI, LoggingLogScopeSpec_ToAPI)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	diffs, changedFields, err := diffs.GoogleAPI.Diff(ctx, maskedActual, desired)
	return diffs, changedFields, err
}
