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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &DiscoveryEngineDataStoreRef{}

// DiscoveryEngineDataStoreRef defines the resource reference to DiscoveryEngineDataStore, which "External" field
// holds the GCP identifier for the KRM object.
type DiscoveryEngineDataStoreRef struct {
	// A reference to an externally managed DiscoveryEngineDataStore resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/datastores/{{datastoreID}}".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineDataStore resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineDataStore resource.
	Namespace string `json:"namespace,omitempty"`
}

// DiscoveryEngineDataStoreRef defines the resource reference to DiscoveryEngineDataStore, which "External" field
// holds the GCP identifier for the KRM object.
type DiscoveryEngineDataStoreID struct {
	*CollectionLink
	DataStore string
}

// NormalizedExternal provision the "External" value for other resource that depends on DiscoveryEngineDataStore.
// If the "External" is given in the other resource's spec.DiscoveryEngineDataStoreRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual DiscoveryEngineDataStore object from the cluster.
func (r *DiscoveryEngineDataStoreRef) NormalizedExternal(ctx context.Context, reader client.Reader, defaultNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineDataStoreGVK.Kind)
	}
	// From given External
	if r.External != "" {
		id, err := ParseDiscoveryEngineDataStoreExternal(r.External)
		if err != nil {
			return "", err
		}
		r.External = id.String()
		return r.External, nil
	}

	// From the Config Connector object
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(DiscoveryEngineDataStoreGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", DiscoveryEngineDataStoreGVK, key, err)
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

// New builds a DiscoveryEngineDataStoreRef from the Config Connector DiscoveryEngineDataStore object.
func NewDiscoveryEngineDataStoreIDFromObject(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStore) (*DiscoveryEngineDataStoreID, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &DiscoveryEngineDataStoreID{
		CollectionLink: &CollectionLink{
			ProjectAndLocation: &ProjectAndLocation{
				ProjectID: projectRef.ProjectID,
				Location:  direct.ValueOf(obj.Spec.Location),
			},
			Collection: direct.ValueOf(obj.Spec.Collection),
		},
		DataStore: resourceID,
	}

	// Validate the status.externalRef, if set
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusID, err := ParseDiscoveryEngineDataStoreExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if statusID.String() != id.String() {
			return nil, fmt.Errorf("cannot change object key after creation; status=%q, new=%q",
				statusID.String(), id.String())
		}
	}
	return id, nil
}

type ProjectAndLocation struct {
	ProjectID string
	Location  string
}

func (p *ProjectAndLocation) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

type CollectionLink struct {
	*ProjectAndLocation
	Collection string
}

func (p *CollectionLink) String() string {
	return p.ProjectAndLocation.String() + "/collections/" + p.Collection
}

func (p *DiscoveryEngineDataStoreID) String() string {
	return p.CollectionLink.String() + "/dataStores/" + p.DataStore
}

func ParseDiscoveryEngineDataStoreExternal(external string) (*DiscoveryEngineDataStoreID, error) {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" {
		projectAndLocation := &ProjectAndLocation{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		collection := &CollectionLink{
			ProjectAndLocation: projectAndLocation,
			Collection:         tokens[5],
		}
		return &DiscoveryEngineDataStoreID{
			CollectionLink: collection,
			DataStore:      tokens[7],
		}, nil
	}
	return nil, fmt.Errorf("format of DiscoveryEngineDataStore external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}})", external)
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
