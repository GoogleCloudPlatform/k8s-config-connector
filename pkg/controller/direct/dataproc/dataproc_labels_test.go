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

package dataproc

import (
	"context"
	"net/http"
	"testing"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataprocSession_KRMPrefixedLabels(t *testing.T) {
	ctx := context.Background()

	session := &krmv1alpha1.DataprocSession{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krmv1alpha1.DataprocSessionGVK.GroupVersion().String(),
			Kind:       krmv1alpha1.DataprocSessionGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-session",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
				"env":                          "prod",
			},
		},
		Spec: krmv1alpha1.DataprocSessionSpec{
			ProjectRef: &refs.ProjectRef{
				External: "projects/test-project",
			},
			Location: direct.LazyPtr("us-central1"),
		},
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(session)
	if err != nil {
		t.Fatalf("ToUnstructured failed: %v", err)
	}
	u := &unstructured.Unstructured{Object: uObj}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(u).Build()

	m := &sessionModel{
		config: config.ControllerConfig{
			HTTPClient: http.DefaultClient,
		},
	}

	adapter, err := m.AdapterForObject(ctx, &directbase.AdapterForObjectOperation{
		Object: u,
		Reader: reader,
	})
	if err != nil {
		t.Fatalf("AdapterForObject failed: %v", err)
	}

	sAdapter, ok := adapter.(*sessionAdapter)
	if !ok {
		t.Fatalf("unexpected adapter type %T", adapter)
	}

	labels := sAdapter.desired.Labels
	if labels == nil {
		t.Fatalf("desired.Labels is nil")
	}

	if _, found := labels["app.kubernetes.io/managed-by"]; found {
		t.Errorf("expected KRM-prefixed label 'app.kubernetes.io/managed-by' to be removed from GCP labels, but found it")
	}
	if labels["env"] != "prod" {
		t.Errorf("expected label env=prod, got %q", labels["env"])
	}
	if labels["managed-by-cnrm"] != "true" {
		t.Errorf("expected managed-by-cnrm=true, got %q", labels["managed-by-cnrm"])
	}
}

func TestDataprocBatch_KRMPrefixedLabels(t *testing.T) {
	batch := &krmv1alpha1.DataprocBatch{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krmv1alpha1.DataprocBatchGVK.GroupVersion().String(),
			Kind:       krmv1alpha1.DataprocBatchGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-batch",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
				"env":                          "prod",
			},
		},
	}

	labels := label.GCPLabels(batch)
	if _, found := labels["app.kubernetes.io/managed-by"]; found {
		t.Errorf("expected KRM-prefixed label 'app.kubernetes.io/managed-by' to be removed from GCP labels, but found it")
	}
	if labels["env"] != "prod" {
		t.Errorf("expected label env=prod, got %q", labels["env"])
	}
	if labels["managed-by-cnrm"] != "true" {
		t.Errorf("expected managed-by-cnrm=true, got %q", labels["managed-by-cnrm"])
	}
}

func TestDataprocCluster_KRMPrefixedLabels(t *testing.T) {
	cluster := &krmv1beta1.DataprocCluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krmv1beta1.DataprocClusterGVK.GroupVersion().String(),
			Kind:       krmv1beta1.DataprocClusterGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cluster",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
				"env":                          "prod",
			},
		},
	}

	labels := label.GCPLabels(cluster)
	if _, found := labels["app.kubernetes.io/managed-by"]; found {
		t.Errorf("expected KRM-prefixed label 'app.kubernetes.io/managed-by' to be removed from GCP labels, but found it")
	}
	if labels["env"] != "prod" {
		t.Errorf("expected label env=prod, got %q", labels["env"])
	}
	if labels["managed-by-cnrm"] != "true" {
		t.Errorf("expected managed-by-cnrm=true, got %q", labels["managed-by-cnrm"])
	}
}
