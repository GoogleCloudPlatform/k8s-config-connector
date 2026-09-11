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

	gcp "cloud.google.com/go/modelarmor/apiv1"
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ModelArmorFloorSettingGVK, NewFloorSettingModel, registry.CannotBeDeleted())
}

func NewFloorSettingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &floorSettingModel{config: *config}, nil
}

var _ directbase.Model = &floorSettingModel{}

type floorSettingModel struct {
	config config.ControllerConfig
}

func (m *floorSettingModel) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "global" {
		endpoint := fmt.Sprintf("modelarmor.%s.rep.googleapis.com:443", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelArmor REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *floorSettingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorFloorSetting{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	floorSettingID := id.(*krm.ModelArmorFloorSettingIdentity)

	gcpClient, err := m.client(ctx, floorSettingID.Location)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ModelArmorFloorSettingSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &floorSettingAdapter{
		id:        floorSettingID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *floorSettingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type floorSettingAdapter struct {
	id        *krm.ModelArmorFloorSettingIdentity
	gcpClient *gcp.Client
	desired   *pb.FloorSetting
	actual    *pb.FloorSetting
}

var _ directbase.Adapter = &floorSettingAdapter{}

func (a *floorSettingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelArmorFloorSetting", "name", a.id)

	req := &pb.GetFloorSettingRequest{Name: a.id.String()}
	floorSettingpb, err := a.gcpClient.GetFloorSetting(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelArmorFloorSetting %q: %w", a.id, err)
	}

	a.actual = floorSettingpb
	return true, nil
}

func (a *floorSettingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ModelArmorFloorSetting", "name", a.id)

	a.desired.Name = a.id.String()

	req := &pb.UpdateFloorSettingRequest{
		FloorSetting: a.desired,
	}

	created, err := a.gcpClient.UpdateFloorSetting(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ModelArmorFloorSetting %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ModelArmorFloorSetting", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *floorSettingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelArmorFloorSetting", "name", a.id.String())

	diffs, updateMask, err := a.diff(ctx)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateFloorSettingRequest{
		FloorSetting: a.desired,
		UpdateMask:   updateMask,
	}

	// Make sure Name is set on update
	req.FloorSetting.Name = a.id.String()

	updated, err := a.gcpClient.UpdateFloorSetting(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ModelArmorFloorSetting %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *floorSettingAdapter) diff(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := ModelArmorFloorSettingSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := ModelArmorFloorSettingSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.FloorSetting)

	return common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
}

func (a *floorSettingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FloorSetting) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ModelArmorFloorSettingStatus{}
	status.ObservedState = ModelArmorFloorSettingObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *floorSettingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.Location = direct.PtrTo(a.id.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Location)
	u.SetGroupVersionKind(krm.ModelArmorFloorSettingGVK)

	u.Object = uObj
	return u, nil
}

func (a *floorSettingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ModelArmorFloorSetting is a no-op", "name", a.id)
	return true, nil
}
