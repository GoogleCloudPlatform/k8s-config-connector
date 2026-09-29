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
	_ identity.IdentityV2 = &CCInsightsAnalysisRuleIdentity{}
	_ identity.Resource   = &CCInsightsAnalysisRule{}
)

var CCInsightsAnalysisRuleIdentityFormat = gcpurls.Template[CCInsightsAnalysisRuleIdentity]("contactcenterinsights.googleapis.com", "projects/{project}/locations/{location}/analysisRules/{analysisRule}")

// CCInsightsAnalysisRuleIdentity is the identity of a GCP CCInsightsAnalysisRule resource.
// +k8s:deepcopy-gen=false
type CCInsightsAnalysisRuleIdentity struct {
	Project      string
	Location     string
	AnalysisRule string
}

func (i *CCInsightsAnalysisRuleIdentity) String() string {
	return CCInsightsAnalysisRuleIdentityFormat.ToString(*i)
}

func (i *CCInsightsAnalysisRuleIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *CCInsightsAnalysisRuleIdentity) FromExternal(ref string) error {
	parsed, match, err := CCInsightsAnalysisRuleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CCInsightsAnalysisRule external=%q was not known (use %s): %w", ref, CCInsightsAnalysisRuleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CCInsightsAnalysisRule external=%q was not known (use %s)", ref, CCInsightsAnalysisRuleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CCInsightsAnalysisRuleIdentity) Host() string {
	return CCInsightsAnalysisRuleIdentityFormat.Host()
}

func getIdentityFromCCInsightsAnalysisRuleSpec(ctx context.Context, reader client.Reader, obj *CCInsightsAnalysisRule) (*CCInsightsAnalysisRuleIdentity, error) {
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

	identity := &CCInsightsAnalysisRuleIdentity{
		Project:      projectID,
		Location:     location,
		AnalysisRule: resourceID,
	}
	return identity, nil
}

func (obj *CCInsightsAnalysisRule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCCInsightsAnalysisRuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CCInsightsAnalysisRuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if !isProjectMatch(statusIdentity.Project, specIdentity.Project) || statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change CCInsightsAnalysisRule identity parent (old=%q, new parent=%s/%s)", externalRef, specIdentity.Project, specIdentity.Location)
		}
		specIdentity.Project = statusIdentity.Project
		specIdentity.AnalysisRule = statusIdentity.AnalysisRule
	}

	return specIdentity, nil
}
