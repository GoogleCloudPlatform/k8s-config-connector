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

package colab

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestColabRuntime_RejectsInvalidLocation(t *testing.T) {
	ctx := context.Background()
	model := &runtimeModel{config: config.ControllerConfig{}}

	maliciousLocations := []string{
		"127.0.0.1:34567/?",
		"127.0.0.1:34567?",
		"attacker.com/?",
		"attacker.com?",
		"attacker.com#",
		"user@attacker.com?",
		"..",
		"",
	}

	for _, loc := range maliciousLocations {
		t.Run("client/"+loc, func(t *testing.T) {
			if _, err := model.client(ctx, loc); err == nil {
				t.Errorf("runtimeModel.client(ctx, %q) succeeded unexpectedly, want error", loc)
			}
		})

		t.Run("AdapterForObject/"+loc, func(t *testing.T) {
			obj := &krm.ColabRuntime{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-runtime",
					Namespace: "default",
				},
				Spec: krm.ColabRuntimeSpec{
					Runtime_Parent: krm.Runtime_Parent{
						ProjectRef: &refs.ProjectRef{
							External: "test-project",
						},
						Location: loc,
					},
				},
			}
			uMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
			if err != nil {
				t.Fatalf("ToUnstructured failed: %v", err)
			}
			u := &unstructured.Unstructured{Object: uMap}
			u.SetGroupVersionKind(krm.ColabRuntimeGVK)

			op := &directbase.AdapterForObjectOperation{
				Object: u,
			}
			if _, err := model.AdapterForObject(ctx, op); err == nil {
				t.Errorf("runtimeModel.AdapterForObject with location=%q succeeded unexpectedly, want error", loc)
			}
		})
	}
}
