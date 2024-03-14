package diffs

import (
	"fmt"
	"io"
	"reflect"
	"sort"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog"
)

type ObjectDiff struct {
	OldObject *unstructured.Unstructured `json:"oldObject"`
	NewObject *unstructured.Unstructured `json:"newObject"`

	fieldDiffs []fieldDiff
}

type fieldDiff struct {
	Path     *FieldPath `json:"path"`
	OldValue any        `json:"oldValue,omitempty"`
	NewValue any        `json:"newValue,omitempty"`
}

func (d *ObjectDiff) walkMap(oldObj, newObj map[string]any, fieldPath *FieldPath) {
	for k, oldValue := range oldObj {
		newValue := newObj[k]
		d.walkAny(oldValue, newValue, fieldPath.With(k))
	}
	for k, newValue := range newObj {
		oldValue, found := oldObj[k]
		if found {
			// Dealt with in previous loop
			continue
		}
		d.walkAny(oldValue, newValue, fieldPath.With(k))
	}
}

func (d *ObjectDiff) walkSlice(oldSlice, newSlice []any, fieldPath *FieldPath) {
	minLen := min(len(oldSlice), len(newSlice))
	for i := 0; i < minLen; i++ {
		oldValue := newSlice[i]
		newValue := newSlice[i]
		d.walkAny(oldValue, newValue, fieldPath.With(fmt.Sprintf("[%d]", i)))
	}
	for i := minLen; i < len(newSlice); i++ {
		newValue := newSlice[i]
		d.walkAny(nil, newValue, fieldPath.With(fmt.Sprintf("[%d]", i)))
	}
	for i := minLen; i < len(oldSlice); i++ {
		oldValue := newSlice[i]
		d.walkAny(oldValue, nil, fieldPath.With(fmt.Sprintf("[%d]", i)))
	}
}

func (d *ObjectDiff) walkAny(oldVal, newVal any, fieldPath *FieldPath) {
	addDiff := true

	switch oldVal := oldVal.(type) {
	case map[string]any:
		newMap, ok := newVal.(map[string]any)
		if ok {
			d.walkMap(oldVal, newMap, fieldPath)
			addDiff = false
		}

	case []any:
		newSlice, ok := newVal.([]any)
		if ok {
			d.walkSlice(oldVal, newSlice, fieldPath)
			addDiff = false
		}

	case string, int64, bool:
		if reflect.DeepEqual(oldVal, newVal) {
			addDiff = false
		}

	default:
		klog.Warningf("type %T not handled", oldVal)
	}

	if addDiff {
		d.fieldDiffs = append(d.fieldDiffs, fieldDiff{Path: fieldPath, OldValue: oldVal, NewValue: newVal})
	}
}

func BuildObjectDiff(oldObj, newObj *unstructured.Unstructured) (*ObjectDiff, error) {
	d := &ObjectDiff{
		OldObject: oldObj,
		NewObject: newObj,
	}
	d.walkMap(oldObj.Object, newObj.Object, nil)
	return d, nil
}

type prettyPrintFieldPath struct {
	fieldDiff
	keyPath []string
}

func (d *ObjectDiff) PrettyPrintTo(out io.Writer) {
	var diffs []prettyPrintFieldPath
	for _, diff := range d.fieldDiffs {
		diffs = append(diffs, prettyPrintFieldPath{
			fieldDiff: diff,
			keyPath:   diff.Path.asSlice(),
		})
	}

	sort.Slice(diffs, func(i, j int) bool {
		return compareStringSlices(diffs[i].keyPath, diffs[j].keyPath) < 0
	})

	// for _, diff := range diffs {
	// 	fmt.Fprintf(out, "%s: %v -> %v\n", strings.Join(diff.keyPath, "."), diff.OldValue, diff.NewValue)
	// }

	var previousKeyPath []string
	for _, diff := range diffs {
		indent := ""
		n := min(len(previousKeyPath), len(diff.keyPath))
		i := 0
		for i < n {
			if previousKeyPath[i] != diff.keyPath[i] {
				break
			}

			indent += "  "
			i++
		}
		for ; i < len(diff.keyPath)-1; i++ {
			fmt.Fprintf(out, "%s%s:\n", indent, diff.keyPath[i])
			indent += "  "
		}
		fmt.Fprintf(out, "%s%s: %v -> %v\n", indent, diff.keyPath[len(diff.keyPath)-1], diff.OldValue, diff.NewValue)
		previousKeyPath = diff.keyPath
	}
}
