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

package updatetypes

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/typeupdater"
	"github.com/spf13/cobra"
)

type insertFieldOptions struct {
	*baseUpdateTypeOptions

	parent string // Fully qualified name of the proto message holding the new field
	field  string // Name of the field to be inserted
}

func buildInsertCommand(baseOptions *baseUpdateTypeOptions) *cobra.Command {
	opt := &insertFieldOptions{
		baseUpdateTypeOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:     "insert",
		Short:   "insert a new field and all of its dependent messages into KRM types",
		PreRunE: validateInsertOptions(opt),
		RunE:    runInsert(opt),
	}

	bindInsertFlags(cmd, opt)

	return cmd
}

func bindInsertFlags(cmd *cobra.Command, opt *insertFieldOptions) {
	opt.BindFlags(cmd)
	cmd.Flags().StringVar(&opt.parent, "parent", opt.parent, "Fully qualified name of the proto message holding the new field. e.g. `google.cloud.bigquery.datatransfer.v1.TransferConfig`")
	cmd.Flags().StringVar(&opt.field, "field", opt.field, "Name of the field to be inserted, e.g. `schedule_options_v2`")
}

func validateInsertOptions(opt *insertFieldOptions) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if opt.apiDirectory == "" {
			return fmt.Errorf("--api-dir is required")
		}
		if opt.parent == "" {
			return fmt.Errorf("--parent is required")
		}
		if opt.field == "" {
			return fmt.Errorf("--field is required")
		}
		return nil
	}
}

func runInsert(opt *insertFieldOptions) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if err := runFieldInserter(ctx, opt); err != nil {
			return err
		}
		return nil
	}
}

func runFieldInserter(_ context.Context, opt *insertFieldOptions) error {
	fieldInserter := typeupdater.NewFieldInserter(&typeupdater.InsertFieldOptions{
		ProtoSourcePath:       opt.GenerateOptions.ProtoSourcePath,
		ParentMessageFullName: opt.parent,
		FieldToInsert:         opt.field,
		APIDirectory:          opt.apiDirectory,
		GoPackagePath:         opt.apiGoPackagePath,
		MetadataDir:           opt.metadataDir,
	})
	return fieldInserter.Run()
}
