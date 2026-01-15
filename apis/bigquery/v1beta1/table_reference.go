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

package v1beta1

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &TableRef{}

// TableRef defines the resource reference to BigQueryTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableRef struct {
	// A reference to an externally-managed BigQueryTable resource.
	// Should be in the format "projects/{{projectID}}/datasets/{{datasetsID}}/tables/{{tableID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryTable resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryTable resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryTable.
// If the "External" is given in the other resource's spec.BigQueryTableRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryTable object from the cluster.
func (r *TableRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryTableGVK.Kind)
	}
	// From given External
	if r.External != "" {
		identity := &TableIdentity{}
		if err := identity.FromExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(BigQueryTableGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryTableGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	// For non-direct resources, there's no status.externalRef
	if !found {
		// Get resourceID
		tableID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return "", err
		}
		// Resolve parent
		obj := &BigQueryTable{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		datasetRef := obj.Spec.DatasetRef
		datasetExternal, err := datasetRef.NormalizedExternal(ctx, reader, otherNamespace)
		if err != nil {
			return "", err
		}
		r.External = fmt.Sprintf("%s/tables/%s", datasetExternal, tableID)
		return r.External, nil
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}
