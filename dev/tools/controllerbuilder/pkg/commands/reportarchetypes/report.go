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

package reportarchetypes

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/version"
)

// Row is one line of the report: a kind, the message it maps to, and what
// the proto says about managing that message.
type Row struct {
	Kind       string
	APIVersion string
	// Message is the full name of the resource message. For an unresolved
	// kind it is the first name that was tried.
	Message string
	// ProtoSource is the mapping's Mapping.ProtoSource: the descriptor set
	// generate.sh builds the kind against, empty for the default one.
	ProtoSource string
	// API is nil when the message is not in the descriptor set.
	API       *protoapi.ResourceAPI
	Archetype protoapi.Archetype
}

// BuildReport classifies each kind in mappings against the descriptors in
// api. It returns one row per kind, sorted by kind, for the mapping
// latestMappings picks.
func BuildReport(api *protoapi.Proto, mappings []Mapping) []Row {
	latest := latestMappings(mappings)
	rows := make([]Row, 0, len(latest))
	for _, m := range latest {
		row := Row{Kind: m.Kind, APIVersion: m.APIVersion, ProtoSource: m.ProtoSource}
		candidates := messageCandidates(m)
		row.Message = candidates[0]
		for _, name := range candidates {
			if resolved, err := api.ResourceAPIByName(protoreflect.FullName(name)); err == nil {
				row.Message, row.API = name, resolved
				break
			}
		}
		row.Archetype = protoapi.ClassifyArchetype(row.API)
		rows = append(rows, row)
	}
	return rows
}

// latestMappings returns one mapping per kind, sorted by kind. A kind mapped
// more than once keeps the mapping with the most mature KRM version (v1 >
// v1beta1 > v1alpha1). Among mappings at that version the last one wins, in
// the order ScanGenerateScripts returns them, because a later generate-types
// run overwrites the types an earlier one generated.
func latestMappings(mappings []Mapping) []Mapping {
	chosen := map[string]Mapping{}
	for _, m := range mappings {
		prev, seen := chosen[m.Kind]
		if !seen || version.CompareKubeAwareVersionStrings(krmVersion(m.APIVersion), krmVersion(prev.APIVersion)) >= 0 {
			chosen[m.Kind] = m
		}
	}
	kinds := make([]string, 0, len(chosen))
	for kind := range chosen {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)
	latest := make([]Mapping, 0, len(kinds))
	for _, kind := range kinds {
		latest = append(latest, chosen[kind])
	}
	return latest
}

// krmVersion returns the version half of a KRM apiVersion.
func krmVersion(apiVersion string) string {
	return apiVersion[strings.LastIndex(apiVersion, "/")+1:]
}

// messageCandidates returns the full names a mapping's message may have, in
// the order to try them. Each --service is tried as a prefix in turn, as
// generate-types does. A dotted name that starts in lower case is a package
// path, so it is tried as written first; other dotted names, such as
// Parent.Child, are nested messages, and are tried as written last.
func messageCandidates(m Mapping) []string {
	var candidates []string
	for _, svc := range m.Services {
		candidates = append(candidates, svc+"."+m.ProtoName)
	}
	if strings.Contains(m.ProtoName, ".") || len(candidates) == 0 {
		if startsWithLower(m.ProtoName) {
			candidates = append([]string{m.ProtoName}, candidates...)
		} else {
			candidates = append(candidates, m.ProtoName)
		}
	}
	return candidates
}

func startsWithLower(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(r)
}

// CountByArchetype returns how many rows have each archetype.
func CountByArchetype(rows []Row) map[protoapi.Archetype]int {
	counts := map[protoapi.Archetype]int{}
	for _, r := range rows {
		counts[r.Archetype]++
	}
	return counts
}

// tsvHeader names the report's columns.
var tsvHeader = []string{
	"kind", "apiVersion", "message", "protoSource", "archetype", "style",
	"defaultHost", "service", "pattern", "patternCount", "get", "create",
	"update", "delete", "list",
}

// WriteTSV writes rows as tab-separated values under a header line. An empty
// value is written as "-", so protoSource is "-" for the default descriptor
// set. The service column holds the short name of the service that hosts the
// standard methods; its package is the message's. A method cell holds the
// method's HTTP verb, or RPC when it has no HTTP binding, then nonstandard
// when its request lacks the AIP shape, then what the method takes and
// returns: id=<field> for a client-assigned ID, mask for an update_mask,
// etag for an etag, and lro for a long-running operation.
func WriteTSV(w io.Writer, rows []Row) error {
	if _, err := fmt.Fprintln(w, strings.Join(tsvHeader, "\t")); err != nil {
		return err
	}
	for _, r := range rows {
		cells := []string{r.Kind, r.APIVersion, r.Message, r.ProtoSource, string(r.Archetype) + "-" + r.Archetype.Name()}
		cells = append(cells, apiCells(r.API)...)
		for i, c := range cells {
			if c == "" {
				cells[i] = "-"
			}
		}
		if _, err := fmt.Fprintln(w, strings.Join(cells, "\t")); err != nil {
			return err
		}
	}
	return nil
}

// apiCells returns the report cells that come from the ResourceAPI, from
// style onwards.
func apiCells(api *protoapi.ResourceAPI) []string {
	if api == nil {
		return []string{"", "", "", "", "0", "", "", "", "", ""}
	}
	service := ""
	if api.Service != nil {
		service = string(api.Service.Name())
	}
	pattern, patternCount := "", 0
	if api.Metadata != nil {
		pattern, patternCount = api.Metadata.Pattern, len(api.Metadata.Patterns)
	}
	return []string{
		string(api.Style),
		api.DefaultHost,
		service,
		pattern,
		strconv.Itoa(patternCount),
		methodCell(api.Get),
		methodCell(api.Create),
		methodCell(api.Update),
		methodCell(api.Delete),
		methodCell(api.List),
	}
}

func methodCell(m *protoapi.StandardMethod) string {
	if m == nil {
		return ""
	}
	verb := m.HTTPVerb
	if verb == "" {
		verb = "RPC"
	}
	parts := []string{verb}
	if m.NonstandardRequest {
		parts = append(parts, "nonstandard")
	}
	if m.IDField != "" {
		parts = append(parts, "id="+m.IDField)
	}
	if m.UpdateMaskField != "" {
		parts = append(parts, "mask")
	}
	if m.EtagField != "" {
		parts = append(parts, "etag")
	}
	if m.LRO != nil {
		parts = append(parts, "lro")
	}
	return strings.Join(parts, " ")
}
