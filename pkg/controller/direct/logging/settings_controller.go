// Copyright 2025 Google LLC
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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.LoggingSettingsGVK, NewSettingsModel)
}

func NewSettingsModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSettings{config: *config}, nil
}

var _ directbase.Model = &modelSettings{}

type modelSettings struct {
	config config.ControllerConfig
}

func (m *modelSettings) client(ctx context.Context) (*gcp.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConfigClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %w", err)
	}
	return gcpClient, err
}

func (m *modelSettings) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingSettings{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSettingsIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get logging GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SettingsAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSettings) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type SettingsAdapter struct {
	id        *krm.SettingsIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.LoggingSettings
	actual    *loggingpb.Settings
}

var _ directbase.Adapter = &SettingsAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SettingsAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Settings", "name", a.id)

	req := &loggingpb.GetSettingsRequest{Name: a.id.String()}
	settingspb, err := a.gcpClient.GetSettings(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Settings %q: %w", a.id, err)
	}

	a.actual = settingspb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SettingsAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Settings", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingSettingsSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &loggingpb.UpdateSettingsRequest{
		Name:     a.id.String(),
		Settings: resource,
	}
	created, err := a.gcpClient.UpdateSettings(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Settings %s: %w", a.id, err)
	}

	if err != nil {
		return fmt.Errorf("Settings %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Settings", "name", a.id)

	status := &krm.LoggingSettingsStatus{}
	status.ObservedState = LoggingSettingsObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SettingsAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Settings", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := LoggingSettingsSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.LoggingSettingsStatus{}
		status.ObservedState = LoggingSettingsObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &loggingpb.UpdateSettingsRequest{
		Name:       a.id.String(),
		UpdateMask: updateMask,
		Settings:   desiredPb,
	}
	updated, err := a.gcpClient.UpdateSettings(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Settings %s: %w", a.id, err)
	}

	if err != nil {
		return fmt.Errorf("Settings %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Settings", "name", a.id)

	status := &krm.LoggingSettingsStatus{}
	status.ObservedState = LoggingSettingsObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SettingsAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingSettings{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingSettingsSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingSettingsGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SettingsAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return true, nil
}
