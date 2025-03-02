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
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

const (
	examples = `
	# export Config Connector resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion export

	# exclude certain namespace prefixes
	kompanion export --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	kompanion export --target-namespaces=my-team

	# target only specific namespace prefixes AND specific object prefixes
	kompanion export --target-namespaces=my-team --target-objects=logging
	`
)

func BuildExportCmd() *cobra.Command {
	var opts ExportOptions

	cmd := &cobra.Command{
		Use:     "export",
		Short:   "export Config Connector resources",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunExport(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, kubeconfigFlag, "", opts.kubeconfig, "path to the kubeconfig file.")
	cmd.Flags().StringVarP(&opts.reportNamePrefix, reportNamePrefixFlag, "", "report", "Prefix for the report name. The tool appends a timestamp to this in the format \"YYYYMMDD-HHMMSS.milliseconds\".")

	cmd.Flags().StringArrayVarP(&opts.targetNamespaces, targetNamespacesFlag, "", []string{}, "namespace prefix to target the export tool. Targets all if empty. Can be specified multiple times.")
	cmd.Flags().StringArrayVarP(&opts.ignoreNamespaces, ignoreNamespacesFlag, "", []string{"kube"}, "namespace prefix to ignore. Excludes nothing if empty. Can be specified multiple times. Defaults to \"kube\".")

	cmd.Flags().StringArrayVarP(&opts.targetObjects, targetObjectsFlag, "", []string{}, "object name prefix to target. Targets all if empty. Can be specified multiple times.")
	cmd.Flags().StringArrayVarP(&opts.ignoreObjects, ignoreObjectsFlag, "", []string{}, "object name prefix to ignore. Excludes nothing if empty. Can be specified multiple times.")

	cmd.Flags().IntVarP(&opts.workerRountines, workerRoutinesFlag, "", 10, "Configure the number of worker routines to export namespaces with. Defaults to 10. ")

	return cmd
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

type ExportOptions struct {
	kubeconfig       string
	reportNamePrefix string

	targetNamespaces []string
	ignoreNamespaces []string

	targetObjects []string
	ignoreObjects []string

	workerRountines int
}

// Task is implemented by our namespace-collection routine, or anything else we want to run in parallel.
type Task interface {
	Run(ctx context.Context) error
}

// tracks the namespaces to be exported.
// thread safe.
type taskQueue struct {
	mu    sync.Mutex
	tasks []Task // will be treated as a FIFO queue
}

func (n *taskQueue) GetWork() Task {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(n.tasks) == 0 {
		return nil
	}

	workItem := n.tasks[0]
	n.tasks = n.tasks[1:]

	return workItem
}

func (n *taskQueue) AddTask(t Task) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.tasks = append(n.tasks, t)
}

type dumpResourcesTask struct {
	// Namespace is the namespace to filter down
	Namespace string

	DynamicClient *dynamic.DynamicClient

	// Resources is the list of resources to query
	Resources []schema.GroupVersionResource

	ReportDir string

	ShouldExcludeObject func(id types.NamespacedName) bool
}

func (t *dumpResourcesTask) Run(ctx context.Context) error {
	log := klog.FromContext(ctx)

	for _, gvr := range t.Resources {
		// todo acpana debug logs
		// log.Printf("Looking for gvr %s in namespace %s", gvr, ns.Name)

		var resources *unstructured.UnstructuredList
		if t.Namespace != "" {
			log.Info(fmt.Sprintf("Fetching %v in namespace %s", gvr.Resource, t.Namespace))
			r, err := t.DynamicClient.Resource(gvr).Namespace(t.Namespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("fetching gvr %s resources in namespace %s: %w", gvr, t.Namespace, err)
			}
			resources = r

		} else {
			log.Info(fmt.Sprintf("Fetching %v", gvr.Resource))
			r, err := t.DynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
			if err != nil {
				return fmt.Errorf("fetching gvr %s resources: %w", gvr, err)
			}
			resources = r
		}

		// todo acpana debug logs
		// log.Printf("Found %d resources of type gvr %s in namespace %s", len(resources.Items), gvr, ns.Name)

		for _, r := range resources.Items {
			id := types.NamespacedName{
				Namespace: r.GetNamespace(),
				Name:      r.GetName(),
			}
			if t.ShouldExcludeObject(id) {
				continue
			}

			data, err := yaml.Marshal(r)
			if err != nil {
				return fmt.Errorf("error marshalling resource %s: %w", id, err)
			}

			// todo acpana if the file exists, log error but move on
			filename := filepath.Join(t.ReportDir, id.Namespace, fmt.Sprintf("%s_%s.yaml", r.GroupVersionKind().Kind, id.Name))
			filedir := filepath.Dir(filename)
			if err := os.MkdirAll(filedir, 0o700); err != nil {
				return fmt.Errorf("error creating directory %s: %w", filedir, err)
			}

			if err := os.WriteFile(filename, data, 0o644); err != nil {
				return fmt.Errorf("error writing file for resource %s in namespace %s: %w", id.Name, id.Namespace, err)
			}
		}
	}

	return nil
}

func (opts *ExportOptions) validateFlags() error {
	if opts.workerRountines <= 0 || opts.workerRountines > 100 {
		return fmt.Errorf("invalid value %d for flag %s. Supported values are [1,100]", opts.workerRountines, workerRoutinesFlag)
	}

	return nil
}

func getRESTConfig(ctx context.Context, opts *ExportOptions) (*rest.Config, error) {
	var loadingRules clientcmd.ClientConfigLoader
	if opts.kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{
			// ClusterInfo: clientcmdapi.Cluster{Server: masterUrl},
		}).ClientConfig()
}

func RunExport(ctx context.Context, opts *ExportOptions) error {
	log.Printf("Running kompanion export with kubeconfig: %s", opts.kubeconfig)

	if err := opts.validateFlags(); err != nil {
		return err
	}

	config, err := getRESTConfig(ctx, opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	// We rely more on server-side rate limiting now, so give it a high client-side QPS
	if config.QPS == 0 {
		config.QPS = 100
		config.Burst = 20
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating Kubernetes clientset: %sw", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	// use the discovery client to iterate over all api resources
	discoveryClient := clientset.Discovery()
	var resources []schema.GroupVersionResource
	resources, err = utils.GetResources(discoveryClient, resources)
	if err != nil {
		return fmt.Errorf("error fetching resources: %w", err)
	}

	// todo acpana debug logs
	// log.Printf("Going to iterate over the following resources %+v", resourcesToName(resources))

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching namespaces: %w", err)
	}

	reportName := timestampedName(opts.reportNamePrefix)
	reportTempDir := filepath.Join(".", reportName)
	reportFile := filepath.Join(".", reportName+".tar.gz")
	if err := os.Mkdir(filepath.Join(reportTempDir), 0o700); err != nil {
		return fmt.Errorf("could not create %q directory: %w", reportTempDir, err)
	}

	shouldExcludeObject := func(id types.NamespacedName) bool {
		if shouldExclude(id.Namespace, opts.ignoreNamespaces, opts.targetNamespaces) {
			return true
		}
		return shouldExclude(id.Name, opts.ignoreObjects, opts.targetObjects)
	}

	// create the work log for go routine workers to use
	q := &taskQueue{}

	// Parallelize across resources, unless we are scoped to a few namespaces
	// The thought is that if users target a particular namespace (or a few), they may not have cluster-wide permission.
	perNamespace := len(opts.targetNamespaces) > 0
	if perNamespace {
		for _, ns := range namespaces.Items {
			if shouldExclude(ns.Name, opts.ignoreNamespaces, opts.targetNamespaces) {
				continue
			}

			q.AddTask(&dumpResourcesTask{
				Namespace:           ns.Name,
				Resources:           resources,
				DynamicClient:       dynamicClient,
				ReportDir:           reportTempDir,
				ShouldExcludeObject: shouldExcludeObject,
			})
		}
	} else {
		for _, resource := range resources {
			q.AddTask(&dumpResourcesTask{
				Resources:           []schema.GroupVersionResource{resource},
				DynamicClient:       dynamicClient,
				ReportDir:           reportTempDir,
				ShouldExcludeObject: shouldExcludeObject,
			})
		}
	}

	log.Printf("Starting worker threads to process Config Connector objects")
	var wg sync.WaitGroup

	var errs []error
	var errsMutex sync.Mutex

	for i := 0; i < opts.workerRountines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				task := q.GetWork()
				if task == nil {
					// no work
					return
				}

				if err := task.Run(ctx); err != nil {
					errsMutex.Lock()
					errs = append(errs, err)
					errsMutex.Unlock()
				}
			}
		}()
	}

	log.Printf("Dumping Config Connector objects to %s", reportFile)
	wg.Wait()

	if err := tarReport(reportFile, reportTempDir); err != nil {
		return fmt.Errorf("failed to create report file %s: %w", reportFile, err)
	}

	if len(errs) != 0 {
		return fmt.Errorf("there have been errors in the export process: %+v", errors.Join(errs...))
	}

	// clean up files
	if err := os.RemoveAll(reportTempDir); err != nil {
		return fmt.Errorf("failed to clean up report folder %q: %w", reportTempDir, err)
	}

	fmt.Fprintf(os.Stderr, "created report file: %v\n", reportFile)

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
