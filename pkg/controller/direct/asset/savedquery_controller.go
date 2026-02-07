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
// proto.service: google.cloud.asset.v1.AssetService
// proto.message: google.cloud.asset.v1.SavedQuery
// crd.type: AssetSavedQuery
// crd.version: v1alpha1

package asset

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/asset/apiv1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AssetSavedQueryGVK, NewSavedQueryModel)
}

func NewSavedQueryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &savedQueryModel{config: *config}, nil
}

var _ directbase.Model = &savedQueryModel{}

type savedQueryModel struct {
	config config.ControllerConfig
}

func (m *savedQueryModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	var opts []option.ClientOption

	config := m.config
	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building asset saved query client: %w", err)
	}

	return gcpClient, err
}

func (m *savedQueryModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AssetSavedQuery{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSavedQueryIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &savedQueryAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *savedQueryModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type savedQueryAdapter struct {
	gcpClient *gcp.Client
	id        *krm.SavedQueryIdentity
	desired   *krm.AssetSavedQuery
	actual    *pb.SavedQuery
	reader    client.Reader
}

var _ directbase.Adapter = &savedQueryAdapter{}

func (a *savedQueryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting asset saved query", "name", a.id)

	name := a.id.String()
	if a.actual != nil {
		name = a.actual.Name
	}
	req := &pb.GetSavedQueryRequest{Name: name}
	actual, err := a.gcpClient.GetSavedQuery(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting asset saved query %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *savedQueryAdapter) normalizeReferences(ctx context.Context) error {
	// No references to resolve in SavedQuery spec currently.
	return nil
}

func (a *savedQueryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating asset saved query", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := AssetSavedQuerySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// resource.Name = a.id.String() // Name is output only for SavedQuery

	req := &pb.CreateSavedQueryRequest{
		Parent:       a.id.Parent().String(),
		SavedQueryId: a.id.ID(),
		SavedQuery:   resource,
	}

	// Attempt to create the saved query
	actual, err := a.gcpClient.CreateSavedQuery(ctx, req)

	// Regardless of CreateSavedQuery result (success or AlreadyExists),
	// we need to ensure a.actual is populated by calling Find.
	found := false
	var findErr error

	if err == nil {
		log.V(2).Info("successfully created asset saved query in gcp, attempting to fetch it", "name", a.id)
		a.actual = actual
		found = true
	} else if errors.IsAlreadyExists(err) {
		// Resource already exists. Log and attempt to Find to populate a.actual.
		log.V(2).Info("asset saved query already exists during creation attempt, attempting to fetch it", "name", a.id.String(), "warning", err)
		found, findErr = a.Find(ctx)
		if findErr != nil {
			return fmt.Errorf("fetching existing asset saved query %q after CreateSavedQuery returned AlreadyExists failed: %w", a.id.String(), findErr)
		}
		if !found {
			// This is unexpected if CreateSavedQuery returned AlreadyExists
			return fmt.Errorf("asset saved query %q not found even though CreateSavedQuery returned AlreadyExists", a.id.String())
		}
	} else {
		// Other error during creation
		return fmt.Errorf("creating asset saved query %s: %w", a.id.String(), err)
	}

	// Update status
	status := AssetSavedQueryStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *savedQueryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating asset saved query", "name", a.id, "actual", a.actual.Name)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := AssetSavedQuerySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.actual.Name // Name must be set for update based on actual state

	// Construct the protobuf message containing only the fields to update
	updateProto := &pb.SavedQuery{Name: a.actual.Name}

	// Create a proto diff and build the update mask
	diffPathsSet, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing proto messages: %w", err)
	}

	// Filter out output-only fields from paths
	mutablePathsSet := sets.New[string]()
	for _, p := range diffPathsSet.UnsortedList() { // Iterate over the set's elements
		// SavedQuery only allows updating 'description', 'labels', 'content'
		if p == "description" || p == "labels" || p == "content" {
			mutablePathsSet.Insert(p) // Insert the string path
		}
	}

	if mutablePathsSet.Len() == 0 { // Check the size of the set
		log.V(2).Info("no mutable fields need update", "name", a.id)
		// Update status even if no fields changed in spec
		status := AssetSavedQueryStatus_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range mutablePathsSet {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMaskPaths := mutablePathsSet.UnsortedList() // Get the slice for the field mask
	log.V(2).Info("updating asset saved query fields", "paths", updateMaskPaths)

	// Populate the updateProto with desired values for the paths being updated
	if mutablePathsSet.Has("description") { // Keep using the set for efficient checks
		updateProto.Description = resource.Description
	}
	if mutablePathsSet.Has("labels") {
		updateProto.Labels = resource.Labels
	}
	if mutablePathsSet.Has("content") {
		updateProto.Content = resource.Content
	}

	updateMask := &fieldmaskpb.FieldMask{Paths: updateMaskPaths} // Use the slice here
	req := &pb.UpdateSavedQueryRequest{
		SavedQuery: updateProto,
		UpdateMask: updateMask,
	}
	actual, err := a.gcpClient.UpdateSavedQuery(ctx, req)
	if err != nil {
		return fmt.Errorf("updating asset saved query %s: %w", a.id.String(), err)
	}
	a.actual = actual // Update internal actual state with the response
	log.V(2).Info("successfully updated asset saved query", "name", a.id)

	status := AssetSavedQueryStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *savedQueryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() must be called before Export()")
	}
	mapCtx := &direct.MapContext{}

	obj := &krm.AssetSavedQuery{}
	obj.Spec = direct.ValueOf(AssetSavedQuerySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set parent reference from the 'name' field
	parentRef, id, err := krm.ParseSavedQueryExternal(a.actual.Name)
	if err != nil {
		return nil, fmt.Errorf("parsing parent from name %q: %w", a.actual.Name, err)
	}
	if parentRef.ProjectID != "" {
		obj.Spec.Parent.ProjectRef = &refs.ProjectRef{External: parentRef.String()}
	} else if parentRef.FolderID != "" {
		obj.Spec.Parent.FolderRef = &refs.FolderRef{External: parentRef.String()}
	} else if parentRef.OrganizationID != "" {
		obj.Spec.Parent.OrganizationRef = &refs.OrganizationRef{External: parentRef.String()}
	} else {
		return nil, fmt.Errorf("unknown parent type in name %q", a.actual.Name)
	}

	// Set observed state
	status := AssetSavedQueryStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Status = *status

	// Convert to unstructured
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting AssetSavedQuery to unstructured: %w", err)
	}
	u := &unstructured.Unstructured{Object: uObj}

	// Set standard Kubernetes metadata
	u.SetGroupVersionKind(krm.AssetSavedQueryGVK)
	u.SetName(id) // Use the parsed ID from the name field
	u.SetNamespace(a.desired.Namespace)
	// Retain labels and annotations from the original object if needed
	u.SetLabels(a.desired.Labels)
	u.SetAnnotations(a.desired.Annotations)

	// Set external reference annotation if available
	if status.ExternalRef != nil {
		annotations := u.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations["cnrm.cloud.google.com/external-ref"] = *status.ExternalRef
		u.SetAnnotations(annotations)
	}

	return u, nil
}

func (a *savedQueryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting asset saved query", "name", a.id)

	req := &pb.DeleteSavedQueryRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSavedQuery(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent asset saved query, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting asset saved query %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted asset saved query", "name", a.id)

	return true, nil
}

// Helper function to get owner reference. Needed for parent resolution.
func ownerFromObject(obj metav1.Object) *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}
