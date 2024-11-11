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
	"fmt"
	"slices"
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

// placeholderForGCPResource returns the placeholder we use for the value, if we recognize the GCP resource type
func (r *Replacements) placeholderForGCPResource(resource string) string {
	switch resource {
	case "addresses":
		return "${addressID}"
	case "tensorboards":
		return "${tensorboardID}"
	case "tagKeys":
		return "${tagKeyID}"
	case "tagValues":
		return "${tagValueID}"
	case "datasets":
		return "${datasetID}"
	case "networks":
		return "${networkID}"
	case "subnetworks":
		return "${subnetworkID}"
	case "notificationChannels":
		return "${notificationChannelID}"
	case "alertPolicies":
		return "${alertPolicyID}"
	case "billingAccounts":
		return "${billingAccountID}"
	case "conditions":
		return "${conditionID}"
	case "exclusions":
		return "${exclusionID}"
	case "forwardingRules":
		return "${forwardingRuleID}"
	case "groups":
		return "${groupID}"
	case "jobs":
		return "${jobID}"
	case "uptimeCheckConfigs":
		return "${uptimeCheckConfigID}"
	case "operations":
		return "${operationID}"
	case "transferConfigs":
		return "${transferConfigID}"
	case "firewallPolicies":
		return "${firewallPolicyID}"
	case "folders":
		return "${folderID}"
	case "sslCertificates":
		return "${sslCertificateID}"
	case "serviceAttachments":
		return "${serviceAttachmentID}"
	case "targetGrpcProxies":
		return "${targetGrpcProxyID}"
	case "targetTcpProxies":
		return "${targetTcpProxyID}"
	case "targetHttpsProxies":
		return "${targetHttpsProxyID}"
	case "targetSslProxies":
		return "${targetSslProxyID}"
	default:
		return ""
	}
}

// ExtractIDsFromLinks parses the URL or partial URL, and extracts generated IDs from it.
func (r *Replacements) ExtractIDsFromLinks(link string) {
	u, _ := ParseGCPLink(link)
	if u != nil {
		for _, item := range u.PathItems {
			placeholder := r.placeholderForGCPResource(item.Resource)
			if placeholder != "" {
				r.PathIDs[item.Name] = placeholder
			}

			// Special case for operations
			// TODO: Can we get rid of this?
			if item.Resource == "operations" {
				r.OperationIDs[item.Name] = true
			}
		}
	}
}

type GCPLink struct {
	PathItems []PathItem
}

type PathItem struct {
	Resource string
	Name     string
}

func ParseGCPLink(link string) (*GCPLink, error) {
	ret := &GCPLink{}

	tokens := strings.Split(link, "/")

	// Consider the last two tokens, in pairs
	for len(tokens) >= 2 {
		n := len(tokens)
		kind := tokens[n-2]
		id := tokens[n-1]
		if id == "" {
			break
		}

		// Remove any "verbs" we might be picking up by mistake
		// e.g. https://cloudresourcemanager.googleapis.com/v3/folders/${folderID}:move?alt=json&prettyPrint=false
		if strings.Contains(id, ":") {
			id = strings.Split(id, ":")[0]
		}

		// Advance by 2 tokens, unless this is one of the special-case GCP resources
		if id == "global" {
			tokens = tokens[:n-1]
			ret.PathItems = append(ret.PathItems, PathItem{Resource: "", Name: "global"})
		} else {
			tokens = tokens[:n-2]
			ret.PathItems = append(ret.PathItems, PathItem{Resource: kind, Name: id})
		}
	}

	if len(ret.PathItems) == 0 {
		return nil, fmt.Errorf("no items found in link %q", link)
	}

	// Return in path order
	slices.Reverse(ret.PathItems)

	return ret, nil
}
