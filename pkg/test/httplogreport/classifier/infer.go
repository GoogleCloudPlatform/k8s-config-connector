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

package classifier

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type updatedLeaf struct {
	normKey   string
	updateVal any
}

// classifyByInference implements Tier 2 offline phase boundary inference for legacy
// _http.log files that do not yet contain explicit "# Phase:" comments.
func (c *Classifier) classifyByInference(classified []*ClassifiedEntry) {
	n := len(classified)
	if n == 0 {
		return
	}

	// 1. Identify the start of the DELETE phase after the last non-delete primary mutation.
	lastPrimaryNonDeleteMutIdx := -1
	for i, ce := range classified {
		if ce.Role == RolePrimary && !ce.IsLROPoll && ce.Entry.IsMutation() && !ce.Entry.IsDelete() {
			lastPrimaryNonDeleteMutIdx = i
		}
	}

	deleteStartIdx := n
	for i := lastPrimaryNonDeleteMutIdx + 1; i < n; i++ {
		ce := classified[i]
		if !ce.IsLROPoll && ce.Entry.IsDelete() {
			deleteStartIdx = i
			break
		}
	}

	// 2. Identify the start of the UPDATE phase among entries [0, deleteStartIdx).
	updateStartIdx := deleteStartIdx
	if c.fix != nil && c.fix.PrimaryUpdate != nil {
		var primaryMutIndices []int
		for i := 0; i < deleteStartIdx; i++ {
			ce := classified[i]
			if ce.Role == RolePrimary && !ce.IsLROPoll && ce.Entry.IsMutation() && !ce.Entry.IsDelete() {
				primaryMutIndices = append(primaryMutIndices, i)
			}
		}

		if len(primaryMutIndices) >= 2 {
			firstCreateMutIdx := primaryMutIndices[0]
			foundUpdateStart := -1
			updatedLeaves := c.extractUpdatedLeaves()

			// Signal 1: Check if any mutation after firstCreateMutIdx carries updated values from update.yaml.
			if len(updatedLeaves) > 0 {
				for _, mutIdx := range primaryMutIndices[1:] {
					if entryCarriesUpdatedLeaf(classified[mutIdx], updatedLeaves) {
						foundUpdateStart = mutIdx
						break
					}
				}
			}

			// Signal 2: Check for a steady-state primary GET between firstCreateMutIdx and the candidate mutation.
			if foundUpdateStart == -1 {
				for _, mutIdx := range primaryMutIndices[1:] {
					if hasSteadyStatePrimaryGetBetween(classified, firstCreateMutIdx, mutIdx) {
						foundUpdateStart = mutIdx
						break
					}
				}
			}

			// Fallback: if update.yaml has changes and there are >= 2 mutations, the second mutation starts UPDATE.
			if foundUpdateStart == -1 && len(updatedLeaves) > 0 {
				foundUpdateStart = primaryMutIndices[1]
			}

			if foundUpdateStart != -1 {
				updateStartIdx = foundUpdateStart
			}
		}
	}

	// 3. Assign phases to all entries, keeping LRO polls in the same phase as their parent mutation.
	for i, ce := range classified {
		switch {
		case i >= deleteStartIdx:
			ce.Phase = logparser.PhaseDelete
		case i >= updateStartIdx:
			ce.Phase = logparser.PhaseUpdate
		default:
			ce.Phase = logparser.PhaseCreate
		}
		ce.Entry.Phase = ce.Phase
	}

	for _, ce := range classified {
		if ce.IsLROPoll && ce.ParentIndex > 0 && ce.ParentIndex <= len(classified) {
			parent := classified[ce.ParentIndex-1]
			ce.Phase = parent.Phase
			ce.Entry.Phase = parent.Phase
		}
	}
}

func (c *Classifier) extractUpdatedLeaves() []updatedLeaf {
	if c.fix == nil || c.fix.PrimaryCreate == nil || c.fix.PrimaryUpdate == nil {
		return nil
	}
	createObj := c.fix.PrimaryCreate.Obj
	updateObj := c.fix.PrimaryUpdate.Obj
	if createObj == nil || updateObj == nil {
		return nil
	}

	var out []updatedLeaf
	cSpec, _, _ := unstructured.NestedMap(createObj.Object, "spec")
	uSpec, _, _ := unstructured.NestedMap(updateObj.Object, "spec")
	collectUpdatedLeavesRecursive(cSpec, uSpec, &out)

	cLabels, _, _ := unstructured.NestedMap(createObj.Object, "metadata", "labels")
	uLabels, _, _ := unstructured.NestedMap(updateObj.Object, "metadata", "labels")
	collectUpdatedLeavesRecursive(cLabels, uLabels, &out)
	return out
}

func collectUpdatedLeavesRecursive(cMap, uMap map[string]any, out *[]updatedLeaf) {
	for k, uVal := range uMap {
		if uVal == nil {
			continue
		}
		cVal, cOk := cMap[k]
		uSubMap, uIsMap := uVal.(map[string]any)
		if uIsMap && !strings.HasSuffix(k, "Ref") {
			var cSubMap map[string]any
			if cOk {
				cSubMap, _ = cVal.(map[string]any)
			}
			collectUpdatedLeavesRecursive(cSubMap, uSubMap, out)
			continue
		}
		if !cOk || !leafValuesEqual(cVal, uVal) {
			*out = append(*out, updatedLeaf{
				normKey:   normalizeInferKey(k),
				updateVal: uVal,
			})
		}
	}
}

func entryCarriesUpdatedLeaf(ce *ClassifiedEntry, leaves []updatedLeaf) bool {
	body, ok := ce.Entry.Request.ParseJSONBody()
	if !ok || body == nil {
		return false
	}
	for _, lf := range leaves {
		if treeContainsKeyAndValue(body, lf.normKey, lf.updateVal) {
			return true
		}
	}
	return false
}

func treeContainsKeyAndValue(curr any, normKey string, expectedVal any) bool {
	switch v := curr.(type) {
	case map[string]any:
		for k, child := range v {
			if normalizeInferKey(k) == normKey && leafValuesEqual(child, expectedVal) {
				return true
			}
			if treeContainsKeyAndValue(child, normKey, expectedVal) {
				return true
			}
		}
	case []any:
		for _, elem := range v {
			if treeContainsKeyAndValue(elem, normKey, expectedVal) {
				return true
			}
		}
	}
	return false
}

func hasSteadyStatePrimaryGetBetween(classified []*ClassifiedEntry, firstMutIdx, nextMutIdx int) bool {
	for i := firstMutIdx + 1; i < nextMutIdx; i++ {
		ce := classified[i]
		if ce.Role == RolePrimary && !ce.IsLROPoll && ce.Entry.IsReadOnly() && ce.Entry.Response.IsSuccess() {
			return true
		}
	}
	return false
}

func normalizeInferKey(k string) string {
	k = strings.TrimSuffix(k, "Ref")
	k = strings.ReplaceAll(k, "_", "")
	k = strings.ReplaceAll(k, "-", "")
	return strings.ToLower(k)
}

func leafValuesEqual(a, b any) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func (c *Classifier) inferLegacyDependencyEntry(e *logparser.Entry) string {
	if c.fix == nil {
		return ""
	}
	primary := c.fix.PrimaryCreate
	primaryKind := ""
	primaryGroup := ""
	if primary != nil {
		primaryKind = primary.Kind
		primaryGroup = primary.Group
	}

	uPath := e.CleanURLPath()
	isIAMPrimary := strings.HasPrefix(primaryKind, "IAM")

	if !isIAMPrimary && primaryKind != "Project" && primaryKind != "Folder" && primaryKind != "Organization" {
		segments := strings.Split(strings.Trim(uPath, "/"), "/")
		if len(segments) >= 2 && isVersionSegment(segments[0]) {
			segments = segments[1:]
		}
		if len(segments) == 2 && (segments[0] == "projects" || segments[0] == "folders" || segments[0] == "organizations" || segments[0] == "billingAccounts") {
			last := segments[1]
			if strings.Contains(last, ":setIamPolicy") || strings.Contains(last, ":getIamPolicy") || strings.Contains(last, ":testIamPermissions") {
				return "IAMPolicy"
			}
			if strings.Contains(e.Request.URL, "cloudresourcemanager.googleapis.com") {
				return "Project"
			}
		}
	}

	if primaryKind != "Service" && primaryGroup != "serviceusage.cnrm.cloud.google.com" {
		if strings.Contains(e.Request.URL, "serviceusage.googleapis.com") {
			return "Service"
		}
	}

	if len(c.fix.Dependencies) == 0 {
		return ""
	}

	if primary != nil && matchesPrimaryResourceDirectly(e, primary, c.fix.Dependencies) {
		return ""
	}

	return c.matchDependencyFromList(e, uPath, primary, primaryKind, isIAMPrimary)
}

func (c *Classifier) inferExplicitDependencyKind(e *logparser.Entry) string {
	if c.fix == nil {
		return "Dependency"
	}
	if strings.Contains(e.Request.URL, "serviceusage.googleapis.com") {
		return "Service"
	}
	uPath := e.CleanURLPath()
	primary := c.fix.PrimaryCreate
	primaryKind := ""
	isIAMPrimary := false
	if primary != nil {
		primaryKind = primary.Kind
		isIAMPrimary = strings.HasPrefix(primaryKind, "IAM")
	}
	if kind := c.matchDependencyFromList(e, uPath, primary, primaryKind, isIAMPrimary); kind != "" {
		return kind
	}
	bestKind := ""
	bestIdx := -1
	for _, dep := range c.fix.Dependencies {
		for _, tok := range []string{dep.ResourceID, dep.Name} {
			if tok == "" {
				continue
			}
			if idx := strings.LastIndex(uPath, tok); idx > bestIdx {
				bestIdx = idx
				bestKind = dep.Kind
			} else if bestIdx == -1 && (strings.Contains(e.Request.URL, tok) || strings.Contains(e.Request.Body, tok)) {
				bestKind = dep.Kind
			}
		}
	}
	if bestKind != "" {
		return bestKind
	}
	if len(c.fix.Dependencies) == 1 {
		return c.fix.Dependencies[0].Kind
	}
	return "Dependency"
}

func (c *Classifier) matchDependencyFromList(e *logparser.Entry, uPath string, primary *fixture.KRMResource, primaryKind string, isIAMPrimary bool) string {
	for _, dep := range c.fix.Dependencies {
		if dep.Kind == primaryKind {
			continue
		}
		// If uPath's leaf collection specifically targets a different dependency in the fixture
		// (e.g. /keyRings/<kr>/cryptoKeys targets KMSCryptoKey rather than parent KMSKeyRing),
		// let the more specific child dependency claim it.
		claimedByChildDep := false
		if !urlTargetsKindCollection(uPath, dep.Kind, dep.Group) {
			for _, other := range c.fix.Dependencies {
				if other.Kind != dep.Kind && other.Kind != primaryKind && urlTargetsKindCollection(uPath, other.Kind, other.Group) {
					claimedByChildDep = true
					break
				}
			}
		}
		if claimedByChildDep {
			continue
		}

		var namesToCheck []string
		if len(dep.Name) >= 3 {
			namesToCheck = append(namesToCheck, dep.Name)
		}
		if len(dep.ResourceID) >= 3 && dep.ResourceID != dep.Name {
			namesToCheck = append(namesToCheck, dep.ResourceID)
		}
		namesToCheck = append(namesToCheck, getPlaceholdersForKind(dep.Kind, dep.Group)...)

		for _, depToken := range namesToCheck {
			if matchesDepTokenInURLOrBody(e, uPath, depToken, dep, primary, isIAMPrimary) {
				return dep.Kind
			}
		}

		if matchesDepCollection(uPath, dep, primary) {
			return dep.Kind
		}
	}

	return ""
}

func matchesPrimaryResourceDirectly(e *logparser.Entry, primary *fixture.KRMResource, deps []*fixture.KRMResource) bool {
	uPath := e.CleanURLPath()
	if urlTargetsKindCollection(uPath, primary.Kind, primary.Group) {
		sharedCol := false
		primCol := kindToCollectionName(primary.Kind, primary.Group)
		for _, d := range deps {
			if d.Kind != primary.Kind && strings.EqualFold(kindToCollectionName(d.Kind, d.Group), primCol) {
				sharedCol = true
				break
			}
		}
		if !sharedCol {
			return true
		}
	}

	var primaryTokens []string
	if len(primary.Name) >= 3 {
		primaryTokens = append(primaryTokens, primary.Name)
	}
	if len(primary.ResourceID) >= 3 && primary.ResourceID != primary.Name {
		primaryTokens = append(primaryTokens, primary.ResourceID)
	}
	if primary.ExternalRef != "" {
		primaryTokens = append(primaryTokens, strings.TrimPrefix(primary.ExternalRef, "/"))
	}

	for _, tok := range primaryTokens {
		sharedWithDep := false
		for _, d := range deps {
			if d.Kind != primary.Kind && (d.Name == tok || d.ResourceID == tok) {
				sharedWithDep = true
				break
			}
		}
		if sharedWithDep {
			continue
		}
		if idx := strings.Index(uPath, tok); idx != -1 {
			suffix := strings.TrimPrefix(uPath[idx+len(tok):], "/")
			if suffix == "" || strings.HasPrefix(suffix, ":") || !strings.Contains(suffix, "/") {
				return true
			}
		}
		if _, query, hasQuery := strings.Cut(e.Request.URL, "?"); hasQuery && strings.Contains(query, tok) {
			return true
		}
	}
	return false
}

func matchesDepTokenInURLOrBody(e *logparser.Entry, uPath, depName string, dep, primary *fixture.KRMResource, isIAMPrimary bool) bool {
	depKind := ""
	depGroup := ""
	if dep != nil {
		depKind = dep.Kind
		depGroup = dep.Group
	}
	primaryKind := ""
	primaryGroup := ""
	primaryName := ""
	primaryResID := ""
	if primary != nil {
		primaryKind = primary.Kind
		primaryGroup = primary.Group
		primaryName = primary.Name
		primaryResID = primary.ResourceID
	}
	if !strings.Contains(uPath, depName) {
		if e.Request.Method == "POST" && (strings.Contains(e.Request.Body, depName) || strings.Contains(e.Response.Body, depName) || strings.Contains(e.Request.URL, depName)) {
			if urlTargetsKindCollection(uPath, depKind, depGroup) && !urlTargetsKindCollection(uPath, primaryKind, primaryGroup) {
				return true
			}
			if !urlTargetsKindCollection(uPath, primaryKind, primaryGroup) &&
				(primaryName == "" || (!strings.Contains(e.Request.URL, primaryName) && !strings.Contains(e.Request.Body, primaryName))) &&
				(primaryResID == "" || (!strings.Contains(e.Request.URL, primaryResID) && !strings.Contains(e.Request.Body, primaryResID))) {
				return true
			}
		}
		if e.Request.Method == "GRPC" {
			parts := strings.Split(e.Request.URL, "/")
			if len(parts) > 0 {
				rpcMethod := parts[len(parts)-1]
				shortDep := trimKindServicePrefix(depKind, depGroup)
				shortPrim := trimKindServicePrefix(primaryKind, primaryGroup)
				if strings.Contains(rpcMethod, shortDep) && !strings.Contains(rpcMethod, shortPrim) &&
					(strings.Contains(e.Request.Body, depName) || strings.Contains(e.Response.Body, depName)) {
					return true
				}
			}
		}
		return false
	}

	idx := strings.LastIndex(uPath, depName)
	if idx > 0 {
		before := uPath[idx-1]
		if before != '/' && before != ':' && before != '=' {
			return false
		}
	}
	suffix := uPath[idx+len(depName):]
	if suffix == "" || suffix == "/" || strings.HasPrefix(suffix, ":") {
		return true
	}
	if strings.HasPrefix(suffix, "/operations") || strings.HasPrefix(suffix, "/locations") || strings.HasPrefix(suffix, "/regions") || strings.HasPrefix(suffix, "/zones") || strings.HasPrefix(suffix, "/global") {
		return true
	}
	trimmedSuffix := strings.TrimPrefix(suffix, "/")
	if strings.Contains(trimmedSuffix, "/") {
		if depKind == "KMSCryptoKey" && strings.HasPrefix(trimmedSuffix, "cryptoKeyVersions") && primaryKind != "KMSCryptoKeyVersion" {
			return true
		}
		return false
	}
	if !isIAMPrimary && strings.HasPrefix(suffix, "/iam") {
		return true
	}
	return !urlTargetsKindCollection(uPath, primaryKind, primaryGroup)
}

func matchesDepCollection(uPath string, dep *fixture.KRMResource, primary *fixture.KRMResource) bool {
	if primary == nil || dep == nil || dep.Kind == primary.Kind {
		return false
	}
	depCol := kindToCollectionName(dep.Kind, dep.Group)
	primCol := kindToCollectionName(primary.Kind, primary.Group)
	if depCol == "" || depCol == primCol {
		return false
	}
	lastCol := extractLastCollectionSegment(uPath)
	if lastCol == "" {
		return false
	}
	return strings.EqualFold(lastCol, depCol) && !strings.EqualFold(lastCol, primCol)
}

func urlTargetsKindCollection(uPath string, kind string, group string) bool {
	col := kindToCollectionName(kind, group)
	if col == "" {
		return false
	}
	lastCol := extractLastCollectionSegment(uPath)
	if lastCol == "" {
		return false
	}
	if strings.EqualFold(lastCol, col) {
		return true
	}
	lowerLast := strings.ToLower(lastCol)
	lowerCol := strings.ToLower(col)
	if len(lowerLast) >= 5 && len(lowerCol) >= 5 {
		if strings.HasSuffix(lowerLast, lowerCol) || strings.HasSuffix(lowerCol, lowerLast) {
			return true
		}
	}
	return false
}

func extractLastCollectionSegment(uPath string) string {
	uPath = strings.Split(uPath, "?")[0]
	if colonIdx := strings.Index(uPath, ":"); colonIdx != -1 {
		uPath = uPath[:colonIdx]
	}
	segments := strings.Split(strings.Trim(uPath, "/"), "/")
	if len(segments) == 0 {
		return ""
	}
	i := 0
	for i < len(segments) {
		seg := segments[i]
		if isVersionSegment(seg) || seg == "compute" || seg == "beta" || seg == "alpha" || seg == "sql" || seg == "storage" || seg == "b" || seg == "o" {
			i++
			continue
		}
		break
	}
	lastCol := ""
	for i < len(segments) {
		seg := segments[i]
		if seg == "global" {
			i++
			continue
		}
		lastCol = seg
		if i+1 < len(segments) {
			i += 2
		} else {
			i++
		}
	}
	return lastCol
}

func isVersionSegment(seg string) bool {
	if len(seg) < 2 || seg[0] != 'v' {
		return false
	}
	return seg[1] >= '0' && seg[1] <= '9'
}

func kindToCollectionName(kind string, group string) string {
	short := trimKindServicePrefix(kind, group)
	if short == "" {
		return ""
	}
	lower := strings.ToLower(short[:1]) + short[1:]
	if strings.HasSuffix(lower, "y") && !strings.HasSuffix(lower, "ey") && !strings.HasSuffix(lower, "ay") && !strings.HasSuffix(lower, "oy") {
		return lower[:len(lower)-1] + "ies"
	}
	if strings.HasSuffix(lower, "s") || strings.HasSuffix(lower, "sh") || strings.HasSuffix(lower, "ch") || strings.HasSuffix(lower, "x") {
		return lower + "es"
	}
	return lower + "s"
}

func trimKindServicePrefix(kind string, group string) string {
	if group != "" {
		svc := strings.Split(group, ".")[0]
		if len(svc) > 0 && len(kind) > len(svc) && strings.EqualFold(kind[:len(svc)], svc) {
			return kind[len(svc):]
		}
	}
	for i := 1; i < len(kind); i++ {
		if kind[i] >= 'A' && kind[i] <= 'Z' {
			if kind[i-1] >= 'a' && kind[i-1] <= 'z' {
				return kind[i:]
			}
			if i+1 < len(kind) && kind[i+1] >= 'a' && kind[i+1] <= 'z' {
				return kind[i:]
			}
		}
	}
	return kind
}

func getPlaceholdersForKind(kind string, group string) []string {
	short := trimKindServicePrefix(kind, group)
	if short == "" {
		return nil
	}
	lowerCamel := strings.ToLower(short[:1]) + short[1:]
	out := []string{
		"${" + lowerCamel + "ID}",
		"${" + lowerCamel + "Number}",
		"${" + lowerCamel + "Email}",
	}
	if group != "" {
		svc := strings.ToLower(strings.Split(group, ".")[0])
		if svc != "" {
			out = append(out, "${"+svc+short+"ID}")
		}
	}
	return out
}
