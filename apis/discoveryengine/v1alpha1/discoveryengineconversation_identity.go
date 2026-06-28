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
	_ identity.IdentityV2 = &DiscoveryEngineConversationIdentity{}
	_ identity.Resource   = &DiscoveryEngineConversation{}
)

var DiscoveryEngineConversationIdentityFormat = gcpurls.Template[DiscoveryEngineConversationIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/dataStores/{dataStore}/conversations/{conversation}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineConversationIdentity struct {
	Project      string
	Location     string
	DataStore    string
	Conversation string
}

func (i *DiscoveryEngineConversationIdentity) String() string {
	return DiscoveryEngineConversationIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineConversationIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineConversationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineConversation external=%q was not known (use %s): %w", ref, DiscoveryEngineConversationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineConversation external=%q was not known (use %s)", ref, DiscoveryEngineConversationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineConversationIdentity) Host() string {
	return DiscoveryEngineConversationIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineConversationSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineConversation) (*DiscoveryEngineConversationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if obj.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("spec.dataStoreRef is required")
	}

	dataStoreRef := *obj.Spec.DataStoreRef
	normalizedExternal, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}

	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(normalizedExternal)
	if err != nil {
		return nil, fmt.Errorf("parsing dataStoreRef.external=%q: %w", normalizedExternal, err)
	}

	identity := &DiscoveryEngineConversationIdentity{
		Project:      projectID,
		Location:     location,
		DataStore:    dataStoreLink.DataStore,
		Conversation: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineConversation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineConversationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineConversationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineConversation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *DiscoveryEngineConversation) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
