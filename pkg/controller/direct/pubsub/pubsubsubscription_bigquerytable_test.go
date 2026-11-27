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
	"testing"

	"github.com/google/go-cmp/cmp"

	krmbigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
)

func TestParsePubSubBigQueryTable(t *testing.T) {
	myTable := &krmbigqueryv1beta1.BigQueryTableIdentity{Project: "my-project", Dataset: "my_dataset", Table: "my_table"}
	domainScopedTable := &krmbigqueryv1beta1.BigQueryTableIdentity{Project: "example.com:my-project", Dataset: "my_dataset", Table: "my_table"}
	tests := []struct {
		name   string
		table  string
		wantID *krmbigqueryv1beta1.BigQueryTableIdentity
		wantOK bool
	}{
		{name: "project, dataset and table separated by dots", table: "my-project.my_dataset.my_table", wantID: myTable, wantOK: true},
		{name: "project separated from the dataset by a colon", table: "my-project:my_dataset.my_table", wantID: myTable, wantOK: true},
		{name: "domain-scoped project followed by a dot", table: "example.com:my-project.my_dataset.my_table", wantID: domainScopedTable, wantOK: true},
		{name: "domain-scoped project followed by a colon", table: "example.com:my-project:my_dataset.my_table", wantID: domainScopedTable, wantOK: true},
		{name: "empty value", table: ""},
		{name: "table without a dataset", table: "my_dataset.my_table"},
		{name: "empty project", table: ".my_dataset.my_table"},
		{name: "empty dataset", table: "my-project..my_table"},
		{name: "empty table", table: "my-project.my_dataset."},
		{name: "canonical resource name", table: "projects/my-project/datasets/my_dataset/tables/my_table"},
		{name: "resource name with the bigquery.googleapis.com host", table: "//bigquery.googleapis.com/projects/my-project/datasets/my_dataset/tables/my_table"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotID, gotOK := ParsePubSubBigQueryTable(tc.table)

			if gotOK != tc.wantOK {
				t.Errorf("ParsePubSubBigQueryTable(%q) ok = %v, want %v", tc.table, gotOK, tc.wantOK)
			}
			if diff := cmp.Diff(tc.wantID, gotID); diff != "" {
				t.Errorf("ParsePubSubBigQueryTable(%q) mismatch (-want +got):\n%s", tc.table, diff)
			}
		})
	}
}
