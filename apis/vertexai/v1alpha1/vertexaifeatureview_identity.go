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
	_ identity.IdentityV2 = &VertexAIFeatureViewIdentity{}
	_ identity.Resource   = &VertexAIFeatureView{}
)

var VertexAIFeatureViewIdentityFormat = gcpurls.Template[VertexAIFeatureViewIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/featureOnlineStores/{featureOnlineStore}/featureViews/{featureView}")

// VertexAIFeatureViewIdentity is the identity of a GCP VertexAIFeatureView resource.
// +k8s:deepcopy-gen=false
type VertexAIFeatureViewIdentity struct {
	Project            string
	Location           string
	FeatureOnlineStore string
	FeatureView        string
}

func (i *VertexAIFeatureViewIdentity) String() string {
	return VertexAIFeatureViewIdentityFormat.ToString(*i)
}

func (i *VertexAIFeatureViewIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIFeatureViewIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIFeatureView external=%q was not known (use %s): %w", ref, VertexAIFeatureViewIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeatureView external=%q was not known (use %s)", ref, VertexAIFeatureViewIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIFeatureViewIdentity) Host() string {
	return VertexAIFeatureViewIdentityFormat.Host()
}

func getIdentityFromVertexAIFeatureViewSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIFeatureViewIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAIFeatureView)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	featureOnlineStoreRef := vertexaiObj.Spec.FeatureOnlineStoreRef
	if featureOnlineStoreRef == nil {
		return nil, fmt.Errorf("spec.featureOnlineStoreRef must be specified")
	}
	if err := featureOnlineStoreRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.featureOnlineStoreRef: %w", err)
	}
	featureOnlineStoreIdentityRaw, err := featureOnlineStoreRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing featureOnlineStoreRef: %w", err)
	}
	featureOnlineStoreIdentity, ok := featureOnlineStoreIdentityRaw.(*VertexAIFeatureOnlineStoreIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAIFeatureOnlineStoreIdentity from featureOnlineStoreRef")
	}
	parentFeatureOnlineStore := featureOnlineStoreIdentity.FeatureOnlineStore
	identity := &VertexAIFeatureViewIdentity{
		Project:            projectID,
		Location:           location,
		FeatureOnlineStore: parentFeatureOnlineStore,
		FeatureView:        resourceID,
	}
	return identity, nil
}

func (obj *VertexAIFeatureView) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIFeatureViewSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIFeatureViewIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIFeatureView identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
