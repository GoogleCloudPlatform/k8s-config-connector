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

package configconnector

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (r *Reconciler) transformForExperiments() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		cc, ok := o.(*corev1beta1.ConfigConnector)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was of type %T", o)
		}

		if err := r.applyExperiments(ctx, cc, m); err != nil {
			return fmt.Errorf("error applying experiment transforms: %w", err)
		}
		return nil
	}
}

func (r *Reconciler) applyExperiments(ctx context.Context, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	log := log.FromContext(ctx)

	for _, experiment := range cc.Spec.Experiments {
		key := strings.ToLower(experiment)
		switch key {
		case "legacy-iam-reconciler":
			if err := r.applyExperimentLegacyIAMReconciler(ctx, cc, m); err != nil {
				return err
			}

		default:
			log.Info("ignoring unknown experiment", "key", key)
			// TODO: add to status?
		}
	}

	return nil
}

func (r *Reconciler) applyExperimentLegacyIAMReconciler(ctx context.Context, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	log := log.FromContext(ctx)
	for _, obj := range m.Items {
		if obj.Kind != "StatefulSet" || !strings.HasPrefix(obj.GetName(), "cnrm-controller-manager") {
			continue
		}

		log.Info("applying 'legacy-iam-reconciler' experiment to StatefulSet", "name", obj.GetName())

		if err := obj.MutateContainers(func(container map[string]interface{}) error {
			name, found, _ := unstructured.NestedString(container, "name")
			if !found || name != "manager" {
				return nil
			}

			args, _, _ := unstructured.NestedStringSlice(container, "args")
			args = append(args, "--legacy-iam-reconciler=true")
			return unstructured.SetNestedStringSlice(container, args, "args")
		}); err != nil {
			return fmt.Errorf("failed to apply 'legacy-iam-reconciler' experiment to %s: %w", obj.GetName(), err)
		}
	}
	return nil
}
