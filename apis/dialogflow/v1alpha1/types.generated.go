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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Deployment
type Deployment struct {
	// The name of the deployment.
	//  Format:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/deployments/<DeploymentID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.name
	Name *string `json:"name,omitempty"`

	// The name of the flow version for this deployment.
	//  Format:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/versions/<VerionID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.flow_version
	FlowVersion *string `json:"flowVersion,omitempty"`

	// The current state of the deployment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.state
	State *string `json:"state,omitempty"`

	// Result of the deployment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.result
	Result *Deployment_Result `json:"result,omitempty"`

	// Start time of this deployment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End time of this deployment.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Deployment.Result
type Deployment_Result struct {
	// Results of test cases running before the deployment.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/testCases/<TestCaseID>/results/<TestCaseResultID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.Result.deployment_test_results
	DeploymentTestResults []string `json:"deploymentTestResults,omitempty"`

	// The name of the experiment triggered by this deployment.
	//  Format:
	//  projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/environments/<EnvironmentID>/experiments/<ExperimentID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Deployment.Result.experiment
	Experiment *string `json:"experiment,omitempty"`
}
