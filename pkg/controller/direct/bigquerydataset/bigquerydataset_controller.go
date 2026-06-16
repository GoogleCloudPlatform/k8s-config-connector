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
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

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

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DatasetIdentity)

	projectID := id.Project
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project ID")
	}

	// Get bigquery GCP client
	gcpService, err := m.service(ctx, projectID)
	if err != nil {
		return nil, err
	}

	// Normalize references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := BigQueryDatasetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(obj)

	// Resolve KMS key reference. KMSCryptoKeyRef does not implement the refs.Ref interface,
	// so it cannot be normalized automatically by common.NormalizeReferences.
	// Therefore, we must resolve it manually using refs.ResolveKMSCryptoKeyRef.
	if obj.Spec.DefaultEncryptionConfiguration != nil {
		kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.DefaultEncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return nil, err
		}
		if desired.DefaultEncryptionConfig == nil {
			desired.DefaultEncryptionConfig = &bigquery.EncryptionConfig{}
		}
		desired.DefaultEncryptionConfig.KMSKeyName = kmsRef.External
	}

	var isCaseInsensitive *bool
	if obj.Spec.IsCaseInsensitive != nil {
		isCaseInsensitive = direct.LazyPtr(*obj.Spec.IsCaseInsensitive)
	}

	return &Adapter{
		id:                id,
		gcpService:        gcpService,
		desired:           desired,
		isCaseInsensitive: isCaseInsensitive,
		reader:            reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id                *krm.DatasetIdentity
	gcpService        *bigquery.Client
	desired           *bigquery.DatasetMetadata
	isCaseInsensitive *bool
	actual            *bigquery.DatasetMetadata
	reader            client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigQueryDataset", "name", a.id.String())

	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)
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

	log := klog.FromContext(ctx)
	log.V(2).Info("creating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)

	if err := dsHandler.Create(ctx, a.desired); err != nil {
		return fmt.Errorf("Error creating Dataset %s: %w", a.id.Dataset, err)
	}
	log.V(2).Info("successfully created Dataset", "name", a.id.Dataset)

	// The bigquery go client Create() does not return the created dataset.
	// Fetching the dataset metadata
	createdMetadata, err := dsHandler.Metadata(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("Error getting the created BigQueryDataset %q: %w", a.id.Dataset, err)
	}
	status := &krm.BigQueryDatasetStatus{}
	status = BigQueryDatasetStatus_FromProto(mapCtx, createdMetadata)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}
	// Write resourceID into spec.
	tokens := strings.Split(createdMetadata.FullID, ":")
	if len(tokens) == 2 {
		resourceID := tokens[1]
		if err := unstructured.SetNestedField(createOp.GetUnstructured().Object, resourceID, "spec", "resourceID"); err != nil {
			return fmt.Errorf("error setting spec.resourceID: %w", err)
		}
	} else {
		return fmt.Errorf("Error getting resourceID: %s. The full ID of the created BigQueryDataset is expected to be in the format of projectID:datasetID", createdMetadata.FullID)
	}

	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	resource := cloneBigQueryDatasetMetadate(a.actual)
	// Find diff
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	updateMask := &fieldmaskpb.FieldMask{}
	if a.desired.Description != "" && !reflect.DeepEqual(a.desired.Description, resource.Description) {
		report.AddField("description", resource.Description, a.desired.Description)
		resource.Description = a.desired.Description
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if a.desired.Name != "" && !reflect.DeepEqual(a.desired.Name, resource.Name) {
		report.AddField("friendly_name", resource.Name, a.desired.Name)
		resource.Name = a.desired.Name
		updateMask.Paths = append(updateMask.Paths, "friendly_name")
	}
	if a.desired.DefaultPartitionExpiration != 0 && !reflect.DeepEqual(a.desired.DefaultPartitionExpiration, resource.DefaultPartitionExpiration) {
		report.AddField("default_partition_expiration", resource.DefaultPartitionExpiration, a.desired.DefaultPartitionExpiration)
		resource.DefaultPartitionExpiration = a.desired.DefaultPartitionExpiration
		updateMask.Paths = append(updateMask.Paths, "default_partition_expiration")
	}
	if a.desired.DefaultTableExpiration != 0 && !reflect.DeepEqual(a.desired.DefaultTableExpiration, resource.DefaultTableExpiration) {
		report.AddField("default_table_expiration", resource.DefaultTableExpiration, a.desired.DefaultTableExpiration)
		resource.DefaultTableExpiration = a.desired.DefaultTableExpiration
		updateMask.Paths = append(updateMask.Paths, "default_table_expiration")
	}
	if a.desired.DefaultCollation != "" && !reflect.DeepEqual(a.desired.DefaultCollation, resource.DefaultCollation) {
		report.AddField("default_collation", resource.DefaultCollation, a.desired.DefaultCollation)
		resource.DefaultCollation = a.desired.DefaultCollation
		updateMask.Paths = append(updateMask.Paths, "default_collation")
	}
	if a.desired.DefaultEncryptionConfig != nil && resource.DefaultEncryptionConfig != nil && !reflect.DeepEqual(a.desired.DefaultEncryptionConfig, resource.DefaultEncryptionConfig) {
		report.AddField("default_encryption_configuration", resource.DefaultEncryptionConfig, a.desired.DefaultEncryptionConfig)
		resource.DefaultEncryptionConfig.KMSKeyName = a.desired.DefaultEncryptionConfig.KMSKeyName
		updateMask.Paths = append(updateMask.Paths, "default_encryption_configuration")
	}
	if a.isCaseInsensitive != nil && !reflect.DeepEqual(*a.isCaseInsensitive, resource.IsCaseInsensitive) {
		report.AddField("is_case_sensitive", resource.IsCaseInsensitive, *a.isCaseInsensitive)
		resource.IsCaseInsensitive = *a.isCaseInsensitive
		updateMask.Paths = append(updateMask.Paths, "is_case_sensitive")
	}
	if a.desired.StorageBillingModel != "" && !reflect.DeepEqual(a.desired.StorageBillingModel, resource.StorageBillingModel) {
		report.AddField("storage_billing_model", resource.StorageBillingModel, a.desired.StorageBillingModel)
		resource.StorageBillingModel = a.desired.StorageBillingModel
		updateMask.Paths = append(updateMask.Paths, "storage_billing_model")
	}
	// If we do not set a value, the GCP service defaults to 168
	// If the existing value is 168, it means that we did not set this field at creation and it defaults to 168.
	// So if the desired value is 0, it means that we do not intend to update this field.
	if a.desired.MaxTimeTravel != 0 && !reflect.DeepEqual(a.desired.MaxTimeTravel, resource.MaxTimeTravel) && (resource.MaxTimeTravel != 168 && a.desired.MaxTimeTravel != 0) {
		report.AddField("max_time_travel", resource.MaxTimeTravel, a.desired.MaxTimeTravel)
		resource.MaxTimeTravel = a.desired.MaxTimeTravel
		updateMask.Paths = append(updateMask.Paths, "max_time_travel")
	}
	if a.desired.Access != nil && resource.Access != nil && len(a.desired.Access) > 0 && !reflect.DeepEqual(a.desired.Access, resource.Access) {
		report.AddField("access", resource.Access, a.desired.Access)
		for _, access := range a.desired.Access {
			resource.Access = append(resource.Access, access)
		}
	}
	if len(updateMask.Paths) == 0 {
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	// Compute the dataset metadate for update request
	datasetMetadataToUpdate := BigQueryDataset_ToMetadataToUpdate(mapCtx, resource, updateMask.Paths)
	for k, v := range a.desired.Labels {
		datasetMetadataToUpdate.SetLabel(k, v)
	}
	datasetMetadataToUpdate.SetLabel("managed-by-cnrm", "true")
	// Call update
	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)
	updated, err := dsHandler.Update(ctx, *datasetMetadataToUpdate, "")
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

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Dataset", "name", a.id.String())

	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)
	annotations := deleteOp.GetUnstructured().GetAnnotations()

	// Support the existing annotation on delete.
	if annotations["cnrm.cloud.google.com/delete-contents-on-destroy"] == "true" {
		if err := dsHandler.DeleteWithContents(ctx); err != nil {
			return false, fmt.Errorf("deleting Dataset %s: %w", a.id.Dataset, err)
		}
	} else {
		if err := dsHandler.Delete(ctx); err != nil {
			return false, fmt.Errorf("deleting Dataset %s: %w", a.id.Dataset, err)
		}
	}
	log.V(2).Info("successfully deleted Dataset", "name", a.id.Dataset)

	return true, nil
}
