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

var _ refsv1beta1.ExternalNormalizer = &DatasetRef{}

// DatasetRef defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
type DatasetRef struct {
	// A reference to an externally-managed BigQueryDataset resource.
	// Should be in the format "projects/{{projectID}}/datasets/{{datasetID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryDataset resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryDataset.
// If the "External" is given in the other resource's spec.BigQueryDatasetRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryDataset object from the cluster.
func (r *DatasetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryDatasetGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseDatasetExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(BigQueryDatasetGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryDatasetGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if !found {
		// BigQueryDataset is still TF based so we resolve the reference from the object
		// BUT only construct the reference if the resource is ready
		resource, err := k8s.NewResource(u)
		if err != nil {
			return "", fmt.Errorf("error converting unstructured to resource: %w", err)
		}
		if !k8s.IsResourceReady(resource) {
			return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}

		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return "", err
		}
		datasetID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return "", err
		}
		actualExternal := fmt.Sprintf("projects/%s/datasets/%s", projectID, datasetID)
		r.External = actualExternal
		return r.External, nil
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}
