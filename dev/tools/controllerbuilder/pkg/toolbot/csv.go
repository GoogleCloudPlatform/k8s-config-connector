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
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type CSVExporter struct {
	enhancers  []Enhancer
	extractor  Extractor
	dataPoints []*DataPoint
}

type Extractor interface {
	Extract(ctx context.Context, b []byte) ([]*DataPoint, error)
}

type Enhancer interface {
	EnhanceDataPoint(ctx context.Context, d *DataPoint) error
}

func NewCSVExporter(extractor Extractor, enhancers ...Enhancer) (*CSVExporter, error) {
	x := &CSVExporter{
		enhancers: enhancers,
		extractor: extractor,
	}

	return x, nil
}

func (x *CSVExporter) visitGoFile(ctx context.Context, p string) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}
	dataPoints, err := x.extractor.Extract(ctx, b)
	if err != nil {
		return err
	}
	x.dataPoints = append(x.dataPoints, dataPoints...)
	return nil
}


func (x *CSVExporter) VisitCodeDir(ctx context.Context, srcDir string) error {
	if err := filepath.WalkDir(srcDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".go":
			// OK
		default:
			return nil
		}
		// klog.Infof("%v", p)
		if err := x.visitGoFile(ctx, p); err != nil {
			return fmt.Errorf("processing file %q: %w", p, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}

func (x *CSVExporter) WriteCSVForAllTools(ctx context.Context, outputDir string) error {
	log := klog.FromContext(ctx)

	for _, dataPoint := range x.dataPoints {
		if err := x.EnhanceDataPoint(ctx, dataPoint); err != nil {
			return err
		}
	}

	toolNames := sets.NewString()
	for _, dataPoint := range x.dataPoints {
		toolNames.Insert(dataPoint.Type)
	}

	for _, toolName := range toolNames.List() {
		outFilePath := filepath.Join(outputDir, toolName+".csv")
		log.Info("writing CSV", "path", outFilePath)
		var bb bytes.Buffer
		if err := x.writeCSVForTool(ctx, toolName, &bb); err != nil {
			return err
		}
		if err := os.WriteFile(outFilePath, bb.Bytes(), 0644); err != nil {
			return fmt.Errorf("writing to file %q: %w", outFilePath, err)
		}
	}

	return nil
}

func (x *CSVExporter) writeCSVForTool(ctx context.Context, toolName string, out io.Writer) error {
	var dataPoints []*DataPoint
	for _, dataPoint := range x.dataPoints {
		if dataPoint.Type != toolName {
			continue
		}
		dataPoints = append(dataPoints, dataPoint)
	}

	var bb bytes.Buffer
	csvFile := csv.NewWriter(&bb)

	columnSet := sets.NewString()
	columnSet.Insert("out")

	for _, dataPoint := range dataPoints {
		for k := range dataPoint.Input {
			columnSet.Insert("in." + k)
		}
	}

	columns := columnSet.List()

	csvFile.Write(columns)
	csvFile.Flush()

	for _, dataPoint := range dataPoints {
		row := make([]string, len(columns))
		for i, column := range columns {
			switch column {
			case "out":
				row[i] = dataPoint.Output

			default:
				if strings.HasPrefix(column, "in.") {
					row[i] = dataPoint.Input[strings.TrimPrefix(column, "in.")]
				} else {
					return fmt.Errorf("unknown column %q", column)
				}
			}
		}
		csvFile.Write(row)
		csvFile.Flush()
	}

	if err := csvFile.Error(); err != nil {
		return fmt.Errorf("writing to csv: %w", err)
	}

	_, err := bb.WriteTo(out)
	if err != nil {
		return err
	}

	return nil
}

func (x *CSVExporter) EnhanceDataPoint(ctx context.Context, d *DataPoint) error {
	for _, enhancer := range x.enhancers {
		if err := enhancer.EnhanceDataPoint(ctx, d); err != nil {
			return err
		}
	}
	return nil
}

func (x *CSVExporter) BuildInputRow(ctx context.Context, b []byte, out io.Writer) error {
	dataPoints, err := x.extractor.Extract(ctx, b)
	if err != nil {
		return err
	}

	for _, dataPoint := range dataPoints {
		if err := x.EnhanceDataPoint(ctx, dataPoint); err != nil {
			return err
		}
	}

	toolNames := sets.NewString()
	for _, dataPoint := range dataPoints {
		toolNames.Insert(dataPoint.Type)
	}

	if toolNames.Len() != 1 {
		return fmt.Errorf("expected exactly one tool name, got %v", toolNames.List())
	}
	for _, toolName := range toolNames.List() {

		var bb bytes.Buffer

		columnSet := sets.NewString()

		for _, dataPoint := range dataPoints {
			if dataPoint.Type != toolName {
				continue
			}
			for k := range dataPoint.Input {
				columnSet.Insert("in." + k)
			}
		}

		columns := columnSet.List()

		for _, dataPoint := range dataPoints {
			if dataPoint.Type != toolName {
				continue
			}

			for _, column := range columns {
				if strings.HasPrefix(column, "in.") {
					fmt.Fprintf(&bb, "%v: %v\n", column, dataPoint.Input[strings.TrimPrefix(column, "in.")])
				} else {
					return fmt.Errorf("unknown column %q", column)
				}
			}
		}

		_, err := bb.WriteTo(out)
		if err != nil {
			return err
		}
	}

	return nil
}
