// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/preview"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/preview/options"
	"github.com/spf13/cobra"
)

const (
	previewExamples = `
	# preview Config Connector resources
	config-connector preview
	`
)

func NewPreviewCmd() *cobra.Command {
	opts := options.NewOptions()
	cmd := &cobra.Command{
		Use:     "preview",
		Short:   "preview Config Connector resources",
		Example: previewExamples,
		Hidden:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return preview.Execute(cmd.Context(), opts)
		},
		Args: cobra.ExactArgs(0),
	}

	opts.AddFlags(cmd.Flags())

	return cmd
}
