// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not not use this file except in compliance with the License.
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

// TODO: We should verify the ComputeImageIdentity before use.
// Currently this struct is not for the ComputeImage direct controller but only used as reference validation (FromExternal)

var _ identity.Identity = &ComputeImageIdentity{}

// +k8s:deepcopy-gen=false
type ComputeImageIdentity struct {
	ProjectID string

	// Name and Family are mutually exclusive.
	Name   string
	Family string
}

func (i *ComputeImageIdentity) String() string {
	if i.Family != "" {
		return fmt.Sprintf("projects/%s/global/images/family/%s", i.ProjectID, i.Family)
	}
	return fmt.Sprintf("projects/%s/global/images/%s", i.ProjectID, i.Name)
}

func (i *ComputeImageIdentity) FromExternal(s string) error {
	if s == "" {
		return fmt.Errorf("value cannot be empty")
	}
	tokens := strings.Split(s, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "images" && tokens[4] == "family" {
		i.ProjectID = tokens[1]
		i.Family = tokens[5]
		i.Name = ""
		return nil
	}

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "images" {
		i.ProjectID = tokens[1]
		i.Family = ""
		i.Name = tokens[4]
		return nil
	}
	return fmt.Errorf("invalid format: %s, expected projects/{project}/global/images/{name} or projects/{project}/global/images/family/{family}", s)
}
