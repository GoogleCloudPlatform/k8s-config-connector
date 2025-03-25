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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.BackupPlan
// crd.type: BackupDRBackupPlan
// crd.version: v1alpha1

package backupdr

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/backupdr/apiv1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BackupDRBackupPlanGVK, NewBackupPlanModel)
}

func NewBackupPlanModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupPlan{config: *config}, nil
}

var _ directbase.Model = &modelBackupPlan{}

type modelBackupPlan struct {
	config config.ControllerConfig
}

func (m *modelBackupPlan) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BackupDRBackupPlan{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupPlanIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Spec.BackupVaultRef != nil {
		if _, err := obj.Spec.BackupVaultRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}

	// Get backupdr GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	backupDRClient, err := gcpClient.newBackupDRClient(ctx)
	if err != nil {
		return nil, err
	}
	return &BackupPlanAdapter{
		id:        id,
		gcpClient: backupDRClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelBackupPlan) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupPlanAdapter struct {
	id        *krm.BackupPlanIdentity
	gcpClient *gcp.Client
	desired   *krm.BackupDRBackupPlan
	actual    *pb.BackupPlan
	reader    client.Reader
}

var _ directbase.Adapter = &BackupPlanAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupPlanAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupPlan", "name", a.id)

	req := &pb.GetBackupPlanRequest{Name: a.id.String()}
	backupplanpb, err := a.gcpClient.GetBackupPlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackupPlan %q: %w", a.id, err)
	}

	a.actual = backupplanpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupPlan", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BackupDRBackupPlanSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBackupPlanRequest{
		Parent:       a.id.Parent().String(),
		BackupPlanId: a.id.ID(),
		BackupPlan:   resource,
	}
	op, err := a.gcpClient.CreateBackupPlan(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BackupPlan %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BackupPlan %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackupPlan", "name", a.id)

	status := &krm.BackupDRBackupPlanStatus{}
	status.ObservedState = BackupDRBackupPlanObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupPlan", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BackupDRBackupPlanSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.BackupRules != nil && !reflect.DeepEqual(resource.BackupRules, a.actual.BackupRules) {
		paths = append(paths, "backup_rules")
	}
	if desired.Spec.ResourceType != nil && !reflect.DeepEqual(resource.ResourceType, a.actual.ResourceType) {
		paths = append(paths, "resource_type")
	}

	if len(paths) != 0 {
		return fmt.Errorf("updating BackupPlan is not supported, fields: %v", paths)
	}

	// still need to update status (in the event of acquiring an existing resource)
	status := &krm.BackupDRBackupPlanStatus{}
	status.ObservedState = BackupDRBackupPlanObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupPlanAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BackupDRBackupPlan{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupDRBackupPlanSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.BackupVaultRef = &krm.BackupVaultRef{External: a.actual.BackupVault}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BackupDRBackupPlanGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupPlanAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupPlan", "name", a.id)

	req := &pb.DeleteBackupPlanRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupPlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BackupPlan, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupPlan %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackupPlan", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupPlan %s: %w", a.id, err)
	}
	return true, nil
}
