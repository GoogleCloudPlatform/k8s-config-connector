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

// Package kindnames reads the hand-maintained table of service prefixes
// (serviceprefixes/service-prefixes.yaml) and builds the Kind it implies for a
// GCP resource.
//
// A Kind is <ServicePrefix><ResourceType>: the service prefix of the API,
// followed by the type in google.api.resource in the singular, e.g. LoggingSetting
// for logging.googleapis.com/Settings.
package kindnames

import (
	"strings"
	"unicode"
)

// ResourceType returns the type part of a google.api.resource type, in the
// singular, because a Kind is always singular: Instance for
// sqladmin.googleapis.com/Instance, and Setting for
// logging.googleapis.com/Settings.
func ResourceType(resourceType string) string {
	_, t, ok := strings.Cut(resourceType, "/")
	if !ok {
		return ""
	}
	return singularLastWord(upperFirst(t))
}

// notPlural lists words ending in s that are not plurals.
var notPlural = map[string]bool{
	"Alias":      true,
	"Atlas":      true,
	"Canvas":     true,
	"Kubernetes": true,
	"News":       true,
	"Series":     true,
	"Species":    true,
}

// irregularPlurals maps plurals the suffix rules get wrong to their singular.
var irregularPlurals = map[string]string{
	"Caches":   "Cache",
	"Indices":  "Index",
	"Sizes":    "Size",
	"Statuses": "Status",
}

// singularLastWord puts the last word of a PascalCase name in the singular:
// ProjectSettings becomes ProjectSetting and SecurityMarks becomes
// SecurityMark, while ServiceClass, RagCorpus, VisionAIAnalysis and an
// all-caps acronym such as DNS are left alone.
func singularLastWord(name string) string {
	i := lastWordStart(name)
	return name[:i] + singularWord(name[i:])
}

func singularWord(w string) string {
	if w == strings.ToUpper(w) || notPlural[w] {
		return w
	}
	if s, ok := irregularPlurals[w]; ok {
		return s
	}
	switch {
	case strings.HasSuffix(w, "ies") && len(w) > 4:
		// Policies -> Policy
		return w[:len(w)-3] + "y"
	case strings.HasSuffix(w, "sses"), strings.HasSuffix(w, "shes"),
		strings.HasSuffix(w, "ches"), strings.HasSuffix(w, "xes"),
		strings.HasSuffix(w, "zes"):
		// Addresses -> Address, Batches -> Batch
		return w[:len(w)-2]
	case strings.HasSuffix(w, "ss"), strings.HasSuffix(w, "us"), strings.HasSuffix(w, "is"):
		// Access, Status, Analysis: not plurals.
		return w
	case strings.HasSuffix(w, "s") && len(w) > 2:
		// Settings -> Setting
		return w[:len(w)-1]
	}
	return w
}

// lastWordStart returns the index where the last word of a PascalCase name
// starts: 7 for ProjectSettings, 0 for Settings.
func lastWordStart(name string) int {
	r := []rune(name)
	start := 0
	for i := 1; i < len(r); i++ {
		prev, cur := r[i-1], r[i]
		if !unicode.IsUpper(cur) {
			continue
		}
		if unicode.IsLower(prev) || unicode.IsDigit(prev) ||
			(unicode.IsUpper(prev) && i+1 < len(r) && unicode.IsLower(r[i+1])) {
			start = i
		}
	}
	return len(string(r[:start]))
}

// APIForResource returns the API name of a resource: from the
// google.api.default_host of its proto package, or, when the descriptor has
// no service for the package (only its message files were compiled in, as
// imports of other APIs), from the host of its google.api.resource type, which
// names the same service: accesscontextmanager for
// accesscontextmanager.googleapis.com/AccessLevel.
func APIForResource(defaultHost, resourceType string) string {
	if defaultHost != "" {
		return apiName(defaultHost)
	}
	host, _, ok := strings.Cut(resourceType, "/")
	if !ok || !strings.HasSuffix(host, ".googleapis.com") {
		return ""
	}
	return apiName(host)
}

// apiName returns the API name for a default host: sqladmin for
// sqladmin.googleapis.com (or sqladmin.googleapis.com:443).
func apiName(host string) string {
	host, _, _ = strings.Cut(host, ":")
	return strings.TrimSuffix(host, ".googleapis.com")
}

func upperFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
