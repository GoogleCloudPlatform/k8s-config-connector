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

// +kcc:proto=google.cloud.dataproc.v1.DriverSchedulingConfig
type DriverSchedulingConfig struct {
	// Required. The amount of memory in MB the driver is requesting.
	// +kcc:proto:field=google.cloud.dataproc.v1.DriverSchedulingConfig.memory_mb
	MemoryMb *int32 `json:"memoryMb,omitempty"`

	// Required. The number of vCPUs the driver is requesting.
	// +kcc:proto:field=google.cloud.dataproc.v1.DriverSchedulingConfig.vcores
	Vcores *int32 `json:"vcores,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobPlacement
type JobPlacement struct {
	// Required. The name of the cluster where the job will be submitted.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobPlacement.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Optional. Cluster labels to identify a cluster where the job will be
	//  submitted.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobPlacement.cluster_labels
	ClusterLabels map[string]string `json:"clusterLabels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobReference
type JobReference struct {
	// Optional. The ID of the Google Cloud Platform project that the job belongs
	//  to. If specified, must match the request project ID.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobReference.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Optional. The job ID, which must be unique within the project.
	//
	//  The ID must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), or hyphens (-). The maximum length is 100 characters.
	//
	//  If not specified by the caller, the job ID will be provided by the server.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobReference.job_id
	JobID *string `json:"jobID,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobScheduling
type JobScheduling struct {
	// Optional. Maximum number of times per hour a driver can be restarted as
	//  a result of driver exiting with non-zero code before job is
	//  reported failed.
	//
	//  A job might be reported as thrashing if the driver exits with a non-zero
	//  code four times within a 10-minute window.
	//
	//  Maximum value is 10.
	//
	//  **Note:** This restartable job option is not supported in Dataproc
	//  [workflow templates]
	//  (https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template).
	// +kcc:proto:field=google.cloud.dataproc.v1.JobScheduling.max_failures_per_hour
	MaxFailuresPerHour *int32 `json:"maxFailuresPerHour,omitempty"`

	// Optional. Maximum total number of times a driver can be restarted as a
	//  result of the driver exiting with a non-zero code. After the maximum number
	//  is reached, the job will be reported as failed.
	//
	//  Maximum value is 240.
	//
	//  **Note:** Currently, this restartable job option is
	//  not supported in Dataproc
	//  [workflow
	//  templates](https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template).
	// +kcc:proto:field=google.cloud.dataproc.v1.JobScheduling.max_failures_total
	MaxFailuresTotal *int32 `json:"maxFailuresTotal,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobStatus
type JobStatus struct {
}

// +kcc:proto=google.cloud.dataproc.v1.LoggingConfig
type LoggingConfig struct {

	// TODO: unsupported map type with key string and value enum

}

// +kcc:proto=google.cloud.dataproc.v1.PrestoJob
type PrestoJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. The format in which query output will be displayed. See the
	//  Presto documentation for supported output formats
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Presto client tags to attach to this query
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.client_tags
	ClientTags []string `json:"clientTags,omitempty"`

	// Optional. A mapping of property names to values. Used to set Presto
	//  [session properties](https://prestodb.io/docs/current/sql/set-session.html)
	//  Equivalent to using the --session flag in the Presto CLI
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.QueryList
type QueryList struct {
	// Required. The queries to execute. You do not need to end a query expression
	//  with a semicolon. Multiple queries can be specified in one
	//  string by separating each with a semicolon. Here is an example of a
	//  Dataproc API snippet that uses a QueryList to specify a HiveJob:
	//
	//      "hiveJob": {
	//        "queryList": {
	//          "queries": [
	//            "query1",
	//            "query2",
	//            "query3;query4",
	//          ]
	//        }
	//      }
	// +kcc:proto:field=google.cloud.dataproc.v1.QueryList.queries
	Queries []string `json:"queries,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.TrinoJob
type TrinoJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. The format in which query output will be displayed. See the
	//  Trino documentation for supported output formats
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Trino client tags to attach to this query
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.client_tags
	ClientTags []string `json:"clientTags,omitempty"`

	// Optional. A mapping of property names to values. Used to set Trino
	//  [session properties](https://trino.io/docs/current/sql/set-session.html)
	//  Equivalent to using the --session flag in the Trino CLI
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.YarnApplication
type YarnApplication struct {
	// Required. The application name.
	// +kcc:proto:field=google.cloud.dataproc.v1.YarnApplication.name
	Name *string `json:"name,omitempty"`

	// Required. The application state.
	// +kcc:proto:field=google.cloud.dataproc.v1.YarnApplication.state
	State *string `json:"state,omitempty"`

	// Required. The numerical progress of the application, from 1 to 100.
	// +kcc:proto:field=google.cloud.dataproc.v1.YarnApplication.progress
	Progress *float32 `json:"progress,omitempty"`

	// Optional. The HTTP URL of the ApplicationMaster, HistoryServer, or
	//  TimelineServer that provides application-specific information. The URL uses
	//  the internal hostname, and requires a proxy server for resolution and,
	//  possibly, access.
	// +kcc:proto:field=google.cloud.dataproc.v1.YarnApplication.tracking_url
	TrackingURL *string `json:"trackingURL,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobPlacement
type JobPlacementObservedState struct {
	// Output only. A cluster UUID generated by the Dataproc service when
	//  the job is submitted.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobPlacement.cluster_uuid
	ClusterUuid *string `json:"clusterUuid,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobStatus
type JobStatusObservedState struct {
	// Output only. A state message specifying the overall job state.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobStatus.state
	State *string `json:"state,omitempty"`

	// Optional. Output only. Job state details, such as an error
	//  description if the state is `ERROR`.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobStatus.details
	Details *string `json:"details,omitempty"`

	// Output only. The time when this state was entered.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobStatus.state_start_time
	StateStartTime *string `json:"stateStartTime,omitempty"`

	// Output only. Additional state information, which includes
	//  status reported by the agent.
	// +kcc:proto:field=google.cloud.dataproc.v1.JobStatus.substate
	Substate *string `json:"substate,omitempty"`
}
