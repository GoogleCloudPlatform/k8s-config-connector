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
	_ identity.IdentityV2 = &CCInsightsViewIdentity{}
	_ identity.Resource   = &CCInsightsView{}
)

var CCInsightsViewIdentityFormat = gcpurls.Template[CCInsightsViewIdentity]("contactcenterinsights.googleapis.com", "projects/{project}/locations/{location}/views/{view}")

// CCInsightsViewIdentity is the identity of a GCP CCInsightsView resource.
// +k8s:deepcopy-gen=false
type CCInsightsViewIdentity struct {
	Project  string
	Location string
	View     string
}

func (i *CCInsightsViewIdentity) String() string {
	return CCInsightsViewIdentityFormat.ToString(*i)
}

func (i *CCInsightsViewIdentity) FromExternal(ref string) error {
	parsed, match, err := CCInsightsViewIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CCInsightsView external=%q was not known (use %s): %w", ref, CCInsightsViewIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CCInsightsView external=%q was not known (use %s)", ref, CCInsightsViewIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CCInsightsViewIdentity) Host() string {
	return CCInsightsViewIdentityFormat.Host()
}

func getIdentityFromCCInsightsViewSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CCInsightsViewIdentity, error) {
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

	identity := &CCInsightsViewIdentity{
		Project:  projectID,
		Location: location,
		View:     resourceID,
	}
	return identity, nil
}

func (obj *CCInsightsView) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCCInsightsViewSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CCInsightsViewIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CCInsightsView identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
