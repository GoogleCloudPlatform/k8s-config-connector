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

package v1alpha2

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var VertexAIWorkbenchInstanceIdentityFormat = gcpurls.Template[VertexAIWorkbenchInstanceIdentity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// VertexAIWorkbenchInstanceIdentity defines the resource reference to VertexAIWorkbenchInstance, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type VertexAIWorkbenchInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *VertexAIWorkbenchInstanceIdentity) String() string {
	return VertexAIWorkbenchInstanceIdentityFormat.ToString(*i)
}

func (i *VertexAIWorkbenchInstanceIdentity) ID() string {
	return i.Instance
}

func (i *VertexAIWorkbenchInstanceIdentity) Parent() *VertexAIWorkbenchInstanceParent {
	return &VertexAIWorkbenchInstanceParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

func (i *VertexAIWorkbenchInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIWorkbenchInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Instance external=%q was not known (use %s): %w", ref, VertexAIWorkbenchInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Instance external=%q was not known (use %s)", ref, VertexAIWorkbenchInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// +k8s:deepcopy-gen=false
type VertexAIWorkbenchInstanceParent struct {
	ProjectID string
	Location  string
}

func (p *VertexAIWorkbenchInstanceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a VertexAIWorkbenchInstanceIdentity from the Config Connector Instance object.
func NewVertexAIWorkbenchInstanceIdentity(ctx context.Context, reader client.Reader, obj *VertexAIWorkbenchInstance) (*VertexAIWorkbenchInstanceIdentity, error) {

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
		actualParent, actualResourceID, err := ParseVertexAIWorkbenchInstanceExternal(externalRef)
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
	return &VertexAIWorkbenchInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}, nil
}

func ParseVertexAIWorkbenchInstanceExternal(external string) (parent *VertexAIWorkbenchInstanceParent, resourceID string, err error) {
	i := &VertexAIWorkbenchInstanceIdentity{}
	if err := i.FromExternal(external); err != nil {
		return nil, "", err
	}
	return i.Parent(), i.ID(), nil
}
