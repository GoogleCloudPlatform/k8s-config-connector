// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DiscoveryEngineServingConfigIdentity{}
	_ identity.Resource   = &DiscoveryEngineServingConfig{}
)

// +k8s:deepcopy-gen=false
type DiscoveryEngineServingConfigIdentity struct {
	Project       string
	Location      string
	Collection    string
	Engine        string // Only one of Engine or DataStore will be non-empty
	DataStore     string // Only one of Engine or DataStore will be non-empty
	ServingConfig string
}

func (i *DiscoveryEngineServingConfigIdentity) String() string {
	if i.Engine != "" {
		return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/servingConfigs/%s",
			i.Project, i.Location, i.Collection, i.Engine, i.ServingConfig)
	}
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/servingConfigs/%s",
		i.Project, i.Location, i.Collection, i.DataStore, i.ServingConfig)
}

func (i *DiscoveryEngineServingConfigIdentity) ParentString() string {
	if i.Engine != "" {
		return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s",
			i.Project, i.Location, i.Collection, i.Engine)
	}
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s",
		i.Project, i.Location, i.Collection, i.DataStore)
}

func (i *DiscoveryEngineServingConfigIdentity) FromExternal(ref string) error {
	s := ref
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[8] == "servingConfigs" {
		if tokens[6] == "engines" {
			*i = DiscoveryEngineServingConfigIdentity{
				Project:       tokens[1],
				Location:      tokens[3],
				Collection:    tokens[5],
				Engine:        tokens[7],
				ServingConfig: tokens[9],
			}
			return nil
		}
		if tokens[6] == "dataStores" {
			*i = DiscoveryEngineServingConfigIdentity{
				Project:       tokens[1],
				Location:      tokens[3],
				Collection:    tokens[5],
				DataStore:     tokens[7],
				ServingConfig: tokens[9],
			}
			return nil
		}
	}
	return fmt.Errorf("format of DiscoveryEngineServingConfig external=%q was not known (use projects/{{project}}/locations/{{location}}/collections/{{collection}}/engines/{{engine}}/servingConfigs/{{servingConfig}} or projects/{{project}}/locations/{{location}}/collections/{{collection}}/dataStores/{{dataStore}}/servingConfigs/{{servingConfig}})", ref)
}

func (i *DiscoveryEngineServingConfigIdentity) Host() string {
	return "discoveryengine.googleapis.com"
}

func getIdentityFromDiscoveryEngineServingConfigSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineServingConfig) (*DiscoveryEngineServingConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location is empty")
	}

	collection := obj.Spec.Collection
	if collection == "" {
		return nil, fmt.Errorf("spec.collection is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if obj.Spec.EngineRef == nil && obj.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("one of spec.engineRef or spec.dataStoreRef must be set")
	}
	if obj.Spec.EngineRef != nil && obj.Spec.DataStoreRef != nil {
		return nil, fmt.Errorf("cannot specify both spec.engineRef and spec.dataStoreRef")
	}

	var engine, dataStore string

	if obj.Spec.EngineRef != nil {
		engineRef := *obj.Spec.EngineRef
		normalizedEngine, err := engineRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.engineRef: %w", err)
		}
		engineID, err := parseDiscoveryEngineEngineExternal(normalizedEngine)
		if err != nil {
			return nil, fmt.Errorf("parsing spec.engineRef external: %w", err)
		}

		if engineID.ProjectID != projectID {
			return nil, fmt.Errorf("resolved spec.engineRef project %q does not match spec.projectRef %q", engineID.ProjectID, projectID)
		}
		if engineID.Location != location {
			return nil, fmt.Errorf("resolved spec.engineRef location %q does not match spec.location %q", engineID.Location, location)
		}
		if engineID.Collection != collection {
			return nil, fmt.Errorf("resolved spec.engineRef collection %q does not match spec.collection %q", engineID.Collection, collection)
		}
		engine = engineID.Engine
	}

	if obj.Spec.DataStoreRef != nil {
		dataStoreRef := *obj.Spec.DataStoreRef
		normalizedDataStore, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
		}
		dataStoreID, err := ParseDiscoveryEngineDataStoreExternal(normalizedDataStore)
		if err != nil {
			return nil, fmt.Errorf("parsing spec.dataStoreRef external: %w", err)
		}

		if dataStoreID.ProjectID != projectID {
			return nil, fmt.Errorf("resolved spec.dataStoreRef project %q does not match spec.projectRef %q", dataStoreID.ProjectID, projectID)
		}
		if dataStoreID.Location != location {
			return nil, fmt.Errorf("resolved spec.dataStoreRef location %q does not match spec.location %q", dataStoreID.Location, location)
		}
		if dataStoreID.Collection != collection {
			return nil, fmt.Errorf("resolved spec.dataStoreRef collection %q does not match spec.collection %q", dataStoreID.Collection, collection)
		}
		dataStore = dataStoreID.DataStore
	}

	identity := &DiscoveryEngineServingConfigIdentity{
		Project:       projectID,
		Location:      location,
		Collection:    collection,
		Engine:        engine,
		DataStore:     dataStore,
		ServingConfig: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineServingConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineServingConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineServingConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineServingConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineServingConfig) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
