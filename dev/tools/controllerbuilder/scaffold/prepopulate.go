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
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/util/sets"
)

// identityFields are proto fields that the KRM object expresses through its own
// identity rather than as spec fields. "name" is the resource's own resource
// name, which KCC models as metadata.name / spec.resourceID plus the parent refs.
var identityFields = map[string]bool{
	"name": true,
}

// JudgementItem represents a finding or field needing human review, recorded in
// the service's needs_judgement_call.txt queue file.
type JudgementItem struct {
	// FieldPath is the KRM JSON path, e.g. ".spec.forwardingRules".
	FieldPath string
	// Reason is an identifier categorized for review (e.g., "possible-reference").
	Reason string
	// Detail provides additional context, such as the target proto type or error explanation.
	Detail string
}

// PrepopulateResult contains rendered struct bodies and review items produced during prepopulation.
type PrepopulateResult struct {
	// SpecFields is the rendered Go source for the Spec struct fields.
	SpecFields string
	// ObservedStateFields is the rendered Go source for the ObservedState struct fields,
	// empty when no fields are marked OUTPUT_ONLY in the proto.
	ObservedStateFields string
	// ExtraImports contains necessary import paths beyond the default template imports.
	ExtraImports []string
	// Judgement lists review items requiring human inspection.
	Judgement []JudgementItem
}

// PrepopulateSpec renders the top-level Spec fields for a resource proto message.
//
// Fields requiring decisions (such as potential references) are emitted using
// their proto-derived types and queued in Judgement for review, ensuring no
// fields are silently dropped during scaffolding.
//
// ObservedState is generated separately by PrepopulateObservedState using output
// fields discovered during proto traversal.
func PrepopulateSpec(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) (*PrepopulateResult, error) {
	if msg == nil {
		return nil, fmt.Errorf("no message descriptor")
	}

	out := &PrepopulateResult{}
	var buf bytes.Buffer

	emitted := 0
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)

		// Output-only fields belong in ObservedState, which the type generator
		// writes separately.
		if codegen.IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			continue
		}
		if identityFields[string(field.Name())] {
			// Identity fields (e.g. "name") are managed via status.externalRef rather
			// than directly in spec fields.
			continue
		}

		// We render each field on its own so we can inspect its output before
		// appending it. When the generator cannot type a field it writes a
		// "// TODO:" comment and moves on, and the field never reaches the CRD.
		var field_ bytes.Buffer
		codegen.WriteField(&field_, field, msg, emitted, false, opts)
		buf.Write(field_.Bytes())
		emitted++

		if _, reason, ok := codegen.UnsupportedFieldMarker(field_.String()); ok {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
				Reason:    "unsupported-field-type",
				Detail:    reason,
			})
		}
		if item, ok := judgementFor(field, opts); ok {
			out.Judgement = append(out.Judgement, item)
		}
	}

	out.SpecFields = buf.String()

	// Add root untriaged marker so the resource is reviewed before graduating.
	out.Judgement = append([]JudgementItem{{
		Reason: "untriaged-bulk-generation",
		Detail: "spec was generated from proto definition; verify refs, omissions, and KRM conventions",
	}}, out.Judgement...)

	return out, nil
}

// PrepopulateObservedState generates the struct body for the resource-level
// <Kind>ObservedState from output message details, along with queue entries
// for any output-only fields that require review or could not be mapped.
func PrepopulateObservedState(details *codegen.OutputMessageDetails, observedStateMessages sets.String, opts codegen.WriteOptions) (fields string, judgement []JudgementItem) {
	if details == nil {
		return "", nil
	}

	var buf bytes.Buffer
	// identityFields leaves out "name", which KCC carries in status.externalRef,
	// even where the proto marks it OUTPUT_ONLY.
	notes := codegen.WriteObservedStateFields(&buf, details, observedStateMessages, identityFields, opts)
	fields = buf.String()

	// A field missing from the struct gets a queue entry, whether the skip map
	// left it out or WriteField could not type it.
	for _, n := range notes {
		switch {
		case n.Skipped:
			judgement = append(judgement, JudgementItem{
				FieldPath: ".status.observedState." + n.JSONName,
				Reason:    "observedstate-identity-field-omitted",
				Detail:    "proto marks field OUTPUT_ONLY; resource name is represented by status.externalRef. Verify applicability for this resource",
			})
		default:
			if _, reason, ok := codegen.UnsupportedFieldMarker(n.Rendered); ok {
				judgement = append(judgement, JudgementItem{
					FieldPath: ".status.observedState." + n.JSONName,
					Reason:    "unsupported-field-type",
					Detail:    reason,
				})
			}
		}
	}

	return fields, judgement
}

// judgementFor checks whether a proto field carries a google.api.resource_reference
// annotation and returns a JudgementItem proposing it as a reference candidate.
func judgementFor(field protoreflect.FieldDescriptor, opts codegen.WriteOptions) (JudgementItem, bool) {
	if field.Options() == nil {
		return JudgementItem{}, false
	}
	v := proto.GetExtension(field.Options(), annotations.E_ResourceReference)
	rr, _ := v.(*annotations.ResourceReference)
	if rr == nil {
		return JudgementItem{}, false
	}
	target := rr.GetType()
	if target == "" {
		target = rr.GetChildType()
	}
	if target == "" {
		return JudgementItem{}, false
	}

	return JudgementItem{
		FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
		Reason:    "possible-reference",
		Detail:    "target=" + target,
	}, true
}

// FormatJudgementEntries renders queue lines for one resource, in the format
// apis/<service>/needs_judgement_call.txt expects.
func FormatJudgementEntries(kind, group string, items []JudgementItem) string {
	var sb strings.Builder
	for _, it := range items {
		// A resource-level item has no field path.
		subject := "resource"
		if it.FieldPath != "" {
			subject = fmt.Sprintf("field %q", it.FieldPath)
		}
		sb.WriteString(fmt.Sprintf("kind=%s group=%s: %s reason=%s",
			kind, group, subject, it.Reason))
		if it.Detail != "" {
			sb.WriteString(" (" + it.Detail + ")")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// ExtraImportsFor scans rendered field bodies and returns required package imports
// for mapped external types (e.g. apiextensionsv1.JSON, common.Status).
func ExtraImportsFor(bodies ...string) []string {
	var out []string
	for qualifier, importPath := range codegen.QualifierImports {
		for _, body := range bodies {
			if strings.Contains(body, qualifier+".") {
				// Always emit with explicit qualifier alias, avoiding mismatches between
				// package import path segments and type qualifiers (e.g. apiextensionsv1).
				out = append(out, fmt.Sprintf("%s %q", qualifier, importPath))
				break
			}
		}
	}
	sort.Strings(out)
	return out
}
