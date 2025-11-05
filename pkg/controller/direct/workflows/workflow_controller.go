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

package workflows

import (
	"context"
	"fmt"
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/workflows/apiv1"

	workflowspb "cloud.google.com/go/workflows/apiv1/workflowspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.WorkflowsWorkflowGVK, NewWorkflowsWorkflowModel)
}

func NewWorkflowsWorkflowModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkflowsWorkflow{config: *config}, nil
}

var _ directbase.Model = &modelWorkflowsWorkflow{}

type modelWorkflowsWorkflow struct {
	config config.ControllerConfig
}

func (m *modelWorkflowsWorkflow) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Workflow client: %w", err)
	}
	return gcpClient, err
}

func normalizeReference(ctx context.Context, reader client.Reader, obj *krm.WorkflowsWorkflow) error {
	if obj.Spec.ServiceAccountRef != nil {
		if err := obj.Spec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return err
		}
	}

	if obj.Spec.CryptoKeyNameRef != nil {
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.CryptoKeyNameRef)
		if err != nil {
			return err
		}
		obj.Spec.CryptoKeyNameRef = kmsKeyRef
	}
	return nil
}

func (m *modelWorkflowsWorkflow) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.WorkflowsWorkflow{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err = normalizeReference(ctx, reader, obj); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()

	desired := WorkflowsWorkflowSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	// Get workflows GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkflowsWorkflowAdapter{
		id:        id.(*krm.WorkflowsWorkflowIdentity),
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *modelWorkflowsWorkflow) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WorkflowsWorkflowAdapter struct {
	id        *krm.WorkflowsWorkflowIdentity
	gcpClient *gcp.Client
	desired   *workflowspb.Workflow
	actual    *workflowspb.Workflow
	reader    client.Reader
}

var _ directbase.Adapter = &WorkflowsWorkflowAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *WorkflowsWorkflowAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Workflow", "name", a.id)

	req := &workflowspb.GetWorkflowRequest{Name: a.id.String()}
	workflowpb, err := a.gcpClient.GetWorkflow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Workflow %q: %w", a.id, err)
	}

	a.actual = workflowpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkflowsWorkflowAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Workflow", "name", a.id)

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &workflowspb.CreateWorkflowRequest{
		Parent:     a.id.Parent().String(),
		Workflow:   a.desired,
		WorkflowId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateWorkflow(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Workflow %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Workflow %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Workflow", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.WorkflowsWorkflowStatus{}
	status.ObservedState = WorkflowsWorkflowObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkflowsWorkflowAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Workflow", "name", a.id)

	paths := []string{}
	if !reflect.DeepEqual(a.desired.Description, a.actual.Description) {
		paths = append(paths, "description")
	}

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}

	if !reflect.DeepEqual(a.desired.ServiceAccount, a.actual.ServiceAccount) {
		paths = append(paths, "service_account")
	}

	if !reflect.DeepEqual(a.desired.SourceCode, a.actual.SourceCode) {
		paths = append(paths, "source_contents")
	}

	if !reflect.DeepEqual(a.desired.CryptoKeyName, a.actual.CryptoKeyName) {
		paths = append(paths, "crypto_key_name")
	}

	if !reflect.DeepEqual(a.desired.CallLogLevel, a.actual.CallLogLevel) {
		paths = append(paths, "call_log_level")
	}

	if !reflect.DeepEqual(a.desired.UserEnvVars, a.actual.UserEnvVars) {
		paths = append(paths, "user_env_vars")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: paths,
	}

	req := &workflowspb.UpdateWorkflowRequest{
		UpdateMask: updateMask,
		Workflow:   a.desired,
	}
	op, err := a.gcpClient.UpdateWorkflow(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Workflow %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Workflow %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Workflow", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.WorkflowsWorkflowStatus{}
	status.ObservedState = WorkflowsWorkflowObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WorkflowsWorkflowAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.WorkflowsWorkflow{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkflowsWorkflowSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.WorkflowsWorkflowGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WorkflowsWorkflowAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Workflow", "name", a.id)

	req := &workflowspb.DeleteWorkflowRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkflow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Workflow, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Workflow %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Workflow", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Workflow %s: %w", a.id, err)
	}
	return true, nil
}
