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

package sqladmin

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "cloud.google.com/go/sql/apiv1/sqlpb"
	apiscommon "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	gax "github.com/googleapis/gax-go/v2"
	api "google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SQLAdminBackupGVK, newSQLAdminBackupModel)
}

func newSQLAdminBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sqlAdminBackupModel{config: config}, nil
}

type sqlAdminBackupModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &sqlAdminBackupModel{}

func (m *sqlAdminBackupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type sqlAdminBackupAdapter struct {
	projectID  string
	instanceID string
	backupID   int64

	desired *pb.BackupRun
	actual  *pb.BackupRun

	sqlBackupRunsClient *api.BackupRunsService
	sqlOperationsClient *api.OperationsService
}

var _ directbase.Adapter = &sqlAdminBackupAdapter{}

func (m *sqlAdminBackupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	obj := &krm.SQLAdminBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("converting to %T failed: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, kube, obj)
	if err != nil {
		return nil, fmt.Errorf("resolving project ID: %w", err)
	}

	if obj.Spec.InstanceRef == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}
	instance, err := refsv1beta1.ResolveSQLInstanceRef(ctx, kube, obj, obj.Spec.InstanceRef)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.instanceRef: %w", err)
	}
	instanceID := instance.SQLInstanceName

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	service, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building sqladmin service client: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desiredProto := SQLAdminBackupSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	adapter := &sqlAdminBackupAdapter{
		projectID:           projectID,
		instanceID:          instanceID,
		desired:             desiredProto,
		sqlBackupRunsClient: api.NewBackupRunsService(service),
		sqlOperationsClient: api.NewOperationsService(service),
	}

	externalRef := apiscommon.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		id := &krm.SQLAdminBackupIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		backupID, err := strconv.ParseInt(id.Backup, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing backup ID %q as int64: %w", id.Backup, err)
		}
		adapter.backupID = backupID
	}

	return adapter, nil
}

func (a *sqlAdminBackupAdapter) Find(ctx context.Context) (bool, error) {
	if a.backupID == 0 {
		return false, nil
	}

	actual, err := a.sqlBackupRunsClient.Get(a.projectID, a.instanceID, a.backupID).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting backup run %d: %w", a.backupID, err)
	}

	actualProto, err := toProto(actual)
	if err != nil {
		return false, err
	}

	a.actual = actualProto
	return true, nil
}

func (a *sqlAdminBackupAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating backup run", "project", a.projectID, "instance", a.instanceID)

	desiredAPI, err := toAPI(a.desired)
	if err != nil {
		return err
	}

	insertOp, err := a.sqlBackupRunsClient.Insert(a.projectID, a.instanceID, desiredAPI).Do()
	if err != nil {
		return fmt.Errorf("inserting backup run: %w", err)
	}

	completedOp, err := a.pollForLROCompletion(ctx, insertOp, "create")
	if err != nil {
		return err
	}

	if completedOp.BackupContext == nil {
		return fmt.Errorf("operation completed but BackupContext is nil")
	}
	backupID := completedOp.BackupContext.BackupId
	if backupID == 0 {
		return fmt.Errorf("operation completed but BackupId is 0")
	}
	a.backupID = backupID
	backupIDStr := strconv.FormatInt(backupID, 10)

	actual, err := a.sqlBackupRunsClient.Get(a.projectID, a.instanceID, a.backupID).Do()
	if err != nil {
		return fmt.Errorf("getting created backup run %d: %w", a.backupID, err)
	}

	actualProto, err := toProto(actual)
	if err != nil {
		return err
	}
	a.actual = actualProto

	identity := &krm.SQLAdminBackupIdentity{
		Project: a.projectID,
		Backup:  backupIDStr,
	}
	externalRef := identity.String()

	return a.updateStatus(ctx, op, externalRef, actualProto)
}

func (a *sqlAdminBackupAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	if a.actual == nil {
		return fmt.Errorf("cannot update: actual resource not found")
	}

	diff, err := compareResource(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diff.HasDiff() {
		return fmt.Errorf("SQLAdminBackup is immutable and cannot be updated. Diff: %s", strings.Join(diff.FieldIDs(), ", "))
	}

	identity := &krm.SQLAdminBackupIdentity{
		Project: a.projectID,
		Backup:  strconv.FormatInt(a.backupID, 10),
	}
	externalRef := identity.String()

	return a.updateStatus(ctx, op, externalRef, a.actual)
}

func (a *sqlAdminBackupAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	if a.backupID == 0 {
		return true, nil
	}

	deleteOp, err := a.sqlBackupRunsClient.Delete(a.projectID, a.instanceID, a.backupID).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting backup run %d: %w", a.backupID, err)
	}

	if _, err := a.pollForLROCompletion(ctx, deleteOp, "delete"); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}

	return true, nil
}

func (a *sqlAdminBackupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("SQLAdminBackup not found")
	}

	mapCtx := &direct.MapContext{}
	spec := SQLAdminBackupSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set instanceRef and projectRef manually from a.projectID and a.instanceID
	spec.ProjectRef = &refsv1beta1.ProjectRef{
		External: a.projectID,
	}
	spec.InstanceRef = &refsv1beta1.SQLInstanceRef{
		External: a.instanceID,
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting SQLAdminBackup spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(strconv.FormatInt(a.backupID, 10))
	u.SetGroupVersionKind(krm.SQLAdminBackupGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *sqlAdminBackupAdapter) pollForLROCompletion(ctx context.Context, op *api.Operation, verb string) (*api.Operation, error) {
	log := klog.FromContext(ctx)
	var err error

	pollingBackoff := gax.Backoff{
		Initial:    time.Second,
		Max:        time.Minute,
		Multiplier: 2,
	}
	for {
		log.V(2).Info("polling", "op", op)

		if op == nil {
			return nil, fmt.Errorf("operation is nil while polling for SQLAdminBackup %d %s", a.backupID, verb)
		}

		if op.Status == "DONE" {
			break
		}
		if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
			return nil, fmt.Errorf("waiting for SQLAdminBackup %d %s failed: %w", a.backupID, verb, err)
		}
		opName := op.Name
		op, err = a.sqlOperationsClient.Get(a.projectID, opName).Do()
		if err != nil {
			return nil, fmt.Errorf("getting SQLAdminBackup %d %s operation %s failed: %w", a.backupID, verb, opName, err)
		}
	}

	if op == nil {
		return nil, fmt.Errorf("operation is nil after polling for SQLAdminBackup %d %s", a.backupID, verb)
	}

	if op.Error != nil && len(op.Error.Errors) > 0 {
		return nil, fmt.Errorf("SQLAdminBackup %d %s operation %s failed: %v", a.backupID, verb, op.Name, op.Error.Errors[0].Message)
	}

	return op, nil
}

func (a *sqlAdminBackupAdapter) updateStatus(ctx context.Context, op directbase.Operation, externalRef string, latest *pb.BackupRun) error {
	mapCtx := &direct.MapContext{}
	observedState := SQLAdminBackupObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.SQLAdminBackupStatus{
		ObservedState: observedState,
		ExternalRef:   &externalRef,
	}
	return op.UpdateStatus(ctx, status, nil)
}

func toProto(in *api.BackupRun) (*pb.BackupRun, error) {
	if in == nil {
		return nil, nil
	}
	out := &pb.BackupRun{}
	if err := common.APIToProto(in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func toAPI(in *pb.BackupRun) (*api.BackupRun, error) {
	if in == nil {
		return nil, nil
	}
	out := &api.BackupRun{}
	if err := common.ProtoToAPI(in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func compareResource(ctx context.Context, actual, desired *pb.BackupRun) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, SQLAdminBackupSpec_FromProto, SQLAdminBackupSpec_ToProto)
	if err != nil {
		return nil, err
	}
	clonedDesired := proto.Clone(desired).(*pb.BackupRun)
	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
