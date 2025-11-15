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

package firestore

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.FirestoreBackupScheduleGVK, NewFirestoreBackupScheduleModel)
}

func NewFirestoreBackupScheduleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreBackupScheduleModel{config: config}, nil
}

var _ directbase.Model = &firestoreBackupScheduleModel{}

type firestoreBackupScheduleModel struct {
	config *config.ControllerConfig
}

func (m *firestoreBackupScheduleModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreBackupSchedule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	var id *krm.FirestoreBackupScheduleIdentity
	{
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		if idFromObject != nil {
			id = idFromObject.(*krm.FirestoreBackupScheduleIdentity)
		}
	}

	var desired *pb.BackupSchedule
	{
		mapCtx := &direct.MapContext{}
		desired = FirestoreBackupScheduleSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &firestoreBackupScheduleAdapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
		desired:              desired,
	}, nil
}

func (m *firestoreBackupScheduleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//firestore.googleapis.com/")

	id := &krm.FirestoreBackupScheduleIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &firestoreBackupScheduleAdapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
	}, nil
}

type firestoreBackupScheduleAdapter struct {
	id                   *krm.FirestoreBackupScheduleIdentity
	firestoreAdminClient *apiv1.FirestoreAdminClient
	desired              *pb.BackupSchedule
	actual               *pb.BackupSchedule
}

var _ directbase.Adapter = &firestoreBackupScheduleAdapter{}

func (a *firestoreBackupScheduleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting FirestoreBackupSchedule", "name", fqn)

	req := &pb.GetBackupScheduleRequest{Name: fqn}
	firestorebackupschedulepb, err := a.firestoreAdminClient.GetBackupSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreBackupSchedule %q: %w", fqn, err)
	}

	a.actual = firestorebackupschedulepb
	return true, nil
}

func (a *firestoreBackupScheduleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// Server-generated id, so no fqn until after creation
	parent := a.parent

	log := klog.FromContext(ctx)
	log.V(0).Info("creating FirestoreBackupSchedule", "parent", parent)

	req := &pb.CreateBackupScheduleRequest{
		BackupSchedule: direct.ProtoClone(a.desired),
	}
	req.Parent = parent

	created, err := a.firestoreAdminClient.CreateBackupSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreBackupSchedule with parent=%s: %w", parent, err)
	}
	log.V(0).Info("successfully created FirestoreBackupSchedule", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *firestoreBackupScheduleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	req := &pb.UpdateBackupScheduleRequest{
		BackupSchedule: direct.ProtoClone(a.desired),
	}
	req.BackupSchedule.Name = fqn

	updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for FirestoreBackupSchedule %q: %w", fqn, err)
	}
	req.UpdateMask = updateMask

	latest := a.desired
	if len(req.UpdateMask.Paths) != 0 {
		log.V(0).Info("updating FirestoreBackupSchedule", "name", fqn)

		updated, err := a.firestoreAdminClient.UpdateBackupSchedule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FirestoreBackupSchedule %q: %w", fqn, err)
		}
		log.V(0).Info("successfully updated FirestoreBackupSchedule", "name", fqn)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *firestoreBackupScheduleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.BackupSchedule) error {
	status := &krm.FirestoreBackupScheduleStatus{}
	{
		mapCtx := &direct.MapContext{}
		status.ObservedState = FirestoreBackupScheduleObservedState_v1alpha1_FromProto(mapCtx, latest)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	status.ExternalRef = direct.PtrTo(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func (a *firestoreBackupScheduleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreBackupSchedule %q not found", fqn)
	}

	obj := &krm.FirestoreBackupSchedule{}

	{
		mapCtx := &direct.MapContext{}
		spec := FirestoreBackupScheduleSpec_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.FirestoreBackupScheduleGVK)
	obj.Name = a.id.Field
	obj.Spec.DatabaseRef.External = a.id.Parent.Parent.String()

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreBackupSchedule to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *firestoreBackupScheduleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id == nil {
		// Nothing to delete
		return false, nil
	}

	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting FirestoreBackupSchedule", "name", fqn)

	req := &pb.DeleteBackupScheduleRequest{}
	req.Name = fqn

	if err := a.firestoreAdminClient.DeleteBackupSchedule(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent FirestoreBackupSchedule, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreBackupSchedule %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted FirestoreBackupSchedule", "name", fqn)
	return true, nil
}

// TODO: Deduplicate
func (a *firestoreBackupScheduleAdapter) changedFields(ctx context.Context) (*fieldmaskpb.FieldMask, error) {
	log := klog.FromContext(ctx)

	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := FirestoreBackupScheduleSpec_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := FirestoreBackupScheduleSpec_v1alpha1_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	var paths []string
	for _, path := range a.allFields().Paths {
		changed, err := fieldHasChanged(path, a.desired.ProtoReflect(), actualMasked)
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
