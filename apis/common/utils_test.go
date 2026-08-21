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

package common

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestToStructuredType(t *testing.T) {
	// 1. Test when object is already of structured type T
	obj := &corev1.Pod{}
	obj.SetName("test-pod")

	res, err := ToStructuredType[*corev1.Pod](obj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GetName() != "test-pod" {
		t.Errorf("expected name to be 'test-pod', got %q", res.GetName())
	}

	// 2. Test when object is unstructured
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Pod",
			"metadata": map[string]interface{}{
				"name": "unstructured-pod",
			},
		},
	}

	res2, err := ToStructuredType[*corev1.Pod](u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res2.GetName() != "unstructured-pod" {
		t.Errorf("expected name to be 'unstructured-pod', got %q", res2.GetName())
	}
}
