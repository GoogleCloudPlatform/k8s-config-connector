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

package common

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type ResourceReference struct {
	// Kind of the referenced resource
	Kind      string `json:"kind"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	// APIVersion of the referenced resource
	APIVersion string `json:"apiVersion,omitempty"`
	// The external name of the referenced resource
	External string `json:"external,omitempty"`
}

func (r *ResourceReference) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ResourceReference) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(r.APIVersion, r.Kind)
}

func (r *ResourceReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	r.APIVersion, r.Kind = gvk.ToAPIVersionAndKind()
}
