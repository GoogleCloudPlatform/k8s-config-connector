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
// proto.service: google.cloud.storageinsights.v1.StorageInsights
// proto.message: google.cloud.storageinsights.v1.ReportConfig
// crd.type: StorageInsightsReportConfig
// crd.version: v1alpha1

package storageinsights

import (
	"context"
	"fmt"
	"slices"

	api "cloud.google.com/go/storageinsights/apiv1"
	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storageinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.StorageInsightsReportConfigGVK, NewStorageInsightsReportConfigModel)
}

func NewStorageInsightsReportConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &storageInsightsReportConfigModel{config: *config}, nil
}

var _ directbase.Model = &storageInsightsReportConfigModel{}

type storageInsightsReportConfigModel struct {
	config config.ControllerConfig
}

func (m *storageInsightsReportConfigModel) client(ctx context.Context) (*api.Client, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newStorageInsightsClient(ctx)
}

func (m *storageInsightsReportConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageInsightsReportConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.StorageInsightsReportConfigIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := StorageInsightsReportConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &storageInsightsReportConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *storageInsightsReportConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.StorageInsightsReportConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &storageInsightsReportConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type storageInsightsReportConfigAdapter struct {
	gcpClient *api.Client
	id        *krm.StorageInsightsReportConfigIdentity
	desired   *pb.ReportConfig
	actual    *pb.ReportConfig
	reader    client.Reader
}

var _ directbase.Adapter = &storageInsightsReportConfigAdapter{}

func (a *storageInsightsReportConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting StorageInsightsReportConfig", "name", a.id)

	if a.id.ReportConfig == "" {
		return false, nil
	}

	req := &pb.GetReportConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetReportConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting StorageInsightsReportConfig %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *storageInsightsReportConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating StorageInsightsReportConfig", "name", a.id)

	if a.id.ReportConfig != "" {
		a.desired.Name = a.id.String()
	}

	req := &pb.CreateReportConfigRequest{
		Parent:       a.id.ParentString(),
		ReportConfig: a.desired,
	}
	created, err := a.gcpClient.CreateReportConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating StorageInsightsReportConfig %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created StorageInsightsReportConfig in gcp", "name", created.GetName())

	if a.id.ReportConfig == "" {
		if err := a.id.FromExternal(created.GetName()); err != nil {
			return fmt.Errorf("parsing created StorageInsightsReportConfig name %q: %w", created.GetName(), err)
		}
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *storageInsightsReportConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating StorageInsightsReportConfig", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareReportConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for StorageInsightsReportConfig", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateReportConfigRequest{
		ReportConfig: a.desired,
		UpdateMask:   updateMask,
	}

	updated, err := a.gcpClient.UpdateReportConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating StorageInsightsReportConfig %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated StorageInsightsReportConfig", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *storageInsightsReportConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.StorageInsightsReportConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageInsightsReportConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.PtrTo(a.id.ReportConfig)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ReportConfig)
	u.SetGroupVersionKind(krm.StorageInsightsReportConfigGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *storageInsightsReportConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting StorageInsightsReportConfig", "name", a.id)

	if a.id.ReportConfig == "" {
		return true, nil
	}

	req := &pb.DeleteReportConfigRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteReportConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting StorageInsightsReportConfig %s: %w", a.id, err)
	}

	return true, nil
}

func compareReportConfig(ctx context.Context, actual, desired *pb.ReportConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, StorageInsightsReportConfigSpec_FromProto, StorageInsightsReportConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	diffs, topLevelUpdateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	updateMask := &fieldmaskpb.FieldMask{}
	for _, path := range topLevelUpdateMask.Paths {
		if path == "object_metadata_report_options" {
			desiredOpt := clonedDesired.GetObjectMetadataReportOptions()
			actualOpt := maskedActual.GetObjectMetadataReportOptions()
			if desiredOpt != nil && actualOpt != nil {
				if !slices.Equal(desiredOpt.GetMetadataFields(), actualOpt.GetMetadataFields()) {
					updateMask.Paths = append(updateMask.Paths, "object_metadata_report_options.metadata_fields")
				}
				if desiredOpt.GetStorageDestinationOptions().GetBucket() != actualOpt.GetStorageDestinationOptions().GetBucket() {
					updateMask.Paths = append(updateMask.Paths, "object_metadata_report_options.storage_destination_options.bucket")
				}
				if desiredOpt.GetStorageDestinationOptions().GetDestinationPath() != actualOpt.GetStorageDestinationOptions().GetDestinationPath() {
					updateMask.Paths = append(updateMask.Paths, "object_metadata_report_options.storage_destination_options.destination_path")
				}
				if desiredOpt.GetStorageFilters().GetBucket() != actualOpt.GetStorageFilters().GetBucket() {
					return nil, nil, fmt.Errorf("objectMetadataReportOptions.storageFilters.bucket cannot be updated")
				}
			} else {
				updateMask.Paths = append(updateMask.Paths, path)
			}
		} else {
			updateMask.Paths = append(updateMask.Paths, path)
		}
	}

	return diffs, updateMask, nil
}

func (a *storageInsightsReportConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ReportConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := StorageInsightsReportConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status := &krm.StorageInsightsReportConfigStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
