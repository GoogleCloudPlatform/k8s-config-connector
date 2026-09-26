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

package codegen

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// SiblingGuessMarker is written above a field whose name matches a resource the
// same service declares. It is a comment marker, so controller-gen strips it
// before the CRD is published and it cannot affect the schema.
const SiblingGuessMarker = "+kcc:guess=possible-reference target="

// SiblingResource reports the resource this service declares whose name matches
// the field's, if any.
//
// siblings maps a lowercased Kind suffix to the Kind: DiscoveryEngineDataStore
// is keyed "datastore", so a field called dataStore matches it. The rule keeps
// no list of known names, so it works on new services.
//
// The field name must match the key exactly (or match after singularizing plural forms).
func SiblingResource(field protoreflect.FieldDescriptor, opts WriteOptions) (string, bool) {
	if len(opts.Siblings) == 0 || field.Kind() != protoreflect.StringKind {
		return "", false
	}
	return SiblingResourceByName(GetJSONForKRM(field, opts), opts.Siblings)
}

// SiblingResourceByName takes the name directly, for callers with no proto
// field to read it from.
func SiblingResourceByName(name string, siblings map[string]string) (string, bool) {
	leaf := strings.ToLower(name)
	if target, ok := siblings[leaf]; ok {
		return target, true
	}
	// Check singular form for repeated fields (e.g. subnetworks -> subnetwork).
	if s := Singular(leaf); s != leaf {
		if target, ok := siblings[s]; ok {
			return target, true
		}
	}
	return "", false
}

// SiblingGuess is one field flagged by the sibling rule while writing a nested
// message, for the judgement queue.
type SiblingGuess struct {
	// Message is the proto message the field belongs to, e.g.
	// google.cloud.discoveryengine.v1.Control.
	Message string
	// Field is the KRM json name.
	Field string
	// Target is the sibling Kind the name matched.
	Target string
}

// scanSiblingGuesses recovers the sibling markers WriteMessage left in a
// rendered body. WriteField writes to an io.Writer, so it has nowhere to
// collect the markers it emits, and we read them back out of the finished text
// instead.
// scanUnsupported recovers the "// TODO:" markers the same way.
func scanSiblingGuesses(msgName, body string) []SiblingGuess {
	var out []SiblingGuess
	pending := ""
	for _, line := range strings.Split(body, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			if _, target, found := strings.Cut(trimmed, SiblingGuessMarker); found {
				pending = target
			}
			continue
		}
		if pending == "" {
			continue
		}
		if _, after, found := strings.Cut(line, "`json:\""); found {
			name, _, _ := strings.Cut(after, ",")
			out = append(out, SiblingGuess{Message: msgName, Field: name, Target: pending})
			pending = ""
		}
	}
	return out
}

// kindDeclaration matches the Spec struct each Kind's types file declares,
// which is how a package states which Kinds it holds.
var kindDeclaration = regexp.MustCompile(`(?m)^type (\w+)Spec struct`)

// SiblingKinds maps a lowercased field-name candidate to the Kind of a resource
// the target package declares. It builds the map SiblingResource reads.
//
// It reads two sources, because neither is complete on its own. The
// invocation's own resource list covers only its slice of a service generated
// by several generate-types calls: discoveryengine runs twice, dialogflow four
// times. The package scan covers the whole service, but finds nothing after a
// wipe-based regeneration, which deletes every _types.go before the generator
// runs. The map is the union of the two.
//
// The key strips the service prefix, so DiscoveryEngineDataStore is keyed
// "datastore" and a field called dataStore matches it.
func SiblingKinds(dir, service string, alsoKnown ...string) map[string]string {
	out := map[string]string{}
	add := func(kind string) {
		trimmed := strings.TrimPrefix(strings.ToLower(kind), strings.ToLower(service))
		if trimmed != "" && trimmed != strings.ToLower(kind) {
			out[trimmed] = kind
		}
	}
	for _, kind := range alsoKnown {
		add(kind)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return out
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), "_types.go") {
			continue
		}
		body, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		for _, m := range kindDeclaration.FindAllStringSubmatch(string(body), -1) {
			add(m[1])
		}
	}
	return out
}
