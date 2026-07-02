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
// proto.message: google.cloud.eventarc.v1.GoogleApiSource
// crd.type: EventarcGoogleAPISource
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"google.golang.org/protobuf/proto"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.EventarcGoogleAPISourceGVK, NewGoogleAPISourceModel)
}

func NewGoogleAPISourceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &googleAPISourceModel{config: *config}, nil
}

var _ directbase.Model = &googleAPISourceModel{}

type googleAPISourceModel struct {
	config config.ControllerConfig
}

func (m *googleAPISourceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcGoogleAPISource{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEventarcGoogleAPISourceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := EventarcGoogleAPISourceSpec_ToProto(mapCtx, &obj.Spec)
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

	return &googleAPISourceAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   desired,
		reader:    reader,
		obj:       obj,
	}, nil
}

func (m *googleAPISourceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type googleAPISourceAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcGoogleAPISourceIdentity
	desired   *pb.GoogleApiSource
	actual    *pb.GoogleApiSource
	reader    client.Reader
	obj       *krm.EventarcGoogleAPISource
}

var _ directbase.Adapter = &googleAPISourceAdapter{}

func (a *googleAPISourceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc GoogleApiSource", "name", a.id)

	req := &pb.GetGoogleApiSourceRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetGoogleApiSource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc GoogleApiSource %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *googleAPISourceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc GoogleApiSource", "name", a.id)

	req := &pb.CreateGoogleApiSourceRequest{
		Parent:            a.id.ParentString(),
		GoogleApiSource:   a.desired,
		GoogleApiSourceId: a.id.Google_api_source,
	}

	op, err := a.gcpClient.CreateGoogleApiSource(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc GoogleApiSource %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc GoogleApiSource creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc GoogleApiSource", "name", a.id)

	// Fetch fully populated resource
	latest, err := a.gcpClient.GetGoogleApiSource(ctx, &pb.GetGoogleApiSourceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created eventarc GoogleApiSource %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *googleAPISourceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc GoogleApiSource", "name", a.id)

	maskedActual, err := mappers.OnlySpecFields(a.actual, EventarcGoogleAPISourceSpec_FromProto, EventarcGoogleAPISourceSpec_ToProto)
	if err != nil {
		return err
	}

	maskedActual.Name = a.desired.Name
	maskedActual.Labels = a.actual.Labels
	maskedActual.Annotations = a.actual.Annotations

	clonedDesired := proto.Clone(a.desired).(*pb.GoogleApiSource)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateGoogleApiSourceRequest{
		GoogleApiSource: a.desired,
		UpdateMask:      updateMask,
	}

	op, err := a.gcpClient.UpdateGoogleApiSource(ctx, req)
	if err != nil {
		return fmt.Errorf("updating eventarc GoogleApiSource %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc GoogleApiSource update %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated eventarc GoogleApiSource", "name", a.id)

	// Fetch fully populated resource
	latest, err := a.gcpClient.GetGoogleApiSource(ctx, &pb.GetGoogleApiSourceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting updated eventarc GoogleApiSource %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *googleAPISourceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcGoogleAPISource{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcGoogleAPISourceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Google_api_source)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetNamespace(obj.Namespace)
	u.SetGroupVersionKind(krm.EventarcGoogleAPISourceGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *googleAPISourceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc GoogleApiSource", "name", a.id)

	req := &pb.DeleteGoogleApiSourceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteGoogleApiSource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc GoogleApiSource not found", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting eventarc GoogleApiSource %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc GoogleApiSource deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc GoogleApiSource", "name", a.id)
	return true, nil
}

func (a *googleAPISourceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.GoogleApiSource) error {
	status := &krm.EventarcGoogleAPISourceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = EventarcGoogleAPISourceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}
