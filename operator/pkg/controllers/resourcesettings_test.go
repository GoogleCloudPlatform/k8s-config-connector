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

package controllers

import (
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
)

func TestComputeResourceSettingsHash(t *testing.T) {
	// Both nil
	hash, err := ComputeResourceSettingsHash(nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash != "" {
		t.Errorf("expected empty hash for nil settings, got %q", hash)
	}

	// CC settings only
	ccSettings := &corev1beta1.ResourceSettings{
		Mode: corev1beta1.ResourceSettingsModeExclude,
		Resources: []corev1beta1.ResourceFilter{
			{Group: ptr.To("pubsub.cnrm.cloud.google.com")},
		},
	}
	hash1, err := ComputeResourceSettingsHash(ccSettings, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash1 == "" {
		t.Errorf("expected non-empty hash, got empty")
	}

	// CCC settings only
	cccSettings := &corev1beta1.ResourceSettings{
		Mode: corev1beta1.ResourceSettingsModeExclude,
		Resources: []corev1beta1.ResourceFilter{
			{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
		},
	}
	hash2, err := ComputeResourceSettingsHash(nil, cccSettings)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash2 == "" || hash2 == hash1 {
		t.Errorf("expected distinct non-empty hash, got %q", hash2)
	}

	// Both CC and CCC settings
	hash3, err := ComputeResourceSettingsHash(ccSettings, cccSettings)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash3 == "" || hash3 == hash1 || hash3 == hash2 {
		t.Errorf("expected distinct non-empty combined hash, got %q", hash3)
	}

	// Determinism
	hash3Again, err := ComputeResourceSettingsHash(ccSettings, cccSettings)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hash3 != hash3Again {
		t.Errorf("expected deterministic hash, got %q vs %q", hash3, hash3Again)
	}
}

func TestApplyResourceSettingsHashToPodTemplate(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"spec": map[string]interface{}{
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"existing-annotation": "true",
						},
					},
				},
			},
		},
	}

	ccSettings := &corev1beta1.ResourceSettings{
		Mode: corev1beta1.ResourceSettingsModeInclude,
		Resources: []corev1beta1.ResourceFilter{
			{Group: ptr.To("iam.cnrm.cloud.google.com"), Kind: ptr.To("IAMServiceAccount")},
		},
	}

	if err := ApplyResourceSettingsHashToPodTemplate(u, ccSettings, nil); err != nil {
		t.Fatalf("unexpected error applying hash: %v", err)
	}

	annotations, _, err := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if err != nil {
		t.Fatalf("unexpected error getting annotations: %v", err)
	}
	if annotations["existing-annotation"] != "true" {
		t.Errorf("expected existing annotation preserved")
	}
	if annotations[ResourceSettingsHashAnnotation] == "" {
		t.Errorf("expected %s annotation to be set", ResourceSettingsHashAnnotation)
	}

	// Setting to nil removes annotation
	if err := ApplyResourceSettingsHashToPodTemplate(u, nil, nil); err != nil {
		t.Fatalf("unexpected error removing hash: %v", err)
	}
	annotations, _, _ = unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if _, exists := annotations[ResourceSettingsHashAnnotation]; exists {
		t.Errorf("expected %s annotation to be removed when settings are nil", ResourceSettingsHashAnnotation)
	}
}
