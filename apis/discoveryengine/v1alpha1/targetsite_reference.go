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

var _ refsv1beta1.ExternalNormalizer = &TargetSiteRef{}

// TargetSiteRef defines the resource reference to DiscoveryEngineDataStoreTargetSite, which "External" field
// holds the GCP identifier for the KRM object.
type TargetSiteRef struct {
	// A reference to an externally managed DiscoveryEngineDataStoreTargetSite resource.
	// Should be in the format "projects/<projectID>/locations/<location>/targetsites/<targetsiteID>".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineDataStoreTargetSite resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineDataStoreTargetSite resource.
	Namespace string `json:"namespace,omitempty"`
}

// TargetSiteLink defines the full identity for a DataStoreTargetSite
//
// +k8s:deepcopy-gen=false
type TargetSiteLink struct {
	*DiscoveryEngineDataStoreID
	TargetSite string
}

func (l *TargetSiteLink) String() string {
	return l.DiscoveryEngineDataStoreID.String() + "/siteSearchEngine/targetSites/" + l.TargetSite
}

// NormalizedExternal provision the "External" value for other resource that depends on DiscoveryEngineDataStoreTargetSite.
// If the "External" is given in the other resource's spec.DiscoveryEngineDataStoreTargetSiteRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual DiscoveryEngineDataStoreTargetSite object from the cluster.
func (r *TargetSiteRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineDataStoreTargetSiteGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseTargetSiteExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(DiscoveryEngineDataStoreTargetSiteGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", DiscoveryEngineDataStoreTargetSiteGVK, key, err)
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

// NewTargetSiteLinkFromObject builds a TargetSiteLink from the Config Connector object.
func NewTargetSiteLinkFromObject(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStoreTargetSite) (*DiscoveryEngineDataStoreID, *TargetSiteLink, error) {
	if obj.Spec.DataStoreRef == nil {
		return nil, nil, fmt.Errorf("spec.dataStoreRef not set")
	}
	dataStoreRef := *obj.Spec.DataStoreRef
	if _, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}
	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(dataStoreRef.External)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing dataStoreRef.external=%q: %w", dataStoreRef.External, err)
	}

	var link *TargetSiteLink

	// Validate the status.externalRef, if set
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		externalLink, err := ParseTargetSiteExternal(externalRef)
		if err != nil {
			return nil, nil, err
		}
		if externalLink.DiscoveryEngineDataStoreID.String() != dataStoreLink.String() {
			return nil, nil, fmt.Errorf("cannot change object key after creation; status=%q, new=%q",
				externalLink.DiscoveryEngineDataStoreID.String(), dataStoreLink.String())
		}
		link = externalLink
	}
	return dataStoreLink, link, nil
}

func ParseTargetSiteExternal(external string) (*TargetSiteLink, error) {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 11 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "siteSearchEngine" && tokens[9] == "targetSites" {
		projectAndLocation := &ProjectAndLocation{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		collection := &CollectionLink{
			ProjectAndLocation: projectAndLocation,
			Collection:         tokens[5],
		}
		dataStoreLink := &DiscoveryEngineDataStoreID{
			CollectionLink: collection,
			DataStore:      tokens[7],
		}
		targetStoreLink := &TargetSiteLink{
			DiscoveryEngineDataStoreID: dataStoreLink,
			TargetSite:                 tokens[10],
		}
		return targetStoreLink, nil
	}
	return nil, fmt.Errorf("format of DiscoveryEngineDataStoreTargetSite external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}}/siteSearchEngine/targetSites/{{targetSiteID}})", external)
}
