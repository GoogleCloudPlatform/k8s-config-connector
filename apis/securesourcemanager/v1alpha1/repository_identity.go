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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RepositoryIdentity struct {
	id     string
	parent *RepositoryParent
}

func (i *RepositoryIdentity) String() string {
	p := i.id
	return i.parent.String() + "/repositories/" + p
}

func (i *RepositoryIdentity) Parent() *RepositoryParent {
	return i.parent
}

func (i *RepositoryIdentity) ID() string {
	return i.id
}

type RepositoryParent struct {
	ProjectID string
	Location  string
}

func (p *RepositoryParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func NewRepositoryIdentity(ctx context.Context, reader client.Reader, obj *SecureSourceManagerRepository, u *unstructured.Unstructured) (*RepositoryIdentity, error) {
	// Get projectID
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}
	// Get Location
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
		actualIdentity, err := parseSecureSourceManagerRepositoryExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.parent.ProjectID, projectID)
		}
		if actualIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.id)
		}
	}

	return &RepositoryIdentity{
		parent: &RepositoryParent{ProjectID: projectID, Location: location},
		id:     resourceID,
	}, nil
}
