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

package notebooks

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/notebooks/apiv1"
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NotebookRuntimeGVK, NewNotebookRuntimeModel)
}

func NewNotebookRuntimeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNotebookRuntime{config: *config}, nil
}

var _ directbase.Model = &modelNotebookRuntime{}

type modelNotebookRuntime struct {
	config config.ControllerConfig
}

func (m *modelNotebookRuntime) client(ctx context.Context) (*gcp.ManagedNotebookClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewManagedNotebookClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NotebookRuntime client: %w", err)
	}
	return gcpClient, err
}

func (m *modelNotebookRuntime) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NotebookRuntime{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*krm.NotebookRuntimeIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &NotebookRuntimeAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelNotebookRuntime) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NotebookRuntimeAdapter struct {
	id        *krm.NotebookRuntimeIdentity
	gcpClient *gcp.ManagedNotebookClient
	desired   *krm.NotebookRuntime
	actual    *pb.Runtime
}

var _ directbase.Adapter = &NotebookRuntimeAdapter{}

func (a *NotebookRuntimeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NotebookRuntime", "name", a.id)

	req := &pb.GetRuntimeRequest{Name: a.id.String()}
	runtimepb, err := a.gcpClient.GetRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NotebookRuntime %q: %w", a.id, err)
	}

	a.actual = runtimepb
	return true, nil
}

func (a *NotebookRuntimeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NotebookRuntime", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := Runtime_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateRuntimeRequest{
		Parent:    parent,
		RuntimeId: a.id.Runtime,
		Runtime:   resource,
	}
	op, err := a.gcpClient.CreateRuntime(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NotebookRuntime %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for NotebookRuntime %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NotebookRuntime", "name", a.id)

	status := &krm.NotebookRuntimeStatus{}
	status.ObservedState = RuntimeObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *NotebookRuntimeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NotebookRuntime status", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.NotebookRuntimeStatus{}
	status.ObservedState = RuntimeObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *NotebookRuntimeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}
	mapCtx := &direct.MapContext{}
	spec := Runtime_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	spec.Location = direct.PtrTo(a.id.Location)
	spec.ResourceID = direct.PtrTo(a.id.Runtime)

	obj := &krm.NotebookRuntime{
		Spec: *spec,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *NotebookRuntimeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NotebookRuntime", "name", a.id)

	req := &pb.DeleteRuntimeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting NotebookRuntime %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for NotebookRuntime %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NotebookRuntime", "name", a.id)
	return true, nil
}
