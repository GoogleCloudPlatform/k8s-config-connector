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

package monitoring

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/monitoringpb"
	"google.golang.org/api/option"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringNotificationChannelGVK, NewNotificationChannelModel)
}

func NewNotificationChannelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNotificationChannel{config: *config}, nil
}

var _ directbase.Model = &modelNotificationChannel{}

type modelNotificationChannel struct {
	config config.ControllerConfig
}

func (m *modelNotificationChannel) client(ctx context.Context) (*monitoringpb.NotificationChannelServiceV3Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := monitoringpb.NewNotificationChannelServiceV3Client(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NotificationChannel client: %w", err)
	}
	return gcpClient, err
}

func (m *modelNotificationChannel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MonitoringNotificationChannel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNotificationChannelIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get monitoring GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &NotificationChannelAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelNotificationChannel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NotificationChannelAdapter struct {
	id        *krm.NotificationChannelIdentity
	gcpClient *monitoringpb.NotificationChannelServiceV3Client
	desired   *krm.MonitoringNotificationChannel
	actual    *monitoringpb.NotificationChannel
}

var _ directbase.Adapter = &NotificationChannelAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *NotificationChannelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NotificationChannel", "name", a.id)

	req := &monitoringpb.GetNotificationChannelRequest{Name: a.id.String()}
	notificationchannelpb, err := a.gcpClient.GetNotificationChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NotificationChannel %q: %w", a.id, err)
	}

	a.actual = notificationchannelpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotificationChannelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NotificationChannel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MonitoringNotificationChannelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &monitoringpb.CreateNotificationChannelRequest{
		Name:              a.id.Parent().String(),
		NotificationChannel: resource,
	}
	op, err := a.gcpClient.CreateNotificationChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NotificationChannel %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NotificationChannel %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NotificationChannel", "name", a.id)

	status := &krm.MonitoringNotificationChannelStatus{}
	status.ObservedState = MonitoringNotificationChannelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotificationChannelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NotificationChannel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MonitoringNotificationChannelSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths sets.Set[string]
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.MonitoringNotificationChannelStatus{}
		status.ObservedState = MonitoringNotificationChannelObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &monitoringpb.UpdateNotificationChannelRequest{
		UpdateMask:          updateMask,
		NotificationChannel: desiredPb,
	}
	req.NotificationChannel.Name = a.id.String()

	op, err := a.gcpClient.UpdateNotificationChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("updating NotificationChannel %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NotificationChannel %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated NotificationChannel", "name", a.id)

	status := &krm.MonitoringNotificationChannelStatus{}
	status.ObservedState = MonitoringNotificationChannelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NotificationChannelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringNotificationChannel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringNotificationChannelSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.MonitoringNotificationChannelGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NotificationChannelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NotificationChannel", "name", a.id)

	req := &monitoringpb.DeleteNotificationChannelRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNotificationChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent NotificationChannel, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting NotificationChannel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NotificationChannel", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NotificationChannel %s: %w", a.id, err)
	}
	return true, nil
}
