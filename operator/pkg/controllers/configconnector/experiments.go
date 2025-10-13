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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
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
	if cc.Spec.Experiments == nil {
		return nil
	}
	// so cc.Spec.Experiments  is non nil

	if cc.Spec.Experiments.LeaderElection != nil {
		// todo acpana we may want to validate the leader election spec here
		// may need to add the identity to the KCC manager
		
		if err := r.applyControllerManagerForcedRestart(ctx, m, "multi-cluster-lease updated"); err != nil {
			return err
		}
	}

	return nil
}

func (r *Reconciler) applyControllerManagerForcedRestart(ctx context.Context, obj *manifest.Objects, reason string) error {
	log := log.FromContext(ctx)
	for _, obj := range obj.Items {
		if obj.Kind != "StatefulSet" || !strings.HasPrefix(obj.GetName(), "cnrm-controller-manager") {
			continue
		}

		log.Info("force restarring 'cnrm-controller-manager' for reason %s", "name", obj.GetName(), reason)
		if err := obj.MutateContainers(func(container map[string]interface{}) error {
			name, found, _ := unstructured.NestedString(container, "name")
			if !found || name != "manager" {
				return nil
			}

			args, _, _ := unstructured.NestedStringSlice(container, "args")
			args = append(args, fmt.Sprintf("--force-restart=%s", reason))
			return unstructured.SetNestedStringSlice(container, args, "args")
		}); err != nil {
			return fmt.Errorf("failed to apply 'http-transport-metrics' experiment to %s: %w", obj.GetName(), err)
		}
	}

	return nil
}
