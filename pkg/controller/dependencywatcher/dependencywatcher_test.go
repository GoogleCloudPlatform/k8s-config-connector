// Copyright 2022 Google LLC
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

package dependencywatcher_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dependencywatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	dynamicfake "k8s.io/client-go/dynamic/fake"
)

func newResourceUnstructured(status map[string]interface{}) *unstructured.Unstructured {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
			"kind":       "PubSubTopic",
			"metadata": map[string]interface{}{
				"name": "test_topic",
			},
			"spec": map[string]interface{}{
				"resourceID": "pubsubtopic-sample",
			},
		},
	}
	if status != nil {
		u.Object["status"] = status
	}
	return u
}

func newStatus(readyStatus corev1.ConditionStatus) map[string]interface{} {
	return map[string]interface{}{
		"conditions": []interface{}{
			map[string]interface{}{
				"type":   "Ready",
				"status": string(readyStatus),
			},
		},
	}
}

func TestIsResourceReady(t *testing.T) {
	tests := []struct {
		name                  string
		unstructuredReference *unstructured.Unstructured
		shouldOk              bool
	}{
		{
			name:                  "referenced resource ready",
			unstructuredReference: newResourceUnstructured(newStatus(corev1.ConditionTrue)),
			shouldOk:              true,
		},
		{
			name:                  "referenced resource does not exist",
			unstructuredReference: nil,
			shouldOk:              false,
		},
		{
			name:                  "referenced resource not ready",
			unstructuredReference: newResourceUnstructured(newStatus(corev1.ConditionFalse)),
			shouldOk:              false,
		},
		{
			name:                  "referenced resource status not set",
			unstructuredReference: newResourceUnstructured(nil),
			shouldOk:              false,
		},
	}
	referencingResource, err := k8s.NewResource(&unstructured.Unstructured{})
	if err != nil {
		t.Fatalf("error creating k8s resource from empty unstructured: %v", err)
	}
	sampleReferencedResource, err := k8s.NewResource(newResourceUnstructured(nil))
	if err != nil {
		t.Fatalf("error creating k8s resource from sample unstructured resource: %v", err)
	}
	refGVK := sampleReferencedResource.GroupVersionKind()
	refNN := sampleReferencedResource.GetNamespacedName()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fakeClient := dynamicfake.NewSimpleDynamicClient(runtime.NewScheme())
			if tc.unstructuredReference != nil {
				fakeClient.
					Resource(k8s.ToGVR(refGVK)).
					Namespace(refNN.Namespace).
					Create(context.TODO(), tc.unstructuredReference, metav1.CreateOptions{})
			}
			depWatcher := dependencywatcher.CreateWatchForResourceWithClient(referencingResource, fakeClient)
			ok, reason, err := depWatcher.IsReferenceReady(context.TODO(), refNN, refGVK)

			if err != nil {
				t.Fatalf("got unexpected error: '%v'", err)
			}
			if ok != tc.shouldOk {
				t.Fatalf("got 'ok' value of '%v', but expected '%v'", ok, tc.shouldOk)
			}
			if ok && reason != "" {
				t.Fatalf("got 'reason' value of '%v', but expected 'reason' to be empty string if ok is true and 'err' is nil", reason)
			}
			if !ok && reason == "" {
				t.Fatalf("got empty string as 'reason', but expected 'reason' to be non-empty string if 'ok is false and 'err' is nil")
			}
		})
	}
}
