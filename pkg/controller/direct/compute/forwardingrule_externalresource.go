/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package compute

import (
	"fmt"
	"strings"
)

const (
	serviceDomain = "//compute.googleapis.com"
)

type ForwardingRuleIdentity struct {
	project        string
	location       string
	forwardingRule string
}

// FullyQualifiedName builds a ForwardingRuleIdentity resource
func (c *ForwardingRuleIdentity) FullyQualifiedName() string {
	if c.location == "global" {
		return fmt.Sprintf("projects/%s/global/forwardingrules/%s", c.project, c.forwardingRule)
	} else {
		return fmt.Sprintf("projects/%s/regions/%s/forwardingrules/%s", c.project, c.location, c.forwardingRule)
	}
}

// AsExternalRef builds a externalRef from a ForwardingRuleIdentity
func (c *ForwardingRuleIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a ForwardingRuleIdentity from a externalRef
func asID(externalRef string) (*ForwardingRuleIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "forwardingrules" {
		return &ForwardingRuleIdentity{
			project:        tokens[1],
			location:       "global",
			forwardingRule: tokens[4],
		}, nil
	}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "forwardingrules" {
		return &ForwardingRuleIdentity{
			project:        tokens[1],
			location:       tokens[3],
			forwardingRule: tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("ExternalRef format invalid: %s", externalRef)
}

// BuildID builds a ForwardingRuleIdentity from resource components.
func BuildID(project, location, forwardingRule string) *ForwardingRuleIdentity {
	return &ForwardingRuleIdentity{
		project:        project,
		location:       location,
		forwardingRule: forwardingRule,
	}
}
