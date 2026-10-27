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

package ha

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	TargetClusterFlag = "target-cluster"
)

// FailoverOptions holds flags and configuration for `config-connector ha failover`.
type FailoverOptions struct {
	*BaseOptions

	TargetCluster string
}

func NewFailoverOptions(baseOpts *BaseOptions) *FailoverOptions {
	if baseOpts == nil {
		baseOpts = NewBaseOptions()
	}
	return &FailoverOptions{
		BaseOptions: baseOpts,
	}
}

func NewFailoverCmd(baseOpts *BaseOptions) *cobra.Command {
	opts := NewFailoverOptions(baseOpts)
	cmd := &cobra.Command{
		Use:   "failover",
		Short: "Trigger a manual failover to a target standby cluster",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Stdout = cmd.OutOrStdout()
			opts.Stderr = cmd.ErrOrStderr()
			if err := opts.Validate(); err != nil {
				return err
			}
			return RunFailover(cmd.Context(), opts)
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

func RunFailover(ctx context.Context, opts *FailoverOptions) error {
	if err := opts.InitClusterHandles(ctx); err != nil {
		return err
	}
	return fmt.Errorf("ha failover is not yet implemented")
}

func (o *FailoverOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.TargetCluster, TargetClusterFlag, o.TargetCluster, "Kubeconfig context name of the standby cluster to promote to leader.")
}

func (o *FailoverOptions) Validate() error {
	if err := o.BaseOptions.Validate(); err != nil {
		return err
	}
	o.TargetCluster = strings.TrimSpace(o.TargetCluster)
	if o.TargetCluster == "" {
		return fmt.Errorf("--%s is required", TargetClusterFlag)
	}
	clusterNames := o.ClusterNames()
	if !slices.Contains(clusterNames, o.TargetCluster) {
		return fmt.Errorf("--%s %q must be one of the contexts specified in --%s (%v)", TargetClusterFlag, o.TargetCluster, ContextsFlag, clusterNames)
	}
	return nil
}
