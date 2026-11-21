// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// reconcileEnd is one Terraform reconcile that has ended, as reported by ReportReconcileEnd.
type reconcileEnd struct {
	// start is the number of Terraform reconciles started for the object when this one started (1-based).
	start  int
	result reconcile.Result
	err    error
}

// notEvaluatedReason returns why the reconcile did not run through diff evaluation, or "" if it did.
//
// The Terraform controller (pkg/controller/tf/controller.go, DoReconcile) returns before ReportDiff in these cases:
//   - err != nil: an error returned before the diff, e.g. an immutable/RequiresNew error. An error returned after the
//     diff was already reported (e.g. while applying it) is also reported here, even though the diff was evaluated.
//   - result.Requeue: the reconcile was requeued (e.g. unresolvable dependencies) before the diff.
//   - result.RequeueAfter <= 0: resourceactuation.ShouldSkip returned early (reconcile-interval-in-seconds: "0").
//
// A completed sync returns RequeueAfter = the jittered re-enqueue period, which is > 0 (the harness uses
// jitter.NewDefaultGenerator). The Paused actuation mode also returns a positive RequeueAfter before the diff, but
// the harness does not use that mode, so this check does not detect it.
func (e reconcileEnd) notEvaluatedReason() string {
	switch {
	case e.err != nil:
		return fmt.Sprintf("it ended with an error: %v", e.err)
	case e.result.Requeue:
		return "it was requeued (e.g. unresolvable dependencies) before the diff"
	case e.result.RequeueAfter <= 0:
		return `it was skipped (resourceactuation.ShouldSkip, e.g. reconcile-interval-in-seconds: "0")`
	}
	return ""
}

// diffRecorder records the diffs reported by controllers, and forwards everything to the debug log listener
// that the harness would otherwise install. It also tracks Terraform reconciles per object, so the
// re-reconciliation check can wait for a reconcile to finish instead of sleeping for a fixed time.
type diffRecorder struct {
	structuredreporting.DebugLogListener

	mu    sync.Mutex
	diffs []string
	// starts counts the Terraform reconciles started per object. ends records, per object, every Terraform reconcile
	// that has ended, in order. A controller runs one reconcile per object at a time, so an end always belongs to the
	// latest start.
	starts map[string]int
	ends   map[string][]reconcileEnd
}

var _ structuredreporting.Listener = &diffRecorder{}

// reconcileKey labels an object the same way the recorded diffs do, e.g. "ComputeInstance ns/name".
func reconcileKey(kind, namespace, name string) string {
	return fmt.Sprintf("%s %s/%s", kind, namespace, name)
}

func (r *diffRecorder) OnDiff(ctx context.Context, diff *structuredreporting.Diff) {
	r.DebugLogListener.OnDiff(ctx, diff)
	if diff == nil || diff.IsNewObject || diff.Controller != k8s.ReconcilerTypeTerraform || diff.Object == nil {
		return
	}
	var fields []string
	for _, f := range diff.Fields {
		fields = append(fields, fmt.Sprintf("%s: %v -> %v", f.ID, f.Old, f.New))
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.diffs = append(r.diffs, fmt.Sprintf("%s %s/%s: %v", diff.Object.GetKind(), diff.Object.GetNamespace(), diff.Object.GetName(), fields))
}

func (r *diffRecorder) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
	// Guard before forwarding: DebugLogListener dereferences u, which is nil on some error paths.
	if u == nil {
		return
	}
	r.DebugLogListener.OnReconcileStart(ctx, u, t)
	if t != k8s.ReconcilerTypeTerraform {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.starts == nil {
		r.starts = make(map[string]int)
	}
	r.starts[reconcileKey(u.GetKind(), u.GetNamespace(), u.GetName())]++
}

func (r *diffRecorder) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
	// Guard before forwarding: DebugLogListener dereferences u, which is nil on some error paths.
	if u == nil {
		return
	}
	r.DebugLogListener.OnReconcileEnd(ctx, u, result, err, t)
	if t != k8s.ReconcilerTypeTerraform {
		return
	}
	key := reconcileKey(u.GetKind(), u.GetNamespace(), u.GetName())
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.ends == nil {
		r.ends = make(map[string][]reconcileEnd)
	}
	r.ends[key] = append(r.ends[key], reconcileEnd{start: r.starts[key], result: result, err: err})
}

// Count returns the number of diffs recorded so far.
func (r *diffRecorder) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.diffs)
}

// Since returns the diffs recorded after the first n.
func (r *diffRecorder) Since(n int) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.diffs[n:]...)
}

// Reconciles returns how many Terraform reconciles have started for the object so far.
func (r *diffRecorder) Reconciles(key string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.starts[key]
}

// WaitForReconcileAfter waits until a Terraform reconcile of the object that started after the first `after`
// reconciles has ended. It returns the first such end and true, or false if the timeout passes first. Diffs are
// reported before the reconcile ends, so once this returns true, Since(...) holds the diffs of that reconcile.
//
// The end can be an early exit that never reached diff evaluation (see reconcileEnd.notEvaluatedReason), and then no
// diffs were recorded for it. The caller must reject such an end, otherwise a no-diff check passes without checking.
func (r *diffRecorder) WaitForReconcileAfter(key string, after int, timeout time.Duration) (reconcileEnd, bool) {
	deadline := time.Now().Add(timeout)
	for {
		r.mu.Lock()
		for _, end := range r.ends[key] {
			if end.start > after {
				r.mu.Unlock()
				return end, true
			}
		}
		r.mu.Unlock()
		if time.Now().After(deadline) {
			return reconcileEnd{}, false
		}
		time.Sleep(100 * time.Millisecond)
	}
}
