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

package e2e

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
)

func normalizeObject(u *unstructured.Unstructured, project testgcp.GCPProject, uniqueID string) error {
	annotations := u.GetAnnotations()
	if annotations["cnrm.cloud.google.com/observed-secret-versions"] != "" {
		// Includes resource versions, very volatile
		annotations["cnrm.cloud.google.com/observed-secret-versions"] = "(removed)"
	}
	u.SetAnnotations(annotations)

	visitor := objectWalker{}

	visitor.removePaths = sets.New[string]()
	visitor.removePaths.Insert(".metadata.creationTimestamp")
	visitor.removePaths.Insert(".metadata.managedFields")
	visitor.removePaths.Insert(".metadata.resourceVersion")
	visitor.removePaths.Insert(".metadata.uid")

	visitor.replacePaths = map[string]any{}
	visitor.replacePaths[".metadata.deletionTimestamp"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.conditions[].lastTransitionTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.uniqueId"] = "12345678"
	visitor.replacePaths[".status.creationTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.updateTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.lastModifiedTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.etag"] = "abcdef123456"

	visitor.sortSlices = sets.New[string]()
	// TODO: This should not be needed, we want to avoid churning the kube objects
	visitor.sortSlices.Insert(".spec.access")

	visitor.stringTransforms = append(visitor.stringTransforms, func(s string) string {
		return strings.ReplaceAll(s, project.ProjectID, "${projectId}")
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(s string) string {
		return strings.ReplaceAll(s, fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}")
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(s string) string {
		return strings.ReplaceAll(s, uniqueID, "${uniqueId}")
	})

	// TODO: Only for some objects?
	visitor.stringTransforms = append(visitor.stringTransforms, func(s string) string {
		r := regexp.MustCompile(regexp.QuoteMeta(`deleted:serviceAccount:gsa-${uniqueId}@${projectId}.iam.gserviceaccount.com?uid=`) + `.*`)
		return r.ReplaceAllLiteralString(s, "deleted:serviceAccount:gsa-${uniqueId}@${projectId}.iam.gserviceaccount.com?uid=12345678")
	})

	return visitor.VisitUnstructued(u)
}

type objectWalker struct {
	removePaths      sets.Set[string]
	sortSlices       sets.Set[string]
	replacePaths     map[string]any
	stringTransforms []func(string) string
}

func (o *objectWalker) visitAny(v any, path string) (any, error) {
	switch v := v.(type) {
	case map[string]any:
		return o.visitMap(v, path)
	case []any:
		return o.visitSlice(v, path)
	case int64, float64, bool:
		return o.visitPrimitive(v, path)
	case string:
		return o.visitString(v, path)
	default:
		return nil, fmt.Errorf("unhandled type %T", v)
	}
}

func (o *objectWalker) visitMap(m map[string]any, path string) (map[string]any, error) {
	for k, v := range m {
		childPath := path + "." + k
		if o.removePaths.Has(childPath) {
			delete(m, k)
			continue // nothing left to process
		}

		if v2, found := o.replacePaths[childPath]; found {
			m[k] = v2
			continue // replacement value is assumed to be normalized
		}

		v2, err := o.visitAny(v, childPath)
		if err != nil {
			return nil, err
		}
		m[k] = v2
		v = v2

		// Note: do sorting "last" so we sort normalized values
		if o.sortSlices.Has(childPath) {
			s, ok := v.([]any)
			if !ok {
				return nil, fmt.Errorf("expected slice at %q, got %T", childPath, v)
			}
			if err := sortSlice(s); err != nil {
				return nil, err
			}
		}
	}

	return m, nil
}

func sortSlice(s []any) error {
	var jsons []string
	for i := range s {
		j, err := json.Marshal(s[i])
		if err != nil {
			return fmt.Errorf("error converting to json: %w", err)
		}
		jsons = append(jsons, string(j))
	}

	sort.Slice(s, func(i, j int) bool {
		return jsons[i] < jsons[j]
	})

	return nil
}

func (o *objectWalker) visitSlice(s []any, path string) (any, error) {
	for i, v := range s {
		v2, err := o.visitAny(v, path+"[]")
		if err != nil {
			return nil, err
		}
		s[i] = v2
	}

	return s, nil
}

func (o *objectWalker) visitPrimitive(v any, _ string) (any, error) {
	return v, nil
}

func (o *objectWalker) visitString(v string, _ string) (string, error) {
	for _, fn := range o.stringTransforms {
		v = fn(v)
	}
	return v, nil
}

func (o *objectWalker) VisitUnstructued(v *unstructured.Unstructured) error {
	v2, err := o.visitMap(v.Object, "")
	if err != nil {
		return err
	}
	v.Object = v2
	return nil
}
