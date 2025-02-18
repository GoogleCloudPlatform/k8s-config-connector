// Copyright 2025 Google LLC
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

package datastream

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/datastream/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	datastreampb "cloud.google.com/go/datastream/apiv1/datastreampb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DatastreamStreamGVK, NewStreamModel)
}

func NewStreamModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelStream{config: *config}, nil
}

var _ directbase.Model = &modelStream{}

type modelStream struct {
	config config.ControllerConfig
}

func (m *modelStream) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Stream client: %w", err)
	}
	return gcpClient, err
}

func (m *modelStream) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DatastreamStream{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewStreamIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get datastream GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &StreamAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelStream) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type StreamAdapter struct {
	id        *krm.StreamIdentity
	gcpClient *gcp.Client
	desired   *krm.DatastreamStream
	actual    *datastreampb.Stream
}

var _ directbase.Adapter = &StreamAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *StreamAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Stream", "name", a.id)

	req := &datastreampb.GetStreamRequest{Name: a.id.String()}
	streampb, err := a.gcpClient.GetStream(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Stream %q: %w", a.id, err)
	}

	a.actual = streampb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StreamAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Stream", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DatastreamStreamSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &datastreampb.CreateStreamRequest{
		Parent: a.id.Parent().String(),
		Stream: resource,
	}
	op, err := a.gcpClient.CreateStream(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Stream %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Stream %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Stream", "name", a.id)

	status := &krm.DatastreamStreamStatus{}
	status.ObservedState = DatastreamStreamObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StreamAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Stream", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := DatastreamStreamSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	{
		var err error
		pathsSet, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
		paths = pathsSet.UnsortedList()
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.DatastreamStreamStatus{}
		status.ObservedState = DatastreamStreamObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: paths}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &datastreampb.UpdateStreamRequest{
		UpdateMask: updateMask,
		Stream:     desiredPb,
	}
	op, err := a.gcpClient.UpdateStream(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Stream %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Stream %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Stream", "name", a.id)

	status := &krm.DatastreamStreamStatus{}
	status.ObservedState = DatastreamStreamObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *StreamAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatastreamStream{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DatastreamStreamSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DatastreamStreamGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *StreamAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Stream", "name", a.id)

	req := &datastreampb.DeleteStreamRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteStream(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Stream, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Stream %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Stream", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Stream %s: %w", a.id, err)
	}
	return true, nil
}
