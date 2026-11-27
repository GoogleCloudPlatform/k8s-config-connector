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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FirestoreIndexGVK, NewFirestoreIndexModel)
}

func NewFirestoreIndexModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreIndexModel{config: config}, nil
}

type firestoreIndexModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &firestoreIndexModel{}

func (m *firestoreIndexModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreIndex{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := FirestoreIndexSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// queryScope defaults to collection (and must be specified)
	if desired.QueryScope == pb.Index_QUERY_SCOPE_UNSPECIFIED {
		desired.QueryScope = pb.Index_COLLECTION
	}

	return &firestoreIndexAdapter{
		id:                   id.(*krm.FirestoreIndexIdentity),
		firestoreAdminClient: firestoreAdminClient,
		desired:              desired,
		acquisition:          direct.ValueOf(obj.Spec.ResourceID) != "",
	}, nil
}

func (m *firestoreIndexModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//firestore.googleapis.com/")

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
	id                   *krm.FirestoreIndexIdentity
	firestoreAdminClient *apiv1.FirestoreAdminClient
	desired              *pb.Index
	actual               *pb.Index
	acquisition          bool
}

var _ directbase.Adapter = &firestoreIndexAdapter{}

func (a *firestoreIndexAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil || a.id.Index == "" {
		return false, nil
	}

	fqn := a.id.String()
	log.V(2).Info("getting FirestoreIndex", "name", fqn)

	req := &pb.GetIndexRequest{Name: fqn}
	obj, err := a.firestoreAdminClient.GetIndex(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			if a.acquisition {
				return false, fmt.Errorf("FirestoreIndex %q can't be acquired: %w; unset 'spec.resourceID' if you want to create it", fqn, err)
			}
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreIndex %q: %w", fqn, err)
	}

	a.actual = obj
	return true, nil
}

func (a *firestoreIndexAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.Parent().String()
	log.V(2).Info("creating FirestoreIndex", "parent", parent)

	resource := proto.CloneOf(a.desired)

	req := &pb.CreateIndexRequest{
		Parent: parent,
		Index:  resource,
	}
	op, err := a.firestoreAdminClient.CreateIndex(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreIndex in %s: %w", parent, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of FirestoreIndex: %w", err)
	}

	// Update the ID since it was generated
	if err := a.id.FromExternal("//firestore.googleapis.com/" + created.Name); err != nil {
		return fmt.Errorf("failed to parse generated index name %q: %w", created.Name, err)
	}

	log.V(2).Info("successfully created FirestoreIndex", "name", created.Name)

	status := &krm.FirestoreIndexStatus{}
	status.Name = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreIndexAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// All fields in FirestoreIndexSpec are immutable. An explicitly acquired
	// index must already have the requested definition because it cannot be
	// updated in place.
	if a.acquisition && !firestoreIndexDefinitionsEqual(a.desired, a.actual) {
		return fmt.Errorf("cannot acquire FirestoreIndex %q: existing index definition does not match spec", a.id.String())
	}

	status := &krm.FirestoreIndexStatus{}
	status.Name = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func firestoreIndexDefinitionsEqual(desired, actual *pb.Index) bool {
	if desired == nil || actual == nil {
		return desired == actual
	}

	desired = normalizeFirestoreIndexDefinition(desired)
	actual = normalizeFirestoreIndexDefinition(actual)
	return proto.Equal(desired, actual)
}

func normalizeFirestoreIndexDefinition(index *pb.Index) *pb.Index {
	index = proto.CloneOf(index)
	index.Name = ""
	index.State = pb.Index_STATE_UNSPECIFIED
	index.Density = pb.Index_DENSITY_UNSPECIFIED
	index.Multikey = false
	index.ShardCount = 0
	index.Unique = false

	if index.QueryScope == pb.Index_QUERY_SCOPE_UNSPECIFIED {
		index.QueryScope = pb.Index_COLLECTION
	}

	if len(index.Fields) > 0 && index.Fields[len(index.Fields)-1].GetFieldPath() != "__name__" {
		order := pb.Index_IndexField_ASCENDING
		lastOrder := index.Fields[len(index.Fields)-1].GetOrder()
		if index.ApiScope != pb.Index_DATASTORE_MODE_API && lastOrder != pb.Index_IndexField_ORDER_UNSPECIFIED {
			order = lastOrder
		}
		index.Fields = append(index.Fields, &pb.Index_IndexField{
			FieldPath: "__name__",
			ValueMode: &pb.Index_IndexField_Order_{Order: order},
		})
	}

	return index
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
	obj.Spec.ResourceID = direct.PtrTo(a.id.Index)
	obj.Spec.Collection = a.id.CollectionGroup
	if a.id.Database != "(default)" {
		obj.Spec.Database = direct.PtrTo(a.id.Database)
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreIndex to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	return u, nil
}

func (a *firestoreIndexAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id == nil || a.id.Index == "" {
		return true, nil
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
