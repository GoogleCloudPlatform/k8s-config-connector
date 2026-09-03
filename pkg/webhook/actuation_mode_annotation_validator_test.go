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

package webhook

import (
	"context"
	"testing"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/yaml"
)

func TestActuationModeAnnotationValidator(t *testing.T) {
	v := &actuationModeAnnotationValidator{}

	tests := []struct {
		name          string
		operation     admissionv1.Operation
		objYaml       string
		oldObjYaml    string
		expectWarning bool
		warningMsg    string
	}{
		{
			name:      "no annotation on create",
			operation: admissionv1.Create,
			objYaml: `
apiVersion: foo.cnrm.cloud.google.com/v1beta1
kind: FooInstance
metadata:
  name: my-instance
`,
			expectWarning: false,
		},
		{
			name:      "reconciling annotation on create",
			operation: admissionv1.Create,
			objYaml: `
apiVersion: foo.cnrm.cloud.google.com/v1beta1
kind: FooInstance
metadata:
  name: my-instance
  annotations:
    cnrm.cloud.google.com/actuation-mode: Reconciling
`,
			expectWarning: false,
		},
		{
			name:      "paused annotation on create",
			operation: admissionv1.Create,
			objYaml: `
apiVersion: foo.cnrm.cloud.google.com/v1beta1
kind: FooInstance
metadata:
  name: my-instance
  annotations:
    cnrm.cloud.google.com/actuation-mode: Paused
`,
			expectWarning: true,
			warningMsg:    "Resource 'my-instance' has cnrm.cloud.google.com/actuation-mode: \"Paused\". All actuation against GCP (including Create, Update, and Delete) is halted until unpaused.",
		},
		{
			name:      "paused annotation on delete",
			operation: admissionv1.Delete,
			oldObjYaml: `
apiVersion: foo.cnrm.cloud.google.com/v1beta1
kind: FooInstance
metadata:
  name: my-instance
  annotations:
    cnrm.cloud.google.com/actuation-mode: Paused
`,
			expectWarning: true,
			warningMsg:    "Resource 'my-instance' has cnrm.cloud.google.com/actuation-mode: \"Paused\". All actuation against GCP (including Create, Update, and Delete) is halted until unpaused.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req admission.Request
			req.Operation = tc.operation

			if tc.objYaml != "" {
				u := &unstructured.Unstructured{}
				if err := yaml.Unmarshal([]byte(tc.objYaml), &u.Object); err != nil {
					t.Fatalf("failed to unmarshal objYaml: %v", err)
				}
				raw, err := yaml.Marshal(u)
				if err != nil {
					t.Fatalf("failed to marshal obj: %v", err)
				}
				req.Object = runtime.RawExtension{Raw: raw}
			}

			if tc.oldObjYaml != "" {
				u := &unstructured.Unstructured{}
				if err := yaml.Unmarshal([]byte(tc.oldObjYaml), &u.Object); err != nil {
					t.Fatalf("failed to unmarshal oldObjYaml: %v", err)
				}
				raw, err := yaml.Marshal(u)
				if err != nil {
					t.Fatalf("failed to marshal old obj: %v", err)
				}
				req.OldObject = runtime.RawExtension{Raw: raw}
			}

			resp := v.Handle(context.Background(), req)
			if !resp.Allowed {
				t.Errorf("expected allowed to be true, got %v", resp.Allowed)
			}

			if tc.expectWarning {
				if len(resp.Warnings) == 0 {
					t.Errorf("expected warning, but got none")
				} else if resp.Warnings[0] != tc.warningMsg {
					t.Errorf("expected warning message '%s', got '%s'", tc.warningMsg, resp.Warnings[0])
				}
			} else {
				if len(resp.Warnings) > 0 {
					t.Errorf("expected no warnings, but got %v", resp.Warnings)
				}
			}
		})
	}
}
