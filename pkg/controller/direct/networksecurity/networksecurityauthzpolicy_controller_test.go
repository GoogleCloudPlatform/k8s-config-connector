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

package networksecurity

import (
	"context"
	"net/http"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAuthzPolicyAdapterForObject_SanitizesKRMLabels(t *testing.T) {
	ctx := context.Background()

	obj := &krm.NetworkSecurityAuthzPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krm.NetworkSecurityAuthzPolicyGVK.GroupVersion().String(),
			Kind:       krm.NetworkSecurityAuthzPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-authz-policy",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by":       "configmanagement.gke.io",
				"configsync.gke.io/declared-version": "v1alpha1",
				"applyset.kubernetes.io/part-of":     "applyset-123",
				"env":                                "stg",
			},
		},
		Spec: krm.NetworkSecurityAuthzPolicySpec{
			Location: direct.LazyPtr("global"),
			ProjectRef: &refs.ProjectRef{
				External: "projects/test-project",
			},
			Action:        direct.LazyPtr("ALLOW"),
			PolicyProfile: direct.LazyPtr("REQUEST_AUTHZ"),
		},
	}

	rawObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		t.Fatalf("ToUnstructured() unexpected error: %v", err)
	}
	u := &unstructured.Unstructured{Object: rawObj}

	reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()
	model, err := NewAuthzPolicyModel(ctx, &config.ControllerConfig{
		HTTPClient: &http.Client{},
	})
	if err != nil {
		t.Fatalf("NewAuthzPolicyModel() unexpected error: %v", err)
	}

	adapterI, err := model.AdapterForObject(ctx, &directbase.AdapterForObjectOperation{
		Reader: reader,
		Object: u,
	})
	if err != nil {
		t.Fatalf("AdapterForObject() unexpected error: %v", err)
	}

	adapter, ok := adapterI.(*authzPolicyAdapter)
	if !ok {
		t.Fatalf("AdapterForObject() returned %T, want *authzPolicyAdapter", adapterI)
	}

	wantLabels := map[string]string{
		"env":             "stg",
		"managed-by-cnrm": "true",
	}
	if diff := cmp.Diff(wantLabels, adapter.desired.GetLabels()); diff != "" {
		t.Errorf("desired.Labels mismatch (-want +got):\n%s", diff)
	}
}
