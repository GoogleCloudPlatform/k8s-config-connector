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
	"io"
	"strings"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/validation"
)

// setupOptions holds flags and configuration for `config-connector ha setup`.
type setupOptions struct {
	baseOptions

	bucket         string
	bucketLocation string
	leaseName      string
	leaseNamespace string
	// mclImage and syncerImage default to the images of this version of Config Connector; see
	// defaultImage.
	mclImage           string
	syncerImage        string
	leaderEndpointType string
	yamlOnly           bool
	dryRun             bool
}

const (
	defaultLeaseName          = "kcc-multicluster-lease"
	defaultLeaseNamespace     = "cnrm-system"
	defaultLeaderEndpointType = operatorv1beta1.LeaderEndpointTypeDefault

	// The repositories that a Config Connector release publishes the multiclusterlease and
	// krmsyncer images to.
	defaultMCLImageRepo    = "gcr.io/gke-release/cnrm/multiclusterlease"
	defaultSyncerImageRepo = "gcr.io/gke-release/cnrm/krmsyncer"
)

func newSetupOptions() setupOptions {
	return setupOptions{
		leaseName:          defaultLeaseName,
		leaseNamespace:     defaultLeaseNamespace,
		mclImage:           defaultImage(defaultMCLImageRepo),
		syncerImage:        defaultImage(defaultSyncerImageRepo),
		leaderEndpointType: defaultLeaderEndpointType,
	}
}

// defaultImage returns the image in repo that belongs to this version of the CLI.
//
// A release tags every image that it builds with the Config Connector version, and it builds
// the multiclusterlease and krmsyncer images from the forks in third_party/github.com/gke-labs.
// So this image is built from the same commit as the manifests that setup applies.
//
// version/VERSION only changes when a release is cut. A CLI built from a later commit still
// defaults to the images of that release.
func defaultImage(repo string) string {
	return repo + ":" + version.GetVersion()
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

const (
	bucketFlag             = "bucket"
	bucketLocationFlag     = "bucket-location"
	leaseNameFlag          = "lease-name"
	leaseNamespaceFlag     = "lease-namespace"
	mclImageFlag           = "mcl-image"
	syncerImageFlag        = "syncer-image"
	leaderEndpointTypeFlag = "leader-endpoint-type"
	yamlOnlyFlag           = "yaml-only"
	dryRunFlag             = "dry-run"
)

func (o *setupOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, bucketFlag, o.bucket, "GCS bucket name used to store the multi-cluster leader election lock.")
	flags.StringVar(&o.bucketLocation, bucketLocationFlag, o.bucketLocation, "GCS bucket location for the multi-cluster lease bucket (e.g. 'us', 'eu', 'asia', a dual-region, or a single region). If not specified, automatically derived from the cluster locations.")
	flags.StringVar(&o.leaseName, leaseNameFlag, o.leaseName, "Name of the MultiClusterLease object across all clusters.")
	flags.StringVar(&o.leaseNamespace, leaseNamespaceFlag, o.leaseNamespace, "Namespace of the MultiClusterLease object across all clusters.")
	flags.StringVar(&o.mclImage, mclImageFlag, o.mclImage, "Container image of the MultiClusterLease controller. The default is the multiclusterlease image of this Config Connector version.")
	flags.StringVar(&o.syncerImage, syncerImageFlag, o.syncerImage, "Container image of the KRM syncer. The default is the krmsyncer image of this Config Connector version.")
	flags.StringVar(&o.leaderEndpointType, leaderEndpointTypeFlag, o.leaderEndpointType, "The control plane endpoint that standby clusters use to reach the leader's cluster: Default, DNS or PrivateIP.")
	flags.BoolVar(&o.yamlOnly, yamlOnlyFlag, o.yamlOnly, "If true, print rendered manifests to stdout without applying them.")
	flags.BoolVar(&o.dryRun, dryRunFlag, o.dryRun, "If true, print planned actions without mutating clusters or GCP resources.")
}

func (o *setupOptions) validate() error {
	var errs []error
	if err := o.baseOptions.validate(); err != nil {
		errs = append(errs, err)
	}

	o.bucket = normalizeBucketName(o.bucket)
	o.bucketLocation = strings.ToLower(strings.TrimSpace(o.bucketLocation))
	o.leaseName = strings.TrimSpace(o.leaseName)
	o.leaseNamespace = strings.TrimSpace(o.leaseNamespace)
	o.mclImage = strings.TrimSpace(o.mclImage)
	o.syncerImage = strings.TrimSpace(o.syncerImage)
	o.leaderEndpointType = strings.TrimSpace(o.leaderEndpointType)

	if o.bucket == "" {
		errs = append(errs, fmt.Errorf("--%s is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --%s=my-kcc-ha-bucket", bucketFlag, bucketFlag))
	} else if err := validateBucketName(o.bucket); err != nil {
		errs = append(errs, err)
	}

	if o.bucketLocation != "" {
		if !isValidBucketLocation(o.bucketLocation) {
			errs = append(errs, fmt.Errorf("--%s %q is not a valid GCS bucket location; use lowercase alphanumeric characters and hyphens (e.g. 'us', 'eu', 'asia', 'us-central1')", bucketLocationFlag, o.bucketLocation))
		}
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

	// An empty image is more likely an unset shell variable than a request for the default.
	if o.mclImage == "" {
		errs = append(errs, fmt.Errorf("--%s cannot be empty; omit the flag to use the default image (%s), or specify the image that runs the MultiClusterLease controller", mclImageFlag, defaultImage(defaultMCLImageRepo)))
	}
	if o.syncerImage == "" {
		errs = append(errs, fmt.Errorf("--%s cannot be empty; omit the flag to use the default image (%s), or specify the image that runs the KRM syncer", syncerImageFlag, defaultImage(defaultSyncerImageRepo)))
	}

	if o.leaderEndpointType != "" {
		switch o.leaderEndpointType {
		case operatorv1beta1.LeaderEndpointTypeDefault, operatorv1beta1.LeaderEndpointTypeDNS, operatorv1beta1.LeaderEndpointTypePrivateIP:
		default:
			errs = append(errs, fmt.Errorf("--%s %q is not a valid endpoint type; use Default, DNS or PrivateIP", leaderEndpointTypeFlag, o.leaderEndpointType))
		}
	}

	return errors.Join(errs...)
}

func isValidBucketLocation(loc string) bool {
	if len(loc) < 2 {
		return false
	}
	for i := 0; i < len(loc); i++ {
		c := loc[i]
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-') {
			return false
		}
	}
	if loc[0] == '-' || loc[len(loc)-1] == '-' {
		return false
	}
	return true
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
	configurers, gkeClusters, err := initClusterConfigurers(ctx, handles, opts)
	if err != nil {
		return err
	}

	if opts.yamlOnly {
		return renderClusterYAMLs(configurers, opts.stdout)
	}

	provisioner := newGCPProvisioner(runner, opts.bucket, opts.bucketLocation, opts.stderr)
	if err := provisioner.provision(ctx, gkeClusters); err != nil {
		return err
	}

	return applyClusterConfigs(ctx, configurers, opts.stdout, opts.dryRun)
}

func initClusterConfigurers(ctx context.Context, handles []*clusterHandle, opts setupOptions) ([]*clusterConfigurer, []gkeClusterInfo, error) {
	if len(handles) == 0 {
		return nil, nil, fmt.Errorf("no cluster handles provided for `config-connector ha setup`")
	}

	var errs []error
	configurers := make([]*clusterConfigurer, 0, len(handles))
	gkeClusters := make([]gkeClusterInfo, 0, len(handles))
	for _, h := range handles {
		cfg, err := newClusterConfigurer(ctx, h, opts.mclImage, opts.syncerImage, opts.leaderEndpointType, opts.leaseName, opts.leaseNamespace, opts.bucket)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		configurers = append(configurers, cfg)
		gkeClusters = append(gkeClusters, cfg.clusterInfo())
	}

	if len(errs) > 0 {
		return nil, nil, errors.Join(errs...)
	}

	for i, cfg := range configurers {
		cfg.setPeerProjects(peerProjectsForCluster(gkeClusters, i))
	}

	return configurers, gkeClusters, nil
}

// peerProjectsForCluster returns the projects of every cluster other than clusters[selfIndex].
// The cluster's own project is among them when another cluster shares it.
func peerProjectsForCluster(clusters []gkeClusterInfo, selfIndex int) []string {
	seen := make(map[string]bool, len(clusters))
	peers := make([]string, 0, len(clusters))
	for i, c := range clusters {
		if i == selfIndex {
			continue
		}
		if !seen[c.project] {
			seen[c.project] = true
			peers = append(peers, c.project)
		}
	}
	return peers
}

func renderClusterYAMLs(configurers []*clusterConfigurer, stdout io.Writer) error {
	for _, cfg := range configurers {
		rendered, err := cfg.renderYAML()
		if err != nil {
			return err
		}

		if stdout != nil {
			fmt.Fprint(stdout, rendered)
		}
	}

	return nil
}

func applyClusterConfigs(ctx context.Context, configurers []*clusterConfigurer, stdout io.Writer, dryRun bool) error {
	if dryRun {
		for _, cfg := range configurers {
			if err := cfg.applyObjects(ctx, stdout, true); err != nil {
				return err
			}
		}
		return nil
	}

	// 1. Apply CRDs on every cluster.
	for _, cfg := range configurers {
		if err := cfg.applyCRDs(ctx, stdout, false); err != nil {
			return err
		}
	}

	// 2. Apply MCL and syncer objects on every cluster.
	for _, cfg := range configurers {
		if err := cfg.applyNonCRDObjects(ctx, stdout, false); err != nil {
			return err
		}
	}

	// 3. Wait until both Deployments are Available on every cluster.
	for _, cfg := range configurers {
		if err := cfg.waitForDeploymentsAvailable(ctx, stdout); err != nil {
			return err
		}
	}

	// 4. Patch ConfigConnector on every cluster. If anything failed before this, HA stays off on every cluster.
	var errs []error
	for _, cfg := range configurers {
		if err := cfg.updateConfigConnector(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
