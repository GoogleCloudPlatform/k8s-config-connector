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

// +tool:controller
// proto.service: google.cloud.metastore.v1.DataprocMetastore
// proto.message: google.cloud.metastore.v1.Backup
// crd.type: MetastoreBackup
// crd.version: v1alpha1

package metastore

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/metastore/apiv1"
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MetastoreBackupGVK, NewMetastoreBackupModel)
}

func NewMetastoreBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &MetastoreBackupModel{config: *config}, nil
}

var _ directbase.Model = &MetastoreBackupModel{}

type MetastoreBackupModel struct {
	config config.ControllerConfig
}

func (m *MetastoreBackupModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MetastoreBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get metastore GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	metastoreClient, err := gcpClient.newDataprocMetastoreClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MetastoreBackupAdapter{
		gcpClient: metastoreClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *MetastoreBackupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MetastoreBackupAdapter struct {
	gcpClient *gcp.DataprocMetastoreClient
	id        *krm.BackupIdentity
	desired   *krm.MetastoreBackup
	actual    *pb.Backup
	reader    client.Reader
}

var _ directbase.Adapter = &MetastoreBackupAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MetastoreBackupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MetastoreBackup", "name", a.id)

	req := &pb.GetBackupRequest{Name: a.id.String()}
	metastorebackuppb, err := a.gcpClient.GetBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MetastoreBackup %q: %w", a.id, err)
	}

	a.actual = metastorebackuppb
	return true, nil
}

func (a *MetastoreBackupAdapter) resolveReferences(ctx context.Context) error {
	obj := a.desired

	if _, err := obj.Spec.ServiceRef.NormalizedExternal(ctx, a.reader, obj.Namespace); err != nil {
		return fmt.Errorf("resolving serviceRef: %w", err)
	}
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetastoreBackupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MetastoreBackup", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MetastoreBackupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBackupRequest{
		Parent:   a.id.Parent().String(),
		BackupId: a.id.ID(),
		Backup:   resource,
	}
	op, err := a.gcpClient.CreateBackup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MetastoreBackup %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetastoreBackup %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MetastoreBackup", "name", a.id)

	status := &krm.MetastoreBackupStatus{}
	status.ObservedState = MetastoreBackupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
// MetastoreBackup resource has no mutable fields as per the proto definition, so this function only updates the status.
func (a *MetastoreBackupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("resource is immutable, skipping update", "name", a.id)

	// No update API call as the resource is immutable or has no mutable fields exposed.
	// We only need to update the status based on the last known state (a.actual).
	mapCtx := &direct.MapContext{}
	status := &krm.MetastoreBackupStatus{}
	status.ObservedState = MetastoreBackupObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MetastoreBackupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MetastoreBackup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MetastoreBackupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate required references from the ID
	obj.Spec.ServiceRef = krm.ServiceRef{
		External: a.id.Parent().String(), // Set external reference to the service name
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MetastoreBackupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MetastoreBackupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MetastoreBackup", "name", a.id)

	req := &pb.DeleteBackupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MetastoreBackup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MetastoreBackup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion for MetastoreBackup", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MetastoreBackup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MetastoreBackup", "name", a.id)
	return true, nil
}
