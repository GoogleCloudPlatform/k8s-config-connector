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
	"strings"
	"unicode"
)

// Plural turns a lowercased singular noun (such as a Kind shortName stem) into
// its English plural. Both the API type template (shortName generation) and
// TestCRDShortNamePluralization in tests/apichecks use this function.
func Plural(singular string) string {
	// Words that are already plural or uncountable stay unchanged.
	for _, word := range []string{"settings", "metrics", "series", "data"} {
		if strings.HasSuffix(singular, word) {
			return singular
		}
	}

	// Corpus pluralizes to corpora (e.g. vertexairagcorpus -> vertexairagcorpora).
	if strings.HasSuffix(singular, "corpus") {
		return strings.TrimSuffix(singular, "corpus") + "corpora"
	}

	// Rule 1: If singular ends with 's', 'x', 'z', 'ch', 'sh', add 'es'.
	if strings.HasSuffix(singular, "s") ||
		strings.HasSuffix(singular, "x") ||
		strings.HasSuffix(singular, "z") ||
		strings.HasSuffix(singular, "ch") ||
		strings.HasSuffix(singular, "sh") {
		return singular + "es"
	}

	// Rule 2: If singular ends with 'y' preceded by a consonant, change 'y' to 'ies'.
	if strings.HasSuffix(singular, "y") && len(singular) > 1 {
		r := rune(singular[len(singular)-2])
		if !isVowel(r) {
			return singular[:len(singular)-1] + "ies"
		}
	}

	// Rule 3: If singular ends with 'f' or 'fe', change to 'ves'.
	if strings.HasSuffix(singular, "f") {
		return singular[:len(singular)-1] + "ves"
	}
	if strings.HasSuffix(singular, "fe") {
		return singular[:len(singular)-2] + "ves"
	}

	// Rule 4: Default case - just add 's'.
	return singular + "s"
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
