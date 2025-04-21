// Copyright 2025 Google LLC
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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Task
// crd.type: DataplexTask
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexTaskGVK, NewTaskModel)
}

func NewTaskModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &taskModel{config: config}, nil
}

var _ directbase.Model = &taskModel{}

type taskModel struct {
	config *config.ControllerConfig
}

func (m *taskModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexTask{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTaskIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	taskAdapter := &taskAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	taskClient, err := gcpClient.client(ctx)
	if err != nil {
		return nil, err
	}
	taskAdapter.gcpClient = taskClient

	return taskAdapter, nil
}

func (m *taskModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type taskAdapter struct {
	gcpClient *gcp.Client
	id        *krm.TaskIdentity
	desired   *krm.DataplexTask
	actual    *pb.Task
	reader    client.Reader
}

var _ directbase.Adapter = &taskAdapter{}

func (a *taskAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex task", "name", a.id)

	req := &pb.GetTaskRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTask(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex task %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *taskAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex task", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	desired.Name = a.id.String()

	task := DataplexTaskSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateTaskRequest{
		Parent: a.id.Parent(),
		Task:   task,
		TaskId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateTask(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex task %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex task %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex task in gcp", "name", a.id)

	status := &krm.DataplexTaskStatus{}
	status.ObservedState = DataplexTaskObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *taskAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex task", "name", a.id)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	task := DataplexTaskSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	task.Name = a.id.String()

	// Set the required trigger type explicitly if unset, default to ON_DEMAND.
	// This ensures comparison works correctly, even if the user didn't specify it.
	if task.TriggerSpec == nil {
		task.TriggerSpec = &pb.Task_TriggerSpec{}
	}
	if task.TriggerSpec.Type == pb.Task_TriggerSpec_TYPE_UNSPECIFIED {
		task.TriggerSpec.Type = pb.Task_TriggerSpec_ON_DEMAND
	}
	// Do the same for the actual resource state for comparison.
	if a.actual.TriggerSpec == nil {
		a.actual.TriggerSpec = &pb.Task_TriggerSpec{}
	}
	if a.actual.TriggerSpec.Type == pb.Task_TriggerSpec_TYPE_UNSPECIFIED {
		a.actual.TriggerSpec.Type = pb.Task_TriggerSpec_ON_DEMAND
	}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(task.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(task.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(task.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	// TriggerSpec.Type is immutable, compare the rest of the spec
	if !reflect.DeepEqual(task.TriggerSpec.StartTime, a.actual.TriggerSpec.StartTime) {
		updateMask.Paths = append(updateMask.Paths, "trigger_spec.start_time")
	}
	if !reflect.DeepEqual(task.TriggerSpec.Disabled, a.actual.TriggerSpec.Disabled) {
		updateMask.Paths = append(updateMask.Paths, "trigger_spec.disabled")
	}
	if !reflect.DeepEqual(task.TriggerSpec.MaxRetries, a.actual.TriggerSpec.MaxRetries) {
		updateMask.Paths = append(updateMask.Paths, "trigger_spec.max_retries")
	}
	if !reflect.DeepEqual(task.TriggerSpec.GetSchedule(), a.actual.TriggerSpec.GetSchedule()) {
		updateMask.Paths = append(updateMask.Paths, "trigger_spec.schedule")
	}

	if !cmp.Equal(task.ExecutionSpec, a.actual.ExecutionSpec, cmpopts.IgnoreUnexported(pb.Task_ExecutionSpec{})) {
		updateMask.Paths = append(updateMask.Paths, "execution_spec")
	}

	// Compare task-specific config (spark or notebook)
	if !cmp.Equal(task.GetSpark(), a.actual.GetSpark(), cmpopts.IgnoreUnexported(pb.Task_SparkTaskConfig{}, pb.Task_InfrastructureSpec{})) {
		updateMask.Paths = append(updateMask.Paths, "spark")
	}
	if !cmp.Equal(task.GetNotebook(), a.actual.GetNotebook(), cmpopts.IgnoreUnexported(pb.Task_NotebookTaskConfig{}, pb.Task_InfrastructureSpec{})) {
		updateMask.Paths = append(updateMask.Paths, "notebook")
	}

	var updated *pb.Task
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)

		// Even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		log.V(2).Info("updating dataplex task fields", "name", a.id, "paths", updateMask.Paths)
		req := &pb.UpdateTaskRequest{
			UpdateMask: updateMask,
			Task:       task,
		}
		op, err := a.gcpClient.UpdateTask(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex task %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex task %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex task", "name", a.id)
	}

	status := &krm.DataplexTaskStatus{}
	status.ObservedState = DataplexTaskObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *taskAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexTask{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexTaskSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set parent references
	obj.Spec.LakeRef = &krm.LakeRef{External: a.id.Parent()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())                // Use the task_id as the KRM resource name
	u.SetNamespace(a.desired.Namespace) // Preserve namespace
	u.SetGroupVersionKind(krm.DataplexTaskGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *taskAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex task", "name", a.id)

	req := &pb.DeleteTaskRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTask(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataplex task %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated dataplex task deletion", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex task %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted dataplex task", "name", a.id)
	return true, nil
}
