// Copyright 2024 Google LLC
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

package resource

import (
	"fmt"
	"regexp"
	"strings"
)

type Variable struct {
	Expression    string
	CELExpression string
	ResolvedValue interface{}
}

var (
	referencesRegex = regexp.MustCompile(`\$\{[a-zA-Z0-9\.\-]*\}`)
)

func extractVariables(raw []byte) ([]*Variable, error) {
	matches := referencesRegex.FindAll(raw, -1)
	variables := make([]*Variable, len(matches))
	for i, match := range matches {
		variables[i] = &Variable{Expression: string(match), CELExpression: trimReferenceSyntax(string(match))}
	}
	return variables, nil
}

func trimReferenceSyntax(reference string) string {
	if !strings.HasPrefix(reference, "${") && !strings.HasSuffix(reference, "}") {
		return reference
	}
	reference = strings.TrimLeft(reference, "${")
	reference = strings.TrimRight(reference, "}")
	return reference
}

func regexExpression(expression string) string {
	expression = trimReferenceSyntax(expression)
	expression = fmt.Sprintf(`\$\{%s\}`, expression)
	return strings.ReplaceAll(expression, ".", `\.`)
}
