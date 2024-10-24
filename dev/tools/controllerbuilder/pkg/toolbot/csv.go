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
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
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

	columnSet := sets.NewString()
	for _, dataPoint := range dataPoints {
		dataPoint.AddCSVColumns(columnSet)
	}

	columns := columnSet.List()

	csvFile := csv.NewWriter(out)

	// write the CSV header
	csvFile.Write(columns)

	for _, dataPoint := range dataPoints {
		if err := dataPoint.WriteCSV(csvFile, columns); err != nil {
			return err
		}
	}

	csvFile.Flush()

	if err := csvFile.Error(); err != nil {
		return fmt.Errorf("writing to csv: %w", err)
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

func (x *CSVExporter) BuildDataPoints(ctx context.Context, b []byte) ([]*DataPoint, error) {
	dataPoints, err := x.extractor.Extract(ctx, b)
	if err != nil {
		return nil, err
	}

	for _, dataPoint := range dataPoints {
		if err := x.EnhanceDataPoint(ctx, dataPoint); err != nil {
			return nil, err
		}
	}

	return dataPoints, nil
}

func (x *CSVExporter) RunGemini(ctx context.Context, input *DataPoint, out io.Writer) error {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return fmt.Errorf("building gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro-002")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	var parts []genai.Part

	for _, dataPoint := range x.dataPoints {
		if dataPoint.Type != input.Type {
			continue
		}
		parts = append(parts, dataPoint.ToGenAIParts()...)
	}

	parts = append(parts, input.ToGenAIParts()...)
	parts = append(parts, genai.Text("out "))

	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return fmt.Errorf("generating content with gemini: %w", err)
	}

	klog.Infof("UsageMetadata: %+v", resp.UsageMetadata)

	for _, candidate := range resp.Candidates {
		content := candidate.Content

		for _, part := range content.Parts {
			if text, ok := part.(genai.Text); ok {
				klog.Infof("TEXT: %+v", text)
			} else {
				klog.Infof("UNKNOWN: %T %+v", part, part)
			}
		}
	}

	return nil
}
