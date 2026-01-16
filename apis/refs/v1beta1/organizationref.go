// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	deprecatedrefs "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
)

// OrganizationRef represents the Organization that this resource belongs to.
type OrganizationRef struct {
	// The 'name' field of an organization, when not managed by Config Connector.
	// +required
	External string `json:"external,omitempty"`
}

// AsOrganizationRef converts a generic ResourceRef into a OrganizationRef.
func AsOrganizationRef(in *deprecatedrefs.ResourceRef) *OrganizationRef {
	if in == nil {
		return nil
	}
	return &OrganizationRef{
		External: in.External,
	}
}

type OrganizationIdentity struct {
	OrganizationID string
}

// Organization is an alias for OrganizationIdentity
// Deprecated: Use OrganizationIdentity instead.
type Organization = OrganizationIdentity

var _ identity.Identity = &OrganizationIdentity{}

var OrganizationFormat = gcpurls.Template[OrganizationIdentity]("cloudresourcemanager.googleapis.com", "organizations/{organizationID}")

func (i *OrganizationIdentity) String() string {
	return OrganizationFormat.ToString(*i)
}

func (i *OrganizationIdentity) FromExternal(ref string) error {
	parsed, match, err := OrganizationFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Organization external=%q was not known (use %s): %w", ref, OrganizationFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Organization external=%q was not known (use %s)", ref, OrganizationFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// ResolveOrganizationFromAnnotation resolves the OrganizationID to use for a
// resource, it should be used for resources which do not have
// 'spec.organizationRef'.
func ResolveOrganizationFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (*OrganizationIdentity, error) {
	if organizationID := src.GetAnnotations()["cnrm.cloud.google.com/organization-id"]; organizationID != "" {
		return &OrganizationIdentity{OrganizationID: organizationID}, nil
	}

	return nil, fmt.Errorf("organization-id annotation not set on resource")
}

// ResolveOrganization will resolve an OrganizationRef to an Organization, with
// the OrganizationID.
func ResolveOrganization(ctx context.Context, reader client.Reader, src client.Object, ref *OrganizationRef) (*OrganizationIdentity, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External == "" {
		return nil, fmt.Errorf("must specify 'external' in 'organizationRef'")
	}

	// We support "organizations/12345"
	tokens := strings.Split(ref.External, "/")
	if len(tokens) == 2 && tokens[0] == "organizations" {
		return &OrganizationIdentity{OrganizationID: tokens[1]}, nil
	}
	return nil, fmt.Errorf("format of 'organizationRef.external'=%q was not known (use organizations/<organizationID>)", ref.External)
}

func ResolveOrganizationID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	organizationRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "organizationRef", "external")
	if organizationRefExternal != "" {
		organizationRef := OrganizationRef{
			External: organizationRefExternal,
		}

		organization, err := ResolveOrganization(ctx, reader, obj, &organizationRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse organizationRef.external %q in %v %v/%v: %w", organizationRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return organization.OrganizationID, nil
	}

	if organizationID := obj.GetAnnotations()["cnrm.cloud.google.com/organization-id"]; organizationID != "" {
		return organizationID, nil
	}

	return "", fmt.Errorf("cannot find organization id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
