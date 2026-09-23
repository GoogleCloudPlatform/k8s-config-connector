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

package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestParentRef pins when a resource gets a parent reference field and when it
// gets only a queue entry. The cases that matter are the ones where the
// generator must decline: no reference type to point at, or several that match
// equally well. TestMissingRefs reads the entry, so a field without one is an
// unflagged guess.
func TestParentRef(t *testing.T) {
	grid := []struct {
		name       string
		pattern    string
		refSource  string
		wantField  string
		wantReason string
		wantPath   string
	}{
		{
			name:       "a reference type in the service package is used",
			pattern:    "projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}",
			refSource:  "package v1alpha1\n\ntype ClusterRef struct{}\n",
			wantField:  "ClusterRef *ClusterRef `json:\"clusterRef,omitempty\"`",
			wantReason: "parent-ref-guessed",
			wantPath:   "projects/{project}/locations/{location}/clusters/{cluster}",
		},
		{
			// PrivateCACertificate: the type carries the Kind's prefix while
			// upstream still calls the field caPoolRef.
			name:       "a kind-prefixed reference type still matches",
			pattern:    "projects/{project}/locations/{location}/caPools/{ca_pool}/certificates/{certificate}",
			refSource:  "package v1alpha1\n\ntype PrivateCACAPoolRef struct{}\n",
			wantField:  "CaPoolRef *PrivateCACAPoolRef `json:\"caPoolRef,omitempty\"`",
			wantReason: "parent-ref-guessed",
			wantPath:   "projects/{project}/locations/{location}/caPools/{ca_pool}",
		},
		{
			// apis/kms/v1beta1 declares both, and counting the unexported one
			// would make this look ambiguous and emit nothing.
			name:       "an unexported type beside the real one is ignored",
			pattern:    "projects/{project}/locations/{location}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{version}",
			refSource:  "package v1alpha1\n\ntype KMSCryptoKeyRef struct{}\n\ntype kmsCryptoKeyRef struct{}\n",
			wantField:  "CryptoKeyRef *KMSCryptoKeyRef `json:\"cryptoKeyRef,omitempty\"`",
			wantReason: "parent-ref-guessed",
			wantPath:   "projects/{project}/locations/{location}/cryptoKeys/{crypto_key}",
		},
		{
			// FirestoreIndex: no CollectionGroupRef exists anywhere.
			name:       "no reference type means no field",
			pattern:    "projects/{project}/databases/{database}/collectionGroups/{collection_group}/indexes/{index}",
			wantField:  "",
			refSource:  "package v1alpha1\n",
			wantReason: "parent-ref-not-modelled",
			wantPath:   "projects/{project}/databases/{database}/collectionGroups/{collection_group}",
		},
		{
			// DiscoveryEngineServingConfig: both really exist upstream.
			name:       "several matching types mean no field",
			pattern:    "projects/{project}/locations/{location}/engines/{engine}/servingConfigs/{config}",
			wantField:  "",
			refSource:  "package v1alpha1\n\ntype DiscoveryEngineEngineRef struct{}\n\ntype DiscoveryEngineSearchEngineRef struct{}\n",
			wantReason: "parent-ref-not-modelled",
			wantPath:   "projects/{project}/locations/{location}/engines/{engine}",
		},
		{
			name:      "a project and location parent is already carried",
			pattern:   "projects/{project}/locations/{location}/foos/{foo}",
			wantField: "",
			refSource: "package v1alpha1\n",
		},
		{
			// rootRef names it, and a second field would not compile.
			name:      "a parent that is the root of the name is left to rootRef",
			pattern:   "properties/{property}/audiences/{audience}",
			wantField: "",
			refSource: "package v1alpha1\n\ntype PropertyRef struct{}\n",
		},
		{
			name:      "an organization parent is left to rootRef",
			pattern:   "organizations/{organization}/policies/{policy}",
			wantField: "",
			refSource: "package v1alpha1\n\ntype OrganizationRef struct{}\n",
		},
		{
			name:      "a resource with no pattern has no parent to name",
			pattern:   "",
			wantField: "",
			refSource: "package v1alpha1\n",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			// Arrange
			dir := t.TempDir()
			scaffolder := &APIScaffolder{BaseDir: dir, GoPackage: "svc/v1alpha1"}
			pkgDir := filepath.Join(dir, scaffolder.GoPackage)
			if err := os.MkdirAll(pkgDir, 0o755); err != nil {
				t.Fatalf("creating the service package: %v", err)
			}
			if err := os.WriteFile(filepath.Join(pkgDir, "refs.go"), []byte(g.refSource), 0o644); err != nil {
				t.Fatalf("writing the reference types: %v", err)
			}

			// Act
			field, item := scaffolder.parentRef(g.pattern)

			// Assert
			if g.wantField == "" && field != "" {
				t.Errorf("emitted a field where none was wanted:\n%s", field)
			}
			if g.wantField != "" && !strings.Contains(field, g.wantField) {
				t.Errorf("field = %q, want it to contain %q", field, g.wantField)
			}
			if g.wantField != "" && !strings.Contains(field, "+kcc:guess") {
				t.Errorf("an emitted field must carry the guess marker, got:\n%s", field)
			}
			switch {
			case g.wantReason == "" && item != nil:
				t.Errorf("queued %q where nothing was wanted", item.Reason)
			case g.wantReason != "" && item == nil:
				t.Errorf("queued nothing, want reason %q", g.wantReason)
			case g.wantReason != "" && item.Reason != g.wantReason:
				t.Errorf("reason = %q, want %q", item.Reason, g.wantReason)
			}
			if g.wantPath != "" && item != nil && !strings.Contains(item.Detail, g.wantPath) {
				t.Errorf("the queue entry must name the parent path %q, got: %s", g.wantPath, item.Detail)
			}
		})
	}
}

// TestRootRef pins which Spec field names the root of a resource's name. A
// resource under an organization, a folder or an Analytics property has no
// project, so a projectRef there is a field its API does not accept, and the
// judgement queue is where an unmodelled root has to show up instead.
func TestRootRef(t *testing.T) {
	grid := []struct {
		name       string
		pattern    string
		sharedRefs string
		wantField  string
		wantReason string
	}{
		{
			name:      "a project",
			pattern:   "projects/{project}/locations/{location}/widgets/{widget}",
			wantField: "ProjectRef *refsv1beta1.ProjectRef `json:\"projectRef\"`",
		},
		{
			name:      "no pattern keeps projectRef",
			pattern:   "",
			wantField: "ProjectRef *refsv1beta1.ProjectRef `json:\"projectRef\"`",
		},
		{
			name:      "an organization",
			pattern:   "organizations/{organization}/policies/{policy}",
			wantField: "OrganizationRef *refsv1beta1.OrganizationRef `json:\"organizationRef\"`",
		},
		{
			name:      "an organization with a location",
			pattern:   "organizations/{organization}/locations/{location}/postures/{posture}",
			wantField: "OrganizationRef *refsv1beta1.OrganizationRef `json:\"organizationRef\"`",
		},
		{
			name:      "a folder",
			pattern:   "folders/{folder}/locations/{location}/settings",
			wantField: "FolderRef *refsv1beta1.FolderRef `json:\"folderRef\"`",
		},
		{
			name:       "another root with a shared reference type",
			pattern:    "billingAccounts/{billing_account}/budgets/{budget}",
			sharedRefs: "package v1beta1\n\ntype BillingAccountRef struct{}\n",
			wantField:  "BillingAccountRef *refsv1beta1.BillingAccountRef `json:\"billingAccountRef,omitempty\"`",
			wantReason: "root-ref-guessed",
		},
		{
			name:       "another root with no reference type",
			pattern:    "properties/{property}/audiences/{audience}",
			wantField:  "",
			wantReason: "root-ref-not-modelled",
		},
		{
			name:      "a resource that is itself a root",
			pattern:   "billingAccounts/{billing_account}",
			wantField: "",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			// Arrange
			repo := t.TempDir()
			scaffolder := &APIScaffolder{BaseDir: filepath.Join(repo, "apis"), GoPackage: "svc/v1alpha1"}
			sharedDir := filepath.Join(repo, sharedRefsPackage)
			if err := os.MkdirAll(sharedDir, 0o755); err != nil {
				t.Fatalf("creating the shared refs package: %v", err)
			}
			if g.sharedRefs != "" {
				if err := os.WriteFile(filepath.Join(sharedDir, "refs.go"), []byte(g.sharedRefs), 0o644); err != nil {
					t.Fatalf("writing the shared reference types: %v", err)
				}
			}

			// Act
			field, item := scaffolder.rootRef(g.pattern)

			// Assert
			if g.wantField == "" && field != "" {
				t.Errorf("emitted a field where none was wanted:\n%s", field)
			}
			if g.wantField != "" && !strings.Contains(field, g.wantField) {
				t.Errorf("field = %q, want it to contain %q", field, g.wantField)
			}
			switch {
			case g.wantReason == "" && item != nil:
				t.Errorf("queued %q where nothing was wanted", item.Reason)
			case g.wantReason != "" && item == nil:
				t.Errorf("queued nothing, want reason %q", g.wantReason)
			case g.wantReason != "" && item.Reason != g.wantReason:
				t.Errorf("reason = %q, want %q", item.Reason, g.wantReason)
			}
		})
	}
}

// TestLocationRef pins when a scaffolded Spec names a location, and that every
// location it writes reaches the judgement queue. A location on a resource
// whose name has none is a field its API rejects, and a reference may replace
// any location the generator does write.
func TestLocationRef(t *testing.T) {
	const field = "Location string `json:\"location\"`"
	grid := []struct {
		name       string
		pattern    string
		wantField  string
		wantReason string
		wantDetail string
	}{
		{
			name:       "the direct parent is a project location",
			pattern:    "projects/{project}/locations/{location}/widgets/{widget}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "parent.ProjectAndLocationRef",
		},
		{
			name:       "a singleton under a location",
			pattern:    "projects/{project}/locations/{location}/settings",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "parent.ProjectAndLocationRef",
		},
		{
			name:       "the direct parent is an organization location",
			pattern:    "organizations/{organization}/locations/{location}/postures/{posture}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "the parent is organizations/{organization}/locations/{location}, and no shared reference type",
		},
		{
			name:       "a location above the direct parent",
			pattern:    "projects/{project}/locations/{location}/clusters/{cluster}/instances/{instance}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "the parent is projects/{project}/locations/{location}/clusters/{cluster}, which already names",
		},
		{
			// Dataplex names a resource collection zones; the location comes first.
			name:       "a zones collection below a location",
			pattern:    "projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/assets/{asset}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "the parent is projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}, which",
		},
		{
			// Location is the canonical name, and a region is a location.
			name:       "a region is still called location",
			pattern:    "projects/{project}/regions/{region}/widgets/{widget}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "the parent is projects/{project}/regions/{region}, and no shared reference type",
		},
		{
			name:       "a zone is still called location",
			pattern:    "projects/{project}/zones/{zone}/widgets/{widget}",
			wantField:  field,
			wantReason: "location-or-parent-ref",
			wantDetail: "the parent is projects/{project}/zones/{zone}, and no shared reference type",
		},
		{
			name:       "a project with no location",
			pattern:    "projects/{project}/topics/{topic}",
			wantField:  "",
			wantReason: "",
		},
		{
			name:       "an organization with no location",
			pattern:    "organizations/{organization}/policies/{policy}",
			wantField:  "",
			wantReason: "",
		},
		{
			name:       "a location that is the resource itself",
			pattern:    "projects/{project}/locations/{location}",
			wantField:  "",
			wantReason: "",
		},
		{
			name:       "a fixed location",
			pattern:    "projects/{project}/locations/global/widgets/{widget}",
			wantField:  "",
			wantReason: "",
		},
		{
			name:       "no pattern keeps the location",
			pattern:    "",
			wantField:  field,
			wantReason: "location-parent-unknown",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			// Arrange
			scaffolder := &APIScaffolder{BaseDir: t.TempDir(), GoPackage: "svc/v1alpha1"}

			// Act
			gotField, item := scaffolder.locationRef(g.pattern)

			// Assert
			if g.wantField == "" && gotField != "" {
				t.Errorf("emitted a field where none was wanted:\n%s", gotField)
			}
			if g.wantField != "" && !strings.Contains(gotField, g.wantField) {
				t.Errorf("field = %q, want it to contain %q", gotField, g.wantField)
			}
			if strings.Contains(gotField, "+kcc:guess") {
				t.Errorf("the location is read from the name, so it must not be marked a guess:\n%s", gotField)
			}
			switch {
			case g.wantReason == "" && item != nil:
				t.Errorf("queued %q where nothing was wanted", item.Reason)
			case g.wantReason != "" && item == nil:
				t.Errorf("queued nothing, want reason %q", g.wantReason)
			case g.wantReason != "" && item.Reason != g.wantReason:
				t.Errorf("reason = %q, want %q", item.Reason, g.wantReason)
			case g.wantDetail != "" && !strings.Contains(item.Detail, g.wantDetail):
				t.Errorf("detail = %q, want it to contain %q", item.Detail, g.wantDetail)
			}
		})
	}
}

// TestRepoRoot pins where the scaffolder looks for apis/refs/v1beta1. With the
// wrong root, parentRef and rootRef never see the shared reference types and
// queue every such parent as having none.
func TestRepoRoot(t *testing.T) {
	for _, tc := range []struct {
		baseDir string
		want    string
	}{
		// generate-types' default, from GenerateCRDOptions.InitDefaults.
		{"/repo/apis/", "/repo"},
		{"/repo/apis", "/repo"},
		{"/tmp/out", "/tmp/out"},
		{"/tmp/out/", "/tmp/out"},
	} {
		t.Run(tc.baseDir, func(t *testing.T) {
			// Arrange
			scaffolder := &APIScaffolder{BaseDir: tc.baseDir}

			// Act
			got := scaffolder.repoRoot()

			// Assert
			if got != tc.want {
				t.Errorf("repoRoot() with BaseDir %q = %q, want %q", tc.baseDir, got, tc.want)
			}
		})
	}
}
