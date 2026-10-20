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

var possibleStatusFieldNames = map[string]bool{
	"state":             true,
	"status":            true,
	"createTime":        true,
	"updateTime":        true,
	"deleteTime":        true,
	"lastUpdateTime":    true,
	"creationTimestamp": true,
	"startTime":         true,
	"endTime":           true,
	"uid":               true,
	"id":                true,
	"selfLink":          true,
	"selfLinkWithID":    true,
	"reconciling":       true,
}

var topLevelLocationFieldNames = map[string]bool{
	"location": true,
	"region":   true,
	"zone":     true,
}

var nestedLocationFieldNames = map[string]bool{
	"region": true,
	"zone":   true,
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
	case codegen.IsFieldAllOutputOnlyMessage(field):
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
				Reason:    judgement.ReasonServerSetFieldPlaced,
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
				Reason:    judgement.ReasonDeprecatedField,
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

		fieldPath := ".spec." + codegen.GetJSONForKRM(field, opts)
		// A string field named after a resource in this service may be a
		// reference to it. WriteField has already written a +kcc:guess marker
		// above the field, and this adds the queue entry. Both use
		// SiblingResource, so every marker gets an entry. The field stays a
		// string. A reviewer decides whether to make it a reference.
		if target, ok := codegen.SiblingResource(field, opts); ok {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: fieldPath,
				Reason:    judgement.ReasonPossibleReferenceBySibling,
				Detail: "the name matches " + target + ", a resource this service declares; " +
					"confirm whether it should be a reference",
			})
		}

		unsupported := false
		if _, reason, ok := codegen.UnsupportedFieldMarker(field_.String()); ok {
			unsupported = true
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: fieldPath,
				Reason:    judgement.ReasonUnsupportedFieldType,
				Detail:    reason,
			})
		} else if opts.EmitMessageMaps {
			if protoMsg, goType, ok := codegen.MessageMapValueInfo(field); ok {
				out.Judgement = append(out.Judgement, JudgementItem{
					FieldPath: fieldPath,
					Reason:    judgement.ReasonMessageMapDerived,
					Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", goType, protoMsg),
				})
			}
		}
		if item, ok := judgementFor(field, fieldPath); ok {
			out.Judgement = append(out.Judgement, item)
		}
		jsonName := codegen.GetJSONForKRM(field, opts)
		if string(field.Name()) == "etag" {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: fieldPath,
				Reason:    judgement.ReasonEtagInSpec,
				Detail:    "etag was generated in the Spec because the proto does not mark it OUTPUT_ONLY; confirm whether it belongs in status.observedState",
			})
		} else if possibleStatusFieldNames[jsonName] && !codegen.IsFieldBehavior(field, annotations.FieldBehavior_REQUIRED) {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: fieldPath,
				Reason:    judgement.ReasonPossibleStatusFieldInSpec,
				Detail:    fmt.Sprintf("%s was generated in the Spec because the proto does not mark it OUTPUT_ONLY; confirm whether it is user-configurable desired state, a duplicate of spec.resourceID, or belongs in status.observedState", jsonName),
			})
		}
		if topLevelLocationFieldNames[jsonName] && field.Kind() == protoreflect.StringKind && !field.IsList() {
			out.Judgement = append(out.Judgement, JudgementItem{
				FieldPath: fieldPath,
				Reason:    judgement.ReasonLocationOrParentRef,
				Detail:    fmt.Sprintf("proto defines a %s field in Spec; confirm whether it should be a reference, replaced by a parent reference, or kept as a plain string", jsonName),
			})
		}
		if !unsupported {
			walkNestedSpecJudgements(field, fieldPath, opts, map[protoreflect.FullName]bool{msg.FullName(): true}, &out.Judgement)
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
func PrepopulateObservedState(details *codegen.OutputMessageDetails, observedStateMessages sets.String, opts codegen.WriteOptions) (fields string, items []JudgementItem) {
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
			items = append(items, JudgementItem{
				Reason: judgement.ReasonIdentityServerGeneratedID,
				Detail: "proto marks the resource .name field OUTPUT_ONLY; confirm whether the resource ID is server-generated or client-specified via a Create request *_id parameter",
			})
		default:
			if _, reason, ok := codegen.UnsupportedFieldMarker(n.Rendered); ok {
				items = append(items, JudgementItem{
					FieldPath: ".status.observedState." + n.JSONName,
					Reason:    judgement.ReasonUnsupportedFieldType,
					Detail:    reason,
				})
				continue
			}
			if opts.EmitMessageMaps && n.Field != nil {
				if protoMsg, goType, ok := codegen.MessageMapValueInfo(n.Field); ok {
					items = append(items, JudgementItem{
						FieldPath: ".status.observedState." + n.JSONName,
						Reason:    judgement.ReasonMessageMapDerived,
						Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", goType, protoMsg),
					})
				}
			}
			isOut := isOutputOrServerSetField(n.Field, details.Message, opts)
			onPath := map[protoreflect.FullName]bool{}
			if details.Message != nil {
				onPath[details.Message.FullName()] = true
			}
			walkObservedStateWellKnownProtoTypes(n.Field, ".status.observedState."+n.JSONName, opts, observedStateMessages, isOut || details.ForceAll, onPath, &items)
		}
	}

	return fields, items
}

func isOutputOrServerSetField(field protoreflect.FieldDescriptor, msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) bool {
	if field == nil {
		return false
	}
	return codegen.IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) ||
		codegen.IsServerSetField(field, msg, opts)
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
		Reason:    judgement.ReasonAuthoritativeProtoReference,
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

// OutputOnlyCandidate is a field the proto documents as output-only in prose
// while carrying no google.api.field_behavior annotation to say so.
type OutputOnlyCandidate struct {
	// FieldPath is the KRM path the field was emitted at, e.g. ".spec.createTime"
	// or, for a nested field, ".spec.config.state".
	FieldPath string
	// Reason is judgement.ReasonWellKnownOutputOnlyPattern when the comment starts with
	// one of outputOnlyPrefixes. It is judgement.ReasonPossibleOutputOnlyPattern, a
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
	targetPath := ".status.observedState." + strings.TrimPrefix(c.FieldPath, ".spec.")
	detail := "proto comment says output only but there is no field_behavior annotation, so it was generated into the Spec at " + c.FieldPath + "."
	if c.Reason == judgement.ReasonPossibleOutputOnlyPattern {
		detail = "proto comment mentions output only but doesn't start with \"Output only.\" or \"[Output Only]\", and there is no field_behavior annotation, so it was generated into the Spec at " + c.FieldPath + ". " +
			"The comment may be a typo, apply only some of the time, or mean something else."
	}
	return JudgementItem{
		FieldPath: targetPath,
		Reason:    c.Reason,
		Detail:    detail + " Move the field to " + targetPath + " if it is confirmed output only. Proto comment: " + c.Comment,
	}
}

// wellKnownProtoTypeItem returns a queue entry if field uses a google.protobuf.*
// message type that is generated as a Go struct rather than mapped to a scalar
// or special KRM type (e.g. google.protobuf.DoubleValue, Int32Value, Empty).
func wellKnownProtoTypeItem(field protoreflect.FieldDescriptor, path string) (JudgementItem, bool) {
	if field.ContainingMessage() != nil && strings.HasPrefix(string(field.ContainingMessage().FullName()), "google.protobuf.") {
		return JudgementItem{}, false
	}
	var subMsg protoreflect.MessageDescriptor
	switch {
	case field.IsMap():
		if v := field.MapValue(); v.Kind() == protoreflect.MessageKind {
			subMsg = v.Message()
		}
	case field.Kind() == protoreflect.MessageKind:
		subMsg = field.Message()
	}
	if subMsg == nil || !codegen.MapsToGoStruct(subMsg) {
		return JudgementItem{}, false
	}
	if !strings.HasPrefix(string(subMsg.FullName()), "google.protobuf.") {
		return JudgementItem{}, false
	}
	return JudgementItem{
		FieldPath: path,
		Reason:    judgement.ReasonWellKnownProtoType,
		Detail: fmt.Sprintf("field uses well-known protobuf type %s, which was generated as struct %s; confirm KRM representation",
			subMsg.FullName(), codegen.GoNameForProtoMessage(subMsg)),
	}, true
}

// WellKnownProtoTypes returns a queue entry for each Spec field, at any depth,
// that uses a google.protobuf.* message type generated as a Go struct.
func WellKnownProtoTypes(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []JudgementItem {
	if msg == nil {
		return nil
	}
	var out []JudgementItem
	walkSpecFields(msg, ".spec", opts, true, map[protoreflect.FullName]bool{}, func(path, _ string, field protoreflect.FieldDescriptor) {
		if item, ok := wellKnownProtoTypeItem(field, path); ok {
			out = append(out, item)
		}
	})
	return out
}

func walkObservedStateWellKnownProtoTypes(field protoreflect.FieldDescriptor, path string, opts codegen.WriteOptions, observedStateMessages sets.String, forceAll bool, onPath map[protoreflect.FullName]bool, out *[]JudgementItem) {
	if field == nil {
		return
	}
	if item, ok := wellKnownProtoTypeItem(field, path); ok {
		*out = append(*out, item)
	}
	var subMsg protoreflect.MessageDescriptor
	nextPrefix := path
	switch {
	case field.IsMap():
		if value := field.MapValue(); value.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(value.Message()) {
			subMsg = value.Message()
			nextPrefix = path + ".KEY"
		}
	case field.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(field.Message()):
		subMsg = field.Message()
		if field.IsList() {
			nextPrefix += "[]"
		}
	}
	if subMsg == nil || strings.HasPrefix(string(subMsg.FullName()), "google.protobuf.") || onPath[subMsg.FullName()] {
		return
	}
	onPath[subMsg.FullName()] = true
	defer delete(onPath, subMsg.FullName())
	subUsesObservedState := !field.IsMap() && observedStateMessages.Has(string(subMsg.FullName()))
	for i := 0; i < subMsg.Fields().Len(); i++ {
		child := subMsg.Fields().Get(i)
		childIsOut := codegen.IsFieldBehavior(child, annotations.FieldBehavior_OUTPUT_ONLY)
		childHasOutputs := child.Kind() == protoreflect.MessageKind && !child.IsMap() && observedStateMessages.Has(string(child.Message().FullName()))
		if subUsesObservedState {
			if !forceAll && !childIsOut && !childHasOutputs {
				continue
			}
		} else {
			if childIsOut || (opts.Prepopulating && codegen.IsFieldAllOutputOnlyMessage(child)) {
				continue
			}
		}
		childPath := nextPrefix + "." + codegen.GetJSONForKRM(child, opts)
		if _, err := codegen.GoTypeForField(child, false, opts); err != nil {
			*out = append(*out, JudgementItem{
				FieldPath: childPath,
				Reason:    judgement.ReasonUnsupportedFieldType,
				Detail:    err.Error(),
			})
			continue
		}
		if opts.EmitMessageMaps {
			if protoMsg, goType, ok := codegen.MessageMapValueInfo(child); ok {
				*out = append(*out, JudgementItem{
					FieldPath: childPath,
					Reason:    judgement.ReasonMessageMapDerived,
					Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", goType, protoMsg),
				})
			}
		}
		childObsMessages := observedStateMessages
		childForceAll := forceAll || childIsOut
		if !subUsesObservedState {
			childObsMessages = nil
			childForceAll = false
		}
		walkObservedStateWellKnownProtoTypes(child, childPath, opts, childObsMessages, childForceAll, onPath, out)
	}
}

func walkNestedSpecJudgements(field protoreflect.FieldDescriptor, path string, opts codegen.WriteOptions, onPath map[protoreflect.FullName]bool, out *[]JudgementItem) {
	if field == nil {
		return
	}
	goType, err := codegen.GoTypeForField(field, false, opts)
	if err != nil || generatesAsReference(field, goType) {
		return
	}
	var subMsg protoreflect.MessageDescriptor
	nextPrefix := path
	switch {
	case field.IsMap():
		if value := field.MapValue(); value.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(value.Message()) {
			subMsg = value.Message()
			nextPrefix = path + ".KEY"
		}
	case field.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(field.Message()):
		subMsg = field.Message()
		if field.IsList() {
			nextPrefix += "[]"
		}
	}
	if subMsg == nil || strings.HasPrefix(string(subMsg.FullName()), "google.protobuf.") || onPath[subMsg.FullName()] {
		return
	}
	onPath[subMsg.FullName()] = true
	defer delete(onPath, subMsg.FullName())

	for i := 0; i < subMsg.Fields().Len(); i++ {
		child := subMsg.Fields().Get(i)
		if codegen.IsFieldBehavior(child, annotations.FieldBehavior_OUTPUT_ONLY) {
			continue
		}
		if codegen.IsFieldAllOutputOnlyMessage(child) {
			continue
		}
		childPath := nextPrefix + "." + codegen.GetJSONForKRM(child, opts)
		childGoType, err := codegen.GoTypeForField(child, false, opts)
		if err != nil {
			*out = append(*out, JudgementItem{
				FieldPath: childPath,
				Reason:    judgement.ReasonUnsupportedFieldType,
				Detail:    err.Error(),
			})
			continue
		}
		if generatesAsReference(child, childGoType) {
			continue
		}
		if target, ok := codegen.SiblingResource(child, opts); ok {
			*out = append(*out, JudgementItem{
				FieldPath: childPath,
				Reason:    judgement.ReasonPossibleReferenceBySibling,
				Detail: "the name matches " + target + ", a resource this service declares; " +
					"confirm whether it should be a reference",
			})
		}
		if opts.EmitMessageMaps {
			if protoMsg, mapGoType, ok := codegen.MessageMapValueInfo(child); ok {
				*out = append(*out, JudgementItem{
					FieldPath: childPath,
					Reason:    judgement.ReasonMessageMapDerived,
					Detail:    fmt.Sprintf("derived as map[string]%s from proto %s; verify Go type and mapping", mapGoType, protoMsg),
				})
			}
		}
		childJSONName := codegen.GetJSONForKRM(child, opts)
		if nestedLocationFieldNames[childJSONName] && child.Kind() == protoreflect.StringKind && !child.IsList() {
			*out = append(*out, JudgementItem{
				FieldPath: childPath,
				Reason:    judgement.ReasonLocationOrParentRef,
				Detail:    fmt.Sprintf("proto defines a %s field in Spec; confirm whether it should be a reference, replaced by a parent reference, or kept as a plain string", childJSONName),
			})
		}
		walkNestedSpecJudgements(child, childPath, opts, onPath, out)
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
// judgement.ReasonWellKnownOutputOnlyPattern. Any other mention of "output only" gets
// judgement.ReasonPossibleOutputOnlyPattern. That covers the words after other text,
// as in websecurityscanner's managed_scan: "Whether the scan config is
// managed by Web Security Scanner, output only." It also covers a comment
// that starts with the words but not with a prefix. That is usually a typo,
// as in automl's "Output only . The", or a condition, as in spanner's
// Backup.name, which is output only for the CreateBackup operation and
// required for UpdateBackup.
func outputOnlyReason(comment string) (string, bool) {
	for _, prefix := range outputOnlyPrefixes {
		if len(comment) >= len(prefix) && strings.EqualFold(comment[:len(prefix)], prefix) {
			return judgement.ReasonWellKnownOutputOnlyPattern, true
		}
	}
	if outputOnlyMention.MatchString(comment) {
		return judgement.ReasonPossibleOutputOnlyPattern, true
	}
	return "", false
}
