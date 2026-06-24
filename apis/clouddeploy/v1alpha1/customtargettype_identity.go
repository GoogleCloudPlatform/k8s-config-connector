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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudDeployCustomTargetTypeIdentity{}
	_ identity.Resource   = &CloudDeployCustomTargetType{}
)

var CloudDeployCustomTargetTypeIdentityFormat = gcpurls.Template[CloudDeployCustomTargetTypeIdentity]("clouddeploy.googleapis.com", "projects/{project}/locations/{location}/customTargetTypes/{customTargetType}")

// +k8s:deepcopy-gen=false
type CloudDeployCustomTargetTypeIdentity struct {
	Project          string
	Location         string
	CustomTargetType string
}

func (i *CloudDeployCustomTargetTypeIdentity) String() string {
	return CloudDeployCustomTargetTypeIdentityFormat.ToString(*i)
}

func (i *CloudDeployCustomTargetTypeIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudDeployCustomTargetTypeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudDeployCustomTargetType external=%q was not known (use %s): %w", ref, CloudDeployCustomTargetTypeIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudDeployCustomTargetType external=%q was not known (use %s)", ref, CloudDeployCustomTargetTypeIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudDeployCustomTargetTypeIdentity) Host() string {
	return CloudDeployCustomTargetTypeIdentityFormat.Host()
}

// CustomTargetTypeIdentity is the identity of a CloudDeployCustomTargetType.
type CustomTargetTypeIdentity struct {
	parent *CustomTargetTypeParent
	id     string
}

func (i *CustomTargetTypeIdentity) String() string {
	return i.parent.String() + "/customTargetTypes/" + i.id
}

func (i *CustomTargetTypeIdentity) ID() string {
	return i.id
}

func (i *CustomTargetTypeIdentity) Parent() *CustomTargetTypeParent {
	return i.parent
}

// CustomTargetTypeParent defines the parent project/location
type CustomTargetTypeParent struct {
	ProjectID string
	Location  string
}

func (p *CustomTargetTypeParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func ParseCustomTargetTypeExternal(external string) (parent *CustomTargetTypeParent, resourceID string, err error) {
	i := &CloudDeployCustomTargetTypeIdentity{}
	if err := i.FromExternal(external); err != nil {
		return nil, "", err
	}
	return &CustomTargetTypeParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}, i.CustomTargetType, nil
}

func NewCustomTargetTypeIdentity(ctx context.Context, reader client.Reader, obj *CloudDeployCustomTargetType) (*CustomTargetTypeIdentity, error) {
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualParent, actualResourceID, err := ParseCustomTargetTypeExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &CustomTargetTypeIdentity{
		parent: &CustomTargetTypeParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func getIdentityFromCloudDeployCustomTargetTypeSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudDeployCustomTargetTypeIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &CloudDeployCustomTargetTypeIdentity{
		Project:          projectID,
		Location:         location,
		CustomTargetType: resourceID,
	}
	return identity, nil
}

func (obj *CloudDeployCustomTargetType) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudDeployCustomTargetTypeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CloudDeployCustomTargetTypeIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudDeployCustomTargetType identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
