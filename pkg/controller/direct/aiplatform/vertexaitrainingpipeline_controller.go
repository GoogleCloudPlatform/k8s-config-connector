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
	registry.RegisterModel(krm.VertexAITrainingPipelineGVK, NewTrainingPipelineModel)
}

func NewTrainingPipelineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &trainingPipelineModel{config: config}, nil
}

var _ directbase.Model = &trainingPipelineModel{}

type trainingPipelineModel struct {
	config *config.ControllerConfig
}

func (m *trainingPipelineModel) client(ctx context.Context) (*gcp.PipelineClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewPipelineClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PipelineClient client: %w", err)
	}
	return gcpClient, err
}

func (m *trainingPipelineModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAITrainingPipeline{}
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

	// Get PipelineClient client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.VertexAITrainingPipelineIdentity)
	if !ok {
		return nil, fmt.Errorf("expected VertexAITrainingPipelineIdentity, got %T", id)
	}

	mapCtx := &direct.MapContext{}
	desiredpb := VertexAITrainingPipelineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	// Post-process missing ParentModel field
	if obj.Spec.ModelRef != nil {
		desiredpb.ParentModel = obj.Spec.ModelRef.External
	}

	return &TrainingPipelineAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *trainingPipelineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type TrainingPipelineAdapter struct {
	id        *krm.VertexAITrainingPipelineIdentity
	gcpClient *gcp.PipelineClient
	desiredpb *pb.TrainingPipeline
	desired   *krm.VertexAITrainingPipeline
	actual    *pb.TrainingPipeline
}

var _ directbase.Adapter = &TrainingPipelineAdapter{}

func (a *TrainingPipelineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAITrainingPipeline", "name", a.id.String())

	req := &pb.GetTrainingPipelineRequest{
		Name: a.id.String(),
	}

	trainingPipelinepb, err := a.gcpClient.GetTrainingPipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAITrainingPipeline %q: %w", a.id.String(), err)
	}

	a.actual = trainingPipelinepb

	mapCtx := &direct.MapContext{}
	observedState := VertexAITrainingPipelineObservedState_FromProto(mapCtx, trainingPipelinepb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *TrainingPipelineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAITrainingPipeline", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	req := &pb.CreateTrainingPipelineRequest{
		Parent:           a.id.ParentString(),
		TrainingPipeline: a.desiredpb,
	}

	created, err := a.gcpClient.CreateTrainingPipeline(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAITrainingPipeline %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created VertexAITrainingPipeline", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *TrainingPipelineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAITrainingPipeline", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, _, err := compareTrainingPipeline(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("VertexAITrainingPipeline is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *TrainingPipelineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAITrainingPipeline{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAITrainingPipelineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Post-process missing ParentModel field
	if a.actual.GetParentModel() != "" {
		obj.Spec.ModelRef = &krm.AIPlatformModelRef{External: a.actual.GetParentModel()}
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.TrainingPipeline

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.TrainingPipeline)
	u.SetGroupVersionKind(krm.VertexAITrainingPipelineGVK)
	return u, nil
}

func (a *TrainingPipelineAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TrainingPipeline) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAITrainingPipelineStatus{}
	status.ObservedState = VertexAITrainingPipelineObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *TrainingPipelineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAITrainingPipeline", "name", a.id.String())

	req := &pb.DeleteTrainingPipelineRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteTrainingPipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAITrainingPipeline %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of VertexAITrainingPipeline %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareTrainingPipeline(ctx context.Context, actual, desired *pb.TrainingPipeline) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAITrainingPipelineSpec_FromProto, VertexAITrainingPipelineSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	// Since mapper.generated.go doesn't map ParentModel, make sure we copy it manually
	maskedActual.ParentModel = actual.ParentModel

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.TrainingPipeline) {
		// populate any defaults if necessary
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
