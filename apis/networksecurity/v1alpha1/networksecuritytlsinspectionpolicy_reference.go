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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	refs.Register(&NetworkSecurityTLSInspectionPolicyRef{})
}

var _ refs.Ref = &NetworkSecurityTLSInspectionPolicyRef{}

// +k8s:deepcopy-gen=true
type NetworkSecurityTLSInspectionPolicyRef struct {
	// A reference to an externally managed NetworkSecurityTLSInspectionPolicy resource.
	// Should be in the format "projects/{projectID}/locations/{location}/tlsInspectionPolicies/{tlsinspectionpolicy}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityTLSInspectionPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityTLSInspectionPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityTLSInspectionPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityTLSInspectionPolicyGVK
}

func (r *NetworkSecurityTLSInspectionPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *NetworkSecurityTLSInspectionPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityTLSInspectionPolicyRef) SetExternal(external string) {
	r.External = external
}

func (r *NetworkSecurityTLSInspectionPolicyRef) ValidateExternal(external string) error {
	id := &NetworkSecurityTLSInspectionPolicyIdentity{}
	return id.FromExternal(external)
}

func (r *NetworkSecurityTLSInspectionPolicyRef) ParseExternalToIdentity() (any, error) {
	id := &NetworkSecurityTLSInspectionPolicyIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

func (r *NetworkSecurityTLSInspectionPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromNetworkSecurityTLSInspectionPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
