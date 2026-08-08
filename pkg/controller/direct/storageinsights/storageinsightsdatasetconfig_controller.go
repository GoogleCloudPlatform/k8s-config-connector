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
// proto.message: google.cloud.storageinsights.v1.DatasetConfig
// crd.type: StorageInsightsDatasetConfig
// crd.version: v1alpha1

package storageinsights

import (
	"context"
	"fmt"
	"strings"

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
	registry.RegisterModel(krm.StorageInsightsDatasetConfigGVK, NewStorageInsightsDatasetConfigModel)
}

func NewStorageInsightsDatasetConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &storageInsightsDatasetConfigModel{config: *config}, nil
}

var _ directbase.Model = &storageInsightsDatasetConfigModel{}

type storageInsightsDatasetConfigModel struct {
	config config.ControllerConfig
}

func (m *storageInsightsDatasetConfigModel) client(ctx context.Context) (*api.Client, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newStorageInsightsClient(ctx)
}

func (m *storageInsightsDatasetConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageInsightsDatasetConfig{}
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
	id := idRaw.(*krm.StorageInsightsDatasetConfigIdentity)

	if strings.Contains(id.DatasetConfig, "-") {
		return nil, fmt.Errorf("invalid StorageInsightsDatasetConfig name/resourceID %q: GCP Storage Insights API does not allow hyphens in the DatasetConfig ID. Please use spec.resourceID to specify a name with underscores or alphanumeric characters only", id.DatasetConfig)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := StorageInsightsDatasetConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &storageInsightsDatasetConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *storageInsightsDatasetConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.StorageInsightsDatasetConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &storageInsightsDatasetConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type storageInsightsDatasetConfigAdapter struct {
	gcpClient *api.Client
	id        *krm.StorageInsightsDatasetConfigIdentity
	desired   *pb.DatasetConfig
	actual    *pb.DatasetConfig
	reader    client.Reader
}

var _ directbase.Adapter = &storageInsightsDatasetConfigAdapter{}

func (a *storageInsightsDatasetConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting StorageInsightsDatasetConfig", "name", a.id)

	req := &pb.GetDatasetConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDatasetConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting StorageInsightsDatasetConfig %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *storageInsightsDatasetConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating StorageInsightsDatasetConfig", "name", a.id)

	a.desired.Name = a.id.String()

	req := &pb.CreateDatasetConfigRequest{
		Parent:          a.id.ParentString(),
		DatasetConfigId: a.id.DatasetConfig,
		DatasetConfig:   a.desired,
	}
	op, err := a.gcpClient.CreateDatasetConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating StorageInsightsDatasetConfig %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for StorageInsightsDatasetConfig creation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created StorageInsightsDatasetConfig in gcp", "name", a.id)

	// Fetch fully populated resource after LRO finishes
	getReq := &pb.GetDatasetConfigRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetDatasetConfig(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching StorageInsightsDatasetConfig after creation: %w", err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *storageInsightsDatasetConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating StorageInsightsDatasetConfig", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareDatasetConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for StorageInsightsDatasetConfig", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateDatasetConfigRequest{
		DatasetConfig: a.desired,
		UpdateMask:    updateMask,
	}

	op, err := a.gcpClient.UpdateDatasetConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating StorageInsightsDatasetConfig %s: %w", a.id, err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for StorageInsightsDatasetConfig update %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated StorageInsightsDatasetConfig", "name", a.id)

	// Fetch fully populated resource after LRO finishes
	getReq := &pb.GetDatasetConfigRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetDatasetConfig(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching StorageInsightsDatasetConfig after update: %w", err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *storageInsightsDatasetConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.StorageInsightsDatasetConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageInsightsDatasetConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.PtrTo(a.id.DatasetConfig)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.DatasetConfig)
	u.SetGroupVersionKind(krm.StorageInsightsDatasetConfigGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *storageInsightsDatasetConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting StorageInsightsDatasetConfig", "name", a.id)

	req := &pb.DeleteDatasetConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDatasetConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting StorageInsightsDatasetConfig %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for StorageInsightsDatasetConfig deletion %s: %w", a.id, err)
	}

	return true, nil
}

func compareDatasetConfig(ctx context.Context, actual, desired *pb.DatasetConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, StorageInsightsDatasetConfigSpec_FromProto, StorageInsightsDatasetConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	// CloudStorageObjectPath is an INPUT_ONLY field in the GCP API and is not returned in GET responses.
	// Overriding/copying it from desired to maskedActual prevents false-positive diffs during reconciliation.
	if _, ok := desired.SourceOptions.(*pb.DatasetConfig_CloudStorageObjectPath); ok {
		if actual.SourceOptions == nil {
			maskedActual.SourceOptions = desired.SourceOptions
		}
	}

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.DatasetConfig) {
		// Populate server defaults if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *storageInsightsDatasetConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DatasetConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := StorageInsightsDatasetConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status := &krm.StorageInsightsDatasetConfigStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
