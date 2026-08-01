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

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type CompareDiff func(fieldName protoreflect.Name, a, b proto.Message) (bool, error)

var BasicDiff = func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
	aField := a.ProtoReflect().Descriptor().Fields().ByName(fieldName)
	bField := b.ProtoReflect().Descriptor().Fields().ByName(fieldName)

	// Skip output-only fields
	if IsFieldBehavior(aField, annotations.FieldBehavior_OUTPUT_ONLY) {
		return false, nil
	}
	// The field is previously unset
	if bField == nil {
		return true, nil
	}
	if !aField.Kind().IsValid() {
		return false, fmt.Errorf("unimplemented kind: %s", aField.Kind().String())
	}

	diff := false
	aVal := a.ProtoReflect().Get(aField)
	bVal := b.ProtoReflect().Get(bField)
	switch aField.Kind() {
	case protoreflect.MessageKind:
		if aField.IsList() || aField.IsMap() {
			if !aVal.Equal(bVal) {
				diff = true
			}
		} else {
			m := aVal.Message().Interface()
			// Compare well-known proto type as a whole otherwise the diffPath (update field mask) could be wrong.
			switch m.(type) {
			case *timestamppb.Timestamp:
				if !aVal.Equal(bVal) {
					diff = true
				}
			case *durationpb.Duration:
				if !aVal.Equal(bVal) {
					diff = true
				}
			default:
				return false, fmt.Errorf("field %s not recursed", fieldName)
			}
		}
	default:
		if !aVal.Equal(bVal) {
			diff = true
		}
	}
	if diff && IsFieldBehavior(aField, annotations.FieldBehavior_IMMUTABLE) {
		return false, fmt.Errorf("change to immutable field %s", fieldName)
	}
	return diff, nil
}

func CompareProtoMessage(a, b proto.Message, compareDiff CompareDiff) (sets.Set[string], error) {
	diffPaths := sets.Set[string]{}
	aDescriptor := a.ProtoReflect().Descriptor()

	for i := 0; i < aDescriptor.Fields().Len(); i++ {
		field := aDescriptor.Fields().Get(i)
		updatePath := updatePathFromField(a, aDescriptor.Fields().Get(i))

		aVal := a.ProtoReflect().Get(field)
		bVal := b.ProtoReflect().Get(field)
		if shouldRecurse(field, a.ProtoReflect()) {
			subPaths, err := CompareProtoMessage(aVal.Message().Interface(), bVal.Message().Interface(), compareDiff)
			if err != nil {
				return nil, err
			}
			for path := range subPaths {
				diffPaths.Insert(updatePath + "." + path)
			}
		} else {
			if diff, err := compareDiff(field.Name(), a, b); err != nil {
				return nil, err
			} else if diff {
				diffPaths.Insert(updatePath)
			}
		}
	}
	return diffPaths, nil
}

// CompareProtoMessageStructuredDiff computes the diff between two proto messages and returns both the set of changed field paths and a structured Diff object.
// Deprecated: CompareProtoMessageStructuredDiff is obsolete; please use common.DiffForTopLevelFields instead.
func CompareProtoMessageStructuredDiff(a, b proto.Message, compareDiff CompareDiff) (sets.Set[string], *structuredreporting.Diff, error) {
	diffPaths := sets.Set[string]{}
	diff := &structuredreporting.Diff{}
	aDescriptor := a.ProtoReflect().Descriptor()

	for i := 0; i < aDescriptor.Fields().Len(); i++ {
		field := aDescriptor.Fields().Get(i)
		updatePath := updatePathFromField(a, field)

		aVal := a.ProtoReflect().Get(field)
		bVal := b.ProtoReflect().Get(field)
		if shouldRecurse(field, a.ProtoReflect()) {
			subPaths, subDiff, err := CompareProtoMessageStructuredDiff(aVal.Message().Interface(), bVal.Message().Interface(), compareDiff)
			if err != nil {
				return nil, nil, err
			}
			for path := range subPaths {
				diffPaths.Insert(updatePath + "." + path)
			}
			if subDiff != nil {
				for _, d := range subDiff.Fields {
					diff.AddField(updatePath+"."+d.ID, d.Old, d.New)
				}
			}
		} else {
			if hasDiff, err := compareDiff(field.Name(), a, b); err != nil {
				return nil, nil, err
			} else if hasDiff {
				diffPaths.Insert(updatePath)
				diff.AddField(updatePath, bVal.Interface(), aVal.Interface())
			}
		}
	}
	return diffPaths, diff, nil
}

func shouldRecurse(field protoreflect.FieldDescriptor, message protoreflect.Message) bool {
	if field.Kind() != protoreflect.MessageKind {
		return false
	}
	if field.IsList() || field.IsMap() {
		return false
	}
	m := message.Get(field).Message().Interface()
	// Compare well-known proto type as a whole otherwise the diffPath (update field mask) could be wrong.
	switch m.(type) {
	case *timestamppb.Timestamp:
		return false
	case *durationpb.Duration:
		return false
	default:
		return true
	}
}

func updatePathFromField(obj proto.Message, field protoreflect.FieldDescriptor) string {
	d := obj.ProtoReflect().Descriptor()
	return strings.TrimPrefix(string(field.FullName()), string(d.FullName())+".")
}

func IsFieldBehavior(field protoreflect.FieldDescriptor, fieldBehavior annotations.FieldBehavior) bool {
	d := field.Options()
	fieldBehaviors := proto.GetExtension(d, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
	for _, f := range fieldBehaviors {
		if f == fieldBehavior {
			return true
		}
	}
	return false
}

func DiffForTopLevelFields(ctx context.Context, desired protoreflect.Message, actual protoreflect.Message) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	diff := &structuredreporting.Diff{
		Controller: k8s.ReconcilerTypeDirect,
	}

	fields := actual.Type().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		path := string(fields.Get(i).Name())
		fieldDiff := fieldHasChanged(ctx, path, desired, actual)
		if fieldDiff == nil {
			continue
		}
		diff.AddProtoField(fieldDiff.FieldPath, fields.Get(i), valToAny(fieldDiff.ActualValue), valToAny(fieldDiff.DesiredValue))
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

type FieldChange struct {
	FieldPath    string
	ActualValue  protoreflect.Value
	DesiredValue protoreflect.Value
}

// fieldHasChanged compares the field at fieldPath in desired and actual messages.
// It returns a FieldChange if the field has changed, or nil if it has not changed.
// If there is an error retrieving the field, it returns the FieldChange with whatever
// values could be retrieved; the error is logged.
// If we can't prove that the field is unchanged, we assume it has changed.
func fieldHasChanged(ctx context.Context, fieldPath string, desired protoreflect.Message, actual protoreflect.Message) *FieldChange {
	log := klog.FromContext(ctx)

	change := &FieldChange{FieldPath: fieldPath}

	actualValue, foundActual, err := commonGetFieldByPath(actual, fieldPath)
	if err != nil {
		log.Error(err, "error fetching previous field value", "field", fieldPath)
		return change
	}
	change.ActualValue = actualValue

	desiredValue, foundDesired, err := commonGetFieldByPath(desired, fieldPath)
	if err != nil {
		log.Error(err, "error fetching desired field value", "field", fieldPath)
		return change
	}
	change.DesiredValue = desiredValue

	if foundActual != foundDesired {
		return change
	}
	if !foundActual && !foundDesired {
		// Both unset
		return change
	}
	if actualValue.Equal(desiredValue) {
		// Note: returning nil to indicate no change
		return nil
	}
	return change
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

func valToAny(v protoreflect.Value) any {
	if !v.IsValid() {
		return nil
	}
	return v.Interface()
}

// MergeUnsetFields recursively copies non-nil values from actual into desired for any pointer/slice fields where desired is nil.
func MergeUnsetFields(desired, actual reflect.Value) {
	if !desired.IsValid() || !actual.IsValid() {
		return
	}

	switch desired.Kind() {
	case reflect.Ptr:
		if desired.IsNil() && !actual.IsNil() {
			desired.Set(actual)
		} else if !desired.IsNil() && !actual.IsNil() {
			MergeUnsetFields(desired.Elem(), actual.Elem())
		}

	case reflect.Struct:
		for i := 0; i < desired.NumField(); i++ {
			field := desired.Type().Field(i)
			if field.IsExported() {
				MergeUnsetFields(desired.Field(i), actual.Field(i))
			}
		}

	case reflect.Slice, reflect.Map:
		if desired.IsNil() && !actual.IsNil() {
			desired.Set(actual)
		}
	}
}

var deterministicJSON = protojson.MarshalOptions{}

// SortRepeatedFields recursively traverses populated list fields in a proto message and sorts them deterministically.
func SortRepeatedFields(msg protoreflect.Message) {
	if msg == nil || !msg.IsValid() {
		return
	}
	msg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		if fd.IsList() {
			sortProtoReflectList(val.List())
		} else if fd.IsMap() {
			val.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				if msg, ok := v.Interface().(protoreflect.Message); ok {
					SortRepeatedFields(msg)
				}
				return true
			})
		} else if fd.Kind() == protoreflect.MessageKind {
			SortRepeatedFields(val.Message())
		}
		return true
	})
}

func sortProtoReflectList(list protoreflect.List) {
	n := list.Len()
	if n <= 1 {
		return
	}

	type itemKey struct {
		val protoreflect.Value
		key string
	}

	items := make([]itemKey, n)
	for i := 0; i < n; i++ {
		val := list.Get(i)
		var key string
		if msgVal, ok := val.Interface().(protoreflect.Message); ok {
			SortRepeatedFields(msgVal)
			if pb, ok := msgVal.Interface().(proto.Message); ok {
				bytes, err := deterministicJSON.Marshal(pb)
				if err == nil {
					key = string(bytes)
				}
			}
			if key == "" {
				key = fmt.Sprintf("%+v", msgVal.Interface())
			}
		} else {
			key = fmt.Sprintf("%+v", val.Interface())
		}
		items[i] = itemKey{val: val, key: key}
	}

	slices.SortFunc(items, func(a, b itemKey) int {
		return strings.Compare(a.key, b.key)
	})

	for i := 0; i < n; i++ {
		list.Set(i, items[i].val)
	}
}

// CompareBrownfieldSpec executes the 5-step brownfield spec comparison pipeline.
func CompareBrownfieldSpec[SpecType any, ProtoT interface {
	proto.Message
}](
	ctx context.Context,
	desiredKRM *SpecType,
	actualProto ProtoT,
	specFromProto func(mapCtx *direct.MapContext, in ProtoT) *SpecType,
	specToProto func(mapCtx *direct.MapContext, in *SpecType) ProtoT,
	normalize func(ctx context.Context, pb ProtoT) error,
) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}

	// Step 1: actualProto -> actualKRM (Spec_FromProto)
	actualKRM := specFromProto(mapCtx, actualProto)
	if err := mapCtx.Err(); err != nil {
		return nil, nil, err
	}

	// Step 2: Merge Unspecified Fields (MergeUnsetFields)
	// Deep copy desiredKRM via JSON to avoid mutating caller state, then merge unspecified fields from actualKRM.
	var clonedDesiredKRM *SpecType
	if desiredKRM != nil {
		bytes, err := json.Marshal(desiredKRM)
		if err != nil {
			return nil, nil, fmt.Errorf("cloning desiredKRM: %w", err)
		}
		clonedDesiredKRM = new(SpecType)
		if err := json.Unmarshal(bytes, clonedDesiredKRM); err != nil {
			return nil, nil, fmt.Errorf("cloning desiredKRM: %w", err)
		}
	}
	MergeUnsetFields(reflect.ValueOf(clonedDesiredKRM), reflect.ValueOf(actualKRM))

	// Step 3: Convert KRM Specs to Protos (Spec_ToProto)
	desiredProtoMasked := specToProto(mapCtx, clonedDesiredKRM)
	if err := mapCtx.Err(); err != nil {
		return nil, nil, err
	}

	actualProtoMasked := specToProto(mapCtx, actualKRM)
	if err := mapCtx.Err(); err != nil {
		return nil, nil, err
	}

	// Step 4: Normalization & Slice Sorting (normalize & SortRepeatedFields)
	if normalize != nil {
		if err := normalize(ctx, desiredProtoMasked); err != nil {
			return nil, nil, err
		}
		if err := normalize(ctx, actualProtoMasked); err != nil {
			return nil, nil, err
		}
	}
	SortRepeatedFields(proto.Message(desiredProtoMasked).ProtoReflect())
	SortRepeatedFields(proto.Message(actualProtoMasked).ProtoReflect())

	// Step 5: Top-Level Diff & FieldMask Generation (DiffForTopLevelFields)
	return DiffForTopLevelFields(ctx, proto.Message(desiredProtoMasked).ProtoReflect(), proto.Message(actualProtoMasked).ProtoReflect())
}
