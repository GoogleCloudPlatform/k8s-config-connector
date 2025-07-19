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
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/stats"
)

var (
	namespace = "configconnector"
	labels    = []string{"namespace", "group_version_kind", "status"}
)

// metrics defined in the format of OpenCensus
var (
	MReconcileOccupiedWorkers = stats.Int64("ReconcileOccupiedWorkers", "The number of occupied reconcile workers", stats.UnitDimensionless)
	MReconcileTotalWorkers    = stats.Int64("ReconcileTotalWorkers", "The number of total reconcile workers", stats.UnitDimensionless)
	MReconcileRequests        = stats.Int64("ReconcileRequests", "WARNING: do not import into GKE; unbound cardinality; The number of reconcile requests", stats.UnitDimensionless)
	MInternalErrors           = stats.Int64("InternalErrorsTotal", "WARNING: do not import into GKE; unbound cardinality; The number of internal errors", stats.UnitDimensionless)
	MReconcileDuration        = stats.Float64("ReconcileDuration", "WARNING: do not import into GKE; unbound cardinality; The duration of reconcile requests", "seconds")
	MProcessStartTime         = stats.Float64("ProcessStartTimeSeconds", "Start time of the process since unix epoch in seconds", "seconds")
)

// metrics defined in the format of prometheus/client_golang
func NewAppliedResourcesCollector() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "applied_resources_total",
		Help:      "The number of applied resources",
	}, labels)
}

func NewBuildInfoCollector(version string) prometheus.Collector {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "build_info",
			Help: fmt.Sprintf(
				"A metric with a constant '1' value labeled by version from which %s was built.",
				namespace,
			),
			ConstLabels: prometheus.Labels{
				"version": version,
			},
		},
		func() float64 { return 1 },
	)
}
