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

var _ identity.Identity = &WorkflowsWorkflowIdentity{}

const (
	WorkflowsWorkflowIdentityURL = parent.ProjectAndLocationURL + "/workflows/{{workflowID}}"
)

// WorkflowsWorkflowIdentity defines the resource reference to WorkflowsWorkflow, which "External" field
// holds the GCP identifier for the KRM object.
type WorkflowsWorkflowIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *WorkflowsWorkflowIdentity) String() string {
	return i.parent.String() + "/workflows/" + i.id
}

func (i *WorkflowsWorkflowIdentity) ID() string {
	return i.id
}

func (i *WorkflowsWorkflowIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *WorkflowsWorkflowIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/workflows/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of WorkflowsWorkflow external=%q was not known (use projects/{{projectID}}/locations/{{location}}/workflows/{{workflowID}})", ref)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("workflowID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &WorkflowsWorkflow{}

func (obj *WorkflowsWorkflow) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &WorkflowsWorkflowIdentity{
		parent: &parent.ProjectAndLocationParent{},
	}

	// Resolve user-configured Parent
	if err := obj.Spec.ProjectAndLocationRef.Build(ctx, reader, obj.GetNamespace(), newIdentity.parent); err != nil {
		return nil, err
	}

	// Get user-configured ID
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" {
		newIdentity.id = obj.GetName()
	}
	if newIdentity.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &WorkflowsWorkflowIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
