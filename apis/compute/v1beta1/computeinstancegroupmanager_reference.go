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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeInstanceGroupManagerRef{}
var ComputeInstanceGroupManagerGVK = GroupVersion.WithKind("ComputeInstanceGroupManager")

// ComputeInstanceGroupManagerRef is a reference to a ComputeInstanceGroupManager.
type ComputeInstanceGroupManagerRef struct {
	// A reference to an externally managed ComputeInstanceGroupManager resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/instanceGroupManagers/{{instanceGroupManagerID}}" or "projects/{{projectID}}/zones/{{zone}}/instanceGroupManagers/{{instanceGroupManagerID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeInstanceGroupManager resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInstanceGroupManager resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provisions the "External" value for resources that depend on ComputeInstanceGroupManager.
func (r *ComputeInstanceGroupManagerRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeInstanceGroupManagerGVK.Kind)
	}
	// From given External
	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeInstanceGroupManagerGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeInstanceGroupManagerGVK, key, err)
	}

	// Try reading externalRef first, then fallback to selfLink
	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		r.External = externalRef
		return r.External, nil
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink or externalRef for referenced %s %s (status is empty)", u.GetKind(), key)
	}
	r.External = selfLink
	return r.External, nil
}
