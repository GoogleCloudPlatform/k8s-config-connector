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
	"context"
	"fmt"
	"io"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/toolbot"

	"github.com/spf13/cobra"
)

type PromptOptions struct {
	*options.GenerateOptions

	ProtoDir string
}

func (o *PromptOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
}

func BuildPromptCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &PromptOptions{
		GenerateOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:   "prompt",
		Short: "build prompt",
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

func RunPrompt(ctx context.Context, o *PromptOptions) error {
	if err := rewriteFilePath(&o.ProtoDir); err != nil {
		return err
	}

	if o.ProtoDir == "" {
		return fmt.Errorf("--proto-dir is required")
	}

	extractor := &toolbot.ExtractToolMarkers{}
	addProtoDefinition := &toolbot.EnhanceWithProtoDefinition{
		ProtoDirectory: o.ProtoDir,
	}
	x, err := toolbot.NewCSVExporter(extractor, addProtoDefinition)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading from stdin: %w", err)
	}

	if err := x.BuildInputRow(ctx, b, os.Stdout); err != nil {
		return err
	}

	return nil
}
