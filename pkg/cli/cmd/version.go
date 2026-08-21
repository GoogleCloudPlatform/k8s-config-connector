// Copyright 2022 Google LLC
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

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type VersionOptions struct {
}

func AddVersionCommand(parent *cobra.Command) {
	var opts VersionOptions
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of config-connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunVersionCommand(cmd.Context(), opts)
		},
		Args: cobra.NoArgs,
	}
	parent.AddCommand(versionCmd)
}

func RunVersionCommand(ctx context.Context, opts VersionOptions) error {
	fmt.Println(version)

	return nil
}
