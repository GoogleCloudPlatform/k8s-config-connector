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

package configconnectorcontext

import (
	"context"
	"fmt"
	"strings"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func (r *Reconciler) transformForExperiments() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was of type %T", o)
		}

		if ccc.Spec.Experiments == nil {
			return nil
		}

		if ccc.Spec.Experiments.ResourceNameLabelMetrics != nil && *ccc.Spec.Experiments.ResourceNameLabelMetrics {
			if err := r.applyResourceNameMetrics(ctx, ccc, m); err != nil {
				return err
			}
		}

		return nil
	}
}

func (r *Reconciler) applyResourceNameMetrics(_ context.Context, _ *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	for _, obj := range m.Items {
		if obj.Kind == "StatefulSet" && strings.Contains(obj.GetName(), "cnrm-controller-manager-") {
			containers, found, err := unstructured.NestedSlice(obj.UnstructuredObject().Object, "spec", "template", "spec", "containers")
			if err != nil {
				return fmt.Errorf("error getting containers from statefulset: %w", err)
			}
			if !found {
				return fmt.Errorf("containers not found in statefulset")
			}
			for i, container := range containers {
				containerMap, ok := container.(map[string]interface{})
				if !ok {
					return fmt.Errorf("container is not a map")
				}
				if containerMap["name"] == "manager" {
					args, _, err := unstructured.NestedStringSlice(containerMap, "args")
					if err != nil {
						return fmt.Errorf("error getting args from container: %w", err)
					}
					args = append(args, "--resource-name-label")
					if err := unstructured.SetNestedStringSlice(containerMap, args, "args"); err != nil {
						return fmt.Errorf("error setting args in container: %w", err)
					}
					containers[i] = containerMap
					if err := unstructured.SetNestedSlice(obj.UnstructuredObject().Object, containers, "spec", "template", "spec", "containers"); err != nil {
						return fmt.Errorf("error setting containers in statefulset: %w", err)
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("cnrm-controller-manager StatefulSet or manager container not found in manifest")
}
