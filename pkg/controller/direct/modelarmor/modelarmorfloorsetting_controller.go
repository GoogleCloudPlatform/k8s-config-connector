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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ModelArmorFloorSettingGVK, NewModelArmorFloorSettingModel)
}

func NewModelArmorFloorSettingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelModelArmorFloorSetting{config: *config}, nil
}

var _ directbase.Model = &modelModelArmorFloorSetting{}

type modelModelArmorFloorSetting struct {
	config config.ControllerConfig
}

func (m *modelModelArmorFloorSetting) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelArmorFloorSetting client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelModelArmorFloorSetting) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorFloorSetting{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ModelArmorFloorSettingSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ModelArmorFloorSettingAdapter{
		id:        id.(*krm.ModelArmorFloorSettingIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelModelArmorFloorSetting) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ModelArmorFloorSettingIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ModelArmorFloorSettingAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ModelArmorFloorSettingAdapter struct {
	id        *krm.ModelArmorFloorSettingIdentity
	gcpClient *gcp.Client
	desired   *pb.FloorSetting
	actual    *pb.FloorSetting
}

var _ directbase.Adapter = &ModelArmorFloorSettingAdapter{}

func (a *ModelArmorFloorSettingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelArmorFloorSetting", "name", a.id)

	req := &pb.GetFloorSettingRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetFloorSetting(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelArmorFloorSetting %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ModelArmorFloorSettingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating (updating) ModelArmorFloorSetting", "name", a.id)

	return a.reconcile(ctx, createOp, createOp.GetUnstructured())
}

func (a *ModelArmorFloorSettingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelArmorFloorSetting", "name", a.id.String())

	return a.reconcile(ctx, updateOp, updateOp.GetUnstructured())
}

func (a *ModelArmorFloorSettingAdapter) reconcile(ctx context.Context, op directbase.Operation, u *unstructured.Unstructured) error {
	diffs, updateMask, err := compareModelArmorFloorSetting(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() || a.actual == nil {
		if diffs.HasDiff() {
			diffs.Object = u
			structuredreporting.ReportDiff(ctx, diffs)
		}

		req := &pb.UpdateFloorSettingRequest{
			FloorSetting: a.desired,
			UpdateMask:   updateMask,
		}
		req.FloorSetting.Name = a.id.String()

		updated, err := a.gcpClient.UpdateFloorSetting(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ModelArmorFloorSetting %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, op, latest)
}

func (a *ModelArmorFloorSettingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FloorSetting) error {
	mapCtx := &direct.MapContext{}
	status := ModelArmorFloorSettingStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ModelArmorFloorSettingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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

	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName("floor-setting")
	u.SetGroupVersionKind(krm.ModelArmorFloorSettingGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *ModelArmorFloorSettingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Singletons usually cannot be deleted, or "deleting" means resetting to defaults.
	// The API does not have a DeleteFloorSetting method.
	return false, fmt.Errorf("ModelArmorFloorSetting cannot be deleted")
}

func compareModelArmorFloorSetting(ctx context.Context, actual, desired *pb.FloorSetting) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	if actual == nil {
		return &structuredreporting.Diff{}, nil, nil
	}
	maskedActual, err := mappers.OnlySpecFields(actual, ModelArmorFloorSettingSpec_FromProto, ModelArmorFloorSettingSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name

	clonedDesired := proto.Clone(desired).(*pb.FloorSetting)
	clonedDesired.Name = actual.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
