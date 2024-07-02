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

package gkehub

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type diff struct {
	path       string
	desiredVal any
}

// objectDiffWalker walks through two objects and record any fields diffs and errors.
type objectDiffWalker struct {
	errs  []error
	diffs []diff
}

// records the left value as the desired value.
func (o *objectDiffWalker) recordDiff(lV any, path string) {
	o.diffs = append(o.diffs, diff{path: path, desiredVal: lV})
}

// To follow the "unmanage and unset" behaviour, the visitAny function does the following:
// lV is the desired state and rV is the actual state.
// * Fields both present in lV and rV, but have different values, record the diff.
// * Fields only present in lV but not in rV, record the diff.
// * Fields not present in lV but in rV, do not record the diff.
func (o *objectDiffWalker) visitAny(lV, rV any, path string) {
	if lV == nil {
		return
	}
	if rV == nil {
		o.recordDiff(lV, path)
	}
	switch lV := lV.(type) {
	case map[string]any:
		m, ok := rV.(map[string]any)
		if !ok {
			o.errs = append(o.errs, fmt.Errorf("unhandled type at path %q: %T", path, rV))
		} else {
			o.visitMap(lV, m, path)
		}
	case []any:
		m, ok := rV.([]any)
		// For the gkehub API, slices only contain primitive types.
		if !ok {
			o.errs = append(o.errs, fmt.Errorf("unhandled type at path %q: %T", path, rV))
		} else {
			if !reflect.DeepEqual(lV, m) {
				o.recordDiff(lV, path)
			}
		}
	case int64, float64, bool, string:
		if !reflect.DeepEqual(lV, rV) {
			o.recordDiff(lV, path)
		}
	default:
		o.errs = append(o.errs, fmt.Errorf("unhandled type at path %q: %T", path, lV))
	}
}

func (o *objectDiffWalker) visitMap(l, r map[string]any, path string) {
	var childPath string
	for k, lV := range l {
		if len(path) == 0 {
			childPath = k
		} else {
			childPath = path + "." + k
		}
		rV, found := r[k]
		if !found {
			o.recordDiff(lV, childPath)
			continue
		}
		o.visitAny(lV, rV, childPath)
	}
}

// ListFieldDiffs assuems the l and r has similar structure after converting to unstructed. otherwise, throw an error.
func ListFieldDiffs(l any, r any) ([]diff, error) {
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

func SetObjWithDiffs(obj *map[string]interface{}, diffs []diff) error {
	for _, diff := range diffs {
		switch val := diff.desiredVal.(type) {
		case map[string]interface{}:
			if err := unstructured.SetNestedField(*obj, val, strings.Split(diff.path, ".")...); err != nil {
				return err
			}
		case []interface{}:
			if err := unstructured.SetNestedSlice(*obj, val, strings.Split(diff.path, ".")...); err != nil {
				return err
			}
		case int64, float64, bool, string:
			if err := unstructured.SetNestedField(*obj, val, strings.Split(diff.path, ".")...); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unhandled type at path %q: %T", diff.path, val)
		}
	}
	return nil
}

// newFeatureMembershipKRMWithDiffs creates a GKEHubFeatureMembershipSpec struct from a GKEHubFeatureMembershipSpec struct with the diffs set.
func newFeatureMembershipKRMWithDiffs(r *krm.GKEHubFeatureMembershipSpec, diffs []diff) (*krm.GKEHubFeatureMembershipSpec, error) {
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		return nil, fmt.Errorf("converting %T to unstructured: %w", r, err)
	}
	if err := SetObjWithDiffs(&u, diffs); err != nil {
		return nil, err
	}
	obj := &krm.GKEHubFeatureMembershipSpec{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	return obj, nil
}
