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

package common

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NormalizeReferences(ctx context.Context, reader client.Reader, obj client.Object, projectRef *refs.Project) error {
	if err := VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: projectRef, kube: reader}); err != nil {
		return err
	}
	return nil
}

func normalizeResourceName(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	if ref == nil {
		return nil, nil
	}

	// For backwards compatibility, infer "Project" kind
	if ref.Kind == "" && ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			ref.Kind = "Project"
		}
	}

	if ref.Kind == "" {
		return nil, fmt.Errorf("must specify kind on reference (%+v)", ref)
	}
	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on reference")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on reference")
	}

	switch ref.Kind {
	case "Project":
		project, err := refs.ResolveProject(ctx, reader, src.GetNamespace(), &refs.ProjectRef{
			Name:      ref.Name,
			Namespace: ref.Namespace,
			External:  ref.External,
			Kind:      ref.Kind,
		})
		if err != nil {
			return nil, err
		}

		ref = &v1alpha1.ResourceRef{
			Kind:     ref.Kind,
			External: fmt.Sprintf("projects/%s", project.ProjectID),
		}

	default:
		return nil, fmt.Errorf("references to kind %q are not supported", ref.Kind)
	}

	tokens := strings.Split(ref.External, "/")
	switch ref.Kind {
	case "Project":
		if len(tokens) == 2 && tokens[0] == "projects" {
			// OK
		} else {
			return nil, fmt.Errorf("resourceName %q should be in the format projects/<projectId>", ref.External)
		}
	default:
		return nil, fmt.Errorf("references to kind %q are not supported", ref.Kind)
	}

	return ref, nil
}

func normalizeProjectRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ProjectRef) (*refs.ProjectRef, error) {
	if ref == nil {
		return nil, nil
	}

	project, err := refs.ResolveProject(ctx, reader, src.GetNamespace(), ref)
	if err != nil {
		return nil, err
	}

	return &refs.ProjectRef{
		External: "projects/" + project.ProjectID,
	}, nil
}

type refNormalizer struct {
	ctx     context.Context
	kube    client.Reader
	src     client.Object
	project *refs.Project
}

func (r *refNormalizer) VisitField(path string, v any) error {
	if logsPanel, ok := v.(*krm.LogsPanel); ok {
		for i := range logsPanel.ResourceNames {
			if ref, err := normalizeResourceName(r.ctx, r.kube, r.src, &logsPanel.ResourceNames[i]); err != nil {
				return err
			} else {
				logsPanel.ResourceNames[i] = *ref
			}
		}
	}
	if alertChart, ok := v.(*krm.AlertChart); ok {
		if r.project == nil {
			return fmt.Errorf("must specify project for alertChart references")
		}
		external, err := alertChart.AlertPolicyRef.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace())
		if err != nil {
			return err
		}
		refined, err := r.RefineMonitoringAlertPolicyRefForAlertChart(external)
		if err != nil {
			return err
		}
		alertChart.AlertPolicyRef.External = refined
	}

	if alertChart, ok := v.(*krm.IncidentList); ok {
		for i, policyRef := range alertChart.PolicyRefs {
			if r.project == nil {
				return fmt.Errorf("must specify project for policyRef references")
			}
			external, err := policyRef.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace())
			if err != nil {
				return err
			}
			refined, err := r.RefineMonitoringAlertPolicyRefForIncidentList(external)
			if err != nil {
				return err
			}
			alertChart.PolicyRefs[i].External = refined
		}
	}

	if projectRef, ok := v.(*refs.ProjectRef); ok {
		if ref, err := normalizeProjectRef(r.ctx, r.kube, r.src, projectRef); err != nil {
			return err
		} else if ref != nil {
			*projectRef = *ref
		}
	}

	return nil
}

// RefineMonitoringAlertPolicyRefForAlertChart refine the alertPolicy format because MonitoringDashboard's AlertChart has a specific format requirement:
// "The format is: projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]"
func (r *refNormalizer) RefineMonitoringAlertPolicyRefForAlertChart(external string) (string, error) {
	// External formats could be alertPolicies/{{alertPolicyId}} or projects/{{projectId}}/alertPolicies/{{alertPolicyId}}
	// Verify project and refine `alertPolicies/{{alertPolicyId}}` to the required format `projects/{{projectId}}/alertPolicies/{{alertPolicyId}}`
	tokens := strings.Split(external, "/")
	projectID := r.project.ProjectID
	// alertPolicies/{{alertPolicyId}}
	if len(tokens) == 2 && tokens[0] == "alertPolicies" {
		return fmt.Sprintf("projects/%s/%s", projectID, external), nil
	} else {
		// projects/{{projectId}}/alertPolicies/{{alertPolicyId}}
		if tokens[1] != projectID {
			return "", fmt.Errorf("resolve alertPolicy (%q) in incidentList was not in same project", external)
		}
	}
	return external, nil
}

// RefineMonitoringAlertPolicyRefForIncidentList refine the alertPolicy format because MonitoringDashboard's IncidentList has a specific format requirement:
// "Don't include the project ID prefix in the policy name. For example, use `alertPolicies/utilization`."
func (r *refNormalizer) RefineMonitoringAlertPolicyRefForIncidentList(external string) (string, error) {
	// External formats could be alertPolicies/{{alertPolicyId}} or projects/{{projectId}}/alertPolicies/{{alertPolicyId}}
	// Verify project and refine `projects/{{projectId}}/alertPolicies/{{alertPolicyId}}` to the required format `alertPolicies/{{alertPolicyId}}`
	tokens := strings.Split(external, "/")
	projectID := r.project.ProjectID
	// alertPolicies/{{alertPolicyId}}
	if len(tokens) == 2 && tokens[0] == "alertPolicies" {
		return external, nil
	} else {
		// projects/{{projectId}}/alertPolicies/{{alertPolicyId}}
		if tokens[1] != projectID {
			return "", fmt.Errorf("resolve alertPolicy (%q) in incidentList was not in same project", external)
		}
	}
	prefix := fmt.Sprintf("projects/%s/", r.project.ProjectID)
	return strings.TrimPrefix(external, prefix), nil
}
