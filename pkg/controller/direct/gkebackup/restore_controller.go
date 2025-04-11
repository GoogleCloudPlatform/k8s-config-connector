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
// proto.message: google.cloud.gkebackup.v1.Restore
// crd.type: GKEBackupRestore
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
	registry.RegisterModel(krm.GKEBackupRestoreGVK, NewRestoreModel)
}

func NewRestoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &restoreModel{config: *config}, nil
}

var _ directbase.Model = &restoreModel{}

type restoreModel struct {
	config config.ControllerConfig
}

func (m *restoreModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.GKEBackupRestore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewRestoreIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// normalize reference fields
	if obj.Spec.BackupRef != nil {
		if _, err := obj.Spec.BackupRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
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
	return &restoreAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *restoreModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type restoreAdapter struct {
	gcpClient         *gcp.BackupForGKEClient
	id                *krm.RestoreIdentity
	desired           *krm.GKEBackupRestore
	actual            *pb.Restore
	resourceOverrides resourceoverrides.ResourceOverrides
}

var _ directbase.Adapter = &restoreAdapter{}

func (a *restoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting gkebackup restore", "name", a.id)

	req := &pb.GetRestoreRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetRestore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting gkebackup restore %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *restoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating gkebackup restore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupRestoreSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateRestoreRequest{
		Parent:    a.id.Parent().String(),
		RestoreId: a.id.ID(),
		Restore:   resource,
	}
	op, err := a.gcpClient.CreateRestore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating gkebackup restore %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("gkebackup restore %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created gkebackup restore in gcp", "name", a.id)

	status := &krm.GKEBackupRestoreStatus{}
	status.ObservedState = GKEBackupRestoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *restoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating gkebackup restore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupRestoreSpec_ToProto(mapCtx, &desired.Spec)
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

	// Immutable fields like filter and volume_data_restore_policy_overrides cannot be updated.

	var updated *pb.Restore
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateRestoreRequest{
			Restore:    resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateRestore(ctx, req)
		if err != nil {
			return fmt.Errorf("updating gkebackup restore %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("gkebackup restore %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated gkebackup restore", "name", a.id)
	}

	status := &krm.GKEBackupRestoreStatus{}
	status.ObservedState = GKEBackupRestoreObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *restoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEBackupRestore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEBackupRestoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// Parent is RestorePlan
	obj.Spec.RestorePlanRef = &krm.RestorePlanRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name) // Name is already set by the parent ID parsing
	u.SetGroupVersionKind(krm.GKEBackupRestoreGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *restoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting gkebackup restore", "name", a.id)

	req := &pb.DeleteRestoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRestore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Restore, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting gkebackup restore %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted gkebackup restore", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Restore %s: %w", a.id, err)
	}
	return true, nil
}
