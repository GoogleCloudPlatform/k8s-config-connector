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
	if cc.Spec.Experiments != nil && cc.Spec.Experiments.ControllersOverrides != nil {
		if err := r.applyControllerManagerExperiments(ctx, cc, m); err != nil {
			return err
		}
	}

	return nil
}

func (r *Reconciler) applyControllerManagerExperiments(ctx context.Context, cc *corev1beta1.ConfigConnector, obj *manifest.Objects) error {
	log := log.FromContext(ctx)
	for _, obj := range obj.Items {
		if obj.Kind != "StatefulSet" || !strings.HasPrefix(obj.GetName(), "cnrm-controller-manager") {
			continue
		}

		log.Info("applying 'controller-overrides' experiment to StatefulSet", "name", obj.GetName())
		for _, override := range cc.Spec.Experiments.ControllersOverrides {
			switch override.Name {
			case "IAMPartialPolicy":
				switch override.Value {
				case "direct":
					if err := obj.MutateContainers(func(container map[string]interface{}) error {
						name, found, _ := unstructured.NestedString(container, "name")
						if !found || name != "manager" {
							return nil
						}

						args, _, _ := unstructured.NestedStringSlice(container, "args")
						args = append(args, "--direct-iam-reconciler=true")
						return unstructured.SetNestedStringSlice(container, args, "args")
					}); err != nil {
						return fmt.Errorf("failed to apply 'direct-iam-reconciler' flag as part of the `controller-overrides` experiment to %s: %w", obj.GetName(), err)
					}
				case "":
					// field is present but empty, do nothing
				default:
					return fmt.Errorf("unknown value %q for controller override on kind %q", override.Value, override.Name)
				}
			default:
				return fmt.Errorf("unknown kind %q for controller override", override.Name)
			}
		}
	}

	return nil
}
