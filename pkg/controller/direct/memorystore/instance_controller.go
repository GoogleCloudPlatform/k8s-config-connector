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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/memorystore/apiv1beta"

	memorystorepb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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

func (m *modelInstance) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MemorystoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	normalize(ctx, reader, obj)

	// Get memorystore GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &InstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func normalize(ctx context.Context, reader client.Reader, obj *krm.MemorystoreInstance) error {
	for _, endpoint := range obj.Spec.Endpoints {
		for _, connection := range endpoint.Connections {
			if connection.PscAutoConnection != nil {
				autoConnection := connection.PscAutoConnection
				if autoConnection.NetworkRef != nil {
					if err := autoConnection.NetworkRef.Normalize(ctx, reader, obj); err != nil {
						return err
					}
				}
			}
			if connection.PscConnection != nil {
				userConnection := connection.PscConnection
				if userConnection.NetworkRef != nil {
					if err := userConnection.NetworkRef.Normalize(ctx, reader, obj); err != nil {
						return err
					}
				}
				err := refs.ResolveComputeServiceAttachment(ctx, reader, obj.GetNamespace(), userConnection.ServiceAttachmentRef)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *modelInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type InstanceAdapter struct {
	id        *krm.InstanceIdentity
	gcpClient *gcp.Client
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
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
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
	mapCtx := &direct.MapContext{}

	desiredPb := MemorystoreInstanceSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])

	// If replica count is unset, the field become unmanaged.
	if a.desired.Spec.ReplicaCount != nil && !reflect.DeepEqual(a.desired.Spec.ReplicaCount, a.actual.ReplicaCount) {
		paths.Insert("replica_count")
	}
	if a.desired.Spec.ShardCount != nil && !reflect.DeepEqual(a.desired.Spec.ShardCount, a.actual.ShardCount) {
		log.V(2).Info("diff shard count", "desired", a.desired.Spec.ShardCount, "actual", a.actual.ShardCount)
		paths.Insert("shard_count")
	}
	if a.desired.Spec.DeletionProtectionEnabled != nil && !reflect.DeepEqual(a.desired.Spec.DeletionProtectionEnabled, a.actual.DeletionProtectionEnabled) {
		paths.Insert("deletion_protection_enabled")
	}
	if a.desired.Spec.PersistenceConfig != nil && !reflect.DeepEqual(a.desired.Spec.PersistenceConfig, a.actual.PersistenceConfig) {
		paths.Insert("persistence_config")
	}
	if a.desired.Spec.EngineConfigs != nil && !reflect.DeepEqual(a.desired.Spec.EngineConfigs, a.actual.EngineConfigs) {
		paths.Insert("engine_configs")
	}
	if a.desired.Spec.Endpoints != nil && !reflect.DeepEqual(a.desired.Spec.Endpoints, a.actual.Endpoints) {
		paths.Insert("endpoints")
	}
	if a.desired.Spec.Labels != nil && !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
		paths.Insert("labels")
	}
	if a.desired.Spec.EngineVersion != nil && !reflect.DeepEqual(a.desired.Spec.EngineVersion, a.actual.EngineVersion) {
		paths.Insert("engine_version")
	}
	if a.desired.Spec.NodeType != nil && !reflect.DeepEqual(a.desired.Spec.NodeType, a.actual.NodeType) {
		paths.Insert("node_type")
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		for _, path := range paths.UnsortedList() {
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
