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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
)

const (
	RolePrimary    = "PRIMARY"
	RoleDependency = "DEPENDENCY"
	RoleOperation  = "OPERATION"
	RoleAuxiliary  = "AUXILIARY"

	PhaseSourceExplicit = "EXPLICIT"
	PhaseSourceInferred = "INFERRED"
)

// LROSummary summarizes Long-Running Operation (LRO) polling associated with a mutation RPC.
type LROSummary struct {
	Verb          string
	OperationName string
	PollIndices   []int // 1-based stanza indices of GET /operations/... calls
	Done          bool
	HasError      bool
	ErrorMsg      string
	FinalResponse map[string]any
}

// ClassifiedEntry wraps a parsed logparser.Entry with lifecycle phase, resource role, and LRO metadata.
type ClassifiedEntry struct {
	Index         int // 1-based stanza index in _http.log
	Entry         *logparser.Entry
	Phase         string
	Role          string
	ResourceKind  string
	IsLROPoll     bool
	ParentIndex   int // 1-based stanza index of the initiating mutation if IsLROPoll
	LRO           *LROSummary
	ShortEndpoint string
	UpdateMask    []string
	Details       string
}

// Classifier attributes resource roles, correlates LRO polling stanzas, and resolves lifecycle phases.
type Classifier struct {
	fix *fixture.Data
}

// New constructs a Classifier for the given fixture Data.
func New(fix *fixture.Data) *Classifier {
	return &Classifier{fix: fix}
}

// Classify classifies each parsed log entry by resource role, LRO correlation, and lifecycle phase.
func (c *Classifier) Classify(entries []*logparser.Entry) ([]*ClassifiedEntry, string) {
	classified := make([]*ClassifiedEntry, len(entries))
	for i, e := range entries {
		classified[i] = &ClassifiedEntry{
			Index:         i + 1,
			Entry:         e,
			Phase:         e.Phase,
			ShortEndpoint: e.ShortEndpoint(),
			UpdateMask:    e.ExtractUpdateMask(),
		}
	}

	c.attributeRolesAndLROs(classified)

	var phaseSource string
	if hasExplicitPhaseComments(entries) {
		phaseSource = PhaseSourceExplicit
		c.classifyWithExplicitPhases(classified)
	} else {
		phaseSource = PhaseSourceInferred
		for _, ce := range classified {
			if ce.Phase == "" {
				ce.Phase = logparser.PhaseUnknown
				ce.Entry.Phase = logparser.PhaseUnknown
			}
		}
	}

	populatePrimaryEntryDetails(classified)
	return classified, phaseSource
}

func (c *Classifier) attributeRolesAndLROs(classified []*ClassifiedEntry) {
	var primary *fixture.KRMResource
	if c.fix != nil {
		primary = c.fix.PrimaryCreate
	}
	var lastMutationWithOp *ClassifiedEntry

	for i, ce := range classified {
		if ce.Entry.IsOperationsPoll() {
			ce.IsLROPoll = true
			if parent, ok := findInitiatingLROMutation(classified, i, lastMutationWithOp); ok {
				ce.ParentIndex = parent.Index
				ce.Role = parent.Role
				ce.ResourceKind = parent.ResourceKind
				if parent.LRO == nil {
					if initLRO, found := extractInitialLROSummary(parent.Entry); found {
						parent.LRO = initLRO
					} else {
						parent.LRO = &LROSummary{Verb: inferLROVerb(parent.Entry, nil, nil)}
					}
				}
				parent.LRO.PollIndices = append(parent.LRO.PollIndices, ce.Index)
				updateLROSummaryFromPoll(parent.LRO, ce.Entry)
			} else if primary != nil {
				ce.Role = RolePrimary
				ce.ResourceKind = primary.Kind
			}
			continue
		}

		if lro, found := extractInitialLROSummary(ce.Entry); found {
			ce.LRO = lro
			lastMutationWithOp = ce
		} else if ce.Entry.IsMutation() {
			lastMutationWithOp = ce
		}

		if depKind := c.matchDependencyEntry(ce.Entry); depKind != "" {
			ce.Role = RoleDependency
			ce.ResourceKind = depKind
		} else {
			ce.Role = RolePrimary
			if primary != nil {
				ce.ResourceKind = primary.Kind
			}
		}
	}
}

func findInitiatingLROMutation(classified []*ClassifiedEntry, currentIdx int, lastMutation *ClassifiedEntry) (*ClassifiedEntry, bool) {
	poll := classified[currentIdx]
	pollBody, _ := poll.Entry.Response.ParseJSONBody()
	pollOpName := ""
	pollTarget := ""
	if pollBody != nil {
		pollOpName, _ = pollBody["name"].(string)
		if meta, ok := pollBody["metadata"].(map[string]any); ok {
			pollTarget, _ = meta["target"].(string)
		}
		if pollTarget == "" {
			pollTarget, _ = pollBody["targetLink"].(string)
		}
	}

	var fallback *ClassifiedEntry
	for j := currentIdx - 1; j >= 0; j-- {
		cand := classified[j]
		if cand.IsLROPoll || cand.LRO == nil {
			continue
		}
		if pollOpName != "" && cand.LRO.OperationName != "" {
			if pollOpName == cand.LRO.OperationName {
				return cand, true
			}
			continue
		}
		if pollTarget != "" && strings.Contains(cand.Entry.Request.URL, pollTarget) {
			return cand, true
		}
		if fallback == nil {
			fallback = cand
		}
	}
	if fallback != nil {
		return fallback, true
	}
	if lastMutation != nil {
		return lastMutation, true
	}
	return nil, false
}

func extractInitialLROSummary(e *logparser.Entry) (*LROSummary, bool) {
	if e == nil || e.IsReadOnly() {
		return nil, false
	}
	body, ok := e.Response.ParseJSONBody()
	if !ok {
		return nil, false
	}

	_, hasDone := body["done"]
	metaMap, hasMeta := body["metadata"].(map[string]any)
	kindStr, _ := body["kind"].(string)
	nameStr, _ := body["name"].(string)

	isLRO := hasDone || strings.HasSuffix(kindStr, "#operation")
	if !isLRO && hasMeta {
		if typeStr, isStr := metaMap["@type"].(string); isStr && strings.Contains(typeStr, "Operation") {
			isLRO = true
		}
		if _, hasVerb := metaMap["verb"]; hasVerb {
			isLRO = true
		}
	}
	if !isLRO && (strings.Contains(nameStr, "/operations/") || strings.HasPrefix(nameStr, "operations/") || strings.HasPrefix(nameStr, "operation-")) {
		isLRO = true
	}
	if !isLRO {
		return nil, false
	}

	lro := &LROSummary{
		OperationName: nameStr,
		Verb:          inferLROVerb(e, body, metaMap),
	}
	updateLROSummaryFromPoll(lro, e)
	return lro, true
}

func inferLROVerb(e *logparser.Entry, body map[string]any, metaMap map[string]any) string {
	if metaMap != nil {
		if v, ok := metaMap["verb"].(string); ok && v != "" {
			return v
		}
	}
	if body != nil {
		if v, ok := body["operationType"].(string); ok && v != "" {
			return v
		}
	}
	switch e.Request.Method {
	case "POST":
		return "create"
	case "PATCH", "PUT":
		return "update"
	case "DELETE":
		return "delete"
	case "GRPC":
		parts := strings.Split(e.Request.URL, "/")
		if len(parts) > 0 {
			m := parts[len(parts)-1]
			switch {
			case strings.HasPrefix(m, "Create") || strings.HasPrefix(m, "Upload") || strings.HasPrefix(m, "Insert"):
				return "create"
			case strings.HasPrefix(m, "Update") || strings.HasPrefix(m, "Patch"):
				return "update"
			case strings.HasPrefix(m, "Delete"):
				return "delete"
			}
			return m
		}
	}
	return strings.ToLower(e.Request.Method)
}

func updateLROSummaryFromPoll(lro *LROSummary, e *logparser.Entry) {
	body, ok := e.Response.ParseJSONBody()
	if !ok {
		return
	}
	if done, isBool := body["done"].(bool); isBool && done {
		lro.Done = true
	}
	if status, isStr := body["status"].(string); isStr && (status == "DONE" || status == "SUCCEEDED") {
		lro.Done = true
	}
	if metaMap, isMap := body["metadata"].(map[string]any); isMap && lro.Verb == "" {
		if v, isStr := metaMap["verb"].(string); isStr && v != "" {
			lro.Verb = v
		}
	}
	if errObj, isMap := body["error"].(map[string]any); isMap && len(errObj) > 0 {
		lro.HasError = true
		if msg, isStr := errObj["message"].(string); isStr {
			lro.ErrorMsg = msg
		}
	}
	if respObj, isMap := body["response"].(map[string]any); isMap && len(respObj) > 0 {
		typeStr, _ := respObj["@type"].(string)
		if !strings.HasSuffix(typeStr, "google.protobuf.Empty") {
			lro.FinalResponse = respObj
		}
	}
}

func (c *Classifier) matchDependencyEntry(e *logparser.Entry) string {
	if e.Phase == logparser.PhaseDependency {
		if c.fix != nil {
			uPath := e.CleanURLPath()
			for _, dep := range c.fix.Dependencies {
				for _, tok := range []string{dep.ResourceID, dep.Name} {
					if tok != "" && (strings.Contains(uPath, tok) || strings.Contains(e.Request.URL, tok) || strings.Contains(e.Request.Body, tok)) {
						return dep.Kind
					}
				}
			}
			if len(c.fix.Dependencies) == 1 {
				return c.fix.Dependencies[0].Kind
			}
		}
		return "Dependency"
	}

	if c.fix == nil || len(c.fix.Dependencies) == 0 {
		return ""
	}

	primary := c.fix.PrimaryCreate
	primaryName := ""
	primaryResID := ""
	primaryKind := ""
	if primary != nil {
		primaryName = primary.Name
		primaryResID = primary.ResourceID
		primaryKind = primary.Kind
	}

	uPath := e.CleanURLPath()
	for _, dep := range c.fix.Dependencies {
		if dep.Kind == primaryKind {
			continue
		}
		for _, depName := range []string{dep.ResourceID, dep.Name} {
			if len(depName) < 3 || depName == primaryName || depName == primaryResID {
				continue
			}
			if strings.Contains(uPath, depName) || strings.Contains(e.Request.URL, depName) {
				return dep.Kind
			}
			if strings.Contains(e.Request.Body, depName) &&
				(primaryName == "" || (!strings.Contains(e.Request.URL, primaryName) && !strings.Contains(e.Request.Body, primaryName))) &&
				(primaryResID == "" || (!strings.Contains(e.Request.URL, primaryResID) && !strings.Contains(e.Request.Body, primaryResID))) {
				return dep.Kind
			}
		}
	}
	return ""
}

func hasExplicitPhaseComments(entries []*logparser.Entry) bool {
	for _, e := range entries {
		if e.Phase != "" {
			return true
		}
	}
	return false
}

func (c *Classifier) classifyWithExplicitPhases(classified []*ClassifiedEntry) {
	curPhase := logparser.PhaseCreate
	for _, ce := range classified {
		if ce.Entry.Phase == logparser.PhaseDependency {
			ce.Role = RoleDependency
			if ce.Entry.IsDelete() {
				ce.Phase = logparser.PhaseDelete
			} else {
				ce.Phase = logparser.PhaseCreate
			}
			continue
		}
		if ce.Phase != "" {
			curPhase = ce.Phase
		} else {
			ce.Phase = curPhase
			ce.Entry.Phase = curPhase
		}
	}
}

func populatePrimaryEntryDetails(classified []*ClassifiedEntry) {
	seenCreateMut := false
	seenUpdateMut := false
	for _, ce := range classified {
		if ce.Role != RolePrimary || ce.IsLROPoll {
			continue
		}
		method := ce.Entry.Request.Method
		code := ce.Entry.Response.ExtractStatusNumber()
		switch ce.Phase {
		case logparser.PhaseCreate:
			if ce.Entry.IsReadOnly() {
				if code == 404 {
					ce.Details = "Pre-create existence check"
				} else if !seenCreateMut {
					ce.Details = "Pre-create read"
				} else {
					ce.Details = "Post-create read (matches `create.yaml`)"
				}
			} else if ce.Entry.IsMutation() {
				if !seenCreateMut {
					seenCreateMut = true
					ce.Details = "Created resource"
				} else if len(ce.UpdateMask) > 0 {
					ce.Details = fmt.Sprintf("Composite create follow-up (updateMask: %s)", strings.Join(ce.UpdateMask, ", "))
				} else {
					ce.Details = fmt.Sprintf("Composite create follow-up (`%s`)", method)
				}
			}
		case logparser.PhaseUpdate:
			if ce.Entry.IsMutation() {
				seenUpdateMut = true
				if len(ce.UpdateMask) > 0 {
					ce.Details = fmt.Sprintf("updateMask: %s", strings.Join(ce.UpdateMask, ", "))
				} else {
					ce.Details = fmt.Sprintf("Updated resource (`%s`)", method)
				}
			} else if ce.Entry.IsReadOnly() {
				if seenUpdateMut {
					ce.Details = "Post-update read (matches `update.yaml`)"
				} else {
					ce.Details = "Pre-update read"
				}
			}
		case logparser.PhaseRereconcile:
			ce.Details = "Re-reconcile read-only check"
		case logparser.PhaseDelete:
			if ce.Entry.IsDelete() {
				ce.Details = "Deleted resource"
			} else if ce.Entry.IsReadOnly() {
				if code == 404 {
					ce.Details = "Post-delete 404 confirmation"
				} else {
					ce.Details = "Pre-delete read"
				}
			} else {
				ce.Details = fmt.Sprintf("Teardown (`%s`)", method)
			}
		}
	}
}
