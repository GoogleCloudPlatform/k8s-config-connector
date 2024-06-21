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

package monitoring

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/references"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

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
		project, err := references.ResolveProject(ctx, reader, src, &refs.ProjectRef{
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

func normalizeMonitoringAlertPolicyRef(ctx context.Context, reader client.Reader, src client.Object, project references.Project, ref *krm.MonitoringAlertPolicyRef) (*krm.MonitoringAlertPolicyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on reference")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on reference")
	}

	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 2 && tokens[0] == "alertPolicies" {
			ref = &krm.MonitoringAlertPolicyRef{
				External: fmt.Sprintf("projects/%s/alertPolicies/%s", project.ProjectID, tokens[1]),
			}
		}
		if len(tokens) == 4 && tokens[0] == "project" && tokens[2] == "alertPolicies" {
			ref = &krm.MonitoringAlertPolicyRef{
				External: fmt.Sprintf("projects/%s/alertPolicies/%s", tokens[1], tokens[3]),
			}
		}
		return nil, fmt.Errorf("format of alertPolicyRef external=%q was not known (use projects/<projectId>/alertPolicies/<alertPolicyId> or alertPolicies/<alertPolicyId>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	alertPolicy := &unstructured.Unstructured{}
	alertPolicy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "monitoring.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MonitoringAlertPolicy",
	})
	if err := reader.Get(ctx, key, alertPolicy); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced MonitoringAlertPolicy %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced MonitoringAlertPolicy %v: %w", key, err)
	}

	alertPolicyResourceID, err := references.GetResourceID(alertPolicy)
	if err != nil {
		return nil, err
	}

	alertPolicyProjectID, err := references.ResolveProjectIDForObject(ctx, reader, alertPolicy)
	if err != nil {
		return nil, err
	}

	ref = &krm.MonitoringAlertPolicyRef{
		External: fmt.Sprintf("projects/%s/alertPolicies/%s", alertPolicyProjectID, alertPolicyResourceID),
	}

	return ref, nil
}

type refNormalizer struct {
	ctx     context.Context
	kube    client.Reader
	src     client.Object
	project references.Project
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
		if ref, err := normalizeMonitoringAlertPolicyRef(r.ctx, r.kube, r.src, r.project, alertChart.AlertPolicyRef); err != nil {
			return err
		} else {
			alertChart.AlertPolicyRef = ref
		}
	}

	return nil
}
