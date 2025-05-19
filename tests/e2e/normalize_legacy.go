// Copyright 2025 Google LLC
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
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// LegacyNormalize is the "legacy" normalization steps,
// we should avoid adding to this function and instead add to the per-service normalization functions.
// Deprecated: add functionality to the per-service normalization instead.
func LegacyNormalize(t *testing.T, h *create.Harness, project testgcp.GCPProject, uniqueID string, events test.LogEntries) (string, []func(string) string) {

	r := NewReplacements()

	// Find "easy" operations and resources by looking for fully-qualified methods
	for _, event := range events {
		u := event.Request.URL
		if index := strings.Index(u, "?"); index != -1 {
			u = u[:index]
		}
		r.ExtractIDsFromLinks(u)
	}

	for _, event := range events {
		id := ""
		body := event.Response.ParseBody()
		val, ok := body["name"]
		if ok {
			s := val.(string)
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
			// SQL operations require a special case.
			if kind, ok := body["kind"]; ok && kind == "sql#operation" {
				id = s
			}
		}
		if id != "" {
			// Avoid marking some well-known values that are not operation ids
			switch id {
			case "projects":
			// Bigtable uses an unusual operation path: "operations/projects/${projectId}/instances/test-instance-${uniqueId}/locations/us-central1-b/operations/${operationID}"
			default:
				r.OperationIDs[id] = true
			}
		}
	}

	for _, event := range events {
		body := event.Response.ParseBody()
		if selfLinkWithId, _, _ := unstructured.NestedString(body, "selfLinkWithId"); selfLinkWithId != "" {
			r.ExtractIDsFromLinks(selfLinkWithId)
		}

		if billingAccountName, _, _ := unstructured.NestedString(body, "billingAccountName"); billingAccountName != "" {
			r.ExtractIDsFromLinks(billingAccountName)
		}

		// if targetId, _, _ := unstructured.NestedString(body, "targetId"); targetId != "" {
		// 	extractIDsFromLinks(selfLinkWithId)
		// }

		if conditions, _, _ := unstructured.NestedSlice(body, "conditions"); conditions != nil {
			for _, conditionAny := range conditions {
				condition := conditionAny.(map[string]any)
				name, _, _ := unstructured.NestedString(condition, "name")
				if name != "" {
					r.ExtractIDsFromLinks(name)
				}
			}
		}

		if val, ok := body["projectNumber"]; ok {
			s := val.(string)
			r.PathIDs[s] = "${projectNumber}"
		}
	}

	// Replace any operation IDs that appear in URLs
	for _, event := range events {
		u := event.Request.URL
		for operationID := range r.OperationIDs {
			u = strings.ReplaceAll(u, operationID, "${operationID}")
		}
		event.Request.URL = u
	}

	for _, event := range events {
		if !isGetOperation(event) {
			continue
		}
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		if name, _, _ := unstructured.NestedString(responseBody, "response", "name"); name != "" {
			r.ExtractIDsFromLinks(name)
		}
		if targetLink, _, _ := unstructured.NestedString(responseBody, "targetLink"); targetLink != "" {
			r.ExtractIDsFromLinks(targetLink)
		}
	}

	// Replace any dynamic IDs that appear in URLs
	for _, event := range events {
		u := event.Request.URL
		for k, v := range r.PathIDs {
			u = strings.ReplaceAll(u, "/"+k, "/"+v)
		}
		event.Request.URL = u
	}

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
		if status, _, _ := unstructured.NestedString(responseBody, "status"); status == "DONE" {
			return true
		}
		// remove if not done - and done can be omitted when false
		return false
	})

	jsonMutators := []test.JSONMutator{}
	addReplacement := func(path string, newValue string) {
		tokens := strings.Split(path, ".")
		jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
			_, found, _ := unstructured.NestedString(obj, tokens...)
			if found {
				if err := unstructured.SetNestedField(obj, newValue, tokens...); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		})
	}

	addSetStringReplacement := func(path string, newValue string) {
		jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
			if err := setStringAtPath(obj, path, newValue); err != nil {
				t.Fatalf("FAIL: error from setStringAtPath(%+v): %v", obj, err)
			}
		})
	}

	addReplacement("id", "000000000000000000000")
	addReplacement("uniqueId", "111111111111111111111")
	addReplacement("oauth2ClientId", "888888888888888888888")
	addReplacement("response.oauth2ClientId", "888888888888888888888")

	addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("expireTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("deleteLockExpireTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.expireTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.deleteTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.deleteLockExpireTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("creationTimestamp", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.createTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".monitoredProjects[].createTime", "2024-04-01T12:34:56.123456Z")

	addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.genericMetadata.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.updateTime", "2024-04-01T12:34:56.123456Z")

	// specific to apigateway
	addReplacement("managedService", "apigatewayapi-minimal-${uniqueId}-{generatedId}.apigateway.${projectId}.cloud.goog")
	addReplacement("response.managedService", "apigatewayapi-minimal-${uniqueId}-{generatedId}.apigateway.${projectId}.cloud.goog")
	// Specific to cloudbuild
	addReplacement("metadata.completeTime", "2024-04-01T12:34:56.123456Z")

	// Specific to spanner
	addReplacement("metadata.startTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.instance.updateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to spanner database
	addReplacement("earliestVersionTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.earliestVersionTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".metadata.progress[].startTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".metadata.progress[].endTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".metadata.commitTimestamps[]", "2024-04-01T12:34:56.123456Z")

	// Specific to redis
	addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.host", "10.1.2.3")
	addReplacement("response.reservedIpRange", "10.1.2.0/24")
	addReplacement("host", "10.1.2.3")
	addReplacement("reservedIpRange", "10.1.2.0/24")
	addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")

	// Specific to Compute
	addReplacement("natIP", "192.0.0.10")
	addReplacement("fingerprint", "abcdef0123A=")

	// Specific to Dataplex
	addReplacement("executionStatus.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.executionStatus.updateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("response.executionStatus.latestJob.uid", "0123456789abcdef")
	addReplacement("executionStatus.latestJob.uid", "0123456789abcdef")
	for _, event := range events {
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		selfLinkWithId, _, _ := unstructured.NestedString(responseBody, "executionStatus", "latestJob", "name")
		if selfLinkWithId != "" {
			tokens := strings.Split(selfLinkWithId, "/")
			n := len(tokens)
			if n >= 2 {
				kind := tokens[n-2]
				id := tokens[n-1]
				switch kind {
				case "jobs":
					r.PathIDs[id] = "0123456789abcdef"
				}
			}
		}
	}

	// Matches the mock ip address of Compute forwarding rule
	addReplacement("IPAddress", "8.8.8.8")
	addReplacement("pscConnectionId", "111111111111")
	for _, event := range events {
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		selfLinkWithId, _, _ := unstructured.NestedString(responseBody, "selfLinkWithId")
		if selfLinkWithId != "" {
			tokens := strings.Split(selfLinkWithId, "/")
			n := len(tokens)
			if n >= 2 {
				kind := tokens[n-2]
				id := tokens[n-1]
				switch kind {
				case "networkEdgeSecurityServices":
					r.PathIDs[id] = "${networkEdgeSecurityServiceID}"
				}
			}
		}
	}

	// Specific to vertexai
	addReplacement("blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
	addReplacement("response.blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
	addReplacement("state.diskUtilizationBytes", "1")
	for _, event := range events {
		responseBody := event.Response.ParseBody()
		if responseBody == nil {
			continue
		}
		metadataArtifact, _, _ := unstructured.NestedString(responseBody, "metadataArtifact")
		if metadataArtifact != "" {
			tokens := strings.Split(metadataArtifact, "/")
			n := len(tokens)
			if n >= 2 {
				kind := tokens[n-2]
				id := tokens[n-1]
				switch kind {
				case "artifacts":
					r.PathIDs[id] = "${artifactId}"
				}
			}
		}
		gcsBucket, _, _ := unstructured.NestedString(responseBody, "metadata", "gcsBucket")
		if gcsBucket != "" && strings.HasPrefix(gcsBucket, "cloud-ai-platform-") {
			r.PathIDs[gcsBucket] = "cloud-ai-platform-${bucketId}"
		}
	}

	// Specific to AlloyDB
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if val, found, _ := unstructured.NestedString(obj, "name"); found {
			if strings.Contains(val, "clusters/alloydb") ||
				strings.Contains(val, "instances/alloydb") ||
				strings.Contains(val, "backups/alloydb") {

				// Explicitly set `reconciling` to `false`.
				if _, found, _ := unstructured.NestedBool(obj, "reconciling"); !found {
					if err := unstructured.SetNestedField(obj, false, "reconciling"); err != nil {
						t.Fatal(err)
					}
				}

				// Replace the IP addresses in `outboundPublicIpAddresses` slice to test IP addresses.
				if _, found, _ := unstructured.NestedSlice(obj, "outboundPublicIpAddresses"); found {
					if err := unstructured.SetNestedStringSlice(obj, []string{"6.6.6.6", "8.8.8.8"}, "outboundPublicIpAddresses"); err != nil {
						t.Fatal(err)
					}
				}
			}
		}
	})
	// Boolean fields in LRO are omitted when false so we need
	// to add them back.
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if _, found, _ := unstructured.NestedMap(obj, "metadata"); found {
			if val, found, err := unstructured.NestedString(obj, "metadata", "@type"); err == nil && found && val == "type.googleapis.com/google.cloud.alloydb.v1beta.OperationMetadata" {
				if _, found, err := unstructured.NestedString(obj, "done"); err == nil && !found {
					// Explicitly set `done` to `false`.
					if err := unstructured.SetNestedField(obj, false, "done"); err != nil {
						t.Fatal(err)
					}
				}

				if _, found, err := unstructured.NestedString(obj, "metadata", "requestedCancellation"); err == nil && !found {
					// Explicitly set `metadata.requestedCancellation` to `false`.
					if err := unstructured.SetNestedField(obj, false, "metadata", "requestedCancellation"); err != nil {
						t.Fatal(err)
					}
				}

				if _, found, _ := unstructured.NestedMap(obj, "response"); found {
					if val, found, _ := unstructured.NestedString(obj, "response", "@type"); found &&
						val == "type.googleapis.com/google.cloud.alloydb.v1beta.Cluster" ||
						val == "type.googleapis.com/google.cloud.alloydb.v1beta.Instance" ||
						val == "type.googleapis.com/google.cloud.alloydb.v1beta.Backup" {
						// Explicitly set `reconciling` in response to `false`.
						if _, found, _ := unstructured.NestedBool(obj, "response", "reconciling"); !found {
							if err := unstructured.SetNestedField(obj, false, "response", "reconciling"); err != nil {
								t.Fatal(err)
							}
						}

						// Replace the IP addresses in `outboundPublicIpAddresses` slice to test IP addresses.
						if _, found, _ := unstructured.NestedSlice(obj, "response", "outboundPublicIpAddresses"); found {
							if err := unstructured.SetNestedStringSlice(obj, []string{"6.6.6.6", "8.8.8.8"}, "response", "outboundPublicIpAddresses"); err != nil {
								t.Fatal(err)
							}
						}
					}
				}
			}
		}
	})
	// Specific to BigQuery
	addSetStringReplacement(".access[].userByEmail", "user@google.com")

	// Specific to Firestore
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if _, found, _ := unstructured.NestedMap(obj, "response"); found {
			// Only run this mutator for firestore database objects.
			if val, found, err := unstructured.NestedString(obj, "response", "@type"); err == nil && found && val == "type.googleapis.com/google.firestore.admin.v1.Database" {
				// Only run this mutator for firestore database objects that have a name set in the response.
				if val, found, err := unstructured.NestedString(obj, "response", "name"); err == nil && found && val != "" {
					// Set name field to use human-readable ID, instead of UID
					// Note: This only works if firestore databases in all resource fixture test cases use the name "firestoredatabase-${uniqueId}"
					if err := unstructured.SetNestedField(obj, "projects/${projectId}/databases/firestoredatabase-${uniqueId}", "response", "name"); err != nil {
						t.Fatalf("FAIL: stting nested field: %v", err)
					}
				}
			}
		}
	})

	// Specific to PAM
	// Boolean fields in LRO are omitted when false so we need
	// to add them back.
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if _, found, _ := unstructured.NestedMap(obj, "metadata"); found {
			if val, found, err := unstructured.NestedString(obj, "metadata", "@type"); err == nil && found && val == "type.googleapis.com/google.cloud.privilegedaccessmanager.v1.OperationMetadata" {
				if _, found, err := unstructured.NestedString(obj, "done"); err == nil && !found {
					// Explicitly set `done` to `false`.
					if err := unstructured.SetNestedField(obj, false, "done"); err != nil {
						t.Fatalf("FAIL: setting nested field: %v", err)
					}
				}

				if _, found, err := unstructured.NestedString(obj, "metadata", "requestedCancellation"); err == nil && !found {
					// Explicitly set `metadata.requestedCancellation` to `false`.
					if err := unstructured.SetNestedField(obj, false, "metadata", "requestedCancellation"); err != nil {
						t.Fatalf("FAIL: setting nested field: %v", err)
					}
				}
			}

		}
	})

	// Specific to pubsub
	addReplacement("revisionCreateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("revisionId", "revision-id-placeholder")

	// Specific to CertificateManager
	addReplacement("response.dnsResourceRecord.data", uniqueID)
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if val, found, err := unstructured.NestedString(obj, "kind"); err != nil || !found || val != "sql#instance" {
			// Only run this mutator for sql instance objects.
			return
		}
		if _, found, _ := unstructured.NestedString(obj, "state"); !found {
			// Only run this mutator for response objects. This is a hack to identify response objects
			// for database instances, because they include the state field (as opposed to requests,
			// which do not).
			return
		}
		if _, found, _ := unstructured.NestedMap(obj, "settings"); found {
			if _, found, _ := unstructured.NestedStringSlice(obj, "settings", "authorizedGaeApplications"); !found {
				// Include settings.authorizedGaeApplications in response, even if it's empty.
				var val []string
				if err := unstructured.SetNestedStringSlice(obj, val, "settings", "authorizedGaeApplications"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
		if _, found, _ := unstructured.NestedMap(obj, "settings", "ipConfiguration"); found {
			if _, found, _ := unstructured.NestedStringSlice(obj, "settings", "ipConfiguration", "authorizedNetworks"); !found {
				// Include settings.ipConfiguration.authorizedNetworks in response, even if it's empty.
				var val []string
				if err := unstructured.SetNestedStringSlice(obj, val, "settings", "ipConfiguration", "authorizedNetworks"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
		if _, found, _ := unstructured.NestedString(obj, "gceZone"); found {
			// Hardcode the zone. GCP chooses this zone within the
			// region, and it varies based on availability.
			if err := unstructured.SetNestedField(obj, "us-central1-a", "gceZone"); err != nil {
				t.Fatalf("FAIL: setting nested field: %v", err)
			}
		}
		if ipConfig, found, _ := unstructured.NestedMap(obj, "settings", "ipConfiguration"); found {
			// Hack fix: remove unpublished field that's suddenly showing up in real gcp proto responses.
			delete(ipConfig, "serverCaMode")
			if err := unstructured.SetNestedMap(obj, ipConfig, "settings", "ipConfiguration"); err != nil {
				t.Fatalf("FAIL: setting nested field: %v", err)
			}
		}
	})
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if val, found, err := unstructured.NestedString(obj, "kind"); err != nil || !found || val != "sql#usersList" {
			// Only run this mutator for sql users list objects.
			return
		}
		if items, found, _ := unstructured.NestedSlice(obj, "items"); found {
			// Include items[].host in response, even if it's empty.
			newItems := []interface{}{}
			for _, item := range items {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if _, found, _ := unstructured.NestedStringSlice(itemMap, "host"); !found {
						if err := unstructured.SetNestedField(itemMap, "", "host"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
					}
					newItems = append(newItems, itemMap)
				}
			}
			if err := unstructured.SetNestedSlice(obj, newItems, "items"); err != nil {
				t.Fatalf("FAIL: setting nested field: %v", err)
			}
		}
	})

	// Specific to KMS
	addReplacement("policy.etag", "abcdef0123A=")
	addSetStringReplacement(".cryptoKeyVersions[].createTime", "2024-04-01T12:34:56.123456Z")
	addSetStringReplacement(".cryptoKeyVersions[].generateTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("destroyTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("generateTime", "2024-04-01T12:34:56.123456Z")

	// Specific to BigQueryConnectionConnection.
	addReplacement("aws.accessRole.identity", "048077221682493034546")
	addReplacement("azure.identity", "117243083562690747295")
	addReplacement("cloudResource.serviceAccountId", "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com")
	addReplacement("cloudSql.serviceAccountId", "service-${projectNumber}@gcp-sa-bigqueryconnection.iam.gserviceaccount.com")
	addReplacement("spark.serviceAccountId", "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com")

	// Specific to BigQueryTable
	addReplacement("materializedView.lastRefreshTime", "123456789")
	addReplacement("materializedViewStatus.refreshWatermark", "2024-04-01T12:34:56.123456Z")

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

	// Specific to Apigee
	addReplacement("lastModifiedAt", strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10))
	addReplacement("createdAt", strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10))

	// Specific to BigQueryDataTransferConfig
	addReplacement("nextRunTime", "2024-04-01T12:34:56.123456Z")
	addReplacement("ownerInfo.email", "user@google.com")
	addReplacement("userId", "0000000000000000000")
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if _, found, err := unstructured.NestedString(obj, "destinationDatasetId"); err != nil || !found {
			// This is a hack to only run this mutator for BigQueryDataTransferConfig objects.
			return
		}
		// special handling because the field includes dot
		if _, found, _ := unstructured.NestedString(obj, "params", "connector.authentication.oauth.clientId"); found {
			if err := unstructured.SetNestedField(obj, "client-id", "params", "connector.authentication.oauth.clientId"); err != nil {
				t.Fatalf("FAIL: setting nested field: %v", err)
			}
		}
		if _, found, _ := unstructured.NestedString(obj, "params", "connector.authentication.oauth.clientSecret"); found {
			if err := unstructured.SetNestedField(obj, "client-secret", "params", "connector.authentication.oauth.clientSecret"); err != nil {
				t.Fatalf("FAIL: setting nested field: %v", err)
			}
		}
		delete(obj, "state") // data transfer run state, which depends on timing
	})

	// Specific to IAPSettings
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if val, found, _ := unstructured.NestedString(obj, "name"); found {
			tokens := strings.Split(val, "/")
			// e.g. "projects/project-id/iap_web/compute-us-central1/services/service-id"
			if len(tokens) >= 6 && tokens[0] == "projects" && tokens[2] == "iap_web" && strings.Contains(tokens[3], "compute") && tokens[4] == "services" {
				tokens[len(tokens)-1] = "${serviceId}"
				if err := unstructured.SetNestedField(obj, strings.Join(tokens, "/"), "name"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
	})

	// Specific to DocumentAIProcessor
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		// normalize the processorVersionAliases
		aliases, found, _ := unstructured.NestedSlice(obj, "processorVersionAliases")
		if !found {
			return
		}
		for i := range aliases {
			aliasMap, ok := aliases[i].(map[string]any)
			if !ok {
				continue
			}
			processorVersion, found, _ := unstructured.NestedString(aliasMap, "processorVersion")
			if !found {
				continue
			}
			tokens := strings.Split(processorVersion, "/")
			// e.g. projects/project-id/locations/us/processors/processor-id/processorVersions/pretrained-ocr-v1.0-2020-09-23
			if len(tokens) >= 2 && tokens[len(tokens)-2] == "processorVersions" {
				tokens[len(tokens)-1] = "${processorVersionID}"
				if err := unstructured.SetNestedField(aliasMap, strings.Join(tokens, "/"), "processorVersion"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
		if err := unstructured.SetNestedField(obj, aliases, "processorVersionAliases"); err != nil {
			t.Fatalf("FAIL: setting nested field: %v", err)
		}

		// normalize the defaultProcessorVersion
		if val, found, _ := unstructured.NestedString(obj, "defaultProcessorVersion"); found {
			tokens := strings.Split(val, "/")
			if len(tokens) >= 2 && tokens[len(tokens)-2] == "processorVersions" {
				tokens[len(tokens)-1] = "${processorVersionID}"
				if err := unstructured.SetNestedField(obj, strings.Join(tokens, "/"), "defaultProcessorVersion"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
	})

	// Specific to VMwareEngineNetwork
	// normalize "vpcNetworks[].network"
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		if val, found, _ := unstructured.NestedString(obj, "name"); found {
			tokens := strings.Split(val, "/")
			if len(tokens) < 2 || tokens[len(tokens)-2] != "vmwareEngineNetworks" {
				return
			}
		}
		vpcNetworks, found, _ := unstructured.NestedSlice(obj, "vpcNetworks")
		if !found {
			return
		}
		for _, vpcNetwork := range vpcNetworks {
			if vpcNetworkMap, ok := vpcNetwork.(map[string]any); ok {
				if val, found, _ := unstructured.NestedString(vpcNetworkMap, "network"); found {
					tokens := strings.Split(val, "/")
					if len(tokens) >= 2 && tokens[len(tokens)-2] == "networks" {
						tokens[len(tokens)-1] = "${networkId}"
						if err := unstructured.SetNestedField(vpcNetworkMap, strings.Join(tokens, "/"), "network"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
					}
				}
			}
		}
		if err := unstructured.SetNestedSlice(obj, vpcNetworks, "vpcNetworks"); err != nil {
			t.Fatalf("FAIL: setting nested field: %v", err)
		}
	})
	// normalize "response.vpcNetworks[].network"
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		responseObj, found, _ := unstructured.NestedMap(obj, "response")
		if !found {
			return
		}
		name, found, _ := unstructured.NestedString(responseObj, "name")
		if !found || !strings.Contains(name, "vmwareEngineNetworks") {
			return
		}
		vpcNetworks, found, _ := unstructured.NestedSlice(responseObj, "vpcNetworks")
		if !found {
			return
		}
		for _, vpcNetwork := range vpcNetworks {
			if vpcNetworkMap, ok := vpcNetwork.(map[string]any); ok {
				if val, found, _ := unstructured.NestedString(vpcNetworkMap, "network"); found {
					tokens := strings.Split(val, "/")
					if len(tokens) >= 2 && tokens[len(tokens)-2] == "networks" {
						tokens[len(tokens)-1] = "${networkId}"
						if err := unstructured.SetNestedField(vpcNetworkMap, strings.Join(tokens, "/"), "network"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
					}
				}
			}
		}
		if err := unstructured.SetNestedSlice(responseObj, vpcNetworks, "vpcNetworks"); err != nil {
			t.Fatalf("FAIL: setting nested field: %v", err)
		}
		if err := unstructured.SetNestedMap(obj, responseObj, "response"); err != nil {
			t.Fatalf("FAIL: setting nested field: %v", err)
		}
	})

	// Specific to BackupPlanDR
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		// normalize "dataSource"
		if val, found, _ := unstructured.NestedString(obj, "dataSource"); found {
			tokens := strings.Split(val, "/")
			if len(tokens) >= 2 && tokens[len(tokens)-2] == "dataSources" {
				tokens[len(tokens)-1] = "${dataSourceID}"
				if err := unstructured.SetNestedField(obj, strings.Join(tokens, "/"), "dataSource"); err != nil {
					t.Fatalf("FAIL: setting nested field: %v", err)
				}
			}
		}
		// normalize "response.dataSource"
		responseObj, found, _ := unstructured.NestedMap(obj, "response")
		if found {
			if val, found, _ := unstructured.NestedString(responseObj, "dataSource"); found {
				tokens := strings.Split(val, "/")
				if len(tokens) >= 2 && tokens[len(tokens)-2] == "dataSources" {
					tokens[len(tokens)-1] = "${dataSourceID}"
					if err := unstructured.SetNestedField(responseObj, strings.Join(tokens, "/"), "dataSource"); err != nil {
						t.Fatalf("FAIL: setting nested field: %v", err)
					}
					if err := unstructured.SetNestedMap(obj, responseObj, "response"); err != nil {
						t.Fatalf("FAIL: setting nested field: %v", err)
					}
				}
			}
		}
	})

	// Remove error details which can contain confidential information
	jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
		response := obj["error"]
		if responseMap, ok := response.(map[string]any); ok {
			delete(responseMap, "details")
		}
	})
	addReplacement("creationTime", "123456789")
	addReplacement("lastModifiedTime", "123456789")

	events.PrettifyJSON(jsonMutators...)

	NormalizeHTTPLog(t, events, h.RegisteredServices(), project, uniqueID, testgcp.TestFolderID.Get(), testgcp.TestOrgID.Get())

	events = RemoveExtraEvents(events)

	got := events.FormatHTTP()
	normalizers := []func(string) string{}
	normalizers = append(normalizers, IgnoreComments)
	normalizers = append(normalizers, ReplaceString(uniqueID, "${uniqueId}"))
	normalizers = append(normalizers, ReplaceString(project.ProjectID, "${projectId}"))
	normalizers = append(normalizers, ReplaceString(fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}"))
	for k, v := range r.PathIDs {
		normalizers = append(normalizers, ReplaceString(k, v))
	}
	for k := range r.OperationIDs {
		normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
	}

	return got, normalizers

}
