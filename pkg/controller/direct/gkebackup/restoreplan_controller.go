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
// proto.message: google.cloud.gkebackup.v1.RestorePlan
// crd.type: GKEBackupRestorePlan
// crd.version: v1alpha1

package gkebackup

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/gkebackup/apiv1"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
)

func init() {
	registry.RegisterModel(krm.GKEBackupRestorePlanGVK, NewRestorePlanModel)
}

func NewRestorePlanModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &restorePlanModel{config: *config}, nil
}

var _ directbase.Model = &restorePlanModel{}

type restorePlanModel struct {
	config config.ControllerConfig
}

func (m *restorePlanModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.GKEBackupRestorePlan{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewRestorePlanIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// normalize required reference fields
	if obj.Spec.BackupPlanRef != nil {
		if _, err := obj.Spec.BackupPlanRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}
	if obj.Spec.ClusterRef != nil {
		if _, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
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
	return &restorePlanAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *restorePlanModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type restorePlanAdapter struct {
	gcpClient         *gcp.BackupForGKEClient
	id                *krm.RestorePlanIdentity
	desired           *krm.GKEBackupRestorePlan
	actual            *pb.RestorePlan
	resourceOverrides resourceoverrides.ResourceOverrides
}

var _ directbase.Adapter = &restorePlanAdapter{}

func (a *restorePlanAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting gkebackup restoreplan", "name", a.id)

	req := &pb.GetRestorePlanRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetRestorePlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting gkebackup restoreplan %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *restorePlanAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating gkebackup restoreplan", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupRestorePlanSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateRestorePlanRequest{
		Parent:        a.id.Parent().String(),
		RestorePlanId: a.id.ID(),
		RestorePlan:   resource,
	}
	op, err := a.gcpClient.CreateRestorePlan(ctx, req)
	if err != nil {
		return fmt.Errorf("creating gkebackup restoreplan %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("gkebackup restoreplan %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created gkebackup restoreplan in gcp", "name", a.id)

	status := &krm.GKEBackupRestorePlanStatus{}
	status.ObservedState = GKEBackupRestorePlanObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *restorePlanAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating gkebackup restoreplan", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEBackupRestorePlanSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}
	if desired.Spec.RestoreConfig != nil && !restoreConfigsEqual(resource.GetRestoreConfig(), a.actual.GetRestoreConfig()) {
		paths = append(paths, "restore_config")
	}

	var updated *pb.RestorePlan
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateRestorePlanRequest{
			RestorePlan: resource,
			UpdateMask:  &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateRestorePlan(ctx, req)
		if err != nil {
			return fmt.Errorf("updating gkebackup restoreplan %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("gkebackup restoreplan %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated gkebackup restoreplan", "name", a.id)
	}

	status := &krm.GKEBackupRestorePlanStatus{}
	status.ObservedState = GKEBackupRestorePlanObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *restorePlanAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEBackupRestorePlan{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEBackupRestorePlanSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.GKEBackupRestorePlanGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *restorePlanAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting gkebackup restoreplan", "name", a.id)

	req := &pb.DeleteRestorePlanRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRestorePlan(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent RestorePlan, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting gkebackup restoreplan %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted gkebackup restoreplan", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete RestorePlan %s: %w", a.id, err)
	}
	return true, nil
}

func restoreConfigsEqual(a, b *pb.RestoreConfig) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Compare VolumeDataRestorePolicy
	if a.GetVolumeDataRestorePolicy() != b.GetVolumeDataRestorePolicy() {
		return false
	}

	// Compare ClusterResourceConflictPolicy
	if a.GetClusterResourceConflictPolicy() != b.GetClusterResourceConflictPolicy() {
		return false
	}

	// Compare NamespacedResourceRestoreMode
	if a.GetNamespacedResourceRestoreMode() != b.GetNamespacedResourceRestoreMode() {
		return false
	}

	// Compare ClusterResourceRestoreScope
	aScope, bScope := a.GetClusterResourceRestoreScope(), b.GetClusterResourceRestoreScope()
	if (aScope == nil) != (bScope == nil) {
		return false
	}
	if aScope != nil {
		if aScope.GetAllGroupKinds() != bScope.GetAllGroupKinds() {
			return false
		}
		if aScope.GetNoGroupKinds() != bScope.GetNoGroupKinds() {
			return false
		}

		// Compare SelectedGroupKinds
		aSelected, bSelected := aScope.GetSelectedGroupKinds(), bScope.GetSelectedGroupKinds()
		if len(aSelected) != len(bSelected) {
			return false
		}
		for i := range aSelected {
			if !proto.Equal(aSelected[i], bSelected[i]) {
				return false
			}
		}

		// Compare ExcludedGroupKinds
		aExcluded, bExcluded := aScope.GetExcludedGroupKinds(), bScope.GetExcludedGroupKinds()
		if len(aExcluded) != len(bExcluded) {
			return false
		}
		for i := range aExcluded {
			if !proto.Equal(aExcluded[i], bExcluded[i]) {
				return false
			}
		}
	}

	// Compare NamespacedResourceRestoreScope (oneof field)
	// We need to check which field is set in the oneof
	switch {
	case a.GetAllNamespaces():
		if !b.GetAllNamespaces() {
			return false
		}
	case a.GetSelectedNamespaces() != nil:
		bSelected := b.GetSelectedNamespaces()
		if bSelected == nil || !proto.Equal(a.GetSelectedNamespaces(), bSelected) {
			return false
		}
	case a.GetSelectedApplications() != nil:
		bSelected := b.GetSelectedApplications()
		if bSelected == nil || !proto.Equal(a.GetSelectedApplications(), bSelected) {
			return false
		}
	case a.GetNoNamespaces():
		if !b.GetNoNamespaces() {
			return false
		}
	case a.GetExcludedNamespaces() != nil:
		bExcluded := b.GetExcludedNamespaces()
		if bExcluded == nil || !proto.Equal(a.GetExcludedNamespaces(), bExcluded) {
			return false
		}
	default:
		// If we get here, the oneof field might not be set in a
		// Check if it's set in b
		if b.GetAllNamespaces() || b.GetSelectedNamespaces() != nil ||
			b.GetSelectedApplications() != nil || b.GetNoNamespaces() ||
			b.GetExcludedNamespaces() != nil {
			return false
		}
	}

	// Compare SubstitutionRules
	if !proto.Equal(&pb.RestoreConfig{SubstitutionRules: a.GetSubstitutionRules()},
		&pb.RestoreConfig{SubstitutionRules: b.GetSubstitutionRules()}) {
		return false
	}

	// Compare TransformationRules
	if !proto.Equal(&pb.RestoreConfig{TransformationRules: a.GetTransformationRules()},
		&pb.RestoreConfig{TransformationRules: b.GetTransformationRules()}) {
		return false
	}

	// Compare VolumeDataRestorePolicyBindings
	if !proto.Equal(&pb.RestoreConfig{VolumeDataRestorePolicyBindings: a.GetVolumeDataRestorePolicyBindings()},
		&pb.RestoreConfig{VolumeDataRestorePolicyBindings: b.GetVolumeDataRestorePolicyBindings()}) {
		return false
	}

	// Compare RestoreOrder
	if !proto.Equal(a.GetRestoreOrder(), b.GetRestoreOrder()) {
		return false
	}

	return true
}
