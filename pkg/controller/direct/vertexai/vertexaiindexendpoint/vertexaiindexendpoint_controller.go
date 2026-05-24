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

package vertexaiindexendpoint

import (
	"context"
	"fmt"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.VertexAIIndexEndpointGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*aiplatform.IndexEndpointClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := aiplatform.NewIndexEndpointClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building IndexEndpoint client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	reader := op.Reader
	u := op.GetUnstructured()
	obj := &krm.VertexAIIndexEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	identity := id.(*krm.VertexAIIndexEndpointIdentity)

	gcpClient, err := m.client(ctx, identity.Location)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        identity,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAIIndexEndpointIdentity
	gcpClient *aiplatform.IndexEndpointClient
	desired   *krm.VertexAIIndexEndpoint
	actual    *pb.IndexEndpoint
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIIndexEndpoint", "name", a.id.String())

	req := &pb.GetIndexEndpointRequest{
		Name: a.id.String(),
	}
	indexendpoint, err := a.gcpClient.GetIndexEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIIndexEndpoint %q: %w", a.id.String(), err)
	}

	a.actual = indexendpoint

	mapCtx := &direct.MapContext{}
	observed := api.VertexAIIndexEndpointObservedState_FromProto(mapCtx, indexendpoint)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}

	a.desired.Status.ObservedState = observed
	a.desired.Status.ExternalRef = direct.PtrTo(a.id.String())
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIIndexEndpoint", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := api.VertexAIIndexEndpointSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateIndexEndpointRequest{
		Parent:        "projects/" + a.id.Project + "/locations/" + a.id.Location,
		IndexEndpoint: resource,
	}
	op, err := a.gcpClient.CreateIndexEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIIndexEndpoint %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("VertexAIIndexEndpoint %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VertexAIIndexEndpoint", "name", a.id.String())

	status := api.VertexAIIndexEndpointObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.desired.Status.ObservedState = status
	a.desired.Status.ExternalRef = direct.PtrTo(a.id.String())
	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIIndexEndpoint", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := api.VertexAIIndexEndpointSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	req := &pb.UpdateIndexEndpointRequest{
		IndexEndpoint: resource,
		UpdateMask:    updateMask,
	}

	updated, err := a.gcpClient.UpdateIndexEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIIndexEndpoint %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated VertexAIIndexEndpoint", "name", a.id.String())

	status := api.VertexAIIndexEndpointObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.desired.Status.ObservedState = status
	return nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("VertexAIIndexEndpoint not found")
	}
	mapCtx := &direct.MapContext{}
	req := &pb.GetIndexEndpointRequest{
		Name: a.id.String(),
	}
	indexendpoint, err := a.gcpClient.GetIndexEndpoint(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting VertexAIIndexEndpoint %q: %w", a.id.String(), err)
	}

	spec := api.VertexAIIndexEndpointSpec_FromProto(mapCtx, indexendpoint)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Convert spec to unstructured
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": krm.VertexAIIndexEndpointGVK.GroupVersion().String(),
			"kind":       krm.VertexAIIndexEndpointGVK.Kind,
			"spec":       obj,
		},
	}
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIIndexEndpoint", "name", a.id.String())

	req := &pb.DeleteIndexEndpointRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteIndexEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting VertexAIIndexEndpoint %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("VertexAIIndexEndpoint %s waiting deletion: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted VertexAIIndexEndpoint", "name", a.id.String())
	return true, nil
}
