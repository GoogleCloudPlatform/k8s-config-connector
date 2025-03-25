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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &TableRef{}
var (
	BigtableTableGVK = GroupVersion.WithKind("BigtableTable")
)

// TableRef defines the resource reference to BigtableTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableRef struct {
	// A reference to an externally managed BigtableTable resource.
	External string `json:"external,omitempty"`

	// The name of a BigtableInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigtableInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on BigtableTable.
// If the "External" is given in the other resource's spec.BigtableTableRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigtableTable object from the cluster.
func (r *TableRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigtableTableGVK.Kind)
	}
	// From given External
	// For backward compatibility, we are not validating the external format.
	// todo: validate external when it's referenced by a pure direct resource
	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(BigtableTableGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigtableTableGVK, key, err)
	}

	// todo: use externalRef for resource that managed by direct controller
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return "", fmt.Errorf("reading spec.resourceID: %w", err)
	}
	if resourceID == "" {
		metadataName, _, err := unstructured.NestedString(u.Object, "metadata", "name")
		if err != nil {
			return "", fmt.Errorf("reading metadata.name: %w", err)
		}
		resourceID = metadataName
	}
	if resourceID == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = resourceID
	return r.External, nil
}
