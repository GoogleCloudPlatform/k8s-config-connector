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

package metrics

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/errors"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/metrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ResourceNameLabel bool

type ReconcilerMetrics struct {
	// atomic counter for occupied workers
	occupiedWorkers   int64
	ResourceNameLabel bool
}

func (r *ReconcilerMetrics) RecordReconcileWorkers(ctx context.Context, gvk schema.GroupVersionKind) {
	atomic.AddInt64(&r.occupiedWorkers, 1)
	openCensusContext, _ := tag.New(ctx, tag.Insert(metrics.KindTag, gvk.GroupKind().String()))
	stats.Record(openCensusContext, metrics.MReconcileTotalWorkers.M(k8s.ControllerMaxConcurrentReconciles))
	stats.Record(openCensusContext, metrics.MReconcileOccupiedWorkers.M(atomic.LoadInt64(&r.occupiedWorkers)))
}

func (r *ReconcilerMetrics) RecordReconcileMetrics(ctx context.Context, gvk schema.GroupVersionKind, ns, name string, startTime time.Time, reconcileErr *error) {
	if reconcileErr == nil {
		log.Println("ERROR: the pointer to reconcile error is nil. Skip recording reconcile metrics")
		return
	}
	status := "OK"
	if *reconcileErr != nil {
		status = "ERROR"
	}
	openCensusContext, _ := tag.New(ctx, tag.Insert(metrics.KindTag, gvk.GroupKind().String()), tag.Insert(metrics.NamespaceTag, ns), tag.Insert(metrics.StatusTag, status))
	if r.ResourceNameLabel {
		openCensusContext, _ = tag.New(openCensusContext, tag.Insert(metrics.ResourceNameTag, name))
	}
	stats.Record(openCensusContext, metrics.MReconcileRequests.M(1), metrics.MReconcileDuration.M(time.Since(startTime).Seconds()))
	r.RecordInternalErrors(ctx, gvk, ns, reconcileErr)
}

func (r *ReconcilerMetrics) RecordInternalErrors(ctx context.Context, gvk schema.GroupVersionKind, ns string, reconcileErr *error) {
	if reconcileErr == nil {
		log.Println("ERROR: the pointer to reconcile error is nil. Skip recording reconcile metrics")
		return
	}
	err := *reconcileErr
	if err == nil {
		return
	}
	if _, ok := errors.AsInternalError(err); !ok {
		return
	}
	openCensusContext, _ := tag.New(ctx, tag.Insert(metrics.KindTag, gvk.GroupKind().String()), tag.Insert(metrics.NamespaceTag, ns))
	stats.Record(openCensusContext, metrics.MInternalErrors.M(1))
}

func (r *ReconcilerMetrics) AfterReconcile() {
	atomic.AddInt64(&r.occupiedWorkers, -1)
}
