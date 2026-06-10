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
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformModelGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelModel{config: *config}, nil
}

var _ directbase.Model = &modelModel{}

type modelModel struct {
	config config.ControllerConfig
}

func (m *modelModel) client(ctx context.Context, location string) (*gcp.ModelClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewModelClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building AIPlatform Model client: %w", err)
	}
	return gcpClient, err
}

func (m *modelModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformModel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idValue, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idValue.(*krm.AIPlatformModelIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idValue)
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &ModelAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ModelAdapter struct {
	id        *krm.AIPlatformModelIdentity
	gcpClient *gcp.ModelClient
	desired   *krm.AIPlatformModel
	actual    *pb.Model
}

var _ directbase.Adapter = &ModelAdapter{}

func (a *ModelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AIPlatformModel", "name", a.id)

	req := &pb.GetModelRequest{Name: a.id.String()}
	modelpb, err := a.gcpClient.GetModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIPlatformModel %q: %w", a.id, err)
	}

	a.actual = modelpb
	return true, nil
}

func (a *ModelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIPlatformModel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformModelSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.UploadModelRequest{
		Parent:  fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Model:   resource,
		ModelId: a.id.Model,
	}
	op, err := a.gcpClient.UploadModel(ctx, req)
	if err != nil {
		return fmt.Errorf("AIPlatformModel %s: %w", a.id, err)
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AIPlatformModel %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully uploaded AIPlatformModel", "name", a.id, "gcpModel", resp.GetModel())

	// Retrieve the created model so we get all fields populated correctly
	obj, err := a.gcpClient.GetModel(ctx, &pb.GetModelRequest{Name: a.id.String()})
	if err != nil {
		return err
	}

	status := &krm.AIPlatformModelStatus{}
	status.ObservedState = AIPlatformModelObservedState_FromProto(mapCtx, obj)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ModelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIPlatformModel", "name", a.id.String())

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformModelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, report, err := common.CompareProtoMessageStructuredDiff(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	// Filter paths and report fields to only include updateable fields
	var updatePaths []string
	var filteredFields []structuredreporting.DiffField
	for _, d := range report.Fields {
		if d.ID == "display_name" || d.ID == "description" || d.ID == "labels" || strings.HasPrefix(d.ID, "labels.") {
			filteredFields = append(filteredFields, d)
		}
	}
	for path := range paths {
		if path == "display_name" || path == "description" || path == "labels" || strings.HasPrefix(path, "labels.") {
			updatePaths = append(updatePaths, path)
		}
	}

	if len(updatePaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.AIPlatformModelStatus{}
		status.ObservedState = AIPlatformModelObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	filteredReport := &structuredreporting.Diff{
		Object: updateOp.GetUnstructured(),
		Fields: filteredFields,
	}
	structuredreporting.ReportDiff(ctx, filteredReport)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: updatePaths,
	}

	resource.Name = a.id.String()
	req := &pb.UpdateModelRequest{
		Model:      resource,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateModel(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AIPlatformModel %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated AIPlatformModel", "name", a.id)

	status := &krm.AIPlatformModelStatus{}
	status.ObservedState = AIPlatformModelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ModelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.AIPlatformModelGVK)

	u.Object = uObj
	return u, nil
}

func (a *ModelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIPlatformModel", "name", a.id)

	req := &pb.DeleteModelRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent AIPlatformModel, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting AIPlatformModel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AIPlatformModel", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete AIPlatformModel %s: %w", a.id, err)
	}
	return true, nil
}
