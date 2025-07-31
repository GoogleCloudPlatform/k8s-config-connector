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

package structuredreporting

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type contextKey string

const (
	listenerContextKey contextKey = "structuredreporting.listener"
)

func GetListenerFromContext(ctx context.Context) (Listener, bool) {
	if ctx == nil {
		return nil, false
	}
	if v := ctx.Value(listenerContextKey); v != nil {
		if l, ok := v.(Listener); ok {
			return l, true
		}
	}
	return nil, false
}

func ContextWithListener(ctx context.Context, listener Listener) context.Context {
	return context.WithValue(ctx, listenerContextKey, listener)
}

// Listener is implemented by listeners to the "structured reporting" event stream.
// Note: these methods are called from performance-sensitive code.
// If you are doing substantial processing (in production paths),
// consider copying the arguments and sending them to another goroutine for any heavy lifting.
type Listener interface {
	// OnError is called when a controller calls ReportError
	OnError(ctx context.Context, err error, args ...any)
	// OnDiff is called when a controller calls ReportDiffs
	OnDiff(ctx context.Context, diffs *Diff)
	// OnReconcileStart is called when a controller calls ReportReconcileStart
	OnReconcileStart(ctx context.Context, u *unstructured.Unstructured)
	// OnReconcileEnd is called when a controller calls ReportReconcileEnd
	OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error)
}
