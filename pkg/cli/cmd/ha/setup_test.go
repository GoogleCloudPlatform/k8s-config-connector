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
		name              string
		contexts          []string
		bucket            string
		leaseName         string
		leaseNamespace    string
		wantErrSubstrings []string
	}{
		{
			name:              "fewer than 2 contexts",
			contexts:          []string{"cluster-a"},
			bucket:            "my-bucket",
			wantErrSubstrings: []string{"at least 2 cluster contexts must be specified"},
		},
		{
			name:              "duplicate contexts padded with whitespace",
			contexts:          []string{"cluster-a", " cluster-a "},
			bucket:            "my-bucket",
			wantErrSubstrings: []string{"duplicate context \"cluster-a\""},
		},
		{
			name:              "empty bucket",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "whitespace-only bucket",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "   ",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "bare gs:// prefix without bucket name",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "gs://   ",
			wantErrSubstrings: []string{"--bucket is required"},
		},
		{
			name:              "invalid DNS-1123 lease-name and lease-namespace",
			contexts:          []string{"cluster-a", "cluster-b"},
			bucket:            "my-bucket",
			leaseName:         "INVALID_LEASE!",
			leaseNamespace:    "BAD_NAMESPACE!",
			wantErrSubstrings: []string{"--lease-name \"INVALID_LEASE!\" is not a valid DNS-1123 subdomain", "--lease-namespace \"BAD_NAMESPACE!\" is not a valid DNS-1123 namespace label"},
		},
		{
			name:              "accumulates multiple flag errors",
			contexts:          []string{"cluster-a"},
			bucket:            "",
			wantErrSubstrings: []string{"at least 2 cluster contexts", "--bucket is required"},
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

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
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
	for _, h := range handles {
		mclDeploy := &appsv1.Deployment{}
		if err := h.client.Get(context.Background(), types.NamespacedName{Namespace: cnrmSystemNamespace, Name: mclDeploymentName}, mclDeploy); err != nil {
			t.Fatalf("expected Deployment %s on %s, got error: %v", mclDeploymentName, h.contextName, err)
		}
		if gotImage := mclDeploy.Spec.Template.Spec.Containers[0].Image; gotImage != kccImage {
			t.Fatalf("got MCL image %q on %s, want %q", gotImage, h.contextName, kccImage)
		}
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetching ConfigConnector on %s: %v", h.contextName, err)
		}
		if cc.Spec.Experiments == nil || cc.Spec.Experiments.MultiClusterLease == nil {
			t.Fatalf("expected MultiClusterLease to be configured on %s, got %+v", h.contextName, cc)
		}
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
	for _, wantSnippet := range []string{
		mclDeploymentName,
		"gke-proj-a-us-central1-cluster-a",
		"mode: cluster",
		cnrmSystemSAGroup,
	} {
		if !strings.Contains(rendered, wantSnippet) {
			t.Fatalf("expected rendered YAML to contain %q, got:\n%s", wantSnippet, rendered)
		}
	}
}

func TestExecuteSetupPreflightErrors(t *testing.T) {
	tests := []struct {
		name              string
		handles           func(t *testing.T) []*clusterHandle
		wantErrSubstrings []string
	}{
		{
			name: "non-GKE context name fails preflight before running gcloud",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "cluster-a", newBaseConfigConnector(), newManagerStatefulSet("img:1")),
					newFakeClusterHandle(t, "cluster-b", newBaseConfigConnector(), newManagerStatefulSet("img:1")),
				}
			},
			wantErrSubstrings: []string{"does not match expected GKE context format"},
		},
		{
			name: "missing ConfigConnector and missing manager StatefulSet accumulate across clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a"),
					newFakeClusterHandle(t, "gke_proj-a_us-east1_cluster-b", newBaseConfigConnector()),
				}
			},
			wantErrSubstrings: []string{
				"missing the \"configconnector.core.cnrm.cloud.google.com\" ConfigConnector object",
				"could not find running Config Connector manager container image",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := tc.handles(t)
			runner := &recordingCommandRunner{}
			opts := newSetupOptions()
			opts.bucket = "my-ha-bucket"

			err := executeSetup(context.Background(), opts, handles, runner)

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
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
			Name: configConnectorResourceName,
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
			Namespace: cnrmSystemNamespace,
			Labels: map[string]string{
				componentLabelKey: managerStatefulSetComponent,
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
