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

package vertexaifeatureonlinestore

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIFeatureOnlineStoreGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.FeatureOnlineStoreAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewFeatureOnlineStoreAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building featureOnlineStoreAdmin client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAIFeatureOnlineStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(reader.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}

	// Get featureOnlineStoreAdmin client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.FeatureOnlineStoreIdentity)
	if !ok {
		return nil, fmt.Errorf("expected FeatureOnlineStoreIdentity, got %T", id)
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
	id        *krm.FeatureOnlineStoreIdentity
	gcpClient *gcp.FeatureOnlineStoreAdminClient
	desired   *krm.VertexAIFeatureOnlineStore
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIFeatureOnlineStore", "name", a.id.String())

	req := &pb.GetFeatureOnlineStoreRequest{
		Name: a.id.String(),
	}

	featureOnlineStorepb, err := a.gcpClient.GetFeatureOnlineStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	observedState := aiplatform.VertexAIFeatureOnlineStoreObservedState_FromProto(mapCtx, featureOnlineStorepb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIFeatureOnlineStore", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	featureOnlineStorepb := aiplatform.VertexAIFeatureOnlineStoreSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from spec to proto: %w", mapCtx.Err())
	}

	req := &pb.CreateFeatureOnlineStoreRequest{
		Parent:               a.id.Parent().String(),
		FeatureOnlineStoreId: a.id.FeatureOnlineStore,
		FeatureOnlineStore:   featureOnlineStorepb,
	}

	op, err := a.gcpClient.CreateFeatureOnlineStore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VertexAIFeatureOnlineStore", "name", a.id.String())

	status := &krm.VertexAIFeatureOnlineStoreStatus{}
	status.ObservedState = aiplatform.VertexAIFeatureOnlineStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIFeatureOnlineStore", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	featureOnlineStorepb := aiplatform.VertexAIFeatureOnlineStoreSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from spec to proto: %w", mapCtx.Err())
	}
	featureOnlineStorepb.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}

	if a.desired.Spec.Labels != nil {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	req := &pb.UpdateFeatureOnlineStoreRequest{
		FeatureOnlineStore: featureOnlineStorepb,
		UpdateMask:         updateMask,
	}

	op, err := a.gcpClient.UpdateFeatureOnlineStore(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated VertexAIFeatureOnlineStore", "name", a.id.String())

	status := &krm.VertexAIFeatureOnlineStoreStatus{}
	status.ObservedState = aiplatform.VertexAIFeatureOnlineStoreObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
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
	log.V(2).Info("deleting VertexAIFeatureOnlineStore", "name", a.id.String())

	req := &pb.DeleteFeatureOnlineStoreRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteFeatureOnlineStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for delete VertexAIFeatureOnlineStore %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted VertexAIFeatureOnlineStore", "name", a.id.String())

	return true, nil
}
