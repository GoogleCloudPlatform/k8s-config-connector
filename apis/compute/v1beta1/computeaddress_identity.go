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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ComputeAddressIdentity struct {
	id       string
	project  string
	location string // region name or "global"
}

func (i *ComputeAddressIdentity) String() string {
	if i.location == "global" {
		return fmt.Sprintf("projects/%s/global/addresses/%s", i.project, i.id)
	}
	return fmt.Sprintf("projects/%s/regions/%s/addresses/%s", i.project, i.location, i.id)
}

func (i *ComputeAddressIdentity) Project() string {
	return i.project
}

func (i *ComputeAddressIdentity) Location() string {
	return i.location
}

func (i *ComputeAddressIdentity) ID() string {
	return i.id
}

func NewComputeAddressIdentity(ctx context.Context, reader client.Reader, obj *ComputeAddress, u *unstructured.Unstructured) (*ComputeAddressIdentity, error) {
	// Get projectID
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}
	// Get Location
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location is required")
	}

	// Get resourceID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.SelfLink)
	if externalRef != "" {
		actualIdentity, err := ParseComputeAddressExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.project, projectID)
		}
		if actualIdentity.location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.location, location)
		}
		if actualIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.id)
		}
	}

	return &ComputeAddressIdentity{
		project:  projectID,
		location: location,
		id:       resourceID,
	}, nil
}

func ParseComputeAddressExternal(external string) (*ComputeAddressIdentity, error) {
	external = strings.TrimPrefix(external, "https://www.googleapis.com/compute/v1/")
	external = strings.TrimPrefix(external, "https://compute.googleapis.com/compute/v1/")
	external = strings.TrimPrefix(external, "https://www.googleapis.com/compute/beta/")
	external = strings.TrimPrefix(external, "https://compute.googleapis.com/compute/beta/")
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses" {
		return &ComputeAddressIdentity{
			project:  tokens[1],
			location: tokens[3],
			id:       tokens[5],
		}, nil
	}
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "addresses" {
		return &ComputeAddressIdentity{
			project:  tokens[1],
			location: "global",
			id:       tokens[4],
		}, nil
	}
	return nil, fmt.Errorf("format of ComputeAddress external=%q was not known (use projects/{{projectID}}/regions/{{region}}/addresses/{{addressID}} or projects/{{projectID}}/global/addresses/{{addressID}})", external)
}
