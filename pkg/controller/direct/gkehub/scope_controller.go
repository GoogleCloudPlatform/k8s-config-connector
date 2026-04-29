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

	gkehubapi "google.golang.org/api/gkehub/v1beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubScopeGVK, getGkeHubScopeModel)
}

func getGkeHubScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubScopeModel{config: config}, nil
}

type gkeHubScopeModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubScopeModel{}

type gkeHubScopeAdapter struct {
	id        *krm.GKEHubScopeIdentity
	desired   *krm.GKEHubScope
	actual    *gkehubapi.Scope
	hubClient *gkeHubClient
}

var _ directbase.Adapter = &gkeHubScopeAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubScopeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.Object
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

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	return &gkeHubScopeAdapter{
		id:        id.(*krm.GKEHubScopeIdentity),
		desired:   obj,
		hubClient: hubClient,
	}, nil
}

func (m *gkeHubScopeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubScopeAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil {
		return false, nil
	}
	actual, err := a.hubClient.scopeClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEHubScope %q: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *gkeHubScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.scopeClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting GKEHubScope %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for GKEHubScope deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEHubScope", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeSpecKRMtoAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()
	op, err := a.hubClient.scopeClient.Create(parent, desired).ScopeId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating GKEHubScope %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubScope creation %q: %w", a.id.String(), err)
	}

	// After creation, we need to get the latest state
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubScope after creation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created GKEHubScope", "id", a.id.String())

	// Update status
	status := GKEHubScopeStatusAPIToKRM(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEHubScope", "id", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeSpecKRMtoAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if !reflect.DeepEqual(a.actual.NamespaceLabels, desired.NamespaceLabels) {
		updateMask := "namespaceLabels"
		op, err := a.hubClient.scopeClient.Patch(a.id.String(), desired).UpdateMask(updateMask).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching GKEHubScope %q: %w", a.id.String(), err)
		}
		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for GKEHubScope update %q: %w", a.id.String(), err)
		}

		// After update, we need to get the latest state
		if _, err := a.Find(ctx); err != nil {
			return fmt.Errorf("getting GKEHubScope after update %q: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully updated GKEHubScope", "id", a.id.String())

	// Update status
	status := GKEHubScopeStatusAPIToKRM(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	mapCtx := &direct.MapContext{}
	spec := GKEHubScopeSpecAPIToKRM(mapCtx, a.actual, a.id)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.GKEHubScope{
		Spec: *spec,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubScopeAdapter) waitForOp(ctx context.Context, op *gkehubapi.Operation) error {
	retryPeriod := 5 * time.Second
	timeoutDuration := 20 * time.Minute
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
		if retryPeriod < 30*time.Second {
			retryPeriod = retryPeriod * 2
		}
	}
}
