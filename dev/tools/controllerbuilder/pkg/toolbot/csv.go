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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

// CSVExporter is an exporter that writes CSV files for each tool.
type CSVExporter struct {
	enhancers  []Enhancer
	extractor  Extractor
	dataPoints []*DataPoint

	// StrictInputColumnKeys ensures that all input datapoints have this shape.
	// This helps detect typos in the examples.
	StrictInputColumnKeys sets.Set[string]
}

// Extractor is an interface for extracting data points from source code.
type Extractor interface {
	Extract(ctx context.Context, description string, b []byte) ([]*DataPoint, error)
}

// Enhancer is an interface for enhancing a data point.
// For example, it might add a computed field to the data point, such as the definition of a proto message,
// given the name of the proto message.
type Enhancer interface {
	EnhanceDataPoint(ctx context.Context, d *DataPoint) error
}

// NewCSVExporter creates a new CSVExporter.
func NewCSVExporter(extractor Extractor, enhancers ...Enhancer) (*CSVExporter, error) {
	x := &CSVExporter{
		enhancers: enhancers,
		extractor: extractor,
	}

	return x, nil
}

// visitGoFile visits a Go file and extracts data points from it.
func (x *CSVExporter) visitGoFile(ctx context.Context, p string) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}
	dataPoints, err := x.BuildDataPoints(ctx, "file://"+p, b)
	if err != nil {
		return err
	}
	x.dataPoints = append(x.dataPoints, dataPoints...)
	return nil
}

// VisitCodeDir visits a directory and extracts data points from all Go files in the directory tree.
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

// WriteCSVForAllTools writes CSV files for all tools.
func (x *CSVExporter) WriteCSVForAllTools(ctx context.Context, outputDir string) error {
	log := klog.FromContext(ctx)

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

// writeCSVForTool writes a CSV file for a single tool.
func (x *CSVExporter) writeCSVForTool(ctx context.Context, toolName string, out io.Writer) error {
	var dataPoints []*DataPoint
	for _, dataPoint := range x.dataPoints {
		if dataPoint.Type != toolName {
			continue
		}
		dataPoints = append(dataPoints, dataPoint)
	}

	columnSet := sets.New[string]()
	for _, dataPoint := range dataPoints {
		dataPoint.AddCSVColumns(columnSet)
	}

	columns := sets.List(columnSet)

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

// EnhanceDataPoint enhances a data point by running all the registered enhancers.
func (x *CSVExporter) EnhanceDataPoint(ctx context.Context, d *DataPoint) error {
	for _, enhancer := range x.enhancers {
		if err := enhancer.EnhanceDataPoint(ctx, d); err != nil {
			return err
		}
	}
	return nil
}

// BuildDataPoints extracts data points from a byte slice representing a Go file.
func (x *CSVExporter) BuildDataPoints(ctx context.Context, description string, src []byte) ([]*DataPoint, error) {
	dataPoints, err := x.extractor.Extract(ctx, description, src)
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

// RunGemini runs a prompt against Gemini, generating context based on the source code.
func (x *CSVExporter) RunGemini(ctx context.Context, input *DataPoint, out io.Writer) error {
	log := klog.FromContext(ctx)

	client, err := llm.NewLLMClientFromEnvVar(ctx) //BuildVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("building gemini client: %w", err)
	}
	defer client.Close()

	var prompt strings.Builder

	// systemPrompt := "" // TODO
	// chat := client.StartChat(systemPrompt)

	// var userParts []string

	// userParts = append(userParts, "I'm implementing a mock for a proto API.  I need to implement go code that implements the proto service.  Here are some examples:")

	fmt.Fprintf(&prompt, "I'm implementing a mock for a proto API.  I need to implement go code that implements the proto service.  Here are some examples:\n")
	// We only include data points for the same tool as the input.
	for _, dataPoint := range x.dataPoints {
		if dataPoint.Type != input.Type {
			continue
		}

		inputColumnKeys := dataPoint.InputColumnKeys()
		if x.StrictInputColumnKeys != nil && !x.StrictInputColumnKeys.Equal(inputColumnKeys) {
			return fmt.Errorf("unexpected input columns for %v; got %v, want %v", dataPoint.Description, inputColumnKeys, x.StrictInputColumnKeys)
		}

		s := dataPoint.ToGenAIFormat()
		// s = "<example>\n" + s + "\n</example>\n\n"
		// s += "\n---\n\n"
		fmt.Fprintf(&prompt, "\n%s\n\n", s)
		// userParts = append(userParts, s)
	}

	// log.Info("context information", "num(parts)", len(userParts))

	{
		// Prompt with the input data point.
		s := input.ToGenAIFormat()
		// We also include a prompt for Gemini to fill in.
		s += "\nout: ```go\n"
		// s = "<example>\n" + s
		// s = "Can you help me implement this?\n" + s
		// userParts = append(userParts, s)
		fmt.Fprintf(&prompt, "\nCan you help me implement this?\n\n%s", s)
	}

	log.Info("sending completion request", "prompt", prompt.String())

	// resp, err := chat.SendMessage(ctx, userParts...)
	// if err != nil {
	// 	return fmt.Errorf("generating content with gemini: %w", err)
	// }
	resp, err := client.GenerateCompletion(ctx, &llm.CompletionRequest{
		Prompt: prompt.String(),
	})
	if err != nil {
		return fmt.Errorf("generating content with gemini: %w", err)
	}

	// Print the usage metadata (includes token count i.e. cost)
	klog.Infof("UsageMetadata: %+v", resp.UsageMetadata())

	// for _, candidate := range resp.Candidates() {
	// 	for _, part := range candidate.Parts() {
	// 		if text, ok := part.AsText(); ok {
	// 			lines := strings.Split(strings.TrimSpace(text), "\n")
	// 			if len(lines) > 2 {
	// 				if lines[0] == "```go" {
	// 					lines = lines[1:]
	// 				}
	// 				if lines[len(lines)-1] == "```" {
	// 					lines = lines[:len(lines)-1]
	// 				}
	// 			}
	// 			text = strings.Join(lines, "\n")
	// 			// klog.Infof("TEXT: %+v", text)
	// 			out.Write([]byte(text + "\n"))
	// 		} else {
	// 			klog.Infof("UNKNOWN: %T %+v", part, part)
	// 		}
	// 	}
	// }

	text := resp.Response()

	lines := strings.Split(strings.TrimSpace(text), "\n")
	if len(lines) > 2 {
		if lines[0] == "```go" {
			lines = lines[1:]
		}
		if lines[len(lines)-1] == "```" {
			lines = lines[:len(lines)-1]
		}
	}
	text = strings.Join(lines, "\n")
	// klog.Infof("TEXT: %+v", text)
	out.Write([]byte(text + "\n"))

	return nil
}
