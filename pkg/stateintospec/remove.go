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

func RemoveStateIntoSpecFields(ctx context.Context, u *unstructured.Unstructured, preserveFieldFunc func(fieldPath fieldpath.Path) bool) error {
	log := klog.FromContext(ctx)

	managedFields, err := parseManagedFields(u.GetManagedFields())
	if err != nil {
		return fmt.Errorf("parsing managed fields: %w", err)
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
				errs = append(errs, removeFieldIfLeaf(ctx, u, fieldPath))
			case ".status", ".metadata":
				// Never part of state-into-spec, ignore
			default:
				errs = append(errs, fmt.Errorf("found unknown field %q in managed fields", fieldPath.String()))
			}
		})
	}

	return errors.Join(errs...)
}

func removeFieldIfLeaf(ctx context.Context, u *unstructured.Unstructured, fieldPath fieldpath.Path) error {
	log := klog.FromContext(ctx)

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
			log.Info("skipping field removal of map field (may be an indication of undetermined ownership)", "path", fieldPath)
			return nil
		case []any:
			log.Info("skipping field removal of array field (may be an indication of undetermined ownership)", "path", fieldPath)
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
		if managedFieldEntry.FieldsV1 == nil {
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
