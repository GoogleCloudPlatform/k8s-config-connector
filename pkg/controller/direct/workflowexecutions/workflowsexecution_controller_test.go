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

package workflowexecutions

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflowexecutions/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestWorkflowsExecution_KRMPrefixedLabels(t *testing.T) {
	execution := &krm.WorkflowsExecution{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krm.WorkflowsExecutionGVK.GroupVersion().String(),
			Kind:       krm.WorkflowsExecutionGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-execution",
			Namespace: "test-ns",
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "configmanagement.gke.io",
				"env":                          "prod",
			},
		},
	}

	labels := label.GCPLabels(execution)
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
