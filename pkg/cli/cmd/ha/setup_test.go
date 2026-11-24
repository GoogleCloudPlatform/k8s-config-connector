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
	"bytes"
	"context"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/version"
	"github.com/spf13/pflag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// newTestSetupOptions returns setup options with the bucket set, and with images other than the
// defaults. The contexts are not set.
func newTestSetupOptions() setupOptions {
	opts := newSetupOptions()
	opts.bucket = "my-ha-bucket"
	opts.mclImage = testMCLImage
	opts.syncerImage = testSyncerImage
	return opts
}

// wantDefaultImages returns the images that setup deploys when --mcl-image and --syncer-image
// are not set: the ones that a release publishes, tagged with the Config Connector version.
func wantDefaultImages() (mclImage, syncerImage string) {
	v := version.GetVersion()
	return "gcr.io/gke-release/cnrm/multiclusterlease:" + v, "gcr.io/gke-release/cnrm/krmsyncer:" + v
}

func TestSetupValidateNormalizesFields(t *testing.T) {
	opts := newSetupOptions()
	opts.contexts = []string{" cluster-a ", "cluster-b "}
	opts.bucket = " gs://my-ha-bucket/ "
	opts.mclImage = " " + testMCLImage + " "
	opts.syncerImage = " " + testSyncerImage + " "

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{"cluster-a", "cluster-b"}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.bucket != "my-ha-bucket" {
		t.Fatalf("got bucket %q, want %q", opts.bucket, "my-ha-bucket")
	}
	if opts.mclImage != testMCLImage {
		t.Fatalf("got mclImage %q, want %q", opts.mclImage, testMCLImage)
	}
	if opts.syncerImage != testSyncerImage {
		t.Fatalf("got syncerImage %q, want %q", opts.syncerImage, testSyncerImage)
	}
}

func TestSetupValidateErrors(t *testing.T) {
	const bucketRequired = "--bucket is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --bucket=my-kcc-ha-bucket"
	defaultMCLImage, defaultSyncerImage := wantDefaultImages()
	tests := []struct {
		name string
		// invalidate makes otherwise valid options invalid.
		invalidate func(o *setupOptions)
		wantErr    string
	}{
		{
			name:       "fewer than 2 contexts",
			invalidate: func(o *setupOptions) { o.contexts = []string{"cluster-a"} },
			wantErr:    "at least 2 cluster contexts must be specified via --contexts",
		},
		{
			name:       "duplicate contexts padded with whitespace",
			invalidate: func(o *setupOptions) { o.contexts = []string{"cluster-a", " cluster-a "} },
			wantErr:    "duplicate context \"cluster-a\" specified in --contexts",
		},
		{
			name:       "empty bucket",
			invalidate: func(o *setupOptions) { o.bucket = "" },
			wantErr:    bucketRequired,
		},
		{
			name:       "whitespace-only bucket",
			invalidate: func(o *setupOptions) { o.bucket = "   " },
			wantErr:    bucketRequired,
		},
		{
			name:       "bare gs:// prefix without bucket name",
			invalidate: func(o *setupOptions) { o.bucket = "gs://   " },
			wantErr:    bucketRequired,
		},
		{
			name:       "invalid bucket name syntax",
			invalidate: func(o *setupOptions) { o.bucket = "INVALID_BUCKET!" },
			wantErr:    "invalid GCS bucket name \"INVALID_BUCKET!\"; bucket names must be 3-63 characters of lowercase letters, numbers, '-', '_', or '.', starting and ending with a letter or number",
		},
		{
			name: "invalid DNS-1123 lease-name and lease-namespace",
			invalidate: func(o *setupOptions) {
				o.leaseName = "INVALID_LEASE!"
				o.leaseNamespace = "BAD_NAMESPACE!"
			},
			wantErr: "--lease-name \"INVALID_LEASE!\" is not a valid DNS-1123 subdomain (a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')); use lowercase alphanumeric characters, '-' or '.'\n--lease-namespace \"BAD_NAMESPACE!\" is not a valid DNS-1123 namespace label (a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')); use lowercase alphanumeric characters or '-'",
		},
		{
			name:       "invalid bucket-location syntax",
			invalidate: func(o *setupOptions) { o.bucketLocation = "invalid location!" },
			wantErr:    "--bucket-location \"invalid location!\" is not a valid GCS bucket location; use lowercase alphanumeric characters and hyphens (e.g. 'us', 'eu', 'asia', 'us-central1')",
		},
		{
			name: "empty mcl-image and whitespace-only syncer-image",
			invalidate: func(o *setupOptions) {
				o.mclImage = ""
				o.syncerImage = "   "
			},
			wantErr: "--mcl-image cannot be empty; omit the flag to use the default image (" + defaultMCLImage + "), or specify the image that runs the MultiClusterLease controller\n--syncer-image cannot be empty; omit the flag to use the default image (" + defaultSyncerImage + "), or specify the image that runs the KRM syncer",
		},
		{
			name: "invalid leader-endpoint-type",
			invalidate: func(o *setupOptions) {
				o.leaderEndpointType = "InvalidType"
			},
			wantErr: `--leader-endpoint-type "InvalidType" is not a valid endpoint type; use Default, DNS or PrivateIP`,
		},
		{
			name: "empty leader-endpoint-type",
			invalidate: func(o *setupOptions) {
				o.leaderEndpointType = ""
			},
			wantErr: `--leader-endpoint-type "" is not a valid endpoint type; use Default, DNS or PrivateIP`,
		},
		{
			name: "accumulates multiple flag errors",
			invalidate: func(o *setupOptions) {
				o.contexts = []string{"cluster-a"}
				o.bucket = ""
			},
			wantErr: "at least 2 cluster contexts must be specified via --contexts\n" + bucketRequired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newTestSetupOptions()
			opts.contexts = []string{"cluster-a", "cluster-b"}
			tc.invalidate(&opts)

			err := opts.validate()

			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("got error %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestSetupImageFlagsDefaultToTheImagesOfThisVersion(t *testing.T) {
	wantMCLImage, wantSyncerImage := wantDefaultImages()

	flags := NewSetupCmd().Flags()

	for flagName, want := range map[string]string{
		"mcl-image":            wantMCLImage,
		"syncer-image":         wantSyncerImage,
		"leader-endpoint-type": defaultLeaderEndpointType,
	} {
		flag := flags.Lookup(flagName)
		if flag == nil {
			t.Fatalf("expected `ha setup` to have a --%s flag", flagName)
		}
		if flag.DefValue != want {
			t.Fatalf("got default --%s %q, want %q", flagName, flag.DefValue, want)
		}
	}
}

func TestSetupOptionsReadFlagsTracksExplicitLeaderEndpointType(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantExplicit bool
		wantType     string
	}{
		{
			name:         "flag omitted leaves explicit false and keeps default",
			args:         []string{},
			wantExplicit: false,
			wantType:     defaultLeaderEndpointType,
		},
		{
			name:         "explicit Default sets explicit true",
			args:         []string{"--" + leaderEndpointTypeFlag + "=Default"},
			wantExplicit: true,
			wantType:     operatorv1beta1.LeaderEndpointTypeDefault,
		},
		{
			name:         "explicit PrivateIP sets explicit true",
			args:         []string{"--" + leaderEndpointTypeFlag + "=PrivateIP"},
			wantExplicit: true,
			wantType:     operatorv1beta1.LeaderEndpointTypePrivateIP,
		},
		{
			name:         "explicit DNS sets explicit true",
			args:         []string{"--" + leaderEndpointTypeFlag + "=DNS"},
			wantExplicit: true,
			wantType:     operatorv1beta1.LeaderEndpointTypeDNS,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newSetupOptions()
			flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
			opts.addFlags(flags)

			if err := flags.Parse(tc.args); err != nil {
				t.Fatalf("flags.Parse failed: %v", err)
			}
			opts.readFlags(flags)

			if opts.leaderEndpointTypeExplicit != tc.wantExplicit {
				t.Fatalf("got leaderEndpointTypeExplicit = %v, want %v", opts.leaderEndpointTypeExplicit, tc.wantExplicit)
			}
			if opts.leaderEndpointType != tc.wantType {
				t.Fatalf("got leaderEndpointType = %q, want %q", opts.leaderEndpointType, tc.wantType)
			}
		})
	}
}

func TestExecuteSetupAppliesManifestsAndConfiguresClusters(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	if len(runner.commands) == 0 {
		t.Fatalf("expected gcloud commands to be recorded, got none")
	}
	assertClusterConfigured(t, handles[0], testMCLImage, testSyncerImage)
	assertClusterConfigured(t, handles[1], testMCLImage, testSyncerImage)
}

func TestExecuteSetupUsesTheDefaultImagesWhenNoImageIsSet(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	// Only the flags that have no default are set.
	opts := newSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.bucket = "my-ha-bucket"
	if err := opts.validate(); err != nil {
		t.Fatalf("validate failed without --mcl-image and --syncer-image: %v", err)
	}

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	wantMCLImage, wantSyncerImage := wantDefaultImages()
	assertClusterConfigured(t, handles[0], wantMCLImage, wantSyncerImage)
	assertClusterConfigured(t, handles[1], wantMCLImage, wantSyncerImage)
}

func TestExecuteSetupWithExplicitBucketLocation(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_europe-west1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return nil, fmt.Errorf("not found")
		},
	}
	opts := newTestSetupOptions()
	opts.bucketLocation = "nam4"

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	var createCmd string
	for _, cmd := range runner.commands {
		if strings.Contains(cmd, "storage buckets create") {
			createCmd = cmd
			break
		}
	}
	if !strings.Contains(createCmd, "--location=nam4") {
		t.Fatalf("expected gcloud buckets create to specify --location=nam4, got %v", runner.commands)
	}
	assertClusterConfigured(t, handles[0], testMCLImage, testSyncerImage)
	assertClusterConfigured(t, handles[1], testMCLImage, testSyncerImage)
}

func TestExecuteSetupFailsOnConflictingClusterLocationsWithoutExplicitLocation(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_europe-west1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return nil, fmt.Errorf("not found")
		},
	}
	opts := newTestSetupOptions()

	err := executeSetup(context.Background(), opts, handles, runner)

	if err == nil {
		t.Fatalf("expected executeSetup to fail on differing cluster locations, got nil")
	}
	if !strings.Contains(err.Error(), "clusters have differing bucket locations") || !strings.Contains(err.Error(), "--bucket-location") {
		t.Fatalf("expected error mentioning differing bucket locations and --bucket-location, got: %v", err)
	}
}

func assertClusterConfigured(t *testing.T, h *clusterHandle, wantMCLImage, wantSyncerImage string) {
	t.Helper()
	for _, crdName := range []string{"multiclusterleases.multicluster.core.cnrm.cloud.google.com", "krmsyncers.syncer.gkelabs.io"} {
		crd := &apiextensionsv1.CustomResourceDefinition{}
		if err := h.client.Get(context.Background(), types.NamespacedName{Name: crdName}, crd); err != nil {
			t.Fatalf("expected CRD %q on %s, got error: %v", crdName, h.contextName, err)
		}
		if crd.Labels[manifestManagedByLabelKey] != manifestManagedByLabelValue {
			t.Fatalf("expected CRD %q on %s to have label %s=%s, got %v", crdName, h.contextName, manifestManagedByLabelKey, manifestManagedByLabelValue, crd.Labels)
		}
	}
	mclDeploy := &appsv1.Deployment{}
	if err := h.client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeploy); err != nil {
		t.Fatalf("expected Deployment multiclusterlease-controller-manager on %s, got error: %v", h.contextName, err)
	}
	if gotImage := mclDeploy.Spec.Template.Spec.Containers[0].Image; gotImage != wantMCLImage {
		t.Fatalf("got MCL image %q on %s, want %q", gotImage, h.contextName, wantMCLImage)
	}
	syncerDeploy := &appsv1.Deployment{}
	if err := h.client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "syncer-controller-manager"}, syncerDeploy); err != nil {
		t.Fatalf("expected Deployment syncer-controller-manager on %s, got error: %v", h.contextName, err)
	}
	if gotImage := syncerDeploy.Spec.Template.Spec.Containers[0].Image; gotImage != wantSyncerImage {
		t.Fatalf("got syncer image %q on %s, want %q", gotImage, h.contextName, wantSyncerImage)
	}
	cc, err := fetchConfigConnector(context.Background(), h)
	if err != nil {
		t.Fatalf("fetching ConfigConnector on %s: %v", h.contextName, err)
	}
	if cc.Spec.Experiments == nil || cc.Spec.Experiments.MultiClusterLease == nil {
		t.Fatalf("expected MultiClusterLease to be configured on %s, got %+v", h.contextName, cc)
	}
	mcl := cc.Spec.Experiments.MultiClusterLease
	if mcl.GKECluster == nil {
		t.Fatalf("expected MultiClusterLease gkeCluster to be set on %s", h.contextName)
	}
	if mcl.LeaderEndpointType != defaultLeaderEndpointType {
		t.Fatalf("expected leaderEndpointType %q on %s, got %q", defaultLeaderEndpointType, h.contextName, mcl.LeaderEndpointType)
	}
	if mcl.ClusterCandidateIdentity != "" {
		t.Fatalf("expected clusterCandidateIdentity to be empty on %s, got %q", h.contextName, mcl.ClusterCandidateIdentity)
	}
	if _, err := mcl.CandidateIdentity(); err != nil {
		t.Fatalf("expected a valid MultiClusterLease gkeCluster on %s: %v", h.contextName, err)
	}
}

func TestExecuteSetupLetsTheSyncersOfTheOtherClustersReadEachCluster(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-west1_cluster-c", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	const (
		syncerOfProjA = "kcc-syncer@proj-a.iam.gserviceaccount.com"
		syncerOfProjB = "kcc-syncer@proj-b.iam.gserviceaccount.com"
	)
	wantReaders := []struct {
		project string
		readers []string
	}{
		// cluster-a and cluster-b share a project, so their syncers share a service account.
		{project: "proj-a", readers: []string{syncerOfProjA, syncerOfProjB}},
		{project: "proj-a", readers: []string{syncerOfProjA, syncerOfProjB}},
		{project: "proj-b", readers: []string{syncerOfProjA}},
	}
	for i, h := range handles {
		binding := &rbacv1.ClusterRoleBinding{}
		if err := h.client.Get(context.Background(), types.NamespacedName{Name: "cnrm-syncer-peers"}, binding); err != nil {
			t.Fatalf("getting ClusterRoleBinding cnrm-syncer-peers on %s: %v", h.contextName, err)
		}
		if binding.RoleRef.Kind != "ClusterRole" || binding.RoleRef.Name != "cnrm-viewer" {
			t.Fatalf("got roleRef %+v on %s, want ClusterRole cnrm-viewer", binding.RoleRef, h.contextName)
		}
		var gotReaders []string
		for _, subject := range binding.Subjects {
			if subject.Kind != "User" {
				t.Fatalf("got subject %+v on %s, want a User", subject, h.contextName)
			}
			gotReaders = append(gotReaders, subject.Name)
		}
		if !slices.Equal(gotReaders, wantReaders[i].readers) {
			t.Fatalf("got readers %v on %s, want %v", gotReaders, h.contextName, wantReaders[i].readers)
		}
		// A reader can only use this access if it can also find and connect to the cluster.
		for _, reader := range gotReaders {
			wantCommand := fmt.Sprintf("gcloud projects add-iam-policy-binding %s --member=serviceAccount:%s --role=roles/container.clusterViewer", wantReaders[i].project, reader)
			if !slices.Contains(runner.commands, wantCommand) {
				t.Fatalf("expected command %q for %s, got commands:\n%s", wantCommand, h.contextName, strings.Join(runner.commands, "\n"))
			}
		}
	}
}

func TestExecuteSetupDryRunCommandRunner(t *testing.T) {
	// In --dry-run, describe commands run via runner.output (non-mutating lookup) and create/bindings run via runner.run.
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			// Simulates describe running and finding nothing
			return nil, fmt.Errorf("not found")
		},
	}
	opts := newTestSetupOptions()
	opts.dryRun = true
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}

	// Verify describe commands went through runner.output (and no mutating commands went through output)
	hasDescribeBucket := false
	hasDescribeSA := false
	for _, cmd := range runner.outputCommands {
		if strings.Contains(cmd, "storage buckets describe") {
			hasDescribeBucket = true
		}
		if strings.Contains(cmd, "service-accounts describe") {
			hasDescribeSA = true
		}
		if strings.Contains(cmd, "create") || strings.Contains(cmd, "add-iam-policy-binding") {
			t.Fatalf("expected outputCommands not to contain mutating command %q", cmd)
		}
	}
	if !hasDescribeBucket {
		t.Fatalf("expected storage buckets describe in outputCommands, got %v", runner.outputCommands)
	}
	if !hasDescribeSA {
		t.Fatalf("expected service-accounts describe in outputCommands, got %v", runner.outputCommands)
	}

	// Verify create and binding commands went through runner.run (and no describe commands went through run)
	hasCreateBucket := false
	hasCreateSA := false
	hasBinding := false
	for _, cmd := range runner.runCommands {
		if strings.Contains(cmd, "storage buckets create") {
			hasCreateBucket = true
		}
		if strings.Contains(cmd, "service-accounts create") {
			hasCreateSA = true
		}
		if strings.Contains(cmd, "add-iam-policy-binding") {
			hasBinding = true
		}
		if strings.Contains(cmd, "describe") {
			t.Fatalf("expected runCommands not to contain describe command %q", cmd)
		}
	}
	if !hasCreateBucket {
		t.Fatalf("expected storage buckets create in runCommands, got %v", runner.runCommands)
	}
	if !hasCreateSA {
		t.Fatalf("expected service-accounts create in runCommands, got %v", runner.runCommands)
	}
	if !hasBinding {
		t.Fatalf("expected add-iam-policy-binding in runCommands, got %v", runner.runCommands)
	}

	// Verify stdout contains the setup dry-run output logging
	outStr := stdout.String()
	if !strings.Contains(outStr, "[dry-run] cluster \"gke_proj-a_us-central1_cluster-a\"") {
		t.Fatalf("expected stdout to contain dry-run message for cluster-a, got %q", outStr)
	}
}

func TestExecuteSetupDryRunPrintsPlannedK8sActionsWithoutMutating(t *testing.T) {
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.dryRun = true
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	manifestCount := len((&manifestBuilder{}).buildClusterObjects())
	wantOutput := fmt.Sprintf(`Lease bucket gs://my-ha-bucket already exists; reusing it.
[dry-run] cluster "gke_proj-a_us-central1_cluster-a": apply %[1]d HA manifests and configure spec.experiments.multiClusterLease (gkeCluster="projects/proj-a/locations/us-central1/clusters/cluster-a", lease=cnrm-system/kcc-multicluster-lease)
[dry-run] cluster "gke_proj-b_us-east1_cluster-b": apply %[1]d HA manifests and configure spec.experiments.multiClusterLease (gkeCluster="projects/proj-b/locations/us-east1/clusters/cluster-b", lease=cnrm-system/kcc-multicluster-lease)
`, manifestCount)
	if stdout.String() != wantOutput {
		t.Fatalf("got dry-run output:\n%s\nwant:\n%s", stdout.String(), wantOutput)
	}
	mclDeploy := &appsv1.Deployment{}
	if err := handles[0].client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeploy); err == nil {
		t.Fatalf("expected Deployment multiclusterlease-controller-manager not to be created in dry-run mode")
	}
}

func TestExecuteSetupYAMLOnlyRendersWithoutMutating(t *testing.T) {
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.yamlOnly = true
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	if len(runner.commands) != 0 {
		t.Fatalf("expected 0 gcloud commands in --yaml-only mode, got %d", len(runner.commands))
	}
	rendered := stdout.String()
	wantGKECluster := "gkeCluster:\n        location: us-central1\n        name: cluster-a\n        project: proj-a\n"
	wantLeaderEndpoint := "leaderEndpointType: Default\n"
	if !strings.Contains(rendered, "multiclusterlease-controller-manager") || !strings.Contains(rendered, wantGKECluster) || !strings.Contains(rendered, wantLeaderEndpoint) {
		t.Fatalf("expected rendered YAML to contain multiclusterlease-controller-manager, %q, and %q, got:\n%s", wantGKECluster, wantLeaderEndpoint, rendered)
	}
	if strings.Contains(rendered, "clusterCandidateIdentity") {
		t.Fatalf("expected rendered YAML not to contain deprecated clusterCandidateIdentity, got:\n%s", rendered)
	}
	// A manifest to apply has no status: only nested fields, such as a CRD's schema, may be
	// named status.
	if strings.Contains(rendered, "\nstatus:") {
		t.Fatalf("expected rendered YAML not to contain a top-level status, got:\n%s", rendered)
	}
}

func TestExecuteSetupPreflightErrors(t *testing.T) {
	tests := []struct {
		name        string
		handles     func(t *testing.T) []*clusterHandle
		wantErrPart string
	}{
		{
			name: "non-GKE context name fails preflight before running gcloud",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "cluster-a", newBaseConfigConnector()),
					newFakeClusterHandle(t, "cluster-b", newBaseConfigConnector()),
				}
			},
			wantErrPart: "does not match expected GKE context format",
		},
		{
			name: "missing ConfigConnector accumulates across clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a"),
					newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b"),
				}
			},
			wantErrPart: "missing the \"configconnector.core.cnrm.cloud.google.com\" ConfigConnector object",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := tc.handles(t)
			runner := &recordingCommandRunner{}
			opts := newTestSetupOptions()

			err := executeSetup(context.Background(), opts, handles, runner)

			if err == nil || !strings.Contains(err.Error(), tc.wantErrPart) {
				t.Fatalf("expected error containing %q, got %v", tc.wantErrPart, err)
			}
			if len(runner.commands) != 0 {
				t.Fatalf("expected 0 gcloud commands when preflight fails, got %v", runner.commands)
			}
		})
	}
}

func TestExecuteSetupReadinessTimeoutLeavesConfigConnectorsUnchanged(t *testing.T) {
	stuckPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "multiclusterlease-controller-manager-stuck",
			Namespace: "cnrm-system",
			Labels: map[string]string{
				"app": "mcl-controller",
			},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Waiting: &corev1.ContainerStateWaiting{Reason: "ImagePullBackOff", Message: "Back-off pulling image"},
					},
				},
			},
		},
	}

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandleWithPendingDeployments(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), stuckPod),
	}

	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.rolloutTimeout = 50 * time.Millisecond
	opts.rolloutPollInterval = 5 * time.Millisecond

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail due to readiness timeout, got nil")
	}
	if !strings.Contains(err.Error(), "ImagePullBackOff") || !strings.Contains(err.Error(), "spec.experiments.multiClusterLease was left unchanged on every cluster") {
		t.Fatalf("expected error with rollout details, got: %v", err)
	}

	// Verify that neither cluster's ConfigConnector was patched (HA stays off on every cluster)
	for _, h := range handles {
		cc, fetchErr := fetchConfigConnector(context.Background(), h)
		if fetchErr != nil {
			t.Fatalf("fetchConfigConnector on %s failed: %v", h.contextName, fetchErr)
		}
		if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
			t.Fatalf("expected ConfigConnector on %s to remain unchanged without MultiClusterLease, got %+v", h.contextName, cc.Spec.Experiments.MultiClusterLease)
		}
	}
}

// rolledOutMCLOnlyClient marks MCL Deployment as rolled out, but leaves syncer unready.
type rolledOutMCLOnlyClient struct {
	client.Client
}

func (c *rolledOutMCLOnlyClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if err := c.Client.Get(ctx, key, obj, opts...); err != nil {
		return err
	}
	if deployment, ok := obj.(*appsv1.Deployment); ok && key.Name == mclDeploymentName {
		replicas := int32(1)
		if deployment.Spec.Replicas != nil {
			replicas = *deployment.Spec.Replicas
		}
		deployment.Status = appsv1.DeploymentStatus{
			ObservedGeneration: deployment.Generation,
			Replicas:           replicas,
			UpdatedReplicas:    replicas,
			ReadyReplicas:      replicas,
			AvailableReplicas:  replicas,
		}
	}
	return nil
}

func TestExecuteSetupSyncerReadinessTimeoutLeavesConfigConnectorsUnchanged(t *testing.T) {
	stuckSyncerPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "syncer-controller-manager-stuck",
			Namespace: "cnrm-system",
			Labels: map[string]string{
				"app": "syncer-controller",
			},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "syncer",
					State: corev1.ContainerState{
						Waiting: &corev1.ContainerStateWaiting{Reason: "CrashLoopBackOff", Message: "Back-off restarting failed container"},
					},
				},
			},
		},
	}

	clusterBHandle := newFakeClusterHandleWithPendingDeployments(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), stuckSyncerPod)
	// Cluster B has MCL rolled out, but syncer stays pending.
	clusterBHandle.client.Client = &rolledOutMCLOnlyClient{Client: clusterBHandle.client.Client}

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		clusterBHandle,
	}

	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.rolloutTimeout = 50 * time.Millisecond
	opts.rolloutPollInterval = 5 * time.Millisecond

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail due to syncer readiness timeout, got nil")
	}
	if !strings.Contains(err.Error(), syncerDeploymentName) || !strings.Contains(err.Error(), "CrashLoopBackOff") || !strings.Contains(err.Error(), "spec.experiments.multiClusterLease was left unchanged on every cluster") {
		t.Fatalf("expected error mentioning syncer rollout failure, got: %v", err)
	}

	// Verify that neither cluster's ConfigConnector was patched (HA stays off on every cluster)
	for _, h := range handles {
		cc, fetchErr := fetchConfigConnector(context.Background(), h)
		if fetchErr != nil {
			t.Fatalf("fetchConfigConnector on %s failed: %v", h.contextName, fetchErr)
		}
		if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
			t.Fatalf("expected ConfigConnector on %s to remain unchanged without MultiClusterLease, got %+v", h.contextName, cc.Spec.Experiments.MultiClusterLease)
		}
	}
}

func TestExecuteSetupLeaderEndpointType(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypePrivateIP
	opts.leaderEndpointTypeExplicit = true

	if err := executeSetup(context.Background(), opts, handles, runner); err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}

	for _, h := range handles {
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetchConfigConnector failed: %v", err)
		}
		mcl := cc.Spec.Experiments.MultiClusterLease
		if mcl.LeaderEndpointType != operatorv1beta1.LeaderEndpointTypePrivateIP {
			t.Fatalf("expected leaderEndpointType %q on %s, got %q", operatorv1beta1.LeaderEndpointTypePrivateIP, h.contextName, mcl.LeaderEndpointType)
		}
	}

	// Re-run with an explicit new value: changes to DNS
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypeDNS
	opts.leaderEndpointTypeExplicit = true
	if err := executeSetup(context.Background(), opts, handles, runner); err != nil {
		t.Fatalf("executeSetup re-run failed: %v", err)
	}
	for _, h := range handles {
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetchConfigConnector failed: %v", err)
		}
		mcl := cc.Spec.Experiments.MultiClusterLease
		if mcl.LeaderEndpointType != operatorv1beta1.LeaderEndpointTypeDNS {
			t.Fatalf("expected leaderEndpointType updated to %q on %s, got %q", operatorv1beta1.LeaderEndpointTypeDNS, h.contextName, mcl.LeaderEndpointType)
		}
	}

	// Re-run with explicit Default: changes from DNS to Default
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypeDefault
	opts.leaderEndpointTypeExplicit = true
	if err := executeSetup(context.Background(), opts, handles, runner); err != nil {
		t.Fatalf("executeSetup re-run failed: %v", err)
	}
	for _, h := range handles {
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetchConfigConnector failed: %v", err)
		}
		mcl := cc.Spec.Experiments.MultiClusterLease
		if mcl.LeaderEndpointType != operatorv1beta1.LeaderEndpointTypeDefault {
			t.Fatalf("expected leaderEndpointType updated to %q on %s, got %q", operatorv1beta1.LeaderEndpointTypeDefault, h.contextName, mcl.LeaderEndpointType)
		}
	}

	// Re-run with flag omitted (leaderEndpointTypeExplicit = false): preserves Default
	opts.leaderEndpointType = defaultLeaderEndpointType
	opts.leaderEndpointTypeExplicit = false
	if err := executeSetup(context.Background(), opts, handles, runner); err != nil {
		t.Fatalf("executeSetup re-run with omitted flag failed: %v", err)
	}
	for _, h := range handles {
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetchConfigConnector failed: %v", err)
		}
		mcl := cc.Spec.Experiments.MultiClusterLease
		if mcl.LeaderEndpointType != operatorv1beta1.LeaderEndpointTypeDefault {
			t.Fatalf("expected leaderEndpointType %q preserved on %s, got %q", operatorv1beta1.LeaderEndpointTypeDefault, h.contextName, mcl.LeaderEndpointType)
		}
	}
}

func TestExecuteSetupRolloutFailureStopsBeforeTouchingNextCluster(t *testing.T) {
	stuckPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "multiclusterlease-controller-manager-stuck",
			Namespace: "cnrm-system",
			Labels: map[string]string{
				"app": "mcl-controller",
			},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Waiting: &corev1.ContainerStateWaiting{Reason: "ImagePullBackOff", Message: "Back-off pulling image"},
					},
				},
			},
		},
	}

	clusterAHandle := newFakeClusterHandleWithPendingDeployments(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), stuckPod)
	clusterBHandle := newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector())

	handles := []*clusterHandle{clusterAHandle, clusterBHandle}

	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.rolloutTimeout = 50 * time.Millisecond
	opts.rolloutPollInterval = 5 * time.Millisecond

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail due to readiness timeout on cluster A, got nil")
	}

	// Verify that cluster B was never touched by step 2: its MCL Deployment was never created
	mclDeploy := &appsv1.Deployment{}
	getErr := clusterBHandle.client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: mclDeploymentName}, mclDeploy)
	if getErr == nil {
		t.Fatalf("expected Deployment %q not to be created on cluster B, but it exists", mclDeploymentName)
	}
	if !apierrors.IsNotFound(getErr) {
		t.Fatalf("expected NotFound error for Deployment on cluster B, got: %v", getErr)
	}

	// Verify that ConfigConnector on both clusters was not patched
	for _, h := range handles {
		cc, fetchErr := fetchConfigConnector(context.Background(), h)
		if fetchErr != nil {
			t.Fatalf("fetchConfigConnector on %s failed: %v", h.contextName, fetchErr)
		}
		if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
			t.Fatalf("expected ConfigConnector on %s to remain unchanged without MultiClusterLease, got %+v", h.contextName, cc.Spec.Experiments.MultiClusterLease)
		}
	}
}

func TestExecuteSetupYAMLOnlyWithLeaderEndpointType(t *testing.T) {
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.yamlOnly = true
	opts.stdout = &stdout
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypePrivateIP

	if err := executeSetup(context.Background(), opts, handles, runner); err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	rendered := stdout.String()
	wantLeaderEndpoint := "leaderEndpointType: PrivateIP\n"
	if !strings.Contains(rendered, wantLeaderEndpoint) {
		t.Fatalf("expected rendered YAML to contain %q, got:\n%s", wantLeaderEndpoint, rendered)
	}
}

type recordingCommandRunner struct {
	commands       []string
	runCommands    []string
	outputCommands []string
	outputFn       func(ctx context.Context, name string, args ...string) ([]byte, error)
	runFn          func(ctx context.Context, name string, args ...string) error
}

func (r *recordingCommandRunner) run(ctx context.Context, name string, args ...string) error {
	cmd := name + " " + strings.Join(args, " ")
	r.commands = append(r.commands, cmd)
	r.runCommands = append(r.runCommands, cmd)
	if r.runFn != nil {
		return r.runFn(ctx, name, args...)
	}
	return nil
}

func (r *recordingCommandRunner) output(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := name + " " + strings.Join(args, " ")
	r.commands = append(r.commands, cmd)
	r.outputCommands = append(r.outputCommands, cmd)
	if r.outputFn != nil {
		return r.outputFn(ctx, name, args...)
	}
	return nil, nil
}

func newBaseConfigConnector() *operatorv1beta1.ConfigConnector {
	return &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configconnector.core.cnrm.cloud.google.com",
		},
		Spec: operatorv1beta1.ConfigConnectorSpec{
			Mode: "cluster",
		},
	}
}
