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

package cmd

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kubectl-plan/pkg/plan"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// buildRESTConfig gets the kube config (rest.Config) for the current kubernetes cluster.
func buildRESTConfig() (*rest.Config, error) {
	restConfig, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("getting kubernetes configuration: %w", err)
	}
	return restConfig, nil
}

// NewPlanCommand builds a cobra command for the plan operation.
func NewPlanCommand() *cobra.Command {
	opt := &plan.PlanOptions{}

	cmd := &cobra.Command{
		Use: "plan [DIR]...",
		// Short:
		// Long:
		// Example:
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			objects, err := readObjects(ctx, cmd, args)
			if err != nil {
				return err
			}
			opt.Objects = objects

			opt.Out = cmd.OutOrStdout()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			restConfig, err := buildRESTConfig()
			if err != nil {
				return fmt.Errorf("building kubernetes configuration: %w", err)
			}
			opt.RESTConfig = restConfig

			httpClient, err := rest.HTTPClientFor(restConfig)
			if err != nil {
				return fmt.Errorf("building HTTP client: %w", err)
			}
			opt.HTTPClient = httpClient

			return plan.RunPlan(cmd.Context(), opt)
		},
	}

	return cmd
}
