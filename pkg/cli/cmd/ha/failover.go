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
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	targetClusterFlag = "target-cluster"
)

// failoverOptions holds flags and configuration for `config-connector ha failover`.
type failoverOptions struct {
	baseOptions

	targetCluster string
}

func newFailoverOptions() failoverOptions {
	return failoverOptions{}
}

func NewFailoverCmd() *cobra.Command {
	opts := newFailoverOptions()
	cmd := &cobra.Command{
		Use:   "failover",
		Short: "Trigger a manual failover to a target standby cluster",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}
			return runFailover(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

func runFailover(ctx context.Context, opts failoverOptions) error {
	if _, err := buildClusterHandles(ctx, opts.contexts); err != nil {
		return err
	}
	return fmt.Errorf("ha failover is not yet implemented")
}

func (o *failoverOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.targetCluster, targetClusterFlag, o.targetCluster, "Kubeconfig context name of the standby cluster to promote to leader.")
}

func (o *failoverOptions) validate() error {
	var errs []error
	if err := o.baseOptions.validate(); err != nil {
		errs = append(errs, err)
	}
	o.targetCluster = strings.TrimSpace(o.targetCluster)
	if o.targetCluster == "" {
		errs = append(errs, fmt.Errorf("--%s is required", targetClusterFlag))
	} else if !slices.Contains(o.contexts, o.targetCluster) {
		errs = append(errs, fmt.Errorf("--%s %q must be one of the contexts specified in --%s (%v)", targetClusterFlag, o.targetCluster, contextsFlag, o.contexts))
	}
	return errors.Join(errs...)
}
