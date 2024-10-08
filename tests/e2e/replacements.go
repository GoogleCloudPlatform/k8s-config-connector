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

package e2e

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

// Replacements manages replacements of dynamic values, like resource IDs
type Replacements struct {
	PathIDs      map[string]string
	OperationIDs map[string]bool
}

// NewReplacements is a constructor for Replacements
func NewReplacements() *Replacements {
	return &Replacements{
		PathIDs:      make(map[string]string),
		OperationIDs: make(map[string]bool),
	}
}

func (r *Replacements) ApplyReplacementsToHTTPEvents(events test.LogEntries) {
	for _, event := range events {
		event.Request.Body = r.ApplyReplacements(event.Request.Body)
		event.Request.URL = r.ApplyReplacements(event.Request.URL)
		event.Response.Body = r.ApplyReplacements(event.Response.Body)
	}
}

func (r *Replacements) ApplyReplacements(s string) string {
	normalizers := []func(string) string{}
	for k, v := range r.PathIDs {
		normalizers = append(normalizers, ReplaceString(k, v))
	}
	for k := range r.OperationIDs {
		normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
	}
	for _, normalizer := range normalizers {
		s = normalizer(s)
	}
	return s
}

// ExtractIDsFromLinks parses the URL or partial URL, and extracts generated IDs from it.
func (r *Replacements) ExtractIDsFromLinks(link string) {
	tokens := strings.Split(link, "/")
	for len(tokens) >= 2 {
		n := len(tokens)
		kind := tokens[n-2]
		id := tokens[n-1]
		if id == "" {
			break
		}
		switch kind {
		case "tensorboards":
			r.PathIDs[id] = "${tensorboardID}"
		case "tagKeys":
			r.PathIDs[id] = "${tagKeyID}"
		case "tagValues":
			r.PathIDs[id] = "${tagValueID}"
		case "datasets":
			r.PathIDs[id] = "${datasetID}"
		case "networks":
			r.PathIDs[id] = "${networkID}"
		case "subnetworks":
			r.PathIDs[id] = "${subnetworkID}"
		case "notificationChannels":
			r.PathIDs[id] = "${notificationChannelID}"
		case "alertPolicies":
			r.PathIDs[id] = "${alertPolicyID}"
		case "billingAccounts":
			r.PathIDs[id] = "${billingAccountID}"
		case "conditions":
			r.PathIDs[id] = "${conditionID}"
		case "exclusions":
			r.PathIDs[id] = "${exclusionID}"
		case "forwardingRules":
			r.PathIDs[id] = "${forwardingRuleID}"
		case "groups":
			r.PathIDs[id] = "${groupID}"
		case "jobs":
			r.PathIDs[id] = "${jobID}"
		case "uptimeCheckConfigs":
			r.PathIDs[id] = "${uptimeCheckConfigId}"
		case "operations":
			r.OperationIDs[id] = true
			r.PathIDs[id] = "${operationID}"
		case "transferConfigs":
			r.PathIDs[id] = "${transferConfigID}"
		}
		tokens = tokens[:n-2]
	}
}
