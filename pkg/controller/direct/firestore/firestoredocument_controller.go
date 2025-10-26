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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	apiv1 "cloud.google.com/go/firestore/apiv1"
	pb "cloud.google.com/go/firestore/apiv1/firestorepb"

	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.FirestoreDocumentGVK, NewFirestoreDocumentModel)
}

func NewFirestoreDocumentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreDocumentModel{config: config}, nil
}

var _ directbase.Model = &firestoreDocumentModel{}

type firestoreDocumentModel struct {
	config *config.ControllerConfig
}

func (m *firestoreDocumentModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	firestoreClient, err := newFirestoreClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreDocument{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := FirestoreDocumentSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &firestoreDocumentAdapter{
		id:              id.(*krm.FirestoreDocumentIdentity),
		firestoreClient: firestoreClient,
		desired:         desired,
	}, nil
}

func (m *firestoreDocumentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//firestore.googleapis.com/")

	id := &krm.FirestoreDocumentIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	firestoreClient, err := newFirestoreClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &firestoreDocumentAdapter{
		id:              id,
		firestoreClient: firestoreClient,
	}, nil
}

type firestoreDocumentAdapter struct {
	id              *krm.FirestoreDocumentIdentity
	firestoreClient *apiv1.Client
	desired         *pb.Document
	actual          *pb.Document
}

var _ directbase.Adapter = &firestoreDocumentAdapter{}

func (a *firestoreDocumentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting FirestoreDocument", "name", fqn)

	if a.id == nil {
		return false, nil
	}

	req := &pb.GetDocumentRequest{Name: fqn}
	obj, err := a.firestoreClient.GetDocument(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreDocument %q: %w", fqn, err)
	}

	a.actual = obj
	return true, nil
}

func (a *firestoreDocumentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirestoreDocument", "name", fqn)

	resource := direct.ProtoClone(a.desired)

	req := &pb.CreateDocumentRequest{
		Parent:       a.id.Parent.String() + "/documents",
		CollectionId: a.id.Collection,
		DocumentId:   a.id.Document,
		Document:     resource,
	}
	created, err := a.firestoreClient.CreateDocument(ctx, req)
	if err != nil {
		log.Info("error creating FirestoreDocument", "req", req, "error", err)
		return fmt.Errorf("creating FirestoreDocument %s: %w", fqn, err)
	}

	log.V(2).Info("successfully created FirestoreDocument", "name", fqn)

	status := &krm.FirestoreDocumentStatus{}
	status.ExternalRef = direct.PtrTo(a.id.String())
	mapCtx := &direct.MapContext{}
	status.ObservedState = FirestoreDocumentObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreDocumentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirestoreDocument", "name", fqn)

	resource := direct.ProtoClone(a.desired)

	latest := a.actual

	changedFieldPaths := changedFieldPaths(a.actual, a.desired)
	if len(changedFieldPaths) > 0 {
		resource.Name = fqn
		req := &pb.UpdateDocumentRequest{
			Document: resource,
		}
		updated, err := a.firestoreClient.UpdateDocument(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FirestoreDocument %s: %w", fqn, err)
		}
		log.V(2).Info("successfully updated FirestoreDocument", "name", fqn)
		latest = updated
	}

	status := &krm.FirestoreDocumentStatus{}
	status.ExternalRef = direct.PtrTo(a.id.String())
	mapCtx := &direct.MapContext{}
	status.ObservedState = FirestoreDocumentObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *firestoreDocumentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreDocument %q not found", fqn)
	}

	mapCtx := &direct.MapContext{}
	objSpec := FirestoreDocumentSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.FirestoreDocument{
		Spec: *objSpec,
	}

	obj.SetGroupVersionKind(krm.FirestoreDocumentGVK)
	obj.Name = a.id.Document
	obj.Spec.DatabaseRef.External = a.id.Parent.String()
	if a.id.Collection != "" {
		obj.Spec.Collection = direct.PtrTo(a.id.Collection)
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreDocument to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *firestoreDocumentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirestoreDocument", "name", fqn)

	req := &pb.DeleteDocumentRequest{
		Name: fqn,
	}
	if err := a.firestoreClient.DeleteDocument(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent FirestoreDocument, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreDocument %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted FirestoreDocument", "name", fqn)

	return true, nil
}

func changedFieldPaths(a, b *pb.Document) []string {
	var fieldPaths []string
	for k, valA := range a.Fields {
		valB, ok := b.Fields[k]
		if !ok {
			fieldPaths = append(fieldPaths, k)
		}
		if !proto.Equal(valA, valB) {
			fieldPaths = append(fieldPaths, k)
		}
	}
	for k := range a.Fields {
		_, ok := b.Fields[k]
		if !ok {
			fieldPaths = append(fieldPaths, k)
		}
	}
	return fieldPaths
}
