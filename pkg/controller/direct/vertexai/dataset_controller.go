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

package vertexai

import (
	"context"
	"fmt"

	refsroot "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIDatasetGVK, NewDatasetModel)
}

func NewDatasetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataset{config: *config}, nil
}

var _ directbase.Model = &modelDataset{}

type modelDataset struct {
	config config.ControllerConfig
}

func (m *modelDataset) client(ctx context.Context, location string) (*gcp.DatasetClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewDatasetRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Dataset client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDataset) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.VertexAIDatasetIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := VertexAIDatasetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &DatasetAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelDataset) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DatasetAdapter struct {
	id        *krm.VertexAIDatasetIdentity
	gcpClient *gcp.DatasetClient
	desired   *pb.Dataset
	actual    *pb.Dataset
}

var _ directbase.Adapter = &DatasetAdapter{}

func (a *DatasetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Dataset", "name", a.id)

	req := &pb.GetDatasetRequest{Name: a.id.String()}
	datasetpb, err := a.gcpClient.GetDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Dataset %q: %w", a.id, err)
	}

	a.actual = datasetpb
	return true, nil
}

func (a *DatasetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Dataset", "name", a.id)

	req := &pb.CreateDatasetRequest{
		Parent:  a.id.ParentString(),
		Dataset: a.desired,
	}
	op, err := a.gcpClient.CreateDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Dataset %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting creation Dataset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Dataset", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *DatasetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Dataset", "name", a.id)

	diffs, updateMask, err := compareDataset(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for Dataset", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	a.desired.Name = a.id.String()
	req := &pb.UpdateDatasetRequest{
		Dataset:    a.desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Dataset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Dataset", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *DatasetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIDatasetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = refsroot.ProjectRef{External: "projects/" + a.id.Project}
	obj.Spec.Region = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Dataset)
	u.SetGroupVersionKind(krm.VertexAIDatasetGVK)

	u.Object = uObj
	return u, nil
}

func (a *DatasetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Dataset", "name", a.id)

	req := &pb.DeleteDatasetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Dataset, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Dataset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Dataset", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Dataset %s: %w", a.id, err)
	}
	return true, nil
}

func compareDataset(ctx context.Context, actual, desired *pb.Dataset) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAIDatasetSpec_FromProto, VertexAIDatasetSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Dataset)

	populateDefaults := func(obj *pb.Dataset) {
		// Populate GCP/server defaults here if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *DatasetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Dataset) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIDatasetStatus{}
	status.ObservedState = VertexAIDatasetObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
