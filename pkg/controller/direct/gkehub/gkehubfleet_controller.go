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
	"strings"
	"time"

	gkehubv1 "google.golang.org/api/gkehub/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.GKEHubFleetGVK, getGkeHubFleetModel)
}

func getGkeHubFleetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubFleetModel{config: config}, nil
}

type gkeHubFleetModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubFleetModel{}

type gkeHubFleetAdapter struct {
	id        *krm.GKEHubFleetIdentity
	reader    client.Reader
	desired   *gkehubv1.Fleet
	actual    *gkehubv1.Fleet
	hubClient *gkeHubClient
}

var _ directbase.Adapter = &gkeHubFleetAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubFleetModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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
	obj := &krm.GKEHubFleet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := GKEHubFleetSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &gkeHubFleetAdapter{
		id:        id.(*krm.GKEHubFleetIdentity),
		reader:    reader,
		desired:   desired,
		hubClient: hubClient,
	}, nil
}

func (m *gkeHubFleetModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubFleetAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil {
		return false, nil
	}
	actual, err := a.hubClient.fleetClientV1.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEHubFleet %q: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *gkeHubFleetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.fleetClientV1.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting GKEHubFleet %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for GKEHubFleet deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubFleetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEHubFleet", "id", a.id.String())

	parent := a.id.Parent()
	op, err := a.hubClient.fleetClientV1.Create(parent, a.desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating GKEHubFleet %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubFleet creation %q: %w", a.id.String(), err)
	}

	// After creation, we need to get the latest state
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubFleet after creation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created GKEHubFleet", "id", a.id.String())

	// Update status
	mapCtx := &direct.MapContext{}
	status := GKEHubFleetStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubFleetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEHubFleet", "id", a.id.String())

	var paths []string
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if a.actual.DisplayName != a.desired.DisplayName {
		paths = append(paths, "displayName")
		report.AddField("displayName", a.actual.DisplayName, a.desired.DisplayName)
	}
	if !reflect.DeepEqual(a.actual.Labels, a.desired.Labels) {
		paths = append(paths, "labels")
		report.AddField("labels", a.actual.Labels, a.desired.Labels)
	}
	if !reflect.DeepEqual(a.actual.DefaultClusterConfig, a.desired.DefaultClusterConfig) {
		paths = append(paths, "defaultClusterConfig")
		report.AddField("defaultClusterConfig", a.actual.DefaultClusterConfig, a.desired.DefaultClusterConfig)
	}

	if len(paths) > 0 {
		structuredreporting.ReportDiff(ctx, report)
		updateMask := strings.Join(paths, ",")
		op, err := a.hubClient.fleetClientV1.Patch(a.id.String(), a.desired).UpdateMask(updateMask).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching GKEHubFleet %q: %w", a.id.String(), err)
		}
		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for GKEHubFleet update %q: %w", a.id.String(), err)
		}

		// After update, we need to get the latest state
		if _, err := a.Find(ctx); err != nil {
			return fmt.Errorf("getting GKEHubFleet after update %q: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully updated GKEHubFleet", "id", a.id.String())

	// Update status
	mapCtx := &direct.MapContext{}
	status := GKEHubFleetStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubFleetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	mapCtx := &direct.MapContext{}
	spec := GKEHubFleetSpec_FromAPI(mapCtx, a.actual, a.id)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.GKEHubFleet{
		Spec: *spec,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubFleetAdapter) waitForOp(ctx context.Context, op *gkehubv1.Operation) error {
	retryPeriod := 5 * time.Second
	timeoutDuration := 20 * time.Minute
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClientV1.Get(op.Name).Context(ctx).Do()
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
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(retryPeriod):
		}
		if retryPeriod < 30*time.Second {
			retryPeriod = retryPeriod * 2
		}
	}
}
