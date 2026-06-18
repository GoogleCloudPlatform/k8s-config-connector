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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/devicestreaming/apiv1"
	devicestreamingpb "cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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
		return nil, fmt.Errorf("building DeviceSession client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DeviceStreamingSession{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identityVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identityVal.(*krm.DeviceStreamingSessionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identityVal)
	}

	// Get devicestreaming GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeviceSessionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DeviceSessionAdapter struct {
	id        *krm.DeviceStreamingSessionIdentity
	gcpClient *gcp.DirectAccessClient
	desired   *krm.DeviceStreamingSession
	actual    *devicestreamingpb.DeviceSession
}

var _ directbase.Adapter = &DeviceSessionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DeviceSessionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeviceSession", "name", a.id)

	req := &devicestreamingpb.GetDeviceSessionRequest{Name: a.id.String()}
	devicesessionpb, err := a.gcpClient.GetDeviceSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeviceSession %q: %w", a.id, err)
	}

	a.actual = devicesessionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeviceSessionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeviceSession", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeviceStreamingSessionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := "projects/" + a.id.Project
	req := &devicestreamingpb.CreateDeviceSessionRequest{
		Parent:          parent,
		DeviceSession:   resource,
		DeviceSessionId: a.id.DeviceSession,
	}
	created, err := a.gcpClient.CreateDeviceSession(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeviceSession %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeviceSession", "name", a.id)

	status := &krm.DeviceStreamingSessionStatus{}
	status.ObservedState = DeviceStreamingSessionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeviceSessionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeviceSession", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := DeviceStreamingSessionSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)

		desiredPb.Name = a.id.String()
		req := &devicestreamingpb.UpdateDeviceSessionRequest{
			DeviceSession: desiredPb,
			UpdateMask:    updateMask,
		}
		res, err := a.gcpClient.UpdateDeviceSession(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DeviceSession %s: %w", a.id, err)
		}
		updated = res
		log.V(2).Info("successfully updated DeviceSession", "name", a.id)
	}

	status := &krm.DeviceStreamingSessionStatus{}
	status.ObservedState = DeviceStreamingSessionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DeviceSessionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DeviceStreamingSession{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DeviceStreamingSessionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.DeviceSession)
	u.SetGroupVersionKind(krm.DeviceStreamingSessionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DeviceSessionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeviceSession", "name", a.id)

	req := &devicestreamingpb.CancelDeviceSessionRequest{Name: a.id.String()}
	err := a.gcpClient.CancelDeviceSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping cancel for non-existent DeviceSession, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("cancelling DeviceSession %s: %w", a.id, err)
	}
	log.V(2).Info("successfully cancelled DeviceSession", "name", a.id)
	return true, nil
}
