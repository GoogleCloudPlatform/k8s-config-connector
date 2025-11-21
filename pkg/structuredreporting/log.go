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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type LogListener struct {
}

var _ Listener = &LogListener{}

// OnError is called when a controller calls ReportError
func (l *LogListener) OnError(ctx context.Context, err error, args ...any) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnError",
		"error", err)
}

// OnDiff is called when a controller calls ReportDiffs
func (l *LogListener) OnDiff(ctx context.Context, diffs *Diff) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnDiff",
		"diff.fields", diffs.Fields,
		"diff.isNewObject", diffs.IsNewObject,
	)
}

// OnReconcileStart is called when a controller calls ReportReconcileStart
func (l *LogListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnReconcileStart",
		"object.kind", u.GroupVersionKind().Kind,
		"object.name", u.GetName())
}

// OnReconcileEnd is called when a controller calls ReportReconcileEnd
func (l *LogListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnReconcileEnd",
		"object.kind", u.GroupVersionKind().Kind,
		"object.name", u.GetName(),
		"result", result,
		"error", err)
}
