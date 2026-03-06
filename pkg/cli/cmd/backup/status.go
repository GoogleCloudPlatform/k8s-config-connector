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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	batchv1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type statusOptions struct {
	kubecli.ClusterOptions
	cluster   string
	location  string
	project   string
	bucket    string
	namespace string
}

func NewStatusCmd() *cobra.Command {
	options := &statusOptions{}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Check the status of recent backup jobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus(cmd, options)
		},
	}

	options.ClusterOptions.AddFlags(cmd)
	cmd.Flags().StringVar(&options.cluster, "cluster", "", "Name of the cluster")
	cmd.Flags().StringVar(&options.location, "location", "", "Region of the cluster")
	cmd.Flags().StringVar(&options.project, "project", "", "GCP project ID")
	cmd.Flags().StringVar(&options.bucket, "bucket", "", "GCS bucket name for backups")
	cmd.Flags().StringVar(&options.namespace, "namespace", "cnrm-system", "Namespace where Config Connector is installed")

	return cmd
}

func runStatus(cmd *cobra.Command, options *statusOptions) error {
	ctx := cmd.Context()
	kubeClient, err := kubecli.NewClient(ctx, options.ClusterOptions)
	if err != nil {
		return fmt.Errorf("creating kubernetes client: %w", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Recent Backup Jobs (Namespace: %s):\n", options.namespace)
	var jobs batchv1.JobList
	if err := kubeClient.List(ctx, &jobs, client.InNamespace(options.namespace)); err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), "Error listing jobs: %v\n", err)
	} else {
		found := false
		for _, job := range jobs.Items {
			// Check for our label or the known name of the CronJob's child jobs
			isBackupJob := job.Labels["app"] == "cnrm-backup"
			if !isBackupJob {
				for _, owner := range job.OwnerReferences {
					if owner.Kind == "CronJob" && owner.Name == "cnrm-backup-daily" {
						isBackupJob = true
						break
					}
				}
			}

			if isBackupJob {
				found = true
				status := "Pending/Running"
				if job.Status.Succeeded > 0 {
					status = "Succeeded"
				} else if job.Status.Failed > 0 {
					status = "Failed"
				}

				completionTime := "In Progress"
				if job.Status.CompletionTime != nil {
					completionTime = job.Status.CompletionTime.Format(time.RFC3339)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "- %s: %s (Completed: %s)\n", job.Name, status, completionTime)
			}
		}
		if !found {
			fmt.Fprintln(cmd.OutOrStdout(), "No backup jobs found.")
		}
	}

	if options.bucket != "" {
		clusterName := options.cluster
		if clusterName == "" {
			clusterName = "default-cluster"
		}
		fmt.Fprintf(cmd.OutOrStdout(), "\nRecent Backups for cluster %s in GCS (gs://%s/%s/):\n", clusterName, options.bucket, clusterName)
		var gcsOptions []option.ClientOption
		if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
			gcsOptions = append(gcsOptions, option.WithHTTPClient(httpClient.(*http.Client)))
		}
		gcsClient, err := storage.NewClient(ctx, gcsOptions...)
		if err != nil {
			return fmt.Errorf("creating GCS client: %w", err)
		}
		defer gcsClient.Close()

		prefix := clusterName + "/"
		it := gcsClient.Bucket(options.bucket).Objects(ctx, &storage.Query{Prefix: prefix, Delimiter: "/"})
		count := 0
		for {
			attrs, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				return fmt.Errorf("iterating GCS objects: %w", err)
			}
			if attrs.Prefix != "" {
				// strip the cluster name prefix for display
				displayPrefix := strings.TrimPrefix(attrs.Prefix, prefix)
				fmt.Fprintf(cmd.OutOrStdout(), "- %s", displayPrefix)

				// Try to read summary.json
				summaryPath := attrs.Prefix + "summary.json"
				rc, err := gcsClient.Bucket(options.bucket).Object(summaryPath).NewReader(ctx)
				if err == nil {
					var stats map[string]int
					if err := json.NewDecoder(rc).Decode(&stats); err == nil {
						total := 0
						for _, count := range stats {
							total += count
						}
						fmt.Fprintf(cmd.OutOrStdout(), " (%d resources)", total)
					}
					rc.Close()
				}
				fmt.Fprintln(cmd.OutOrStdout())
				count++
			}
			if count >= 10 {
				fmt.Fprintln(cmd.OutOrStdout(), "... (limited to 10)")
				break
			}
		}
		if count == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No backup artifacts found for this cluster.")
		}
	}

	return nil
}
