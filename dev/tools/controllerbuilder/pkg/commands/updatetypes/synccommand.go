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

type syncProtoPackageOptions struct {
	*baseUpdateTypeOptions

	legacyMode bool
}

func buildSyncCommand(baseOptions *baseUpdateTypeOptions) *cobra.Command {
	opt := &syncProtoPackageOptions{
		baseUpdateTypeOptions: baseOptions,
	}

	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync the KRM types with the proto package",
		Long: `Sync the KRM types with the proto package. This command will update the KRM types 
to match the proto package. If --message is specified, only the specified message and its 
dependent messages will be synced. If --message is not specified, all messages in the proto 
package indicated by --service will be synced.`,
		PreRunE: validateSyncOptions(opt),
		RunE:    runSync(opt),
	}

	bindSyncFlags(cmd, opt)

	return cmd
}

func bindSyncFlags(cmd *cobra.Command, opt *syncProtoPackageOptions) {
	opt.BindFlags(cmd)
	cmd.Flags().BoolVar(&opt.legacyMode, "legacy-mode", false, "Set to true if the resource has KRM fields that are missing proto annotations.")
}

func validateSyncOptions(opt *syncProtoPackageOptions) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := validateRequiredFlags(opt); err != nil {
			return err
		}
		return nil
	}
}

func validateRequiredFlags(opt *syncProtoPackageOptions) error {
	if opt.apiDirectory == "" {
		return fmt.Errorf("--api-dir is required")
	}
	if opt.apiGoPackagePath == "" {
		return fmt.Errorf("--api-go-package-path is required")
	}
	if opt.ServiceName == "" {
		return fmt.Errorf("--service is required")
	}
	return nil
}

func runSync(opt *syncProtoPackageOptions) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if err := runPackageSyncer(ctx, opt); err != nil {
			return err
		}
		return nil
	}
}

func runPackageSyncer(ctx context.Context, opt *syncProtoPackageOptions) error {
	syncer := typeupdater.NewProtoPackageSyncer(&typeupdater.SyncProtoPackageOptions{
		ServiceName:     opt.ServiceName,
		APIVersion:      opt.APIVersion,
		ProtoSourcePath: opt.GenerateOptions.ProtoSourcePath,
		APIDirectory:    opt.apiDirectory,
		GoPackagePath:   opt.apiGoPackagePath,
		LegacyMode:      opt.legacyMode,
	})
	return syncer.Run()
}
