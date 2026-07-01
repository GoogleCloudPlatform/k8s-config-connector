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

// +tool:controller
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.ChannelConnection
// crd.type: EventarcChannelConnection
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.EventarcChannelConnectionGVK, NewChannelConnectionModel)
}

func NewChannelConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &channelConnectionModel{config: *config}, nil
}

var _ directbase.Model = &channelConnectionModel{}

type channelConnectionModel struct {
	config config.ControllerConfig
}

func (m *channelConnectionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcChannelConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewEventarcChannelConnectionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get eventarc GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := EventarcChannelConnectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &channelConnectionAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *channelConnectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.EventarcChannelConnectionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	return &channelConnectionAdapter{
		gcpClient: eventarcClient,
		id:        id,
	}, nil
}

type channelConnectionAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcChannelConnectionIdentity
	desired   *pb.ChannelConnection
	actual    *pb.ChannelConnection
	reader    client.Reader
}

var _ directbase.Adapter = &channelConnectionAdapter{}

func (a *channelConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc channel connection", "name", a.id)

	req := &pb.GetChannelConnectionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetChannelConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc channel connection %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *channelConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc channel connection", "name", a.id)

	req := &pb.CreateChannelConnectionRequest{
		Parent:              a.id.ParentString(),
		ChannelConnection:   a.desired,
		ChannelConnectionId: a.id.ChannelConnection,
	}

	op, err := a.gcpClient.CreateChannelConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc channel connection %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc channel connection creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc channel connection", "name", a.id)

	// Fetch fully-populated resource to populate status completely
	latest, err := a.gcpClient.GetChannelConnection(ctx, &pb.GetChannelConnectionRequest{Name: created.Name})
	if err != nil {
		return fmt.Errorf("getting eventarc channel connection %q after creation: %w", created.Name, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *channelConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc channel connection", "name", a.id)

	// ChannelConnection is immutable. Let's compare spec fields to report errors for updates.
	diffs := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if a.desired.Channel != a.actual.Channel {
		diffs.AddField("spec.channelRef", a.desired.Channel, a.actual.Channel)
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("EventarcChannelConnection is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *channelConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcChannelConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcChannelConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ChannelConnection)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetNamespace(obj.Namespace)
	u.SetGroupVersionKind(krm.EventarcChannelConnectionGVK)
	u.Object = uObj

	return u, nil
}

func (a *channelConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc channel connection", "name", a.id)

	req := &pb.DeleteChannelConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteChannelConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc channel connection already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting eventarc channel connection %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc channel connection deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc channel connection", "name", a.id)
	return true, nil
}

func (a *channelConnectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ChannelConnection) error {
	mapCtx := &direct.MapContext{}
	status := &krm.EventarcChannelConnectionStatus{}
	status.ObservedState = EventarcChannelConnectionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}
