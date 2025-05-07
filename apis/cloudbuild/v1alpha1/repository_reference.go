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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	//apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &RepositoryRef{}

var RespositoryGVK = GroupVersion.WithKind("CloudBuildRepository")

// RepositoryRef defines the resource reference to CloudBuildRepository, which "External" field
// holds the GCP identifier for the KRM object.
type RepositoryRef struct {
	// A reference to an externally managed CloudBuildRepository resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/connections/{{connectionID}}/repositories/{{RepositoryID}}".
	External string `json:"external,omitempty"`

	// The name of a CloudBuildRepository resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudBuildRepository resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on CloudBuildRepository.
// If the "External" is given in the other resource's spec.CloudBuildRepositoryRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual CloudBuildRepository object from the cluster.
func (r *RepositoryRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", RespositoryGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseRepositoryExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// NOT YET
	// // From the Config Connector object
	// if r.Namespace == "" {
	// 	r.Namespace = otherNamespace
	// }
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(RespositoryGVK)
	// if err := reader.Get(ctx, key, u); err != nil {
	// 	if apierrors.IsNotFound(err) {
	// 		return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	// 	}
	// 	return "", fmt.Errorf("reading referenced %s %s: %w", RespositoryGVK, key, err)
	// }
	// // Get external from status.externalRef. This is the most trustworthy place.
	// actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	// if err != nil {
	// 	return "", fmt.Errorf("reading status.externalRef: %w", err)
	// }
	// if actualExternalRef == "" {
	// 	return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	// }
	// r.External = actualExternalRef
	// return r.External, nil
	return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
}

func ParseRepositoryExternal(external string) (parent *RepositoryParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" || tokens[6] != "repositories" {
		return nil, "", fmt.Errorf("format of DeployRepository external=%q was not known (use projects/{{projectID}}/locations/{{location}}/connections/{{connectionID}}/repositories/{{repositoryID}})", external)
	}
	parent = &RepositoryParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

type RepositoryParent struct {
	ProjectID string
	Location  string
}
