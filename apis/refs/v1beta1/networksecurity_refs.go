// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NetworkSecurityGatewaySecurityPolicyRef is a reference to a NetworkSecurityGatewaySecurityPolicy.
type NetworkSecurityGatewaySecurityPolicyRef struct {
	// Name of the referenced object.
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace of the referenced object.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// The GatewaySecurityPolicy selfLink, when not managed by Config Connector.
	// +optional
	External string `json:"external,omitempty"`
}

// Normalize resolves the reference and sets the External field.
func (r *NetworkSecurityGatewaySecurityPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name == "" {
		return nil
	}
	if r.Name == "" {
		return fmt.Errorf("must specify either name or external")
	}

	ns := r.Namespace
	if ns == "" {
		ns = defaultNamespace
	}

	key := types.NamespacedName{
		Namespace: ns,
		Name:      r.Name,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "networksecurity.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "NetworkSecurityGatewaySecurityPolicy",
	})
	if err := reader.Get(ctx, key, u); err != nil {
		return fmt.Errorf("reading NetworkSecurityGatewaySecurityPolicy %s: %w", key, err)
	}

	// Try to get externalRef from status
	val, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("getting status.externalRef: %w", err)
	}
	if found && val != "" {
		r.External = val
		return nil
	}

	return fmt.Errorf("NetworkSecurityGatewaySecurityPolicy %s is not ready, externalRef is not set", key)
}
