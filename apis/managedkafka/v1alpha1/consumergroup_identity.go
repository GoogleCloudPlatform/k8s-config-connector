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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ConsumerGroupIdentity defines the resource reference to ManagedKafkaConsumerGroup, which "External" field
// holds the GCP identifier for the KRM object.
type ConsumerGroupIdentity struct {
	parent *ConsumerGroupParent
	id     string
}

func (i *ConsumerGroupIdentity) String() string {
	return i.parent.String() + "/consumerGroups/" + i.id
}

func (i *ConsumerGroupIdentity) ID() string {
	return i.id
}

func (i *ConsumerGroupIdentity) Parent() *ConsumerGroupParent {
	return i.parent
}

type ConsumerGroupParent struct {
	ProjectID   string
	Location    string
	ClusterName string
}

func (p *ConsumerGroupParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/clusters/" + p.ClusterName
}

// New builds a ConsumerGroupIdentity from the Config Connector ConsumerGroup object.
func NewConsumerGroupIdentity(ctx context.Context, reader client.Reader, obj *ManagedKafkaConsumerGroup) (*ConsumerGroupIdentity, error) {

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
	clusterName := obj.Spec.ClusterRef.Name

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
		actualParent, actualResourceID, err := ParseConsumerGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.ClusterName != clusterName {
			return nil, fmt.Errorf("spec.clusterRef.name changed, expect %s, got %s", actualParent.ClusterName, clusterName)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ConsumerGroupIdentity{
		parent: &ConsumerGroupParent{
			ProjectID:   projectID,
			Location:    location,
			ClusterName: clusterName,
		},
		id: resourceID,
	}, nil
}

func ParseConsumerGroupExternal(external string) (parent *ConsumerGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "clusters" || tokens[6] != "consumerGroups" {
		return nil, "", fmt.Errorf("format of ManagedKafkaConsumerGroup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/clusters/{{cluster}}/consumerGroups/{{consumergroupID}})", external)
	}
	parent = &ConsumerGroupParent{
		ProjectID:   tokens[1],
		Location:    tokens[3],
		ClusterName: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
