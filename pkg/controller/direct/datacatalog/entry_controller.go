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

func (m *entryModel) client(ctx context.Context, obj *krm.DataCatalogEntry) (*api.Client, error) {
	var opts []option.ClientOption

	projectID := obj.Spec.EntryGroupRef.Project()

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
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

func (m *entryModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataCatalogEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEntryIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, obj)
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
	reader    client.Reader
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

	desired, err := DataCatalogEntrySpec_ToProto(mapCtx, a.reader, ctx, &a.desired.Spec)
	if err != nil {
		return err
	}
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

	desired, err := DataCatalogEntrySpec_ToProto(mapCtx, a.reader, ctx, &a.desired.Spec)
	if err != nil {
		return err
	}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String() // Set the name field for update requests

	diffs, err := common.CompareProtoMessage(desired, a.actual, common.DefaultDiff)
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
	if diffs.Has("gcs_fileset_spec") && !direct.SliceContains(updateMask.Paths, "gcs_fileset_spec") {
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
	spec, err := DataCatalogEntrySpec_FromProto(mapCtx, a.actual)
	if err != nil {
		return nil, err
	}
	obj.Spec = direct.ValueOf(spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate the reference from the identity
	obj.Spec.EntryGroupRef = a.id.ParentRef(a.reader.Scheme())

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
			case pb.IntegratedSystem_BIGQUERY, pb.IntegratedSystem_PUBSUB, pb.IntegratedSystem_DATAPROC_METASTORE, pb.IntegratedSystem_DATAPLEX, pb.IntegratedSystem_CLOUD_SPANNER, pb.IntegratedSystem_CLOUD_BIGTABLE, pb.IntegratedSystem_CLOUD_SQL, pb.IntegratedSystem_LOOKER, pb.IntegratedSystem_VERTEX_AI:
				// These are typically synced and not deletable via DataCatalog API directly.
				log.V(1).Info("Skipping deletion of synced datacatalog entry", "name", a.id, "system", a.actual.GetIntegratedSystem())
				return false, fmt.Errorf("entries synced from %s cannot be deleted directly via the Data Catalog API", a.actual.GetIntegratedSystem())
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

// Placeholder for the conversion functions. These need to be implemented based on the KRM struct definitions.
func DataCatalogEntrySpec_ToProto(mapCtx *direct.MapContext, reader client.Reader, ctx context.Context, k *krm.DataCatalogEntrySpec) (*pb.Entry, error) {
	if k == nil {
		return nil, nil
	}
	// TODO: Implement conversion logic from krm.DataCatalogEntrySpec to pb.Entry
	// This involves mapping fields and handling oneofs.
	proto := &pb.Entry{}

	// --- Oneof entry_type ---
	if k.Type != nil {
		proto.EntryType = &pb.Entry_Type{Type: pb.EntryType(pb.EntryType_value[direct.ValueOf(k.Type)])}
	} else if k.UserSpecifiedType != nil {
		proto.EntryType = &pb.Entry_UserSpecifiedType{UserSpecifiedType: direct.ValueOf(k.UserSpecifiedType)}
	} else {
		mapCtx.Errorf("entry_type: exactly one of type or userSpecifiedType must be set")
	}

	// --- Oneof system ---
	if k.UserSpecifiedSystem != nil {
		proto.System = &pb.Entry_UserSpecifiedSystem{UserSpecifiedSystem: direct.ValueOf(k.UserSpecifiedSystem)}
	}
	// IntegratedSystem is output only, so no mapping from spec

	// --- Simple Fields ---
	proto.LinkedResource = direct.ValueOf(k.LinkedResource)
	proto.FullyQualifiedName = direct.ValueOf(k.FullyQualifiedName) // Settable only during creation
	proto.DisplayName = direct.ValueOf(k.DisplayName)
	proto.Description = direct.ValueOf(k.Description)
	proto.Labels = k.Labels

	// --- Complex Fields / Specs ---
	proto.Schema = Schema_ToProto(mapCtx, k.Schema)
	proto.SourceSystemTimestamps = SystemTimestamps_ToProto(mapCtx, k.SourceSystemTimestamps)

	// --- Oneof type_spec ---
	if k.GcsFilesetSpec != nil {
		proto.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: GcsFilesetSpec_ToProto(mapCtx, k.GcsFilesetSpec)}
	}
	// BigQuery specs are output only

	// --- Oneof system_spec ---
	if k.SqlDatabaseSystemSpec != nil {
		proto.SystemSpec = &pb.Entry_SqlDatabaseSystemSpec{SqlDatabaseSystemSpec: SqlDatabaseSystemSpec_ToProto(mapCtx, k.SqlDatabaseSystemSpec)}
	}
	if k.LookerSystemSpec != nil {
		proto.SystemSpec = &pb.Entry_LookerSystemSpec{LookerSystemSpec: LookerSystemSpec_ToProto(mapCtx, k.LookerSystemSpec)}
	}
	// CloudBigtableSystemSpec is output only

	// --- Oneof spec ---
	if k.DatabaseTableSpec != nil {
		proto.Spec = &pb.Entry_DatabaseTableSpec{DatabaseTableSpec: DatabaseTableSpec_ToProto(mapCtx, k.DatabaseTableSpec)}
	} else if k.DataSourceConnectionSpec != nil {
		proto.Spec = &pb.Entry_DataSourceConnectionSpec{DataSourceConnectionSpec: DataSourceConnectionSpec_ToProto(mapCtx, k.DataSourceConnectionSpec)}
	} else if k.RoutineSpec != nil {
		proto.Spec = &pb.Entry_RoutineSpec{RoutineSpec: RoutineSpec_ToProto(mapCtx, k.RoutineSpec)}
	} else if k.DatasetSpec != nil {
		proto.Spec = &pb.Entry_DatasetSpec{DatasetSpec: DatasetSpec_ToProto(mapCtx, k.DatasetSpec)}
	} else if k.FilesetSpec != nil {
		proto.Spec = &pb.Entry_FilesetSpec{FilesetSpec: FilesetSpec_ToProto(mapCtx, k.FilesetSpec)}
	} else if k.ServiceSpec != nil {
		proto.Spec = &pb.Entry_ServiceSpec{ServiceSpec: ServiceSpec_ToProto(mapCtx, k.ServiceSpec)}
	} else if k.ModelSpec != nil {
		proto.Spec = &pb.Entry_ModelSpec{ModelSpec: ModelSpec_ToProto(mapCtx, k.ModelSpec)}
	} else if k.FeatureOnlineStoreSpec != nil {
		proto.Spec = &pb.Entry_FeatureOnlineStoreSpec{FeatureOnlineStoreSpec: FeatureOnlineStoreSpec_ToProto(mapCtx, k.FeatureOnlineStoreSpec)}
	}

	// BusinessContext has separate update methods, typically not set/updated directly here.
	// proto.BusinessContext = BusinessContext_ToProto(mapCtx, k.BusinessContext)

	// UsageSignal is output only.
	// DataSource is output only.
	// PersonalDetails is output only.

	return proto, mapCtx.Err()
}

func DataCatalogEntryObservedState_FromProto(mapCtx *direct.MapContext, p *pb.Entry) *krm.DataCatalogEntryObservedState {
	if p == nil {
		return nil
	}
	// TODO: Implement conversion logic from pb.Entry to krm.DataCatalogEntryObservedState
	// This involves mapping fields back, including output-only fields.
	k := &krm.DataCatalogEntryObservedState{}
	k.Name = &p.Name // Use pointer type for string fields mapped from proto
	k.IntegratedSystem = direct.EnumPtrFromVal(pb.IntegratedSystem_name[int32(p.GetIntegratedSystem())])
	k.BigqueryTableSpec = BigQueryTableSpec_FromProto(mapCtx, p.GetBigqueryTableSpec())
	k.BigqueryDateShardedSpec = BigQueryDateShardedSpec_FromProto(mapCtx, p.GetBigqueryDateShardedSpec())
	k.DataSource = DataSource_FromProto(mapCtx, p.GetDataSource())
	k.PersonalDetails = PersonalDetails_FromProto(mapCtx, p.GetPersonalDetails())
	k.UsageSignal = UsageSignal_FromProto(mapCtx, p.GetUsageSignal())
	k.BusinessContext = BusinessContext_FromProto(mapCtx, p.GetBusinessContext()) // Also include observed business context
	k.CloudBigtableSystemSpec = CloudBigtableSystemSpec_FromProto(mapCtx, p.GetCloudBigtableSystemSpec())

	// Map fields also present in Spec
	k.LinkedResource = &p.LinkedResource
	k.FullyQualifiedName = &p.FullyQualifiedName
	k.DisplayName = &p.DisplayName
	k.Description = &p.Description
	k.Labels = p.Labels
	k.Schema = Schema_FromProto(mapCtx, p.GetSchema())
	k.SourceSystemTimestamps = SystemTimestamps_FromProto(mapCtx, p.GetSourceSystemTimestamps())

	// Map oneofs back
	switch x := p.EntryType.(type) {
	case *pb.Entry_Type:
		k.Type = direct.EnumPtrFromVal(pb.EntryType_name[int32(x.Type)])
	case *pb.Entry_UserSpecifiedType:
		k.UserSpecifiedType = &x.UserSpecifiedType
	}

	switch x := p.System.(type) {
	case *pb.Entry_IntegratedSystem:
		k.IntegratedSystem = direct.EnumPtrFromVal(pb.IntegratedSystem_name[int32(x.IntegratedSystem)])
	case *pb.Entry_UserSpecifiedSystem:
		k.UserSpecifiedSystem = &x.UserSpecifiedSystem
	}

	switch x := p.TypeSpec.(type) {
	case *pb.Entry_GcsFilesetSpec:
		k.GcsFilesetSpec = GcsFilesetSpec_FromProto(mapCtx, x.GcsFilesetSpec)
	case *pb.Entry_BigqueryTableSpec:
		k.BigqueryTableSpec = BigQueryTableSpec_FromProto(mapCtx, x.BigqueryTableSpec)
	case *pb.Entry_BigqueryDateShardedSpec:
		k.BigqueryDateShardedSpec = BigQueryDateShardedSpec_FromProto(mapCtx, x.BigqueryDateShardedSpec)
	}

	switch x := p.SystemSpec.(type) {
	case *pb.Entry_SqlDatabaseSystemSpec:
		k.SqlDatabaseSystemSpec = SqlDatabaseSystemSpec_FromProto(mapCtx, x.SqlDatabaseSystemSpec)
	case *pb.Entry_LookerSystemSpec:
		k.LookerSystemSpec = LookerSystemSpec_FromProto(mapCtx, x.LookerSystemSpec)
	case *pb.Entry_CloudBigtableSystemSpec:
		k.CloudBigtableSystemSpec = CloudBigtableSystemSpec_FromProto(mapCtx, x.CloudBigtableSystemSpec)
	}

	switch x := p.Spec.(type) {
	case *pb.Entry_DatabaseTableSpec:
		k.DatabaseTableSpec = DatabaseTableSpec_FromProto(mapCtx, x.DatabaseTableSpec)
	case *pb.Entry_DataSourceConnectionSpec:
		k.DataSourceConnectionSpec = DataSourceConnectionSpec_FromProto(mapCtx, x.DataSourceConnectionSpec)
	case *pb.Entry_RoutineSpec:
		k.RoutineSpec = RoutineSpec_FromProto(mapCtx, x.RoutineSpec)
	case *pb.Entry_DatasetSpec:
		k.DatasetSpec = DatasetSpec_FromProto(mapCtx, x.DatasetSpec)
	case *pb.Entry_FilesetSpec:
		k.FilesetSpec = FilesetSpec_FromProto(mapCtx, x.FilesetSpec)
	case *pb.Entry_ServiceSpec:
		k.ServiceSpec = ServiceSpec_FromProto(mapCtx, x.ServiceSpec)
	case *pb.Entry_ModelSpec:
		k.ModelSpec = ModelSpec_FromProto(mapCtx, x.ModelSpec)
	case *pb.Entry_FeatureOnlineStoreSpec:
		k.FeatureOnlineStoreSpec = FeatureOnlineStoreSpec_FromProto(mapCtx, x.FeatureOnlineStoreSpec)
	}

	return k
}

func DataCatalogEntrySpec_FromProto(mapCtx *direct.MapContext, p *pb.Entry) (*krm.DataCatalogEntrySpec, error) {
	if p == nil {
		return nil, nil
	}
	// TODO: Implement conversion logic from pb.Entry to krm.DataCatalogEntrySpec
	// This involves mapping only the fields relevant to the spec (non-output-only).
	k := &krm.DataCatalogEntrySpec{}

	// --- Oneof entry_type ---
	if t, ok := p.EntryType.(*pb.Entry_Type); ok {
		k.Type = direct.EnumPtrFromVal(pb.EntryType_name[int32(t.Type)])
	} else if u, ok := p.EntryType.(*pb.Entry_UserSpecifiedType); ok {
		k.UserSpecifiedType = &u.UserSpecifiedType
	}

	// --- Oneof system ---
	if u, ok := p.System.(*pb.Entry_UserSpecifiedSystem); ok {
		k.UserSpecifiedSystem = &u.UserSpecifiedSystem
	}
	// IntegratedSystem is output only

	// --- Simple Fields ---
	k.LinkedResource = &p.LinkedResource
	// FullyQualifiedName is immutable after creation, but we include it here for export consistency
	k.FullyQualifiedName = &p.FullyQualifiedName
	k.DisplayName = &p.DisplayName
	k.Description = &p.Description
	k.Labels = p.Labels

	// --- Complex Fields / Specs ---
	k.Schema = Schema_FromProto(mapCtx, p.GetSchema())
	k.SourceSystemTimestamps = SystemTimestamps_FromProto(mapCtx, p.GetSourceSystemTimestamps())

	// --- Oneof type_spec ---
	if gcsSpec, ok := p.TypeSpec.(*pb.Entry_GcsFilesetSpec); ok {
		k.GcsFilesetSpec = GcsFilesetSpec_FromProto(mapCtx, gcsSpec.GcsFilesetSpec)
	}
	// BigQuery specs are output only

	// --- Oneof system_spec ---
	if sqlSpec, ok := p.SystemSpec.(*pb.Entry_SqlDatabaseSystemSpec); ok {
		k.SqlDatabaseSystemSpec = SqlDatabaseSystemSpec_FromProto(mapCtx, sqlSpec.SqlDatabaseSystemSpec)
	}
	if lookerSpec, ok := p.SystemSpec.(*pb.Entry_LookerSystemSpec); ok {
		k.LookerSystemSpec = LookerSystemSpec_FromProto(mapCtx, lookerSpec.LookerSystemSpec)
	}
	// CloudBigtableSystemSpec is output only

	// --- Oneof spec ---
	switch x := p.Spec.(type) {
	case *pb.Entry_DatabaseTableSpec:
		k.DatabaseTableSpec = DatabaseTableSpec_FromProto(mapCtx, x.DatabaseTableSpec)
	case *pb.Entry_DataSourceConnectionSpec:
		k.DataSourceConnectionSpec = DataSourceConnectionSpec_FromProto(mapCtx, x.DataSourceConnectionSpec)
	case *pb.Entry_RoutineSpec:
		k.RoutineSpec = RoutineSpec_FromProto(mapCtx, x.RoutineSpec)
	case *pb.Entry_DatasetSpec:
		k.DatasetSpec = DatasetSpec_FromProto(mapCtx, x.DatasetSpec)
	case *pb.Entry_FilesetSpec:
		k.FilesetSpec = FilesetSpec_FromProto(mapCtx, x.FilesetSpec)
	case *pb.Entry_ServiceSpec:
		k.ServiceSpec = ServiceSpec_FromProto(mapCtx, x.ServiceSpec)
	case *pb.Entry_ModelSpec:
		k.ModelSpec = ModelSpec_FromProto(mapCtx, x.ModelSpec)
	case *pb.Entry_FeatureOnlineStoreSpec:
		k.FeatureOnlineStoreSpec = FeatureOnlineStoreSpec_FromProto(mapCtx, x.FeatureOnlineStoreSpec)
	}

	// BusinessContext has separate update methods, typically not set/updated directly here.
	// k.BusinessContext = BusinessContext_FromProto(mapCtx, p.GetBusinessContext())

	return k, mapCtx.Err()
}

// NOTE: The detailed mapping functions (e.g., Schema_ToProto, GcsFilesetSpec_FromProto, etc.)
// are assumed to exist based on the KRM struct definitions and corresponding proto messages.
// These would need to be generated or implemented separately.
