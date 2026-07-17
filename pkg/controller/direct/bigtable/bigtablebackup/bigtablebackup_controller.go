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

package bigtablebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigtable"
	adminpb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableBackupGVK, NewBigtableBackupModel)
}

func NewBigtableBackupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableBackup{config: *config}, nil
}

var _ directbase.Model = &modelBigtableBackup{}

type modelBigtableBackup struct {
	config config.ControllerConfig
}

func (m *modelBigtableBackup) client(ctx context.Context, project, instance string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableBackup client options: %w", err)
	}
	gcpClient, err := gcp.NewAdminClient(ctx, project, instance, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableBackup admin client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBigtableBackup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableBackup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Project, id.Instance)
	if err != nil {
		return nil, err
	}

	return &BigtableBackupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelBigtableBackup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type BigtableBackupAdapter struct {
	id        *krm.BackupIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableBackup
	actual    *gcp.BackupInfo
	reader    client.Reader
}

var _ directbase.Adapter = &BigtableBackupAdapter{}

func (a *BigtableBackupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableBackup", "name", a.id)

	info, err := a.gcpClient.BackupInfo(ctx, a.id.Cluster, a.id.Backup)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NotFound") {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableBackup %q: %w", a.id, err)
	}

	a.actual = info
	return true, nil
}

func parseTableID(tableRef string) (string, error) {
	tokens := strings.Split(tableRef, "/")
	if len(tokens) != 6 || tokens[4] != "tables" {
		return "", fmt.Errorf("unexpected format for table reference: %s", tableRef)
	}
	return tokens[5], nil
}

func (a *BigtableBackupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableBackup", "name", a.id)

	// Resolve sourceTableRef
	if a.desired.Spec.SourceTableRef == nil {
		return fmt.Errorf("spec.sourceTableRef is required")
	}
	tableRef, err := a.desired.Spec.SourceTableRef.NormalizedExternal(ctx, a.reader, a.desired.GetNamespace())
	if err != nil {
		return fmt.Errorf("resolving sourceTableRef: %w", err)
	}
	tableID, err := parseTableID(tableRef)
	if err != nil {
		return err
	}

	// Resolve ExpireTime
	if a.desired.Spec.ExpireTime == nil {
		return fmt.Errorf("spec.expireTime is required")
	}
	expireTime, err := time.Parse(time.RFC3339, *a.desired.Spec.ExpireTime)
	if err != nil {
		return fmt.Errorf("parsing spec.expireTime: %w", err)
	}

	var opts []gcp.BackupOption
	opts = append(opts, gcp.WithExpiry(expireTime))

	if a.desired.Spec.BackupType != nil {
		switch *a.desired.Spec.BackupType {
		case "HOT":
			opts = append(opts, gcp.WithHotBackup())
		}
	}
	if a.desired.Spec.HotToStandardTime != nil {
		hotToStandardTime, err := time.Parse(time.RFC3339, *a.desired.Spec.HotToStandardTime)
		if err != nil {
			return fmt.Errorf("parsing spec.hotToStandardTime: %w", err)
		}
		opts = append(opts, gcp.WithHotToStandardBackup(hotToStandardTime))
	}

	err = a.gcpClient.CreateBackupWithOptions(ctx, tableID, a.id.Cluster, a.id.Backup, opts...)
	if err != nil {
		return fmt.Errorf("creating BigtableBackup %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created BigtableBackup", "name", a.id)

	info, err := a.gcpClient.BackupInfo(ctx, a.id.Cluster, a.id.Backup)
	if err != nil {
		return fmt.Errorf("getting created BigtableBackup info: %w", err)
	}

	status := &krm.BigtableBackupStatus{}
	if err := a.updateStatus(ctx, status, info); err != nil {
		return err
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *BigtableBackupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableBackup", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := BigtableBackupSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	actualPb := convertBackupInfoToProto(a.actual)

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, actualPb, func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
		if fieldName != "expire_time" && fieldName != "hot_to_standard_time" {
			return false, nil
		}
		return common.BasicDiff(fieldName, a, b)
	})
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	if paths.Has("expire_time") {
		desiredExpireTime, err := time.Parse(time.RFC3339, *a.desired.Spec.ExpireTime)
		if err != nil {
			return fmt.Errorf("parsing spec.expireTime: %w", err)
		}
		err = a.gcpClient.UpdateBackup(ctx, a.id.Cluster, a.id.Backup, desiredExpireTime)
		if err != nil {
			return fmt.Errorf("updating BigtableBackup expireTime %s: %w", a.id, err)
		}
	}

	if paths.Has("hot_to_standard_time") {
		if a.desired.Spec.HotToStandardTime != nil {
			desiredHotTime, err := time.Parse(time.RFC3339, *a.desired.Spec.HotToStandardTime)
			if err != nil {
				return fmt.Errorf("parsing spec.hotToStandardTime: %w", err)
			}
			err = a.gcpClient.UpdateBackupHotToStandardTime(ctx, a.id.Cluster, a.id.Backup, desiredHotTime)
			if err != nil {
				return fmt.Errorf("updating BigtableBackup hotToStandardTime %s: %w", a.id, err)
			}
		} else {
			err := a.gcpClient.UpdateBackupRemoveHotToStandardTime(ctx, a.id.Cluster, a.id.Backup)
			if err != nil {
				return fmt.Errorf("removing BigtableBackup hotToStandardTime %s: %w", a.id, err)
			}
		}
	}

	info, err := a.gcpClient.BackupInfo(ctx, a.id.Cluster, a.id.Backup)
	if err != nil {
		return fmt.Errorf("getting updated BigtableBackup info: %w", err)
	}

	status := &krm.BigtableBackupStatus{}
	if err := a.updateStatus(ctx, status, info); err != nil {
		return err
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *BigtableBackupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableBackup", "name", a.id)

	err := a.gcpClient.DeleteBackup(ctx, a.id.Cluster, a.id.Backup)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NotFound") {
			return false, nil
		}
		return false, fmt.Errorf("deleting BigtableBackup %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableBackup", "name", a.id)
	return true, nil
}

func (a *BigtableBackupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func convertBackupInfoToProto(info *gcp.BackupInfo) *adminpb.Backup {
	if info == nil {
		return nil
	}
	pbBackup := &adminpb.Backup{
		Name:         info.Name,
		SourceTable:  info.SourceTable,
		SourceBackup: info.SourceBackup,
		SizeBytes:    info.SizeBytes,
		ExpireTime:   timestamppb.New(info.ExpireTime),
		StartTime:    timestamppb.New(info.StartTime),
		EndTime:      timestamppb.New(info.EndTime),
	}

	switch info.State {
	case "CREATING":
		pbBackup.State = adminpb.Backup_CREATING
	case "READY":
		pbBackup.State = adminpb.Backup_READY
	default:
		pbBackup.State = adminpb.Backup_STATE_UNSPECIFIED
	}

	switch info.BackupType {
	case gcp.BackupTypeStandard:
		pbBackup.BackupType = adminpb.Backup_STANDARD
	case gcp.BackupTypeHot:
		pbBackup.BackupType = adminpb.Backup_HOT
	}

	if info.HotToStandardTime != nil {
		pbBackup.HotToStandardTime = timestamppb.New(*info.HotToStandardTime)
	}

	if info.EncryptionInfo != nil {
		pbBackup.EncryptionInfo = &adminpb.EncryptionInfo{
			KmsKeyVersion: info.EncryptionInfo.KMSKeyVersion,
		}
		switch info.EncryptionInfo.Type {
		case gcp.GoogleDefaultEncryption:
			pbBackup.EncryptionInfo.EncryptionType = adminpb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION
		case gcp.CustomerManagedEncryption:
			pbBackup.EncryptionInfo.EncryptionType = adminpb.EncryptionInfo_CUSTOMER_MANAGED_ENCRYPTION
		}
	}

	return pbBackup
}

func (a *BigtableBackupAdapter) updateStatus(ctx context.Context, status *krm.BigtableBackupStatus, info *gcp.BackupInfo) error {
	mapCtx := &direct.MapContext{}
	pbBackup := convertBackupInfoToProto(info)
	status.ObservedState = BigtableBackupObservedState_FromProto(mapCtx, pbBackup)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return nil
}
