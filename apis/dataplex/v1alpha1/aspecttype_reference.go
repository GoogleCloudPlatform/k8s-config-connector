// Copyright 2025 Google LLC
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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &AspectTypeRef{}
var DataplexAspectTypeGVK = GroupVersion.WithKind("DataplexAspectType")

// AspectTypeRef defines the resource reference to DataplexAspectTypeRef, which "External" field
// holds the GCP identifier for the KRM object.
type AspectTypeRef struct {
	// A reference to an externally managed DataplexAspectType resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/aspectTypes/{{aspecttypeID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a DataplexAspectType resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataplexAspectType resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *AspectTypeRef) GetGVK() schema.GroupVersionKind {
	return DataplexAspectTypeGVK
}

func (r *AspectTypeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *AspectTypeRef) GetExternal() string {
	return r.External
}

func (r *AspectTypeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AspectTypeRef) ValidateExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("external reference is empty")
	}
	// Format: projects/{{projectID}}/locations/{{location}}/aspectTypes/{{aspecttypeID}}
	tokens := strings.Split(ref, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "aspectTypes" {
		return fmt.Errorf("format of DataplexAspectType external=%q was not known (use projects/{{projectID}}/locations/{{location}}/aspectTypes/{{aspecttypeID}})", ref)
	}
	return nil
}

func (r *AspectTypeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", DataplexAspectTypeGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
