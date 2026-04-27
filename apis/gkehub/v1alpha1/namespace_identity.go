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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.Identity   = &GKEHubNamespaceIdentity{}
	_ identity.IdentityV2 = &GKEHubNamespaceIdentity{}
	_ identity.Resource   = &GKEHubNamespace{}

	namespaceURL = gcpurls.Template[GKEHubNamespaceIdentity](
		"gkehub.googleapis.com",
		"projects/{projectID}/locations/{location}/scopes/{scopeID}/namespaces/{namespaceID}",
	)
)

// GKEHubNamespaceIdentity defines the resource reference to GKEHubNamespace, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubNamespaceIdentity struct {
	ProjectID   string
	Location    string
	ScopeID     string
	NamespaceID string
}

func (i *GKEHubNamespaceIdentity) String() string {
	return namespaceURL.ToString(*i)
}

func (i *GKEHubNamespaceIdentity) ID() string {
	return i.NamespaceID
}

func (i *GKEHubNamespaceIdentity) Host() string {
	return namespaceURL.Host()
}

func (i *GKEHubNamespaceIdentity) Parent() *GKEHubScopeIdentity {
	return NewGKEHubScopeIdentity(i.ProjectID, i.Location, i.ScopeID)
}

func (i *GKEHubNamespaceIdentity) FromExternal(external string) error {
	out, match, err := namespaceURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubNamespace external=%q was not known (use %s)", external, namespaceURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubNamespaceIdentity(project, location, scopeID, namespaceID string) *GKEHubNamespaceIdentity {
	return &GKEHubNamespaceIdentity{
		ProjectID:   project,
		Location:    location,
		ScopeID:     scopeID,
		NamespaceID: namespaceID,
	}
}

func (obj *GKEHubNamespace) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	if obj.Spec.ScopeRef == nil {
		return nil, fmt.Errorf("spec.scopeRef is required")
	}
	scopeRef, err := ResolveGKEHubScopeRef(ctx, reader, obj, obj.Spec.ScopeRef)
	if err != nil {
		return nil, err
	}
	projectID := scopeRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("could not derive projectID from scopeRef")
	}
	location := scopeRef.Location
	if location == "" {
		return nil, fmt.Errorf("could not derive location from scopeRef")
	}
	scopeID := scopeRef.ID()
	if scopeID == "" {
		return nil, fmt.Errorf("could not derive scopeID from scopeRef")
	}

	namespaceID := direct.ValueOf(obj.Spec.NamespaceID)
	if namespaceID == "" {
		return nil, fmt.Errorf("spec.namespaceID is required")
	}

	return NewGKEHubNamespaceIdentity(projectID, location, scopeID, namespaceID), nil
}
