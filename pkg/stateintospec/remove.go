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

package stateintospec

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

type Warning struct {
	Message string
}

type Warnings struct {
	Warnings []Warning
}

func (w *Warnings) AddWarningf(msg string, args ...any) {
	w.Warnings = append(w.Warnings, Warning{
		Message: fmt.Sprintf(msg, args...),
	})
}

func RemoveStateIntoSpecFields(ctx context.Context, u *unstructured.Unstructured, preserveFieldFunc func(fieldPath fieldpath.Path) bool) (*Warnings, error) {
	log := klog.FromContext(ctx)

	warnings := &Warnings{}

	managedFields, err := parseManagedFields(u.GetManagedFields())
	if err != nil {
		return warnings, fmt.Errorf("parsing managed fields: %w", err)
	}

	var errs []error

	for manager, fields := range managedFields {
		if manager != k8s.ControllerManagedFieldManager {
			continue
		}

		fields.Iterate(func(fieldPath fieldpath.Path) {
			switch fieldPath[0].String() {
			case ".spec":
				pathName := fieldPath.String()
				if preserveFieldFunc(fieldPath) {
					log.Info("preserving field as requested", "field", pathName)
					return
				}
				log.Info("removing field", "field", pathName)
				errs = append(errs, removeFieldIfNotArray(ctx, u, fieldPath, warnings))
			case ".status", ".metadata":
				// Never part of state-into-spec, ignore
			default:
				errs = append(errs, fmt.Errorf("found unknown field %q in managed fields", fieldPath.String()))
			}
		})
	}

	return warnings, errors.Join(errs...)
}

func removeFieldIfNotArray(ctx context.Context, u *unstructured.Unstructured, fieldPath fieldpath.Path, warnings *Warnings) error {
	// log := klog.FromContext(ctx)

	pos := u.Object
	n := len(fieldPath)
	for i := 0; i < n-1; i++ {
		element := fieldPath[i]

		if element.FieldName != nil {
			v, found := pos[*element.FieldName]
			if !found {
				return nil
			}
			m, ok := v.(map[string]any)
			if ok {
				pos = m
				continue
			}
			return fmt.Errorf("unexpected type for %q: got %T, expected map", fieldPath, v)
		}
		return fmt.Errorf("removal of fieldPath %v not implemented", fieldPath)
	}

	last := fieldPath[n-1]
	if last.FieldName != nil {
		v, found := pos[*last.FieldName]
		if !found {
			// Already removed
			return nil
		}
		switch v := v.(type) {
		case map[string]any:
			warnings.AddWarningf("deleting map field %q (all the subfields are considered managed)", fieldPath)
			delete(pos, *last.FieldName)
			return nil
		case []any:
			warnings.AddWarningf("skipping field removal of array field %q (may be an indication of undetermined ownership)", fieldPath)
			return nil
		case string, int, int32, int64, float32, float64, bool:
			delete(pos, *last.FieldName)
			return nil
		default:
			return fmt.Errorf("unhandled type for field %q: got %T", fieldPath, v)
		}

	}

	return fmt.Errorf("removal of fieldPath %v not implemented", fieldPath)
}

// parseManagedFields takes the given managed field entries and constructs a
// set of all the k8s-managed fields from the spec, grouping by manager name.
func parseManagedFields(managedFields []metav1.ManagedFieldsEntry) (map[string]*fieldpath.Set, error) {
	res := make(map[string]*fieldpath.Set)
	for _, managedFieldEntry := range managedFields {
		if managedFieldEntry.FieldsType != k8s.ManagedFieldsTypeFieldsV1 {
			return nil, fmt.Errorf(
				"expected managed field entry for manager '%v' and operation '%v' of type '%v', got type '%v'",
				managedFieldEntry.Manager, managedFieldEntry.Operation, k8s.ManagedFieldsTypeFieldsV1,
				managedFieldEntry.FieldsType)
		}
		fieldsV1 := managedFieldEntry.FieldsV1
		if fieldsV1 == nil {
			return nil, fmt.Errorf("managed field entry for manager '%v' and operation '%v' has empty fieldsV1",
				managedFieldEntry.Manager, managedFieldEntry.Operation)
		}
		entrySet := fieldpath.NewSet()
		if err := entrySet.FromJSON(bytes.NewReader(fieldsV1.Raw)); err != nil {
			return nil, fmt.Errorf("error marshaling managed fields for manager '%v' and operation '%v' from JSON: %w",
				managedFieldEntry.Manager, managedFieldEntry.Operation, err)
		}

		fields := res[managedFieldEntry.Manager]
		if fields == nil {
			fields = fieldpath.NewSet()
		}
		fields = fields.Union(entrySet)
		res[managedFieldEntry.Manager] = fields
	}
	return res, nil
}
