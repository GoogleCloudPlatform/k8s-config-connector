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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &FeatureOnlineStoreIdentity{}
	_ identity.Resource   = &VertexAIFeatureOnlineStore{}
)

var FeatureOnlineStoreIdentityFormat = gcpurls.Template[FeatureOnlineStoreIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/featureOnlineStores/{featureonlinestore}")

// +k8s:deepcopy-gen=false
type FeatureOnlineStoreIdentity struct {
	Project            string
	Location           string
	FeatureOnlineStore string
}

func (i *FeatureOnlineStoreIdentity) String() string {
	return FeatureOnlineStoreIdentityFormat.ToString(*i)
}

func (i *FeatureOnlineStoreIdentity) FromExternal(ref string) error {
	parsed, match, err := FeatureOnlineStoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIFeatureOnlineStore external=%q was not known (use %s): %w", ref, FeatureOnlineStoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeatureOnlineStore external=%q was not known (use %s)", ref, FeatureOnlineStoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FeatureOnlineStoreIdentity) Host() string {
	return FeatureOnlineStoreIdentityFormat.Host()
}

func getIdentityFromFeatureOnlineStoreSpec(ctx context.Context, reader client.Reader, obj *VertexAIFeatureOnlineStore) (*FeatureOnlineStoreIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &FeatureOnlineStoreIdentity{
		Project:            projectID,
		Location:           location,
		FeatureOnlineStore: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIFeatureOnlineStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromFeatureOnlineStoreSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &FeatureOnlineStoreIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIFeatureOnlineStore identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (i *FeatureOnlineStoreIdentity) Parent() *FeatureOnlineStoreParent {
	return &FeatureOnlineStoreParent{
		Project:  i.Project,
		Location: i.Location,
	}
}

type FeatureOnlineStoreParent struct {
	Project  string
	Location string
}

func (p *FeatureOnlineStoreParent) String() string {
	return "projects/" + p.Project + "/locations/" + p.Location
}
