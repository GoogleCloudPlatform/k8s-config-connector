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
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func normalizeKRMObject(t *testing.T, u *unstructured.Unstructured, project testgcp.GCPProject, uniqueID string) error {
	replacements := NewReplacements()
	findLinksInKRMObject(t, replacements, u)

	annotations := u.GetAnnotations()
	if annotations["cnrm.cloud.google.com/observed-secret-versions"] != "" {
		// Includes resource versions, very volatile
		annotations["cnrm.cloud.google.com/observed-secret-versions"] = "(removed)"
	}
	if annotations["test.cnrm.cloud.google.com/reconcile-cookie"] != "" {
		// Deliberately volatile, ignore
		annotations["test.cnrm.cloud.google.com/reconcile-cookie"] = "(removed)"
	}
	u.SetAnnotations(annotations)

	visitor := newObjectWalker()

	// Apply replacements
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		return replacements.ApplyReplacements(s)
	})

	visitor.removePaths.Insert(".metadata.creationTimestamp")
	visitor.removePaths.Insert(".metadata.managedFields")
	visitor.removePaths.Insert(".metadata.resourceVersion")
	visitor.removePaths.Insert(".metadata.uid")

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
	visitor.replacePaths[".status.observedState.oauth2ClientID"] = "888888888888888888888"

	// Apigee
	visitor.replacePaths[".status.expiresAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	visitor.replacePaths[".status.createdAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	visitor.replacePaths[".status.lastModifiedAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	visitor.replacePaths[".status.observedState.createdAt"] = time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix()
	visitor.replacePaths[".status.observedState.lastModifiedAt"] = time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix()

	// Specific to AlloyDB
	visitor.replacePaths[".status.continuousBackupInfo[].enabledTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.ipAddress"] = "10.1.2.3"
	visitor.replacePaths[".status.outboundPublicIpAddresses"] = []string{"6.6.6.6", "8.8.8.8"}

	// Specific to CloudKMS
	visitor.replacePaths[".primary.createTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".primary.generateTime"] = "2024-04-01T12:34:56.123456Z"

	// Specific to BigQuery
	visitor.replacePaths[".spec.access[].userByEmail"] = "user@google.com"

	// Specific to Dataflow
	visitor.sortAndDeduplicateSlices.Insert(".spec.additionalExperiments")

	// Specific to Firestore
	visitor.replacePaths[".status.observedState.earliestVersionTime"] = "1970-01-01T00:00:00Z"

	// Specific to Sql
	visitor.replacePaths[".items[].etag"] = "abcdef0123A="
	visitor.replacePaths[".status.firstIpAddress"] = "10.1.2.3"
	visitor.replacePaths[".status.publicIpAddress"] = "10.1.2.3"
	visitor.replacePaths[".status.ipAddress"] = "10.1.2.3"
	visitor.replacePaths[".status.serverCaCert.cert"] = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n"
	visitor.replacePaths[".status.serverCaCert.commonName"] = "common-name"
	visitor.replacePaths[".status.serverCaCert.createTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.serverCaCert.expirationTime"] = "1970-01-01T00:00:00Z"
	visitor.replacePaths[".status.serverCaCert.sha1Fingerprint"] = "12345678"
	visitor.replacePaths[".status.serviceAccountEmailAddress"] = "p${projectNumber}-abcdef@gcp-sa-cloud-sql.iam.gserviceaccount.com"

	// Specific to Redis
	visitor.replacePaths[".status.observedState.uid"] = "0123456789abcdef"
	visitor.replacePaths[".status.observedState.pscConnections[].pscConnectionID"] = "${pscConnectionID}"
	visitor.replacePaths[".status.observedState.pscConnections[].address"] = "10.11.12.13"
	visitor.replacePaths[".status.observedState.discoveryEndpoints[].address"] = "10.11.12.13"

	// Specific to VertexAI
	visitor.replacePaths[".status.blobStoragePathPrefix"] = "cloud-ai-platform-00000000-1111-2222-3333-444444444444"
	visitor.replacePaths[".status.state[].diskUtilizationBytes"] = "1"
	visitor.replacePaths[".creator"] = "${creatorID}"

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
	visitor.ReplacePath(".spec.softDeletePolicy.effectiveTime", "1970-01-01T00:00:00Z")
	visitor.ReplacePath(".status.observedState.softDeletePolicy.effectiveTime", "1970-01-01T00:00:00Z")

	// Specific to Compute
	visitor.replacePaths[".status.observedState.certificateID"] = 1111111111111111
	visitor.replacePaths[".status.instanceId"] = "1111111111111111"
	visitor.replacePaths[".status.gatewayId"] = 1111111111111111
	visitor.replacePaths[".status.proxyId"] = 1111111111111111
	visitor.replacePaths[".status.mapId"] = 1111111111111111
	visitor.replacePaths[".status.id"] = 1111111111111111
	visitor.replacePaths[".status.certificateId"] = 1111111111111111
	visitor.replacePaths[".status.labelFingerprint"] = "abcdef0123A="
	visitor.replacePaths[".status.fingerprint"] = "abcdef0123A="

	// Specific to Certificate Manager
	visitor.replacePaths[".status.dnsResourceRecord[].data"] = "${uniqueId}"

	// Specific to Secret Manager
	visitor.replacePaths[".spec.expireTime"] = "2025-10-03T15:01:23Z"

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

	// Specific to DataFlow
	visitor.replacePaths[".status.jobId"] = "${jobID}"

	// Specific to SecretManager
	visitor.replacePaths[".expireTime"] = "2024-04-01T12:34:56.123456Z"

	// Specific to CloudIdentityMembership
	visitor.replacePaths[".membership.createTime"] = "2025-01-17T18:51:02.320337735Z"
	visitor.replacePaths[".membership.updateTime"] = "2025-01-17T18:51:02.320337735Z"

	// Specific to BigQueryConnectionConnection.
	visitor.replacePaths[".status.observedState.aws.accessRole.identity"] = "048077221682493034546"
	visitor.replacePaths[".status.observedState.azure.identity"] = "117243083562690747295"
	visitor.replacePaths[".status.observedState.cloudResource.serviceAccountID"] = "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com"
	visitor.replacePaths[".status.observedState.cloudSQL.serviceAccountID"] = "service-${projectNumber}@gcp-sa-bigqueryconnection.iam.gserviceaccount.com"
	visitor.replacePaths[".status.observedState.spark.serviceAccountID"] = "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com"

	// Specific to BigQueryDataTransferConfig
	if u.GetKind() == "BigQueryDataTransferConfig" {
		visitor.replacePaths[".status.observedState.nextRunTime"] = "1970-01-01T00:00:00Z"
		visitor.replacePaths[".status.observedState.ownerInfo.email"] = "user@google.com"
		visitor.replacePaths[".status.observedState.userID"] = "0000000000000000000"
		visitor.removePaths.Insert(".status.observedState.state") // data transfer run state, which depends on timing
	}

	// Specific to WorflowsWorkflow
	visitor.replacePaths[".status.observedState.revisionId"] = "revision-id-placeholder"
	visitor.replacePaths[".status.observedState.revisionCreateTime"] = "2024-04-01T12:34:56.123456Z"

	// Specific to DocumentAIProcessor
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.HasSuffix(path, ".status.observedState.processorVersionAliases[].processorVersion") {
			tokens := strings.Split(s, "/")
			if len(tokens) >= 2 {
				switch tokens[len(tokens)-2] {
				case "processorVersions":
					s = strings.ReplaceAll(s, tokens[len(tokens)-1], "${processorVersionID}")
				}
			}
		}
		return s
	})
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.HasSuffix(path, ".status.observedState.defaultProcessorVersion") {
			tokens := strings.Split(s, "/")
			if len(tokens) >= 2 && tokens[len(tokens)-2] == "processorVersions" {
				tokens[len(tokens)-1] = "${processorVersionID}"
				s = strings.Join(tokens, "/")
			}
		}
		return s
	})

	// Specific to VMwareEngineNetwork
	// normalize "observedState.vpcNetworks[].network"
	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if strings.HasSuffix(path, ".observedState.vpcNetworks[].network") {
			tokens := strings.Split(s, "/")
			if len(tokens) >= 2 && tokens[len(tokens)-2] == "networks" {
				tokens[len(tokens)-1] = "${networkId}"
				s = strings.Join(tokens, "/")
			}
		}
		return s
	})

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
		if len(tokens) >= 2 {
			typeName := tokens[len(tokens)-2]
			id := tokens[len(tokens)-1]

			// Remove any "verbs" we might be picking up by mistake
			// e.g. https://cloudresourcemanager.googleapis.com/v3/folders/${folderID}:move?alt=json&prettyPrint=false
			if strings.Contains(id, ":") {
				id = strings.Split(id, ":")[0]
			}

			if typeName == "datasets" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${datasetId}")
				})
			}
			if typeName == "folders" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${folderId}")
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
			if typeName == "transferConfigs" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${transferConfigID}")
				})
			}
			if typeName == "processors" {
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, id, "${processorID}")
				})
			}
		}

		id, _, _ := unstructured.NestedString(u.Object, "status", "selfLinkWithId")
		if id != "" {
			tokens := strings.Split(id, "/")
			n := len(tokens)
			if n >= 2 {
				typeName := tokens[len(tokens)-2]
				id := tokens[len(tokens)-1]
				if typeName == "targetGrpcProxies" {
					visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
						return strings.ReplaceAll(s, id, "${targetGrpcProxiesID}")
					})
				}
			}
		}

		// Get firewall policy id from firewall policy rule's externalRef and replace it
		externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
		if externalRef != "" {
			tokens := strings.Split(externalRef, "/")
			n := len(tokens)
			if n >= 3 {
				// e.g. "locations/global/firewallPolicies/${firewallPolicyID}/rules/9000"
				typeName := tokens[len(tokens)-2]
				firewallPolicyId := tokens[len(tokens)-3]
				if typeName == "rules" {
					visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
						return strings.ReplaceAll(s, firewallPolicyId, "${firewallPolicyID}")
					})
				}
			}
		}

		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID != "" {
			switch u.GroupVersionKind() {
			case schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringUptimeCheckConfig"}:
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, resourceID, "${uptimeCheckConfigId}")
				})

			case schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringGroup"}:
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, resourceID, "${monitoringGroupID}")
				})

			case schema.GroupVersionKind{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIdentityGroup"}:
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, resourceID, "${groupID}")
				})

			case schema.GroupVersionKind{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIdentityMembership"}:
				visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
					return strings.ReplaceAll(s, resourceID, "${membershipID}")
				})
			}
		}

		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			switch u.GroupVersionKind() {
			case schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeFirewallPolicy"}:
				// https://www.googleapis.com/compute/beta/locations/global/firewallPolicies/1059732409893
				selfLink = strings.TrimPrefix(selfLink, "https://www.googleapis.com/compute/v1/locations/global/")
				tokens := strings.Split(selfLink, "/")
				n := len(tokens)
				if n >= 2 {
					visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
						return strings.ReplaceAll(s, tokens[len(tokens)-1], "${firewallPolicyID}")
					})
				}
			}
		}
	}

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		return strings.ReplaceAll(s, "organizations/"+testgcp.TestOrgID.Get(), "organizations/${organizationID}")

	})

	return visitor.VisitUnstructured(u)
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
	removePaths              sets.Set[string]
	sortSlices               sets.Set[string]
	sortAndDeduplicateSlices sets.Set[string]
	replacePaths             map[string]any
	stringTransforms         []func(path string, value string) string
	objectTransforms         []func(path string, value map[string]any)
	sliceTransforms          []func(path string, value []any) []any

	stringReplacements []stringReplacement
}

type stringReplacement struct {
	Find    string
	Replace string
}

func newObjectWalker() *objectWalker {
	return &objectWalker{
		removePaths:              sets.New[string](),
		sortSlices:               sets.New[string](),
		sortAndDeduplicateSlices: sets.New[string](),
		replacePaths:             make(map[string]any),
	}
}

func (o *objectWalker) ReplacePath(path string, v any) {
	if v2, found := o.replacePaths[path]; found && !reflect.DeepEqual(v, v2) {
		klog.Fatalf("objectWalker has duplicate ReplacePath %q", path)
	}

	o.replacePaths[path] = v
}

func (o *objectWalker) ReplaceStringValue(oldValue string, newValue string) {
	// Check for duplicates
	for _, replacement := range o.stringReplacements {
		if replacement.Find == oldValue {
			if replacement.Replace == newValue {
				// Already have this replacement, no point adding it twice
				return
			}
			klog.Fatalf("objectWalker has duplicate ReplaceStringValue %q=%q and %q=%q", oldValue, replacement.Replace, oldValue, newValue)
		}
	}

	o.stringReplacements = append(o.stringReplacements, stringReplacement{Find: oldValue, Replace: newValue})

	// Make sure the biggest replacements are first, to avoid substring replacement non-determinism
	// e.g. subnetwork-a => ${subnet} but network-a => ${network}
	sort.Slice(o.stringReplacements, func(i, j int) bool {
		return len(o.stringReplacements[i].Find) > len(o.stringReplacements[j].Find)
	})
}

func (o *objectWalker) RemovePath(path string) {
	o.removePaths.Insert(path)
}

func (o *objectWalker) SortSlice(path string) {
	o.sortSlices.Insert(path)
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
	for _, fn := range o.objectTransforms {
		fn(path, m)
	}

	for k, v := range m {
		childPath := path + "." + k
		if o.removePaths.Has(childPath) {
			delete(m, k)
			continue // nothing left to process
		}

		if v2, found := o.replacePaths[childPath]; found {
			if v != nil && v != "null" { // don't replace nil/null values
				m[k] = v2
				continue // replacement value is assumed to be normalized
			}
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

func sortSlice(s []any, deduplicate bool) ([]any, error) {
	type entry struct {
		o       any
		sortKey string
	}

	var entries []entry
	for i := range s {
		j, err := json.Marshal(s[i])
		if err != nil {
			return nil, fmt.Errorf("error converting to json: %w", err)
		}
		entries = append(entries, entry{o: s[i], sortKey: string(j)})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].sortKey < entries[j].sortKey
	})

	out := make([]any, 0, len(s))
	for i := range s {
		if deduplicate && i > 0 && entries[i].sortKey == entries[i-1].sortKey {
			continue
		}
		out = append(out, entries[i].o)
	}

	return out, nil
}

func (o *objectWalker) visitSlice(s []any, path string) (any, error) {
	for _, fn := range o.sliceTransforms {
		s = fn(path+"[]", s)
	}

	for i, v := range s {
		v2, err := o.visitAny(v, path+"[]")
		if err != nil {
			return nil, err
		}
		s[i] = v2
	}

	// Note: do sorting "last" so we sort normalized values
	if o.sortSlices.Has(path) {
		sorted, err := sortSlice(s, false)
		if err != nil {
			return s, err
		}
		s = sorted
	}

	if o.sortAndDeduplicateSlices.Has(path) {
		sorted, err := sortSlice(s, true)
		if err != nil {
			return s, err
		}
		s = sorted
	}

	return s, nil
}

func (o *objectWalker) visitPrimitive(v any, _ string) (any, error) {
	return v, nil
}

func (o *objectWalker) visitString(s string, path string) (string, error) {
	for _, stringReplacement := range o.stringReplacements {
		s = strings.ReplaceAll(s, stringReplacement.Find, stringReplacement.Replace)
	}
	for _, fn := range o.stringTransforms {
		s = fn(path, s)
	}
	return s, nil
}

func (o *objectWalker) VisitUnstructured(v *unstructured.Unstructured) error {
	if err := o.visitMap(v.Object, ""); err != nil {
		return err
	}
	return nil
}

// findLinksInEvent looks for link paths and feeds the values into replacement.ExtractIDsFromLinks
func findLinksInEvent(t *testing.T, replacement *Replacements, event *test.LogEntry) {
	linkPaths := sets.New(
		".response.pscConnections[].forwardingRule",
		".response.pscConnections[].network",
		".selfLink",
	)

	wellKnownPaths := map[string]string{
		".pscConnections[].pscConnectionId":          "${pscConnectionID}",
		".response.pscConnections[].pscConnectionId": "${pscConnectionID}",
	}

	s := event.Response.Body
	if s == "" {
		return
	}

	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &obj); err != nil {
		t.Fatalf("error from json.Unmarshal(%q): %v", s, err)
		return
	}

	visitor := objectWalker{}

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if linkPaths.Has(path) {
			replacement.ExtractIDsFromLinks(s)
		}
		if v := wellKnownPaths[path]; v != "" {
			replacement.PathIDs[s] = v
		}
		return s
	})

	if err := visitor.visitMap(obj, ""); err != nil {
		t.Fatalf("visiting response object: %v", err)
	}
}

// findLinksInKRMObject looks for link paths and feeds the values into replacement.ExtractIDsFromLinks
func findLinksInKRMObject(t *testing.T, replacement *Replacements, u *unstructured.Unstructured) {
	linkPaths := sets.New(
		".status.observedState.pscConnections[].forwardingRule",
		".status.observedState.pscConnections[].network",
	)

	visitor := objectWalker{}

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if linkPaths.Has(path) {
			replacement.ExtractIDsFromLinks(s)
		}
		return s
	})

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		if s == "" {
			return s
		}

		switch path {
		case ".spec.organizationRef.external":
			id := strings.TrimPrefix(s, "organizations/")
			replacement.PathIDs[id] = "${organizationID}"
		case ".status.writerIdentity":
			if strings.HasPrefix(s, "serviceAccount:service-org-") && strings.HasSuffix(s, "@gcp-sa-logging.iam.gserviceaccount.com") {
				id := strings.TrimSuffix(strings.TrimPrefix(s, "serviceAccount:service-org-"), "@gcp-sa-logging.iam.gserviceaccount.com")
				replacement.PathIDs[id] = "${organizationID}"
			}
			if strings.HasPrefix(s, "serviceAccount:service-folder-") && strings.HasSuffix(s, "@gcp-sa-logging.iam.gserviceaccount.com") {
				id := strings.TrimSuffix(strings.TrimPrefix(s, "serviceAccount:service-folder-"), "@gcp-sa-logging.iam.gserviceaccount.com")
				replacement.PathIDs[id] = "${folderID}"
			}
		}
		return s
	})

	if err := visitor.visitMap(u.Object, ""); err != nil {
		t.Fatalf("visiting KRM object: %v", err)
	}
}

func NormalizeHTTPLog(t *testing.T, events test.LogEntries, services mockgcpregistry.Normalizer, project testgcp.GCPProject, uniqueID string, folderID string, organizationID string) {
	normalizer := NewNormalizer(uniqueID, project)

	normalizer.Preprocess(events)

	if organizationID != "" {
		normalizer.Replacements.PathIDs[organizationID] = "${organizationID}"
	}
	if folderID != "" {
		normalizer.Replacements.PathIDs[folderID] = "${folderID}"
	}
	if uniqueID != "" {
		normalizer.Replacements.PathIDs[uniqueID] = "${uniqueId}"
	}

	// Find any URLs
	for _, event := range events {
		findLinksInEvent(t, normalizer.Replacements, event)
	}

	// Remove idempotency tokens
	events.ReplaceRequestQueryParameter("requestId", "123456")

	// Remove headers that just aren't very relevant to testing
	// Remove headers in request.
	events.RemoveHTTPRequestHeader("X-Goog-Api-Client")
	// Remove header in response.
	events.RemoveHTTPResponseHeader("Date")
	events.RemoveHTTPResponseHeader("Alt-Svc")
	events.RemoveHTTPResponseHeader("Server-Timing")
	events.RemoveHTTPResponseHeader("X-Debug-Tracking-Id")
	events.RemoveHTTPResponseHeader("X-Guploader-Uploadid")
	events.RemoveHTTPResponseHeader("Etag")
	events.RemoveHTTPResponseHeader("Content-Length") // an artifact of encoding
	events.RemoveHTTPResponseHeader("Cache-Control")  // not really relevant to us

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

	normalizeHTTPResponses(t, services, events)

	// Normalize using the KRM normalization function
	events.PrettifyJSON(func(requestURL string, obj map[string]any) {
		u := &unstructured.Unstructured{}
		u.Object = obj
		if err := normalizeKRMObject(t, u, project, uniqueID); err != nil {
			t.Fatalf("error from normalizeObject: %v", err)
		}
	})

	// Apply replacements
	normalizer.Replacements.ApplyReplacementsToHTTPEvents(events)
}

func normalizeHTTPResponses(t *testing.T, normalizer mockgcpregistry.Normalizer, events test.LogEntries) {
	visitor := newObjectWalker()

	// If we get detailed info, don't record it - it's not part of the API contract
	visitor.removePaths.Insert(".error.errors[].debugInfo")

	// Common variables
	visitor.replacePaths[".etag"] = "abcdef0123A="
	visitor.replacePaths[".response.etag"] = "abcdef0123A="
	visitor.replacePaths[".serviceAccount.etag"] = "abcdef0123A="
	visitor.replacePaths[".response.uniqueId"] = "12345678"
	visitor.replacePaths[".response.startTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".response.endTime"] = "2024-04-01T12:34:56.123456Z"

	// Misc Operations
	visitor.replacePaths[".insertTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".endTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".user"] = "user@example.com"

	// Compute operations
	visitor.replacePaths[".fingerprint"] = "abcdef0123A="
	visitor.replacePaths[".startTime"] = "2024-04-01T12:34:56.123456Z"

	// Specific to Apigee
	visitor.replacePaths[".response.createdAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	visitor.replacePaths[".response.lastModifiedAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	visitor.replacePaths[".response.expiresAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	{
		visitor.sortSlices.Insert(".response.properties.property")
		visitor.sortSlices.Insert(".properties.property")
		visitor.replacePaths[".expiresAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
		visitor.replacePaths[".createdAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
		visitor.replacePaths[".lastModifiedAt"] = strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10)
	}

	for _, event := range events {
		// Compute URLs: Replace any compute beta URLs with v1 URLs
		// Terraform uses the /beta/ endpoints, but mocks and direct controller should use /v1/
		// This special handling to avoid diffs in http logs.
		// This can be removed once all Compute resources are migrated to direct controller.
		event.Request.URL = rewriteComputeURL(event.Request.URL)

		// Normalize etags in URLS
		event.Request.URL = normalizeEtagsInURL(event.Request.URL)
	}

	visitor.stringTransforms = append(visitor.stringTransforms, func(path string, s string) string {
		switch path {
		case ".network", ".region", ".selfLink", ".selfLinkWithId", ".sourceImage", ".subnetworks[]", ".target", ".targetLink", ".zone":
			return rewriteComputeURL(s)
		}
		return s
	})

	// Specific to LROs
	{
		// For reasons unclear, operations emit done: false and cancelRequested: false.
		// This seems to violate the normal behaviour of proto bool fields with implicit presence.
		// Easiest just to normalize away the GCP responses that are hard to produce!
		visitor.objectTransforms = append(visitor.objectTransforms, func(path string, m map[string]any) {
			if path == "." {
				if m["done"] == false {
					delete(m, "done")
				}
			}

			if path == ".metadata" {
				if m["cancelRequested"] == false {
					delete(m, "cancelRequested")
				}
			}
		})
	}

	// Specific to DataFlow
	{
		visitor.ReplacePath(".job.startTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".job.createTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".currentStateTime", "2024-04-01T12:34:56.123456Z")
		// The pipelineUrl includes a long random ID that does not appear elsewhere
		visitor.ReplacePath(".environment.sdkPipelineOptions.options.pipelineUrl", "${pipelineUrl}")
		visitor.sortAndDeduplicateSlices.Insert(".environment.experiments")
		visitor.sortAndDeduplicateSlices.Insert(".environment.sdkPipelineOptions.options.experiments")

		visitor.objectTransforms = append(visitor.objectTransforms, func(path string, m map[string]any) {
			switch path {
			case ".environment.userAgent",
				".environment.version",
				".jobMetadata",
				".pipelineDescription",
				".environment.sdkPipelineOptions.serialized_fn":
				// These fields have a _lot_ of information, but it isn't surfaced in KCC
				clear(m)
				m["removed"] = "simplicity"
			}
		})

		visitor.sliceTransforms = append(visitor.sliceTransforms, func(path string, a []any) []any {
			switch path {
			case ".environment.workerPools[]",
				".stageStates[]",
				".steps[]",
				".environment.sdkPipelineOptions.display_data[]":
				// These fields have a _lot_ of information, but it isn't surfaced in KCC
				return []any{}
			}
			return a
		})
	}

	// Specific to Redis
	{
		visitor.ReplacePath(".pscConnections[].address", "10.11.12.13")
		visitor.ReplacePath(".response.pscConnections[].address", "10.11.12.13")
		visitor.ReplacePath(".discoveryEndpoints[].address", "10.11.12.13")
		visitor.ReplacePath(".response.discoveryEndpoints[].address", "10.11.12.13")
	}

	// Specific to Sql
	{
		visitor.ReplacePath(".ipAddresses[].ipAddress", "10.1.2.3")
		visitor.ReplacePath(".serverCaCert.cert", "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n")
		visitor.ReplacePath(".serverCaCert.commonName", "common-name")
		visitor.ReplacePath(".serverCaCert.createTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".serverCaCert.expirationTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".serverCaCert.sha1Fingerprint", "12345678")
		visitor.ReplacePath(".serviceAccountEmailAddress", "p${projectNumber}-abcdef@gcp-sa-cloud-sql.iam.gserviceaccount.com")
		visitor.ReplacePath(".settings.backupConfiguration.startTime", "12:00")
		visitor.ReplacePath(".settings.settingsVersion", "123")
	}

	// Specific to BigQuery
	{
		visitor.SortSlice(".access")
		visitor.ReplacePath(".access[].userByEmail", "user@google.com")
	}

	// BigQueryConnection
	{
		visitor.ReplacePath(".cloudResource.serviceAccountId", "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com")
		visitor.ReplacePath(".creationTime", "123456789")
		visitor.ReplacePath(".lastModifiedTime", "123456789")
	}

	// Workflows
	{
		visitor.ReplacePath(".revisionCreateTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".response.revisionCreateTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".revisionId", "revision-id-placeholder")
		visitor.ReplacePath(".response.revisionId", "revision-id-placeholder")
	}

	// DocumentAI
	{
		visitor.ReplacePath(".metadata.commonMetadata.createTime", "2025-01-01T12:34:56.123456Z")
		visitor.ReplacePath(".metadata.commonMetadata.updateTime", "2025-01-02T12:34:56.123456Z")
	}

	// AI Platform
	{
		visitor.ReplacePath(".updateTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".nextRunTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".expirationTime", "2024-09-01T12:34:56.123456Z")
		visitor.ReplacePath(".schedules[].createTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".schedules[].nextRunTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".schedules[].startTime", "2024-04-01T12:34:56.123456Z")
	}

	// KMS
	{
		visitor.ReplacePath(".expireTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".generateTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".importJobs[].createTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".importJobs[].expireTime", "2024-04-01T12:34:56.123456Z")
		visitor.ReplacePath(".importJobs[].generateTime", "2024-04-01T12:34:56.123456Z")
	}

	// Network Management
	{
		visitor.ReplacePath(".reachabilityDetails.verifyTime", "2025-01-01T12:34:56.123456Z")
		visitor.ReplacePath(".response.reachabilityDetails.verifyTime", "2025-01-01T12:34:56.123456Z")
	}

	// Dataplex
	visitor.replacePaths[".response.metastoreStatus.updateTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".metastoreStatus.updateTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".lakes[].updateTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".lakes[].createTime"] = "2024-04-01T12:34:56.123456Z"
	visitor.replacePaths[".lakes[].metastoreStatus.updateTime"] = "2024-04-01T12:34:56.123456Z"

	// Run visitors
	events.PrettifyJSON(func(requestURL string, obj map[string]any) {
		// Deprecated: try to move these into mockgcp normalizers
		if err := visitor.visitMap(obj, ""); err != nil {
			t.Fatalf("error normalizing response: %v", err)
		}
	})

	// Run per-service replaceres
	{
		replacements := newObjectWalker()

		for _, entry := range events {
			normalizer.ConfigureVisitor(entry.Request.URL, replacements)
		}

		for _, entry := range events {
			normalizer.Previsit(entry, replacements)
		}

		events.PrettifyJSON(func(requestURL string, obj map[string]any) {
			if err := replacements.visitMap(obj, ""); err != nil {
				t.Fatalf("error normalizing response: %v", err)
			}
		})
	}
}

// Compute URLs: Replace any compute beta URLs with v1 URLs
func rewriteComputeURL(u string) string {
	for _, basePath := range []string{"https://compute.googleapis.com/compute", "https://www.googleapis.com/compute"} {
		if strings.HasPrefix(u, basePath+"/beta/") {
			u = basePath + "/v1/" + strings.TrimPrefix(u, basePath+"/beta/")
		}
	}
	return u
}

func normalizeEtagsInURL(u string) string {
	re := regexp.MustCompile(`etag=[a-zA-Z0-9%]+`)
	return re.ReplaceAllString(u, "etag=abcdef0123A")
}

// isGetOperation returns true if this is an operation poll request
func isGetOperation(e *test.LogEntry) bool {
	if strings.Contains(e.Request.URL, "/operations/${operationID}") {
		return true
	}
	if e.Request.URL == "/google.longrunning.Operations/GetOperation" {
		return true
	}
	return false
}
