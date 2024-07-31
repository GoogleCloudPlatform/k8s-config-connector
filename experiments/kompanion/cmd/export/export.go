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

package export

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

const (
	examples = `
	# export KCC resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion export

	# exclude certain namespace prefixes
	kompanion export --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	kompanion export --include-namespaces=my-team

	# target only specific namespace prefixes AND specific object prefixes
	kompanion export --target-namespaces=my-team --target-objects=logging
	`
)

var ExportCmd = &cobra.Command{
	Use:     "export",
	Short:   "export Config Connector resources",
	Example: examples,
	RunE:    runE,
	Args:    cobra.ExactArgs(0),
}

const (
	// flag names.
	kubeconfigFlag       = "kubeconfig"
	reportNamePrefixFlag = "report-prefix"

	targetNamespacesFlag = "target-namespaces"
	ignoreNamespacesFlag = "exclude-namespaces"

	targetObjectsFlag = "target-objects"
	ignoreObjectsFlag = "exclude-objects"

	workerRoutinesFlag = "worker-routines"
)

var (
	kubeconfig       string
	reportNamePrefix string

	targetNamespaces []string
	ignoreNamespaces []string

	targetObjects []string
	ignoreObjects []string

	workerRountines int
)

func init() {
	ExportCmd.Flags().StringVarP(&kubeconfig, kubeconfigFlag, "", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "path to the kubeconfig file. Defaults to ~HOME/.kube/config")
	ExportCmd.Flags().StringVarP(&reportNamePrefix, reportNamePrefixFlag, "", "report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")

	ExportCmd.Flags().StringArrayVarP(&targetNamespaces, targetNamespacesFlag, "", []string{}, "namespace prefix to target the export tool. Targets all if empty. Can be specified multiple times.")
	ExportCmd.Flags().StringArrayVarP(&ignoreNamespaces, ignoreNamespacesFlag, "", []string{"kube"}, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	ExportCmd.Flags().StringArrayVarP(&targetObjects, targetObjectsFlag, "", []string{}, "object name prefix to target. Targets all if empty. Can be specified multiple times.")
	ExportCmd.Flags().StringArrayVarP(&ignoreObjects, ignoreObjectsFlag, "", []string{}, "object name prefix to ignore. Excludes nothing if empty. Can be specified multiple times.")

	ExportCmd.Flags().IntVarP(&workerRountines, workerRoutinesFlag, "", 10, "Configure the number of worker rountines to export namespaces with. Defaults to 10. ")
}

var (
	reportName string
	resources  []*metav1.APIResource
)

// tracks the namespaces to be exported.
// thread safe.
type namespacesWorkLog struct {
	mu         sync.Mutex
	namespaces []string // will be treated as a FIFO queue
	errs       []error  // for err accumulaiton while working on work items
}

func (n *namespacesWorkLog) AddError(e error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.errs = append(n.errs, e)
}

func (n *namespacesWorkLog) GetErrors() []error {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.errs
}

func (n *namespacesWorkLog) GetWork() string {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(n.namespaces) == 0 {
		return ""
	}

	workItem := n.namespaces[0]
	n.namespaces = n.namespaces[1:]

	return workItem
}

func processNamespace(namespaceName string, dynamicClient *dynamic.DynamicClient) error {
	log.Printf("Processing namespace: %s", namespaceName)
	if shouldExclude(namespaceName, ignoreNamespaces, targetNamespaces) {
		return nil
	}

	if err := os.Mkdir(filepath.Join(reportName, namespaceName), 0o700); err != nil {
		return fmt.Errorf("error creating directory for namespace %s:, err: %w; skipping", namespaceName, err)
	}

	for _, res := range resources {
		gvr := schema.GroupVersionResource{
			Group:    res.Group,
			Version:  res.Version,
			Resource: res.Name,
		}
		// todo acpana debug logs
		// log.Printf("Looking for gvr %s in namespace %s", gvr, ns.Name)

		resources, err := dynamicClient.Resource(gvr).Namespace(namespaceName).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error fetching gvr %s resources in namespace %s: %w", gvr, namespaceName, err)
		}

		// todo acpana debug logs
		// log.Printf("Found %d resources of type gvr %s in namespace %s", len(resources.Items), gvr, ns.Name)

		for _, r := range resources.Items {
			if shouldExclude(r.GetName(), ignoreObjects, targetObjects) {
				continue
			}

			// todo acpana if the file exists, log error but move on
			filename := filepath.Join(reportName, namespaceName, fmt.Sprintf("%s_%s.yaml", r.GroupVersionKind().Kind, r.GetName()))
			data, err := yaml.Marshal(r)
			if err != nil {
				return fmt.Errorf("error marshalling resource %s in namespace %s: %w", r.GetName(), namespaceName, err)
			}
			err = ioutil.WriteFile(filename, data, 0o644)
			if err != nil {
				return fmt.Errorf("error writing file for resource %s in namespace %s: %w", r.GetName(), namespaceName, err)
			}
		}
	}

	return nil
}

func validateFlags() error {
	if workerRountines <= 0 || workerRountines > 100 {
		return fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", workerRountines, workerRoutinesFlag)
	}

	return nil
}

func runE(_ *cobra.Command, _ []string) error {
	log.Printf("Running kompanion export with kubeconfig: %s", kubeconfig)

	if err := validateFlags(); err != nil {
		return err
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %sw", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	// use the discovery client to iterate over all api resoruces
	discoveryClient := clientset.Discovery()
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("failed to get preferred resources: %w", err)
	}

	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, ".cnrm.cloud.google.com/") {
			// todo acpana log debug level
			// log.Printf("ApiResource %s group doesn't contain \"cnrm\"; skipping", apiResourceList.GroupVersion)
			continue
		}

		for _, apiResource := range apiResourceList.APIResources {
			if !apiResource.Namespaced {
				// todo acpana log debug level
				// log.Printf("ApiResource %s is not namespaced; skipping", apiResource.SingularName)
				continue
			}
			if !contains(apiResource.Verbs, "list") {
				// todo acpana log debug level
				// log.Printf("ApiResource %s is not listabble; skipping", apiResource.SingularName)
				continue
			}

			resources = append(resources, &apiResource)
		}
	}

	// incentiveze determinism for debuggability and idempotency
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].SingularName < resources[j].SingularName
	})

	// todo acpana debug logs
	// log.Printf("Going to iterate over the following resources %+v", resourcesToName(resources))

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}

	reportName = timestampedName(reportNamePrefix)
	if err := os.Mkdir(filepath.Join(reportName), 0o700); err != nil {
		return fmt.Errorf("could not create \"%s\" directory: %w", reportName, err)
	}

	// create the work log for go routine workers to use
	workLog := &namespacesWorkLog{}
	for _, ns := range namespaces.Items {
		workLog.namespaces = append(workLog.namespaces, ns.Name)
	}

	log.Printf("Starting worker threads to process namespaces")
	var wg sync.WaitGroup
	for i := 0; i < workerRountines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				ns := workLog.GetWork()
				if ns == "" {
					// no work
					return
				}

				if err := processNamespace(ns, dynamicClient); err != nil {
					workLog.AddError(err)
				}
			}
		}()
	}

	log.Printf("Dumping Config Connector objects to %s.tar.gz", reportName)
	wg.Wait()

	if err := tarReport("./"+reportName+".tar.gz", "./"+reportName); err != nil {
		return fmt.Errorf("failed to tar report %s: %w", reportName, err)
	}

	errors := workLog.GetErrors()
	if len(errors) != 0 {
		return fmt.Errorf("there have been errors in the export process: %+v", errors)
	}

	// clean up files
	if err := os.RemoveAll(reportName); err != nil {
		return fmt.Errorf("failed to clean up report folder: %w", err)
	}

	return nil
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
