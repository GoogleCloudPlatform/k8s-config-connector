// Copyright 2022 Google LLC
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

package k8s

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const maxSubdomainLength = 253

var notLegalRegex = regexp.MustCompile("[^a-z0-9\\-\\.]+")

func ValueToDNSSubdomainName(value string) string {
	// As per https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names, the
	// following rules apply:
	// * contain no more than 253 characters
	// * contain only lowercase alphanumeric characters, '-' or '.'
	// * start with an alphanumeric character
	// * end with an alphanumeric character
	value = strings.ToLower(value)
	value = replaceIllegalPunctuation(value)
	value = toAlphaNumeric(value)
	value = ensureFirstAndLastCharactersAreAlphaNumeric(value)
	if len(value) > maxSubdomainLength {
		value = value[0:maxSubdomainLength]
	}
	return value
}

func replaceIllegalPunctuation(value string) string {
	sb := strings.Builder{}
	for _, runeValue := range value {
		if unicode.IsPunct(runeValue) {
			if !isLegalPunctuation(runeValue) {
				runeValue = '-'
			}
		}
		sb.WriteRune(runeValue)
	}
	return sb.String()
}

func isLegalPunctuation(r rune) bool {
	return r == '-' || r == '.'
}

func toAlphaNumeric(value string) string {
	return notLegalRegex.ReplaceAllString(value, "")
}

func ensureFirstAndLastCharactersAreAlphaNumeric(value string) string {
	defaultCharValue := "a"
	if value == "" {
		return defaultCharValue
	}
	runes := []rune(value)
	if !isAlphaNumeric(runes[0]) {
		value = fmt.Sprintf("%v%v", defaultCharValue, value)
	}
	if !isAlphaNumeric(runes[len(runes)-1]) {
		value = fmt.Sprintf("%v%v", value, defaultCharValue)
	}
	return value
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
