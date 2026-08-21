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

package datalabelinginstruction

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/datalabeling/apiv1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DataLabelingInstructionGVK, NewModel)
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

	obj := &krm.DataLabelingInstruction{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.DataLabelingInstructionIdentity)
	if !ok {
		return nil, fmt.Errorf("expected DataLabelingInstructionIdentity, got %T", id)
	}

	return &Adapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DataLabelingInstructionIdentity
	gcpClient *gcp.Client
	desired   *krm.DataLabelingInstruction
	actual    *pb.Instruction
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	if a.id.Instruction == "" {
		return false, nil
	}

	req := &pb.GetInstructionRequest{
		Name: a.id.String(),
	}
	instruction, err := a.gcpClient.GetInstruction(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = instruction
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataLabelingInstruction", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	instructionProto := DataLabelingInstructionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateInstructionRequest{
		Parent:      a.id.ParentString(),
		Instruction: instructionProto,
	}

	op, err := a.gcpClient.CreateInstruction(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataLabelingInstruction: %w", err)
	}

	createdInstruction, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for DataLabelingInstruction creation: %w", err)
	}

	log.V(2).Info("successfully created DataLabelingInstruction", "name", a.id.String())

	return a.updateStatus(ctx, createOp, createdInstruction)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating/diffing DataLabelingInstruction", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredProto := DataLabelingInstructionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Because Instruction is immutable, let's verify if there is any diff in mutable fields (which is actually all spec fields).
	// Let's compare desired against actual
	diff := false
	if desiredProto.DisplayName != a.actual.DisplayName {
		diff = true
	}
	if desiredProto.Description != a.actual.Description {
		diff = true
	}
	if desiredProto.DataType != a.actual.DataType {
		diff = true
	}
	if !reflect.DeepEqual(desiredProto.CsvInstruction, a.actual.CsvInstruction) {
		diff = true
	}
	if !reflect.DeepEqual(desiredProto.PdfInstruction, a.actual.PdfInstruction) {
		diff = true
	}

	if diff {
		return fmt.Errorf("DataLabelingInstruction is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataLabelingInstruction", "name", a.id.String())

	req := &pb.DeleteInstructionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteInstruction(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting DataLabelingInstruction: %w", err)
	}

	log.V(2).Info("successfully deleted DataLabelingInstruction", "name", a.id.String())
	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Instruction) error {
	mapCtx := &direct.MapContext{}
	observedState := DataLabelingInstructionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.DataLabelingInstructionStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(latest.Name)

	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("DataLabelingInstruction actual state not found")
	}

	mapCtx := &direct.MapContext{}
	spec := DataLabelingInstructionSpec_FromProto(mapCtx, a.actual)
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
	u.SetName(a.id.Instruction)
	u.SetGroupVersionKind(krm.DataLabelingInstructionGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func DataLabelingInstructionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingInstructionSpec) *pb.Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Instruction{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	if in.DataType != nil {
		if val, ok := pb.DataType_value[*in.DataType]; ok {
			out.DataType = pb.DataType(val)
		} else {
			mapCtx.Errorf("invalid DataType: %q", *in.DataType)
		}
	}
	if in.CsvInstruction != nil {
		out.CsvInstruction = &pb.CsvInstruction{
			GcsFileUri: direct.ValueOf(in.CsvInstruction.GcsFileURI),
		}
	}
	if in.PdfInstruction != nil {
		out.PdfInstruction = &pb.PdfInstruction{
			GcsFileUri: direct.ValueOf(in.PdfInstruction.GcsFileURI),
		}
	}
	return out
}

func DataLabelingInstructionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instruction) *krm.DataLabelingInstructionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingInstructionSpec{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.Description = direct.LazyPtr(in.Description)

	dataTypeStr := in.DataType.String()
	out.DataType = &dataTypeStr

	if in.CsvInstruction != nil {
		out.CsvInstruction = &krm.InstructionCsvInstruction{
			GcsFileURI: direct.LazyPtr(in.CsvInstruction.GcsFileUri),
		}
	}
	if in.PdfInstruction != nil {
		out.PdfInstruction = &krm.InstructionPdfInstruction{
			GcsFileURI: direct.LazyPtr(in.PdfInstruction.GcsFileUri),
		}
	}
	return out
}

func DataLabelingInstructionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instruction) *krm.DataLabelingInstructionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataLabelingInstructionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.UpdateTime)
	out.BlockingResources = in.BlockingResources
	return out
}

func DataLabelingInstructionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataLabelingInstructionObservedState) *pb.Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Instruction{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.BlockingResources = in.BlockingResources
	return out
}
