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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/validation"
)

const (
	bucketFlag         = "bucket"
	leaseNameFlag      = "lease-name"
	leaseNamespaceFlag = "lease-namespace"
	yamlOnlyFlag       = "yaml-only"

	defaultLeaseName      = "kcc-multicluster-lease"
	defaultLeaseNamespace = "cnrm-system"
)

// setupOptions holds flags and configuration for `config-connector ha setup`.
type setupOptions struct {
	baseOptions

	bucket         string
	leaseName      string
	leaseNamespace string
	yamlOnly       bool
}

func newSetupOptions() setupOptions {
	return setupOptions{
		leaseName:      defaultLeaseName,
		leaseNamespace: defaultLeaseNamespace,
	}
}

func NewSetupCmd() *cobra.Command {
	opts := newSetupOptions()
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Configure multi-cluster high availability across Config Connector clusters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}

			return runSetup(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

func (o *setupOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, bucketFlag, o.bucket, "GCS bucket name used to store the multi-cluster leader election lock.")
	flags.StringVar(&o.leaseName, leaseNameFlag, o.leaseName, "Name of the MultiClusterLease object across all clusters.")
	flags.StringVar(&o.leaseNamespace, leaseNamespaceFlag, o.leaseNamespace, "Namespace of the MultiClusterLease object across all clusters.")
	flags.BoolVar(&o.yamlOnly, yamlOnlyFlag, o.yamlOnly, "If true, print rendered manifests to stdout without applying them.")
}

func (o *setupOptions) validate() error {
	var errs []error
	if err := o.baseOptions.validate(); err != nil {
		errs = append(errs, err)
	}

	o.bucket = normalizeBucketName(o.bucket)
	o.leaseName = strings.TrimSpace(o.leaseName)
	o.leaseNamespace = strings.TrimSpace(o.leaseNamespace)

	if o.bucket == "" {
		errs = append(errs, fmt.Errorf("--%s is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --%s=my-kcc-ha-bucket", bucketFlag, bucketFlag))
	}

	if o.leaseName == "" {
		errs = append(errs, fmt.Errorf("--%s is required; specify a DNS-1123 resource name, e.g. --%s=%s", leaseNameFlag, leaseNameFlag, defaultLeaseName))
	} else if dnsErrs := validation.IsDNS1123Subdomain(o.leaseName); len(dnsErrs) > 0 {
		errs = append(errs, fmt.Errorf("--%s %q is not a valid DNS-1123 subdomain (%s); use lowercase alphanumeric characters, '-' or '.'", leaseNameFlag, o.leaseName, strings.Join(dnsErrs, "; ")))
	}

	if o.leaseNamespace == "" {
		errs = append(errs, fmt.Errorf("--%s is required; specify a DNS-1123 namespace name, e.g. --%s=%s", leaseNamespaceFlag, leaseNamespaceFlag, defaultLeaseNamespace))
	} else if dnsErrs := validation.IsDNS1123Label(o.leaseNamespace); len(dnsErrs) > 0 {
		errs = append(errs, fmt.Errorf("--%s %q is not a valid DNS-1123 namespace label (%s); use lowercase alphanumeric characters or '-'", leaseNamespaceFlag, o.leaseNamespace, strings.Join(dnsErrs, "; ")))
	}

	return errors.Join(errs...)
}

func runSetup(ctx context.Context, opts setupOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}

	runner := newExecCommandRunner(opts.stdout, opts.stderr, opts.dryRun)
	return executeSetup(ctx, opts, handles, runner)
}

func executeSetup(ctx context.Context, opts setupOptions, handles []*clusterHandle, runner commandRunner) error {
	configurers, err := preflightClusters(ctx, handles)
	if err != nil {
		return err
	}

	builder := newManifestBuilder(opts.leaseName, opts.leaseNamespace, opts.bucket)
	if opts.yamlOnly {
		for _, cfg := range configurers {
			rendered, err := cfg.renderYAML(builder)
			if err != nil {
				return err
			}

			if opts.stdout != nil {
				fmt.Fprint(opts.stdout, rendered)
			}
		}

		return nil
	}

	gkeClusters := make([]gkeClusterInfo, 0, len(configurers))
	for _, cfg := range configurers {
		gkeClusters = append(gkeClusters, cfg.gkeInfo)
	}

	provisioner := newGCPProvisioner(runner, opts.bucket)
	if err := provisioner.provision(ctx, gkeClusters); err != nil {
		return err
	}

	var errs []error
	for _, cfg := range configurers {
		if err := cfg.apply(ctx, builder, opts.stdout, opts.dryRun); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
