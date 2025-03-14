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

// TableIdentity defines the resource reference to BigtableTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableIdentity struct {
	Parent *InstanceIdentity
	Id     string
}

func (i *TableIdentity) String() string {
	return i.Parent.String() + "/tables/" + i.Id
}

func (i *TableIdentity) ID() string {
	return i.Id
}

func ParseTableExternal(external string) (*InstanceIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" {
		return nil, "", fmt.Errorf("format of BigtableTable external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}})", external)
	}
	p := &InstanceIdentity{
		Parent: &parent.ProjectParent{
			ProjectID: tokens[1],
		},
		Id: tokens[3],
	}
	resourceID := tokens[5]
	return p, resourceID, nil
}
