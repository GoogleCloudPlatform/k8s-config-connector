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

package examplestore

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
	vertexaipb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIExampleStoreGVK, NewExampleStoreModel)
}

func NewExampleStoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelExampleStore{config: *config}, nil
}

var _ directbase.Model = &modelExampleStore{}

type modelExampleStore struct {
	config config.ControllerConfig
}

func (m *modelExampleStore) client(ctx context.Context, location string) (*gcp.ExampleStoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	aiplatformurl := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(aiplatformurl))
	gcpClient, err := gcp.NewExampleStoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building ExampleStore client: %w", err)
	}
	return gcpClient, err
}

func (m *modelExampleStore) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIExampleStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.VertexAIExampleStoreIdentity)

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &ExampleStoreAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelExampleStore) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ExampleStoreAdapter struct {
	id        *krm.VertexAIExampleStoreIdentity
	gcpClient *gcp.ExampleStoreClient
	desired   *krm.VertexAIExampleStore
	actual    *vertexaipb.ExampleStore
}

var _ directbase.Adapter = &ExampleStoreAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ExampleStoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIExampleStore", "name", a.id)

	req := &vertexaipb.GetExampleStoreRequest{Name: a.id.String()}
	examplestorepb, err := a.gcpClient.GetExampleStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIExampleStore %q: %w", a.id, err)
	}

	a.actual = examplestorepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ExampleStoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIExampleStore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAIExampleStoreSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &vertexaipb.CreateExampleStoreRequest{
		Parent:       parent,
		ExampleStore: resource,
	}
	op, err := a.gcpClient.CreateExampleStore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIExampleStore %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting VertexAIExampleStore %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created VertexAIExampleStore", "name", a.id)

	status := &krm.VertexAIExampleStoreStatus{}
	status.ObservedState = VertexAIExampleStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ExampleStoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIExampleStore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := VertexAIExampleStoreSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.VertexAIExampleStoreStatus{}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &vertexaipb.UpdateExampleStoreRequest{
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		ExampleStore: desiredPb,
	}
	op, err := a.gcpClient.UpdateExampleStore(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIExampleStore %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting VertexAIExampleStore %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated VertexAIExampleStore", "name", a.id.String())

	status := &krm.VertexAIExampleStoreStatus{}
	status.ObservedState = VertexAIExampleStoreObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ExampleStoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIExampleStore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIExampleStoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VertexAIExampleStoreGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ExampleStoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIExampleStore", "name", a.id)

	req := &vertexaipb.DeleteExampleStoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteExampleStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VertexAIExampleStore, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAIExampleStore %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted VertexAIExampleStore", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete VertexAIExampleStore %s: %w", a.id, err)
	}
	return true, nil
}
