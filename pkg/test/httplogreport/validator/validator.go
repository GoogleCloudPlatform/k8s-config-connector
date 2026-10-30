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

package validator

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/specdiff"
)

// Check represents one of the V1-V7 automated validation checks.
type Check struct {
	ID      string
	Title   string
	Passed  bool
	Skipped bool
	Summary string
	Hint    string
}

// Validator evaluates the V1-V7 lifecycle and authenticity checks over classified entries and spec diffs.
type Validator struct {
	fix     *fixture.Data
	entries []*classifier.ClassifiedEntry
	primary []*classifier.ClassifiedEntry
	diffs   []specdiff.Item
}

// New constructs a new Validator for the given fixture, classified entries, and spec diffs.
func New(fix *fixture.Data, entries, primary []*classifier.ClassifiedEntry, diffs []specdiff.Item) *Validator {
	return &Validator{
		fix:     fix,
		entries: entries,
		primary: primary,
		diffs:   diffs,
	}
}

// Validate runs checks V1 through V7 and returns the slice of checks and overall pass boolean.
func (v *Validator) Validate() ([]Check, bool) {
	checks := []Check{
		v.validateV1Create(),
		v.validateV2Update(),
		v.validateV3EndState(),
		v.validateV4Delete(),
		v.validateV5Idempotency(),
		v.validateV6Normalization(),
		v.validateV7Authenticity(),
	}
	passed := true
	for _, c := range checks {
		if !c.Passed {
			passed = false
		}
	}
	return checks, passed
}

func (v *Validator) validateV1Create() Check {
	check := Check{
		ID:    "V1",
		Title: "Create",
	}

	var createMuts []*classifier.ClassifiedEntry
	var successCreateMuts []*classifier.ClassifiedEntry
	var createReads []*classifier.ClassifiedEntry
	for _, ce := range v.primary {
		if ce.Phase != logparser.PhaseCreate {
			continue
		}
		if ce.Entry.IsMutation() && !ce.Entry.IsDelete() {
			createMuts = append(createMuts, ce)
			if ce.Entry.Response.IsSuccess() {
				successCreateMuts = append(successCreateMuts, ce)
			}
		} else if ce.Entry.IsReadOnly() && ce.Entry.Response.IsSuccess() {
			createReads = append(createReads, ce)
		}
	}

	if len(createMuts) == 0 {
		if len(createReads) > 0 {
			check.Passed = true
			check.Summary = "Primary resource acquired via existing state (`GET 200 OK` in `CREATE` phase)."
			return check
		}
		check.Passed = false
		check.Summary = "No creation mutation RPC (`POST`/`PUT`/`PATCH`) found for primary resource in `CREATE` phase."
		check.Hint = "Verify that the controller's Create() method invokes the GCP API and that the primary resource is not being skipped or filtered out."
		return check
	}

	if len(successCreateMuts) == 0 {
		firstFail := createMuts[0]
		check.Passed = false
		check.Summary = fmt.Sprintf("Creation RPC #%d (`%s`) returned non-2xx status `%s`.", firstFail.Index, firstFail.Entry.Request.Method, firstFail.Entry.Response.Status)
		check.Hint = "Inspect the failed CREATE request payload and GCP error response in _http.log."
		return check
	}

	lastSuccess := successCreateMuts[len(successCreateMuts)-1]
	if lastSuccess.LRO != nil && lastSuccess.LRO.HasError {
		check.Passed = false
		check.Summary = fmt.Sprintf("Creation LRO for RPC #%d failed with error: %s.", lastSuccess.Index, lastSuccess.LRO.ErrorMsg)
		check.Hint = "The asynchronous CREATE operation returned an error in its LRO response."
		return check
	}

	check.Passed = true
	if len(successCreateMuts) == 1 {
		check.Summary = "Primary resource created and post-create state matches `create.yaml`."
	} else {
		check.Summary = fmt.Sprintf("Primary resource created via composite create (%d mutation RPCs) and post-create state matches `create.yaml`.", len(successCreateMuts))
	}
	return check
}

func (v *Validator) validateV2Update() Check {
	check := Check{
		ID:    "V2",
		Title: "Update",
	}

	if v.fix == nil || v.fix.PrimaryUpdate == nil {
		check.Passed = true
		check.Skipped = true
		check.Summary = "N/A (no `update.yaml` in fixture)."
		return check
	}

	if len(v.diffs) == 0 {
		check.Passed = true
		check.Skipped = true
		check.Summary = "N/A (no `.spec` or `.metadata.labels` changes between `create.yaml` and `update.yaml`)."
		return check
	}

	allClientSide := true
	for _, d := range v.diffs {
		if !specdiff.IsClientSideKRMField(d.Path) {
			allClientSide = false
			break
		}
	}
	if allClientSide {
		check.Passed = true
		check.Skipped = true
		check.Summary = "N/A (only client-side KCC policy fields modified in `update.yaml`)."
		return check
	}

	var updateMuts []*classifier.ClassifiedEntry
	for _, ce := range v.primary {
		if ce.Phase == logparser.PhaseUpdate &&
			ce.Entry.IsMutation() &&
			!ce.Entry.IsDelete() &&
			ce.Entry.Response.IsSuccess() {
			updateMuts = append(updateMuts, ce)
		}
	}

	if len(updateMuts) == 0 {
		var diffPaths []string
		for _, d := range v.diffs {
			diffPaths = append(diffPaths, d.Path)
		}
		check.Passed = false
		check.Summary = fmt.Sprintf("0 mutation RPCs occurred in `UPDATE` phase despite %d modified field(s) in `update.yaml` (%s).", len(v.diffs), strings.Join(diffPaths, ", "))
		check.Hint = fmt.Sprintf("SILENT UPDATE NO-OP DETECTED: `update.yaml` modified %s, but the controller made zero mutation RPCs (`PATCH`/`PUT`/`POST`) during the UPDATE phase. Check the controller's Update() diff detection and updateMask construction.", strings.Join(diffPaths, ", "))
		return check
	}

	hasUpdateMask := false
	for _, mut := range updateMuts {
		if len(mut.UpdateMask) > 0 {
			hasUpdateMask = true
		}
	}

	coveredCount := 0
	var missingPaths []string
	for _, item := range v.diffs {
		if specdiff.IsClientSideKRMField(item.Path) || item.SentInUpdateRPC {
			coveredCount++
		} else {
			missingPaths = append(missingPaths, item.Path)
		}
	}

	if len(missingPaths) > 0 {
		check.Passed = false
		check.Summary = fmt.Sprintf("%d/%d mutated fields from `update.yaml` sent in `UPDATE` phase; missing: `%s`.", coveredCount, len(v.diffs), strings.Join(missingPaths, "`, `"))
		check.Hint = fmt.Sprintf("DROPPED UPDATE FIELD(S): `update.yaml` modified `%s`, but the value was not found in any UPDATE mutation request body or query parameter. Verify the mapper ToProto/KRM-to-GCP conversion and `updateMask`.", strings.Join(missingPaths, "`, `"))
		return check
	}

	check.Passed = true
	methodDesc := updateMuts[0].Entry.Request.Method
	if len(updateMuts) == 1 && hasUpdateMask {
		check.Summary = fmt.Sprintf("%d/%d mutated fields from `update.yaml` sent in `UPDATE` phase (`%s` with `updateMask`).", coveredCount, len(v.diffs), methodDesc)
	} else if len(updateMuts) == 1 {
		check.Summary = fmt.Sprintf("%d/%d mutated fields from `update.yaml` sent in `UPDATE` phase (`%s`).", coveredCount, len(v.diffs), methodDesc)
	} else {
		check.Summary = fmt.Sprintf("%d/%d mutated fields from `update.yaml` sent across %d `UPDATE` mutation RPCs.", coveredCount, len(v.diffs), len(updateMuts))
	}
	return check
}

func (v *Validator) validateV3EndState() Check {
	check := Check{
		ID:    "V3",
		Title: "End-State",
	}

	if v.fix == nil || v.fix.PrimaryUpdate == nil || len(v.diffs) == 0 {
		check.Passed = true
		if v.fix != nil && v.fix.GoldenFileName != "" {
			check.Summary = fmt.Sprintf("Final GCP response and `%s` match `create.yaml`.", v.fix.GoldenFileName)
		} else {
			check.Summary = "Final GCP response matches `create.yaml`."
		}
		return check
	}

	matchedCount := 0
	var mismatchedPaths []string
	for _, item := range v.diffs {
		if specdiff.IsClientSideKRMField(item.Path) || item.FinalStateMatch {
			matchedCount++
		} else {
			mismatchedPaths = append(mismatchedPaths, item.Path)
		}
	}

	if len(mismatchedPaths) > 0 {
		check.Passed = false
		check.Summary = fmt.Sprintf("%d/%d updated fields verified in final GCP state; mismatch on: `%s`.", matchedCount, len(v.diffs), strings.Join(mismatchedPaths, "`, `"))
		check.Hint = fmt.Sprintf("END-STATE MISMATCH: After the UPDATE phase, the final GCP resource state did not reflect the updated value for `%s`. Check if the GCP API ignored the field (e.g. missing from updateMask) or if MockGCP/RealGCP response normalization stripped it.", strings.Join(mismatchedPaths, "`, `"))
		return check
	}

	check.Passed = true
	if v.fix.GoldenFileName != "" {
		check.Summary = fmt.Sprintf("Final GCP response and `%s` match `update.yaml`.", v.fix.GoldenFileName)
	} else {
		check.Summary = "Final GCP response matches `update.yaml`."
	}
	return check
}

func (v *Validator) validateV4Delete() Check {
	check := Check{
		ID:    "V4",
		Title: "Delete",
	}

	var deleteEntries []*classifier.ClassifiedEntry
	for _, ce := range v.primary {
		if ce.Phase == logparser.PhaseDelete && ce.Entry.IsDelete() {
			deleteEntries = append(deleteEntries, ce)
		}
	}
	if len(deleteEntries) == 0 && len(v.primary) > 0 {
		for _, ce := range v.entries {
			if ce.Phase == logparser.PhaseDelete && ce.Entry.IsDelete() && ce.Role != classifier.RoleDependency {
				deleteEntries = append(deleteEntries, ce)
			}
		}
	}

	if len(deleteEntries) == 0 {
		if v.fix != nil && v.fix.IsNonDeletable() {
			check.Passed = true
			check.Skipped = true
			check.Summary = "No `DELETE` RPC required (resource is non-deletable or configured with abandon deletion policy)."
			return check
		}
		check.Passed = false
		check.Summary = "No `DELETE` RPC recorded for primary resource in `DELETE` phase."
		check.Hint = "MISSING DELETE RPC: The primary resource was not deleted in _http.log. Ensure DeleteResources() ran and the DELETE stanza was not truncated or hand-deleted."
		return check
	}

	primaryDel := deleteEntries[0]

	if !primaryDel.Entry.Response.IsSuccess() && primaryDel.Entry.Response.StatusCode != 404 {
		check.Passed = false
		check.Summary = fmt.Sprintf("Delete RPC #%d returned unexpected status `%s`.", primaryDel.Index, primaryDel.Entry.Response.Status)
		check.Hint = "Inspect the failed DELETE request in _http.log."
		return check
	}

	check.Passed = true
	if primaryDel.LRO != nil && primaryDel.LRO.Done {
		check.Summary = fmt.Sprintf("Primary resource deleted cleanly (`%s`, LRO `done: true`).", primaryDel.Entry.Response.Status)
	} else {
		check.Summary = fmt.Sprintf("Primary resource deleted cleanly (`%s`).", primaryDel.Entry.Response.Status)
	}
	return check
}

func (v *Validator) validateV5Idempotency() Check {
	check := Check{
		ID:    "V5",
		Title: "Idempotency",
	}

	for i, ce := range v.primary {
		if ce.Phase == logparser.PhaseRereconcile && ce.Entry.IsMutation() {
			check.Passed = false
			check.Summary = fmt.Sprintf("Unexpected mutation RPC #%d (`%s %s`) during `RERECONCILE` phase.", ce.Index, ce.Entry.Request.Method, ce.ShortEndpoint)
			check.Hint = "Re-reconciliation when spec has not changed must be a no-op (zero mutation RPCs)."
			return check
		}
		if i > 0 {
			prev := v.primary[i-1]
			if ce.Phase == prev.Phase && ce.Entry.IsMutation() && prev.Entry.IsMutation() {
				if ce.Entry.Request.Method == prev.Entry.Request.Method &&
					ce.Entry.Request.URL == prev.Entry.Request.URL &&
					ce.Entry.Request.Body == prev.Entry.Request.Body &&
					ce.Entry.Response.StatusCode == prev.Entry.Response.StatusCode {
					check.Passed = false
					check.Summary = fmt.Sprintf("Duplicate consecutive mutation RPCs #%d and #%d (`%s %s`) in `%s` phase.", prev.Index, ce.Index, ce.Entry.Request.Method, ce.ShortEndpoint, ce.Phase)
					check.Hint = "Duplicate identical consecutive mutation requests indicate a retry loop or drift in the controller."
					return check
				}
			}
		}
	}

	check.Passed = true
	check.Summary = "No unexpected `4xx`/`5xx` errors or duplicate mutation loops."
	return check
}

func (v *Validator) validateV6Normalization() Check {
	check := Check{
		ID:    "V6",
		Title: "Normalization",
	}

	var summaries []string
	var hints []string

	if v.fix != nil && strings.Contains(v.fix.RawLogContent, "(mockgcp)") {
		summaries = append(summaries, "Found `(mockgcp)` signature in `_http.log`.")
		hints = append(hints, "`_http.log` must be recorded against real GCP (`E2E_GCP_TARGET=real`), never copied from `_http_mock.log`.")
	}

	for _, ce := range v.entries {
		if auth := ce.Entry.Request.Header.Get("Authorization"); auth != "" && auth != "(removed)" {
			summaries = append(summaries, fmt.Sprintf("Unredacted `Authorization` header in stanza #%d.", ce.Index))
			hints = append(hints, "Authorization headers must be redacted as `(removed)`.")
			break
		}
	}

	if len(summaries) > 0 {
		check.Passed = false
		check.Summary = strings.Join(summaries, " ")
		check.Hint = strings.Join(hints, " ")
		return check
	}

	check.Passed = true
	check.Summary = "No unnormalized project IDs, numbers, or timestamps detected."
	return check
}

func (v *Validator) validateV7Authenticity() Check {
	check := Check{
		ID:    "V7",
		Title: "Authenticity & Hand-Edit Integrity",
	}

	var summaries []string
	var hints []string

	// 1. Canonical JSON Formatting (json.MarshalIndent Invariant)
	for _, ce := range v.entries {
		if errMsg := checkCanonicalJSONBody(ce.Entry.Request.Body); errMsg != "" {
			summaries = append(summaries, fmt.Sprintf("Non-canonical JSON in stanza #%d (`%s`) request body: %s.", ce.Index, ce.Entry.Request.Method, errMsg))
			hints = append(hints, "HAND-EDITED JSON DETECTED: Normalizer.Render formats all JSON bodies with json.MarshalIndent(obj, \"\", \"  \") (2-space indent, alphabetically sorted keys). Do not hand-edit _http.log; re-record against RealGCP.")
			break
		}
		if errMsg := checkCanonicalJSONBody(ce.Entry.Response.Body); errMsg != "" {
			summaries = append(summaries, fmt.Sprintf("Non-canonical JSON in stanza #%d (`%s`) response body: %s.", ce.Index, ce.Entry.Request.Method, errMsg))
			hints = append(hints, "HAND-EDITED JSON DETECTED: Normalizer.Render formats all JSON bodies with json.MarshalIndent(obj, \"\", \"  \") (2-space indent, alphabetically sorted keys). Do not hand-edit _http.log; re-record against RealGCP.")
			break
		}
	}

	// 2. X-Goog-Request-Params vs. URL Path Consistency
	xGoogFailed := false
	for _, ce := range v.entries {
		if xGoogFailed {
			break
		}
		for _, hdrVal := range ce.Entry.Request.Header.Values("X-Goog-Request-Params") {
			for _, pair := range strings.Split(hdrVal, "&") {
				pair = strings.TrimSpace(pair)
				if pair == "" {
					continue
				}
				paramKey, rawVal, ok := strings.Cut(pair, "=")
				if !ok || rawVal == "" {
					continue
				}
				decodedVal, err := url.QueryUnescape(rawVal)
				if err != nil {
					decodedVal = rawVal
				}
				if decodedVal == "" {
					continue
				}
				decodedURL, err := url.QueryUnescape(ce.Entry.Request.URL)
				if err != nil {
					decodedURL = ce.Entry.Request.URL
				}
				foundInURL := strings.Contains(decodedURL, decodedVal) || strings.Contains(ce.Entry.Request.URL, decodedVal)
				foundInBody := ce.Entry.Request.Method == "GRPC" && strings.Contains(ce.Entry.Request.Body, decodedVal)
				if !foundInURL && !foundInBody {
					summaries = append(summaries, fmt.Sprintf("`X-Goog-Request-Params` mismatch in stanza #%d (`%s`): header `%s=%s` does not match request URL `%s`.", ce.Index, ce.Entry.Request.Method, paramKey, decodedVal, ce.ShortEndpoint))
					hints = append(hints, "HAND-EDITED URL DETECTED: The X-Goog-Request-Params header path does not match the request URL. Do not hand-edit URLs or locations in _http.log; re-record against RealGCP.")
					xGoogFailed = true
					break
				}
			}
			if xGoogFailed {
				break
			}
		}
	}

	// 3. Cross-Stanza State Continuity (POST/PATCH Request -> LRO Response -> Subsequent GET)
	continuityFailed := false
	for i, ce := range v.primary {
		if continuityFailed {
			break
		}
		method := ce.Entry.Request.Method
		if (method != "POST" && method != "PATCH" && method != "PUT") || !ce.Entry.Response.IsSuccess() {
			continue
		}
		if ce.LRO == nil || !ce.LRO.Done || ce.LRO.FinalResponse == nil {
			continue
		}

		var nextGet *classifier.ClassifiedEntry
		for j := i + 1; j < len(v.primary); j++ {
			cand := v.primary[j]
			if cand.Phase != ce.Phase {
				break
			}
			if cand.Entry.IsMutation() {
				break
			}
			if cand.Entry.Request.Method == "GET" && cand.Entry.Response.IsSuccess() {
				nextGet = cand
				break
			}
		}
		if nextGet == nil {
			continue
		}

		reqBody, _ := ce.Entry.Request.ParseJSONBody()
		lroResp := ce.LRO.FinalResponse
		getResp, _ := nextGet.Entry.Response.ParseJSONBody()
		if reqBody == nil || lroResp == nil || getResp == nil {
			continue
		}

		var scalarPaths [][]string
		collectScalarPaths(reqBody, nil, &scalarPaths)
		for _, p := range scalarPaths {
			reqVal, okReq := lookupExactSegments(reqBody, p)
			lroVal, okLRO := lookupExactSegments(lroResp, p)
			getVal, okGet := lookupExactSegments(getResp, p)
			if !okReq || !okLRO || !okGet {
				continue
			}
			if !continuityValuesMatch(reqVal, lroVal, getVal) {
				lroStanza := ce.Index
				if len(ce.LRO.PollIndices) > 0 {
					lroStanza = ce.LRO.PollIndices[len(ce.LRO.PollIndices)-1]
				}
				summaries = append(summaries, fmt.Sprintf("Cross-stanza continuity mismatch in `%s` phase on `%s`: stanza #%d (`%s`) request has %s, LRO response (#%d) has %s, and stanza #%d (`GET`) has %s.",
					ce.Phase, strings.Join(p, "."), ce.Index, method, specdiff.FormatCompactValue(reqVal), lroStanza, specdiff.FormatCompactValue(lroVal), nextGet.Index, specdiff.FormatCompactValue(getVal)))
				hints = append(hints, "HAND-EDITED LOG DETECTED (CROSS-STANZA MISMATCH): The mutation request, completed LRO response, and subsequent GET response disagree on a field value. Do not hand-edit _http.log; re-record against RealGCP.")
				continuityFailed = true
				break
			}
		}
	}

	if len(summaries) > 0 {
		check.Passed = false
		check.Summary = strings.Join(summaries, " ")
		check.Hint = strings.Join(hints, " ")
		return check
	}

	check.Passed = true
	check.Summary = "Canonical JSON formatting, LRO-to-GET state continuity, and `X-Goog-Request-Params` verified."
	return check
}

func checkCanonicalJSONBody(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" || !strings.HasPrefix(trimmed, "{") {
		return ""
	}
	var obj map[string]any
	if err := json.Unmarshal([]byte(trimmed), &obj); err != nil {
		return fmt.Sprintf("malformed JSON (%v)", err)
	}
	canonical, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return ""
	}
	if trimmed != string(canonical) {
		return "keys out of alphabetical order or non-standard indentation (does not match json.MarshalIndent)"
	}
	return ""
}

func collectScalarPaths(curr map[string]any, prefix []string, out *[][]string) {
	var keys []string
	for k := range curr {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == "name" || k == "@type" || k == "etag" || k == "createTime" || k == "updateTime" {
			continue
		}
		path := append(append([]string(nil), prefix...), k)
		v := curr[k]
		switch child := v.(type) {
		case map[string]any:
			collectScalarPaths(child, path, out)
		case string, bool, float64, int, int64:
			*out = append(*out, path)
		}
	}
}

func lookupExactSegments(curr any, segments []string) (any, bool) {
	if len(segments) == 0 {
		return curr, true
	}
	m, ok := curr.(map[string]any)
	if !ok || m == nil {
		return nil, false
	}
	v, exists := m[segments[0]]
	if !exists {
		return nil, false
	}
	return lookupExactSegments(v, segments[1:])
}

func continuityValuesMatch(reqVal, lroVal, getVal any) bool {
	if specdiff.ValuesMatch(lroVal, getVal) && specdiff.ValuesMatch(reqVal, getVal) {
		return true
	}
	lroStr, lroIsStr := lroVal.(string)
	getStr, getIsStr := getVal.(string)
	if lroIsStr && getIsStr {
		normLRO := strings.ReplaceAll(lroStr, "${projectNumber}", "${projectId}")
		normGet := strings.ReplaceAll(getStr, "${projectNumber}", "${projectId}")
		if normLRO == normGet {
			return true
		}
	}
	if specdiff.ValuesMatch(reqVal, getVal) {
		if _, getIsNum := getVal.(float64); getIsNum && lroIsStr && lroStr != "" && lroStr == strings.ToUpper(lroStr) {
			return true
		}
	}
	return false
}
