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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &PrivateCloudRef{}
var _ refsv1beta1.Ref = &PrivateCloudRef{}

// PrivateCloudRef is a reference to a VMwareEnginePrivateCloud.
type PrivateCloudRef struct {
	// A reference to an externally managed VMwareEnginePrivateCloud resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/privateClouds/{{privatecloudID}}".
	External string `json:"external,omitempty"`

	// The name of a VMwareEnginePrivateCloud resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VMwareEnginePrivateCloud resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *PrivateCloudRef) GetGVK() schema.GroupVersionKind {
	return VMwareEnginePrivateCloudGVK
}

func (r *PrivateCloudRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *PrivateCloudRef) GetExternal() string {
	return r.External
}

func (r *PrivateCloudRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *PrivateCloudRef) ValidateExternal(ref string) error {
	if _, _, err := ParsePrivateCloudExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *PrivateCloudRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj := &VMwareEnginePrivateCloud{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return ""
		}
		id, err := NewPrivateCloudIdentity(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value for other resource that depends on VMwareEnginePrivateCloud.
// If the "External" is given in the other resource's spec.VMwareEnginePrivateCloudRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual VMwareEnginePrivateCloud object from the cluster.
func (r *PrivateCloudRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", VMwareEnginePrivateCloudGVK.Kind)
	}
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
