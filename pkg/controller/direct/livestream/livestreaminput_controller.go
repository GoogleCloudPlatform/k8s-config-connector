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

package livestream

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/video/livestream/apiv1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/livestream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	registry.RegisterModel(krm.LiveStreamInputGVK, NewInputModel)
}

func NewInputModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &inputModel{config: *config}, nil
}

var _ directbase.Model = &inputModel{}

type inputModel struct {
	config config.ControllerConfig
}

func (m *inputModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building livestream client: %w", err)
	}
	return gcpClient, nil
}

func (m *inputModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LiveStreamInput{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.LiveStreamInputIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := LiveStreamInputSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &inputAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		model:     m,
	}, nil
}

func (m *inputModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.LiveStreamInputIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &inputAdapter{
		id:        id,
		gcpClient: gcpClient,
		model:     m,
	}, nil
}

type inputAdapter struct {
	id        *krm.LiveStreamInputIdentity
	gcpClient *gcp.Client
	desired   *pb.Input
	actual    *pb.Input
	model     *inputModel
}

var _ directbase.Adapter = &inputAdapter{}

func (a *inputAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding LiveStreamInput", "id", a.id)

	req := &pb.GetInputRequest{
		Name: a.id.String(),
	}
	input, err := a.gcpClient.GetInput(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LiveStreamInput %s: %w", a.id.String(), err)
	}

	a.actual = input
	return true, nil
}

func (a *inputAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LiveStreamInput", "id", a.id)

	req := &pb.CreateInputRequest{
		Parent:  a.id.ParentString(),
		InputId: a.id.Input,
		Input:   a.desired,
	}
	op, err := a.gcpClient.CreateInput(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LiveStreamInput %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for LiveStreamInput %s creation: %w", a.id.String(), err)
	}

	// Fetch fully-populated resource after creation
	refetched, err := a.gcpClient.GetInput(ctx, &pb.GetInputRequest{Name: a.id.String()})
	if err != nil {
		refetched = created
	}
	a.actual = refetched

	return a.updateStatus(ctx, createOp, refetched)
}

func (a *inputAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LiveStreamInput", "id", a.id)

	diffs, updateMask, err := a.compareInput(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*pb.Input)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateInputRequest{
			Input:      desiredCopy,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateInput(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LiveStreamInput %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting LiveStreamInput %s update: %w", a.id.String(), err)
		}

		// Fetch fully-populated resource after update
		refetched, err := a.gcpClient.GetInput(ctx, &pb.GetInputRequest{Name: a.id.String()})
		if err != nil {
			refetched = latest
		}
		latest = refetched
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *inputAdapter) compareInput(ctx context.Context, actual, desired *pb.Input) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LiveStreamInputSpec_FromProto, LiveStreamInputSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	if desired.Tier == pb.Input_TIER_UNSPECIFIED && maskedActual.Tier == pb.Input_HD {
		desired.Tier = pb.Input_HD
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *inputAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Input) error {
	mapCtx := &direct.MapContext{}
	status := krm.LiveStreamInputStatus{}
	status.ObservedState = LiveStreamInputObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *inputAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LiveStreamInput{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LiveStreamInputSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Input)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Input)
	u.SetGroupVersionKind(krm.LiveStreamInputGVK)

	return u, nil
}

func (a *inputAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LiveStreamInput", "id", a.id)

	req := &pb.DeleteInputRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInput(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LiveStreamInput, assuming it was already deleted", "id", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LiveStreamInput %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete LiveStreamInput %s: %w", a.id.String(), err)
	}
	return true, nil
}
