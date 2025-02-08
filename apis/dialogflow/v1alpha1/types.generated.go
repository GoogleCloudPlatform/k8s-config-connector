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

package v1alpha1


// +kcc:proto=google.cloud.dialogflow.cx.v3.ContinuousTestResult
type ContinuousTestResult struct {
	// The resource name for the continuous test result. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/continuousTestResults/<ContinuousTestResultID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ContinuousTestResult.name
	Name *string `json:"name,omitempty"`

	// The result of this continuous test run, i.e. whether all the tests in this
	//  continuous test run pass or not.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ContinuousTestResult.result
	Result *string `json:"result,omitempty"`

	// A list of individual test case results names in this continuous test run.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ContinuousTestResult.test_case_results
	TestCaseResults []string `json:"testCaseResults,omitempty"`

	// Time when the continuous testing run starts.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ContinuousTestResult.run_time
	RunTime *string `json:"runTime,omitempty"`
}
