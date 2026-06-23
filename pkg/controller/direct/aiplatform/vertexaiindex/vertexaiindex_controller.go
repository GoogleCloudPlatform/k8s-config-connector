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

package vertexaiindex

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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIIndexGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.IndexClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewIndexClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Index client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.VertexAIIndex{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}

	// Get gcp client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id.(*krm.VertexAIIndexIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support Discovery
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAIIndexIdentity
	gcpClient *gcp.IndexClient
	desired   *krm.VertexAIIndex
	actual    *pb.Index
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIIndex", "name", a.id.String())

	req := &pb.GetIndexRequest{
		Name: a.id.String(),
	}
	index, err := a.gcpClient.GetIndex(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIIndex %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	observedState := aiplatform.VertexAIIndexObservedState_FromProto(mapCtx, index)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping to observed state: %w", mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.actual = index
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIIndex", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := aiplatform.VertexAIIndexSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateIndexRequest{
		Parent: parent,
		Index:  resource,
	}
	op, err := a.gcpClient.CreateIndex(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIIndex %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for VertexAIIndex %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VertexAIIndex", "name", a.id.String())

	status := &krm.VertexAIIndexStatus{}
	status.ObservedState = aiplatform.VertexAIIndexObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIIndex", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := aiplatform.VertexAIIndexSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	paths := []string{"display_name", "description", "labels"}
	updateMask := &fieldmaskpb.FieldMask{Paths: paths}

	req := &pb.UpdateIndexRequest{
		Index:      resource,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateIndex(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIIndex %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for VertexAIIndex %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated VertexAIIndex", "name", a.id.String())

	status := &krm.VertexAIIndexStatus{}
	status.ObservedState = aiplatform.VertexAIIndexObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	obj := &krm.VertexAIIndex{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(aiplatform.VertexAIIndexSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: uObj}
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIIndex", "name", a.id.String())

	req := &pb.DeleteIndexRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteIndex(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting VertexAIIndex %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for VertexAIIndex %s deletion: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted VertexAIIndex", "name", a.id.String())
	return true, nil
}
