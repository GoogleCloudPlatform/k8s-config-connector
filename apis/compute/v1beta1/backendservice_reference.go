// Copyright 2024 Google LLC
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

var _ refsv1beta1.ExternalNormalizer = &ComputeBackendServiceRef{}
var ComputeBackendServiceGVK = GroupVersion.WithKind("ComputeBackendService")

// ComputeBackendServiceRef defines the resource reference to ComputeBackendService, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeBackendServiceRef struct {
	// For backward compatibility, we are not enforcing the external format.

	// The value of an externally managed ComputeBackendService resource.
	External string `json:"external,omitempty"`

	// The name of a ComputeBackendService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeBackendService resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeBackendService.
// If the "External" is given in the other resource's spec.ComputeBackendServiceRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeBackendService object from the cluster.
func (r *ComputeBackendServiceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeBackendServiceGVK.Kind)
	}
	// From given External
	// For backward compatibility, we are not validating the external format.
	// todo(yuhou): validate external when it's referenced by a pure direct resource
	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeBackendServiceGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeBackendServiceGVK, key, err)
	}

	// targetField: self_link
	// See compute servicemappings for details
	// todo(yuhou): use externalRef for resource that managed by direct controller
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", u.GetKind(), u.GetNamespace())
	}
	r.External = selfLink
	return r.External, nil
}
