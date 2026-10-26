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

// +generated:controller
// krm.group: notebooks.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NotebooksRuntime

package notebooks

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/notebooks/apiv1"
	notebookspb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NotebooksRuntimeGVK, NewRuntimeModel)
}

func NewRuntimeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelRuntime{config: *config}, nil
}

var _ directbase.Model = &modelRuntime{}

type modelRuntime struct {
	config config.ControllerConfig
}

func (m *modelRuntime) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NotebooksRuntime{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NotebooksRuntimeIdentity)

	// Get notebooks GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &RuntimeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelRuntime) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//notebooks.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NotebooksRuntimeIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &RuntimeAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelRuntime) client(ctx context.Context) (*gcp.ManagedNotebookClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewManagedNotebookClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Runtime client: %w", err)
	}
	return gcpClient, err
}

type RuntimeAdapter struct {
	id        *krm.NotebooksRuntimeIdentity
	gcpClient *gcp.ManagedNotebookClient
	desired   *krm.NotebooksRuntime
	actual    *notebookspb.Runtime
}

var _ directbase.Adapter = &RuntimeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *RuntimeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Runtime", "name", a.id)

	req := &notebookspb.GetRuntimeRequest{
		Name: a.id.String(),
	}
	runtimepb, err := a.gcpClient.GetRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Runtime %q: %w", a.id, err)
	}

	a.actual = runtimepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RuntimeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Runtime", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NotebooksRuntimeSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &notebookspb.CreateRuntimeRequest{
		Parent:    a.id.ParentString(),
		RuntimeId: a.id.Runtime,
		Runtime:   resource,
	}
	op, err := a.gcpClient.CreateRuntime(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Runtime %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Runtime %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Runtime", "name", a.id)

	status := &krm.NotebooksRuntimeStatus{}
	status.ObservedState = NotebooksRuntimeObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RuntimeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Runtime", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NotebooksRuntimeSpec_v1alpha1_FromProto,
		NotebooksRuntimeSpec_v1alpha1_ToProto,
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
		req := &notebookspb.UpdateRuntimeRequest{
			UpdateMask: updateMask,
			Runtime:    resource,
		}
		op, err := a.gcpClient.UpdateRuntime(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Runtime %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Runtime %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Runtime", "name", a.id)
	}

	status := &krm.NotebooksRuntimeStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NotebooksRuntimeObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *RuntimeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NotebooksRuntime{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NotebooksRuntimeSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NotebooksRuntimeGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *RuntimeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Runtime", "name", a.id)

	req := &notebookspb.DeleteRuntimeRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Runtime, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Runtime %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Runtime", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Runtime %s: %w", a.id, err)
	}
	return true, nil
}
