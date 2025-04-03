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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.Backup
// crd.type: GKEBackupBackup
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
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
)

func init() {
	registry.RegisterModel(krm.GKEBackupBackupGVK, NewBackupModel)
}

func NewBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backupModel{config: *config}, nil
}

var _ directbase.Model = &backupModel{}

type backupModel struct {
	config config.ControllerConfig
}

func (m *backupModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.GKEBackupBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupIdentity(ctx, reader, obj)
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
	return &backupAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *backupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type backupAdapter struct {
	gcpClient         *gcp.BackupForGKEClient
	id                *krm.BackupIdentity
	desired           *krm.GKEBackupBackup
	actual            *pb.Backup
	resourceOverrides resourceoverrides.ResourceOverrides
}

var _ directbase.Adapter = &backupAdapter{}

func (a *backupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting gkebackup backup", "name", a.id)

	req := &pb.GetBackupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting gkebackup backup %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *backupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating gkebackup backup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBackupRequest{
		Parent:   a.id.Parent().String(),
		BackupId: a.id.ID(),
		Backup:   resource,
	}
	op, err := a.gcpClient.CreateBackup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating gkebackup backup %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("gkebackup backup %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created gkebackup backup in gcp", "name", a.id)

	status := &krm.GKEBackupBackupStatus{}
	status.ObservedState = GKEBackupBackupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *backupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating gkebackup backup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}
	if desired.Spec.DeleteLockDays != nil && !reflect.DeepEqual(resource.GetDeleteLockDays(), a.actual.GetDeleteLockDays()) {
		paths = append(paths, "delete_lock_days")
	}
	if desired.Spec.RetainDays != nil && !reflect.DeepEqual(resource.GetRetainDays(), a.actual.GetRetainDays()) {
		paths = append(paths, "retain_days")
	}

	var updated *pb.Backup
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateBackupRequest{
			Backup:     resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateBackup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating gkebackup backup %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("gkebackup backup %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated gkebackup backup", "name", a.id)
	}

	status := &krm.GKEBackupBackupStatus{}
	status.ObservedState = GKEBackupBackupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *backupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEBackupBackup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEBackupBackupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate required Spec fields from identity
	obj.Spec.BackupPlanRef = &krm.BackupPlanRef{
		External: a.id.Parent().String(),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID()) // Name should be the K8s resource name, which matches the Backup ID
	u.SetGroupVersionKind(krm.GKEBackupBackupGVK)
	u.Object = uObj

	// Populate status
	status := GKEBackupBackupObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	statusMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&krm.GKEBackupBackupStatus{
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

// Delete implements the Adapter interface.
func (a *backupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting gkebackup backup", "name", a.id)

	req := &pb.DeleteBackupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Backup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting gkebackup backup %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion gkebackup backup", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Backup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted gkebackup backup", "name", a.id)
	return true, nil
}
