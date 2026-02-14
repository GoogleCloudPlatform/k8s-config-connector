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
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	processStartTime = &view.View{
		Name:        "process_start_time_seconds",
		Measure:     MProcessStartTime,
		Description: MProcessStartTime.Description(),
		Aggregation: view.LastValue(),
	}
	sharedControllerViews = []*view.View{
		{
			Name:        "reconcile_occupied_workers_total",
			Measure:     MReconcileOccupiedWorkers,
			Description: MReconcileOccupiedWorkers.Description(),
			TagKeys:     []tag.Key{KindTag},
			Aggregation: view.LastValue(),
		},
		{
			Name:        "reconcile_workers_total",
			Measure:     MReconcileTotalWorkers,
			Description: MReconcileTotalWorkers.Description(),
			TagKeys:     []tag.Key{KindTag},
			Aggregation: view.LastValue(),
		},
		{
			Name:        "internal_errors_total",
			Measure:     MInternalErrors,
			Description: MInternalErrors.Description(),
			TagKeys:     []tag.Key{KindTag, NamespaceTag},
			Aggregation: view.Count(),
		},
		processStartTime,
	}
	controllerViewsWithResourceNameLabel = []*view.View{
		{
			Name:        "reconcile_requests_total",
			Measure:     MReconcileRequests,
			Description: MReconcileRequests.Description(),
			TagKeys:     []tag.Key{KindTag, NamespaceTag, StatusTag, ResourceNameTag},
			Aggregation: view.Count(),
		},
		{
			Name:        "reconcile_request_duration_seconds",
			Measure:     MReconcileDuration,
			Description: MReconcileDuration.Description(),
			TagKeys:     []tag.Key{KindTag, NamespaceTag, StatusTag, ResourceNameTag},
			// Latency in buckets:
			// [>=0s, >=5s, >=10s, >=25s, >=1min, >=5min, >=10min, >=15min, >=30min, >=45min, >1h]
			Aggregation: view.Distribution(0, 5, 10, 25, 60, 5*60, 10*60, 15*60, 30*60, 45*60, 60*60),
		},
	}
)

func GetControllerViewsWithResourceNameLabel() []*view.View {
	views := make([]*view.View, 0)
	views = append(views, sharedControllerViews...)
	views = append(views, controllerViewsWithResourceNameLabel...)
	return views
}
