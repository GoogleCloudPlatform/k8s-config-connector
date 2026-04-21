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

package tags

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/klog/v2"
)

type FieldChange struct {
	FieldPath    string
	ActualValue  protoreflect.Value
	DesiredValue protoreflect.Value
}

func DiffForTopLevelFields(ctx context.Context, desired protoreflect.Message, actual protoreflect.Message) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	diff := &structuredreporting.Diff{}

	fields := actual.Type().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		path := string(fields.Get(i).Name())
		fieldDiff := fieldHasChanged(ctx, path, desired, actual)
		if fieldDiff == nil {
			continue
		}
		diff.AddProtoField(fieldDiff.FieldPath, fields.Get(i), fieldDiff.ActualValue, fieldDiff.DesiredValue)
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
