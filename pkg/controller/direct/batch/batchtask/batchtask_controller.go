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
// proto.service: google.cloud.batch.v1.BatchService
// proto.message: google.cloud.batch.v1.Task
// crd.type: BatchTask
// crd.version: v1alpha1

package batchtask

import (
	"context"
	"fmt"

	gcpbatch "cloud.google.com/go/batch/apiv1"
	batchpb "cloud.google.com/go/batch/apiv1/batchpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/batch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.BatchTaskGVK, NewTaskModel)
}

func NewTaskModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &taskModel{config: *config}, nil
}

var _ directbase.Model = &taskModel{}

type taskModel struct {
	config config.ControllerConfig
}

func (m *taskModel) Client(ctx context.Context, projectID string) (*gcpbatch.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcpbatch.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building batch task client: %w", err)
	}

	return gcpClient, err
}

func (m *taskModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BatchTask{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.TaskIdentity)

	gcpClient, err := m.Client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &taskAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

func (m *taskModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type taskAdapter struct {
	gcpClient *gcpbatch.Client
	id        *krm.TaskIdentity
	actual    *batchpb.Task
}

var _ directbase.Adapter = &taskAdapter{}

func (a *taskAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting batch task", "name", a.id)

	req := &batchpb.GetTaskRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTask(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting batch task %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *taskAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("BatchTask %s does not exist on GCP; Tasks are read-only and cannot be created directly", a.id.String())
}

func (a *taskAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating batch task status", "name", a.id)

	status := &krm.BatchTaskStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = batch.BatchTaskObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.actual.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *taskAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting batch task (read-only, no GCP API call)", "name", a.id)
	return true, nil
}

func (a *taskAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BatchTask{}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Task)
	u.SetGroupVersionKind(krm.BatchTaskGVK)

	u.Object = uObj
	return u, nil
}
