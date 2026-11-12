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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// newTestSetupOptions returns setup options with every required flag set, except the contexts.
func newTestSetupOptions() setupOptions {
	opts := newSetupOptions()
	opts.bucket = "my-ha-bucket"
	opts.mclImage = testMCLImage
	opts.syncerImage = testSyncerImage
	return opts
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
			name: "missing mcl-image and whitespace-only syncer-image",
			invalidate: func(o *setupOptions) {
				o.mclImage = ""
				o.syncerImage = "   "
			},
			wantErr: "--mcl-image is required; specify the multiclusterlease image that runs the MultiClusterLease controller, e.g. --mcl-image=gcr.io/my-project/multiclusterlease:my-tag\n--syncer-image is required; specify the krmsyncer image that runs the KRM syncer, e.g. --syncer-image=gcr.io/my-project/krmsyncer:my-tag",
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

func TestSetupHasNoDefaultImages(t *testing.T) {
	opts := newSetupOptions()

	if opts.mclImage != "" || opts.syncerImage != "" {
		t.Fatalf("got default images %q and %q, want none: the images must be passed with --mcl-image and --syncer-image", opts.mclImage, opts.syncerImage)
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

func assertClusterConfigured(t *testing.T, h *clusterHandle, wantMCLImage, wantSyncerImage string) {
	t.Helper()
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
	if _, err := cc.Spec.Experiments.MultiClusterLease.CandidateIdentity(); err != nil {
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
	wantOutput := fmt.Sprintf(`[dry-run] cluster "gke_proj-a_us-central1_cluster-a": apply %[1]d HA manifests and configure spec.experiments.multiClusterLease (gkeCluster="projects/proj-a/locations/us-central1/clusters/cluster-a", lease=cnrm-system/kcc-multicluster-lease)
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
	if !strings.Contains(rendered, "multiclusterlease-controller-manager") || !strings.Contains(rendered, wantGKECluster) {
		t.Fatalf("expected rendered YAML to contain multiclusterlease-controller-manager and %q, got:\n%s", wantGKECluster, rendered)
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

type recordingCommandRunner struct {
	commands []string
}

func (r *recordingCommandRunner) run(_ context.Context, name string, args ...string) error {
	r.commands = append(r.commands, name+" "+strings.Join(args, " "))
	return nil
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
