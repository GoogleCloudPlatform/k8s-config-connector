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
	"encoding/json"
	"testing"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func TestErrorHandlingModeAnnotationValidator(t *testing.T) {
	tests := []struct {
		name          string
		operation     admissionv1.Operation
		annotations   map[string]string
		wantAllowed   bool
		wantErrStatus bool
	}{
		{
			name:        "no annotations",
			operation:   admissionv1.Create,
			wantAllowed: true,
		},
		{
			name:      "ContinuousRetry - valid value",
			operation: admissionv1.Create,
			annotations: map[string]string{
				"cnrm.cloud.google.com/error-handling-mode": "ContinuousRetry",
			},
			wantAllowed: true,
		},
		{
			name:      "PauseOnTerminalError - valid value",
			operation: admissionv1.Create,
			annotations: map[string]string{
				"cnrm.cloud.google.com/error-handling-mode": "PauseOnTerminalError",
			},
			wantAllowed: true,
		},
		{
			name:      "invalid value - should be denied",
			operation: admissionv1.Create,
			annotations: map[string]string{
				"cnrm.cloud.google.com/error-handling-mode": "InvalidValue",
			},
			wantAllowed: false,
		},
		{
			name:      "invalid value on Update - should be denied",
			operation: admissionv1.Update,
			annotations: map[string]string{
				"cnrm.cloud.google.com/error-handling-mode": "InvalidValue",
			},
			wantAllowed: false,
		},
	}

	validator := &errorHandlingModeAnnotationValidator{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			objMap := map[string]interface{}{
				"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
				"kind":       "APIGatewayAPI",
				"metadata": map[string]interface{}{
					"name":      "test-resource",
					"namespace": "test-namespace",
				},
			}
			if tc.annotations != nil {
				objMap["metadata"].(map[string]interface{})["annotations"] = tc.annotations
			}

			rawBytes, err := json.Marshal(objMap)
			if err != nil {
				t.Fatalf("failed to marshal test object: %v", err)
			}

			req := admission.Request{
				AdmissionRequest: admissionv1.AdmissionRequest{
					Operation: tc.operation,
				},
			}

			if tc.operation == admissionv1.Delete {
				req.AdmissionRequest.OldObject = runtime.RawExtension{Raw: rawBytes}
			} else {
				req.AdmissionRequest.Object = runtime.RawExtension{Raw: rawBytes}
			}

			resp := validator.Handle(context.Background(), req)
			if resp.Allowed != tc.wantAllowed {
				t.Errorf("expected allowed=%v, got %v (Result: %v)", tc.wantAllowed, resp.Allowed, resp.Result)
			}
		})
	}
}
