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
	"context"
	"fmt"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

type configureOptions struct {
	cluster        string
	location       string
	bucket         string
	bucketLocation string
	frequency      string
	project        string
}

func NewConfigureCmd() *cobra.Command {
	options := &configureOptions{}

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure scheduled backups for Config Connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigure(cmd.Context(), options)
		},
	}

	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")
	cmd.Flags().StringVar(&options.bucketLocation, "bucket-location", "us-central1", "GCS bucket location")
	cmd.Flags().StringVar(&options.frequency, "frequency", "daily", "Backup frequency (e.g. daily)")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")

	return cmd
}

func runConfigure(ctx context.Context, options *configureOptions) error {
	if options.project == "" {
		return fmt.Errorf("--project is required")
	}
	if options.bucket == "" {
		return fmt.Errorf("--bucket is required")
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

	data := struct {
		ProjectID      string
		Bucket         string
		BucketLocation string
		Schedule       string
	}{
		ProjectID:      options.project,
		Bucket:         options.bucket,
		BucketLocation: options.bucketLocation,
		Schedule:       schedule,
	}

	tmpl, err := template.New("configure").Parse(configureTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	kubeClient, err := kubecli.NewClient(ctx, kubecli.ClusterOptions{})
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	fmt.Println("Applying backup configuration resources...")

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
		fmt.Printf("- Applied %s/%s (%s)\n", obj.GetNamespace(), obj.GetName(), obj.GetKind())
	}

	fmt.Println("\nBackup configuration successful.")
	return nil
}

const configureTemplate = `
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: {{.Bucket}}
  namespace: cnrm-system
spec:
  location: {{.BucketLocation}}
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMServiceAccount
metadata:
  name: cnrm-backup
  namespace: cnrm-system
spec:
  displayName: Config Connector Backup Service Account
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-wi
  namespace: cnrm-system
spec:
  member: serviceAccount:{{.ProjectID}}.svc.id.goog[cnrm-system/cnrm-backup-manager]
  role: roles/iam.workloadIdentityUser
  resourceRef:
    apiVersion: iam.cnrm.cloud.google.com/v1beta1
    kind: IAMServiceAccount
    name: cnrm-backup
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPolicyMember
metadata:
  name: cnrm-backup-bucket-admin
  namespace: cnrm-system
spec:
  member: serviceAccount:cnrm-backup@{{.ProjectID}}.iam.gserviceaccount.com
  role: roles/storage.admin
  resourceRef:
    apiVersion: storage.cnrm.cloud.google.com/v1beta1
    kind: StorageBucket
    name: {{.Bucket}}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnrm-backup-manager
  namespace: cnrm-system
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
  namespace: cnrm-system
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: cnrm-backup-daily
  namespace: cnrm-system
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
          containers:
          - name: backup
            image: gcr.io/gke-release/cnrm/controller:1.143.0
            command: ["config-connector", "backup", "create", "--bucket", "{{.Bucket}}", "--project", "{{.ProjectID}}"]
          restartPolicy: OnFailure
`
