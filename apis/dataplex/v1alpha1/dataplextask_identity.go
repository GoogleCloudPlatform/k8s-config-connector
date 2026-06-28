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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &TaskIdentity{}
	_ identity.Resource   = &DataplexTask{}
)

var TaskIdentityFormat = gcpurls.Template[TaskIdentity]("dataplex.googleapis.com", "projects/{project}/locations/{location}/lakes/{lake}/tasks/{task}")

// +k8s:deepcopy-gen=false
type TaskIdentity struct {
	Project  string
	Location string
	Lake     string
	Task     string
}

func (i *TaskIdentity) String() string {
	return TaskIdentityFormat.ToString(*i)
}

func (i *TaskIdentity) FromExternal(ref string) error {
	parsed, match, err := TaskIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataplexTask external=%q was not known (use %s): %w", ref, TaskIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataplexTask external=%q was not known (use %s)", ref, TaskIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TaskIdentity) Host() string {
	return TaskIdentityFormat.Host()
}

func getIdentityFromDataplexTaskSpec(ctx context.Context, reader client.Reader, obj *DataplexTask) (*TaskIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	lakeRef := obj.Spec.LakeRef
	if lakeRef == nil {
		return nil, fmt.Errorf("LakeRef is required")
	}

	if err := lakeRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot normalize LakeRef: %w", err)
	}

	lakeIdentity := &LakeIdentity{}
	if err := lakeIdentity.FromExternal(lakeRef.External); err != nil {
		return nil, fmt.Errorf("cannot parse LakeRef external: %w", err)
	}

	identity := &TaskIdentity{
		Project:  lakeIdentity.Project,
		Location: lakeIdentity.Location,
		Lake:     lakeIdentity.Lake,
		Task:     resourceID,
	}
	return identity, nil
}

func (obj *DataplexTask) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataplexTaskSpec(ctx, reader, obj)
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
			return nil, fmt.Errorf("cannot change DataplexTask identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
