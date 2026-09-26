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

var imagesBuiltInProcess bool

func setupSmoketestCluster(t *testing.T, clusterPrefix string) (context.Context, string, *Harness) {
	t.Helper()
	if os.Getenv("E2E") != "1" {
		t.Skip("skipping smoketest; E2E=1 not set")
	}

	ctx := context.Background()

	repoRoot, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatalf("failed to get repo root: %v", err)
	}
	root := strings.TrimSpace(string(repoRoot))

	// Prepend GOPATH/bin and .build/bin to PATH so modern tools (e.g., kubectl v1.31+ with Kustomize v5)
	// take precedence over legacy binaries like /usr/local/kubebuilder/bin/kubectl (v1.16).
	if goPathOut, err := exec.Command("go", "env", "GOPATH").Output(); err == nil {
		goBin := filepath.Join(strings.TrimSpace(string(goPathOut)), "bin")
		buildBin := filepath.Join(root, ".build", "bin")
		t.Setenv("PATH", fmt.Sprintf("%s:%s:%s", goBin, buildBin, os.Getenv("PATH")))
	}

	logDiskUsage := func(label string) {
		out, err := exec.Command("df", "-h", "/").Output()
		if err == nil {
			t.Logf("[DISK USAGE - %s]\n%s", label, string(out))
		}
	}
	logDiskUsage("Test Start (" + t.Name() + ")")

	clusterName := fmt.Sprintf("%s-%s", clusterPrefix, strings.ToLower(time.Now().Format("20060102-150405")))

	// Cleanup cluster at the end and collect logs if failed
	t.Cleanup(func() {
		if t.Failed() {
			artifactsDir := os.Getenv("ARTIFACTS")
			if artifactsDir != "" {
				testArtifactsDir := filepath.Join(artifactsDir, t.Name())
				h, err := NewHarnessNoFatal(ctx, t)
				if err != nil {
					t.Logf("failed to create harness for artifact dumping: %v", err)
				} else {
					h.DumpArtifacts(testArtifactsDir)
				}
			}
		}

		t.Logf("Deleting kind cluster %q", clusterName)
		cmd := exec.CommandContext(ctx, "kind", "delete", "cluster", "--name", clusterName)
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Logf("failed to delete kind cluster: %v\nOutput: %s", err, string(output))
		}
	})

	t.Logf("[PHASE START] Creating kind cluster %q", clusterName)
	tKind := time.Now()
	if err := runCommand(ctx, t, root, "kind", "create", "cluster", "--name", clusterName); err != nil {
		t.Fatalf("failed to create kind cluster: %v", err)
	}
	t.Logf("[PHASE DONE] Kind cluster creation took %v", time.Since(tKind))
	logDiskUsage("After Kind Creation")

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
	originalManifestBytes := make(map[string][]byte)
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
				if strings.HasSuffix(path, "manager_sidecar_patch.yaml") || strings.HasSuffix(path, "recorder_sidecar_patch.yaml") {
					// Replace the prom-to-sd sidecar patch with a valid RFC 6902 no-op test so Kind does not fail with ImagePullBackOff on registry.kind/prometheus-to-sd
					newContent = "- op: test\n  path: /apiVersion\n  value: apps/v1\n"
				} else {
					newContent = strings.ReplaceAll(newContent, "imagePullPolicy: Always", "imagePullPolicy: IfNotPresent")
					newContent = strings.ReplaceAll(newContent, ":"+stableVersion, ":"+imageTag)
					// Inject METRICS_VERSION=v2 env var into all manager containers using our custom rewriter (issue #10260)
					newContent = InjectEnvVar(newContent, "METRICS_VERSION", "v2")
					if strings.HasSuffix(path, "per-namespace-components.yaml") || strings.HasSuffix(path, "components/manager/base/manager.yaml") {
						newContent = InjectNamespacedSecretVolume(newContent, "kcc-google-service-account")
					}
				}
				if newContent != string(content) {
					if _, exists := originalManifestBytes[path]; !exists {
						originalManifestBytes[path] = content
					}
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
	patchManifests(filepath.Join(root, "operator/config/channels"))
	patchManifests(filepath.Join(root, "config/installbundle/components"))

	// Revert patches at the end of the test to keep workspace clean
	t.Cleanup(func() {
		t.Logf("Reverting manifest patches")
		for path, origContent := range originalManifestBytes {
			if err := os.WriteFile(path, origContent, 0644); err != nil {
				t.Errorf("failed to revert manifest %s: %v", path, err)
			}
		}
	})

	// Always ensure kustomize image patch files exist (even when SKIP_BUILD=1 or BUILD_OPERATOR_ONLY=1 is set)
	ensureImagePatchFiles := func() {
		patches := []struct {
			tmpl string
			dst  string
			img  string
		}{
			{"config/installbundle/components/manager/base/manager_image_patch_template.yaml", "config/installbundle/components/manager/base/manager_image_patch.yaml", imagePrefix + "controller:" + imageTag},
			{"config/installbundle/components/recorder/recorder_image_patch_template.yaml", "config/installbundle/components/recorder/recorder_image_patch.yaml", imagePrefix + "recorder:" + imageTag},
			{"config/installbundle/components/webhook/webhook_image_patch_template.yaml", "config/installbundle/components/webhook/webhook_image_patch.yaml", imagePrefix + "webhook:" + imageTag},
			{"config/installbundle/components/deletiondefender/deletiondefender_image_patch_template.yaml", "config/installbundle/components/deletiondefender/deletiondefender_image_patch.yaml", imagePrefix + "deletiondefender:" + imageTag},
			{"config/installbundle/components/unmanageddetector/unmanageddetector_image_patch_template.yaml", "config/installbundle/components/unmanageddetector/unmanageddetector_image_patch.yaml", imagePrefix + "unmanageddetector:" + imageTag},
			{"operator/config/manager/manager_image_patch_template.yaml", "operator/config/manager/manager_image_patch.yaml", imagePrefix + "operator:" + imageTag},
			{"operator/config/autopilot-manager/manager_image_patch_template.yaml", "operator/config/autopilot-manager/manager_image_patch.yaml", imagePrefix + "operator:" + imageTag},
		}
		reImage := regexp.MustCompile(`image:\s+.*`)
		for _, p := range patches {
			tmplBytes, err := os.ReadFile(filepath.Join(root, p.tmpl))
			if err != nil {
				t.Fatalf("failed to read image patch template %s: %v", p.tmpl, err)
			}
			patched := reImage.ReplaceAllString(string(tmplBytes), "image: "+p.img)
			if err := os.WriteFile(filepath.Join(root, p.dst), []byte(patched), 0644); err != nil {
				t.Fatalf("failed to write image patch %s: %v", p.dst, err)
			}
		}
	}
	ensureImagePatchFiles()

	imagesToLoad := []string{
		"operator",
		"controller",
		"recorder",
		"webhook",
		"deletiondefender",
		"unmanageddetector",
	}

	allImagesExistLocally := func() bool {
		for _, img := range imagesToLoad {
			fullImage := imagePrefix + img + ":" + imageTag
			if err := exec.CommandContext(ctx, "docker", "image", "inspect", fullImage).Run(); err != nil {
				return false
			}
		}
		return true
	}

	if (os.Getenv("SKIP_BUILD") == "1" || os.Getenv("REUSE_IMAGES") == "1" || imagesBuiltInProcess) && os.Getenv("BUILD_OPERATOR_ONLY") != "1" && allImagesExistLocally() {
		t.Logf("Skipping image build step because SKIP_BUILD/REUSE_IMAGES is set or images were already built in this test run; reusing local images with tag %q", imageTag)
	} else if os.Getenv("BUILD_OPERATOR_ONLY") == "1" {
		t.Logf("[PHASE START] Building operator image only with tag %q", imageTag)
		tBuild := time.Now()
		buildCmd := exec.CommandContext(ctx, "docker", "buildx", "bake", "--load", "operator")
		buildCmd.Dir = root
		buildCmd.Env = append(os.Environ(),
			"IMAGE_TAG="+imageTag,
			"IMAGE_PREFIX="+imagePrefix,
		)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			t.Fatalf("failed to build operator image: %v", err)
		}
		imagesBuiltInProcess = true
		t.Logf("[PHASE DONE] Building operator image took %v", time.Since(tBuild))
		logDiskUsage("After Operator Docker Build")
	} else {
		if os.Getenv("SKIP_BUILD") == "1" || os.Getenv("REUSE_IMAGES") == "1" {
			t.Logf("SKIP_BUILD/REUSE_IMAGES was set, but one or more images with tag %q are missing locally; proceeding with image build", imageTag)
		}
		t.Logf("[PHASE START] Building images with tag %q", imageTag)
		tBuild := time.Now()
		buildCmd := exec.CommandContext(ctx, filepath.Join(root, "dev/tasks/build-images"))
		buildCmd.Dir = root
		buildCmd.Env = append(os.Environ(),
			"IMAGE_TAG="+imageTag,
			"IMAGE_PREFIX="+imagePrefix,
		)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			t.Fatalf("failed to build images: %v", err)
		}
		imagesBuiltInProcess = true
		t.Logf("[PHASE DONE] Building images took %v", time.Since(tBuild))
		logDiskUsage("After Docker Build")
	}

	t.Logf("Loading images into kind")
	kindArgs := []string{"load", "--name", clusterName, "docker-image"}
	for _, img := range imagesToLoad {
		fullImage := imagePrefix + img + ":" + imageTag
		kindArgs = append(kindArgs, fullImage)
	}
	t.Logf("[PHASE START] Bulk loading all %d images into kind cluster %q...", len(imagesToLoad), clusterName)
	tLoad := time.Now()
	if err := runCommand(ctx, t, root, "kind", kindArgs...); err != nil {
		t.Fatalf("failed to bulk load images into kind: %v", err)
	}
	t.Logf("[PHASE DONE] Bulk loading images took %v", time.Since(tLoad))
	logDiskUsage("After Kind Image Load")

	t.Logf("Deploying operator to kind")
	kustomizeCmd := exec.CommandContext(ctx, "kubectl", "kustomize", filepath.Join(root, "operator/config/default"))
	kustomizeOutput, err := kustomizeCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run kustomize: %v\nOutput: %s", err, string(kustomizeOutput))
	}

	manifests := string(kustomizeOutput)
	manifests = strings.ReplaceAll(manifests, "imagePullPolicy: Always", "imagePullPolicy: IfNotPresent")
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

	h := NewHarness(ctx, t)
	return ctx, root, h
}

func TestSmoketest(t *testing.T) {
	ctx, root, h := setupSmoketestCluster(t, "kcc-smoketest")

	// -------------------------------------------------------------------------
	// Phase 1 (Control Group): Default mode without partial CRD (all controllers)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 1 - CONTROL GROUP] Configuring ConfigConnector in default cluster mode (without resourceSettings)")
	defaultCCManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: cluster
  stateIntoSpec: Absent
  credentialSecretName: kcc-google-service-account
`
	applyDefaultCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyDefaultCC.Stdin = strings.NewReader(defaultCCManifest)
	if output, err := applyDefaultCC.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply default ConfigConnector: %v\nOutput: %s", err, string(output))
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

	t.Logf("[PHASE START] Waiting for KCC system components (webhook and controller managers) to be ready")
	tSystem := time.Now()
	if err := h.WaitForDeploymentAvailable("cnrm-system", "cnrm-webhook-manager", 5*time.Minute); err != nil {
		t.Fatalf("cnrm-webhook-manager failed to become ready: %v", err)
	}
	if err := h.WaitForStatefulSetReady("cnrm-system", "cnrm-controller-manager", 5*time.Minute); err != nil {
		t.Fatalf("cnrm-controller-manager failed to become ready: %v", err)
	}
	t.Logf("[PHASE DONE] KCC system components ready in %v", time.Since(tSystem))

	t.Logf("[CONTROL GROUP] Verifying the number of active watches / registered controllers in default mode without Partial CRD")
	defaultReconcileLines, defaultRawMetrics, err := h.GetControllerWorkerMetrics(1 * time.Minute)
	if err != nil {
		t.Fatalf("failed to get controller worker metrics in default mode: %v", err)
	}
	t.Logf("Found %d registered controller worker entries in default mode without Partial CRD", len(defaultReconcileLines))
	if len(defaultReconcileLines) < 100 {
		t.Errorf("expected >= 100 controller workers registered in default mode without partial CRD, but found %d. Raw metrics:\n%s", len(defaultReconcileLines), defaultRawMetrics)
	}

	// -------------------------------------------------------------------------
	// Phase 2 (Exclusive Mode): Partial CRD in Exclusive Mode (excluding StorageBucket)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 2 - EXCLUSIVE MODE] Updating ConfigConnector to Exclusive Partial CRD mode (excluding StorageBucket)")
	exclusiveCCManifest := `
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
      mode: exclude
      resources:
      - group: storage.cnrm.cloud.google.com
        kind: StorageBucket
`
	applyExclusiveCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyExclusiveCC.Stdin = strings.NewReader(exclusiveCCManifest)
	if output, err := applyExclusiveCC.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply exclusive ConfigConnector: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for operator to automatically roll out cnrm-controller-manager with updated cc-config-hash (exclusive mode)")
	exclusiveHash, err := h.WaitForStatefulSetConfigHashRollout("cnrm-system", "cnrm-controller-manager", "", 5*time.Minute)
	if err != nil {
		t.Fatalf("cnrm-controller-manager failed automatic rollout in exclusive mode: %v", err)
	}

	t.Logf("[EXCLUSIVE MODE] Verifying active watches / registered controllers in exclusive mode (StorageBucket excluded)")
	exclusiveReconcileLines, exclusiveRawMetrics, err := h.GetControllerWorkerMetrics(1 * time.Minute)
	if err != nil {
		t.Fatalf("failed to get controller worker metrics in exclusive mode: %v", err)
	}
	t.Logf("Found %d registered controller worker entries in exclusive mode", len(exclusiveReconcileLines))

	// Verify StorageBucket is excluded and other controllers (like PubSubTopic) remain registered
	hasStorageBucket := false
	hasPubSubTopic := false
	for _, line := range exclusiveReconcileLines {
		if strings.Contains(line, "storagebucket-parent-controller") {
			hasStorageBucket = true
		}
		if strings.Contains(line, "pubsubtopic-parent-controller") {
			hasPubSubTopic = true
		}
	}
	if hasStorageBucket {
		t.Errorf("expected storagebucket-parent-controller to be excluded in exclusive mode, but it was found in metrics:\n%s", strings.Join(exclusiveReconcileLines, "\n"))
	}
	if !hasPubSubTopic {
		t.Errorf("expected pubsubtopic-parent-controller to be registered in exclusive mode, but it was missing from metrics:\n%s", strings.Join(exclusiveReconcileLines, "\n"))
	}
	if len(exclusiveReconcileLines) != len(defaultReconcileLines)-1 {
		t.Errorf("expected controller count in exclusive mode (%d) to be exactly 1 less than default mode (%d). Raw metrics:\n%s", len(exclusiveReconcileLines), len(defaultReconcileLines), exclusiveRawMetrics)
	}

	// -------------------------------------------------------------------------
	// Phase 3 (Inclusive Mode): Partial CRD in Inclusive Mode (StorageBucket & PubSubTopic)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 3 - INCLUSIVE MODE] Updating ConfigConnector to Inclusive Partial CRD mode")
	inclusiveCCManifest := `
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
	applyInclusiveCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyInclusiveCC.Stdin = strings.NewReader(inclusiveCCManifest)
	if output, err := applyInclusiveCC.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply inclusive ConfigConnector: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for operator to automatically roll out cnrm-controller-manager with updated cc-config-hash (inclusive mode)")
	if _, err := h.WaitForStatefulSetConfigHashRollout("cnrm-system", "cnrm-controller-manager", exclusiveHash, 5*time.Minute); err != nil {
		t.Fatalf("cnrm-controller-manager failed automatic rollout in inclusive mode: %v", err)
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
	applyInterval := 500 * time.Millisecond
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
		describeOutput, _ := exec.CommandContext(ctx, "kubectl", "describe", "storagebucket", "-n", ns, bucketName).CombinedOutput()
		t.Logf("StorageBucket describe output:\n%s", string(describeOutput))
		t.Errorf("expected StorageBucket Ready status to be False, got %q", status)
	}

	// Verify the number of active watches / registered controllers via Prometheus metrics.
	t.Logf("Verifying the number of active watches / registered controllers via Prometheus metrics (issue #9651)")
	partialReconcileLines, partialRawMetrics, err := h.GetControllerWorkerMetrics(1 * time.Minute)
	if err != nil {
		t.Fatalf("failed to get controller worker metrics in partial CRD mode: %v", err)
	}

	t.Logf("Found %d kind worker entries in metrics in inclusive mode:\n%s", len(partialReconcileLines), strings.Join(partialReconcileLines, "\n"))

	if len(partialReconcileLines) == 0 {
		t.Fatalf("expected some kind worker entries to be registered, but found none in metrics. Raw metrics response:\n%s", partialRawMetrics)
	}
	if len(partialReconcileLines) > 5 {
		t.Errorf("expected only a few kind worker entries (due to inclusive mode specifying only StorageBucket and PubSubTopic), but found %d in metrics:\n%s", len(partialReconcileLines), strings.Join(partialReconcileLines, "\n"))
	}

	t.Logf("Cluster-mode smoketest completed successfully")
}

// TestSmoketestNamespaced verifies Partial CRD (Controlled CR Reconciliation) in Namespaced Mode
// across multiple tenant namespaces with their own ConfigConnectorContext (CCC) objects, including:
//  1. Initial dual-hash configuration (`cc-config-hash` & `ccc-config-hash`) and additive controller registration.
//  2. Namespace-scoped CCC updates triggering a rolling restart ONLY for the target namespace's StatefulSet.
//  3. Global CC updates triggering a rolling restart across ALL per-namespace StatefulSets.
func TestSmoketestNamespaced(t *testing.T) {
	ctx, root, h := setupSmoketestCluster(t, "kcc-smoketest-ns")

	// -------------------------------------------------------------------------
	// Phase 1: Initial Namespaced Mode Configuration (CC + Multiple Tenant CCCs)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 1 - INITIAL SETUP] Configuring ConfigConnector in namespaced mode with initial resourceSettings")
	namespacedCCInitialManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  stateIntoSpec: Absent
  experiments:
    resourceSettings:
      mode: include
      resources:
      - group: storage.cnrm.cloud.google.com
        kind: StorageBucket
`
	applyNamespacedCCInitial := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyNamespacedCCInitial.Stdin = strings.NewReader(namespacedCCInitialManifest)
	if output, err := applyNamespacedCCInitial.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply initial namespaced ConfigConnector: %v\nOutput: %s", err, string(output))
	}

	t.Logf("Waiting for required CRDs and cnrm-webhook-manager to be ready")
	for _, crd := range []string{
		"crd/storagebuckets.storage.cnrm.cloud.google.com",
		"crd/pubsubtopics.pubsub.cnrm.cloud.google.com",
		"crd/pubsubsubscriptions.pubsub.cnrm.cloud.google.com",
		"crd/bigquerydatasets.bigquery.cnrm.cloud.google.com",
		"crd/computenetworks.compute.cnrm.cloud.google.com",
	} {
		if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=create", crd, "--timeout=5m"); err != nil {
			t.Fatalf("%s not created: %v", crd, err)
		}
		if err := runCommand(ctx, t, root, "kubectl", "wait", "--for=condition=Established", crd, "--timeout=5m"); err != nil {
			t.Fatalf("%s not established: %v", crd, err)
		}
	}
	if err := h.WaitForDeploymentAvailable("cnrm-system", "cnrm-webhook-manager", 5*time.Minute); err != nil {
		t.Fatalf("cnrm-webhook-manager failed to become ready: %v", err)
	}

	type tenantConfig struct {
		namespace string
		group     string
		kind      string
		metricKey string
	}
	tenants := []tenantConfig{
		{
			namespace: "tenant-ns-1",
			group:     "pubsub.cnrm.cloud.google.com",
			kind:      "PubSubTopic",
			metricKey: "pubsubtopic-parent-controller",
		},
		{
			namespace: "tenant-ns-2",
			group:     "bigquery.cnrm.cloud.google.com",
			kind:      "BigQueryDataset",
			metricKey: "bigquerydataset-parent-controller",
		},
	}

	for _, tenant := range tenants {
		t.Logf("Creating namespace %q and its ConfigConnectorContext", tenant.namespace)
		if err := runCommand(ctx, t, root, "kubectl", "create", "ns", tenant.namespace); err != nil && !strings.Contains(err.Error(), "already exists") {
			t.Fatalf("failed to create namespace %s: %v", tenant.namespace, err)
		}
		cccManifest := fmt.Sprintf(`
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: %s
spec:
  googleServiceAccount: "fake-sa-%s@fake-project-id.iam.gserviceaccount.com"
  experiments:
    resourceSettings:
      mode: include
      resources:
      - group: %s
        kind: %s
`, tenant.namespace, tenant.namespace, tenant.group, tenant.kind)
		applyCCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
		applyCCC.Stdin = strings.NewReader(cccManifest)
		if output, err := applyCCC.CombinedOutput(); err != nil {
			t.Fatalf("failed to apply ConfigConnectorContext in %s: %v\nOutput: %s", tenant.namespace, err, string(output))
		}
	}

	initialStates := make(map[string]*NamespacedRolloutState, len(tenants))
	for _, tenant := range tenants {
		state, err := h.WaitForNamespacedStatefulSetConfigHashRollout(tenant.namespace, "", "", 5*time.Minute)
		if err != nil {
			t.Fatalf("failed waiting for initial manager StatefulSet rollout for namespace %s: %v", tenant.namespace, err)
		}
		initialStates[tenant.namespace] = state
		t.Logf("Initial StatefulSet ready for %s: sts=%s cc-config-hash=%s ccc-config-hash=%s rev=%s podUID=%s",
			tenant.namespace, state.StatefulSetName, state.CCHash, state.CCCHash, state.Revision, state.PodUID)
	}

	if initialStates["tenant-ns-1"].CCHash != initialStates["tenant-ns-2"].CCHash {
		t.Errorf("expected both CCC StatefulSets to share the same initial cc-config-hash, got %q vs %q",
			initialStates["tenant-ns-1"].CCHash, initialStates["tenant-ns-2"].CCHash)
	}
	if initialStates["tenant-ns-1"].CCCHash == initialStates["tenant-ns-2"].CCCHash {
		t.Errorf("expected distinct ccc-config-hash values for different CCC configs, got identical %q",
			initialStates["tenant-ns-1"].CCCHash)
	}

	// -------------------------------------------------------------------------
	// Phase 2: Local CCC Update (Scoped Restart for tenant-ns-1 Only)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 2 - CCC UPDATE] Updating ConfigConnectorContext in tenant-ns-1 to include ComputeNetwork")
	updatedTenant1CCCManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: tenant-ns-1
spec:
  googleServiceAccount: "fake-sa-tenant-ns-1@fake-project-id.iam.gserviceaccount.com"
  experiments:
    resourceSettings:
      mode: include
      resources:
      - group: pubsub.cnrm.cloud.google.com
        kind: PubSubTopic
      - group: compute.cnrm.cloud.google.com
        kind: ComputeNetwork
`
	applyUpdatedCCC := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyUpdatedCCC.Stdin = strings.NewReader(updatedTenant1CCCManifest)
	if output, err := applyUpdatedCCC.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply updated ConfigConnectorContext in tenant-ns-1: %v\nOutput: %s", err, string(output))
	}

	prevTenant1 := initialStates["tenant-ns-1"]
	afterCCCUpdateTenant1, err := h.WaitForNamespacedStatefulSetConfigHashRollout("tenant-ns-1", "", prevTenant1.CCCHash, 5*time.Minute)
	if err != nil {
		t.Fatalf("manager StatefulSet for tenant-ns-1 did not roll out after CCC update: %v", err)
	}
	t.Logf("StatefulSet for tenant-ns-1 rolled out after CCC update: sts=%s cc-config-hash=%s ccc-config-hash=%s rev=%s podUID=%s",
		afterCCCUpdateTenant1.StatefulSetName, afterCCCUpdateTenant1.CCHash, afterCCCUpdateTenant1.CCCHash, afterCCCUpdateTenant1.Revision, afterCCCUpdateTenant1.PodUID)

	if afterCCCUpdateTenant1.CCCHash == prevTenant1.CCCHash {
		t.Errorf("expected ccc-config-hash to change for tenant-ns-1 after CCC update, remained %q", afterCCCUpdateTenant1.CCCHash)
	}
	if afterCCCUpdateTenant1.CCHash != prevTenant1.CCHash {
		t.Errorf("expected cc-config-hash for tenant-ns-1 to remain %q when only CCC changed, got %q",
			prevTenant1.CCHash, afterCCCUpdateTenant1.CCHash)
	}
	if afterCCCUpdateTenant1.PodUID == prevTenant1.PodUID {
		t.Errorf("expected manager pod for tenant-ns-1 to restart (new PodUID) after CCC update, remained %q", afterCCCUpdateTenant1.PodUID)
	}

	tenant1LinesAfterCCC, _, err := h.GetNamespacedControllerWorkerMetrics("tenant-ns-1", 1*time.Minute)
	if err != nil {
		t.Fatalf("failed to get controller worker metrics for tenant-ns-1 after CCC update: %v", err)
	}
	hasComputeNetwork := false
	for _, line := range tenant1LinesAfterCCC {
		if strings.Contains(line, "computenetwork-parent-controller") {
			hasComputeNetwork = true
			break
		}
	}
	if !hasComputeNetwork {
		t.Errorf("expected computenetwork-parent-controller (from updated CCC) to be registered in tenant-ns-1 metrics:\n%s",
			strings.Join(tenant1LinesAfterCCC, "\n"))
	}

	// Verify tenant-ns-2 was NOT restarted by tenant-ns-1's CCC update
	prevTenant2 := initialStates["tenant-ns-2"]
	currentTenant2, err := h.WaitForNamespacedStatefulSetConfigHashRollout("tenant-ns-2", "", "", 1*time.Minute)
	if err != nil {
		t.Fatalf("failed to inspect StatefulSet state for tenant-ns-2: %v", err)
	}
	if currentTenant2.PodUID != prevTenant2.PodUID || currentTenant2.Revision != prevTenant2.Revision {
		t.Errorf("expected tenant-ns-2 manager pod not to restart when tenant-ns-1 CCC changed (prev PodUID=%s rev=%s, got PodUID=%s rev=%s)",
			prevTenant2.PodUID, prevTenant2.Revision, currentTenant2.PodUID, currentTenant2.Revision)
	}

	// Update baseline state for Phase 3
	initialStates["tenant-ns-1"] = afterCCCUpdateTenant1

	// -------------------------------------------------------------------------
	// Phase 3: Global CC Update (Rolling Restart Across All CCC StatefulSets)
	// -------------------------------------------------------------------------
	t.Logf("[PHASE 3 - CC UPDATE] Updating global ConfigConnector resourceSettings to trigger rolling restart for each CCC StatefulSet")
	namespacedCCUpdatedManifest := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  mode: namespaced
  stateIntoSpec: Absent
  experiments:
    resourceSettings:
      mode: include
      resources:
      - group: storage.cnrm.cloud.google.com
        kind: StorageBucket
      - group: pubsub.cnrm.cloud.google.com
        kind: PubSubSubscription
`
	applyNamespacedCCUpdated := exec.CommandContext(ctx, "kubectl", "apply", "--server-side", "-f", "-")
	applyNamespacedCCUpdated.Stdin = strings.NewReader(namespacedCCUpdatedManifest)
	if output, err := applyNamespacedCCUpdated.CombinedOutput(); err != nil {
		t.Fatalf("failed to apply updated namespaced ConfigConnector: %v\nOutput: %s", err, string(output))
	}

	updatedStates := make(map[string]*NamespacedRolloutState, len(tenants))
	for _, tenant := range tenants {
		prevState := initialStates[tenant.namespace]
		newState, err := h.WaitForNamespacedStatefulSetConfigHashRollout(tenant.namespace, prevState.CCHash, "", 5*time.Minute)
		if err != nil {
			t.Fatalf("manager StatefulSet for CCC in namespace %s did not roll out after CC update: %v", tenant.namespace, err)
		}
		updatedStates[tenant.namespace] = newState
		t.Logf("Updated StatefulSet rolled out for %s: sts=%s cc-config-hash=%s ccc-config-hash=%s rev=%s podUID=%s",
			tenant.namespace, newState.StatefulSetName, newState.CCHash, newState.CCCHash, newState.Revision, newState.PodUID)

		if newState.CCHash == prevState.CCHash {
			t.Errorf("expected cc-config-hash to change for %s after CC update, remained %q", tenant.namespace, newState.CCHash)
		}
		if newState.CCCHash != prevState.CCCHash {
			t.Errorf("expected ccc-config-hash for %s to remain %q when only CC changed, got %q",
				tenant.namespace, prevState.CCCHash, newState.CCCHash)
		}
		if newState.Revision == prevState.Revision {
			t.Errorf("expected StatefulSet revision for %s to advance after CC update, remained %q", tenant.namespace, newState.Revision)
		}
		if newState.PodUID == prevState.PodUID {
			t.Errorf("expected manager pod for %s to be restarted (new Pod UID) after CC update, remained %q", tenant.namespace, newState.PodUID)
		}

		// Verify controller metrics on each restarted per-namespace manager pod
		nsLines, _, err := h.GetNamespacedControllerWorkerMetrics(tenant.namespace, 1*time.Minute)
		if err != nil {
			t.Fatalf("failed to get controller worker metrics for namespace %s after CC update: %v", tenant.namespace, err)
		}
		hasGlobalPubSubSub := false
		hasTenantController := false
		for _, line := range nsLines {
			if strings.Contains(line, "pubsubsubscription-parent-controller") {
				hasGlobalPubSubSub = true
			}
			if strings.Contains(line, tenant.metricKey) {
				hasTenantController = true
			}
		}
		if !hasGlobalPubSubSub {
			t.Errorf("expected pubsubsubscription-parent-controller (from updated CC) to be registered in namespace %s metrics:\n%s",
				tenant.namespace, strings.Join(nsLines, "\n"))
		}
		if !hasTenantController {
			t.Errorf("expected %s (from namespace CCC) to be registered in namespace %s metrics:\n%s",
				tenant.metricKey, tenant.namespace, strings.Join(nsLines, "\n"))
		}
	}

	if updatedStates["tenant-ns-1"].CCHash != updatedStates["tenant-ns-2"].CCHash {
		t.Errorf("expected all CCC StatefulSets to converge on the same updated cc-config-hash, got %q vs %q",
			updatedStates["tenant-ns-1"].CCHash, updatedStates["tenant-ns-2"].CCHash)
	}

	t.Logf("Namespaced-mode smoketest completed successfully")
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
		time.Sleep(250 * time.Millisecond)
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
		time.Sleep(250 * time.Millisecond)
	}
}

func (h *Harness) WaitForStatefulSetConfigHashRollout(ns, name, prevHash string, timeout time.Duration) (string, error) {
	h.Logf("Waiting for statefulset %s/%s to roll out with a new cnrm.cloud.google.com/cc-config-hash (previous=%q)", ns, name, prevHash)
	stsGVR := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"}
	podGVR := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	deadline := time.Now().Add(timeout)
	for {
		if h.ctx.Err() != nil {
			return "", h.ctx.Err()
		}
		obj, err := h.dynamicClient.Resource(stsGVR).Namespace(ns).Get(h.ctx, name, metav1.GetOptions{})
		if err == nil {
			annotations, _, _ := unstructured.NestedStringMap(obj.Object, "spec", "template", "metadata", "annotations")
			newHash := annotations["cnrm.cloud.google.com/cc-config-hash"]
			gen := obj.GetGeneration()
			obsGen, _, _ := unstructured.NestedInt64(obj.Object, "status", "observedGeneration")
			currentRev, _, _ := unstructured.NestedString(obj.Object, "status", "currentRevision")
			updateRev, _, _ := unstructured.NestedString(obj.Object, "status", "updateRevision")
			updatedReplicas, _, _ := unstructured.NestedInt64(obj.Object, "status", "updatedReplicas")
			readyReplicas, _, _ := unstructured.NestedInt64(obj.Object, "status", "readyReplicas")

			if newHash != "" && newHash != prevHash && obsGen >= gen && currentRev != "" && currentRev == updateRev && updatedReplicas >= 1 && readyReplicas >= 1 {
				podName := fmt.Sprintf("%s-0", name)
				pod, err := h.dynamicClient.Resource(podGVR).Namespace(ns).Get(h.ctx, podName, metav1.GetOptions{})
				if err == nil && pod.GetDeletionTimestamp() == nil && pod.GetAnnotations()["cnrm.cloud.google.com/cc-config-hash"] == newHash {
					conditions, _, _ := unstructured.NestedSlice(pod.Object, "status", "conditions")
					for _, c := range conditions {
						if cond, ok := c.(map[string]interface{}); ok && cond["type"] == "Ready" && cond["status"] == "True" {
							h.Logf("statefulset %s/%s rolled out with cc-config-hash=%q", ns, name, newHash)
							return newHash, nil
						}
					}
				}
			}
		}

		if time.Now().After(deadline) {
			return "", fmt.Errorf("timeout waiting for statefulset %s/%s rollout with new cc-config-hash (previous=%q)", ns, name, prevHash)
		}
		time.Sleep(250 * time.Millisecond)
	}
}

type NamespacedRolloutState struct {
	StatefulSetName string
	CCHash          string
	CCCHash         string
	Revision        string
	PodUID          string
}

func (h *Harness) WaitForNamespacedStatefulSetConfigHashRollout(scopedNamespace, prevCCHash, prevCCCHash string, timeout time.Duration) (*NamespacedRolloutState, error) {
	h.Logf("Waiting for namespaced manager StatefulSet (scoped-namespace=%s) to roll out with cc-config-hash != %q and ccc-config-hash != %q", scopedNamespace, prevCCHash, prevCCCHash)
	stsGVR := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"}
	podGVR := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	labelSelector := fmt.Sprintf("cnrm.cloud.google.com/component=cnrm-controller-manager,cnrm.cloud.google.com/scoped-namespace=%s", scopedNamespace)
	deadline := time.Now().Add(timeout)

	for {
		if h.ctx.Err() != nil {
			return nil, h.ctx.Err()
		}
		stsList, err := h.dynamicClient.Resource(stsGVR).Namespace("cnrm-system").List(h.ctx, metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err == nil && len(stsList.Items) == 1 {
			sts := &stsList.Items[0]
			annotations, _, _ := unstructured.NestedStringMap(sts.Object, "spec", "template", "metadata", "annotations")
			ccHash := annotations["cnrm.cloud.google.com/cc-config-hash"]
			cccHash := annotations["cnrm.cloud.google.com/ccc-config-hash"]
			gen := sts.GetGeneration()
			obsGen, _, _ := unstructured.NestedInt64(sts.Object, "status", "observedGeneration")
			currentRev, _, _ := unstructured.NestedString(sts.Object, "status", "currentRevision")
			updateRev, _, _ := unstructured.NestedString(sts.Object, "status", "updateRevision")
			updatedReplicas, _, _ := unstructured.NestedInt64(sts.Object, "status", "updatedReplicas")
			readyReplicas, _, _ := unstructured.NestedInt64(sts.Object, "status", "readyReplicas")

			if ccHash != "" && ccHash != prevCCHash && cccHash != "" && cccHash != prevCCCHash && obsGen >= gen && currentRev != "" && currentRev == updateRev && updatedReplicas >= 1 && readyReplicas >= 1 {
				podName := fmt.Sprintf("%s-0", sts.GetName())
				pod, err := h.dynamicClient.Resource(podGVR).Namespace("cnrm-system").Get(h.ctx, podName, metav1.GetOptions{})
				if err == nil && pod.GetDeletionTimestamp() == nil &&
					pod.GetAnnotations()["cnrm.cloud.google.com/cc-config-hash"] == ccHash &&
					pod.GetAnnotations()["cnrm.cloud.google.com/ccc-config-hash"] == cccHash {
					conditions, _, _ := unstructured.NestedSlice(pod.Object, "status", "conditions")
					for _, c := range conditions {
						if cond, ok := c.(map[string]interface{}); ok && cond["type"] == "Ready" && cond["status"] == "True" {
							return &NamespacedRolloutState{
								StatefulSetName: sts.GetName(),
								CCHash:          ccHash,
								CCCHash:         cccHash,
								Revision:        currentRev,
								PodUID:          string(pod.GetUID()),
							}, nil
						}
					}
				}
			}
		}

		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout waiting for namespaced StatefulSet rollout for namespace %s (prevCCHash=%q, prevCCCHash=%q)", scopedNamespace, prevCCHash, prevCCCHash)
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func (h *Harness) GetNamespacedControllerWorkerMetrics(scopedNamespace string, timeout time.Duration) ([]string, string, error) {
	deadline := time.Now().Add(timeout)
	labelSelector := fmt.Sprintf("cnrm.cloud.google.com/component=cnrm-controller-manager,cnrm.cloud.google.com/scoped-namespace=%s", scopedNamespace)
	for {
		if h.ctx.Err() != nil {
			return nil, "", h.ctx.Err()
		}
		podNameCmd := exec.CommandContext(h.ctx, "kubectl", "get", "pods", "-n", "cnrm-system", "-l", labelSelector, "-o", "jsonpath={.items[0].metadata.name}")
		podNameBytes, err := podNameCmd.Output()
		if err != nil || strings.TrimSpace(string(podNameBytes)) == "" {
			if time.Now().After(deadline) {
				return nil, "", fmt.Errorf("timeout waiting for manager pod in namespace %s: %w", scopedNamespace, err)
			}
			time.Sleep(250 * time.Millisecond)
			continue
		}
		podName := strings.TrimSpace(string(podNameBytes))

		metricsCmd := exec.CommandContext(h.ctx, "kubectl", "get", "--raw", fmt.Sprintf("/api/v1/namespaces/cnrm-system/pods/%s:8888/proxy/metrics.v2", podName))
		metricsBytes, err := metricsCmd.Output()
		if err != nil {
			if time.Now().After(deadline) {
				return nil, "", fmt.Errorf("timeout getting metrics from pod %s for namespace %s: %w", podName, scopedNamespace, err)
			}
			time.Sleep(250 * time.Millisecond)
			continue
		}
		rawMetrics := string(metricsBytes)

		var reconcileLines []string
		for _, line := range strings.Split(rawMetrics, "\n") {
			if strings.Contains(line, "controller_runtime_max_concurrent_reconciles") && !strings.HasPrefix(line, "#") {
				reconcileLines = append(reconcileLines, line)
			}
		}
		if len(reconcileLines) > 0 {
			return reconcileLines, rawMetrics, nil
		}
		if time.Now().After(deadline) {
			return nil, rawMetrics, fmt.Errorf("timeout: no controller worker metric entries found for namespace %s", scopedNamespace)
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func (h *Harness) GetControllerWorkerMetrics(timeout time.Duration) ([]string, string, error) {
	deadline := time.Now().Add(timeout)
	for {
		if h.ctx.Err() != nil {
			return nil, "", h.ctx.Err()
		}
		// 1. Get the pod name of cnrm-controller-manager
		podNameCmd := exec.CommandContext(h.ctx, "kubectl", "get", "pods", "-n", "cnrm-system", "-l", "cnrm.cloud.google.com/component=cnrm-controller-manager", "-o", "jsonpath={.items[0].metadata.name}")
		podNameBytes, err := podNameCmd.Output()
		if err != nil {
			h.Logf("failed to get manager pod name: %v\nOutput: %s", err, string(podNameBytes))
			if time.Now().After(deadline) {
				return nil, "", fmt.Errorf("timeout waiting for cnrm-controller-manager pod name: %w", err)
			}
			time.Sleep(250 * time.Millisecond)
			continue
		}
		podName := strings.TrimSpace(string(podNameBytes))
		if podName == "" {
			if time.Now().After(deadline) {
				return nil, "", fmt.Errorf("timeout waiting for cnrm-controller-manager pod to start")
			}
			time.Sleep(250 * time.Millisecond)
			continue
		}

		// 2. Query metrics from the pod via kubectl API proxy. Since we use METRICS_VERSION=v2,
		// the standard controller-runtime metrics are served on /metrics.v2.
		metricsCmd := exec.CommandContext(h.ctx, "kubectl", "get", "--raw", fmt.Sprintf("/api/v1/namespaces/cnrm-system/pods/%s:8888/proxy/metrics.v2", podName))
		metricsBytes, err := metricsCmd.Output()
		if err != nil {
			h.Logf("failed to get metrics from pod proxy: %v, retrying...", err)
			if time.Now().After(deadline) {
				return nil, "", fmt.Errorf("timeout getting metrics from pod proxy: %w", err)
			}
			time.Sleep(250 * time.Millisecond)
			continue
		}
		rawMetrics := string(metricsBytes)

		// Filter for lines containing "controller_runtime_max_concurrent_reconciles"
		var reconcileLines []string
		for _, line := range strings.Split(rawMetrics, "\n") {
			if strings.Contains(line, "controller_runtime_max_concurrent_reconciles") && !strings.HasPrefix(line, "#") {
				reconcileLines = append(reconcileLines, line)
			}
		}

		if len(reconcileLines) > 0 {
			return reconcileLines, rawMetrics, nil
		}

		if time.Now().After(deadline) {
			return nil, rawMetrics, fmt.Errorf("timeout: no controller worker metric entries found in metrics:\n%s", rawMetrics)
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func (h *Harness) MustWaitForObservedGeneration(gvr schema.GroupVersionResource, ns, name string, expectedGeneration int64) {
	h.Logf("Waiting for %s/%s in namespace %s to have observedGeneration >= %d", gvr.Resource, name, ns, expectedGeneration)
	timeout := 5 * time.Minute
	pollInterval := 250 * time.Millisecond
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
