// Copyright 2025 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &OrganizationRef{}
var OrganizationGVK = GroupVersion.WithKind("Organization")

// OrganizationRef represents the Organization that this resource belongs to.
type OrganizationRef struct {
	// A reference to an externally managed Organization resource.
	// Should be in the format "organizations/{{organizationID}}" or "{{organizationID}}".
	External string `json:"external,omitempty"`
}

func (r *OrganizationRef) GetGVK() schema.GroupVersionKind {
	return OrganizationGVK
}

func (r *OrganizationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *OrganizationRef) GetExternal() string {
	return r.External
}

func (r *OrganizationRef) SetExternal(external string) {
	r.External = external
}

func (r *OrganizationRef) ValidateExternal(external string) error {
	if _, err := ParseOrganizationExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *OrganizationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("must specify 'external' in 'organizationRef'")
	}
	organizationIdentity, err := ParseOrganizationExternal(r.External)
	if err != nil {
		return err
	}
	r.SetExternal(organizationIdentity.String())
	return nil
}

// ResolveOrganizationFromAnnotation resolves the OrganizationID for resources do not have 'spec.organizationRef'.
// todo: Identify the resources or use cases where this function is necessary.
func ResolveOrganizationFromAnnotation(ctx context.Context, reader client.Reader, src client.Object) (string, error) {
	if organizationID := src.GetAnnotations()["cnrm.cloud.google.com/organization-id"]; organizationID != "" {
		return organizationID, nil
	}
	return "", fmt.Errorf("organization-id annotation not set on resource")
}
