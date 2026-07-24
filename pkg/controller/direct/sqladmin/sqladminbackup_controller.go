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

// +tool:controller
// proto.service: google.cloud.sql.v1beta4
// proto.message: google.cloud.sql.v1beta4.BackupRun
// crd.type: SQLAdminBackup
// crd.version: v1alpha1

package sqladmin

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	api "google.golang.org/api/sqladmin/v1beta4"
	"google.golang.org/protobuf/types/known/timestamppb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.SQLAdminBackupGVK, NewBackupModel)
}

func NewBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backupModel{config: *config}, nil
}

var _ directbase.Model = &backupModel{}

type backupModel struct {
	config config.ControllerConfig
}

type gcpClient struct {
	config  config.ControllerConfig
	service *api.Service
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient.service, err = api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building gcp service client: %w", err)
	}

	return gcpClient, nil
}

func (m *gcpClient) sqlBackupRunsClient() *api.BackupRunsService {
	return api.NewBackupRunsService(m.service)
}

func (m *gcpClient) sqlOperationsClient() *api.OperationsService {
	return api.NewOperationsService(m.service)
}

func (m *backupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SQLAdminBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	backupId := id.(*krm.SQLAdminBackupIdentity)

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}

	resolvedInstance, err := refsv1beta1.ResolveSQLInstanceRef(ctx, reader, obj, obj.Spec.InstanceRef)
	if err != nil {
		return nil, fmt.Errorf("resolving instanceRef: %w", err)
	}
	obj.Spec.InstanceRef.External = resolvedInstance.String()

	if obj.Spec.DiskEncryptionConfiguration != nil && obj.Spec.DiskEncryptionConfiguration.KMSKeyRef != nil {
		resolvedKMSKey, err := refsv1beta1.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.DiskEncryptionConfiguration.KMSKeyRef)
		if err != nil {
			return nil, fmt.Errorf("resolving KMSCryptoKeyRef: %w", err)
		}
		obj.Spec.DiskEncryptionConfiguration.KMSKeyRef.External = resolvedKMSKey.External
	}

	project := resolvedInstance.ProjectID
	instance := resolvedInstance.SQLInstanceName

	mapCtx := &direct.MapContext{}
	desired := SQLAdminBackupSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	var actual *pb.BackupRun
	if obj.Status.ObservedState != nil {
		actual = SQLAdminBackupObservedState_ToProto(mapCtx, obj.Status.ObservedState)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &backupAdapter{
		sqlBackupRunsClient: gcpClient.sqlBackupRunsClient(),
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		id:                  backupId,
		project:             project,
		instance:            instance,
		desired:             desired,
		actual:              actual,
	}, nil
}

func (m *backupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.SQLAdminBackupIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}

	return &backupAdapter{
		sqlBackupRunsClient: gcpClient.sqlBackupRunsClient(),
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		id:                  id,
		project:             id.Project,
	}, nil
}

type backupAdapter struct {
	sqlBackupRunsClient *api.BackupRunsService
	sqlOperationsClient *api.OperationsService
	id                  *krm.SQLAdminBackupIdentity
	project             string
	instance            string
	desired             *pb.BackupRun
	actual              *pb.BackupRun
}

var _ directbase.Adapter = &backupAdapter{}

func (a *backupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding SQLAdminBackup", "id", a.id)

	if a.instance == "" {
		return false, fmt.Errorf("instance name is not specified or resolved")
	}

	// Try using numeric ID if specified as Backup name or present in status
	idInt, err := strconv.ParseInt(a.id.Backup, 10, 64)
	var run *api.BackupRun
	if err == nil {
		run, err = a.sqlBackupRunsClient.Get(a.project, a.instance, idInt).Do()
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting SQLAdminBackup %d: %w", idInt, err)
		}
	} else {
		// If Backup name is a string, check if we have the numeric ID in actual status
		if a.actual != nil && a.actual.Id != 0 {
			idInt = a.actual.Id
			run, err = a.sqlBackupRunsClient.Get(a.project, a.instance, idInt).Do()
			if err != nil {
				if direct.IsNotFound(err) {
					// Fall back to description search if not found
					run = nil
				} else {
					return false, fmt.Errorf("getting SQLAdminBackup %d from status: %w", idInt, err)
				}
			}
		}

		if run == nil {
			// Find by description
			run, err = a.findBackupByDescription(ctx, a.id.Backup)
			if err != nil {
				return false, err
			}
		}
	}

	if run == nil {
		log.V(2).Info("SQLAdminBackup not found", "id", a.id)
		return false, nil
	}

	pbRun, err := convertAPIToPBBackupRun(run)
	if err != nil {
		return false, fmt.Errorf("converting backupRun: %w", err)
	}

	a.actual = pbRun
	return true, nil
}

func (a *backupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SQLAdminBackup", "id", a.id)

	apiRun, err := convertPBToAPIBackupRun(a.desired)
	if err != nil {
		return fmt.Errorf("converting desired backupRun: %w", err)
	}

	// Force type to ON_DEMAND for user created backups, and set description to identify it
	apiRun.Type = "ON_DEMAND"
	if apiRun.Description == "" {
		apiRun.Description = a.id.Backup
	}

	op, err := a.sqlBackupRunsClient.Insert(a.project, a.instance, apiRun).Do()
	if err != nil {
		return fmt.Errorf("inserting SQLAdminBackup: %w", err)
	}

	op, err = a.pollForLROCompletion(ctx, op, "CREATE")
	if err != nil {
		return err
	}

	// Retrieve the created backup run
	var idInt int64
	if op.TargetId != "" {
		idInt, err = strconv.ParseInt(op.TargetId, 10, 64)
	}
	if err != nil || idInt == 0 {
		idInt, err = parseBackupRunID(op.TargetLink)
	}
	if err != nil || idInt == 0 {
		// Fallback to searching by description
		run, err := a.findBackupByDescription(ctx, apiRun.Description)
		if err != nil || run == nil {
			return fmt.Errorf("failed to retrieve created backupRun: %w", err)
		}
		idInt = run.Id
	}

	run, err := a.sqlBackupRunsClient.Get(a.project, a.instance, idInt).Do()
	if err != nil {
		return fmt.Errorf("getting created SQLAdminBackup %d: %w", idInt, err)
	}

	pbRun, err := convertAPIToPBBackupRun(run)
	if err != nil {
		return fmt.Errorf("converting created backupRun: %w", err)
	}

	a.actual = pbRun
	log.V(2).Info("successfully created SQLAdminBackup", "id", a.id, "gcpId", idInt)
	return a.updateStatus(ctx, createOp, pbRun)
}

func (a *backupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SQLAdminBackup", "id", a.id)

	maskedActual, err := mappers.OnlySpecFields(a.actual, SQLAdminBackupSpec_FromProto, SQLAdminBackupSpec_ToProto)
	if err != nil {
		return err
	}

	diffs, _, err := tags.DiffForTopLevelFields(ctx, a.desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("SQLAdminBackup is immutable and cannot be updated")
	}

	log.V(2).Info("no diff detected for SQLAdminBackup", "id", a.id)
	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *backupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting SQLAdminBackup", "id", a.id)

	idInt, err := strconv.ParseInt(a.id.Backup, 10, 64)
	if err != nil {
		// Look up by description
		run, err := a.findBackupByDescription(ctx, a.id.Backup)
		if err != nil {
			return false, err
		}
		if run == nil {
			return true, nil
		}
		idInt = run.Id
	}

	op, err := a.sqlBackupRunsClient.Delete(a.project, a.instance, idInt).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting SQLAdminBackup %d: %w", idInt, err)
	}

	_, err = a.pollForLROCompletion(ctx, op, "DELETE")
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}

	log.Info("successfully deleted SQLAdminBackup", "id", a.id, "gcpId", idInt)
	return true, nil
}

func (a *backupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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

	if a.id.Project != "" {
		obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	}
	if a.instance != "" {
		obj.Spec.InstanceRef = &refsv1beta1.SQLInstanceRef{External: fmt.Sprintf("projects/%s/locations/%s/instances/%s", a.project, a.id.Backup, a.instance)}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Backup)
	u.SetGroupVersionKind(krm.SQLAdminBackupGVK)

	u.Object = uObj
	return u, nil
}

func (a *backupAdapter) pollForLROCompletion(ctx context.Context, op *api.Operation, verb string) (*api.Operation, error) {
	for {
		latestOp, err := a.sqlOperationsClient.Get(a.project, op.Name).Do()
		if err != nil {
			return nil, fmt.Errorf("getting SQLAdminBackup %s operation %s failed: %w", verb, op.Name, err)
		}
		if latestOp.Status == "DONE" {
			if latestOp.Error != nil && len(latestOp.Error.Errors) > 0 {
				return latestOp, fmt.Errorf("SQLAdminBackup %s operation %s failed: %s", verb, op.Name, latestOp.Error.Errors[0].Message)
			}
			return latestOp, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}

func (a *backupAdapter) findBackupByDescription(ctx context.Context, desc string) (*api.BackupRun, error) {
	list, err := a.sqlBackupRunsClient.List(a.project, a.instance).Do()
	if err != nil {
		return nil, fmt.Errorf("listing backup runs: %w", err)
	}
	for _, run := range list.Items {
		if run.Description == desc {
			return run, nil
		}
	}
	return nil, nil
}

func (a *backupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.BackupRun) error {
	mapCtx := &direct.MapContext{}
	observedState := SQLAdminBackupObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.SQLAdminBackupStatus{
		ObservedState: observedState,
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func getProjectAndInstance(id *krm.SQLAdminBackupIdentity, instanceRef *refsv1beta1.SQLInstanceRef) (string, string, error) {
	if instanceRef == nil || instanceRef.External == "" {
		return "", "", fmt.Errorf("instanceRef.external is required")
	}
	if !strings.Contains(instanceRef.External, "/") {
		return id.Project, instanceRef.External, nil
	}
	tokens := strings.Split(instanceRef.External, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		return tokens[1], tokens[5], nil
	}
	// Also support projects/{project}/instances/{instance} format if any
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instances" {
		return tokens[1], tokens[3], nil
	}
	return "", "", fmt.Errorf("invalid sqlinstance external identifier format %q", instanceRef.External)
}

func parseBackupRunID(targetLink string) (int64, error) {
	u, err := url.Parse(targetLink)
	if err != nil {
		return 0, err
	}
	path := u.Path
	idx := strings.LastIndex(path, "/backupRuns/")
	if idx == -1 {
		return 0, fmt.Errorf("invalid backupRun targetLink: %s", targetLink)
	}
	idStr := path[idx+len("/backupRuns/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing backupRun ID from %s: %w", idStr, err)
	}
	return id, nil
}

func convertAPIToPBBackupRun(in *api.BackupRun) (*pb.BackupRun, error) {
	if in == nil {
		return nil, nil
	}
	out := &pb.BackupRun{
		Id:          in.Id,
		Instance:    in.Instance,
		Description: in.Description,
		Location:    in.Location,
		TimeZone:    in.TimeZone,
	}

	if in.Status != "" {
		if val, ok := pb.SqlBackupRunStatus_value[in.Status]; ok {
			out.Status = pb.SqlBackupRunStatus(val)
		}
	}
	if in.Type != "" {
		if val, ok := pb.SqlBackupRunType_value[in.Type]; ok {
			out.Type = pb.SqlBackupRunType(val)
		}
	}
	if in.BackupKind != "" {
		if val, ok := pb.SqlBackupKind_value[in.BackupKind]; ok {
			out.BackupKind = pb.SqlBackupKind(val)
		}
	}

	if in.StartTime != "" {
		t, err := time.Parse(time.RFC3339, in.StartTime)
		if err == nil {
			out.StartTime = timestamppb.New(t)
		}
	}
	if in.EndTime != "" {
		t, err := time.Parse(time.RFC3339, in.EndTime)
		if err == nil {
			out.EndTime = timestamppb.New(t)
		}
	}
	if in.EnqueuedTime != "" {
		t, err := time.Parse(time.RFC3339, in.EnqueuedTime)
		if err == nil {
			out.EnqueuedTime = timestamppb.New(t)
		}
	}
	if in.WindowStartTime != "" {
		t, err := time.Parse(time.RFC3339, in.WindowStartTime)
		if err == nil {
			out.WindowStartTime = timestamppb.New(t)
		}
	}

	if in.DiskEncryptionConfiguration != nil {
		out.DiskEncryptionConfiguration = &pb.DiskEncryptionConfiguration{
			Kind:       in.DiskEncryptionConfiguration.Kind,
			KmsKeyName: in.DiskEncryptionConfiguration.KmsKeyName,
		}
	}

	if in.DiskEncryptionStatus != nil {
		out.DiskEncryptionStatus = &pb.DiskEncryptionStatus{
			Kind:              in.DiskEncryptionStatus.Kind,
			KmsKeyVersionName: in.DiskEncryptionStatus.KmsKeyVersionName,
		}
	}

	if in.Error != nil {
		out.Error = &pb.OperationError{
			Kind:    in.Error.Kind,
			Code:    in.Error.Code,
			Message: in.Error.Message,
		}
	}

	return out, nil
}

func convertPBToAPIBackupRun(in *pb.BackupRun) (*api.BackupRun, error) {
	if in == nil {
		return nil, nil
	}
	out := &api.BackupRun{
		Id:          in.Id,
		Instance:    in.Instance,
		Description: in.Description,
		Location:    in.Location,
		Status:      in.Status.String(),
		Type:        in.Type.String(),
		TimeZone:    in.TimeZone,
		BackupKind:  in.BackupKind.String(),
	}

	if in.StartTime != nil {
		out.StartTime = in.StartTime.AsTime().Format(time.RFC3339)
	}
	if in.EndTime != nil {
		out.EndTime = in.EndTime.AsTime().Format(time.RFC3339)
	}
	if in.EnqueuedTime != nil {
		out.EnqueuedTime = in.EnqueuedTime.AsTime().Format(time.RFC3339)
	}
	if in.WindowStartTime != nil {
		out.WindowStartTime = in.WindowStartTime.AsTime().Format(time.RFC3339)
	}

	if in.DiskEncryptionConfiguration != nil {
		out.DiskEncryptionConfiguration = &api.DiskEncryptionConfiguration{
			Kind:       in.DiskEncryptionConfiguration.Kind,
			KmsKeyName: in.DiskEncryptionConfiguration.KmsKeyName,
		}
	}

	if in.DiskEncryptionStatus != nil {
		out.DiskEncryptionStatus = &api.DiskEncryptionStatus{
			Kind:              in.DiskEncryptionStatus.Kind,
			KmsKeyVersionName: in.DiskEncryptionStatus.KmsKeyVersionName,
		}
	}

	if in.Error != nil {
		out.Error = &api.OperationError{
			Kind:    in.Error.Kind,
			Code:    in.Error.Code,
			Message: in.Error.Message,
		}
	}

	return out, nil
}
