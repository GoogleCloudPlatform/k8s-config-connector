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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DeployPolicyRef{}

func init() {
	refs.Register(&DeployPolicyRef{})
}

// DeployPolicyRef is a reference to a DeployDeployPolicy.
type DeployPolicyRef struct {
	// A reference to an externally managed CloudDeployDeployPolicy resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/deployPolicies/{{deployPolicy}}".
	External string `json:"external,omitempty"`

	// The name of a CloudDeployDeployPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudDeployDeployPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *DeployPolicyRef) GetGVK() schema.GroupVersionKind {
	return DeployDeployPolicyGVK
}

func (r *DeployPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DeployPolicyRef) GetExternal() string {
	return r.External
}

func (r *DeployPolicyRef) SetExternal(external string) {
	r.External = external
}

func (r *DeployPolicyRef) ValidateExternal(ref string) error {
	i := &DeployPolicyIdentity{}
	return i.FromExternal(ref)
}

func (r *DeployPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	i := &DeployPolicyIdentity{}
	err := i.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (r *DeployPolicyRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		i, err := getIdentityFromDeployPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return i.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
