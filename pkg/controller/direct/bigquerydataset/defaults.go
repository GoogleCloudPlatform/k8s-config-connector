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

package bigquerydataset

import (
	"strconv"
	"time"

	bigquery "cloud.google.com/go/bigquery"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

const defaultMaxTimeTravelHours = "168" // 7 days (in hours)

func ApplyBigQueryDatasetGCPDefaults(mapCtx *direct.MapContext, in *krm.BigQueryDatasetSpec, out *bigquery.DatasetMetadata, actual *bigquery.DatasetMetadata) {
	if in == nil {
		return
	}
	if in.MaxTimeTravelHours == nil {
		maxHours, _ := strconv.ParseInt(defaultMaxTimeTravelHours, 10, 64)
		out.MaxTimeTravel = time.Duration(maxHours) * time.Hour
	}
	if in.Access == nil && actual != nil {
		out.Access = actual.Access
	}
}
