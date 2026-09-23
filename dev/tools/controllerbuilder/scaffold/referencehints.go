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
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/refs"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ReferenceHints inspects spec fields using heuristic rules (name matching,
// URI templates, loose descriptions) and proposes potential references for review.
func ReferenceHints(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []JudgementItem {
	var out []JudgementItem
	walkSpecFields(msg, ".spec", opts, true, map[protoreflect.FullName]bool{}, func(path, desc string) {
		if item, ok := referenceHint(path, desc); ok {
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
// depth, and at the top level the identity fields and server-set fields that
// PrepopulateSpec drops. A field the generator cannot type is absent from the
// CRD too, so it is skipped with its subtree. So is a field the generator
// already writes as a reference, which needs no hint.
//
// onPath holds the messages between msg and the root. Proto messages can
// contain themselves, and the generated struct breaks the cycle with a
// pointer, so a message already on the path is not entered again.
func walkSpecFields(msg protoreflect.MessageDescriptor, prefix string, opts codegen.WriteOptions, top bool, onPath map[protoreflect.FullName]bool, visit func(path, desc string)) {
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
		// Skip exactly what PrepopulateSpec drops, or a hint names a path the
		// CRD does not have. Today that is the identity fields and, where the
		// flag is on, the server-set ones. Both are only dropped from the
		// resource's own message, so both are checked only at the top.
		if top && identityFields[string(field.Name())] {
			continue
		}
		if top && codegen.IsServerSetField(field, msg, opts) {
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
		visit(path, fieldComment(field))

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
// strict description rule, then the loose one, then the name rules. Only an
// IsReference verdict ends the search, so a field Classify calls
// NotRepresentable can still match a later rule.
func referenceHint(path, desc string) (JudgementItem, bool) {
	if refs.IsReferenceFieldPath(path) {
		return JudgementItem{}, false
	}
	if verdict, _ := refs.Classify(path, desc); verdict == refs.IsReference {
		return JudgementItem{FieldPath: path, Reason: "possible-reference-by-description"}, true
	}
	if refs.MatchDescriptionLoose(path, desc) {
		return JudgementItem{FieldPath: path, Reason: "possible-reference-by-description-loose"}, true
	}
	if target, ok := refs.MatchName(path); ok {
		return JudgementItem{FieldPath: path, Reason: "possible-reference-by-name", Detail: target}, true
	}
	return JudgementItem{}, false
}
