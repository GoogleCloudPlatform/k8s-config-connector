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

package v1alpha1

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

var _ refsv1beta1.ExternalNormalizer = &BigQueryAnalyticsHubDataExchangeRef{}

// BigQueryAnalyticsHubDataExchangeRef defines the resource reference to BigQueryAnalyticsHubDataExchange, which "External" field
// holds the GCP identifier for the KRM object.
type BigQueryAnalyticsHubDataExchangeRef struct {
	// A reference to an externally managed BigQueryAnalyticsHubDataExchange resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/dataexchanges/{{dataexchangeID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryAnalyticsHubDataExchange resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryAnalyticsHubDataExchange resource.
	Namespace string `json:"namespace,omitempty"`

	parent *BigQueryAnalyticsHubDataExchangeParent
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryAnalyticsHubDataExchange.
// If the "External" is given in the other resource's spec.BigQueryAnalyticsHubDataExchangeRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryAnalyticsHubDataExchange object from the cluster.
func (r *BigQueryAnalyticsHubDataExchangeRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryAnalyticsHubDataExchangeGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseBigQueryAnalyticsHubDataExchangeExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(BigQueryAnalyticsHubDataExchangeGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryAnalyticsHubDataExchangeGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return actualExternalRef, nil
}

// New builds a BigQueryAnalyticsHubDataExchangeRef from the Config Connector BigQueryAnalyticsHubDataExchange object.
func NewBigQueryAnalyticsHubDataExchangeRef(ctx context.Context, reader client.Reader, obj *BigQueryAnalyticsHubDataExchange) (*BigQueryAnalyticsHubDataExchangeRef, error) {
	id := &BigQueryAnalyticsHubDataExchangeRef{}

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location
	id.parent = &BigQueryAnalyticsHubDataExchangeParent{ProjectID: projectID, Location: location}

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
		id.External = asBigQueryAnalyticsHubDataExchangeExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseBigQueryAnalyticsHubDataExchangeExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.ProjectID != projectID {
		return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
	}
	if actualParent.Location != location {
		return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
	}
	if actualResourceID != resourceID {
		return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
			resourceID, actualResourceID)
	}
	id.External = externalRef
	id.parent = &BigQueryAnalyticsHubDataExchangeParent{ProjectID: projectID, Location: location}
	return id, nil
}

func (r *BigQueryAnalyticsHubDataExchangeRef) Parent() (*BigQueryAnalyticsHubDataExchangeParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseBigQueryAnalyticsHubDataExchangeExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("BigQueryAnalyticsHubDataExchangeRef not initialized from `NewBigQueryAnalyticsHubDataExchangeRef` or `NormalizedExternal`")
}

type BigQueryAnalyticsHubDataExchangeParent struct {
	ProjectID string
	Location  string
}

func (p *BigQueryAnalyticsHubDataExchangeParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func asBigQueryAnalyticsHubDataExchangeExternal(parent *BigQueryAnalyticsHubDataExchangeParent, resourceID string) (external string) {
	return parent.String() + "/dataexchanges/" + resourceID
}

func parseBigQueryAnalyticsHubDataExchangeExternal(external string) (parent *BigQueryAnalyticsHubDataExchangeParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "dataexchange" {
		return nil, "", fmt.Errorf("format of BigQueryAnalyticsHubDataExchange external=%q was not known (use projects/{{projectID}}/locations/{{location}}/dataexchanges/{{dataexchangeID}})", external)
	}
	parent = &BigQueryAnalyticsHubDataExchangeParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
