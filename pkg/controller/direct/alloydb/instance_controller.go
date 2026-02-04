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

package alloydb

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	gcp "cloud.google.com/go/alloydb/apiv1beta"
	alloydbpb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AlloyDBInstanceGVK, NewInstanceModel)
}

func NewInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &instanceModel{config: *config}, nil
}

var _ directbase.Model = &instanceModel{}

type instanceModel struct {
	config config.ControllerConfig
}

func (m *instanceModel) client(ctx context.Context) (*gcp.AlloyDBAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAlloyDBAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building AlloyDB client for Instance: %w", err)
	}
	return gcpClient, err
}

func resolveInstanceType(ctx context.Context, reader client.Reader, obj *krm.AlloyDBInstance, isDeletion bool) error {
	if obj.Spec.InstanceType == nil && obj.Spec.InstanceTypeRef == nil {
		return fmt.Errorf("at least one of 'spec.InstanceTypeRef' " +
			"and 'spec.InstanceType' should be configured: neither is configured")
	}

	var instanceType *string
	if obj.Spec.InstanceType != nil {
		instanceType = obj.Spec.InstanceType
	}
	if obj.Spec.InstanceTypeRef != nil {
		plainTextInstanceType := instanceType
		var err error
		instanceType, err = refsv1beta1.ResolveAlloyDBClusterType(ctx, reader, obj, obj.Spec.InstanceTypeRef)
		if err != nil {
			if !isDeletion {
				// TODO: Read instance type from observed state because it's necessary during deletion.
				return fmt.Errorf("cannot resolve `spec.InstanceTypeRef`: %w", err)
			}
		}
		if plainTextInstanceType != nil && *plainTextInstanceType != *instanceType {
			return fmt.Errorf("'spec.InstanceTypeRef' and 'spec.InstanceType' "+
				"resolve into different values: spec.InstanceTypeRef resolves to "+
				"instanceType %q, spec.InstanceType is %q: they must be the same",
				*instanceType, *plainTextInstanceType)
		}
	}
	if *instanceType == "" {
		return fmt.Errorf("instance type should be a non-empty string, " +
			"ensure the referenced AlloyDBCluster in `spec.instancerTypeRef` " +
			"has the correct cluster type or `spec.instancerTypeRef.external` " +
			"has a correct instance type")
	}
	if obj.Spec.InstanceTypeRef == nil {
		obj.Spec.InstanceTypeRef = &refsv1beta1.AlloyDBClusterTypeRef{}
	}
	obj.Spec.InstanceTypeRef.External = *instanceType
	return nil
}

func validateConfig(spec *krm.AlloyDBInstanceSpec) error {
	if spec.ObservabilityInstanceConfig != nil && spec.QueryInsightsInstanceConfig != nil {
		return fmt.Errorf("cannot specify both observabilityConfig and queryInsightsConfig; please use only one")
	}
	return nil
}

func (m *instanceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AlloyDBInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get alloydb GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &instanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		reader:    reader,
		desired:   obj,
	}, nil
}

func (m *instanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type instanceAdapter struct {
	id        *krm.InstanceIdentity
	gcpClient *gcp.AlloyDBAdminClient
	reader    client.Reader
	desired   *krm.AlloyDBInstance
	actual    *alloydbpb.Instance
}

var _ directbase.Adapter = &instanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting instance", "name", a.id)

	req := &alloydbpb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		log.V(2).Info("error getting instance", "name", a.id, "error", err)
		if direct.IsNotFound(err) {
			return false, nil
		}

		return false, fmt.Errorf("getting instance %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := resolveInstanceType(ctx, a.reader, a.desired, false); err != nil {
		return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, err)
	}

	if err := validateConfig(&a.desired.Spec); err != nil {
		return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, err)
	}

	desired := a.desired.DeepCopy()
	resource := AlloyDBInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, mapCtx.Err())
	}
	resource.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	resource.Labels["managed-by-cnrm"] = "true"

	var created *alloydbpb.Instance
	instanceType := a.desired.Spec.InstanceTypeRef.External
	if instanceType == "SECONDARY" {
		req := &alloydbpb.CreateSecondaryInstanceRequest{
			Parent:     a.id.Parent().String(),
			InstanceId: a.id.ID(),
			Instance:   resource,
		}
		op, err := a.gcpClient.CreateSecondaryInstance(ctx, req)
		if err != nil {
			log.V(2).Info("error creating secondary instance", "name", a.id, "error", err)
			return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("creating secondary instance %s: %w", a.id, err))
		}
		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting secondary instance creation", "name", a.id, "error", err)
			return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("secondary instance %s waiting creation: %w", a.id, err))
		}
		log.V(2).Info("successfully created secondary instance", "name", a.id)
	} else {
		req := &alloydbpb.CreateInstanceRequest{
			Parent:     a.id.Parent().String(),
			InstanceId: a.id.ID(),
			Instance:   resource,
		}
		op, err := a.gcpClient.CreateInstance(ctx, req)
		if err != nil {
			log.V(2).Info("error creating instance", "name", a.id, "error", err)
			return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("creating instance %s: %w", a.id, err))
		}
		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting instance creation", "name", a.id, "error", err)
			return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("instance %s waiting creation: %w", a.id, err))
		}
		log.V(2).Info("successfully created instance", "name", a.id)
	}

	status := AlloyDBInstanceStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return createOp.RecordCreateError(ctx, &krm.AlloyDBInstanceStatus{}, mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := resolveInstanceType(ctx, a.reader, a.desired, false); err != nil {
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, err)
	}

	if err := validateConfig(&a.desired.Spec); err != nil {
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, err)
	}

	parsedActual := AlloyDBInstanceSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, mapCtx.Err())
	}

	// TODO: Change to CompareProtoMessage once we support all the files in the instance pb.
	updatePaths, err := compareInstance(ctx, parsedActual, &a.desired.Spec)
	if err != nil {
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, err)
	}
	desiredLabels := label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	if desiredLabels == nil {
		desiredLabels = make(map[string]string)
	}
	desiredLabels["managed-by-cnrm"] = "true"
	if !reflect.DeepEqual(a.actual.GetLabels(), desiredLabels) {
		log.V(2).Info("'metadata.labels' field is updated (-old +new)", cmp.Diff(a.actual.GetLabels(), desiredLabels))
		updatePaths = append(updatePaths, "labels")
	}

	if len(updatePaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			// If it is the first reconciliation after switching to direct controller,
			// or is an acquisition, then update Status to fill out the ExternalRef
			// and ObservedState.
			status := AlloyDBInstanceStatus_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, mapCtx.Err())
			}
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: updatePaths,
	}
	desiredPb := AlloyDBInstanceSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	desiredPb.Labels = desiredLabels
	desiredPb.Name = a.id.String()
	req := &alloydbpb.UpdateInstanceRequest{
		UpdateMask: updateMask,
		Instance:   desiredPb,
	}
	op, err := a.gcpClient.UpdateInstance(ctx, req)
	if err != nil {
		log.V(2).Info("error updating instance", "name", a.id, "error", err)
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("updating instance %s: %w", a.id, err))
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting instance update", "name", a.id, "error", err)
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, fmt.Errorf("instance %s waiting update: %w", a.id, err))
	}
	log.V(2).Info("successfully updated instance", "name", a.id)

	status := AlloyDBInstanceStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return updateOp.RecordUpdateError(ctx, &krm.AlloyDBInstanceStatus{}, mapCtx.Err())
	}
	if a.desired.Status.ExternalRef == nil {
		// If it is the first reconciliation after switching to direct controller,
		// or is an acquisition with updates, then fill out the ExternalRef.
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// HELPER FUNCTION STARTS

// Helper function to compare boolean pointers
func boolPtrChanged(actual, desired *bool, basePath string, log klog.Logger) (string, bool) {
	if desired == nil {
		return "", false // Desired is not set, so no update for this field
	}
	if actual == nil || *actual != *desired {
		log.V(2).Info(fmt.Sprintf("'%s' field is updated (-old +new)", basePath), "diff", cmp.Diff(actual, desired))
		return basePath, true
	}
	return "", false
}

// Helper function to compare int32 pointers
func int32PtrChanged(actual, desired *int32, basePath string, log klog.Logger) (string, bool) {
	if desired == nil {
		return "", false // Desired is not set, so no update for this field
	}
	if actual == nil || *actual != *desired {
		log.V(2).Info(fmt.Sprintf("'%s' field is updated (-old +new)", basePath), "diff", cmp.Diff(actual, desired))
		return basePath, true
	}
	return "", false
}

func unsignedInt32PtrChanged(actual, desired *uint32, basePath string, log klog.Logger) (string, bool) {
	if desired == nil {
		return "", false
	}
	if actual == nil || *actual != *desired {
		log.V(2).Info(fmt.Sprintf("'%s' field is updated (-old +new)", basePath), "diff", cmp.Diff(actual, desired))
		return basePath, true
	}
	return "", false
}

// Helper function to compare string pointers
func stringPtrChanged(actual, desired *string, basePath string, log klog.Logger) (string, bool) {
	if desired == nil {
		return "", false // Desired is not set, so no update for this field
	}
	if actual == nil || *actual != *desired {
		log.V(2).Info(fmt.Sprintf("'%s' field is updated (-old +new)", basePath), "diff", cmp.Diff(actual, desired))
		return basePath, true
	}
	return "", false
}

// HELPER FUNCTION ENDS

func compareInstance(ctx context.Context, actual, desired *krm.AlloyDBInstanceSpec) (updatePaths []string, err error) {
	log := klog.FromContext(ctx)
	updatePaths = make([]string, 0)
	if !reflect.DeepEqual(actual.Annotations, desired.Annotations) {
		log.V(2).Info("'spec.annotations' field is updated (-old +new)", cmp.Diff(actual.Annotations, desired.Annotations))
		updatePaths = append(updatePaths, "annotations")
	}
	// TODO: Test case with availability type unset.
	if desired.AvailabilityType != nil && !reflect.DeepEqual(actual.AvailabilityType, desired.AvailabilityType) {
		log.V(2).Info("'spec.availabilityType' field is updated (-old +new)", cmp.Diff(actual.AvailabilityType, desired.AvailabilityType))
		updatePaths = append(updatePaths, "availability_type")
	}
	if desired.ConnectionPoolConfig != nil {
		if desired.ConnectionPoolConfig.Enabled != nil && !reflect.DeepEqual(actual.ConnectionPoolConfig.Enabled, desired.ConnectionPoolConfig.Enabled) {
			log.V(2).Info("'spec.connectionPoolConfig.enabled' field is updated (-old +new)", cmp.Diff(actual.ConnectionPoolConfig.Enabled, desired.ConnectionPoolConfig.Enabled))
			updatePaths = append(updatePaths, "connection_pool_config.enabled")
		}
		if desired.ConnectionPoolConfig.Flags != nil && !reflect.DeepEqual(actual.ConnectionPoolConfig.Flags, desired.ConnectionPoolConfig.Flags) {
			log.V(2).Info("'spec.connectionPoolConfig.flags' field is updated (-old +new)", cmp.Diff(actual.ConnectionPoolConfig.Flags, desired.ConnectionPoolConfig.Flags))
			updatePaths = append(updatePaths, "connection_pool_config.flags")
		}
	}
	// TODO: Test "copied" behavior for read pool
	// TODO: Test "overridden" behavior for read pool
	// Default value of databaseFlags is unknown for a read instance unless we
	// make API calls to get the database flags of the primary instance.
	if desired.DatabaseFlags != nil && !reflect.DeepEqual(actual.DatabaseFlags, desired.DatabaseFlags) {
		log.V(2).Info("'spec.databaseFlags' field is updated (-old +new)", cmp.Diff(actual.DatabaseFlags, desired.DatabaseFlags))
		updatePaths = append(updatePaths, "database_flags")
	}
	if desired.DisplayName != nil && !reflect.DeepEqual(actual.DisplayName, desired.DisplayName) {
		log.V(2).Info("'spec.displayName' field is updated (-old +new)", cmp.Diff(actual.DisplayName, desired.DisplayName))
		updatePaths = append(updatePaths, "display_name")
	}
	if desired.GCEZone != nil && !reflect.DeepEqual(actual.GCEZone, desired.GCEZone) {
		log.V(2).Info("'spec.gceZone' field is updated (-old +new)", cmp.Diff(actual.GCEZone, desired.GCEZone))
		updatePaths = append(updatePaths, "gce_zone")
	}
	if desired.InstanceTypeRef != nil && !reflect.DeepEqual(actual.InstanceTypeRef.External, desired.InstanceTypeRef.External) {
		log.V(2).Info("'spec.instanceTypeRef' field is updated (-old +new)", cmp.Diff(actual.InstanceTypeRef.External, desired.InstanceTypeRef.External))
		return nil, fmt.Errorf("cannot change immutable field %s from %v to %v", "'spec.instanceTypeRef'", actual.InstanceTypeRef.External, desired.InstanceTypeRef.External)
	}
	// TODO: Test machineConfig unset and empty struct
	if desired.MachineConfig != nil {
		if desired.MachineConfig.CPUCount != nil && !reflect.DeepEqual(actual.MachineConfig.CPUCount, desired.MachineConfig.CPUCount) {
			log.V(2).Info("'spec.machineConfig.cpuCount' field is updated (-old +new)", cmp.Diff(actual.MachineConfig.CPUCount, desired.MachineConfig.CPUCount))
			updatePaths = append(updatePaths, "machine_config.cpu_count")
		}
		if desired.MachineConfig.MachineType != nil && !reflect.DeepEqual(actual.MachineConfig.MachineType, desired.MachineConfig.MachineType) {
			log.V(2).Info("'spec.machineConfig.machineType' field is updated (-old +new)", cmp.Diff(actual.MachineConfig.MachineType, desired.MachineConfig.MachineType))
			updatePaths = append(updatePaths, "machine_config.machine_type")
		}
	}
	if desired.NetworkConfig != nil {
		if desired.NetworkConfig.EnablePublicIP != nil && !reflect.DeepEqual(actual.NetworkConfig.EnablePublicIP, desired.NetworkConfig.EnablePublicIP) {
			log.V(2).Info("'spec.networkConfig.enablePublicIp' field is updated (-old +new)", cmp.Diff(actual.NetworkConfig.EnablePublicIP, desired.NetworkConfig.EnablePublicIP))
			updatePaths = append(updatePaths, "network_config.enable_public_ip")
		}
		if desired.NetworkConfig.EnableOutboundPublicIP != nil && !reflect.DeepEqual(actual.NetworkConfig.EnableOutboundPublicIP, desired.NetworkConfig.EnableOutboundPublicIP) {
			log.V(2).Info("'spec.networkConfig.enableOutboundPublicIp' field is updated (-old +new)", cmp.Diff(actual.NetworkConfig.EnableOutboundPublicIP, desired.NetworkConfig.EnableOutboundPublicIP))
			updatePaths = append(updatePaths, "network_config.enable_outbound_public_ip")
		}
		if desired.NetworkConfig.AuthorizedExternalNetworks != nil && !reflect.DeepEqual(actual.NetworkConfig.AuthorizedExternalNetworks, desired.NetworkConfig.AuthorizedExternalNetworks) {
			log.V(2).Info("'spec.networkConfig.authorizedExternalNetworks' field is updated (-old +new)", cmp.Diff(actual.NetworkConfig.AuthorizedExternalNetworks, desired.NetworkConfig.AuthorizedExternalNetworks))
			updatePaths = append(updatePaths, "network_config.authorized_external_networks")
		}
	}
	if desired.ReadPoolConfig != nil {
		if desired.ReadPoolConfig.NodeCount != nil && !reflect.DeepEqual(actual.ReadPoolConfig.NodeCount, desired.ReadPoolConfig.NodeCount) {
			log.V(2).Info("'spec.readPoolConfig.nodeCount' field is updated (-old +new)", cmp.Diff(actual.ReadPoolConfig.NodeCount, desired.ReadPoolConfig.NodeCount))
			updatePaths = append(updatePaths, "read_pool_config.node_count")
		}
	}

	// TODO clean the above as per below in later pr
	if desired.ObservabilityInstanceConfig != nil {
		actualObs := actual.ObservabilityInstanceConfig
		if actualObs == nil {
			actualObs = &krm.Instance_ObservabilityInstanceConfig{}
		}
		desiredObs := desired.ObservabilityInstanceConfig

		if path, changed := boolPtrChanged(actualObs.Enabled, desiredObs.Enabled, "observability_config.enabled", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := boolPtrChanged(actualObs.PreserveComments, desiredObs.PreserveComments, "observability_config.preserve_comments", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := boolPtrChanged(actualObs.TrackWaitEvents, desiredObs.TrackWaitEvents, "observability_config.track_wait_events", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := int32PtrChanged(actualObs.MaxQueryStringLength, desiredObs.MaxQueryStringLength, "observability_config.max_query_string_length", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := boolPtrChanged(actualObs.RecordApplicationTags, desiredObs.RecordApplicationTags, "observability_config.record_application_tags", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := int32PtrChanged(actualObs.QueryPlansPerMinute, desiredObs.QueryPlansPerMinute, "observability_config.query_plans_per_minute", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := boolPtrChanged(actualObs.TrackActiveQueries, desiredObs.TrackActiveQueries, "observability_config.track_active_queries", log); changed {
			updatePaths = append(updatePaths, path)
		}
		// TrackClientAcddress is in v1beta but is not in v1
		if path, changed := boolPtrChanged(actualObs.TrackClientAddress, desiredObs.TrackClientAddress, "observability_config.track_client_address", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := boolPtrChanged(actualObs.TrackClientAddress, desiredObs.TrackClientAddress, "observability_config.assistive_experiences_enabled", log); changed {
			updatePaths = append(updatePaths, path)
		}
	}

	if desired.QueryInsightsInstanceConfig != nil {
		actualInsights := actual.QueryInsightsInstanceConfig
		if actualInsights == nil {
			actualInsights = &krm.Instance_QueryInsightsInstanceConfig{}
		}

		desiredInsights := desired.QueryInsightsInstanceConfig

		if path, changed := boolPtrChanged(actualInsights.RecordApplicationTags, desiredInsights.RecordApplicationTags, "query_insights_config.record_application_tags", log); changed {
			updatePaths = append(updatePaths, path)
		}

		if path, changed := boolPtrChanged(actualInsights.RecordClientAddress, desiredInsights.RecordClientAddress, "query_insights_config.record_client_address", log); changed {
			updatePaths = append(updatePaths, path)
		}

		if path, changed := unsignedInt32PtrChanged(actualInsights.QueryStringLength, desiredInsights.QueryStringLength, "query_insights_config.query_string_length", log); changed {
			updatePaths = append(updatePaths, path)
		}
		if path, changed := unsignedInt32PtrChanged(actualInsights.QueryPlansPerMinute, desiredInsights.QueryPlansPerMinute, "query_insights_config.query_plans_per_minute", log); changed {
			updatePaths = append(updatePaths, path)
		}
	}

	return updatePaths, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *instanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AlloyDBInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AlloyDBInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Split name into tokens and use ID.
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.AlloyDBInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *instanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting instance", "name", a.id)

	// instanceType must be resolved before calling DELETE.
	if err := resolveInstanceType(ctx, a.reader, a.desired, true); err != nil {
		return false, err
	}

	// Returning true directly if it is to delete a secondary instance.
	// This is because deletion of secondary instance is not supported. Instead,
	// users should delete the secondary cluster which will forcefully delete
	// the associated secondary instance.
	instanceType := a.desired.Spec.InstanceTypeRef.External
	if instanceType == "SECONDARY" {
		log.V(2).Info("This operation didn't delete the secondary instance. You need to delete the associated secondary cluster to delete the secondary instance (and the entire secondary cluster).", "name", a.id)
		return true, nil
	}

	req := &alloydbpb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		log.V(2).Info("error deleting instance", "name", a.id, "error", err)
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting instance %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		// "(line 15:3): missing \"value\" field" is likely a protolib error but
		// not a real server side error. When it happens, it usually means the
		// deletion has been completed successfully but the returned empty
		// struct can't be parsed by the protolib.
		if !strings.Contains(err.Error(), "(line 15:3): missing \"value\" field") {
			log.V(2).Info("error waiting instance delete", "name", a.id, "error", err)
			return false, fmt.Errorf("waiting delete instance %s: %w", a.id, err)
		}
	}

	log.V(2).Info("successfully deleted instance", "name", a.id)
	return true, nil
}
