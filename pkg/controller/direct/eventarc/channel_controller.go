// Copyright 2024 Google LLC
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
// proto.message: google.cloud.eventarc.v1.Channel
// crd.type: EventarcChannel
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.EventarcChannelGVK, NewChannelModel)
}

func NewChannelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &channelModel{config: *config}, nil
}

var _ directbase.Model = &channelModel{}

type channelModel struct {
	config config.ControllerConfig
}

func (m *channelModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.EventarcChannel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewChannelIdentity(ctx, reader, obj)
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
	return &channelAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *channelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type channelAdapter struct {
	gcpClient *gcp.Client
	id        *krm.ChannelIdentity
	desired   *krm.EventarcChannel
	actual    *pb.Channel
	reader    client.Reader
}

var _ directbase.Adapter = &channelAdapter{}

func (a *channelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc channel", "name", a.id)

	req := &pb.GetChannelRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc channel %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *channelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc channel", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := EventarcChannelSpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateChannelRequest{
		Parent:    a.id.Parent().String(),
		Channel:   resource,
		ChannelId: a.id.ID(),
	}

	op, err := a.gcpClient.CreateChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc channel %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc channel creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc channel", "name", a.id)

	status := &krm.EventarcChannelStatus{}
	status.ObservedState = EventarcChannelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *channelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc channel", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := EventarcChannelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if resource.Provider != a.actual.Provider {
		paths = append(paths, "provider")
	}
	if resource.CryptoKeyName != a.actual.CryptoKeyName {
		paths = append(paths, "crypto_key_name")
	}

	var updated *pb.Channel
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateChannelRequest{
			Channel:    resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}

		op, err := a.gcpClient.UpdateChannel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating eventarc channel %s: %w", a.id.String(), err)
		}

		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for eventarc channel update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated eventarc channel", "name", a.id)
	}

	status := &krm.EventarcChannelStatus{}
	status.ObservedState = EventarcChannelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *channelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcChannel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcChannelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetNamespace(obj.Namespace) // This is required KCC controller code convention
	u.SetGroupVersionKind(krm.EventarcChannelGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *channelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc channel", "name", a.id)

	req := &pb.DeleteChannelRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc channel not found", "name", a.id)
			return false, nil // Resource is gone, consider the delete successful.
		}
		if apierrors.IsNotFound(err) { // Handle potential K8S not found errors if applicable
			log.V(2).Info("eventarc channel resource not found in K8S", "name", a.id)
			return false, nil
		}
		return false, fmt.Errorf("deleting eventarc channel %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc channel deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc channel", "name", a.id)
	return true, nil
}

func (a *channelAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.Provider != nil {
		providerRef, err := obj.Spec.Provider.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.Provider.External = providerRef
	}
	if obj.Spec.KmsKeyRef != nil {
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.KmsKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.KmsKeyRef = kmsKeyRef
	}
	return nil
}
