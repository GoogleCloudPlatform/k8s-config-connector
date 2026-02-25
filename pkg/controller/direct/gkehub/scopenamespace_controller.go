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
	"strings"
	"time"

	gkehubv1 "google.golang.org/api/gkehub/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubNamespaceGVK, func(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
		return &namespaceModel{config: config}, nil
	})
}

type namespaceModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &namespaceModel{}

type namespaceAdapter struct {
	projectID        string
	location         string
	scopeID          string
	scopeNamespaceID string

	desired *krm.GKEHubNamespace
	actual  *gkehubv1.Namespace

	hubClient *gkeHubClient
}

var _ directbase.Adapter = &namespaceAdapter{}

func (m *namespaceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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
	obj := &krm.GKEHubNamespace{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID, err := resolveProjectID(ctx, reader, obj.Namespace, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	scope, err := resolveScopeRef(ctx, reader, obj, projectID)
	if err != nil {
		return nil, err
	}

	scopeID := ""
	location := "global"
	// Parse scope ID and location from scope.id (format: projects/{project}/locations/{location}/scopes/{scope})
	tokens := strings.Split(scope.id, "/")
	if len(tokens) == 6 {
		location = tokens[3]
		scopeID = tokens[5]
	} else {
		return nil, fmt.Errorf("unexpected format for scope reference: %s", scope.id)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.Name
	}

	return &namespaceAdapter{
		projectID:        projectID,
		location:         location,
		scopeID:          scopeID,
		scopeNamespaceID: resourceID,
		desired:          obj,
		hubClient:        hubClient,
	}, nil
}

func (m *namespaceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *namespaceAdapter) Find(ctx context.Context) (bool, error) {
	if a.scopeNamespaceID == "" {
		return false, nil
	}
	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	actual, err := a.hubClient.namespaceClient.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting namespace %q: %w", name, err)
	}
	a.actual = actual
	return true, nil
}

func (a *namespaceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	op, err := a.hubClient.namespaceClient.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting namespace %q: %w", name, err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for namespace %q deletion: %w", name, err)
	}
	return true, nil
}

func (a *namespaceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := GKEHubNamespaceSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", a.projectID, a.location, a.scopeID)
	op, err := a.hubClient.namespaceClient.Create(parent, desired).ScopeNamespaceId(a.scopeNamespaceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating namespace: %w", err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for namespace creation: %w", err)
	}

	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	actual, err := a.hubClient.namespaceClient.Get(name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting namespace after creation: %w", err)
	}
	a.actual = actual

	status := GKEHubNamespaceStatus_FromAPI(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *namespaceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := GKEHubNamespaceSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	op, err := a.hubClient.namespaceClient.Patch(name, desired).UpdateMask("labels,namespaceLabels").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating namespace %q: %w", name, err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for namespace update %s: %w", name, err)
	}

	actual, err := a.hubClient.namespaceClient.Get(name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting namespace after update: %w", err)
	}
	a.actual = actual

	status := GKEHubNamespaceStatus_FromAPI(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", a.projectID, a.location, a.scopeID, a.scopeNamespaceID)
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *namespaceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil // TODO
}

func (a *namespaceAdapter) waitForOp(ctx context.Context, op *gkehubv1.Operation) error {
	return a.waitForOpName(ctx, op.Name)
}

func (a *namespaceAdapter) waitForOpName(ctx context.Context, opName string) error {
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
