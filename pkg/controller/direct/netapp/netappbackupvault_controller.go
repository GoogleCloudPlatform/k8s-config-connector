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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/netapp/apiv1"
	netapppb "cloud.google.com/go/netapp/apiv1/netapppb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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
		return nil, fmt.Errorf("building NetApp client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelBackupVault) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetAppBackupVault{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupVaultIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := BackupVaultSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &backupVaultAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelBackupVault) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type backupVaultAdapter struct {
	id        *krm.BackupVaultIdentity
	gcpClient *gcp.Client
	desired   *netapppb.BackupVault
	actual    *netapppb.BackupVault
}

var _ directbase.Adapter = &backupVaultAdapter{}

func (a *backupVaultAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding NetAppBackupVault", "id", a.id)

	req := &netapppb.GetBackupVaultRequest{
		Name: a.id.String(),
	}
	backupVault, err := a.gcpClient.GetBackupVault(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetAppBackupVault %s: %w", a.id.String(), err)
	}

	a.actual = backupVault
	return true, nil
}

func (a *backupVaultAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetAppBackupVault", "id", a.id)

	req := &netapppb.CreateBackupVaultRequest{
		Parent:        a.id.Parent().String(),
		BackupVaultId: a.id.ID(),
		BackupVault:   a.desired,
	}
	op, err := a.gcpClient.CreateBackupVault(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetAppBackupVault %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for NetAppBackupVault %s creation: %w", a.id.String(), err)
	}

	// Fetch fully-populated resource after creation
	refetched, err := a.gcpClient.GetBackupVault(ctx, &netapppb.GetBackupVaultRequest{Name: a.id.String()})
	if err != nil {
		refetched = created
	}
	a.actual = refetched

	return a.updateStatus(ctx, createOp, refetched)
}

func (a *backupVaultAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetAppBackupVault", "id", a.id)

	diffs, updateMask, err := a.compareBackupVault(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*netapppb.BackupVault)
		desiredCopy.Name = a.id.String()

		req := &netapppb.UpdateBackupVaultRequest{
			BackupVault: desiredCopy,
			UpdateMask:  updateMask,
		}

		op, err := a.gcpClient.UpdateBackupVault(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetAppBackupVault %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting NetAppBackupVault %s update: %w", a.id.String(), err)
		}

		// Fetch fully-populated resource after update
		refetched, err := a.gcpClient.GetBackupVault(ctx, &netapppb.GetBackupVaultRequest{Name: a.id.String()})
		if err != nil {
			refetched = latest
		}
		latest = refetched
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *backupVaultAdapter) compareBackupVault(ctx context.Context, actual, desired *netapppb.BackupVault) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	specActual := BackupVaultSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := BackupVaultSpec_ToProto(mapCtx, specActual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*netapppb.BackupVault)

	populateDefaults := func(obj *netapppb.BackupVault) {
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *backupVaultAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *netapppb.BackupVault) error {
	mapCtx := &direct.MapContext{}
	status := krm.BackupVaultStatus{}
	status.ObservedState = BackupVaultObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *backupVaultAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetAppBackupVault{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BackupVaultSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetAppBackupVaultGVK)

	return u, nil
}

func (a *backupVaultAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetAppBackupVault", "id", a.id)

	req := &netapppb.DeleteBackupVaultRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteBackupVault(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent NetAppBackupVault, assuming it was already deleted", "id", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting NetAppBackupVault %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NetAppBackupVault %s: %w", a.id.String(), err)
	}
	return true, nil
}
