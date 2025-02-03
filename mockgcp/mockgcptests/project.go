// Copyright 2025 Google LLC
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

package mockgcptests

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"golang.org/x/oauth2/google"
)

type GCPProject struct {
	ProjectID     string
	ProjectNumber int64
}

// GetDefaultProject returns the ID of user's configured default GCP project.
func GetDefaultProject(t *testing.T) GCPProject {
	t.Helper()
	ctx := context.TODO()

	projectID := GetDefaultProjectID(t)

	projectNumber, err := GetProjectNumber(ctx, projectID)
	if err != nil {
		t.Fatalf("error getting project number for %q: %v", projectID, err)
	}
	return GCPProject{ProjectID: projectID, ProjectNumber: projectNumber}
}

// GetDefaultProjectID returns the ID of user's configured default GCP project.
func GetDefaultProjectID(t *testing.T) string {
	t.Helper()

	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		s, err := getDefaultProjectID()
		if err != nil {
			t.Fatalf("error getting default project: %v", err)
		}
		projectID = s
	}

	return projectID
}

func GetProjectNumber(ctx context.Context, projectID string) (int64, error) {
	client, err := gcp.NewCloudResourceManagerClient(ctx)
	if err != nil {
		return 0, fmt.Errorf("error creating resource manager client: %w", err)
	}
	project, err := client.Projects.Get(projectID).Do()
	if err != nil {
		return 0, fmt.Errorf("error getting project with id %q: %w", projectID, err)
	}

	return project.ProjectNumber, nil
}

const (
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// GetDefaultProjectID tries to retrieve the default project id through the following:
//  1. Grabbing the project id specified in the application-default GCP credentials on the host machine. This often
//     returns an error, for example when the application-default credentials are expired. Also, the default credentials
//     often do not have the project id set (it's set when the credentials are for a service account).
//  2. If, in step 1 above, there is an error or the project id field is blank, then silently ignore the failure, and
//     fall back to shelling out to gcloud to get the default project id from the local gcloud config.
func getDefaultProjectID() (string, error) {
	creds, err := google.FindDefaultCredentials(context.Background(), CloudPlatformScope)
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
