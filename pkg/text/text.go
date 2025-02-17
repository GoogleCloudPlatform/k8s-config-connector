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

package text

import (
	"regexp"
	"strings"
	"unicode"
)

// TODO: there must be some package out there that does this well
func Pluralize(singular string) string {
	// special case for IAPSettings, as "settings" is already a plural of "setting", adding another "es" would be unnecessary
	if singular == "IAPSettings" {
		return "IAPSettings"
	}

	var plural string
	if strings.HasSuffix(singular, "y") {
		if strings.HasSuffix(singular, "ay") || strings.HasSuffix(singular, "ey") {
			plural = singular + "s"
		} else {
			plural = singular[:len(singular)-1] + "ies"
		}
	} else if strings.HasSuffix(singular, "s") || strings.HasSuffix(singular, "x") || strings.HasSuffix(singular, "sh") {
		plural = singular + "es"
	} else if singular != "" {
		plural = singular + "s"
	}
	return plural
}

func SnakeCaseStrsToLowerCamelCaseStrs(strs []string) []string {
	out := make([]string, 0)
	for _, s := range strs {
		out = append(out, SnakeCaseToLowerCamelCase(s))
	}
	return out
}

func SnakeCaseToLowerCamelCase(s string) string {
	return snakeCaseToUpperCamelCase(s)
}

func SnakeCaseToUpperCamelCase(s string) string {
	return strings.Title(snakeCaseToUpperCamelCase(s))
}

// Convert a snake_case_string to CamelCaseString.
func snakeCaseToUpperCamelCase(s string) string {
	split := strings.Split(s, "_")
	ret := ""
	for i, v := range split {
		if i == 0 {
			ret += v
		} else {
			ret += strings.Title(v)
		}
	}
	return ret
}

func SnakeCaseToLowerCase(s string) string {
	split := strings.Split(s, "_")
	ret := ""
	for _, v := range split {
		ret += v
	}
	return ret
}

// AsSnakeCase returns the given string converted to lowercase snake_case. If the input is already snake_case, no
// change is made. Any transitions in the input from lowercase to uppercase are interpreted as camelCase-style word
// transitions, and are replaced with an underscore.
func AsSnakeCase(s string) string {
	res := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(res, "${1}_${2}"))
}

// Convert a CamelCaseString to kebab-case-string.
func CamelCaseToKebabCase(s string) string {
	kebabed := ""
	for i, r := range s {
		if i != 0 && i != len(s)-1 && unicode.IsUpper(r) {
			kebabed += "-"
		}
		kebabed += strings.ToLower(string(r))
	}
	return kebabed
}

// Convert a kebab-case-string to lower_snake_case string.
func KebabCaseToLowerSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "-", "_"))
}

func SnakeCaseToKebabCase(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

func BeginsWithVowel(s string) bool {
	if s == "" {
		return false
	}
	switch strings.ToLower(string(s[0])) {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

func UppercaseInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LowercaseInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func IndefiniteArticleFor(s string) string {
	if BeginsWithVowel(s) {
		return "an"
	}
	return "a"
}

func AppendStrAsNewParagraph(base, str string) string {
	if base == "" {
		return str
	}
	return base + "\n\n" + str
}

func IsPascalCase(s string) bool {
	sampleRegex := regexp.MustCompile("^[A-Z][a-z]*([A-Z][a-z]*)*$")
	return sampleRegex.Match([]byte(s))
}

func IsSnakeCase(s string) bool {
	stringRegex := regexp.MustCompile("^[a-z1-9]+(_[a-z1-9]+)*$")
	return stringRegex.Match([]byte(s))
}

func RemoveSpecialCharacters(s string) string {
	specialCharRegex := regexp.MustCompile(`[^0-9A-Za-z ]+`)
	return specialCharRegex.ReplaceAllString(s, "")
}
