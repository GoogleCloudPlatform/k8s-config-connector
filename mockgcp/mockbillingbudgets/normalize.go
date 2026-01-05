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

package mockbillingbudgets

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const EtagPlaceholder = "abcdef0123A="

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Budget
	replacements.ReplacePath(".etag", EtagPlaceholder)
	replacements.ReplacePath(".budgets[].etag", EtagPlaceholder)

	// // We randomize the order of projects in the budget filter, so we need to sort them for normalization.
	// // Note this affects http.log, not what the client sees.
	// replacements.SortSlice(".budgetFilter.projects")
}

// matchLink will match the link to the pattern, and add any replacements if the link matches the pattern.
func matchLink(link string, pattern string, replacements mockgcpregistry.NormalizingVisitor) (bool, map[string]string) {
	tokens := strings.Split(link, "/")
	patternTokens := strings.Split(pattern, "/")
	if len(tokens) != len(patternTokens) {
		return false, nil
	}
	m := make(map[string]string)
	for i, token := range tokens {
		patternToken := patternTokens[i]
		if strings.HasPrefix(patternToken, "{") {
			key := strings.Trim(patternToken, "{}")
			m[key] = token
		} else if patternToken != token {
			return false, nil
		}
	}

	for k, v := range m {
		replacements.ReplaceStringValue(v, "${"+k+"}")
	}
	return true, m
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	previsitLink := func(link string) {
		matchLink(link, "billingAccounts/{billingAccountId}/budgets/{billingAccountBudgetId}", replacements)
	}
	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".name":
			previsitLink(value)
		}
	})
}
