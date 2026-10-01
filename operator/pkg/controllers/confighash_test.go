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
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
)

func TestComputeCCConfigHash(t *testing.T) {
	t.Run("nil or unset ResourceSettings returns empty string", func(t *testing.T) {
		if got := ComputeCCConfigHash(nil); got != "" {
			t.Errorf("expected empty hash for nil CC, got %q", got)
		}
		cc := &corev1beta1.ConfigConnector{}
		if got := ComputeCCConfigHash(cc); got != "" {
			t.Errorf("expected empty hash for CC without experiments, got %q", got)
		}
	})

	t.Run("order and duplicate invariance", func(t *testing.T) {
		cc1 := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Mode: corev1beta1.ResourceSettingsModeInclude,
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
							{Group: ptr.To("pubsub.cnrm.cloud.google.com"), Kind: ptr.To("PubSubTopic")},
						},
					},
				},
			},
		}
		cc2 := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Mode: corev1beta1.ResourceSettingsModeInclude,
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("pubsub.cnrm.cloud.google.com"), Kind: ptr.To("PubSubTopic")},
							{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
							{Group: ptr.To("pubsub.cnrm.cloud.google.com"), Kind: ptr.To("PubSubTopic")},
						},
					},
				},
			},
		}

		h1 := ComputeCCConfigHash(cc1)
		h2 := ComputeCCConfigHash(cc2)
		if h1 == "" {
			t.Fatal("expected non-empty hash")
		}
		if h1 != h2 {
			t.Errorf("expected identical hashes for reordered/duplicate resources, got %q vs %q", h1, h2)
		}
	})

	t.Run("default mode exclude matches explicit mode exclude", func(t *testing.T) {
		ccImplicit := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("storage.cnrm.cloud.google.com")},
						},
					},
				},
			},
		}
		ccExplicit := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Mode: corev1beta1.ResourceSettingsModeExclude,
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("storage.cnrm.cloud.google.com")},
						},
					},
				},
			},
		}
		if ComputeCCConfigHash(ccImplicit) != ComputeCCConfigHash(ccExplicit) {
			t.Errorf("expected implicit and explicit exclude mode to produce identical hash")
		}
	})

	t.Run("mode transition produces distinct hash", func(t *testing.T) {
		ccExclude := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Mode: corev1beta1.ResourceSettingsModeExclude,
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
						},
					},
				},
			},
		}
		ccInclude := &corev1beta1.ConfigConnector{
			Spec: corev1beta1.ConfigConnectorSpec{
				Experiments: &corev1beta1.CCExperiments{
					ResourceSettings: &corev1beta1.ResourceSettings{
						Mode: corev1beta1.ResourceSettingsModeInclude,
						Resources: []corev1beta1.ResourceFilter{
							{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
						},
					},
				},
			},
		}
		if ComputeCCConfigHash(ccExclude) == ComputeCCConfigHash(ccInclude) {
			t.Errorf("expected exclude and include modes to produce different hashes")
		}
	})
}

func TestComputeCCCConfigHash(t *testing.T) {
	ccc1 := &corev1beta1.ConfigConnectorContext{
		Spec: corev1beta1.ConfigConnectorContextSpec{
			Experiments: &corev1beta1.Experiments{
				ResourceSettings: &corev1beta1.ResourceSettings{
					Mode: corev1beta1.ResourceSettingsModeExclude,
					Resources: []corev1beta1.ResourceFilter{
						{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
						{Group: ptr.To("bigquery.cnrm.cloud.google.com"), Kind: ptr.To("BigQueryDataset")},
					},
				},
			},
		},
	}
	ccc2 := &corev1beta1.ConfigConnectorContext{
		Spec: corev1beta1.ConfigConnectorContextSpec{
			Experiments: &corev1beta1.Experiments{
				ResourceSettings: &corev1beta1.ResourceSettings{
					Mode: corev1beta1.ResourceSettingsModeExclude,
					Resources: []corev1beta1.ResourceFilter{
						{Group: ptr.To("bigquery.cnrm.cloud.google.com"), Kind: ptr.To("BigQueryDataset")},
						{Group: ptr.To("storage.cnrm.cloud.google.com"), Kind: ptr.To("StorageBucket")},
					},
				},
			},
		},
	}
	if ComputeCCCConfigHash(ccc1) != ComputeCCCConfigHash(ccc2) {
		t.Errorf("expected identical CCC hashes for reordered resources")
	}
}

func TestSetPodTemplateAnnotation(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
		},
	}

	if err := SetPodTemplateAnnotation(u, k8s.CCConfigHashAnnotation, "hash-cc-1"); err != nil {
		t.Fatalf("unexpected error setting CC hash: %v", err)
	}
	if err := SetPodTemplateAnnotation(u, k8s.CCCConfigHashAnnotation, "hash-ccc-1"); err != nil {
		t.Fatalf("unexpected error setting CCC hash: %v", err)
	}

	annotations, _, _ := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if annotations[k8s.CCConfigHashAnnotation] != "hash-cc-1" {
		t.Errorf("expected %s=hash-cc-1, got %q", k8s.CCConfigHashAnnotation, annotations[k8s.CCConfigHashAnnotation])
	}
	if annotations[k8s.CCCConfigHashAnnotation] != "hash-ccc-1" {
		t.Errorf("expected %s=hash-ccc-1, got %q", k8s.CCCConfigHashAnnotation, annotations[k8s.CCCConfigHashAnnotation])
	}

	// Remove CC hash while preserving CCC hash
	if err := SetPodTemplateAnnotation(u, k8s.CCConfigHashAnnotation, ""); err != nil {
		t.Fatalf("unexpected error removing CC hash: %v", err)
	}
	annotations, _, _ = unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if _, exists := annotations[k8s.CCConfigHashAnnotation]; exists {
		t.Errorf("expected %s to be removed", k8s.CCConfigHashAnnotation)
	}
	if annotations[k8s.CCCConfigHashAnnotation] != "hash-ccc-1" {
		t.Errorf("expected %s=hash-ccc-1 to remain intact, got %q", k8s.CCCConfigHashAnnotation, annotations[k8s.CCCConfigHashAnnotation])
	}
}
