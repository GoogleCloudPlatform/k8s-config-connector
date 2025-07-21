// Copyright 2025 Google LLC
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

package llm

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/kubectl-ai/gollm"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

// Options are the configuration options for using LLMs
type Options struct {
	Project  string
	Location string

	Model string
}

const DefaultModel = "gemini-2.5-pro"

func (o *Options) InitDefaults() {
	model := os.Getenv("LLM_MODEL")
	if model == "" {
		model = DefaultModel
	}
	o.Model = model
}

func (o *Options) InitDefaultsWithLatestModelIfUnset(ctx context.Context, llmClient gollm.Client) error {
	model := os.Getenv("LLM_MODEL")
	if model == "" {

		models, err := llmClient.ListModels(ctx)
		if err != nil {
			return fmt.Errorf("listing models: %w", err)
		}

		for _, m := range models {
			// There are many old or experimental models in the list.
			// Let's use the first `gemini-2.5-pro` model for now.
			if strings.Contains(m, "gemini-2.5-pro") {
				model = m
				break
			}
		}
	}

	klog.Infof("Using model %v", model)
	o.Model = model
	return nil
}

func (o *Options) AddFlags(flagset *flag.FlagSet) {
	flagset.StringVar(&o.Project, "project", o.Project, "the GCP project that the LLM service files billing for, Default to gcloud config")
	flagset.StringVar(&o.Location, "location", o.Location, "the GCP location. Default to gcloud config")
	flagset.StringVar(&o.Model, "model", o.Model, "The LLM model to use")
}

func (o *Options) AddCobraFlags(flags *pflag.FlagSet) {
	flagset := flag.NewFlagSet("", flag.ContinueOnError)
	o.AddFlags(flagset)
	flags.AddGoFlagSet(flagset)
}

// NewLLMClient creates a gollm.Client with the provided configuration
func (o *Options) NewLLMClient(ctx context.Context) (gollm.Client, error) {
	vertexAIOptions := gollm.VertexAIClientOptions{
		Project:  o.Project,
		Location: o.Location,
	}
	return gollm.NewVertexAIClient(ctx, vertexAIOptions)
}
