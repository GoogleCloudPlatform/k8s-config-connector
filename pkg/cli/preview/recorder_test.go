// Copyright 2025 Google LLC
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

package preview

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestToTrackedGVR(t *testing.T) {
	tests := []struct {
		name                        string
		apiResource                 metav1.APIResource
		apiResourceListGroupVersion schema.GroupVersion
		wantGVR                     schema.GroupVersionResource
		wantOk                      bool
	}{
		{
			name: "Valid CNRM resource",
			apiResource: metav1.APIResource{
				Name:    "storagebuckets",
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "storage.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: true,
		},
		{
			name: "Valid CNRM resource - inherit Group/Version",
			apiResource: metav1.APIResource{
				Name: "storagebuckets",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "storage.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: true,
		},
		{
			name: "Core CNRM resource (ignored)",
			apiResource: metav1.APIResource{
				Name:    "configconnectors",
				Group:   "core.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "core.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "core.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "configconnectors",
			},
			wantOk: false,
		},
		{
			name: "Non-CNRM resource (ignored)",
			apiResource: metav1.APIResource{
				Name:    "deployments",
				Group:   "apps",
				Version: "v1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "apps",
				Version: "v1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "apps",
				Version:  "v1",
				Resource: "deployments",
			},
			wantOk: false,
		},
		{
			name: "Ignored CRD (gameservicesrealms)",
			apiResource: metav1.APIResource{
				Name:    "gameservicesrealms",
				Group:   "gameservices.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "gameservices.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "gameservices.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "gameservicesrealms",
			},
			wantOk: false,
		},
		{
			name: "Fake CNRM group (ignored)",
			apiResource: metav1.APIResource{
				Name:    "storagebuckets",
				Group:   "fake-cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "fake-cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "fake-cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotGVR, gotOk := toTrackedGVR(tc.apiResource, tc.apiResourceListGroupVersion)
			if diff := cmp.Diff(tc.wantGVR, gotGVR); diff != "" {
				t.Errorf("toTrackedGVR() GVR mismatch (-want +got):\n%s", diff)
			}
			if gotOk != tc.wantOk {
				t.Errorf("toTrackedGVR() ok = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}

func TestRecorder_OnError_NonBlockedError(t *testing.T) {
	ctx := context.Background()
	recorder := NewRecorder()
	listener := recorder.NewStructuredReportingListener()

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "spanner.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SpannerInstance",
	})
	u.SetName("test-instance")
	u.SetNamespace("test-ns")

	testErr := fmt.Errorf("mapping error: invalid proto conversion")
	listener.OnError(ctx, testErr, u)

	gknn := GKNN{
		Group:     "spanner.cnrm.cloud.google.com",
		Kind:      "SpannerInstance",
		Namespace: "test-ns",
		Name:      "test-instance",
	}

	info := recorder.getObjectInfo(gknn)
	if len(info.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(info.events))
	}
	event := info.events[0]
	if event.eventType != EventTypeError {
		t.Errorf("expected event type %v, got %v", EventTypeError, event.eventType)
	}
	if event.err == nil || event.err.Error() != testErr.Error() {
		t.Errorf("expected error %v, got %v", testErr, event.err)
	}
}

func TestRecorder_OnError_ResourceArg(t *testing.T) {
	ctx := context.Background()
	recorder := NewRecorder()
	listener := recorder.NewStructuredReportingListener()

	resource := &k8s.Resource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "storage.cnrm.cloud.google.com/v1beta1",
			Kind:       "StorageBucket",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "test-bucket",
		},
	}

	testErr := fmt.Errorf("update failed: reference not found")
	listener.OnError(ctx, testErr, resource)

	gknn := GKNN{
		Group:     "storage.cnrm.cloud.google.com",
		Kind:      "StorageBucket",
		Namespace: "test-ns",
		Name:      "test-bucket",
	}

	info := recorder.getObjectInfo(gknn)
	if len(info.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(info.events))
	}
	event := info.events[0]
	if event.eventType != EventTypeError {
		t.Errorf("expected event type %v, got %v", EventTypeError, event.eventType)
	}
	if event.err == nil || event.err.Error() != testErr.Error() {
		t.Errorf("expected error %v, got %v", testErr, event.err)
	}
}

func TestRecorder_RecordReconcileEnd_WithError(t *testing.T) {
	ctx := context.Background()
	recorder := NewRecorder()
	listener := recorder.NewStructuredReportingListener()

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInstance",
	})
	u.SetName("test-compute")
	u.SetNamespace("test-ns")

	testErr := fmt.Errorf("reconcile failed: early validation error")
	listener.OnReconcileEnd(ctx, u, reconcile.Result{}, testErr, k8s.ReconcilerTypeDirect)

	gknn := GKNN{
		Group:     "compute.cnrm.cloud.google.com",
		Kind:      "ComputeInstance",
		Namespace: "test-ns",
		Name:      "test-compute",
	}

	info := recorder.getObjectInfo(gknn)
	if len(info.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(info.events))
	}
	event := info.events[0]
	if event.eventType != EventTypeReconcileEnd {
		t.Errorf("expected event type %v, got %v", EventTypeReconcileEnd, event.eventType)
	}
	if event.err == nil || event.err.Error() != testErr.Error() {
		t.Errorf("expected error %v, got %v", testErr, event.err)
	}
}

func TestRecorder_OnError_BlockedErrors_Ignored(t *testing.T) {
	ctx := context.Background()
	recorder := NewRecorder()
	listener := recorder.NewStructuredReportingListener()

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "spanner.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SpannerInstance",
	})
	u.SetName("test-instance")
	u.SetNamespace("test-ns")

	gknn := GKNN{
		Group:     "spanner.cnrm.cloud.google.com",
		Kind:      "SpannerInstance",
		Namespace: "test-ns",
		Name:      "test-instance",
	}

	// 1. Blocked GCP error should record GCPAction, not EventTypeError
	blockedGCP := BlockedGCPError{
		Method: "POST",
		URL:    "https://spanner.googleapis.com/v1/projects/p/instances",
	}
	listener.OnError(ctx, blockedGCP, u)

	info := recorder.getObjectInfo(gknn)
	if len(info.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(info.events))
	}
	if info.events[0].eventType != EventTypeGCPAction {
		t.Errorf("expected event type %v, got %v", EventTypeGCPAction, info.events[0].eventType)
	}

	// 2. Blocked Kube error should be ignored by OnError
	blockedKubeErr := fmt.Errorf("\"update\" blocked in preview mode")
	listener.OnError(ctx, blockedKubeErr, u)

	// Events count should still be 1 (no EventTypeError added)
	if len(info.events) != 1 {
		t.Fatalf("expected 1 event after blocked kube error, got %d", len(info.events))
	}
}
