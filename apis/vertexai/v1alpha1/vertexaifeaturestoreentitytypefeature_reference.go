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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAIFeaturestoreEntityTypeFeatureRef{}

// VertexAIFeaturestoreEntityTypeFeatureRef is a reference to a GCP VertexAIFeaturestoreEntityTypeFeature.
type VertexAIFeaturestoreEntityTypeFeatureRef struct {
	// A reference to an externally managed VertexAIFeaturestoreEntityTypeFeature resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/featurestores/{{featurestore}}/entityTypes/{{entityType}}/features/{{feature}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIFeaturestoreEntityTypeFeature resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIFeaturestoreEntityTypeFeature resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIFeaturestoreEntityTypeFeatureRef{}, &VertexAIFeaturestoreEntityTypeFeature{})
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) GetGVK() schema.GroupVersionKind {
	return VertexAIFeaturestoreEntityTypeFeatureGVK
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) GetExternal() string {
	return r.External
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) ValidateExternal(ref string) error {
	id := &VertexAIFeaturestoreEntityTypeFeatureIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIFeaturestoreEntityTypeFeatureIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIFeaturestoreEntityTypeFeatureRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		ready := false
		conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
		if err == nil && found {
			for _, condVal := range conditions {
				cond, ok := condVal.(map[string]interface{})
				if !ok {
					continue
				}
				if cond["type"] == "Ready" && cond["status"] == "True" {
					ready = true
					break
				}
			}
		}
		if !ready {
			return ""
		}
		obj, err := common.ToStructuredType[*VertexAIFeaturestoreEntityTypeFeature](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromVertexAIFeaturestoreEntityTypeFeatureSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
