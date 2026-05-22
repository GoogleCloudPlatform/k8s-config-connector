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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var AppProfileIdentityFormat = gcpurls.Template[AppProfileIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}/appProfiles/{appProfile}")

// AppProfileIdentity defines the resource reference to BigtableAppProfile, which "External" field
// holds the GCP identifier for the KRM object.
type AppProfileIdentity struct {
	Project    string
	Instance   string
	AppProfile string
}

func (i *AppProfileIdentity) String() string {
	return AppProfileIdentityFormat.ToString(*i)
}

func (i *AppProfileIdentity) FromExternal(ref string) error {
	parsed, match, err := AppProfileIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AppProfile external=%q was not known (use %s): %w", ref, AppProfileIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AppProfile external=%q was not known (use %s)", ref, AppProfileIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AppProfileIdentity) ID() string {
	return i.AppProfile
}

func (i *AppProfileIdentity) Parent() *BigtableInstanceIdentity {
	return &BigtableInstanceIdentity{
		Project:  i.Project,
		Instance: i.Instance,
	}
}

func (i *AppProfileIdentity) ParentString() string {
	return i.Parent().String()
}

func (i *AppProfileIdentity) ParentInstanceIdString() string {
	return i.Instance
}

// NewAppProfileIdentity builds a AppProfileIdentity from the Config Connector AppProfile object.
func NewAppProfileIdentity(ctx context.Context, reader client.Reader, obj *BigtableAppProfile) (*AppProfileIdentity, error) {
	// Get Parent
	if err := obj.Spec.BigtableInstanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	instanceRef := obj.Spec.BigtableInstanceRef.External
	instanceID := &BigtableInstanceIdentity{}
	if err := instanceID.FromExternal(instanceRef); err != nil {
		return nil, err
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
		actualIdentity := &AppProfileIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != instanceID.Project {
			return nil, fmt.Errorf("ProjectID in spec.instanceRef changed, expect %s, got %s", actualIdentity.Project, instanceID.Project)
		}
		if actualIdentity.Instance != instanceID.Instance {
			return nil, fmt.Errorf("instanceID in spec.instanceRef changed, expect %s, got %s", actualIdentity.Instance, instanceID.Instance)
		}
		if actualIdentity.AppProfile != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.AppProfile)
		}
	}
	return &AppProfileIdentity{
		Project:    instanceID.Project,
		Instance:   instanceID.Instance,
		AppProfile: resourceID,
	}, nil
}
