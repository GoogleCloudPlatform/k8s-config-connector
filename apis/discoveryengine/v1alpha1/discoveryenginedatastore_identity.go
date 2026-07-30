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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DataStoreIdentity is the identity of a DiscoveryEngineDataStore.
type DataStoreIdentity struct {
	parent     *DataStoreParent
	collection string
	id         string
}

func (i *DataStoreIdentity) String() string {
	return i.parent.String() + "/collections/" + i.collection + "/dataStores/" + i.id
}

func (i *DataStoreIdentity) ID() string {
	return i.id
}

func (i *DataStoreIdentity) Parent() *DataStoreParent {
	return i.parent
}

type DataStoreParent struct {
	ProjectID string
	Location  string
}

func (p *DataStoreParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a DataStoreIdentity from the Config Connector DataStore object.
func NewDataStoreIdentity(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataStore) (*DataStoreIdentity, error) {

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

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualCollection, actualResourceID, err := ParseDataStoreExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualCollection != obj.Spec.Collection {
			return nil, fmt.Errorf("spec.collection changed, expect %s, got %s", actualCollection, obj.Spec.Collection)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &DataStoreIdentity{
		parent: &DataStoreParent{
			ProjectID: projectID,
			Location:  location,
		},
		collection: obj.Spec.Collection,
		id:         resourceID,
	}, nil
}

func ParseDataStoreExternal(external string) (parent *DataStoreParent, collection string, resourceID string, err error) {
	id, err := ParseDiscoveryEngineDataStoreExternal(external)
	if err != nil {
		return nil, "", "", err
	}
	parent = &DataStoreParent{
		ProjectID: id.ProjectID,
		Location:  id.Location,
	}
	return parent, id.Collection, id.DataStore, nil
}

func (i *DataStoreIdentity) FromExternal(ref string) error {
	s := ref
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	parent, collection, resourceID, err := ParseDataStoreExternal(s)
	if err != nil {
		return err
	}
	i.parent = parent
	i.collection = collection
	i.id = resourceID
	return nil
}

func (obj *DiscoveryEngineDataStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewDataStoreIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
