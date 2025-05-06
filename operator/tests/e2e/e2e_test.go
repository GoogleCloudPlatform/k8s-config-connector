// Copyright 2022 Google LLC
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
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/scripts/utils"
	kcck8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"

	"github.com/blang/semver"
	"github.com/cenkalti/backoff"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/cloudresourcemanager/v1"
	containerBeta "google.golang.org/api/container/v1beta1"
	"google.golang.org/api/iam/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // Load the gcp auth plugin
	"k8s.io/client-go/tools/clientcmd"
)

const (
	ServiceAccID = "cnrm-system"
	SecretName   = "gsa-key"

	// projects, by default, are limited to 3 GKE clusters per region, for that reason we run the tests in two regions
	GkeClusterZone1      = "us-central1-a"
	GkeClusterZone2      = "us-west1-a"
	GkeClusterZone3      = "us-west2-a"
	KubectlDeleteTimeout = 5 * time.Minute

	OperatorReleaseBucket  = "kcc-operator-internal"
	OperatorReleaseTarball = "release-bundle.tar.gz"
	KccReleaseBucket       = "cnrm"
	KccReleaseTarball      = "release-bundle.tar.gz"
	// Use ConfigConnector operator 1.95.0 as the base version for upgrade test
	// because previous versions cannot be installed onto K8s 1.26+ (default GKE
	// cluster version as of 6/27/2023) due to the removal of
	// autoscaling/v2beta2.
	BaseVersionSha = "4119846"
)

var (
	services = []string{
		"container.googleapis.com",
		"iamcredentials.googleapis.com",
		"artifactregistry.googleapis.com",
	}
	organization        = testgcp.TestOrgID.Get()
	billingAccount      = testgcp.TestBillingAccountID.Get()
	f                   = &flags{}
	defaultBackOff      = wait.Backoff{Steps: 5, Duration: 500 * time.Millisecond, Factor: 1.5}
	longIntervalBackOff = wait.Backoff{Steps: 3, Duration: 2 * time.Minute, Factor: 1}
)

type TestOptions struct {
	OrganizationID     string
	BillingAccountID   string
	ServiceAccountID   string
	GKEClusterLocation string
	BaseVersionSHA     string
	ProjectID          string
	SecretName         string
}

type flags struct {
	projectID string
	cleanup   bool
	version   string
}

type cluster struct {
	kubectl   *kubectl
	clientset *kubernetes.Clientset
	log       logr.Logger
}

type kubectl struct {
	kubeconfigPath string
	deleteTimeout  time.Duration
}

type configConnectorSample struct {
	configConnectorClusterModeWorkloadIdentityYAMLPath string
	configConnectorClusterModeGCPIdentityYAMLPath      string
	configConnectorNamespacedModeYAMLPath              string
	configConnectorContextYAMLPath                     string
}

type cleanupFunc func() error

func TestMain(m *testing.M) {
	flag.StringVar(&f.projectID, "project-id", "", "Project ID that will be used for the project created for E2E tests.")
	flag.BoolVar(&f.cleanup, "cleanup", true, "If true, project and clusters created for testing will be deleted before exiting the test suite. "+
		"Set to false if you want to keep clusters for debugging when running the test locally.")
	flag.StringVar(&f.version, "version", "latest", "Version of the KCC Operator to use for E2E tests. "+
		"The version of the KCC Operator is defined by SHORT_SHA in release.sh. Use the corresponding SHORT_SHA value to run the e2e test against a particular commit. "+
		"The default value is 'latest', which represents the version of the last green canary candidate promoted by the periodic-kcc-operator-release prow job.")
	flag.Parse()
	if f.projectID == "" {
		fmt.Println("error parsing command line flags: project-id is required")
		os.Exit(1)
	}

	log, err := newLogger("TestMain")
	if err != nil {
		fmt.Print(fmt.Errorf("error creating logger: %w", err))
		os.Exit(1)
	}

	log.Info("Setting up a project for E2E tests...")
	deleteProject, err := setupProject(organization, f.projectID, billingAccount, ServiceAccID, log)
	if err != nil {
		log.Error(err, "error setting up project\r\n",
			"Organization", organization, "ProjectID", f.projectID,
			"BillingAccount", billingAccount, "ServiceId", ServiceAccID)
		cleanUpProject(deleteProject, f.cleanup, log)
		os.Exit(1)
	}
	log.Info("Beginning tests...")
	exitCode := m.Run()
	cleanUpProject(deleteProject, f.cleanup, log)
	os.Exit(exitCode)
}

func TestKCCInstallAndUninstall_Namespaced(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone1
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}

	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC...")
	if err := cluster.installKCC(sample.configConnectorNamespacedModeYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.enableKCCForNamespace(namespace, sample.configConnectorContextYAMLPath, testOptions.ServiceAccountID, testOptions.ProjectID); err != nil {
		t.Fatal(fmt.Errorf("error enabling KCC for namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(err)
	}
	log.Info("Uninstalling KCC...")
	if err := cluster.uninstallKCC(); err != nil {
		t.Fatal(fmt.Errorf("error uninstalling KCC: %w", err))
	}
}

func TestKCCInstallAnd_Delete_Namespace_In_Namespaced_Mode(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone3
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}
	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC...")
	if err := cluster.installKCC(sample.configConnectorNamespacedModeYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.enableKCCForNamespace(namespace, sample.configConnectorContextYAMLPath, testOptions.ServiceAccountID, testOptions.ProjectID); err != nil {
		t.Fatal(fmt.Errorf("error enabling KCC for namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	// add an extra finalizer to ensure resources are not deleted, the config-connector-operator should wait
	// until all KCC resources are deleted before deleting the related KCC pods
	log.Info("Adding custom finalizer to prevent deletion...")
	extraFinalizer := "extra-finalizer"
	if err := cluster.addFinalizerToArtifactRegistryRepository(namespace, repoName, extraFinalizer); err != nil {
		t.Fatal(fmt.Errorf("error adding finalizer to ArtifactRegistryRepository: %w", err))
	}
	log.Info("Deleting Namespace...")
	if err := cluster.deleteNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error deleting namespace: %w", err))
	}
	// Sometimes, it takes a long time for k8s to cascade delete KCC resource CRs under the deleted namespace.
	// Therefore we perform a direct deletion on the ArtifactRegistryRepository object to speed things up.
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName, "--wait=false"); err != nil {
		t.Fatal(err)
	}
	// The CNRM manager pod should still be running as the operator should wait until all CNRM resources deleted before
	// deleting the manager pods. This check ensures the manager is able to remove its finalizer from the ArtifactRegistryRepository.
	log.Info("Waiting for CNRM finalizer to be removed from ArtifactRegistryRepository...")
	if err := cluster.waitForCNRMFinalizersToBeRemovedFromArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(fmt.Errorf("error waiting for CNRM finalizer to be removed from ArtifactRegistryRepository: %w", err))
	}
	// The config connector context should NOT be removed as the ArtifactRegistryRepository has not yet been removed due to its extra finalizer
	log.Info("Verifying the ConfigConnectorContext still exists but is unhealthy...")
	if err := cluster.waitForConfigConnectorContextToBeUnhealthy(namespace, corev1beta1.ConfigConnectorContextAllowedName); err != nil {
		t.Fatal(fmt.Errorf("error verifying the ConfigConnectorContext's health: %w", err))
	}
	log.Info("Removing custom finalizer to enable deletion...")
	if err := cluster.removeFinalizerToArtifactRegistryRepository(namespace, repoName, extraFinalizer); err != nil {
		t.Fatal(fmt.Errorf("error removing finalizer from ArtifactRegistryRepository: %w", err))
	}
	log.Info("Waiting for ConfigConnectorContext to be removed...")
	if err := cluster.waitForConfigConnectorContextToBeRemoved(namespace, corev1beta1.ConfigConnectorContextAllowedName); err != nil {
		t.Fatal(fmt.Errorf("error waiting for ConfigConnectorContextToBeRemoved: %w", err))
	}
	log.Info("Waiting for namespace to be deleted...")
	if err := cluster.waitForNamespaceToBeDeleted(namespace); err != nil {
		t.Fatalf("error waiting for namespace to be deleted")
	}
}

func TestKCCInstallAndUninstall_Cluster_WorkloadIdentity(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone2
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}

	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC in cluster mode with workload identity...")
	if err := cluster.installKCC(sample.configConnectorClusterModeWorkloadIdentityYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(err)
	}
	log.Info("Uninstalling KCC...")
	if err := cluster.uninstallKCC(); err != nil {
		t.Fatal(fmt.Errorf("error uninstalling KCC: %w", err))
	}
}

func TestKCCInstallAndUninstall_Cluster_GCPIdentity(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone2
	testOptions.SecretName = SecretName
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}

	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC in cluster mode with GCP identity...")
	if err := cluster.installKCC(sample.configConnectorClusterModeGCPIdentityYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(err)
	}
	log.Info("Uninstalling KCC...")
	if err := cluster.uninstallKCC(); err != nil {
		t.Fatal(fmt.Errorf("error uninstalling KCC: %w", err))
	}
}

func TestKCCInstallAndUninstallWithoutDeletingKCCResources(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone1
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}

	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC...")
	if err := cluster.installKCC(sample.configConnectorNamespacedModeYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.enableKCCForNamespace(namespace, sample.configConnectorContextYAMLPath, testOptions.ServiceAccountID, testOptions.ProjectID); err != nil {
		t.Fatal(fmt.Errorf("error enabling KCC for namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	log.Info("Uninstalling KCC...")
	if err := cluster.uninstallKCC(); err != nil {
		t.Fatal(fmt.Errorf("error uninstalling KCC: %w", err))
	}
	if err := checkArtifactRegistryRepositoryExistsOnGCP(repoName, f.projectID); err != nil {
		t.Fatal(err)
	}
}

func TestShouldNotBeAbleToCreateKCCResourcesIfKCCNotEnabledForNamespace(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone1
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}

	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the operator: %w", err))
	}
	log.Info("Installing KCC...")
	if err := cluster.installKCC(sample.configConnectorNamespacedModeYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepositoryShouldFail(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	ok, err := cluster.doesArtifactRegistryRepositoryHaveFinalizer(namespace, repoName, k8s.KCCFinalizer)
	if err != nil {
		t.Fatal(fmt.Errorf("error checking if ArtifactRegistryRepository has finalizer: %w", err))
	}
	if ok {
		t.Fatalf("expected ArtifactRegistryRepository to not have finalizer '%v', but it does", k8s.KCCFinalizer)
	}
	ok, err = cluster.doesArtifactRegistryRepositoryHaveStatusUnmanaged(namespace, repoName)
	if err != nil {
		t.Fatal(fmt.Errorf("error checking if ArtifactRegistryRepository has status '%v': %w", kcck8s.Unmanaged, err))
	}
	if !ok {
		t.Fatalf("expected ArtifactRegistryRepository to have status '%v', but it does not", kcck8s.Unmanaged)
	}
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(err)
	}
}

func TestUpgrade(t *testing.T) {
	t.Parallel()
	testOptions := newTestOptions()
	testOptions.GKEClusterLocation = GkeClusterZone2
	testID, log, cluster, teardown := setup(t, testOptions)
	if f.cleanup {
		defer teardown()
	}
	//Get older version of the operator to perform an upgrade against
	manifestsDir, sample, err := getOperatorReleaseAssetsForVersion(testOptions.BaseVersionSHA, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version %v: %w", testOptions.BaseVersionSHA, err))
	}
	log.Info("Installing the base version operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the base version operator: %w", err))
	}
	log.Info("Installing KCC...")
	if err := cluster.installKCC(sample.configConnectorNamespacedModeYAMLPath); err != nil {
		t.Fatal(fmt.Errorf("error installing KCC: %w", err))
	}
	namespace := "e2e-test-namespace"
	if err := cluster.createNamespace(namespace); err != nil {
		t.Fatal(fmt.Errorf("error creating namespace '%v': %w", namespace, err))
	}
	if err := cluster.enableKCCForNamespace(namespace, sample.configConnectorContextYAMLPath, testOptions.ServiceAccountID, testOptions.ProjectID); err != nil {
		t.Fatal(fmt.Errorf("error enabling KCC for namespace '%v': %w", namespace, err))
	}
	if err := cluster.addProjectIDAnnotationToNamespace(namespace, f.projectID); err != nil {
		t.Fatal(fmt.Errorf("error annotating namespace '%v' with the project ID: %w", namespace, err))
	}
	kccVersion, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error determining KCC version: %w", err))
	}
	log.Info("Downloading and extracting KCC release tarball ...", "version", kccVersion)
	kccReleaseAssetsDir, err := createTempDir("e2e-kcc-release-assets")
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temporary directory for KCC release assets: %w", err))
	}
	if err := downloadAndExtractKCCReleaseTarball(kccVersion, kccReleaseAssetsDir); err != nil {
		t.Fatal(fmt.Errorf("error downloading and extracting KCC with version '%v': %w", kccVersion, err))
	}
	repoName, repoYAMLDir, err := getArtifactRegistryRepositorySample(kccReleaseAssetsDir, testID, f.projectID)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting ArtifactRegistryRepository sample from KCC release assets: %w", err))
	}
	log.Info("Creating ArtifactRegistryRepository...")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error creating ArtifactRegistryRepository: %w", err))
	}
	log.Info("Upgrading the operator to the latest version")
	manifestsDir, _, err = getOperatorReleaseAssetsForVersion(f.version, testOptions.ServiceAccountID, testOptions.ProjectID, log)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting operator release assets for version '%v': %w", f.version, err))
	}
	log.Info("Installing the latest operator...")
	if err := cluster.installOperator(manifestsDir); err != nil {
		t.Fatal(fmt.Errorf("error installing the latest operator: %w", err))
	}
	time.Sleep(120 * time.Second) // Some buffer time for the operator to reconcile on the existing ConfigConnector
	if err := cluster.waitForConfigConnectorToBeHealthy(corev1beta1.ConfigConnectorAllowedName); err != nil {
		t.Fatal(fmt.Errorf("error waitting for ConfigConnector to be healthy: %w", err))
	}
	checkIfKCCHasUpgradedToTheLatestVersion(t, cluster, log)
	log.Info("Re-applying ArtifactRegistryRepository")
	if err := cluster.createArtifactRegistryRepository(namespace, repoName, repoYAMLDir); err != nil {
		t.Fatal(fmt.Errorf("error re-applying ArtifactRegistryRepository: %w", err))
	}
	log.Info("Deleting ArtifactRegistryRepository...")
	if err := cluster.deleteArtifactRegistryRepository(namespace, repoName); err != nil {
		t.Fatal(err)
	}
	log.Info("Deleting ConfigConnectorContext...")
	if err := cluster.deleteConfigConnectorContext(namespace, corev1beta1.ConfigConnectorContextAllowedName); err != nil {
		t.Fatal(err)
	}
	log.Info("Uninstalling KCC...")
	if err := cluster.uninstallKCC(); err != nil {
		t.Fatal(fmt.Errorf("error uninstalling KCC: %w", err))
	}
}

func newTestOptions() TestOptions {
	return TestOptions{
		OrganizationID:   organization,
		BillingAccountID: billingAccount,
		ServiceAccountID: ServiceAccID,
		BaseVersionSHA:   BaseVersionSha,
		ProjectID:        f.projectID,
	}
}

func checkIfKCCHasUpgradedToTheLatestVersion(t *testing.T, cluster *cluster, log logr.Logger) {
	curKccVersionRaw, err := cluster.getKCCVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error getting the current KCC version: %w", err))
	}
	currentKccVersion, err := semver.ParseTolerant(curKccVersionRaw)
	if err != nil {
		t.Fatalf("current KCC version %v is not a valid semantic version", curKccVersionRaw)
	}
	latestOperatorVersionRaw, err := cluster.getOperatorVersion()
	if err != nil {
		t.Fatal(fmt.Errorf("error getting the latest operator version: %w", err))
	}
	latestOperatorVersion, err := semver.ParseTolerant(latestOperatorVersionRaw)
	if err != nil {
		t.Fatalf("latest Operator version %v is not a valid semantic version", curKccVersionRaw)
	}
	log.Info("Version checking", "currentKCCVersion", currentKccVersion, "latestOperatorVersion", latestOperatorVersion)
	// only compare major.minor.patch as the operator version might contain operator specific extension e.g. 1.6.0-operator.x
	if currentKccVersion.Major != latestOperatorVersion.Major || currentKccVersion.Minor != latestOperatorVersion.Minor || currentKccVersion.Patch != latestOperatorVersion.Patch {
		t.Fatalf("expect to have KCC upgraded to %v, but it's still on version %v", latestOperatorVersion, currentKccVersion)
	}
}

func setup(t *testing.T, testOptions TestOptions) (testID string, log logr.Logger, cluster *cluster, teardown func()) {
	testID = newUniqueTestID()
	log, err := newLogger(t.Name())
	if err != nil {
		t.Fatal(fmt.Errorf("error creating logger: %w", err))
	}
	clusterName := "e2e-test-" + testID
	cluster, cleanup, err := setupCluster(clusterName, testOptions.ProjectID, testOptions.GKEClusterLocation, log)
	teardown = func() {
		if cleanup != nil {
			log.Info("Beginning cluster cleanup...")
			if err := cleanup(); err != nil {
				t.Fatal(fmt.Errorf("error during cluster cleanup: %w", err))
			}
		}
	}
	if err != nil {
		teardown()
		t.Fatal(fmt.Errorf("error setting up cluster: %w", err))
	}
	if err := setupIdentity(testOptions, cluster.kubectl, log); err != nil {
		teardown()
		t.Fatal(fmt.Errorf("error setting up cluster: %w", err))
	}
	return testID, log, cluster, teardown
}

func cleanUpProject(deleteFunc cleanupFunc, shouldCleanUp bool, log logr.Logger) {
	if shouldCleanUp {
		log.Info("Beginning project cleanup...")
		if err := deleteFunc(); err != nil {
			log.Error(err, "error during project cleanup")
		}
	}
}

func getOperatorReleaseAssetsForVersion(version, serviceAccountID, projectID string, log logr.Logger) (manifestsDir string, sample configConnectorSample, err error) {
	log.Info("Downloading and extracting operator release tarball...", "version", version)
	emptySample := configConnectorSample{}
	releaseAssetsDir, err := createTempDir("e2e-operator-release-assets")
	if err != nil {
		return "", emptySample, fmt.Errorf("error creating temporary directory for operator release assets: %w", err)
	}
	if err := downloadAndExtractOperatorReleaseTarball(version, releaseAssetsDir); err != nil {
		return "", emptySample, fmt.Errorf("error downloading and extracting operator release tarball with version '%v': %w", version, err)
	}
	manifestsDir = path.Join(releaseAssetsDir, "operator-system")
	sample, err = getConfigConnectorSample(releaseAssetsDir, serviceAccountID, projectID)
	if err != nil {
		return "", emptySample, fmt.Errorf("error getting ConfigConnector sample from operator release assets: %w", err)
	}
	return manifestsDir, sample, nil
}

func (c *cluster) installOperator(operatorManifestsDir string) error {
	if _, err := c.kubectl.apply("-f", operatorManifestsDir); err != nil {
		return fmt.Errorf("error applying operator manifests: %w", err)
	}
	time.Sleep(30 * time.Second) // Wait for the operator's controllers and webhooks to come up and be registered
	return nil
}

func getConfigConnectorSample(operatorReleaseAssetsDir, serviceAccountID, projectID string) (sample configConnectorSample, err error) {
	emptySample := configConnectorSample{}
	samplesDir := path.Join(operatorReleaseAssetsDir, "samples")
	var yamlPaths []string
	sample = configConnectorSample{
		configConnectorClusterModeWorkloadIdentityYAMLPath: path.Join(samplesDir, "configconnector_cluster_mode_workload_identity.yaml"),
		configConnectorClusterModeGCPIdentityYAMLPath:      path.Join(samplesDir, "configconnector_cluster_mode_gcp_identity.yaml"),
		configConnectorNamespacedModeYAMLPath:              path.Join(samplesDir, "configconnector_namespaced_mode.yaml"),
		configConnectorContextYAMLPath:                     path.Join(samplesDir, "configconnectorcontext_sample.yaml"),
	}
	yamlPaths = []string{sample.configConnectorClusterModeWorkloadIdentityYAMLPath, sample.configConnectorNamespacedModeYAMLPath, sample.configConnectorContextYAMLPath}
	for _, yamlPath := range yamlPaths {
		content, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			return emptySample, fmt.Errorf("error reading YAML: %w", err)
		}
		s := string(content)

		// Replace vars (e.g. {GSA?})
		vars := map[string]string{
			"${GSA?}":        serviceAccountID,
			"${PROJECT_ID?}": projectID,
			"{GSA?}":         serviceAccountID,
			"{PROJECT_ID?}":  projectID,
		}
		processOrder := []string{"${GSA?}", "${PROJECT_ID?}", "{GSA?}", "{PROJECT_ID?}"}
		for _, k := range processOrder {
			v := vars[k]
			s = strings.ReplaceAll(s, k, v)
		}

		// Write back modified YAML to disk
		if err := writeToFile(s, yamlPath); err != nil {
			return emptySample, fmt.Errorf("error updating YAML file: %w", err)
		}
	}
	return sample, nil
}

func (c *cluster) installKCC(configConnectorYAMLPath string) error {
	content, err := ioutil.ReadFile(configConnectorYAMLPath)
	if err != nil {
		return fmt.Errorf("error reading ConfigConnector YAML: %w", err)
	}
	c.log.Info("Applying ConfigConnector YAML", "content", string(content))
	if _, err := c.kubectl.apply("-f", configConnectorYAMLPath); err != nil {
		return fmt.Errorf("error applying ConfigConnector YAML: %w", err)
	}
	if err := c.waitForConfigConnectorToBeHealthy(corev1beta1.ConfigConnectorAllowedName); err != nil {
		return err
	}

	// Wait for KCC's components to come up and be registered
	return c.waitForAllComponentPodsReady()
}

func (c *cluster) enableKCCForNamespace(namespace, configConnectorContextYAMLPath, serviceAccountID, projectID string) error {
	c.log.Info("Setting up Workload Identity binding for namespace...", "namespace", namespace)
	serviceAccEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountID, projectID)
	if err := setupWorkloadIdentityForNamespace(namespace, serviceAccEmail, projectID); err != nil {
		return fmt.Errorf("error setting up Workload Identity binding for namespace '%v': %w", namespace, err)
	}

	content, err := ioutil.ReadFile(configConnectorContextYAMLPath)
	if err != nil {
		return fmt.Errorf("error reading ConfigConnectorContext YAML for namespace '%v': %w", namespace, err)
	}
	c.log.Info("Applying ConfigConnectorContext YAML", "namespace", namespace, "content", string(content))
	if _, err := c.kubectl.apply("-n", namespace, "-f", configConnectorContextYAMLPath); err != nil {
		return fmt.Errorf("error applying ConfigConnectorContext YAML for namespace '%v': %w", namespace, err)
	}
	if err := c.waitForConfigConnectorToBeHealthy(corev1beta1.ConfigConnectorAllowedName); err != nil {
		return err
	}
	if err := c.waitForConfigConnectorContextToBeHealthy(namespace, corev1beta1.ConfigConnectorContextAllowedName); err != nil {
		return err
	}
	time.Sleep(90 * time.Second) // Wait for a KCC controller to come up and be registered for the given namespace
	return nil
}

func (c *cluster) waitForConfigConnectorToBeHealthy(name string) error {
	err := wait.PollImmediate(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		f := func() (interface{}, error) {
			return c.kubectl.get("configconnector", name, "-o", "yaml")
		}
		res, err := c.retry(f, longIntervalBackOff)
		if err != nil {
			return false, fmt.Errorf("error getting ConfigConnector '%v': %w", name, err)
		}
		c.log.Info("Waiting for ConfigConnector to reach a healthy state...", "name", name)
		return strings.Contains(res.(string), "healthy: true"), nil
	})
	if err != nil {
		if !errors.Is(err, wait.ErrWaitTimeout) {
			return err
		}
		out, _ := c.kubectl.get("configconnector", name, "-o", "yaml")
		return fmt.Errorf("timed out waiting for ConfigConnector '%v' to reach a healthy state:\n%v", name, out)
	}
	c.log.Info("ConfigConnector has reached a healthy state", "name", name)
	return nil
}

func (c *cluster) waitForAllComponentPodsReady() error {
	c.log.Info("waiting for all component pods in 'cnrm-system' namespace to be ready...")
	out, err := c.kubectl.exec("wait", "", "-n", "cnrm-system", "--for=condition=Ready", "pod", "--all", "--timeout=180s")
	if err != nil {
		return fmt.Errorf("error waiting for all component pods in 'cnrm-system' namespace to be ready: %w, output: %v", err, out)
	}
	return nil
}

func (c *cluster) waitForConfigConnectorContextToBeRemoved(namespace, name string) error {
	err := wait.PollImmediate(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		var isDeleted bool
		f := func() (interface{}, error) {
			c.log.Info("Getting ConfigConnectorContext...", "namespace", namespace, "name", name)
			out, err := c.kubectl.get("-n", namespace, "configconnectorcontext", name)
			if err != nil {
				// Quick exit if the ConfigConnectorContext object is deleted.
				if strings.Contains(err.Error(), "Error from server (NotFound)") {
					isDeleted = true
					return nil, nil
				}
				return nil, err
			}
			return out, nil
		}
		out, err := c.retry(f, defaultBackOff)
		if err != nil {
			return false, fmt.Errorf("unexpected error getting ConfigConnectorContext: %w; command output: %v", err, out)
		}
		if isDeleted {
			return true, nil
		}
		c.log.Info("Waiting for ConfigConnectorContext to be deleted...", "namespace", namespace, "name", name, "output", out)
		return false, nil
	})
	if err == nil {
		return nil
	}
	if !errors.Is(err, wait.ErrWaitTimeout) {
		return err
	}
	out, _ := c.kubectl.get("-n", namespace, "configconnectorcontext", name, "-o", "yaml")
	return fmt.Errorf("timed out waiting for ConfigConnectorContext '%v/%v' to be removed:\n%v", namespace, name, out)
}

func (c *cluster) waitForConfigConnectorContextToBeHealthy(namespace, name string) error {
	return c.waitForConfigConnectorContextToBeHealthyOrUnhealthy(namespace, name, true)
}

func (c *cluster) waitForConfigConnectorContextToBeUnhealthy(namespace, name string) error {
	return c.waitForConfigConnectorContextToBeHealthyOrUnhealthy(namespace, name, false)
}

func (c *cluster) waitForConfigConnectorContextToBeHealthyOrUnhealthy(namespace, name string, healthy bool) error {
	desiredState := "unhealthy"
	if healthy {
		desiredState = "healthy"
	}
	err := wait.PollImmediate(5*time.Second, 10*time.Minute, func() (done bool, err error) {
		f := func() (interface{}, error) {
			return c.kubectl.get("-n", namespace, "configconnectorcontext", name, "-o", "yaml")
		}
		out, err := c.retry(f, longIntervalBackOff)
		if err != nil {
			return false, fmt.Errorf("error getting ConfigConnectorContext '%v' for namespace '%v': %w", name, namespace, err)
		}
		c.log.Info(fmt.Sprintf("Waiting for ConfigConnectorContext to reach an %v state...", desiredState), "namespace", namespace, "name", name)
		return strings.Contains(out.(string), fmt.Sprintf("healthy: %v", healthy)), nil
	})
	if err != nil {
		if !errors.Is(err, wait.ErrWaitTimeout) {
			return err
		}
		out, _ := c.kubectl.get("-n", namespace, "configconnectorcontext", name, "-o", "yaml")
		return fmt.Errorf("timed out waiting for ConfigConnectorContext '%v/%v' to reach an %v state:\n%v", namespace, name, desiredState, out)
	}
	c.log.Info(fmt.Sprintf("ConfigConnectorContext has reached an %v state", desiredState), "namespace", namespace, "name", name)
	return nil
}

func setupWorkloadIdentityForNamespace(namespace, serviceAccEmail, projectID string) error {
	member := fmt.Sprintf("serviceAccount:%v.svc.id.goog[cnrm-system/%v%v]", projectID, k8s.ServiceAccountNamePrefix, namespace)
	role := "roles/iam.workloadIdentityUser"
	if err := addIAMBindingForServiceAcc(serviceAccEmail, member, role, projectID); err != nil {
		return fmt.Errorf("error setting up Workload Identity binding: %w", err)
	}
	return nil
}

func getArtifactRegistryRepositorySample(kccReleaseAssetsDir, uniqueID string, projectID string) (repoName string, repoYAMLDir string, err error) {
	repoYAMLDir = path.Join(kccReleaseAssetsDir, "samples", "resources", "artifactregistryrepository")
	yamlPaths, err := getYAMLFilesInDir(repoYAMLDir)
	if err != nil {
		return "", "", fmt.Errorf("error getting paths to YAML files in ArtifactRegistryRepository sample directory '%v': %w", repoYAMLDir, err)
	}
	for _, yamlPath := range yamlPaths {
		b, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			return "", "", fmt.Errorf("error reading file '%v': %w", yamlPath, err)
		}
		s := string(b)
		s = strings.ReplaceAll(s, "sample", "sample"+uniqueID)
		s = strings.ReplaceAll(s, "dep", "dep"+uniqueID)

		s = strings.ReplaceAll(s, "${PROJECT_ID?}", projectID)

		// Write back modified YAML to disk
		if err := writeToFile(s, yamlPath); err != nil {
			return "", "", fmt.Errorf("error updating file '%v': %w", yamlPath, err)
		}
	}
	repoName, err = getRepoNameFromArtifactRegistryRepositorySampleDir(repoYAMLDir)
	if err != nil {
		return "", "", fmt.Errorf("error getting name of ArtifactRegistryRepository for ArtifactRegistryRepository sample directory '%v': %w", repoYAMLDir, err)
	}
	return repoName, repoYAMLDir, nil
}

func getRepoNameFromArtifactRegistryRepositorySampleDir(repoYAMLDir string) (string, error) {
	unstructs := make([]*unstructured.Unstructured, 0)
	yamlPaths, err := getYAMLFilesInDir(repoYAMLDir)
	if err != nil {
		return "", fmt.Errorf("error getting paths to YAML files in directory '%v': %w", repoYAMLDir, err)
	}
	for _, yamlPath := range yamlPaths {
		u, err := utils.ReadFileToUnstructs(yamlPath)
		if err != nil {
			return "", fmt.Errorf("error converting file '%v' to unstructs: %w", yamlPath, err)
		}
		unstructs = append(unstructs, u...)
	}
	repoNames := make([]string, 0)
	for _, u := range unstructs {
		if u.GetKind() == "ArtifactRegistryRepository" {
			repoNames = append(repoNames, u.GetName())
		}
	}
	switch len(repoNames) {
	case 0:
		return "", fmt.Errorf("no ArtifactRegistryRepository found in directory '%v'", repoYAMLDir)
	case 1:
		return repoNames[0], nil
	default:
		return "", fmt.Errorf("multiple ArtifactRegistryRepositories found in directory '%v'", repoYAMLDir)
	}
}

func getYAMLFilesInDir(dir string) (yamlPaths []string, err error) {
	yamlPaths = make([]string, 0)
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, fmt.Errorf("error reading directory '%v': %w", dir, err)
	}
	for _, fi := range fileInfos {
		if fi.IsDir() {
			continue
		}
		if !strings.HasSuffix(fi.Name(), ".yaml") {
			continue
		}
		yamlPaths = append(yamlPaths, path.Join(dir, fi.Name()))
	}
	return yamlPaths, nil
}

func (c *cluster) createArtifactRegistryRepository(namespace, repoName, repoYAMLDir string) error {
	if err := c.createArtifactRegistryRepositoryAndWait(namespace, repoName, repoYAMLDir); err != nil {
		if errors.Is(err, wait.ErrWaitTimeout) {
			out, _ := c.kubectl.get("-n", namespace, "artifactregistryrepository", repoName, "-o", "yaml")
			return fmt.Errorf("timed out waiting for ArtifactRegistryRepository to reach an UpToDate state:\n%v", out)
		}
		return err
	}
	return nil
}

func (c *cluster) createArtifactRegistryRepositoryShouldFail(namespace, repoName, repoYAMLDir string) error {
	if err := c.createArtifactRegistryRepositoryAndWait(namespace, repoName, repoYAMLDir); err != nil {
		if errors.Is(err, wait.ErrWaitTimeout) {
			return nil // i.e. ArtifactRegistryRepository never reached an "UpToDate" state as expected
		}
		return err
	}
	// ArtifactRegistryRepository ended up being created successfully contrary to expectations, so return an error
	out, _ := c.kubectl.get("-n", namespace, "artifactregistryrepository", repoName, "-o", "yaml")
	return fmt.Errorf("expected creation of ArtifactRegistryRepository to fail, but got:\n%v", out)
}

func (c *cluster) createArtifactRegistryRepositoryAndWait(namespace, repoName, repoYAMLDir string) error {
	if _, err := c.kubectl.apply("-n", namespace, "-f", repoYAMLDir); err != nil {
		return fmt.Errorf("error applying ArtifactRegistryRepository: %w", err)
	}
	return wait.PollImmediate(5*time.Second, 2*time.Minute, func() (done bool, err error) {
		c.log.Info("Getting ArtifactRegistryRepository...", "name", repoName)
		f := func() (interface{}, error) {
			return c.kubectl.get("-n", namespace, "artifactregistryrepository", repoName, "-o", "yaml")
		}
		// Sometime, polling on object returns some transient-ish connection errors;
		// here we want to be more tolerant/robust by retrying a little more with a longer interval.
		out, err := c.retry(f, longIntervalBackOff)
		if err != nil {
			return false, fmt.Errorf("error getting ArtifactRegistryRepository '%v/%v': %w", namespace, repoName, err)
		}
		c.log.Info("Waiting for ArtifactRegistryRepository to reach an UpToDate state...", "name", repoName)
		return strings.Contains(out.(string), "UpToDate"), nil
	})
}

func (c *cluster) waitForCNRMFinalizersToBeRemovedFromArtifactRegistryRepository(namespace, repoName string) error {
	waitFunc := func() (done bool, err error) {
		ok, err := c.doesArtifactRegistryRepositoryHaveFinalizer(namespace, repoName, k8s.KCCFinalizer)
		if err != nil {
			return false, fmt.Errorf("error checking for the finalizer on ArtifactRegistryRepository: %w", err)
		}
		return !ok, nil
	}
	if err := wait.PollImmediate(5*time.Second, 5*time.Minute, waitFunc); err != nil {
		if !errors.Is(err, wait.ErrWaitTimeout) {
			return err
		}
		out, _ := c.kubectl.get("-n", namespace, "artifactregistryrepository", repoName, "-o", "yaml")
		return fmt.Errorf("timed out waiting for the CNRM finalizers to be removed from ArtifactRegistryRepository '%v/%v':\n%v", namespace, repoName, out)
	}
	return nil
}

func (c *cluster) getArtifactRegistryRepositoryUnstructured(namespace, repoName string) (*unstructured.Unstructured, error) {
	out, err := c.kubectl.get("-n", namespace, "artifactregistryrepository", repoName, "-o", "yaml")
	if err != nil {
		return nil, fmt.Errorf("error getting ArtifactRegistryRepository '%v/%v': %w", namespace, repoName, err)
	}
	repoUnstruct, err := utils.BytesToUnstruct([]byte(out))
	if err != nil {
		return nil, fmt.Errorf("error converting '%v' to unstruct: %w", out, err)
	}
	return repoUnstruct, err
}

func (c *cluster) applyUnstructured(u *unstructured.Unstructured) error {
	bytes, err := utils.UnstructToYaml(u)
	if err != nil {
		return fmt.Errorf("error converting unstruct to yaml: %w", err)
	}
	if _, err := c.kubectl.applyStdin(string(bytes), "-f", "-"); err != nil {
		return fmt.Errorf("error applying %v after adding finalizer: %w", u.GetKind(), err)
	}
	return nil
}

func (c *cluster) removeFinalizerToArtifactRegistryRepository(namespace, repoName, finalizer string) error {
	repoUnstruct, err := c.getArtifactRegistryRepositoryUnstructured(namespace, repoName)
	if err != nil {
		return err
	}
	finalizers := append(repoUnstruct.GetFinalizers(), finalizer)
	var newFinalizers []string
	for _, f := range finalizers {
		if f == finalizer {
			continue
		}
		newFinalizers = append(newFinalizers, f)
	}
	repoUnstruct.SetFinalizers(newFinalizers)
	return c.applyUnstructured(repoUnstruct)
}

func (c *cluster) addFinalizerToArtifactRegistryRepository(namespace, repoName, finalizer string) error {
	repoUnstruct, err := c.getArtifactRegistryRepositoryUnstructured(namespace, repoName)
	if err != nil {
		return err
	}
	finalizers := append(repoUnstruct.GetFinalizers(), finalizer)
	repoUnstruct.SetFinalizers(finalizers)
	return c.applyUnstructured(repoUnstruct)
}

func (c *cluster) doesArtifactRegistryRepositoryHaveFinalizer(namespace, repoName, finalizer string) (ok bool, err error) {
	repoUnstruct, err := c.getArtifactRegistryRepositoryUnstructured(namespace, repoName)
	if err != nil {
		return false, err
	}
	for _, f := range repoUnstruct.GetFinalizers() {
		if finalizer == f {
			return true, nil
		}
	}
	return false, nil
}

func (c *cluster) doesArtifactRegistryRepositoryHaveStatusUnmanaged(namespace, repoName string) (ok bool, err error) {
	repoUnstruct, err := c.getArtifactRegistryRepositoryUnstructured(namespace, repoName)
	if err != nil {
		return false, err
	}
	r, err := kcck8s.NewResource(repoUnstruct)
	if err != nil {
		return false, err
	}
	condition, ok := kcck8s.GetReadyCondition(r)
	if !ok {
		return false, nil
	}
	return condition.Reason == kcck8s.Unmanaged, nil
}

func (c *cluster) deleteArtifactRegistryRepository(namespace, repoName string, extraArgs ...string) error {
	args := []string{"-n", namespace, "artifactregistryrepository", repoName}
	args = append(args, extraArgs...)
	f := func() (interface{}, error) {
		return c.kubectl.delete(args...)
	}
	_, err := c.retry(f, defaultBackOff)
	if err != nil {
		return fmt.Errorf("error deleting ArtifactRegistryRepository: %w", err)
	}
	return nil
}

func (c *cluster) deleteConfigConnectorContext(namespace, name string) error {
	args := []string{"-n", namespace, "configconnectorcontext", name}
	f := func() (interface{}, error) {
		if _, err := c.kubectl.delete(args...); err != nil {
			c.log.Info("error deleting ConfigConnectorContext...", "error", err)
			return nil, err
		}
		return nil, nil
	}
	if _, err := c.retry(f, longIntervalBackOff); err != nil {
		return fmt.Errorf("error deleting ConfigConnectorContext: %w", err)
	}
	return c.waitForConfigConnectorContextToBeRemoved(namespace, name)
}

func (c *cluster) uninstallKCC() error {
	c.log.Info("deleting the ConfigConnector object")
	f := func() (interface{}, error) {
		if _, err := c.kubectl.delete("configconnector", corev1beta1.ConfigConnectorAllowedName); err != nil {
			c.log.Info("error deleting ConfigConnector...", "error", err)
			return nil, err
		}
		return nil, nil
	}
	if _, err := c.retry(f, longIntervalBackOff); err != nil {
		return fmt.Errorf("error deleting ConfigConnectors: %w", err)
	}
	c.log.Info("Asserting that the ConfigConnector object is gone")
	out, err := c.kubectl.get("configconnector")
	if err != nil {
		return fmt.Errorf("error getting ConfigConnectors: %w", err)
	}
	if !strings.Contains(out, "No resources found") {
		return fmt.Errorf("expected no ConfigConnectors to exist, but got:\n%v", out)
	}

	// As the uninstallation is no longer blocked by the deletion of the ignored
	// CRDs, the following assertion might fail in the beginning. But the ignored
	// CRDs have the ownerReferences to the ConfigConnector object, so after the
	// ConfigConnector object is deleted, the ignored CRDs will eventually be garbage
	// collected.
	// Retrying the assertion to simulate the latest UX.
	crdAssertionFunc := func() (interface{}, error) {
		c.log.Info("Asserting that resource CRDs are deleted")
		out, err = c.kubectl.get("crds", "--selector", "cnrm.cloud.google.com/managed-by-kcc=true")
		if err != nil {
			return nil, fmt.Errorf("error getting KCC CRDs: %w", err)
		}
		if !strings.Contains(out, "No resources found") {
			return nil, fmt.Errorf("expected KCC CRDs to not exist, but got:\n%v", out)
		}
		return out, nil
	}
	_, err = c.retry(crdAssertionFunc, defaultBackOff)
	if err != nil {
		return fmt.Errorf("unexpected error asserting that resource CRDs are deleted: %w", err)
	}

	out, err = c.kubectl.get("validatingwebhookconfiguration")
	if err != nil {
		return fmt.Errorf("error getting ValidatingWebhookConfigurations: %w", err)
	}
	if strings.Contains(out, k8s.CNRMDomain) {
		return fmt.Errorf("expected KCC validating webhooks to not exist, but got:\n%v", out)
	}
	out, err = c.kubectl.get("mutatingwebhookconfiguration")
	if err != nil {
		return fmt.Errorf("error getting MutatingWebhookConfigurations: %w", err)
	}
	if strings.Contains(out, k8s.CNRMDomain) {
		return fmt.Errorf("expected KCC mutating webhooks to not exist, but got:\n%v", out)
	}
	c.log.Info("Asserting that `cnrm-system` namespace is deleted")
	return c.waitForNamespaceToBeDeleted(k8s.CNRMSystemNamespace)
}

func (c *cluster) waitForNamespaceToBeDeleted(namespace string) error {
	// Deleting a namespace can take a long time. Give some buffer time for k8s api server to process.
	time.Sleep(5 * time.Minute)

	err := wait.PollImmediate(20*time.Second, 10*time.Minute, func() (done bool, err error) {
		c.log.Info("Getting namespace...", "namespace", namespace)
		var isDeleted bool
		f := func() (interface{}, error) {
			ns, err := c.getNamespace(namespace)
			if err != nil {
				// Quick exit if the namespace is deleted already.
				if apierrors.IsNotFound(err) {
					isDeleted = true
					return nil, nil
				}
				return nil, err
			}
			return ns, nil
		}
		// Sometime, polling on object returns some transient-ish connection errors;
		// here we want to be more tolerant/robust by retrying a little more with a longer interval.
		res, err := c.retry(f, longIntervalBackOff)
		if err != nil {
			return false, fmt.Errorf("error getting namespace '%v': %w", namespace, err)
		}
		if isDeleted {
			return true, nil
		}
		ns := res.(*v1.Namespace)
		c.log.Info("Waiting for namespace to be deleted...", "namespace", namespace, "status", ns.Status)
		return false, nil
	})
	if err != nil {
		if !errors.Is(err, wait.ErrWaitTimeout) {
			return err
		}
		return fmt.Errorf("timed out waiting for namespace '%v' to be deleted", namespace)
	}
	return nil
}

func checkArtifactRegistryRepositoryExistsOnGCP(repoName, projectID string) error {
	cmd := exec.Command("gcloud", "artifacts", "repositories", "describe", repoName,
		"--location", "us-west1",
		"--project", projectID)
	_, err := utils.ExecuteAndCaptureOutput(cmd)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			return fmt.Errorf("expected project '%v' to have Artifact Registry Repository named '%v', but got:\n%w", projectID, repoName, err)
		}
		return fmt.Errorf("error checking if Artifact Registry Repository exists on GCP: %w", err)
	}
	return nil
}

func setupCluster(clusterName, projectID, location string, log logr.Logger) (*cluster, cleanupFunc, error) {
	var cleanup cleanupFunc
	log.Info("Creating a Container client...")
	ctx := context.Background()
	container, err := containerBeta.NewService(ctx)
	if err != nil {
		return nil, cleanup, fmt.Errorf("error creating Container client: %w", err)
	}
	log.Info("Creating a GKE cluster...", "name", clusterName)
	if err := createGKECluster(container, clusterName, projectID, location, log); err != nil {
		return nil, cleanup, fmt.Errorf("error creating GKE cluster with name '%v': %w", clusterName, err)
	}
	cleanup = func() error {
		log.Info("Deleting GKE cluster...", "name", clusterName)
		if err := deleteGKECluster(container, clusterName, projectID, location); err != nil {
			return fmt.Errorf("error deleting cluster with name '%v': %w", clusterName, err)
		}
		return nil
	}

	log.Info("Getting the cluster's kubeconfig...")
	outDirForKubeconfig, err := createTempDir("e2e-" + clusterName + "-kubeconfig")
	if err != nil {
		return nil, cleanup, fmt.Errorf("error creating temporary directory for cluster's kubeconfig: %w", err)
	}
	outPathForKubeconfig := path.Join(outDirForKubeconfig, "kubeconfig.yaml")
	if err := getKubeconfigForGKECluster(projectID, location, clusterName, outPathForKubeconfig); err != nil {
		return nil, cleanup, fmt.Errorf("error getting cluster's kubeconfig: %w", err)
	}
	log.Info("Setting up a client-go Clientset...")
	config, err := clientcmd.BuildConfigFromFlags("", outPathForKubeconfig)
	if err != nil {
		return nil, cleanup, fmt.Errorf("error building REST client config from kubeconfig: %w", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, cleanup, fmt.Errorf("error creating client-go Clientset: %w", err)
	}
	cluster := &cluster{
		kubectl: &kubectl{
			kubeconfigPath: outPathForKubeconfig,
			deleteTimeout:  KubectlDeleteTimeout,
		},
		clientset: clientset,
		log:       log,
	}
	return cluster, cleanup, nil
}

func setupIdentity(testOptions TestOptions, k *kubectl, log logr.Logger) error {
	if testOptions.SecretName != "" {
		log.Info("Creating a secret containing service account key..")
		serviceAccEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", testOptions.ServiceAccountID, testOptions.ProjectID)

		return createCredentialSecret(serviceAccEmail, testOptions.ProjectID, testOptions.SecretName, k)
	}

	log.Info("Setting up Workload Identity binding...")
	serviceAccEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", testOptions.ServiceAccountID, testOptions.ProjectID)
	member := fmt.Sprintf("serviceAccount:%v.svc.id.goog[cnrm-system/cnrm-controller-manager]", testOptions.ProjectID)
	role := "roles/iam.workloadIdentityUser"

	return addIAMBindingForServiceAcc(serviceAccEmail, member, role, testOptions.ProjectID)
}

func setupProject(organizationID, projectID, billingAccountID, serviceAccountID string, log logr.Logger) (cleanupFunc, error) {
	var cleanup cleanupFunc
	log.Info("Creating GCP clients...")
	ctx := context.Background()
	resourceManagerClient, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating ResourceManager client: %w", err)
	}
	billingClient, err := cloudbilling.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating Billing client: %w", err)
	}
	iamClient, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating IAM client: %w", err)
	}
	log.Info("Creating project...", "projectID", projectID)
	if err := createProject(resourceManagerClient, organizationID, projectID, log); err != nil {
		return cleanup, fmt.Errorf("error creating project with project ID '%v': %w", projectID, err)
	}
	cleanup = func() error {
		log.Info("Deleting project...", "projectID", projectID)
		if err := deleteProject(resourceManagerClient, projectID); err != nil {
			return fmt.Errorf("error deleting project with project ID '%v': %w", projectID, err)
		}
		return nil
	}
	log.Info("Linking project to billing account...", "billingAccount", billingAccountID)
	if err := linkProjectToBillingAccount(billingClient, projectID, billingAccountID); err != nil {
		return cleanup, fmt.Errorf("error linking project to billing account '%v'", billingAccountID)
	}
	log.Info("Enabling services for project...")
	if err := enableServicesForProject(projectID, services, log); err != nil {
		return cleanup, fmt.Errorf("error enabling services for project: %w", err)
	}
	log.Info("Setting up IAM service account...")
	if err := createServiceAccount(iamClient, serviceAccountID, projectID); err != nil {
		return cleanup, fmt.Errorf("error creating service account: %w", err)
	}
	serviceAccEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountID, projectID)
	if err := addIAMBindingForProject(projectID, "serviceAccount:"+serviceAccEmail, "roles/owner"); err != nil {
		return cleanup, fmt.Errorf("error granting service account project owner role: %w", err)
	}
	return cleanup, nil
}

func createProject(resourceManagerClient *cloudresourcemanager.Service, organizationID, projectID string, log logr.Logger) error {
	project := &cloudresourcemanager.Project{
		ProjectId: projectID,
		Labels: map[string]string{
			"cnrm-test": "true",
		},
		Parent: &cloudresourcemanager.ResourceId{
			Type: "organization",
			Id:   organizationID,
		},
	}
	op, err := resourceManagerClient.Projects.Create(project).Do()
	if err != nil {
		return err
	}
	// Wait for project creation operation to finish
	return wait.PollImmediate(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		res, err := resourceManagerClient.Operations.Get(op.Name).Do()
		if err != nil {
			return false, err
		}
		if res.Done {
			if res.Error != nil {
				return true, fmt.Errorf("project creation operation failed: %v", res.Error)
			}
			return true, nil
		}
		log.Info("Waiting for project creation operation to finish...")
		return false, nil
	})
}

func deleteProject(resourceManagerClient *cloudresourcemanager.Service, projectID string) error {
	_, err := resourceManagerClient.Projects.Delete(projectID).Do()
	if err != nil {
		return err
	}
	return nil
}

func linkProjectToBillingAccount(billingClient *cloudbilling.APIService, projectID, billingAccount string) error {
	ba := &cloudbilling.ProjectBillingInfo{
		BillingAccountName: "billingAccounts/" + billingAccount,
	}
	_, err := billingClient.Projects.UpdateBillingInfo("projects/"+projectID, ba).Do()
	return err
}

func enableServicesForProject(projectID string, services []string, log logr.Logger) error {
	for _, service := range services {
		log.Info("Enabling service...", "service", service)
		cmd := exec.Command("gcloud", "services", "enable", service, "--project", projectID)
		if err := utils.Execute(cmd); err != nil {
			return err
		}
	}
	return nil
}

func createServiceAccount(iamClient *iam.Service, serviceAccountID, projectID string) error {
	req := &iam.CreateServiceAccountRequest{
		AccountId: serviceAccountID,
	}
	_, err := iamClient.Projects.ServiceAccounts.Create("projects/"+projectID, req).Do()
	if err != nil {
		return err
	}
	return nil
}

func addIAMBindingForProject(projectID, member, role string) error {
	cmd := exec.Command(
		"gcloud", "projects", "add-iam-policy-binding", projectID,
		"--member", member,
		"--role", role,
	)
	return utils.Execute(cmd)
}

func createCredentialSecret(serviceAccEmail, projectID, secretName string, k *kubectl) error {
	cmd := exec.Command(
		"gcloud", "iam", "service-accounts", "keys", "create",
		"--iam-account", serviceAccEmail,
		"--project", projectID,
		"./key.json",
	)
	if err := utils.Execute(cmd); err != nil {
		return fmt.Errorf("error creating a service account key: %w", err)
	}
	if _, err := k.exec("create", "", "ns", "cnrm-system"); err != nil {
		return fmt.Errorf("error creating cnrm-system namespace: %w", err)
	}
	if _, err := k.exec("create", "", "secret", "generic", secretName, "--from-file", "./key.json", "--namespace", "cnrm-system"); err != nil {
		return fmt.Errorf("error creating a secret containing service account key: %w", err)
	}
	rm := exec.Command("rm", "./key.json")
	if err := utils.Execute(rm); err != nil {
		return fmt.Errorf("error removing the service account key: %w", err)
	}
	return nil
}

func addIAMBindingForServiceAcc(serviceAccEmail, member, role, projectID string) error {
	addIAMBindingFunc := func() error {
		cmd := exec.Command(
			"gcloud", "iam", "service-accounts", "add-iam-policy-binding", serviceAccEmail,
			"--member", member,
			"--role", role,
			"--project", projectID,
		)
		return utils.Execute(cmd)
	}
	return backoff.Retry(addIAMBindingFunc, backoff.NewExponentialBackOff())
}

func createGKECluster(containerClient *containerBeta.Service, clusterName, projectID, location string, log logr.Logger) error {
	cluster := &containerBeta.Cluster{
		Name: clusterName,
		WorkloadIdentityConfig: &containerBeta.WorkloadIdentityConfig{
			IdentityNamespace: projectID + ".svc.id.goog",
		},
		InitialNodeCount: 6,
	}
	req := &containerBeta.CreateClusterRequest{
		Cluster: cluster,
	}
	parent := fmt.Sprintf("projects/%s/locations/%s", projectID, location)
	op, err := containerClient.Projects.Locations.Clusters.Create(parent, req).Do()
	if err != nil {
		return err
	}
	// Wait for cluster creation operation to finish
	err = wait.PollImmediate(10*time.Second, 10*time.Minute, func() (done bool, err error) {
		name := containerOpFullName(projectID, location, op.Name)
		res, err := containerClient.Projects.Locations.Operations.Get(name).Do()
		if err != nil {
			return false, err
		}
		if res.Status == "DONE" {
			if res.StatusMessage != "" {
				return true, fmt.Errorf("cluster creation operation failed: %v", res.StatusMessage)
			}
			return true, nil
		}
		log.Info("Waiting for cluster creation operation to finish...")
		return false, nil
	})
	if err != nil {
		return err
	}
	// Wait for cluster to be in a RUNNING state
	err = wait.PollImmediate(5*time.Second, 5*time.Minute, func() (done bool, err error) {
		name := clusterFullName(projectID, location, clusterName)
		cluster, err := containerClient.Projects.Locations.Clusters.Get(name).Do()
		if err != nil {
			return false, err
		}
		if cluster.Status == "RUNNING" {
			return true, nil
		}
		log.Info("Waiting for cluster to be in RUNNING state...", "currentState", cluster.Status)
		return false, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func deleteGKECluster(containerClient *containerBeta.Service, clusterName, projectID, location string) error {
	name := clusterFullName(projectID, location, clusterName)
	_, err := containerClient.Projects.Locations.Clusters.Delete(name).Do()
	if err != nil {
		return err
	}
	return nil
}

func containerOpFullName(projectID, location, opName string) string {
	return fmt.Sprintf("projects/%s/locations/%s/operations/%s", projectID, location, opName)
}

func clusterFullName(projectID, location, clusterName string) string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectID, location, clusterName)
}

func getKubeconfigForGKECluster(projectID, location, clusterName, outputFilePath string) error {
	cmd := exec.Command(
		"gcloud", "container", "clusters", "get-credentials", clusterName,
		"--zone", location,
		"--project", projectID,
	)
	// Override the file onto which the retrieved GKE credentials are to be
	// written in. Note that this overrides the KUBECONFIG env var even if it
	// has already been set.
	envVarOverride := "KUBECONFIG=" + outputFilePath
	cmd.Env = append(os.Environ(), envVarOverride)
	return utils.Execute(cmd)
}

func (k *kubectl) applyStdin(stdin string, args ...string) (string, error) {
	return k.exec("apply", stdin, args...)
}

func (k *kubectl) apply(args ...string) (stdout string, err error) {
	return k.applyStdin("", args...)
}

func (k *kubectl) delete(args ...string) (stdout string, err error) {
	timeout := fmt.Sprintf("%vs", k.deleteTimeout.Seconds())
	args = append(args, "--timeout", timeout)
	stdout, err = k.exec("delete", "", args...)
	if err != nil && strings.Contains(err.Error(), "Error from server (NotFound)") {
		// The resource is already gone.
		return "", nil
	}
	return stdout, err
}

func (k *kubectl) get(args ...string) (stdout string, err error) {
	return k.exec("get", "", args...)
}

func (k *kubectl) exec(command, stdin string, args ...string) (stdout string, err error) {
	if k.kubeconfigPath == "" {
		return "", fmt.Errorf("attempted to execute a kubectl command without a specified kubeconfig")
	}
	args = append([]string{command}, args...)
	args = append(args, "--kubeconfig", k.kubeconfigPath)
	cmd := exec.Command("kubectl", args...)
	if stdin != "" {
		cmd.Stdin = bytes.NewBufferString(stdin)
	}
	return utils.ExecuteAndCaptureOutput(cmd)
}

func downloadAndExtractOperatorReleaseTarball(version, outputDir string) error {
	tarballGCSPath := fmt.Sprintf("gs://%v/%v/%v", OperatorReleaseBucket, version, OperatorReleaseTarball)
	return utils.DownloadAndExtractTarballAt(tarballGCSPath, outputDir)
}

func downloadAndExtractKCCReleaseTarball(version, outputDir string) error {
	tarballGCSPath := fmt.Sprintf("gs://%v/%v/%v", KccReleaseBucket, version, KccReleaseTarball)
	return utils.DownloadAndExtractTarballAt(tarballGCSPath, outputDir)
}

func createTempDir(namePrefix string) (path string, err error) {
	// Creates a directory in /tmp whose name starts with
	// the given namePrefix.
	return ioutil.TempDir("", namePrefix)
}

func writeToFile(content string, filePath string) error {
	fileMode := os.FileMode(0644) // -rw-r--r--
	return ioutil.WriteFile(filePath, []byte(content), fileMode)
}

func (c *cluster) createNamespace(namespace string) error {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err := c.clientset.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	return err
}

func (c *cluster) deleteNamespace(namespace string) error {
	return c.clientset.CoreV1().Namespaces().Delete(context.Background(), namespace, metav1.DeleteOptions{})
}

func (c *cluster) addProjectIDAnnotationToNamespace(namespace, projectID string) error {
	getFunc := func() (interface{}, error) {
		return c.getNamespace(namespace)
	}
	res, err := c.retry(getFunc, defaultBackOff)
	if err != nil {
		return err
	}
	ns := res.(*v1.Namespace)
	annotations := getAnnotationsForNS(ns)
	annotations[k8s.ProjectIDAnnotation] = projectID
	ns.SetAnnotations(annotations)
	updateFunc := func() (interface{}, error) {
		ns, err = c.clientset.CoreV1().Namespaces().Update(context.Background(), ns, metav1.UpdateOptions{})
		if err != nil {
			return nil, fmt.Errorf("error updating namespace '%v': %w", namespace, err)
		}
		return ns, nil
	}
	_, err = c.retry(updateFunc, defaultBackOff)
	return err
}

// retry is a helper function to retry a function with the given backoff policy.
// It will return the original error from the input function if it still fails after retries.
//
// Note that we have frequently observed transient connection lost issues in operator e2e test,
// the mitigation is to retry on almost all the requests to the k8s API server. Use a backoff policy with
// a long interval if the observed transient connection issue seemingly takes a long time to recover.
func (c *cluster) retry(f func() (interface{}, error), backoff wait.Backoff) (interface{}, error) {
	var funcError error
	var res interface{}
	if err := wait.ExponentialBackoff(backoff, func() (bool, error) {
		res, funcError = f()
		if funcError != nil {
			c.log.Info("Retrying after encountering error", "error", funcError)
			return false, nil
		}
		return true, nil
	}); err != nil {
		return nil, funcError
	}
	return res, nil
}

func (c *cluster) getKCCVersion() (string, error) {
	ns, err := c.getNamespace(k8s.CNRMSystemNamespace)
	if err != nil {
		return "", err
	}
	annotations := getAnnotationsForNS(ns)
	version, ok := annotations[k8s.VersionAnnotation]
	if !ok {
		return "", fmt.Errorf("KCC version annotation ('%v') not found in namespace '%v'", k8s.VersionAnnotation, k8s.CNRMSystemNamespace)
	}
	return version, nil
}

func (c *cluster) getOperatorVersion() (string, error) {
	ns, err := c.getNamespace(k8s.OperatorSystemNamespace)
	if err != nil {
		return "", err
	}
	annotations := getAnnotationsForNS(ns)
	version, ok := annotations[k8s.OperatorVersionAnnotation]
	if !ok {
		return "", fmt.Errorf("KCC operator version annotation ('%v') not found in namespace '%v'", k8s.OperatorVersionAnnotation, k8s.OperatorSystemNamespace)
	}
	return version, nil
}

func (c *cluster) getNamespace(namespace string) (*v1.Namespace, error) {
	return c.clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
}

func getAnnotationsForNS(ns *v1.Namespace) map[string]string {
	annotations := ns.GetAnnotations()
	if annotations == nil {
		return make(map[string]string)
	}
	return annotations
}

func newUniqueTestID() string {
	return randomid.New().String()
}

func newLogger(name string) (logr.Logger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableCaller = true
	zapLog, err := zapConfig.Build()
	if err != nil {
		return logr.Discard(), err
	}
	log := zapr.NewLogger(zapLog)
	return log.WithName(name), nil
}
