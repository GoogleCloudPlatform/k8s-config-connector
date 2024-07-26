// Copyright 2024 Google LLC
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
	"archive/tar"
	"compress/gzip"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfigFlag *string

	targetNamespacesFlag *string // todo acpana repeatable
	ignoreNamespacesFlag *string // todo acpana repeatable

	targetObjectsFlag *string // todo acpana repeatable
	ignoreObjectsFlag *string // todo acpana repeatable
)

func main() {
	kubeconfigFlag = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "path to the kubeconfig file")

	targetNamespacesFlag = flag.String("target-namespaces", "", "comma-separated namespace prefixes to target the export tool. Targets all if empty.")
	ignoreNamespacesFlag = flag.String("exclude-namespaces", "kube", "comma-separated namespace prefixes to ignore. Excludes nothing if empty.")

	targetObjectsFlag = flag.String("target-objects", "", "comma-separated object name prefixes to target. Targets all if empty.")
	ignoreObjectsFlag = flag.String("exclude-objects", "", "comma-separated object name prefixes to ignore. Excludes nothing if empty.")
	flag.Parse()

	log.Printf("Running kompanion export with kubeconfig: %s", *kubeconfigFlag)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigFlag)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s\n", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes clientset: %s\n", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %s\n", err)
	}

	// use the discovery client to iterate over all api resoruces
	discoveryClient := clientset.Discovery()
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		log.Fatalf("Failed to get preferred resources: %s", err)
	}

	resources := []*metav1.APIResource{}
	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, "cnrm.") {
			log.Printf("ApiResource %s group doesn't contain \"cnrm\"; skipping", apiResourceList.GroupVersion)
		}
		for _, apiResource := range apiResourceList.APIResources {
			if !apiResource.Namespaced {
				log.Printf("ApiResource %s is not namespaced; skipping", apiResource.SingularName)
				continue
			}
			if !contains(apiResource.Verbs, "list") {
				log.Printf("ApiResource %s is not listabble; skipping", apiResource.SingularName)
				continue
			}

			resources = append(resources, &apiResource)
		}
	}

	// incentiveze determinism for debuggability and idempotency
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].SingularName < resources[j].SingularName
	})
	// log.Printf("Going to iterate over the following resources %+v", resourcesToName(resources))

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error fetching namespaces: %s\n", err)
	}

	var namespacesToIgnore []string
	if *ignoreNamespacesFlag != "" {
		namespacesToIgnore = strings.Split(*ignoreNamespacesFlag, ",")
	}
	var namespacesToTarget []string
	if *targetNamespacesFlag != "" {
		namespacesToTarget = strings.Split(*targetNamespacesFlag, ",")
	}

	var objectsToIgnore []string
	if *ignoreObjectsFlag != "" {
		objectsToIgnore = strings.Split(*ignoreObjectsFlag, ",")
	}
	var objectsToTarget []string
	if *targetObjectsFlag != "" {
		objectsToTarget = strings.Split(*targetObjectsFlag, ",")
	}

	reportName := timestampedName("report")
	if err := os.Mkdir(filepath.Join(reportName), os.ModePerm); err != nil {
		log.Fatalf("Could not create \"report\" directory: %w", err)
	}
	for _, ns := range namespaces.Items {

		log.Printf("Processing namespace: %s", ns.Name)
		if shouldExclude(ns.Name, namespacesToIgnore, namespacesToTarget) {
			continue
		}

		if err := os.Mkdir(filepath.Join(reportName, ns.Name), os.ModePerm); err != nil {
			log.Printf("Error creating directory for namespace %s:, err: %w; skipping", ns.Name, err)
			continue
		}

		for _, res := range resources {
			gvr := schema.GroupVersionResource{
				Group:    res.Group,
				Version:  res.Version,
				Resource: res.Name,
			}

			resources, err := dynamicClient.Resource(gvr).Namespace(ns.Name).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				if apierrors.IsNotFound(err) {
					log.Printf("gvr %s not found in namespace %s", gvr, ns.Name)
					continue
				} else {
					log.Printf("Error fetching gvr %s resources in namespace %s: %s\n", gvr, ns.Name, err)
					continue
				}
			}

			for _, r := range resources.Items {
				if shouldExclude(r.GetName(), objectsToIgnore, objectsToTarget) {
					continue
				}

				filename := filepath.Join(reportName, ns.Name, r.GetName()+".yaml")
				data, err := yaml.Marshal(r)
				if err != nil {
					log.Printf("Error marshalling resource %s in namespace %s: %s\n", r.GetName(), ns.Name, err)
					continue
				}
				err = ioutil.WriteFile(filename, data, 0o644)
				if err != nil {
					log.Printf("Error writing file for resource %s in namespace %s: %s\n", r.GetName(), ns.Name, err)
				}
			}
		}
	}

	if err := tarReport("./"+reportName+".tar.gz", "./"+reportName); err != nil {
		log.Fatalf("Failed to tar report %s: %w", reportName, err)
	}
}

func timestampedName(name string) string {
	now := time.Now()

	// Format the time into a string. Example: "20060102-150405.000"
	// This format is YYYYMMDD-HHMMSS.milliseconds
	timestamp := now.Format("20060102-150405.000")

	return fmt.Sprintf("%s-%s", name, timestamp)
}

func resourcesToName(resources []*metav1.APIResource) []string {
	names := make([]string, len(resources))

	for i, resource := range resources {
		names[i] = resource.SingularName
	}

	return names
}

// contains checks if a slice contains a specific string.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

func shouldExclude(name string, excludes []string, includes []string) bool {
	// todo acpana: maps
	for _, exclude := range excludes {
		if strings.Contains(name, exclude) {
			log.Printf("Excluding %s as it contains %s", name, exclude)
			return true
		}
	}

	if len(includes) == 0 {
		return false // no includes means includes all in this case
	}

	for _, include := range includes {
		if strings.Contains(name, include) {
			log.Printf("Including %s as it contains %s", name, include)
			return false
		}
	}

	// by default exclude if nothing that has been defined included this namespace.
	log.Printf("Excluding %s as nothing targets it %s", name, includes)
	return true
}

// tarReport creates a .tar.gz file at targetPath from the contents of sourceDir.
func tarReport(targetPath, sourceDir string) error {
	outFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	gzWriter := gzip.NewWriter(outFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	err = filepath.Walk(sourceDir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(sourceDir, file)
		if err != nil {
			return err
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If not a directory, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			defer data.Close()
			if _, err := io.Copy(tarWriter, data); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
