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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeBackendServiceGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeBackendService",
}

var _ refs.Ref = &ComputeBackendServiceRef{}
var _ refs.ExternalNormalizer = &ComputeBackendServiceRef{}

// ComputeBackendServiceRef is a reference to a ComputeBackendService.
type ComputeBackendServiceRef struct {
	// A reference to an externally managed ComputeBackendService resource.
	// Should be in the format "projects/{{projectID}}/global/backendServices/{{backendservice}}" or "projects/{{projectID}}/regions/{{location}}/backendServices/{{backendservice}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeBackendService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeBackendService resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeBackendServiceRef{}, &ComputeBackendService{})
}

func (r *ComputeBackendServiceRef) GetGVK() schema.GroupVersionKind {
	return ComputeBackendServiceGVK
}

func (r *ComputeBackendServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeBackendServiceRef) GetExternal() string {
	return r.External
}

func (r *ComputeBackendServiceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeBackendServiceRef) ValidateExternal(ref string) error {
	id := &ComputeBackendServiceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeBackendServiceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeBackendServiceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeBackendServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ComputeBackendService](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeBackendServiceSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	if err := refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback); err != nil {
		return err
	}
	if r.External != "" && !strings.HasPrefix(r.External, "https://") && strings.HasPrefix(r.External, "projects/") {
		r.External = "https://www.googleapis.com/compute/v1/" + r.External
	}
	return nil
}

// NormalizedExternal is a deprecated method that assigns and returns the "External" field.
// Deprecated: Use Normalize instead.
func (r *ComputeBackendServiceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
