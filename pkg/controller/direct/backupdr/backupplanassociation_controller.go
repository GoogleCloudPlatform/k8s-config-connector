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
// proto.message: google.cloud.backupdr.v1.BackupPlanAssociation
// crd.type: BackupDRBackupPlanAssociation
// crd.version: v1alpha1

package backupdr

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.BackupDRBackupPlanAssociationGVK, NewBackupPlanAssociationModel)
}

func NewBackupPlanAssociationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupPlanAssociation{config: *config}, nil
}

var _ directbase.Model = &modelBackupPlanAssociation{}

type modelBackupPlanAssociation struct {
	config config.ControllerConfig
}

func (m *modelBackupPlanAssociation) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BackupDRBackupPlanAssociation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupPlanAssociationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
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
	return &BackupPlanAssociationAdapter{
		id:        id,
		gcpClient: backupDRClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelBackupPlanAssociation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupPlanAssociationAdapter struct {
	id        *krm.BackupPlanAssociationIdentity
	gcpClient *gcp.Client
	desired   *krm.BackupDRBackupPlanAssociation
	actual    *pb.BackupPlanAssociation
	reader    client.Reader
}

var _ directbase.Adapter = &BackupPlanAssociationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupPlanAssociationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupPlanAssociation", "name", a.id)

	req := &pb.GetBackupPlanAssociationRequest{Name: a.id.String()}
	backupplanassociationpb, err := a.gcpClient.GetBackupPlanAssociation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackupPlanAssociation %q: %w", a.id, err)
	}

	a.actual = backupplanassociationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAssociationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupPlanAssociation", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BackupDRBackupPlanAssociationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBackupPlanAssociationRequest{
		Parent:                  a.id.Parent().String(),
		BackupPlanAssociationId: a.id.ID(),
		BackupPlanAssociation:   resource,
	}
	if desired.Spec.Resource.ComputeInstanceRef != nil {
		fmt.Printf("[debug] desired.spec.resource.computeInstanceRef: %+v\n", desired.Spec.Resource.ComputeInstanceRef)
		fmt.Printf("[debug] desired.spec.resource.computeInstanceRef.external: %+v\n", desired.Spec.Resource.ComputeInstanceRef.External)
	}
	fmt.Printf("[debug] desired.spec.resource_type: %s\n", *desired.Spec.ResourceType)
	fmt.Printf("[debug] resource: %+v\n", resource)
	fmt.Printf("[debug] create request: %+v\n", req)
	op, err := a.gcpClient.CreateBackupPlanAssociation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BackupPlanAssociation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BackupPlanAssociation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackupPlanAssociation", "name", a.id)
	fmt.Printf("[debug] created: %+v\n", created)

	status := &krm.BackupDRBackupPlanAssociationStatus{}
	status.ObservedState = BackupDRBackupPlanAssociationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupPlanAssociationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupPlan", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BackupDRBackupPlanAssociationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.ResourceType != nil && !reflect.DeepEqual(resource.ResourceType, a.actual.ResourceType) {
		paths = append(paths, "resource_type")
	}
	if desired.Spec.Resource != nil && !reflect.DeepEqual(resource.Resource, a.actual.Resource) {
		paths = append(paths, "resource")
	}
	if desired.Spec.BackupPlanRef != nil && !reflect.DeepEqual(resource.BackupPlan, a.actual.BackupPlan) {
		paths = append(paths, "backup_plan")
	}

	if len(paths) != 0 {
		return fmt.Errorf("updating BackupPlan is not supported, fields: %v", paths)
	}

	// still need to update status (in the event of acquiring an existing resource)
	status := &krm.BackupDRBackupPlanAssociationStatus{}
	status.ObservedState = BackupDRBackupPlanAssociationObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupPlanAssociationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BackupDRBackupPlanAssociation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupDRBackupPlanAssociationSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.BackupDRBackupPlanAssociationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupPlanAssociationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupPlanAssociation", "name", a.id)

	req := &pb.DeleteBackupPlanAssociationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupPlanAssociation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BackupPlanAssociation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupPlanAssociation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackupPlanAssociation", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupPlanAssociation %s: %w", a.id, err)
	}
	return true, nil
}

func (a *BackupPlanAssociationAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.BackupPlanRef != nil {
		if _, err := obj.Spec.BackupPlanRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.Resource != nil {
		if obj.Spec.Resource.ComputeInstanceRef != nil {
			if _, err := obj.Spec.Resource.ComputeInstanceRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}

	return nil
}
