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

package workflowexecutions

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflowexecutions/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/workflows/executions/apiv1"
	executionpb "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.WorkflowsExecutionGVK, NewExecutionModel)
}

func NewExecutionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelExecution{config: *config}, nil
}

var _ directbase.Model = &modelExecution{}

type modelExecution struct {
	config config.ControllerConfig
}

func (m *modelExecution) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Execution client: %w", err)
	}
	return gcpClient, err
}

func (m *modelExecution) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.WorkflowsExecution{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	id, err := krm.NewExecutionRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get workflows GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ExecutionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelExecution) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ExecutionAdapter struct {
	id        *krm.ExecutionRef
	gcpClient *gcp.Client
	desired   *krm.WorkflowsExecution
	actual    *executionpb.Execution
}

var _ directbase.Adapter = &ExecutionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ExecutionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Execution", "name", a.id)

	_, idIsSet, err := a.id.ExecutionID()
	if err != nil {
		return false, err
	}
	if !idIsSet { // resource is not yet created
		return false, nil
	}

	req := &executionpb.GetExecutionRequest{Name: a.id.External}
	executionpb, err := a.gcpClient.GetExecution(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Execution %q: %w", a.id, err)
	}

	a.actual = executionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ExecutionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Execution", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkflowsExecutionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"
	resource.Name = a.id.External
	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	req := &executionpb.CreateExecutionRequest{
		Parent:    parent,
		Execution: resource,
	}
	created, err := a.gcpClient.CreateExecution(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Execution %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Execution", "name", a.id)

	status := &krm.WorkflowsExecutionStatus{}
	status.ObservedState = WorkflowsExecutionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	tokens := strings.Split(created.Name, "/")
	externalRef := parent + "/executions/" + tokens[7]
	status.ExternalRef = direct.LazyPtr(externalRef)
	return createOp.UpdateStatus(ctx, status, nil)
}

// WorkflowsExecution does not support Update.
func (a *ExecutionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Execution", "name", a.id)
	status := &krm.WorkflowsExecutionStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ExecutionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.WorkflowsExecution{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkflowsExecutionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workflows" {
			obj.Spec.ProjectRef = &refs.ProjectRef{Name: tokens[1]}
			obj.Spec.Location = tokens[3]
		}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.External)
	u.SetGroupVersionKind(krm.WorkflowsExecutionGVK)

	u.Object = uObj
	return u, nil
}

// Cancel the execution resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ExecutionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("cancelling Execution", "name", a.id)

	// Check the state of the execution.
	// If the execution is in SUCCEED state, abort cancel.
	getReq := &executionpb.GetExecutionRequest{Name: a.id.External}
	obj, err := a.gcpClient.GetExecution(ctx, getReq)
	if obj.State == executionpb.Execution_SUCCEEDED {
		// Return success if not found (assume it was already deleted).
		log.V(2).Info("skipping cancel for SUCCEEDED Execution", "name", a.id)
		return false, nil
	}

	req := &executionpb.CancelExecutionRequest{Name: a.id.External}
	_, err = a.gcpClient.CancelExecution(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Execution, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Execution %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Execution", "name", a.id)
	return true, nil
}
