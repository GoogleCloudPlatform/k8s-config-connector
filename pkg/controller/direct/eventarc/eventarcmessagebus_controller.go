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
// proto.message: google.cloud.eventarc.v1.MessageBus
// crd.type: EventarcMessageBus
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
	registry.RegisterModel(krm.EventarcMessageBusGVK, NewMessageBusModel)
}

func NewMessageBusModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &messageBusModel{config: *config}, nil
}

var _ directbase.Model = &messageBusModel{}

type messageBusModel struct {
	config config.ControllerConfig
}

func (m *messageBusModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcMessageBus{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEventarcMessageBusIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := EventarcMessageBusSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	// Get eventarc GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	return &messageBusAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   desired,
		reader:    reader,
		obj:       obj,
	}, nil
}

func (m *messageBusModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type messageBusAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcMessageBusIdentity
	desired   *pb.MessageBus
	actual    *pb.MessageBus
	reader    client.Reader
	obj       *krm.EventarcMessageBus
}

var _ directbase.Adapter = &messageBusAdapter{}

func (a *messageBusAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc MessageBus", "name", a.id)

	req := &pb.GetMessageBusRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMessageBus(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc MessageBus %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *messageBusAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc MessageBus", "name", a.id)

	req := &pb.CreateMessageBusRequest{
		Parent:       a.id.ParentString(),
		MessageBus:   a.desired,
		MessageBusId: a.id.MessageBus,
	}

	op, err := a.gcpClient.CreateMessageBus(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc MessageBus %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc MessageBus creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc MessageBus", "name", a.id)

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetMessageBus(ctx, &pb.GetMessageBusRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created eventarc MessageBus %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *messageBusAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc MessageBus", "name", a.id)

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.obj.Spec,
		a.actual,
		EventarcMessageBusSpec_FromProto,
		EventarcMessageBusSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateMessageBusRequest{
		MessageBus: a.desired,
		UpdateMask: updateMask,
	}

	op, err := a.gcpClient.UpdateMessageBus(ctx, req)
	if err != nil {
		return fmt.Errorf("updating eventarc MessageBus %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc MessageBus update %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated eventarc MessageBus", "name", a.id)

	// Fetch fully-populated resource after update
	latest, err := a.gcpClient.GetMessageBus(ctx, &pb.GetMessageBusRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting updated eventarc MessageBus %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *messageBusAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcMessageBus{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcMessageBusSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.MessageBus)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.MessageBus)
	u.SetNamespace(obj.Namespace)
	u.SetGroupVersionKind(krm.EventarcMessageBusGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *messageBusAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc MessageBus", "name", a.id)

	req := &pb.DeleteMessageBusRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMessageBus(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc MessageBus not found", "name", a.id)
			return false, nil
		}
		return false, fmt.Errorf("deleting eventarc MessageBus %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc MessageBus deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc MessageBus", "name", a.id)
	return true, nil
}

func (a *messageBusAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MessageBus) error {
	mapCtx := &direct.MapContext{}
	status := &krm.EventarcMessageBusStatus{}
	status.ObservedState = EventarcMessageBusObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}
