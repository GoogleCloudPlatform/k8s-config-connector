package toolbot

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type DataPoint struct {
	Type   string
	Input  map[string]string
	Output string
}

func (p *DataPoint) SetInput(k, v string) {
	if p.Input == nil {
		p.Input = make(map[string]string)
	}
	p.Input[k] = v
}

func (p *DataPoint) AddCSVColumns(columnSet sets.String) {
	if p.Output != "" {
		columnSet.Insert("out")
	}

	for k := range p.Input {
		columnSet.Insert("in." + k)
	}
}

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
