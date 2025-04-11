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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AppProfileIdentity defines the resource reference to BigtableAppProfile, which "External" field
// holds the GCP identifier for the KRM object.
type AppProfileIdentity struct {
	parent *InstanceIdentity
	id     string
}

func (i *AppProfileIdentity) String() string {
	return i.ParentString() + "/appProfiles/" + i.id
}

func (i *AppProfileIdentity) ID() string {
	return i.id
}

<<<<<<< HEAD
func (i *AppProfileIdentity) ParentString() string {
	return i.parent.String()
=======
func (i *AppProfileIdentity) Parent() *AppProfileParent {
	return i.parent
}

type AppProfileParent struct {
	ProjectID        string
	BigtableInstance string
}

func (p *AppProfileParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.BigtableInstance
>>>>>>> b1aac8820fed97c13e2b9c0bf52037772b60199f
}

// New builds a AppProfileIdentity from the Config Connector AppProfile object.
func NewAppProfileIdentity(ctx context.Context, reader client.Reader, obj *BigtableAppProfile) (*AppProfileIdentity, error) {
	// Get Parent
	instanceRef, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	instanceParent, instanceID, err := ParseInstanceExternal(instanceRef)
	if err != nil {
		return nil, err
	}
=======
	tokens := strings.Split(instanceRef, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "instances" {
		return nil, fmt.Errorf("Invalid format of BigtableAppProfile external=%q was not known (expected projects/{{projectID}}/instances/{{instance}})", instanceRef)
	}
	parentProjectId := tokens[1]
	parentBigtableInstance := tokens[3]
>>>>>>> b1aac8820fed97c13e2b9c0bf52037772b60199f

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
		actualParent, actualResourceID, err := ParseAppProfileExternal(externalRef)
		if err != nil {
			return nil, err
		}
<<<<<<< HEAD
		if actualParent.Parent.ProjectID != instanceParent.ProjectID {
			return nil, fmt.Errorf("ProjectID in spec.instanceRef changed, expect %s, got %s", actualParent.Parent.ProjectID, instanceParent.ProjectID)
		}
		if actualParent.Id != instanceID {
			return nil, fmt.Errorf("instanceID in spec.instanceRef changed, expect %s, got %s", actualParent.Id, instanceID)
=======
		if actualParent.ProjectID != parentProjectId {
			return nil, fmt.Errorf("ProjectID in spec.instanceRef changed, expect %s, got %s", actualParent.ProjectID, parentProjectId)
		}
		if actualParent.BigtableInstance != parentBigtableInstance {
			return nil, fmt.Errorf("BigtableInstance in spec.instanceRef changed, expect %s, got %s", actualParent.BigtableInstance, parentBigtableInstance)
>>>>>>> b1aac8820fed97c13e2b9c0bf52037772b60199f
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &AppProfileIdentity{
<<<<<<< HEAD
		parent: &InstanceIdentity{
			Parent: instanceParent,
			Id:     instanceID,
=======
		parent: &AppProfileParent{
			ProjectID:        parentProjectId,
			BigtableInstance: parentBigtableInstance,
>>>>>>> b1aac8820fed97c13e2b9c0bf52037772b60199f
		},
		id: resourceID,
	}, nil
}

func ParseAppProfileExternal(external string) (*InstanceIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "appProfiles" {
		return nil, "", fmt.Errorf("format of BigtableAppProfile external=%q was not known (use projects/{{projectID}}/instances/{{instance}}/appProfiles/{{appprofileID}})", external)
	}
	return &InstanceIdentity{Parent: &parent.ProjectParent{ProjectID: tokens[1]}, Id: tokens[3]}, tokens[5], nil
}
