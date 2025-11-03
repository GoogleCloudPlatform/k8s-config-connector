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

package firestore

import (
	"context"
	"fmt"
	"strings"

	api "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.FirestoreIndexGVK, NewFirestoreIndexModel)
}

func NewFirestoreIndexModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreIndexModel{config: config}, nil
}

var _ directbase.Model = &firestoreIndexModel{}

type firestoreIndexModel struct {
	config *config.ControllerConfig
}

func (m *firestoreIndexModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreIndex{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references for FirestoreIndex %q: %w", u.GetName(), err)
	}

	adapter := &firestoreIndexAdapter{
		firestoreAdminClient: firestoreAdminClient,
	}

	{
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		// Server-generated identity; id may be nil.
		if idFromObject != nil {
			adapter.id = idFromObject.(*krm.FirestoreIndexIdentity)
		}
	}

	// Resolve the parent, needed for create
	{
		projectID, err := v1beta1.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("resolving project ID for FirestoreIndex %q: %w", u.GetName(), err)
		}
		database := direct.ValueOf(obj.Spec.Database)
		collectionGroup := obj.Spec.Collection

		if projectID == "" {
			return nil, fmt.Errorf("project ID is required for FirestoreIndex %q", u.GetName())
		}
		if database == "" {
			return nil, fmt.Errorf("database is required for FirestoreIndex %q", u.GetName())
		}
		if collectionGroup == "" {
			return nil, fmt.Errorf("collectionGroup is required for FirestoreIndex %q", u.GetName())
		}

		parent := "projects/" + projectID + "/databases/" + database + "/collectionGroups/" + collectionGroup
		parentID := &krm.CollectionGroupIdentity{}
		if err := parentID.FromExternal(parent); err != nil {
			return nil, fmt.Errorf("parsing parent from FirestoreIndex.Parent=%q: %w", parent, err)
		}
		if adapter.id != nil && parentID.String() != adapter.id.Parent.String() {
			return nil, fmt.Errorf("mismatched parent: %q vs %q", parentID.String(), adapter.id.String())
		}
		adapter.parent = parentID
	}

	mapCtx := &direct.MapContext{}
	desired := FirestoreIndexSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set (required) default value for QueryScope if not set.
	if desired.QueryScope == pb.Index_QUERY_SCOPE_UNSPECIFIED {
		desired.QueryScope = pb.Index_COLLECTION
	}

	adapter.desired = desired

	return adapter, nil
}

func (m *firestoreIndexModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	id := &krm.FirestoreIndexIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &firestoreIndexAdapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
	}, nil
}

type firestoreIndexAdapter struct {
	id *krm.FirestoreIndexIdentity

	// parent is used when creating a new index (because this is server-generated)
	parent *krm.CollectionGroupIdentity

	firestoreAdminClient *api.FirestoreAdminClient
	desired              *pb.Index
	actual               *pb.Index
}

var _ directbase.Adapter = &firestoreIndexAdapter{}

func (a *firestoreIndexAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(2).Info("getting FirestoreIndex", "name", fqn)

	req := &pb.GetIndexRequest{Name: fqn}
	obj, err := a.firestoreAdminClient.GetIndex(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreIndex %q: %w", fqn, err)
	}

	a.actual = obj
	return true, nil
}

func (a *firestoreIndexAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// Server generated identity, no id until after creation.

	parent := a.parent.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirestoreIndex", "parent", parent)

	desired := direct.ProtoClone(a.desired)

	req := &pb.CreateIndexRequest{
		Parent: parent,
		Index:  desired,
	}
	op, err := a.firestoreAdminClient.CreateIndex(ctx, req)
	if err != nil {
		log.Info("error creating FirestoreIndex", "req", req, "error", err)
		return fmt.Errorf("creating FirestoreIndex in %s: %w", parent, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of FirestoreIndex in %s: %w", parent, err)
	}

	log.V(2).Info("successfully created FirestoreIndex", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *firestoreIndexAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirestoreIndex", "name", fqn)

	latest := a.actual

	changedFieldPaths, err := common.CompareProtoMessage(a.actual, a.desired, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing FirestoreIndex %s: %w", fqn, err)
	}
	if len(changedFieldPaths) > 0 {
		return fmt.Errorf("firestore indexes cannot be updated; changed fields: %v", sets.List(changedFieldPaths))
		// resource := direct.ProtoClone(a.desired)
		// resource.Name = fqn
		// req := &pb.Patch{
		// 	Document: resource,
		// }
		// updated, err := a.firestoreAdminClient.UpdateIndex(ctx, req)
		// if err != nil {
		// 	return fmt.Errorf("updating FirestoreIndex %s: %w", fqn, err)
		// }
		// log.V(2).Info("successfully updated FirestoreIndex", "name", fqn)
		// latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *firestoreIndexAdapter) updateStatus(ctx context.Context, updateOp directbase.Operation, latest *pb.Index) error {
	status := &krm.FirestoreIndexStatus{}
	// NOTYET: status.ExternalRef not yet supported in firestoreindex terraform resource
	// status.ExternalRef = direct.PtrTo(a.id.String())
	mapCtx := &direct.MapContext{}
	status = FirestoreIndexStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreIndexAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreIndex %q not found", fqn)
	}

	mapCtx := &direct.MapContext{}
	objSpec := FirestoreIndexSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.FirestoreIndex{
		Spec: *objSpec,
	}

	obj.SetGroupVersionKind(krm.FirestoreIndexGVK)
	obj.Name = a.id.Index
	obj.Spec.Database = direct.PtrTo(a.id.Parent.Parent.Database)
	obj.Spec.Collection = a.id.Parent.CollectionGroup

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreIndex to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *firestoreIndexAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id == nil {
		// Nothing to delete
		return false, nil
	}

	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirestoreIndex", "name", fqn)

	req := &pb.DeleteIndexRequest{
		Name: fqn,
	}
	if err := a.firestoreAdminClient.DeleteIndex(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FirestoreIndex, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreIndex %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted FirestoreIndex", "name", fqn)

	return true, nil
}
