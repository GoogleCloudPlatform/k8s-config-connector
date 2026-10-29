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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestSetupValidateNormalizesFields(t *testing.T) {
	opts := newSetupOptions()
	opts.contexts = []string{" cluster-a ", "cluster-b "}
	opts.bucket = " gs://my-ha-bucket/ "

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
}

func TestSetupValidateErrors(t *testing.T) {
	tests := []struct {
		name           string
		contexts       []string
		bucket         string
		leaseName      string
		leaseNamespace string
		wantErr        string
	}{
		{
			name:     "fewer than 2 contexts",
			contexts: []string{"cluster-a"},
			bucket:   "my-bucket",
			wantErr:  "at least 2 cluster contexts must be specified via --contexts",
		},
		{
			name:     "duplicate contexts padded with whitespace",
			contexts: []string{"cluster-a", " cluster-a "},
			bucket:   "my-bucket",
			wantErr:  "duplicate context \"cluster-a\" specified in --contexts",
		},
		{
			name:     "empty bucket",
			contexts: []string{"cluster-a", "cluster-b"},
			bucket:   "",
			wantErr:  "--bucket is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --bucket=my-kcc-ha-bucket",
		},
		{
			name:     "whitespace-only bucket",
			contexts: []string{"cluster-a", "cluster-b"},
			bucket:   "   ",
			wantErr:  "--bucket is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --bucket=my-kcc-ha-bucket",
		},
		{
			name:     "bare gs:// prefix without bucket name",
			contexts: []string{"cluster-a", "cluster-b"},
			bucket:   "gs://   ",
			wantErr:  "--bucket is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --bucket=my-kcc-ha-bucket",
		},
		{
			name:     "invalid bucket name syntax",
			contexts: []string{"cluster-a", "cluster-b"},
			bucket:   "INVALID_BUCKET!",
			wantErr:  "invalid GCS bucket name \"INVALID_BUCKET!\"; bucket names must be 3-63 characters of lowercase letters, numbers, '-', '_', or '.', starting and ending with a letter or number",
		},
		{
			name:           "invalid DNS-1123 lease-name and lease-namespace",
			contexts:       []string{"cluster-a", "cluster-b"},
			bucket:         "my-bucket",
			leaseName:      "INVALID_LEASE!",
			leaseNamespace: "BAD_NAMESPACE!",
			wantErr:        "--lease-name \"INVALID_LEASE!\" is not a valid DNS-1123 subdomain (a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')); use lowercase alphanumeric characters, '-' or '.'\n--lease-namespace \"BAD_NAMESPACE!\" is not a valid DNS-1123 namespace label (a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')); use lowercase alphanumeric characters or '-'",
		},
		{
			name:     "accumulates multiple flag errors",
			contexts: []string{"cluster-a"},
			bucket:   "",
			wantErr:  "at least 2 cluster contexts must be specified via --contexts\n--bucket is required; specify a GCS bucket name for the multi-cluster lease lock, e.g. --bucket=my-kcc-ha-bucket",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newSetupOptions()
			opts.contexts = slices.Clone(tc.contexts)
			opts.bucket = tc.bucket
			if tc.leaseName != "" {
				opts.leaseName = tc.leaseName
			}
			if tc.leaseNamespace != "" {
				opts.leaseNamespace = tc.leaseNamespace
			}

			err := opts.validate()

			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("got error %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestExecuteSetupAppliesManifestsAndConfiguresClusters(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
	}
	runner := &recordingCommandRunner{}
	opts := newSetupOptions()
	opts.bucket = "my-ha-bucket"

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	if len(runner.commands) == 0 {
		t.Fatalf("expected gcloud commands to be recorded, got none")
	}
	assertClusterConfigured(t, handles[0], kccImage)
	assertClusterConfigured(t, handles[1], kccImage)
}

func assertClusterConfigured(t *testing.T, h *clusterHandle, wantImage string) {
	t.Helper()
	mclDeploy := &appsv1.Deployment{}
	if err := h.client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeploy); err != nil {
		t.Fatalf("expected Deployment multiclusterlease-controller-manager on %s, got error: %v", h.contextName, err)
	}
	if gotImage := mclDeploy.Spec.Template.Spec.Containers[0].Image; gotImage != wantImage {
		t.Fatalf("got MCL image %q on %s, want %q", gotImage, h.contextName, wantImage)
	}
	cc, err := fetchConfigConnector(context.Background(), h)
	if err != nil {
		t.Fatalf("fetching ConfigConnector on %s: %v", h.contextName, err)
	}
	if cc.Spec.Experiments == nil || cc.Spec.Experiments.MultiClusterLease == nil {
		t.Fatalf("expected MultiClusterLease to be configured on %s, got %+v", h.contextName, cc)
	}
}

func TestExecuteSetupDryRunPrintsPlannedK8sActionsWithoutMutating(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
	}
	runner := &recordingCommandRunner{}
	opts := newSetupOptions()
	opts.bucket = "my-ha-bucket"
	opts.dryRun = true
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)

	if err != nil {
		t.Fatalf("executeSetup failed: %v", err)
	}
	const wantOutput = `[dry-run] cluster "gke_proj-a_us-central1_cluster-a": apply 8 HA manifests and configure spec.experiments.multiClusterLease (identity="gke-proj-a-us-central1-cluster-a", lease=cnrm-system/kcc-multicluster-lease)
[dry-run] cluster "gke_proj-b_us-east1_cluster-b": apply 8 HA manifests and configure spec.experiments.multiClusterLease (identity="gke-proj-b-us-east1-cluster-b", lease=cnrm-system/kcc-multicluster-lease)
`
	if stdout.String() != wantOutput {
		t.Fatalf("got dry-run output:\n%s\nwant:\n%s", stdout.String(), wantOutput)
	}
	mclDeploy := &appsv1.Deployment{}
	if err := handles[0].client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeploy); err == nil {
		t.Fatalf("expected Deployment multiclusterlease-controller-manager not to be created in dry-run mode")
	}
}

func TestExecuteSetupYAMLOnlyRendersWithoutMutating(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	var stdout bytes.Buffer
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
		newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector(), newManagerStatefulSet(kccImage)),
	}
	runner := &recordingCommandRunner{}
	opts := newSetupOptions()
	opts.bucket = "my-ha-bucket"
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
	if !strings.Contains(rendered, "multiclusterlease-controller-manager") || !strings.Contains(rendered, "gke-proj-a-us-central1-cluster-a") {
		t.Fatalf("expected rendered YAML to contain multiclusterlease-controller-manager and cluster identity, got:\n%s", rendered)
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
					newFakeClusterHandle(t, "cluster-a", newBaseConfigConnector(), newManagerStatefulSet("img:1")),
					newFakeClusterHandle(t, "cluster-b", newBaseConfigConnector(), newManagerStatefulSet("img:1")),
				}
			},
			wantErrPart: "does not match expected GKE context format",
		},
		{
			name: "missing ConfigConnector and missing manager StatefulSet accumulate across clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a"),
					newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector()),
				}
			},
			wantErrPart: "could not find running Config Connector manager container image",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := tc.handles(t)
			runner := &recordingCommandRunner{}
			opts := newSetupOptions()
			opts.bucket = "my-ha-bucket"

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
	commands     []string
	failDescribe bool
}

func (r *recordingCommandRunner) run(_ context.Context, name string, args ...string) error {
	full := name + " " + strings.Join(args, " ")
	r.commands = append(r.commands, full)
	if r.failDescribe && slices.Contains(args, "describe") {
		return fmt.Errorf("resource not found")
	}

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

func newManagerStatefulSet(image string) *appsv1.StatefulSet {
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: "cnrm-system",
			Labels: map[string]string{
				"cnrm.cloud.google.com/component": "cnrm-controller-manager",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "manager",
							Image: image,
						},
					},
				},
			},
		},
	}
}
