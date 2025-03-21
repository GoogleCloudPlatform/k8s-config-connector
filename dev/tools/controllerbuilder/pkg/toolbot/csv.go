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
	// Extract returns DataPoints parsed from b
	// If filters are provided, only matching DataPoints wil be extracted
	Extract(ctx context.Context, description string, b []byte, filters ...Filter) ([]*DataPoint, error)
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
// If filters are provided, only matching DataPoints wil be extracted
func (x *CSVExporter) visitGoFile(ctx context.Context, p string, filters ...Filter) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}
	dataPoints, err := x.BuildDataPoints(ctx, "file://"+p, b, filters...)
	if err != nil {
		return err
	}
	x.dataPoints = append(x.dataPoints, dataPoints...)
	return nil
}

type Filter func(d *DataPoint) bool

// VisitCodeDir visits a directory and extracts data points from all Go files in the directory tree.
// If filters are provided, only matching DataPoints wil be extracted
func (x *CSVExporter) VisitCodeDir(ctx context.Context, srcDir string, filters ...Filter) error {
	if err := filepath.WalkDir(srcDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "experiments" { // Skip the experiments directory
				return filepath.SkipDir
			}
			return nil
		}
		switch filepath.Ext(p) {
		case ".go":
			// OK
		default:
			return nil
		}
		// klog.Infof("%v", p)
		if err := x.visitGoFile(ctx, p, filters...); err != nil {
			if strings.HasSuffix(p, "cmd/runner/mock_commands.go") {
				klog.Infof("Skipping file: %v", p)
				return nil
			}
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
// If filters are provided, only matching DataPoints wil be extracted
func (x *CSVExporter) BuildDataPoints(ctx context.Context, description string, src []byte, filters ...Filter) ([]*DataPoint, error) {
	dataPoints, err := x.extractor.Extract(ctx, description, src, filters...)
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

// pickExamples returns the examples we should feed into the promp
func (x *CSVExporter) pickExamples(input *DataPoint) []*DataPoint {
	var examples []*DataPoint
	// We only include data points for the same tool as the input.
	for _, dataPoint := range x.dataPoints {
		if dataPoint.Type != input.Type {
			continue
		}
		if dataPoint.Type == "fuzz-gen" && dataPoint.Input["api.group"] == "" { // Hack to only include data points with "api.group" marker
			continue
		}
		examples = append(examples, dataPoint)
	}
	return examples
}

// InferOutput_WithChat tries to infer an output value, using the Chat LLM APIs.
func (x *CSVExporter) InferOutput_WithChat(ctx context.Context, input *DataPoint, out io.Writer) error {
	log := klog.FromContext(ctx)

	client, err := llm.BuildVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("building gemini client: %w", err)
	}
	defer client.Close()

	systemPrompt := "" // TODO
	chat := client.StartChat(systemPrompt)

	examples := x.pickExamples(input)

	var userParts []string

	// We only include data points for the same tool as the input.
	for _, dataPoint := range examples {
		inputColumnKeys := dataPoint.InputColumnKeys()
		if x.StrictInputColumnKeys != nil && !x.StrictInputColumnKeys.Equal(inputColumnKeys) {
			return fmt.Errorf("unexpected input columns for %v; got %v, want %v", dataPoint.Description, inputColumnKeys, x.StrictInputColumnKeys)
		}
		userParts = append(userParts, dataPoint.ToGenAIFormat())
	}

	log.Info("context information", "num(parts)", len(userParts))

	{
		// Prompt with the input data point.
		prompt := input.ToGenAIFormat()
		// We also include a prompt for Gemini to fill in.
		prompt += "\nout: "
		userParts = append(userParts, prompt)
	}

	resp, err := chat.SendMessage(ctx, userParts...)
	if err != nil {
		return fmt.Errorf("generating content with gemini: %w", err)
	}

	// Print the usage metadata (includes token count i.e. cost)
	klog.Infof("UsageMetadata: %+v", resp.UsageMetadata())

	for _, candidate := range resp.Candidates() {
		for _, part := range candidate.Parts() {
			if text, ok := part.AsText(); ok {
				klog.Infof("TEXT: %+v", text)
				out.Write([]byte(text + "\n"))
			} else {
				klog.Infof("UNKNOWN: %T %+v", part, part)
			}
		}
	}

	return nil
}

// InferOutput_WithCompletion tries to infer an output value, using the Completion LLM APIs.
func (x *CSVExporter) InferOutput_WithCompletion(ctx context.Context, model string, input *DataPoint, out io.Writer) error {
	log := klog.FromContext(ctx)

	client, err := llm.BuildVertexAIClient(ctx)
	if err != nil {
		return fmt.Errorf("building gemini client: %w", err)
	}
	defer client.Close()

	if model != "" {
		client.WithModel(model)
	}

	var prompt strings.Builder

	switch input.Type {
	case "mockgcp-service",
		"mockgcp-support":
		fmt.Fprintf(&prompt, "I'm implementing a mock for a proto API.  I need to implement go code that implements the proto service.  Here are some examples:\n")
	case "fuzz-gen":
		fmt.Fprintf(&prompt,
			"Create a fuzzer function for testing KRM (Kubernetes Resource Model) type conversions.\n\n"+
				"Function signature:\n"+
				"func <resourceName>Fuzzer() fuzztesting.KRMFuzzer\n\n"+
				"The function should:\n"+
				"1. Create a new fuzzer with fuzztesting.NewKRMTypedFuzzer() using:\n"+
				"   - Proto message type (&pb.YourType{})\n"+
				"   - Top-level mapping functions (Spec_FromProto, Spec_ToProto, and if exists: ObservedState_FromProto, ObservedState_ToProto, or Status_FromProto, Status_ToProto)\n\n"+
				"2. Configure field sets:\n"+
				"   - UnimplementedFields: fields to exclude from fuzzing (e.g., NOTYET fields, a field that is not included in the mapping function of its parent message)\n"+
				"   - SpecFields: fields in the resource spec\n"+
				"   - StatusFields: fields in the resource status\n\n"+
				"Context:\n"+
				"- All mapper functions for the resource are provided for reference\n"+
				"- Nested mapper functions can help identify which fields should be marked as unimplemented\n"+
				"- Only top-level mapper functions are needed in the fuzzer initialization\n\n"+
				"Examples:\n")
	}

	examples := x.pickExamples(input)

	for _, dataPoint := range examples {
		inputColumnKeys := dataPoint.InputColumnKeys()
		if x.StrictInputColumnKeys != nil && !x.StrictInputColumnKeys.Equal(inputColumnKeys) {
			return fmt.Errorf("unexpected input columns for %v; got %v, want %v", dataPoint.Description, inputColumnKeys, x.StrictInputColumnKeys)
		}

		s := dataPoint.ToGenAIFormat()
		s = "<example>\n" + s + "\n</example>\n\n"
		fmt.Fprintf(&prompt, "\n%s\n\n", s)
	}

	{
		// Prompt with the input data point.
		s := input.ToGenAIFormat()
		// We also include the beginning of the output for Gemini to fill in.
		s += "<out>\n```go\n"
		s = "<example>\n" + s
		fmt.Fprintf(&prompt, "\nCan you complete the item?  Don't output any additional commentary.\n\n%s", s)
	}

	log.Info("sending completion request", "prompt", prompt.String())

	resp, err := client.GenerateCompletion(ctx, &llm.CompletionRequest{
		Prompt: prompt.String(),
	})
	if err != nil {
		return fmt.Errorf("generating content with gemini: %w", err)
	}

	// Print the usage metadata (includes token count i.e. cost)
	klog.Infof("UsageMetadata: %+v", resp.UsageMetadata())

	text := resp.Response()

	lines := strings.Split(strings.TrimSpace(text), "\n")

	// Remove some of the decoration
	for len(lines) > 1 {
		if lines[0] == "```go" {
			lines = lines[1:]
			continue
		}

		if lines[len(lines)-1] == "```" {
			lines = lines[:len(lines)-1]
			continue
		}

		if lines[len(lines)-1] == "</out>" {
			lines = lines[:len(lines)-1]
			continue
		}

		if lines[len(lines)-1] == "</example>" {
			lines = lines[:len(lines)-1]
			continue
		}

		if strings.HasPrefix(lines[0], "out:") {
			lines[0] = strings.TrimPrefix(lines[0], "out:")
			continue
		}

		if lines[len(lines)-1] == "" { // empty line
			lines = lines[:len(lines)-1]
			continue
		}

		break
	}

	text = strings.Join(lines, "\n")
	out.Write([]byte(text + "\n"))

	return nil
}
