// Copyright 2026 Google LLC
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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/memorystore/apiv1"

	memorystorepb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MemorystoreInstanceBackupGVK, NewBackupModel)
}

func NewBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackup{config: *config}, nil
}

var _ directbase.Model = &modelBackup{}

type modelBackup struct {
	config config.ControllerConfig
}

func (m *modelBackup) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Backup client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBackup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	obj := &krm.MemorystoreInstanceBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if obj.Spec.InstanceRef == nil && obj.Spec.BackupCollection == nil {
		return nil, fmt.Errorf("instanceRef and backupCollection are not set for %s", obj.Name)
	}

	// Get memorystore GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &BackupAdapter{
		gcpClient: gcpClient,
		reader:    op.Reader,
		desired:   obj,
	}, nil
}

func resolveBackupName(obj *krm.MemorystoreInstanceBackup) string {
	if obj.Spec.ResourceID != nil {
		return *obj.Spec.ResourceID
	}
	return obj.GetName()
}

func resolveBackupRefereces(ctx context.Context, reader client.Reader, gcpClient *gcp.Client, obj *krm.MemorystoreInstanceBackup) error {
	if obj.Spec.InstanceRef != nil {
		if obj.Spec.InstanceRef.External == "" {
			instanceRef, err := refs.ResolveMemorystoreInstance(ctx, reader, obj, obj.Spec.InstanceRef)
			if err != nil {
				return err
			}
			obj.Spec.InstanceRef.External = instanceRef.External
		}
		if obj.Spec.BackupCollection == nil {
			instance, err := gcpClient.GetInstance(ctx, &memorystorepb.GetInstanceRequest{
				Name: obj.Spec.InstanceRef.External,
			})
			if instance == nil || err != nil {
				return fmt.Errorf("failed to get instance %q: %w", obj.Spec.InstanceRef.External, err)
			}
			if instance.BackupCollection != nil {
				obj.Spec.BackupCollection = instance.BackupCollection
			}
		}
	}
	return nil
}

func (m *modelBackup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupAdapter struct {
	id        *krm.BackupIdentity
	gcpClient *gcp.Client
	reader    client.Reader
	desired   *krm.MemorystoreInstanceBackup
	actual    *memorystorepb.Backup
}

var _ directbase.Adapter = &BackupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupAdapter) Find(ctx context.Context) (bool, error) {
	if err := resolveBackupRefereces(ctx, a.reader, a.gcpClient, a.desired); err != nil {
		return false, err
	}

	id, err := krm.NewBackupIdentity(ctx, a.desired)
	if err != nil {
		return false, err
	}

	a.id = id

	log := klog.FromContext(ctx)
	log.V(2).Info("getting Backup", "name", a.id.String())
	req := &memorystorepb.GetBackupRequest{Name: a.id.String()}
	backuppb, err := a.gcpClient.GetBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Backup %q: %w", a.id, err)
	}

	a.actual = backuppb
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	name := resolveBackupName(a.desired)

	log := klog.FromContext(ctx)
	log.V(2).Info("creating Backup", "name", name)

	if a.desired.Spec.InstanceRef == nil || a.desired.Spec.InstanceRef.External == "" {
		return fmt.Errorf("instanceRef is not set for %s", name)
	}

	desired := a.desired.DeepCopy()
	desired.Spec.ResourceID = &name
	mapCtx := &direct.MapContext{}
	reqCreate := MemorystoreInstanceBackupSpec_ToProtoBackupInstanceRequest(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	opCreate, err := a.gcpClient.BackupInstance(ctx, reqCreate)
	if err != nil {
		return fmt.Errorf("creating Backup %s: %w", name, err)
	}
	instance, err := opCreate.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Backup %s waiting creation: %w", name, err)
	}
	log.V(2).Info("successfully created Backup", "name", name)

	a.desired.Spec.BackupCollection = instance.BackupCollection
	id, err := krm.NewBackupIdentity(ctx, a.desired)
	if err != nil {
		return err
	}

	reqGet := &memorystorepb.GetBackupRequest{Name: id.String()}
	backup, err := a.gcpClient.GetBackup(ctx, reqGet)
	if err != nil {
		return fmt.Errorf("getting Backup %s: %w", name, err)
	}

	a.id = id

	status := &krm.MemorystoreInstanceBackupStatus{}
	status.ObservedState = MemorystoreInstanceBackupObservedState_FromProto(mapCtx, backup)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Backup", "name", a.id)

	status := &krm.MemorystoreInstanceBackupStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = MemorystoreInstanceBackupObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MemorystoreInstanceBackup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MemorystoreInstanceBackupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MemorystoreInstanceBackupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Backup", "name", a.id)

	req := &memorystorepb.DeleteBackupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Backup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Backup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Backup", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Backup %s: %w", a.id, err)
	}
	return true, nil
}
