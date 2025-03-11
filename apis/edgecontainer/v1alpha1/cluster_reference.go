<<<<<<< HEAD
// Copyright 2024 Google LLC
=======
// Copyright 2025 Google LLC
>>>>>>> e9922e25c (conductor: "Generated types and mapper for EdgeContainerCluster")
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
<<<<<<< HEAD
	"strings"
=======
>>>>>>> e9922e25c (conductor: "Generated types and mapper for EdgeContainerCluster")

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ClusterRef{}
var EdgeContainerClusterGVK = GroupVersion.WithKind("EdgeContainerCluster")

// ClusterRef defines the resource reference to EdgeContainerCluster, which "External" field
// holds the GCP identifier for the KRM object.
type ClusterRef struct {
	// A reference to an externally managed EdgeContainerCluster resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}}".
	External string `json:"external,omitempty"`

	// The name of a AlloyDBCluster resource.
	Name string `json:"name,omitempty"`

	// The namespace of a AlloyDBCluster resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on EdgeContainerCluster.
// If the "External" is given in the other resource's spec.EdgeContainerClusterRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual EdgeContainerCluster object from the cluster.
func (r *ClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", EdgeContainerClusterGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseClusterExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(EdgeContainerClusterGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", EdgeContainerClusterGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

func ParseClusterExternal(external string) (fullResourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "clusters" {
		return "", fmt.Errorf("format of EdgeContainerCluster external=%q was not known (use projects/{{projectID}}/locations/{{location}}/clusters/{{clusterID}})", external)
	}
	resourceID := tokens[5]
	return resourceID, nil
}
