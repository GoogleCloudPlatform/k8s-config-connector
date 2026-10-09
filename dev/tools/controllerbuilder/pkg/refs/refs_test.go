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

package refs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestClassify pins one field per branch of Classify. TestMissingRefs turns an
// IsReference into a missingrefs.txt entry and a NotRepresentable into a
// refs_not_representable.txt entry, so a branch that changes verdict moves
// entries between a ratchet and a golden file.
func TestClassify(t *testing.T) {
	for _, tc := range []struct {
		name        string
		fieldPath   string
		desc        string
		wantVerdict Verdict
		wantReason  string
	}{
		{
			name:        "a brace template names a resource",
			fieldPath:   ".spec.network",
			desc:        "Format: projects/{project}/global/networks/{network}",
			wantVerdict: IsReference,
		},
		{
			name:        "an angle-bracket template after a backtick counts too",
			fieldPath:   ".spec.notificationChannels[]",
			desc:        "Must be of the format `projects/<project_id_or_number>/notificationChannels/<channel_id>`",
			wantVerdict: IsReference,
		},
		{
			name:        "locations/ without a placeholder is prose",
			fieldPath:   ".spec.allowedLocations[]",
			desc:        "Resources may only be created in locations/us-central1.",
			wantVerdict: NotAReference,
		},
		{
			name:        "a service account is a reference by name alone",
			fieldPath:   ".spec.serviceAccount",
			wantVerdict: IsReference,
		},
		{
			name:        "a bucket described as Cloud Storage is a reference",
			fieldPath:   ".spec.bucket",
			desc:        "The Cloud Storage bucket to write to.",
			wantVerdict: IsReference,
		},
		{
			name:        "a zone is a coordinate, not a resource",
			fieldPath:   ".spec.placement.zone",
			desc:        "Format: projects/{project}/zones/{zone}",
			wantVerdict: NotAReference,
		},
		{
			// The plural forms are excluded for the same reason as the
			// singular: a field naming a set of coordinates is no more a
			// reference than one naming a single coordinate.
			name:        "a set of zones is still coordinates",
			fieldPath:   ".spec.placement.zones",
			desc:        "Format: projects/{project}/zones/{zone}",
			wantVerdict: NotAReference,
		},
		{
			name:        "a region is a coordinate",
			fieldPath:   ".spec.region",
			desc:        "Format: projects/{project}/regions/{region}",
			wantVerdict: NotAReference,
		},
		{
			name:        "a set of regions is still coordinates",
			fieldPath:   ".spec.regions",
			desc:        "Format: projects/{project}/regions/{region}",
			wantVerdict: NotAReference,
		},
		{
			name:        "a set of locations is still coordinates",
			fieldPath:   ".spec.locations",
			desc:        "Format: projects/{project}/locations/{location}",
			wantVerdict: NotAReference,
		},
		{
			name:        "an accelerator type is a catalogue entry",
			fieldPath:   ".spec.acceleratorType",
			desc:        "Format: projects/{project}/zones/{zone}/acceleratorTypes/{type}",
			wantVerdict: NotAReference,
		},
		{
			name:        "a regex matches a set of resources",
			fieldPath:   ".spec.cloudStorageRegex.bucketNameRegex",
			desc:        "Regex to test the bucket name against.",
			wantVerdict: NotAReference,
		},
		{
			name:        "a wildcard path matches a set of objects",
			fieldPath:   ".spec.gcsSource.uris[]",
			desc:        "Google Cloud Storage URI(-s) to the input file(s). May contain wildcards.",
			wantVerdict: NotAReference,
		},
		{
			name:        "an image URI is not a storage object",
			fieldPath:   ".spec.executorImageURI",
			desc:        "An image in Artifact Registry. Google Cloud Storage paths will be mapped to local paths.",
			wantVerdict: NotRepresentable,
			wantReason:  "container-image-uri-not-a-storage-object",
		},
		{
			name:        "bq:// has no KCC parsing",
			fieldPath:   ".spec.outputTable",
			desc:        "The table, as bq://project.dataset.table.",
			wantVerdict: NotRepresentable,
			wantReason:  "bq-scheme-not-a-gcp-resource-name",
		},
		{
			name:        "a BigQuery URI is not a resource name",
			fieldPath:   ".spec.tableURI",
			desc:        "The BigQuery table to read.",
			wantVerdict: NotRepresentable,
			wantReason:  "bigquery-uri-not-a-gcp-resource-name",
		},
		{
			name:        "a storage prefix needs a bucket reference plus a path",
			fieldPath:   ".spec.outputURIPrefix",
			desc:        "The Cloud Storage location to write to.",
			wantVerdict: NotRepresentable,
			wantReason:  "gcs-prefix-needs-bucket-ref-plus-path",
		},
		{
			name:        "a storage object path stays a string for now",
			fieldPath:   ".spec.configURI",
			desc:        "A gs:// path to the config file.",
			wantVerdict: NotRepresentable,
			wantReason:  "gcs-object-path-string-for-now-decomposable-as-bucketref-plus-path",
		},
		{
			name:        "a gs:// path in a field not named uri is not a resource name",
			fieldPath:   ".spec.terraformBlueprint.gcsSource",
			desc:        "URI of an object in Google Cloud Storage. Format: gs://{bucket}/{object}",
			wantVerdict: NotRepresentable,
			wantReason:  "gcs-scheme-not-a-gcp-resource-name",
		},
		{
			name:        "a gs:// list in a field not named uris is not a resource name either",
			fieldPath:   ".spec.spark.infrastructureSpec.containerImage.javaJars",
			desc:        "Optional. A list of Java JARS to add to the classpath. For example, gs://bucket-name/my/path/to/file.jar",
			wantVerdict: NotRepresentable,
			wantReason:  "gcs-scheme-not-a-gcp-resource-name",
		},
		{
			name:        "a bucket field that mentions gs:// stays a reference",
			fieldPath:   ".spec.loggingSettings.audioRecordingConfig.gcsBucket",
			desc:        "Optional. The Cloud Storage bucket to store the session audio recordings. The URI must start with \"gs://\".",
			wantVerdict: IsReference,
		},
		{
			name:        "a resource-name template wins over gs://",
			fieldPath:   ".spec.source",
			desc:        "Either projects/{project}/locations/{location}/datasets/{dataset} or a gs:// path.",
			wantVerdict: IsReference,
		},
		{
			name:        "a service account that mentions gs:// stays a reference",
			fieldPath:   ".spec.serviceAccount",
			desc:        "The service account that reads gs://my-bucket.",
			wantVerdict: IsReference,
		},
		{
			name:        "a gs:// pattern still matches a set of objects",
			fieldPath:   ".spec.gcsFilesetSpec.filePatterns",
			desc:        "Patterns to identify a set of files in Google Cloud Storage. See Wildcard Names. Example: gs://bucket_name/*",
			wantVerdict: NotAReference,
		},
		{
			name:        "bq:// comes before gs://",
			fieldPath:   ".spec.inputSource",
			desc:        "A gs:// path or a bq:// table.",
			wantVerdict: NotRepresentable,
			wantReason:  "bq-scheme-not-a-gcp-resource-name",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			verdict, reason := Classify(tc.fieldPath, tc.desc)

			// Assert
			if verdict != tc.wantVerdict || reason != tc.wantReason {
				t.Errorf("Classify(%q) = %v, %q, want %v, %q",
					tc.fieldPath, verdict, reason, tc.wantVerdict, tc.wantReason)
			}
		})
	}
}

// TestReferenceRules checks the phrases the generator puts in the detail of a
// possible-reference-by-description entry. The template phrase quotes the
// whole word that holds the template, without the backticks, quotes,
// parentheses or punctuation around it. When several rules match, all of
// them are listed.
func TestReferenceRules(t *testing.T) {
	for _, tc := range []struct {
		name      string
		fieldPath string
		desc      string
		want      []string
	}{
		{
			name:      "a brace template in backticks",
			fieldPath: ".spec.ragResources[].ragCorpus",
			desc:      "Optional. RagCorpora resource name. Format: `projects/{project}/locations/{location}/ragCorpora/{rag_corpus}`",
			want:      []string{`the description has the resource-name template "projects/{project}/locations/{location}/ragCorpora/{rag_corpus}"`},
		},
		{
			name:      "an angle-bracket template",
			fieldPath: ".spec.notificationChannels[]",
			desc:      "Must be of the format `projects/<project_id_or_number>/notificationChannels/<channel_id>`",
			want:      []string{`the description has the resource-name template "projects/<project_id_or_number>/notificationChannels/<channel_id>"`},
		},
		{
			name:      "a template in parentheses at the end of a sentence",
			fieldPath: ".spec.network",
			desc:      "The network to join (projects/{project}/global/networks/{network}).",
			want:      []string{`the description has the resource-name template "projects/{project}/global/networks/{network}"`},
		},
		{
			name:      "the whole word is quoted when a later segment matched",
			fieldPath: ".spec.policy",
			desc:      "Format: organizations/{organization}/locations/{location}/policies/{policy}",
			want:      []string{`the description has the resource-name template "organizations/{organization}/locations/{location}/policies/{policy}"`},
		},
		{
			name:      "projects/ after a space needs no placeholder",
			fieldPath: ".spec.topic",
			desc:      "The topic, such as projects/my-project/topics/my-topic.",
			want:      []string{`the description has the resource-name template "projects/my-project/topics/my-topic"`},
		},
		{
			name:      "a service account",
			fieldPath: ".spec.runAsServiceAccount",
			desc:      "Email of the service account to run as.",
			want:      []string{"the field name ends in serviceAccount"},
		},
		{
			name:      "a Cloud Storage bucket",
			fieldPath: ".spec.artifactsGCSBucket",
			desc:      "The Cloud Storage bucket that holds the artifacts.",
			want:      []string{"the field name has bucket and the description mentions Cloud Storage"},
		},
		{
			name:      "every matching rule is listed",
			fieldPath: ".spec.serviceAccount",
			desc:      "Format: projects/{project}/serviceAccounts/{email}",
			want: []string{
				`the description has the resource-name template "projects/{project}/serviceAccounts/{email}"`,
				"the field name ends in serviceAccount",
			},
		},
		{
			name:      "no rule",
			fieldPath: ".spec.displayName",
			desc:      "A name people can read.",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := ReferenceRules(tc.fieldPath, tc.desc)

			// Assert
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ReferenceRules(%q) mismatch (-want +got):\n%s", tc.fieldPath, diff)
			}
		})
	}
}

// TestIsReferenceFieldPath pins which paths Classify's callers skip as already
// being references. A path it misses gets classified, and the fields inside a
// reference would then show up in missingrefs.txt.
func TestIsReferenceFieldPath(t *testing.T) {
	for _, tc := range []struct {
		fieldPath string
		want      bool
	}{
		{".spec.networkRef", true},
		{".spec.subnetworkRefs", true},
		{".spec.subnetworkRefs[]", true},
		{".spec.networkRef.external", true},
		{".spec.subnetworkRefs[].external", true},
		{".spec.networkRef.name", true},
		{".spec.networkRef.namespace", false},
		{".spec.network", false},
		{".spec.referrer", false},
	} {
		t.Run(tc.fieldPath, func(t *testing.T) {
			// Act
			got := IsReferenceFieldPath(tc.fieldPath)

			// Assert
			if got != tc.want {
				t.Errorf("IsReferenceFieldPath(%q) = %v, want %v", tc.fieldPath, got, tc.want)
			}
		})
	}
}

// TestMatchName pins the name rules. The generator files what they match as
// possible-reference-by-name, naming the target type, so a rule that widens
// puts questions about ordinary fields in front of a reviewer.
func TestMatchName(t *testing.T) {
	for _, tc := range []struct {
		fieldPath  string
		wantTarget string
	}{
		{".spec.userTokenSecretVersion", "SecretManagerSecretVersionRef"},
		{".spec.network", "ComputeNetworkRef"},
		{".spec.peering.vpc", "ComputeNetworkRef"},
		{".spec.kmsKeyName", "KMSCryptoKeyRef"},
		{".spec.clientSecret", "SecretManagerSecretVersionRef"},
		{".spec.project", "ProjectRef"},
		{".spec.hosts[]", ""},
		{".spec.networkConfig", ""},
		{".spec.projectNumber", ""},
	} {
		t.Run(tc.fieldPath, func(t *testing.T) {
			// Act
			got, ok := MatchName(tc.fieldPath)

			// Assert
			if got != tc.wantTarget || ok != (tc.wantTarget != "") {
				t.Errorf("MatchName(%q) = %q, %v, want %q", tc.fieldPath, got, ok, tc.wantTarget)
			}
		})
	}
}

// TestMatchDescriptionLoose pins the prose forms of "this is a resource name".
// The generator files a match as possible-reference-by-description-loose.
func TestMatchDescriptionLoose(t *testing.T) {
	for _, tc := range []struct {
		name      string
		fieldPath string
		desc      string
		want      bool
	}{
		{"resource name of", ".spec.destination", "The resource name (URI) of the destination connection profile.", true},
		{"resource URL", ".spec.securityService", "The resource URL for the network edge security service.", true},
		{"square-bracket template", ".spec.dataset", "Its format is projects/[project_id]/datasets/[bigquery_dataset_id].", true},
		{"a sub-message's own name", ".spec.workspace.name", "The resource name of the conversion workspace.", false},
		{"a catalogue entry", ".spec.machineTypeURI", "The resource URL for the machine type.", false},
		{"ordinary prose", ".spec.displayName", "A name for people to read.", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := MatchDescriptionLoose(tc.fieldPath, tc.desc)

			// Assert
			if got != tc.want {
				t.Errorf("MatchDescriptionLoose(%q, %q) = %v, want %v", tc.fieldPath, tc.desc, got, tc.want)
			}
		})
	}
}
