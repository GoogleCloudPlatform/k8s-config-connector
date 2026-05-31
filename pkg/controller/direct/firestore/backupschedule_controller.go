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

package firestore

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FirestoreBackupScheduleGVK, NewFirestoreBackupScheduleModel)
}

func NewFirestoreBackupScheduleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreBackupScheduleModel{config: config}, nil
}

type firestoreBackupScheduleModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &firestoreBackupScheduleModel{}

func (m *firestoreBackupScheduleModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreBackupSchedule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := FirestoreBackupScheduleSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &firestoreBackupScheduleAdapter{
		id:                   id.(*krm.FirestoreBackupScheduleIdentity),
		firestoreAdminClient: firestoreAdminClient,
		desired:              desired,
	}, nil
}

func (m *firestoreBackupScheduleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
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

	if a.id == nil || a.id.BackupSchedule == "" {
		return false, nil
	}

	fqn := a.id.String()
	log.V(2).Info("getting FirestoreBackupSchedule", "name", fqn)

	req := &pb.GetBackupScheduleRequest{Name: fqn}
	obj, err := a.firestoreAdminClient.GetBackupSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreBackupSchedule %q: %w", fqn, err)
	}

	a.actual = obj
	return true, nil
}

func (a *firestoreBackupScheduleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := fmt.Sprintf("projects/%s/databases/%s", a.id.Project, a.id.Database)
	log.V(2).Info("creating FirestoreBackupSchedule", "parent", parent)

	resource := proto.Clone(a.desired).(*pb.BackupSchedule)

	req := &pb.CreateBackupScheduleRequest{
		Parent:         parent,
		BackupSchedule: resource,
	}
	created, err := a.firestoreAdminClient.CreateBackupSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreBackupSchedule in %s: %w", parent, err)
	}

	// Update the ID since it was generated
	if err := a.id.FromExternal("//firestore.googleapis.com/" + created.Name); err != nil {
		return fmt.Errorf("failed to parse generated backup schedule name %q: %w", created.Name, err)
	}

	log.V(2).Info("successfully created FirestoreBackupSchedule", "name", created.Name)

	mapCtx := &direct.MapContext{}
	status := &krm.FirestoreBackupScheduleStatus{}
	status.ObservedState = FirestoreBackupScheduleObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreBackupScheduleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating FirestoreBackupSchedule", "name", fqn)

	resource := proto.Clone(a.desired).(*pb.BackupSchedule)
	resource.Name = fqn

	var paths []string
	if !proto.Equal(a.desired.Retention, a.actual.Retention) {
		paths = append(paths, "retention")
	}
	if !proto.Equal(a.desired.GetDailyRecurrence(), a.actual.GetDailyRecurrence()) || !proto.Equal(a.desired.GetWeeklyRecurrence(), a.actual.GetWeeklyRecurrence()) {
		paths = append(paths, "recurrence")
	}

	if len(paths) == 0 {
		mapCtx := &direct.MapContext{}
		status := &krm.FirestoreBackupScheduleStatus{}
		status.ObservedState = FirestoreBackupScheduleObservedState_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.PtrTo(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateBackupScheduleRequest{
		BackupSchedule: resource,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths,
		},
	}

	updated, err := a.firestoreAdminClient.UpdateBackupSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FirestoreBackupSchedule %q: %w", fqn, err)
	}

	log.V(2).Info("successfully updated FirestoreBackupSchedule", "name", updated.Name)

	mapCtx := &direct.MapContext{}
	status := &krm.FirestoreBackupScheduleStatus{}
	status.ObservedState = FirestoreBackupScheduleObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreBackupScheduleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreBackupSchedule %q not found", fqn)
	}

	mapCtx := &direct.MapContext{}
	objSpec := FirestoreBackupScheduleSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.FirestoreBackupSchedule{
		Spec: *objSpec,
	}

	obj.SetGroupVersionKind(krm.FirestoreBackupScheduleGVK)
	obj.Name = a.id.BackupSchedule

	obj.Spec.DatabaseRef = v1beta1.FirestoreDatabaseRef{
		External: fmt.Sprintf("projects/%s/databases/%s", a.id.Project, a.id.Database),
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreBackupSchedule to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	return u, nil
}

func (a *firestoreBackupScheduleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id == nil || a.id.BackupSchedule == "" {
		return true, nil
	}

	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirestoreBackupSchedule", "name", fqn)

	req := &pb.DeleteBackupScheduleRequest{
		Name: fqn,
	}
	if err := a.firestoreAdminClient.DeleteBackupSchedule(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent FirestoreBackupSchedule, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreBackupSchedule %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted FirestoreBackupSchedule", "name", fqn)

	return true, nil
}
