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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CCInsightsIssueModelIdentity{}
	_ identity.Resource   = &CCInsightsIssueModel{}
)

var CCInsightsIssueModelIdentityFormat = gcpurls.Template[CCInsightsIssueModelIdentity]("contactcenterinsights.googleapis.com", "projects/{project}/locations/{location}/issueModels/{issue_model}")

// +k8s:deepcopy-gen=false
type CCInsightsIssueModelIdentity struct {
	Project     string
	Location    string
	Issue_model string
}

func (i *CCInsightsIssueModelIdentity) String() string {
	return CCInsightsIssueModelIdentityFormat.ToString(*i)
}

func (i *CCInsightsIssueModelIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func (i *CCInsightsIssueModelIdentity) FromExternal(ref string) error {
	parsed, match, err := CCInsightsIssueModelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CCInsightsIssueModel external=%q was not known (use %s): %w", ref, CCInsightsIssueModelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CCInsightsIssueModel external=%q was not known (use %s)", ref, CCInsightsIssueModelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CCInsightsIssueModelIdentity) Host() string {
	return CCInsightsIssueModelIdentityFormat.Host()
}

func getIdentityFromCCInsightsIssueModelSpec(ctx context.Context, reader client.Reader, obj *CCInsightsIssueModel) (*CCInsightsIssueModelIdentity, error) {
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

	identity := &CCInsightsIssueModelIdentity{
		Project:     projectID,
		Location:    location,
		Issue_model: resourceID,
	}
	return identity, nil
}

func (obj *CCInsightsIssueModel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCCInsightsIssueModelSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CCInsightsIssueModelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if !isProjectMatch(statusIdentity.Project, specIdentity.Project) || statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change CCInsightsIssueModel identity parent (old=%q, new parent=%s/%s)", externalRef, specIdentity.Project, specIdentity.Location)
		}
		specIdentity.Project = statusIdentity.Project
		specIdentity.Issue_model = statusIdentity.Issue_model
	}

	return specIdentity, nil
}

type CCInsightsIssueModelParent struct {
	ProjectID string
	Location  string
}

func (p *CCInsightsIssueModelParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func ParseCCInsightsIssueModelExternal(external string) (parent *CCInsightsIssueModelParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "issueModels" {
		return nil, "", fmt.Errorf("format of CCInsightsIssueModel external=%q was not known (use projects/{{projectID}}/locations/{{location}}/issueModels/{{issueModelID}})", external)
	}
	parent = &CCInsightsIssueModelParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
