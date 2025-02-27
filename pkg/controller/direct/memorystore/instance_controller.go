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
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MemorystoreInstanceGVK, NewinstanceModel)
}

func NewinstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelinstance{config: *config}, nil
}

var _ directbase.Model = &modelinstance{}

type modelinstance struct {
	config config.ControllerConfig
}

type instanceAdapter struct {
	id        *krm.MemoryStoreInstanceIdentity
	gcpClient *gcp.Client
	desired   *krm.MemorystoreInstance
	actual    *memorystorepb.Instance
}

var _ directbase.Adapter = &instanceAdapter{}

func (m *modelinstance) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building instance client: %w", err)
	}
	return gcpClient, err
}

func (m *modelinstance) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MemorystoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMemoryStoreInstanceIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	// Get memorystore GCP client
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

func (m *modelinstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting instance", "name", a.id)

	req := &memorystorepb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting instance %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MemorystoreInstanceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &memorystorepb.CreateInstanceRequest{
		Instance:   resource,
		InstanceId: a.id.ID(),
		Parent:     a.id.Parent().String(),
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("instance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created instance", "name", a.id)

	status := &krm.MemorystoreInstanceStatus{}
	status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *instanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating instance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MemorystoreInstanceSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}

	// If replica count is unset, the field become unmanaged.
	if a.desired.Spec.ReplicaCount != nil && !reflect.DeepEqual(a.desired.Spec.ReplicaCount, a.actual.ReplicaCount) {
		paths = append(paths, "replica_count")
	}
	if a.desired.Spec.ShardCount != nil && !reflect.DeepEqual(a.desired.Spec.ShardCount, a.actual.ShardCount) {
		paths = append(paths, "shard_count")
	}
	if a.desired.Spec.DeletionProtectionEnabled != nil && !reflect.DeepEqual(a.desired.Spec.DeletionProtectionEnabled, a.actual.DeletionProtectionEnabled) {
		paths = append(paths, "deletion_protection_enabled")
	}
	if a.desired.Spec.PersistenceConfig != nil && !reflect.DeepEqual(a.desired.Spec.PersistenceConfig, a.actual.PersistenceConfig) {
		paths = append(paths, "persistence_config")
	}
	if a.desired.Spec.EngineConfigs != nil && !reflect.DeepEqual(a.desired.Spec.EngineConfigs, a.actual.EngineConfigs) {
		paths = append(paths, "engine_configs")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.MemorystoreInstanceStatus{}
		status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	var latest *memorystorepb.Instance
	if len(paths) > 0 {
		// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
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
			latest, err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("instance %s waiting update: %w", a.id, err)
			}
		}
	} else {
		latest = a.actual
	}

	log.V(2).Info("successfully updated instance", "name", a.id)

	status := &krm.MemorystoreInstanceStatus{}
	status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *instanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.MemorystoreInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *instanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting instance", "name", a.id)

	req := &memorystorepb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted instance", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete instance %s: %w", a.id, err)
	}
	return true, nil
}
