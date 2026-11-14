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

package gcp

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var (
	partitionedEmailDomainRegex  = regexp.MustCompile(`@([a-z0-9-]+):([a-z0-9-]+)\.iam\.gserviceaccount\.com`)
	partitionedWorkloadPoolRegex = regexp.MustCompile(`\b([a-z0-9-]+):([a-z0-9-]+)\.svc\.id\.goog\b`)
)

// SplitPartitionedProjectID splits a partitioned project ID ("partition:project")
// into its partition prefix and project name. For standard commercial projects,
// partitioned is false and project is the original projectID.
func SplitPartitionedProjectID(projectID string) (partition, project string, partitioned bool) {
	p, proj, ok := strings.Cut(projectID, ":")
	if !ok || p == "" || proj == "" {
		return "", projectID, false
	}
	return p, proj, true
}

// FormatProjectDomainPrefix formats a project ID for use in DNS/RFC 5322 domain labels.
// Partitioned project IDs ("partition:project") are formatted using reversed dot-syntax
// ("project.partition") because colons are prohibited in RFC 5322 email domains and DNS names.
func FormatProjectDomainPrefix(projectID string) string {
	if partition, project, ok := SplitPartitionedProjectID(projectID); ok {
		return fmt.Sprintf("%s.%s", project, partition)
	}
	return projectID
}

// FormatServiceAccountEmail constructs the RFC 5322-compliant IAM service account email
// for a given project ID and service account short name.
func FormatServiceAccountEmail(projectID, saName string) string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", saName, FormatProjectDomainPrefix(projectID))
}

// FormatWorkloadPool constructs the GKE Workload Identity pool domain for a given project ID.
func FormatWorkloadPool(projectID string) string {
	return fmt.Sprintf("%s.svc.id.goog", FormatProjectDomainPrefix(projectID))
}

// NormalizePartitionedDomains rewrites colon-delimited partitioned project IDs inside
// IAM service account email domains (@partition:project.iam.gserviceaccount.com) and
// GKE Workload Identity pools (partition:project.svc.id.goog) to their reversed dot-syntax
// equivalents (@project.partition.iam.gserviceaccount.com and project.partition.svc.id.goog).
func NormalizePartitionedDomains(input string) string {
	if !strings.Contains(input, ":") {
		return input
	}
	out := partitionedEmailDomainRegex.ReplaceAllString(input, "@$2.$1.iam.gserviceaccount.com")
	out = partitionedWorkloadPoolRegex.ReplaceAllString(out, "$2.$1.svc.id.goog")
	return out
}

// GetDefaultProjectID tries to retrieve the default project id through the following:
//  1. Grabbing the project id specified in the application-default GCP credentials on the host machine. This often
//     returns an error, for example when the application-default credentials are expired. Also, the default credentials
//     often do not have the project id set (it's set when the credentials are for a service account).
//  2. If, in step 1 above, there is an error or the project id field is blank, then silently ignore the failure, and
//     fall back to shelling out to gcloud to get the default project id from the local gcloud config.
func GetDefaultProjectID() (string, error) {
	creds, err := google.FindDefaultCredentials(context.Background(), sqladmin.CloudPlatformScope)
	if err == nil && creds.ProjectID != "" {
		return creds.ProjectID, nil
	}
	return getGCloudDefaultProjectID()
}

func getGCloudDefaultProjectID() (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", "project")
	bytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing command '%v': %w'", cmd, err)
	}
	value := string(bytes)
	if value == "" {
		return "", fmt.Errorf("error getting default project: gcloud config value for 'project' is empty")
	}
	return strings.TrimSpace(string(bytes)), nil
}
