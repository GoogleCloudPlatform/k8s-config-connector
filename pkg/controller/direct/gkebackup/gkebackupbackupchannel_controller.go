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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupChannel
// crd.type: GKEBackupBackupChannel
// crd.version: v1alpha1

package gkebackup

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/gkebackup/apiv1"
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.GKEBackupBackupChannelGVK, NewBackupChannelModel)
}

func NewBackupChannelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backupChannelModel{config: *config}, nil
}

var _ directbase.Model = &backupChannelModel{}

type backupChannelModel struct {
	config config.ControllerConfig
}

func (m *backupChannelModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GKEBackupBackupChannel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupChannelIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get gkebackup GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newBackupForGKEClient(ctx)
	if err != nil {
		return nil, err
	}
	return &backupChannelAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *backupChannelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type backupChannelAdapter struct {
	gcpClient         *gcp.BackupForGKEClient
	id                *krm.BackupChannelIdentity
	desired           *krm.GKEBackupBackupChannel
	actual            *pb.BackupChannel
	resourceOverrides resourceoverrides.ResourceOverrides
}

var _ directbase.Adapter = &backupChannelAdapter{}

func (a *backupChannelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting gkebackup backupchannel", "name", a.id)

	req := &pb.GetBackupChannelRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBackupChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting gkebackup backupchannel %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *backupChannelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating gkebackup backupchannel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupChannelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBackupChannelRequest{
		Parent:          a.id.Parent().String(),
		BackupChannelId: a.id.ID(),
		BackupChannel:   resource,
	}
	op, err := a.gcpClient.CreateBackupChannel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating gkebackup backupchannel %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("gkebackup backupchannel %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created gkebackup backupchannel in gcp", "name", a.id)

	status := &krm.GKEBackupBackupChannelStatus{}
	status.ObservedState = GKEBackupBackupChannelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *backupChannelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating gkebackup backupchannel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupChannelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.GetDescription(), a.actual.GetDescription()) {
		report.AddField("description", a.actual.GetDescription(), resource.GetDescription())
		paths = append(paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.GetLabels(), a.actual.GetLabels()) {
		report.AddField("labels", a.actual.GetLabels(), resource.GetLabels())
		paths = append(paths, "labels")
	}

	var updated *pb.BackupChannel
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateBackupChannelRequest{
			BackupChannel: resource,
			UpdateMask:    &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateBackupChannel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating gkebackup backupchannel %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("gkebackup backupchannel %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated gkebackup backupchannel in gcp", "name", a.id)
	}

	status := &krm.GKEBackupBackupChannelStatus{}
	status.ObservedState = GKEBackupBackupChannelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *backupChannelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEBackupBackupChannel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEBackupBackupChannelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate required Spec fields from identity
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{
		External: a.id.Parent().ProjectID,
	}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.GKEBackupBackupChannelGVK)
	u.Object = uObj

	// Populate status
	status := GKEBackupBackupChannelObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	statusMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&krm.GKEBackupBackupChannelStatus{
		ObservedState: status,
		ExternalRef:   direct.LazyPtr(a.id.String()),
	})
	if err != nil {
		return nil, err
	}
	if err := unstructured.SetNestedField(u.Object, statusMap, "status"); err != nil {
		return nil, err
	}

	return u, nil
}

func (a *backupChannelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting gkebackup backupchannel", "name", a.id)

	req := &pb.DeleteBackupChannelRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupChannel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting gkebackup backupchannel %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting deletion of gkebackup backupchannel %s: %w", a.id.String(), err)
	}
	return true, nil
}
