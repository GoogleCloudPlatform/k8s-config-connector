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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GKEHubNamespaceIdentity{}
	_ identity.Resource   = &GKEHubNamespace{}
)

var GKEHubNamespaceIdentityFormat = gcpurls.Template[GKEHubNamespaceIdentity](
	"gkehub.googleapis.com",
	"projects/{project}/locations/{location}/scopes/{scope}/namespaces/{namespace}",
)

// GKEHubNamespaceIdentity is the identity of a GKEHubNamespace.
// +k8s:deepcopy-gen=false
type GKEHubNamespaceIdentity struct {
	Project   string
	Location  string
	Scope     string
	Namespace string
}

func (i *GKEHubNamespaceIdentity) String() string {
	return GKEHubNamespaceIdentityFormat.ToString(*i)
}

func (i *GKEHubNamespaceIdentity) Host() string {
	return GKEHubNamespaceIdentityFormat.Host()
}

func (i *GKEHubNamespaceIdentity) FromExternal(ref string) error {
	parsed, match, err := GKEHubNamespaceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of GKEHubNamespace external=%q was not known (use %s): %w", ref, GKEHubNamespaceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of GKEHubNamespace external=%q was not known (use %s)", ref, GKEHubNamespaceIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func NewGKEHubNamespaceIdentity(project, location, scope, namespace string) *GKEHubNamespaceIdentity {
	return &GKEHubNamespaceIdentity{
		Project:   project,
		Location:  location,
		Scope:     scope,
		Namespace: namespace,
	}
}

func getIdentityFromGKEHubNamespaceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*GKEHubNamespaceIdentity, error) {
	gkeHubNamespace := &GKEHubNamespace{}
	switch t := obj.(type) {
	case *GKEHubNamespace:
		gkeHubNamespace = t
	case *unstructured.Unstructured:
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(t.Object, gkeHubNamespace); err != nil {
			return nil, fmt.Errorf("failed to convert unstructured to GKEHubNamespace: %w", err)
		}
	default:
		return nil, fmt.Errorf("expected *GKEHubNamespace or *unstructured.Unstructured, got %T", obj)
	}

	if gkeHubNamespace.Spec.ScopeRef == nil {
		return nil, fmt.Errorf("spec.scopeRef is required")
	}
	scopeRef, err := ResolveGKEHubScopeRef(ctx, reader, gkeHubNamespace, gkeHubNamespace.Spec.ScopeRef)
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

	namespaceID := direct.ValueOf(gkeHubNamespace.Spec.NamespaceID)
	if namespaceID == "" {
		namespaceID = direct.ValueOf(gkeHubNamespace.Spec.ResourceID)
	}
	if namespaceID == "" {
		namespaceID = gkeHubNamespace.GetName()
	}

	return NewGKEHubNamespaceIdentity(projectID, location, scopeID, namespaceID), nil
}

func (obj *GKEHubNamespace) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromGKEHubNamespaceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &GKEHubNamespaceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change GKEHubNamespace identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
