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
	_ identity.IdentityV2 = &DiscoveryEngineChatEngineIdentity{}
	_ identity.Resource   = &DiscoveryEngineChatEngine{}
)

var DiscoveryEngineChatEngineIdentityFormat = gcpurls.Template[DiscoveryEngineChatEngineIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/engines/{chatEngine}")

// DiscoveryEngineChatEngineIdentity is the identity of a GCP DiscoveryEngineChatEngine resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineChatEngineIdentity struct {
	Project    string
	Location   string
	Collection string
	ChatEngine string
}

func (i *DiscoveryEngineChatEngineIdentity) String() string {
	return DiscoveryEngineChatEngineIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineChatEngineIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineChatEngineIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineChatEngine external=%q was not known (use %s): %w", ref, DiscoveryEngineChatEngineIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineChatEngine external=%q was not known (use %s)", ref, DiscoveryEngineChatEngineIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineChatEngineIdentity) Host() string {
	return DiscoveryEngineChatEngineIdentityFormat.Host()
}

func (i *DiscoveryEngineChatEngineIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/collections/" + i.Collection
}

func getIdentityFromDiscoveryEngineChatEngineSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineChatEngine) (*DiscoveryEngineChatEngineIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	collection := obj.Spec.Collection
	if collection == "" {
		return nil, fmt.Errorf("cannot resolve collection")
	}

	identity := &DiscoveryEngineChatEngineIdentity{
		Project:    projectID,
		Location:   location,
		Collection: collection,
		ChatEngine: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineChatEngine) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineChatEngineSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &DiscoveryEngineChatEngineIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineChatEngine identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
