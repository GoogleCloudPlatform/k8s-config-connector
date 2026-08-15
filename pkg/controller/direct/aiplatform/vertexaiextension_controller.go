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

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	projects "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	registry.RegisterModel(krm.VertexAIExtensionGVK, NewExtensionModel)
}

func NewExtensionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &extensionModel{config: config}, nil
}

var _ directbase.Model = &extensionModel{}

type extensionModel struct {
	config *config.ControllerConfig
}

func (m *extensionModel) client(ctx context.Context, location string) (*gcp.ExtensionRegistryClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewExtensionRegistryClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ExtensionRegistryClient: %w", err)
	}
	return gcpClient, nil
}

func (m *extensionModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAIExtension{}
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

	typedID, ok := id.(*krm.VertexAIExtensionIdentity)
	if !ok {
		return nil, fmt.Errorf("expected VertexAIExtensionIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx, typedID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := VertexAIExtensionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &ExtensionAdapter{
		id:            typedID,
		gcpClient:     gcpClient,
		desiredpb:     desiredpb,
		desired:       obj,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *extensionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type ExtensionAdapter struct {
	id            *krm.VertexAIExtensionIdentity
	gcpClient     *gcp.ExtensionRegistryClient
	desiredpb     *aiplatformpb.Extension
	desired       *krm.VertexAIExtension
	actual        *aiplatformpb.Extension
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &ExtensionAdapter{}

func (a *ExtensionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIExtension", "name", a.id.String())

	if a.id.Extension == "" {
		return false, nil
	}

	req := &aiplatformpb.GetExtensionRequest{
		Name: a.id.String(),
	}

	extensionpb, err := a.gcpClient.GetExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIExtension %q: %w", a.id.String(), err)
	}

	a.actual = extensionpb
	return true, nil
}

func (a *ExtensionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIExtension", "name", a.desired.GetName())

	req := &aiplatformpb.ImportExtensionRequest{
		Parent:    a.id.ParentString(),
		Extension: a.desiredpb,
	}

	op, err := a.gcpClient.ImportExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("importing VertexAIExtension %q: %w", a.desired.GetName(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of VertexAIExtension %q: %w", a.desired.GetName(), err)
	}

	log.V(2).Info("successfully created VertexAIExtension", "name", created.Name)

	// Rule: Always perform a GET operation (Get<Resource>) immediately after a Create or Update LRO
	// successfully completes to fetch the fully-populated resource before calling updateStatus.
	getReq := &aiplatformpb.GetExtensionRequest{
		Name: created.Name,
	}
	latest, err := a.gcpClient.GetExtension(ctx, getReq)
	if err != nil {
		latest = created // Fallback if GET fails
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *ExtensionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIExtension", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, updateMask, err := compareResource(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &aiplatformpb.UpdateExtensionRequest{
		Extension:  a.desiredpb,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIExtension %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated VertexAIExtension", "name", a.id.String())

	// Rule: Always perform a GET operation (Get<Resource>) immediately after a Create or Update LRO
	// successfully completes to fetch the fully-populated resource before calling updateStatus.
	getReq := &aiplatformpb.GetExtensionRequest{
		Name: updated.Name,
	}
	latest, err := a.gcpClient.GetExtension(ctx, getReq)
	if err != nil {
		latest = updated // Fallback if GET fails
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ExtensionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIExtension{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIExtensionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Extension

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.Extension)
	u.SetGroupVersionKind(krm.VertexAIExtensionGVK)
	return u, nil
}

func (a *ExtensionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *aiplatformpb.Extension) error {
	if latest.Name != "" {
		normalized, err := a.projectMapper.ReplaceProjectNumberWithIDInLink(ctx, latest.Name)
		if err != nil {
			return fmt.Errorf("normalizing latest extension name %q: %w", latest.Name, err)
		}
		latest.Name = normalized
	}
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIExtensionStatus{}
	status.ObservedState = VertexAIExtensionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ExtensionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIExtension", "name", a.id.String())

	req := &aiplatformpb.DeleteExtensionRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteExtension(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAIExtension %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of VertexAIExtension %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareResource(ctx context.Context, actual, desired *aiplatformpb.Extension) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAIExtensionSpec_FromProto, VertexAIExtensionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *aiplatformpb.Extension) {
		// populate defaults if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
