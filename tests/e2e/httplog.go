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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
)

type Normalizer struct {
	uniqueID string
	project  testgcp.GCPProject

	pathIDs      map[string]string
	operationIDs map[string]bool
}

func NewNormalizer(uniqueID string, project testgcp.GCPProject) *Normalizer {
	return &Normalizer{
		uniqueID: uniqueID,
		project:  project,

		operationIDs: map[string]bool{},
		pathIDs:      map[string]string{},
	}
}

func (x *Normalizer) Render(events test.LogEntries) string {

	// Replace any dynamic IDs that appear in URLs
	for _, event := range events {
		url := event.Request.URL
		for k, v := range x.pathIDs {
			url = strings.ReplaceAll(url, "/"+k, "/"+v)
		}
		event.Request.URL = url
	}

	// Remove operation polling requests (ones where the operation is not ready)
	events = events.KeepIf(func(e *test.LogEntry) bool {
		if !strings.Contains(e.Request.URL, "/operations/${operationID}") {
			return true
		}
		responseBody := e.Response.ParseBody()
		if responseBody == nil {
			return true
		}
		if done, _, _ := unstructured.NestedBool(responseBody, "done"); done {
			return true
		}
		// remove if not done - and done can be omitted when false
		return false
	})

	jsonMutators := []test.JSONMutator{}
	addReplacement := func(path string, newValue string) {
		tokens := strings.Split(path, ".")
		jsonMutators = append(jsonMutators, func(obj map[string]any) {
			_, found, _ := unstructured.NestedString(obj, tokens...)
			if found {
				if err := unstructured.SetNestedField(obj, newValue, tokens...); err != nil {
					klog.Fatalf("error setting field: %v", err)
				}
			}
		})
	}

	addReplacement("id", "000000000000000000000")
	addReplacement("uniqueId", "111111111111111111111")
	addReplacement("oauth2ClientId", "888888888888888888888")

	addReplacement("etag", "abcdef0123A=")
	addReplacement("serviceAccount.etag", "abcdef0123A=")
	addReplacement("response.etag", "abcdef0123A=")

	addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("creationTimestamp", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.createTime", "2024-04-01T12:34:56.123456Z")

	addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.updateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to spanner
	addReplacement("metadata.startTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.updateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to vertexai
	addReplacement("blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
	addReplacement("response.blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")

	// Replace any empty values in LROs; this is surprisingly difficult to fix in mockgcp
	//
	//     "response": {
	// 	-    "@type": "type.googleapis.com/google.protobuf.Empty"
	// 	+    "@type": "type.googleapis.com/google.protobuf.Empty",
	// 	+    "value": {}
	// 	   }
	jsonMutators = append(jsonMutators, func(obj map[string]any) {
		response := obj["response"]
		if responseMap, ok := response.(map[string]any); ok {
			if responseMap["@type"] == "type.googleapis.com/google.protobuf.Empty" {
				value := responseMap["value"]
				if valueMap, ok := value.(map[string]any); ok && len(valueMap) == 0 {
					delete(responseMap, "value")
				}
			}
		}
	})

	events.PrettifyJSON(jsonMutators...)

	// Remove headers that just aren't very relevant to testing
	// Remove headers in request.
	events.RemoveHTTPRequestHeader("X-Goog-Api-Client")
	// Remove headers in response.
	events.RemoveHTTPResponseHeader("Date")
	events.RemoveHTTPResponseHeader("Alt-Svc")
	events.RemoveHTTPResponseHeader("Server-Timing")

	got := events.FormatHTTP()
	normalizers := []func(string) string{}
	normalizers = append(normalizers, ReplaceString(x.uniqueID, "${uniqueId}"))
	normalizers = append(normalizers, ReplaceString(x.project.ProjectID, "${projectId}"))
	normalizers = append(normalizers, ReplaceString(fmt.Sprintf("%d", x.project.ProjectNumber), "${projectNumber}"))
	for k, v := range x.pathIDs {
		normalizers = append(normalizers, ReplaceString(k, v))
	}
	for k := range x.operationIDs {
		normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
	}

	for _, normalizer := range normalizers {
		got = normalizer(got)
	}
	return got
}

func (x *Normalizer) Preprocess(events []*test.LogEntry) {

	// Find "easy" operations and resources by looking for fully-qualified methods
	for _, event := range events {
		u := event.Request.URL
		if index := strings.Index(u, "?"); index != -1 {
			u = u[:index]
		}
		tokens := strings.Split(u, "/")
		n := len(tokens)
		if n >= 2 {
			kind := tokens[n-2]
			id := tokens[n-1]
			switch kind {
			case "tensorboards":
				x.pathIDs[id] = "${tensorboardID}"
			case "operations":
				x.operationIDs[id] = true
				x.pathIDs[id] = "${operationID}"
			}
		}
	}

	for _, event := range events {
		id := ""
		body := event.Response.ParseBody()
		val, ok := body["name"]
		if ok {
			s := val.(string)
			// operation name format: operations/{operationId}
			if strings.HasPrefix(s, "operations/") {
				id = strings.TrimPrefix(s, "operations/")
			}
			// operation name format: {prefix}/operations/{operationId}
			if ix := strings.Index(s, "/operations/"); ix != -1 {
				id = strings.TrimPrefix(s[ix:], "/operations/")
			}
			// operation name format: operation-{operationId}
			if strings.HasPrefix(s, "operation") {
				id = s
			}
		}
		if id != "" {
			x.operationIDs[id] = true
		}
	}

	for _, event := range events {
		if !strings.Contains(event.Request.URL, "/operations/${operationID}") {
			continue
		}
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		name, _, _ := unstructured.NestedString(responseBody, "response", "name")
		if strings.HasPrefix(name, "tagKeys/") {
			x.pathIDs[name] = "tagKeys/${tagKeyID}"
		}
	}

}

// ReplaceString is a normalization function that replaces a string, useful for e.g. project IDs.
func ReplaceString(from, to string) func(string) string {
	return func(s string) string {
		return strings.ReplaceAll(s, from, to)
	}
}

// IgnoreComments is a normalization function that strips comments.
func IgnoreComments(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			lines[i] = ""
		}
	}
	s = strings.Join(lines, "\n")
	return strings.TrimSpace(s)
}

func IgnoreAnnotations(annotations map[string]struct{}) func(string) string {
	return func(s string) string {
		lines := strings.Split(s, "\n")
		sb := strings.Builder{}
		for _, line := range lines {
			ignore := false
			// todo(acpana): only operate on annotations and actually look up in map
			for anon := range annotations {
				if strings.Contains(line, anon) {
					ignore = true
					break
				}
			}

			if !ignore {
				sb.WriteString(line)
				sb.WriteString("\n")
			}
		}

		return sb.String()
	}
}
