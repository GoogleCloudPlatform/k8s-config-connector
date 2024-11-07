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

package bigquerydataset

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	bigquery "cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "bigquery-controller"
	serviceDomain = "//bigquery.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.BigQueryDatasetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) service(ctx context.Context, projectID string) (*bigquery.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpService, err := bigquery.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Dataset client: %w", err)
	}
	return gcpService, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewDatasetIdentity(ctx, reader, obj)
	hasUpdateIsCaseInsensitive := false
	if obj.Spec.IsCaseInsensitive != nil {
		hasUpdateIsCaseInsensitive = true
	}
	id, err := krm.NewBigQueryDatasetRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	parent, _, err := krm.ParseBigQueryDatasetExternal(id.External)
	projectID := parent.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project ID")
	}

	// Get bigquery GCP client
	gcpService, err := m.service(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:                         id,
		gcpService:                 gcpService,
		desired:                    obj,
		reader:                     reader,
		hasUpdateIsCaseInsensitive: hasUpdateIsCaseInsensitive,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id                         *krm.DatasetIdentity
	gcpService                 *api.Service
	desired                    *krm.BigQueryDataset
	actual                     *api.Dataset
	reader                     client.Reader
	id                         *krm.BigQueryDatasetRef
	gcpService                 *bigquery.Client
	desired                    *krm.BigQueryDataset
	actual                     *bigquery.DatasetMetadata
	reader                     client.Reader
	hasUpdateIsCaseInsensitive bool
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting BigQueryDataset", "name", a.id.String())

	datasetGetCall := a.gcpService.Datasets.Get(a.id.Parent().ProjectID, a.id.ID())
	datasetpb, err := datasetGetCall.Do()
	parent, datasetId, err := krm.ParseBigQueryDatasetExternal(a.id.External)
	if err != nil {
		return false, fmt.Errorf("failed to parse bigquery dataset full name, %w", err)
	}
	dsHandler := a.gcpService.DatasetInProject(parent.ProjectID, datasetId)
	datasetpb, err := dsHandler.Metadata(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryDataset %q: %w", a.id.String(), err)
	}
	a.actual = datasetpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredDataset := BigQueryDatasetSpec_ToProto(mapCtx, &a.desired.Spec)
	desiredDataset.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		desiredDataset.Labels[k] = v
	}
	desiredDataset.Labels["managed-by-cnrm"] = "true"

	parent, datasetId, err := krm.ParseBigQueryDatasetExternal(a.id.External)
	if err != nil {
		return fmt.Errorf("failed to parse bigquery dataset full name, %w", err)
	}
	// Resolve KMS key reference
	if a.desired.Spec.DefaultEncryptionConfiguration != nil {
		kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, a.desired, a.desired.Spec.DefaultEncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return err
		}
		desiredDataset.DefaultEncryptionConfig.KMSKeyName = kmsRef.External
	}
	insertDatasetCall := a.gcpService.Datasets.Insert(a.id.Parent().ProjectID, desiredDataset)
	inserted, err := insertDatasetCall.Do()
	if err != nil {
		return fmt.Errorf("inserting Dataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully inserted Dataset", "name", a.id.String())
	dsHandler := a.gcpService.DatasetInProject(parent.ProjectID, datasetId)
	if err := dsHandler.Create(ctx, desiredDataset); err != nil {
		return fmt.Errorf("Error creating Dataset %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created Dataset", "name", a.id.External)

	// The bigquery go client Create() does not return the created dataset.
	// Fetching the dataset metadata
	createdMetadata, err := dsHandler.Metadata(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("Error getting the created BigQueryDataset %q: %w", a.id.External, err)
	}
	status := &krm.BigQueryDatasetStatus{}
	status = BigQueryDatasetStatus_FromProto(mapCtx, createdMetadata)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	external := a.id.String()
	status.ExternalRef = &external
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Convert KRM object to proto message
	desiredKRM := a.desired.DeepCopy()
	desired := BigQueryDatasetSpec_ToProto(mapCtx, &desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource := cloneBigQueryDatasetMetadate(a.actual)
	// Check for immutable fields
	if desiredKRM.Spec.Location != nil && !reflect.DeepEqual(desired.Location, resource.Location) {
		return fmt.Errorf("BigQueryDataset %s/%s location cannot be changed, actual: %s, desired: %s", u.GetNamespace(), u.GetName(), resource.Location, desired.Location)
	}
	// Find diff
	updateMask := &fieldmaskpb.FieldMask{}
	if desired.Description != "" && !reflect.DeepEqual(desired.Description, resource.Description) {
		resource.Description = desired.Description
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if desired.Name != "" && !reflect.DeepEqual(desired.Name, resource.Name) {
		resource.Name = desired.Name
		updateMask.Paths = append(updateMask.Paths, "friendly_name")
	}
	if desired.DefaultPartitionExpiration != 0 && !reflect.DeepEqual(desired.DefaultPartitionExpiration, resource.DefaultPartitionExpiration) {
		resource.DefaultPartitionExpiration = desired.DefaultPartitionExpiration
		updateMask.Paths = append(updateMask.Paths, "default_partition_expiration")
	}
	if desired.DefaultTableExpiration != 0 && !reflect.DeepEqual(desired.DefaultTableExpiration, resource.DefaultTableExpiration) {
		resource.DefaultTableExpiration = desired.DefaultTableExpiration
		updateMask.Paths = append(updateMask.Paths, "default_table_expiration")
	}
	if desired.DefaultCollation != "" && !reflect.DeepEqual(desired.DefaultCollation, resource.DefaultCollation) {
		resource.DefaultCollation = desired.DefaultCollation
		updateMask.Paths = append(updateMask.Paths, "default_collation")
	}
	if desired.DefaultEncryptionConfig != nil && resource.DefaultEncryptionConfig != nil && !reflect.DeepEqual(desired.DefaultEncryptionConfig, resource.DefaultEncryptionConfig) {
		// Resolve KMS key reference
		if a.desired.Spec.DefaultEncryptionConfiguration != nil {
			kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, a.desired, a.desired.Spec.DefaultEncryptionConfiguration.KmsKeyRef)
			if err != nil {
				return err
			}
			desired.DefaultEncryptionConfig.KMSKeyName = kmsRef.External
		}
		resource.DefaultEncryptionConfig.KMSKeyName = desired.DefaultEncryptionConfig.KMSKeyName
		updateMask.Paths = append(updateMask.Paths, "default_encryption_configuration")
	}
	if a.hasUpdateIsCaseInsensitive && !reflect.DeepEqual(desired.IsCaseInsensitive, resource.IsCaseInsensitive) {
		resource.IsCaseInsensitive = desired.IsCaseInsensitive
		updateMask.Paths = append(updateMask.Paths, "is_case_sensitive")
	}
	if desired.StorageBillingModel != "" && !reflect.DeepEqual(desired.StorageBillingModel, resource.StorageBillingModel) {
		resource.StorageBillingModel = desired.StorageBillingModel
		updateMask.Paths = append(updateMask.Paths, "storage_billing_model")
	}
	// If we do not set a value, the GCP service defaults to 168
	// If the existing value is 168, it means that we did not set this field at creation and it defaults to 168.
	// So if the desired value is 0, it means that we do not intend to update this field.
	if desired.MaxTimeTravel != 0 && !reflect.DeepEqual(desired.MaxTimeTravel, resource.MaxTimeTravel) && (resource.MaxTimeTravel != 168 && desired.MaxTimeTravel != 0) {
		resource.MaxTimeTravel = desired.MaxTimeTravel
		updateMask.Paths = append(updateMask.Paths, "max_time_travel")
	}
	if desired.Access != nil && resource.Access != nil && len(desired.Access) > 0 && !reflect.DeepEqual(desired.Access, resource.Access) {
		for _, access := range desired.Access {
			resource.Access = append(resource.Access, access)
		}
	}
	if len(updateMask.Paths) == 0 {
		return nil
	}

	// Get parent information
	parent, datasetId, err := krm.ParseBigQueryDatasetExternal(a.id.External)
	if err != nil {
		return fmt.Errorf("failed to parse bigquery dataset full name, %w", err)
	}

	// Compute the dataset metadate for update request
	datasetMetadataToUpdate := BigQueryDataset_ToMetadataToUpdate(mapCtx, resource, updateMask.Paths)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		datasetMetadataToUpdate.SetLabel(k, v)
	}
	updateDatasetCall := a.gcpService.Datasets.Update(a.id.Parent().ProjectID, a.id.ID(), resource)
	updated, err := updateDatasetCall.Do()
	datasetMetadataToUpdate.SetLabel("managed-by-cnrm", "true")
	// Call update
	dsHandler := a.gcpService.DatasetInProject(parent.ProjectID, datasetId)
	updated, err := dsHandler.Update(ctx, *datasetMetadataToUpdate, resource.ETag)
	if err != nil {
		return fmt.Errorf("updating Dataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Dataset", "name", a.id.String())

	status := &krm.BigQueryDatasetStatus{}
	status = BigQueryDatasetStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryDatasetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting Dataset", "name", a.id.String())

	deleteDatasetCall := a.gcpService.Datasets.Delete(a.id.Parent().ProjectID, a.id.ID())
	err := deleteDatasetCall.Do()
	if err != nil {
		return false, fmt.Errorf("deleting Dataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Dataset", "name", a.id.String())
	dsHandler := a.gcpService.DatasetInProject(parent.ProjectID, datasetId)
	if err := dsHandler.DeleteWithContents(ctx); err != nil {
		return false, fmt.Errorf("deleting Dataset %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted Dataset", "name", a.id.External)

	return true, nil
}
