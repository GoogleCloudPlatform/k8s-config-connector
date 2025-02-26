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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TopicIdentity defines the resource reference to ManagedKafkaTopic, which "External" field
// holds the GCP identifier for the KRM object.
type TopicIdentity struct {
	parent *TopicParent
	id     string
}

func (i *TopicIdentity) String() string {
	return i.parent.String() + "/topics/" + i.id
}

func (i *TopicIdentity) ID() string {
	return i.id
}

func (i *TopicIdentity) Parent() *TopicParent {
	return i.parent
}

type TopicParent struct {
	ProjectID string
	Location  string
	Cluster   string
}

func (p *TopicParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/clusters/" + p.Cluster
}

// New builds a TopicIdentity from the Config Connector Topic object.
func NewTopicIdentity(ctx context.Context, reader client.Reader, obj *ManagedKafkaTopic) (*TopicIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location
	clusterExternalRef, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	_, cluster, err := ParseClusterExternal(clusterExternalRef)
	if err != nil {
		return nil, err
	}
	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTopicExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.Cluster != cluster {
			return nil, fmt.Errorf("spec.cluster changed, expect %s, got %s", actualParent.Cluster, cluster)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &TopicIdentity{
		parent: &TopicParent{
			ProjectID: projectID,
			Location:  location,
			Cluster:   cluster,
		},
		id: resourceID,
	}, nil
}

func ParseTopicExternal(external string) (parent *TopicParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "clusters" || tokens[6] != "topics" {
		return nil, "", fmt.Errorf("format of ManagedKafkaTopic external=%q was not known (use projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}}/topics/{{topicID}})", external)
	}
	parent = &TopicParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Cluster:   tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
