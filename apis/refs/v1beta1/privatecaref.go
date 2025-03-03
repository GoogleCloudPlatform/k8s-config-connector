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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PrivateCACAPoolRef struct {
	// A reference to an externally managed PrivateCACAPool.
	// Should be in the format `projects/{project_id}/locations/{region}/caPools/{caPool}`.
	External string `json:"external,omitempty"`

	// The `name` of a `PrivateCACAPool` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `PrivateCACAPool` resource.
	Namespace string `json:"namespace,omitempty"`
}

type PrivateCACAPool struct {
	Ref        *PrivateCACAPoolRef
	ResourceID string
}

// ResolvePrivateCACAPoolRef will resolve a PrivateCACAPoolRef to a PrivateCACAPool.
func ResolvePrivateCACAPoolRef(ctx context.Context, reader client.Reader, src client.Object, ref *PrivateCACAPoolRef) (*PrivateCACAPoolRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on PrivateCACAPoolRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on PrivateCACAPoolRef")
	}

	// External should be in the `projects/{project_id}/locations/{region}/caPools/{caPool}` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" {
			ref = &PrivateCACAPoolRef{
				External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of PrivateCACAPoolRef external=%q was not known (use projects/{project_id}/locations/{region}/caPools/{caPool})", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	caPool := &unstructured.Unstructured{}
	caPool.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "privateca.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "PrivateCACAPool",
	})
	if err := reader.Get(ctx, key, caPool); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced PrivateCACAPool %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced PrivateCACAPool %v: %w", key, err)
	}
	resource, err := k8s.NewResource(caPool)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(caPool.GroupVersionKind(), key)
	}

	caPoolResourceID, err := GetResourceID(caPool)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, caPool)
	if err != nil {
		return nil, err
	}

	location, err := GetLocation(caPool)
	if err != nil {
		return nil, err
	}

	ref = &PrivateCACAPoolRef{
		External: fmt.Sprintf("projects/%s/locations/%s/caPools/%s", projectID, location, caPoolResourceID),
	}

	return ref, nil
}
