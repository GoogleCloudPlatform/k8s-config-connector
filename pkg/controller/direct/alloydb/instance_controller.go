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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

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

func (m *instanceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AlloyDBInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Spec.InstanceType != nil && obj.Spec.InstanceTypeRef != nil {
		return nil, fmt.Errorf("one and only one of 'spec.InstanceTypeRef' " +
			"and 'spec.InstanceType' should be configured: both are configured")
	}
	if obj.Spec.InstanceType == nil && obj.Spec.InstanceTypeRef == nil {
		return nil, fmt.Errorf("one and only one of 'spec.InstanceTypeRef' " +
			"and 'spec.InstanceType' should be configured: neither is configured")
	}

	var instanceType *string
	if obj.Spec.InstanceType != nil {
		if *instanceType == "" {
			return nil, fmt.Errorf("'spec.InstanceType' should be configured with a non-empty string")
		}
		instanceType = obj.Spec.InstanceType
	}
	if obj.Spec.InstanceTypeRef != nil {
		instanceType, err = refsv1beta1.ResolveAlloyDBClusterType(ctx, reader, obj, obj.Spec.InstanceTypeRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve `spec.InstanceTypeRef`: %w", err)
		}
	}
	obj.Spec.InstanceType = instanceType

	// Get alloydb GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &instanceAdapter{
		id:        id,
		gcpClient: gcpClient,
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
	desired   *krm.AlloyDBInstance
	actual    *alloydbpb.Instance
}

var _ directbase.Adapter = &instanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return true means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting instance", "name", a.id)
	fmt.Printf("getting instance: %v\n", a.id)

	req := &alloydbpb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		log.V(2).Info("error getting instance", "name", a.id, "error", err)
		fmt.Printf("instance error: %+v\n", err)
		if direct.IsNotFound(err) {
			return false, nil
		}

		return false, fmt.Errorf("getting instance %q: %w", a.id, err)
	}

	fmt.Printf("retrieved instance: %+v\n", instancepb)
	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating instance", "name", a.id)
	fmt.Printf("creating instance: %v\n", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AlloyDBInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	var created *alloydbpb.Instance
	instanceType := a.desired.Spec.InstanceType
	if instanceType != nil && *instanceType == "SECONDARY" {
		req := &alloydbpb.CreateSecondaryInstanceRequest{
			Parent:     a.id.Parent().String(),
			InstanceId: a.id.ID(),
			Instance:   resource,
		}
		op, err := a.gcpClient.CreateSecondaryInstance(ctx, req)
		if err != nil {
			log.V(2).Info("error creating secondary instance", "name", a.id, "error", err)
			return fmt.Errorf("creating secondary instance %s: %w", a.id, err)
		}
		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting secondary instance creation", "name", a.id, "error", err)
			return fmt.Errorf("secondary instance %s waiting creation: %w", a.id, err)
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
			return fmt.Errorf("creating instance %s: %w", a.id, err)
		}
		created, err = op.Wait(ctx)
		if err != nil {
			log.V(2).Info("error waiting instance creation", "name", a.id, "error", err)
			return fmt.Errorf("instance %s waiting creation: %w", a.id, err)
		}
		log.V(2).Info("successfully created instance", "name", a.id)
	}

	status := AlloyDBInstanceStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	parsedActual := AlloyDBInstanceSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updatePaths, err := compareInstance(ctx, parsedActual, &a.desired.Spec)
	if err != nil {
		return err
	}
	desiredLabels := a.desired.GetObjectMeta().GetLabels()
	desiredLabels["managed-by-cnrm"] = "true"
	if !reflect.DeepEqual(a.actual.GetLabels(), desiredLabels) {
		log.V(2).Info("'metadata.labels' field is updated (-old +new)", cmp.Diff(a.actual.GetLabels(), desiredLabels))
		updatePaths = append(updatePaths, "availability_type")
	}

	if len(updatePaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	fmt.Printf("maqiuyu... updateMasks: %+v\n", updatePaths)
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
		return fmt.Errorf("updating instance %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting instance update", "name", a.id, "error", err)
		return fmt.Errorf("instance %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated instance", "name", a.id)

	status := AlloyDBInstanceStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

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
	if desired.GceZone != nil && !reflect.DeepEqual(actual.GceZone, desired.GceZone) {
		log.V(2).Info("'spec.gceZone' field is updated (-old +new)", cmp.Diff(actual.GceZone, desired.GceZone))
		updatePaths = append(updatePaths, "gce_zone")
	}
	if desired.InstanceType != nil && !reflect.DeepEqual(actual.InstanceType, desired.InstanceType) {
		log.V(2).Info("'spec.instanceType' field is updated (-old +new)", cmp.Diff(actual.InstanceType, desired.InstanceType))
		return nil, fmt.Errorf("cannot change immutable field %s from %v to %v", "'spec.instanceType'", actual.InstanceType, desired.InstanceType)
	}
	// TODO: Test machineConfig unset and empty struct
	if desired.MachineConfig != nil {
		if desired.MachineConfig.CpuCount != nil && !reflect.DeepEqual(actual.MachineConfig.CpuCount, desired.MachineConfig.CpuCount) {
			log.V(2).Info("'spec.machineConfig.cpuCount' field is updated (-old +new)", cmp.Diff(actual.MachineConfig.CpuCount, desired.MachineConfig.CpuCount))
			updatePaths = append(updatePaths, "machine_config.cpu_count")
		}
	}
	if desired.NetworkConfig != nil {
		if desired.NetworkConfig.EnablePublicIp != nil && !reflect.DeepEqual(actual.NetworkConfig.EnablePublicIp, desired.NetworkConfig.EnablePublicIp) {
			log.V(2).Info("'spec.networkConfig.enablePublicIp' field is updated (-old +new)", cmp.Diff(actual.NetworkConfig.EnablePublicIp, desired.NetworkConfig.EnablePublicIp))
			updatePaths = append(updatePaths, "network_config.enable_public_ip")
		}
		if desired.NetworkConfig.EnableOutboundPublicIp != nil && !reflect.DeepEqual(actual.NetworkConfig.EnableOutboundPublicIp, desired.NetworkConfig.EnableOutboundPublicIp) {
			log.V(2).Info("'spec.networkConfig.enableOutboundPublicIp' field is updated (-old +new)", cmp.Diff(actual.NetworkConfig.EnableOutboundPublicIp, desired.NetworkConfig.EnableOutboundPublicIp))
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

	// Returning true directly if it is to delete a secondary instance.
	// Technically the secondary instance is only abandoned but not deleted.
	// This is because deletion of secondary instance is not supported. Instead,
	// users should delete the secondary cluster which will forcefully delete
	// the associated secondary instance.
	instanceType := a.desired.Spec.InstanceType
	if instanceType != nil && *instanceType == "SECONDARY" {
		log.V(2).Info("This operation didn't delete the secondary instance. You need to delete the associated secondary cluster to delete the secondary instance (and the entire secondary cluster).", "name", a.id)
		return true, nil
	}

	req := &alloydbpb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if op != nil {
		opMetadata, opErr := op.Metadata()
		fmt.Printf("maqiuyu... delete operation: %v\n%v\nMetadata:\n%+v\nErr while getting metadata\n%v\n", op.Name(), op.Done(), opMetadata, opErr)
	} else {
		fmt.Printf("maqiuyu... delete operation not triggered. Maybe there is an error? %+v\n", err)
	}
	if err != nil {
		log.V(2).Info("error deleting instance", "name", a.id, "error", err)
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting instance %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting instance delete", "name", a.id, "error", err)
		return false, fmt.Errorf("waiting delete instance %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted instance", "name", a.id)
	return true, nil
}
