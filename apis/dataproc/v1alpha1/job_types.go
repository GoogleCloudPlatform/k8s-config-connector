// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataprocJobGVK = GroupVersion.WithKind("DataprocJob")

// DataprocJobSpec defines the desired state of DataprocJob
// +kcc:proto=google.cloud.dataproc.v1.Job
type DataprocJobSpec struct {
	// Optional. The fully qualified reference to the job, which can be used to
	//  obtain the equivalent REST path of the job resource. If this property
	//  is not specified when a job is created, the server generates a
	//  <code>job_id</code>.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.reference
	Reference *JobReference `json:"reference,omitempty"`

	// Required. Job information, including how, when, and where to
	//  run the job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.placement
	Placement *JobPlacement `json:"placement,omitempty"`

	// Optional. Job is a Hadoop job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.hadoop_job
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// Optional. Job is a Spark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_job
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// Optional. Job is a PySpark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.pyspark_job
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Optional. Job is a Hive job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.hive_job
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// Optional. Job is a Pig job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.pig_job
	PigJob *PigJob `json:"pigJob,omitempty"`

	// Optional. Job is a SparkR job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_r_job
	SparkRJob *SparkRJob `json:"sparkRJob,omitempty"`

	// Optional. Job is a SparkSql job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.spark_sql_job
	SparkSQLJob *SparkSQLJob `json:"sparkSQLJob,omitempty"`

	// Optional. Job is a Presto job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.presto_job
	PrestoJob *PrestoJob `json:"prestoJob,omitempty"`

	// Optional. Job is a Trino job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.trino_job
	TrinoJob *TrinoJob `json:"trinoJob,omitempty"`

	// Optional. Job is a Flink job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.flink_job
	FlinkJob *FlinkJob `json:"flinkJob,omitempty"`

	// Optional. The labels to associate with this job.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** can be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with a job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Job scheduling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.scheduling
	Scheduling *JobScheduling `json:"scheduling,omitempty"`

	// Optional. Driver scheduling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_scheduling_config
	DriverSchedulingConfig *DriverSchedulingConfig `json:"driverSchedulingConfig,omitempty"`

	// Required.
	*DataprocJobParent `json:"parent,omitempty"`

	// The DataprocJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type DataprocJobParent struct {

	// Required. The ID of the Google Cloud Platform project that the job belongs to.
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required. The Dataproc region in which to handle the request.
	Region string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.FlinkJob
type FlinkJob struct {
	// The HCFS URI of the jar file that contains the main class.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the class
	//  must be in the default CLASSPATH or specified in
	//  [jarFileUris][google.cloud.dataproc.v1.FlinkJob.jar_file_uris].
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision
	//  might occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Flink driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS URI of the savepoint, which contains the last saved progress
	//  for starting the current job.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.savepoint_uri
	SavepointURI *string `json:"savepointURI,omitempty"`

	// Optional. A mapping of property names to values, used to configure Flink.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  `/etc/flink/conf/flink-defaults.conf` and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.HadoopJob
type HadoopJob struct {
	// The HCFS URI of the jar file containing the main class.
	//  Examples:
	//      'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar'
	//      'hdfs:/tmp/test-samples/custom-wordcount.jar'
	//      'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file containing the class
	//  must be in the default CLASSPATH or specified in `jar_file_uris`.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not
	//  include arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as
	//  job properties, since a collision might occur that causes an incorrect job
	//  submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.args
	Args []string `json:"args,omitempty"`

	// Optional. Jar file URIs to add to the CLASSPATHs of the
	//  Hadoop driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS (Hadoop Compatible Filesystem) URIs of files to be copied
	//  to the working directory of Hadoop drivers and distributed tasks. Useful
	//  for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted in the working directory of
	//  Hadoop drivers and tasks. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, or .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`

	// Optional. A mapping of property names to values, used to configure Hadoop.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site` and
	//  classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.HiveJob
type HiveJob struct {
	// The HCFS URI of the script that contains Hive queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the
	//  Hive command: `SET name="value";`).
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names and values, used to configure Hive.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site.xml`,
	//  /etc/hive/conf/hive-site.xml, and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATH of the
	//  Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes
	//  and UDFs.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PigJob
type PigJob struct {
	// The HCFS URI of the script that contains the Pig queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the Pig
	//  command: `name=[value]`).
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names to values, used to configure Pig.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site.xml`,
	//  /etc/pig/conf/pig.properties, and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATH of
	//  the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PySparkJob
type PySparkJob struct {
	// Required. The HCFS URI of the main Python file to use as the driver. Must
	//  be a .py file.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.main_python_file_uri
	MainPythonFileURI *string `json:"mainPythonFileURI,omitempty"`

	// Optional. The arguments to pass to the driver.  Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS file URIs of Python files to pass to the PySpark
	//  framework. Supported file types: .py, .egg, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.python_file_uris
	PythonFileURIs []string `json:"pythonFileURIs,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Python driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`

	// Optional. A mapping of property names to values, used to configure PySpark.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkJob
type SparkJob struct {
	// The HCFS URI of the jar file that contains the main class.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the class
	//  must be in the default CLASSPATH or specified in
	//  SparkJob.jar_file_uris.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Spark driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`

	// Optional. A mapping of property names to values, used to configure Spark.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkRJob
type SparkRJob struct {
	// Required. The HCFS URI of the main R file to use as the driver.
	//  Must be a .R file.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.main_r_file_uri
	MainRFileURI *string `json:"mainRFileURI,omitempty"`

	// Optional. The arguments to pass to the driver.  Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.file_uris
	FileURIs []string `json:"fileURIs,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.archive_uris
	ArchiveURIs []string `json:"archiveURIs,omitempty"`

	// Optional. A mapping of property names to values, used to configure SparkR.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkSqlJob
type SparkSQLJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the
	//  Spark SQL command: SET `name="value";`).
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names to values, used to configure
	//  Spark SQL's SparkConf. Properties that conflict with values set by the
	//  Dataproc API might be overwritten.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.jar_file_uris
	JarFileURIs []string `json:"jarFileURIs,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// DataprocJobStatus defines the config connector machine state of DataprocJob
type DataprocJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocJobObservedState `json:"observedState,omitempty"`
}

// DataprocJobObservedState is the state of the DataprocJob resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataproc.v1.Job
type DataprocJobObservedState struct {
	// Required. Job information, including how, when, and where to
	//  run the job.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.placement
	Placement *JobPlacementObservedState `json:"placement,omitempty"`

	// Output only. The job status. Additional application-specific
	//  status information might be contained in the <code>type_job</code>
	//  and <code>yarn_applications</code> fields.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.status
	Status *JobStatusObservedState `json:"status,omitempty"`

	// Output only. The previous job status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.status_history
	StatusHistory []JobStatusObservedState `json:"statusHistory,omitempty"`

	// Output only. The collection of YARN applications spun up by this job.
	//
	//  **Beta** Feature: This report is available for testing purposes only. It
	//  might be changed before final release.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.yarn_applications
	YarnApplications []YarnApplication `json:"yarnApplications,omitempty"`

	// Output only. A URI pointing to the location of the stdout of the job's
	//  driver program.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_output_resource_uri
	DriverOutputResourceURI *string `json:"driverOutputResourceURI,omitempty"`

	// Output only. If present, the location of miscellaneous control files
	//  which can be used as part of job setup and handling. If not present,
	//  control files might be placed in the same location as `driver_output_uri`.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.driver_control_files_uri
	DriverControlFilesURI *string `json:"driverControlFilesURI,omitempty"`

	// Output only. A UUID that uniquely identifies a job within the project
	//  over time. This is in contrast to a user-settable reference.job_id that
	//  might be reused over time.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.job_uuid
	JobUUid *string `json:"jobUUid,omitempty"`

	// Output only. Indicates whether the job is completed. If the value is
	//  `false`, the job is still in progress. If `true`, the job is completed, and
	//  `status.state` field will indicate if it was successful, failed,
	//  or cancelled.
	// +kcc:proto:field=google.cloud.dataproc.v1.Job.done
	Done *bool `json:"done,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocjob;gcpdataprocjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocJob is the Schema for the DataprocJob API
// +k8s:openapi-gen=true
type DataprocJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocJobSpec   `json:"spec,omitempty"`
	Status DataprocJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocJobList contains a list of DataprocJob
type DataprocJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocJob{}, &DataprocJobList{})
}

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
