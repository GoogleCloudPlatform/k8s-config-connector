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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	api "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MonitoringNotificationChannelGVK, NewMonitoringNotificationChannelModel)
}

func NewMonitoringNotificationChannelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &MonitoringNotificationChannelModel{config: config}, nil
}

var _ directbase.Model = &MonitoringNotificationChannelModel{}

type MonitoringNotificationChannelModel struct {
	config *config.ControllerConfig
}

func (m *MonitoringNotificationChannelModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	notificationChannelsClient, err := newNotificationChannelsClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringNotificationChannel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	var id *krm.MonitoringNotificationChannelIdentity
	if obj.Spec.ResourceID != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.MonitoringNotificationChannelIdentity)
	}

	var desired *pb.NotificationChannel
	{
		mapCtx := &direct.MapContext{}
		desired = MonitoringNotificationChannelSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &MonitoringNotificationChannelAdapter{
		id:                         id,
		notificationChannelsClient: notificationChannelsClient,
		desired:                    desired,
	}, nil
}

func (m *MonitoringNotificationChannelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	id := &krm.MonitoringNotificationChannelIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	notificationChannelsClient, err := newNotificationChannelsClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &MonitoringNotificationChannelAdapter{
		id:                         id,
		notificationChannelsClient: notificationChannelsClient,
	}, nil
}

type MonitoringNotificationChannelAdapter struct {
	id                         *krm.MonitoringNotificationChannelIdentity
	notificationChannelsClient *api.NotificationChannelClient
	desired                    *pb.NotificationChannel
	actual                     *pb.NotificationChannel
}

var _ directbase.Adapter = &MonitoringNotificationChannelAdapter{}

func (a *MonitoringNotificationChannelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting MonitoringNotificationChannel", "name", fqn)

	req := &pb.GetNotificationChannelRequest{Name: fqn}
	actual, err := a.notificationChannelsClient.GetNotificationChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MonitoringNotificationChannel %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *MonitoringNotificationChannelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// There no FQN until after creation (server generated id).

	log := klog.FromContext(ctx)
	log.V(0).Info("creating MonitoringNotificationChannel")

	req := &pb.CreateNotificationChannelRequest{
		NotificationChannel: direct.ProtoClone(a.desired),
	}

	created, err := a.notificationChannelsClient.CreateNotificationChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MonitoringNotificationChannel: %w", err)
	}
	log.V(0).Info("created MonitoringNotificationChannel", "name", created.GetName())

	// Set resourceID
	resourceID := strings.TrimPrefix(created.GetName(), "notificationChannels/")
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *MonitoringNotificationChannelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	req := &pb.UpdateNotificationChannelRequest{
		NotificationChannel: direct.ProtoClone(a.desired),
	}
	req.NotificationChannel.Name = fqn

	updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for MonitoringNotificationChannel %q: %w", fqn, err)
	}
	req.UpdateMask = updateMask

	latest := a.actual
	if len(req.UpdateMask.Paths) != 0 {
		log.V(0).Info("updating MonitoringNotificationChannel", "name", fqn)

		updated, err := a.notificationChannelsClient.UpdateNotificationChannel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MonitoringNotificationChannel %q: %w", fqn, err)
		}
		log.V(0).Info("updated MonitoringNotificationChannel", "name", fqn)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *MonitoringNotificationChannelAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.NotificationChannel) error {
	status := &krm.MonitoringNotificationChannelStatus{}

	// NOTYET: observedState
	// {
	// 	mapCtx := &direct.MapContext{}
	// 	status.ObservedState = MonitoringNotificationChannelObservedState_v1alpha1_FromProto(mapCtx, latest)
	// 	if mapCtx.Err() != nil {
	// 		return mapCtx.Err()
	// 	}
	// }

	// Legacy status fields
	status.Name = direct.PtrTo(strings.TrimPrefix(latest.GetName(), "notificationChannels/"))
	verificationStatus := latest.GetVerificationStatus()
	switch verificationStatus {
	case pb.NotificationChannel_VERIFICATION_STATUS_UNSPECIFIED:
		status.VerificationStatus = nil
	default:
		status.VerificationStatus = direct.PtrTo(verificationStatus.String())
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *MonitoringNotificationChannelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("MonitoringNotificationChannel %q not found", fqn)
	}

	obj := &krm.MonitoringNotificationChannel{}

	{
		mapCtx := &direct.MapContext{}
		spec := MonitoringNotificationChannelSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.MonitoringNotificationChannelGVK)
	obj.Name = a.id.NotificationChannel

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting MonitoringNotificationChannel to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *MonitoringNotificationChannelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting MonitoringNotificationChannel", "name", fqn)

	req := &pb.DeleteNotificationChannelRequest{}
	req.Name = fqn

	if err := a.notificationChannelsClient.DeleteNotificationChannel(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent MonitoringNotificationChannel, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting MonitoringNotificationChannel %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted MonitoringNotificationChannel", "name", fqn)

	return true, nil
}

// TODO: Make generic
func (a *MonitoringNotificationChannelAdapter) changedFields(ctx context.Context) (*fieldmaskpb.FieldMask, error) {
	log := klog.FromContext(ctx)

	// Compute the actual with only the spec fields populated.
	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := MonitoringNotificationChannelSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := MonitoringNotificationChannelSpec_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	var paths []string
	fields := actualMasked.Type().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		path := string(fields.Get(i).Name())
		changed, err := direct.FieldHasChanged(ctx, path, a.desired.ProtoReflect(), actualMasked)
		if err != nil {
			log.Error(err, "error determining if field has changed", "field", path)
			// If we can't determine if the field has changed, include it in the update.
		} else if !changed {
			continue
		}
		paths = append(paths, path)
	}
	return &fieldmaskpb.FieldMask{Paths: paths}, nil
}
