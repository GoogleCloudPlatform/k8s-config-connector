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

package predicate

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// AnnotationKeyAlphaReconciler allows customers to opt-in to using the direct reconciler.
const AnnotationKeyAlphaReconciler = "alpha.cnrm.cloud.google.com/reconciler"

// ChooseReconciler allows users to opt in/out of direct reconciliation
// by specifying an AnnotationKeyAlphaReconciler annotation.
type ChooseReconciler struct {
}

var _ ReconcileGate = &ChooseReconciler{}

// ShouldReconcile returns true if the reconciler should be used to for the resource.
func (r *ChooseReconciler) ShouldReconcile(o *unstructured.Unstructured) bool {
	v := o.GetAnnotations()[AnnotationKeyAlphaReconciler]
	gvkLabels := supportedgvks.SupportedGVKs[o.GroupVersionKind()].Labels

	defaultLegacy := gvkLabels[k8s.TF2CRDLabel] == "true" || gvkLabels[k8s.DCL2CRDLabel] == "true"
	defaultDirect := gvkLabels[k8s.LegacyControllerLabel] != ""

	if defaultLegacy {
		// If default controller is legacy, allow opt-in to direct controller.
		return v == "direct"
	} else if defaultDirect {
		// If default controller is direct (with a known legacy controller), allow opt-out of direct controller.
		return v != "legacy"
	} else {
		// If no known legacy controller, always use direct controller.
		return true
	}
}
