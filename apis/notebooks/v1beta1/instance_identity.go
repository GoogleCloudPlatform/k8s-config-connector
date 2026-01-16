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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var InstanceIdentityFormat = gcpurls.Template[InstanceIdentity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// InstanceIdentity defines the resource reference to NotebookInstance, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *InstanceIdentity) String() string {
	return InstanceIdentityFormat.ToString(*i)
}

func (i *InstanceIdentity) ID() string {
	return i.Instance
}

func (i *InstanceIdentity) Parent() *InstanceParent {
	return &InstanceParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

func (i *InstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := InstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Instance external=%q was not known (use %s): %w", ref, InstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Instance external=%q was not known (use %s)", ref, InstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

type InstanceParent struct {
	ProjectID string
	Location  string
}

func (p *InstanceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a InstanceIdentity from the Config Connector Instance object.
func NewInstanceIdentity(ctx context.Context, reader client.Reader, obj *NotebookInstance) (*InstanceIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Zone

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
		actualParent, actualResourceID, err := ParseInstanceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.zone changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}, nil
}

func ParseInstanceExternal(external string) (parent *InstanceParent, resourceID string, err error) {
	i := &InstanceIdentity{}
	if err := i.FromExternal(external); err != nil {
		return nil, "", err
	}
	return i.Parent(), i.ID(), nil
}
