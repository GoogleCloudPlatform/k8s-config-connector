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

package mcp

import (
	"context"
	"fmt"

	mcpserver "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/mcp"
	"github.com/spf13/cobra"
)

func BuildMCPCmd() *cobra.Command {
	var opts MCPOptions

	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "start an MCP server to interact with KCC",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMCP(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, "kubeconfig", "", opts.kubeconfig, "path to the kubeconfig file.")
	cmd.Flags().StringVarP(&opts.context, "context", "", opts.context, "the name of the kubeconfig context to use.")

	return cmd
}

func RunMCP(ctx context.Context, opts *MCPOptions) error {
	restConfig, err := opts.getRESTConfig(ctx)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	server, err := mcpserver.NewServer(restConfig)
	if err != nil {
		return fmt.Errorf("error creating MCP server: %w", err)
	}

	return server.Serve(ctx)
}
