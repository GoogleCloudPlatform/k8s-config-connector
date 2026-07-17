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

package smoketest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

func TestSmoketest(t *testing.T) {
	if os.Getenv("E2E") != "1" {
		t.Skip("skipping smoketest; E2E=1 not set")
	}

	ctx := context.Background()

	repoRoot, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatalf("failed to get repo root: %v", err)
	}
	root := strings.TrimSpace(string(repoRoot))

	clusterName := "kcc-smoketest-" + strings.ToLower(time.Now().Format("20060102-150405"))

	// Cleanup cluster at the end and collect logs if failed
	t.Cleanup(func() {
		if t.Failed() {
			artifactsDir := os.Getenv("ARTIFACTS")
			if artifactsDir != "" {
				h, err := NewHarnessNoFatal(ctx, t)
				if err != nil {
					t.Logf("failed to create harness for artifact dumping: %v", err)
				} else {
					h.DumpArtifacts(artifactsDir)
				}
			}
		}

		t.Logf("Deleting kind cluster %q", clusterName)
		cmd := exec.CommandContext(ctx, "kind", "delete", "cluster", "--name", clusterName)
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Logf("failed to delete kind cluster: %v\nOutput: %s", err, string(output))
		}
	})

	t.Logf("Creating kind cluster %q", clusterName)
	if err := runCommand(ctx, t, root, "kind", "create", "cluster", "--name", clusterName); err != nil {
		t.Fatalf("failed to create kind cluster: %v", err)
	}

	imagePrefix := "registry.kind/"

	// Read current stable version to patch manifests
	stableVersionFile := filepath.Join(root, "operator/channels/stable")
	stableVersionBytes, err := os.ReadFile(stableVersionFile)
	if err != nil {
		t.Fatalf("failed to read stable version: %v", err)
	}
	re := regexp.MustCompile(`version:\s+([0-9.]+)`)
	matches := re.FindStringSubmatch(string(stableVersionBytes))
	if len(matches) < 2 {
		t.Fatalf("failed to extract stable version from %s", string(stableVersionBytes))
	}
	stableVersion := matches[1]

	// Use stableVersion as the image tag so the operator doesn't fail with ImagePullBackOff (issue #10260)
	imageTag := stableVersion

	t.Logf("Detected stable version %q, patching manifests to %q", stableVersion, imageTag)

	// Update pull policy to IfNotPresent for all components and update image tags
	patchManifests := func(dir string) {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && (strings.HasSuffix(info.Name(), ".yaml") || strings.HasSuffix(info.Name(), ".yml")) {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				newContent := string(content)
				newContent = strings.ReplaceAll(newContent, "imagePullPolicy: Always", "imagePullPolicy: IfNotPresent")
				newContent = strings.ReplaceAll(newContent, ":"+stableVersion, ":"+imageTag)
				// Inject METRICS_VERSION=v2 env var into all manager containers using our custom rewriter (issue #10260)
				newContent = InjectEnvVar(newContent, "METRICS_VERSION", "v2")
				if newContent != string(content) {
					t.Logf("Patched manifest %s, diff:\n%s", path, cmp.Diff(string(content), newContent))
					if err := os.WriteFile(path, []byte(newContent), info.Mode()); err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("failed to patch manifests in %s: %v", dir, err)
		}
	}

	patchManifests(filepath.Join(root, "operator/channels/packages/configconnector"))
	patchManifests(filepath.Join(root, "operator/autopilot-channels/packages/configconnector"))
	patchManifests(filepath.Join(root, "config/installbundle/components"))

	// Revert patches at the end of the test to keep workspace clean
	t.Cleanup(func() {
		t.Logf("Reverting manifest patches")
		revertManifests := func(dir string) {
			if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && (strings.HasSuffix(info.Name(), ".yaml") || strings.HasSuffix(info.Name(), ".yml")) {
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					newContent := string(content)
					newContent = strings.ReplaceAll(newContent, "imagePullPolicy: IfNotPresent", "imagePullPolicy: Always")
					newContent = strings.ReplaceAll(newContent, ":"+imageTag, ":"+stableVersion)
					newContent = RemoveEnvVar(newContent, "METRICS_VERSION")
					if newContent != string(content) {
						if err := os.WriteFile(path, []byte(newContent), info.Mode()); err != nil {
							return err
						}
					}
				}
				return nil
			}); err != nil {
				t.Errorf("failed to revert manifests in %s: %v", dir, err)
			}
		}
		revertManifests(filepath.Join(root, "operator/channels/packages/configconnector"))
		revertManifests(filepath.Join(root, "operator/autopilot-channels/packages/configconnector"))
		revertManifests(filepath.Join(root, "config/installbundle/components"))
	})

	t.Logf("Building images with tag %q", imageTag)
	buildCmd := exec.CommandContext(ctx, filepath.Join(root, "dev/tasks/build-images"))
	buildCmd.Dir = root
	buildCmd.Env = append(os.Environ(),
		"IMAGE_TAG="+imageTag,
		"IMAGE_PREFIX="+imagePrefix,
	)
	if output, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build images: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Loading images into kind")
	imagesToLoad := []string{
		"operator",
		"controller",
		"recorder",
		"webhook",
		"deletiondefender",
		"unmanageddetector",
	}
	for _, img := range imagesToLoad {
		fullImage := imagePrefix + img + ":" + imageTag
		t.Logf("Loading image %q into kind", fullImage)
		if err := runCommand(ctx, t, root, "kind", "load", "--name", clusterName, "docker-image", fullImage); err != nil {
			t.Fatalf("failed to load image %q into kind: %v", fullImage, err)
		}
	}

	t.Logf("Deploying operator to kind")
	kustomizeCmd := exec.CommandContext(ctx, "kubectl", "kustomize", filepath.Join(root, "operator/config/default"))
	kustomizeOutput, err := kustomizeCmd.Output()
	if err != nil {
		t.Fatalf("failed to run kustomize: %v", err)
	}

	manifests := string(kustomizeOutput)
	// Replace operator image and pull policy
	// The kustomize output should have the image we set during build if we ran make docker-build
	// But let's be safe and do a replacement here too if needed, or just ensure we use the built one.
	// Actually, dev/tasks/build-images calls make -C operator docker-build which updates manager_image_patch.yaml
	// So kustomize build should already have the right image.
	// However, we want to ensure imagePullPolicy is IfNotPresent so kind uses the loaded image.
	manifests = strings.ReplaceAll(manifests, "imagePullPolicy: Always", "imagePullPolicy: IfNotPresent")
	// Set --image-prefix so the operator uses local kind registry
	manifests = strings.ReplaceAll(manifests, "- --local-repo=/configconnector-operator/channels", "- --local-repo=/configconnector-operator/channels\n        - --image-prefix="+imagePrefix)

	applyCmd := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyCmd.Stdin = strings.NewReader(manifests)
	if output, err := applyCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply operator manifests: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for operator to be ready")
	if err := runCommand(ctx, t, root, "kubectl", "wait", "-n", "configconnector-operator-system", "--for=jsonpath={.status.readyReplicas}=1", "statefulset/configconnector-operator", "--timeout=5m"); err != nil {
		t.Fatalf("operator failed to become ready: %v", err)
	}

	t.Logf("Creating cnrm-system namespace")
	if err := runCommand(ctx, t, root, "kubectl", "create", "ns", "cnrm-system"); err != nil && !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("failed to create cnrm-system namespace: %v", err)
	}

	t.Logf("Creating fake google service account secret")
	secretManifest := `
apiVersion: v1
kind: Secret
metadata:
  name: kcc-google-service-account
  namespace: cnrm-system
type: Opaque
stringData:
  key.json: |
    {
      "type": "service_account",
      "project_id": "fake-project-id",
      "private_key_id": "fake-private-key-id",
      "private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAsGHDAdHZfi81LgVeeMHXYLgNDpcFYhoBykYtTDdNyA5AixID\n8JdKlCmZ6qLNnZrbs4JlBJfmzw6rjUC5bVBFg5NwYVBu3+3Msa4rgLsTGsjPH9rt\nC+QFnFhcmzg3zz8eeXBqJdhw7wmn1Xa9SsC3h6YWveBk98ecyE7yGe8J8xGphjk7\nEQ/KBmRK/EJD0ZwuYW1W4Bv5f5fca7qvi9rCprEmL8//uy0qCwoJj2jU3zc5p72M\npkSZb1XlYxxTEo/h9WCEvWS9pGhy6fJ0sA2RsBHqU4Y5O7MJEei9yu5fVSZUi05f\n/ggfUID+cFEq0Z/A98whKPEBBJ/STdEaqEEkBwIDAQABAoIBAED6EsvF0dihbXbh\ntXbI+h4AT5cTXYFRUV2B0sgkC3xqe65/2YG1Sl0gojoE9bhcxxjvLWWuy/F1Vw93\nS5gQnTsmgpzm86F8yg6euhn3UMdqOJtknDToMITzLFJmOHEZsJFOL1x3ysrUhMan\nsn4qVrIbJn+WfbumBoToSFnzbHflacOh06ZRbYa2bpSPMfGGFtwqQjRadn5+pync\nlCjaupcg209sM0qEk/BDSzHvWL1VgLMdiKBx574TSwS0o569+7vPNt92Ydi7kARo\reOzkkF4L3xNhKZnmls2eGH6A8cp1KZXoMLFuO+IwvBMA0O29LsUlKJU4PjBrf+7\nwaslnMECgYEA5bJv0L6DKZQD3RCBLue4/mDg0GHZqAhJBS6IcaXeaWeH6PgGZggV\nMGkWnULltJIYFwtaueTfjWqciAeocKx+rqoRjuDMOGgcrEf6Y+b5AqF+IjQM66Ll\nIYPUt3FCIc69z5LNEtyP4DSWsFPJ5UhAoG4QRlDTqT5q0gKHFjeLdeECgYEAxJRk\nkrsWmdmUs5NH9pyhTdEDIc59EuJ8iOqOLzU8xUw6/s2GSClopEFJeeEoIWhLuPY3\nX3bFt4ppl/ksLh05thRs4wXRxqhnokjD3IcGu3l6Gb5QZTYwb0VfN+q2tWVEE8Qc\nPQURheUsM2aP/gpJVQvNsWVmkT0Ijc3J8bR2hucCgYEAjOF4e0ueHu5NwFTTJvWx\nHTRGLwkU+l66ipcT0MCvPW7miRk2s3XZqSuLV0Ekqi/A3sF0D/g0tQPipfwsb48c\n0/wzcLKoDyCsFW7AQG315IswVcIe+peaeYfl++1XZmzrNlkPtrXY+ObIVbXOavZ5\nzOw0xyvj5jYGRnCOci33N4ECgYA91EKx2ABq0YGw3aEj0u31MMlgZ7b1KqFq2wNv\nm7oKgEiJ/hC/P673AsXefNAHeetfOKn/77aOXQ2LTEb2FiEhwNjiquDpL+ywoVxh\nT2LxsmqSEEbvHpUrWlFxn/Rpp3k7ElKjaqWxTHyTii2+BHQ+OKEwq6kQA3deSpy6\n1jz1fwKBgQDLqbdq5FA63PWqApfNVykXukg9MASIcg/0fjADFaHTPDvJjhFutxRP\nppI5Q95P12CQ/eRBZKJnRlkhkL8tfPaWPzzOpCTjID7avRhx2oLmstmYuXx0HluE\ncqXLbAV9WDpIJ3Bpa/S8tWujWhLDmixn2JeAdurWS+naH9U9e4I6Rw==\n-----END RSA PRIVATE KEY-----\n",
      "client_email": "fake-service-account@fake-project-id.iam.gserviceaccount.com",
      "client_id": "fake-client-id",
      "auth_uri": "https://accounts.google.com/o/oauth2/auth",
      "token_uri": "https://oauth2.googleapis.com/token",
      "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
      "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/fake-service-account@fake-project-id.iam.gserviceaccount.com"
    }
`
	applySecret := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applySecret.Stdin = strings.NewReader(secretManifest)
	if output, err := applySecret.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply Secret: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Configuring ConfigConnector in cluster mode with partial CRD installation (inclusive mode)")
	// Use cluster mode with inclusive resourceSettings to verify partial CRD installation (issue #9651)
	ccManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  stateIntoSpec: Absent
  credentialSecretName: kcc-google-service-account
  experiments:
    resourceSettings:
      mode: include
      resources:
      - group: storage.cnrm.cloud.google.com
        kind: StorageBucket
      - group: pubsub.cnrm.cloud.google.com
        kind: PubSubTopic
`
	applyCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyCC.Stdin = strings.NewReader(ccManifest)
	if output, err := applyCC.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply ConfigConnector: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for StorageBucket CRD")
	if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=create", "crd/storagebuckets.storage.cnrm.cloud.google.com", "--timeout=5m"); err != nil {
		t.Fatalf("StorageBucket CRD not created: %v", err)
	}
	if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=condition=Established", "crd/storagebuckets.storage.cnrm.cloud.google.com", "--timeout=5m"); err != nil {
		t.Fatalf("StorageBucket CRD not established: %v", err)
	}

	t.Logf("Waiting for PubSubTopic CRD")
	if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=create", "crd/pubsubtopics.pubsub.cnrm.cloud.google.com", "--timeout=5m"); err != nil {
		t.Fatalf("PubSubTopic CRD not created: %v", err)
	}
	if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=condition=Established", "crd/pubsubtopics.pubsub.cnrm.cloud.google.com", "--timeout=5m"); err != nil {
		t.Fatalf("PubSubTopic CRD not established: %v", err)
	}

	h := NewHarness(ctx, t)

	t.Logf("Waiting for KCC system components (webhook and controller managers) to be ready")
	if err := h.WaitForDeploymentAvailable("cnrm-system", "cnrm-webhook-manager", 5*time.Minute); err != nil {
		t.Fatalf("cnrm-webhook-manager failed to become ready: %v", err)
	}
	if err := h.WaitForStatefulSetReady("cnrm-system", "cnrm-controller-manager", 5*time.Minute); err != nil {
		t.Fatalf("cnrm-controller-manager failed to become ready: %v", err)
	}

	t.Logf("Creating namespace")
	ns := "config-control"
	if err := runCommand(ctx, t, root, "kubectl", "create", "ns", ns); err != nil && !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("failed to create namespace: %v", err)
	}

	bucketManifest := fmt.Sprintf(`
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: kcc-test-%s
  namespace: %s
  annotations:
    cnrm.cloud.google.com/project-id: "test-project-1"
spec:
  lifecycleRule:
    - action:
        type: Delete
      condition:
        age: 7
        withState: ANY
  versioning:
    enabled: true
  uniformBucketLevelAccess: true
`, ns, ns)

	// Try to apply StorageBucket with retries because of validating webhook propagation delay (issue #10260)
	t.Logf("Creating StorageBucket (with retries for validating webhook propagation delay)")
	applyTimeout := 3 * time.Minute
	applyInterval := 5 * time.Second
	applyDeadline := time.Now().Add(applyTimeout)
	for {
		applyCmd := exec.CommandContext(ctx, "kubectl", "apply", "-f", "-")
		applyCmd.Stdin = strings.NewReader(bucketManifest)
		output, err := applyCmd.CombinedOutput()
		if err == nil {
			t.Logf("StorageBucket applied successfully")
			break
		}

		if time.Now().After(applyDeadline) {
			t.Fatalf("failed to apply StorageBucket after %v: %v\nOutput: %s", applyTimeout, err, string(output))
		}

		t.Logf("Apply failed, likely validating webhook propagation delay. Retrying in %v... Error: %v\nOutput: %s", applyInterval, err, string(output))
		time.Sleep(applyInterval)
	}

	t.Logf("Waiting for StorageBucket reconciliation (expected to fail with permission error)")
	bucketName := "kcc-test-" + ns
	gvr := schema.GroupVersionResource{
		Group:    "storage.cnrm.cloud.google.com",
		Version:  "v1beta1",
		Resource: "storagebuckets",
	}

	h.MustWaitForObservedGeneration(gvr, ns, bucketName, 1)

	// Verify StorageBucket failed as expected
	status := h.MustGetReadyConditionStatus(gvr, ns, bucketName)
	if status != "False" {
		// Log more info for debugging
		describeOutput, _ := exec.CommandContext(ctx, "kubectl", "describe", "storagebucket", "-n", ns, bucketName).CombinedOutput()
		t.Logf("StorageBucket describe output:\n%s", string(describeOutput))
		t.Errorf("expected StorageBucket Ready status to be False, got %q", status)
	}

	// Verify the number of active watches / registered controllers via Prometheus metrics.
	// In inclusive mode specifying only StorageBucket and PubSubTopic, only these two controllers
	// should be registered, but since we only reconcile a StorageBucket, only its metrics will be recorded.
	t.Logf("Verifying the number of active watches / registered controllers via Prometheus metrics (issue #9651)")

	var metricsStr string
	var reconcileLines []string
	timeoutAt := time.Now().Add(1 * time.Minute)
	for {
		// 1. Get the pod name of cnrm-controller-manager
		podNameCmd := exec.CommandContext(ctx, "kubectl", "get", "pods", "-n", "cnrm-system", "-l", "cnrm.cloud.google.com/component=cnrm-controller-manager", "-o", "jsonpath={.items[0].metadata.name}")
		podNameBytes, err := podNameCmd.Output()
		if err != nil {
			t.Fatalf("failed to get manager pod name: %v\nOutput: %s", err, string(podNameBytes))
		}
		podName := strings.TrimSpace(string(podNameBytes))
		if podName == "" {
			if time.Now().After(timeoutAt) {
				t.Fatalf("timeout waiting for cnrm-controller-manager pod to start")
			}
			t.Log("manager pod name is empty, retrying...")
			time.Sleep(2 * time.Second)
			continue
		}

		// 2. Query metrics from the pod via kubectl API proxy. Since we use METRICS_VERSION=v2,
		// the standard controller-runtime metrics are served on /metrics.v2.
		metricsCmd := exec.CommandContext(ctx, "kubectl", "get", "--raw", fmt.Sprintf("/api/v1/namespaces/cnrm-system/pods/%s:8888/proxy/metrics.v2", podName))
		metricsBytes, err := metricsCmd.Output()
		if err != nil {
			t.Logf("failed to get metrics from pod proxy: %v, retrying...", err)
			time.Sleep(2 * time.Second)
			continue
		}
		metricsStr = string(metricsBytes)
		t.Logf("Raw metrics response:\n%s", metricsStr)

		// Filter for lines containing "controller_runtime_max_concurrent_reconciles"
		reconcileLines = nil
		for _, line := range strings.Split(metricsStr, "\n") {
			if strings.Contains(line, "controller_runtime_max_concurrent_reconciles") && !strings.HasPrefix(line, "#") {
				reconcileLines = append(reconcileLines, line)
			}
		}

		// Since we specified only "StorageBucket" and "PubSubTopic" to be included, there should be very few kinds of workers registered.
		// If the metrics are fully populated and we have at least 1 but not too many, we are good.
		if len(reconcileLines) >= 1 && len(reconcileLines) <= 5 {
			break
		}

		if time.Now().After(timeoutAt) {
			break
		}
		t.Logf("Metrics not fully matching or populated yet (found %d lines). Retrying...", len(reconcileLines))
		time.Sleep(2 * time.Second)
	}

	t.Logf("Found %d kind worker entries in metrics:\n%s", len(reconcileLines), strings.Join(reconcileLines, "\n"))

	if len(reconcileLines) == 0 {
		t.Fatalf("expected some kind worker entries to be registered, but found none in metrics. Raw metrics response:\n%s", metricsStr)
	}
	if len(reconcileLines) > 5 {
		t.Errorf("expected only a few kind worker entries (due to inclusive mode specifying only StorageBucket and PubSubTopic), but found %d in metrics:\n%s", len(reconcileLines), strings.Join(reconcileLines, "\n"))
	}

	t.Logf("Smoketest completed successfully (failed as expected)")
}

func runCommand(ctx context.Context, t *testing.T, dir string, name string, args ...string) error {
	t.Helper()
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("command %s %v failed: %w\nOutput: %s", name, args, err, string(output))
	}
	return nil
}

type Harness struct {
	*testing.T
	ctx           context.Context
	dynamicClient dynamic.Interface
}

func NewHarnessNoFatal(ctx context.Context, t *testing.T) (*Harness, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to build kubeconfig: %w", err)
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build dynamic client: %w", err)
	}
	return &Harness{
		T:             t,
		ctx:           ctx,
		dynamicClient: dynamicClient,
	}, nil
}

func NewHarness(ctx context.Context, t *testing.T) *Harness {
	h, err := NewHarnessNoFatal(ctx, t)
	if err != nil {
		t.Fatalf("%v", err)
	}
	return h
}

func (h *Harness) WaitForDeploymentAvailable(ns, name string, timeout time.Duration) error {
	h.Logf("Waiting for deployment %s/%s to be available", ns, name)
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	deadline := time.Now().Add(timeout)
	for {
		if h.ctx.Err() != nil {
			return h.ctx.Err()
		}
		obj, err := h.dynamicClient.Resource(gvr).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
		if err != nil {
			h.Logf("deployment %s/%s not found or error occurred: %v. Retrying...", ns, name, err)
		} else {
			// Check status.conditions for Available == True
			status, found, err := unstructured.NestedMap(obj.Object, "status")
			if err == nil && found {
				conditions, foundConditions, err := unstructured.NestedSlice(status, "conditions")
				if err == nil && foundConditions {
					available := false
					for _, condVal := range conditions {
						cond, ok := condVal.(map[string]any)
						if !ok {
							continue
						}
						cType, _ := cond["type"].(string)
						cStatus, _ := cond["status"].(string)
						if cType == "Available" && cStatus == "True" {
							available = true
							break
						}
					}
					if available {
						h.Logf("deployment %s/%s is now available", ns, name)
						return nil
					}
				}
			}
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for deployment %s/%s to be available", ns, name)
		}
		time.Sleep(2 * time.Second)
	}
}

func (h *Harness) WaitForStatefulSetReady(ns, name string, timeout time.Duration) error {
	h.Logf("Waiting for statefulset %s/%s to be ready (readyReplicas >= 1)", ns, name)
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"}
	deadline := time.Now().Add(timeout)
	for {
		if h.ctx.Err() != nil {
			return h.ctx.Err()
		}
		obj, err := h.dynamicClient.Resource(gvr).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
		if err != nil {
			h.Logf("statefulset %s/%s not found or error occurred: %v. Retrying...", ns, name, err)
		} else {
			readyReplicas, found, err := unstructured.NestedInt64(obj.Object, "status", "readyReplicas")
			if err == nil && found && readyReplicas >= 1 {
				h.Logf("statefulset %s/%s is now ready (readyReplicas: %d)", ns, name, readyReplicas)
				return nil
			}
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for statefulset %s/%s to be ready", ns, name)
		}
		time.Sleep(2 * time.Second)
	}
}

func (h *Harness) MustWaitForObservedGeneration(gvr schema.GroupVersionResource, ns, name string, expectedGeneration int64) {
	h.Logf("Waiting for %s/%s in namespace %s to have observedGeneration >= %d", gvr.Resource, name, ns, expectedGeneration)
	timeout := 5 * time.Minute
	pollInterval := 2 * time.Second
	deadline := time.Now().Add(timeout)

	for {
		if time.Now().After(deadline) {
			obj, err := h.dynamicClient.Resource(gvr).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
			if err == nil {
				if y, err := yaml.Marshal(obj); err == nil {
					h.Logf("Last state of %s/%s in namespace %s:\n%s", gvr.Resource, name, ns, string(y))
				}
			} else {
				h.Logf("failed to fetch last state of %s/%s on timeout: %v", gvr.Resource, name, err)
			}
			h.Fatalf("timeout waiting for %s/%s in namespace %s to reach observedGeneration %d", gvr.Resource, name, ns, expectedGeneration)
		}

		obj, err := h.dynamicClient.Resource(gvr).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
		if err != nil {
			h.Logf("error getting resource %s/%s: %v, retrying...", gvr.Resource, name, err)
			time.Sleep(pollInterval)
			continue
		}

		observedGen, found, err := unstructured.NestedInt64(obj.Object, "status", "observedGeneration")
		if err != nil {
			h.Fatalf("failed to get status.observedGeneration from %s/%s: %v", gvr.Resource, name, err)
		}

		if found && observedGen >= expectedGeneration {
			h.Logf("Resource %s/%s reached observedGeneration %d", gvr.Resource, name, observedGen)
			return
		}

		time.Sleep(pollInterval)
	}
}

func (h *Harness) MustGetReadyConditionStatus(gvr schema.GroupVersionResource, ns, name string) string {
	obj, err := h.dynamicClient.Resource(gvr).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
	if err != nil {
		h.Fatalf("failed to get resource %s/%s: %v", gvr.Resource, name, err)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil {
		h.Fatalf("failed to get status.conditions from %s/%s: %v", gvr.Resource, name, err)
	}
	if !found {
		return ""
	}

	for _, condObj := range conditions {
		cond, ok := condObj.(map[string]interface{})
		if !ok {
			continue
		}
		condType, _, _ := unstructured.NestedString(cond, "type")
		if condType == "Ready" {
			condStatus, _, _ := unstructured.NestedString(cond, "status")
			return condStatus
		}
	}
	return ""
}

func (h *Harness) DumpArtifacts(artifactsDir string) {
	h.Logf("Collecting test artifacts into directory %q", artifactsDir)
	if err := os.MkdirAll(artifactsDir, 0755); err != nil {
		h.Logf("failed to create artifacts directory %q: %v", artifactsDir, err)
		return
	}

	// 1. Collect objects
	h.collectObjects(artifactsDir)

	// 2. Collect pod logs
	h.collectPodLogs(artifactsDir)
}

func (h *Harness) collectObjects(artifactsDir string) {
	kinds := []string{
		"namespaces",
		"configconnectors",
		"configconnectorcontexts",
		"storagebuckets",
		"pods",
		"deployments",
		"statefulsets",
		"services",
	}

	for _, kind := range kinds {
		kindDir := filepath.Join(artifactsDir, "objects", kind)
		if err := os.MkdirAll(kindDir, 0755); err != nil {
			h.Logf("failed to create directory %q: %v", kindDir, err)
			continue
		}

		// Check if kind is cluster-scoped.
		isClusterScoped := kind == "namespaces" || kind == "configconnectors"

		if isClusterScoped {
			// Write yaml
			yamlCmd := exec.CommandContext(h.ctx, "kubectl", "get", kind, "-o", "yaml")
			if yamlOut, err := yamlCmd.CombinedOutput(); err == nil {
				_ = os.WriteFile(filepath.Join(kindDir, "_cluster_scoped.yaml"), yamlOut, 0644)
			}
			// Write plaintext (wide output)
			txtCmd := exec.CommandContext(h.ctx, "kubectl", "get", kind, "-o", "wide")
			if txtOut, err := txtCmd.CombinedOutput(); err == nil {
				_ = os.WriteFile(filepath.Join(kindDir, "_cluster_scoped.txt"), txtOut, 0644)
			}
		} else {
			// Get list of namespaces to loop over
			nsCmd := exec.CommandContext(h.ctx, "kubectl", "get", "namespaces", "-o", "jsonpath={.items[*].metadata.name}")
			nsBytes, err := nsCmd.Output()
			if err != nil {
				h.Logf("failed to list namespaces: %v", err)
				continue
			}
			namespaces := strings.Fields(string(nsBytes))

			for _, ns := range namespaces {
				// Query if there are any objects of this kind in this namespace to avoid empty files
				checkCmd := exec.CommandContext(h.ctx, "kubectl", "get", kind, "-n", ns, "-o", "jsonpath={.items}")
				checkBytes, _ := checkCmd.Output()
				if strings.TrimSpace(string(checkBytes)) == "[]" || strings.TrimSpace(string(checkBytes)) == "" {
					continue
				}

				// Write yaml
				yamlCmd := exec.CommandContext(h.ctx, "kubectl", "get", kind, "-n", ns, "-o", "yaml")
				if yamlOut, err := yamlCmd.CombinedOutput(); err == nil {
					_ = os.WriteFile(filepath.Join(kindDir, ns+".yaml"), yamlOut, 0644)
				}
				// Write plaintext (wide output)
				txtCmd := exec.CommandContext(h.ctx, "kubectl", "get", kind, "-n", ns, "-o", "wide")
				if txtOut, err := txtCmd.CombinedOutput(); err == nil {
					_ = os.WriteFile(filepath.Join(kindDir, ns+".txt"), txtOut, 0644)
				}
			}
		}
	}
}

func (h *Harness) collectPodLogs(artifactsDir string) {
	podCmd := exec.CommandContext(h.ctx, "kubectl", "get", "pods", "-A", "-o", "json")
	podBytes, err := podCmd.Output()
	if err != nil {
		h.Logf("failed to get pods for logs collection: %v", err)
		return
	}

	var pods struct {
		Items []struct {
			Metadata struct {
				Name      string `json:"name"`
				Namespace string `json:"namespace"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Name string `json:"name"`
				} `json:"containers"`
				InitContainers []struct {
					Name string `json:"name"`
				} `json:"initContainers"`
			} `json:"spec"`
		} `json:"items"`
	}

	if err := json.Unmarshal(podBytes, &pods); err != nil {
		h.Logf("failed to unmarshal pods json: %v", err)
		return
	}

	for _, pod := range pods.Items {
		ns := pod.Metadata.Namespace
		name := pod.Metadata.Name

		nsDir := filepath.Join(artifactsDir, "logs", ns)
		if err := os.MkdirAll(nsDir, 0755); err != nil {
			h.Logf("failed to create log directory %q: %v", nsDir, err)
			continue
		}

		var allContainers []string
		for _, c := range pod.Spec.Containers {
			allContainers = append(allContainers, c.Name)
		}
		for _, initC := range pod.Spec.InitContainers {
			allContainers = append(allContainers, initC.Name)
		}

		// Filter duplicates if any
		containerNames := make(map[string]bool)
		var uniqueContainers []string
		for _, c := range allContainers {
			if !containerNames[c] {
				containerNames[c] = true
				uniqueContainers = append(uniqueContainers, c)
			}
		}

		if len(uniqueContainers) == 1 {
			// Save directly as <podname>.log
			logFile := filepath.Join(nsDir, name+".log")
			logCmd := exec.CommandContext(h.ctx, "kubectl", "logs", "-n", ns, name)
			if logOut, err := logCmd.CombinedOutput(); err == nil {
				_ = os.WriteFile(logFile, logOut, 0644)
			}
		} else {
			// Save each container log as <podname>_<containername>.log
			for _, container := range uniqueContainers {
				logFile := filepath.Join(nsDir, fmt.Sprintf("%s_%s.log", name, container))
				logCmd := exec.CommandContext(h.ctx, "kubectl", "logs", "-n", ns, name, "-c", container)
				if logOut, err := logCmd.CombinedOutput(); err == nil {
					_ = os.WriteFile(logFile, logOut, 0644)
				}
			}
		}
	}
}
