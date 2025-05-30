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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &BigQueryAnalyticsHubListingRef{}

// BigQueryAnalyticsHubListingRef defines the resource reference to BigQueryAnalyticsHubListing, which "External" field
// holds the GCP identifier for the KRM object.
type BigQueryAnalyticsHubListingRef struct {
	// A reference to an externally managed BigQueryAnalyticsHubListing resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/listings/{{listingID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryAnalyticsHubListing resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryAnalyticsHubListing resource.
	Namespace string `json:"namespace,omitempty"`

	parent *BigQueryAnalyticsHubListingParent
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryAnalyticsHubListing.
// If the "External" is given in the other resource's spec.BigQueryAnalyticsHubListingRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryAnalyticsHubListing object from the cluster.
func (r *BigQueryAnalyticsHubListingRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryAnalyticsHubListingGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseBigQueryAnalyticsHubListingExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(BigQueryAnalyticsHubListingGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryAnalyticsHubListingGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("BigQueryAnalyticsHubListing is not ready yet")
	}
	r.External = actualExternalRef
	return r.External, nil
}

// New builds a BigQueryAnalyticsHubListingRef from the Config Connector BigQueryAnalyticsHubListing object.
func NewBigQueryAnalyticsHubListingRef(ctx context.Context, reader client.Reader, obj *BigQueryAnalyticsHubListing) (*BigQueryAnalyticsHubListingRef, error) {
	id := &BigQueryAnalyticsHubListingRef{}

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
	if location == "" {
		return nil, fmt.Errorf("location cannot be empty")
	}

	if obj.Spec.DataExchangeRef == nil {
		return nil, fmt.Errorf("spec.dataExchangeRef cannot be empty")
	}
	dataExchangeExternal, err := obj.Spec.DataExchangeRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot normalize dataExchangeRef for listing ref: %w", err)
	}

	dataExchangeID, err := v1beta1.ParseDataExchangeIdentity(dataExchangeExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse dataExchangeRef for listing ref: %w", err)
	}

	deID := dataExchangeID.ID()
	id.parent = &BigQueryAnalyticsHubListingParent{ProjectID: projectID, Location: location, DataExchangeID: deID}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if deID == "" {
		return nil, fmt.Errorf("cannot resolve dataExchange ID")
	}

	// Use approved External
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asBigQueryAnalyticsHubListingExternal(id.parent, resourceID)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualResourceID, err := parseBigQueryAnalyticsHubListingExternal(externalRef)
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
	if actualParent.DataExchangeID != deID {
		return nil, fmt.Errorf("spec.dataExchangeRef changed, expect %s, got %s", actualParent.DataExchangeID, deID)
	}

	id.External = externalRef
	id.parent = &BigQueryAnalyticsHubListingParent{ProjectID: projectID, Location: location, DataExchangeID: deID}
	return id, nil
}

func (r *BigQueryAnalyticsHubListingRef) Parent() (*BigQueryAnalyticsHubListingParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseBigQueryAnalyticsHubListingExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("BigQueryAnalyticsHubListingRef not initialized from `NewBigQueryAnalyticsHubListingRef` or `NormalizedExternal`")
}

type BigQueryAnalyticsHubListingParent struct {
	ProjectID      string
	Location       string
	DataExchangeID string
}

func (p *BigQueryAnalyticsHubListingParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/dataExchanges/" + p.DataExchangeID
}

func asBigQueryAnalyticsHubListingExternal(parent *BigQueryAnalyticsHubListingParent, resourceID string) (external string) {
	return parent.String() + "/listings/" + resourceID
}

func parseBigQueryAnalyticsHubListingExternal(external string) (parent *BigQueryAnalyticsHubListingParent, resourceID string, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "dataExchanges" || tokens[6] != "listings" {
		return nil, "", fmt.Errorf("format of BigQueryAnalyticsHubListing external=%q was not known (use projects/{{projectID}}/locations/{{location}}/dataExchanges/{{dataExchangeID}}/listings/{{listingID}})", external)
	}
	parent = &BigQueryAnalyticsHubListingParent{
		ProjectID:      tokens[1],
		Location:       tokens[3],
		DataExchangeID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
