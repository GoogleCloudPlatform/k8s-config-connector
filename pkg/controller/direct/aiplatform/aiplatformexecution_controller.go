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

// +generated:controller
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformExecution

package aiplatform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformExecutionGVK, NewExecutionModel)
}

func NewExecutionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelExecution{config: *config}, nil
}

var _ directbase.Model = &modelExecution{}

type modelExecution struct {
	config config.ControllerConfig
}

func (m *modelExecution) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformExecution{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformExecutionIdentity)

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
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
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformExecutionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &ExecutionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelExecution) client(ctx context.Context, location string) (*gcp.MetadataClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewMetadataClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Execution client: %w", err)
	}
	return gcpClient, err
}

type ExecutionAdapter struct {
	id        *krm.AIPlatformExecutionIdentity
	gcpClient *gcp.MetadataClient
	desired   *krm.AIPlatformExecution
	actual    *aiplatformpb.Execution
}

var _ directbase.Adapter = &ExecutionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ExecutionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Execution", "name", a.id)

	req := &aiplatformpb.GetExecutionRequest{
		Name: a.id.String(),
	}
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

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ExecutionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Execution", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformExecutionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateExecutionRequest{
		Parent:      a.id.ParentString(),
		ExecutionId: a.id.Execution,
		Execution:   resource,
	}
	created, err := a.gcpClient.CreateExecution(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Execution %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Execution", "name", a.id)

	status := &krm.AIPlatformExecutionStatus{}
	status.ObservedState = AIPlatformExecutionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ExecutionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Execution", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformExecutionSpec_FromProto,
		AIPlatformExecutionSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &aiplatformpb.UpdateExecutionRequest{
			UpdateMask: updateMask,
			Execution:  resource,
		}
		updated, err = a.gcpClient.UpdateExecution(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Execution %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Execution", "name", a.id)
	}

	status := &krm.AIPlatformExecutionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformExecutionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ExecutionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformExecution{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformExecutionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformExecutionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ExecutionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Execution", "name", a.id)

	req := &aiplatformpb.DeleteExecutionRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteExecution(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Execution, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Execution %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Execution", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Execution %s: %w", a.id, err)
	}
	return true, nil
}
