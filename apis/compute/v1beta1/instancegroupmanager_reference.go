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
	"strings"

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
	// Should be in the format "projects/{{projectID}}/zones/{{zone}}/instanceGroupManagers/{{instanceGroupManagerID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeInstanceGroupManager resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInstanceGroupManager resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeInstanceGroupManager.
func (r *ComputeInstanceGroupManagerRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeInstanceGroupManagerGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, _, err := ParseInstanceGroupManagerExternal(r.External); err != nil {
			return "", err
		}
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

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", u.GetKind(), u.GetNamespace())
	}
	identity, err := ParseInstanceGroupManagerSelfLink(selfLink)
	if err != nil {
		return "", fmt.Errorf("failed to parse selfLink: %w", err)
	}
	r.External = identity
	return r.External, nil
}

func ParseInstanceGroupManagerExternal(external string) (projectID, zone, resourceID string, err error) {
	// e.g. projects/my-project/zones/us-central1-a/instanceGroupManagers/my-mig
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "zones" || tokens[4] != "instanceGroupManagers" {
		return "", "", "", fmt.Errorf("format of ComputeInstanceGroupManager external=%q was not known (use projects/{{projectID}}/zones/{{zone}}/instanceGroupManagers/{{instanceGroupManagerID}})", external)
	}
	return tokens[1], tokens[3], tokens[5], nil
}

func ParseInstanceGroupManagerSelfLink(selfLink string) (string, error) {
	// Remove the service prefix if present
	path := selfLink
	if strings.HasPrefix(selfLink, "https://") {
		parts := strings.SplitN(selfLink, "/projects/", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid selfLink format: %s", selfLink)
		}
		path = "projects/" + parts[1]
	}

	if _, _, _, err := ParseInstanceGroupManagerExternal(path); err != nil {
		return "", fmt.Errorf("failed to parse selfLink: %w", err)
	}
	return path, nil
}
