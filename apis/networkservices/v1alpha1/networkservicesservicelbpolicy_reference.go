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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkServicesServiceLBPolicyRef{}

// NetworkServicesServiceLBPolicyRef is a reference to a GCP NetworkServicesServiceLBPolicy.
type NetworkServicesServiceLBPolicyRef struct {
	// A reference to an externally managed NetworkServicesServiceLBPolicy resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/serviceLbPolicies/{{serviceLbPolicyID}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkServicesServiceLBPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkServicesServiceLBPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&NetworkServicesServiceLBPolicyRef{}, &NetworkServicesServiceLBPolicy{})
}

func (r *NetworkServicesServiceLBPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkServicesServiceLBPolicyGVK
}

func (r *NetworkServicesServiceLBPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesServiceLBPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesServiceLBPolicyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkServicesServiceLBPolicyRef) ValidateExternal(ref string) error {
	id := &NetworkServicesServiceLBPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkServicesServiceLBPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkServicesServiceLBPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkServicesServiceLBPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", r.GetGVK().Kind)
	}
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
