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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var NetworkSecurityInterceptDeploymentGroupGVK = schema.GroupVersionKind{
	Group:   "networksecurity.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "NetworkSecurityInterceptDeploymentGroup",
}

type NetworkSecurityInterceptDeploymentGroupRef struct {
	/* A reference to an externally managed NetworkSecurityInterceptDeploymentGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptDeploymentGroups/{{interceptDeploymentGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityInterceptDeploymentGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityInterceptDeploymentGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

type NetworkSecurityMirroringDeploymentGroupRef struct {
	/* A reference to an externally managed NetworkSecurityMirroringDeploymentGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroringDeploymentGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityMirroringDeploymentGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityMirroringDeploymentGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityInterceptDeploymentGroupRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", NetworkSecurityInterceptDeploymentGroupGVK.Kind)
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Name == "" {
		return "", fmt.Errorf("either name or external must be specified on %s reference", NetworkSecurityInterceptDeploymentGroupGVK.Kind)
	}

	nn := types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	if nn.Namespace == "" {
		nn.Namespace = otherNamespace
	}

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(NetworkSecurityInterceptDeploymentGroupGVK)
	if err := reader.Get(ctx, nn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("%s %v not found", NetworkSecurityInterceptDeploymentGroupGVK.Kind, nn)
		}
		return "", fmt.Errorf("error reading %s %v: %w", NetworkSecurityInterceptDeploymentGroupGVK.Kind, nn, err)
	}

	externalRef, _, _ := unstructured.NestedString(obj.Object, "status", "externalRef")
	if externalRef == "" {
		return "", fmt.Errorf("%s %v is not ready, or externalRef is not set", NetworkSecurityInterceptDeploymentGroupGVK.Kind, nn)
	}
	return externalRef, nil
}

var NetworkSecurityInterceptDeploymentGVK = schema.GroupVersionKind{
	Group:   "networksecurity.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "NetworkSecurityInterceptDeployment",
}

type NetworkSecurityInterceptDeploymentRef struct {
	/* A reference to an externally managed NetworkSecurityInterceptDeployment resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptDeployments/{{interceptDeploymentID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityInterceptDeployment resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityInterceptDeployment resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityInterceptDeploymentRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", NetworkSecurityInterceptDeploymentGVK.Kind)
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Name == "" {
		return "", fmt.Errorf("either name or external must be specified on %s reference", NetworkSecurityInterceptDeploymentGVK.Kind)
	}

	nn := types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	if nn.Namespace == "" {
		nn.Namespace = otherNamespace
	}

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(NetworkSecurityInterceptDeploymentGVK)
	if err := reader.Get(ctx, nn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("%s %v not found", NetworkSecurityInterceptDeploymentGVK.Kind, nn)
		}
		return "", fmt.Errorf("error reading %s %v: %w", NetworkSecurityInterceptDeploymentGVK.Kind, nn, err)
	}

	externalRef, _, _ := unstructured.NestedString(obj.Object, "status", "externalRef")
	if externalRef == "" {
		return "", fmt.Errorf("%s %v is not ready, or externalRef is not set", NetworkSecurityInterceptDeploymentGVK.Kind, nn)
	}
	return externalRef, nil
}

var NetworkSecurityInterceptEndpointGroupGVK = schema.GroupVersionKind{
	Group:   "networksecurity.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "NetworkSecurityInterceptEndpointGroup",
}

type NetworkSecurityInterceptEndpointGroupRef struct {
	/* A reference to an externally managed NetworkSecurityInterceptEndpointGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptEndpointGroups/{{interceptEndpointGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityInterceptEndpointGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityInterceptEndpointGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityInterceptEndpointGroupRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", NetworkSecurityInterceptEndpointGroupGVK.Kind)
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Name == "" {
		return "", fmt.Errorf("either name or external must be specified on %s reference", NetworkSecurityInterceptEndpointGroupGVK.Kind)
	}

	nn := types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	if nn.Namespace == "" {
		nn.Namespace = otherNamespace
	}

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(NetworkSecurityInterceptEndpointGroupGVK)
	if err := reader.Get(ctx, nn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("%s %v not found", NetworkSecurityInterceptEndpointGroupGVK.Kind, nn)
		}
		return "", fmt.Errorf("error reading %s %v: %w", NetworkSecurityInterceptEndpointGroupGVK.Kind, nn, err)
	}

	externalRef, _, _ := unstructured.NestedString(obj.Object, "status", "externalRef")
	if externalRef == "" {
		return "", fmt.Errorf("%s %v is not ready, or externalRef is not set", NetworkSecurityInterceptEndpointGroupGVK.Kind, nn)
	}
	return externalRef, nil
}
