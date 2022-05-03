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

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/scripts/utils"
)

const (
	dirMode        = os.FileMode(0755) //drwxr-x---
	fileMode       = os.FileMode(0644) // -rw-r--r--
	gcsPathTmpl    = "gs://cnrm/%v/release-bundle.tar.gz"
	baseDir        = "scripts/update-kcc-manifest"
	channelDir     = "channels/packages/configconnector"
	managerPatch   = "manager_sidecar_patch.yaml"
	recorderPatch  = "recorder_sidecar_patch.yaml"
	finalizerPatch = "finalizer_patch.yaml"
)

var (
	version string
)

// Download the latest KCC manifest, kustomize and upload it to the stable channel
// Usage: go run scripts/update-kcc-manifest/main.go --version latest
func main() {
	flag.StringVar(&version, "version", "latest", "Version of the KCC core to download.")
	flag.Parse()

	// download the KCC manifest
	operatorSrcRoot := paths.GetOperatorSrcRootOrLogFatal()
	outputDir := path.Join(operatorSrcRoot, baseDir, "kcc")
	if err := os.Mkdir(outputDir, dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir %v: %v", outputDir, err)
	}
	gcsPath := fmt.Sprintf(gcsPathTmpl, version)
	if err := utils.DownloadAndExtractTarballAt(gcsPath, outputDir); err != nil {
		log.Fatalf("error downloading and extracting the tarball %v: %v", gcsPath, err)
	}

	kustomizeBuild(operatorSrcRoot)

	// swap container registry
	wiSystemManifest := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-workload-identity/0-cnrm-system.yaml")
	gcpSystemManifest := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-gcp-identity/0-cnrm-system.yaml")
	namespacedSystemManifest := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-namespaced/0-cnrm-system.yaml")
	pnc := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-namespaced/per-namespace-components.yaml")
	manifests := []string{wiSystemManifest, gcpSystemManifest, namespacedSystemManifest, pnc}
	for _, manifest := range manifests {
		if err := swapContainerRegistry(manifest); err != nil {
			log.Fatalf("error swapping container registry: %v", err)
		}
	}

	// upload the new manifest under channels dir
	manifestFile := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-namespaced/0-cnrm-system.yaml")
	version, err := extractVersionFromManifest(manifestFile)
	if err != nil {
		log.Fatalf("error extracting version from manifest %v: %v", manifestFile, err)
	}
	manifestDir := path.Join(operatorSrcRoot, channelDir, version)
	if err := os.Mkdir(manifestDir, dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir %v: %v", manifestDir, err)
	}

	// copy crds.yaml
	crds := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-namespaced/crds.yaml")
	destCrds := path.Join(manifestDir, "crds.yaml")
	if err := utils.Copy(crds, destCrds); err != nil {
		log.Fatalf("error copying %v to %v: %v", crds, destCrds, err)
	}
	// copy install-bundle-workload-identity/0-cnrm-system.yaml
	if err := os.Mkdir(path.Join(manifestDir, "cluster"), dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir: %v", err)
	}
	if err := os.Mkdir(path.Join(manifestDir, "cluster", "workload-identity"), dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir: %v", err)
	}
	destWiSystemManifest := path.Join(manifestDir, "cluster", "workload-identity", "0-cnrm-system.yaml")
	if err := utils.Copy(wiSystemManifest, destWiSystemManifest); err != nil {
		log.Fatalf("error copying %v to %v: %v", wiSystemManifest, destWiSystemManifest, err)
	}
	// copy install-bundle-gcp-identity/0-cnrm-system.yaml
	if err := os.Mkdir(path.Join(manifestDir, "cluster", "gcp-identity"), dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir: %v", err)
	}
	destGcpSystemManifest := path.Join(manifestDir, "cluster", "gcp-identity", "0-cnrm-system.yaml")
	if err := utils.Copy(gcpSystemManifest, destGcpSystemManifest); err != nil {
		log.Fatalf("error copying %v to %v: %v", wiSystemManifest, destWiSystemManifest, err)
	}
	// copy install-bundle-namespaced/0-cnrm-system.yaml and install-bundle-namespaced/per-namespace-components.yaml
	if err := os.Mkdir(path.Join(manifestDir, "namespaced"), dirMode); err != nil && !os.IsExist(err) {
		log.Fatalf("error creating dir %v: %v", outputDir, err)
	}
	destNamespacedSystemManifest := path.Join(manifestDir, "namespaced", "0-cnrm-system.yaml")
	if err := utils.Copy(namespacedSystemManifest, destNamespacedSystemManifest); err != nil {
		log.Fatalf("error copying %v to %v: %v", namespacedSystemManifest, destNamespacedSystemManifest, err)
	}
	destPnc := path.Join(manifestDir, "namespaced", "per-namespace-components.yaml")
	if err := utils.Copy(pnc, destPnc); err != nil {
		log.Fatalf("error copying %v to %v: %v", pnc, destPnc, err)
	}
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatalf("error deleting dir %v: %v", outputDir, err)
	}

	// update the operator version
	kustomizationFilePath := path.Join(operatorSrcRoot, "config", "default", "kustomization.yaml")
	b, err := ioutil.ReadFile(kustomizationFilePath)
	if err != nil {
		log.Fatalf("error reading %v: %v", kustomizationFilePath, err)
	}
	kustomization := string(b)
	m := regexp.MustCompile("cnrm.cloud.google.com/operator-version: (\".*\")")
	kustomization = m.ReplaceAllString(kustomization, fmt.Sprintf("cnrm.cloud.google.com/operator-version: \"%v\"", version))
	if err := ioutil.WriteFile(kustomizationFilePath, []byte(kustomization), fileMode); err != nil {
		log.Fatalf("error updating file %v", kustomizationFilePath)
	}
	log.Printf("successfully updated the version annotation in %v\n", kustomizationFilePath)

	//remove the stale manifest
	r := loaders.NewFSRepository(path.Join(operatorSrcRoot, loaders.FlagChannel))
	channel, err := r.LoadChannel(context.TODO(), k8s.StableChannel)
	if err != nil {
		log.Fatalf("error loading %v channel: %v", k8s.StableChannel, err)
	}
	currentVersion, err := channel.Latest("configconnector")
	if err != nil {
		log.Fatalf("error resolving the current version: %v", err)
	}
	if currentVersion.Version == version {
		log.Printf("the current KCC version is the same as the latest version %v\n", version)
		return
	}
	stableFilePath := path.Join(operatorSrcRoot, "channels", "stable")
	b, err = ioutil.ReadFile(stableFilePath)
	if err != nil {
		log.Fatalf("error reading %v: %v", stableFilePath, err)
	}
	stable := string(b)
	stable = strings.ReplaceAll(stable, fmt.Sprintf("- version: %v", currentVersion.Version), fmt.Sprintf("- version: %v", version))
	if err := ioutil.WriteFile(stableFilePath, []byte(stable), fileMode); err != nil {
		log.Fatalf("error updating file %v", stableFilePath)
	}
	staleManifestDir := path.Join(operatorSrcRoot, "channels", "packages", "configconnector", currentVersion.Version)
	log.Printf("removing stale manifest %v", staleManifestDir)
	if err := os.RemoveAll(staleManifestDir); err != nil {
		log.Fatalf("error deleting dir %v: %v", staleManifestDir, err)
	}
}

func kustomizeBuild(operatorSrcRoot string) {
	// workload-identity cluster mode
	buildPath := path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-workload-identity")
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, "kustomizations", "kustomization_workload-identity.yaml"), path.Join(buildPath, "kustomization.yaml")); err != nil {
		log.Fatalf("error copying kustomization: %v", err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, managerPatch), path.Join(buildPath, managerPatch)); err != nil {
		log.Fatalf("error copying %v: %v", managerPatch, err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, recorderPatch), path.Join(buildPath, recorderPatch)); err != nil {
		log.Fatalf("error copying %v: %v", recorderPatch, err)
	}
	output := path.Join(buildPath, "0-cnrm-system.yaml")
	if err := utils.KustomizeBuild(buildPath, output); err != nil {
		log.Fatalf("error running kustomize build: %v", err)
	}

	// namespaced mode
	buildPath = path.Join(operatorSrcRoot, baseDir, "kcc", "install-bundle-namespaced")
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, managerPatch), path.Join(buildPath, managerPatch)); err != nil {
		log.Fatalf("error copying %v: %v", managerPatch, err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, recorderPatch), path.Join(buildPath, recorderPatch)); err != nil {
		log.Fatalf("error copying %v: %v", recorderPatch, err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, finalizerPatch), path.Join(buildPath, finalizerPatch)); err != nil {
		log.Fatalf("error copying %v: %v", finalizerPatch, err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, "kustomizations", "kustomization_namespaced_0-cnrm-system.yaml"), path.Join(buildPath, "kustomization.yaml")); err != nil {
		log.Fatalf("error copying kustomization: %v", err)
	}
	output = path.Join(buildPath, "0-cnrm-system.yaml")
	if err := utils.KustomizeBuild(buildPath, output); err != nil {
		log.Fatalf("error running kustomize build: %v", err)
	}
	if err := utils.Copy(path.Join(operatorSrcRoot, baseDir, "kustomizations", "kustomization_namespaced_per-namespace-components.yaml"), path.Join(buildPath, "kustomization.yaml")); err != nil {
		log.Fatalf("error copying kustomization: %v", err)
	}
	output = path.Join(buildPath, "per-namespace-components.yaml")
	if err := utils.KustomizeBuild(buildPath, output); err != nil {
		log.Fatalf("error running kustomize build: %v", err)
	}
}

// This step can be removed once we switch KCC core to also use gcr.io/gke-release container registry
func swapContainerRegistry(manifestPath string) error {
	content, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("error reading manifestPath: %v", err)
	}
	manifest := string(content)
	updatedManifest := strings.ReplaceAll(manifest, "gcr.io/cnrm-eap/", "gcr.io/gke-release/cnrm/")
	fileMode := os.FileMode(0644) // -rw-r--r--
	return ioutil.WriteFile(manifestPath, []byte(updatedManifest), fileMode)
}

func extractVersionFromManifest(filePath string) (string, error) {
	objs, err := utils.ReadFileToUnstructs(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %v and converting to unstructs: %v", filePath, err)
	}
	for _, obj := range objs {
		if obj.GetKind() == "Namespace" && obj.GetName() == k8s.CNRMSystemNamespace {
			for key, val := range obj.GetAnnotations() {
				if key == k8s.VersionAnnotation {
					return val, nil
				}
			}
		}
	}
	return "", fmt.Errorf("couldn't extract the version from the manifest %v", filePath)
}
