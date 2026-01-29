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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigateway/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// APIGatewayAPIRef supports using an APIGatewayAPI resource to specify the API.
type APIGatewayAPIRef struct {
	/* The name of the APIGatewayAPI resource. */
	// +optional
	Name string `json:"name,omitempty"`

	/* The namespace of the APIGatewayAPI resource. */
	// +optional
	Namespace string `json:"namespace,omitempty"`

	/* The external name of the APIGatewayAPI resource. */
	// +optional
	External string `json:"external,omitempty"`
}

func (r *APIGatewayAPIRef) GroupVersionKind() schema.GroupVersionKind {
	return v1alpha1.APIGatewayAPIGVK
}

func (r *APIGatewayAPIRef) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	// No-op
}

func (r *APIGatewayAPIRef) ExternalName() string {
	return r.External
}

func (r *APIGatewayAPIRef) SetExternalName(name string) {
	r.External = name
}

func (r *APIGatewayAPIRef) Reference() *refs.ResourceRef {
	if r == nil {
		return nil
	}
	return &refs.ResourceRef{
		Name:      r.Name,
		Namespace: r.Namespace,
		External:  r.External,
	}
}
