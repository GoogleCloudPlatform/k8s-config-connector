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

package lint

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestE2ESuiteServiceOverrides(t *testing.T) {
	e2eSuitePath := filepath.Join("..", "..", "dev", "ci", "periodics", "e2e-suite")

	testCases := []struct {
		service              string
		expectedExtraService string
		expectedOutputSubstr string
	}{
		{
			service:              "asset",
			expectedExtraService: "cloudasset.googleapis.com",
		},
		{
			service:              "billing",
			expectedExtraService: "cloudbilling.googleapis.com",
		},
		{
			service:              "cloudquota",
			expectedExtraService: "cloudquotas.googleapis.com",
		},
		{
			service:              "vertexai",
			expectedExtraService: "aiplatform.googleapis.com",
		},
		{
			service:              "colab",
			expectedExtraService: "aiplatform.googleapis.com compute.googleapis.com dataform.googleapis.com notebooks.googleapis.com",
		},
		{
			service:              "identityplatform",
			expectedExtraService: "identitytoolkit.googleapis.com",
		},
		{
			service:              "cloudiot",
			expectedOutputSubstr: "Cloud IoT product was shut down in August 2023",
		},
		{
			service:              "pubsub",
			expectedExtraService: "pubsub.googleapis.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.service, func(t *testing.T) {
			cmd := exec.Command(e2eSuitePath, tc.service)
			cmd.Env = append(os.Environ(), "E2E_SUITE_DRY_RUN=1")
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("e2e-suite %s failed: %v\nOutput: %s", tc.service, err, string(out))
			}

			outStr := string(out)
			if tc.expectedOutputSubstr != "" {
				if !strings.Contains(outStr, tc.expectedOutputSubstr) {
					t.Errorf("e2e-suite %s output missing expected substring %q, got: %s", tc.service, tc.expectedOutputSubstr, outStr)
				}
			}

			if tc.expectedExtraService != "" {
				expectedLine := "EXTRA_GCP_SERVICES=" + tc.expectedExtraService
				if !strings.Contains(outStr, expectedLine) {
					t.Errorf("e2e-suite %s output missing expected %q, got: %s", tc.service, expectedLine, outStr)
				}
			}
		})
	}
}
