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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
)

func normalizeKRMObject(u *unstructured.Unstructured, project testgcp.GCPProject, uniqueID string) error {
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
	visitor.replacePaths[".status.uid"] = "12345678"
	visitor.replacePaths[".status.creationTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.observedState.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.observedState.updateTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.updateTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.lastModifiedTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.etag"] = "abcdef123456"
	visitor.replacePaths[".status.observedState.etag"] = "abcdef123456"
	visitor.replacePaths[".status.observedState.creationTimestamp"] = "1970-01-01T00:00:00Z"

	// Specific to Sql
	visitor.replacePaths[".items[].etag"] = "abcdef0123A="

	// Specific to AlloyDB
	visitor.replacePaths[".status.continuousBackupInfo[].enabledTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.ipAddress"] = "10.1.2.3"

	// Specific to BigQuery
	visitor.replacePaths[".spec.access[].userByEmail"] = "user@google.com"

	// Specific to postgresinstance
	visitor.replacePaths[".status.serverCaCert.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.serverCaCert.expirationTime"] = "1970-01-01T00:00:00Z"

	// Specific to VertexAI
	visitor.replacePaths[".status.blobStoragePathPrefix"] = "cloud-ai-platform-00000000-1111-2222-3333-444444444444"

	// Specific to Monitoring
	visitor.replacePaths[".status.creationRecord[].mutateTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.creationRecord[].mutatedBy"] = "user@google.com"
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if path == ".spec.conditions[].name" {
			tokens := strings.Split(s, "/")
			if len(tokens) == 6 && tokens[4] == "conditions" {
				tokens[5] = "${conditionId}"
			}
			s = strings.Join(tokens, "/")
		}
		return s
	})

	// Specific to GCS
	visitor.replacePaths[".softDeletePolicy.effectiveTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".timeCreated"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".updated"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".acl[].etag"] = "abcdef0123A"
	visitor.replacePaths[".defaultObjectAcl[].etag"] = "abcdef0123A="
	visitor.replacePaths[".spec.softDeletePolicy.effectiveTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.observedState.softDeletePolicy.effectiveTime"] = "1970-01-01T00:00:00Z"

	// Specific to Compute
	visitor.replacePaths[".status.observedState.certificateID"] = 1111111111111111
	visitor.replacePaths[".status.instanceId"] = "1111111111111111"
	visitor.replacePaths[".status.gatewayId"] = 1111111111111111
	visitor.replacePaths[".status.proxyId"] = 1111111111111111
	visitor.replacePaths[".status.mapId"] = 1111111111111111

	// Specific to MonitoringDashboard
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.HasSuffix(path, ".alertChart.alertPolicyRef.external") {
			tokens := strings.Split(s, "/")
			if len(tokens) >= 2 {
				switch tokens[len(tokens)-2] {
				case "alertPolicies":
					s = strings.ReplaceAll(s, tokens[len(tokens)-1], "${alertPolicyID}")
				}
			}
		}
		return s
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.HasSuffix(path, ".policyRefs[].external") {
			tokens := strings.Split(s, "/")
			if len(tokens) >= 2 {
				switch tokens[len(tokens)-2] {
				case "alertPolicies":
					s = strings.ReplaceAll(s, tokens[len(tokens)-1], "${alertPolicyID}")
				}
			}
		}
		return s
	})

	visitor.sortSlices = sets.New[string]()
	// TODO: This should not be needed, we want to avoid churning the kube objects
	visitor.sortSlices.Insert(".spec.access")
	visitor.sortSlices.Insert(".spec.nodeConfig.oauthScopes")

	if u.GetKind() == "Project" {
		// For some tests that talk to the Mock Resource Manager, the Project object's ProjectID and ProjectNumber are dynamcially generated.
		// We do not want to overrride this with the default mocked Project "mock-project".
		visitor.replacePaths[".status.number"] = "${projectNumber}"
	}

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

	// Try to extract resource IDs from links and replace them
	{
		name, _, _ := unstructured.NestedString(u.Object, "status", "observedState", "name")
		if name == "" {
			name, _, _ = unstructured.NestedString(u.Object, "status", "name")
		}
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		tokens := strings.Split(name, "/")
		if len(tokens) == 1 {
			switch u.GetKind() {
			case "TagsTagKey", "TagsTagValue":
				// TODO: The mock TagKey server returns the correct format `tagKeys/{number}`, but the golden object `status.name`
				// only has {number}. Need to triage the tf/dcl controller.
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, name, "${uniqueId}")
				})
			}
		}
		if len(tokens) > 2 {
			typeName := tokens[len(tokens)-2]
			id := tokens[len(tokens)-1]
			if typeName == "datasets" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${datasetId}")
				})
			}
			if typeName == "alertPolicies" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${alertPolicyId}")
				})
			}
			if typeName == "tensorboards" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${tensorboardId}")
				})
			}
			if typeName == "notificationChannels" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${notificationChannelID}")
				})
			}
		}

		switch u.GroupVersionKind() {
		case schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringUptimeCheckConfig"}:
			visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
				return strings.ReplaceAll(s, resourceID, "${uptimeCheckConfigId}")
			})
		}
	}

	return visitor.VisitUnstructued(u)
}

func setStringAtPath(m map[string]any, atPath string, newValue string) error {
	visitor := objectWalker{}

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if path == atPath {
			return newValue
		}
		return s
	})

	if err := visitor.visitMap(m, ""); err != nil {
		return err
	}
	return nil
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
	// Remove headers in request.
	events.RemoveHTTPRequestHeader("X-Goog-Api-Client")
	// Remove header in response.
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

	normalizeHTTPResponses(t, events)

	// Normalize using the KRM normalization function
	events.PrettifyJSON(func(obj map[string]any) {
		u := &unstructured.Unstructured{}
		u.Object = obj
		if err := normalizeKRMObject(u, project, uniqueID); err != nil {
			t.Fatalf("error from normalizeObject: %v", err)
		}
	})
}

func normalizeHTTPResponses(t *testing.T, events test.LogEntries) {
	visitor := objectWalker{}

	visitor.removePaths = sets.New[string]()
	visitor.replacePaths = make(map[string]any)

	// If we get detailed info, don't record it - it's not part of the API contract
	visitor.removePaths.Insert(".error.errors[].debugInfo")

	// Common variables
	visitor.replacePaths[".etag"] = "abcdef0123A="
	visitor.replacePaths[".response.etag"] = "abcdef0123A="
	visitor.replacePaths[".serviceAccount.etag"] = "abcdef0123A="

	// Compute operations
	visitor.replacePaths[".fingerprint"] = "abcdef0123A="
	visitor.replacePaths[".startTime"] = "2024-04-01T12:34:56.123456Z"

	events.PrettifyJSON(func(obj map[string]any) {
		if err := visitor.visitMap(obj, ""); err != nil {
			t.Fatalf("error normalizing response: %v", err)
		}
	})
}
