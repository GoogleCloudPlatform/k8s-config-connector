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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &CloudRunInstanceIdentity{}

var cloudRunInstanceURL = gcpurls.Template[CloudRunInstanceIdentity](
	"run.googleapis.com",
	"projects/{project}/locations/{location}/instances/{instance}",
)

// CloudRunInstanceIdentity defines the resource reference to CloudRunInstance.
// +k8s:deepcopy-gen=false
type CloudRunInstanceIdentity struct {
	project  string
	location string
	instance string
}

func (i *CloudRunInstanceIdentity) FromExternal(ref string) error {
	out, match, err := cloudRunInstanceURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of CloudRunInstance external=%q was not known (use %s)", ref, cloudRunInstanceURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *CloudRunInstanceIdentity) String() string {
	return cloudRunInstanceURL.ToString(*i)
}

func (i *CloudRunInstanceIdentity) Host() string {
	return cloudRunInstanceURL.Host()
}

// ExternalIdentifier implements identity.IdentityV2
func (i *CloudRunInstanceIdentity) ExternalIdentifier() string {
	return i.String()
}

// BuildCloudRunInstanceIdentity builds a CloudRunInstanceIdentity from the Config Connector CloudRunInstance object.
func BuildCloudRunInstanceIdentity(ctx context.Context, reader client.Reader, obj *CloudRunInstance) (*CloudRunInstanceIdentity, error) {

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
	if location == nil || *location == "" {
		return nil, fmt.Errorf("cannot resolve location")
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
		actualIdentity := &CloudRunInstanceIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.project, projectID)
		}
		if actualIdentity.location != *location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.location, *location)
		}
		if actualIdentity.instance != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.instance)
		}
	}
	return &CloudRunInstanceIdentity{
		project:  projectID,
		location: *location,
		instance: resourceID,
	}, nil
}
