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
// proto.message: google.cloud.eventarc.v1.GoogleChannelConfig
// crd.type: EventarcGoogleChannelConfig
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
	registry.RegisterModel(krm.EventarcGoogleChannelConfigGVK, NewGoogleChannelConfigModel)
}

func NewGoogleChannelConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &googleChannelConfigModel{config: *config}, nil
}

var _ directbase.Model = &googleChannelConfigModel{}

type googleChannelConfigModel struct {
	config config.ControllerConfig
}

func (m *googleChannelConfigModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.EventarcGoogleChannelConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewGoogleChannelConfigIdentity(ctx, reader, obj)
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
	return &googleChannelConfigAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *googleChannelConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type googleChannelConfigAdapter struct {
	gcpClient *gcp.Client
	id        *krm.GoogleChannelConfigIdentity
	desired   *krm.EventarcGoogleChannelConfig
	actual    *pb.GoogleChannelConfig
	reader    client.Reader
}

var _ directbase.Adapter = &googleChannelConfigAdapter{}

func (a *googleChannelConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc googlechannelconfig", "name", a.id)

	req := &pb.GetGoogleChannelConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetGoogleChannelConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc googlechannelconfig %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *googleChannelConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// GoogleChannelConfig is a singleton resource that cannot be created.
	// It can only be fetched and updated.
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc googlechannelconfig is a no-op", "name", a.id)
	return nil
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *googleChannelConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc googlechannelconfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := EventarcGoogleChannelConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	paths = append(paths, "crypto_key_name")

	var updated *pb.GoogleChannelConfig
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateGoogleChannelConfigRequest{
			GoogleChannelConfig: resource,
			UpdateMask:          &fieldmaskpb.FieldMask{Paths: paths},
		}
		var err error
		updated, err = a.gcpClient.UpdateGoogleChannelConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating eventarc googlechannelconfig %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated eventarc googlechannelconfig", "name", a.id)
	}

	status := &krm.EventarcGoogleChannelConfigStatus{}
	status.ObservedState = EventarcGoogleChannelConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *googleChannelConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcGoogleChannelConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcGoogleChannelConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Use the static name for the singleton resource.
	u.SetName("googlechannelconfig")
	u.SetGroupVersionKind(krm.EventarcGoogleChannelConfigGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *googleChannelConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc googlechannelconfig is a no-op", "name", a.id)
	// GoogleChannelConfig is a singleton resource that cannot be deleted.
	return true, nil
}

func (a *googleChannelConfigAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.CryptoKeyRef != nil {
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.CryptoKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.CryptoKeyRef = kmsKeyRef
	}
	return nil
}
