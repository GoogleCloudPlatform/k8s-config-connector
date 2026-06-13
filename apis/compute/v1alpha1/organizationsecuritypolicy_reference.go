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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeOrganizationSecurityPolicyGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "ComputeOrganizationSecurityPolicy",
}

var _ refsv1beta1.Ref = &ComputeOrganizationSecurityPolicyRef{}

var ComputeOrganizationSecurityPolicyIdentityFormat = gcpurls.Template[ComputeOrganizationSecurityPolicyIdentity]("compute.googleapis.com", "organizations/{organization}/locations/global/securityPolicies/{name}")

// +k8s:deepcopy-gen=false
type ComputeOrganizationSecurityPolicyIdentity struct {
	Organization string
	Name         string
}

func (i *ComputeOrganizationSecurityPolicyIdentity) String() string {
	return ComputeOrganizationSecurityPolicyIdentityFormat.ToString(*i)
}

func (i *ComputeOrganizationSecurityPolicyIdentity) FromExternal(ref string) error {
	ref = refsv1beta1.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeOrganizationSecurityPolicyIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeOrganizationSecurityPolicy external=%q was not known (use %s)", ref, ComputeOrganizationSecurityPolicyIdentityFormat.CanonicalForm())
}

func (i *ComputeOrganizationSecurityPolicyIdentity) Host() string {
	return ComputeOrganizationSecurityPolicyIdentityFormat.Host()
}

// ComputeOrganizationSecurityPolicyRef is a reference to a ComputeOrganizationSecurityPolicy.
type ComputeOrganizationSecurityPolicyRef struct {
	// A reference to an externally managed ComputeOrganizationSecurityPolicy resource.
	// Should be in the format "organizations/{{organization}}/locations/global/securityPolicies/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeOrganizationSecurityPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeOrganizationSecurityPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&ComputeOrganizationSecurityPolicyRef{})
}

func (r *ComputeOrganizationSecurityPolicyRef) GetGVK() schema.GroupVersionKind {
	return ComputeOrganizationSecurityPolicyGVK
}

func (r *ComputeOrganizationSecurityPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeOrganizationSecurityPolicyRef) GetExternal() string {
	return r.External
}

func (r *ComputeOrganizationSecurityPolicyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeOrganizationSecurityPolicyRef) ValidateExternal(ref string) error {
	id := &ComputeOrganizationSecurityPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeOrganizationSecurityPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeOrganizationSecurityPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeOrganizationSecurityPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, _ := refsv1beta1.GetResourceID(u)
		if resourceID != "" {
			return "locations/global/securityPolicies/" + resourceID
		}
		return ""
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeOrganizationSecurityPolicy.
func (r *ComputeOrganizationSecurityPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
