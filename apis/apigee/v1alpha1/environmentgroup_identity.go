// Copyright 2024 Google LLC
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
	"fmt"
	"strings"
)

// GoogleCloudApigeeV1EnvironmentGroupIdentity defines the resource reference to ApigeeEnvgroup, which "External" field
// holds the GCP identifier for the KRM object.
type GoogleCloudApigeeV1EnvironmentGroupIdentity struct {
	parent *GoogleCloudApigeeV1EnvironmentGroupParent
	id     string
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) String() string {
	return i.parent.String() + "/googlecloudapigeev1environmentgroups/" + i.id
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) ID() string {
	return i.id
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) Parent() *GoogleCloudApigeeV1EnvironmentGroupParent {
	return i.parent
}

type GoogleCloudApigeeV1EnvironmentGroupParent struct {
	ProjectID string
	Location  string
}

func (p *GoogleCloudApigeeV1EnvironmentGroupParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// // New builds a GoogleCloudApigeeV1EnvironmentGroupIdentity from the Config Connector GoogleCloudApigeeV1EnvironmentGroup object.
// func NewGoogleCloudApigeeV1EnvironmentGroupIdentity(ctx context.Context, reader client.Reader, obj *ApigeeEnvgroup) (*GoogleCloudApigeeV1EnvironmentGroupIdentity, error) {

// 	// Get Parent
// 	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
// 	if err != nil {
// 		return nil, err
// 	}
// 	projectID := projectRef.ProjectID
// 	if projectID == "" {
// 		return nil, fmt.Errorf("cannot resolve project")
// 	}
// 	location := obj.Spec.Location

// 	// Get desired ID
// 	resourceID := common.ValueOf(obj.Spec.ResourceID)
// 	if resourceID == "" {
// 		resourceID = obj.GetName()
// 	}
// 	if resourceID == "" {
// 		return nil, fmt.Errorf("cannot resolve resource ID")
// 	}

// 	// Use approved External
// 	externalRef := common.ValueOf(obj.Status.ExternalRef)
// 	if externalRef != "" {
// 		// Validate desired with actual
// 		actualParent, actualResourceID, err := ParseGoogleCloudApigeeV1EnvironmentGroupExternal(externalRef)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if actualParent.ProjectID != projectID {
// 			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
// 		}
// 		if actualParent.Location != location {
// 			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
// 		}
// 		if actualResourceID != resourceID {
// 			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
// 				resourceID, actualResourceID)
// 		}
// 	}
// 	return &GoogleCloudApigeeV1EnvironmentGroupIdentity{
// 		parent: &GoogleCloudApigeeV1EnvironmentGroupParent{
// 			ProjectID: projectID,
// 			Location:  location,
// 		},
// 		id: resourceID,
// 	}, nil
// }

func ParseGoogleCloudApigeeV1EnvironmentGroupExternal(external string) (parent *GoogleCloudApigeeV1EnvironmentGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "googlecloudapigeev1environmentgroups" {
		return nil, "", fmt.Errorf("format of ApigeeEnvgroup external=%q was not known (use projects/<projectId>/locations/<location>/googlecloudapigeev1environmentgroups/<googlecloudapigeev1environmentgroupID>)", external)
	}
	parent = &GoogleCloudApigeeV1EnvironmentGroupParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
