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

package tensorboard

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAITensorBoardGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*gcp.TensorboardClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewTensorboardRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Tensorboard client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAITensorBoard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	vertexaiID := id.(*krm.VertexAITensorBoardIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, vertexaiID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := VertexAITensorBoardSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.GCPLabels(obj)

	return &Adapter{
		id:        vertexaiID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAITensorBoardIdentity
	gcpClient *gcp.TensorboardClient
	desired   *pb.Tensorboard
	actual    *pb.Tensorboard
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAITensorBoard", "name", a.id)

	req := &pb.GetTensorboardRequest{
		Name: a.id.String(),
	}
	resp, err := a.gcpClient.GetTensorboard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) || direct.IsBadRequest(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAITensorBoard %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAITensorBoard", "name", a.id)

	req := &pb.CreateTensorboardRequest{
		Parent:      a.id.ParentString(),
		Tensorboard: a.desired,
	}
	op, err := a.gcpClient.CreateTensorboard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAITensorBoard %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting creation VertexAITensorBoard %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created VertexAITensorBoard", "name", a.id)

	if err := a.id.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created VertexAITensorBoard name: %w", err)
	}

	actual, err := a.gcpClient.GetTensorboard(ctx, &pb.GetTensorboardRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting VertexAITensorBoard after creation %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAITensorBoard", "name", a.id)

	diffs, updateMask, err := compareResource(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	clonedDesired := proto.Clone(a.desired).(*pb.Tensorboard)
	clonedDesired.Name = a.id.String()

	req := &pb.UpdateTensorboardRequest{
		Tensorboard: clonedDesired,
		UpdateMask:  updateMask,
	}
	op, err := a.gcpClient.UpdateTensorboard(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAITensorBoard %s: %w", a.id, err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting update VertexAITensorBoard %s: %w", a.id, err)
	}

	actual, err := a.gcpClient.GetTensorboard(ctx, &pb.GetTensorboardRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting VertexAITensorBoard after update %s: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAITensorBoard", "name", a.id)

	req := &pb.DeleteTensorboardRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTensorboard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VertexAITensorBoard, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAITensorBoard %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted VertexAITensorBoard, waiting for operation to complete", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete VertexAITensorBoard %s: %w", a.id, err)
	}
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAITensorBoard{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAITensorBoardSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	obj.Spec.Region = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Tensorboard)
	u.SetGroupVersionKind(krm.VertexAITensorBoardGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Tensorboard) error {
	mapCtx := &direct.MapContext{}
	observedState := VertexAITensorBoardObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.VertexAITensorBoardStatus{
		ObservedState: observedState,
		ExternalRef:   direct.LazyPtr(a.id.String()),
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareResource(ctx context.Context, actual, desired *pb.Tensorboard) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual := &pb.Tensorboard{
		DisplayName:    actual.DisplayName,
		Description:    actual.Description,
		EncryptionSpec: actual.EncryptionSpec,
		IsDefault:      actual.IsDefault,
		Labels:         actual.Labels,
	}

	clonedDesired := proto.Clone(desired).(*pb.Tensorboard)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
