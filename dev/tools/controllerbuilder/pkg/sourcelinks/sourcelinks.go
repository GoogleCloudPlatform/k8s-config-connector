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

// Package sourcelinks finds the links a reviewer needs for a generated
// resource: the proto it came from, the service's docs, and the resource's
// REST reference page.
//
// The proto link is always right. The two docs links are guesses, so they are
// checked online. A link that does not check out is still returned, marked
// with the judgement queue reason a person should work.
package sourcelinks

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
)

// The keys of the +kcc:source:<key>=<url> markers.
const (
	KeyProto        = "proto"
	KeyServiceDocs  = "service-docs"
	KeyResourceDocs = "resource-docs"
)

// GuessKind is the <what> in the "+kcc:guess=<what> reason=<why>" line written
// above a link that was not verified.
const GuessKind = "source-link"

// DiscoveryDirectoryURL lists every API Google publishes, with a
// documentationLink for each.
const DiscoveryDirectoryURL = "https://www.googleapis.com/discovery/v1/apis"

// docsHost serves real 404s for missing pages. cloud.google.com answers 200
// for any path, so a check against it proves nothing.
const docsHost = "docs.cloud.google.com"

// Link is one marker for a generated types file.
type Link struct {
	Key string
	URL string
	// Guess is the queue reason when the link was not verified. It is empty
	// for a verified link and for the proto link.
	Guess string
	// Detail says what was tried, for the queue entry.
	Detail string
}

// Resource is what Resolve needs to know about one generated resource.
type Resource struct {
	Message protoreflect.MessageDescriptor
	// Pattern is the primary google.api.resource pattern. It may be empty.
	Pattern string
	// Files is used to find the service's default_host when the message's own
	// file declares no service. It may be nil.
	Files *protoregistry.Files
}

// Resolver finds and checks links. It is safe to reuse for every resource in
// a run; the Discovery directory is fetched once.
type Resolver struct {
	Client *http.Client
	// GoogleapisSHA is the googleapis commit the protos were compiled from.
	GoogleapisSHA string
	// GoogleapisDir is a local googleapis checkout, used to read service
	// config YAMLs. Optional.
	GoogleapisDir string
	// DiscoveryURL defaults to DiscoveryDirectoryURL.
	DiscoveryURL string

	once       sync.Once
	docLinks   map[string]string // "name:version" -> documentationLink
	preferred  map[string]string // name -> documentationLink of the preferred version
	discoveryE error
}

// NewResolver reads the pinned googleapis commit from apis/git.versions and
// uses the googleapis checkout that generate-proto.sh leaves in .build.
func NewResolver(repoRoot string) (*Resolver, error) {
	sha, err := googleapisSHA(filepath.Join(repoRoot, "apis", "git.versions"))
	if err != nil {
		return nil, err
	}
	return &Resolver{
		Client:        &http.Client{Timeout: 20 * time.Second},
		GoogleapisSHA: sha,
		GoogleapisDir: filepath.Join(repoRoot, ".build", "third_party", "googleapis"),
	}, nil
}

func googleapisSHA(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) == 2 && fields[0] == "https://github.com/googleapis/googleapis" {
			return fields[1], nil
		}
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("%s has no googleapis line", path)
}

// input is what the links are derived from, taken out of the descriptor so
// tests can build it directly.
type input struct {
	// protoPath is the proto file path, such as
	// google/bigtable/admin/v2/instance.proto.
	protoPath string
	// version is the API version, such as v2, or "".
	version string
	// api is the Discovery name, such as bigtableadmin, or "".
	api string
	// rpath is the REST resource path, such as projects.instances, or "".
	rpath string
}

// Resolve returns the proto, service-docs and resource-docs links, in that
// order.
func (r *Resolver) Resolve(ctx context.Context, res Resource) []Link {
	file := res.Message.ParentFile()
	return r.links(ctx, input{
		protoPath: file.Path(),
		version:   apiVersion(string(file.Package())),
		api:       apiName(defaultHost(file, res.Files)),
		rpath:     restPath(res.Pattern),
	})
}

func (r *Resolver) links(ctx context.Context, in input) []Link {
	var cfg serviceConfig
	if r.GoogleapisDir != "" {
		cfg = readServiceConfig(filepath.Join(r.GoogleapisDir, filepath.Dir(in.protoPath)))
	}
	discoveryLink := r.documentationLink(ctx, in.api, in.version)
	return []Link{
		r.protoLink(in.protoPath),
		r.serviceDocsLink(ctx, discoveryLink, cfg),
		r.resourceDocsLink(ctx, in, discoveryLink, cfg),
	}
}

// kccProtoOverrides are the protos generate-proto.sh replaces with a copy
// from mockgcp/apis before compiling. Types generated from them follow the
// KCC copy, so the link does too.
var kccProtoOverrides = map[string]bool{
	"google/cloud/config/v1/config.proto":    true,
	"google/cloud/dataplex/v1/catalog.proto": true,
}

func (r *Resolver) protoLink(path string) Link {
	if kccProtoOverrides[path] {
		return Link{Key: KeyProto, URL: "https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/" + path}
	}
	return Link{Key: KeyProto, URL: "https://github.com/googleapis/googleapis/blob/" + r.GoogleapisSHA + "/" + path}
}

func (r *Resolver) serviceDocsLink(ctx context.Context, discoveryLink string, cfg serviceConfig) Link {
	link := Link{Key: KeyServiceDocs, URL: discoveryLink}
	if link.URL == "" {
		link.URL = cfg.Publishing.DocumentationURI
	}
	switch {
	case link.URL == "":
		link.Guess = judgement.ReasonVerifyServiceDocsLink
		link.Detail = "neither the Discovery directory nor the service config names a docs page; add one"
		if r.discoveryE != nil {
			link.Detail += " (the Discovery directory could not be read: " + r.discoveryE.Error() + ")"
		}
	case isBareLandingPage(link.URL):
		link.Guess = judgement.ReasonVerifyServiceDocsLink
		link.Detail = link.URL + " is a generic landing page, not the service's docs"
	default:
		if _, ok := r.fetch(ctx, link.URL); !ok {
			link.Guess = judgement.ReasonVerifyServiceDocsLink
			link.Detail = link.URL + " did not load"
		}
	}
	return link
}

// isBareLandingPage reports a link with no path, such as
// https://cloud.google.com/, which the Discovery directory gives for some APIs.
func isBareLandingPage(link string) bool {
	u, err := url.Parse(link)
	return err == nil && strings.Trim(u.Path, "/") == ""
}

func (r *Resolver) resourceDocsLink(ctx context.Context, in input, discoveryLink string, cfg serviceConfig) Link {
	roots := restRoots(in.api, in.rpath, discoveryLink, cfg)
	link := Link{Key: KeyResourceDocs}
	if len(roots) == 0 {
		link.Guess = judgement.ReasonVerifyResourceDocsLink
		link.Detail = "found no docs root for this API; find the REST reference page by hand"
		return link
	}
	if in.version == "" || in.rpath == "" {
		link.URL = roots[0].url
		link.Guess = judgement.ReasonVerifyResourceDocsLink
		link.Detail = "the proto declares no google.api.resource pattern, so the page for this resource could not be derived; " +
			"this is a guess at the API's REST reference, find the resource in it"
		return link
	}
	tail := "/" + in.version + "/" + in.rpath
	var tried []string
	for _, root := range roots {
		candidate := root.url + tail
		tried = append(tried, candidate)
		if final, ok := r.fetch(ctx, candidate); ok && final.Host == docsHost &&
			strings.HasSuffix(strings.TrimSuffix(final.Path, "/"), tail) {
			if root.override {
				// Say so, so we learn which overrides are still needed.
				why := "no root could be derived"
				if derived := tried[:len(tried)-1]; len(derived) > 0 {
					why = "none of the derived pages loaded: " + strings.Join(derived, ", ")
				}
				klog.Infof("resource-docs for %s %s came from restRootOverrides; %s", in.api, in.rpath, why)
			}
			link.URL = candidate
			return link
		}
	}
	link.URL = tried[0]
	link.Guess = judgement.ReasonVerifyResourceDocsLink
	link.Detail = "none of these loaded as a page ending in " + tail + ": " + strings.Join(tried, ", ")
	return link
}

// restRootOverrides are docs roots for APIs whose REST reference cannot be
// derived from the service config or the Discovery directory. A key is an API
// name, or an API name and a resource path for an API documented under more
// than one product. Every entry was checked by hand.
//
// They are tried last, and using one is logged. The goal is to remove them:
// an entry whose API now verifies through a derived root is no longer needed.
var restRootOverrides = map[string]string{
	"bigtableadmin": "bigtable/docs/reference/admin/rest",
	"privateca":     "certificate-authority-service/docs/reference/rest",
	"sqladmin":      "sql/docs/mysql/admin-api/rest",
	"storage":       "storage/docs/json_api",

	"networkservices/projects.locations.lbRouteExtensions":   "service-extensions/docs/reference/rest",
	"networkservices/projects.locations.lbTrafficExtensions": "service-extensions/docs/reference/rest",
	"networkservices/projects.locations.meshes":              "service-mesh/docs/reference/network-services/rest",
}

// root is a candidate REST reference root, as
// https://docs.cloud.google.com/<path>.
type root struct {
	url string
	// override is set for a root from restRootOverrides.
	override bool
}

// restRoots returns candidate REST reference roots, best first: the ones
// derived from the service config and the Discovery directory, then the
// hand-written overrides.
func restRoots(api, rpath, discoveryLink string, cfg serviceConfig) []root {
	var out []root
	addRoot := func(p string, override bool) {
		p = strings.Trim(p, "/")
		if p == "" {
			return
		}
		u := "https://" + docsHost + "/" + p
		for _, existing := range out {
			if existing.url == u {
				return
			}
		}
		out = append(out, root{url: u, override: override})
	}
	add := func(p string) { addRoot(p, false) }

	// The service config names its reference for some APIs. It usually
	// points at the RPC reference, and the REST reference sits beside it;
	// notebooks v2 points at the REST reference directly.
	switch p := docsPath(cfg.Publishing.ProtoReferenceDocumentationURI); {
	case strings.HasSuffix(p, "/rpc"):
		add(strings.TrimSuffix(p, "/rpc") + "/rest")
	case strings.HasSuffix(p, "/rest"):
		add(p)
	}
	add(restRootFromDocumentationLink(discoveryLink))
	// The first segment of the service config's docs page names the product.
	if p := docsPath(cfg.Publishing.DocumentationURI); p != "" {
		add(strings.SplitN(p, "/", 2)[0] + "/docs/reference/rest")
	}
	// Last resort: a hand-checked override.
	addRoot(restRootOverrides[api+"/"+rpath], true)
	addRoot(restRootOverrides[api], true)
	return out
}

// restRootFromDocumentationLink turns a Discovery documentationLink such as
// https://cloud.google.com/kubernetes-engine/docs/ into
// kubernetes-engine/docs/reference/rest.
func restRootFromDocumentationLink(link string) string {
	p := docsPath(link)
	// compute's link is https://developers.google.com/compute/docs/reference/latest/.
	p = strings.TrimSuffix(p, "/reference/latest")
	if p == "" {
		return ""
	}
	if !strings.Contains("/"+p+"/", "/docs/") {
		p += "/docs"
	}
	return p + "/reference/rest"
}

// docsPath returns the path of a Google docs URL without surrounding slashes,
// or "" for a URL on another host.
func docsPath(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	switch u.Host {
	case "cloud.google.com", docsHost, "developers.google.com":
		return strings.Trim(u.Path, "/")
	}
	return ""
}

var versionRE = regexp.MustCompile(`^v\d+((alpha|beta)\d*)?$`)

// apiVersion returns the version segment of a proto package, such as v1alpha
// for google.cloud.discoveryengine.v1alpha, or "".
func apiVersion(pkg string) string {
	parts := strings.Split(pkg, ".")
	if v := parts[len(parts)-1]; versionRE.MatchString(v) {
		return v
	}
	return ""
}

// restPath turns a resource pattern into the path the REST reference uses:
// projects/{project}/locations/{location}/dataStores/{data_store} becomes
// projects.locations.dataStores.
func restPath(pattern string) string {
	var segs []string
	for _, s := range strings.Split(pattern, "/") {
		if s != "" && !strings.HasPrefix(s, "{") {
			segs = append(segs, s)
		}
	}
	return strings.Join(segs, ".")
}

// defaultHost returns the google.api.default_host of a service in the file,
// or else in another file of the same package.
func defaultHost(file protoreflect.FileDescriptor, files *protoregistry.Files) string {
	if h := fileDefaultHost(file); h != "" {
		return h
	}
	if files == nil {
		return ""
	}
	var host string
	files.RangeFilesByPackage(file.Package(), func(f protoreflect.FileDescriptor) bool {
		host = fileDefaultHost(f)
		return host == ""
	})
	return host
}

func fileDefaultHost(file protoreflect.FileDescriptor) string {
	services := file.Services()
	for i := 0; i < services.Len(); i++ {
		opts := services.Get(i).Options()
		if opts == nil {
			continue
		}
		if h, ok := proto.GetExtension(opts, annotations.E_DefaultHost).(string); ok && h != "" {
			return h
		}
	}
	return ""
}

// apiName returns the Discovery name for a host: discoveryengine for
// discoveryengine.googleapis.com.
func apiName(host string) string {
	host, _, _ = strings.Cut(host, ":")
	return strings.TrimSuffix(host, ".googleapis.com")
}

// serviceConfig is the part of a googleapis service config YAML this package
// reads.
type serviceConfig struct {
	Type       string `yaml:"type"`
	Publishing struct {
		DocumentationURI               string `yaml:"documentation_uri"`
		ProtoReferenceDocumentationURI string `yaml:"proto_reference_documentation_uri"`
	} `yaml:"publishing"`
}

// readServiceConfig returns the service config in dir, or an empty one.
func readServiceConfig(dir string) serviceConfig {
	paths, _ := filepath.Glob(filepath.Join(dir, "*.yaml"))
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		var cfg serviceConfig
		if yaml.Unmarshal(data, &cfg) == nil && cfg.Type == "google.api.Service" {
			return cfg
		}
	}
	return serviceConfig{}
}

// documentationLink returns the Discovery documentationLink for name:version,
// falling back to the preferred version of the API. It returns "" when the
// directory cannot be fetched.
func (r *Resolver) documentationLink(ctx context.Context, name, version string) string {
	r.once.Do(func() { r.discoveryE = r.loadDiscovery(ctx) })
	if name == "" {
		return ""
	}
	if l := r.docLinks[name+":"+version]; l != "" {
		return l
	}
	return r.preferred[name]
}

func (r *Resolver) loadDiscovery(ctx context.Context) error {
	u := r.DiscoveryURL
	if u == "" {
		u = DiscoveryDirectoryURL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := r.client().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", u, resp.Status)
	}
	var dir struct {
		Items []struct {
			Name              string `json:"name"`
			Version           string `json:"version"`
			DocumentationLink string `json:"documentationLink"`
			Preferred         bool   `json:"preferred"`
		} `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dir); err != nil {
		return fmt.Errorf("decoding %s: %w", u, err)
	}
	r.docLinks = map[string]string{}
	r.preferred = map[string]string{}
	for _, it := range dir.Items {
		r.docLinks[it.Name+":"+it.Version] = it.DocumentationLink
		if it.Preferred {
			r.preferred[it.Name] = it.DocumentationLink
		}
	}
	return nil
}

// fetch GETs a URL, following redirects, and returns where it ended up and
// whether that page answered 200.
func (r *Resolver) fetch(ctx context.Context, u string) (*url.URL, bool) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, false
	}
	resp, err := r.client().Do(req)
	if err != nil {
		return nil, false
	}
	resp.Body.Close()
	return resp.Request.URL, resp.StatusCode == http.StatusOK
}

func (r *Resolver) client() *http.Client {
	if r.Client != nil {
		return r.Client
	}
	return http.DefaultClient
}
