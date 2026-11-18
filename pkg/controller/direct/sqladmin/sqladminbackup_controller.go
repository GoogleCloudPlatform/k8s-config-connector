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

// +generated:controller
// krm.group: sqladmin.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SQLAdminBackup

package sqladmin

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/sql/apiv1"
	sqlpb "cloud.google.com/go/sql/apiv1/sqlpb"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SQLAdminBackupGVK, NewBackupRunModel)
}

func NewBackupRunModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBackupRun{config: *config}, nil
}

var _ directbase.Model = &modelBackupRun{}

type modelBackupRun struct {
	config config.ControllerConfig
}

func (m *modelBackupRun) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SQLAdminBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SQLAdminBackupIdentity)

	instanceName := ""
	if obj.Spec.InstanceRef != nil {
		sqlInstance, err := refsv1beta1.ResolveSQLInstanceRef(ctx, reader, obj, obj.Spec.InstanceRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.instanceRef: %w", err)
		}
		if sqlInstance != nil {
			instanceName = sqlInstance.SQLInstanceName
			obj.Spec.InstanceRef.External = instanceName
		}
	}

	// Get sqladmin GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	operationsClient, err := m.operationsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &BackupRunAdapter{
		id:               id,
		instance:         instanceName,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		desired:          obj,
	}, nil
}

func (m *modelBackupRun) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//sqladmin.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SQLAdminBackupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	operationsClient, err := m.operationsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &BackupRunAdapter{
		id:               id,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
	}, nil
}

func (m *modelBackupRun) client(ctx context.Context) (*gcp.SqlBackupRunsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSqlBackupRunsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BackupRun client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelBackupRun) operationsClient(ctx context.Context) (*gcp.SqlOperationsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	operationsClient, err := gcp.NewSqlOperationsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SqlOperations client: %w", err)
	}
	return operationsClient, nil
}

type BackupRunAdapter struct {
	id               *krm.SQLAdminBackupIdentity
	instance         string
	gcpClient        *gcp.SqlBackupRunsClient
	operationsClient *gcp.SqlOperationsClient
	desired          *krm.SQLAdminBackup
	actual           *sqlpb.BackupRun
}

var _ directbase.Adapter = &BackupRunAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BackupRunAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BackupRun", "name", a.id)

	if a.id.Backup == "" {
		return false, nil
	}

	backupID, err := strconv.ParseInt(a.id.Backup, 10, 64)
	if err != nil {
		return false, fmt.Errorf("parsing backup ID %q: %w", a.id.Backup, err)
	}

	if a.instance != "" {
		req := &sqlpb.SqlBackupRunsGetRequest{
			Project:  a.id.Project,
			Instance: a.instance,
			Id:       backupID,
		}
		actual, err := a.gcpClient.Get(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting BackupRun %q: %w", a.id.String(), err)
		}
		a.actual = actual
		return true, nil
	}

	// If instance is not known (e.g. from AdapterForURL), search across all instances in the project.
	it := a.gcpClient.List(ctx, &sqlpb.SqlBackupRunsListRequest{
		Project:  a.id.Project,
		Instance: "-",
	})
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("listing BackupRuns in project %q: %w", a.id.Project, err)
		}
		if item.GetId() == backupID {
			a.actual = item
			a.instance = item.GetInstance()
			return true, nil
		}
	}

	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupRunAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackupRun", "instance", a.instance)

	if a.instance == "" {
		return fmt.Errorf("cannot create BackupRun: instanceRef is required")
	}

	mapCtx := &direct.MapContext{}
	desired := SQLAdminBackupSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Instance must be set on the request body
	desired.Instance = a.instance

	req := &sqlpb.SqlBackupRunsInsertRequest{
		Project:  a.id.Project,
		Instance: a.instance,
		Body:     desired,
	}

	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("inserting BackupRun: %w", err)
	}

	completedOp, err := a.pollOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for insert BackupRun operation: %w", err)
	}

	var backupID int64
	if completedOp.GetBackupContext() != nil && completedOp.GetBackupContext().GetBackupId() != 0 {
		backupID = completedOp.GetBackupContext().GetBackupId()
	} else if completedOp.GetTargetId() != "" {
		if id, parseErr := strconv.ParseInt(completedOp.GetTargetId(), 10, 64); parseErr == nil {
			backupID = id
		}
	}

	if backupID == 0 {
		return fmt.Errorf("could not determine backup run ID from completed operation: %+v", completedOp)
	}

	a.id.Backup = strconv.FormatInt(backupID, 10)

	latest, err := a.gcpClient.Get(ctx, &sqlpb.SqlBackupRunsGetRequest{
		Project:  a.id.Project,
		Instance: a.instance,
		Id:       backupID,
	})
	if err != nil {
		return fmt.Errorf("fetching newly created BackupRun %q: %w", a.id.String(), err)
	}
	a.actual = latest

	status := &krm.SQLAdminBackupStatus{}
	status.ObservedState = SQLAdminBackupObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BackupRunAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BackupRun", "name", a.id)

	normalize := func(ctx context.Context, pb *sqlpb.BackupRun) error {
		if pb.GetInstance() == "" && a.instance != "" {
			pb.Instance = a.instance
		}
		return nil
	}
	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SQLAdminBackupSpec_FromProto,
		SQLAdminBackupSpec_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating SQLAdminBackup is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.SQLAdminBackupStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SQLAdminBackupObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BackupRunAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SQLAdminBackup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SQLAdminBackupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	if a.actual.GetInstance() != "" {
		obj.Spec.InstanceRef = &refsv1beta1.SQLInstanceRef{External: a.actual.GetInstance()}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Backup)
	u.SetGroupVersionKind(krm.SQLAdminBackupGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BackupRunAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackupRun", "name", a.id)

	if a.id.Backup == "" {
		return true, nil
	}

	backupID, err := strconv.ParseInt(a.id.Backup, 10, 64)
	if err != nil {
		return false, fmt.Errorf("parsing backup ID %q: %w", a.id.Backup, err)
	}

	if a.instance == "" {
		found, err := a.Find(ctx)
		if err != nil {
			return false, err
		}
		if !found {
			return true, nil
		}
	}

	req := &sqlpb.SqlBackupRunsDeleteRequest{
		Project:  a.id.Project,
		Instance: a.instance,
		Id:       backupID,
	}

	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BackupRun %q: %w", a.id.String(), err)
	}

	if _, err := a.pollOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for delete BackupRun operation: %w", err)
	}

	return true, nil
}

func (a *BackupRunAdapter) pollOperation(ctx context.Context, op *sqlpb.Operation) (*sqlpb.Operation, error) {
	log := klog.FromContext(ctx)
	var err error

	pollingBackoff := gax.Backoff{
		Initial:    2 * time.Second,
		Max:        30 * time.Second,
		Multiplier: 1.5,
	}

	for {
		log.V(2).Info("polling operation", "op", op.GetName(), "status", op.GetStatus())

		if op.GetStatus() == sqlpb.Operation_DONE {
			if opError := op.GetError(); opError != nil && len(opError.GetErrors()) > 0 {
				var errMsgs []string
				for _, e := range opError.GetErrors() {
					errMsgs = append(errMsgs, e.GetMessage())
				}
				return op, fmt.Errorf("operation failed: %s", strings.Join(errMsgs, "; "))
			}
			return op, nil
		}

		if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
			return op, fmt.Errorf("waiting for operation %s: %w", op.GetName(), err)
		}

		op, err = a.operationsClient.Get(ctx, &sqlpb.SqlOperationsGetRequest{
			Project:   a.id.Project,
			Operation: op.GetName(),
		})
		if err != nil {
			return op, fmt.Errorf("getting operation %s: %w", op.GetName(), err)
		}
	}
}
