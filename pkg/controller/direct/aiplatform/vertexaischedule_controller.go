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

package aiplatform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
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
	registry.RegisterModel(krm.VertexAIScheduleGVK, NewScheduleModel)
}

func NewScheduleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &scheduleModel{config: config}, nil
}

var _ directbase.Model = &scheduleModel{}

type scheduleModel struct {
	config *config.ControllerConfig
}

func (m *scheduleModel) client(ctx context.Context) (*gcp.ScheduleClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewScheduleClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ScheduleClient client: %w", err)
	}
	return gcpClient, err
}

func (m *scheduleModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAISchedule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(reader.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get ScheduleClient client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.VertexAIScheduleIdentity)
	if !ok {
		return nil, fmt.Errorf("expected VertexAIScheduleIdentity, got %T", id)
	}

	mapCtx := &direct.MapContext{}
	desiredpb := VertexAIScheduleSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &ScheduleAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *scheduleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type ScheduleAdapter struct {
	id        *krm.VertexAIScheduleIdentity
	gcpClient *gcp.ScheduleClient
	desiredpb *pb.Schedule
	desired   *krm.VertexAISchedule
	actual    *pb.Schedule
}

var _ directbase.Adapter = &ScheduleAdapter{}

func (a *ScheduleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAISchedule", "name", a.id.String())

	req := &pb.GetScheduleRequest{
		Name: a.id.String(),
	}

	schedulepb, err := a.gcpClient.GetSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAISchedule %q: %w", a.id.String(), err)
	}

	a.actual = schedulepb

	mapCtx := &direct.MapContext{}
	observedState := VertexAIScheduleObservedState_FromProto(mapCtx, schedulepb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *ScheduleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAISchedule", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	req := &pb.CreateScheduleRequest{
		Parent:   a.id.ParentString(),
		Schedule: a.desiredpb,
	}

	created, err := a.gcpClient.CreateSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAISchedule %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created VertexAISchedule", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *ScheduleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAISchedule", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, updateMask, err := compareSchedule(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateScheduleRequest{
		Schedule:   a.desiredpb,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAISchedule %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated VertexAISchedule", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ScheduleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAISchedule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIScheduleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Schedule

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.Schedule)
	u.SetGroupVersionKind(krm.VertexAIScheduleGVK)
	return u, nil
}

func (a *ScheduleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Schedule) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIScheduleStatus{}
	status.ObservedState = VertexAIScheduleObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ScheduleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAISchedule", "name", a.id.String())

	req := &pb.DeleteScheduleRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAISchedule %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of VertexAISchedule %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareSchedule(ctx context.Context, actual, desired *pb.Schedule) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAIScheduleSpec_FromProto, VertexAIScheduleSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Schedule)

	populateDefaults := func(obj *pb.Schedule) {
		if obj.StartTime == nil && maskedActual.StartTime != nil {
			obj.StartTime = maskedActual.StartTime
		}
		if obj.EndTime == nil && maskedActual.EndTime != nil {
			obj.EndTime = maskedActual.EndTime
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
