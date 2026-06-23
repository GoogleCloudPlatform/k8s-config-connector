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

var (
	_ identity.IdentityV2 = &TaskIdentity{}
	_ identity.Resource   = &BatchTask{}
)

var TaskIdentityFormat = gcpurls.Template[TaskIdentity]("batch.googleapis.com", "projects/{project}/locations/{location}/jobs/{job}/taskGroups/{taskgroup}/tasks/{task}")

// +k8s:deepcopy-gen=false
type TaskIdentity struct {
	Project   string
	Location  string
	Job       string
	TaskGroup string
	Task      string
}

func (i *TaskIdentity) String() string {
	return TaskIdentityFormat.ToString(*i)
}

func (i *TaskIdentity) FromExternal(ref string) error {
	parsed, match, err := TaskIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BatchTask external=%q was not known (use %s): %w", ref, TaskIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BatchTask external=%q was not known (use %s)", ref, TaskIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TaskIdentity) Host() string {
	return TaskIdentityFormat.Host()
}

func (i *TaskIdentity) ExternalIdentifier() *string {
	return &i.Task
}

func getIdentityFromBatchTaskSpec(ctx context.Context, reader client.Reader, obj *BatchTask) (*TaskIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if err := obj.Spec.JobRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot resolve JobRef: %w", err)
	}

	jobIdentity := &JobIdentity{}
	if err := jobIdentity.FromExternal(obj.Spec.JobRef.External); err != nil {
		return nil, fmt.Errorf("invalid JobRef: %w", err)
	}

	identity := &TaskIdentity{
		Project:   projectRef.ProjectID,
		Location:  obj.Spec.Location,
		Job:       jobIdentity.id,
		TaskGroup: obj.Spec.TaskGroup,
		Task:      resourceID,
	}
	return identity, nil
}

func (obj *BatchTask) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBatchTaskSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &TaskIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BatchTask identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
