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
	"net/http"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
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
	visitor.replacePaths[".status.creationTimestamp"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.conditions[].lastTransitionTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.uniqueId"] = "12345678"
	visitor.replacePaths[".status.creationTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.updateTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.lastModifiedTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.etag"] = "abcdef123456"

	// Specific to BigQuery
	visitor.replacePaths[".spec.access[].userByEmail"] = "user@google.com"

	// Specific to GCS
	visitor.replacePaths[".softDeletePolicy.effectiveTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".timeCreated"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".updated"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".acl[].etag"] = "abcdef0123A"
	visitor.replacePaths[".defaultObjectAcl[].etag"] = "abcdef0123A="

	visitor.sortSlices = sets.New[string]()
	// TODO: This should not be needed, we want to avoid churning the kube objects
	visitor.sortSlices.Insert(".spec.access")

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		return strings.ReplaceAll(s, project.ProjectID, "${projectId}")
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		return strings.ReplaceAll(s, fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}")
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		return strings.ReplaceAll(s, uniqueID, "${uniqueId}")
	})

	// TODO: Only for some objects?
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		r := regexp.MustCompile(regexp.QuoteMeta(`deleted:serviceAccount:gsa-${uniqueId}@${projectId}.iam.gserviceaccount.com?uid=`) + `.*`)
		return r.ReplaceAllLiteralString(s, "deleted:serviceAccount:gsa-${uniqueId}@${projectId}.iam.gserviceaccount.com?uid=12345678")
	})

	return visitor.VisitUnstructued(u)
}

type objectWalker struct {
	removePaths      sets.Set[string]
	sortSlices       sets.Set[string]
	replacePaths     map[string]any
	stringTransforms []func(path string, value string) string
}

func (o *objectWalker) visitAny(v any, path string) (any, error) {
	if v == nil {
		return v, nil
	}
	switch v := v.(type) {
	case map[string]any:
		if err := o.visitMap(v, path); err != nil {
			return nil, err
		}
		return v, nil
	case []any:
		return o.visitSlice(v, path)
	case int64, float64, bool:
		return o.visitPrimitive(v, path)
	case string:
		return o.visitString(v, path)
	default:
		return nil, fmt.Errorf("unhandled type at path %q: %T", path, v)
	}
}

func (o *objectWalker) visitMap(m map[string]any, path string) error {
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
			return err
		}
		m[k] = v2
		v = v2
	}

	return nil
}

func sortSlice(s []any) error {
	type entry struct {
		o       any
		sortKey string
	}

	var entries []entry
	for i := range s {
		j, err := json.Marshal(s[i])
		if err != nil {
			return fmt.Errorf("error converting to json: %w", err)
		}
		entries = append(entries, entry{o: s[i], sortKey: string(j)})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].sortKey < entries[j].sortKey
	})

	for i := range s {
		s[i] = entries[i].o
	}

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

	// Note: do sorting "last" so we sort normalized values
	if o.sortSlices.Has(path) {
		if err := sortSlice(s); err != nil {
			return s, err
		}
	}

	return s, nil
}

func (o *objectWalker) visitPrimitive(v any, _ string) (any, error) {
	return v, nil
}

func (o *objectWalker) visitString(v string, path string) (string, error) {
	for _, fn := range o.stringTransforms {
		v = fn(path, v)
	}
	return v, nil
}

func (o *objectWalker) VisitUnstructued(v *unstructured.Unstructured) error {
	if err := o.visitMap(v.Object, ""); err != nil {
		return err
	}
	return nil
}

func NormalizeHTTPLog(t *testing.T, events test.LogEntries, project testgcp.GCPProject, uniqueID string) {
	// Remove headers that just aren't very relevant to testing
	events.RemoveHTTPResponseHeader("Date")
	events.RemoveHTTPResponseHeader("Alt-Svc")
	events.RemoveHTTPResponseHeader("Server-Timing")
	events.RemoveHTTPResponseHeader("X-Guploader-Uploadid")
	events.RemoveHTTPResponseHeader("Etag")
	events.RemoveHTTPResponseHeader("Content-Length") // an artifact of encoding

	// Replace any expires headers with (rounded) relative offsets
	for _, event := range events {
		expires := event.Response.Header.Get("Expires")
		if expires == "" {
			continue
		}

		if expires == "Mon, 01 Jan 1990 00:00:00 GMT" {
			// Magic value meaning no-cache; don't change
			continue
		}

		expiresTime, err := time.Parse(http.TimeFormat, expires)
		if err != nil {
			t.Fatalf("parsing Expires header %q: %v", expires, err)
		}
		now := time.Now()
		delta := expiresTime.Sub(now)
		if delta > (55 * time.Minute) {
			delta = delta.Round(time.Hour)
			event.Response.Header.Set("Expires", fmt.Sprintf("{now+%vh}", delta.Hours()))
		} else {
			delta = delta.Round(time.Minute)
			event.Response.Header.Set("Expires", fmt.Sprintf("{now+%vm}", delta.Minutes()))
		}
	}

	events.PrettifyJSON(func(obj map[string]any) {
		u := &unstructured.Unstructured{}
		u.Object = obj
		if err := normalizeObject(u, project, uniqueID); err != nil {
			t.Fatalf("error from normalizeObject: %v", err)
		}
	})

}
