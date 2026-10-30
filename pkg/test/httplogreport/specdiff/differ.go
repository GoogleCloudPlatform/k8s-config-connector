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

package specdiff

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Item represents a single field difference between create.yaml and update.yaml.
type Item struct {
	Path             string
	CreateVal        any
	UpdateVal        any
	SentInUpdateRPC  bool
	UpdateRPCDetail  string
	FinalStateMatch  bool
	FinalStateDetail string
}

// Differ computes deterministic KRM field diffs and correlates them against
// UPDATE-phase RPC requests and final GCP resource state.
type Differ struct {
	fix     *fixture.Data
	primary []*classifier.ClassifiedEntry
}

// New constructs a new Differ for the given fixture data and classified primary entries.
func New(fix *fixture.Data, primary []*classifier.ClassifiedEntry) *Differ {
	return &Differ{
		fix:     fix,
		primary: primary,
	}
}

// Compute returns the deterministic slice of spec/label diffs along with their
// update RPC coverage and final GCP state verification details.
func (d *Differ) Compute() []Item {
	if d.fix == nil || d.fix.PrimaryCreate == nil || d.fix.PrimaryUpdate == nil {
		return []Item{}
	}
	diffs := DiffObjects(d.fix.PrimaryCreate.Obj, d.fix.PrimaryUpdate.Obj)
	if len(diffs) == 0 {
		return []Item{}
	}

	d.populateUpdateCoverage(diffs)
	d.populateFinalStateCoverage(diffs)
	return diffs
}

// DiffObjects computes the deterministic list of field changes between createObj and updateObj
// across .spec and .metadata.labels.
func DiffObjects(createObj, updateObj *unstructured.Unstructured) []Item {
	if createObj == nil || updateObj == nil {
		return []Item{}
	}
	diffs := make([]Item, 0)

	createSpec, _, _ := unstructured.NestedMap(createObj.Object, "spec")
	updateSpec, _, _ := unstructured.NestedMap(updateObj.Object, "spec")
	diffMapsRecursive(createSpec, updateSpec, "spec", &diffs)

	createLabels, _, _ := unstructured.NestedMap(createObj.Object, "metadata", "labels")
	updateLabels, _, _ := unstructured.NestedMap(updateObj.Object, "metadata", "labels")
	diffMapsRecursive(createLabels, updateLabels, "metadata.labels", &diffs)

	return diffs
}

func diffMapsRecursive(cMap, uMap map[string]any, prefix string, out *[]Item) {
	keySet := make(map[string]struct{})
	for k := range cMap {
		keySet[k] = struct{}{}
	}
	for k := range uMap {
		keySet[k] = struct{}{}
	}
	var keys []string
	for k := range keySet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		path := prefix + "." + k
		cVal, cOk := cMap[k]
		uVal, uOk := uMap[k]

		if !cOk && uOk {
			if uSubMap, ok := uVal.(map[string]any); ok && !strings.HasSuffix(k, "Ref") {
				diffMapsRecursive(nil, uSubMap, path, out)
			} else {
				*out = append(*out, Item{Path: path, CreateVal: nil, UpdateVal: simplifyRefValue(k, uVal)})
			}
			continue
		}
		if cOk && !uOk {
			if cSubMap, ok := cVal.(map[string]any); ok && !strings.HasSuffix(k, "Ref") {
				diffMapsRecursive(cSubMap, nil, path, out)
			} else {
				*out = append(*out, Item{Path: path, CreateVal: simplifyRefValue(k, cVal), UpdateVal: nil})
			}
			continue
		}

		if strings.HasSuffix(k, "Ref") {
			cRef := simplifyRefValue(k, cVal)
			uRef := simplifyRefValue(k, uVal)
			if !ValuesMatch(cRef, uRef) {
				*out = append(*out, Item{Path: path, CreateVal: cRef, UpdateVal: uRef})
			}
			continue
		}

		cSubMap, cIsMap := cVal.(map[string]any)
		uSubMap, uIsMap := uVal.(map[string]any)
		if cIsMap && uIsMap {
			diffMapsRecursive(cSubMap, uSubMap, path, out)
			continue
		}

		if !ValuesMatch(cVal, uVal) {
			*out = append(*out, Item{Path: path, CreateVal: cVal, UpdateVal: uVal})
		}
	}
}

func simplifyRefValue(key string, val any) any {
	if !strings.HasSuffix(key, "Ref") {
		return val
	}
	m, ok := val.(map[string]any)
	if !ok {
		return val
	}
	if ext, ok := m["external"].(string); ok && ext != "" {
		return ext
	}
	if name, ok := m["name"].(string); ok && name != "" {
		return name
	}
	return val
}

func (d *Differ) populateUpdateCoverage(diffs []Item) {
	var updateMuts []*classifier.ClassifiedEntry
	for _, ce := range d.primary {
		if ce.Phase == logparser.PhaseUpdate &&
			ce.Entry.IsMutation() &&
			!ce.Entry.IsDelete() &&
			ce.Entry.Response.IsSuccess() {
			updateMuts = append(updateMuts, ce)
		}
	}

	allClientSide := true
	for _, item := range diffs {
		if !IsClientSideKRMField(item.Path) {
			allClientSide = false
			break
		}
	}
	if allClientSide {
		for i := range diffs {
			diffs[i].SentInUpdateRPC = true
			diffs[i].UpdateRPCDetail = "N/A (KCC client-side policy)"
		}
		return
	}

	if len(updateMuts) == 0 {
		for i := range diffs {
			diffs[i].SentInUpdateRPC = false
			diffs[i].UpdateRPCDetail = "No (0 Update RPCs)"
		}
		return
	}

	for i := range diffs {
		item := &diffs[i]
		if IsClientSideKRMField(item.Path) {
			item.SentInUpdateRPC = true
			item.UpdateRPCDetail = "N/A (KCC client-side policy)"
			continue
		}
		var matchedRpcs []string
		for _, mut := range updateMuts {
			body, _ := mut.Entry.Request.ParseJSONBody()
			if requestMatchesDiffItem(mut, body, *item) {
				matchedRpcs = append(matchedRpcs, fmt.Sprintf("#%d `%s`", mut.Index, mut.Entry.Request.Method))
			}
		}
		if len(matchedRpcs) > 0 {
			item.SentInUpdateRPC = true
			item.UpdateRPCDetail = fmt.Sprintf("Yes (%s)", strings.Join(matchedRpcs, ", "))
		} else {
			item.SentInUpdateRPC = false
			item.UpdateRPCDetail = "No"
		}
	}
}

func requestMatchesDiffItem(ce *classifier.ClassifiedEntry, body map[string]any, item Item) bool {
	hasValueAtPath := false
	// 1. Check JSON request body
	if body != nil {
		if foundVal, ok := findValueForDiffPath(body, item.Path); ok {
			hasValueAtPath = true
			if item.UpdateVal == nil {
				return foundVal == nil
			}
			if valuesCompatible(foundVal, item.UpdateVal, item.CreateVal, item.Path) {
				return true
			}
		} else if isZeroOrEmptyValue(item.UpdateVal) {
			return true
		}

		if !hasValueAtPath {
			if item.UpdateVal != nil && containsLeafValue(body, item.Path, item.UpdateVal) {
				return true
			}
			if item.UpdateVal != nil && !ValuesMatch(item.CreateVal, item.UpdateVal) && containsDistinctValueAnywhere(body, item.UpdateVal) {
				return true
			}
		}
	}

	// 2. Check updateMask coverage (important for cleared fields or custom/wrapped structures)
	if len(ce.UpdateMask) > 0 && updateMaskCoversPath(ce.UpdateMask, item.Path) {
		if item.UpdateVal == nil {
			return true
		}
		if body != nil && !hasValueAtPath && !containsOldValueAtPath(body, item.Path, item.CreateVal) {
			return true
		}
	}

	// 3. Check URL query parameters or custom action verb matching the field name
	uPath, query, hasQuery := strings.Cut(ce.Entry.Request.URL, "?")
	leafKey := lastMeaningfulPathSegment(item.Path)
	if colonIdx := strings.LastIndex(uPath, ":"); colonIdx != -1 {
		actionVerb := uPath[colonIdx+1:]
		if strings.Contains(normalizeKey(actionVerb), normalizeKey(leafKey)) {
			return true
		}
	}
	if hasQuery && query != "" {
		if vals, err := url.ParseQuery(query); err == nil {
			for k, vs := range vals {
				if normalizeKey(k) == normalizeKey(leafKey) && len(vs) > 0 {
					return true
				}
			}
		}
	}

	return false
}

func (d *Differ) populateFinalStateCoverage(diffs []Item) {
	finalBodies := collectFinalPhaseResponseBodies(d.primary, logparser.PhaseUpdate)
	for i := range diffs {
		item := &diffs[i]
		if IsClientSideKRMField(item.Path) {
			item.FinalStateMatch = true
			item.FinalStateDetail = "Yes"
			continue
		}
		matched, detail := responseBodiesMatchDiffItem(finalBodies, *item)
		item.FinalStateMatch = matched
		item.FinalStateDetail = detail
	}
}

func collectFinalPhaseResponseBodies(primary []*classifier.ClassifiedEntry, phase string) []map[string]any {
	var bodies []map[string]any
	for i := len(primary) - 1; i >= 0; i-- {
		ce := primary[i]
		if ce.Phase != phase {
			continue
		}
		if ce.LRO != nil && ce.LRO.FinalResponse != nil {
			bodies = append(bodies, ce.LRO.FinalResponse)
		}
		if b, ok := ce.Entry.Response.ParseJSONBody(); ok && b != nil {
			bodies = append(bodies, b)
		}
	}
	return bodies
}

func responseBodiesMatchDiffItem(bodies []map[string]any, item Item) (bool, string) {
	if len(bodies) == 0 {
		return true, "Yes"
	}

	foundIncompatibleAtPath := false
	for _, body := range bodies {
		if foundVal, ok := findValueForDiffPath(body, item.Path); ok {
			if item.UpdateVal == nil {
				if foundVal == nil {
					return true, "Yes"
				}
				foundIncompatibleAtPath = true
				continue
			}
			if valuesCompatible(foundVal, item.UpdateVal, item.CreateVal, item.Path) {
				return true, "Yes"
			}
			foundIncompatibleAtPath = true
		} else if isZeroOrEmptyValue(item.UpdateVal) {
			return true, "Yes"
		}

		if !foundIncompatibleAtPath && item.UpdateVal != nil && containsLeafValue(body, item.Path, item.UpdateVal) {
			return true, "Yes"
		}
	}

	if foundIncompatibleAtPath {
		return false, "No (mismatch)"
	}

	return true, "Yes"
}

func findValueForDiffPath(obj map[string]any, diffPath string) (any, bool) {
	if obj == nil {
		return nil, false
	}

	var segments []string
	if strings.HasPrefix(diffPath, "spec.") {
		segments = strings.Split(strings.TrimPrefix(diffPath, "spec."), ".")
	} else if strings.HasPrefix(diffPath, "metadata.labels.") {
		labelKey := strings.TrimPrefix(diffPath, "metadata.labels.")
		for _, rootKey := range []string{"labels", "gcpLabels", "userLabels"} {
			if v, ok := lookupNormalizedSegments(obj, []string{rootKey, labelKey}); ok {
				return v, true
			}
		}
		segments = []string{"labels", labelKey}
	} else {
		segments = strings.Split(diffPath, ".")
	}

	candidateSegmentLists := [][]string{segments}
	if len(segments) >= 2 && (segments[len(segments)-1] == "value" || segments[len(segments)-1] == "valueFrom") {
		candidateSegmentLists = append(candidateSegmentLists, segments[:len(segments)-1])
	} else if len(segments) >= 3 && segments[len(segments)-2] == "valueFrom" {
		candidateSegmentLists = append(candidateSegmentLists, segments[:len(segments)-2])
	}

	for _, segs := range candidateSegmentLists {
		if v, ok := lookupNormalizedSegments(obj, segs); ok {
			return v, true
		}
		for k, v := range obj {
			if k == "metadata" || k == "error" {
				continue
			}
			if subMap, ok := v.(map[string]any); ok {
				if val, found := lookupNormalizedSegments(subMap, segs); found {
					return val, true
				}
			}
		}
	}

	return nil, false
}

func lookupNormalizedSegments(curr any, segments []string) (any, bool) {
	if len(segments) == 0 {
		return curr, true
	}
	m, ok := curr.(map[string]any)
	if !ok || m == nil {
		return nil, false
	}

	target := normalizeKey(segments[0])
	for k, v := range m {
		if normalizeKey(k) == target {
			return lookupNormalizedSegments(v, segments[1:])
		}
	}
	return nil, false
}

func valuesCompatible(foundVal, expectedVal, createVal any, diffPath string) bool {
	if ValuesMatch(foundVal, expectedVal) {
		return true
	}
	if foundVal == nil || expectedVal == nil {
		return false
	}
	if createVal != nil && ValuesMatch(foundVal, createVal) {
		return false
	}
	// Allow proto enum integer encoding vs string enum name
	if isNumericValue(foundVal) && isEnumLikeString(expectedVal) {
		return true
	}
	if isNumericValue(expectedVal) && isEnumLikeString(foundVal) {
		return true
	}
	if strings.HasSuffix(diffPath, "Ref") {
		return true
	}
	fStr, fIsStr := foundVal.(string)
	eStr, eIsStr := expectedVal.(string)
	if fIsStr && eIsStr {
		fl := strings.ToLower(fStr)
		el := strings.ToLower(eStr)
		if fl != "" && el != "" && (strings.Contains(fl, el) || strings.Contains(el, fl)) {
			return true
		}
		return false
	}
	return false
}

func isNumericValue(v any) bool {
	switch v.(type) {
	case int, int32, int64, float32, float64:
		return true
	}
	return false
}

func isEnumLikeString(v any) bool {
	s, ok := v.(string)
	return ok && s != ""
}

func containsLeafValue(curr any, diffPath string, expectedVal any) bool {
	leafKey := normalizeKey(lastMeaningfulPathSegment(diffPath))
	return searchTreeForKeyAndValue(curr, leafKey, expectedVal)
}

func searchTreeForKeyAndValue(curr any, normLeafKey string, expectedVal any) bool {
	switch v := curr.(type) {
	case map[string]any:
		for k, child := range v {
			if normalizeKey(k) == normLeafKey && ValuesMatch(child, expectedVal) {
				return true
			}
			if searchTreeForKeyAndValue(child, normLeafKey, expectedVal) {
				return true
			}
		}
	case []any:
		for _, elem := range v {
			if searchTreeForKeyAndValue(elem, normLeafKey, expectedVal) {
				return true
			}
		}
	}
	return false
}

func containsDistinctValueAnywhere(curr any, expectedVal any) bool {
	if _, isBool := expectedVal.(bool); isBool {
		return false
	}
	switch v := curr.(type) {
	case map[string]any:
		for _, child := range v {
			if ValuesMatch(child, expectedVal) {
				return true
			}
			if containsDistinctValueAnywhere(child, expectedVal) {
				return true
			}
		}
	case []any:
		if ValuesMatch(v, expectedVal) {
			return true
		}
		for _, elem := range v {
			if ValuesMatch(elem, expectedVal) || containsDistinctValueAnywhere(elem, expectedVal) {
				return true
			}
		}
	}
	return false
}

func containsOldValueAtPath(body map[string]any, diffPath string, createVal any) bool {
	if createVal == nil {
		return false
	}
	found, ok := findValueForDiffPath(body, diffPath)
	if !ok {
		return false
	}
	return ValuesMatch(found, createVal)
}

func updateMaskCoversPath(mask []string, diffPath string) bool {
	clean := strings.TrimPrefix(diffPath, "spec.")
	if strings.HasPrefix(diffPath, "metadata.labels.") {
		clean = "labels"
	}
	parts := strings.Split(clean, ".")
	if len(parts) >= 2 && (parts[len(parts)-1] == "value" || parts[len(parts)-1] == "valueFrom") {
		parts = parts[:len(parts)-1]
	}
	normParts := make([]string, len(parts))
	for i, p := range parts {
		normParts[i] = normalizeKey(p)
	}

	for _, m := range mask {
		if m == "*" {
			return true
		}
		mParts := strings.Split(m, ".")
		// A mask path with more segments than diffPath only covers a child sub-field,
		// not the target field itself (critical when checking whether a parent field was cleared).
		if len(mParts) > len(normParts) {
			continue
		}
		match := true
		for i, mp := range mParts {
			if normalizeKey(mp) != normParts[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func normalizeKey(k string) string {
	k = strings.TrimSuffix(k, "Ref")
	k = strings.ReplaceAll(k, "_", "")
	k = strings.ReplaceAll(k, "-", "")
	return strings.ToLower(k)
}

func isZeroOrEmptyValue(v any) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case string:
		return val == ""
	case bool:
		return !val
	case int, int64, float64:
		return fmt.Sprintf("%v", val) == "0"
	case []any:
		return len(val) == 0
	case map[string]any:
		return len(val) == 0
	}
	return false
}

func lastMeaningfulPathSegment(path string) string {
	parts := strings.Split(path, ".")
	if len(parts) >= 3 && parts[len(parts)-2] == "valueFrom" {
		return parts[len(parts)-3]
	}
	if len(parts) >= 2 && (parts[len(parts)-1] == "value" || parts[len(parts)-1] == "valueFrom") {
		return parts[len(parts)-2]
	}
	return parts[len(parts)-1]
}

// IsClientSideKRMField returns true if path is a KCC-internal policy field not sent to GCP APIs.
func IsClientSideKRMField(path string) bool {
	switch path {
	case "spec.deletionPolicy", "spec.skipDelete", "spec.retainOnDelete", "spec.managementPolicy":
		return true
	}
	return false
}

// ValuesMatch returns true if a and b represent equivalent scalar or composite JSON/YAML values.
func ValuesMatch(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if reflect.DeepEqual(a, b) {
		return true
	}
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)
	if aStr == bStr {
		return true
	}
	aj, err1 := json.Marshal(a)
	bj, err2 := json.Marshal(b)
	if err1 == nil && err2 == nil && string(aj) == string(bj) {
		return true
	}
	return false
}

// FormatCompactValue renders a value as a compact backtick-wrapped JSON snippet for tables and diagnostics.
func FormatCompactValue(v any) string {
	if v == nil {
		return "*(unset)*"
	}
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("`%v`", v)
	}
	s := formatJSONWithSpaces(string(b))
	if len(s) > 64 {
		s = s[:61] + "..."
	}
	return fmt.Sprintf("`%s`", s)
}

func formatJSONWithSpaces(s string) string {
	var out strings.Builder
	inString := false
	escaped := false
	for i := 0; i < len(s); i++ {
		ch := s[i]
		out.WriteByte(ch)
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' && inString {
			escaped = true
			continue
		}
		if ch == '"' {
			inString = !inString
			continue
		}
		if !inString && (ch == ',' || ch == ':') {
			if i+1 < len(s) && s[i+1] != ' ' {
				out.WriteByte(' ')
			}
		}
	}
	return out.String()
}
