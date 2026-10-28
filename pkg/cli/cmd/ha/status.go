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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// statusOptions holds flags and configuration for `config-connector ha status`.
type statusOptions struct {
	baseOptions

	bucket string
}

func newStatusOptions() statusOptions {
	return statusOptions{}
}

func NewStatusCmd() *cobra.Command {
	opts := newStatusOptions()
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display multi-cluster high availability leadership and replication status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}
			return runStatus(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

func runStatus(ctx context.Context, opts statusOptions) error {
	if _, err := buildClusterHandles(ctx, opts.contexts); err != nil {
		return err
	}
	return fmt.Errorf("ha status is not yet implemented")
}

func (o *statusOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, bucketFlag, o.bucket, "Optional GCS bucket name to inspect the global lock object directly.")
}

func (o *statusOptions) validate() error {
	if err := o.baseOptions.validate(); err != nil {
		return err
	}
	o.bucket = normalizeBucketName(o.bucket)
	return nil
}
