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

package backup

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"cloud.google.com/go/storage"
)

type createOptions struct {
	kubecli.ClusterOptions
	cluster        string
	location       string
	project              string
	clusterProject       string
	bucket               string
	replicaBucket        string
	replicaProject       string
	outputDir            string
	namespace            string
	includeClusterBackup bool
	gkeBackupPlan        string
}

func NewCreateCmd() *cobra.Command {
	options := &createOptions{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Trigger an immediate backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(cmd, options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&options.clusterProject, "cluster-project", "", "GCP project ID where the cluster resides (defaults to --project)")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")
	cmd.Flags().StringVar(&options.replicaBucket, "replica-bucket", "", "Secondary replica GCS bucket for cross-region disaster recovery")
	cmd.Flags().StringVar(&options.replicaProject, "replica-project", "", "GCP project ID for replica bucket (defaults to --project)")
	cmd.Flags().StringVar(&options.outputDir, "output-dir", "", "Local directory path to save backups")
	cmd.Flags().StringVar(&options.namespace, "namespace", "cnrm-system", "Namespace where Config Connector is installed")
	cmd.Flags().BoolVar(&options.includeClusterBackup, "include-cluster-backup", false, "Trigger Backup for GKE (gkebackup.googleapis.com) to backup cluster workloads and persistent volumes")
	cmd.Flags().StringVar(&options.gkeBackupPlan, "gke-backup-plan", "", "Name of the GKE BackupPlan to use (defaults to <cluster>-backup-plan)")

	return cmd
}

type SummaryManifest struct {
	Counts        map[string]int         `json:"counts"`
	ClusterBackup *ClusterBackupMetadata `json:"clusterBackup,omitempty"`
	Integrity     map[string]string      `json:"integrity,omitempty"`
}

func runCreate(cmd *cobra.Command, options *createOptions) error {
	ctx := context.Background()
	if cmd != nil {
		ctx = cmd.Context()
	}
	if options.bucket == "" && options.outputDir == "" {
		return fmt.Errorf("either --bucket or --output-dir is required")
	}

	if options.outputDir != "" {
		if strings.Contains(options.outputDir, "..") {
			return fmt.Errorf("invalid --output-dir %q: relative path traversal sequences ('..') are prohibited", options.outputDir)
		}
		options.outputDir = filepath.Clean(options.outputDir)
	}

	if options.ClusterOptions.Context == "" && options.cluster != "" {
		rawConfig, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
		if err == nil {
			for ctxName := range rawConfig.Contexts {
				if ctxName == options.cluster || strings.Contains(ctxName, "_"+options.cluster) || strings.HasSuffix(ctxName, options.cluster) {
					options.ClusterOptions.Context = ctxName
					break
				}
			}
		}
	}

	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	var gcsClient *storage.Client
	if options.bucket != "" {
		var gcsOptions []option.ClientOption
		if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
			gcsOptions = append(gcsOptions, option.WithHTTPClient(httpClient.(*http.Client)))
		}
		client, err := storage.NewClient(ctx, gcsOptions...)
		if err != nil {
			return fmt.Errorf("creating GCS client: %w", err)
		}
		defer client.Close()
		gcsClient = client
	}

	var replicaClient *storage.Client
	if options.replicaBucket != "" {
		var replicaOptions []option.ClientOption
		if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
			replicaOptions = append(replicaOptions, option.WithHTTPClient(httpClient.(*http.Client)))
		}
		client, err := storage.NewClient(ctx, replicaOptions...)
		if err != nil {
			return fmt.Errorf("creating replica GCS client: %w", err)
		}
		defer client.Close()
		replicaClient = client
	}

	timestamp := time.Now().UTC().Format("2006_01_02_15_04_05")
	clusterName := options.cluster
	if clusterName == "" {
		clusterName = "default-cluster"
	}

	destinations := []string{}
	if options.bucket != "" {
		destinations = append(destinations, fmt.Sprintf("GCS: gs://%s", options.bucket))
	}
	if options.replicaBucket != "" {
		destinations = append(destinations, fmt.Sprintf("Replica GCS: gs://%s", options.replicaBucket))
	}
	if options.outputDir != "" {
		destinations = append(destinations, fmt.Sprintf("Directory: %s", options.outputDir))
	}
	fmt.Fprintf(cmd.OutOrStdout(), "Starting backup for cluster %q to [%s] (timestamp: %s)...\n", clusterName, strings.Join(destinations, ", "), timestamp)

	// Discover KCC resources
	_, resourceLists, err := kubeClient.DiscoveryClient.ServerGroupsAndResources()
	if err != nil {
		return fmt.Errorf("discovering server resources: %w", err)
	}

	stats := make(map[string]int)
	integrity := make(map[string]string)
	var integrityMu sync.Mutex
	backedUpGKs := make(map[string]bool)

	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			continue
		}

		if !strings.HasSuffix(gv.Group, ".cnrm.cloud.google.com") {
			continue
		}
		// Skip operator configurations
		if gv.Group == "core.cnrm.cloud.google.com" {
			continue
		}

		for _, resource := range resourceList.APIResources {
			if strings.Contains(resource.Name, "/status") || strings.Contains(resource.Name, "/finalizers") {
				continue
			}

			gk := fmt.Sprintf("%s/%s", gv.Group, resource.Kind)
			if backedUpGKs[gk] {
				continue
			}

			gvk := gv.WithKind(resource.Kind)
			count, err := backupResource(ctx, kubeClient, gcsClient, replicaClient, options, clusterName, timestamp, gvk, integrity, &integrityMu)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to backup %v: %v\n", gvk, err)
			}
			if count > 0 {
				stats[gk] = count
			}
			backedUpGKs[gk] = true
		}
	}

	var clusterBackupMeta *ClusterBackupMetadata
	if options.includeClusterBackup {
		planName := options.gkeBackupPlan
		if planName == "" {
			planName = fmt.Sprintf("%s-backup-plan", clusterName)
		}
		gkeLoc := options.location
		if gkeLoc == "" {
			gkeLoc = "us-central1"
		}
		gkeProj := options.clusterProject
		if gkeProj == "" {
			gkeProj = options.project
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Triggering cluster backup via Backup for GKE (plan: %s)...\n", planName)
		gkeMgr, err := NewGKEBackupManager(ctx)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to initialize GKE Backup manager: %v\n", err)
			clusterBackupMeta = &ClusterBackupMetadata{
				Provider:     "gkebackup.googleapis.com",
				State:        "FAILED",
				ErrorMessage: err.Error(),
			}
		} else {
			meta, err := gkeMgr.TriggerBackup(ctx, gkeProj, gkeLoc, planName, 0)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to trigger Backup for GKE: %v\n", err)
				clusterBackupMeta = &ClusterBackupMetadata{
					Provider:     "gkebackup.googleapis.com",
					State:        "SKIPPED_AGENT_OR_PERMISSION_ERROR",
					ErrorMessage: err.Error(),
				}
			} else {
				clusterBackupMeta = meta
				fmt.Fprintf(cmd.OutOrStdout(), "GKE cluster backup initiated: %s\n", meta.BackupName)
			}
		}
	}

	if err := writeSummary(ctx, gcsClient, replicaClient, options, clusterName, timestamp, stats, integrity, clusterBackupMeta); err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to write summary.json: %v\n", err)
	}

	total := 0
	for _, count := range stats {
		total += count
	}
	fmt.Fprintf(cmd.OutOrStdout(), "\nBackup completed successfully. Total resources backed up: %d\n", total)

	return nil
}

func writeSummary(ctx context.Context, gcsClient, replicaClient *storage.Client, options *createOptions, cluster, timestamp string, stats map[string]int, integrity map[string]string, clusterBackup *ClusterBackupMetadata) error {
	summary := SummaryManifest{
		Counts:        stats,
		ClusterBackup: clusterBackup,
		Integrity:     integrity,
	}
	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}

	if gcsClient != nil && options.bucket != "" {
		objectName := fmt.Sprintf("%s/%s/summary.json", cluster, timestamp)
		wc := gcsClient.Bucket(options.bucket).Object(objectName).NewWriter(ctx)
		if _, err := wc.Write(data); err != nil {
			_ = wc.Close()
			return err
		}
		if err := wc.Close(); err != nil {
			return err
		}
	}

	if replicaClient != nil && options.replicaBucket != "" {
		objectName := fmt.Sprintf("%s/%s/summary.json", cluster, timestamp)
		wc := replicaClient.Bucket(options.replicaBucket).Object(objectName).NewWriter(ctx)
		if _, err := wc.Write(data); err != nil {
			_ = wc.Close()
			return err
		}
		if err := wc.Close(); err != nil {
			return err
		}
	}

	if options.outputDir != "" {
		cleanBase := filepath.Clean(options.outputDir)
		summaryDir := filepath.Join(cleanBase, cluster, timestamp)
		if err := os.MkdirAll(summaryDir, 0755); err != nil {
			return fmt.Errorf("creating summary dir %s: %w", summaryDir, err)
		}
		summaryPath := filepath.Join(summaryDir, "summary.json")
		if err := os.WriteFile(summaryPath, data, 0644); err != nil {
			return fmt.Errorf("writing summary file %s: %w", summaryPath, err)
		}
	}

	return nil
}

func backupResource(ctx context.Context, kubeClient *kubecli.Client, gcsClient, replicaClient *storage.Client, options *createOptions, cluster, timestamp string, gvk schema.GroupVersionKind, integrity map[string]string, integrityMu *sync.Mutex) (int, error) {
	limit := int64(500)
	continueToken := ""
	count := 0

	for {
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)

		listOptions := []client.ListOption{
			client.Limit(limit),
		}
		if continueToken != "" {
			listOptions = append(listOptions, client.Continue(continueToken))
		}

		if err := kubeClient.List(ctx, list, listOptions...); err != nil {
			return count, err
		}

		for _, item := range list.Items {
			if err := backupObject(ctx, gcsClient, replicaClient, options, cluster, timestamp, item, integrity, integrityMu); err != nil {
				fmt.Printf("Warning: failed to backup object %s/%s (%s): %v\n", item.GetNamespace(), item.GetName(), item.GetKind(), err)
				continue
			}
			count++
		}

		continueToken = list.GetContinue()
		if continueToken == "" {
			break
		}
	}

	return count, nil
}

func backupObject(ctx context.Context, gcsClient, replicaClient *storage.Client, options *createOptions, cluster, timestamp string, obj unstructured.Unstructured, integrity map[string]string, integrityMu *sync.Mutex) error {
	sanitizeObject(&obj)

	data, err := yaml.Marshal(obj.Object)
	if err != nil {
		return err
	}

	name := obj.GetName()
	namespace := obj.GetNamespace()
	if namespace == "" {
		namespace = "_cluster_scoped"
	}
	kind := strings.ToLower(obj.GetKind())
	relPath := fmt.Sprintf("%s/%s/%s.yaml", namespace, kind, name)

	// Record SHA-256 cryptographic digest
	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("sha256:%x", hash)
	if integrity != nil && integrityMu != nil {
		integrityMu.Lock()
		integrity[relPath] = hashStr
		integrityMu.Unlock()
	}

	if gcsClient != nil && options.bucket != "" {
		objectName := fmt.Sprintf("%s/%s/%s/%s/%s.yaml", cluster, timestamp, namespace, kind, name)
		wc := gcsClient.Bucket(options.bucket).Object(objectName).NewWriter(ctx)
		if _, err := wc.Write(data); err != nil {
			_ = wc.Close()
			return err
		}
		if err := wc.Close(); err != nil {
			return err
		}
	}

	if replicaClient != nil && options.replicaBucket != "" {
		objectName := fmt.Sprintf("%s/%s/%s/%s/%s.yaml", cluster, timestamp, namespace, kind, name)
		wc := replicaClient.Bucket(options.replicaBucket).Object(objectName).NewWriter(ctx)
		if _, err := wc.Write(data); err != nil {
			_ = wc.Close()
			return fmt.Errorf("writing to replica bucket %s: %w", options.replicaBucket, err)
		}
		if err := wc.Close(); err != nil {
			return fmt.Errorf("closing replica object %s: %w", objectName, err)
		}
	}

	if options.outputDir != "" {
		cleanBase := filepath.Clean(options.outputDir)
		targetDir := filepath.Join(cleanBase, cluster, timestamp, namespace, kind)
		cleanTarget := filepath.Clean(targetDir)
		if !strings.HasPrefix(cleanTarget, cleanBase+string(filepath.Separator)) && cleanTarget != cleanBase {
			return fmt.Errorf("path traversal violation detected: %s is outside %s", cleanTarget, cleanBase)
		}
		if err := os.MkdirAll(cleanTarget, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", cleanTarget, err)
		}
		targetFile := filepath.Join(cleanTarget, name+".yaml")
		if err := os.WriteFile(targetFile, data, 0644); err != nil {
			return fmt.Errorf("writing object file %s: %w", targetFile, err)
		}
	}

	return nil
}

func sanitizeObject(obj *unstructured.Unstructured) {
	unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
	unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(obj.Object, "metadata", "generation")
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(obj.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(obj.Object, "metadata", "ownerReferences")
	unstructured.RemoveNestedField(obj.Object, "metadata", "finalizers")
	unstructured.RemoveNestedField(obj.Object, "status")

	annotations := obj.GetAnnotations()
	if annotations != nil {
		delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
		delete(annotations, "deployment.kubernetes.io/revision")
		if len(annotations) == 0 {
			unstructured.RemoveNestedField(obj.Object, "metadata", "annotations")
		} else {
			obj.SetAnnotations(annotations)
		}
	}
}
