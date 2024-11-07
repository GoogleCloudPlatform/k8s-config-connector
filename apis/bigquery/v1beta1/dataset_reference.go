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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &BigQueryDatasetRef{}

// BigQueryDatasetRef defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
type BigQueryDatasetRef struct {
	// A reference to an externally managed BigQueryDataset resource.
	// Should be in the format "projects/<projectID>/locations/<location>/datasets/<datasetID>".
	External string `json:"external,omitempty"`

	// The name of a BigQueryDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryDataset resource.
	Namespace string `json:"namespace,omitempty"`

	parent *BigQueryDatasetParent
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryDataset.
// If the "External" is given in the other resource's spec.BigQueryDatasetRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryDataset object from the cluster.
func (r *BigQueryDatasetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryDatasetGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseBigQueryDatasetExternal(r.External); err != nil {
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

// New builds a BigQueryDatasetRef from the Config Connector BigQueryDataset object.
func NewBigQueryDatasetRef(ctx context.Context, reader client.Reader, obj *BigQueryDataset) (*BigQueryDatasetRef, error) {
	id := &BigQueryDatasetRef{}

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	id.parent = &BigQueryDatasetParent{ProjectID: projectID}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asBigQueryDatasetExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := ParseBigQueryDatasetExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &BigQueryDatasetParent{ProjectID: projectID}
	return id, nil
}

func (r *BigQueryDatasetRef) Parent() (*BigQueryDatasetParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := ParseBigQueryDatasetExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("BigQueryDatasetRef not initialized from `NewBigQueryDatasetRef` or `NormalizedExternal`")
}

type BigQueryDatasetParent struct {
	ProjectID string
}

func (p *BigQueryDatasetParent) String() string {
	return "projects/" + p.ProjectID
}

func asBigQueryDatasetExternal(parent *BigQueryDatasetParent, resourceID string) (external string) {
	// Link Reference https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets/get
	return parent.String() + "/datasets/" + resourceID
}

func ParseBigQueryDatasetExternal(external string) (parent *BigQueryDatasetParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "datasets" {
		return nil, "", fmt.Errorf("format of BigQueryDataset external=%q was not known (use projects/<projectId>/datasets/<datasetID>)", external)
	}
	parent = &BigQueryDatasetParent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
