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
	"errors"
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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIStudyGVK, NewStudyModel)
}

func NewStudyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &studyModel{config: config}, nil
}

var _ directbase.Model = &studyModel{}

type studyModel struct {
	config *config.ControllerConfig
}

func (m *studyModel) client(ctx context.Context) (*gcp.VizierClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewVizierClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building VizierClient client: %w", err)
	}
	return gcpClient, err
}

func (m *studyModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAIStudy{}
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

	// Get VizierClient client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.VertexAIStudyIdentity)
	if !ok {
		return nil, fmt.Errorf("expected VertexAIStudyIdentity, got %T", id)
	}

	mapCtx := &direct.MapContext{}
	desiredpb := VertexAIStudySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &StudyAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *studyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type StudyAdapter struct {
	id        *krm.VertexAIStudyIdentity
	gcpClient *gcp.VizierClient
	desiredpb *pb.Study
	desired   *krm.VertexAIStudy
	actual    *pb.Study
}

var _ directbase.Adapter = &StudyAdapter{}

func (a *StudyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	// 1. If we have status.externalRef, we use that directly.
	externalRef := ""
	if a.desired.Status.ExternalRef != nil {
		externalRef = *a.desired.Status.ExternalRef
	}

	if externalRef == "" {
		// Try to find the study by listing and matching DisplayName
		displayName := ""
		if a.desired.Spec.DisplayName != nil {
			displayName = *a.desired.Spec.DisplayName
		}
		if displayName != "" {
			log.V(2).Info("searching for existing Study by displayName", "displayName", displayName)
			req := &pb.ListStudiesRequest{
				Parent: a.id.ParentString(),
			}
			it := a.gcpClient.ListStudies(ctx, req)
			for {
				study, err := it.Next()
				if errors.Is(err, iterator.Done) {
					break
				}
				if err != nil {
					return false, fmt.Errorf("listing studies in parent %q: %w", a.id.ParentString(), err)
				}
				if study.GetDisplayName() == displayName {
					log.V(2).Info("found matching Study by displayName", "displayName", displayName, "name", study.GetName())
					externalRef = study.GetName()
					break
				}
			}
		}
	}

	if externalRef == "" {
		return false, nil
	}

	log.V(2).Info("getting VertexAIStudy", "name", externalRef)

	req := &pb.GetStudyRequest{
		Name: externalRef,
	}

	studypb, err := a.gcpClient.GetStudy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIStudy %q: %w", externalRef, err)
	}

	a.actual = studypb

	mapCtx := &direct.MapContext{}
	observedState := VertexAIStudyObservedState_FromProto(mapCtx, studypb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(externalRef)
	return true, nil
}

func (a *StudyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIStudy", "parent", a.id.ParentString())

	req := &pb.CreateStudyRequest{
		Parent: a.id.ParentString(),
		Study:  a.desiredpb,
	}

	created, err := a.gcpClient.CreateStudy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIStudy: %w", err)
	}

	log.V(2).Info("successfully created VertexAIStudy", "name", created.GetName())

	return a.updateStatus(ctx, createOp, created)
}

func (a *StudyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIStudy", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, _, err := compareStudy(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("VertexAIStudy is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *StudyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIStudy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIStudySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Study

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.Study)
	u.SetGroupVersionKind(krm.VertexAIStudyGVK)
	return u, nil
}

func (a *StudyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Study) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIStudyStatus{}
	status.ObservedState = VertexAIStudyObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *StudyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	externalRef := ""
	if a.desired.Status.ExternalRef != nil {
		externalRef = *a.desired.Status.ExternalRef
	}

	if externalRef == "" {
		// It was never created or found
		return true, nil
	}

	log.V(2).Info("deleting VertexAIStudy", "name", externalRef)

	req := &pb.DeleteStudyRequest{
		Name: externalRef,
	}

	err := a.gcpClient.DeleteStudy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAIStudy %q: %w", externalRef, err)
	}

	return true, nil
}

func compareStudy(ctx context.Context, actual, desired *pb.Study) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAIStudySpec_FromProto, VertexAIStudySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Study) {
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
