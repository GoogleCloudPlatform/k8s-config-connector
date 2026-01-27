// Copyright 2024 Google LLC
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
// proto.service: google.cloud.aiplatform.v1.NotebookService
// proto.message: google.cloud.aiplatform.v1.NotebookRuntime
// crd.type: ColabRuntime
// crd.version: v1alpha1

package colab

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ColabRuntimeGVK, NewColabRuntimeModel)
}

func NewColabRuntimeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &runtimeModel{config: *config}, nil
}

var _ directbase.Model = &runtimeModel{}

type runtimeModel struct {
	config config.ControllerConfig
}

func (m *runtimeModel) client(ctx context.Context, location string) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption
	config := m.config
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewNotebookRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building colabruntimetemplate client: %w", err)
	}

	return gcpClient, err
}

func (m *runtimeModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ColabRuntime{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNotebookRuntimeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().Location)
	if err != nil {
		return nil, err
	}

	return &runtimeAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *runtimeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type runtimeAdapter struct {
	gcpClient *gcp.NotebookClient
	id        *krm.NotebookRuntimeIdentity
	desired   *krm.ColabRuntime
	actual    *pb.NotebookRuntime
	reader    client.Reader
}

var _ directbase.Adapter = &runtimeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *runtimeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ColabRuntime", "name", a.id)

	req := &pb.GetNotebookRuntimeRequest{Name: a.id.String()}
	notebookRuntimePb, err := a.gcpClient.GetNotebookRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ColabRuntime %q: %w", a.id, err)
	}

	a.actual = notebookRuntimePb
	return true, nil
}

func normalizeReferences(ctx context.Context, reader client.Reader, obj *krm.ColabRuntime) error {
	if obj.Spec.ColabRuntimeTemplateRef != nil {
		if _, err := obj.Spec.ColabRuntimeTemplateRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *runtimeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ColabRuntime", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := normalizeReferences(ctx, a.reader, a.desired); err != nil {
		return fmt.Errorf("resolving references: %w", err)
	}

	desiredPb := ColabRuntimeSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.AssignNotebookRuntimeRequest{
		Parent:                  a.id.Parent().String(),
		NotebookRuntimeTemplate: a.desired.Spec.ColabRuntimeTemplateRef.External,
		NotebookRuntime:         desiredPb,
		NotebookRuntimeId:       a.id.ID(),
	}
	op, err := a.gcpClient.AssignNotebookRuntime(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ColabRuntime %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ColabRuntime %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ColabRuntime in gcp", "name", a.id)

	status := &krm.ColabRuntimeStatus{}
	status.ObservedState = ColabRuntimeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *runtimeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ColabRuntime", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := normalizeReferences(ctx, a.reader, a.desired); err != nil {
		return fmt.Errorf("resolving references: %w", err)
	}

	desiredPb := ColabRuntimeSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Still update the status to cover the use case of acquisition.
		if a.desired.Status.ExternalRef == nil {
			observedState := ColabRuntimeObservedState_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}

			a.desired.Status.ExternalRef = direct.PtrTo(a.id.String())
			a.desired.Status.ObservedState = observedState

			return updateOp.UpdateStatus(ctx, a.desired.Status, nil)
		}
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	log.V(2).Info("ColabRuntime doesn't support update", "name", a.id)
	return fmt.Errorf("updating ColabRuntime %s: update is not supported", a.id)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *runtimeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ColabRuntime{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ColabRuntimeSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.ColabRuntimeGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *runtimeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ColabRuntime", "name", a.id)

	req := &pb.DeleteNotebookRuntimeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNotebookRuntime(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ColabRuntime, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ColabRuntime %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ColabRuntime", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ColabRuntime %s: %w", a.id, err)
	}
	return true, nil
}
