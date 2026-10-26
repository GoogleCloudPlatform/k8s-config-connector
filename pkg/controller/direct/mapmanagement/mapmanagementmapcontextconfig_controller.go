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
// krm.group: mapmanagement.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: MapmanagementMapContextConfig

package mapmanagement

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/maps/mapmanagement/apiv2"
	mapmanagementpb "cloud.google.com/go/maps/mapmanagement/apiv2/mapmanagementpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.MapmanagementMapContextConfigGVK, NewMapContextConfigModel)
}

func NewMapContextConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMapContextConfig{config: *config}, nil
}

var _ directbase.Model = &modelMapContextConfig{}

type modelMapContextConfig struct {
	config config.ControllerConfig
}

func (m *modelMapContextConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MapmanagementMapContextConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.MapmanagementMapContextConfigIdentity)

	// Get mapmanagement GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MapContextConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelMapContextConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//mapmanagement.googleapis.com/") {
		return nil, nil
	}
	id := &krm.MapmanagementMapContextConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MapContextConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelMapContextConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MapContextConfig client: %w", err)
	}
	return gcpClient, err
}

type MapContextConfigAdapter struct {
	id        *krm.MapmanagementMapContextConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.MapmanagementMapContextConfig
	actual    *mapmanagementpb.MapContextConfig
}

var _ directbase.Adapter = &MapContextConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MapContextConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MapContextConfig", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &mapmanagementpb.GetMapContextConfigRequest{
		Name: a.id.String(),
	}
	mapcontextconfigpb, err := a.gcpClient.GetMapContextConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MapContextConfig %q: %w", a.id, err)
	}

	a.actual = mapcontextconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MapContextConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MapContextConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MapmanagementMapContextConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &mapmanagementpb.CreateMapContextConfigRequest{
		Parent:           a.id.ParentString(),
		MapContextConfig: resource,
	}
	created, err := a.gcpClient.CreateMapContextConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MapContextConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created MapContextConfig", "name", a.id)

	status := &krm.MapmanagementMapContextConfigStatus{}
	status.ObservedState = MapmanagementMapContextConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.MapmanagementMapContextConfigIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.MapContextConfig = parsedID.MapContextConfig
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MapContextConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MapContextConfig", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		MapmanagementMapContextConfigSpec_FromProto,
		MapmanagementMapContextConfigSpec_ToProto,
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
		req := &mapmanagementpb.UpdateMapContextConfigRequest{
			UpdateMask:       updateMask,
			MapContextConfig: resource,
		}
		updated, err = a.gcpClient.UpdateMapContextConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MapContextConfig %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated MapContextConfig", "name", a.id)
	}

	status := &krm.MapmanagementMapContextConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = MapmanagementMapContextConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MapContextConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MapmanagementMapContextConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MapmanagementMapContextConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.MapmanagementMapContextConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MapContextConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MapContextConfig", "name", a.id)

	req := &mapmanagementpb.DeleteMapContextConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteMapContextConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MapContextConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MapContextConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MapContextConfig", "name", a.id)
	return true, nil
}
