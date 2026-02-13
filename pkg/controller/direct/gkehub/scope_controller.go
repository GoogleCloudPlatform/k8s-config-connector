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
	"reflect"
	"time"

	api "google.golang.org/api/gkehub/v1beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubScopeGVK, getGKEHubScopeModel)
}

func getGKEHubScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubScopeModel{config: config}, nil
}

type gkeHubScopeModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubScopeModel{}

type gkeHubScopeAdapter struct {
	id      *GKEHubScopeIdentity
	desired *krm.GKEHubScope
	actual  *api.Scope

	hubClient *gkeHubClient
}

var _ directbase.Adapter = &gkeHubScopeAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubScopeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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

	projectID, err := direct.ResolveProjectID(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	location := "global"

	return &gkeHubScopeAdapter{
		id: &GKEHubScopeIdentity{
			Project:  projectID,
			Location: location,
			ID:       resourceID,
		},
		desired:   obj,
		hubClient: hubClient,
	}, nil
}

func (m *gkeHubScopeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubScopeAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ID == "" {
		return false, nil
	}
	scope, err := a.hubClient.scopeClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting scope %q: %w", a.id.String(), err)
	}
	a.actual = scope
	return true, nil
}

func (a *gkeHubScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.scopeClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting scope %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for scope deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating scope", "id", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.krmToApi(mapCtx)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	op, err := a.hubClient.scopeClient.Create(parent, desired).ScopeId(a.id.ID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating scope %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for scope creation %q: %w", a.id.String(), err)
	}

	actual, err := a.hubClient.scopeClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting scope %q after creation: %w", a.id.String(), err)
	}
	a.actual = actual

	return a.setID(createOp.GetUnstructured())
}

func (a *gkeHubScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating scope", "id", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.krmToApi(mapCtx)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := ""
	if !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
		updateMask = "labels"
	}

	if updateMask == "" {
		return nil
	}

	op, err := a.hubClient.scopeClient.Patch(a.id.String(), desired).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating scope %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for scope update %q: %w", a.id.String(), err)
	}

	actual, err := a.hubClient.scopeClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting scope %q after update: %w", a.id.String(), err)
	}
	a.actual = actual

	return nil
}

func (a *gkeHubScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	mapCtx := &direct.MapContext{}
	obj := &krm.GKEHubScope{}
	obj.Spec = a.apiToKrm(mapCtx)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubScopeAdapter) waitForOp(ctx context.Context, op *api.Operation) error {
	retryPeriod := baseDelay
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClient.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q failed: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("operation timed out waiting for LRO after %s", timeoutDuration.String())
		}
		time.Sleep(retryPeriod)
		retryPeriod = retryPeriod * 2
	}
}

func (a *gkeHubScopeAdapter) setID(u *unstructured.Unstructured) error {
	if err := unstructured.SetNestedField(u.Object, a.id.String(), "status", "externalRef"); err != nil {
		return err
	}
	return nil
}

func (a *gkeHubScopeAdapter) krmToApi(mapCtx *direct.MapContext) *api.Scope {
	out := &api.Scope{}
	out.Labels = a.desired.Spec.Labels
	return out
}

func (a *gkeHubScopeAdapter) apiToKrm(mapCtx *direct.MapContext) krm.GKEHubScopeSpec {
	out := krm.GKEHubScopeSpec{}
	out.ProjectRef = refs.ProjectRef{External: a.id.Project}
	out.ResourceID = direct.LazyPtr(a.id.ID)
	out.Labels = a.actual.Labels
	return out
}
