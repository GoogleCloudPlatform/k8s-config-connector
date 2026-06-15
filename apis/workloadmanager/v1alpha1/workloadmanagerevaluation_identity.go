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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &WorkloadManagerEvaluationIdentity{}
	_ identity.Resource   = &WorkloadManagerEvaluation{}
)

var WorkloadManagerEvaluationIdentityFormat = gcpurls.Template[WorkloadManagerEvaluationIdentity]("workloadmanager.googleapis.com", "projects/{project}/locations/{location}/evaluations/{evaluation}")

// +k8s:deepcopy-gen=false
type WorkloadManagerEvaluationIdentity struct {
	Project    string
	Location   string
	Evaluation string
}

func (i *WorkloadManagerEvaluationIdentity) String() string {
	return WorkloadManagerEvaluationIdentityFormat.ToString(*i)
}

func (i *WorkloadManagerEvaluationIdentity) FromExternal(ref string) error {
	parsed, match, err := WorkloadManagerEvaluationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of WorkloadManagerEvaluation external=%q was not known (use %s): %w", ref, WorkloadManagerEvaluationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of WorkloadManagerEvaluation external=%q was not known (use %s)", ref, WorkloadManagerEvaluationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *WorkloadManagerEvaluationIdentity) Host() string {
	return WorkloadManagerEvaluationIdentityFormat.Host()
}

func getIdentityFromWorkloadManagerEvaluationSpec(ctx context.Context, reader client.Reader, obj client.Object) (*WorkloadManagerEvaluationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := resolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &WorkloadManagerEvaluationIdentity{
		Project:    projectID,
		Location:   location,
		Evaluation: resourceID,
	}
	return identity, nil
}

func resolveProjectID(ctx context.Context, reader client.Reader, obj client.Object) (string, error) {
	if cc, ok := obj.(*WorkloadManagerEvaluation); ok {
		if cc.Spec.ProjectRef != nil {
			project, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), cc.Spec.ProjectRef)
			if err != nil {
				return "", err
			}
			return project.ProjectID, nil
		}
	} else if u, ok := obj.(*unstructured.Unstructured); ok {
		return refs.ResolveProjectID(ctx, reader, u)
	}
	return "", fmt.Errorf("projectRef is required")
}

func (obj *WorkloadManagerEvaluation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromWorkloadManagerEvaluationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &WorkloadManagerEvaluationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change WorkloadManagerEvaluation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
