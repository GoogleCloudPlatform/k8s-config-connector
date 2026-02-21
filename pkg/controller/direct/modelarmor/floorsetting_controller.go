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

package modelarmor

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/modelarmor/apiv1"
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ModelArmorFloorSettingGVK, NewFloorSettingModel)
}

func NewFloorSettingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFloorSetting{config: *config}, nil
}

var _ directbase.Model = &modelFloorSetting{}

type modelFloorSetting struct {
	config config.ControllerConfig
}

func (m *modelFloorSetting) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building FloorSetting client: %w", err)
	}
	return gcpClient, err
}

func (m *modelFloorSetting) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorFloorSetting{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFloorSettingIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get modelarmor GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FloorSettingAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelFloorSetting) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type FloorSettingAdapter struct {
	id        *krm.FloorSettingIdentity
	gcpClient *gcp.Client
	desired   *krm.ModelArmorFloorSetting
	actual    *pb.FloorSetting
	reader    client.Reader
}

var _ directbase.Adapter = &FloorSettingAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *FloorSettingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FloorSetting", "name", a.id)

	req := &pb.GetFloorSettingRequest{Name: a.id.String()}
	floorsettingpb, err := a.gcpClient.GetFloorSetting(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FloorSetting %q: %w", a.id, err)
	}

	a.actual = floorsettingpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FloorSettingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FloorSetting (Update)", "name", a.id)

	// FloorSetting is a singleton, so "Create" is actually an "Update" in GCP
	return a.ensureFloorSetting(ctx, createOp)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *FloorSettingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return a.ensureFloorSetting(ctx, updateOp)
}

type modelArmorOperation interface {
	directbase.Operation
	GetUnstructured() *unstructured.Unstructured
}

func (a *FloorSettingAdapter) ensureFloorSetting(ctx context.Context, op modelArmorOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FloorSetting", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := ModelArmorFloorSettingSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set the name field to ensure the GCP API can identity the resource during UpdateFloorSetting().
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.ModelArmorFloorSettingStatus{}
		status.ObservedState = ModelArmorFloorSettingObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return op.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: op.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateFloorSettingRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths)},
		FloorSetting: desiredPb,
	}
	updated, err := a.gcpClient.UpdateFloorSetting(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FloorSetting %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated FloorSetting", "name", a.id.String())

	status := &krm.ModelArmorFloorSettingStatus{}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	status.ObservedState = ModelArmorFloorSettingObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *FloorSettingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ModelArmorFloorSetting{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ModelArmorFloorSettingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ModelArmorFloorSettingGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *FloorSettingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// FloorSetting is a singleton, we probably shouldn't "delete" it in the traditional sense,
	// but maybe disable enforcement?
	// Based on the requirement, we'll just return true and not do anything in GCP,
	// or we could reset it to default.
	// For now, let's just return true.
	return true, nil
}
