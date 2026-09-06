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
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"cloud.google.com/go/storage"
)

type restoreOptions struct {
	kubecli.ClusterOptions
	cluster                  string
	targetClusterLocation    string
	sourceBucket             string
	fromDir                  string
	backupTimestamp          string
	project                  string
	targetProject            string
	targetRegion             string
	regionMapping            string
	filterNamespace          string
	skipMissingCRDs          bool
	namespace                string
	autoCreateNamespaces     bool
	deletionPolicy           string
	managementConflictPolicy string
	dryRun                   bool
}

func NewRestoreCmd() *cobra.Command {
	options := &restoreOptions{}

	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore resources from GCS or local directory to a target cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRestore(cmd.Context(), options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.targetClusterLocation, "target-cluster-location", "", "Location of the target cluster")
	cmd.Flags().StringVar(&options.sourceBucket, "source-bucket", "", "Source GCS bucket name")
	cmd.Flags().StringVar(&options.fromDir, "from-dir", "", "Local directory containing backup files to restore from")
	cmd.Flags().StringVar(&options.backupTimestamp, "backup-timestamp", "latest", "Backup timestamp (YYYY_MM_DD_HH_MM_SS or 'latest')")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&options.targetProject, "target-project", "", "DR target project ID override")
	cmd.Flags().StringVar(&options.targetRegion, "target-region", "", "DR target region for regional resources")
	cmd.Flags().StringVar(&options.regionMapping, "region-mapping", "", "Source to target region mapping (e.g. us-central1=us-east4)")
	cmd.Flags().StringVar(&options.filterNamespace, "filter-namespace", "", "Only restore resources belonging to this specific namespace")
	cmd.Flags().BoolVar(&options.skipMissingCRDs, "skip-missing-crds", false, "Skip resources whose CRDs are not installed in the target cluster")
	cmd.Flags().StringVar(&options.namespace, "namespace", "cnrm-system", "Namespace where Config Connector is installed")
	cmd.Flags().BoolVar(&options.autoCreateNamespaces, "auto-create-namespaces", true, "Automatically create target namespaces if they do not exist")
	cmd.Flags().StringVar(&options.deletionPolicy, "deletion-policy", "abandon", "Deletion policy annotation (abandon or delete)")
	cmd.Flags().StringVar(&options.managementConflictPolicy, "management-conflict-policy", "none", "Management conflict policy (none, resource, or priority)")
	cmd.Flags().BoolVar(&options.dryRun, "dry-run", false, "Perform a dry-run validation")

	return cmd
}

func runRestore(ctx context.Context, options *restoreOptions) error {
	if options.sourceBucket == "" && options.fromDir == "" {
		return fmt.Errorf("either --source-bucket or --from-dir is required")
	}

	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	clusterName := options.cluster
	if clusterName == "" {
		clusterName = "default-cluster"
	}

	var objects []*unstructured.Unstructured
	if options.fromDir != "" {
		objs, err := loadObjectsFromDir(options.fromDir, clusterName, options.backupTimestamp)
		if err != nil {
			return fmt.Errorf("loading objects from directory %s: %w", options.fromDir, err)
		}
		objects = objs
	} else {
		var gcsOptions []option.ClientOption
		if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
			gcsOptions = append(gcsOptions, option.WithHTTPClient(httpClient.(*http.Client)))
		}
		gcsClient, err := storage.NewClient(ctx, gcsOptions...)
		if err != nil {
			return fmt.Errorf("creating GCS client: %w", err)
		}
		defer gcsClient.Close()

		objs, err := loadObjectsFromGCS(ctx, gcsClient, options.sourceBucket, clusterName, options.backupTimestamp)
		if err != nil {
			return fmt.Errorf("loading objects from GCS: %w", err)
		}
		objects = objs
	}

	if len(objects) == 0 {
		return fmt.Errorf("no resources found in backup")
	}

	if options.filterNamespace != "" {
		var filtered []*unstructured.Unstructured
		for _, obj := range objects {
			if obj.GetNamespace() == options.filterNamespace {
				filtered = append(filtered, obj)
			}
		}
		objects = filtered
		fmt.Printf("Filtered to %d resources in namespace %q.\n", len(objects), options.filterNamespace)
	}

	fmt.Printf("Found %d resources in backup. Validating against target cluster...\n", len(objects))
	validObjects, err := validateResources(ctx, kubeClient, objects, options.skipMissingCRDs)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	objects = validObjects

	// Auto-create namespaces if enabled
	if options.autoCreateNamespaces {
		if err := ensureNamespaces(ctx, kubeClient, objects, options.dryRun); err != nil {
			fmt.Printf("Warning: failed to ensure all namespaces: %v\n", err)
		}
	}

	fmt.Println("Sorting resources for restore DAG dependencies...")
	sortResources(objects)

	fmt.Println("Initiating restore to target cluster...")
	successCount := 0
	failureCount := 0
	for _, obj := range objects {
		if err := applyObject(ctx, kubeClient, obj, options); err != nil {
			fmt.Printf("Warning: failed to restore %s/%s (%s): %v\n", obj.GetNamespace(), obj.GetName(), obj.GetKind(), err)
			failureCount++
		} else {
			successCount++
		}
	}

	fmt.Printf("\nRestore complete. %d succeeded, %d failed.\n", successCount, failureCount)
	if failureCount > 0 {
		return fmt.Errorf("restore encountered %d errors", failureCount)
	}
	return nil
}

func loadObjectsFromDir(rootDir, cluster, timestamp string) ([]*unstructured.Unstructured, error) {
	searchDir := rootDir
	// Check if rootDir contains cluster/timestamp subdirectories
	clusterDir := filepath.Join(rootDir, cluster)
	if fi, err := os.Stat(clusterDir); err == nil && fi.IsDir() {
		if timestamp == "latest" || timestamp == "" {
			entries, err := os.ReadDir(clusterDir)
			if err != nil {
				return nil, err
			}
			var timestamps []string
			for _, e := range entries {
				if e.IsDir() {
					timestamps = append(timestamps, e.Name())
				}
			}
			if len(timestamps) > 0 {
				sort.Strings(timestamps)
				searchDir = filepath.Join(clusterDir, timestamps[len(timestamps)-1])
			} else {
				searchDir = clusterDir
			}
		} else {
			searchDir = filepath.Join(clusterDir, timestamp)
		}
	}

	fmt.Printf("Loading backup files from directory: %s\n", searchDir)
	var objects []*unstructured.Unstructured

	err := filepath.WalkDir(searchDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || (!strings.HasSuffix(d.Name(), ".yaml") && !strings.HasSuffix(d.Name(), ".yml")) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Warning: failed to read file %s: %v\n", path, err)
			return nil
		}

		obj := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(data, &obj.Object); err != nil {
			fmt.Printf("Warning: failed to unmarshal YAML in %s: %v\n", path, err)
			return nil
		}
		if obj.GetKind() != "" && obj.GetAPIVersion() != "" {
			objects = append(objects, obj)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return objects, nil
}

func loadObjectsFromGCS(ctx context.Context, gcsClient *storage.Client, bucket, cluster, timestamp string) ([]*unstructured.Unstructured, error) {
	fmt.Printf("Loading backup from gs://%s/%s/%s/...\n", bucket, cluster, timestamp)

	backupTimestamp := timestamp
	if backupTimestamp == "latest" || backupTimestamp == "" {
		latest, err := findLatestBackup(ctx, gcsClient, bucket, cluster)
		if err != nil {
			return nil, fmt.Errorf("finding latest backup: %w", err)
		}
		backupTimestamp = latest
		fmt.Printf("Resolved 'latest' to timestamp: %s\n", backupTimestamp)
	}

	prefix := fmt.Sprintf("%s/%s/", cluster, backupTimestamp)
	it := gcsClient.Bucket(bucket).Objects(ctx, &storage.Query{Prefix: prefix})

	var objects []*unstructured.Unstructured
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("iterating GCS objects: %w", err)
		}

		if !strings.HasSuffix(attrs.Name, ".yaml") {
			continue
		}

		obj, err := loadObject(ctx, gcsClient, bucket, attrs.Name)
		if err != nil {
			fmt.Printf("Warning: failed to load %s: %v\n", attrs.Name, err)
			continue
		}
		objects = append(objects, obj)
	}

	return objects, nil
}

func ensureNamespaces(ctx context.Context, kubeClient *kubecli.Client, objects []*unstructured.Unstructured, dryRun bool) error {
	seenNamespaces := make(map[string]bool)
	for _, obj := range objects {
		ns := obj.GetNamespace()
		if ns == "" || ns == "default" || ns == "kube-system" || ns == "kube-public" || ns == "kube-node-lease" {
			continue
		}
		seenNamespaces[ns] = true
	}

	for ns := range seenNamespaces {
		nsObj := &unstructured.Unstructured{}
		nsObj.SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"})
		nsObj.SetName(ns)

		err := kubeClient.Get(ctx, client.ObjectKey{Name: ns}, nsObj)
		if err == nil {
			// Namespace already exists
			continue
		}

		if dryRun {
			fmt.Printf("Dry-run: would auto-create target namespace %q\n", ns)
			continue
		}

		fmt.Printf("Auto-creating missing target namespace %q...\n", ns)
		if err := kubeClient.Create(ctx, nsObj); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				fmt.Printf("Warning: failed to create namespace %s: %v\n", ns, err)
			}
		}
	}
	return nil
}

func findLatestBackup(ctx context.Context, gcsClient *storage.Client, bucket, cluster string) (string, error) {
	prefix := cluster + "/"
	it := gcsClient.Bucket(bucket).Objects(ctx, &storage.Query{Prefix: prefix, Delimiter: "/"})

	var timestamps []string
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return "", err
		}
		if attrs.Prefix != "" {
			timestamp := strings.TrimSuffix(strings.TrimPrefix(attrs.Prefix, prefix), "/")
			if timestamp != "" {
				timestamps = append(timestamps, timestamp)
			}
		}
	}

	if len(timestamps) == 0 {
		return "", fmt.Errorf("no backups found in gs://%s/%s/", bucket, cluster)
	}

	sort.Strings(timestamps)
	return timestamps[len(timestamps)-1], nil
}

func loadObject(ctx context.Context, gcsClient *storage.Client, bucket, objectName string) (*unstructured.Unstructured, error) {
	rc, err := gcsClient.Bucket(bucket).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	obj := &unstructured.Unstructured{}
	if err := yaml.Unmarshal(data, &obj.Object); err != nil {
		return nil, fmt.Errorf("unmarshaling YAML: %w", err)
	}

	return obj, nil
}

func validateResources(ctx context.Context, kubeClient *kubecli.Client, objects []*unstructured.Unstructured, skipMissingCRDs bool) ([]*unstructured.Unstructured, error) {
	// Check if GVKs exist in the target cluster
	_, resourceLists, err := kubeClient.DiscoveryClient.ServerGroupsAndResources()
	if err != nil {
		return nil, fmt.Errorf("discovering server resources: %w", err)
	}

	supportedGVKs := make(map[string]bool)
	for _, list := range resourceLists {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			continue
		}
		for _, r := range list.APIResources {
			gvk := gv.WithKind(r.Kind)
			supportedGVKs[gvk.String()] = true
		}
	}

	var validObjects []*unstructured.Unstructured
	for _, obj := range objects {
		gvk := obj.GroupVersionKind()
		if !supportedGVKs[gvk.String()] {
			if skipMissingCRDs {
				fmt.Printf("Notice: skipping %s/%s (%v) because CRD is not installed in target cluster\n", obj.GetNamespace(), obj.GetName(), gvk)
				continue
			}
			return nil, fmt.Errorf("resource type %v not supported in target cluster (is the CRD installed?)", gvk)
		}
		validObjects = append(validObjects, obj)
	}
	return validObjects, nil
}

func sortResources(objects []*unstructured.Unstructured) {
	// Priority-based sort to handle dependencies
	getPriority := func(obj *unstructured.Unstructured) int {
		kind := obj.GetKind()
		group := obj.GroupVersionKind().Group

		// Priority 0-2: Containers / Resource Manager hierarchy
		if group == "resourcemanager.cnrm.cloud.google.com" {
			switch kind {
			case "Organization":
				return 0
			case "Folder":
				return 1
			case "Project":
				return 2
			}
		}

		// Priority 3-4: KMS KeyRings & CryptoKeys
		if group == "kms.cnrm.cloud.google.com" {
			if kind == "KMSKeyRing" {
				return 3
			}
			return 4
		}

		// Priority 5: Identity (Service Accounts, Workload Identity Pools)
		if group == "iam.cnrm.cloud.google.com" && (kind == "IAMServiceAccount" || kind == "IAMWorkloadIdentityPool") {
			return 5
		}

		// Priority 6-7: Networking foundational
		if group == "compute.cnrm.cloud.google.com" {
			switch kind {
			case "ComputeNetwork":
				return 6
			case "ComputeSubnetwork", "ComputeRouter", "ComputeRoute", "ComputeFirewall":
				return 7
			}
		}

		// Priority 8: Primary state stores & databases
		if (group == "storage.cnrm.cloud.google.com" && kind == "StorageBucket") ||
			(group == "sql.cnrm.cloud.google.com" && kind == "SQLInstance") ||
			(group == "spanner.cnrm.cloud.google.com" && kind == "SpannerInstance") ||
			(group == "bigtable.cnrm.cloud.google.com" && kind == "BigtableInstance") ||
			(group == "redis.cnrm.cloud.google.com" && kind == "RedisInstance") {
			return 8
		}

		// Priority 9: Roots for messaging & secret storage
		if (group == "secretmanager.cnrm.cloud.google.com" && kind == "SecretManagerSecret") ||
			(group == "pubsub.cnrm.cloud.google.com" && kind == "PubSubTopic") ||
			(group == "artifactregistry.cnrm.cloud.google.com" && kind == "ArtifactRegistryRepository") {
			return 9
		}

		// Priority 10: Dependent child resources
		if (group == "secretmanager.cnrm.cloud.google.com" && kind == "SecretManagerSecretVersion") ||
			(group == "pubsub.cnrm.cloud.google.com" && kind == "PubSubSubscription") ||
			(group == "sql.cnrm.cloud.google.com" && (kind == "SQLDatabase" || kind == "SQLUser" || kind == "SQLSSLCert")) {
			return 10
		}

		// Priority 12: IAM policy attachments and member bindings (must be applied last)
		if group == "iam.cnrm.cloud.google.com" {
			if strings.HasSuffix(kind, "PolicyMember") || strings.HasSuffix(kind, "Policy") || strings.HasSuffix(kind, "Binding") {
				return 12
			}
		}
		if strings.HasSuffix(kind, "PolicyMember") || strings.HasSuffix(kind, "Binding") {
			return 12
		}

		// Priority 11: General resources
		return 11
	}

	sort.SliceStable(objects, func(i, j int) bool {
		pi := getPriority(objects[i])
		pj := getPriority(objects[j])
		if pi != pj {
			return pi < pj
		}
		// Within the same priority, sort by kind then name for consistency
		if objects[i].GetKind() != objects[j].GetKind() {
			return objects[i].GetKind() < objects[j].GetKind()
		}
		return objects[i].GetName() < objects[j].GetName()
	})
}

func applyObject(ctx context.Context, kubeClient *kubecli.Client, obj *unstructured.Unstructured, options *restoreOptions) error {
	// Sanitization
	unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
	unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(obj.Object, "metadata", "ownerReferences")
	unstructured.RemoveNestedField(obj.Object, "metadata", "finalizers")
	unstructured.RemoveNestedField(obj.Object, "metadata", "generation")
	unstructured.RemoveNestedField(obj.Object, "metadata", "creationTimestamp")

	// Remove status as it is derived from GCP and can cause issues during apply (e.g. SSA emulation)
	unstructured.RemoveNestedField(obj.Object, "status")

	// Acquisition and Conflict Prevention
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
	delete(annotations, "deployment.kubernetes.io/revision")

	deletionPolicy := options.deletionPolicy
	if deletionPolicy == "" {
		deletionPolicy = "abandon"
	}
	annotations["cnrm.cloud.google.com/deletion-policy"] = deletionPolicy

	managementConflictPolicy := options.managementConflictPolicy
	if managementConflictPolicy == "" {
		managementConflictPolicy = "none"
	}
	annotations["cnrm.cloud.google.com/management-conflict-prevention-policy"] = managementConflictPolicy

	// DR Transformations: Target Project
	if options.targetProject != "" {
		annotations["cnrm.cloud.google.com/project-id"] = options.targetProject
		if _, ok, _ := unstructured.NestedFieldNoCopy(obj.Object, "spec", "projectRef"); ok {
			_ = unstructured.SetNestedField(obj.Object, options.targetProject, "spec", "projectRef", "external")
		}
	}

	// DR Transformations: Region Remapping
	remapRegion := func(currentRegion string) string {
		if currentRegion == "" {
			return ""
		}
		if options.regionMapping != "" {
			parts := strings.Split(options.regionMapping, "=")
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == currentRegion {
				return strings.TrimSpace(parts[1])
			}
		}
		if options.targetRegion != "" {
			return options.targetRegion
		}
		return currentRegion
	}

	if location, ok, _ := unstructured.NestedString(obj.Object, "spec", "location"); ok {
		newLoc := remapRegion(location)
		if newLoc != location {
			_ = unstructured.SetNestedField(obj.Object, newLoc, "spec", "location")
		}
	}
	if region, ok, _ := unstructured.NestedString(obj.Object, "spec", "region"); ok {
		newReg := remapRegion(region)
		if newReg != region {
			_ = unstructured.SetNestedField(obj.Object, newReg, "spec", "region")
		}
	}

	obj.SetAnnotations(annotations)

	if options.dryRun {
		fmt.Printf("Dry-run: would restore %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
		return nil
	}

	// Apply to cluster using Server-Side Apply
	if err := kubeClient.Patch(ctx, obj, client.Apply, client.FieldOwner("config-connector-backup"), client.ForceOwnership); err != nil {
		return fmt.Errorf("applying object: %w", err)
	}

	fmt.Printf("Restored %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
	return nil
}
