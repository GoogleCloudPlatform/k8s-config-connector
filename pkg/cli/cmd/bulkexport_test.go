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

package cmd

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	testos "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/os"
)

func TestBulkExportParseFlags(t *testing.T) {
	testCases := []struct {
		Name           string
		IAMFormat      string
		Input          string
		Output         string
		StorageKey     string
		OnError        string
		ProjectID      string
		FolderID       string
		OrganizationID string
		OAuth2Token    string
		ResourceFormat string
		Verbose        string
		Error          string
	}{
		{
			Name:           "empty arguments should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "filename should succeed",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "all arguments should succeed",
			IAMFormat:      "none",
			Input:          "",
			Output:         "",
			StorageKey:     "my-key",
			ProjectID:      "my-project",
			FolderID:       "1234",
			OrganizationID: "5678",
			OAuth2Token:    "ya29.a0Ae4lvC1p3hHnvX2ECJOSUUBr_JQFDUP7o6pta6leAGg95k58XLQwuISGU9it0bnwmTvN-sFp-6stn42xTwbUxH-FzWbrM1NohqBuhpd6J_0oimF4ZAc5pIb27PJ8mHBFPUZFqSauNu9TQfxgQukKIBm7IJ8xKdCVUxWtJC71u1EF",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "non-numeric folder id should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "my-key",
			OnError:        "",
			ProjectID:      "my-project",
			FolderID:       "a",
			OrganizationID: "5678",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "invalid argument \"a\" for \"--folder\" flag: strconv.ParseInt: parsing \"a\": invalid syntax",
		},
		{
			Name:           "non-numeric organization id should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "my-key",
			OnError:        "",
			ProjectID:      "my-project",
			FolderID:       "1234",
			OrganizationID: "b",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "invalid argument \"b\" for \"--organization\" flag: strconv.ParseInt: parsing \"b\": invalid syntax",
		},
	}
	defaultParams := bulkExportParams
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			args := buildArgs(t, tc.Verbose, "input", tc.Input, "output", tc.Output, "storage-key", tc.StorageKey,
				"on-error", tc.OnError, "project", tc.ProjectID, "folder", tc.FolderID, "organization", tc.OrganizationID,
				"oauth2-token", tc.OAuth2Token, "iam-format", tc.IAMFormat, "resource-format", tc.ResourceFormat)
			// by retrieving the Inherited flags we force the cobra library to add them to the entire flaglist
			bulkExportParams = defaultParams
			err := bulkExportCmd.Flags().Parse(args)
			if (err == nil && tc.Error != "") || (err != nil && err.Error() != tc.Error) {
				t.Fatalf("error string mismatch: got '%v', want '%v'", err, tc.Error)
			}
		})
	}
	bulkExportParams = defaultParams
}

func TestBulkExportValidateFlags(t *testing.T) {
	testCases := []struct {
		Name           string
		IAMFormat      string
		Input          string
		Output         string
		StorageKey     string
		OnError        string
		ProjectID      string
		FolderID       string
		OrganizationID string
		OAuth2Token    string
		ResourceFormat string
		Verbose        string
		Stdin          string
		Error          string
	}{
		{
			Name:           "empty arguments should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "no input or export parameters supplied, must supply an asset inventory on 'stdin' or the 'input' parameter or supply one of 'project', 'folder', or 'organization' to perform an export",
		},
		{
			Name:           "storage key with no container / parent resource should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "gs://my-key",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "one of the 'project', 'folder', or 'organization' parameters must be defined to perform an export",
		},
		{
			Name:           "storage key with invalid gcs URI should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "my-key",
			OnError:        "",
			ProjectID:      "my-project-id",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "invalid storage-key value of 'my-key': must be a valid cloud storage URI",
		},
		{
			Name:           "storage key with project id should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "gs://my-key",
			OnError:        "",
			ProjectID:      "my-project-id",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "storage key with folder should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "gs://my-key",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "2345",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "storage key with organization id should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "gs://my-key",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "5678",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "storage key with stdin should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "gs://my-key",
			OnError:        "",
			ProjectID:      "project-id",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "cannot supply input on stdin with the 'storage-key' parameter",
		},
		{
			Name:           "filename with stdin should fail",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "cannot supply input on stdin with the 'input' parameter",
		},
		{
			Name:           "filename with storage-key should fail",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "my-storage-key",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "cannot supply both 'input' and 'storage-key': the parameters are mutually exclusive",
		},
		{
			Name:           "filename with project should fail",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "my-project-id",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "cannot supply both 'input' and 'project': the parameters are mutually exclusive",
		},
		{
			Name:           "filename with folder should fail",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "1234",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "cannot supply both 'input' and 'folder': the parameters are mutually exclusive",
		},
		{
			Name:           "filename with organization should fail",
			IAMFormat:      "",
			Input:          "/tmp/my-file",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "5678",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "cannot supply both 'input' and 'organization': the parameters are mutually exclusive",
		},
		{
			Name:           "only projectid should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "my-project-id",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "only folder should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "1234",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "only organization should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "5678",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "",
			Error:          "",
		},
		{
			Name:           "only stdin should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "on-error 'continue' should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "continue",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "on-error 'halt' should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "halt",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "on-error 'ignore' should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "ignore",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "on-error with garbage value should error",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "garbage",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "invalid on-error value of 'garbage': must be one of {continue, halt, ignore}",
		},
		{
			Name:           "oauth2-token should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "ya29.a0Ae4lvC1p3hHnvX2ECJOSUUBr_JQFDUP7o6pta6leAGg95k58XLQwuISGU9it0bnwmTvN-sFp-6stn42xTwbUxH-FzWbrM1NohqBuhpd6J_0oimF4ZAc5pIb27PJ8mHBFPUZFqSauNu9TQfxgQukKIBm7IJ8xKdCVUxWtJC71u1EF",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "iam-format 'partialpolicy' should succeed",
			IAMFormat:      "partialpolicy",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "iam-format 'policy' should succeed",
			IAMFormat:      "policy",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "iam-format 'policymember' should succeed",
			IAMFormat:      "policymember",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "iam-format 'none' should succeed",
			IAMFormat:      "none",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "iam-format with garbage value should error",
			IAMFormat:      "garbage",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Stdin:          "value",
			Error:          "invalid iam-format value of 'garbage': must be one of {partialpolicy, policy, policymember, none}",
		},
		{
			Name:           "resource-format 'krm' should succeed",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "krm",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "resource-format 'hcl' with default iam-format should fail",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Stdin:          "value",
			Error:          "unsupported value of 'policy' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'none' should succeed",
			IAMFormat:      "none",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Stdin:          "value",
			Error:          "",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'policy' should fail",
			IAMFormat:      "policy",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Stdin:          "value",
			Error:          "unsupported value of 'policy' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'policymember' should fail",
			IAMFormat:      "policymember",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Stdin:          "value",
			Error:          "unsupported value of 'policymember' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
		{
			Name:           "resource-format with garbage value should error",
			IAMFormat:      "",
			Input:          "",
			Output:         "",
			StorageKey:     "",
			OnError:        "",
			ProjectID:      "",
			FolderID:       "",
			OrganizationID: "",
			OAuth2Token:    "",
			ResourceFormat: "garbage",
			Verbose:        "",
			Stdin:          "value",
			Error:          "invalid resource-format value of 'garbage': must be one of {krm, hcl}",
		},
	}
	defaultParams := bulkExportParams
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			args := buildArgs(t, tc.Verbose, "input", tc.Input, "output", tc.Output, "storage-key", tc.StorageKey,
				"on-error", tc.OnError, "project", tc.ProjectID, "folder", tc.FolderID, "organization", tc.OrganizationID,
				"oauth2-token", tc.OAuth2Token, "iam-format", tc.IAMFormat, "resource-format", tc.ResourceFormat)
			// reset the package local global variable back to the default values
			bulkExportParams = defaultParams
			err := bulkExportCmd.Flags().Parse(args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			cleanup, tmpFile := testos.GetStdin(t, tc.Stdin)
			defer cleanup()
			err = parameters.Validate(&bulkExportParams, tmpFile)
			if (err == nil && tc.Error != "") || (err != nil && err.Error() != tc.Error) {
				t.Fatalf("error string mismatch:\ngot\n\t'%v'\nwant\n\t'%v'", err, tc.Error)
			}
		})
	}
	bulkExportParams = defaultParams
}

func buildArgs(t *testing.T, verbose string, args ...string) []string {
	if (len(args) % 2) != 0 {
		t.Fatalf("unexpected number of args, needs to be an even number")
	}
	results := make([]string, 0)
	for i := 0; i < len(args); i += 2 {
		if args[i+1] == "" {
			continue
		}
		results = append(results, fmt.Sprintf("--%v", args[i]))
		results = append(results, args[i+1])
	}
	if verbose != "" {
		results = append(results, verbose)
	}
	return results
}
