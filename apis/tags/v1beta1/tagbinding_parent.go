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

package v1beta1

import (
	"fmt"
	"strings"
)

// +k8s:deepcopy-gen=false
type TagBindingParent interface {
	String() string
	Prefix() string

	FromExternal(ref string) error
}

var ParentFullResourceMap = map[FullResourcePrefix]TagBindingParent{
	ProjectPrefix: &TagBindingProject{},
	OrgPrefix:     &TagBindingOrganization{},
}

type FullResourcePrefix string

var ProjectPrefix FullResourcePrefix = "//cloudresourcemanager.googleapis.com/projects"

var _ TagBindingParent = &TagBindingProject{}

// TagBindingProject represents a TagBinding parent of the form:
// `//cloudresourcemanager.googleapis.com/projects/{{projectNumber}}`
type TagBindingProject struct {
	ProjectNumber string
}

func (p *TagBindingProject) String() string {
	return fmt.Sprintf("%s/%s", ProjectPrefix, p.ProjectNumber)
}

func (p *TagBindingProject) Prefix() string {
	return fmt.Sprintf("%s", ProjectPrefix)
}

func (p *TagBindingProject) FromExternal(ref string) error {
	var found bool
	p.ProjectNumber, found = strings.CutPrefix(ref, p.Prefix()+"/")
	if !found {
		return fmt.Errorf("format of TagBindingProject missing prefix %q. got=%q", p.Prefix(), ref)
	}
	if p.ProjectNumber == "" {
		return fmt.Errorf("projectID was empty in TagBinding parent external=%q", ref)
	}
	return nil
}

var OrgPrefix FullResourcePrefix = "//cloudresourcemanager.googleapis.com/organizations"

var _ TagBindingParent = &TagBindingOrganization{}

// TagBindingOrganization represents a TagBinding parent of the form:
// `//cloudresourcemanager.googleapis.com/projects/{{OrgNumber}}`
type TagBindingOrganization struct {
	OrgNumber string
}

func (p *TagBindingOrganization) String() string {
	return fmt.Sprintf("%s/%s", OrgPrefix, p.OrgNumber)
}

func (p *TagBindingOrganization) Prefix() string {
	return fmt.Sprintf("%s", OrgPrefix)
}

func (p *TagBindingOrganization) FromExternal(ref string) error {
	var found bool
	p.OrgNumber, found = strings.CutPrefix(ref, p.Prefix()+"/")
	if !found {
		return fmt.Errorf("format of TagBindingOrganization missing prefix %q. got=%q", p.Prefix(), ref)
	}
	if p.OrgNumber == "" {
		return fmt.Errorf("organizationID was empty in TagBinding parent external=%q", ref)
	}
	return nil
}
