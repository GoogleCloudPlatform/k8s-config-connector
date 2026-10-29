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
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
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

type setupClusterTarget struct {
	handle          *clusterHandle
	gkeInfo         gkeClusterInfo
	configConnector *operatorv1beta1.ConfigConnector
	controllerImage string
}

func (t setupClusterTarget) manifestParams(opts setupOptions) clusterManifestParams {
	return clusterManifestParams{
		clusterIdentity: t.gkeInfo.identity,
		project:         t.gkeInfo.project,
		bucket:          opts.bucket,
		leaseName:       opts.leaseName,
		leaseNamespace:  opts.leaseNamespace,
		controllerImage: t.controllerImage,
	}
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
	if len(handles) == 0 {
		return nil, fmt.Errorf("no cluster handles provided for `config-connector ha setup`; specify at least 2 contexts via --%s", contextsFlag)
	}
	var errs []error
	targets := make([]setupClusterTarget, 0, len(handles))
	for i, h := range handles {
		target, err := preflightCluster(ctx, i, h)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		targets = append(targets, target)
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return targets, nil
}

func preflightCluster(ctx context.Context, index int, handle *clusterHandle) (setupClusterTarget, error) {
	if handle == nil || handle.client == nil {
		return setupClusterTarget{}, fmt.Errorf("cluster handle at index %d has a nil Kubernetes client", index)
	}
	gkeInfo, err := parseGKEContext(handle.contextName)
	if err != nil {
		return setupClusterTarget{}, err
	}
	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		return setupClusterTarget{}, err
	}
	if cc == nil {
		return setupClusterTarget{}, fmt.Errorf(
			"cluster %q is missing the %q ConfigConnector object; install Config Connector and create the %q resource before running `config-connector ha setup`",
			handle.contextName, configConnectorResourceName, configConnectorResourceName,
		)
	}
	image, err := discoverKCCControllerImage(ctx, handle)
	if err != nil {
		return setupClusterTarget{}, err
	}
	return setupClusterTarget{
		handle:          handle,
		gkeInfo:         gkeInfo,
		configConnector: cc,
		controllerImage: image,
	}, nil
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

	if image, ok := findPrimaryManagerImage(stsList.Items); ok {
		return image, nil
	}
	return "", fmt.Errorf(
		"could not find running Config Connector manager container image on cluster %q (expected StatefulSet with label %s=%s in namespace %q); verify Config Connector is installed and healthy",
		handle.contextName, componentLabelKey, managerStatefulSetComponent, cnrmSystemNamespace,
	)
}

func findPrimaryManagerImage(items []appsv1.StatefulSet) (string, bool) {
	for _, sts := range items {
		if sts.Name == managerStatefulSetComponent {
			if image, ok := managerImageFromContainers(sts.Spec.Template.Spec.Containers); ok {
				return image, true
			}
		}
	}
	for _, sts := range items {
		if image, ok := managerImageFromContainers(sts.Spec.Template.Spec.Containers); ok {
			return image, true
		}
	}
	return "", false
}

func managerImageFromContainers(containers []corev1.Container) (string, bool) {
	for _, c := range containers {
		if c.Name == "manager" && c.Image != "" {
			return c.Image, true
		}
	}
	return "", false
}

func renderSetupManifestsToStdout(opts setupOptions, targets []setupClusterTarget) error {
	for _, t := range targets {
		params := t.manifestParams(opts)
		objs := buildClusterHAManifests(params)
		rendered, err := renderClusterYAML(t.handle.contextName, objs, t.configConnector, params)
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
	bucketURI := "gs://" + opts.bucket
	if err := createLeaseBucket(ctx, runner, bucketURI, targets[0].gkeInfo); err != nil {
		return err
	}

	projects := uniqueProjectInfos(targets)
	for _, project := range projects {
		if err := configureProjectIAM(ctx, runner, bucketURI, project, projects); err != nil {
			return err
		}
	}
	return nil
}

func createLeaseBucket(ctx context.Context, runner commandRunner, bucketURI string, primary gkeClusterInfo) error {
	location := primary.gcsBucketLocation()
	if err := runner.run(ctx, "gcloud", "storage", "buckets", "create", bucketURI, "--project="+primary.project, "--location="+location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q in location %q (project %q): %w", bucketURI, location, primary.project, err)
	}
	return nil
}

func uniqueProjectInfos(targets []setupClusterTarget) []gkeClusterInfo {
	seenProjects := make(map[string]bool, len(targets))
	projects := make([]gkeClusterInfo, 0, len(targets))
	for _, t := range targets {
		if !seenProjects[t.gkeInfo.project] {
			seenProjects[t.gkeInfo.project] = true
			projects = append(projects, t.gkeInfo)
		}
	}
	return projects
}

func configureProjectIAM(ctx context.Context, runner commandRunner, bucketURI string, project gkeClusterInfo, allProjects []gkeClusterInfo) error {
	if err := grantBucketObjectAdmin(ctx, runner, bucketURI, project.mclGSAEmail()); err != nil {
		return err
	}
	if err := bindProjectWorkloadIdentities(ctx, runner, project); err != nil {
		return err
	}
	return grantPeerContainerViewerAccess(ctx, runner, project, allProjects)
}

func grantBucketObjectAdmin(ctx context.Context, runner commandRunner, bucketURI, gsaEmail string) error {
	if err := runner.run(ctx, "gcloud", "storage", "buckets", "add-iam-policy-binding", bucketURI,
		"--member=serviceAccount:"+gsaEmail,
		"--role=roles/storage.objectAdmin",
	); err != nil {
		return fmt.Errorf("granting storage.objectAdmin on %q to %q: %w", bucketURI, gsaEmail, err)
	}
	return nil
}

func bindProjectWorkloadIdentities(ctx context.Context, runner commandRunner, project gkeClusterInfo) error {
	bindings := []struct {
		gsa string
		ksa string
	}{
		{gsa: project.mclGSAEmail(), ksa: mclServiceAccountName},
		{gsa: project.syncerGSAEmail(), ksa: syncerServiceAccountName},
	}
	for _, b := range bindings {
		wiMember := fmt.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]", project.project, cnrmSystemNamespace, b.ksa)
		if err := runner.run(ctx, "gcloud", "iam", "service-accounts", "add-iam-policy-binding", b.gsa,
			"--project="+project.project,
			"--role=roles/iam.workloadIdentityUser",
			"--member="+wiMember,
		); err != nil {
			return fmt.Errorf("binding Workload Identity on %q for %q: %w", b.gsa, wiMember, err)
		}
	}
	return nil
}

func grantPeerContainerViewerAccess(ctx context.Context, runner commandRunner, project gkeClusterInfo, allProjects []gkeClusterInfo) error {
	syncerGSA := project.syncerGSAEmail()
	for _, peer := range allProjects {
		if err := runner.run(ctx, "gcloud", "projects", "add-iam-policy-binding", peer.project,
			"--member=serviceAccount:"+syncerGSA,
			"--role=roles/container.viewer",
		); err != nil {
			return fmt.Errorf("granting container.viewer on project %q to %q: %w", peer.project, syncerGSA, err)
		}
	}
	return nil
}

func applySetupToClusters(ctx context.Context, opts setupOptions, targets []setupClusterTarget) error {
	var errs []error
	for _, t := range targets {
		if err := applyClusterSetup(ctx, opts, t); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func applyClusterSetup(ctx context.Context, opts setupOptions, target setupClusterTarget) error {
	params := target.manifestParams(opts)
	objs := buildClusterHAManifests(params)
	if opts.dryRun {
		if opts.stdout != nil {
			fmt.Fprintf(opts.stdout, "[dry-run] cluster %q: apply %d HA manifests and configure spec.experiments.multiClusterLease (identity=%q, lease=%s/%s)\n",
				target.handle.contextName, len(objs), params.clusterIdentity, params.leaseNamespace, params.leaseName)
		}
		return nil
	}

	var errs []error
	for _, obj := range objs {
		if err := upsertObject(ctx, target.handle.client.Client, obj); err != nil {
			errs = append(errs, fmt.Errorf("applying %s on cluster %q: %w", formatObjectID(target.handle.client.Client, obj), target.handle.contextName, err))
		}
	}
	if err := configureClusterMultiClusterLease(ctx, target.handle, params); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func upsertObject(ctx context.Context, c client.Client, obj client.Object) error {
	err := c.Create(ctx, obj)
	if err == nil {
		return nil
	}
	if !apierrors.IsAlreadyExists(err) {
		return err
	}

	gvk, gvkErr := apiutil.GVKForObject(obj, c.Scheme())
	if gvkErr != nil {
		return gvkErr
	}
	existing := &unstructured.Unstructured{}
	existing.SetGroupVersionKind(gvk)
	key := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
	if getErr := c.Get(ctx, key, existing); getErr != nil {
		return getErr
	}
	obj.SetResourceVersion(existing.GetResourceVersion())
	return c.Update(ctx, obj)
}

func formatObjectID(c client.Client, obj client.Object) string {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	if kind == "" && c.Scheme() != nil {
		if gvk, err := apiutil.GVKForObject(obj, c.Scheme()); err == nil {
			kind = gvk.Kind
		}
	}
	if obj.GetNamespace() != "" {
		return fmt.Sprintf("%s %q (namespace %q)", kind, obj.GetName(), obj.GetNamespace())
	}
	return fmt.Sprintf("%s %q", kind, obj.GetName())
}

func configureClusterMultiClusterLease(ctx context.Context, handle *clusterHandle, params clusterManifestParams) error {
	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		return err
	}
	if cc == nil {
		return fmt.Errorf("ConfigConnector %q disappeared on cluster %q before update", configConnectorResourceName, handle.contextName)
	}
	setMultiClusterLeaseSpec(&cc.Spec, params)
	if err := handle.client.Update(ctx, cc); err != nil {
		return fmt.Errorf("updating ConfigConnector %q on cluster %q: %w", configConnectorResourceName, handle.contextName, err)
	}
	if !hasMultiClusterLeaseConfigured(cc) {
		return fmt.Errorf(
			"ConfigConnector %q on cluster %q silently dropped spec.experiments.multiClusterLease during update; the cluster's ConfigConnector CRD schema is too old (requires KCC v1.148.0+)",
			configConnectorResourceName, handle.contextName,
		)
	}
	return nil
}

func hasMultiClusterLeaseConfigured(cc *operatorv1beta1.ConfigConnector) bool {
	return cc != nil && cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil
}
