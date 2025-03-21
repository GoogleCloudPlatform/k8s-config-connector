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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ImportJobIdentity defines the resource reference to KMSImportJob, which "External" field
// holds the GCP identifier for the KRM object.
type ImportJobIdentity struct {
	parent *ImportJobParent
	id     string
}

func (i *ImportJobIdentity) String() string {
	return i.parent.String() + "/importJobs/" + i.id
}

func (i *ImportJobIdentity) ID() string {
	return i.id
}

func (i *ImportJobIdentity) Parent() *ImportJobParent {
	return i.parent
}

type ImportJobParent struct {
	ProjectID string
	Location  string
	KeyRingID string
}

func (p *ImportJobParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/keyRings/" + p.KeyRingID
}

// New builds a ImportJobIdentity from the Config Connector ImportJob object.
func NewImportJobIdentity(ctx context.Context, reader client.Reader, obj *KMSImportJob) (*ImportJobIdentity, error) {

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

	keyRingRef := obj.Spec.KeyRingRef
	keyRingID := keyRingRef.Name
	if keyRingID == "" {
		if keyRingRef.External == "" {
			return nil, fmt.Errorf("one of spec.keyRingRef.name or spec.keyRingRef.external must be set")
		}
		keyRingID = keyRingRef.External
	}

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
		actualParent, actualResourceID, err := ParseImportJobExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.KeyRingID != keyRingID {
			return nil, fmt.Errorf("spec.keyRingRef changed, expect %s, got %s", actualParent.KeyRingID, keyRingID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ImportJobIdentity{
		parent: &ImportJobParent{
			ProjectID: projectID,
			Location:  location,
			KeyRingID: keyRingID,
		},
		id: resourceID,
	}, nil
}

func ParseImportJobExternal(external string) (parent *ImportJobParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyRings" || tokens[6] != "importJobs" {
		return nil, "", fmt.Errorf("format of KMSImportJob external=%q was not known (use projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRingID}}/importJobs/{{importJobID}})", external)
	}
	parent = &ImportJobParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		KeyRingID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
