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
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
)

// parentRef renders the Spec field naming a resource's direct parent and the
// queue entry that goes with it, as referenceTo describes. Both are empty where
// the parent is a location or the root of the name, which the location field
// and rootRef already carry.
func (a *APIScaffolder) parentRef(pattern string) (field string, item *JudgementItem) {
	collection, _ := protoapi.ParentPair(pattern)
	switch collection {
	case "", "locations", "regions", "zones", "global":
		return "", nil
	}
	path := protoapi.ParentPattern(pattern)
	if strings.Count(path, "/") == 1 {
		return "", nil
	}

	// The collection segment identifies the parent resource type.
	return a.referenceTo("parent", collection, path, pattern)
}

// rootRef renders the Spec field naming the root of a resource's name, and the
// queue entry that goes with it when the root is one KCC has no fixed field
// for, or when the name has no root at all.
//
// A project, and a resource with no pattern, get projectRef; an organization
// or folder gets organizationRef or folderRef. Any other root, such as
// properties/{property} in Analytics, is looked up the way parentRef looks up
// a parent. A pattern that is only the resource itself, such as
// folders/{folder}, names no root to point at, so it gets no field, only a
// queue entry: such a resource usually still lives in an organization, folder
// or project, given in a parent field rather than in its name.
//
// It reads the first segment rather than protoapi.ParentStyle, which calls
// "organizations/{organization}/locations/{location}/..." ParentOther.
func (a *APIScaffolder) rootRef(pattern string) (field string, item *JudgementItem) {
	segs := strings.Split(pattern, "/")
	switch {
	case pattern == "":
		return fixedRootField("ProjectRef", "projectRef", "project"), nil
	case len(segs) < 3:
		return "", &JudgementItem{
			Reason: "root-not-in-name",
			Detail: fmt.Sprintf("the name %s has nothing above the resource, so no root reference was emitted; "+
				"if the resource or its Create request takes a parent, as Folder, Project and TagKey do, "+
				"model that parent as a reference", pattern),
		}
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

// fixedRootField renders a required reference to one of the roots every KCC
// resource can point at, using the shared type of that name.
func fixedRootField(refType, jsonName, noun string) string {
	return fmt.Sprintf("\t// The %s that this resource belongs to.\n"+
		"\t%s *refsv1beta1.%s `json:%q`", noun, refType, refType, jsonName)
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
	collection, placeholder, at := "", "", 0
	for i := 0; i+1 < len(segs); i++ {
		switch segs[i] {
		case "locations", "regions", "zones":
			if strings.HasPrefix(segs[i+1], "{") {
				collection, placeholder, at = segs[i], strings.Trim(segs[i+1], "{}"), i
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
	case parentCollection == "":
		// ParentPair finds no parent in a name with one pair, such as
		// locations/{location}/policy, or in one that ends in two placeholders.
		detail = fmt.Sprintf("no parent was found in %s; "+
			"keep location unless a reference to the parent replaces it", pattern)
	case parentCollection != collection || parentPlaceholder != placeholder:
		detail = fmt.Sprintf("the parent is %s, which already names the location; "+
			"a reference to it could replace location and the root reference",
			protoapi.ParentPattern(pattern))
	case segs[0] == "projects" && collection == "locations":
		// Its String() builds projects/.../locations/..., so a region or zone
		// parent falls through to the next case.
		detail = "parent.ProjectAndLocationRef in apis/common/parent, inlined, could replace " +
			"projectRef and location; the CRD keeps the same keys"
	default:
		// The location need not be the second pair, as in
		// organizations/{organization}/sources/{source}/locations/{location}/findings/{finding}.
		detail = fmt.Sprintf("the parent is %s, and no shared reference type covers it yet; "+
			"keep location unless one is added", strings.Join(segs[:at+2], "/"))
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

// referenceTo renders a Spec field pointing at the resource at path, whose
// collection segment is collection, and the queue entry for it. role says
// which ancestor this is, "parent" or "root", in the entry.
//
// The field is written only when exactly one matching reference type exists.
func (a *APIScaffolder) referenceTo(role, collection, path, pattern string) (field string, item *JudgementItem) {
	name := codegen.Singular(collection)
	currentService := a.serviceName()

	var candidates map[string]string
	// 1. Try google.api.resource / google.api.resource_definition from proto first.
	if a.Proto != nil && path != "" {
		if targetType := a.Proto.TargetForPath(path); targetType != "" {
			candidates = a.candidatesForTarget(targetType, filepath.Join(a.BaseDir, a.GoPackage))
		}
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
