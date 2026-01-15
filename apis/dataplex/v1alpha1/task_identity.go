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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &TaskIdentity{}

const (
	TaskIDURL = LakeIDURL + "/tasks/{{taskID}}"
)

// TaskIdentity defines the resource reference to DataplexTask, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type TaskIdentity struct {
	parent *LakeIdentity
	id     string
}

func (i *TaskIdentity) String() string {
	// Uses the pattern: projects/{project}/locations/{location}/lakes/{lake}/tasks/{task}
	return i.parent.String() + "/tasks/" + i.id
}

func (i *TaskIdentity) ID() string {
	return i.id
}

func (i *TaskIdentity) Parent() *LakeIdentity {
	return i.parent
}

func (i *TaskIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/tasks/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of DataplexTask external=%q was not known (use %s)", ref, TaskIDURL)
	}
	i.parent = &LakeIdentity{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("taskID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &DataplexTask{}

func (obj *DataplexTask) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &TaskIdentity{
		parent: &LakeIdentity{
			parent: &parent.ProjectAndLocationParent{},
		},
	}

	// Resolve Parent
	if err := obj.Spec.LakeRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	if err := newIdentity.parent.FromExternal(obj.Spec.LakeRef.External); err != nil {
		return nil, err
	}

	// Get desired Resource ID
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" {
		newIdentity.id = obj.GetName()
	}
	if newIdentity.id == "" {
		return nil, fmt.Errorf("cannot determine resource ID (metadata.name or spec.resourceID)")
	}

	// Validate against existing ExternalRef if present
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &TaskIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}

	return newIdentity, nil
}
