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

package alloydbbackup

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	gcp "cloud.google.com/go/alloydb/apiv1beta"
	alloydbpb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AlloyDBBackupGVK, NewBackupModel)
}

func NewBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backupModel{config: *config}, nil
}

var _ directbase.Model = &backupModel{}

type backupModel struct {
	config config.ControllerConfig
}

func (m *backupModel) client(ctx context.Context) (*gcp.AlloyDBAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAlloyDBAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building AlloyDB client for Backup: %w", err)
	}
	return gcpClient, err
}

func (m *backupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AlloyDBBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AlloyDBBackupIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &backupAdapter{
		id:        id,
		gcpClient: gcpClient,
		reader:    reader,
		desired:   obj,
	}, nil
}

func (m *backupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type backupAdapter struct {
	id        *krm.AlloyDBBackupIdentity
	gcpClient *gcp.AlloyDBAdminClient
	reader    client.Reader
	desired   *krm.AlloyDBBackup
	actual    *alloydbpb.Backup
}

var _ directbase.Adapter = &backupAdapter{}

func (a *backupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting backup", "name", a.id.String())

	req := &alloydbpb.GetBackupRequest{Name: a.id.String()}
	backuppb, err := a.gcpClient.GetBackup(ctx, req)
	if err != nil {
		log.V(2).Info("error getting backup", "name", a.id, "error", err)
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting backup %q: %w", a.id, err)
	}

	a.actual = backuppb
	return true, nil
}

func (a *backupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating backup", "name", a.id)
	mapCtx := &direct.MapContext{}

	// Resolve the reference fields.
	if a.desired.Spec.ClusterNameRef != nil {
		if err := a.desired.Spec.ClusterNameRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}

	desired := a.desired.DeepCopy()
	resource := AlloyDBBackupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	resource.Labels["managed-by-cnrm"] = "true"

	req := &alloydbpb.CreateBackupRequest{
		Parent:   a.id.ParentString(),
		BackupId: a.id.ID(),
		Backup:   resource,
	}
	op, err := a.gcpClient.CreateBackup(ctx, req)
	if err != nil {
		log.V(2).Info("error creating backup", "name", a.id, "error", err)
		return fmt.Errorf("creating backup %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting backup creation", "name", a.id, "error", err)
		return fmt.Errorf("backup %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created backup", "name", a.id)

	status := AlloyDBBackupStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *backupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating backup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredLabels := label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	if desiredLabels == nil {
		desiredLabels = make(map[string]string)
	}
	desiredLabels["managed-by-cnrm"] = "true"

	var updatePaths []string
	if !reflect.DeepEqual(a.actual.GetLabels(), desiredLabels) {
		log.V(2).Info("'metadata.labels' field is updated (-old +new)", cmp.Diff(a.actual.GetLabels(), desiredLabels))
		updatePaths = append(updatePaths, "labels")
	}

	if len(updatePaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := AlloyDBBackupStatus_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	updateMask := &fieldmaskpb.FieldMask{
		Paths: updatePaths,
	}
	desiredPb := &alloydbpb.Backup{
		Name:   a.id.String(),
		Labels: desiredLabels,
	}
	req := &alloydbpb.UpdateBackupRequest{
		UpdateMask: updateMask,
		Backup:     desiredPb,
	}
	op, err := a.gcpClient.UpdateBackup(ctx, req)
	if err != nil {
		log.V(2).Info("error updating backup", "name", a.id, "error", err)
		return fmt.Errorf("updating backup %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		log.V(2).Info("error waiting backup update", "name", a.id, "error", err)
		return fmt.Errorf("backup %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated backup", "name", a.id)

	status := AlloyDBBackupStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *backupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting backup", "name", a.id)

	req := &alloydbpb.DeleteBackupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteBackup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting backup %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting backup deletion: %w", err)
	}
	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *backupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AlloyDBBackup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AlloyDBBackupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.AlloyDBBackupGVK)

	raw, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = raw
	return u, nil
}
