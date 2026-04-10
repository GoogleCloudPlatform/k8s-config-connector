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

package bigquery

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	bigquery "google.golang.org/api/bigquery/v2"
)

func TestTableEq(t *testing.T) {
	t.Run("basic fields", func(t *testing.T) {
		a := &bigquery.Table{
			Description:  "description a",
			FriendlyName: "friendly name a",
		}
		b := &bigquery.Table{
			Description:  "description b",
			FriendlyName: "friendly name b",
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if eq {
			t.Errorf("expected TableEq to return false, got true")
		}

		expectedDiffCount := 2
		if len(diff.Fields) != expectedDiffCount {
			t.Errorf("expected %d diff fields, got %d", expectedDiffCount, len(diff.Fields))
		}
	})

	t.Run("nested schema fields", func(t *testing.T) {
		a := &bigquery.Table{
			Schema: &bigquery.TableSchema{
				Fields: []*bigquery.TableFieldSchema{
					{
						Name:        "field1",
						Description: "desc1 a",
						Type:        "STRING",
					},
				},
			},
		}
		b := &bigquery.Table{
			Schema: &bigquery.TableSchema{
				Fields: []*bigquery.TableFieldSchema{
					{
						Name:        "field1",
						Description: "desc1 b",
						Type:        "INTEGER",
					},
				},
			},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if eq {
			t.Errorf("expected TableEq to return false, got true")
		}

		expectedDiffCount := 2
		if len(diff.Fields) != expectedDiffCount {
			t.Errorf("expected %d diff fields, got %d. Diffs: %+v", expectedDiffCount, len(diff.Fields), diff.Fields)
		}
	})

	t.Run("labels nil vs empty", func(t *testing.T) {
		a := &bigquery.Table{
			Labels: nil,
		}
		b := &bigquery.Table{
			Labels: map[string]string{},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if !eq {
			t.Errorf("expected TableEq to return true for nil vs empty labels, got false")
		}
	})

	t.Run("policy tags nil vs empty Names", func(t *testing.T) {
		a := &bigquery.TableFieldSchemaPolicyTags{
			Names: nil,
		}
		b := &bigquery.TableFieldSchemaPolicyTags{
			Names: []string{},
		}

		if !policyTagsEqual(a, b) {
			t.Errorf("expected policyTagsEqual to return true for nil vs empty Names, got false")
		}
	})

	t.Run("time partitioning missing check", func(t *testing.T) {
		a := &bigquery.Table{
			TimePartitioning: &bigquery.TimePartitioning{
				Type: "DAY",
			},
		}
		b := &bigquery.Table{
			TimePartitioning: &bigquery.TimePartitioning{
				Type: "MONTH",
			},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if eq {
			t.Errorf("expected TableEq to return false for different TimePartitioning, got true")
		}
	})

	t.Run("table constraints sorting", func(t *testing.T) {
		a := &bigquery.Table{
			TableConstraints: &bigquery.TableConstraints{
				ForeignKeys: []*bigquery.TableConstraintsForeignKeys{
					{Name: "fk1"},
					{Name: "fk2"},
				},
			},
		}
		b := &bigquery.Table{
			TableConstraints: &bigquery.TableConstraints{
				ForeignKeys: []*bigquery.TableConstraintsForeignKeys{
					{Name: "fk2"},
					{Name: "fk1"},
				},
			},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if !eq {
			t.Errorf("expected TableEq to return true for reordered foreign keys, got false. Diffs: %+v", diff.Fields)
		}
	})

	t.Run("encryption configuration nil vs empty", func(t *testing.T) {
		a := &bigquery.Table{
			EncryptionConfiguration: nil,
		}
		b := &bigquery.Table{
			EncryptionConfiguration: &bigquery.EncryptionConfiguration{
				KmsKeyName: "",
			},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if !eq {
			t.Errorf("expected TableEq to return true for nil vs empty KmsKeyName, got false. Diffs: %+v", diff.Fields)
		}
	})

	t.Run("csv options defaults", func(t *testing.T) {
		a := &bigquery.Table{
			ExternalDataConfiguration: &bigquery.ExternalDataConfiguration{
				CsvOptions: &bigquery.CsvOptions{
					FieldDelimiter: ",",
					Encoding:       "UTF8",
				},
			},
		}
		b := &bigquery.Table{
			ExternalDataConfiguration: &bigquery.ExternalDataConfiguration{
				CsvOptions: &bigquery.CsvOptions{
					// Defaults from mapper should match
				},
			},
		}

		diff := &structuredreporting.Diff{}
		eq, err := TableEq(a, b, diff)
		if err != nil {
			t.Fatalf("TableEq failed: %v", err)
		}

		if !eq {
			t.Errorf("expected TableEq to return true for default csv options, got false. Diffs: %+v", diff.Fields)
		}
	})
}
