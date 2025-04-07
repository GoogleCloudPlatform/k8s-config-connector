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

// +tool:controller
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.EntryGroup
// crd.type: DataplexEntryGroup
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/dataplex/apiv1"
	catalogpb "cloud.google.com/go/dataplex/apiv1/catalogpb"
	pb "cloud.google.com/go/dataplex/apiv1/catalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexEntryGroupGVK, NewEntryGroupModel)
}

func NewEntryGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &entryGroupModel{config: config}, nil
}

var _ directbase.Model = &entryGroupModel{}

type entryGroupModel struct {
	config *config.ControllerConfig
}

func (m *entryGroupModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexEntryGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEntryGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	catalogClient, err := gcpClient.newCatalogClient(ctx)
	if err != nil {
		return nil, err
	}

	return &entryGroupAdapter{
		gcpClient: catalogClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *entryGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type entryGroupAdapter struct {
	gcpClient *gcp.CatalogClient
	id        *krm.EntryGroupIdentity
	desired   *krm.DataplexEntryGroup
	actual    *pb.EntryGroup
	reader    client.Reader
}

var _ directbase.Adapter = &entryGroupAdapter{}

func (a *entryGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex entry group", "name", a.id)

	req := &catalogpb.GetEntryGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEntryGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex entry group %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *entryGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex entry group", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DataplexEntryGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &catalogpb.CreateEntryGroupRequest{
		Parent:       a.id.Parent().String(),
		EntryGroupId: a.id.ID(),
		EntryGroup:   resource,
	}
	op, err := a.gcpClient.CreateEntryGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex entry group %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex entry group %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex entry group in gcp", "name", a.id)

	status := &krm.DataplexEntryGroupStatus{}
	status.ObservedState = DataplexEntryGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *entryGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex entry group", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := DataplexEntryGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String() // Add name for update request

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if !reflect.DeepEqual(resource.Etag, a.actual.Etag) && resource.Etag != "" { // etag is optional
		// Etag cannot be directly updated, it's used for optimistic concurrency control.
		// It seems the intention here is to match the actual etag for the update.
		// If the spec includes an etag, and it doesn't match the actual, we should fail.
		// The update request itself doesn't use the etag field directly in the resource,
		// but rather the Etag field in the UpdateEntryGroupRequest if needed (which we aren't using here).
		// For KCC, we generally don't set etag in Spec, but if it is set, it should match.
		// If it doesn't match, it implies the resource was modified outside KCC.
		log.V(2).Info("etag mismatch", "desired", resource.Etag, "actual", a.actual.Etag)
		// We don't add 'etag' to the updateMask.
		// If the user provided an etag in spec and it differs from actual, the update might implicitly fail
		// due to concurrency control on the server side if the server uses it.
		// However, the UpdateEntryGroupRequest doesn't seem to have an etag field for this purpose.
		// Let's proceed with other updates if any, or return if only etag differs.
	}

	// Remove Etag from the comparison as it's not part of the update mask paths.
	pathsWithoutEtag := make([]string, 0, len(updateMask.Paths))
	for _, p := range updateMask.Paths {
		if p != "etag" {
			pathsWithoutEtag = append(pathsWithoutEtag, p)
		}
	}
	updateMask.Paths = pathsWithoutEtag

	var updated *pb.EntryGroup
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		req := &catalogpb.UpdateEntryGroupRequest{
			UpdateMask: updateMask,
			EntryGroup: resource,
		}
		op, err := a.gcpClient.UpdateEntryGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex entry group %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex entry group %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex entry group", "name", a.id)
	}

	status := &krm.DataplexEntryGroupStatus{}
	status.ObservedState = DataplexEntryGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *entryGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexEntryGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexEntryGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID()) // Use the Entry Group ID as the K8s resource name
	u.SetGroupVersionKind(krm.DataplexEntryGroupGVK)
	u.Object = uObj

	return u, nil
}

func (a *entryGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex entry group", "name", a.id)

	req := &catalogpb.DeleteEntryGroupRequest{
		Name: a.id.String(),
		Etag: direct.ValueOf(a.desired.Spec.Etag), // Use etag from spec if provided for concurrency control
	}
	op, err := a.gcpClient.DeleteEntryGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent dataplex entry group, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex entry group %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion of dataplex entry group", "name", a.id)

	// Wait for the LRO to complete.
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex entry group %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex entry group", "name", a.id)

	return true, nil
}
