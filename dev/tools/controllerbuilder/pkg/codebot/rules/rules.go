// Copyright 2025 Google LLC
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

// holds code style and guidance coding
package rules

import "strings"

var (
	NoNakedReturns = Rule("Never use naked returns.")
	CRDShortNames = Rule (`
	For API type generation it is very important to add the right shortname labels. Good example:
	// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeenvgroup;gcpapigeeenvgroups

	Figure out the right pluralization for your resources!
	`)
)

type Rule string

func ApplyRule(prompt string, rule Rule) string {
	trimmedRule := strings.TrimSpace(string(rule))
	if trimmedRule != "" {
		return prompt + "\n" + string(rule)
	}
	return prompt
}

// ApplyRules applies multiple rules to a system prompt.
func ApplyRules(prompt string, rules []Rule) string {
    modifiedPrompt := prompt
	modifiedPrompt += "\n Here are a few rules to keep in mind as you code in this codebase" + 
	"in order to match the style and other coding conventions"
    for _, rule := range rules {
        modifiedPrompt = ApplyRule(modifiedPrompt, rule)
    }
    return modifiedPrompt
}