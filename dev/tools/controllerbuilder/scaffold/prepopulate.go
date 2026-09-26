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
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/util/sets"
)

// identityFields are proto fields that the KRM object expresses through its own
// identity rather than as spec fields. "name" is the resource's own resource
// name, which KCC models as metadata.name / spec.resourceID plus the parent refs.
var identityFields = map[string]bool{
	"name": true,
}

// PrepopulateResult contains rendered struct bodies produced during prepopulation.
type PrepopulateResult struct {
	// SpecFields is the rendered Go source for the Spec struct fields.
	SpecFields string
	// ObservedStateFields is the rendered Go source for the ObservedState struct fields,
	// empty when no fields are marked OUTPUT_ONLY in the proto.
	ObservedStateFields string
	// ExtraImports contains necessary import paths beyond the default template imports.
	ExtraImports []string
}

// PrepopulateSpec renders the top-level Spec fields for a resource proto message.
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

		codegen.WriteField(&buf, field, msg, emitted, false, opts)
		emitted++
	}

	out.SpecFields = buf.String()
	return out, nil
}

// PrepopulateObservedState generates the struct body for the resource-level
// <Kind>ObservedState from output message details.
func PrepopulateObservedState(details *codegen.OutputMessageDetails, observedStateMessages sets.String, opts codegen.WriteOptions) string {
	if details == nil {
		return ""
	}

	var buf bytes.Buffer
	// identityFields leaves out "name", which KCC carries in status.externalRef,
	// even where the proto marks it OUTPUT_ONLY.
	codegen.WriteObservedStateFields(&buf, details, observedStateMessages, identityFields, opts)
	return buf.String()
}

// ExtraImportsFor returns import lines for any packages referenced by the
// rendered Spec and ObservedState fields that are not in the template defaults.
func ExtraImportsFor(specFields, observedStateFields string) []string {
	body := specFields + "\n" + observedStateFields
	var imports []string
	for qualifier, pkgPath := range codegen.QualifierImports {
		if strings.Contains(body, qualifier+".") {
			imports = append(imports, fmt.Sprintf("%s %q", qualifier, pkgPath))
		}
	}
	sort.Strings(imports)
	return imports
}
