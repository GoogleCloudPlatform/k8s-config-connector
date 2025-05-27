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

package bigquerydatatransfer

import (
	"context"
	"fmt"
	"reflect"

	bigquerykrmapi "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/datatransfer/apiv1"
	bigquerydatatransferpb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	serviceDomain = "//bigquerydatatransfer.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.BigQueryDataTransferConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building bigquerydatatransfer client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryDataTransferConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		serviceGeneratedID, err := parseServiceGeneratedIDFromExternalRef(obj)
		if err != nil {
			return nil, err
		}
		resourceID = serviceGeneratedID
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Resolve PubSubTopic Ref
	if obj.Spec.PubSubTopicRef != nil {
		topic, err := refv1beta1.ResolvePubSubTopic(ctx, reader, obj, obj.Spec.PubSubTopicRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.PubSubTopicRef.External = topic.String()
	}

	// Resolve PubSubSubscription Ref
	if obj.Spec.ScheduleOptionsV2 != nil &&
		obj.Spec.ScheduleOptionsV2.EventDrivenSchedule != nil &&
		obj.Spec.ScheduleOptionsV2.EventDrivenSchedule.PubSubSubscriptionRef != nil {
		subscription, err := refv1beta1.ResolvePubSubSubscription(ctx, reader, obj, obj.Spec.ScheduleOptionsV2.EventDrivenSchedule.PubSubSubscriptionRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.ScheduleOptionsV2.EventDrivenSchedule.PubSubSubscriptionRef.External = subscription.String()
	}

	// Resolve BigQueryDataSet Ref
	if obj.Spec.DatasetRef != nil {
		dataset, err := obj.Spec.DatasetRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, err
		}

		// for backwards compatibility and to satisfy the GCP API constraints, we must overrite the
		// external reference in the payloads to just the resource ID of the dataset.
		_, id, err := bigquerykrmapi.ParseDatasetExternal(dataset)
		if err != nil {
			return nil, err
		}
		obj.Spec.DatasetRef.External = id
	}

	// Resolve KMSCryptoKey Ref
	if obj.Spec.EncryptionConfiguration != nil {
		if ref := obj.Spec.EncryptionConfiguration.KMSKeyRef; ref != nil {
			_, err := ref.NormalizedExternal(ctx, reader, obj.GetNamespace())
			if err != nil {
				return nil, err
			}
		}
	}

	// Resolve ServiceAccount Ref
	if obj.Spec.ServiceAccountRef != nil {
		err := obj.Spec.ServiceAccountRef.Resolve(ctx, reader, obj)
		if err != nil {
			return nil, err
		}
	}

	// Resolve Project Ref
	projectRef, err := refv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	var id *BigQueryDataTransferConfigIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.projectID != projectID {
			return nil, fmt.Errorf("BigQueryDataTransferConfig %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.projectID, projectID)
		}
		if id.location != location {
			return nil, fmt.Errorf("BigQueryDataTransferConfig %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, location)
		}
		if id.transferConfigID != resourceID {
			return nil, fmt.Errorf("BigQueryDataTransferConfig %s/%s spec.resourceID changed or does not match the service generated ID, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.transferConfigID, resourceID)
		}
	}

	// Get bigquerydatatransfer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *BigQueryDataTransferConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryDataTransferConfig
	actual    *bigquerydatatransferpb.TransferConfig
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id.transferConfigID == "" { // resource ID is not yet generated by the GCP service
		return false, nil
	}

	log.V(2).Info("getting BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())
	req := &bigquerydatatransferpb.GetTransferConfigRequest{Name: a.id.FullyQualifiedName()}
	transferconfigpb, err := a.gcpClient.GetTransferConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryDataTransferConfig %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = transferconfigpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	projectID := a.id.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}

	desired := a.desired.DeepCopy()
	resource := BigQueryDataTransferConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigquerydatatransferpb.CreateTransferConfigRequest{
		Parent:         a.id.Parent(),
		TransferConfig: resource,
	}
	if desired.Spec.ServiceAccountRef != nil { // special handling for service account field which is not present in GCP proto
		req.ServiceAccountName = desired.Spec.ServiceAccountRef.External
	}
	created, err := a.gcpClient.CreateTransferConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigQueryDataTransferConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())

	status := &krm.BigQueryDataTransferConfigStatus{}
	status.ObservedState = BigQueryDataTransferConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// The UUID is service generated. e.g. "projects/{project_id}/locations/{region}/transferConfigs/{config_id}"
	serviceGeneratedID, err := parseServiceGeneratedIDFromName(created.Name)
	if err != nil {
		return fmt.Errorf("error converting %s to BigQueryDataTransferConfigIdentity: %w", created.Name, err)
	}
	a.id.transferConfigID = serviceGeneratedID
	status.ExternalRef = a.id.AsExternalRef()

	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	// Convert KRM object to proto message
	desiredKRM := a.desired.DeepCopy()
	desired := BigQueryDataTransferConfigSpec_ToProto(mapCtx, &desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	actual := a.actual
	resource := proto.Clone(a.actual).(*bigquerydatatransferpb.TransferConfig) // this is the proto resource we are passing to GCP API update call.

	// Check for immutable fields
	if !reflect.DeepEqual(desired.DataSourceId, a.actual.DataSourceId) {
		return fmt.Errorf("BigQueryDataTransferConfig %s/%s data source ID cannot be changed", u.GetNamespace(), u.GetName())
	}
	if desired.Destination != nil && !reflect.DeepEqual(desired.Destination, a.actual.Destination) {
		return fmt.Errorf("BigQueryDataTransferConfig %s/%s destination dataset cannot be changed", u.GetNamespace(), u.GetName())
	}

	// Find diff
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(desired.DataRefreshWindowDays, actual.DataRefreshWindowDays) {
		resource.DataRefreshWindowDays = desired.DataRefreshWindowDays
		updateMask.Paths = append(updateMask.Paths, "data_refresh_window_days")
	}
	if !reflect.DeepEqual(desired.Disabled, actual.Disabled) {
		resource.Disabled = desired.Disabled
		updateMask.Paths = append(updateMask.Paths, "disabled")
	}
	if !reflect.DeepEqual(desired.DisplayName, actual.DisplayName) {
		resource.DisplayName = desired.DisplayName
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if desired.EmailPreferences != nil && !reflect.DeepEqual(desired.EmailPreferences, actual.EmailPreferences) {
		resource.EmailPreferences = desired.EmailPreferences
		updateMask.Paths = append(updateMask.Paths, "email_preferences")
	}
	if desired.EncryptionConfiguration != nil && !reflect.DeepEqual(desired.EncryptionConfiguration, actual.EncryptionConfiguration) {
		resource.EncryptionConfiguration = desired.EncryptionConfiguration
		updateMask.Paths = append(updateMask.Paths, "encryption_configuration")
	}
	if !reflect.DeepEqual(desired.NotificationPubsubTopic, actual.NotificationPubsubTopic) {
		resource.NotificationPubsubTopic = desired.NotificationPubsubTopic
		updateMask.Paths = append(updateMask.Paths, "notification_pubsub_topic")
	}
	if desired.Params != nil && !reflect.DeepEqual(desired.Params, actual.Params) {
		// TODO: sensitive fields maybe masked by the service, leading to constant diff.
		resource.Params = desired.Params
		updateMask.Paths = append(updateMask.Paths, "params")
	}
	if !reflect.DeepEqual(desired.Schedule, actual.Schedule) {
		resource.Schedule = desired.Schedule
		updateMask.Paths = append(updateMask.Paths, "schedule")
	}
	if desired.ScheduleOptions != nil && !reflect.DeepEqual(desired.ScheduleOptions, actual.ScheduleOptions) {
		resource.ScheduleOptions = desired.ScheduleOptions
		updateMask.Paths = append(updateMask.Paths, "schedule_options")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	resource.Name = a.id.FullyQualifiedName() // need to pass service generated ID to GCP API to identify the GCP resource
	req := &bigquerydatatransferpb.UpdateTransferConfigRequest{
		TransferConfig: resource,
		UpdateMask:     updateMask,
	}
	if a.desired.Spec.ServiceAccountRef != nil { // special handling for service account field which is not present in GCP proto
		req.ServiceAccountName = a.desired.Spec.ServiceAccountRef.External
	}
	updated, err := a.gcpClient.UpdateTransferConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BigQueryDataTransferConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())

	status := &krm.BigQueryDataTransferConfigStatus{}
	status.ObservedState = BigQueryDataTransferConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryDataTransferConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *BigQueryDataTransferConfigSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refv1beta1.ProjectRef{Name: a.id.projectID}
	obj.Spec.Location = a.id.location
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
	log.V(2).Info("deleting BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())

	if a.id.transferConfigID == "" {
		return false, nil
	}
	req := &bigquerydatatransferpb.DeleteTransferConfigRequest{Name: a.id.FullyQualifiedName()}
	err := a.gcpClient.DeleteTransferConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting BigQueryDataTransferConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted BigQueryDataTransferConfig", "name", a.id.FullyQualifiedName())

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
