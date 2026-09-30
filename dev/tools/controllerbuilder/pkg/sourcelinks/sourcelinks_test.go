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

package sourcelinks

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
)

const testSHA = "1765b559c42386788ff0c6412491277b4791107a"

// fakeWeb answers every request in-process, so no test touches the network.
// Pages maps a full URL to a status, or to "->"+URL for a redirect. Anything
// missing is a 404.
type fakeWeb struct {
	pages     map[string]string
	discovery string
	offline   bool
}

func (w *fakeWeb) RoundTrip(req *http.Request) (*http.Response, error) {
	if w.offline {
		return nil, errors.New("offline")
	}
	rec := httptest.NewRecorder()
	u := req.URL.String()
	switch page, ok := w.pages[u]; {
	case u == DiscoveryDirectoryURL:
		rec.WriteHeader(http.StatusOK)
		rec.WriteString(w.discovery)
	case ok && strings.HasPrefix(page, "->"):
		rec.Header().Set("Location", strings.TrimPrefix(page, "->"))
		rec.WriteHeader(http.StatusMovedPermanently)
	case ok:
		rec.WriteHeader(http.StatusOK)
	default:
		rec.WriteHeader(http.StatusNotFound)
	}
	resp := rec.Result()
	resp.Request = req
	return resp, nil
}

const testDiscovery = `{"items": [
  {"name": "bigtableadmin", "version": "v2", "preferred": true, "documentationLink": "https://cloud.google.com/bigtable/"},
  {"name": "dataproc", "version": "v1", "preferred": true, "documentationLink": "https://cloud.google.com/dataproc/"},
  {"name": "container", "version": "v1", "preferred": true, "documentationLink": "https://cloud.google.com/kubernetes-engine/docs/"},
  {"name": "container", "version": "v1beta1", "preferred": false, "documentationLink": "https://cloud.google.com/kubernetes-engine/docs/"},
  {"name": "privateca", "version": "v1", "preferred": true, "documentationLink": "https://cloud.google.com/"},
  {"name": "storage", "version": "v1", "preferred": true, "documentationLink": "https://developers.google.com/storage/docs/json_api/"}
]}`

func newTestResolver(t *testing.T, web *fakeWeb) *Resolver {
	t.Helper()
	if web.discovery == "" {
		web.discovery = testDiscovery
	}
	return &Resolver{
		Client:        &http.Client{Transport: web},
		GoogleapisSHA: testSHA,
	}
}

func TestLinks(t *testing.T) {
	grid := []struct {
		name        string
		in          input
		pages       map[string]string
		offline     bool
		cfgYAML     string
		wantService Link
		wantRes     Link
	}{
		{
			name: "derived from Discovery and verified",
			in:   input{protoPath: "google/container/v1/cluster_service.proto", version: "v1", api: "container", rpath: "projects.locations.clusters"},
			pages: map[string]string{
				"https://cloud.google.com/kubernetes-engine/docs/":                                                   "200",
				"https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/kubernetes-engine/docs/"},
			wantRes:     Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters"},
		},
		{
			name: "override for an irregular API",
			in:   input{protoPath: "google/bigtable/admin/v2/instance.proto", version: "v2", api: "bigtableadmin", rpath: "projects.instances"},
			pages: map[string]string{
				"https://cloud.google.com/bigtable/":                                                     "200",
				"https://docs.cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/bigtable/"},
			wantRes:     Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances"},
		},
		{
			name: "renamed product still verifies through the redirect",
			in:   input{protoPath: "google/cloud/dataproc/v1/clusters.proto", version: "v1", api: "dataproc", rpath: "projects.regions.clusters"},
			pages: map[string]string{
				"https://cloud.google.com/dataproc/":                                                           "200",
				"https://docs.cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.clusters":      "->https://docs.cloud.google.com/managed-spark/docs/reference/rest/v1/projects.regions.clusters",
				"https://docs.cloud.google.com/managed-spark/docs/reference/rest/v1/projects.regions.clusters": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/dataproc/"},
			wantRes:     Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.clusters"},
		},
		{
			name: "redirect to a general page does not verify",
			in:   input{protoPath: "google/container/v1/cluster_service.proto", version: "v1", api: "container", rpath: "projects.locations.clusters"},
			pages: map[string]string{
				"https://cloud.google.com/kubernetes-engine/docs/":                                                   "200",
				"https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters": "->https://docs.cloud.google.com/kubernetes-engine/docs/apis",
				"https://docs.cloud.google.com/kubernetes-engine/docs/apis":                                          "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/kubernetes-engine/docs/"},
			wantRes: Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters",
				Guess: judgement.ReasonVerifyResourceDocsLink},
		},
		{
			name: "404 is a guess and bare landing page is flagged",
			in:   input{protoPath: "google/cloud/security/privateca/v1/resources.proto", version: "v1", api: "privateca", rpath: "projects.locations.caPools"},
			pages: map[string]string{
				"https://cloud.google.com/": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/", Guess: judgement.ReasonVerifyServiceDocsLink},
			wantRes: Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/certificate-authority-service/docs/reference/rest/v1/projects.locations.caPools",
				Guess: judgement.ReasonVerifyResourceDocsLink},
		},
		{
			name: "no pattern falls back to the API's reference root",
			in:   input{protoPath: "google/container/v1/cluster_service.proto", version: "v1", api: "container"},
			pages: map[string]string{
				"https://cloud.google.com/kubernetes-engine/docs/": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/kubernetes-engine/docs/"},
			wantRes: Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest",
				Guess: judgement.ReasonVerifyResourceDocsLink},
		},
		{
			name: "service config proto reference is tried before Discovery",
			in:   input{protoPath: "google/cloud/notebooks/v2/instance.proto", version: "v2", api: "notebooks", rpath: "projects.locations.instances"},
			cfgYAML: `type: google.api.Service
publishing:
  documentation_uri: https://cloud.google.com/vertex-ai/docs/workbench/
  proto_reference_documentation_uri: https://cloud.google.com/vertex-ai/docs/workbench/reference/rpc
`,
			pages: map[string]string{
				"https://cloud.google.com/vertex-ai/docs/workbench/":                                                    "200",
				"https://docs.cloud.google.com/vertex-ai/docs/workbench/reference/rest/v2/projects.locations.instances": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/vertex-ai/docs/workbench/"},
			wantRes:     Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/vertex-ai/docs/workbench/reference/rest/v2/projects.locations.instances"},
		},
		{
			name: "service config proto reference that is already the REST reference",
			in:   input{protoPath: "google/cloud/notebooks/v2/instance.proto", version: "v2", api: "notebooks", rpath: "projects.locations.instances"},
			cfgYAML: `type: google.api.Service
publishing:
  documentation_uri: https://cloud.google.com/vertex-ai/docs/workbench/instances/introduction
  proto_reference_documentation_uri: https://cloud.google.com/vertex-ai/docs/workbench/reference/rest
`,
			pages: map[string]string{
				"https://cloud.google.com/vertex-ai/docs/workbench/instances/introduction":                              "200",
				"https://docs.cloud.google.com/vertex-ai/docs/workbench/reference/rest/v2/projects.locations.instances": "200",
			},
			wantService: Link{Key: KeyServiceDocs, URL: "https://cloud.google.com/vertex-ai/docs/workbench/instances/introduction"},
			wantRes:     Link{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/vertex-ai/docs/workbench/reference/rest/v2/projects.locations.instances"},
		},
		{
			name:        "offline writes links as guesses",
			in:          input{protoPath: "google/container/v1/cluster_service.proto", version: "v1", api: "container", rpath: "projects.locations.clusters"},
			offline:     true,
			wantService: Link{Key: KeyServiceDocs, Guess: judgement.ReasonVerifyServiceDocsLink},
			wantRes:     Link{Key: KeyResourceDocs, Guess: judgement.ReasonVerifyResourceDocsLink},
		},
	}
	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			r := newTestResolver(t, &fakeWeb{pages: g.pages, offline: g.offline})
			if g.cfgYAML != "" {
				r.GoogleapisDir = t.TempDir()
				dir := filepath.Join(r.GoogleapisDir, filepath.Dir(g.in.protoPath))
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(dir, "service_v2.yaml"), []byte(g.cfgYAML), 0644); err != nil {
					t.Fatal(err)
				}
			}
			got := r.links(context.Background(), g.in)
			if len(got) != 3 {
				t.Fatalf("got %d links, want 3", len(got))
			}
			wantProto := Link{Key: KeyProto, URL: "https://github.com/googleapis/googleapis/blob/" + testSHA + "/" + g.in.protoPath}
			if got[0] != wantProto {
				t.Errorf("proto link = %+v, want %+v", got[0], wantProto)
			}
			for i, want := range []Link{g.wantService, g.wantRes} {
				l := got[i+1]
				if (l.Guess == "") != (l.Detail == "") {
					t.Errorf("%s: a guess needs a detail and only a guess has one: %+v", l.Key, l)
				}
				l.Detail = ""
				if l != want {
					t.Errorf("%s link = %+v, want %+v", want.Key, l, want)
				}
			}
		})
	}
}

func TestProtoLinkForKCCOverride(t *testing.T) {
	r := &Resolver{GoogleapisSHA: testSHA}
	got := r.protoLink("google/cloud/config/v1/config.proto")
	want := "https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/config/v1/config.proto"
	if got.URL != want {
		t.Errorf("got %s, want %s", got.URL, want)
	}
}

// TestRestRootsTriesOverridesLast checks that derived roots come first, so an
// override is only used when nothing derived verifies.
func TestRestRootsTriesOverridesLast(t *testing.T) {
	got := restRoots("networkservices", "projects.locations.meshes", "https://cloud.google.com/networking/", serviceConfig{})
	want := []root{
		{url: "https://docs.cloud.google.com/networking/docs/reference/rest"},
		{url: "https://docs.cloud.google.com/service-mesh/docs/reference/network-services/rest", override: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestRestRootFromDocumentationLink(t *testing.T) {
	for link, want := range map[string]string{
		"https://cloud.google.com/kubernetes-engine/docs/":             "kubernetes-engine/docs/reference/rest",
		"https://cloud.google.com/bigtable/":                           "bigtable/docs/reference/rest",
		"https://developers.google.com/compute/docs/reference/latest/": "compute/docs/reference/rest",
		"https://cloud.google.com/generative-ai-app-builder/docs/":     "generative-ai-app-builder/docs/reference/rest",
		"https://firebase.google.com/docs/hosting/":                    "",
		"https://cloud.google.com/":                                    "",
	} {
		if got := restRootFromDocumentationLink(link); got != want {
			t.Errorf("%s: got %q, want %q", link, got, want)
		}
	}
}

func TestRestPathAndVersion(t *testing.T) {
	if got := restPath("projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}"); got != "projects.locations.collections.dataStores" {
		t.Errorf("restPath = %q", got)
	}
	if got := restPath(""); got != "" {
		t.Errorf("restPath of empty = %q", got)
	}
	for pkg, want := range map[string]string{
		"google.cloud.discoveryengine.v1alpha": "v1alpha",
		"google.bigtable.admin.v2":             "v2",
		"google.cloud.notebooks.v1beta1":       "v1beta1",
		"google.storage.control":               "",
	} {
		if got := apiVersion(pkg); got != want {
			t.Errorf("apiVersion(%s) = %q, want %q", pkg, got, want)
		}
	}
}

// TestResolveReadsDescriptor checks the descriptor plumbing: the proto path,
// the version from the package and the API name from default_host.
func TestResolveReadsDescriptor(t *testing.T) {
	svcOpts := &descriptorpb.ServiceOptions{}
	proto.SetExtension(svcOpts, annotations.E_DefaultHost, "container.googleapis.com")
	fdp := &descriptorpb.FileDescriptorProto{
		Name:        proto.String("google/container/v1/cluster_service.proto"),
		Package:     proto.String("google.container.v1"),
		Syntax:      proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{Name: proto.String("Cluster")}},
		Service:     []*descriptorpb.ServiceDescriptorProto{{Name: proto.String("ClusterManager"), Options: svcOpts}},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newTestResolver(t, &fakeWeb{pages: map[string]string{
		"https://cloud.google.com/kubernetes-engine/docs/":                                                   "200",
		"https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters": "200",
	}})
	got := r.Resolve(context.Background(), Resource{
		Message: fd.Messages().ByName("Cluster"),
		Pattern: "projects/{project}/locations/{location}/clusters/{cluster}",
	})
	for _, l := range got {
		if l.Guess != "" {
			t.Errorf("%s was not verified: %s", l.Key, l.Detail)
		}
	}
	if want := "https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters"; got[2].URL != want {
		t.Errorf("resource-docs = %s, want %s", got[2].URL, want)
	}
}

func TestRenderAndParse(t *testing.T) {
	links := []Link{
		{Key: KeyProto, URL: "https://github.com/googleapis/googleapis/blob/abc/google/bigtable/admin/v2/instance.proto"},
		{Key: KeyServiceDocs, URL: "https://cloud.google.com/bigtable/"},
		{Key: KeyResourceDocs, URL: "https://docs.cloud.google.com/bigtable/docs/reference/rest/v2/projects.instances", Guess: judgement.ReasonVerifyResourceDocsLink, Detail: "ignored"},
	}
	block := Render("BigtableInstance", links)
	want := `// API sources for BigtableInstance, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/abc/google/bigtable/admin/v2/instance.proto
// +kcc:source:service-docs=https://cloud.google.com/bigtable/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/bigtable/docs/reference/rest/v2/projects.instances`
	if block != want {
		t.Errorf("Render =\n%s\nwant\n%s", block, want)
	}

	src := "// Copyright\n\n" + block + "\n\npackage v1alpha1\n\n// +kcc:source:proto=after-package-is-ignored\n"
	h := Parse([]byte(src))
	if h.Kind != "BigtableInstance" {
		t.Errorf("Kind = %q", h.Kind)
	}
	links[2].Detail = ""
	if !reflect.DeepEqual(h.Links, links) {
		t.Errorf("Parse links = %+v, want %+v", h.Links, links)
	}
	if got := h.Guesses(); !reflect.DeepEqual(got, []string{judgement.ReasonVerifyResourceDocsLink}) {
		t.Errorf("Guesses = %q", got)
	}
}
