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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/lint"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/lint/options"
	"github.com/spf13/cobra"
)

const (
	lintExamples = `
	# lint Config Connector resources across all namespaces, excludes "kube" namespaces by default
	config-connector lint

	# exclude certain namespace prefixes
	config-connector lint --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	config-connector lint --target-namespaces=my-team

	# include more resource kind in linting
	config-connector lint --include-resources=ComputeNetwork

	# ignore more resource kind in linting
	config-connector lint --ignore-resources=StorageBucket
	`
)

func NewLintCmd() *cobra.Command {
	opts := options.NewOptions()
	cmd := &cobra.Command{
		Use:     "lint",
		Short:   "lint Config Connector resources",
		Example: lintExamples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return lint.Execute(cmd.Context(), opts)
		},
		Args: cobra.ExactArgs(0),
	}

	opts.AddFlags(cmd.Flags())

	return cmd
}
