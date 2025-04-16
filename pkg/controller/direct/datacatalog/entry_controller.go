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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.Entry
// crd.type: DataCatalogEntry
// crd.version: v1alpha1

package datacatalog

import (
	"context"
	"fmt"
	"slices"

	api "cloud.google.com/go/datacatalog/apiv1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DataCatalogEntryGVK, NewEntryModel)
}

func NewEntryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &entryModel{config: *config}, nil
}

var _ directbase.Model = &entryModel{}

type entryModel struct {
	config config.ControllerConfig
}

func (m *entryModel) client(ctx context.Context, projectID string) (*api.Client, error) {
	var opts []option.ClientOption

	// projectID is now passed as argument

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID // Use passed projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog entry client: %w", err)
	}

	return gcpClient, nil
}

func (m *entryModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) { // Reverted back to client.Reader
	obj := &krm.DataCatalogEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEntryIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &entryAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *entryModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type entryAdapter struct {
	gcpClient *api.Client
	id        *krm.EntryIdentity
	desired   *krm.DataCatalogEntry
	actual    *pb.Entry
	reader    client.Reader // Reverted back to client.Reader
}

var _ directbase.Adapter = &entryAdapter{}

func (a *entryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting datacatalog entry", "name", a.id)

	req := &pb.GetEntryRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		// PermissionDenied is also observed for non-existent entries
		if direct.IsPermissionDenied(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting datacatalog entry %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *entryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating datacatalog entry", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogEntrySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEntryRequest{
		Parent:  a.id.Parent().String(),
		EntryId: a.id.ID(),
		Entry:   desired,
	}
	created, err := a.gcpClient.CreateEntry(ctx, req)
	if err != nil {
		return fmt.Errorf("creating datacatalog entry %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created datacatalog entry in gcp", "name", a.id)

	status := &krm.DataCatalogEntryStatus{}
	status.ObservedState = DataCatalogEntryObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *entryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating datacatalog entry", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogEntrySpec_ToProto(mapCtx, &a.desired.Spec)
	// Removed err variable from assignment above as function now returns 1 value
	if mapCtx.Err() != nil { // Check mapCtx error immediately after call
		return mapCtx.Err()
	}
	// Removed redundant mapCtx check below
	desired.Name = a.id.String() // Set the name field for update requests

	diffs, err := common.CompareProtoMessage(desired, a.actual, nil)
	if err != nil {
		return fmt.Errorf("comparing desired and actual proto: %w", err)
	}

	updateMask := &fieldmaskpb.FieldMask{}

	// These fields are specified as mutable in the API documentation or by convention
	// FQN is immutable after creation. EntryType and System oneofs are generally immutable.
	updatableFields := []string{
		"linked_resource",
		"display_name",
		"description",
		"schema",
		"source_system_timestamps",
		"labels",
		"gcs_fileset_spec", // Assuming this can be updated if type is FILESET
		// Specs within the 'spec' oneof
		"database_table_spec",
		"data_source_connection_spec",
		"routine_spec",
		"dataset_spec",
		"fileset_spec",
		"service_spec",
		"model_spec",
		"feature_online_store_spec",
		// Specs within the 'system_spec' oneof (usually tied to user_specified_system which is immutable)
		"sql_database_system_spec",
		"looker_system_spec",
		// BusinessContext has its own modify methods, likely not updatable here.
	}

	for _, field := range updatableFields {
		if diffs.Has(field) {
			updateMask.Paths = append(updateMask.Paths, field)
		}
	}

	// Handle oneof fields explicitly if needed, though changing type/system is usually disallowed.
	// For example, if GcsFilesetSpec needs update, ensure the path is added.
	if diffs.Has("gcs_fileset_spec") && !slices.Contains(updateMask.Paths, "gcs_fileset_spec") {
		updateMask.Paths = append(updateMask.Paths, "gcs_fileset_spec")
	}
	// Add similar checks for other oneof fields if they are mutable

	var updated *pb.Entry
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		log.V(2).Info("updating fields", "name", a.id, "fields", updateMask.Paths)
		req := &pb.UpdateEntryRequest{
			Entry:      desired,
			UpdateMask: updateMask,
		}
		updated, err = a.gcpClient.UpdateEntry(ctx, req)
		if err != nil {
			return fmt.Errorf("updating datacatalog entry %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated datacatalog entry in gcp", "name", a.id)
	}

	status := &krm.DataCatalogEntryStatus{}
	status.ObservedState = DataCatalogEntryObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *entryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataCatalogEntry{}
	mapCtx := &direct.MapContext{}
	specProto := DataCatalogEntrySpec_FromProto(mapCtx, a.actual) // Renamed to avoid conflict if 'spec' name is used later
	obj.Spec = direct.ValueOf(specProto)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate the reference from the identity
	// Assuming krm.EntryParent has an EntryGroupRef field holding the KRM reference.
	if parent := a.id.Parent(); parent != nil { // Check parent and nested ref
		obj.Spec.EntryGroupRef.External = parent.String()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.SetName(a.id.ID())
	u.SetNamespace(a.desired.GetNamespace()) // Preserve namespace
	u.SetGroupVersionKind(krm.DataCatalogEntryGVK)
	// Set labels and annotations from desired?
	// u.SetLabels(a.desired.GetLabels())
	// u.SetAnnotations(a.desired.GetAnnotations())

	u.Object = uObj
	// Clear status fields from the exported object
	unstructured.RemoveNestedField(u.Object, "status")
	return u, nil
}

func (a *entryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting datacatalog entry", "name", a.id)

	// Check if the entry type is one that can be deleted.
	// "You can delete only the entries created by the CreateEntry method."
	// This typically means custom types or specific supported types like FILESET.
	// Entries synced from other systems (BigQuery, PubSub etc.) usually cannot be deleted via API.
	isDeletable := false
	if a.actual != nil { // If Find was successful
		switch a.actual.System.(type) {
		case *pb.Entry_UserSpecifiedSystem:
			isDeletable = true // Assume user-specified system entries are deletable
		case *pb.Entry_IntegratedSystem:
			// Check if the integrated system entry type allows deletion
			switch a.actual.GetIntegratedSystem() {
			case pb.IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED:
				// Unspecified might imply custom or deletable? Needs clarification. Let's assume not deletable for safety.
			case pb.IntegratedSystem_BIGQUERY, pb.IntegratedSystem_CLOUD_PUBSUB, pb.IntegratedSystem_DATAPROC_METASTORE, pb.IntegratedSystem_DATAPLEX, pb.IntegratedSystem_CLOUD_SPANNER, pb.IntegratedSystem_CLOUD_BIGTABLE, pb.IntegratedSystem_CLOUD_SQL, pb.IntegratedSystem_LOOKER, pb.IntegratedSystem_VERTEX_AI:
				// These are typically synced and not deletable via DataCatalog API directly.
				log.V(1).Info("Skipping deletion of synced datacatalog entry", "name", a.id, "system", a.actual.GetIntegratedSystem())
				return false, fmt.Errorf("entries synced from %s cannot be deleted directly via the Data Catalog API", a.actual.GetIntegratedSystem().String()) // Use String() for better error message
			}
		}
		// Also check entry type if needed, e.g., FILESET might be deletable even if integrated.
		switch a.actual.EntryType.(type) {
		case *pb.Entry_UserSpecifiedType:
			isDeletable = true
		case *pb.Entry_Type:
			// Per CreateEntry docs, only FILESET, CLUSTER, DATA_STREAM or custom types are creatable/deletable.
			switch a.actual.GetType() {
			case pb.EntryType_FILESET, pb.EntryType_CLUSTER, pb.EntryType_DATA_STREAM:
				isDeletable = true
			}
		}
	} else {
		// If Find failed, assume it might be deletable (or already deleted).
		// Let the API call determine deletability.
		isDeletable = true
	}

	if !isDeletable {
		// This scenario might occur if the resource was somehow created outside KCC for a non-deletable type.
		log.V(1).Info("Skipping deletion attempt for non-user-managed datacatalog entry", "name", a.id)
		// Returning 'false' indicates KCC should abandon the deletion.
		// Returning an error might cause retries. Abandoning seems safer.
		return false, fmt.Errorf("entry %q is not of a type that can be deleted via the Data Catalog API", a.id.String())
	}

	req := &pb.DeleteEntryRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent datacatalog entry, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		// PermissionDenied can also mean not found
		if direct.IsPermissionDenied(err) {
			log.V(2).Info("skipping delete for datacatalog entry due to permission denied, potentially already deleted or access revoked", "name", a.id)
			// Treat as deleted to avoid endless loops if permissions are gone.
			return true, nil
		}
		return false, fmt.Errorf("deleting datacatalog entry %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted datacatalog entry", "name", a.id)

	return true, nil
}
