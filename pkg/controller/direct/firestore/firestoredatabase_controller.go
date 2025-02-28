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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	gcp "cloud.google.com/go/firestore/apiv1"
	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	firestorepb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "firestoredatabase-controller"
	serviceDomain = "//firestore.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.FirestoreDatabaseGVK, NewModel)
	fuzztesting.RegisterKRMFuzzer(firestoreDatabaseFuzzer())
}

func firestoreDatabaseFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Database{},
		FirestoreDatabaseSpec_FromProto, FirestoreDatabaseSpec_ToProto,
		FirestoreDatabaseObservedState_FromProto, FirestoreDatabaseObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".delete_time")
	f.UnimplementedFields.Insert(".key_prefix")
	f.UnimplementedFields.Insert(".cmek_config")
	f.UnimplementedFields.Insert(".previous_id")
	f.UnimplementedFields.Insert(".source_info")

	// Default value fields set by controller
	f.UnimplementedFields.Insert(".type")
	f.UnimplementedFields.Insert(".app_engine_integration_mode")
	f.UnimplementedFields.Insert(".delete_protection_state")

	f.SpecFields.Insert(".location_id")
	f.SpecFields.Insert(".concurrency_mode")
	f.SpecFields.Insert(".point_in_time_recovery_enablement")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".version_retention_period")
	f.StatusFields.Insert(".earliest_version_time")
	f.StatusFields.Insert(".etag")

	return f
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building firestore client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	firestoreAdminClient, err := gcpClient.newFirestoreAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get Resource ID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get Project ID
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	var id *FirestoreDatabaseIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != projectID {
			return nil, fmt.Errorf("FirestoreDatabase %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, projectID)
		}
		if id.firestoredatabase != resourceID {
			return nil, fmt.Errorf("FirestoreDatabase  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.firestoredatabase, resourceID)
		}
	}

	return &Adapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
		desired:              obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id                   *FirestoreDatabaseIdentity
	firestoreAdminClient *apiv1.FirestoreAdminClient
	desired              *krm.FirestoreDatabase
	actual               *firestorepb.Database
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting FirestoreDatabase", "name", a.id.FullyQualifiedName())

	if a.id == nil {
		return false, nil
	}

	req := &firestorepb.GetDatabaseRequest{Name: a.id.FullyQualifiedName()}
	firestoredatabasepb, err := a.firestoreAdminClient.GetDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreDatabase %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = firestoredatabasepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating FirestoreDatabase", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	projectID := a.id.project
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.id.firestoredatabase == "" {
		return fmt.Errorf("resourceID is empty")
	}

	desired := a.desired.DeepCopy()
	resource := FirestoreDatabaseSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.FullyQualifiedName()

	// Apply default values.
	ApplyFirestoreDatabaseDefaults(resource)

	req := &firestorepb.CreateDatabaseRequest{
		Parent:     a.id.Parent(),
		Database:   resource,
		DatabaseId: a.id.firestoredatabase,
	}
	op, err := a.firestoreAdminClient.CreateDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreDatabase %s: %w", a.id.FullyQualifiedName(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FirestoreDatabase %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created FirestoreDatabase", "name", a.id.FullyQualifiedName())

	status := &krm.FirestoreDatabaseStatus{}
	status.ObservedState = FirestoreDatabaseObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating FirestoreDatabase", "name", a.id.FullyQualifiedName())
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

	if len(updateMask.Paths) == 0 {
		return nil
	}

	req := &firestorepb.UpdateDatabaseRequest{
		Database:   newDb,
		UpdateMask: updateMask,
	}
	op, err := a.firestoreAdminClient.UpdateDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("updating FirestoreDatabase %q: %w", a.id.FullyQualifiedName(), err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FirestoreDatabase %s waiting update: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated FirestoreDatabase", "name", a.id.FullyQualifiedName())

	status := &krm.FirestoreDatabaseStatus{}
	status.ObservedState = FirestoreDatabaseObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreDatabase %q not found", a.id.FullyQualifiedName())
	}

	mapCtx := &direct.MapContext{}
	dbSpec := FirestoreDatabaseSpec_FromProto(mapCtx, a.actual)

	db := &krm.FirestoreDatabase{
		Spec: *dbSpec,
	}
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(db)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreDatabase to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: obj,
	}
	u.SetName(a.id.firestoredatabase)
	u.SetGroupVersionKind(krm.FirestoreDatabaseGVK)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting FirestoreDatabase", "name", a.id.FullyQualifiedName())

	req := &firestorepb.DeleteDatabaseRequest{
		Name: a.id.FullyQualifiedName(),
		Etag: a.actual.Etag,
	}
	op, err := a.firestoreAdminClient.DeleteDatabase(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting FirestoreDatabase %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted FirestoreDatabase", "name", a.id.FullyQualifiedName())

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete FirestoreDatabase %s: %w", a.id.FullyQualifiedName(), err)
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
