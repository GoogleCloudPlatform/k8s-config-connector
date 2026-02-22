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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var RunServiceIdentityFormat = gcpurls.Template[RunServiceIdentity]("run.googleapis.com", "projects/{project}/locations/{location}/services/{service}")

var _ identity.Identity = &RunServiceIdentity{}

// +k8s:deepcopy-gen=false
type RunServiceIdentity struct {
	Project  string
	Location string
	Service  string
}

func (i *RunServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := RunServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of RunService external=%q was not known (use %s): %w", ref, RunServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of RunService external=%q was not known (use %s)", ref, RunServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *RunServiceIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", i.Project, i.Location, i.Service)
}

func (i *RunServiceIdentity) Host() string {
	return RunServiceIdentityFormat.Host()
}

// NewServiceIdentity builds a RunServiceIdentity from the Config Connector RunService object.
func NewServiceIdentity(ctx context.Context, reader client.Reader, obj *RunService) (*RunServiceIdentity, error) {

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
		actualIdentity := &RunServiceIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Location != *location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, *location)
		}
		if actualIdentity.Service != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.Service)
		}
	}
	return &RunServiceIdentity{
		Project:  projectID,
		Location: *location,
		Service:  resourceID,
	}, nil
}
