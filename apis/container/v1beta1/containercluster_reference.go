// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"regexp"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ContainerClusterRef{}
var ContainerClusterGVK = GroupVersion.WithKind("ContainerCluster")

type ContainerClusterRef struct {
	// The GKE cluster. Valid formats:
	//  `projects/{projectID}/locations/{location}/clusters/{clusterID}`
	//  `projects/{projectID}/zones/{zone}/clusters/{clusterID}`
	External string `json:"external,omitempty"`

	// Name of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	// Namespace of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ContainerCluster.
// If the "External" is given in the other resource's spec.ContainerClusterRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ContainerCluster object from the cluster.
// NOTE: ContainerCluster is currently a TF-based resource, so we need to rely on "status.selfLink" instead of "status.externalRef".
func (r *ContainerClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ContainerClusterGVK.Kind)
	}

	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ContainerClusterGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ContainerClusterGVK, key, err)
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", u.GetKind(), u.GetNamespace())
	}

	// Example "status.selfLink":
	//  - https://container.googleapis.com/v1beta1/projects/${projectId}/zones/us-central1-a/clusters/cluster-${uniqueId}
	//  - https://container.googleapis.com/v1/projects/${projectId}/locations/us-central1/clusters/cluster-${uniqueId}

	// remove service prefix
	r.External = strings.TrimPrefix(selfLink, "https://container.googleapis.com/")
	// remove version prefix
	r.External = regexp.MustCompile(`^(v1beta1/|v1/)`).ReplaceAllString(r.External, "")
	return r.External, nil
}
