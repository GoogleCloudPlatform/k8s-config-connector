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

package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"sigs.k8s.io/yaml"
)

// generateTypesResource is one --resource <Kind>:<ProtoMessage> argument of a
// generate-types call in apis/<service>/generate.sh.
type generateTypesResource struct {
	Kind    string
	Message string
	Service string // the --service flag of the call, e.g. google.cloud.vectorsearch.v1
	Script  string // e.g. apis/vectorsearch/generate.sh
}

var (
	// serviceFlagRe matches --service pkg[,pkg...], with optional quotes.
	serviceFlagRe = regexp.MustCompile(`--service[= ]+["']?([A-Za-z0-9_.,]+)["']?`)
	// resourceFlagRe matches --resource Kind:Message and --resource=Kind:Message,
	// with optional quotes.
	resourceFlagRe = regexp.MustCompile(`--resource[= ]+["']?([A-Za-z0-9]+):["']?([A-Za-z0-9_.]+)["']?`)
)

// parseGenerateTypesResources returns the --resource arguments of every
// generate-types call in a generate.sh script.
func parseGenerateTypesResources(script, content string) []generateTypesResource {
	// Join backslash-continued lines so each command is on one line.
	joined := strings.ReplaceAll(content, "\\\r\n", " ")
	joined = strings.ReplaceAll(joined, "\\\n", " ")

	var out []generateTypesResource
	for _, line := range strings.Split(joined, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "generate-types") {
			continue
		}
		service := ""
		if m := serviceFlagRe.FindStringSubmatch(line); m != nil {
			service = m[1]
		}
		for _, m := range resourceFlagRe.FindAllStringSubmatch(line, -1) {
			msg := m[2]
			if i := strings.LastIndex(msg, "."); i >= 0 {
				msg = msg[i+1:]
			}
			out = append(out, generateTypesResource{Kind: m[1], Message: msg, Service: service, Script: script})
		}
	}
	return out
}

// The Kind naming files in dev/tools/controllerbuilder/serviceprefixes.
const (
	servicePrefixesDir         = "../../dev/tools/controllerbuilder/serviceprefixes"
	servicePrefixesFile        = "service-prefixes.yaml"
	servicePrefixesPendingFile = "service-prefix-pending.yaml"
	acronymsFile               = "acronyms.yaml"
)

// servicePrefixEntry is the service prefix of one API.
type servicePrefixEntry struct {
	API           string   `json:"api"`
	ServicePrefix string   `json:"servicePrefix"`
	ProtoPackages []string `json:"protoPackages"`
	// Pending is true for an entry in servicePrefixesPendingFile.
	Pending bool `json:"-"`
}

// loadServicePrefixes reads both service-prefix files in dir and returns the
// entry of each proto package listed in protoPackages.
func loadServicePrefixes(dir string) (map[string]servicePrefixEntry, error) {
	out := map[string]servicePrefixEntry{}
	for _, file := range []struct {
		name    string
		pending bool
	}{
		{servicePrefixesFile, false},
		{servicePrefixesPendingFile, true},
	} {
		path := filepath.Join(dir, file.name)
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var f struct {
			Services []servicePrefixEntry `json:"services"`
		}
		if err := yaml.Unmarshal(data, &f); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, err)
		}
		for _, e := range f.Services {
			e.Pending = file.pending
			for _, pkg := range e.ProtoPackages {
				if prev, dup := out[pkg]; dup {
					return nil, fmt.Errorf("%s: proto package %s is listed for both api %q and api %q", path, pkg, prev.API, e.API)
				}
				out[pkg] = e
			}
		}
	}
	return out, nil
}

// loadAcronyms reads the "acronyms" map of acronymsFile in dir, e.g.
// Api: API.
func loadAcronyms(dir string) (map[string]string, error) {
	path := filepath.Join(dir, acronymsFile)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f struct {
		Acronyms map[string]string `json:"acronyms"`
	}
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return f.Acronyms, nil
}

// pascalWordRe matches one word of a PascalCase name: Api, Hub, URL, V2.
var pascalWordRe = regexp.MustCompile(`[A-Z]+[a-z0-9]*|[a-z0-9]+`)

// applyAcronyms replaces each whole word of a PascalCase name that is a key of
// acronyms with its value: ApiDeployment gives APIDeployment, and
// QaQuestion gives QAQuestion. Words are matched whole,
// so Apis and Rapid are left alone.
func applyAcronyms(name string, acronyms map[string]string) string {
	return pascalWordRe.ReplaceAllStringFunc(name, func(w string) string {
		if a, ok := acronyms[w]; ok {
			return a
		}
		return w
	})
}

// knownKindNames lists resources whose Kind cannot be built directly from the proto
// message name, with the Kind they should have. An entry allows that Kind in
// addition to <servicePrefix><message>, so another Kind generated from the same
// message (AIPlatformNotebookRuntimeTemplate next to ColabRuntimeTemplate) is
// still checked normally.
var knownKindNames = []struct {
	ServicePrefix string // the servicePrefix of the resource's API
	Message       string // the proto message passed to generate-types --resource
	Want          string // the correct Kind
}{
	// Kinds that were first implemented with the Terraform controller are
	// named after their Terraform resource, not the proto message, and keep
	// that name so existing users are not broken
	{ServicePrefix: "Apigee", Message: "GoogleCloudApigeeV1ApiProduct", Want: "ApigeeAPIProduct"},
	{ServicePrefix: "Apigee", Message: "GoogleCloudApigeeV1NatAddress", Want: "ApigeeNATAddress"},
	{ServicePrefix: "Apigee", Message: "GoogleCloudApigeeV1SyncAuthorization", Want: "ApigeeSyncAuthorization"},
	{ServicePrefix: "Apigee", Message: "GoogleCloudApigeeV1AddonsConfig", Want: "ApigeeAddonsConfig"},
	{ServicePrefix: "AppEngine", Message: "Service", Want: "AppEngineServiceSplitTraffic"},
	{ServicePrefix: "AppEngine", Message: "Version", Want: "AppEngineFlexibleAppVersion"},
	{ServicePrefix: "AppEngine", Message: "Version", Want: "AppEngineStandardAppVersion"},
	{ServicePrefix: "BigQuery", Message: "Access", Want: "BigQueryDatasetAccess"},
	{ServicePrefix: "Bigtable", Message: "GcRule", Want: "BigtableGCPolicy"},
	{ServicePrefix: "AIPlatform", Message: "NotebookRuntime", Want: "ColabRuntime"},
	{ServicePrefix: "AIPlatform", Message: "NotebookRuntimeTemplate", Want: "ColabRuntimeTemplate"},
	{ServicePrefix: "Compute", Message: "SignedUrlKey", Want: "ComputeBackendServiceSignedURLKey"},
	{ServicePrefix: "Compute", Message: "FirewallPolicy", Want: "ComputeNetworkFirewallPolicy"},
	{ServicePrefix: "Compute", Message: "SecurityPolicy", Want: "ComputeOrganizationSecurityPolicy"},
	{ServicePrefix: "Compute", Message: "SecurityPolicyAssociation", Want: "ComputeOrganizationSecurityPolicyAssociation"},
	{ServicePrefix: "Compute", Message: "Metadata", Want: "ComputeProjectMetadata"},
	{ServicePrefix: "Compute", Message: "Autoscaler", Want: "ComputeRegionAutoscaler"},
	{ServicePrefix: "Compute", Message: "PerInstanceConfig", Want: "ComputeRegionPerInstanceConfig"},
	{ServicePrefix: "Compute", Message: "SslPolicy", Want: "ComputeRegionSSLPolicy"},
	{ServicePrefix: "Compute", Message: "SslCertificate", Want: "ComputeManagedSSLCertificate"},
	{ServicePrefix: "GKEMultiCloud", Message: "AttachedCluster", Want: "ContainerAttachedCluster"},
	{ServicePrefix: "DocumentAI", Message: "SetDefaultProcessorVersionRequest", Want: "DocumentAIProcessorDefaultVersion"},
	{ServicePrefix: "IAM", Message: "Policy", Want: "IAMDenyPolicy"},
	{ServicePrefix: "KMS", Message: "ImportJob", Want: "KMSKeyRingImportJob"},
	{ServicePrefix: "CloudResourceManager", Message: "TagBinding", Want: "TagsLocationTagBinding"},
	{ServicePrefix: "CloudResourceManager", Message: "TagBinding", Want: "TagsTagBinding"},
	{ServicePrefix: "CloudResourceManager", Message: "TagKey", Want: "TagsTagKey"},
	{ServicePrefix: "CloudResourceManager", Message: "TagValue", Want: "TagsTagValue"},

	// The message name starts with the words the service prefix ends with, so
	// <servicePrefix><message> would repeat them (BigQueryDataPolicyDataPolicy).
	// The Kind says them once.
	{ServicePrefix: "APIHub", Message: "ApiHubInstance", Want: "APIHubInstance"},
	{ServicePrefix: "BigQueryDataPolicy", Message: "DataPolicy", Want: "BigQueryDataPolicy"},
	{ServicePrefix: "Billing", Message: "BillingAccount", Want: "BillingAccount"},
	{ServicePrefix: "CloudBuild", Message: "BuildTrigger", Want: "CloudBuildTrigger"},
	{ServicePrefix: "ConfigDeployment", Message: "DeploymentGroup", Want: "ConfigDeploymentGroup"},
	{ServicePrefix: "IAP", Message: "IapSettings", Want: "IAPSettings"},
	{ServicePrefix: "VMwareEngine", Message: "VmwareEngineNetwork", Want: "VMwareEngineNetwork"},

	// Others
	{ServicePrefix: "CloudResourceManager", Message: "Folder", Want: "Folder"},
	{ServicePrefix: "CloudResourceManager", Message: "Project", Want: "Project"},
}

// isKnownKindName reports whether knownKindNames allows kind for a message of
// an API with the given service prefix.
func isKnownKindName(servicePrefix, message, kind string) bool {
	for _, k := range knownKindNames {
		if k.ServicePrefix == servicePrefix && k.Message == message && k.Want == kind {
			return true
		}
	}
	return false
}

// kindNameViolations checks the Kind of every resource that a generate.sh
// passes to generate-types against the service-prefix table: the Kind must be
// a CRD Kind, its API must have an approved service prefix, and the Kind must
// be that prefix followed by the proto message name, with the
// acronyms in acronymsFile applied to the
// message words (UrlMap gives URLMap).
//
// The API is the one whose protoPackages lists a package of the
// generate-types --service flag (the first listed package wins).
func kindNameViolations(apisDir string, prefixes map[string]servicePrefixEntry, acronyms map[string]string, crdKinds map[string]bool) ([]string, error) {
	scripts, err := filepath.Glob(filepath.Join(apisDir, "*", "generate.sh"))
	if err != nil {
		return nil, err
	}

	var resources []generateTypesResource
	for _, path := range scripts {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		script := filepath.ToSlash(filepath.Join("apis", filepath.Base(filepath.Dir(path)), "generate.sh"))
		resources = append(resources, parseGenerateTypesResources(script, string(data))...)
	}

	seen := map[string]bool{}
	var errs []string
	for _, r := range resources {
		key := r.Kind + ":" + r.Message
		if seen[key] {
			continue
		}
		seen[key] = true

		var e servicePrefixEntry
		found := false
		for _, pkg := range strings.Split(r.Service, ",") {
			if e, found = prefixes[pkg]; found {
				break
			}
		}
		api := e.API
		if !found {
			api = "unknown"
		}
		prefix := fmt.Sprintf("[kind-name-violation] kind=%s api=%s message=%s", r.Kind, api, r.Message)
		if !crdKinds[r.Kind] {
			errs = append(errs, fmt.Sprintf("%s: no CRD has this kind (%s)", prefix, r.Script))
			continue
		}
		if !found {
			errs = append(errs, fmt.Sprintf("%s: no API lists proto package %q of %s in protoPackages in %s or %s", prefix, r.Service, r.Script, servicePrefixesFile, servicePrefixesPendingFile))
			continue
		}
		// A Kind listed in knownKindNames keeps its name whatever prefix the
		// API ends up with, so it is accepted even while the prefix is pending.
		if isKnownKindName(e.ServicePrefix, r.Message, r.Kind) {
			continue
		}
		if e.Pending {
			errs = append(errs, fmt.Sprintf("%s: the service prefix of API %s (%s) is pending approval in %s", prefix, api, e.ServicePrefix, servicePrefixesPendingFile))
			continue
		}
		typeName := applyAcronyms(r.Message, acronyms)
		want := e.ServicePrefix + typeName
		if r.Kind != want {
			errs = append(errs, fmt.Sprintf("%s: expected kind name is %s", prefix, want))
		}
	}
	sort.Strings(errs)
	return errs, nil
}

// kindNameViolationsFile lists the known violations of
// TestKindNamesFollowServicePrefixes.
const kindNameViolationsFile = "testdata/exceptions/kind_name_violations.txt"

// TestKindNamesFollowServicePrefixes checks that every Kind generated by an
// apis/<service>/generate.sh is <servicePrefix><proto message>, using
// dev/tools/controllerbuilder/serviceprefixes. Existing violations are listed in
// testdata/exceptions/kind_name_violations.txt; new Kinds must not add to it.
//
// Run with WRITE_GOLDEN_OUTPUT=1 to rewrite the exceptions list with the
// current violations; the test then passes.
func TestKindNamesFollowServicePrefixes(t *testing.T) {
	t.Parallel()

	prefixes, err := loadServicePrefixes(servicePrefixesDir)
	if err != nil {
		t.Fatalf("loading service-prefix table: %v", err)
	}
	acronyms, err := loadAcronyms(servicePrefixesDir)
	if err != nil {
		t.Fatalf("loading acronyms: %v", err)
	}
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading CRDs: %v", err)
	}
	crdKinds := map[string]bool{}
	for _, crd := range crds {
		crdKinds[crd.Spec.Names.Kind] = true
	}

	errs, err := kindNameViolations("../../apis", prefixes, acronyms, crdKinds)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(errs, "\n")
	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		if err := os.WriteFile(kindNameViolationsFile, []byte(got), 0644); err != nil {
			t.Fatalf("writing %s: %v", kindNameViolationsFile, err)
		}
		t.Logf("wrote %d violation(s) to %s", len(errs), kindNameViolationsFile)
		return
	}
	test.CompareGoldenFile(t, kindNameViolationsFile, got)
}
