// Copyright 2025 Google LLC
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

package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	kindImage = "kindest/node:v1.29.2"
)

func TestMultiClusterLeaderElection(t *testing.T) {
	g := gomega.NewWithT(t)
	ctx := context.Background()

	uniqueID := uuid.New().String()
	clusterAName := "cluster-a-" + uniqueID
	clusterBName := "cluster-b-" + uniqueID

	// Create two kind clusters
	clusterAClient := createKindCluster(ctx, t, clusterAName)
	clusterBClient := createKindCluster(ctx, t, clusterBName)

	defer deleteKindCluster(ctx, t, clusterAName)
	defer deleteKindCluster(ctx, t, clusterBName)

	// Start MinIO
	minioContainerName := "minio-" + uniqueID
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "-d", "--name", minioContainerName, "-p", "9000:9000", "-p", "9001:9001", "minio/minio", "server", "/data", "--console-address", ":9001")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to start minio container: %v\nOutput:\n%s", err, string(output))
	}
	defer func() {
		cmd := exec.Command("docker", "stop", minioContainerName)
		_ = cmd.Run()
	}()

	// Configure MinIO
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	g.Expect(err).NotTo(gomega.HaveOccurred())

	g.Eventually(func() error {
		return minioClient.MakeBucket(ctx, "leader-election", minio.MakeBucketOptions{})
	}, 30*time.Second, 1*time.Second).Should(gomega.Succeed(), "should be able to create minio bucket")

	// Deploy KCC and the election controller to both clusters
	deployToCluster(ctx, g, clusterAClient)
	deployToCluster(ctx, g, clusterBClient)

	// Verify that only one KCC instance becomes the leader
	var leader string
	g.Eventually(func() (string, error) {
		leaderA, _ := getLeader(ctx, clusterAClient)
		leaderB, _ := getLeader(ctx, clusterBClient)

		if leaderA != "" && leaderB != "" && leaderA != leaderB {
			return "", fmt.Errorf("split brain detected: leaderA=%s, leaderB=%s", leaderA, leaderB)
		}
		if leaderA != "" {
			return leaderA, nil
		}
		if leaderB != "" {
			return leaderB, nil
		}
		return "", nil
	}, 2*time.Minute, 5*time.Second).ShouldNot(gomega.BeEmpty(), "leader should be elected")

	leader, err = getLeader(ctx, clusterAClient)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	if leader == "" {
		leader, err = getLeader(ctx, clusterBClient)
		g.Expect(err).NotTo(gomega.HaveOccurred())
	}

	// Simulate leader failure
	var leaderCluster client.Client
	var standbyCluster client.Client
	if strings.Contains(leader, "kcc-a") { // Assuming pod name contains cluster identifier
		leaderCluster = clusterAClient
		standbyCluster = clusterBClient
	} else {
		leaderCluster = clusterBClient
		standbyCluster = clusterAClient
	}

	podToDelete := &corev1.Pod{}
	podToDelete.Name = leader
	podToDelete.Namespace = "kcc-system"
	err = leaderCluster.Delete(ctx, podToDelete)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Verify failover
	g.Eventually(func() (string, error) {
		return getLeader(ctx, standbyCluster)
	}, 2*time.Minute, 5*time.Second).ShouldNot(gomega.Or(gomega.BeEmpty(), gomega.Equal(leader)), "new leader should be elected")
}

func createKindCluster(ctx context.Context, t *testing.T, name string) client.Client {
	cmd := exec.CommandContext(ctx, "kind", "create", "cluster", "--name", name, "--image", kindImage)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to create kind cluster %s: %v\nOutput:\n%s", name, err, string(output))
	}

	cmd = exec.CommandContext(ctx, "kind", "get", "kubeconfig", "--name", name)
	kubeconfig, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to get kubeconfig for %s: %v", name, err)
	}

	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		t.Fatalf("failed to get rest config for %s: %v", name, err)
	}

	c, err := client.New(restConfig, client.Options{})
	if err != nil {
		t.Fatalf("failed to create client for %s: %v", name, err)
	}
	return c
}

func deleteKindCluster(ctx context.Context, t *testing.T, name string) {
	cmd := exec.CommandContext(ctx, "kind", "delete", "cluster", "--name", name)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Logf("failed to delete kind cluster %s: %v\nOutput:\n%s", name, err, string(output))
	}
}

func applyManifest(ctx context.Context, g *gomega.WithT, c client.Client, manifest string) {
	objs := splitYAML(g, []byte(manifest))
	for _, obj := range objs {
		err := c.Create(ctx, obj)
		g.Expect(err).NotTo(gomega.HaveOccurred())
	}
}

func applyManifestFile(ctx context.Context, g *gomega.WithT, c client.Client, manifestPath string) {
	content, err := os.ReadFile(manifestPath)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	applyManifest(ctx, g, c, string(content))
}

func deployToCluster(ctx context.Context, g *gomega.WithT, c client.Client) {
	// Deploy the election controller
	applyManifestFile(ctx, g, c, filepath.Join("..", "..", "..", "multicluster-leader-election", "config", "crd", "bases", "multicluster.core.cnrm.cloud.google.com_multiclusterleases.yaml"))
	applyManifestFile(ctx, g, c, filepath.Join("..", "..", "..", "multicluster-leader-election", "config", "rbac", "role.yaml"))
	applyManifestFile(ctx, g, c, filepath.Join("..", "..", "..", "multicluster-leader-election", "config", "rbac", "role_binding.yaml"))
	applyManifestFile(ctx, g, c, filepath.Join("..", "..", "..", "multicluster-leader-election", "config", "manager", "manager.yaml"))

	// Deploy KCC
	applyManifestFile(ctx, g, c, filepath.Join("..", "..", "config", "install-bundle", "install-bundle-namespaced.yaml"))

	// Create ConfigConnector object
	ccYAML := `
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
  leaderElection:
    multiClusterLease:
      leaseName: kcc-leader-lease
      namespace: kcc-system
      globalLockName: kcc-global-lock
`
	applyManifest(ctx, g, c, ccYAML)
}

func getLeader(ctx context.Context, c client.Client) (string, error) {
	pods := &corev1.PodList{}
	err := c.List(ctx, pods, client.InNamespace("kcc-system"), client.MatchingLabels{"cnrm.cloud.google.com/component": "cnrm-controller-manager"})
	if err != nil {
		return "", err
	}
	for _, pod := range pods.Items {
		if val, ok := pod.Annotations["cnrm.cloud.google.com/leader"]; ok && val == "true" {
			return pod.Name, nil
		}
	}
	return "", nil
}

func splitYAML(g *gomega.WithT, content []byte) []*unstructured.Unstructured {
	var result []*unstructured.Unstructured
	parts := strings.Split(string(content), "---")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		obj := &unstructured.Unstructured{}
		g.Expect(yaml.Unmarshal([]byte(part), obj)).To(gomega.Succeed())
		result = append(result, obj)
	}
	return result
}
