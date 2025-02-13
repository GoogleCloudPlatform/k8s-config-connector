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

package detectnewfields

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/newfieldsdetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type DetectNewFieldsOptions struct {
	*options.GenerateOptions

	targetMessages    string // comma-separated list of proto message names
	ignoredFieldsFile string // path to ignored fields YAML file
	outputFormat      string // optional: json, yaml, or text
}

func (o *DetectNewFieldsOptions) InitDefaults() error {
	o.outputFormat = "text"

	// Set default ignored fields file path
	_, err := options.RepoRoot()
	if err != nil {
		return err
	}
	// TODO: create this file
	// o.ignoredFieldsFile = filepath.Join(repoRoot, "dev", "tools", "controllerbuilder", "config", "ignored_fields.yaml")

	return nil
}

func (o *DetectNewFieldsOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.targetMessages, "target-messages", o.targetMessages, "Comma-separated list of target fully qualified proto message names to check")
	cmd.Flags().StringVar(&o.ignoredFieldsFile, "ignored-fields-file", o.ignoredFieldsFile, "Path to YAML file containing ignored fields configuration")
	cmd.Flags().StringVar(&o.outputFormat, "output-format", o.outputFormat, "Output format: text, json, or yaml")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &DetectNewFieldsOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "detect-new-fields",
		Short: "Detect new fields between pinned and HEAD versions of proto definitions",
		Long: `Detect new fields by comparing the pinned version of proto definitions with the current HEAD version.
The pinned version is determined by the version specified in mockgcp/git.versions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := runNewFieldDetector(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func runNewFieldDetector(ctx context.Context, opt *DetectNewFieldsOptions) error {
	ignoredFields, err := newfieldsdetector.LoadIgnoredFields(opt.ignoredFieldsFile)
	if err != nil {
		return fmt.Errorf("loading ignored fields: %w", err)
	}

	targetMessages := sets.NewString()
	if opt.targetMessages != "" {
		targetMessages = sets.NewString(strings.Split(opt.targetMessages, ",")...)
	}
	newFieldDetector, err := newfieldsdetector.NewFieldDetector(&newfieldsdetector.DetectorOptions{
		TargetMessages: targetMessages,
		IgnoredFields:  ignoredFields,
	})
	if err != nil {
		return fmt.Errorf("creating new field detector: %w", err)
	}

	diffs, err := newFieldDetector.DetectNewFields()
	if err != nil {
		return fmt.Errorf("detecting new fields: %w", err)
	}

	return outputResults(diffs, opt.outputFormat)
}

func outputResults(diffs []newfieldsdetector.MessageDiff, format string) error {
	if len(diffs) == 0 {
		klog.Info("No changes detected in the fields")
		return nil
	}

	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i].MessageName < diffs[j].MessageName
	})

	switch format {
	case "text":
		for _, diff := range diffs {
			fmt.Printf("Changes detected in message: %s\n", diff.MessageName)
			for _, field := range diff.NewFields {
				fmt.Printf("  New field: %v\n", field)
			}
			for _, field := range diff.RemovedFields {
				fmt.Printf("  Removed field: %v\n", field)
			}
			for field, change := range diff.ChangedFields {
				fmt.Printf("  Changed field %s: %v -> %v (repeated: %v)\n",
					field, change.OldType, change.NewType, change.IsRepeated)
			}
		}
	case "json":
		// TODO
		return fmt.Errorf("JSON output not yet implemented")
	case "yaml":
		// TODO
		return fmt.Errorf("YAML output not yet implemented")
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}

	return nil
}
