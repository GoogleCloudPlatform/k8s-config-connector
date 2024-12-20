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

	"cloud.google.com/go/vertexai/genai"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

// DataPoint holds the input and output for a tool.
type DataPoint struct {
	Type   string
	Input  map[string]string
	Output string
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

// ToGenAIParts converts the data point to the input format for Gemini.
func (p *DataPoint) ToGenAIParts() []genai.Part {
	columnSet := sets.NewString()
	if p.Output != "" {
		columnSet.Insert("out")
	}
	for k := range p.Input {
		columnSet.Insert("in." + k)
	}

	var parts []genai.Part
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

		s := fmt.Sprintf("%s %s", column, v)
		parts = append(parts, genai.Text(s))
	}

	return parts
}
