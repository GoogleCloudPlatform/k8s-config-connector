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

package caistesting

import (
	"strings"

	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
)

func InitializeFakeGCPVariables() {
	testgcp.TestFolderID.Set("123451001")
	testgcp.TestFolder2ID.Set("123451002")
	testgcp.TestOrgID.Set("123450001")
	testgcp.IsolatedTestOrgName.Set("isolated-test-org.example.com")
	testgcp.TestBillingAccountID.Set("123456-777777-000001")
	testgcp.TestBillingAccountIDForBillingResources.Set("123456-777777-000003")
	testgcp.IAMIntegrationTestsOrganizationID.Set("123450002")
	testgcp.IAMIntegrationTestsBillingAccountID.Set("123456-777777-000002")
	testgcp.TestAttachedClusterName.Set("xks-cluster")
	testgcp.TestDependentNoNetworkProjectID.Set("mock-project")
	testgcp.TestDependentOrgProjectID.Set("example-project-01")
	testgcp.TestDependentFolderProjectID.Set("example-project-02")
	testgcp.IdentityPlatformTestProject.Set("kcc-identity-platform")
	testgcp.RecaptchaEnterpriseTestProject.Set("kcc-recaptcha-enterprise")
	testgcp.TestKCCAlloyDBProject.Set("mock-project")
	testgcp.TestKCCAlloyDBProjectNumber.Set("518915279")
	testgcp.TestKCCAttachedClusterProject.Set("mock-project")
	testgcp.TestKCCAttachedClusterPlatformVersion.Set("1.30.0-gke.1")
	testgcp.TestSharedReservationsProject.Set("shared-reservations-project")
}

func ReplacePlaceholdersInCAIS(caisYAMLStr string, dir string, createBytes []byte, depBytes []byte) string {
	// Normalize placeholders back
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "puxvndidajatl5i", "${uniqueId}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestSharedReservationsProject.Get(), "${TEST_SHARED_RESERVATIONS_PROJECT}")

	// Context-aware/Directory-specific project ID replacements to preserve original template variable names
	hasAlloyDBProject := strings.Contains(string(createBytes), "KCC_ALLOYDB_TEST_PROJECT") || strings.Contains(string(depBytes), "KCC_ALLOYDB_TEST_PROJECT")

	if hasAlloyDBProject {
		caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "mock-project", "${KCC_ALLOYDB_TEST_PROJECT}")
	} else if strings.Contains(dir, "containerattached") {
		caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "mock-project", "${KCC_ATTACHED_CLUSTER_TEST_PROJECT}")
	} else if strings.Contains(dir, "resourcemanagerlien") || strings.Contains(dir, "gkehubfeature") {
		caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "mock-project", "${TEST_DEPENDENT_NO_NETWORK_PROJECT_ID}")
	} else {
		caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "mock-project", "${projectId}")
	}

	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, "1234567890", "${projectNumber}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestDependentOrgProjectID.Get(), "${TEST_DEPENDENT_ORG_PROJECT_ID}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestDependentFolderProjectID.Get(), "${TEST_DEPENDENT_FOLDER_PROJECT_ID}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestOrgID.Get(), "${TEST_ORG_ID}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestFolderID.Get(), "${TEST_FOLDER_ID}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestBillingAccountIDForBillingResources.Get(), "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES}")
	caisYAMLStr = strings.ReplaceAll(caisYAMLStr, testgcp.TestBillingAccountID.Get(), "${TEST_BILLING_ACCOUNT_ID}")

	return caisYAMLStr
}

// NormalizeDynamicIDs replaces dynamic server-assigned IDs before comparing
func NormalizeDynamicIDs(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		// Normalize BigQuery Connection IDs: locations/.../connections/<connectionId>
		if idx := strings.Index(line, "/connections/"); idx != -1 {
			lines[i] = line[:idx+len("/connections/")]
		}
		// Normalize Monitoring Notification Channel numeric IDs
		if idx := strings.Index(line, "/notificationChannels/"); idx != -1 {
			lines[i] = line[:idx+len("/notificationChannels/")]
		}
		// Normalize Folder IDs: folders/<folderId>
		if idx := strings.Index(line, "/folders/"); idx != -1 {
			start := idx + len("/folders/")
			end := strings.Index(line[start:], "/")
			if end != -1 {
				lines[i] = line[:start] + "${folderId}" + line[start+end:]
			} else {
				lines[i] = line[:start] + "${folderId}"
			}
		}
	}
	return strings.Join(lines, "\n")
}
