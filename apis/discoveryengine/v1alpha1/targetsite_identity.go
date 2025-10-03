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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TargetSiteIdentity defines the full identity for a DataStoreTargetSite
//
// +k8s:deepcopy-gen=false
type TargetSiteIdentity struct {
	*DiscoveryEngineDataStoreID
	TargetSite string
}

func (l *TargetSiteIdentity) String() string {
	return l.DiscoveryEngineDataStoreID.String() + "/siteSearchEngine/targetSites/" + l.TargetSite
}

// NewTargetSiteIdentityFromObject builds a TargetSiteIdentity from the Config Connector object.
func NewTargetSiteIdentityFromObject(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStoreTargetSite) (*DiscoveryEngineDataStoreID, *TargetSiteIdentity, error) {
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

	var link *TargetSiteIdentity

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

func ParseTargetSiteExternal(external string) (*TargetSiteIdentity, error) {
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
		targetStoreLink := &TargetSiteIdentity{
			DiscoveryEngineDataStoreID: dataStoreLink,
			TargetSite:                 tokens[10],
		}
		return targetStoreLink, nil
	}
	return nil, fmt.Errorf("format of DiscoveryEngineDataStoreTargetSite external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/dataStores/{{dataStoreID}}/siteSearchEngine/targetSites/{{targetSiteID}})", external)
}
