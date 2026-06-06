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

package v1beta1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CustomClassRef{}

// CustomClassRef is a reference to a SpeechCustomClass.
type CustomClassRef struct {
	// A reference to an externally managed SpeechCustomClass resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/customClasses/{{customclassID}}".
	External string `json:"external,omitempty"`

	// The name of a SpeechCustomClass resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SpeechCustomClass resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CustomClassRef{})
}

func (r *CustomClassRef) GetGVK() schema.GroupVersionKind {
	return SpeechCustomClassGVK
}

func (r *CustomClassRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CustomClassRef) GetExternal() string {
	return r.External
}

func (r *CustomClassRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *CustomClassRef) ValidateExternal(ref string) error {
	id := &CustomClassIdentity{}
	return id.FromExternal(ref)
}

func (r *CustomClassRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CustomClassIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CustomClassRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromSpeechCustomClassSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
