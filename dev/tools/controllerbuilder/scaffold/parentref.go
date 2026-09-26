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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
)

// sharedRefsPackage holds the reference types services use in common, relative
// to the repo root.
const sharedRefsPackage = "apis/refs/v1beta1"

// parentRef renders the Spec field naming a resource's direct parent and the
// queue entry that goes with it, as referenceTo describes. Both are empty where
// the parent is a location or the root of the name, which the location field
// and rootRef already carry.
func (a *APIScaffolder) parentRef(pattern string) (field string, item *JudgementItem) {
	collection, placeholder := protoapi.ParentPair(pattern)
	switch collection {
	case "", "locations", "regions", "zones", "global":
		return "", nil
	}
	path := parentPath(pattern, collection, placeholder)
	if strings.Count(path, "/") == 1 {
		return "", nil
	}

	// The collection segment identifies the parent resource type.
	return a.referenceTo("parent", collection, path, pattern)
}

// rootRef renders the Spec field naming the root of a resource's name, and the
// queue entry that goes with it when the root is one KCC has no fixed field
// for.
//
// A project, and a resource with no pattern, get projectRef; an organization
// or folder gets organizationRef or folderRef. Any other root, such as
// properties/{property} in Analytics, is looked up the way parentRef looks up
// a parent. A pattern that is only the root itself, such as
// billingAccounts/{billing_account}, names no root to point at, so it gets no
// field.
//
// It reads the first segment rather than protoapi.ParentStyle, which calls
// "organizations/{organization}/locations/{location}/..." ParentOther.
func (a *APIScaffolder) rootRef(pattern string) (field string, item *JudgementItem) {
	segs := strings.Split(pattern, "/")
	switch {
	case pattern == "":
		return fixedRootField("ProjectRef", "projectRef", "project"), nil
	case len(segs) < 3:
		return "", nil
	}
	switch segs[0] {
	case "projects":
		return fixedRootField("ProjectRef", "projectRef", "project"), nil
	case "organizations":
		return fixedRootField("OrganizationRef", "organizationRef", "organization"), nil
	case "folders":
		return fixedRootField("FolderRef", "folderRef", "folder"), nil
	}
	return a.referenceTo("root", segs[0], strings.Join(segs[:2], "/"), pattern)
}

// locationRef renders the Spec's location field and corresponding queue entry
// when the resource name pattern contains a location, region, or zone placeholder.
func (a *APIScaffolder) locationRef(pattern string) (field string, item *JudgementItem) {
	if pattern == "" {
		return locationField, &JudgementItem{
			FieldPath: ".spec.location",
			Reason:    "location-parent-unknown",
			Detail: "the proto declares no google.api.resource, so the parent shape is unknown; " +
				"drop location if the resource is not regional",
		}
	}

	segs := strings.Split(pattern, "/")
	collection, placeholder := "", ""
	for i := 0; i+1 < len(segs); i++ {
		switch segs[i] {
		case "locations", "regions", "zones":
			if strings.HasPrefix(segs[i+1], "{") {
				collection, placeholder = segs[i], strings.Trim(segs[i+1], "{}")
			}
		}
		if collection != "" {
			// The resource's own ID, as for projects/{project}/locations/{location},
			// is resourceID, not a location field.
			if i+2 == len(segs) {
				return "", nil
			}
			break
		}
	}
	if collection == "" {
		// The name has no location, or only a fixed one such as locations/global.
		return "", nil
	}

	var detail string
	parentCollection, parentPlaceholder := protoapi.ParentPair(pattern)
	switch {
	case parentCollection != collection || parentPlaceholder != placeholder:
		detail = fmt.Sprintf("the parent is %s, which already names the location; "+
			"a reference to it could replace location and the root reference",
			parentPath(pattern, parentCollection, parentPlaceholder))
	case segs[0] == "projects" && collection == "locations":
		// Its String() builds projects/.../locations/..., so a region or zone
		// parent falls through to the next case.
		detail = "parent.ProjectAndLocationRef in apis/common/parent, inlined, could replace " +
			"projectRef and location; the CRD keeps the same keys"
	default:
		detail = fmt.Sprintf("the parent is %s, and no shared reference type covers it yet; "+
			"keep location unless one is added", strings.Join(segs[:4], "/"))
	}
	return locationField, &JudgementItem{
		FieldPath: ".spec.location",
		Reason:    "location-or-parent-ref",
		Detail:    detail,
	}
}

// locationField is the Spec's location field, exactly as the types template
// wrote it before, so a resource with a location scaffolds unchanged.
const locationField = "\t// The location of this resource.\n" +
	"\tLocation string `json:\"location\"`"

// fixedRootField renders a required reference to one of the roots every KCC
// resource can point at, using the shared type of that name.
func fixedRootField(refType, jsonName, noun string) string {
	return fmt.Sprintf("\t// The %s that this resource belongs to.\n"+
		"\t%s *refsv1beta1.%s `json:%q`", noun, refType, refType, jsonName)
}

// referenceTo renders a Spec field pointing at the resource at path, whose
// collection segment is collection, and the queue entry for it. role says
// which ancestor this is, "parent" or "root", in the entry.
//
// The field is written only when exactly one matching reference type exists.
func (a *APIScaffolder) referenceTo(role, collection, path, pattern string) (field string, item *JudgementItem) {
	name := codegen.Singular(collection)
	currentService := a.serviceName()

	var candidates map[string]string
	// 1. Try google.api.resource / google.api.resource_reference from proto first.
	if targetType := a.targetForPath(path); targetType != "" {
		candidates = a.candidatesForTarget(targetType, filepath.Join(a.BaseDir, a.GoPackage))
	}
	// 2. Fall back to name-based matching if no candidates found from proto.
	if len(candidates) == 0 {
		candidates = parentRefTypes(a.repoRoot(), filepath.Join(a.BaseDir, a.GoPackage), currentService, name)
	}

	if len(candidates) != 1 {
		detail := fmt.Sprintf("the %s is %s, and no %sRef type exists to point at; "+
			"add the %s resource first, then model this as a reference", role, path, exportedName(name), role)
		if len(candidates) > 1 {
			detail = fmt.Sprintf("the %s is %s, and several reference types match it (%s); "+
				"pick one and add the field by hand", role, path, strings.Join(sortedNames(candidates), ", "))
		}
		return "", &JudgementItem{
			Reason: role + "-ref-not-modelled",
			Detail: detail,
		}
	}

	typeName := sortedNames(candidates)[0]
	goType := typeName
	if qualifier := candidates[typeName]; qualifier != "" {
		goType = qualifier + "." + typeName
	}

	field = fmt.Sprintf("\t// A reference to the %s this resource belongs to.\n"+
		"\t// +kcc:guess\n"+
		"\t// %sRef *%s `json:%q`", path, exportedName(name), goType, name+"Ref,omitempty")
	return field, &JudgementItem{
		FieldPath: ".spec." + name + "Ref",
		Reason:    role + "-ref-guessed",
		Detail: fmt.Sprintf("emitted as a reference to %s, read from the %s segment of %s; "+
			"the field name comes from the pattern rather than the proto, so confirm both the name and the target",
			typeName, collection, pattern),
	}
}

// targetForPath finds the google.api.resource or google.api.resource_definition target for path in the proto.
func (a *APIScaffolder) targetForPath(path string) string {
	if a.Proto == nil || path == "" {
		return ""
	}
	for _, f := range a.Proto.SortedFiles() {
		// Check file-level resource_definition annotations
		if v := proto.GetExtension(f.Options(), annotations.E_ResourceDefinition); v != nil {
			if rds, ok := v.([]*annotations.ResourceDescriptor); ok {
				for _, rd := range rds {
					if rd != nil {
						for _, p := range rd.GetPattern() {
							if p == path {
								return rd.GetType()
							}
						}
					}
				}
			}
		}
		// Check message-level resource annotations
		for i := 0; i < f.Messages().Len(); i++ {
			msg := f.Messages().Get(i)
			if v := proto.GetExtension(msg.Options(), annotations.E_Resource); v != nil {
				if rd, ok := v.(*annotations.ResourceDescriptor); ok && rd != nil {
					for _, p := range rd.GetPattern() {
						if p == path {
							return rd.GetType()
						}
					}
				}
			}
		}
	}
	return ""
}

// parseResourceTarget parses a resource type into its service and kind components.
func parseResourceTarget(target string) (service, kind string) {
	parts := strings.Split(target, "/")
	if len(parts) == 0 {
		return "", ""
	}
	kind = parts[len(parts)-1]
	domain := parts[0]
	service = strings.TrimSuffix(domain, ".googleapis.com")
	return service, kind
}

// candidatesForTarget finds reference types matching a known proto resource target type.
func (a *APIScaffolder) candidatesForTarget(targetType, serviceDir string) map[string]string {
	targetService, targetKind := parseResourceTarget(targetType)
	if targetKind == "" {
		return nil
	}
	currentService := a.serviceName()
	want := strings.ToLower(targetKind) + "ref"
	re := regexp.MustCompile(`(?m)^type ([A-Z]\w*Ref) struct`)

	// 1. Scan serviceDir if service matches
	if serviceMatchesPrefix(currentService, targetService) || serviceMatchesPrefix(targetService, currentService) {
		found := map[string]string{}
		entries, err := os.ReadDir(serviceDir)
		if err == nil {
			for _, e := range entries {
				if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
					continue
				}
				body, err := os.ReadFile(filepath.Join(serviceDir, e.Name()))
				if err != nil {
					continue
				}
				for _, m := range re.FindAllStringSubmatch(string(body), -1) {
					if strings.HasSuffix(strings.ToLower(m[1]), want) {
						found[m[1]] = ""
					}
				}
			}
		}
		if len(found) > 0 {
			return found
		}
	}

	// 2. Scan shared refs package
	sharedDir := filepath.Join(a.repoRoot(), sharedRefsPackage)
	found := map[string]string{}
	entries, err := os.ReadDir(sharedDir)
	if err == nil {
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			body, err := os.ReadFile(filepath.Join(sharedDir, e.Name()))
			if err != nil {
				continue
			}
			for _, m := range re.FindAllStringSubmatch(string(body), -1) {
				typeName := m[1]
				lowerType := strings.ToLower(typeName)
				if !strings.HasSuffix(lowerType, want) {
					continue
				}
				// Exact match (e.g. ProjectRef, FolderRef, OrganizationRef)
				if strings.EqualFold(typeName, want) {
					found[typeName] = "refsv1beta1"
					continue
				}
				// Service prefix match
				prefix := strings.TrimSuffix(lowerType, want)
				if serviceMatchesPrefix(targetService, prefix) || serviceMatchesPrefix(currentService, prefix) {
					found[typeName] = "refsv1beta1"
				}
			}
		}
	}
	return found
}

// serviceName extracts the service name from APIScaffolder's GoPackage or Group.
func (a *APIScaffolder) serviceName() string {
	if a.GoPackage != "" {
		clean := filepath.Clean(a.GoPackage)
		parts := strings.Split(clean, string(filepath.Separator))
		if len(parts) > 0 && parts[0] != "" && parts[0] != "." {
			return parts[0]
		}
	}
	if a.Group != "" {
		parts := strings.Split(a.Group, ".")
		if len(parts) > 0 {
			return parts[0]
		}
	}
	return ""
}

// serviceMatchesPrefix checks if a service name matches a prefix on a reference type.
func serviceMatchesPrefix(service, prefix string) bool {
	service = strings.ToLower(service)
	prefix = strings.ToLower(prefix)
	if service == "" || prefix == "" {
		return false
	}
	if service == prefix {
		return true
	}
	switch service {
	case "sql", "sqladmin":
		return prefix == "sql" || prefix == "sqladmin"
	case "iam":
		return prefix == "iam" || prefix == "gcpserviceaccount"
	case "secretmanager":
		return prefix == "secretmanager" || prefix == "secret"
	case "bigquery":
		return prefix == "bigquery"
	case "compute":
		return prefix == "compute"
	case "alloydb":
		return prefix == "alloydb"
	case "appengine":
		return prefix == "appengine"
	case "kms":
		return prefix == "kms"
	}
	return false
}

// isAllowedSharedRef ensures that types in apis/refs/v1beta1 only match if
// they are generic un-prefixed references or their prefix matches the current service.
func isAllowedSharedRef(typeName, currentService, want string) bool {
	if strings.EqualFold(typeName, want) {
		return true
	}
	if currentService == "" {
		return false
	}
	prefix := strings.TrimSuffix(strings.ToLower(typeName), want)
	return serviceMatchesPrefix(currentService, prefix)
}

// parentRefTypes returns the reference types that could name a parent whose
// collection singularises to name, as a map of type name to import qualifier.
//
// Services spell a reference type either as a bare noun, ClusterRef in alloydb,
// or with the Kind's own prefix, as in DNSManagedZoneRef, so the match is on
// the suffix. The service's own package wins over the shared one, and
// unexported types are skipped: apis/kms/v1beta1 declares kmsCryptoKeyRef
// beside the real KMSCryptoKeyRef, which would make a single match look
// ambiguous.
//
// Shared reference types in apis/refs/v1beta1 are only matched if they are
// un-prefixed generic references or their prefix matches the current service,
// preventing foreign services (e.g. Chronicle) from accidentally matching
// service-specific types (e.g. SQLInstanceRef).
func parentRefTypes(repoRoot, serviceDir, currentService, name string) map[string]string {
	want := strings.ToLower(name) + "ref"
	re := regexp.MustCompile(`(?m)^type ([A-Z]\w*Ref) struct`)

	scan := func(dir, qualifier string, isShared bool) map[string]string {
		found := map[string]string{}
		entries, err := os.ReadDir(dir)
		if err != nil {
			return found
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			body, err := os.ReadFile(filepath.Join(dir, e.Name()))
			if err != nil {
				continue
			}
			for _, m := range re.FindAllStringSubmatch(string(body), -1) {
				typeName := m[1]
				lowerType := strings.ToLower(typeName)
				if !strings.HasSuffix(lowerType, want) {
					continue
				}
				if isShared {
					if !isAllowedSharedRef(typeName, currentService, want) {
						continue
					}
				}
				found[typeName] = qualifier
			}
		}
		return found
	}

	if found := scan(serviceDir, "", false); len(found) > 0 {
		return found
	}
	return scan(filepath.Join(repoRoot, sharedRefsPackage), "refsv1beta1", true)
}

// repoRoot returns the directory above BaseDir when BaseDir is an apis
// directory, and BaseDir otherwise, so the shared refs package can be found.
//
// generate-types defaults BaseDir to "<repo>/apis/", with a trailing slash,
// and filepath.Dir of that is "<repo>/apis", so the path is cleaned first.
func (a *APIScaffolder) repoRoot() string {
	dir := filepath.Clean(a.BaseDir)
	if filepath.Base(dir) == "apis" {
		return filepath.Dir(dir)
	}
	return dir
}

// parentPath is the pattern truncated at the parent, which is the whole path a
// reference to it carries in its External field, per AIP-122.
func parentPath(pattern, collection, placeholder string) string {
	segs := strings.Split(pattern, "/")
	for i := 0; i+1 < len(segs); i++ {
		if segs[i] == collection && segs[i+1] == "{"+placeholder+"}" {
			return strings.Join(segs[:i+2], "/")
		}
	}
	return pattern
}

func sortedNames(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func exportedName(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
