// Copyright 2025 Google LLC
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

package memorystore

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1beta1"
	refsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/memorystore/apiv1"

	memorystorepb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MemorystoreInstanceGVK, NewInstanceModel)
}

func NewInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstance{config: *config}, nil
}

var _ directbase.Model = &modelInstance{}

type modelInstance struct {
	config config.ControllerConfig
}

func (m *modelInstance) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Instance client: %w", err)
	}
	return gcpClient, err
}

func (m *modelInstance) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MemorystoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get memorystore GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		reader:    reader,
		desired:   obj,
	}, nil
}

func resolveReferences(ctx context.Context, reader client.Reader, obj *krm.MemorystoreInstance) error {
	for _, endpoint := range obj.Spec.Endpoints {
		for _, connection := range endpoint.Connections {
			if connection.PscAutoConnection != nil {
				autoConnection := connection.PscAutoConnection
				if autoConnection.NetworkRef != nil {
					if err := autoConnection.NetworkRef.Normalize(ctx, reader, obj.Namespace); err != nil {
						return err
					}
				}
				if autoConnection.ProjectRef != nil {
					if err := autoConnection.ProjectRef.Normalize(ctx, reader, obj.Namespace); err != nil {
						return err
					}
				}
			}
		}
	}
	if obj.Spec.CrossInstanceReplicationConfig != nil {
		crr := obj.Spec.CrossInstanceReplicationConfig
		if err := resolveRemoteInstanceRef(ctx, reader, obj, crr.PrimaryInstance); err != nil {
			return err
		}
		for _, secondaryInstance := range crr.SecondaryInstances {
			if err := resolveRemoteInstanceRef(ctx, reader, obj, &secondaryInstance); err != nil {
				return err
			}
		}
	}
	if obj.Spec.ManagedBackupSource != nil && obj.Spec.ManagedBackupSource.BackupRef != nil {
		ref, err := refsv1alpha1.ResolveMemorystoreInstanceBackup(ctx, reader, obj, obj.Spec.ManagedBackupSource.BackupRef)
		if err != nil {
			return err
		}
		obj.Spec.ManagedBackupSource.BackupRef.External = ref.External
	}
	return nil
}

func resolveRemoteInstanceRef(ctx context.Context, reader client.Reader, obj *krm.MemorystoreInstance, remoteInstance *krm.CrossInstanceReplicationConfig_RemoteInstance) error {
	if remoteInstance == nil {
		return nil
	}
	if remoteInstance.InstanceRef == nil {
		return fmt.Errorf("InstanceRef is nil")
	}
	ref, err := refs.ResolveMemorystoreInstance(ctx, reader, obj, remoteInstance.InstanceRef)
	if err != nil {
		return err
	}
	remoteInstance.InstanceRef.External = ref.External
	return nil
}

func (m *modelInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type InstanceAdapter struct {
	id        *krm.InstanceIdentity
	gcpClient *gcp.Client
	reader    client.Reader
	desired   *krm.MemorystoreInstance
	actual    *memorystorepb.Instance
}

var _ directbase.Adapter = &InstanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Instance", "name", a.id)

	req := &memorystorepb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Instance %q: %w", a.id, err)
	}

	a.actual = instancepb
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Instance", "name", a.id)

	desired := a.desired.DeepCopy()
	if err := resolveReferences(ctx, a.reader, desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	resource := MemorystoreInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &memorystorepb.CreateInstanceRequest{
		Parent:     a.id.Parent().String(),
		Instance:   resource,
		InstanceId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Instance %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Instance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Instance", "name", a.id)

	status := &krm.MemorystoreInstanceStatus{}
	status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Instance", "name", a.id)

	desired := a.desired.DeepCopy()
	if err := resolveReferences(ctx, a.reader, desired); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := MemorystoreInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Mask out some fields in the persistenceConfig, to accommodate KRM "discriminated union" semantics
	if desiredPb.PersistenceConfig != nil {
		switch desiredPb.GetPersistenceConfig().GetMode() {
		case memorystorepb.PersistenceConfig_DISABLED:
			desiredPb.PersistenceConfig.AofConfig = nil
			desiredPb.PersistenceConfig.RdbConfig = nil
		case memorystorepb.PersistenceConfig_RDB:
			desiredPb.PersistenceConfig.AofConfig = nil
		case memorystorepb.PersistenceConfig_AOF:
			desiredPb.PersistenceConfig.RdbConfig = nil
		}
	}

	// Need to unset the fields to allow for switchover in cross region replication.
	if desiredPb.CrossInstanceReplicationConfig != nil {
		switch desiredPb.GetCrossInstanceReplicationConfig().GetInstanceRole() {
		case memorystorepb.CrossInstanceReplicationConfig_PRIMARY:
			desiredPb.CrossInstanceReplicationConfig.PrimaryInstance =
				&memorystorepb.CrossInstanceReplicationConfig_RemoteInstance{Instance: ""}
		case memorystorepb.CrossInstanceReplicationConfig_SECONDARY:
			desiredPb.CrossInstanceReplicationConfig.SecondaryInstances = nil
		default:
			desiredPb.CrossInstanceReplicationConfig.PrimaryInstance =
				&memorystorepb.CrossInstanceReplicationConfig_RemoteInstance{Instance: ""}
			desiredPb.CrossInstanceReplicationConfig.SecondaryInstances = nil
		}
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	var paths []string

	// If a field is unset, the field become unmanaged.
	// see https://docs.cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields
	if a.desired.Spec.Labels != nil && !reflect.DeepEqual(desiredPb.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, desiredPb.Labels)
		paths = append(paths, "labels")
	}
	if a.desired.Spec.ReplicaCount != nil && !reflect.DeepEqual(desiredPb.ReplicaCount, a.actual.ReplicaCount) {
		report.AddField("replica_count", a.actual.ReplicaCount, desiredPb.ReplicaCount)
		paths = append(paths, "replica_count")
	}
	if a.desired.Spec.ShardCount != nil && !reflect.DeepEqual(desiredPb.ShardCount, a.actual.ShardCount) {
		report.AddField("shard_count", a.actual.ShardCount, desiredPb.ShardCount)
		paths = append(paths, "shard_count")
	}
	if a.desired.Spec.NodeType != nil && !reflect.DeepEqual(desiredPb.NodeType, a.actual.NodeType) {
		report.AddField("node_type", a.actual.NodeType, desiredPb.NodeType)
		paths = append(paths, "node_type")
	}
	if a.desired.Spec.PersistenceConfig != nil && !reflect.DeepEqual(desiredPb.PersistenceConfig, a.actual.PersistenceConfig) {
		report.AddField("persistence_config", a.actual.PersistenceConfig, desiredPb.PersistenceConfig)
		paths = append(paths, "persistence_config")
	}
	if a.desired.Spec.EngineVersion != nil && !reflect.DeepEqual(desiredPb.EngineVersion, a.actual.EngineVersion) {
		report.AddField("engine_version", a.actual.EngineVersion, desiredPb.EngineVersion)
		paths = append(paths, "engine_version")
	}
	if a.desired.Spec.EngineConfigs != nil && !reflect.DeepEqual(desiredPb.EngineConfigs, a.actual.EngineConfigs) {
		report.AddField("engine_configs", a.actual.EngineConfigs, desiredPb.EngineConfigs)
		paths = append(paths, "engine_configs")
	}
	if a.desired.Spec.DeletionProtectionEnabled != nil && !reflect.DeepEqual(desiredPb.DeletionProtectionEnabled, a.actual.DeletionProtectionEnabled) {
		report.AddField("deletion_protection_enabled", a.actual.DeletionProtectionEnabled, desiredPb.DeletionProtectionEnabled)
		paths = append(paths, "deletion_protection_enabled")
	}
	if a.desired.Spec.MaintenancePolicy != nil && !reflect.DeepEqual(desiredPb.MaintenancePolicy, a.actual.MaintenancePolicy) {
		report.AddField("maintenance_policy", a.actual.MaintenancePolicy, desiredPb.MaintenancePolicy)
		paths = append(paths, "maintenance_policy")
	}
	if a.desired.Spec.CrossInstanceReplicationConfig != nil && !reflect.DeepEqual(desiredPb.CrossInstanceReplicationConfig, a.actual.CrossInstanceReplicationConfig) {
		report.AddField("cross_instance_replication_config", a.actual.CrossInstanceReplicationConfig, desiredPb.CrossInstanceReplicationConfig)
		paths = append(paths, "cross_instance_replication_config")
	}
	if a.desired.Spec.AutomatedBackupConfig != nil && !reflect.DeepEqual(desiredPb.AutomatedBackupConfig, a.actual.AutomatedBackupConfig) {
		report.AddField("automated_backup_config", a.actual.AutomatedBackupConfig, desiredPb.AutomatedBackupConfig)
		paths = append(paths, "automated_backup_config")
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		structuredreporting.ReportDiff(ctx, report)
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		for _, path := range paths {
			updateMask := &fieldmaskpb.FieldMask{
				Paths: []string{path},
			}
			req := &memorystorepb.UpdateInstanceRequest{
				UpdateMask: updateMask,
				Instance:   desiredPb,
			}
			req.Instance.Name = a.id.String()
			op, err := a.gcpClient.UpdateInstance(ctx, req)
			if err != nil {
				return fmt.Errorf("updating instance %s: %w", a.id, err)
			}
			updated, err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("instance %s waiting update: %w", a.id, err)
			}
			log.V(2).Info("successfully updated Instance", "name", a.id, "updateMask", path)
		}
		log.V(2).Info("successfully updated Instance", "name", a.id)
	}

	status := &krm.MemorystoreInstanceStatus{}
	status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MemorystoreInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MemorystoreInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MemorystoreInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Instance", "name", a.id)

	req := &memorystorepb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Instance, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Instance", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Instance %s: %w", a.id, err)
	}
	return true, nil
}
