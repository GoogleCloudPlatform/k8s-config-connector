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

package gkehub

import (
	"context"
	"fmt"
	"time"

	gkehubv1 "google.golang.org/api/gkehub/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubScopeGVK, func(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
		return &scopeModel{config: config}, nil
	})
}

type scopeModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &scopeModel{}

type scopeAdapter struct {
	projectID string
	location  string
	scopeID   string

	desired *krm.GKEHubScope
	actual  *gkehubv1.Scope

	hubClient *gkeHubClient
}

var _ directbase.Adapter = &scopeAdapter{}

func (m *scopeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	hubClient, err := gcpClient.newGkeHubClient(ctx)
	if err != nil {
		return nil, err
	}
	obj := &krm.GKEHubScope{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectRef := &refs.ProjectRef{
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
		External:  obj.Spec.ProjectRef.External,
	}
	project, err := refs.ResolveProject(ctx, reader, u.GetNamespace(), projectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := "global"
	if obj.Spec.Location != nil {
		location = *obj.Spec.Location
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.Name
	}

	return &scopeAdapter{
		projectID: projectID,
		location:  location,
		scopeID:   resourceID,
		desired:   obj,
		hubClient: hubClient,
	}, nil
}

func (m *scopeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *scopeAdapter) Find(ctx context.Context) (bool, error) {
	if a.scopeID == "" {
		return false, nil
	}
	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	actual, err := a.hubClient.scopeClient.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting scope %q: %w", name, err)
	}
	a.actual = actual
	return true, nil
}

func (a *scopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	op, err := a.hubClient.scopeClient.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting scope %q: %w", name, err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for scope %q deletion: %w", name, err)
	}
	return true, nil
}

func (a *scopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)
	op, err := a.hubClient.scopeClient.Create(parent, desired).ScopeId(a.scopeID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating scope: %w", err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for scope creation: %w", err)
	}

	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	actual, err := a.hubClient.scopeClient.Get(name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting scope after creation: %w", err)
	}
	a.actual = actual

	status := GKEHubScopeStatus_FromAPI(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *scopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	op, err := a.hubClient.scopeClient.Patch(name, desired).UpdateMask("labels,namespaceLabels").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating scope %q: %w", name, err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for scope update %q: %w", name, err)
	}

	actual, err := a.hubClient.scopeClient.Get(name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting scope after update: %w", err)
	}
	a.actual = actual

	status := GKEHubScopeStatus_FromAPI(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *scopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *scopeAdapter) waitForOp(ctx context.Context, op *gkehubv1.Operation) error {
	return a.waitForOpName(ctx, op.Name)
}

func (a *scopeAdapter) waitForOpName(ctx context.Context, opName string) error {
	retryPeriod := 5 * time.Second
	for {
		op, err := a.hubClient.v1OperationClient.Get(opName).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", opName, err)
		}
		if op.Done {
			if op.Error != nil {
				return fmt.Errorf("operation %q failed: %s", opName, op.Error.Message)
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(retryPeriod):
			if retryPeriod < 30*time.Second {
				retryPeriod = retryPeriod * 2
			}
		}
	}
}
