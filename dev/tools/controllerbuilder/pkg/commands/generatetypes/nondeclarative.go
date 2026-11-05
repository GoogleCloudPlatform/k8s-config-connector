// Copyright 2026 Google LLC
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

package generatetypes

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Why non-declarative detection uses a hybrid of dynamic RPC checks and
// explicit collection maps (see docs/ai/declarative-resource-criteria.md):
//
//  1. Dynamic proto RPC checks (nonDeclarativeRPCReason) reliably catch:
//     - Read-only / discovered resources that expose Get/List (and sometimes
//       Delete) but neither Create nor Update (Criterion 1.1).
//     - One-shot executions that expose Create without Update or Delete, or
//       pair Create (!Update) with a Cancel<Resource> RPC (Criterion 1.2).
//     - High-cardinality data-plane collections that expose AIP-165 /
//       AIP-231–235 bulk mutation RPCs such as Purge<Plural>, BulkCreate<Plural>,
//       BatchCreate<Plural>, BatchDelete<Plural>, or DeleteAll<Plural>
//       (Criterion 3.1).
//
//  2. Explicit collection maps (ephemeralCollectionSegments and
//     nonDeclarativeLeafCollections) cover the remaining ~35% of cases where
//     the protobuf service exposes normal Create/Get/Update/Delete RPCs, so
//     CRUD inspection alone cannot distinguish them from control-plane
//     infrastructure. Attempting to infer these purely dynamically from state
//     enum names, cross-RPC request fields (e.g. RevertVolumeRequest), or
//     ancestor package lookups proved brittle (causing false positives on
//     resources whose parents live in another proto package, such as
//     StorageRapidCache under storage.googleapis.com/Bucket, or resources with
//     secondary state enums like NetappReplication.MirrorState). Each entry
//     below documents the concrete GCP proto resource(s) and rationale that
//     placed it in the map.

// ephemeralCollectionSegments lists google.api.resource pattern collection
// segments that represent ephemeral runtime sessions, conversations, job runs,
// or code-review/issue-tracker activity (Criterion 2 in
// docs/ai/declarative-resource-criteria.md). Any resource whose pattern
// includes one of these segments (as parent or leaf) is rejected by default.
var ephemeralCollectionSegments = map[string]string{
	// Encountered in google.cloud.aiplatform.v1.Session (AIPlatformSession),
	// google.cloud.aiplatform.v1beta1.SessionEvent (.../sessions/{session}/events/{event}),
	// google.cloud.dialogflow.v2.Context (projects/{project}/agent/sessions/{session}/contexts/{context}),
	// and google.cloud.dialogflow.cx.v3.SessionEntityType (.../sessions/{session}/entityTypes/{entity_type}).
	// Live end-user chat/reasoning sessions carry short TTLs (ttl/expire_time) or
	// have no control-plane CreateSession RPC at all (Dialogflow sessions are
	// implicit runtime IDs passed to DetectIntent), so neither the session nor
	// child resources scoped under a session ID can be reconciled via GitOps.
	"sessions": "ephemeral runtime",

	// Encountered in google.cloud.contactcenterinsights.v1.Conversation
	// (CCInsightsConversation) and google.cloud.dialogflow.v2.Participant
	// (projects/{project}/conversations/{conversation}/participants/{participant}).
	// Represents individual call/chat transcripts or active support conversations
	// (start_time/end_time, CompleteConversation), in contrast to durable
	// control-plane conversationProfiles or conversationDatasets.
	"conversations": "ephemeral runtime",

	// Encountered in google.cloud.datacatalog.lineage.v1.Run (DatalineageRun,
	// .../processes/{process}/runs/{run}) and
	// google.cloud.aiplatform.v1.TensorboardTimeSeries
	// (.../tensorboards/{tensorboard}/experiments/{experiment}/runs/{run}/timeSeries/{time_series}).
	// Exposes full CRUD in proto, but represents an individual execution run of a
	// pipeline or training process (start_time/end_time, COMPLETED/FAILED/ABORTED)
	// rather than a reusable job definition.
	"runs": "ephemeral execution",

	// Encountered in google.cloud.aiplatform.v1.Execution (AIPlatformExecution,
	// .../metadataStores/{metadata_store}/executions/{execution}). Represents a
	// single workflow or ML pipeline execution record (RUNNING/COMPLETE/FAILED/CANCELLED)
	// rather than a declarative job template.
	"executions": "ephemeral execution",

	// Standard AIP-151 long-running operation collection (.../operations/{operation});
	// tracks asynchronous RPC progress rather than durable desired state.
	"operations": "ephemeral execution",

	// Encountered in Vertex AI hyperparameter tuning (.../studies/{study}/trials/{trial});
	// represents individual optimization trial runs generated by the service.
	"trials": "ephemeral execution",

	// Encountered in google.cloud.securesourcemanager.v1.PullRequest
	// (SecuresourcemanagerPullRequest) and PullRequestComment
	// (.../repositories/{repository}/pullRequests/{pull_request}/pullRequestComments/{comment}).
	// Exposes CreatePullRequest/UpdatePullRequest alongside MergePullRequest,
	// OpenPullRequest, and ClosePullRequest (State.OPEN/CLOSED/MERGED); represents
	// developer code-review activity rather than repository configuration.
	"pullRequests": "ephemeral activity",

	// Encountered in google.cloud.securesourcemanager.v1.Issue and IssueComment
	// (.../repositories/{repository}/issues/{issue}/issueComments/{comment}).
	// Exposes Create/Update/Delete alongside OpenIssue/CloseIssue; represents
	// issue-tracker activity rather than repository infrastructure.
	"issues": "ephemeral activity",

	// Encountered in google.cloud.securesourcemanager.v1.IssueComment
	// (.../issues/{issue}/issueComments/{comment}). Exposes standard
	// Create/Get/Update/Delete in proto with no bulk RPCs, so CRUD inspection
	// alone cannot distinguish it; excluded as collaborative discussion activity.
	"issueComments": "ephemeral activity",

	// Encountered in google.cloud.securesourcemanager.v1.PullRequestComment
	// (.../pullRequests/{pull_request}/pullRequestComments/{comment}). Represents
	// individual code-review comments on a pull request.
	"pullRequestComments": "ephemeral activity",
}

// nonDeclarativeLeafCollections lists google.api.resource leaf collection
// segments that represent high-cardinality data-plane records, immutable
// version snapshots / eval fixtures, or point-in-time backups (Criterion 3 in
// docs/ai/declarative-resource-criteria.md). These are checked only on the
// leaf collection segment so that declarative resources nested under a version
// segment (e.g. ApihubSpec at "projects/*/locations/*/apis/*/versions/*/specs/*")
// are not false-flagged.
var nonDeclarativeLeafCollections = map[string]string{
	// Criterion 3.1: High-cardinality data-plane rows, documents, SKUs & telemetry records

	// Encountered in google.cloud.chronicle.v1.DataTableRow (ChronicleDataTableRow,
	// .../dataTables/{data_table}/dataTableRows/{data_table_row}). While the parent
	// ChronicleDataTable is control-plane schema configuration, DataTableRow
	// represents individual table rows (row_time_to_live, BulkCreateDataTableRows).
	"dataTableRows": "data-plane record",

	// Encountered in google.cloud.dialogflow.v2.Document (DialogflowDocument,
	// .../knowledgeBases/{knowledge_base}/documents/{document}) and
	// google.cloud.discoveryengine.v1.Document (DiscoveryengineDocument,
	// .../dataStores/{data_store}/branches/{branch}/documents/{document}).
	// Note that dialogflow.v2.Document exposes standard Create/Get/Update/Delete
	// plus ImportDocuments/ReloadDocument (without a Purge/Batch prefix), so the
	// leaf collection rule is needed alongside the dynamic bulk-RPC check. The
	// parent KnowledgeBase/DataStore is control-plane infrastructure, whereas
	// Document is ingested search/knowledge content.
	"documents": "data-plane document",

	// Encountered in google.cloud.retail.v2.Product (RetailProduct,
	// .../catalogs/{catalog}/branches/{branch}/products/{product}). Represents
	// individual e-commerce SKU catalog items (price_info, available_quantity,
	// ttl, PurgeProducts, ImportProducts) inside a Catalog.
	"products": "data-plane catalog item",

	// Encountered in grafeas.v1.Occurrence (GrafeasOccurrence,
	// projects/{project}/occurrences/{occurrence}). Represents high-volume
	// per-image vulnerability, build, and attestation scan occurrences
	// (BatchCreateOccurrences), in contrast to control-plane Note definitions.
	"occurrences": "data-plane telemetry",

	// Encountered in google.cloud.aiplatform.v1.Artifact (AIPlatformArtifact,
	// .../metadataStores/{metadata_store}/artifacts/{artifact}). Represents
	// individual ML Metadata lineage artifact records (PurgeArtifacts) logged by
	// pipelines inside a control-plane MetadataStore.
	"artifacts": "data-plane lineage",

	// Encountered in google.cloud.aiplatform.v1.Context (AIPlatformContext,
	// .../metadataStores/{metadata_store}/contexts/{context}) and
	// google.cloud.dialogflow.v2.Context (.../sessions/{session}/contexts/{context}).
	// Represents ML lineage grouping records (PurgeContexts) or conversational
	// turn state (lifespan_count, DeleteAllContexts).
	"contexts": "data-plane lineage",

	// Encountered in google.cloud.aiplatform.v1beta1.SessionEvent
	// (.../sessions/{session}/events/{event}) and
	// google.cloud.datacatalog.lineage.v1.LineageEvent. Represents immutable
	// timestamped event log entries.
	"events": "data-plane telemetry",

	// Encountered in google.cloud.aiplatform.v1.TensorboardTimeSeries
	// (.../tensorboards/{tensorboard}/experiments/{experiment}/runs/{run}/timeSeries/{time_series}).
	// Represents logged metric time-series streams (ReadTensorboardTimeSeriesData,
	// BatchCreateTensorboardTimeSeries).
	"timeSeries": "data-plane telemetry",

	// Encountered in google.cloud.dialogflow.v2.Participant
	// (projects/{project}/conversations/{conversation}/participants/{participant}).
	// Represents a live human or virtual agent participant in an active
	// conversation (AnalyzeContent, no DeleteParticipant RPC).
	"participants": "data-plane runtime participant",

	// Criterion 3.2: Immutable version snapshots, image tags, experiments & test cases

	// Encountered in google.cloud.dialogflow.cx.v3.Version (DialogflowVersion,
	// .../agents/{agent}/flows/{flow}/versions/{version}). Mints an immutable
	// point-in-time snapshot of a Flow (LoadVersion, CompareVersions,
	// State.SUCCEEDED). Checked only on the leaf segment so resources nested
	// under an ancestor version (e.g. ApihubSpec) remain allowed.
	"versions": "version snapshot",

	// Encountered in google.cloud.aiplatform.v1.DatasetVersion
	// (AIPlatformDatasetVersion, .../datasets/{dataset}/datasetVersions/{dataset_version}).
	// Exposes standard Create/Get/Update/Delete in proto (plus RestoreDatasetVersion)
	// with a server-generated ID, so CRUD checks alone cannot distinguish it;
	// excluded because it mints point-in-time dataset snapshots.
	"datasetVersions": "version snapshot",

	// Encountered in google.devtools.artifactregistry.v1.Tag (ArtifactRegistryTag,
	// .../repositories/{repository}/packages/{package}/tags/{tag}). Exposes
	// clean CreateTag/GetTag/UpdateTag/DeleteTag in proto with no bulk/purge RPCs,
	// so dynamic RPC checks cannot catch it. Excluded because its parent
	// packages/{package} has no CreatePackage RPC (packages are populated by
	// data-plane artifact uploads like `docker push`) and tags are per-build
	// mutable pointers managed by CI/CD pipelines.
	"tags": "artifact tag",

	// Encountered in google.cloud.dialogflow.cx.v3.Experiment (DialogflowExperiment,
	// .../agents/{agent}/environments/{environment}/experiments/{experiment}).
	// Exposes standard CRUD plus StartExperiment/StopExperiment (start_time/end_time,
	// State.DONE/ROLLOUT_FAILED); represents time-bounded A/B rollout experiments.
	"experiments": "evaluation/experiment",

	// Encountered in google.cloud.dialogflow.cx.v3.TestCase (DialogflowTestCase,
	// .../agents/{agent}/testCases/{test_case}). Represents conversational
	// evaluation fixtures (RunTestCase, BatchRunTestCases, BatchDeleteTestCases,
	// no single-resource DeleteTestCase).
	"testCases": "evaluation/test-case",

	// Encountered in google.cloud.dialogflow.cx.v3.Example (DialogflowExample,
	// .../agents/{agent}/playbooks/{playbook}/examples/{example}). Exposes
	// standard Create/Get/Update/Delete in proto with no bulk or custom lifecycle
	// RPCs, so dynamic RPC inspection cannot distinguish it; excluded because
	// few-shot conversation trace examples (playbook_input/playbook_output) are
	// prompt/training data rather than cloud infrastructure configuration.
	"examples": "training/evaluation example",

	// Encountered in google.cloud.metastore.v1.MetadataImport
	// (MetastoreMetadataImport, .../services/{service}/metadataImports/{metadata_import}).
	// Exposes Create/Get/List/Update (no Delete) for a one-shot database dump
	// import operation (database_dump, end_time, State.SUCCEEDED/FAILED).
	"metadataImports": "one-shot import",

	// Criterion 3.3: Point-in-time backups & snapshots

	// Encountered in google.cloud.netapp.v1.Backup (NetappBackup,
	// .../backupVaults/{backup_vault}/backups/{backup}) and
	// google.spanner.admin.database.v1.Backup (SpannerBackup,
	// .../instances/{instance}/backups/{backup}). Both expose standard
	// Create/Get/Update/Delete in proto (unlike declarative backupPolicies or
	// backupVaults), but represent point-in-time data captures of a live volume
	// or database (RestoreBackupFiles, CopyBackup, expire_time). Requiring
	// --allow-non-declarative ensures point-in-time backup CRDs are opted into
	// deliberately rather than swept into control-plane batches.
	"backups": "point-in-time backup",

	// Encountered in google.cloud.netapp.v1.Snapshot (NetappSnapshot,
	// .../volumes/{volume}/snapshots/{snapshot}). Exposes clean
	// CreateSnapshot/GetSnapshot/UpdateSnapshot/DeleteSnapshot in proto with no
	// special RPCs on Snapshot itself (the revert action lives on the parent
	// Volume's RevertVolumeRequest.snapshot_id), so dynamic RPC inspection on
	// the snapshots collection cannot distinguish it from control-plane config;
	// gated behind --allow-non-declarative for explicit opt-in.
	"snapshots": "point-in-time snapshot",
}

// nonDeclarativeResourceReason reports whether a proto message appears
// non-declarative or data-plane according to Criteria 1, 2, and 3 in
// docs/ai/declarative-resource-criteria.md.
func nonDeclarativeResourceReason(api *protoapi.Proto, msg protoreflect.MessageDescriptor) (string, bool) {
	md := protoapi.GetResourceMetadata(msg)
	if md == nil {
		return "", false
	}
	for _, pattern := range md.Patterns {
		if reason, ok := nonDeclarativePatternReason(pattern); ok {
			return reason, true
		}
	}
	if api != nil {
		if reason, ok := nonDeclarativeRPCReason(api, msg, md); ok {
			return reason, true
		}
	}
	return "", false
}

// nonDeclarativePatternReason inspects a single google.api.resource pattern
// string and returns a human-readable explanation if any literal collection
// segment is an ephemeral runtime collection (Criterion 2) or if the leaf
// collection segment is a data-plane, version-snapshot, or backup collection
// (Criterion 3).
func nonDeclarativePatternReason(pattern string) (string, bool) {
	literals := patternLiterals(pattern)
	if len(literals) == 0 {
		return "", false
	}
	for _, seg := range literals {
		if category, ok := ephemeralCollectionSegments[seg]; ok {
			return fmt.Sprintf("pattern %q contains %s collection %q", pattern, category, seg), true
		}
	}
	leaf := literals[len(literals)-1]
	if category, ok := nonDeclarativeLeafCollections[leaf]; ok {
		return fmt.Sprintf("pattern %q has %s leaf collection %q", pattern, category, leaf), true
	}
	return "", false
}

// bulkDataPlaneRPCPrefixes lists AIP-165 / AIP-231–235 collection-level bulk
// mutation and purge RPC prefixes that indicate a high-cardinality data-plane
// collection rather than a declarative control-plane resource. We intentionally
// include BatchCreate/BatchUpdate/BatchDelete and custom methods (AIP-136) so
// that suspected batch/data-plane resources are surfaced in the rejection
// report rather than slipping through as false negatives.
var bulkDataPlaneRPCPrefixes = []string{
	"Purge",
	"BulkCreate",
	"BulkUpdate",
	"BulkDelete",
	"BulkReplace",
	"BatchCreate",
	"BatchUpdate",
	"BatchDelete",
	"DeleteAll",
}

// nonDeclarativeRPCReason inspects service methods in the resource message's
// proto package for:
//   - Criterion 1.1: Read-only / discovered resources that define Get or Delete
//     RPCs but neither Create nor Update.
//   - Criterion 1.2: One-shot execution resources that define Create without
//     Update and either lack Delete or expose a Cancel<Resource> RPC.
//   - Criterion 3.1: High-cardinality data-plane collections that expose
//     Purge/Bulk*/Batch*/DeleteAll RPCs on the collection.
func nonDeclarativeRPCReason(api *protoapi.Proto, msg protoreflect.MessageDescriptor, md *protoapi.ResourceMetadata) (string, bool) {
	pkg := msg.ParentFile().Package()
	name := string(msg.Name())
	plural := md.Collection
	capPlural := ""
	if len(plural) > 0 {
		capPlural = strings.ToUpper(plural[:1]) + plural[1:]
	}

	var hasCreate, hasUpdate, hasDelete, hasGet, hasCancel bool
	var bulkRPCs []string

	api.Files().RangeFilesByPackage(pkg, func(fd protoreflect.FileDescriptor) bool {
		for i := 0; i < fd.Services().Len(); i++ {
			svc := fd.Services().Get(i)
			for j := 0; j < svc.Methods().Len(); j++ {
				m := svc.Methods().Get(j)
				mName := string(m.Name())
				httpRule, _ := proto.GetExtension(m.Options(), annotations.E_Http).(*annotations.HttpRule)
				bindings := httpVerbAndPaths(httpRule)
				anyPathMatches := false
				for _, b := range bindings {
					if matchesAnyResourcePattern(b.path, md.Patterns) {
						anyPathMatches = true
						switch {
						case b.verb == "POST" && !strings.Contains(b.path, ":") && strings.HasPrefix(mName, "Create"):
							hasCreate = true
						case (b.verb == "PATCH" || b.verb == "PUT") && strings.HasPrefix(mName, "Update"):
							hasUpdate = true
						case b.verb == "DELETE" && strings.HasPrefix(mName, "Delete"):
							hasDelete = true
						case b.verb == "GET" && strings.HasPrefix(mName, "Get"):
							hasGet = true
						}
					}
				}

				switch mName {
				case "Create" + name:
					hasCreate = true
				case "Update" + name:
					hasUpdate = true
				case "Delete" + name:
					hasDelete = true
				case "Get" + name:
					hasGet = true
				case "Cancel" + name:
					hasCancel = true
				}

				for _, prefix := range bulkDataPlaneRPCPrefixes {
					if mName == prefix+name || (capPlural != "" && mName == prefix+capPlural) || (anyPathMatches && strings.HasPrefix(mName, prefix)) {
						bulkRPCs = append(bulkRPCs, mName)
					}
				}
			}
		}
		return true
	})

	if !hasCreate && !hasUpdate && (hasGet || hasDelete) {
		return fmt.Sprintf("service defines neither Create%s nor Update%s (read-only/discovered resource)", name, name), true
	}
	if hasCreate && !hasUpdate && (!hasDelete || hasCancel) {
		return fmt.Sprintf("service defines Create%s without Update%s (delete=%v, cancel=%v; one-shot execution)", name, name, hasDelete, hasCancel), true
	}
	if len(bulkRPCs) > 0 {
		return fmt.Sprintf("service exposes bulk/purge data-plane RPCs %v on %q", bulkRPCs, plural), true
	}
	return "", false
}

func matchesAnyResourcePattern(httpPath string, resourcePatterns []string) bool {
	for _, p := range resourcePatterns {
		if matchesResourcePattern(httpPath, p) {
			return true
		}
	}
	return false
}

func matchesResourcePattern(httpPath, resourcePattern string) bool {
	if httpPath == "" || resourcePattern == "" {
		return false
	}
	base := httpPath
	if idx := strings.IndexByte(base, ':'); idx >= 0 {
		base = base[:idx]
	}
	base = strings.ReplaceAll(base, "{", "")
	base = strings.ReplaceAll(base, "}", "")
	if idx := strings.LastIndexByte(base, '='); idx >= 0 {
		base = base[idx+1:]
	}
	hLits := patternLiterals(base)
	pLits := patternLiterals(resourcePattern)
	if len(hLits) == 0 || len(hLits) != len(pLits) {
		return false
	}
	for i := range pLits {
		if hLits[i] != pLits[i] {
			return false
		}
	}
	return true
}

func patternLiterals(p string) []string {
	var out []string
	for _, s := range strings.Split(strings.Trim(p, "/"), "/") {
		if s == "" || s == "*" || strings.HasPrefix(s, "{") {
			continue
		}
		if len(out) == 0 && len(s) >= 2 && s[0] == 'v' && s[1] >= '0' && s[1] <= '9' {
			continue
		}
		out = append(out, s)
	}
	return out
}

type httpBinding struct {
	verb string
	path string
}

func httpVerbAndPaths(r *annotations.HttpRule) []httpBinding {
	if r == nil {
		return nil
	}
	var out []httpBinding
	if verb, path := httpVerbAndPath(r); verb != "" || path != "" {
		out = append(out, httpBinding{verb: verb, path: path})
	}
	for _, b := range r.GetAdditionalBindings() {
		if verb, path := httpVerbAndPath(b); verb != "" || path != "" {
			out = append(out, httpBinding{verb: verb, path: path})
		}
	}
	return out
}

func httpVerbAndPath(r *annotations.HttpRule) (string, string) {
	if r == nil {
		return "", ""
	}
	switch p := r.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return "GET", p.Get
	case *annotations.HttpRule_Post:
		return "POST", p.Post
	case *annotations.HttpRule_Put:
		return "PUT", p.Put
	case *annotations.HttpRule_Patch:
		return "PATCH", p.Patch
	case *annotations.HttpRule_Delete:
		return "DELETE", p.Delete
	case *annotations.HttpRule_Custom:
		if p.Custom != nil {
			return p.Custom.Kind, p.Custom.Path
		}
	}
	return "", ""
}
