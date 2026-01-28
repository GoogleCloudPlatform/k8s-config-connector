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

package spanner

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/spanner/admin/database/apiv1"
	spannerbackupschedulespb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SpannerBackupScheduleGVK, NewBackupScheduleModel)
}

func NewBackupScheduleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupSchedule{config: *config}, nil
}

var _ directbase.Model = &modelBackupSchedule{}

var outputOnlyFields = []string{"name", "spec.cron_spec.creation_window", "spec.cron_spec.creation_window"}
var defaultOnEmptyFields = []string{"encryption_config.encryption_type"}

type modelBackupSchedule struct {
	config config.ControllerConfig
}

func (m *modelBackupSchedule) client(ctx context.Context) (*gcp.DatabaseAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDatabaseAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BackupSchedule client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBackupSchedule) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpannerBackupSchedule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupScheduleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get spannerbackupschedules GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BackupScheduleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBackupSchedule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackupScheduleAdapter struct {
	id        *krm.BackupScheduleIdentity
	gcpClient *gcp.DatabaseAdminClient
	desired   *krm.SpannerBackupSchedule
	actual    *spannerbackupschedulespb.BackupSchedule
}

var _ directbase.Adapter = &BackupScheduleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupScheduleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupSchedule", "name", a.id)

	req := &spannerbackupschedulespb.GetBackupScheduleRequest{Name: a.id.String()}
	backupschedulepb, err := a.gcpClient.GetBackupSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackupSchedule %q: %w", a.id, err)
	}

	a.actual = backupschedulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupScheduleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupSchedule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SpannerBackupScheduleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &spannerbackupschedulespb.CreateBackupScheduleRequest{
		Parent:           a.id.Parent().String(),
		BackupSchedule:   resource,
		BackupScheduleId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateBackupSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BackupSchedule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BackupSchedule", "name", a.id)

	status := &krm.SpannerBackupScheduleStatus{}
	status.ObservedState = SpannerBackupScheduleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupScheduleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupSchedule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SpannerBackupScheduleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	// Remove output only fields from paths
	paths = paths.Delete(outputOnlyFields...)

	// Check default-on-empty fields
	if desiredPb.EncryptionConfig == nil {
		paths = paths.Delete(defaultOnEmptyFields...)
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.SpannerBackupScheduleStatus{}
		status.ObservedState = SpannerBackupScheduleObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	desiredPb.Name = a.id.String()
	req := &spannerbackupschedulespb.UpdateBackupScheduleRequest{
		UpdateMask:     updateMask,
		BackupSchedule: desiredPb,
	}
	updated, err := a.gcpClient.UpdateBackupSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating BackupSchedule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated BackupSchedule", "name", a.id)

	status := &krm.SpannerBackupScheduleStatus{}
	status.ObservedState = SpannerBackupScheduleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupScheduleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpannerBackupSchedule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpannerBackupScheduleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.DatabaseRef = &krm.SpannerDatabaseRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.SpannerBackupScheduleGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupScheduleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupSchedule", "name", a.id)

	req := &spannerbackupschedulespb.DeleteBackupScheduleRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteBackupSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BackupSchedule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupSchedule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BackupSchedule", "name", a.id)
	return true, nil
}
