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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/ha"
	"github.com/spf13/cobra"
)

const (
	haCommandName = "ha"
	haExamples    = `
	# Set up multi-cluster HA across two clusters
	config-connector ha setup --contexts=cluster-a,cluster-b --bucket=my-lease-bucket

	# Check HA leadership and syncer health across clusters
	config-connector ha status --contexts=cluster-a,cluster-b

	# Trigger a manual failover to cluster-b
	config-connector ha failover --contexts=cluster-a,cluster-b --target-cluster=cluster-b
	`
)

func NewHACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     haCommandName,
		Short:   "Manage multi-cluster High Availability (HA) for Config Connector",
		Example: haExamples,
		Hidden:  true,
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(ha.NewSetupCmd())
	cmd.AddCommand(ha.NewStatusCmd())
	cmd.AddCommand(ha.NewFailoverCmd())

	return cmd
}
