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
	"strconv"
	"strings"
	"time"
	"unicode"

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
	rawDiffs := DiffObjects(d.fix.PrimaryCreate.Obj, d.fix.PrimaryUpdate.Obj)
	if len(rawDiffs) == 0 {
		return []Item{}
	}

	var diffs []Item
	for _, item := range rawDiffs {
		// When a test runs with non-SSA create, omitting a field or map key in update.yaml
		// does not remove it from the Kubernetes object in kube-apiserver (it remains in GoldenObj).
		// Skip such retained fields so we do not falsely blame the controller for not clearing them.
		if item.UpdateVal == nil && d.fix.GoldenObj != nil && goldenObjectRetainsValue(d.fix.GoldenObj, item.Path, item.CreateVal) {
			continue
		}
		diffs = append(diffs, item)
	}
	if len(diffs) == 0 {
		return []Item{}
	}

	d.populateUpdateCoverage(diffs)
	d.populateFinalStateCoverage(diffs)
	return diffs
}

func goldenObjectRetainsValue(goldenObj *unstructured.Unstructured, diffPath string, createVal any) bool {
	if goldenObj == nil || createVal == nil {
		return false
	}
	val, found := lookupExactGoldenField(goldenObj.Object, diffPath)
	if !found || val == nil {
		return false
	}
	parts := strings.Split(diffPath, ".")
	if last := parts[len(parts)-1]; strings.HasSuffix(last, "Ref") {
		val = simplifyRefValue(last, val)
	}
	return ValuesMatch(val, createVal)
}

func lookupExactGoldenField(obj map[string]any, diffPath string) (any, bool) {
	if obj == nil || diffPath == "" {
		return nil, false
	}
	if strings.HasPrefix(diffPath, "metadata.labels.") {
		labelKey := strings.TrimPrefix(diffPath, "metadata.labels.")
		labels, found, err := unstructured.NestedMap(obj, "metadata", "labels")
		if err != nil || !found {
			return nil, false
		}
		v, ok := labels[labelKey]
		return v, ok
	}
	return lookupExactDottedSegments(obj, strings.Split(diffPath, "."))
}

func lookupExactDottedSegments(curr any, segments []string) (any, bool) {
	if len(segments) == 0 {
		return curr, true
	}
	m, ok := curr.(map[string]any)
	if !ok || m == nil {
		return nil, false
	}
	if joined := strings.Join(segments, "."); joined != segments[0] {
		if v, exists := m[joined]; exists {
			return v, true
		}
	}
	v, exists := m[segments[0]]
	if !exists {
		return nil, false
	}
	return lookupExactDottedSegments(v, segments[1:])
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
	if createSpec != nil && updateSpec != nil {
		cResID := createSpec["resourceID"]
		uResID := updateSpec["resourceID"]
		if isResourceIDEquivalent(cResID, uResID, createObj.GetName(), updateObj.GetName()) {
			cCopy := make(map[string]any, len(createSpec))
			for k, v := range createSpec {
				if k != "resourceID" {
					cCopy[k] = v
				}
			}
			createSpec = cCopy
			uCopy := make(map[string]any, len(updateSpec))
			for k, v := range updateSpec {
				if k != "resourceID" {
					uCopy[k] = v
				}
			}
			updateSpec = uCopy
		}
	}
	diffMapsRecursive(createSpec, updateSpec, "spec", &diffs)

	createLabels, _, _ := unstructured.NestedMap(createObj.Object, "metadata", "labels")
	updateLabels, _, _ := unstructured.NestedMap(updateObj.Object, "metadata", "labels")
	diffMapsRecursive(createLabels, updateLabels, "metadata.labels", &diffs)

	return diffs
}

func isResourceIDEquivalent(cVal, uVal any, cName, uName string) bool {
	cStr, _ := cVal.(string)
	uStr, _ := uVal.(string)
	if cStr == "" {
		cStr = cName
	}
	if uStr == "" {
		uStr = uName
	}
	return cStr != "" && cStr == uStr
}

func isLabelOrAnnotationPath(path string) bool {
	return strings.HasPrefix(path, "metadata.labels.") ||
		strings.HasPrefix(path, "spec.labels.") ||
		strings.HasPrefix(path, "spec.annotations.") ||
		strings.HasPrefix(path, "metadata.annotations.")
}

func isRefKey(k string) bool {
	return strings.HasSuffix(k, "Ref") || strings.HasSuffix(k, "Refs")
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
			if uSubMap, ok := uVal.(map[string]any); ok && !isRefKey(k) {
				diffMapsRecursive(nil, uSubMap, path, out)
			} else {
				*out = append(*out, Item{Path: path, CreateVal: nil, UpdateVal: simplifyRefValue(k, uVal)})
			}
			continue
		}
		if cOk && !uOk {
			if isLabelOrAnnotationPath(path) {
				// Removed label or annotation: in KCC/GCP, omitting an existing label/annotation
				// in update.yaml preserves the existing value and is not treated as a mutation.
				continue
			}
			if cSubMap, ok := cVal.(map[string]any); ok && !isRefKey(k) {
				diffMapsRecursive(cSubMap, nil, path, out)
			} else {
				*out = append(*out, Item{Path: path, CreateVal: simplifyRefValue(k, cVal), UpdateVal: nil})
			}
			continue
		}

		if isRefKey(k) {
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
	if !isRefKey(key) {
		return val
	}
	if sList, ok := val.([]any); ok {
		simplified := make([]any, len(sList))
		for i, item := range sList {
			simplified[i] = simplifyRefValue(key, item)
		}
		return simplified
	}
	if m, ok := val.(map[string]any); ok {
		if ext, ok := m["external"].(string); ok && ext != "" {
			return ext
		}
		if name, ok := m["name"].(string); ok && name != "" {
			return name
		}
		if len(m) == 1 {
			for _, v := range m {
				if s, ok := v.(string); ok {
					return s
				}
			}
		}
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
	if isLabelOrAnnotationPath(item.Path) && item.UpdateVal == nil {
		return true
	}
	if item.Path == "spec.resourceID" {
		return true
	}

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
	if isLabelOrAnnotationPath(item.Path) && item.UpdateVal == nil {
		return true, "Yes"
	}
	if item.Path == "spec.resourceID" {
		return true, "Yes"
	}
	if len(bodies) == 0 {
		if item.SentInUpdateRPC || isZeroOrEmptyValue(item.UpdateVal) {
			return true, "Yes"
		}
		return false, "No (mismatch)"
	}

	foundIncompatibleAtPath := false
	foundNonRedactedMismatch := false
	for _, body := range bodies {
		if foundVal, ok := findValueForDiffPath(body, item.Path); ok {
			if item.UpdateVal == nil {
				if foundVal == nil {
					return true, "Yes"
				}
				foundIncompatibleAtPath = true
				foundNonRedactedMismatch = true
				continue
			}
			if valuesCompatible(foundVal, item.UpdateVal, item.CreateVal, item.Path) {
				return true, "Yes"
			}
			foundIncompatibleAtPath = true
			if !isRedactedOrEmptySecretValue(foundVal) {
				foundNonRedactedMismatch = true
			}
		} else if isZeroOrEmptyValue(item.UpdateVal) {
			return true, "Yes"
		}

		if !foundIncompatibleAtPath && item.UpdateVal != nil && containsLeafValue(body, item.Path, item.UpdateVal) {
			return true, "Yes"
		}
		if !foundIncompatibleAtPath && item.UpdateVal != nil && containsDistinctValueAnywhere(body, item.UpdateVal) {
			return true, "Yes"
		}
	}

	if item.SentInUpdateRPC && !foundNonRedactedMismatch && isWriteOnlyFieldPath(item.Path, item.UpdateVal) {
		return true, "Yes (write-only)"
	}

	if foundIncompatibleAtPath || !item.SentInUpdateRPC {
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
			if subSlice, ok := v.([]any); ok {
				for _, elem := range subSlice {
					if elemMap, ok := elem.(map[string]any); ok {
						if val, found := lookupNormalizedSegments(elemMap, segs); found {
							return val, true
						}
					}
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
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if normalizeKey(k) == target {
			if val, found := lookupNormalizedSegments(m[k], segments[1:]); found {
				return val, true
			}
		}
	}
	if strings.HasSuffix(segments[0], "Ref") {
		for _, k := range keys {
			if refKeysMatch(target, normalizeKey(k)) {
				if val, found := lookupNormalizedSegments(m[k], segments[1:]); found {
					return val, true
				}
			}
		}
	}
	return nil, false
}

func refKeysMatch(normRefKey, normProtoKey string) bool {
	if normRefKey == "" || normProtoKey == "" {
		return false
	}
	trimAffixes := func(s string) string {
		for _, p := range []string{"kms", "gcp", "cloud"} {
			if strings.HasPrefix(s, p) && len(s)-len(p) >= 3 {
				s = strings.TrimPrefix(s, p)
			}
		}
		for _, suf := range []string{"name", "id", "uri", "url", "path"} {
			if strings.HasSuffix(s, suf) && len(s)-len(suf) >= 3 {
				s = strings.TrimSuffix(s, suf)
			}
		}
		return s
	}
	a := trimAffixes(normRefKey)
	b := trimAffixes(normProtoKey)
	return a != "" && a == b
}

func isWriteOnlyFieldPath(diffPath string, expectedVal any) bool {
	switch expectedVal.(type) {
	case bool, int, int32, int64, float32, float64:
		return false
	}
	leaf := strings.ToLower(lastMeaningfulPathSegment(diffPath))
	if leaf == "" || strings.HasSuffix(leaf, "ref") {
		return false
	}
	return strings.HasSuffix(leaf, "password") ||
		strings.HasSuffix(leaf, "passphrase") ||
		strings.HasSuffix(leaf, "privatekey") ||
		strings.HasSuffix(leaf, "secret") ||
		strings.HasSuffix(leaf, "secretkey") ||
		strings.HasSuffix(leaf, "apikey") ||
		strings.HasSuffix(leaf, "token")
}

func isRedactedOrEmptySecretValue(v any) bool {
	if v == nil {
		return true
	}
	s, ok := v.(string)
	if !ok {
		return false
	}
	trimmed := strings.TrimSpace(s)
	if trimmed == "" || strings.Trim(trimmed, "*") == "" {
		return true
	}
	lower := strings.ToLower(trimmed)
	return lower == "redacted" || lower == "[redacted]" || lower == "<redacted>"
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
	if isEnumLikeOrSlice(expectedVal) && isEnumLikeOrSlice(createVal) && !reflect.DeepEqual(expectedVal, createVal) {
		if jsonContainsExpected(foundVal, expectedVal) {
			return true
		}
	} else if jsonContainsExpected(foundVal, expectedVal) && (createVal == nil || !jsonContainsExpected(foundVal, createVal)) {
		return true
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
		if len(fStr) > 0 && strings.Trim(fStr, "*") == "" {
			return true
		}
		if isPlaceholder(fStr) || isPlaceholder(eStr) {
			return true
		}
		fl := strings.ToLower(fStr)
		el := strings.ToLower(eStr)
		if fl != "" && el != "" && (strings.Contains(fl, el) || strings.Contains(el, fl)) {
			return true
		}
		return false
	}
	return false
}

func isPlaceholder(s string) bool {
	return strings.HasPrefix(s, "${") && strings.HasSuffix(s, "}") && len(s) > 3 && !strings.Contains(s[2:len(s)-1], " ")
}

func isEnumLikeOrSlice(v any) bool {
	if isEnumLikeString(v) {
		return true
	}
	if slice, ok := v.([]any); ok && len(slice) > 0 {
		for _, elem := range slice {
			if !isEnumLikeString(elem) {
				return false
			}
		}
		return true
	}
	return false
}

func jsonContainsExpected(actual, expected any) bool {
	if actual == nil && expected == nil {
		return true
	}
	if actual == nil || expected == nil {
		return false
	}
	if aMap, aOk := actual.(map[string]any); aOk {
		if eSlice, eIsSlice := expected.([]any); eIsSlice {
			return mapContainsKeyedSlice(aMap, eSlice)
		}
		eMap, eOk := expected.(map[string]any)
		if !eOk {
			return false
		}
		aNorm := make(map[string]any, len(aMap))
		for k, v := range aMap {
			aNorm[normalizeKey(k)] = v
		}
		for k, eVal := range eMap {
			nk := normalizeKey(k)
			aVal, exists := aNorm[nk]
			if !exists {
				if isOmissibleEmptyCollection(eVal) || isZeroOrEmptyValue(eVal) {
					continue
				}
				return false
			}
			if isOmissibleEmptyCollection(eVal) {
				if !isOmissibleEmptyCollection(aVal) && !isZeroOrEmptyValue(aVal) {
					return false
				}
				continue
			}
			eVal = simplifyRefValue(k, eVal)
			if !jsonContainsExpected(aVal, eVal) {
				return false
			}
		}
		return true
	}
	if aSlice, aOk := actual.([]any); aOk {
		eSlice, eOk := expected.([]any)
		if !eOk || len(aSlice) != len(eSlice) {
			return false
		}
		for i := range aSlice {
			if !jsonContainsExpected(aSlice[i], eSlice[i]) {
				return false
			}
		}
		return true
	}
	if isNumericValue(actual) && isEnumLikeString(expected) {
		return true
	}
	if isNumericValue(expected) && isEnumLikeString(actual) {
		return true
	}
	return ValuesMatch(actual, expected)
}

func mapContainsKeyedSlice(actualMap map[string]any, expectedSlice []any) bool {
	if len(expectedSlice) == 0 {
		return len(actualMap) == 0
	}
	if len(actualMap) != len(expectedSlice) {
		return false
	}
	seen := make(map[string]struct{}, len(expectedSlice))
	for _, elem := range expectedSlice {
		elemMap, ok := elem.(map[string]any)
		if !ok {
			return false
		}
		var keyField, keyVal string
		for _, kField := range []string{"id", "name", "key"} {
			if s, isStr := elemMap[kField].(string); isStr && s != "" {
				keyField = kField
				keyVal = s
				break
			}
		}
		if keyVal == "" {
			return false
		}
		if _, dup := seen[keyVal]; dup {
			return false
		}
		seen[keyVal] = struct{}{}
		actualElem, exists := actualMap[keyVal]
		if !exists {
			return false
		}
		if actualElemMap, isMap := actualElem.(map[string]any); isMap {
			expectedForCompare := elemMap
			if _, hasKeyInActual := actualElemMap[keyField]; !hasKeyInActual {
				stripped := make(map[string]any, len(elemMap))
				for k, v := range elemMap {
					if k != keyField {
						stripped[k] = v
					}
				}
				expectedForCompare = stripped
			}
			if !jsonContainsExpected(actualElemMap, expectedForCompare) {
				return false
			}
			continue
		}
		if valField, hasVal := elemMap["value"]; hasVal && len(elemMap) == 2 {
			if !jsonContainsExpected(actualElem, valField) {
				return false
			}
			continue
		}
		return false
	}
	return true
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
	if !ok || s == "" {
		return false
	}
	hasLetter := false
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
		if unicode.IsLetter(r) {
			hasLetter = true
		}
	}
	return hasLetter
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
			nmp := normalizeKey(mp)
			if nmp != normParts[i] && !(strings.HasSuffix(parts[i], "Ref") && refKeysMatch(normParts[i], nmp)) {
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
	case "spec.deletionPolicy", "spec.skipDelete", "spec.retainOnDelete", "spec.managementPolicy", "spec.resourceID":
		return true
	}
	return false
}

// IsWriteOnlyField returns true if path corresponds to a write-only or secret field
// (e.g. passwords, secrets, private keys) that GCP APIs do not return in GET responses.
func IsWriteOnlyField(path string) bool {
	lower := strings.ToLower(path)
	segments := strings.Split(lower, ".")
	for i, seg := range segments {
		if strings.Contains(seg, "password") {
			if !strings.Contains(seg, "policy") {
				return true
			}
		}
		if seg == "secret" || seg == "secretkey" || seg == "clientsecret" || seg == "webhooksecret" {
			return true
		}
		if strings.Contains(seg, "privatekey") || strings.Contains(seg, "private_key") {
			return true
		}
		if seg == "value" && i > 0 {
			parent := segments[i-1]
			if strings.Contains(parent, "password") || strings.Contains(parent, "secret") || strings.Contains(parent, "credential") || strings.Contains(parent, "token") {
				return true
			}
		}
	}
	return false
}

func parseDuration(v any) (time.Duration, bool) {
	switch val := v.(type) {
	case string:
		val = strings.TrimSpace(val)
		if val == "" {
			return 0, false
		}
		if d, err := time.ParseDuration(val); err == nil {
			return d, true
		}
		if strings.HasSuffix(val, "s") {
			secStr := strings.TrimSuffix(val, "s")
			if f, err := strconv.ParseFloat(secStr, 64); err == nil {
				return time.Duration(f * float64(time.Second)), true
			}
		}
	}
	return 0, false
}

func toFloat(v any) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case int32:
		return float64(val), true
	case string:
		val = strings.TrimSpace(val)
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f, true
		}
	}
	return 0, false
}

func areDurationsEqual(a, b any) bool {
	d1, ok1 := parseDuration(a)
	d2, ok2 := parseDuration(b)
	if ok1 && ok2 {
		return d1 == d2
	}
	if ok1 {
		if f, isNum := toFloat(b); isNum {
			return d1 == time.Duration(f*float64(time.Second))
		}
	}
	if ok2 {
		if f, isNum := toFloat(a); isNum {
			return d2 == time.Duration(f*float64(time.Second))
		}
	}
	return false
}

func areNumbersEqual(a, b any) bool {
	f1, ok1 := toFloat(a)
	f2, ok2 := toFloat(b)
	if ok1 && ok2 {
		return f1 == f2
	}
	return false
}

func stripResourcePrefixes(s string) string {
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "//")
	if slashIdx := strings.Index(s, "/"); slashIdx != -1 && strings.Contains(s[:slashIdx], ".googleapis.com") {
		s = s[slashIdx+1:]
	}
	return strings.TrimPrefix(s, "/")
}

func isReferenceMatch(expected, actual string) bool {
	if expected == actual {
		return true
	}
	if expected == "" || actual == "" {
		return false
	}
	cleanExpected := stripResourcePrefixes(expected)
	cleanActual := stripResourcePrefixes(actual)
	if cleanExpected == cleanActual {
		return true
	}
	if strings.HasSuffix(cleanActual, "/"+cleanExpected) || strings.HasSuffix(cleanExpected, "/"+cleanActual) {
		return true
	}
	if strings.HasPrefix(cleanActual, cleanExpected+"@") || strings.HasPrefix(cleanExpected, cleanActual+"@") {
		return true
	}
	return false
}

func areReferencesEqual(a, b any) bool {
	aStr, aOk := a.(string)
	bStr, bOk := b.(string)
	if !aOk || !bOk {
		return false
	}
	if strings.Contains(aStr, "/") || strings.Contains(bStr, "/") ||
		strings.Contains(aStr, "@") || strings.Contains(bStr, "@") ||
		strings.Contains(aStr, "${") || strings.Contains(bStr, "${") {
		return isReferenceMatch(aStr, bStr) || isReferenceMatch(bStr, aStr)
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
	if aStr, aOk := a.(string); aOk {
		if bStr, bOk := b.(string); bOk {
			if isPlaceholder(aStr) || isPlaceholder(bStr) {
				return true
			}
		}
	}
	if areDurationsEqual(a, b) {
		return true
	}
	if areNumbersEqual(a, b) {
		return true
	}
	if areReferencesEqual(a, b) {
		return true
	}
	aj, err1 := json.Marshal(a)
	bj, err2 := json.Marshal(b)
	if err1 == nil && err2 == nil && string(aj) == string(bj) {
		return true
	}
	return jsonValuesEquivalent(a, b)
}

func jsonValuesEquivalent(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if aMap, aOk := a.(map[string]any); aOk {
		bMap, bOk := b.(map[string]any)
		if !bOk {
			return false
		}
		aNorm := make(map[string]any, len(aMap))
		for k, v := range aMap {
			aNorm[normalizeKey(k)] = v
		}
		bNorm := make(map[string]any, len(bMap))
		for k, v := range bMap {
			bNorm[normalizeKey(k)] = v
		}
		for k, aVal := range aNorm {
			bVal, exists := bNorm[k]
			if !exists {
				if isOmissibleEmptyCollection(aVal) {
					continue
				}
				return false
			}
			if !jsonValuesEquivalent(aVal, bVal) {
				return false
			}
		}
		for k, bVal := range bNorm {
			if _, exists := aNorm[k]; !exists {
				if isOmissibleEmptyCollection(bVal) {
					continue
				}
				return false
			}
		}
		return true
	}
	if aSlice, aOk := a.([]any); aOk {
		bSlice, bOk := b.([]any)
		if !bOk || len(aSlice) != len(bSlice) {
			return false
		}
		for i := range aSlice {
			if !jsonValuesEquivalent(aSlice[i], bSlice[i]) {
				return false
			}
		}
		return true
	}
	if areDurationsEqual(a, b) {
		return true
	}
	if areNumbersEqual(a, b) {
		return true
	}
	if areReferencesEqual(a, b) {
		return true
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func isOmissibleEmptyCollection(v any) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case []any:
		return len(val) == 0
	case map[string]any:
		return len(val) == 0
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
