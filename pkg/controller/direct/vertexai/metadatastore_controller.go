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

package vertexai

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	vertexaipb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	cloudresourcemanager "cloud.google.com/go/resourcemanager/apiv3"

	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.VertexAIMetadataStoreGVK, NewMetadataStoreModel)
}

func NewMetadataStoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMetadataStore{config: *config}, nil
}

var _ directbase.Model = &modelMetadataStore{}

type modelMetadataStore struct {
	config config.ControllerConfig
}

func (m *modelMetadataStore) client(ctx context.Context, location string) (*gcp.MetadataClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	aiplatformurl := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(aiplatformurl))
	gcpClient, err := gcp.NewMetadataRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building MetadataStore client: %w", err)
	}
	return gcpClient, err
}

func (m *modelMetadataStore) projectsClient(ctx context.Context) (*cloudresourcemanager.ProjectsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	crmClient, err := cloudresourcemanager.NewProjectsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudresourcemanager client: %w", err)
	}
	return crmClient, err
}

func (m *modelMetadataStore) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VertexAIMetadataStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMetadataStoreIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err = normalizeExternal(ctx, reader, u, obj); err != nil {
		return nil, err
	}

	// Get Project GCP client
	projectClient, err := m.projectsClient(ctx)
	if err != nil {
		return nil, err
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Parent().Location)
	if err != nil {
		return nil, err
	}

	return &MetadataStoreAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       obj,
		projectClient: projectClient,
	}, nil
}

func (m *modelMetadataStore) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func normalizeExternal(ctx context.Context, reader client.Reader, src client.Object, store *krm.VertexAIMetadataStore) error {
	// Normalize the kmsKeyRef in the encryptionSpec.
	if store.Spec.EncryptionSpec != nil && store.Spec.EncryptionSpec.KMSKeyRef != nil {
		kmsKey, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, src.GetNamespace(), store.Spec.EncryptionSpec.KMSKeyRef)
		if err != nil {
			return err
		}
		store.Spec.EncryptionSpec.KMSKeyRef = kmsKey
	}
	return nil
}

type MetadataStoreAdapter struct {
	id            *krm.MetadataStoreIdentity
	gcpClient     *gcp.MetadataClient
	desired       *krm.VertexAIMetadataStore
	actual        *vertexaipb.MetadataStore
	projectClient *cloudresourcemanager.ProjectsClient
}

var _ directbase.Adapter = &MetadataStoreAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MetadataStoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MetadataStore", "name", a.id)

	// metadataStores requires project number instead of project ID
	projectNumber, err := getProjectNumberFromID(ctx, a.id.Parent().ProjectID, a.projectClient)
	if err != nil {
		return false, fmt.Errorf("error converting project ID %s to project number: %w", a.id.Parent().ProjectID, err)
	}
	id := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", projectNumber, a.id.Parent().Location, a.id.ID())
	req := &vertexaipb.GetMetadataStoreRequest{Name: id}
	metadatastorepb, err := a.gcpClient.GetMetadataStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MetadataStore %q: %w", a.id, err)
	}

	a.actual = metadatastorepb
	return true, nil
}

func getProjectNumberFromID(ctx context.Context, projectID string, projectsClient *cloudresourcemanager.ProjectsClient) (string, error) {
	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + projectID,
	}
	project, err := projectsClient.GetProject(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error getting project %q: %w", req.Name, err)
	}
	n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
	if err != nil {
		return "", fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
	}
	return fmt.Sprintf("%d", n), nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetadataStoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MetadataStore", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAIMetadataStoreSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &vertexaipb.CreateMetadataStoreRequest{
		Parent:          a.id.Parent().String(),
		MetadataStore:   resource,
		MetadataStoreId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateMetadataStore(ctx, req)
	if err != nil {
		return fmt.Errorf("MetadataStore %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetadataStore %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MetadataStore", "name", a.id)

	obj, err := a.gcpClient.GetMetadataStore(ctx, &vertexaipb.GetMetadataStoreRequest{Name: a.id.String()})
	if err != nil {
		return err
	}
	created.State = obj.State
	status := &krm.VertexAIMetadataStoreStatus{}
	status.ObservedState = VertexAIMetadataStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetadataStoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// This resource does not support update functions.
	return nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MetadataStoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIMetadataStore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIMetadataStoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Region = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VertexAIMetadataStoreGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MetadataStoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MetadataStore", "name", a.id)

	req := &vertexaipb.DeleteMetadataStoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMetadataStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MetadataStore, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting MetadataStore %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MetadataStore", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MetadataStore %s: %w", a.id, err)
	}
	return true, nil
}
