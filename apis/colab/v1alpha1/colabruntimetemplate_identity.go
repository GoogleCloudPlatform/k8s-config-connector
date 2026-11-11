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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NotebookRuntimeTemplateIdentity{}
	_ identity.Resource   = &ColabRuntimeTemplate{}
)

// +k8s:deepcopy-gen=false
type notebookRuntimeTemplateURL struct {
	ProjectID                 string
	Location                  string
	NotebookRuntimeTemplateID string
}

var NotebookRuntimeTemplateIdentityFormat = gcpurls.Template[notebookRuntimeTemplateURL](
	"aiplatform.googleapis.com",
	"projects/{projectID}/locations/{location}/notebookRuntimeTemplates/{notebookRuntimeTemplateID}",
)

// NotebookRuntimeTemplateIdentity is the identity of a ColabRuntimeTemplate.
type NotebookRuntimeTemplateIdentity struct {
	parent *NotebookRuntimeTemplateParent
	id     string
}

func (i *NotebookRuntimeTemplateIdentity) String() string {
	return NotebookRuntimeTemplateIdentityFormat.ToString(notebookRuntimeTemplateURL{
		ProjectID:                 i.parent.ProjectID,
		Location:                  i.parent.Location,
		NotebookRuntimeTemplateID: i.id,
	})
}

func (i *NotebookRuntimeTemplateIdentity) FromExternal(ref string) error {
	parsed, err := ParseNotebookRuntimeTemplateExternal(ref)
	if err != nil {
		return err
	}
	i.parent = parsed.parent
	i.id = parsed.id
	return nil
}

func (i *NotebookRuntimeTemplateIdentity) Host() string {
	return NotebookRuntimeTemplateIdentityFormat.Host()
}

func (i *NotebookRuntimeTemplateIdentity) ID() string {
	return i.id
}

func (i *NotebookRuntimeTemplateIdentity) Parent() *NotebookRuntimeTemplateParent {
	return i.parent
}

type NotebookRuntimeTemplateParent struct {
	ProjectID string
	Location  string
}

func (p *NotebookRuntimeTemplateParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// GetIdentity builds a NotebookRuntimeTemplateIdentity from the ColabRuntimeTemplate object.
func (obj *ColabRuntimeTemplate) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewNotebookRuntimeTemplateIdentity(ctx, reader, obj)
}

// New builds a NotebookRuntimeTemplateIdentity from the Config Connector NotebookRuntimeTemplate object.
func NewNotebookRuntimeTemplateIdentity(ctx context.Context, reader client.Reader, obj *ColabRuntimeTemplate) (*NotebookRuntimeTemplateIdentity, error) {

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

	if err := NotebookRuntimeTemplateIdentityFormat.Validate(notebookRuntimeTemplateURL{
		ProjectID:                 projectID,
		Location:                  location,
		NotebookRuntimeTemplateID: resourceID,
	}); err != nil {
		return nil, err
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualID, err := ParseNotebookRuntimeTemplateExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualID.Parent().ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualID.Parent().ProjectID, projectID)
		}
		if actualID.Parent().Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualID.Parent().Location, location)
		}
		if actualID.ID() != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualID.ID())
		}
	}
	return &NotebookRuntimeTemplateIdentity{
		parent: &NotebookRuntimeTemplateParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseNotebookRuntimeTemplateExternal(external string) (*NotebookRuntimeTemplateIdentity, error) {
	parsed, match, err := NotebookRuntimeTemplateIdentityFormat.Parse(external)
	if err != nil {
		return nil, fmt.Errorf("format of ColabRuntimeTemplate external=%q was not known (use projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}}): %w", external, err)
	}
	if !match {
		return nil, fmt.Errorf("format of ColabRuntimeTemplate external=%q was not known (use projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}})", external)
	}
	parent := &NotebookRuntimeTemplateParent{
		ProjectID: parsed.ProjectID,
		Location:  parsed.Location,
	}
	return &NotebookRuntimeTemplateIdentity{
		parent: parent,
		id:     parsed.NotebookRuntimeTemplateID,
	}, nil
}
