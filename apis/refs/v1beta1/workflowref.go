// Copyright 2024 Google LLC
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
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WorkflowRef struct {
	/* The StorageBucket selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `StorageBucket` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `StorageBucket` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type Workflow struct {
	ProjectID    string
	Location     string
	WorkflowName string
}

func (s *Workflow) String() string {
	return "projects/" + s.ProjectID + "locations/" + s.Location + "/workflows/" + s.WorkflowName
}

func ResolveWorkflowRef(ctx context.Context, reader client.Reader, obj client.Object, ref *WorkflowRef) (*Workflow, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on workflowRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.workflowRef.name and spec.workflowRef.external")
	}

	if ref.External != "" {
		// External should be in the `projects/[projectID]/locations/[Location]/instances/[instanceName]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workflows" {
			return &Workflow{
				ProjectID:    tokens[1],
				Location:     tokens[3],
				WorkflowName: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of workflow external=%q was not known (use projects/<projectId>/locations/[Location]/workflows/<workflowName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	workflow := &unstructured.Unstructured{}
	workflow.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "workflow.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "WorkflowsWorkflow",
	})
	if err := reader.Get(ctx, key, workflow); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced workflow %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced workflow %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(workflow.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from workflow %s/%s: %w", workflow.GetNamespace(), workflow.GetName(), err)
	}
	if resourceID == "" {
		resourceID = workflow.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, workflow)
	if err != nil {
		return nil, err
	}

	location, _, err := unstructured.NestedString(workflow.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from workflow %s/%s: %w", workflow.GetNamespace(), workflow.GetName(), err)
	}

	return &Workflow{
		ProjectID:    projectID,
		Location:     location,
		WorkflowName: resourceID,
	}, nil
}
