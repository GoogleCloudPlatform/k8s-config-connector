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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NotebooksEnvironmentIdentity{}
	_ identity.Resource   = &NotebooksEnvironment{}
)

var NotebooksEnvironmentIdentityFormat = gcpurls.Template[NotebooksEnvironmentIdentity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/environments/{environment}")

// NotebooksEnvironmentIdentity is the identity of a Google Cloud NotebooksEnvironment resource.
// +k8s:deepcopy-gen=false
type NotebooksEnvironmentIdentity struct {
	Project     string
	Location    string
	Environment string
}

func (i *NotebooksEnvironmentIdentity) String() string {
	return NotebooksEnvironmentIdentityFormat.ToString(*i)
}

func (i *NotebooksEnvironmentIdentity) FromExternal(ref string) error {
	parsed, match, err := NotebooksEnvironmentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NotebooksEnvironment external=%q was not known (use %s): %w", ref, NotebooksEnvironmentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NotebooksEnvironment external=%q was not known (use %s)", ref, NotebooksEnvironmentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NotebooksEnvironmentIdentity) Host() string {
	return NotebooksEnvironmentIdentityFormat.Host()
}

func (i *NotebooksEnvironmentIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func (i *NotebooksEnvironmentIdentity) ID() string {
	return i.Environment
}

// NewEnvironmentIdentity builds an NotebooksEnvironmentIdentity from the Config Connector NotebooksEnvironment object.
func NewEnvironmentIdentity(ctx context.Context, reader client.Reader, obj *NotebooksEnvironment) (*NotebooksEnvironmentIdentity, error) {
	// Get Parent
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
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
		actual := &NotebooksEnvironmentIdentity{}
		if err := actual.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actual.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actual.Project, projectID)
		}
		if actual.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actual.Location, location)
		}
		if actual.Environment != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actual.Environment)
		}
	}
	return &NotebooksEnvironmentIdentity{
		Project:     projectID,
		Location:    location,
		Environment: resourceID,
	}, nil
}

func getIdentityFromNotebooksEnvironmentSpec(ctx context.Context, reader client.Reader, obj *NotebooksEnvironment) (*NotebooksEnvironmentIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NotebooksEnvironmentIdentity{
		Project:     projectID,
		Location:    location,
		Environment: resourceID,
	}
	return identity, nil
}

func (obj *NotebooksEnvironment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNotebooksEnvironmentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &NotebooksEnvironmentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NotebooksEnvironment identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
