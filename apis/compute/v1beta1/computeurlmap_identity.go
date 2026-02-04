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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ComputeURLMapIdentity{}

// ComputeURLMapIdentity defines the resource reference to ComputeURLMap, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeURLMapIdentity struct {
	ParentID   *parent.ComputeParent
	ResourceID string
}

func (i *ComputeURLMapIdentity) String() string {
	return i.ParentID.String() + "/urlMaps/" + i.ResourceID
}

func (i *ComputeURLMapIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	p, err := parent.ParseComputeParent(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return err
	}
	if tokens[len(tokens)-2] != "urlMaps" {
		return fmt.Errorf("format of ComputeURLMap external=%q was not known (use %s/urlMaps/{{urlMapID}}", ref, p)
	}
	i.ResourceID = tokens[len(tokens)-1]
	i.ParentID = p
	return nil
}

var _ identity.Resource = &ComputeURLMap{}

func (obj *ComputeURLMap) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &ComputeURLMapIdentity{
		ParentID:   parentID,
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ComputeURLMapIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ComputeURLMap identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ComputeURLMap) GetParentIdentity(ctx context.Context, reader client.Reader) (*parent.ComputeParent, error) {
	projectID, err := refsv1beta1.ResolveProjectFromAnnotation(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get Location
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("location is required")
	}

	return &parent.ComputeParent{ProjectID: projectID.ProjectID, Location: location}, nil
}
