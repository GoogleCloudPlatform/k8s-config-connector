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
// proto.message: google.cloud.datacatalog.v1.Tag
// crd.type: DataCatalogTag
// crd.version: v1alpha1

package datacatalog

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	api "cloud.google.com/go/datacatalog/apiv1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	// refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1" // Not directly used here but identity might need it
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DataCatalogTagGVK, NewTagModel)
}

func NewTagModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &tagModel{config: *config}, nil
}

var _ directbase.Model = &tagModel{}

type tagModel struct {
	config config.ControllerConfig
}

func (m *tagModel) client(ctx context.Context, projectID string) (*api.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID // Use project from identity
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building datacatalog tag client: %w", err)
	}

	return gcpClient, nil
}

func (m *tagModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataCatalogTag{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve the tag template reference first to get the full template name.
	tagTemplateName, err := obj.Spec.TemplateRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("resolving template reference: %w", err)
	}

	// DataCatalogTagIdentity requires the parent KRM object (Entry or EntryGroup)
	// and the TagTemplate name to construct the potential tag name.
	// However, the actual Tag ID is server-generated. We'll use the identity mostly
	// for the parent resource name.
	id, err := krm.NewTagIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	projectID := id.Parent().ProjectID

	gcpClient, err := m.client(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &tagAdapter{
		gcpClient:       gcpClient,
		id:              id, // Primarily used for parent resolution
		desired:         obj,
		reader:          reader,
		tagTemplateName: tagTemplateName, // Store resolved template name
	}, nil
}

func (m *tagModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type tagAdapter struct {
	gcpClient       *api.Client
	id              *krm.TagIdentity // Primarily for parent resolution
	desired         *krm.DataCatalogTag
	actual          *pb.Tag
	reader          client.Reader
	tagTemplateName string // Resolved TagTemplate resource name
}

var _ directbase.Adapter = &tagAdapter{}

// Find implements directbase.Adapter.
// Data Catalog Tags don't have a Get RPC. We list tags on the parent and filter.
func (a *tagAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	mapCtx := &direct.MapContext{}

	parentName := a.id.Parent().String()
	log.V(2).Info("listing datacatalog tags", "parent", parentName)

	// Get desired scope for matching
	desiredProto := DataCatalogTagSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	desiredColumn := desiredProto.GetColumn() // Works even if nil

	// Use the fully resolved template name for matching
	desiredTemplate := a.tagTemplateName

	req := &pb.ListTagsRequest{Parent: parentName}
	it := a.gcpClient.ListTags(ctx, req)
	for {
		tag, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			// PermissionDenied might occur if parent doesn't exist
			if direct.IsPermissionDenied(err) {
				log.V(1).Info("permission denied listing tags, parent likely missing", "parent", parentName, "error", err)
				return false, nil // Treat as not found
			}
			return false, fmt.Errorf("listing datacatalog tags for parent %q: %w", parentName, err)
		}

		// Match based on template and scope (column)
		if tag.Template == desiredTemplate {
			actualColumn := tag.GetColumn() // Works even if nil
			if actualColumn == desiredColumn {
				// Found a match based on template and scope. Assume this is our tag.
				// In rare cases of multiple identical tags (same template, same scope),
				// this might pick the wrong one, but the API doesn't provide better filtering.
				log.V(2).Info("found matching datacatalog tag", "name", tag.Name)
				a.actual = tag
				return true, nil
			}
		}
	}

	log.V(2).Info("no matching datacatalog tag found", "parent", parentName, "template", desiredTemplate, "column", desiredColumn)
	return false, nil
}

// Create implements directbase.Adapter.
func (a *tagAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	mapCtx := &direct.MapContext{}

	parentName := a.id.Parent().String()
	log.V(2).Info("creating datacatalog tag", "parent", parentName)

	desiredProto := DataCatalogTagSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set the fully resolved template name
	desiredProto.Template = a.tagTemplateName

	req := &pb.CreateTagRequest{
		Parent: parentName,
		Tag:    desiredProto,
	}
	created, err := a.gcpClient.CreateTag(ctx, req)
	if err != nil {
		return fmt.Errorf("creating datacatalog tag for parent %q: %w", parentName, err)
	}
	log.V(2).Info("successfully created datacatalog tag in gcp", "name", created.Name)

	status := &krm.DataCatalogTagStatus{}
	status.ObservedState = DataCatalogTagObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Store the generated name for future Find/Update/Delete operations.
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update implements directbase.Adapter.
// Only the 'fields' map is mutable. Template and scope (column) are immutable.
func (a *tagAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	mapCtx := &direct.MapContext{}

	// The name (including the system-generated ID) is required for update.
	if a.actual == nil || a.actual.Name == "" {
		// This should not happen if Find ran successfully.
		// Could also fetch from status.ExternalRef if needed.
		return fmt.Errorf("cannot update tag: actual state or name not found")
	}
	tagName := a.actual.Name
	log.V(2).Info("updating datacatalog tag", "name", tagName)

	desiredProto := DataCatalogTagSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Set required fields for the update request payload
	desiredProto.Name = tagName
	desiredProto.Template = a.actual.Template // Template is immutable, use actual value

	// Check if only mutable fields ('fields') have changed.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(desiredProto.Fields, a.actual.Fields) {
		updateMask.Paths = append(updateMask.Paths, "fields")
	}

	// Check for immutable field changes (should be rejected by KCC webhook ideally)
	if desiredProto.GetColumn() != a.actual.GetColumn() {
		return fmt.Errorf("attempting to update immutable field 'column' for tag %q", tagName)
	}
	// No need to check 'template' as we force it to the actual value above.

	var updated *pb.Tag
	var err error
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no mutable field needs update", "name", tagName)
		updated = a.actual // Use current actual state for status update
	} else {
		log.V(2).Info("updating fields", "name", tagName, "fields", updateMask.Paths)
		req := &pb.UpdateTagRequest{
			Tag:        desiredProto,
			UpdateMask: updateMask,
		}
		updated, err = a.gcpClient.UpdateTag(ctx, req)
		if err != nil {
			return fmt.Errorf("updating datacatalog tag %q: %w", tagName, err)
		}
		log.V(2).Info("successfully updated datacatalog tag in gcp", "name", tagName)
	}

	status := &krm.DataCatalogTagStatus{}
	status.ObservedState = DataCatalogTagObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Ensure ExternalRef is preserved/updated if somehow missing
	if status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(tagName)
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export implements directbase.Adapter.
func (a *tagAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("actual state is nil, Find() might not have been called or failed")
	}
	mapCtx := &direct.MapContext{}

	// Convert actual proto to KRM spec
	spec := DataCatalogTagSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Populate the reference from the identity
	// Assuming krm.EntryParent has an EntryGroupRef field holding the KRM reference.
	if parent := a.id.Parent(); parent != nil { // Check parent and nested ref
		spec.EntryRef.External = parent.String()
	}
	// Populate references
	spec.TemplateRef = &krm.TagTemplateRef{External: a.actual.Template}

	// Build unstructured object
	obj := &krm.DataCatalogTag{}
	obj.Spec = direct.ValueOf(spec)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	// Set metadata
	// Name for KRM is tricky as TagID is server-generated. Use a combination?
	// Or rely on labels/annotations? Let's use metadata.name from desired.
	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.desired.GetName())
	u.SetNamespace(a.desired.GetNamespace())
	u.SetGroupVersionKind(krm.DataCatalogTagGVK)
	// Add annotation for the full resource name?
	// common.SetAnnotation(u, common.AnnotationExternalResourceName, a.actual.Name)

	// Clear output-only fields from spec (already done by _FromProto generally)
	// unstructured.RemoveNestedField(u.Object, "spec", "templateDisplayName") // Should be handled by ObservedState mapping
	// unstructured.RemoveNestedField(u.Object, "spec", "dataplexTransferStatus") // Should be handled by ObservedState mapping

	// Clear status
	unstructured.RemoveNestedField(u.Object, "status")

	return u, nil
}

// Delete implements directbase.Adapter.
func (a *tagAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	// Get the tag name (including system-generated ID) from actual state or status
	tagName := ""
	if a.actual != nil && a.actual.Name != "" {
		tagName = a.actual.Name
	} else if a.desired.Status.ExternalRef != nil {
		tagName = *a.desired.Status.ExternalRef
	}

	if tagName == "" {
		// Resource was likely never created or Find failed. Assume deleted.
		log.V(2).Info("skipping delete for datacatalog tag: resource name not found", "namespace", a.desired.Namespace, "name", a.desired.Name)
		return true, nil
	}

	log.V(2).Info("deleting datacatalog tag", "name", tagName)

	req := &pb.DeleteTagRequest{Name: tagName}
	err := a.gcpClient.DeleteTag(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent datacatalog tag, assuming it was already deleted", "name", tagName)
			return true, nil
		}
		// PermissionDenied might also indicate not found or access removed
		if direct.IsPermissionDenied(err) {
			log.V(2).Info("skipping delete for datacatalog tag due to permission denied, potentially already deleted or access revoked", "name", tagName)
			return true, nil // Treat as deleted
		}
		return false, fmt.Errorf("deleting datacatalog tag %q: %w", tagName, err)
	}
	log.V(2).Info("successfully deleted datacatalog tag", "name", tagName)

	return true, nil
}
