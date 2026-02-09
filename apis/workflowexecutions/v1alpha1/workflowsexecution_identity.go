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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	workflow "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &WorkflowsExecutionIdentity{}
	_ identity.Resource   = &WorkflowsExecution{}
)

var WorkflowsExecutionIdentityFormat = gcpurls.Template[WorkflowsExecutionIdentity]("workflowexecutions.googleapis.com", "projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}")

// +k8s:deepcopy-gen=false
type WorkflowsExecutionIdentity struct {
	Project   string
	Location  string
	Workflow  string
	Execution string
}

func (i *WorkflowsExecutionIdentity) String() string {
	return WorkflowsExecutionIdentityFormat.ToString(*i)
}

func (i *WorkflowsExecutionIdentity) FromExternal(ref string) error {
	parsed, match, err := WorkflowsExecutionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of WorkflowsExecution external=%q was not known (use %s): %w", ref, WorkflowsExecutionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of WorkflowsExecution external=%q was not known (use %s)", ref, WorkflowsExecutionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *WorkflowsExecutionIdentity) Host() string {
	return WorkflowsExecutionIdentityFormat.Host()
}

func getIdentityFromWorkflowsExecutionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*WorkflowsExecutionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	var workflowRef *workflow.WorkflowsWorkflowRef
	if u, ok := obj.(*WorkflowsExecution); ok {
		if u.Spec.WorkflowExecutionParent != nil {
			workflowRef = u.Spec.WorkflowRef
		}
	} else {
		unstr := obj.(*unstructured.Unstructured)
		if v, ok, _ := unstructured.NestedMap(unstr.Object, "spec", "workflowRef"); ok {
			workflowRef = &workflow.WorkflowsWorkflowRef{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(v, workflowRef); err != nil {
				return nil, fmt.Errorf("error converting workflowRef: %w", err)
			}
		}
	}

	if workflowRef == nil {
		return nil, fmt.Errorf("cannot resolve workflowRef")
	}

	workflowExternal, err := workflowRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	_, workflowID, err := workflow.ParseWorkflowsWorkflowExternal(workflowExternal)
	if err != nil {
		return nil, err
	}

	identity := &WorkflowsExecutionIdentity{
		Project:   projectID,
		Location:  location,
		Workflow:  workflowID,
		Execution: resourceID,
	}
	return identity, nil
}

func (obj *WorkflowsExecution) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromWorkflowsExecutionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &WorkflowsExecutionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change WorkflowsExecution identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}