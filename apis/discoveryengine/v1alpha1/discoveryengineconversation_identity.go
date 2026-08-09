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
	"unicode"

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

var DiscoveryEngineConversationIdentityFormat = gcpurls.Template[DiscoveryEngineConversationIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}/conversations/{conversation}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineConversationIdentity struct {
	Project      string
	Location     string
	Collection   string
	DataStore    string
	Conversation string
}

func (i *DiscoveryEngineConversationIdentity) String() string {
	return DiscoveryEngineConversationIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineConversationIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", i.Project, i.Location, i.Collection, i.DataStore)
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

	// Validation checks: parent's project/location should match conversation's project/location
	if !IsProjectIDMatch(dataStoreLink.ProjectID, projectID) {
		return nil, fmt.Errorf("resolved spec.dataStoreRef project %q does not match spec.projectRef %q", dataStoreLink.ProjectID, projectID)
	}
	if dataStoreLink.Location != location {
		return nil, fmt.Errorf("resolved spec.dataStoreRef location %q does not match spec.location %q", dataStoreLink.Location, location)
	}

	identity := &DiscoveryEngineConversationIdentity{
		Project:      projectID,
		Location:     location,
		Collection:   dataStoreLink.Collection,
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

		if statusIdentity.Location != specIdentity.Location || statusIdentity.Collection != specIdentity.Collection || statusIdentity.DataStore != specIdentity.DataStore {
			return nil, fmt.Errorf("cannot change DiscoveryEngineConversation parent identity (old=%q, new parent=%s/%s/%s/%s)", statusIdentity.String(), specIdentity.Project, specIdentity.Location, specIdentity.Collection, specIdentity.DataStore)
		}

		specIdentity.Project = statusIdentity.Project
		specIdentity.Conversation = statusIdentity.Conversation
	}

	return specIdentity, nil
}

func (obj *DiscoveryEngineConversation) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}

// IsProjectIDMatch returns true if two project identifiers are considered matching,
// or if we can't reliably compare them because one is an alphanumeric project ID
// and the other is a numeric project number.
func IsProjectIDMatch(p1, p2 string) bool {
	if p1 == p2 {
		return true
	}
	if p1 == "" || p2 == "" {
		return false
	}
	p1IsNumeric := isNumeric(p1)
	p2IsNumeric := isNumeric(p2)
	if p1IsNumeric != p2IsNumeric {
		// Skip strict matching if one is project ID (alphanumeric) and the other is project number (numeric).
		return true
	}
	return p1 == p2
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
