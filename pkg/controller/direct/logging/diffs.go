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

package logging

import (
	"errors"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// objectDiffWalker walks two objects, recording any diffs (and errors encountered)
type objectDiffWalker struct {
	errs  []error
	diffs []string
}

func (o *objectDiffWalker) foundDifference(lV, rV any, path string) {
	o.diffs = append(o.diffs, path)
}

func (o *objectDiffWalker) visitAny(lV, rV any, path string) {
	if lV == nil {
		if rV != nil {
			o.foundDifference(lV, rV, path)
		}
		return
	}
	if rV == nil {
		if lV != nil {
			o.foundDifference(lV, rV, path)
		}
		return
	}
	switch lV := lV.(type) {
	case map[string]any:
		m, ok := rV.(map[string]any)
		if !ok {
			o.foundDifference(lV, rV, path)
		} else {
			o.visitMap(lV, m, path)
		}

	case []any:
		m, ok := rV.([]any)
		if !ok {
			o.foundDifference(lV, rV, path)
		} else {
			o.visitSlice(lV, m, path)
		}
	case int64, float64, bool, string:
		if !reflect.DeepEqual(lV, rV) {
			o.foundDifference(lV, rV, path)
		}
	default:
		o.errs = append(o.errs, fmt.Errorf("unhandled type at path %q: %T", path, lV))
	}
}

func (o *objectDiffWalker) visitMap(l, r map[string]any, path string) {
	for k, lV := range l {
		childPath := path + "." + k

		rV, found := r[k]
		if !found {
			o.foundDifference(lV, rV, childPath)
			continue
		}

		o.visitAny(lV, rV, childPath)
	}

	for k, rV := range r {
		childPath := path + "." + k

		lV, found := l[k]
		if !found {
			o.foundDifference(lV, rV, childPath)
			continue
		}
	}
}

func (o *objectDiffWalker) visitSlice(l, r []any, path string) {
	for i, lV := range l {
		if i >= len(r) {
			o.foundDifference(lV, nil, path)
			continue
		}
		rV := r[i]
		o.visitAny(lV, rV, path+"[]")
	}

	for i := len(l); i < len(r); i++ {
		rV := r[i]
		o.foundDifference(nil, rV, path)
	}
}

func (o *objectDiffWalker) visitUnstructured(l, r *unstructured.Unstructured) {
	o.visitMap(l.Object, r.Object, "")
}

// ListFieldDiffsForUnstructured returns a list of field paths where the values differ, for two unstructured objects.
func ListFieldDiffsForUnstructured(l, r *unstructured.Unstructured) ([]string, error) {
	o := &objectDiffWalker{}
	o.visitUnstructured(l, r)
	return o.diffs, errors.Join(o.errs...)
}

// ListFieldDiffs returns a list of field paths where the values differ, for two arbitrary objects.
// If the objects are unstructured.Unstructured, use ListFieldDiffsForUnstructured for efficiency.
func ListFieldDiffs(l, r any) ([]string, error) {
	lObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&l)
	if err != nil {
		return nil, fmt.Errorf("converting %T to unstructured: %w", l, err)
	}
	rObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&r)
	if err != nil {
		return nil, fmt.Errorf("converting %T to unstructured: %w", r, err)
	}
	o := &objectDiffWalker{}
	o.visitMap(lObj, rObj, "")
	return o.diffs, errors.Join(o.errs...)
}
