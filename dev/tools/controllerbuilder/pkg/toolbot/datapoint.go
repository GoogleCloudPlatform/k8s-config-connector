// Copyright 2024 Google LLC
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

package toolbot

import (
	"encoding/csv"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

// DataPoint holds the input and output for a tool.
type DataPoint struct {
	Description string
	Type        string
	Input       map[string]string
	Output      string
}

// InputColumnKeys returns a set of the input column names
func (p *DataPoint) InputColumnKeys() sets.Set[string] {
	keys := sets.New[string]()
	for k := range p.Input {
		keys.Insert(k)
	}
	return keys
}

// SetInput sets an input value for the data point.
func (p *DataPoint) SetInput(k, v string) {
	if p.Input == nil {
		p.Input = make(map[string]string)
	}
	p.Input[k] = v
}

// AddCSVColumns adds the columns for the data point to the columnSet.
func (p *DataPoint) AddCSVColumns(columnSet sets.Set[string]) {
	if p.Output != "" {
		columnSet.Insert("out")
	}

	for k := range p.Input {
		columnSet.Insert("in." + k)
	}
}

// WriteCSV writes the data point to the CSV writer.
func (p *DataPoint) WriteCSV(csvWriter *csv.Writer, columns []string) error {
	row := make([]string, len(columns))
	for i, column := range columns {
		switch column {
		case "out":
			row[i] = p.Output

		default:
			if strings.HasPrefix(column, "in.") {
				row[i] = p.Input[strings.TrimPrefix(column, "in.")]
			} else {
				return fmt.Errorf("unknown column %q", column)
			}
		}
	}
	return csvWriter.Write(row)
}

// ToGenAIFormat converts the data point to the input format for Gemini (or other LLMs).
func (p *DataPoint) ToGenAIFormat() string {
	columnSet := sets.NewString()
	if p.Output != "" {
		columnSet.Insert("out")
	}
	for k := range p.Input {
		columnSet.Insert("in." + k)
	}

	var part strings.Builder
	columns := columnSet.List()

	for _, column := range columns {
		v := ""

		switch column {
		case "out":
			v = p.Output

		default:
			if strings.HasPrefix(column, "in.") {
				v = p.Input[strings.TrimPrefix(column, "in.")]
			} else {
				klog.Fatalf("unknown column %q", column)
			}
		}

		if strings.Contains(v, "\n") {
			fmt.Fprintf(&part, "<%s>\n%s\n</%s>\n", column, v, column)

		} else {
			fmt.Fprintf(&part, "<%s>%s</%s>\n", column, v, column)
		}
	}

	return part.String()
}
