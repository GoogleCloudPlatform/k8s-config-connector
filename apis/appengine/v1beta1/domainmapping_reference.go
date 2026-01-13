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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &DomainMappingRef{}

// DomainMappingRef is a reference to an AppEngineDomainMapping resource.
type DomainMappingRef struct {
	// A reference to an externally managed AppEngineDomainMapping resource.
	// Should be in the format "apps/{{projectID}}/domainMappings/{{domainMappingID}}".
	External *string `json:"external,omitempty"`

	// The name of an AppEngineDomainMapping resource.
	Name *string `json:"name,omitempty"`

	// The namespace of an AppEngineDomainMapping resource.
	Namespace *string `json:"namespace,omitempty"`
}

func (r *DomainMappingRef) GetGVK() schema.GroupVersionKind {
	return AppEngineDomainMappingGVK
}

func (r *DomainMappingRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      direct.ValueOf(r.Name),
		Namespace: direct.ValueOf(r.Namespace),
	}
}

func (r *DomainMappingRef) GetExternal() string {
	return direct.ValueOf(r.External)
}

func (r *DomainMappingRef) SetExternal(ref string) {
	r.External = direct.LazyPtr(ref)
}

func (r *DomainMappingRef) ValidateExternal(ref string) error {
	_, _, err := ParseDomainMappingExternal(ref)
	if err != nil {
		return err
	}
	return nil
}

func (r *DomainMappingRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Try to get the external ref from status
		external, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
		return external
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
