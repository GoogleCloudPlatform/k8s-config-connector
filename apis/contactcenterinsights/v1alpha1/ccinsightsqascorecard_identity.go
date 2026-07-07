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
	_ identity.IdentityV2 = &CCInsightsQAScorecardIdentity{}
	_ identity.Resource   = &CCInsightsQAScorecard{}
)

var CCInsightsQAScorecardIdentityFormat = gcpurls.Template[CCInsightsQAScorecardIdentity]("contactcenterinsights.googleapis.com", "projects/{project}/locations/{location}/qaScorecards/{qaScorecard}")

// CCInsightsQAScorecardIdentity is the identity of a GCP CCInsightsQAScorecard resource.
// +k8s:deepcopy-gen=false
type CCInsightsQAScorecardIdentity struct {
	Project     string
	Location    string
	QaScorecard string
}

func (i *CCInsightsQAScorecardIdentity) String() string {
	return CCInsightsQAScorecardIdentityFormat.ToString(*i)
}

func (i *CCInsightsQAScorecardIdentity) FromExternal(ref string) error {
	parsed, match, err := CCInsightsQAScorecardIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CCInsightsQAScorecard external=%q was not known (use %s): %w", ref, CCInsightsQAScorecardIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CCInsightsQAScorecard external=%q was not known (use %s)", ref, CCInsightsQAScorecardIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CCInsightsQAScorecardIdentity) Host() string {
	return CCInsightsQAScorecardIdentityFormat.Host()
}

func (i *CCInsightsQAScorecardIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromCCInsightsQAScorecardSpec(ctx context.Context, reader client.Reader, obj *CCInsightsQAScorecard) (*CCInsightsQAScorecardIdentity, error) {
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

	identity := &CCInsightsQAScorecardIdentity{
		Project:     projectID,
		Location:    location,
		QaScorecard: resourceID,
	}
	return identity, nil
}

func (obj *CCInsightsQAScorecard) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCCInsightsQAScorecardSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CCInsightsQAScorecardIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CCInsightsQAScorecard identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
