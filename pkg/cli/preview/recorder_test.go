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
	"errors"
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

func TestRecordReconcileEnd_HealthStatus(t *testing.T) {
	ctx := context.Background()

	makeObj := func(name string) *unstructured.Unstructured {
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "monitoring.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "MonitoringAlertPolicy",
		})
		u.SetNamespace("default")
		u.SetName(name)
		return u
	}

	tests := []struct {
		name             string
		reconcileErr     error
		wantStatus       ReconcileStatus
		wantReconcileErr string
	}{
		{
			name:         "nil error is healthy",
			reconcileErr: nil,
			wantStatus:   ReconcileStatusHealthy,
		},
		{
			name:         "blocked GCP error is healthy",
			reconcileErr: &BlockedGCPError{Method: "POST", URL: "https://monitoring.googleapis.com/v3/projects/p/alertPolicies"},
			wantStatus:   ReconcileStatusHealthy,
		},
		{
			name:         "wrapped blocked GCP error is healthy",
			reconcileErr: fmt.Errorf("failed gcp call: %w", &BlockedGCPError{Method: "PATCH", URL: "https://monitoring.googleapis.com/v3/projects/p/alertPolicies/123"}),
			wantStatus:   ReconcileStatusHealthy,
		},
		{
			name:         "blocked kube error is healthy",
			reconcileErr: &BlockedKubeError{Method: "status.patch"},
			wantStatus:   ReconcileStatusHealthy,
		},
		{
			name:         "wrapped blocked kube error is healthy",
			reconcileErr: fmt.Errorf("updating status: %w", &BlockedKubeError{Method: "update"}),
			wantStatus:   ReconcileStatusHealthy,
		},
		{
			name:             "non-blocked error marks resource unhealthy",
			reconcileErr:     fmt.Errorf("unknown enum value \"warning\""),
			wantStatus:       ReconcileStatusUnhealthy,
			wantReconcileErr: "unknown enum value \"warning\"",
		},
		{
			name:             "joined error containing non-blocked error marks resource unhealthy",
			reconcileErr:     errors.Join(&BlockedKubeError{Method: "status.patch"}, fmt.Errorf("unknown enum value \"warning\"")),
			wantStatus:       ReconcileStatusUnhealthy,
			wantReconcileErr: "\"status.patch\" blocked in preview mode\nunknown enum value \"warning\"",
		},
		{
			name:         "joined error containing only blocked errors is healthy",
			reconcileErr: errors.Join(&BlockedKubeError{Method: "status.patch"}, &BlockedGCPError{Method: "POST", URL: "https://example.com"}),
			wantStatus:   ReconcileStatusHealthy,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recorder := NewRecorder()
			obj := makeObj("test-policy")
			gknn := gknnFromUnstructured(obj)

			recorder.recordReconcileStart(ctx, obj, k8s.ReconcilerTypeDirect)
			recorder.recordReconcileEnd(ctx, obj, reconcile.Result{}, tc.reconcileErr, k8s.ReconcilerTypeDirect)

			results := recorder.GenerateRecorderReconciledResults()
			res, ok := results.results[gknn]
			if !ok {
				t.Fatalf("expected result for %v", gknn)
			}
			if res.ReconcileStatus != tc.wantStatus {
				t.Errorf("ReconcileStatus = %v, want %v", res.ReconcileStatus, tc.wantStatus)
			}
			if res.ReconcileError != tc.wantReconcileErr {
				t.Errorf("ReconcileError = %q, want %q", res.ReconcileError, tc.wantReconcileErr)
			}
		})
	}
}

func TestOnError_NonBlockedErrorMarksUnhealthy(t *testing.T) {
	ctx := context.Background()
	recorder := NewRecorder()
	listener := recorder.NewStructuredReportingListener()

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "monitoring.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MonitoringAlertPolicy",
	})
	u.SetNamespace("default")
	u.SetName("alert-policy-mapping-fail")

	resObj := &k8s.Resource{}
	resObj.APIVersion = "monitoring.cnrm.cloud.google.com/v1beta1"
	resObj.Kind = "MonitoringAlertPolicy"
	resObj.Namespace = "default"
	resObj.Name = "alert-policy-mapping-fail"

	listener.OnReconcileStart(ctx, u, k8s.ReconcilerTypeDirect)
	listener.OnError(ctx, fmt.Errorf("unknown enum value \"warning\""), resObj)
	listener.OnReconcileEnd(ctx, u, reconcile.Result{}, fmt.Errorf("Update call failed: unknown enum value \"warning\""), k8s.ReconcilerTypeDirect)

	results := recorder.GenerateRecorderReconciledResults()
	gknn := gknnFromUnstructured(u)
	got, ok := results.results[gknn]
	if !ok {
		t.Fatalf("expected result for %v", gknn)
	}
	if got.ReconcileStatus != ReconcileStatusUnhealthy {
		t.Errorf("ReconcileStatus = %v, want %v", got.ReconcileStatus, ReconcileStatusUnhealthy)
	}
	if len(got.GCPActions) != 0 {
		t.Errorf("expected no GCPActions recorded when mapping/comparison fails before GCP call, got %v", got.GCPActions)
	}
}
