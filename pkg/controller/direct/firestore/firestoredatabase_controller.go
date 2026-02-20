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

package firestore

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FirestoreDatabaseGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreDatabaseModel{config: config}, nil
}

var _ directbase.Model = &firestoreDatabaseModel{}

type firestoreDatabaseModel struct {
	config *config.ControllerConfig
}

func (m *firestoreDatabaseModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:                   id.(*krm.FirestoreDatabaseIdentity),
		firestoreAdminClient: firestoreAdminClient,
		desired:              obj,
	}, nil
}

func (m *firestoreDatabaseModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//firestore.googleapis.com/")

	id := &krm.FirestoreDatabaseIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
	}, nil
}

type Adapter struct {
	id                   *krm.FirestoreDatabaseIdentity
	firestoreAdminClient *apiv1.FirestoreAdminClient
	desired              *krm.FirestoreDatabase
	actual               *pb.Database
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()

	log.V(2).Info("getting FirestoreDatabase", "name", fqn)
	req := &pb.GetDatabaseRequest{Name: fqn}
	actual, err := a.firestoreAdminClient.GetDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreDatabase %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirestoreDatabase", "name", fqn)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirestoreDatabaseSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = fqn

	// Type is required, and we default it to FIRESTORE_NATIVE if unspecified.
	if resource.Type == pb.Database_DATABASE_TYPE_UNSPECIFIED {
		resource.Type = pb.Database_FIRESTORE_NATIVE
	}

	req := &pb.CreateDatabaseRequest{
		Parent:     a.id.Parent.String(),
		Database:   resource,
		DatabaseId: a.id.Database,
	}
	op, err := a.firestoreAdminClient.CreateDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreDatabase %s: %w", fqn, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FirestoreDatabase %s waiting creation: %w", fqn, err)
	}
	log.V(2).Info("successfully created FirestoreDatabase", "name", fqn)

	// HACK: We wait for the negative-cache in the API server(s) to expire.
	time.Sleep(15 * time.Second)
	for {
		_, err := a.firestoreAdminClient.GetDatabase(ctx, &pb.GetDatabaseRequest{Name: fqn})
		if err == nil {
			break
		}
		log.Info("waiting for FirestoreDatabase to be queryable after creation", "name", fqn, "error", err)
		time.Sleep(15 * time.Second)
	}

	return a.setStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirestoreDatabase", "name", fqn)
	mapCtx := &direct.MapContext{}

	desired := FirestoreDatabaseSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = fqn

	// Type is required, and we default it to FIRESTORE_NATIVE if unspecified.
	if desired.Type == pb.Database_DATABASE_TYPE_UNSPECIFIED {
		desired.Type = pb.Database_FIRESTORE_NATIVE
	}

	// Simulate server-side defaulting, so we don't issue updates for fields that
	// are different only because of defaulting.
	expected := direct.ProtoClone(desired)
	ApplyServerSideDefaults(expected)

	actual := a.actual

	diff := &structuredreporting.Diff{}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(actual.ConcurrencyMode, expected.ConcurrencyMode) {
		diff.AddField("concurrency_mode", actual.ConcurrencyMode, desired.ConcurrencyMode)
		updateMask.Paths = append(updateMask.Paths, "concurrency_mode")
	}
	if !reflect.DeepEqual(actual.PointInTimeRecoveryEnablement, expected.PointInTimeRecoveryEnablement) {
		diff.AddField("point_in_time_recovery_enablement", actual.PointInTimeRecoveryEnablement, desired.PointInTimeRecoveryEnablement)
		updateMask.Paths = append(updateMask.Paths, "point_in_time_recovery_enablement")
	}
	if !reflect.DeepEqual(actual.AppEngineIntegrationMode, expected.AppEngineIntegrationMode) {
		diff.AddField("app_engine_integration_mode", actual.AppEngineIntegrationMode, desired.AppEngineIntegrationMode)
		updateMask.Paths = append(updateMask.Paths, "app_engine_integration_mode")
	}
	if !reflect.DeepEqual(actual.DeleteProtectionState, expected.DeleteProtectionState) {
		diff.AddField("delete_protection_state", actual.DeleteProtectionState, desired.DeleteProtectionState)
		updateMask.Paths = append(updateMask.Paths, "delete_protection_state")
	}

	structuredreporting.ReportDiff(ctx, diff)

	latest := direct.ProtoClone(actual)
	if len(updateMask.Paths) != 0 {
		req := &pb.UpdateDatabaseRequest{
			Database:   desired,
			UpdateMask: updateMask,
		}
		op, err := a.firestoreAdminClient.UpdateDatabase(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FirestoreDatabase %q: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("FirestoreDatabase %s waiting update: %w", fqn, err)
		}
		log.V(2).Info("successfully updated FirestoreDatabase", "name", fqn)
		latest = updated
	}

	return a.setStatus(ctx, updateOp, latest)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreDatabase %q not found", fqn)
	}

	mapCtx := &direct.MapContext{}
	dbSpec := FirestoreDatabaseSpec_v1beta1_FromProto(mapCtx, a.actual)

	db := &krm.FirestoreDatabase{
		Spec: *dbSpec,
	}

	db.SetGroupVersionKind(krm.FirestoreDatabaseGVK)
	db.Name = a.id.Database
	db.Spec.ProjectRef.External = a.id.Parent.String()

	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(db)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreDatabase to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: obj,
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirestoreDatabase", "name", fqn)

	req := &pb.DeleteDatabaseRequest{
		Name: fqn,
		Etag: a.actual.Etag,
	}
	op, err := a.firestoreAdminClient.DeleteDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FirestoreDatabase, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreDatabase %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted FirestoreDatabase", "name", fqn)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete FirestoreDatabase %s: %w", fqn, err)
	}
	return true, nil
}

func (a *Adapter) setStatus(ctx context.Context, op directbase.Operation, latest *pb.Database) error {
	status := &krm.FirestoreDatabaseStatus{}
	status.ExternalRef = direct.PtrTo(a.id.String())

	mapCtx := &direct.MapContext{}
	status.ObservedState = FirestoreDatabaseObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}
