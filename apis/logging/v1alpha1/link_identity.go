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

// LinkIdentity defines the resource reference to LoggingLink, which "External" field
// holds the GCP identifier for the KRM object.
type LinkIdentity struct {
	parent *LinkParent
	id     string
}

func (i *LinkIdentity) String() string {
	return i.parent.String() + "/links/" + i.id
}

func (i *LinkIdentity) ID() string {
	return i.id
}

func (i *LinkIdentity) Parent() *LinkParent {
	return i.parent
}

type LinkParent struct {
	ProjectID          string
	Location           string
	LoggingLogBucketID string
}

func (p *LinkParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/buckets/" + p.LoggingLogBucketID
}

// New builds a LinkIdentity from the Config Connector Link object.
func NewLinkIdentity(ctx context.Context, reader client.Reader, obj *LoggingLink) (*LinkIdentity, error) {

	// Get Parent
	bucketRef, err := refsv1beta1.ResolveLoggingLogBucketRef(ctx, reader, obj, obj.Spec.LoggingLogBucketRef)
	if err != nil {
		return nil, err
	}
	projectID := bucketRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := bucketRef.Location
	bucketID := bucketRef.LoggingLogBucketID

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
		actualParent, actualResourceID, err := ParseLinkExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != bucketRef.String() {
			return nil, fmt.Errorf("actualParent changed, expect %s, got %s", actualParent.String(), bucketRef.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s", resourceID, actualResourceID)
		}

	}
	return &LinkIdentity{
		parent: &LinkParent{
			ProjectID:          projectID,
			Location:           location,
			LoggingLogBucketID: bucketID,
		},
		id: resourceID,
	}, nil
}

func ParseLinkExternal(external string) (parent *LinkParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "buckets" || tokens[6] != "links" {
		return nil, "", fmt.Errorf("format of LoggingLink external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}}/links/{{linkID}})", external)
	}
	parent = &LinkParent{
		ProjectID:          tokens[1],
		Location:           tokens[3],
		LoggingLogBucketID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
