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

package networkconnectivityrefs

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var NetworkConnectivityServiceConnectionPolicyGVK = schema.GroupVersionKind{
	Group:   "networkconnectivity.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "NetworkConnectivityServiceConnectionPolicy",
}

var (
	_ refs.Ref         = &NetworkConnectivityServiceConnectionPolicyRef{}
	_ refs.ExternalRef = &NetworkConnectivityServiceConnectionPolicyRef{}
)

type NetworkConnectivityServiceConnectionPolicyRef struct {
	// A reference to an externally managed NetworkConnectivityServiceConnectionPolicy resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/serviceConnectionPolicies/{{serviceConnectionPolicyID}}".
	External string `json:"external,omitempty"`

	// The name field of a NetworkConnectivityServiceConnectionPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace field of a NetworkConnectivityServiceConnectionPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	kccscheme.RegisterRef(&NetworkConnectivityServiceConnectionPolicyRef{}, NetworkConnectivityServiceConnectionPolicyGVK)
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkConnectivityServiceConnectionPolicyGVK
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) ValidateExternal(ref string) error {
	id := &NetworkConnectivityServiceConnectionPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkConnectivityServiceConnectionPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkConnectivityServiceConnectionPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			return ""
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", projectID, location, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
