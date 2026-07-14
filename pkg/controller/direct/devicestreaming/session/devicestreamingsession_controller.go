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

package session

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/devicestreaming/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	mapper "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/devicestreaming"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/devicestreaming/apiv1"
	pb "cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DeviceStreamingSessionGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.DirectAccessClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDirectAccessRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building devicestreaming client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DeviceStreamingSession{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.DeviceStreamingSessionIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := mapper.DeviceStreamingSessionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &DeviceStreamingSessionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type DeviceStreamingSessionAdapter struct {
	id        *krm.DeviceStreamingSessionIdentity
	gcpClient *gcp.DirectAccessClient
	desired   *pb.DeviceSession
	actual    *pb.DeviceSession
}

var _ directbase.Adapter = &DeviceStreamingSessionAdapter{}

func (a *DeviceStreamingSessionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeviceStreamingSession", "name", a.id)

	req := &pb.GetDeviceSessionRequest{Name: a.id.String()}
	session, err := a.gcpClient.GetDeviceSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeviceStreamingSession %q: %w", a.id, err)
	}

	a.actual = session
	return true, nil
}

func (a *DeviceStreamingSessionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeviceStreamingSession", "name", a.id)

	req := &pb.CreateDeviceSessionRequest{
		Parent:          a.id.ParentString(),
		DeviceSessionId: a.id.DeviceSession,
		DeviceSession:   a.desired,
	}

	created, err := a.gcpClient.CreateDeviceSession(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeviceStreamingSession %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeviceStreamingSession", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *DeviceStreamingSessionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeviceStreamingSession", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired

	maskedActualSpec := mapper.DeviceStreamingSessionSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := mapper.DeviceStreamingSessionSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(desired).(*pb.DeviceSession)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateDeviceSessionRequest{
		DeviceSession: clonedDesired,
		UpdateMask:    updateMask,
	}

	updated, err := a.gcpClient.UpdateDeviceSession(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DeviceStreamingSession %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *DeviceStreamingSessionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeviceStreamingSession", "name", a.id)

	req := &pb.CancelDeviceSessionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.CancelDeviceSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DeviceStreamingSession %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DeviceStreamingSession", "name", a.id)
	return true, nil
}

func (a *DeviceStreamingSessionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("exporting DeviceStreamingSession", "name", a.id)

	mapCtx := &direct.MapContext{}
	spec := mapper.DeviceStreamingSessionSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.DeviceStreamingSession{}
	obj.Spec = *spec
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting to unstructured: %w", err)
	}

	uObj := &unstructured.Unstructured{Object: u}
	uObj.SetGroupVersionKind(krm.DeviceStreamingSessionGVK)
	return uObj, nil
}

func (a *DeviceStreamingSessionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DeviceSession) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DeviceStreamingSessionStatus{}
	status.ObservedState = mapper.DeviceStreamingSessionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
