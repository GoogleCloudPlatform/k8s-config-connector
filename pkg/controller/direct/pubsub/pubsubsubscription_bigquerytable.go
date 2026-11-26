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

package pubsub

import (
	"strings"

	krmbigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

// ParsePubSubBigQueryTable parses a BigQuery table name in the formats that
// Pub/Sub accepts for bigquery_config.table: "{project}.{dataset}.{table}" and
// "{project}:{dataset}.{table}". Like Pub/Sub, it splits from the right, so
// domain-scoped project IDs such as "example.com:my-project" are supported.
func ParsePubSubBigQueryTable(table string) (*krmbigqueryv1beta1.BigQueryTableIdentity, bool) {
	tableSep := strings.LastIndexByte(table, '.')
	if tableSep < 0 {
		return nil, false
	}
	datasetSep := strings.LastIndexAny(table[:tableSep], ".:")
	if datasetSep < 0 {
		return nil, false
	}
	id := &krmbigqueryv1beta1.BigQueryTableIdentity{
		Project: table[:datasetSep],
		Dataset: table[datasetSep+1 : tableSep],
		Table:   table[tableSep+1:],
	}
	// IsValidSegment rejects empty parts and "/", so canonical resource names
	// such as "projects/p/datasets/d/tables/t" are never treated as this format.
	if !gcpurls.IsValidSegment(id.Project) || !gcpurls.IsValidSegment(id.Dataset) || !gcpurls.IsValidSegment(id.Table) {
		return nil, false
	}
	return id, true
}
