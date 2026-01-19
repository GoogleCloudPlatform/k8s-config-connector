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

var jobURL = gcpurls.Template[JobIdentity](
	"run.googleapis.com",
	"projects/{projectID}/locations/{location}/jobs/{jobID}",
)

// JobIdentity defines the resource reference to RunJob, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type JobIdentity struct {
	ProjectID string
	Location  string
	JobID     string
}

func (i *JobIdentity) FromExternal(ref string) error {
	out, match, err := jobURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of RunJob external=%q was not known (use %s)", ref, jobURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *JobIdentity) String() string {
	return jobURL.ToString(*i)
}

func (i *JobIdentity) ID() string {
	return i.JobID
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
		actualIdentity := &JobIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.ProjectID, projectID)
		}
		if actualIdentity.Location != *location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, *location)
		}
		if actualIdentity.JobID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.JobID)
		}
	}
	return &JobIdentity{
		ProjectID: projectID,
		Location:  *location,
		JobID:     resourceID,
	}, nil
}