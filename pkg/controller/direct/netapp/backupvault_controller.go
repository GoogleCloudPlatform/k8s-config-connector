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

package netapp

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/netapp/apiv1"
	netapppb "cloud.google.com/go/netapp/apiv1/netapppb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetAppBackupVaultGVK, NewBackupVaultModel)
}

func NewBackupVaultModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupVault{config: *config}, nil
}

var _ directbase.Model = &modelBackupVault{}

type modelBackupVault struct {
	config config.ControllerConfig
}

func (m *modelBackupVault) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BackupVault client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBackupVault) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetAppBackupVault{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupVaultIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get netapp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BackupVaultAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBackupVault) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupVaultAdapter struct {
	id        *krm.BackupVaultIdentity
	gcpClient *gcp.Client
	desired   *krm.NetAppBackupVault
	actual    *netapppb.BackupVault
}

var _ directbase.Adapter = &BackupVaultAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupVaultAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupVault", "name", a.id)

	req := &netapppb.GetBackupVaultRequest{Name: a.id.String()}
	backupvaultpb, err := a.gcpClient.GetBackupVault(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackupVault %q: %w", a.id, err)
	}

	a.actual = backupvaultpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupVaultAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupVault", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BackupVault_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &netapppb.CreateBackupVaultRequest{
		Parent:        a.id.Parent().String(),
		BackupVault:   resource,
		BackupVaultId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateBackupVault(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BackupVault %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BackupVault %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackupVault", "name", a.id)

	status := &krm.BackupVaultStatus{}
	status.ObservedState = BackupVaultObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupVaultAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupVault", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := BackupVault_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &netapppb.UpdateBackupVaultRequest{
		UpdateMask:  updateMask,
		BackupVault: desiredPb,
	}
	op, err := a.gcpClient.UpdateBackupVault(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BackupVault %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BackupVault %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated BackupVault", "name", a.id)

	status := &krm.BackupVaultStatus{}
	status.ObservedState = BackupVaultObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupVaultAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetAppBackupVault{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupVault_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetAppBackupVaultGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupVaultAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupVault", "name", a.id)

	req := &netapppb.DeleteBackupVaultRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackupVault(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupVault %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackupVault", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}
