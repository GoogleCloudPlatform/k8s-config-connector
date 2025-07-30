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

package resourcewatcher

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	dynamicfake "k8s.io/client-go/dynamic/fake"
)

func TestWatchResourceTimeout(t *testing.T) {
	unreadyResourceUnstructured := newResourceUnstructured(newStatus(corev1.ConditionFalse))
	gvk, nn, err := getResourceInformation(unreadyResourceUnstructured)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	fake := dynamicfake.NewSimpleDynamicClient(runtime.NewScheme())
	logger := log.Log.WithName("resourcewatcher-test-timeout")
	watch, err := NewWithClient(fake, logger).WatchResource(context.TODO(), nn, gvk)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	if err := waitForResourceToBeReadyOrDeletedViaWatch(ctx, watch, logger); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got error '%v', expected '%v'", err, context.DeadlineExceeded)
	}
}

func TestWatchResourceSuccess(t *testing.T) {
	readyResourceUnstructured := newResourceUnstructured(newStatus(corev1.ConditionTrue))
	gvk, nn, err := getResourceInformation(readyResourceUnstructured)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	fake := dynamicfake.NewSimpleDynamicClient(runtime.NewScheme())
	logger := log.Log.WithName("resourcewatcher-test-success")
	watch, err := NewWithClient(fake, logger).WatchResource(context.TODO(), nn, gvk)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if _, err := fake.Resource(k8s.ToGVR(gvk)).
		Namespace(nn.Namespace).
		Create(context.TODO(), readyResourceUnstructured, metav1.CreateOptions{}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	if err := waitForResourceToBeReadyOrDeletedViaWatch(ctx, watch, logger); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
}

func newResourceUnstructured(status map[string]interface{}) *unstructured.Unstructured {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
			"kind":       "PubSubTopic",
			"metadata": map[string]interface{}{
				"name":      "test_topic",
				"namespace": "test_namespace",
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

func getResourceInformation(u *unstructured.Unstructured) (schema.GroupVersionKind, types.NamespacedName, error) {
	resource, err := k8s.NewResource(u)
	if err != nil {
		return schema.GroupVersionKind{}, types.NamespacedName{}, fmt.Errorf("error creating k8s resource from unstructured: %w", err)
	}
	return resource.GroupVersionKind(), resource.GetNamespacedName(), nil
}
