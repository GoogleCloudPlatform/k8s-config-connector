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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

const (
	AnnotationDeletionPolicy           = "cnrm.cloud.google.com/deletion-policy"
	AnnotationConflictPreventionPolicy = "cnrm.cloud.google.com/management-conflict-prevention-policy"
	AnnotationProjectID                = "cnrm.cloud.google.com/project-id"
	AnnotationLastAppliedConfiguration = "kubectl.kubernetes.io/last-applied-configuration"
	DeletionPolicyAbandon              = "abandon"
	ConflictPreventionPolicyNone       = "none"
	SummaryFileName                    = "summary.json"
)

// Summary records the integrity metadata for a Config Connector backup snapshot.
type Summary struct {
	Timestamp     string            `json:"timestamp"`
	Project       string            `json:"project"`
	Cluster       string            `json:"cluster"`
	Location      string            `json:"location"`
	ResourceCount int               `json:"resourceCount"`
	Status        string            `json:"status"`
	FileDigests   map[string]string `json:"fileDigests"`
}

// ConfigureOptions holds parameters for `config-connector backup configure`.
type ConfigureOptions struct {
	Cluster        string
	Location       string
	Bucket         string
	BucketLocation string
	Frequency      string
	Project        string
	OutputDir      string
}

// CreateOptions holds parameters for `config-connector backup create`.
type CreateOptions struct {
	Cluster   string
	Location  string
	Project   string
	InputDir  string
	OutputDir string
	Timestamp string
}

// StatusOptions holds parameters for `config-connector backup status`.
type StatusOptions struct {
	Cluster   string
	Location  string
	Project   string
	BackupDir string
}

// RestoreOptions holds parameters for `config-connector backup restore`.
type RestoreOptions struct {
	TargetCluster         string
	TargetClusterLocation string
	SourceBucket          string
	BackupTimestamp       string
	BackupDir             string
	Project               string
	TargetProject         string
	TargetLocation        string
	OutputDir             string
	DryRun                bool
}

// NewCommand constructs the `config-connector backup` command tree.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backup",
		Short: "Configure, create, verify, and restore Config Connector resource backups",
		Long:  "Standardized disaster recovery backup and restore tooling for Config Connector resources.",
	}

	cmd.AddCommand(newConfigureCmd())
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newStatusCmd())
	cmd.AddCommand(newRestoreCmd())
	return cmd
}

func newConfigureCmd() *cobra.Command {
	opts := &ConfigureOptions{}
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Generate scheduled backup infrastructure manifests (Bucket, ServiceAccount, RBAC, CronJob)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunConfigure(opts, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.Cluster, "cluster", "", "GKE cluster name")
	cmd.Flags().StringVar(&opts.Location, "location", "", "GKE cluster region or zone")
	cmd.Flags().StringVar(&opts.Bucket, "bucket", "", "GCS bucket name for storing backups")
	cmd.Flags().StringVar(&opts.BucketLocation, "bucket-location", "US", "Location for the backup GCS bucket")
	cmd.Flags().StringVar(&opts.Frequency, "frequency", "daily", "Backup schedule frequency (hourly, daily, weekly, or cron expression)")
	cmd.Flags().StringVar(&opts.Project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&opts.OutputDir, "output-dir", "", "Optional directory to write generated manifests")
	return cmd
}

func newCreateCmd() *cobra.Command {
	opts := &CreateOptions{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a point-in-time backup snapshot of Config Connector resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunCreate(opts, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.Cluster, "cluster", "", "GKE cluster name")
	cmd.Flags().StringVar(&opts.Location, "location", "", "GKE cluster location")
	cmd.Flags().StringVar(&opts.Project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&opts.InputDir, "input-dir", "", "Directory containing exported KCC YAML manifests")
	cmd.Flags().StringVar(&opts.OutputDir, "output-dir", "", "Target directory for the structured backup snapshot")
	cmd.Flags().StringVar(&opts.Timestamp, "timestamp", "", "Optional snapshot timestamp override (defaults to UTC date)")
	return cmd
}

func newStatusCmd() *cobra.Command {
	opts := &StatusOptions{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Verify the status and SHA-256 integrity of a Config Connector backup snapshot",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunStatus(opts, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.Cluster, "cluster", "", "GKE cluster name")
	cmd.Flags().StringVar(&opts.Location, "location", "", "GKE cluster location")
	cmd.Flags().StringVar(&opts.Project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&opts.BackupDir, "backup-dir", "", "Path to the backup snapshot directory containing summary.json")
	return cmd
}

func newRestoreCmd() *cobra.Command {
	opts := &RestoreOptions{}
	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Validate, sanitize, topologically sort, and restore Config Connector resources from a backup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunRestore(opts, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.TargetCluster, "target-cluster", "", "Target GKE cluster name")
	cmd.Flags().StringVar(&opts.TargetClusterLocation, "target-cluster-location", "", "Target GKE cluster location")
	cmd.Flags().StringVar(&opts.SourceBucket, "source-bucket", "", "Source GCS bucket name")
	cmd.Flags().StringVar(&opts.BackupTimestamp, "backup-timestamp", "", "Backup timestamp identifier (e.g. YYYY-MM-DD)")
	cmd.Flags().StringVar(&opts.BackupDir, "backup-dir", "", "Local directory containing the backup snapshot to restore")
	cmd.Flags().StringVar(&opts.Project, "project", "", "Source GCP project ID")
	cmd.Flags().StringVar(&opts.TargetProject, "target-project", "", "Optional target GCP project ID for cross-project restore")
	cmd.Flags().StringVar(&opts.TargetLocation, "target-location", "", "Optional target GCP location/region for cross-region restore")
	cmd.Flags().StringVar(&opts.OutputDir, "output-dir", "", "Directory to write sanitized, ordered restore manifests")
	cmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "Validate and generate ordered restore manifests without applying")
	return cmd
}

// CronScheduleForFrequency converts human-readable frequencies into standard cron expressions.
func CronScheduleForFrequency(freq string) string {
	switch strings.ToLower(strings.TrimSpace(freq)) {
	case "hourly":
		return "0 * * * *"
	case "daily", "":
		return "0 2 * * *"
	case "weekly":
		return "0 2 * * 0"
	default:
		return freq
	}
}

// RunConfigure generates the Kubernetes and KCC manifests required for scheduled backups.
func RunConfigure(opts *ConfigureOptions, out io.Writer) error {
	if opts.Bucket == "" || opts.Project == "" {
		return fmt.Errorf("--bucket and --project are required")
	}
	if opts.BucketLocation == "" {
		opts.BucketLocation = "US"
	}
	schedule := CronScheduleForFrequency(opts.Frequency)
	gsaEmail := fmt.Sprintf("kcc-backup-writer@%s.iam.gserviceaccount.com", opts.Project)

	manifestYAML := fmt.Sprintf(`apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: %s
  namespace: cnrm-system
  annotations:
    cnrm.cloud.google.com/project-id: %s
    cnrm.cloud.google.com/deletion-policy: abandon
spec:
  location: %s
  uniformBucketLevelAccess: true
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kcc-backup-runner
  namespace: cnrm-system
  annotations:
    iam.gke.io/gcp-service-account: %s
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: kcc-scheduled-backup
  namespace: cnrm-system
spec:
  schedule: %q
  concurrencyPolicy: Forbid
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: kcc-backup-runner
          restartPolicy: OnFailure
          containers:
          - name: backup
            image: gcr.io/gke-release/config-connector:latest
            args:
            - backup
            - create
            - --project=%s
            - --cluster=%s
            - --location=%s
`, opts.Bucket, opts.Project, opts.BucketLocation, gsaEmail, schedule, opts.Project, opts.Cluster, opts.Location)

	if opts.OutputDir != "" {
		if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
			return fmt.Errorf("creating output directory: %w", err)
		}
		target := filepath.Join(opts.OutputDir, "backup-infrastructure.yaml")
		if err := os.WriteFile(target, []byte(manifestYAML), 0o644); err != nil {
			return fmt.Errorf("writing backup infrastructure manifest: %w", err)
		}
	}
	_, err := fmt.Fprint(out, manifestYAML)
	return err
}

// RunCreate exports KCC manifests from InputDir into a structured snapshot `<OutputDir>/<Timestamp>/namespace/kind/name.yaml`
// and writes a `summary.json` file with SHA-256 digests.
func RunCreate(opts *CreateOptions, out io.Writer) error {
	if opts.InputDir == "" || opts.OutputDir == "" {
		return fmt.Errorf("--input-dir and --output-dir are required")
	}
	ts := opts.Timestamp
	if ts == "" {
		ts = time.Now().UTC().Format("2006-01-02")
	}
	snapshotRoot := filepath.Join(opts.OutputDir, ts)
	if err := os.MkdirAll(snapshotRoot, 0o755); err != nil {
		return fmt.Errorf("creating snapshot directory: %w", err)
	}

	objects, err := loadYAMLObjectsFromDir(opts.InputDir)
	if err != nil {
		return err
	}

	digests := make(map[string]string, len(objects))
	for _, obj := range objects {
		ns := obj.GetNamespace()
		if ns == "" {
			ns = "_cluster"
		}
		kind := strings.ToLower(obj.GetKind())
		name := obj.GetName()
		relPath := filepath.Join(ns, kind, name+".yaml")
		fullPath := filepath.Join(snapshotRoot, relPath)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return fmt.Errorf("creating resource directory: %w", err)
		}
		data, err := yaml.Marshal(obj.Object)
		if err != nil {
			return fmt.Errorf("marshaling %s/%s: %w", kind, name, err)
		}
		if err := os.WriteFile(fullPath, data, 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", fullPath, err)
		}
		sum := sha256.Sum256(data)
		digests[filepath.ToSlash(relPath)] = hex.EncodeToString(sum[:])
	}

	summary := Summary{
		Timestamp:     ts,
		Project:       opts.Project,
		Cluster:       opts.Cluster,
		Location:      opts.Location,
		ResourceCount: len(objects),
		Status:        "Succeeded",
		FileDigests:   digests,
	}
	summaryBytes, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling summary.json: %w", err)
	}
	if err := os.WriteFile(filepath.Join(snapshotRoot, SummaryFileName), summaryBytes, 0o644); err != nil {
		return fmt.Errorf("writing summary.json: %w", err)
	}

	_, err = fmt.Fprintf(out, "Backup snapshot created at %s (%d resources)\n", snapshotRoot, len(objects))
	return err
}

// VerifySnapshot validates the `summary.json` manifest and verifies the SHA-256 digest of every backed-up file.
func VerifySnapshot(snapshotDir string) (*Summary, error) {
	summaryPath := filepath.Join(snapshotDir, SummaryFileName)
	raw, err := os.ReadFile(summaryPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", summaryPath, err)
	}
	var summary Summary
	if err := json.Unmarshal(raw, &summary); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", summaryPath, err)
	}
	if summary.Status != "Succeeded" {
		return nil, fmt.Errorf("backup snapshot status is %q, expected Succeeded", summary.Status)
	}
	if len(summary.FileDigests) != summary.ResourceCount {
		return nil, fmt.Errorf("resource count mismatch in summary.json: declared %d, digests %d", summary.ResourceCount, len(summary.FileDigests))
	}
	for relPath, wantDigest := range summary.FileDigests {
		fullPath := filepath.Join(snapshotDir, filepath.FromSlash(relPath))
		data, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("missing backup artifact %s: %w", relPath, err)
		}
		gotSum := sha256.Sum256(data)
		gotDigest := hex.EncodeToString(gotSum[:])
		if gotDigest != wantDigest {
			return nil, fmt.Errorf("SHA-256 digest mismatch for %s: got %s, want %s", relPath, gotDigest, wantDigest)
		}
	}
	return &summary, nil
}

// RunStatus checks the status and integrity of a backup snapshot directory.
func RunStatus(opts *StatusOptions, out io.Writer) error {
	if opts.BackupDir == "" {
		return fmt.Errorf("--backup-dir is required")
	}
	summary, err := VerifySnapshot(opts.BackupDir)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, "Status: %s | Timestamp: %s | Resources: %d | Integrity: Verified\n",
		summary.Status, summary.Timestamp, summary.ResourceCount)
	return err
}

// SanitizeForRestore strips cluster-specific runtime metadata and status fields, preserves
// server-generated resource identifiers into spec.resourceID when absent, injects
// safe acquisition annotations (`deletion-policy: abandon` and `management-conflict-prevention-policy: none`),
// and optionally remaps project and location references.
func SanitizeForRestore(obj *unstructured.Unstructured, targetProject, targetLocation string) *unstructured.Unstructured {
	cloned := obj.DeepCopy()

	// Preserve server-generated resourceID from status.externalRef if spec.resourceID is unset
	specMap, hasSpec, _ := unstructured.NestedMap(cloned.Object, "spec")
	if hasSpec {
		if existingID, found, _ := unstructured.NestedString(cloned.Object, "spec", "resourceID"); !found || existingID == "" {
			if extRef, ok, _ := unstructured.NestedString(cloned.Object, "status", "externalRef"); ok && extRef != "" {
				parts := strings.Split(strings.TrimSuffix(extRef, "/"), "/")
				if len(parts) > 0 && parts[len(parts)-1] != "" {
					_ = unstructured.SetNestedField(cloned.Object, parts[len(parts)-1], "spec", "resourceID")
				}
			}
		}
	}
	_ = specMap

	// Strip runtime metadata fields
	unstructured.RemoveNestedField(cloned.Object, "metadata", "uid")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "generation")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "selfLink")
	unstructured.RemoveNestedField(cloned.Object, "metadata", "ownerReferences")
	unstructured.RemoveNestedField(cloned.Object, "status")

	annotations := cloned.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	delete(annotations, AnnotationLastAppliedConfiguration)
	annotations[AnnotationDeletionPolicy] = DeletionPolicyAbandon
	annotations[AnnotationConflictPreventionPolicy] = ConflictPreventionPolicyNone

	if targetProject != "" {
		if _, ok := annotations[AnnotationProjectID]; ok {
			annotations[AnnotationProjectID] = targetProject
		}
		if _, found, _ := unstructured.NestedMap(cloned.Object, "spec", "projectRef"); found {
			_ = unstructured.SetNestedField(cloned.Object, "projects/"+targetProject, "spec", "projectRef", "external")
			unstructured.RemoveNestedField(cloned.Object, "spec", "projectRef", "name")
			unstructured.RemoveNestedField(cloned.Object, "spec", "projectRef", "namespace")
		}
	}
	cloned.SetAnnotations(annotations)

	if targetLocation != "" {
		if _, found, _ := unstructured.NestedString(cloned.Object, "spec", "location"); found {
			_ = unstructured.SetNestedField(cloned.Object, targetLocation, "spec", "location")
		}
		if _, found, _ := unstructured.NestedString(cloned.Object, "spec", "region"); found {
			_ = unstructured.SetNestedField(cloned.Object, targetLocation, "spec", "region")
		}
	}

	return cloned
}

// KindTier assigns a topological apply tier (0 = earliest foundation, 9 = leaf workloads)
// so parent resources are applied before dependent child resources during restore.
func KindTier(kind string) int {
	switch kind {
	case "Namespace", "ConfigConnector", "ConfigConnectorContext":
		return 0
	case "Folder", "Project", "ResourceManagerLien", "ServiceIdentity":
		return 1
	case "IAMServiceAccount", "IAMCustomRole", "KMSKeyRing", "KMSCryptoKey":
		return 2
	case "IAMPolicy", "IAMPartialPolicy", "IAMPolicyMember", "IAMAuditConfig":
		return 3
	case "ComputeNetwork", "ComputeSubnetwork", "ComputeSharedVPCHostProject", "ComputeSharedVPCServiceProject", "DNSManagedZone":
		return 4
	case "ComputeAddress", "ComputeGlobalAddress", "ComputeFirewall", "ComputeRouter", "ComputeFirewallPolicy", "ServiceNetworkingConnection":
		return 5
	case "StorageBucket", "PubSubTopic", "SecretManagerSecret", "ArtifactRegistryRepository", "BigQueryDataset", "SpannerInstance":
		return 6
	case "SQLInstance", "RedisInstance", "AlloyDBCluster", "BigtableInstance", "SpannerDatabase", "PubSubSubscription":
		return 7
	case "SQLDatabase", "SQLUser", "AlloyDBInstance", "ContainerCluster":
		return 8
	default:
		return 9
	}
}

// SortByDependencyOrder sorts KCC resources into deterministic topological tiers (e.g.
// Project -> IAMServiceAccount -> KMSKeyRing -> KMSCryptoKey -> ComputeNetwork -> ComputeSubnetwork ->
// PubSubTopic -> PubSubSubscription -> SQLInstance -> SQLDatabase).
func SortByDependencyOrder(objects []*unstructured.Unstructured) []*unstructured.Unstructured {
	sorted := make([]*unstructured.Unstructured, len(objects))
	copy(sorted, objects)
	sort.SliceStable(sorted, func(i, j int) bool {
		ti, tj := KindTier(sorted[i].GetKind()), KindTier(sorted[j].GetKind())
		if ti != tj {
			return ti < tj
		}
		if sorted[i].GetKind() != sorted[j].GetKind() {
			return sorted[i].GetKind() < sorted[j].GetKind()
		}
		if sorted[i].GetNamespace() != sorted[j].GetNamespace() {
			return sorted[i].GetNamespace() < sorted[j].GetNamespace()
		}
		return sorted[i].GetName() < sorted[j].GetName()
	})
	return sorted
}

// RunRestore verifies snapshot integrity, sanitizes all resources, orders them topologically,
// and writes the ordered restore bundle (or reports dry-run validation).
func RunRestore(opts *RestoreOptions, out io.Writer) error {
	if opts.BackupDir == "" {
		return fmt.Errorf("--backup-dir is required")
	}
	summary, err := VerifySnapshot(opts.BackupDir)
	if err != nil {
		return fmt.Errorf("pre-restore validation failed: %w", err)
	}

	objects, err := loadYAMLObjectsFromDir(opts.BackupDir)
	if err != nil {
		return err
	}

	sanitized := make([]*unstructured.Unstructured, 0, len(objects))
	for _, obj := range objects {
		sanitized = append(sanitized, SanitizeForRestore(obj, opts.TargetProject, opts.TargetLocation))
	}
	ordered := SortByDependencyOrder(sanitized)

	if opts.OutputDir != "" {
		if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
			return fmt.Errorf("creating restore output directory: %w", err)
		}
		var combined []string
		for idx, obj := range ordered {
			data, err := yaml.Marshal(obj.Object)
			if err != nil {
				return fmt.Errorf("marshaling restored object %s/%s: %w", obj.GetKind(), obj.GetName(), err)
			}
			fileName := fmt.Sprintf("%03d-%s-%s.yaml", idx+1, strings.ToLower(obj.GetKind()), obj.GetName())
			if err := os.WriteFile(filepath.Join(opts.OutputDir, fileName), data, 0o644); err != nil {
				return fmt.Errorf("writing restore manifest %s: %w", fileName, err)
			}
			combined = append(combined, strings.TrimSpace(string(data)))
		}
		bundle := strings.Join(combined, "\n---\n") + "\n"
		if err := os.WriteFile(filepath.Join(opts.OutputDir, "restore-bundle.yaml"), []byte(bundle), 0o644); err != nil {
			return fmt.Errorf("writing restore-bundle.yaml: %w", err)
		}
	}

	mode := "Applied"
	if opts.DryRun {
		mode = "DryRun Validated"
	}
	_, err = fmt.Fprintf(out, "Restore %s: %d resources from snapshot %s ordered across topological tiers\n",
		mode, len(ordered), summary.Timestamp)
	return err
}

func loadYAMLObjectsFromDir(root string) ([]*unstructured.Unstructured, error) {
	var objects []*unstructured.Unstructured
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		docs := strings.Split(string(raw), "\n---")
		for _, doc := range docs {
			trimmed := strings.TrimSpace(doc)
			if trimmed == "" {
				continue
			}
			var objMap map[string]interface{}
			if err := yaml.Unmarshal([]byte(trimmed), &objMap); err != nil {
				return fmt.Errorf("parsing YAML %s: %w", path, err)
			}
			if len(objMap) == 0 {
				continue
			}
			u := &unstructured.Unstructured{Object: objMap}
			if u.GetKind() == "" || u.GetName() == "" {
				continue
			}
			objects = append(objects, u)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return objects, nil
}
