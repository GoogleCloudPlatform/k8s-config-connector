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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	bigquery "cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
	ghttptransport "google.golang.org/api/transport/http"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	return &Adapter{
		id:         id,
		gcpService: gcpService,
		desired:    obj,
		reader:     reader,
		config:     m.config,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id             *krm.DatasetIdentity
	gcpService     *bigquery.Client
	desired        *krm.BigQueryDataset
	actual         *bigquery.DatasetMetadata
	actualReplicas []krm.DatasetReplicaStatus
	reader         client.Reader
	config         config.ControllerConfig
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

	// Fetch actual replicas if desired replicas are specified
	if len(a.desired.Spec.Replicas) > 0 {
		replicas, err := a.getReplicas(ctx)
		if err != nil {
			return true, fmt.Errorf("getting replicas: %w", err)
		}
		a.actualReplicas = replicas
	}
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {

	log := klog.FromContext(ctx)
	log.V(2).Info("creating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredDataset := BigQueryDatasetSpec_ToProto(mapCtx, &a.desired.Spec)
	desiredDataset.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		desiredDataset.Labels[k] = v
	}
	desiredDataset.Labels["managed-by-cnrm"] = "true"

	// Resolve KMS key reference
	if a.desired.Spec.DefaultEncryptionConfiguration != nil {
		kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, a.desired, a.desired.Spec.DefaultEncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return err
		}
		desiredDataset.DefaultEncryptionConfig.KMSKeyName = kmsRef.External
	}
	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)

	if err := dsHandler.Create(ctx, desiredDataset); err != nil {
		return fmt.Errorf("Error creating Dataset %s: %w", a.id.Dataset, err)
	}
	log.V(2).Info("successfully created Dataset", "name", a.id.Dataset)

	if err := a.patchReplicas(ctx); err != nil {
		return fmt.Errorf("Error setting replicas for Dataset %s: %w", a.id.Dataset, err)
	}

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
	readyCond, err := a.populateReplicasAndCondition(ctx, status)
	if err != nil {
		return fmt.Errorf("error populating replicas and conditions: %w", err)
	}
	if err := createOp.UpdateStatus(ctx, status, readyCond); err != nil {
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
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating Dataset", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Convert KRM object to proto message
	desiredKRM := a.desired.DeepCopy()
	desired := BigQueryDatasetSpec_ToProto(mapCtx, &desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Resolve KMS key reference
	if a.desired.Spec.DefaultEncryptionConfiguration != nil {
		kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, a.desired, a.desired.Spec.DefaultEncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return err
		}
		desired.DefaultEncryptionConfig.KMSKeyName = kmsRef.External
	}

	resource := cloneBigQueryDatasetMetadate(a.actual)
	// Check for immutable fields
	if desiredKRM.Spec.Location != nil && !reflect.DeepEqual(desired.Location, resource.Location) {
		return fmt.Errorf("BigQueryDataset %s/%s location cannot be changed, actual: %s, desired: %s", u.GetNamespace(), u.GetName(), resource.Location, desired.Location)
	}
	// Check for immutable replicas field
	if len(desiredKRM.Spec.Replicas) > 0 {
		desiredLocations := []string{}
		for _, r := range desiredKRM.Spec.Replicas {
			desiredLocations = append(desiredLocations, r.Location)
		}
		sort.Strings(desiredLocations)

		actualLocations := []string{}
		for _, r := range a.actualReplicas {
			if r.Location != nil {
				actualLocations = append(actualLocations, *r.Location)
			}
		}
		sort.Strings(actualLocations)

		if !reflect.DeepEqual(desiredLocations, actualLocations) {
			return fmt.Errorf("BigQueryDataset %s/%s replicas cannot be changed, actual: %v, desired: %v", u.GetNamespace(), u.GetName(), actualLocations, desiredLocations)
		}
	} else if len(a.actualReplicas) > 0 {
		return fmt.Errorf("BigQueryDataset %s/%s replicas cannot be removed, actual replicas exist but spec.replicas is omitted", u.GetNamespace(), u.GetName())
	}
	// Find diff
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	updateMask := &fieldmaskpb.FieldMask{}
	if desired.Description != "" && !reflect.DeepEqual(desired.Description, resource.Description) {
		report.AddField("description", resource.Description, desired.Description)
		resource.Description = desired.Description
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if desired.Name != "" && !reflect.DeepEqual(desired.Name, resource.Name) {
		report.AddField("friendly_name", resource.Name, desired.Name)
		resource.Name = desired.Name
		updateMask.Paths = append(updateMask.Paths, "friendly_name")
	}
	if desired.DefaultPartitionExpiration != 0 && !reflect.DeepEqual(desired.DefaultPartitionExpiration, resource.DefaultPartitionExpiration) {
		report.AddField("default_partition_expiration", resource.DefaultPartitionExpiration, desired.DefaultPartitionExpiration)
		resource.DefaultPartitionExpiration = desired.DefaultPartitionExpiration
		updateMask.Paths = append(updateMask.Paths, "default_partition_expiration")
	}
	if desired.DefaultTableExpiration != 0 && !reflect.DeepEqual(desired.DefaultTableExpiration, resource.DefaultTableExpiration) {
		report.AddField("default_table_expiration", resource.DefaultTableExpiration, desired.DefaultTableExpiration)
		resource.DefaultTableExpiration = desired.DefaultTableExpiration
		updateMask.Paths = append(updateMask.Paths, "default_table_expiration")
	}
	if desired.DefaultCollation != "" && !reflect.DeepEqual(desired.DefaultCollation, resource.DefaultCollation) {
		report.AddField("default_collation", resource.DefaultCollation, desired.DefaultCollation)
		resource.DefaultCollation = desired.DefaultCollation
		updateMask.Paths = append(updateMask.Paths, "default_collation")
	}
	if desired.DefaultEncryptionConfig != nil && resource.DefaultEncryptionConfig != nil && !reflect.DeepEqual(desired.DefaultEncryptionConfig, resource.DefaultEncryptionConfig) {
		report.AddField("default_encryption_configuration", resource.DefaultEncryptionConfig, desired.DefaultEncryptionConfig)
		resource.DefaultEncryptionConfig.KMSKeyName = desired.DefaultEncryptionConfig.KMSKeyName
		updateMask.Paths = append(updateMask.Paths, "default_encryption_configuration")
	}
	if desiredKRM.Spec.IsCaseInsensitive != nil && !reflect.DeepEqual(desired.IsCaseInsensitive, resource.IsCaseInsensitive) {
		report.AddField("is_case_sensitive", resource.IsCaseInsensitive, desired.IsCaseInsensitive)
		resource.IsCaseInsensitive = desired.IsCaseInsensitive
		updateMask.Paths = append(updateMask.Paths, "is_case_sensitive")
	}
	if desired.StorageBillingModel != "" && !reflect.DeepEqual(desired.StorageBillingModel, resource.StorageBillingModel) {
		report.AddField("storage_billing_model", resource.StorageBillingModel, desired.StorageBillingModel)
		resource.StorageBillingModel = desired.StorageBillingModel
		updateMask.Paths = append(updateMask.Paths, "storage_billing_model")
	}
	// If we do not set a value, the GCP service defaults to 168
	// If the existing value is 168, it means that we did not set this field at creation and it defaults to 168.
	// So if the desired value is 0, it means that we do not intend to update this field.
	if desired.MaxTimeTravel != 0 && !reflect.DeepEqual(desired.MaxTimeTravel, resource.MaxTimeTravel) && (resource.MaxTimeTravel != 168 && desired.MaxTimeTravel != 0) {
		report.AddField("max_time_travel", resource.MaxTimeTravel, desired.MaxTimeTravel)
		resource.MaxTimeTravel = desired.MaxTimeTravel
		updateMask.Paths = append(updateMask.Paths, "max_time_travel")
	}
	if desired.Access != nil && resource.Access != nil && len(desired.Access) > 0 && !reflect.DeepEqual(desired.Access, resource.Access) {
		report.AddField("access", resource.Access, desired.Access)
		for _, access := range desired.Access {
			resource.Access = append(resource.Access, access)
		}
	}
	dsHandler := a.gcpService.DatasetInProject(a.id.Project, a.id.Dataset)
	var updated *bigquery.DatasetMetadata
	var err error

	if len(updateMask.Paths) > 0 {
		structuredreporting.ReportDiff(ctx, report)

		// Compute the dataset metadate for update request
		datasetMetadataToUpdate := BigQueryDataset_ToMetadataToUpdate(mapCtx, resource, updateMask.Paths)
		for k, v := range a.desired.GetObjectMeta().GetLabels() {
			datasetMetadataToUpdate.SetLabel(k, v)
		}
		datasetMetadataToUpdate.SetLabel("managed-by-cnrm", "true")
		updated, err = dsHandler.Update(ctx, *datasetMetadataToUpdate, "")
		if err != nil {
			return fmt.Errorf("updating Dataset %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated Dataset", "name", a.id.String())
	} else {
		updated, err = dsHandler.Metadata(ctx)
		if err != nil {
			return fmt.Errorf("getting Dataset metadata %s: %w", a.id.String(), err)
		}
	}

	if err := a.patchReplicas(ctx); err != nil {
		return fmt.Errorf("Error updating replicas for Dataset %s: %w", a.id.Dataset, err)
	}

	status := &krm.BigQueryDatasetStatus{}
	status = BigQueryDatasetStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	readyCond, err := a.populateReplicasAndCondition(ctx, status)
	if err != nil {
		return fmt.Errorf("error populating replicas and conditions: %w", err)
	}
	return updateOp.UpdateStatus(ctx, status, readyCond)
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

func (a *Adapter) getReplicas(ctx context.Context) ([]krm.DatasetReplicaStatus, error) {
	opts, err := a.config.RESTClientOptions()
	if err != nil {
		return nil, fmt.Errorf("getting REST client options: %w", err)
	}
	httpClient, _, err := ghttptransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building HTTP client: %w", err)
	}

	url := fmt.Sprintf("https://bigquery.googleapis.com/bigquery/v2/projects/%s/datasets/%s", a.id.Project, a.id.Dataset)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending GET request to BigQuery: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get dataset raw, status: %s, body: %s", resp.Status, string(respBody))
	}

	var raw struct {
		Replicas []struct {
			ID                              string `json:"id"`
			Location                        string `json:"location"`
			PrimaryState                    string `json:"primaryState"`
			CreationTime                    string `json:"creation_time"`
			CompletionTime                  string `json:"completion_time"`
			PrimaryAssignmentTime           string `json:"primary_assignment_time"`
			PrimaryAssignmentCompletionTime string `json:"primary_assignment_completion_time"`
			SyncStatus                      []struct {
				ReplicationTime string `json:"replication_time"`
			} `json:"sync_status"`
		} `json:"replicas"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decoding dataset response: %w", err)
	}

	var replicas []krm.DatasetReplicaStatus
	for _, r := range raw.Replicas {
		rep := krm.DatasetReplicaStatus{
			ID:                              direct.LazyPtr(r.ID),
			Location:                        direct.LazyPtr(r.Location),
			PrimaryState:                    direct.LazyPtr(r.PrimaryState),
			CreationTime:                    direct.LazyPtr(r.CreationTime),
			CompletionTime:                  direct.LazyPtr(r.CompletionTime),
			PrimaryAssignmentTime:           direct.LazyPtr(r.PrimaryAssignmentTime),
			PrimaryAssignmentCompletionTime: direct.LazyPtr(r.PrimaryAssignmentCompletionTime),
		}
		for _, s := range r.SyncStatus {
			rep.SyncStatus = append(rep.SyncStatus, krm.DatasetReplicaSyncStatus{
				ReplicationTime: direct.LazyPtr(s.ReplicationTime),
			})
		}
		replicas = append(replicas, rep)
	}
	return replicas, nil
}

func (a *Adapter) patchReplicas(ctx context.Context) error {
	if len(a.desired.Spec.Replicas) == 0 {
		return nil
	}
	opts, err := a.config.RESTClientOptions()
	if err != nil {
		return fmt.Errorf("getting REST client options: %w", err)
	}
	httpClient, _, err := ghttptransport.NewClient(ctx, opts...)
	if err != nil {
		return fmt.Errorf("building HTTP client: %w", err)
	}

	url := fmt.Sprintf("https://bigquery.googleapis.com/bigquery/v2/projects/%s/datasets/%s", a.id.Project, a.id.Dataset)

	type replicaPayload struct {
		Location string `json:"location"`
	}
	type patchPayload struct {
		Replicas []replicaPayload `json:"replicas"`
	}

	payload := patchPayload{}
	for _, r := range a.desired.Spec.Replicas {
		payload.Replicas = append(payload.Replicas, replicaPayload{Location: r.Location})
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling patch payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending PATCH request to BigQuery for replicas: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to patch replicas, status: %s, body: %s", resp.Status, string(respBody))
	}
	return nil
}

func (a *Adapter) populateReplicasAndCondition(ctx context.Context, status *krm.BigQueryDatasetStatus) (*v1alpha1.Condition, error) {
	if len(a.desired.Spec.Replicas) == 0 {
		return nil, nil
	}
	replicas, err := a.getReplicas(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting replicas: %w", err)
	}

	if len(replicas) > 0 {
		if status.ObservedState == nil {
			status.ObservedState = &krm.BigQueryDatasetObservedState{}
		}
		status.ObservedState.Replicas = replicas

		primaryLocation := ""
		for _, r := range replicas {
			if direct.ValueOf(r.PrimaryState) == "PRIMARY" {
				primaryLocation = direct.ValueOf(r.Location)
				break
			}
		}
		if primaryLocation != "" {
			status.PrimaryLocation = direct.LazyPtr(primaryLocation)

			desiredPrimaryLocation := direct.ValueOf(a.desired.Spec.Location)
			if desiredPrimaryLocation != "" && desiredPrimaryLocation != primaryLocation {
				return &v1alpha1.Condition{
					Type:               "Ready",
					Status:             "False",
					Reason:             "PrimaryLocationDrifted",
					Message:            fmt.Sprintf("Primary location has drifted from the desired location %q to %q", desiredPrimaryLocation, primaryLocation),
					LastTransitionTime: metav1.Now().UTC().Format(time.RFC3339),
				}, nil
			}
		}
	}
	return nil, nil
}
