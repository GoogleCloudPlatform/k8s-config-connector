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
	"bytes"
	"fmt"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/GoogleCloudPlatform/k8s-config-connector/version"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

type configureOptions struct {
	kubecli.ClusterOptions
	cluster          string
	location         string
	bucket           string
	bucketLocation   string
	frequency        string
	project          string
	clusterProject   string
	namespace        string
	dualRegion       string
	turboReplication bool
	versioning       bool
	retentionDays    int
	lockRetention    bool
	autopilot            bool
	replicaBucket        string
	replicaLocation      string
	dryRun               bool
	includeClusterBackup bool
	gkeBackupPlan        string
}

func NewConfigureCmd() *cobra.Command {
	options := &configureOptions{
		versioning: true,
	}

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure scheduled backups for Config Connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigure(cmd, options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")
	cmd.Flags().StringVar(&options.bucketLocation, "bucket-location", "us-central1", "GCS bucket location")
	cmd.Flags().StringVar(&options.frequency, "frequency", "daily", "Backup frequency (e.g. daily)")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID where the backup bucket/SA reside")
	cmd.Flags().StringVar(&options.clusterProject, "cluster-project", "", "GCP project ID where the cluster resides (for Workload Identity). Defaults to --project if not specified.")
	cmd.Flags().StringVar(&options.namespace, "namespace", "cnrm-system", "Namespace where Config Connector is installed")
	cmd.Flags().StringVar(&options.dualRegion, "dual-region", "", "Dual-region bucket pair (e.g. nam4, eur4, asia1)")
	cmd.Flags().BoolVar(&options.turboReplication, "turbo-replication", false, "Enable GCS Turbo Replication (RPO < 15m) for dual-region buckets")
	cmd.Flags().BoolVar(&options.versioning, "versioning", true, "Enable object versioning on backup bucket")
	cmd.Flags().IntVar(&options.retentionDays, "retention-days", 0, "Retention period in days for WORM compliance")
	cmd.Flags().BoolVar(&options.lockRetention, "lock-retention", false, "Permanently lock the bucket retention policy (WORM compliance)")
	cmd.Flags().BoolVar(&options.autopilot, "autopilot", false, "Configure CronJob with GKE Autopilot-compliant securityContext and resources")
	cmd.Flags().StringVar(&options.replicaBucket, "replica-bucket", "", "Secondary replica GCS bucket for cross-region disaster recovery")
	cmd.Flags().StringVar(&options.replicaLocation, "replica-location", "us-east1", "GCS location for secondary replica bucket")
	cmd.Flags().BoolVar(&options.dryRun, "dry-run", false, "Output the generated Kubernetes manifests without applying them to the cluster")
	cmd.Flags().BoolVar(&options.includeClusterBackup, "include-cluster-backup", false, "Configure Backup for GKE (gkebackup.googleapis.com) to backup cluster workloads and persistent volumes")
	cmd.Flags().StringVar(&options.gkeBackupPlan, "gke-backup-plan", "", "Name of the GKE BackupPlan to configure (defaults to <cluster>-backup-plan)")

	return cmd
}

func runConfigure(cmd *cobra.Command, options *configureOptions) error {
	ctx := cmd.Context()
	if options.project == "" {
		return fmt.Errorf("--project is required")
	}
	if options.bucket == "" {
		return fmt.Errorf("--bucket is required")
	}

	clusterProject := options.clusterProject
	if clusterProject == "" {
		clusterProject = options.project
	}

	schedule := "0 0 * * *"
	switch options.frequency {
	case "daily":
		schedule = "0 0 * * *"
	case "weekly":
		schedule = "0 0 * * 0"
	case "hourly":
		schedule = "0 * * * *"
	default:
		// assume it's a cron expression if it contains spaces
		schedule = options.frequency
	}

	bucketLoc := options.bucketLocation
	if options.dualRegion != "" {
		bucketLoc = options.dualRegion
	}

	retentionSeconds := 0
	if options.retentionDays > 0 {
		retentionSeconds = options.retentionDays * 86400
	}

	data := struct {
		ProjectID              string
		ClusterProjectID       string
		Cluster                string
		Bucket                 string
		BucketLocation         string
		Schedule               string
		Namespace              string
		Version                string
		DualRegion             string
		TurboReplication       bool
		Versioning             bool
		RetentionPeriodSeconds int
		LockRetention          bool
		Autopilot              bool
		ReplicaBucket          string
		ReplicaLocation        string
		IncludeClusterBackup   bool
		GKEBackupPlan          string
	}{
		ProjectID:              options.project,
		ClusterProjectID:       clusterProject,
		Cluster:                options.cluster,
		Bucket:                 options.bucket,
		BucketLocation:         bucketLoc,
		Schedule:               schedule,
		Namespace:              options.namespace,
		Version:                version.GetVersion(),
		DualRegion:             options.dualRegion,
		TurboReplication:       options.turboReplication,
		Versioning:             options.versioning,
		RetentionPeriodSeconds: retentionSeconds,
		LockRetention:          options.lockRetention,
		Autopilot:              options.autopilot,
		ReplicaBucket:          options.replicaBucket,
		ReplicaLocation:        options.replicaLocation,
		IncludeClusterBackup:   options.includeClusterBackup,
		GKEBackupPlan:          options.gkeBackupPlan,
	}

	tmpl, err := template.New("configure").Parse(configureTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	if options.dryRun {
		fmt.Fprint(cmd.OutOrStdout(), buf.String())
		return nil
	}

	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), "Applying backup configuration resources...")

	parts := bytes.Split(buf.Bytes(), []byte("\n---\n"))
	for _, part := range parts {
		if len(bytes.TrimSpace(part)) == 0 {
			continue
		}

		obj := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(part, &obj.Object); err != nil {
			return fmt.Errorf("unmarshaling YAML part: %w", err)
		}

		if err := kubeClient.Patch(ctx, obj, client.Apply, client.FieldOwner("config-connector-backup"), client.ForceOwnership); err != nil {
			return fmt.Errorf("applying resource %s/%s (%s): %w", obj.GetNamespace(), obj.GetName(), obj.GetKind(), err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "- Applied %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
	}

	fmt.Fprintln(cmd.OutOrStdout(), "\nBackup configuration successful.")
	return nil
}

const configureTemplate = `
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: {{.Bucket}}
  namespace: {{.Namespace}}
spec:
  location: {{.BucketLocation}}
  uniformBucketLevelAccess: true
  publicAccessPrevention: enforced
{{- if .Versioning}}
  versioning:
    enabled: true
{{- end}}
{{- if .RetentionPeriodSeconds}}
  retentionPolicy:
    retentionPeriod: {{.RetentionPeriodSeconds}}
    {{- if .LockRetention}}
    isLocked: true
    {{- end}}
{{- end}}
{{- if .ReplicaBucket}}
---
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: {{.ReplicaBucket}}
  namespace: {{.Namespace}}
spec:
  location: {{.ReplicaLocation}}
  uniformBucketLevelAccess: true
  publicAccessPrevention: enforced
{{- if .Versioning}}
  versioning:
    enabled: true
{{- end}}
{{- if .RetentionPeriodSeconds}}
  retentionPolicy:
    retentionPeriod: {{.RetentionPeriodSeconds}}
    {{- if .LockRetention}}
    isLocked: true
    {{- end}}
{{- end}}
{{- end}}
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  name: cnrm-backup
  namespace: {{.Namespace}}
spec:
  displayName: Config Connector Backup Service Account
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-wi
  namespace: {{.Namespace}}
spec:
  member: serviceAccount:{{.ClusterProjectID}}.svc.id.goog[{{.Namespace}}/cnrm-backup-manager]
  role: roles/iam.workloadIdentityUser
  resourceRef:
    apiVersion: iam.cnrm.cloud.google.com/v1beta1
    kind: IAMServiceAccount
    name: cnrm-backup
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-bucket-creator
  namespace: {{.Namespace}}
spec:
  member: serviceAccount:cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
  role: roles/storage.objectCreator
  resourceRef:
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    name: {{.Bucket}}
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-bucket-viewer
  namespace: {{.Namespace}}
spec:
  member: serviceAccount:cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
  role: roles/storage.objectViewer
  resourceRef:
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    name: {{.Bucket}}
{{- if .ReplicaBucket}}
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-replica-creator
  namespace: {{.Namespace}}
spec:
  member: serviceAccount:cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
  role: roles/storage.objectCreator
  resourceRef:
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    name: {{.ReplicaBucket}}
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-replica-viewer
  namespace: {{.Namespace}}
spec:
  member: serviceAccount:cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
  role: roles/storage.objectViewer
  resourceRef:
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    name: {{.ReplicaBucket}}
{{- end}}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnrm-backup-manager
  namespace: {{.Namespace}}
  annotations:
    iam.gke.io/gcp-service-account: cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cnrm-backup-reader
rules:
- apiGroups: ["*.cnrm.cloud.google.com"]
  resources: ["*"]
  verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cnrm-backup-reader-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cnrm-backup-reader
subjects:
- kind: ServiceAccount
  name: cnrm-backup-manager
  namespace: {{.Namespace}}
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: cnrm-backup-daily
  namespace: {{.Namespace}}
spec:
  schedule: "{{.Schedule}}"
  jobTemplate:
    metadata:
      labels:
        app: cnrm-backup
    spec:
      template:
        metadata:
          labels:
            app: cnrm-backup
        spec:
          serviceAccountName: cnrm-backup-manager
          restartPolicy: OnFailure
          securityContext:
            runAsNonRoot: true
            runAsUser: 65534
            runAsGroup: 65534
            fsGroup: 65534
            seccompProfile:
              type: RuntimeDefault
          containers:
          - name: backup
            image: gcr.io/gke-release/cnrm/config-connector-cli:{{.Version}}
            command:
            - config-connector
            - backup
            - create
            - --bucket
            - {{.Bucket}}
            - --project
            - {{.ProjectID}}
            - --namespace
            - {{.Namespace}}
            {{- if .Cluster}}
            - --cluster
            - {{.Cluster}}
            {{- end}}
            {{- if .ReplicaBucket}}
            - --replica-bucket
            - {{.ReplicaBucket}}
            {{- end}}
            {{- if .IncludeClusterBackup}}
            - --include-cluster-backup
            {{- if .GKEBackupPlan}}
            - --gke-backup-plan
            - {{.GKEBackupPlan}}
            {{- end}}
            {{- end}}
            securityContext:
              allowPrivilegeEscalation: false
              capabilities:
                drop:
                - ALL
            resources:
              requests:
                cpu: 250m
                memory: 512Mi
                ephemeral-storage: 1Gi
              limits:
                cpu: 1000m
                memory: 1Gi
                ephemeral-storage: 2Gi
`
