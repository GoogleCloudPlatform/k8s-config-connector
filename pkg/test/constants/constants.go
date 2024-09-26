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

// testsconstants contains constants used in tests.
package testconstants

// TestNameSubstringsToSkip contains a list of substrings
// that test names are matched against, to skip.
//
// Typically these are skipped temporarily, and each string
// should include a comment with a ticket to denote the
// temporary issue being addressed, and if possibly a date
// to remove.
var TestNameRegexesToSkip = []string{
	// Uncomment lines below when project-related tests need to be disabled
	// (e.g. due to issues like b/196437940).
	// ".*project(movedfoldertofolder|inorg|infolder).*",
	// ".*(acmfeature|projectorgpolicy).*",
	// TODO(b/220357089): re-enable eventfunction test in long-running CRUD test suite.
	".*(eventfunction).*",
	// TODO(b/215781076): re-enable gke hub tests in long-running CRUD test suite.
	".*(mcifeature).*",
	// TODO(b/228525841): re-enable streamingdataflowjobupdateparameters test.
	".*(streamingdataflowjobupdateparameters).*",
	// TODO(b/267510222): re-enable calendarbudget test when test GCP org configconnector.net is allowlisted.
	".*(calendarbudget).*",
	// Creating Remote Control Plane Clusters is no longer supported.
	// Need to allowlist test project to bypass this: b/359316882.
	".*(edgecontainercluster).*",
	// Edge Network tests require a mocked GCP API.
	".*(edgenetwork).*",
	// Edge container node pool tests require a mocked GCP API.
	".*(edgecontainernodepool).*",
	// Edge container vpn connection tests require a mocked GCP API.
	".*(edgecontainervpnconnection).*",
	// Disable due to b/315013168, covered by mocked GCP test.
	".*(containerattachedcluster).*",
	// Disable due to TF bug https://github.com/hashicorp/terraform-provider-google/issues/16255.
	// We can't specify labels in the create operation, that causes AssertLabelsMatchAndHaveManagedLabel check to fail.
	".*(regionalforwardingrulepsc).*",
	// This test only works with direct controller. Re-enable it when we turn on direct ComputeForwardingRule.
	".*(globalforwardingrulepscgoogleapis).*",
}

// TestNameRegexToSkipForTestCRUD is similar to
// TestNameRegexToSkip, but is specific for the
// TestCRUD tests.
//
// The same recommendations for documentation as
// TestNameRegexesToSkip apply.
var TestNameRegexToSkipForTestCRUD = []string{
	// TODO(b/184899399): Remove the if condition once the name-regex is supported.
	// Disable orgrole test in TestCRUD() to avoid the quota issue.
	// Org-level IAM roles have a hard limit of 300 roles per organization,
	// and the deleted role will consume the quota for another 7 days. Because
	// of that we frequently run into the quota issue which caused presubmit
	// test failures these days.
	// The test cases are selected by GetFilteredSetCover(). After orgrole
	// test is disabled, projectrole test is selected to test IAMRole resource.
	".*orgrole.*",
	// Creating Remote Control Plane Clusters is no longer supported.
	// Need to allowlist test project to bypass this: b/359316882
	".*(edgecontainercluster).*",
	// Edge Network tests require a mocked GCP API.
	".*(edgenetwork).*",
	// Edge container node pool tests require a mocked GCP API.
	".*(edgecontainernodepool).*",
	// Edge container vpn connection tests require a mocked GCP API.
	".*(edgecontainervpnconnection).*",
	// Disable due to b/315013168, covered by mocked GCP test.
	".*(containerattachedcluster).*",
	// Disable due to TF bug https://github.com/hashicorp/terraform-provider-google/issues/16255.
	// We can't specify labels in the create operation, that causes AssertLabelsMatchAndHaveManagedLabel check to fail.
	".*(privateserviceconnectforwardingrule).*",
}
