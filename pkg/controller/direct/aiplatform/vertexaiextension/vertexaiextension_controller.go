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

package vertexaiextension

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	registry.RegisterModel(krm.VertexAIExtensionGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*gcp.ExtensionRegistryClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}
	gcpClient, err := gcp.NewExtensionRegistryClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ExtensionRegistryClient: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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

	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.VertexAIExtensionIdentity)
	if !ok {
		return nil, fmt.Errorf("expected VertexAIExtensionIdentity, got %T", id)
	}

	return &Adapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAIExtensionIdentity
	gcpClient *gcp.ExtensionRegistryClient
	desired   *krm.VertexAIExtension
	actual    *pb.Extension
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIExtension", "name", a.id.String())

	req := &pb.GetExtensionRequest{
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

	mapCtx := &direct.MapContext{}
	observedState := aiplatform.VertexAIExtensionObservedState_FromProto(mapCtx, extensionpb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIExtension", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	extensionpb := aiplatform.VertexAIExtensionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from spec to proto: %w", mapCtx.Err())
	}
	extensionpb.Name = a.id.String()

	req := &pb.ImportExtensionRequest{
		Parent:    a.id.ParentString(),
		Extension: extensionpb,
	}

	op, err := a.gcpClient.ImportExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIExtension %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create VertexAIExtension %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VertexAIExtension", "name", a.id.String())

	// Perform a GET operation to fetch the fully-populated resource
	latest, err := a.gcpClient.GetExtension(ctx, &pb.GetExtensionRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting VertexAIExtension after create %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIExtension", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	extensionpb := aiplatform.VertexAIExtensionSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from spec to proto: %w", mapCtx.Err())
	}
	extensionpb.Name = a.id.String()

	diffs, updateMask, err := compareExtension(ctx, a.actual, extensionpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateExtensionRequest{
		Extension:  extensionpb,
		UpdateMask: updateMask,
	}

	_, err = a.gcpClient.UpdateExtension(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIExtension %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated VertexAIExtension", "name", a.id.String())

	// Perform a GET operation to fetch the fully-populated resource
	latest, err := a.gcpClient.GetExtension(ctx, &pb.GetExtensionRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting VertexAIExtension after update %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("adapter not initialized")
	}
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.desired)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: obj}, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIExtension", "name", a.id.String())

	req := &pb.DeleteExtensionRequest{
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
		return false, fmt.Errorf("waiting for delete VertexAIExtension %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted VertexAIExtension", "name", a.id.String())

	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Extension) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIExtensionStatus{}
	status.ObservedState = aiplatform.VertexAIExtensionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareExtension(ctx context.Context, actual, desired *pb.Extension) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, aiplatform.VertexAIExtensionSpec_FromProto, aiplatform.VertexAIExtensionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Extension) {
		// Populate defaults if any
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	// Refine UpdateMask to only allowed paths
	var allowedPaths []string
	for _, path := range updateMask.Paths {
		switch path {
		case "display_name", "description", "runtime_config", "tool_use_examples":
			allowedPaths = append(allowedPaths, path)
		case "manifest":
			if actual.GetManifest().GetName() != desired.GetManifest().GetName() {
				return nil, nil, fmt.Errorf("manifest.name is immutable and cannot be updated")
			}
			if !proto.Equal(actual.GetManifest().GetApiSpec(), desired.GetManifest().GetApiSpec()) {
				return nil, nil, fmt.Errorf("manifest.api_spec is immutable and cannot be updated")
			}
			if !proto.Equal(actual.GetManifest().GetAuthConfig(), desired.GetManifest().GetAuthConfig()) {
				return nil, nil, fmt.Errorf("manifest.auth_config is immutable and cannot be updated")
			}
			if actual.GetManifest().GetDescription() != desired.GetManifest().GetDescription() {
				allowedPaths = append(allowedPaths, "manifest.description")
			}
		default:
			return nil, nil, fmt.Errorf("field %q is immutable and cannot be updated", path)
		}
	}
	updateMask.Paths = allowedPaths

	return diffs, updateMask, nil
}
