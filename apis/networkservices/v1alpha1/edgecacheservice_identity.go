// Copyright 2026 Google LLC
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

// EdgeCacheServiceIdentity defines the resource reference to NetworkServicesEdgeCacheService, which "External" field
// holds the GCP identifier for the KRM object.
type EdgeCacheServiceIdentity struct {
	parent *EdgeCacheServiceParent
	id     string
}

func (i *EdgeCacheServiceIdentity) String() string {
	return i.parent.String() + "/edgeCacheServices/" + i.id
}

func (i *EdgeCacheServiceIdentity) ID() string {
	return i.id
}

func (i *EdgeCacheServiceIdentity) Parent() *EdgeCacheServiceParent {
	return i.parent
}

type EdgeCacheServiceParent struct {
	ProjectID string
}

func (p *EdgeCacheServiceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/global"
}

// NewEdgeCacheServiceIdentity builds a EdgeCacheServiceIdentity from the Config Connector EdgeCacheService object.
func NewEdgeCacheServiceIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesEdgeCacheService) (*EdgeCacheServiceIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
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
		actualParent, actualResourceID, err := ParseEdgeCacheServiceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &EdgeCacheServiceIdentity{
		parent: &EdgeCacheServiceParent{
			ProjectID: projectID,
		},
		id: resourceID,
	}, nil
}

func ParseEdgeCacheServiceExternal(external string) (parent *EdgeCacheServiceParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "edgeCacheServices" {
		return nil, "", fmt.Errorf("format of NetworkServicesEdgeCacheService external=%q was not known (use projects/{{projectID}}/locations/global/edgeCacheServices/{{edgeCacheServiceID}})", external)
	}
	if tokens[3] != "global" {
		return nil, "", fmt.Errorf("NetworkServicesEdgeCacheService external=%q has invalid location %q (must be global)", external, tokens[3])
	}
	parent = &EdgeCacheServiceParent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
