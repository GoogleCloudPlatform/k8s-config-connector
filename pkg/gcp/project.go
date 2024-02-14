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
	"strings"

	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

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
