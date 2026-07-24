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

func safeReplaceAll(s, old, new string) string {
	if old == "" {
		return s
	}
	return strings.ReplaceAll(s, old, new)
}

func ReplacePlaceholdersInCAIS(caisYAMLStr string, dir string, createBytes []byte, depBytes []byte) string {
	// Normalize placeholders back
	caisYAMLStr = safeReplaceAll(caisYAMLStr, "puxvndidajatl5i", "${uniqueId}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestSharedReservationsProject.Get(), "${TEST_SHARED_RESERVATIONS_PROJECT}")

	// Context-aware/Directory-specific project ID replacements to preserve original template variable names
	hasAlloyDBProject := strings.Contains(string(createBytes), "KCC_ALLOYDB_TEST_PROJECT") || strings.Contains(string(depBytes), "KCC_ALLOYDB_TEST_PROJECT")

	if hasAlloyDBProject {
		caisYAMLStr = safeReplaceAll(caisYAMLStr, "mock-project", "${KCC_ALLOYDB_TEST_PROJECT}")
	} else if strings.Contains(dir, "containerattached") {
		caisYAMLStr = safeReplaceAll(caisYAMLStr, "mock-project", "${KCC_ATTACHED_CLUSTER_TEST_PROJECT}")
	} else if strings.Contains(dir, "resourcemanagerlien") || strings.Contains(dir, "gkehubfeature") {
		caisYAMLStr = safeReplaceAll(caisYAMLStr, "mock-project", "${TEST_DEPENDENT_NO_NETWORK_PROJECT_ID}")
	} else {
		caisYAMLStr = safeReplaceAll(caisYAMLStr, "mock-project", "${projectId}")
	}

	caisYAMLStr = safeReplaceAll(caisYAMLStr, "1234567890", "${projectNumber}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestDependentOrgProjectID.Get(), "${TEST_DEPENDENT_ORG_PROJECT_ID}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestDependentFolderProjectID.Get(), "${TEST_DEPENDENT_FOLDER_PROJECT_ID}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestOrgID.Get(), "${TEST_ORG_ID}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestFolderID.Get(), "${TEST_FOLDER_ID}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestBillingAccountIDForBillingResources.Get(), "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES}")
	caisYAMLStr = safeReplaceAll(caisYAMLStr, testgcp.TestBillingAccountID.Get(), "${TEST_BILLING_ACCOUNT_ID}")

	return caisYAMLStr
}

// NormalizeDynamicIDs replaces dynamic server-assigned IDs before comparing
func NormalizeDynamicIDs(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		// Normalize ComputeFirewallPolicy IDs: locations/global/firewallPolicies/<firewallPolicyId>
		// Since it has a server-generated ID, we normalize it to unknown to match static unit tests consistently.
		if idx := strings.Index(line, "/firewallPolicies/"); idx != -1 {
			if strings.HasPrefix(strings.TrimSpace(line), "- ") {
				lines[i] = "- caisURL: unknown"
			} else {
				lines[i] = "  caisURL: unknown"
			}
		}
		// Normalize BigQuery Connection IDs: locations/.../connections/<connectionId>
		if idx := strings.Index(line, "/connections/"); idx != -1 {
			lines[i] = line[:idx+len("/connections/")]
		}
		// Normalize KMSKeyHandle keyHandles IDs: locations/.../keyHandles/<keyHandleId>
		if idx := strings.Index(line, "/keyHandles/"); idx != -1 {
			start := idx + len("/keyHandles/")
			lines[i] = line[:start] + "${keyHandleID}"
		}
		// Normalize Monitoring Notification Channel numeric IDs
		if idx := strings.Index(line, "/notificationChannels/"); idx != -1 {
			lines[i] = line[:idx+len("/notificationChannels/")]
		}
		// Normalize BillingBudgetsBudget server-generated budget IDs
		if idx := strings.Index(line, "/budgets/"); idx != -1 {
			lines[i] = line[:idx+len("/budgets/")] + "billingbudgetsbudget-${uniqueId}"
		}
		// Normalize Monitoring Alert Policy numeric IDs
		if idx := strings.Index(line, "/alertPolicies/"); idx != -1 {
			lines[i] = line[:idx+len("/alertPolicies/")]
		}
		// Normalize reCAPTCHA Enterprise Key IDs: projects/.../keys/<keyId>
		if idx := strings.Index(line, "/keys/"); idx != -1 && strings.Contains(line, "recaptchaenterprise") {
			lines[i] = line[:idx+len("/keys/")] + "${keyID}"
		}
		// Normalize IAP Brand numeric IDs: projects/.../brands/<brandId>
		if idx := strings.Index(line, "/brands/"); idx != -1 {
			lines[i] = line[:idx+len("/brands/")]
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
		// Normalize VertexAICustomJob server-generated customJob IDs and project number
		if idx := strings.Index(line, "/customJobs/"); idx != -1 {
			lines[i] = line[:idx+len("/customJobs/")] + "vertexaicustomjob-${uniqueId}"
			lines[i] = strings.ReplaceAll(lines[i], "projects/${projectNumber}/", "projects/${projectId}/")
		}
	}
	return strings.Join(lines, "\n")
}
