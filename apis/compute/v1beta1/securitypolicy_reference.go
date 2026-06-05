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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeSecurityPolicyGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeSecurityPolicy",
}

var _ refsv1beta1.Ref = &ComputeSecurityPolicyRef{}
var _ refsv1beta1.ExternalNormalizer = &ComputeSecurityPolicyRef{}

// ComputeSecurityPolicyRef is a reference to a ComputeSecurityPolicy.
type ComputeSecurityPolicyRef struct {
	// A reference to an externally managed ComputeSecurityPolicy resource.
	// Should be in the format "projects/{{projectID}}/global/securityPolicies/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeSecurityPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeSecurityPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&ComputeSecurityPolicyRef{})
}

func (r *ComputeSecurityPolicyRef) GetGVK() schema.GroupVersionKind {
	return ComputeSecurityPolicyGVK
}

func (r *ComputeSecurityPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeSecurityPolicyRef) GetExternal() string {
	return r.External
}

func (r *ComputeSecurityPolicyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeSecurityPolicyRef) ValidateExternal(ref string) error {
	id := &ComputeSecurityPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeSecurityPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeSecurityPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeSecurityPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromComputeSecurityPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeSecurityPolicy.
// If the "External" is given in the other resource's spec.ComputeSecurityPolicyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeSecurityPolicy object from the cluster.
func (r *ComputeSecurityPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
