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

package configconnector

import (
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	k8sreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestCCControllerOverridesField(t *testing.T) {
	t.Parallel()
	cc := &corev1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: corev1beta1.ConfigConnectorAllowedName,
		},
		Spec: corev1beta1.ConfigConnectorSpec{
			Experiments: &corev1beta1.CCExperiments{
				ControllerOverrides: map[string]k8sreconciler.ReconcilerType{
					"BigQueryDataset.bigquery.cnrm.cloud.google.com": "direct",
				},
			},
		},
	}

	unstructuredCC, err := runtime.DefaultUnstructuredConverter.ToUnstructured(cc)
	if err != nil {
		t.Fatalf("error converting to unstructured: %v", err)
	}

	_, found, err := unstructured.NestedMap(unstructuredCC, "spec", "experiments", "controllerOverrides")
	if err != nil {
		t.Fatalf("error getting nested map: %v", err)
	}
	if !found {
		t.Fatalf("field .spec.experiments.controllerOverrides not found in unstructured object")
	}
}
