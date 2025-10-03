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
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/toolbot"

	"github.com/spf13/cobra"
)

// ExportCSVOptions are the options for the export-csv command.
type ExportCSVOptions struct {
	*options.GenerateOptions

	ProtoDir  string
	SrcDir    string
	OutputDir string
}

// BindFlags binds the flags to the command.
func (o *ExportCSVOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for checkout of proto API definitions")
	cmd.Flags().StringVar(&o.SrcDir, "src-dir", o.SrcDir, "base directory for source code")
	cmd.Flags().StringVar(&o.OutputDir, "output-dir", o.OutputDir, "base directory for writing CSVs")
}

// BuildCommand builds the export-csv command.
func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &ExportCSVOptions{
		GenerateOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:   "export-csv",
		Short: "generate CSV from tool annotations",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunExportCSV(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

// rewriteFilePath rewrites the file path to the user's home directory if it starts with "~".
func rewriteFilePath(p *string) error {
	if strings.HasPrefix(*p, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("getting home directory: %w", err)
		}
		*p = strings.Replace(*p, "~", homeDir, 1)
	}
	return nil
}

// RunExportCSV runs the export-csv command.
func RunExportCSV(ctx context.Context, o *ExportCSVOptions) error {
	if err := rewriteFilePath(&o.ProtoDir); err != nil {
		return err
	}

	if o.ProtoDir == "" {
		return fmt.Errorf("--proto-dir is required")
	}
	if o.SrcDir == "" {
		return fmt.Errorf("--src-dir is required")
	}
	if o.OutputDir == "" {
		return fmt.Errorf("--output-dir is required")
	}

	extractor := &toolbot.ExtractToolMarkers{}
	addProtoDefinition, err := toolbot.NewEnhanceWithProtoDefinition(o.ProtoDir)
	if err != nil {
		return err
	}
	x, err := toolbot.NewCSVExporter(extractor, addProtoDefinition)
	if err != nil {
		return err
	}
	if err := x.VisitCodeDir(ctx, o.SrcDir); err != nil {
		return err
	}

	if err := x.WriteCSVForAllTools(ctx, o.OutputDir); err != nil {
		return err
	}

	return nil
}
