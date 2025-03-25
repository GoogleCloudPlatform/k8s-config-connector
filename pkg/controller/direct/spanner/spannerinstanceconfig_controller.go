// Copyright 2024 Google LLC
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

// +tool:controller
// proto.service: google.spanner.admin.instance.v1.InstanceAdmin
// proto.message: google.spanner.admin.instance.v1.InstanceConfig
// crd.type: SpannerInstanceConfig
// crd.version: v1alpha1

package spanner

import (
	"context"
	"fmt"
	"reflect"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.SpannerInstanceConfigGVK, NewInstanceConfigModel)
}

func NewInstanceConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &instanceConfigModel{config: *config}, nil
}

var _ directbase.Model = &instanceConfigModel{}

type instanceConfigModel struct {
	config config.ControllerConfig
}

func (m *instanceConfigModel) Client(ctx context.Context, projectID string) (*instance.InstanceAdminClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := instance.NewInstanceAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building spanner instance config client: %w", err)
	}

	return gcpClient, err
}

func (m *instanceConfigModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpannerInstanceConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := SpannerInstanceConfigSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.Client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &instanceConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *instanceConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type instanceConfigAdapter struct {
	gcpClient *instance.InstanceAdminClient
	id        *krm.InstanceConfigIdentity
	desired   *instancepb.InstanceConfig
	actual    *instancepb.InstanceConfig
}

var _ directbase.Adapter = &instanceConfigAdapter{}

func (a *instanceConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting spanner instance config", "name", a.id)

	req := &instancepb.GetInstanceConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetInstanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting spanner instance config %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *instanceConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating spanner instance config", "name", a.id)

	req := &instancepb.CreateInstanceConfigRequest{
		Parent:            a.id.Parent().String(),
		InstanceConfigId:  a.id.ID(),
		InstanceConfig:    a.desired,
		ValidateOnly:      a.desired.GetValidateOnly(),
		ValidateOnlyProto: direct.ValueOf(a.desired.ValidateOnly),
	}
	op, err := a.gcpClient.CreateInstanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating spanner instance config %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("spanner instance config %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created spanner instance config in gcp", "name", a.id)

	status := &krm.SpannerInstanceConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SpannerInstanceConfigObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// SpannerInstanceConfig supports update.
func (a *instanceConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating spanner instance config", "name", a.id)

	// Convert the KCC spec to proto format
	desiredpb := SpannerInstanceConfigSpec_ToProto(&direct.MapContext{}, &a.desired)

	// Check the spec for changes
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if !reflect.DeepEqual(a.desired.Etag, a.actual.Etag) {
		updateMask.Paths = append(updateMask.Paths, "etag")
	}
	if !reflect.DeepEqual(a.desired.Replicas, a.actual.Replicas) {
		updateMask.Paths = append(updateMask.Paths, "replicas")
	}
	if len(updateMask.Paths) == 0 {
		log.Info("no field to update", "name", a.id)
		return nil
	}

	// Update
	req := &instancepb.UpdateInstanceConfigRequest{
		InstanceConfig: desiredpb,
		UpdateMask:     updateMask,
		ValidateOnly:   a.desired.GetValidateOnly(),
	}
	_, err := a.gcpClient.UpdateInstanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating spanner instance config %s: %w", a.id.String(), err)
	}
	log.Info("successfully updated spanner instance config in gcp", "name", a.id)

	// Update status
	status := &krm.SpannerInstanceConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SpannerInstanceConfigObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *instanceConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting spanner instance config", "name", a.id)

	req := &instancepb.DeleteInstanceConfigRequest{
		Name:         a.id.String(),
		Etag:         direct.PtrTo(a.desired.Etag),
		ValidateOnly: a.desired.GetValidateOnly(),
	}
	err := a.gcpClient.DeleteInstanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.Info("skipping delete for non-existent spanner instance config, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting spanner instance config %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted spanner instance config", "name", a.id)

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *instanceConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpannerInstanceConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpannerInstanceConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &krm.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.SpannerInstanceConfigGVK)

	u.Object = uObj
	return u, nil
}
