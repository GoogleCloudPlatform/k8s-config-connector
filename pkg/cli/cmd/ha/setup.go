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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	BucketFlag         = "bucket"
	LeaseNameFlag      = "lease-name"
	LeaseNamespaceFlag = "lease-namespace"
	YAMLOnlyFlag       = "yaml-only"

	defaultLeaseName      = "kcc-multicluster-lease"
	defaultLeaseNamespace = "cnrm-system"
)

// SetupOptions holds flags and configuration for `config-connector ha setup`.
type SetupOptions struct {
	*BaseOptions

	Bucket         string
	LeaseName      string
	LeaseNamespace string
	YAMLOnly       bool
}

func NewSetupOptions(baseOpts *BaseOptions) *SetupOptions {
	if baseOpts == nil {
		baseOpts = NewBaseOptions()
	}
	return &SetupOptions{
		BaseOptions:    baseOpts,
		LeaseName:      defaultLeaseName,
		LeaseNamespace: defaultLeaseNamespace,
	}
}

func NewSetupCmd(baseOpts *BaseOptions) *cobra.Command {
	opts := NewSetupOptions(baseOpts)
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Configure multi-cluster high availability across Config Connector clusters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Stdout = cmd.OutOrStdout()
			opts.Stderr = cmd.ErrOrStderr()
			if err := opts.Validate(); err != nil {
				return err
			}
			return RunSetup(cmd.Context(), opts)
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

func RunSetup(ctx context.Context, opts *SetupOptions) error {
	if err := opts.InitClusterHandles(ctx); err != nil {
		return err
	}
	return fmt.Errorf("ha setup is not yet implemented")
}

func (o *SetupOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.Bucket, BucketFlag, o.Bucket, "GCS bucket name used to store the multi-cluster leader election lock.")
	flags.StringVar(&o.LeaseName, LeaseNameFlag, o.LeaseName, "Name of the MultiClusterLease object across all clusters.")
	flags.StringVar(&o.LeaseNamespace, LeaseNamespaceFlag, o.LeaseNamespace, "Namespace of the MultiClusterLease object across all clusters.")
	flags.BoolVar(&o.YAMLOnly, YAMLOnlyFlag, o.YAMLOnly, "If true, print rendered manifests to stdout without applying them.")
}

func (o *SetupOptions) Validate() error {
	if err := o.BaseOptions.Validate(); err != nil {
		return err
	}
	o.Bucket = strings.TrimSpace(o.Bucket)
	o.LeaseName = strings.TrimSpace(o.LeaseName)
	o.LeaseNamespace = strings.TrimSpace(o.LeaseNamespace)

	if o.Bucket == "" {
		return fmt.Errorf("--%s is required", BucketFlag)
	}
	if o.LeaseName == "" {
		return fmt.Errorf("--%s cannot be empty", LeaseNameFlag)
	}
	if o.LeaseNamespace == "" {
		return fmt.Errorf("--%s cannot be empty", LeaseNamespaceFlag)
	}
	return nil
}
