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

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CompareGreenfieldSpec executes the 5-step greenfield spec comparison pipeline:
//  1. actualProto -> actualKRM (specFromProto)
//  2. mergeUnsetFieldsGreenfield(clonedDesiredKRM, actualKRM) - adopts actual values for top-level/nested fields unspecified in KRM
//  3. Specs -> Protos (specToProto for desired & actual)
//  4. Optional Normalization (normalize callback; unlike CompareBrownfieldSpec, repeated fields are NOT globally sorted by default)
//  5. DiffForTopLevelFieldsGreenfield (generates structured Diff and FieldMask, skipping OUTPUT_ONLY and INPUT_ONLY fields and rejecting IMMUTABLE field changes)
//
// Unlike CompareBrownfieldSpec, CompareGreenfieldSpec returns mergedDesiredProto alongside the diff and field mask
// so callers can pass mergedDesiredProto in the Update request without wiping out server-defaulted sibling sub-fields
// inside nested structs (when UpdateMask is top-level) or top-level server defaults (when the Update RPC does not accept an UpdateMask).
func CompareGreenfieldSpec[SpecType any, ProtoT interface {
	proto.Message
}](
	ctx context.Context,
	desiredKRM *SpecType,
	actualProto ProtoT,
	specFromProto func(mapCtx *direct.MapContext, in ProtoT) *SpecType,
	specToProto func(mapCtx *direct.MapContext, in *SpecType) ProtoT,
	normalize func(ctx context.Context, pb ProtoT) error,
) (ProtoT, *structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var zeroProto ProtoT
	mapCtx := &direct.MapContext{}

	actualKRM := specFromProto(mapCtx, actualProto)
	if err := mapCtx.Err(); err != nil {
		return zeroProto, nil, nil, err
	}

	rawDesiredProto := specToProto(mapCtx, desiredKRM)
	if err := mapCtx.Err(); err != nil {
		return zeroProto, nil, nil, err
	}

	clonedDesiredKRM, err := cloneKRMSpec(desiredKRM)
	if err != nil {
		return zeroProto, nil, nil, err
	}
	mergeUnsetFieldsGreenfield(reflect.ValueOf(clonedDesiredKRM), reflect.ValueOf(actualKRM))

	desiredProtoMasked := specToProto(mapCtx, clonedDesiredKRM)
	if err := mapCtx.Err(); err != nil {
		return zeroProto, nil, nil, err
	}
	preserveDesiredOneofs(protoReflectOrNil(rawDesiredProto), protoReflectOrNil(desiredProtoMasked))

	actualProtoMasked := specToProto(mapCtx, actualKRM)
	if err := mapCtx.Err(); err != nil {
		return zeroProto, nil, nil, err
	}

	// Clone mergedDesiredProto before running diff-only normalization hooks so the caller receives
	// the exact merged proto (including INPUT_ONLY fields and original customer slice ordering) to send in the Update RPC.
	var mergedDesiredProto ProtoT
	if !isNilProto(desiredProtoMasked) {
		mergedDesiredProto = proto.Clone(desiredProtoMasked).(ProtoT)
		clearNonComparedProtoFields(proto.Message(desiredProtoMasked).ProtoReflect())
	}
	if !isNilProto(actualProtoMasked) {
		clearNonComparedProtoFields(proto.Message(actualProtoMasked).ProtoReflect())
	}

	if normalize != nil {
		if !isNilProto(desiredProtoMasked) {
			if err := normalize(ctx, desiredProtoMasked); err != nil {
				return zeroProto, nil, nil, err
			}
		}
		if !isNilProto(actualProtoMasked) {
			if err := normalize(ctx, actualProtoMasked); err != nil {
				return zeroProto, nil, nil, err
			}
		}
	}

	diff, updateMask, err := DiffForTopLevelFieldsGreenfield(
		ctx,
		protoReflectOrNil(desiredProtoMasked),
		protoReflectOrNil(actualProtoMasked),
	)
	if err != nil {
		return zeroProto, nil, nil, err
	}

	return mergedDesiredProto, diff, updateMask, nil
}

// CompareGreenfieldSpecAndLabels executes CompareGreenfieldSpec, populates the desired GCP labels on mergedDesiredProto,
// and compares metadata.labels against actualProto labels using the specified labels field name in the proto message.
func CompareGreenfieldSpecAndLabels[SpecType any, ProtoT interface {
	proto.Message
}](
	ctx context.Context,
	u *unstructured.Unstructured,
	desiredKRM *SpecType,
	actualProto ProtoT,
	labelsFieldName string,
	specFromProto func(mapCtx *direct.MapContext, in ProtoT) *SpecType,
	specToProto func(mapCtx *direct.MapContext, in *SpecType) ProtoT,
	normalize func(ctx context.Context, pb ProtoT) error,
) (ProtoT, *structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mergedDesiredProto, diff, updateMask, err := CompareGreenfieldSpec(ctx, desiredKRM, actualProto, specFromProto, specToProto, normalize)
	if err != nil {
		var zeroProto ProtoT
		return zeroProto, nil, nil, err
	}

	desiredLabels := label.GCPLabels(u)
	setProtoLabels(mergedDesiredProto, labelsFieldName, desiredLabels)
	actualLabels := extractProtoLabels(actualProto, labelsFieldName)
	diff, updateMask = appendLabelDiffIfChanged(diff, updateMask, labelsFieldName, actualLabels, desiredLabels)

	return mergedDesiredProto, diff, updateMask, nil
}

// DiffForTopLevelFieldsGreenfield compares top-level fields of desired and actual messages,
// skipping OUTPUT_ONLY and INPUT_ONLY fields (including nested ones) and returning an error if an IMMUTABLE field has changed.
func DiffForTopLevelFieldsGreenfield(ctx context.Context, desired protoreflect.Message, actual protoreflect.Message) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	diff := &structuredreporting.Diff{
		Controller: k8s.ReconcilerTypeDirect,
	}
	if (desired == nil || !desired.IsValid()) && (actual == nil || !actual.IsValid()) {
		return diff, &fieldmaskpb.FieldMask{}, nil
	}
	desired, actual = prepareMessagesForGreenfieldDiff(desired, actual)

	fields := actual.Type().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if IsFieldBehavior(fd, annotations.FieldBehavior_OUTPUT_ONLY) || IsFieldBehavior(fd, annotations.FieldBehavior_INPUT_ONLY) {
			continue
		}
		path := string(fd.Name())
		fieldDiff := fieldHasChanged(ctx, path, desired, actual)
		if fieldDiff == nil {
			continue
		}
		if IsFieldBehavior(fd, annotations.FieldBehavior_IMMUTABLE) {
			return nil, nil, fmt.Errorf("change to immutable field %s", path)
		}
		diff.AddProtoField(fieldDiff.FieldPath, fd, valToAny(fieldDiff.ActualValue), valToAny(fieldDiff.DesiredValue))
	}

	slices.SortFunc(diff.Fields, func(a, b structuredreporting.DiffField) int {
		return strings.Compare(a.ID, b.ID)
	})

	var paths []string
	for _, field := range diff.Fields {
		paths = append(paths, field.ID)
	}

	return diff, &fieldmaskpb.FieldMask{Paths: paths}, nil
}

func prepareMessagesForGreenfieldDiff(desired, actual protoreflect.Message) (protoreflect.Message, protoreflect.Message) {
	if desired == nil || !desired.IsValid() {
		desired = actual.Type().New()
	} else {
		desired = proto.Clone(desired.Interface()).ProtoReflect()
	}
	if actual == nil || !actual.IsValid() {
		actual = desired.Type().New()
	} else {
		actual = proto.Clone(actual.Interface()).ProtoReflect()
	}
	clearNonComparedProtoFields(desired)
	clearNonComparedProtoFields(actual)
	return desired, actual
}

func mergeUnsetFieldsGreenfield(desired, actual reflect.Value) {
	if !desired.IsValid() || !actual.IsValid() {
		return
	}

	switch desired.Kind() {
	case reflect.Ptr:
		if desired.IsNil() && !actual.IsNil() {
			if desired.CanSet() {
				desired.Set(actual)
			}
		} else if !desired.IsNil() && !actual.IsNil() {
			mergeUnsetFieldsGreenfield(desired.Elem(), actual.Elem())
		}

	case reflect.Struct:
		for i := 0; i < desired.NumField(); i++ {
			field := desired.Type().Field(i)
			if field.IsExported() {
				mergeUnsetFieldsGreenfield(desired.Field(i), actual.Field(i))
			}
		}

	case reflect.Slice:
		if desired.IsNil() && !actual.IsNil() && desired.CanSet() {
			desired.Set(actual)
		} else if !desired.IsNil() && !actual.IsNil() && desired.Len() == actual.Len() {
			for i := 0; i < desired.Len(); i++ {
				if sliceElemHasNoConflicts(desired.Index(i), actual.Index(i)) {
					mergeUnsetFieldsGreenfield(desired.Index(i), actual.Index(i))
				}
			}
		}

	case reflect.Map:
		mergeUnsetMapFields(desired, actual)
	}
}

// sliceElemHasNoConflicts reports whether every non-nil field set in d is also
// present and equal in a. This ensures server-populated defaults are only merged
// into slice elements that match the existing element at the same index, avoiding
// cross-contamination when list elements are reordered or modified.
func sliceElemHasNoConflicts(d, a reflect.Value) bool {
	if !d.IsValid() || !a.IsValid() {
		return true
	}

	switch d.Kind() {
	case reflect.Ptr:
		if d.IsNil() {
			return true
		}
		if a.IsNil() {
			return false
		}
		return sliceElemHasNoConflicts(d.Elem(), a.Elem())

	case reflect.Struct:
		for i := 0; i < d.NumField(); i++ {
			if d.Type().Field(i).IsExported() {
				if !sliceElemHasNoConflicts(d.Field(i), a.Field(i)) {
					return false
				}
			}
		}
		return true

	case reflect.Slice:
		if d.IsNil() {
			return true
		}
		if a.IsNil() || d.Len() != a.Len() {
			return false
		}
		for i := 0; i < d.Len(); i++ {
			if !sliceElemHasNoConflicts(d.Index(i), a.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Map:
		if d.IsNil() {
			return true
		}
		if a.IsNil() {
			return false
		}
		for _, k := range d.MapKeys() {
			aVal := a.MapIndex(k)
			if !aVal.IsValid() || !sliceElemHasNoConflicts(d.MapIndex(k), aVal) {
				return false
			}
		}
		return true

	default:
		return reflect.DeepEqual(d.Interface(), a.Interface())
	}
}

func mergeUnsetMapFields(desired, actual reflect.Value) {
	if desired.IsNil() && !actual.IsNil() {
		if desired.CanSet() {
			desired.Set(actual)
		}
		return
	}
	if desired.IsNil() || actual.IsNil() {
		return
	}
	elemKind := desired.Type().Elem().Kind()
	isStructElem := elemKind == reflect.Struct || (elemKind == reflect.Ptr && desired.Type().Elem().Elem().Kind() == reflect.Struct)
	if !isStructElem {
		return
	}
	for _, k := range desired.MapKeys() {
		dVal := desired.MapIndex(k)
		aVal := actual.MapIndex(k)
		if !dVal.IsValid() || !aVal.IsValid() {
			continue
		}
		if elemKind == reflect.Ptr {
			if dVal.IsNil() && !aVal.IsNil() {
				desired.SetMapIndex(k, aVal)
			} else if !dVal.IsNil() && !aVal.IsNil() {
				mergeUnsetFieldsGreenfield(dVal, aVal)
			}
		} else {
			dCopy := reflect.New(dVal.Type()).Elem()
			dCopy.Set(dVal)
			mergeUnsetFieldsGreenfield(dCopy, aVal)
			desired.SetMapIndex(k, dCopy)
		}
	}
}

func cloneKRMSpec[SpecType any](desiredKRM *SpecType) (*SpecType, error) {
	if desiredKRM == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(desiredKRM)
	if err != nil {
		return nil, fmt.Errorf("cloning desiredKRM: %w", err)
	}
	cloned := new(SpecType)
	if err := json.Unmarshal(bytes, cloned); err != nil {
		return nil, fmt.Errorf("cloning desiredKRM: %w", err)
	}
	preserveNonNilEmptyCollections(reflect.ValueOf(desiredKRM), reflect.ValueOf(cloned))
	return cloned, nil
}

// preserveNonNilEmptyCollections walks src and dst in lockstep after JSON round-trip cloning
// and restores non-nil empty slices ([]T{}) and non-nil empty maps (map[K]V{}) that json.Marshal
// omitted due to `omitempty` struct tags.
func preserveNonNilEmptyCollections(src, dst reflect.Value) {
	if !src.IsValid() || !dst.IsValid() {
		return
	}
	switch src.Kind() {
	case reflect.Ptr:
		if !src.IsNil() && !dst.IsNil() {
			preserveNonNilEmptyCollections(src.Elem(), dst.Elem())
		}
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			if src.Type().Field(i).IsExported() {
				preserveNonNilEmptyCollections(src.Field(i), dst.Field(i))
			}
		}
	case reflect.Slice:
		if !src.IsNil() && src.Len() == 0 && dst.IsNil() && dst.CanSet() {
			dst.Set(reflect.MakeSlice(src.Type(), 0, 0))
		} else if !src.IsNil() && !dst.IsNil() && src.Len() == dst.Len() {
			for i := 0; i < src.Len(); i++ {
				preserveNonNilEmptyCollections(src.Index(i), dst.Index(i))
			}
		}
	case reflect.Map:
		if !src.IsNil() && src.Len() == 0 && dst.IsNil() && dst.CanSet() {
			dst.Set(reflect.MakeMap(src.Type()))
		} else if !src.IsNil() && !dst.IsNil() && src.Len() > 0 {
			preserveNonNilEmptyCollectionsInMap(src, dst)
		}
	}
}

func preserveNonNilEmptyCollectionsInMap(src, dst reflect.Value) {
	for _, k := range src.MapKeys() {
		sVal := src.MapIndex(k)
		dVal := dst.MapIndex(k)
		if !sVal.IsValid() || !dVal.IsValid() {
			continue
		}
		if sVal.Kind() == reflect.Struct {
			dCopy := reflect.New(dVal.Type()).Elem()
			dCopy.Set(dVal)
			preserveNonNilEmptyCollections(sVal, dCopy)
			dst.SetMapIndex(k, dCopy)
		} else if sVal.Kind() == reflect.Slice && !sVal.IsNil() && sVal.Len() == 0 && dVal.IsNil() {
			dst.SetMapIndex(k, reflect.MakeSlice(sVal.Type(), 0, 0))
		} else if sVal.Kind() == reflect.Map && !sVal.IsNil() && sVal.Len() == 0 && dVal.IsNil() {
			dst.SetMapIndex(k, reflect.MakeMap(sVal.Type()))
		} else {
			preserveNonNilEmptyCollections(sVal, dVal)
		}
	}
}

// preserveDesiredOneofs ensures that any protobuf oneof explicitly selected in rawDesired
// wins over stale sibling oneof fields that mergeUnsetFieldsGreenfield copied from actualKRM into clonedDesiredKRM.
func preserveDesiredOneofs(rawDesired, mergedDesired protoreflect.Message) {
	if rawDesired == nil || !rawDesired.IsValid() || mergedDesired == nil || !mergedDesired.IsValid() {
		return
	}
	oneofs := rawDesired.Descriptor().Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneof := oneofs.Get(i)
		if oneof.IsSynthetic() {
			continue
		}
		rawField := rawDesired.WhichOneof(oneof)
		mergedField := mergedDesired.WhichOneof(oneof)
		if rawField != nil && mergedField != rawField {
			restoreOneofField(rawDesired, mergedDesired, rawField)
		}
	}
	rawDesired.Range(func(fd protoreflect.FieldDescriptor, rawVal protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind && mergedDesired.Has(fd) {
			preserveDesiredOneofsInField(fd, rawVal, mergedDesired.Get(fd))
		}
		return true
	})
}

func restoreOneofField(rawDesired, mergedDesired protoreflect.Message, fd protoreflect.FieldDescriptor) {
	rawVal := rawDesired.Get(fd)
	if fd.Kind() == protoreflect.MessageKind && !fd.IsList() && !fd.IsMap() {
		clonedMsg := proto.Clone(rawVal.Message().Interface())
		mergedDesired.Set(fd, protoreflect.ValueOfMessage(clonedMsg.ProtoReflect()))
		return
	}
	mergedDesired.Set(fd, rawVal)
}

func preserveDesiredOneofsInField(fd protoreflect.FieldDescriptor, rawVal, mergedVal protoreflect.Value) {
	if fd.IsList() {
		rawList := rawVal.List()
		mergedList := mergedVal.List()
		for i := 0; i < rawList.Len() && i < mergedList.Len(); i++ {
			preserveDesiredOneofs(rawList.Get(i).Message(), mergedList.Get(i).Message())
		}
		return
	}
	if fd.IsMap() {
		if fd.MapValue().Kind() == protoreflect.MessageKind {
			rawMap := rawVal.Map()
			mergedMap := mergedVal.Map()
			rawMap.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				if mergedMap.Has(k) {
					preserveDesiredOneofs(v.Message(), mergedMap.Get(k).Message())
				}
				return true
			})
		}
		return
	}
	preserveDesiredOneofs(rawVal.Message(), mergedVal.Message())
}

func isNilProto(msg proto.Message) bool {
	if msg == nil {
		return true
	}
	rv := reflect.ValueOf(msg)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}

func protoReflectOrNil(msg proto.Message) protoreflect.Message {
	if isNilProto(msg) {
		return nil
	}
	return msg.ProtoReflect()
}

// clearNonComparedProtoFields recursively clears any populated field marked OUTPUT_ONLY or INPUT_ONLY
// on msg and its nested messages, lists of messages, and maps of messages.
func clearNonComparedProtoFields(msg protoreflect.Message) {
	if msg == nil || !msg.IsValid() {
		return
	}
	var toClear []protoreflect.FieldDescriptor
	msg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		if IsFieldBehavior(fd, annotations.FieldBehavior_OUTPUT_ONLY) || IsFieldBehavior(fd, annotations.FieldBehavior_INPUT_ONLY) {
			toClear = append(toClear, fd)
			return true
		}
		if fd.Kind() == protoreflect.MessageKind {
			clearNonComparedNestedField(fd, val)
		}
		return true
	})
	for _, fd := range toClear {
		msg.Clear(fd)
	}
}

func clearNonComparedNestedField(fd protoreflect.FieldDescriptor, val protoreflect.Value) {
	if fd.IsList() {
		list := val.List()
		for i := 0; i < list.Len(); i++ {
			clearNonComparedProtoFields(list.Get(i).Message())
		}
		return
	}
	if fd.IsMap() {
		if fd.MapValue().Kind() == protoreflect.MessageKind {
			val.Map().Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool {
				clearNonComparedProtoFields(v.Message())
				return true
			})
		}
		return
	}
	clearNonComparedProtoFields(val.Message())
}

func appendLabelDiffIfChanged(
	diff *structuredreporting.Diff,
	updateMask *fieldmaskpb.FieldMask,
	labelsFieldName string,
	actualLabels, desiredLabels map[string]string,
) (*structuredreporting.Diff, *fieldmaskpb.FieldMask) {
	if maps.Equal(actualLabels, desiredLabels) {
		return diff, updateMask
	}
	if updateMask == nil {
		updateMask = &fieldmaskpb.FieldMask{}
	}
	if !slices.Contains(updateMask.Paths, labelsFieldName) {
		updateMask.Paths = append(updateMask.Paths, labelsFieldName)
		slices.Sort(updateMask.Paths)
	}
	if diff == nil {
		diff = &structuredreporting.Diff{}
	}
	hasLabelDiffField := slices.ContainsFunc(diff.Fields, func(f structuredreporting.DiffField) bool {
		return f.ID == labelsFieldName
	})
	if !hasLabelDiffField {
		diff.AddField(labelsFieldName, actualLabels, desiredLabels)
		slices.SortFunc(diff.Fields, func(a, b structuredreporting.DiffField) int {
			return strings.Compare(a.ID, b.ID)
		})
	}
	return diff, updateMask
}

func setProtoLabels(msg proto.Message, labelsFieldName string, labels map[string]string) {
	if isNilProto(msg) || labelsFieldName == "" {
		return
	}
	fd := msg.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(labelsFieldName))
	if fd == nil || !fd.IsMap() || fd.MapKey().Kind() != protoreflect.StringKind || fd.MapValue().Kind() != protoreflect.StringKind {
		return
	}
	if len(labels) == 0 {
		msg.ProtoReflect().Clear(fd)
		return
	}
	m := msg.ProtoReflect().Mutable(fd).Map()
	m.Range(func(k protoreflect.MapKey, _ protoreflect.Value) bool {
		m.Clear(k)
		return true
	})
	for k, v := range labels {
		m.Set(protoreflect.ValueOfString(k).MapKey(), protoreflect.ValueOfString(v))
	}
}

// NormalizeGreenfieldReferences resolves all KRM references on obj, including:
//   - *refs.ProjectRef
//   - refs.Ref
//   - refs.ExternalNormalizer (skipping already-normalized and zero-value references so NormalizedExternal calls are idempotent)
func NormalizeGreenfieldReferences(ctx context.Context, reader client.Reader, obj client.Object, projectRef *refs.ProjectIdentity, opts ...NormalizeOption) error {
	base := &refNormalizer{
		ctx:     ctx,
		kube:    reader,
		src:     obj,
		project: projectRef,
	}
	for _, opt := range opts {
		opt(base)
	}
	return VisitFields(obj, &greenfieldRefNormalizer{refNormalizer: base})
}

type greenfieldRefNormalizer struct {
	*refNormalizer
}

func (r *greenfieldRefNormalizer) VisitField(path string, v any) error {
	if r.skipPaths != nil {
		if r.skipPaths[path] || r.skipPaths[strings.TrimPrefix(path, ".")] {
			return nil
		}
	}

	if _, ok := v.(*refs.ProjectRef); ok {
		return r.refNormalizer.VisitField(path, v)
	}

	if _, ok := v.(refs.Ref); ok {
		return r.refNormalizer.VisitField(path, v)
	}

	if extNorm, ok := v.(refs.ExternalNormalizer); ok {
		if isNonEmptyExternalNormalizer(v) {
			if _, err := extNorm.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace()); err != nil {
				return err
			}
		}
		return nil
	}

	return nil
}

func isNonEmptyExternalNormalizer(v any) bool {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return true
	}
	nameField := rv.FieldByName("Name")
	extField := rv.FieldByName("External")
	if nameField.IsValid() || extField.IsValid() {
		if nameField.IsValid() && extField.IsValid() && !nameField.IsZero() && !extField.IsZero() {
			// Legacy ExternalNormalizer implementations set r.External in-place without clearing r.Name,
			// and return "cannot specify both name and external" if NormalizedExternal is called a second time
			// (both when VisitFields visits the pointer and then its dereferenced struct, and when
			// GetIdentity/New*Identity normalizes the parent ref before NormalizeGreenfieldReferences runs).
			return false
		}
		return (nameField.IsValid() && !nameField.IsZero()) || (extField.IsValid() && !extField.IsZero())
	}
	for i := 0; i < rv.NumField(); i++ {
		if !rv.Field(i).IsZero() {
			return true
		}
	}
	return false
}
