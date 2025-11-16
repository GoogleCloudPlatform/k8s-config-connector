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
	"slices"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.FirestoreFieldGVK, NewFirestoreFieldModel)
}

func NewFirestoreFieldModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &firestoreFieldModel{config: config}, nil
}

var _ directbase.Model = &firestoreFieldModel{}

type firestoreFieldModel struct {
	config *config.ControllerConfig
}

func (m *firestoreFieldModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.FirestoreField{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	var desired *pb.Field
	{
		mapCtx := &direct.MapContext{}
		desired = FirestoreFieldSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &firestoreFieldAdapter{
		id:                   id.(*krm.FirestoreFieldIdentity),
		firestoreAdminClient: firestoreAdminClient,
		desired:              desired,
	}, nil
}

func (m *firestoreFieldModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//firestore.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//firestore.googleapis.com/")

	id := &krm.FirestoreFieldIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	firestoreAdminClient, err := newFirestoreAdminClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &firestoreFieldAdapter{
		id:                   id,
		firestoreAdminClient: firestoreAdminClient,
	}, nil
}

type firestoreFieldAdapter struct {
	id                   *krm.FirestoreFieldIdentity
	firestoreAdminClient *apiv1.FirestoreAdminClient
	desired              *pb.Field
	actual               *pb.Field
}

var _ directbase.Adapter = &firestoreFieldAdapter{}

func (a *firestoreFieldAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting FirestoreField", "name", fqn)

	req := &pb.GetFieldRequest{Name: fqn}
	firestorefieldpb, err := a.firestoreAdminClient.GetField(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirestoreField %q: %w", fqn, err)
	}

	a.actual = firestorefieldpb
	return true, nil
}

func (a *firestoreFieldAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating FirestoreField", "name", fqn)

	req := &pb.UpdateFieldRequest{
		Field: direct.ProtoClone(a.desired),
	}
	req.Field.Name = fqn

	req.UpdateMask = a.allFields()

	op, err := a.firestoreAdminClient.UpdateField(ctx, req)
	if err != nil {
		return fmt.Errorf("creating FirestoreField %s: %w", fqn, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("FirestoreField %s waiting creation: %w", fqn, err)
	}
	log.V(0).Info("successfully created FirestoreField", "name", fqn)

	return a.updateStatus(ctx, createOp, created)
}

func (a *firestoreFieldAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	req := &pb.UpdateFieldRequest{
		Field: direct.ProtoClone(a.desired),
	}
	req.Field.Name = fqn

	updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for FirestoreField %q: %w", fqn, err)
	}

	req.UpdateMask = updateMask

	latest := a.desired
	if len(req.UpdateMask.Paths) != 0 {
		log.V(0).Info("updating FirestoreField", "name", fqn)

		op, err := a.firestoreAdminClient.UpdateField(ctx, req)
		if err != nil {
			return fmt.Errorf("updating FirestoreField %q: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of FirestoreField %s: %w", fqn, err)
		}
		log.V(0).Info("successfully updated FirestoreField", "name", fqn)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *firestoreFieldAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Field) error {
	status := &krm.FirestoreFieldStatus{}
	{
		mapCtx := &direct.MapContext{}
		status.ObservedState = FirestoreFieldObservedState_v1alpha1_FromProto(mapCtx, latest)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *firestoreFieldAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("FirestoreField %q not found", fqn)
	}

	obj := &krm.FirestoreField{}

	{
		mapCtx := &direct.MapContext{}
		spec := FirestoreFieldSpec_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.FirestoreFieldGVK)
	obj.Name = a.id.Field
	obj.Spec.DatabaseRef.External = a.id.Parent.Parent.String()
	obj.Spec.CollectionGroup = &a.id.Parent.CollectionGroup

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting FirestoreField to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *firestoreFieldAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting FirestoreField", "name", fqn)

	// There is no delete method per se, instead we clear indexConfig and ttlConfig via UpdateField.

	req := &pb.UpdateFieldRequest{}
	req.Field = &pb.Field{}
	req.Field.Name = fqn

	req.UpdateMask = a.allFields()

	// The index configuration for this field. If unset, field indexing will revert to the configuration defined by the ancestorField.
	req.Field.IndexConfig = nil

	// The TTL configuration for this field.
	req.Field.TtlConfig = nil

	op, err := a.firestoreAdminClient.UpdateField(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent FirestoreField, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting FirestoreField %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted FirestoreField", "name", fqn)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for delete of FirestoreField %s: %w", fqn, err)
	}

	// Primarily for debugging purposes, we get the field after deletion attempt.
	afterDelete, err := a.firestoreAdminClient.GetField(ctx, &pb.GetFieldRequest{Name: fqn})
	if err != nil {
		return false, fmt.Errorf("getting FirestoreField %q after deletion attempt: %w", fqn, err)
	}
	if !afterDelete.GetIndexConfig().GetUsesAncestorConfig() {
		return false, fmt.Errorf("FirestoreField %q still exists after deletion attempt", fqn)
	}

	return true, nil
}

func (a *firestoreFieldAdapter) allFields() *fieldmaskpb.FieldMask {
	return &fieldmaskpb.FieldMask{
		Paths: []string{
			"index_config",
			"ttl_config",
		},
	}
}

func (a *firestoreFieldAdapter) changedFields(ctx context.Context) (*fieldmaskpb.FieldMask, error) {
	log := klog.FromContext(ctx)

	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := FirestoreFieldSpec_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := FirestoreFieldSpec_v1alpha1_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	var paths []string
	for _, path := range a.allFields().Paths {
		changed, err := fieldHasChanged(path, a.desired.ProtoReflect(), actualMasked)
		if err != nil {
			log.Error(err, "error determining if field has changed", "field", path)
			// If we can't determine if the field has changed, include it in the update.
		} else if !changed {
			continue
		}
		paths = append(paths, path)
	}

	// Special case: the presence of ttl_config is used to enable ttl; our normal comparison treats nil and empty as the same.
	if !slices.Contains(paths, "ttl_config") {
		ttlConfigField := actualMasked.Descriptor().Fields().ByName("ttl_config")
		if actualMasked.Has(ttlConfigField) != a.desired.ProtoReflect().Has(ttlConfigField) {
			paths = append(paths, "ttl_config")
		}
	}

	return &fieldmaskpb.FieldMask{Paths: paths}, nil
}

func fieldHasChanged(fieldPath string, desired protoreflect.Message, actual protoreflect.Message) (bool, error) {
	actualField, foundActual, err := commonGetFieldByPath(actual, fieldPath)
	if err != nil {
		return true, err
	}
	desiredField, foundDesired, err := commonGetFieldByPath(desired, fieldPath)
	if err != nil {
		return true, err
	}
	if foundActual != foundDesired {
		klog.Infof("Field changed %q: foundActual=%v foundDesired=%v", fieldPath, foundActual, foundDesired)
		return true, nil
	}
	if !foundActual && !foundDesired {
		// Both unset
		return false, nil
	}
	if actualField.Equal(desiredField) {
		return false, nil
	}
	klog.Infof("Field changed %q: actual=%v desired=%v", fieldPath, format(actualField), format(desiredField))
	return true, nil
}

func commonGetFieldByPath(msg protoreflect.Message, fieldPath string) (protoreflect.Value, bool, error) {
	if msg == nil {
		return protoreflect.Value{}, false, nil
	}
	tokens := strings.SplitN(fieldPath, ".", 2)
	fieldName := protoreflect.Name(tokens[0])
	field := msg.Descriptor().Fields().ByName(fieldName)
	if field == nil {
		return protoreflect.Value{}, false, fmt.Errorf("field %q not found in %T", fieldName, msg)
	}
	v := msg.Get(field)
	if len(tokens) == 1 {
		return v, true, nil
	}
	switch field.Kind() {
	case protoreflect.MessageKind:
		return commonGetFieldByPath(v.Message(), tokens[1])
	default:
		return protoreflect.Value{}, false, fmt.Errorf("field %q in %T is not a message", fieldName, msg)
	}
}

func format(v protoreflect.Value) string {
	o := v.Interface()
	if msg, ok := o.(protoreflect.Message); ok {
		return prototext.Format(msg.Interface())
	}
	return fmt.Sprintf("[%T]:%v", o, o)
}
