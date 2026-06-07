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

package vertexai

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIReasoningEngineGVK, NewReasoningEngineModel)
}

func NewReasoningEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReasoningEngine{config: *config}, nil
}

var _ directbase.Model = &modelReasoningEngine{}

type modelReasoningEngine struct {
	config config.ControllerConfig
}

func (m *modelReasoningEngine) client(ctx context.Context, location string) (*gcp.ReasoningEngineClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewReasoningEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReasoningEngine client: %w", err)
	}
	return gcpClient, err
}

func (m *modelReasoningEngine) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIReasoningEngine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewReasoningEngineIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Parent().Location)
	if err != nil {
		return nil, err
	}
	return &ReasoningEngineAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReasoningEngine) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ReasoningEngineAdapter struct {
	id        *krm.ReasoningEngineIdentity
	gcpClient *gcp.ReasoningEngineClient
	desired   *krm.VertexAIReasoningEngine
	actual    *pb.ReasoningEngine
}

var _ directbase.Adapter = &ReasoningEngineAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReasoningEngineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReasoningEngine", "name", a.id)

	req := &pb.GetReasoningEngineRequest{Name: a.id.String()}
	reasoningenginepb, err := a.gcpClient.GetReasoningEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReasoningEngine %q: %w", a.id, err)
	}

	a.actual = reasoningenginepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReasoningEngineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReasoningEngine", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAIReasoningEngineSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateReasoningEngineRequest{
		Parent:          a.id.Parent().String(),
		ReasoningEngine: resource,
	}
	op, err := a.gcpClient.CreateReasoningEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReasoningEngine %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting creation ReasoningEngine %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ReasoningEngine", "name", a.id)
	status := &krm.VertexAIReasoningEngineStatus{}
	status.ObservedState = VertexAIReasoningEngineObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	_, actualResourceID, err := krm.ParseReasoningEngineExternal(created.Name)
	if err != nil {
		return fmt.Errorf("parsing created ReasoningEngine name: %w", err)
	}
	externalRef := fmt.Sprintf("%s/reasoningEngines/%s", a.id.Parent().String(), actualResourceID)
	status.ExternalRef = direct.LazyPtr(externalRef)

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReasoningEngineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReasoningEngine", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := VertexAIReasoningEngineSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Name = a.id.String() // populate the name field before comparison

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	paths.Delete("name")
	paths.Delete("etag")

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	desiredPb.Name = a.id.String() // populate the name field
	req := &pb.UpdateReasoningEngineRequest{
		UpdateMask:      updateMask,
		ReasoningEngine: desiredPb,
	}
	op, err := a.gcpClient.UpdateReasoningEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ReasoningEngine %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting update ReasoningEngine %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ReasoningEngine", "name", a.id)

	status := &krm.VertexAIReasoningEngineStatus{}
	status.ObservedState = VertexAIReasoningEngineObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReasoningEngineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIReasoningEngine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIReasoningEngineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.VertexAIReasoningEngineGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReasoningEngineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReasoningEngine", "name", a.id)

	req := &pb.DeleteReasoningEngineRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteReasoningEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ReasoningEngine, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ReasoningEngine %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ReasoningEngine", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ReasoningEngine %s: %w", a.id, err)
	}
	return true, nil
}
