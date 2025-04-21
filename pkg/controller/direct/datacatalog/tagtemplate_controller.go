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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.TagTemplate
// crd.type: DataCatalogTagTemplate
// crd.version: v1alpha1

package datacatalog

import (
	"context"
	"fmt"
	"reflect"

	api "cloud.google.com/go/datacatalog/apiv1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DataCatalogTagTemplateGVK, NewTagTemplateModel)
}

func NewTagTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &tagTemplateModel{config: *config}, nil
}

var _ directbase.Model = &tagTemplateModel{}

type tagTemplateModel struct {
	config config.ControllerConfig
}

func (m *tagTemplateModel) client(ctx context.Context, projectID string) (*api.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog tagtemplate client: %w", err)
	}

	return gcpClient, nil
}

func (m *tagTemplateModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataCatalogTagTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTagTemplateIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &tagTemplateAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *tagTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type tagTemplateAdapter struct {
	gcpClient *api.Client
	id        *krm.TagTemplateIdentity
	desired   *krm.DataCatalogTagTemplate
	actual    *pb.TagTemplate
}

var _ directbase.Adapter = &tagTemplateAdapter{}

func (a *tagTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := log.FromContext(ctx)
	log.V(2).Info("getting datacatalog tagtemplate", "name", a.id)

	req := &pb.GetTagTemplateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTagTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		if direct.IsPermissionDenied(err) {
			// Seeing DatacatalogTagTemplate returning 403 when the tagtemplate is not found.
			return false, nil
		}
		return false, fmt.Errorf("getting datacatalog tagtemplate %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *tagTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := log.FromContext(ctx)
	log.V(2).Info("creating datacatalog tagtemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogTagTemplateSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateTagTemplateRequest{
		Parent:        a.id.Parent().String(),
		TagTemplateId: a.id.ID(),
		TagTemplate:   desired,
	}
	created, err := a.gcpClient.CreateTagTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating datacatalog tagtemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created datacatalog tagtemplate in gcp", "name", a.id)

	status := &krm.DataCatalogTagTemplateStatus{}
	status.ObservedState = DataCatalogTagTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *tagTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := log.FromContext(ctx)
	log.V(2).Info("updating datacatalog tagtemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := DataCatalogTagTemplateSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	// The UpdateTagTemplate RPC can only update certain top-level fields.
	// Template fields (the 'fields' map) are managed via separate RPCs (Create/Update/DeleteTagTemplateField).
	// Therefore, we only compare and update the mutable top-level fields.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(desired.IsPubliclyReadable, a.actual.IsPubliclyReadable) {
		updateMask.Paths = append(updateMask.Paths, "is_publicly_readable")
	}
	// Note: 'fields' map is intentionally not included in the updateMask as it's managed separately.
	// Note: 'dataplex_transfer_status' might be updatable, but often such fields are immutable or have specific state transitions. Assuming immutable for now unless specified otherwise.

	var updated *pb.TagTemplate
	var err error
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Even though there is no update to the TagTemplate itself, we still want to update KRM status.
		// We also need to ensure that the fields are reconciled.
		// TODO: Implement reconciliation logic for TagTemplateFields if they are managed inline.
		// For now, just return the current actual state.
		updated = a.actual
	} else {
		req := &pb.UpdateTagTemplateRequest{
			TagTemplate: desired,
			UpdateMask:  updateMask,
		}
		updated, err = a.gcpClient.UpdateTagTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating datacatalog tagtemplate %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated datacatalog tagtemplate in gcp", "name", a.id)
		// TODO: After successful update, reconcile TagTemplateFields if managed inline.
	}

	// Regardless of whether an update call was made, reconcile status.
	// Need the latest state which might include output-only fields not present in `desired`.
	// Re-fetch if no update was performed or if the update response might be incomplete.
	if len(updateMask.Paths) == 0 {
		refetched, findErr := a.gcpClient.GetTagTemplate(ctx, &pb.GetTagTemplateRequest{Name: a.id.String()})
		if findErr != nil {
			// Log the error but continue with the potentially stale 'actual' for status update
			log.Error(findErr, "failed to re-fetch tagtemplate after no-op update", "name", a.id)
		} else {
			updated = refetched
			a.actual = refetched // Update internal 'actual' state
		}
	} else {
		// If an update was made, the 'updated' response should be sufficient.
		a.actual = updated // Update internal 'actual' state
	}

	status := &krm.DataCatalogTagTemplateStatus{}
	status.ObservedState = DataCatalogTagTemplateObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *tagTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataCatalogTagTemplate{}
	mapCtx := &direct.MapContext{}
	// Export the observed state, including fields.
	obj.Spec = direct.ValueOf(DataCatalogTagTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.SetName(a.id.ID())
	u.SetNamespace(a.desired.GetNamespace()) // Preserve namespace
	u.SetGroupVersionKind(krm.DataCatalogTagTemplateGVK)

	// Clear output-only fields from spec before setting object
	// 'dataplex_transfer_status' is marked Optional but behaves like output-only in practice for export.
	unstructured.RemoveNestedField(uObj, "spec", "dataplexTransferStatus")
	// Fields within TagTemplateField might have output-only subfields, ensure mapper handles this.

	u.Object = uObj

	// Clear status fields from the exported object
	unstructured.RemoveNestedField(u.Object, "status")

	return u, nil
}

func (a *tagTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := log.FromContext(ctx)
	log.V(2).Info("deleting datacatalog tagtemplate", "name", a.id)

	req := &pb.DeleteTagTemplateRequest{
		Name:  a.id.String(),
		Force: true, // Force deletion of template and associated tags.
	}
	err := a.gcpClient.DeleteTagTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent datacatalog tagtemplate, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		if direct.IsPermissionDenied(err) {
			// Treat as deleted if permission denied, might be gone already.
			log.V(2).Info("skipping delete for datacatalog tagtemplate due to permission denied, potentially already deleted or access revoked", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting datacatalog tagtemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted datacatalog tagtemplate", "name", a.id)

	return true, nil
}
