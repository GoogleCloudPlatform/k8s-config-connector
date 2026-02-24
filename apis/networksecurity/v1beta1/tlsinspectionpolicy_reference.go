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

package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NetworkSecurityTlsInspectionPolicyRef struct {
	/* The `name` of a `NetworkSecurityTlsInspectionPolicy` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` of a `NetworkSecurityTlsInspectionPolicy` resource. */
	Namespace string `json:"namespace,omitempty"`
	/* The NetworkSecurityTlsInspectionPolicy selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
}

func ResolveNetworkSecurityTlsInspectionPolicy(ctx context.Context, reader client.Reader, obj client.Object, ref *NetworkSecurityTlsInspectionPolicyRef) (string, error) {
	if ref == nil {
		return "", nil
	}
	if ref.External != "" {
		if ref.Name != "" {
			return "", fmt.Errorf("cannot specify both name and external on %s reference", "NetworkSecurityTlsInspectionPolicy")
		}
		return ref.External, nil
	}
	if ref.Name == "" {
		return "", fmt.Errorf("must specify either name or external on %s reference", "NetworkSecurityTlsInspectionPolicy")
	}

	key := types.NamespacedName{
		Name:      ref.Name,
		Namespace: ref.Namespace,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	endpoint := &unstructured.Unstructured{}
	endpoint.SetGroupVersionKind(GroupVersion.WithKind("NetworkSecurityTlsInspectionPolicy"))
	if err := reader.Get(ctx, key, endpoint); err != nil {
		return "", fmt.Errorf("getting %s %s: %w", "NetworkSecurityTlsInspectionPolicy", key, err)
	}

	externalRef, _, err := unstructured.NestedString(endpoint.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("getting status.externalRef: %w", err)
	}
	if externalRef == "" {
		return "", fmt.Errorf("%s %s has not yet been reconciled (status.externalRef is empty)", "NetworkSecurityTlsInspectionPolicy", key)
	}
	return externalRef, nil
}
