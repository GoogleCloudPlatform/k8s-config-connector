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

package exportcsv

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	kccio "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/io"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/toolbot"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
)

// PromptOptions are the options for the prompt command.
type PromptOptions struct {
	*options.GenerateOptions

	ProtoDir  string
	SrcDir    string
	Output    string
	InputFile string

	// StrictInputColumnKeys ensures that all input datapoints have this shape.
	// This helps detect typos in the examples.
	StrictInputColumnKeys []string
}

func (o *PromptOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.SrcDir = root
	o.ProtoDir = filepath.Join(root, ".build/third_party/googleapis/google")
	return nil
}

// BindFlags binds the flags to the command.
func (o *PromptOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.SrcDir, "src-dir", o.SrcDir, "base directory for source code")
	cmd.Flags().StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
	cmd.Flags().StringVar(&o.Output, "output", o.Output, "the directory to store the prompt outcome")
	cmd.Flags().StringVar(&o.InputFile, "input-file", o.InputFile, "the input file to get input from")
	cmd.Flags().StringSliceVar(&o.StrictInputColumnKeys, "strict-input-columns", o.StrictInputColumnKeys, "return an error if we see an irregular datapoint for this tool")
}

// BuildPromptCommand builds the `prompt` command.
func BuildPromptCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &PromptOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "prompt",
		Short: "executes a prompt against Gemini, generating context based on the source code.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunPrompt(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

// RunPrompt runs the `prompt` command.
func RunPrompt(ctx context.Context, o *PromptOptions) error {
	log := klog.FromContext(ctx)

	if err := rewriteFilePath(&o.ProtoDir); err != nil {
		return err
	}

	if o.ProtoDir == "" {
		return fmt.Errorf("--proto-dir is required")
	}

	extractor := &toolbot.ExtractToolMarkers{}
	addProtoDefinition, err := toolbot.NewEnhanceWithProtoDefinition(o.ProtoDir)
	if err != nil {
		return err
	}
	apiDir := o.SrcDir + "/apis/"
	addGoStruct, err := toolbot.NewEnhanceWithGoStruct(apiDir)
	if err != nil {
		return err
	}
	mapperDir := o.SrcDir + "/pkg/controller/direct/" // direct controller directory contains all mapper functions
	addMapperFunctions, err := toolbot.NewEnhanceWithMappers(mapperDir)
	if err != nil {
		return err
	}
	x, err := toolbot.NewCSVExporter(extractor, addProtoDefinition, addGoStruct, addMapperFunctions)
	if err != nil {
		return err
	}

	if len(o.StrictInputColumnKeys) != 0 {
		x.StrictInputColumnKeys = sets.New(o.StrictInputColumnKeys...)
	}

	var b []byte
	if o.InputFile == "" {
		if b, err = io.ReadAll(os.Stdin); err != nil {
			return fmt.Errorf("reading from stdin: %w", err)
		}
	} else {
		if b, err = os.ReadFile(o.InputFile); err != nil {
			return fmt.Errorf("reading from %s: %w", o.InputFile, err)
		}
	}

	dataPoints, err := x.BuildDataPoints(ctx, "<prompt>", b)
	if err != nil {
		return err
	}

	if len(dataPoints) != 1 {
		return fmt.Errorf("expected exactly one data point, got %d", len(dataPoints))
	}

	dataPoint := dataPoints[0]
	dataPoint.Output = ""

	log.Info("built data point", "dataPoint", dataPoint)

	if o.SrcDir != "" {
		filterByType := func(p *toolbot.DataPoint) bool {
			return p.Type == dataPoint.Type
		}
		if err := x.VisitCodeDir(ctx, o.SrcDir, filterByType); err != nil {
			return err
		}
	}

	model := os.Getenv("LLM_MODEL")
	if model == "" {
		model = "gemini-2.5-pro-exp-03-25"
	}
	log.Info("using model", "model", model)

	out := &bytes.Buffer{}
	if err := x.InferOutput_WithCompletion(ctx, model, dataPoint, out); err != nil {
		return fmt.Errorf("running LLM inference: %w", err)
	}

	if o.Output == "" {
		fmt.Println(out)
		return nil
	}

	if tmpF, err := kccio.WriteToCache(ctx, o.Output, out.String(), fileNamePattern(dataPoint)); err != nil {
		return err
	} else {
		fmt.Println(tmpF)
	}
	return nil
}

func fileNamePattern(dataPoint *toolbot.DataPoint) string {
	for k, _ := range dataPoint.Input {
		return strings.Replace(k, " ", "-", -1)
	}
	return ""
}
