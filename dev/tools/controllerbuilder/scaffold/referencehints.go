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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/refs"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ReferenceHints inspects spec fields using heuristic rules (name matching,
// URI templates, loose descriptions) and proposes potential references for review.
//
// It also queues a field that names another resource in a way KCC cannot
// express as a reference, so a reviewer sees why it stays a string, and a
// field whose name says it holds a password. A nested field with a
// google.api.resource_reference annotation gets the same entry that
// PrepopulateSpec gives a top-level one.
func ReferenceHints(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []JudgementItem {
	var out []JudgementItem
	walkSpecFields(msg, ".spec", opts, true, map[protoreflect.FullName]bool{}, func(path, desc string, field protoreflect.FieldDescriptor) {
		// PrepopulateSpec already checks the annotation on top-level fields.
		// The walk never re-enters msg, so any field of msg is top-level.
		if field.ContainingMessage().FullName() != msg.FullName() {
			if item, ok := judgementFor(field, path); ok {
				out = append(out, item)
			}
		}
		if item, ok := referenceHint(path, desc); ok {
			out = append(out, item)
		}
		if item, ok := sensitiveField(path); ok {
			out = append(out, item)
		}
	})
	return out
}

// walkSpecFields calls visit with the CRD path and proto comment of every field
// of msg that the generator writes into a Spec struct, then descends into
// message fields the same way.
//
// It skips what the generator leaves out of the Spec: OUTPUT_ONLY fields at any
// depth, and at the top level the identity, server-set and deprecated fields
// PrepopulateSpec drops. A field the generator cannot type is absent from the
// CRD too, so it is skipped with its subtree. So is a field the generator
// already writes as a reference, which needs no hint.
//
// onPath holds the messages between msg and the root. Proto messages can
// contain themselves, and the generated struct breaks the cycle with a
// pointer, so a message already on the path is not entered again.
func walkSpecFields(msg protoreflect.MessageDescriptor, prefix string, opts codegen.WriteOptions, top bool, onPath map[protoreflect.FullName]bool, visit func(path, desc string, field protoreflect.FieldDescriptor)) {
	if onPath[msg.FullName()] {
		return
	}
	onPath[msg.FullName()] = true
	defer delete(onPath, msg.FullName())

	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if codegen.IsFieldBehavior(field, annotations.FieldBehavior_OUTPUT_ONLY) {
			continue
		}
		// Skip the fields PrepopulateSpec drops, or a hint names a path the CRD
		// does not have: identity fields, server-set fields when
		// --place-server-set-fields is enabled, and deprecated fields.
		// PrepopulateSpec only drops them from the resource's top-level
		// message, so they are only checked at the top.
		if top && topLevelSkip(field, msg, opts) != inSpec {
			continue
		}
		goType, err := codegen.GoTypeForField(field, false, opts)
		if err != nil {
			continue
		}
		if generatesAsReference(field, goType) {
			continue
		}

		path := prefix + "." + codegen.GetJSONForKRM(field, opts)
		visit(path, fieldComment(field), field)

		switch {
		case field.IsMap():
			if value := field.MapValue(); value.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(value.Message()) {
				walkSpecFields(value.Message(), path+".KEY", opts, false, onPath, visit)
			}
		case field.Kind() == protoreflect.MessageKind && codegen.MapsToGoStruct(field.Message()):
			if field.IsList() {
				path += "[]"
			}
			walkSpecFields(field.Message(), path, opts, false, onPath, visit)
		}
	}
}

// generatesAsReference reports whether the generator writes field as a KCC
// reference type instead of a struct of its own, as it writes
// google.cloud.connectors.v1.Secret as *secretmanagerv1beta1.SecretRef.
// goType is what GoTypeForField returned for field.
//
// The Ref suffix alone is not enough. DeployedModelRef and
// NotebookRuntimeTemplateRef are ordinary generated structs whose proto names
// end in Ref, and their fields still need hints.
func generatesAsReference(field protoreflect.FieldDescriptor, goType string) bool {
	return field.Kind() == protoreflect.MessageKind &&
		!codegen.MapsToGoStruct(field.Message()) &&
		strings.HasSuffix(goType, "Ref")
}

// fieldComment returns a field's leading proto comment on one line, which is
// the text the CRD carries as the field's description.
func fieldComment(field protoreflect.FieldDescriptor) string {
	loc := field.ParentFile().SourceLocations().ByDescriptor(field)
	return strings.Join(strings.Fields(loc.LeadingComments), " ")
}

// referenceHint returns the queue entry the rules give a field, trying the
// strict description rule, then the loose one, then the name rules. A verdict
// from the strict rule ends the search. For IsReference, the entry says which
// of Classify's rules matched. For NotRepresentable, the entry says the field
// cannot be a reference, instead of a looser hint that would contradict that.
func referenceHint(path, desc string) (JudgementItem, bool) {
	if refs.IsReferenceFieldPath(path) {
		return JudgementItem{}, false
	}
	switch verdict, reason := refs.Classify(path, desc); verdict {
	case refs.IsReference:
		return JudgementItem{
			FieldPath: path,
			Reason:    judgement.ReasonPossibleReferenceByDescription,
			Detail:    strings.Join(refs.ReferenceRules(path, desc), "; "),
		}, true
	case refs.NotRepresentable:
		return JudgementItem{
			FieldPath: path,
			Reason:    judgement.ReasonReferenceNotRepresentable,
			Detail:    reason + ": names another resource, but KCC cannot express it as a reference today, so it stays a string",
		}, true
	}
	if refs.MatchDescriptionLoose(path, desc) {
		return JudgementItem{FieldPath: path, Reason: judgement.ReasonPossibleReferenceByDescriptionLoose}, true
	}
	if target, ok := refs.MatchName(path); ok {
		return JudgementItem{FieldPath: path, Reason: judgement.ReasonPossibleReferenceByName, Detail: target}, true
	}
	return JudgementItem{}, false
}

// sensitiveField returns a queue entry for a field whose name says it holds a
// password. It uses the rule from TestNoSensitiveField in tests/apichecks: the
// KRM path, lowercased, ends in "password". The generator still writes the
// field as a plain value. We haven't settled how a new Kind should take a
// secret, so the entry only flags the field.
func sensitiveField(path string) (JudgementItem, bool) {
	if !strings.HasSuffix(strings.ToLower(path), "password") {
		return JudgementItem{}, false
	}
	return JudgementItem{
		FieldPath: path,
		Reason:    "sensitive-field",
		Detail:    "the name says this holds a password, and it is generated as a plain value. Leave this open until we settle how new Kinds take secrets",
	}, true
}
