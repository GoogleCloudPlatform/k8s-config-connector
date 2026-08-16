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

package datalabelingannotationspecset

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/datalabeling/apiv1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataLabelingAnnotationSpecSetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datalabeling client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.DataLabelingAnnotationSpecSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.DataLabelingAnnotationSpecSetIdentity)
	if !ok {
		return nil, fmt.Errorf("expected DataLabelingAnnotationSpecSetIdentity, got %T", id)
	}

	mapCtx := &direct.MapContext{}
	desiredProto := DataLabelingAnnotationSpecSetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DataLabelingAnnotationSpecSetIdentity
	gcpClient *gcp.Client
	desired   *pb.AnnotationSpecSet
	actual    *pb.AnnotationSpecSet
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	if a.id.Annotation_spec_set == "" {
		return false, nil
	}

	req := &pb.GetAnnotationSpecSetRequest{
		Name: a.id.String(),
	}
	annotationSpecSet, err := a.gcpClient.GetAnnotationSpecSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = annotationSpecSet
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataLabelingAnnotationSpecSet", "name", a.id.String())

	req := &pb.CreateAnnotationSpecSetRequest{
		Parent:            a.id.ParentString(),
		AnnotationSpecSet: a.desired,
	}

	createdAnnotationSpecSet, err := a.gcpClient.CreateAnnotationSpecSet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataLabelingAnnotationSpecSet: %w", err)
	}

	log.V(2).Info("successfully created DataLabelingAnnotationSpecSet", "name", a.id.String())

	return a.updateStatus(ctx, createOp, createdAnnotationSpecSet)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating/diffing DataLabelingAnnotationSpecSet", "name", a.id.String())

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) > 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("DataLabelingAnnotationSpecSet is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataLabelingAnnotationSpecSet", "name", a.id.String())

	req := &pb.DeleteAnnotationSpecSetRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAnnotationSpecSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting DataLabelingAnnotationSpecSet: %w", err)
	}

	log.V(2).Info("successfully deleted DataLabelingAnnotationSpecSet", "name", a.id.String())
	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AnnotationSpecSet) error {
	mapCtx := &direct.MapContext{}
	observedState := DataLabelingAnnotationSpecSetObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.DataLabelingAnnotationSpecSetStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(latest.Name)

	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("DataLabelingAnnotationSpecSet actual state not found")
	}

	mapCtx := &direct.MapContext{}
	spec := DataLabelingAnnotationSpecSetSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("converting spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Annotation_spec_set)
	u.SetGroupVersionKind(krm.DataLabelingAnnotationSpecSetGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func DataLabelingAnnotationSpecSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingAnnotationSpecSetSpec) *pb.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpecSet{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	if in.AnnotationSpecs != nil {
		out.AnnotationSpecs = make([]*pb.AnnotationSpec, len(in.AnnotationSpecs))
		for i, v := range in.AnnotationSpecs {
			out.AnnotationSpecs[i] = &pb.AnnotationSpec{
				DisplayName: direct.ValueOf(v.DisplayName),
				Description: direct.ValueOf(v.Description),
			}
		}
	}
	return out
}

func DataLabelingAnnotationSpecSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpecSet) *krm.DataLabelingAnnotationSpecSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingAnnotationSpecSetSpec{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.Description = direct.LazyPtr(in.Description)
	if in.AnnotationSpecs != nil {
		out.AnnotationSpecs = make([]krm.AnnotationSpec, len(in.AnnotationSpecs))
		for i, v := range in.AnnotationSpecs {
			out.AnnotationSpecs[i] = krm.AnnotationSpec{
				DisplayName: direct.LazyPtr(v.DisplayName),
				Description: direct.LazyPtr(v.Description),
			}
		}
	}
	return out
}

func DataLabelingAnnotationSpecSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpecSet) *krm.DataLabelingAnnotationSpecSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingAnnotationSpecSetObservedState{}
	out.BlockingResources = in.BlockingResources
	return out
}

func DataLabelingAnnotationSpecSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingAnnotationSpecSetObservedState) *pb.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpecSet{}
	out.BlockingResources = in.BlockingResources
	return out
}
