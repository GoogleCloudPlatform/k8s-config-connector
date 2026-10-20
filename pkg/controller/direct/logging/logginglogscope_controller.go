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
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/logging/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LoggingLogScopeGVK, NewLoggingLogScopeModel)
}

func NewLoggingLogScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogScope{config: config}, nil
}

type modelLoggingLogScope struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &modelLoggingLogScope{}

func (m *modelLoggingLogScope) client(ctx context.Context) (*api.ProjectsLocationsLogScopesService, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newProjectsLocationsLogScopesService(ctx)
}

func (m *modelLoggingLogScope) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.LoggingLogScope{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idVal.(*krm.LoggingLogScopeIdentity)

	client, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &loggingLogScopeAdapter{
		client:  client,
		id:      id,
		desired: obj,
	}, nil
}

func (m *modelLoggingLogScope) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type loggingLogScopeAdapter struct {
	client  *api.ProjectsLocationsLogScopesService
	id      *krm.LoggingLogScopeIdentity
	desired *krm.LoggingLogScope
	actual  *api.LogScope
}

var _ directbase.Adapter = &loggingLogScopeAdapter{}

func (a *loggingLogScopeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	name := a.id.String()
	log.V(2).Info("getting LoggingLogScope", "name", name)

	actual, err := a.client.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogScope %q: %w", name, err)
	}

	a.actual = actual
	return true, nil
}

func (a *loggingLogScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	name := a.id.String()
	log.V(2).Info("creating LoggingLogScope", "name", name)

	parent := a.id.ParentString()
	logScopeID := a.id.LogScope

	toCreate := &api.LogScope{
		ResourceNames: a.desired.Spec.ResourceNames,
	}
	if a.desired.Spec.Description != nil {
		toCreate.Description = *a.desired.Spec.Description
	}

	created, err := a.client.Create(parent, toCreate).LogScopeId(logScopeID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating LoggingLogScope %q: %w", name, err)
	}

	a.actual = created

	u := createOp.GetUnstructured()
	if err := unstructured.SetNestedField(u.Object, logScopeID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *loggingLogScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	name := a.id.String()
	log.V(2).Info("updating LoggingLogScope", "name", name)

	var updateMask []string
	toUpdate := &api.LogScope{}

	desiredDescription := ""
	if a.desired.Spec.Description != nil {
		desiredDescription = *a.desired.Spec.Description
	}
	if desiredDescription != a.actual.Description {
		toUpdate.Description = desiredDescription
		updateMask = append(updateMask, "description")
	}

	if !reflect.DeepEqual(a.desired.Spec.ResourceNames, a.actual.ResourceNames) {
		toUpdate.ResourceNames = a.desired.Spec.ResourceNames
		updateMask = append(updateMask, "resourceNames")
	}

	latest := a.actual
	if len(updateMask) > 0 {
		call := a.client.Patch(name, toUpdate).UpdateMask(strings.Join(updateMask, ",")).Context(ctx)
		updated, err := call.Do()
		if err != nil {
			return fmt.Errorf("updating LoggingLogScope %q: %w", name, err)
		}
		latest = updated
	}

	a.actual = latest
	return a.updateStatus(ctx, updateOp, latest)
}

func (a *loggingLogScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	name := a.id.String()
	log.V(2).Info("deleting LoggingLogScope", "name", name)

	_, err := a.client.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting LoggingLogScope %q: %w", name, err)
	}

	return true, nil
}

func (a *loggingLogScopeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.LogScope) error {
	status := &krm.LoggingLogScopeStatus{}

	if latest != nil {
		status.ObservedState = &krm.LoggingLogScopeObservedState{
			CreateTime: direct.LazyPtr(latest.CreateTime),
			UpdateTime: direct.LazyPtr(latest.UpdateTime),
		}
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *loggingLogScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("LoggingLogScope %q not found", a.id.String())
	}

	obj := &krm.LoggingLogScope{}
	obj.SetGroupVersionKind(krm.LoggingLogScopeGVK)
	obj.SetName(a.id.LogScope)
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceNames = a.actual.ResourceNames
	if a.actual.Description != "" {
		obj.Spec.Description = direct.LazyPtr(a.actual.Description)
	}

	un, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting LoggingLogScope to unstructured: %w", err)
	}
	return &unstructured.Unstructured{Object: un}, nil
}
