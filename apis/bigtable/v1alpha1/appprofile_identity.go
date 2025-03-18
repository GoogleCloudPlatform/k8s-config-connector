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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AppProfileIdentity defines the resource reference to BigtableAppProfile, which "External" field
// holds the GCP identifier for the KRM object.
type AppProfileIdentity struct {
	parent *AppProfileParent
	id     string
}

func (i *AppProfileIdentity) String() string {
	return i.parent.String() + "/appProfiles/" + i.id
}

func (i *AppProfileIdentity) ID() string {
	return i.id
}

func (i *AppProfileIdentity) Parent() *AppProfileParent {
	return i.parent
}

type AppProfileParent struct {
	ProjectID   string
	InstanceRef string
}

func (p *AppProfileParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.InstanceRef
}

// New builds a AppProfileIdentity from the Config Connector AppProfile object.
// No changes were needed as the existing logic already resolves the only field in AppProfileParent (InstanceRef).
func NewAppProfileIdentity(ctx context.Context, reader client.Reader, obj *BigtableAppProfile) (*AppProfileIdentity, error) {
	// Get Parent
	instanceRef, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}

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
		actualParent, actualResourceID, err := ParseAppProfileExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.InstanceRef != instanceRef {
			return nil, fmt.Errorf("spec.instanceRef changed, expect %s, got %s", actualParent.InstanceRef, instanceRef)
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRefRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &AppProfileIdentity{
		parent: &AppProfileParent{
			ProjectID:   projectID,
			InstanceRef: instanceRef,
		},
		id: resourceID,
	}, nil
}

func ParseAppProfileExternal(external string) (parent *AppProfileParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "appProfiles" {
		return nil, "", fmt.Errorf("format of BigtableAppProfile external=%q was not known (use projects/{{projectID}}/instances/{{instance}}/appProfiles/{{appProfileId}})", external)
	}
	parent = &AppProfileParent{
		ProjectID:   tokens[1],
		InstanceRef: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
