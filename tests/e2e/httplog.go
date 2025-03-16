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
	"github.com/GoogleCloudPlatform/k8s-config-connector/version"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
)

type Normalizer struct {
	uniqueID string
	project  testgcp.GCPProject

	*Replacements
}

func NewNormalizer(uniqueID string, project testgcp.GCPProject) *Normalizer {
	return &Normalizer{
		uniqueID:     uniqueID,
		project:      project,
		Replacements: NewReplacements(),
	}
}

// RemoveExtraEvents removes events that are not as relevant to our golden logs
// In particular, we remove LRO polling operations (and things that look like LROs)
func RemoveExtraEvents(events test.LogEntries) test.LogEntries {
	// Remove operation polling requests (ones where the operation is not ready)
	events = events.KeepIf(func(e *test.LogEntry) bool {
		if !isGetOperation(e) {
			return true
		}
		responseBody := e.Response.ParseBody()
		if responseBody == nil {
			return true
		}
		if done, _, _ := unstructured.NestedBool(responseBody, "done"); done {
			return true
		}
		// compute operations return status==DONE
		if status, _, _ := unstructured.NestedString(responseBody, "status"); status == "DONE" {
			return true
		}
		// remove if not done - and done can be omitted when false
		return false
	})

	// Remove dataflow operation polling requests (ones where the job is PENDING or QUEUED)
	events = events.KeepIf(func(e *test.LogEntry) bool {
		if !strings.HasPrefix(e.Request.URL, "https://dataflow.googleapis.com/v1b3/") {
			return true
		}
		responseBody := e.Response.ParseBody()
		if responseBody == nil {
			return true
		}
		currentState, _, _ := unstructured.NestedString(responseBody, "currentState")
		switch currentState {
		case "JOB_STATE_PENDING", "JOB_STATE_QUEUED":
			return false
		}
		// Also handle when we're encoding enums as integers
		currentStateEnum, _, _ := unstructured.NestedInt64(responseBody, "currentState")
		switch currentStateEnum {
		case 9 /* JOB_STATE_PENDING */, 11 /* JOB_STATE_QUEUED */ :
			return false
		}
		return true
	})

	return events
}

// RewriteUserAgent removes volatile values from the user agent:
// it replaces the version with ${kccVersion}.
func RewriteUserAgent(events test.LogEntries) test.LogEntries {
	// Remove operation polling requests (ones where the operation is not ready)
	for _, event := range events {
		userAgent := event.Request.Header.Get("User-Agent")
		if userAgent != "" {
			currentVersion := version.GetVersion()
			userAgent = strings.ReplaceAll(userAgent, currentVersion, "${kccVersion}")
			event.Request.Header.Set("User-Agent", userAgent)
		}
	}

	return events
}

func (x *Normalizer) Render(events test.LogEntries) string {

	// Replace any dynamic IDs that appear in URLs
	for _, event := range events {
		url := event.Request.URL
		for k, v := range x.PathIDs {
			url = strings.ReplaceAll(url, "/"+k, "/"+v)
		}
		event.Request.URL = url
	}

	events = RemoveExtraEvents(events)

	jsonMutators := []test.JSONMutator{}
	addReplacement := func(path string, newValue string) {
		tokens := strings.Split(path, ".")
		jsonMutators = append(jsonMutators, func(url string, obj map[string]any) {
			_, found, _ := unstructured.NestedString(obj, tokens...)
			if found {
				if err := unstructured.SetNestedField(obj, newValue, tokens...); err != nil {
					klog.Fatalf("error setting field: %v", err)
				}
			}
		})
	}

	addSetStringReplacement := func(path string, newValue string) {
		jsonMutators = append(jsonMutators, func(url string, obj map[string]any) {
			if err := setStringAtPath(obj, path, newValue); err != nil {
				klog.Fatalf("error from setStringAtPath(%+v): %v", obj, err)
			}
		})
	}

	addReplacement("id", "000000000000000000000")
	addReplacement("uniqueId", "111111111111111111111")
	addReplacement("oauth2ClientId", "888888888888888888888")
	addReplacement("response.oauth2ClientId", "888888888888888888888")

	addReplacement("etag", "abcdef0123A=")
	addReplacement("serviceAccount.etag", "abcdef0123A=")
	addReplacement("response.etag", "abcdef0123A=")

	addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("creationTimestamp", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.createTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".monitoredProjects[].createTime", "2024-04-01T12:34:56.123456Z")

	addReplacement("lastUpdateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("dataCatalogTimestamps.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("dataCatalogTimestamps.updateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to cloudbuild
	addReplacement("metadata.completeTime", "2024-04-01T12:34:56.123456Z")

	// Specific to spanner
	addReplacement("metadata.startTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.updateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to vertexai
	addReplacement("blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
	addReplacement("response.blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")

	// Specific to BigTable
	addSetStringReplacement(".instances[].createTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".metadata.requestTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".metadata.finishTime", "2024-04-01T12:34:56.123456Z")

	// Specific to Sql
	addSetStringReplacement(".ipAddresses[].ipAddress", "10.1.2.3")
	addReplacement("serverCaCert.cert", "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n")
	addReplacement("serverCaCert.commonName", "common-name")
	addReplacement("serverCaCert.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("serverCaCert.expirationTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("serverCaCert.sha1Fingerprint", "12345678")
	addReplacement("serviceAccountEmailAddress", "p${projectNumber}-abcdef@gcp-sa-cloud-sql.iam.gserviceaccount.com")
	addReplacement("settings.backupConfiguration.startTime", "12:00")
	addReplacement("settings.settingsVersion", "123")

	// Replace any empty values in LROs; this is surprisingly difficult to fix in mockgcp
	//
	//     "response": {
	// 	-    "@type": "type.googleapis.com/google.protobuf.Empty"
	// 	+    "@type": "type.googleapis.com/google.protobuf.Empty",
	// 	+    "value": {}
	// 	   }
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
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

	// Add Essential Contacts specific normalizations
	addReplacement("validateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.validateTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".contacts[].validateTime", "2024-04-01T12:34:56.123456Z")

	events.PrettifyJSON(jsonMutators...)

	// Remove headers that just aren't very relevant to testing
	// Remove headers in request.
	events.RemoveHTTPRequestHeader("X-Goog-Api-Client")
	events.RemoveHTTPRequestHeader("X-Goog-User-Project")
	// Remove headers in response.
	events.RemoveHTTPResponseHeader("Date")
	events.RemoveHTTPResponseHeader("Alt-Svc")
	events.RemoveHTTPResponseHeader("Server-Timing")

	got := events.FormatHTTP()
	normalizers := []func(string) string{}
	normalizers = append(normalizers, ReplaceString(x.uniqueID, "${uniqueId}"))
	normalizers = append(normalizers, ReplaceString(x.project.ProjectID, "${projectId}"))
	normalizers = append(normalizers, ReplaceString(fmt.Sprintf("%d", x.project.ProjectNumber), "${projectNumber}"))
	for k, v := range x.PathIDs {
		normalizers = append(normalizers, ReplaceString(k, v))
	}
	for k := range x.OperationIDs {
		normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
	}

	for _, normalizer := range normalizers {
		got = normalizer(got)
	}
	return got
}

func (x *Normalizer) Preprocess(events []*test.LogEntry) {
	events = RewriteUserAgent(events)

	// Find "easy" operations and resources by looking for fully-qualified methods
	for _, event := range events {
		u := event.Request.URL
		if index := strings.Index(u, "?"); index != -1 {
			u = u[:index]
		}
		x.ExtractIDsFromLinks(u)
	}

	for _, event := range events {
		id := ""
		body := event.Response.ParseBody()
		val, ok := body["name"]
		if ok {
			s := val.(string)
			x.ExtractIDsFromLinks(s)

			// Also check for operations
			tokens := strings.Split(s, "/")
			// operation name format: operations/{operationId}
			if len(tokens) == 2 && tokens[0] == "operations" {
				id = strings.TrimPrefix(s, "operations/")
			}
			// operation name format: {prefix}/operations/{operationId}
			if len(tokens) > 2 && tokens[len(tokens)-2] == "operations" {
				id = tokens[len(tokens)-1]
			}
			// operation name format: operation-{operationId}
			if len(tokens) == 1 && strings.HasPrefix(tokens[0], "operation") {
				id = s
			}
		}
		if id != "" {
			x.OperationIDs[id] = true
		}
	}

	for _, event := range events {
		if !isGetOperation(event) {
			continue
		}
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		name, _, _ := unstructured.NestedString(responseBody, "response", "name")
		if strings.HasPrefix(name, "tagKeys/") {
			x.PathIDs[name] = "tagKeys/${tagKeyID}"
		}
	}

	// TODO: Remove this, it should now be done in normalize in mockcompute
	// Extract resource IDs / numbers from compute operations.
	// The number / id is in the targetID field, we infer the type from the targetLink field.
	for _, event := range events {
		if !isGetOperation(event) {
			continue
		}
		body := event.Response.ParseBody()
		targetLink, _, _ := unstructured.NestedString(body, "targetLink")
		targetId, _, _ := unstructured.NestedString(body, "targetId")
		if targetLink != "" && targetId != "" {
			u, _ := ParseGCPLink(targetLink)
			if u != nil {
				kind := u.PathItems[len(u.PathItems)-1].Resource

				placeholder := x.placeholderForGCPResource(kind, targetId)
				if placeholder != "" {
					// We _should_ differentiate between ID and number.
					// But this causes too many diffs right now.
					// if isNumber(targetId) {
					// 	x.PathIDs[targetId] = strings.Replace(placeholder, "ID", "Number", 1)
					// } else {
					// 	x.PathIDs[targetId] = placeholder
					// }
					x.PathIDs[targetId] = placeholder
				}
			}
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
