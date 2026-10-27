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

// StatusOptions holds flags and configuration for `config-connector ha status`.
type StatusOptions struct {
	*BaseOptions

	Bucket string
}

func NewStatusOptions(baseOpts *BaseOptions) *StatusOptions {
	if baseOpts == nil {
		baseOpts = NewBaseOptions()
	}
	return &StatusOptions{
		BaseOptions: baseOpts,
	}
}

func NewStatusCmd(baseOpts *BaseOptions) *cobra.Command {
	opts := NewStatusOptions(baseOpts)
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display multi-cluster high availability leadership and replication status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Stdout = cmd.OutOrStdout()
			opts.Stderr = cmd.ErrOrStderr()
			if err := opts.Validate(); err != nil {
				return err
			}
			return RunStatus(cmd.Context(), opts)
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

func RunStatus(ctx context.Context, opts *StatusOptions) error {
	if err := opts.InitClusterHandles(ctx); err != nil {
		return err
	}
	return fmt.Errorf("ha status is not yet implemented")
}

func (o *StatusOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.Bucket, BucketFlag, o.Bucket, "Optional GCS bucket name to inspect the global lock object directly.")
}

func (o *StatusOptions) Validate() error {
	if err := o.BaseOptions.Validate(); err != nil {
		return err
	}
	o.Bucket = strings.TrimSpace(o.Bucket)
	return nil
}
