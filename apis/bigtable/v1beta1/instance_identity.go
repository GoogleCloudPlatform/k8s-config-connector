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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

// InstanceIdentity defines the resource reference to BigtableInstance, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceIdentity struct {
	Parent *parent.ProjectParent
	Id     string
}

func (i *InstanceIdentity) String() string {
	return i.ParentString() + "/instances/" + i.Id
}

func (i *InstanceIdentity) ID() string {
	return i.Id
}

func (i *InstanceIdentity) ParentString() string {
	return i.Parent.String()
}

func ParseInstanceExternal(external string) (*parent.ProjectParent, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "instances" {
		return nil, "", fmt.Errorf("format of BigtableInstance external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}})", external)
	}

	return &parent.ProjectParent{ProjectID: tokens[1]}, tokens[3], nil
}
