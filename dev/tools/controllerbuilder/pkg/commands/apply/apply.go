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

package apply

import (
	"bytes"
	"context"
	"fmt"
	"os"

	kccio "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/io"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/spf13/cobra"
)

type ApplyOptions struct {
	*options.GenerateOptions
	SrcFile  string
	DestFile string
}

func (o *ApplyOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.DestFile, "dest-file", o.DestFile, "destination file path to write the gocode to, default to ./apis/ to update the Config Connector types")
	cmd.Flags().StringVar(&o.SrcFile, "src-file", o.SrcFile, "src file path to read the new gocode from")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &ApplyOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "[ALPHA] Write go code from src dir to dest dir",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opt.SrcFile == "" {
				return fmt.Errorf("--src-file is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := Apply(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}
	opt.BindFlags(cmd)

	return cmd
}

func Apply(ctx context.Context, o *ApplyOptions) error {
	rawData, err := os.ReadFile(o.SrcFile)
	if err != nil {
		return err
	}
	return kccio.UpdateGoFile(ctx, o.DestFile, bytes.NewBuffer(rawData))

}
