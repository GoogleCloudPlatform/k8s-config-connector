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

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	firestorepb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "firestoredatabase-controller"
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

func (m *firestoreDatabaseModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
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
	actual               *firestorepb.Database
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	fqn := a.id.String()
	log.V(2).Info("getting FirestoreDatabase", "name", fqn)

	if a.id == nil {
		return false, nil
	}

	req := &firestorepb.GetDatabaseRequest{Name: fqn}
	firestoredatabasepb, err := a.firestoreAdminClient.GetDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreDatabase %q: %w", fqn, err)
	}

	a.actual = firestoredatabasepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()

	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating FirestoreDatabase", "name", fqn)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirestoreDatabaseSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = fqn

	// Apply default values.
	ApplyFirestoreDatabaseDefaults(resource)

	req := &firestorepb.CreateDatabaseRequest{
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
		_, err := a.firestoreAdminClient.GetDatabase(ctx, &firestorepb.GetDatabaseRequest{Name: fqn})
		if err == nil {
			break
		}
		log.Info("waiting for FirestoreDatabase to be queryable after creation", "name", fqn, "error", err)
		time.Sleep(15 * time.Second)
	}

	status := &krm.FirestoreDatabaseStatus{}
	status.ObservedState = FirestoreDatabaseObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating FirestoreDatabase", "name", fqn)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FirestoreDatabaseSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Apply default values.
	ApplyFirestoreDatabaseDefaults(resource)

	newDb := proto.Clone(a.actual).(*firestorepb.Database)

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.ConcurrencyMode, a.actual.ConcurrencyMode) {
		// Skip update if concurrency_mode is unspecified
		if resource.ConcurrencyMode != firestorepb.Database_CONCURRENCY_MODE_UNSPECIFIED {
			newDb.ConcurrencyMode = resource.ConcurrencyMode
			updateMask.Paths = append(updateMask.Paths, "concurrency_mode")
		}
	}
	if !reflect.DeepEqual(resource.PointInTimeRecoveryEnablement, a.actual.PointInTimeRecoveryEnablement) {
		// Skip update if point_in_time_recovery_enablement is unspecified
		if resource.PointInTimeRecoveryEnablement != firestorepb.Database_POINT_IN_TIME_RECOVERY_ENABLEMENT_UNSPECIFIED {
			newDb.PointInTimeRecoveryEnablement = resource.PointInTimeRecoveryEnablement
			updateMask.Paths = append(updateMask.Paths, "point_in_time_recovery_enablement")
		}
	}
	if !reflect.DeepEqual(resource.DeleteProtectionState, a.actual.DeleteProtectionState) {
		// Skip update if delete_protection_state is unspecified
		if resource.DeleteProtectionState != firestorepb.Database_DELETE_PROTECTION_STATE_UNSPECIFIED {
			newDb.DeleteProtectionState = resource.DeleteProtectionState
			updateMask.Paths = append(updateMask.Paths, "delete_protection_state")
		}
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	req := &firestorepb.UpdateDatabaseRequest{
		Database:   newDb,
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

	status := &krm.FirestoreDatabaseStatus{}
	status.ObservedState = FirestoreDatabaseObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreDatabase %q not found", fqn)
	}

	mapCtx := &direct.MapContext{}
	dbSpec := FirestoreDatabaseSpec_FromProto(mapCtx, a.actual)

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

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting FirestoreDatabase", "name", fqn)

	req := &firestorepb.DeleteDatabaseRequest{
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

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
