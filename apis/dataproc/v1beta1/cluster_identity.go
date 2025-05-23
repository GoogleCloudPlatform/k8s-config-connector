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

// ClusterIdentity defines the resource reference to DataprocCluster, which "External" field
// holds the GCP identifier for the KRM object.
type ClusterIdentity struct {
	parent *ClusterParent
	id     string
}

func (i *ClusterIdentity) String() string {
	return i.parent.String() + "/clusters/" + i.id
}

func (i *ClusterIdentity) ID() string {
	return i.id
}

func (i *ClusterIdentity) Parent() *ClusterParent {
	return i.parent
}

type ClusterParent struct {
	ProjectID string
	Region    string
}

func (p *ClusterParent) String() string {
	return "projects/" + p.ProjectID + "/regions/" + p.Region
}

func ParseClusterExternal(external string) (id *ClusterIdentity, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "clusters" {
		return nil, fmt.Errorf("format of DataprocCluster external=%q was not known (use projects/<projectID>/regions/<region>/clusters/<clusterName>)", external)
	}

	return &ClusterIdentity{parent: &ClusterParent{ProjectID: tokens[1], Region: tokens[3]}, id: tokens[5]}, nil
}
