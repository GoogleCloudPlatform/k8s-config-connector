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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupPlan
// crd.type: GKEBackupBackupPlan
// crd.version: v1alpha1

package gkebackup

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/gkebackup/apiv1"
	"cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEBackupBackupPlanGVK, NewBackupPlanModel)
}

func NewBackupPlanModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backupPlanModel{config: *config}, nil
}

var _ directbase.Model = &backupPlanModel{}

type backupPlanModel struct {
	config config.ControllerConfig
}

func (m *backupPlanModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.GKEBackupBackupPlan{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupPlanIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get gkebackup GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newBackupForGKEClient(ctx)
	if err != nil {
		return nil, err
	}
	return &backupPlanAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *backupPlanModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type backupPlanAdapter struct {
	gcpClient *gcp.BackupForGKEClient
	id        *krm.BackupPlanIdentity
	desired   *krm.GKEBackupBackupPlan
	actual    *pb.BackupPlan
	reader    client.Reader
}

var _ directbase.Adapter = &backupPlanAdapter{}

func (a *backupPlanAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting gkebackup backupplan", "name", a.id)

	req := &pb.GetBackupPlanRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBackupPlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting gkebackup backupplan %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *backupPlanAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating gkebackup backupplan", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupPlanSpec_ToProto(mapCtx, &desired.Spec)
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
		return fmt.Errorf("creating gkebackup backupplan %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("gkebackup backupplan %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created gkebackup backupplan in gcp", "name", a.id)

	status := &krm.GKEBackupBackupPlanStatus{}
	status.ObservedState = GKEBackupBackupPlanObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *backupPlanAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating gkebackup backupplan", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := GKEBackupBackupPlanSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}
	if desired.Spec.RetentionPolicy != nil && !reflect.DeepEqual(resource.GetRetentionPolicy(), a.actual.GetRetentionPolicy()) {
		paths = append(paths, "retention_policy")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}
	if desired.Spec.BackupSchedule != nil && !reflect.DeepEqual(resource.GetBackupSchedule(), a.actual.GetBackupSchedule()) {
		paths = append(paths, "backup_schedule")
	}
	if desired.Spec.Deactivated != nil && !reflect.DeepEqual(resource.GetDeactivated(), a.actual.GetDeactivated()) {
		paths = append(paths, "deactivated")
	}

	// cannot use reflect.DeepEqual here because BackupScope is a oneof field which results in unexpected diffs
	if desired.Spec.BackupConfig != nil && !backupConfigsEqual(resource.GetBackupConfig(), a.actual.GetBackupConfig()) {
		paths = append(paths, "backup_config")
	}

	var updated *pb.BackupPlan
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateBackupPlanRequest{
			BackupPlan: resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateBackupPlan(ctx, req)
		if err != nil {
			return fmt.Errorf("updating gkebackup backupplan %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("gkebackup backupplan %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated gkebackup backupplan", "name", a.id)
	}

	status := &krm.GKEBackupBackupPlanStatus{}
	status.ObservedState = GKEBackupBackupPlanObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *backupPlanAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEBackupBackupPlan{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEBackupBackupPlanSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.GKEBackupBackupPlanGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *backupPlanAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting gkebackup backupplan", "name", a.id)

	req := &pb.DeleteBackupPlanRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupPlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting gkebackup backupplan %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted gkebackup backupplan", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}

func backupConfigsEqual(a, b *gkebackuppb.BackupPlan_BackupConfig) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Compare BackupScope (oneof field)
	aScope, bScope := a.GetBackupScope(), b.GetBackupScope()
	if (aScope == nil) != (bScope == nil) {
		return false
	}
	if aScope != nil {
		// Check which type is set in the oneof
		switch aTyped := aScope.(type) {
		case *gkebackuppb.BackupPlan_BackupConfig_AllNamespaces:
			bTyped, ok := bScope.(*gkebackuppb.BackupPlan_BackupConfig_AllNamespaces)
			if !ok || aTyped.AllNamespaces != bTyped.AllNamespaces {
				return false
			}
		case *gkebackuppb.BackupPlan_BackupConfig_SelectedNamespaces:
			bTyped, ok := bScope.(*gkebackuppb.BackupPlan_BackupConfig_SelectedNamespaces)
			if !ok || !proto.Equal(aTyped.SelectedNamespaces, bTyped.SelectedNamespaces) {
				return false
			}
		case *gkebackuppb.BackupPlan_BackupConfig_SelectedApplications:
			bTyped, ok := bScope.(*gkebackuppb.BackupPlan_BackupConfig_SelectedApplications)
			if !ok || !proto.Equal(aTyped.SelectedApplications, bTyped.SelectedApplications) {
				return false
			}
		default:
			// Unknown type
			return false
		}
	}

	// Compare other fields
	if a.GetIncludeVolumeData() != b.GetIncludeVolumeData() {
		return false
	}
	if a.GetIncludeSecrets() != b.GetIncludeSecrets() {
		return false
	}
	if !proto.Equal(a.GetEncryptionKey(), b.GetEncryptionKey()) {
		return false
	}
	if a.GetPermissiveMode() != b.GetPermissiveMode() {
		return false
	}

	return true
}

func (a *backupPlanAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.ClusterRef != nil {
		if _, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.BackupConfig != nil &&
		obj.Spec.BackupConfig.EncryptionKey != nil &&
		obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef != nil {
		if _, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.BackupConfig.EncryptionKey.KMSKeyRef); err != nil {
			return err
		}
	}

	return nil
}
