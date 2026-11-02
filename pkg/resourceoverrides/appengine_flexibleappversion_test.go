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

package resourceoverrides_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestGetAppEngineFlexibleAppVersionResourceOverrides_PreActuationTransform(t *testing.T) {
	tests := []struct {
		name        string
		initialVal  string
		expectedVal string
	}{
		{
			name:        "URI format with apps and services prefix",
			initialVal:  "apps/test-project/services/default",
			expectedVal: "default",
		},
		{
			name:        "URI format with custom service name",
			initialVal:  "apps/my-app/services/my-custom-service",
			expectedVal: "my-custom-service",
		},
		{
			name:        "Plain service name without prefix",
			initialVal:  "default",
			expectedVal: "default",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ro := resourceoverrides.GetAppEngineFlexibleAppVersionResourceOverrides()
			if len(ro.Overrides) != 1 {
				t.Fatalf("expected 1 override, got %d", len(ro.Overrides))
			}
			override := ro.Overrides[0]
			res := &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					Kind: "AppEngineFlexibleAppVersion",
				},
				Spec: map[string]interface{}{
					"serviceRef": map[string]interface{}{
						"external": tc.initialVal,
					},
				},
			}

			// Act
			if err := override.PreActuationTransform(res); err != nil {
				t.Fatalf("PreActuationTransform failed: %v", err)
			}

			// Assert
			val, found, err := unstructured.NestedString(res.Spec, "serviceRef", "external")
			if err != nil {
				t.Fatalf("unexpected error reading serviceRef.external: %v", err)
			}
			if !found {
				t.Fatalf("serviceRef.external not found")
			}
			if val != tc.expectedVal {
				t.Errorf("expected serviceRef.external = %q, got %q", tc.expectedVal, val)
			}
		})
	}
}

func TestGetAppEngineFlexibleAppVersionResourceOverrides_PostActuationTransform(t *testing.T) {
	// Arrange
	ro := resourceoverrides.GetAppEngineFlexibleAppVersionResourceOverrides()
	override := ro.Overrides[0]
	original := &k8s.Resource{
		TypeMeta: metav1.TypeMeta{
			Kind: "AppEngineFlexibleAppVersion",
		},
		Spec: map[string]interface{}{
			"serviceRef": map[string]interface{}{
				"external": "apps/test-project/services/default",
			},
		},
	}
	reconciled := &k8s.Resource{
		TypeMeta: metav1.TypeMeta{
			Kind: "AppEngineFlexibleAppVersion",
		},
		Spec: map[string]interface{}{
			"serviceRef": map[string]interface{}{
				"external": "default",
			},
		},
	}

	// Act
	if err := override.PostActuationTransform(original, reconciled, nil, nil); err != nil {
		t.Fatalf("PostActuationTransform failed: %v", err)
	}

	// Assert
	val, found, err := unstructured.NestedString(reconciled.Spec, "serviceRef", "external")
	if err != nil {
		t.Fatalf("unexpected error reading serviceRef.external: %v", err)
	}
	if !found {
		t.Fatalf("serviceRef.external not found in reconciled resource")
	}
	expected := "apps/test-project/services/default"
	if val != expected {
		t.Errorf("expected serviceRef.external preserved as %q, got %q", expected, val)
	}
}

func TestGetAppEngineFlexibleAppVersionResourceOverrides_PostUpdateStatusTransform(t *testing.T) {
	// Arrange
	ro := resourceoverrides.GetAppEngineFlexibleAppVersionResourceOverrides()
	override := ro.Overrides[0]
	res := &k8s.Resource{
		TypeMeta: metav1.TypeMeta{
			Kind: "AppEngineFlexibleAppVersion",
		},
		Spec: map[string]interface{}{
			"serviceRef": map[string]interface{}{
				"external": "apps/test-project/services/default",
			},
		},
	}

	// Act
	if err := override.PostUpdateStatusTransform(res); err != nil {
		t.Fatalf("PostUpdateStatusTransform failed: %v", err)
	}

	// Assert
	val, found, err := unstructured.NestedString(res.Spec, "serviceRef", "external")
	if err != nil {
		t.Fatalf("unexpected error reading serviceRef.external: %v", err)
	}
	if !found {
		t.Fatalf("serviceRef.external not found")
	}
	expected := "default"
	if val != expected {
		t.Errorf("expected serviceRef.external = %q, got %q", expected, val)
	}
}
