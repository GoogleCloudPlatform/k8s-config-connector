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

package preview

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	k8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func newReconcilerOverrideTransformer(namespace string, reconcilerOverride map[schema.GroupKind]k8s.ReconcilerType) ObjectTransformer {
	return func(ctx context.Context, obj client.Object) error {
		if ccc, ok := obj.(*corev1beta1.ConfigConnectorContext); ok {
			if namespace != "" && ccc.GetNamespace() != namespace {
				return nil
			}
			if len(reconcilerOverride) == 0 {
				return nil
			}
			if ccc.Spec.Experiments == nil {
				ccc.Spec.Experiments = &corev1beta1.Experiments{}
			}
			if ccc.Spec.Experiments.ControllerOverrides == nil {
				ccc.Spec.Experiments.ControllerOverrides = make(map[string]k8s.ReconcilerType)
			}
			for gk, rt := range reconcilerOverride {
				ccc.Spec.Experiments.ControllerOverrides[gk.String()] = rt
			}
		}
		return nil
	}
}
