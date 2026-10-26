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
// krm.group: dataform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DataformReleaseConfig

package dataform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dataform/apiv1"
	dataformpb "cloud.google.com/go/dataform/apiv1/dataformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataformReleaseConfigGVK, NewReleaseConfigModel)
}

func NewReleaseConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReleaseConfig{config: *config}, nil
}

var _ directbase.Model = &modelReleaseConfig{}

type modelReleaseConfig struct {
	config config.ControllerConfig
}

func (m *modelReleaseConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataformReleaseConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataformReleaseConfigIdentity)

	// Get dataform GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReleaseConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReleaseConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dataform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DataformReleaseConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReleaseConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelReleaseConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReleaseConfig client: %w", err)
	}
	return gcpClient, err
}

type ReleaseConfigAdapter struct {
	id        *krm.DataformReleaseConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.DataformReleaseConfig
	actual    *dataformpb.ReleaseConfig
}

var _ directbase.Adapter = &ReleaseConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReleaseConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReleaseConfig", "name", a.id)

	req := &dataformpb.GetReleaseConfigRequest{
		Name: a.id.String(),
	}
	releaseconfigpb, err := a.gcpClient.GetReleaseConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReleaseConfig %q: %w", a.id, err)
	}

	a.actual = releaseconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReleaseConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReleaseConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataformReleaseConfigSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataformpb.CreateReleaseConfigRequest{
		Parent:          a.id.ParentString(),
		ReleaseConfigId: a.id.ReleaseConfig,
		ReleaseConfig:   resource,
	}
	created, err := a.gcpClient.CreateReleaseConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReleaseConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ReleaseConfig", "name", a.id)

	status := &krm.DataformReleaseConfigStatus{}
	status.ObservedState = DataformReleaseConfigObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReleaseConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReleaseConfig", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataformReleaseConfigSpec_v1alpha1_FromProto,
		DataformReleaseConfigSpec_v1alpha1_ToProto,
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
		req := &dataformpb.UpdateReleaseConfigRequest{
			UpdateMask:    updateMask,
			ReleaseConfig: resource,
		}
		updated, err = a.gcpClient.UpdateReleaseConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ReleaseConfig %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ReleaseConfig", "name", a.id)
	}

	status := &krm.DataformReleaseConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataformReleaseConfigObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReleaseConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataformReleaseConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataformReleaseConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataformReleaseConfigGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReleaseConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReleaseConfig", "name", a.id)

	req := &dataformpb.DeleteReleaseConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteReleaseConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ReleaseConfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ReleaseConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ReleaseConfig", "name", a.id)
	return true, nil
}
