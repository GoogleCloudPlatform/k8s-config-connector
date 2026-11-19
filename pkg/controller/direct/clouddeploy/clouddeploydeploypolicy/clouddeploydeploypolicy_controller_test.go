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

package clouddeploydeploypolicy

import (
	"context"
	"net/http"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCloudDeployDeployPolicy_KRMPrefixedLabels(t *testing.T) {
	ctx := context.Background()

	policy := &krm.CloudDeployDeployPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krm.DeployDeployPolicyGVK.GroupVersion().String(),
			Kind:       krm.DeployDeployPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-policy",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
				"env":                          "prod",
			},
		},
		Spec: krm.DeployPolicySpec{
			Parent: krm.Parent{
				ProjectRef: &refs.ProjectRef{
					External: "projects/test-project",
				},
				Location: "us-central1",
			},
		},
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(policy)
	if err != nil {
		t.Fatalf("ToUnstructured failed: %v", err)
	}
	u := &unstructured.Unstructured{Object: uObj}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).Build()

	m := &model{
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

	deployAdapter, ok := adapter.(*Adapter)
	if !ok {
		t.Fatalf("unexpected adapter type %T", adapter)
	}

	if err := deployAdapter.resolveReferences(ctx); err != nil {
		t.Fatalf("resolveReferences failed: %v", err)
	}

	labels := deployAdapter.desiredPb.Labels
	if labels == nil {
		t.Fatalf("desiredPb.Labels is nil")
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
