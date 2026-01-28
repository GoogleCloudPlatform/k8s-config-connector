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

	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.SpannerInstanceConfigGVK, NewInstanceConfigModel)
}

func NewInstanceConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &instanceConfigModel{config: config}, nil
}

var _ directbase.Model = &instanceConfigModel{}

type instanceConfigModel struct {
	config *config.ControllerConfig
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

	// Get spanner GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	instanceConfigClient, err := gcpClient.newInstanceAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	return &instanceConfigAdapter{
		gcpClient: instanceConfigClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *instanceConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type instanceConfigAdapter struct {
	gcpClient *instance.InstanceAdminClient
	id        *krm.InstanceConfigIdentity
	desired   *krm.SpannerInstanceConfig
	actual    *instancepb.InstanceConfig
}

var _ directbase.Adapter = &instanceConfigAdapter{}

func (a *instanceConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting spanner instance config", "name", a.id)

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
	log.V(2).Info("creating spanner instance config", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := SpannerInstanceConfigSpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()

	req := &instancepb.CreateInstanceConfigRequest{
		Parent:           a.id.Parent().String(),
		InstanceConfigId: a.id.ID(),
		InstanceConfig:   resource,
		// An option to validate, but not actually execute, a request,
		// and provide the same response.
		// todo: shall we support this field in spec?
		ValidateOnly: false,
	}
	op, err := a.gcpClient.CreateInstanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating spanner instance config %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("spanner instance config %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created spanner instance config in gcp", "name", a.id)

	status := &krm.SpannerInstanceConfigStatus{}
	status.ObservedState = SpannerInstanceConfigObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// SpannerInstanceConfig supports update.
func (a *instanceConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating spanner instance config", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := SpannerInstanceConfigSpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()

	// Check the spec for changes
	// Only display_name and labels can be updated.
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		report.AddField("display_name", a.actual.DisplayName, resource.DisplayName)
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, resource.Labels)
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field to update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	// Update
	req := &instancepb.UpdateInstanceConfigRequest{
		InstanceConfig: resource,
		UpdateMask:     updateMask,
		// An option to validate, but not actually execute, a request,
		// and provide the same response.
		// todo: shall we support this field in spec?
		ValidateOnly: false,
	}
	_, err := a.gcpClient.UpdateInstanceConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating spanner instance config %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated spanner instance config in gcp", "name", a.id)

	// Update status
	status := &krm.SpannerInstanceConfigStatus{}
	status.ObservedState = SpannerInstanceConfigObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *instanceConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting spanner instance config", "name", a.id)

	req := &instancepb.DeleteInstanceConfigRequest{
		Name: a.id.String(),
		// An option to validate, but not actually execute, a request,
		// and provide the same response.
		// todo: shall we support this field in spec?
		ValidateOnly: false,
	}
	err := a.gcpClient.DeleteInstanceConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent spanner instance config, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting spanner instance config %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted spanner instance config", "name", a.id)

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
	obj.Spec.ProjectRef = &refv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.SpannerInstanceConfigGVK)

	u.Object = uObj
	return u, nil
}
