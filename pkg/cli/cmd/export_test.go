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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"
)

func TestExportParseFlags(t *testing.T) {
	testCases := []struct {
		Name           string
		IAMFormat      string
		Output         string
		OAuth2Token    string
		ResourceFormat string
		Verbose        string
		Error          string
	}{
		{
			Name:           "empty arguments should succeed",
			IAMFormat:      "",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "all arguments should succeed",
			IAMFormat:      "none",
			Output:         "/path/to/my/file",
			OAuth2Token:    "ya29.a0Ae4lvC1p3hHnvX2ECJOSUUBr_JQFDUP7o6pta6leAGg95k58XLQwuISGU9it0bnwmTvN-sFp-6stn42xTwbUxH-FzWbrM1NohqBuhpd6J_0oimF4ZAc5pIb27PJ8mHBFPUZFqSauNu9TQfxgQukKIBm7IJ8xKdCVUxWtJC71u1EF",
			ResourceFormat: "",
			Verbose:        "true",
			Error:          "",
		},
	}
	defaultParams := exportParams
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			args := buildArgs(t, tc.Verbose, "output", tc.Output, "oauth2-token", tc.OAuth2Token,
				"iam-format", tc.IAMFormat, "resource-format", tc.ResourceFormat)
			exportParams = defaultParams
			err := exportCmd.Flags().Parse(args)
			if (err == nil && tc.Error != "") || (err != nil && err.Error() != tc.Error) {
				t.Fatalf("error string mismatch: got '%v', want '%v'", err, tc.Error)
			}
		})
	}
	exportParams = defaultParams
}

func TestExportValidateFlags(t *testing.T) {
	testCases := []struct {
		Name           string
		IAMFormat      string
		Output         string
		OAuth2Token    string
		ResourceFormat string
		Verbose        string
		Error          string
	}{
		{
			Name:           "empty arguments should succeed",
			IAMFormat:      "",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "oauth2-token should succeed",
			IAMFormat:      "",
			Output:         "",
			OAuth2Token:    "ya29.a0Ae4lvC1p3hHnvX2ECJOSUUBr_JQFDUP7o6pta6leAGg95k58XLQwuISGU9it0bnwmTvN-sFp-6stn42xTwbUxH-FzWbrM1NohqBuhpd6J_0oimF4ZAc5pIb27PJ8mHBFPUZFqSauNu9TQfxgQukKIBm7IJ8xKdCVUxWtJC71u1EF",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "iam-format 'partialpolicy' should succeed",
			IAMFormat:      "partialpolicy",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "iam-format 'policy' should succeed",
			IAMFormat:      "policy",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "iam-format 'policymember' should succeed",
			IAMFormat:      "policymember",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "iam-format 'none' should succeed",
			IAMFormat:      "none",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "iam-format with garbage value should error",
			IAMFormat:      "garbage",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "",
			Verbose:        "",
			Error:          "invalid iam-format value of 'garbage': must be one of {partialpolicy, policy, policymember, none}",
		},
		{
			Name:           "resource-format 'krm' should succeed",
			IAMFormat:      "",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "krm",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "resource-format 'hcl' with default iam-format should fail",
			IAMFormat:      "",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Error:          "unsupported value of 'policy' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'none' should succeed",
			IAMFormat:      "none",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Error:          "",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'policy' should fail",
			IAMFormat:      "policy",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Error:          "unsupported value of 'policy' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
		{
			Name:           "resource-format 'hcl' with iam-format 'policymember' should fail",
			IAMFormat:      "policymember",
			Output:         "",
			OAuth2Token:    "",
			ResourceFormat: "hcl",
			Verbose:        "",
			Error:          "unsupported value of 'policymember' for flag 'iam-format': when 'resource-format' is 'hcl' the 'iam-format' flag must have a value of 'none'",
		},
	}
	defaultParams := exportParams
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			args := buildArgs(t, tc.Verbose, "output", tc.Output, "oauth2-token", tc.OAuth2Token,
				"iam-format", tc.IAMFormat, "resource-format", tc.ResourceFormat)
			// reset the package local global variable back to the default values
			exportParams = defaultParams
			err := exportCmd.Flags().Parse(args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			err = parameters.Validate(&exportParams)
			if (err == nil && tc.Error != "") || (err != nil && err.Error() != tc.Error) {
				t.Fatalf("error string mismatch:\ngot\n\t'%v'\nwant\n\t'%v'", err, tc.Error)
			}
		})
	}
	exportParams = defaultParams
}

func TestURLArgument(t *testing.T) {
	testCases := []struct {
		Name  string
		Args  []string
		Error string
	}{
		{
			Name:  "empty URL should fail",
			Args:  []string{},
			Error: "expected a URL argument",
		},
		{
			Name:  "URL with no host should fail",
			Args:  []string{"www.googleapis.com/path/to/my/resource"},
			Error: "invalid URL argument 'www.googleapis.com/path/to/my/resource': host portion is empty",
		},
		{
			Name:  "URL with no path should fail",
			Args:  []string{"https://www.googleapis.com"},
			Error: "invalid URL argument 'https://www.googleapis.com': path portion is empty",
		},
		{
			Name:  "URL with no path, but a trailing slash, should fail",
			Args:  []string{"https://www.googleapis.com/"},
			Error: "invalid URL argument 'https://www.googleapis.com/': path portion is empty",
		},
		{
			Name:  "valid URL without scheme should succeed",
			Args:  []string{"//www.googleapis.com/path/to/my/resource"},
			Error: "",
		},
		{
			Name:  "valid URL with scheme should succeed",
			Args:  []string{"https://www.googleapis.com/path/to/my/resource"},
			Error: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := exportCmd.Args(exportCmd, tc.Args)
			if (err == nil && tc.Error != "") || (err != nil && err.Error() != tc.Error) {
				t.Fatalf("error string mismatch:\ngot\n\t'%v'\nwant\n\t'%v'", err, tc.Error)
			}
		})
	}
}
