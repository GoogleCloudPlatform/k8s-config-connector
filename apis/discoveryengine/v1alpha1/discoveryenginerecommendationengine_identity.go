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
	_ identity.IdentityV2 = &DiscoveryEngineRecommendationEngineIdentity{}
	_ identity.Resource   = &DiscoveryEngineRecommendationEngine{}
)

var DiscoveryEngineRecommendationEngineIdentityFormat = gcpurls.Template[DiscoveryEngineRecommendationEngineIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/engines/{recommendationengine}")

// DiscoveryEngineRecommendationEngineIdentity is the identity of a GCP DiscoveryEngineRecommendationEngine resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineRecommendationEngineIdentity struct {
	Project              string
	Location             string
	Collection           string
	RecommendationEngine string
}

func (i *DiscoveryEngineRecommendationEngineIdentity) String() string {
	return DiscoveryEngineRecommendationEngineIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineRecommendationEngineIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineRecommendationEngineIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineRecommendationEngine external=%q was not known (use %s): %w", ref, DiscoveryEngineRecommendationEngineIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineRecommendationEngine external=%q was not known (use %s)", ref, DiscoveryEngineRecommendationEngineIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineRecommendationEngineIdentity) Host() string {
	return DiscoveryEngineRecommendationEngineIdentityFormat.Host()
}

func (i *DiscoveryEngineRecommendationEngineIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/collections/" + i.Collection
}

func getIdentityFromDiscoveryEngineRecommendationEngineSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineRecommendationEngine) (*DiscoveryEngineRecommendationEngineIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	collection := obj.Spec.Collection
	if collection == "" {
		return nil, fmt.Errorf("cannot resolve collection")
	}

	identity := &DiscoveryEngineRecommendationEngineIdentity{
		Project:              projectID,
		Location:             location,
		Collection:           collection,
		RecommendationEngine: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineRecommendationEngine) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineRecommendationEngineSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &DiscoveryEngineRecommendationEngineIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineRecommendationEngine identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
