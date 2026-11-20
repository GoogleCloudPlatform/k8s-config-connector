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

package e2e

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcileEndNotEvaluatedReason(t *testing.T) {
	tests := []struct {
		name          string
		end           reconcileEnd
		wantEvaluated bool
	}{
		{
			name:          "ended with an error",
			end:           reconcileEnd{err: errors.New("RequiresNew")},
			wantEvaluated: false,
		},
		{
			name:          "requeued",
			end:           reconcileEnd{result: reconcile.Result{Requeue: true}},
			wantEvaluated: false,
		},
		{
			name:          "zero result",
			end:           reconcileEnd{result: reconcile.Result{}},
			wantEvaluated: false,
		},
		{
			name:          "requeued after the jittered period",
			end:           reconcileEnd{result: reconcile.Result{RequeueAfter: time.Second}},
			wantEvaluated: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reason := tc.end.notEvaluatedReason()
			if tc.wantEvaluated && reason != "" {
				t.Errorf("notEvaluatedReason() = %q, want \"\" (diff evaluated)", reason)
			}
			if !tc.wantEvaluated && reason == "" {
				t.Errorf("notEvaluatedReason() = \"\", want a reason (diff not evaluated)")
			}
		})
	}
}

func TestDiffRecorderWaitForReconcileAfter(t *testing.T) {
	ctx := context.Background()
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeInstance"})
	u.SetNamespace("ns")
	u.SetName("name")
	key := reconcileKey(u.GetKind(), u.GetNamespace(), u.GetName())

	r := &diffRecorder{}
	// First Terraform reconcile ends with an error.
	r.OnReconcileStart(ctx, u, k8s.ReconcilerTypeTerraform)
	r.OnReconcileEnd(ctx, u, reconcile.Result{}, errors.New("immutable field changed"), k8s.ReconcilerTypeTerraform)
	// Second Terraform reconcile completes normally.
	r.OnReconcileStart(ctx, u, k8s.ReconcilerTypeTerraform)
	r.OnReconcileEnd(ctx, u, reconcile.Result{RequeueAfter: time.Minute}, nil, k8s.ReconcilerTypeTerraform)
	// A Direct reconcile of the same object must be ignored.
	r.OnReconcileStart(ctx, u, k8s.ReconcilerTypeDirect)
	r.OnReconcileEnd(ctx, u, reconcile.Result{}, errors.New("direct error"), k8s.ReconcilerTypeDirect)

	end, ok := r.WaitForReconcileAfter(key, 0, time.Second)
	if !ok {
		t.Fatalf("WaitForReconcileAfter(after=0) timed out, want the first reconcile's end")
	}
	if end.start != 1 || end.err == nil {
		t.Errorf("WaitForReconcileAfter(after=0) = {start: %d, err: %v}, want the first reconcile (start 1) with an error", end.start, end.err)
	}

	end, ok = r.WaitForReconcileAfter(key, 1, time.Second)
	if !ok {
		t.Fatalf("WaitForReconcileAfter(after=1) timed out, want the second reconcile's end")
	}
	if end.start != 2 || end.err != nil || end.result.RequeueAfter != time.Minute {
		t.Errorf("WaitForReconcileAfter(after=1) = {start: %d, result: %+v, err: %v}, want the second reconcile (start 2) without error", end.start, end.result, end.err)
	}

	if _, ok := r.WaitForReconcileAfter(key, 2, 50*time.Millisecond); ok {
		t.Errorf("WaitForReconcileAfter(after=2) returned an end, want a timeout: the Direct reconcile must not count")
	}
}
