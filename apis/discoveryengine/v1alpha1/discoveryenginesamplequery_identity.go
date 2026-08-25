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
	_ identity.IdentityV2 = &DiscoveryEngineSampleQueryIdentity{}
	_ identity.Resource   = &DiscoveryEngineSampleQuery{}
)

var DiscoveryEngineSampleQueryIdentityFormat = gcpurls.Template[DiscoveryEngineSampleQueryIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/sampleQuerySets/{sampleQuerySet}/sampleQueries/{sampleQuery}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineSampleQueryIdentity struct {
	Project        string
	Location       string
	SampleQuerySet string
	SampleQuery    string
}

func (i *DiscoveryEngineSampleQueryIdentity) String() string {
	return DiscoveryEngineSampleQueryIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineSampleQueryIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineSampleQueryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineSampleQuery external=%q was not known (use %s): %w", ref, DiscoveryEngineSampleQueryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineSampleQuery external=%q was not known (use %s)", ref, DiscoveryEngineSampleQueryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineSampleQueryIdentity) Host() string {
	return DiscoveryEngineSampleQueryIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineSampleQuerySpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveryEngineSampleQueryIdentity, error) {
	sampleQuery, ok := obj.(*DiscoveryEngineSampleQuery)
	if !ok {
		return nil, fmt.Errorf("object is not a DiscoveryEngineSampleQuery")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := common.ValueOf(sampleQuery.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("spec.location is not set")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if sampleQuery.Spec.SampleQuerySetRef == nil {
		return nil, fmt.Errorf("spec.sampleQuerySetRef is not set")
	}

	sampleQuerySetRef := *sampleQuery.Spec.SampleQuerySetRef
	normalizedSampleQuerySet, err := sampleQuerySetRef.NormalizedExternal(ctx, reader, sampleQuery.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.sampleQuerySetRef: %w", err)
	}

	sampleQuerySetParent, sampleQuerySetID, err := ParseSampleQuerySetExternal(normalizedSampleQuerySet)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.sampleQuerySetRef external: %w", err)
	}

	// Validation checks: parent's project/location should match sampleQuery's project/location
	if sampleQuerySetParent.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.sampleQuerySetRef project %q does not match spec.projectRef %q", sampleQuerySetParent.ProjectID, projectID)
	}
	if sampleQuerySetParent.Location != location {
		return nil, fmt.Errorf("resolved spec.sampleQuerySetRef location %q does not match spec.location %q", sampleQuerySetParent.Location, location)
	}

	identity := &DiscoveryEngineSampleQueryIdentity{
		Project:        projectID,
		Location:       location,
		SampleQuerySet: sampleQuerySetID,
		SampleQuery:    resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineSampleQuery) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineSampleQuerySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineSampleQueryIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineSampleQuery identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DiscoveryEngineSampleQuery) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
