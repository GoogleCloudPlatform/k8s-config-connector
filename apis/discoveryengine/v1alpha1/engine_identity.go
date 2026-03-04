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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// New builds a DiscoveryEngineEngineID from the Config Connector Engine object.
func NewDiscoveryEngineEngineIDFromObject(ctx context.Context, reader client.Reader, obj *DiscoveryEngineEngine) (*DiscoveryEngineEngineID, error) {

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
	collectionID := obj.Spec.Collection

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &DiscoveryEngineEngineID{
		CollectionLink: &CollectionLink{
			ProjectAndLocation: &ProjectAndLocation{
				ProjectID: projectID,
				Location:  location,
			},
			Collection: collectionID,
		},
		Engine: resourceID,
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusID, err := parseDiscoveryEngineEngineExternal(externalRef)
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
