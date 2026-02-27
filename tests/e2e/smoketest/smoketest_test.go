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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
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

	// Cleanup cluster at the end
	t.Cleanup(func() {
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

	imageTag := "dev-" + time.Now().Format("20060102T150405")
	imagePrefix := "registry.kind/"

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

	t.Logf("Loading operator image into kind")
	operatorImage := imagePrefix + "operator:" + imageTag
	if err := runCommand(ctx, t, root, "kind", "load", "--name", clusterName, "docker-image", operatorImage); err != nil {
		t.Fatalf("failed to load image into kind: %v", err)
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

	t.Logf("Configuring ConfigConnector in cluster mode")
	// Use cluster mode for simplicity in smoke test as in PR 4506
	ccManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  stateIntoSpec: Absent
  credentialSecretName: kcc-google-service-account
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

	t.Logf("Creating namespace and StorageBucket")
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

	applyBucket := exec.CommandContext(ctx, "kubectl", "apply", "-f", "-")
	applyBucket.Stdin = strings.NewReader(bucketManifest)
	if output, err := applyBucket.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply StorageBucket: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for StorageBucket reconciliation (expected to fail with permission error)")
	// We wait for the Ready condition to be present
	bucketName := "kcc-test-" + ns
	if err := runCommand(ctx, t, root, "kubectl", "wait", "storagebucket", "-n", ns, bucketName, "--for=jsonpath={.status.conditions[].type}=Ready", "--timeout=5m"); err != nil {
		t.Fatalf("StorageBucket failed to reconcile: %v", err)
	}

	// Verify it failed as expected
	output, err := exec.CommandContext(ctx, "kubectl", "get", "storagebucket", "-n", ns, bucketName, "-o", "jsonpath={.status.conditions[?(@.type=='Ready')].status}").Output()
	if err != nil {
		t.Fatalf("failed to get StorageBucket status: %v", err)
	}
	status := strings.TrimSpace(string(output))
	if status != "False" {
		// Log more info for debugging
		describeOutput, _ := exec.CommandContext(ctx, "kubectl", "describe", "storagebucket", "-n", ns, bucketName).CombinedOutput()
		t.Logf("StorageBucket describe output:\n%s", string(describeOutput))
		t.Errorf("expected StorageBucket Ready status to be False, got %q", status)
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
