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

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
)

type GCPProject struct {
	ProjectID      string
	ProjectNumber  int64
	OrganizationID string
}

// GetDefaultProject returns the ID of user's configured default GCP project.
func GetDefaultProject(t *testing.T) GCPProject {
	t.Helper()
	ctx := context.TODO()

	projectID := GetDefaultProjectID(t)

	projectInfo, err := GetProjectInfo(ctx, projectID)
	if err != nil {
		t.Fatalf("error getting project number for %q: %v", projectID, err)
	}

	ancestry, err := GetProjectAncestry(ctx, projectID)
	if err != nil {
		t.Fatalf("error getting project ancestry for %q: %v", projectID, err)
	}

	var organizationID string
	for _, ancestor := range ancestry.Ancestor {
		if ancestor.ResourceId.Type == "organization" {
			organizationID = strings.TrimPrefix(ancestor.ResourceId.Id, "organizations/")
			break
		}
	}

	return GCPProject{ProjectID: projectID, ProjectNumber: projectInfo.ProjectNumber, OrganizationID: organizationID}
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

// NewCloudResourceManagerClient returns a GCP Cloud Resource Manager service.
func NewCloudResourceManagerClient(ctx context.Context) (*cloudresourcemanager.Service, error) {
	client, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return nil, err
	}
	// client.UserAgent = KCCUserAgent()
	return client, nil
}

// GetProjectInfo returns the full project for the given project ID.
func GetProjectInfo(ctx context.Context, projectID string) (*cloudresourcemanager.Project, error) {
	client, err := NewCloudResourceManagerClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating resource manager client: %w", err)
	}
	project, err := client.Projects.Get(projectID).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error getting project with id %q: %w", projectID, err)
	}

	return project, nil
}

// GetProjectAncestry returns the project ancestry for the given project ID.
func GetProjectAncestry(ctx context.Context, projectID string) (*cloudresourcemanager.GetAncestryResponse, error) {
	client, err := NewCloudResourceManagerClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating resource manager client: %w", err)
	}
	req := &cloudresourcemanager.GetAncestryRequest{}
	ancestry, err := client.Projects.GetAncestry(projectID, req).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error getting project ancestry for project %q: %w", projectID, err)
	}

	return ancestry, nil
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
