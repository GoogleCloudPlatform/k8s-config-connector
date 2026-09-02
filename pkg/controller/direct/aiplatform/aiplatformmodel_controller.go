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

package aiplatform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformModelGVK, NewAIPlatformModelModel)
}

func NewAIPlatformModelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &aiplatformModelModel{config: config}, nil
}

var _ directbase.Model = &aiplatformModelModel{}

type aiplatformModelModel struct {
	config *config.ControllerConfig
}

func (m *aiplatformModelModel) client(ctx context.Context, location string) (*gcp.ModelClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewModelClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelClient client: %w", err)
	}
	return gcpClient, nil
}

func (m *aiplatformModelModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.AIPlatformModel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(reader.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	typedID, ok := id.(*krm.AIPlatformModelIdentity)
	if !ok {
		return nil, fmt.Errorf("expected AIPlatformModelIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx, typedID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := AIPlatformModelSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &AIPlatformModelAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *aiplatformModelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AIPlatformModelIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &AIPlatformModelAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AIPlatformModelAdapter struct {
	id        *krm.AIPlatformModelIdentity
	gcpClient *gcp.ModelClient
	desiredpb *pb.Model
	desired   *krm.AIPlatformModel
	actual    *pb.Model
}

var _ directbase.Adapter = &AIPlatformModelAdapter{}

func (a *AIPlatformModelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AIPlatformModel", "name", a.id.String())

	req := &pb.GetModelRequest{
		Name: a.id.String(),
	}

	modelpb, err := a.gcpClient.GetModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIPlatformModel %q: %w", a.id.String(), err)
	}

	a.actual = modelpb

	mapCtx := &direct.MapContext{}
	observedState := AIPlatformModelObservedState_FromProto(mapCtx, modelpb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}

	if a.desired != nil {
		a.desired.Status.ObservedState = observedState
		a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return true, nil
}

func (a *AIPlatformModelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIPlatformModel", "name", a.id.String())

	// Ensure Name is empty/set correctly according to GCP REST API patterns
	a.desiredpb.Name = ""

	req := &pb.UploadModelRequest{
		Parent:  a.id.ParentString(),
		ModelId: a.id.Model,
		Model:   a.desiredpb,
	}

	op, err := a.gcpClient.UploadModel(ctx, req)
	if err != nil {
		return fmt.Errorf("uploading AIPlatformModel %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully started upload of AIPlatformModel", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformModel %q upload: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully completed upload of AIPlatformModel", "name", a.id.String())

	// Fetch fully-populated resource immediately after LRO success
	getReq := &pb.GetModelRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetModel(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created AIPlatformModel %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *AIPlatformModelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIPlatformModel", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, updateMask, err := compareModel(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateModelRequest{
		Model:      a.desiredpb,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateModel(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AIPlatformModel %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated AIPlatformModel", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *AIPlatformModelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformModel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformModelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if obj.Spec.Parent == nil {
		obj.Spec.Parent = &krm.Parent{}
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Model

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.Model)
	u.SetGroupVersionKind(krm.AIPlatformModelGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *AIPlatformModelAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Model) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformModelStatus{}
	status.ObservedState = AIPlatformModelObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AIPlatformModelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIPlatformModel", "name", a.id.String())

	req := &pb.DeleteModelRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AIPlatformModel %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of AIPlatformModel %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareModel(ctx context.Context, actual, desired *pb.Model) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AIPlatformModelSpec_FromProto, AIPlatformModelSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	// Copy immutable fields from actual to desired to avoid invalid update requests.
	// Only displayName, description, and labels are mutable via UpdateModel.
	clonedDesired.ContainerSpec = maskedActual.ContainerSpec
	clonedDesired.ArtifactUri = maskedActual.ArtifactUri
	clonedDesired.VersionDescription = maskedActual.VersionDescription
	clonedDesired.VersionAliases = maskedActual.VersionAliases
	clonedDesired.PredictSchemata = maskedActual.PredictSchemata
	clonedDesired.MetadataSchemaUri = maskedActual.MetadataSchemaUri
	clonedDesired.Metadata = maskedActual.Metadata
	clonedDesired.PipelineJob = maskedActual.PipelineJob
	clonedDesired.DataStats = maskedActual.DataStats
	clonedDesired.EncryptionSpec = maskedActual.EncryptionSpec
	clonedDesired.BaseModelSource = maskedActual.BaseModelSource
	clonedDesired.ExplanationSpec = maskedActual.ExplanationSpec

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
