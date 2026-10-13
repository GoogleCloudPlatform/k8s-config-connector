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
	"regexp"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"

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
// the service's judgement_queue.yaml.
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

// SpecFields returns the fields of msg that PrepopulateSpec writes into the
// Spec, in proto order. generate-types uses it to plan the structs before
// the Spec exists. Both use topLevelSkip, and a test checks that they agree.
func SpecFields(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []protoreflect.FieldDescriptor {
	var out []protoreflect.FieldDescriptor
	for i := 0; i < msg.Fields().Len(); i++ {
		if field := msg.Fields().Get(i); topLevelSkip(field, msg, opts) == inSpec {
			out = append(out, field)
		}
	}
	return out
}

// specSkip says why PrepopulateSpec leaves a top-level field out of the Spec.
type specSkip int

const (
	// inSpec means the field goes in the Spec.
	inSpec specSkip = iota
	skipOutputOnly
	skipServerSet
	skipIdentity
	skipDeprecated
)

// topLevelSkip returns why PrepopulateSpec leaves a top-level field of the
// resource message msg out of the Spec, or inSpec if the field stays.
// SpecFields and walkSpecFields use it too, so all three drop the same
// fields.
func topLevelSkip(field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) specSkip {
	switch {
	case codegen.IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY):
		return skipOutputOnly
	case codegen.IsServerSetField(field, msg, opts):
		return skipServerSet
	case identityFields[string(field.Name())]:
		return skipIdentity
	case isDeprecated(field):
		return skipDeprecated
	default:
		return inSpec
	}
}

// PrepopulateSpec renders the top-level Spec fields for a resource proto message.
//
// Fields requiring decisions (such as potential references) are emitted using
// their proto-derived types and queued in Judgement for review, ensuring no
// fields are silently dropped during scaffolding.
//
// A field the proto marks deprecated is left out and queued, so a new Kind does
// not start with it.
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

		switch topLevelSkip(field, msg, opts) {
		case skipOutputOnly:
			// Output-only fields belong in ObservedState, which the type
			// generator writes separately.
			continue
		case skipServerSet:
			// A server-set field goes to ObservedState, like an OUTPUT_ONLY
			// one. This check and the type generator's have to agree, or the
			// field lands in both structs.
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: ".status.observedState." + codegen.GetJSONForKRM(field, opts),
				Reason:    "server-set-field-placed",
				Detail: "moved to ObservedState because its name is on the " +
					"server-set allowlist, not because the proto says so: no " +
					"field on this message carries field_behavior. Confirm GCP " +
					"sets this field instead of a user",
			})
			continue
		case skipIdentity:
			// Identity fields (e.g. "name") are managed via status.externalRef
			// rather than directly in spec fields.
			continue
		case skipDeprecated:
			// Leave out a field the proto marks deprecated, so a new Kind does
			// not start with it. walkSpecFields and DetectOutputOnlyInComments
			// skip it too, or their entries would name a path the CRD does not
			// have.
			//
			// A types file scaffolded before this check can still have the
			// field, so the detail says what to do in both cases.
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
				Reason:    "deprecated-field",
				Detail: "the proto marks this field deprecated, so new Kinds leave it out of the Spec. " +
					"If this Spec still has it, remove it by hand; add it back only if users still need it",
			})
			continue
		}

		// We render each field on its own so we can inspect its output before
		// appending it. When the generator cannot type a field it writes a
		// "// TODO:" comment and moves on, and the field never reaches the CRD.
		var field_ bytes.Buffer
		codegen.WriteField(&field_, field, msg, emitted, false, opts, "")
		buf.Write(field_.Bytes())
		emitted++

		// A string field named after a resource in this service may be a
		// reference to it. WriteField has already written a +kcc:guess marker
		// above the field, and this adds the queue entry. Both use
		// SiblingResource, so every marker gets an entry. The field stays a
		// string. A reviewer decides whether to make it a reference.
		if target, ok := codegen.SiblingResource(field, opts); ok {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
				Reason:    judgement.ReasonPossibleReferenceBySibling,
				Detail: "the name matches " + target + ", a resource this service declares; " +
					"confirm whether it should be a reference",
			})
		}

		if _, reason, ok := codegen.UnsupportedFieldMarker(field_.String()); ok {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
				Reason:    "unsupported-field-type",
				Detail:    reason,
			})
		} else if opts.EmitMessageMaps {
			if protoMsg, goType, ok := codegen.MessageMapValueInfo(field); ok {
				out.Judgement = append(out.Judgement, JudgementItem{
					FieldPath: ".spec." + codegen.GetJSONForKRM(field, opts),
					Reason:    "message-map-derived",
					Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", goType, protoMsg),
				})
			}
		}
		if item, ok := judgementFor(field, ".spec."+codegen.GetJSONForKRM(field, opts)); ok {
			out.Judgement = append(out.Judgement, item)
		}
	}

	out.SpecFields = buf.String()

	// The root marker records that nobody has reviewed this resource yet. It
	// does not suppress any check.
	out.Judgement = append([]JudgementItem{{
		Reason: judgement.ReasonUntriagedBulkGeneration,
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
			} else if opts.EmitMessageMaps && n.Field != nil {
				if protoMsg, goType, ok := codegen.MessageMapValueInfo(n.Field); ok {
					judgement = append(judgement, JudgementItem{
						FieldPath: ".status.observedState." + n.JSONName,
						Reason:    "message-map-derived",
						Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", goType, protoMsg),
					})
				}
			}
		}
	}

	return fields, judgement
}

// judgementFor checks whether a proto field carries a google.api.resource_reference
// annotation and returns a JudgementItem proposing it as a reference candidate.
// path is the field's KRM path. PrepopulateSpec calls this for top-level
// fields and ReferenceHints for nested ones.
func judgementFor(field protoreflect.FieldDescriptor, path string) (JudgementItem, bool) {
	if field.Options() == nil {
		return JudgementItem{}, false
	}
	v := proto.GetExtension(field.Options(), annotations.E_ResourceReference)
	rr, _ := v.(*annotations.ResourceReference)
	if rr == nil {
		return JudgementItem{}, false
	}
	var detail string
	switch {
	case rr.GetType() != "":
		detail = "the proto marks this field as a reference to " + rr.GetType() +
			" (google.api.resource_reference); confirm whether it should be a KCC reference"
	case rr.GetChildType() != "":
		detail = "the proto marks this field as the parent of a " + rr.GetChildType() +
			" (google.api.resource_reference child_type); confirm whether it should be a KCC reference"
	default:
		return JudgementItem{}, false
	}

	return JudgementItem{
		FieldPath: path,
		Reason:    judgement.ReasonPossibleReference,
		Detail:    detail,
	}, true
}

// JudgementEntries turns one resource's items into open queue entries.
func JudgementEntries(kind, group string, items []JudgementItem) []judgement.Entry {
	out := make([]judgement.Entry, 0, len(items))
	for _, it := range items {
		out = append(out, judgement.Entry{
			Kind:   kind,
			Group:  group,
			Field:  it.FieldPath,
			Reason: it.Reason,
			Detail: it.Detail,
			Status: judgement.StatusOpen,
		})
	}
	return out
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

// Reasons for an OutputOnlyCandidate.
const (
	reasonWellKnownOutputOnlyPattern = "well-known-output-only-pattern-in-comment"
	reasonPossibleOutputOnlyPattern  = "possible-output-only-pattern-in-comment"
)

// OutputOnlyCandidate is a field the proto documents as output-only in prose
// while carrying no google.api.field_behavior annotation to say so.
type OutputOnlyCandidate struct {
	// FieldPath is the KRM path the field was emitted at, e.g. ".spec.createTime"
	// or, for a nested field, ".spec.config.state".
	FieldPath string
	// Reason is reasonWellKnownOutputOnlyPattern when the comment starts with
	// one of outputOnlyPrefixes. It is reasonPossibleOutputOnlyPattern, a
	// weaker signal, when the comment says "output only" any other way.
	Reason string
	// Comment is the proto's leading comment, so a reviewer can decide without
	// opening the proto.
	Comment string
}

// Item returns the queue entry for c.
//
// The entry names where the field belongs, not where it sits now. The entry
// is about a field missing from ObservedState, and a path under .spec would
// not point there. A nested field keeps the rest of its path, so
// .spec.config.state becomes .status.observedState.config.state.
func (c OutputOnlyCandidate) Item() JudgementItem {
	detail := "proto comment says output only but there is no field_behavior annotation, so it was generated into the Spec."
	if c.Reason == reasonPossibleOutputOnlyPattern {
		detail = "proto comment mentions output only but doesn't start with \"Output only.\" or \"[Output Only]\", and there is no field_behavior annotation, so it was generated into the Spec. " +
			"The comment may be a typo, apply only some of the time, or mean something else."
	}
	return JudgementItem{
		FieldPath: ".status.observedState." + strings.TrimPrefix(c.FieldPath, ".spec."),
		Reason:    c.Reason,
		Detail:    detail + " Move the field to status.observedState if it is confirmed output only. Proto comment: " + c.Comment,
	}
}

// outputOnlyPrefixes are the markers Google API comments use to say a field is
// output only when the proto has no field_behavior annotation. They match
// regardless of case, because compute writes both "[Output Only]" and
// "[Output only]".
var outputOnlyPrefixes = []string{"Output only.", "[Output Only]"}

// outputOnlyMention matches "output only" in any case, with a space or a
// hyphen between the words, as in dlp's "Output-only field, populated by the
// system".
var outputOnlyMention = regexp.MustCompile(`(?i)\boutput[\s-]+only\b`)

// DetectOutputOnlyInComments finds spec fields whose leading proto comments describe
// them as output-only, but lack explicit google.api.field_behavior annotations.
// It reports candidates for manual review rather than moving them automatically.
//
// It checks nested Spec fields too, with the same rules.
func DetectOutputOnlyInComments(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []OutputOnlyCandidate {
	if msg == nil {
		return nil
	}
	var out []OutputOnlyCandidate
	// walkSpecFields skips the fields PrepopulateSpec leaves out of the Spec:
	// OUTPUT_ONLY fields, identity fields, deprecated top-level fields, and
	// server-set fields when --place-server-set-fields is on. A server-set
	// field is already in ObservedState with its own queue entry. The walk
	// also skips fields the generator cannot type, since they never reach the
	// CRD, and fields the generator writes as references.
	walkSpecFields(msg, ".spec", opts, true, map[protoreflect.FullName]bool{}, func(path, comment string, _ protoreflect.FieldDescriptor) {
		if reason, ok := outputOnlyReason(comment); ok {
			out = append(out, OutputOnlyCandidate{FieldPath: path, Reason: reason, Comment: comment})
		}
	})
	return out
}

// outputOnlyReason reports whether a field's comment, joined into one line,
// says the field is output only, and how strongly.
//
// A comment that starts with one of outputOnlyPrefixes gets
// reasonWellKnownOutputOnlyPattern. Any other mention of "output only" gets
// reasonPossibleOutputOnlyPattern. That covers the words after other text,
// as in websecurityscanner's managed_scan: "Whether the scan config is
// managed by Web Security Scanner, output only." It also covers a comment
// that starts with the words but not with a prefix. That is usually a typo,
// as in automl's "Output only . The", or a condition, as in spanner's
// Backup.name, which is output only for the CreateBackup operation and
// required for UpdateBackup.
func outputOnlyReason(comment string) (string, bool) {
	for _, prefix := range outputOnlyPrefixes {
		if len(comment) >= len(prefix) && strings.EqualFold(comment[:len(prefix)], prefix) {
			return reasonWellKnownOutputOnlyPattern, true
		}
	}
	if outputOnlyMention.MatchString(comment) {
		return reasonPossibleOutputOnlyPattern, true
	}
	return "", false
}
