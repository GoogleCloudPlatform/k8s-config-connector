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

package configconnectorcontext

import (
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
)

func TestApplyResourceSettingsHashInNamespacedMode(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"spec": map[string]interface{}{
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{},
					},
				},
			},
		},
	}

	cc := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Experiments: &corev1beta1.CCExperiments{
				ResourceSettings: &corev1beta1.ResourceSettings{
					Mode: corev1beta1.ResourceSettingsModeExclude,
					Resources: []corev1beta1.ResourceFilter{
						{Group: ptr.To("pubsub.cnrm.cloud.google.com")},
					},
				},
			},
		},
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		Spec: corev1beta1.ConfigConnectorContextSpec{
			Experiments: &corev1beta1.Experiments{
				ResourceSettings: &corev1beta1.ResourceSettings{
					Mode: corev1beta1.ResourceSettingsModeExclude,
					Resources: []corev1beta1.ResourceFilter{
						{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
					},
				},
			},
		},
	}

	if err := applyExperimentsToManagerContainer(u, cc, ccc); err != nil {
		t.Fatalf("applyExperimentsToManagerContainer failed: %v", err)
	}

	templateAnnotations, _, err := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if err != nil {
		t.Fatalf("failed to get pod template annotations: %v", err)
	}

	hash := templateAnnotations[controllers.ResourceSettingsHashAnnotation]
	if hash == "" {
		t.Fatalf("expected %s annotation to be set, but was empty", controllers.ResourceSettingsHashAnnotation)
	}

	// Change CCC settings and verify hash changes
	ccc.Spec.Experiments.ResourceSettings.Resources = append(ccc.Spec.Experiments.ResourceSettings.Resources, corev1beta1.ResourceFilter{
		Group: ptr.To("compute.cnrm.cloud.google.com"),
	})

	if err := applyExperimentsToManagerContainer(u, cc, ccc); err != nil {
		t.Fatalf("applyExperimentsToManagerContainer with updated CCC failed: %v", err)
	}

	templateAnnotations, _, _ = unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	newHash := templateAnnotations[controllers.ResourceSettingsHashAnnotation]
	if newHash == "" || newHash == hash {
		t.Fatalf("expected different hash after updating CCC, got %q vs %q", newHash, hash)
	}
}
