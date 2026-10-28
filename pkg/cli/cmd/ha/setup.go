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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func runSetup(ctx context.Context, opts setupOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}
	runner := newExecCommandRunner(opts.stdout, opts.stderr, opts.dryRun)
	return executeSetup(ctx, opts, handles, runner)
}

type setupClusterTarget struct {
	handle          *clusterHandle
	gkeInfo         gkeClusterInfo
	configConnector *operatorv1beta1.ConfigConnector
	controllerImage string
}

func executeSetup(ctx context.Context, opts setupOptions, handles []*clusterHandle, runner commandRunner) error {
	targets, err := preflightSetupClusters(ctx, handles)
	if err != nil {
		return err
	}

	if opts.yamlOnly {
		return renderSetupManifestsToStdout(opts, targets)
	}

	if err := provisionGCPResources(ctx, opts, targets, runner); err != nil {
		return err
	}

	return applySetupToClusters(ctx, opts, targets)
}

func preflightSetupClusters(ctx context.Context, handles []*clusterHandle) ([]setupClusterTarget, error) {
	var errs []error
	targets := make([]setupClusterTarget, 0, len(handles))
	for _, h := range handles {
		if h == nil || h.client == nil {
			errs = append(errs, fmt.Errorf("cluster handle or client is nil"))
			continue
		}
		gkeInfo, err := parseGKEContext(h.contextName)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		cc, err := fetchConfigConnector(ctx, h)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if cc == nil {
			errs = append(errs, fmt.Errorf("cluster %q is missing the %q ConfigConnector object; install Config Connector and create the %q resource before running `config-connector ha setup`", h.contextName, configConnectorResourceName, configConnectorResourceName))
			continue
		}
		image, err := discoverKCCControllerImage(ctx, h)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		targets = append(targets, setupClusterTarget{
			handle:          h,
			gkeInfo:         gkeInfo,
			configConnector: cc,
			controllerImage: image,
		})
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return targets, nil
}

func discoverKCCControllerImage(ctx context.Context, handle *clusterHandle) (string, error) {
	stsList := &appsv1.StatefulSetList{}
	if err := handle.client.List(
		ctx,
		stsList,
		client.InNamespace(cnrmSystemNamespace),
		client.MatchingLabels{componentLabelKey: managerStatefulSetComponent},
	); err != nil {
		return "", fmt.Errorf("listing Config Connector controller StatefulSets in namespace %q on cluster %q: %w", cnrmSystemNamespace, handle.contextName, err)
	}
	for _, sts := range stsList.Items {
		for _, c := range sts.Spec.Template.Spec.Containers {
			if c.Name == "manager" && c.Image != "" {
				return c.Image, nil
			}
		}
	}
	return "", fmt.Errorf(
		"could not find running Config Connector manager container image on cluster %q (expected StatefulSet with label %s=%s in namespace %q); verify Config Connector is installed and healthy",
		handle.contextName, componentLabelKey, managerStatefulSetComponent, cnrmSystemNamespace,
	)
}

func renderSetupManifestsToStdout(opts setupOptions, targets []setupClusterTarget) error {
	for _, t := range targets {
		params := clusterManifestParams{
			clusterIdentity: t.gkeInfo.identity,
			project:         t.gkeInfo.project,
			bucket:          opts.bucket,
			leaseName:       opts.leaseName,
			leaseNamespace:  opts.leaseNamespace,
			controllerImage: t.controllerImage,
		}
		objs := buildClusterHAManifests(params)
		rendered, err := renderClusterYAML(t.handle.contextName, objs, params)
		if err != nil {
			return err
		}
		if opts.stdout != nil {
			fmt.Fprint(opts.stdout, rendered)
		}
	}
	return nil
}

func provisionGCPResources(ctx context.Context, opts setupOptions, targets []setupClusterTarget, runner commandRunner) error {
	primary := targets[0].gkeInfo
	bucketURI := "gs://" + opts.bucket
	if err := runner.run(ctx, "gcloud", "storage", "buckets", "create", bucketURI, "--project="+primary.project, "--location="+primary.location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q: %w", bucketURI, err)
	}

	seenProjects := make(map[string]bool, len(targets))
	projects := make([]gkeClusterInfo, 0, len(targets))
	for _, t := range targets {
		if !seenProjects[t.gkeInfo.project] {
			seenProjects[t.gkeInfo.project] = true
			projects = append(projects, t.gkeInfo)
		}
	}

	for _, p := range projects {
		if err := runner.run(ctx, "gcloud", "storage", "buckets", "add-iam-policy-binding", bucketURI,
			"--member=serviceAccount:"+p.mclGSAEmail(),
			"--role=roles/storage.objectAdmin",
		); err != nil {
			return fmt.Errorf("granting storage.objectAdmin on %q to %q: %w", bucketURI, p.mclGSAEmail(), err)
		}

		for _, binding := range []struct {
			gsa string
			ksa string
		}{
			{gsa: p.mclGSAEmail(), ksa: mclServiceAccountName},
			{gsa: p.syncerGSAEmail(), ksa: syncerServiceAccountName},
		} {
			wiMember := fmt.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]", p.project, cnrmSystemNamespace, binding.ksa)
			if err := runner.run(ctx, "gcloud", "iam", "service-accounts", "add-iam-policy-binding", binding.gsa,
				"--project="+p.project,
				"--role=roles/iam.workloadIdentityUser",
				"--member="+wiMember,
			); err != nil {
				return fmt.Errorf("binding Workload Identity on %q for %q: %w", binding.gsa, wiMember, err)
			}
		}

		for _, peer := range projects {
			if err := runner.run(ctx, "gcloud", "projects", "add-iam-policy-binding", peer.project,
				"--member=serviceAccount:"+p.syncerGSAEmail(),
				"--role=roles/container.viewer",
			); err != nil {
				return fmt.Errorf("granting container.viewer on project %q to %q: %w", peer.project, p.syncerGSAEmail(), err)
			}
		}
	}
	return nil
}

func applySetupToClusters(ctx context.Context, opts setupOptions, targets []setupClusterTarget) error {
	var errs []error
	for _, t := range targets {
		params := clusterManifestParams{
			clusterIdentity: t.gkeInfo.identity,
			project:         t.gkeInfo.project,
			bucket:          opts.bucket,
			leaseName:       opts.leaseName,
			leaseNamespace:  opts.leaseNamespace,
			controllerImage: t.controllerImage,
		}
		objs := buildClusterHAManifests(params)
		if opts.dryRun {
			if opts.stdout != nil {
				fmt.Fprintf(opts.stdout, "[dry-run] cluster %q: apply %d HA manifests and configure spec.experiments.multiClusterLease (identity=%q, lease=%s/%s)\n",
					t.handle.contextName, len(objs), params.clusterIdentity, params.leaseNamespace, params.leaseName)
			}
			continue
		}

		for _, obj := range objs {
			if err := upsertObject(ctx, t.handle.client.Client, obj); err != nil {
				errs = append(errs, fmt.Errorf("applying %s %q on cluster %q: %w", obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), t.handle.contextName, err))
			}
		}

		if err := configureClusterMultiClusterLease(ctx, t.handle, t.configConnector, params); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func upsertObject(ctx context.Context, c client.Client, obj client.Object) error {
	if err := c.Create(ctx, obj); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		existing := &unstructured.Unstructured{}
		existing.SetGroupVersionKind(obj.GetObjectKind().GroupVersionKind())
		key := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
		if getErr := c.Get(ctx, key, existing); getErr != nil {
			return getErr
		}
		obj.SetResourceVersion(existing.GetResourceVersion())
		return c.Update(ctx, obj)
	}
	return nil
}

func configureClusterMultiClusterLease(ctx context.Context, handle *clusterHandle, cc *operatorv1beta1.ConfigConnector, params clusterManifestParams) error {
	if cc.Spec.Experiments == nil {
		cc.Spec.Experiments = &operatorv1beta1.CCExperiments{}
	}
	cc.Spec.Experiments.MultiClusterLease = &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                params.leaseName,
		Namespace:                params.leaseNamespace,
		ClusterCandidateIdentity: params.clusterIdentity,
		ResourceReplicationMode:  "Status",
	}
	if err := handle.client.Update(ctx, cc); err != nil {
		return fmt.Errorf("updating ConfigConnector %q on cluster %q: %w", configConnectorResourceName, handle.contextName, err)
	}
	return nil
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
		errs = append(errs, fmt.Errorf("--%s is required", bucketFlag))
	}
	if o.leaseName == "" {
		errs = append(errs, fmt.Errorf("--%s cannot be empty", leaseNameFlag))
	}
	if o.leaseNamespace == "" {
		errs = append(errs, fmt.Errorf("--%s cannot be empty", leaseNamespaceFlag))
	}
	return errors.Join(errs...)
}
