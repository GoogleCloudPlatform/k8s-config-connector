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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var NotebooksScheduleIdentityFormat = gcpurls.Template[NotebooksScheduleIdentity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/schedules/{schedule}")

// NotebooksScheduleIdentity is the identity of a NotebooksSchedule.
// +k8s:deepcopy-gen=false
type NotebooksScheduleIdentity struct {
	Project  string
	Location string
	Schedule string
}

func (i *NotebooksScheduleIdentity) String() string {
	return NotebooksScheduleIdentityFormat.ToString(*i)
}

func (i *NotebooksScheduleIdentity) ID() string {
	return i.Schedule
}

func (i *NotebooksScheduleIdentity) Parent() *NotebooksScheduleParent {
	return &NotebooksScheduleParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

func (i *NotebooksScheduleIdentity) FromExternal(ref string) error {
	parsed, match, err := NotebooksScheduleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NotebooksSchedule external=%q was not known (use %s): %w", ref, NotebooksScheduleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NotebooksSchedule external=%q was not known (use %s)", ref, NotebooksScheduleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// +k8s:deepcopy-gen=false
type NotebooksScheduleParent struct {
	ProjectID string
	Location  string
}

func (p *NotebooksScheduleParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a NotebooksScheduleIdentity from the Config Connector NotebooksSchedule object.
func NewNotebooksScheduleIdentity(ctx context.Context, reader client.Reader, obj *NotebooksSchedule) (*NotebooksScheduleIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
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
		actualParent, actualResourceID, err := ParseNotebooksScheduleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &NotebooksScheduleIdentity{
		Project:  projectID,
		Location: location,
		Schedule: resourceID,
	}, nil
}

func ParseNotebooksScheduleExternal(external string) (parent *NotebooksScheduleParent, resourceID string, err error) {
	i := &NotebooksScheduleIdentity{}
	if err := i.FromExternal(external); err != nil {
		return nil, "", err
	}
	return i.Parent(), i.ID(), nil
}
