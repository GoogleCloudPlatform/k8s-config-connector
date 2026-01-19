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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &JobIdentity{}

var jobURL = gcpurls.Template[jobURLIdentity](
	"run.googleapis.com",
	"projects/{projectID}/locations/{location}/jobs/{jobID}",
)

type jobURLIdentity struct {
	ProjectID string
	Location  string
	JobID     string
}

// JobIdentity defines the resource reference to RunJob, which "External" field
// holds the GCP identifier for the KRM object.
type JobIdentity struct {
	parent *JobParent
	id     string
}

func (i *JobIdentity) FromExternal(ref string) error {
	parent, resourceID, err := ParseJobExternal(ref)
	if err != nil {
		return err
	}
	i.parent = parent
	i.id = resourceID
	return nil
}

func (i *JobIdentity) String() string {
	return jobURL.ToString(jobURLIdentity{
		ProjectID: i.parent.ProjectID,
		Location:  i.parent.Location,
		JobID:     i.id,
	})
}

func (i *JobIdentity) ID() string {
	return i.id
}

func (i *JobIdentity) Parent() *JobParent {
	return i.parent
}

type JobParent struct {
	ProjectID string
	Location  string
}

func (p *JobParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a JobIdentity from the Config Connector Job object.
func NewJobIdentity(ctx context.Context, reader client.Reader, obj *RunJob) (*JobIdentity, error) {

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
		actualParent, actualResourceID, err := ParseJobExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != *location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, *location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &JobIdentity{
		parent: &JobParent{
			ProjectID: projectID,
			Location:  *location,
		},
		id: resourceID,
	}, nil
}

func ParseJobExternal(external string) (parent *JobParent, resourceID string, err error) {
	out, match, err := jobURL.Parse(external)
	if err != nil {
		return nil, "", err
	}
	if !match {
		return nil, "", fmt.Errorf("format of RunJob external=%q was not known (use %s)", external, jobURL.CanonicalForm())
	}
	parent = &JobParent{
		ProjectID: out.ProjectID,
		Location:  out.Location,
	}
	resourceID = out.JobID
	return parent, resourceID, nil
}
