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

const (
	// LakeIDLabel is the label key used to store the lake ID.
	// This is needed because DataplexTask doesn't have a LakeRef field.
	LakeIDLabel = "cnrm.cloud.google.com/lake-id"
)

// TaskIdentity defines the resource reference to DataplexTask, which "External" field
// holds the GCP identifier for the KRM object.
type TaskIdentity struct {
	parent *TaskParent
	id     string
}

func (i *TaskIdentity) String() string {
	// Uses the pattern: projects/{project}/locations/{location}/lakes/{lake}/tasks/{task}
	return i.parent.String() + "/tasks/" + i.id
}

func (i *TaskIdentity) ID() string {
	return i.id
}

func (i *TaskIdentity) Parent() *TaskParent {
	return i.parent
}

// TaskParent includes the IDs for the parent resources of a DataplexTask.
// Based on the pattern: projects/{project}/locations/{location}/lakes/{lake}/tasks/{task}
type TaskParent struct {
	ProjectID string
	Location  string
	LakeID    string
}

func (p *TaskParent) String() string {
	// Returns the parent string based on the pattern: projects/{project}/locations/{location}/lakes/{lake}
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s", p.ProjectID, p.Location, p.LakeID)
}

// NewTaskIdentity builds a TaskIdentity from the Config Connector Task object.
func NewTaskIdentity(ctx context.Context, reader client.Reader, obj *DataplexTask) (*TaskIdentity, error) {

	// Get Parent fields
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, fmt.Errorf("resolving projectRef: %w", err)
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("projectID is empty after resolving projectRef")
	}
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location is empty")
	}

	// Get LakeID from label (Assumption: label exists and is correctly populated)
	lakeID, ok := obj.Labels[LakeIDLabel]
	if !ok || lakeID == "" {
		// Attempt to extract from potential lakeRef if available (though not defined in provided types)
		// Since LakeRef is not available in DataplexTaskSpec, we rely solely on the label.
		// If this fails often, consider adding LakeRef to DataplexTaskSpec.
		return nil, fmt.Errorf("lake ID label '%s' not found or empty on DataplexTask %s/%s", LakeIDLabel, obj.Namespace, obj.Name)
	}

	// Get desired Resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot determine resource ID (metadata.name or spec.resourceID)")
	}

	// Validate against existing ExternalRef if present
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTaskExternal(externalRef)
		if err != nil {
			return nil, fmt.Errorf("parsing existing externalRef %q: %w", externalRef, err)
		}

		// Check immutable fields
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("project ID mismatch: spec.projectRef resolves to %q but existing externalRef has %q", projectID, actualParent.ProjectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("location mismatch: spec.location is %q but existing externalRef has %q", location, actualParent.Location)
		}
		if actualParent.LakeID != lakeID {
			// This validation relies on the label, which might not be immutable in KCC's view.
			// If the label can change, this might cause unexpected errors.
			return nil, fmt.Errorf("lake ID mismatch: label '%s' is %q but existing externalRef has %q", LakeIDLabel, lakeID, actualParent.LakeID)
		}

		// Check user-modifiable ID
		if actualResourceID != resourceID {
			// Allow ID change only if externalRef was based on the old name and resourceID is now set
			oldNameBasedID := obj.GetName() != resourceID && actualResourceID == obj.GetName() && common.IsManagedByKCC(obj.GetAnnotations())

			if !oldNameBasedID || common.ValueOf(obj.Spec.ResourceID) == "" {
				return nil, fmt.Errorf("resource ID mismatch: cannot change from %q (in externalRef) to %q (from metadata.name or spec.resourceID)",
					actualResourceID, resourceID)
			}
			// If we reach here, it means resourceID is set and differs from the name, and the externalRef matched the name.
			// This implies the user is setting resourceID for the first time on an existing resource. We allow this.
		}
	}

	return &TaskIdentity{
		parent: &TaskParent{
			ProjectID: projectID,
			Location:  location,
			LakeID:    lakeID,
		},
		id: resourceID,
	}, nil
}

// ParseTaskExternal parses the "external" format into its components.
// Expected format: projects/{{projectID}}/locations/{{location}}/lakes/{{lakeID}}/tasks/{{taskID}}
func ParseTaskExternal(external string) (parent *TaskParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expected pattern: projects/P/locations/L/lakes/LA/tasks/T (8 segments)
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "lakes" || tokens[6] != "tasks" {
		return nil, "", fmt.Errorf("format of DataplexTask externalRef %q was not known (expected projects/<project>/locations/<location>/lakes/<lake>/tasks/<task>)", external)
	}
	parent = &TaskParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		LakeID:    tokens[5], // Lake ID is at index 5
	}
	resourceID = tokens[7] // Task ID is at index 7
	return parent, resourceID, nil
}
