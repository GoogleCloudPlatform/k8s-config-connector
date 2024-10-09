// Copyright 2024 Google LLC
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

package mockdataflow

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/dataflow/v1beta3"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *MockService) StopJob(fqn string) error {
	ctx := context.TODO()

	now := time.Now()

	obj := &pb.Job{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return fmt.Errorf("error getting job %q: %w", fqn, err)
	}

	switch obj.CurrentState {
	case pb.JobState_JOB_STATE_CANCELLING:
		obj.CurrentState = pb.JobState_JOB_STATE_CANCELLED
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_UNKNOWN
	case pb.JobState_JOB_STATE_QUEUED:
		obj.CurrentState = pb.JobState_JOB_STATE_CANCELLED
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_UNKNOWN
	case pb.JobState_JOB_STATE_DRAINING:
		obj.CurrentState = pb.JobState_JOB_STATE_DRAINED
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_UNKNOWN
	default:
		return fmt.Errorf("unexpected state for job, got=%q, expected=CANCELLING or DRAINING: %v", fqn, obj.CurrentState)
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return fmt.Errorf("error updating job %q: %v", fqn, err)
	}

	return nil
}

func (r *MockService) StartJob(fqn string, project *projects.ProjectData, req *pb.LaunchFlexTemplateRequest) error {
	ctx := context.TODO()

	containerSpecGCSPath := req.GetLaunchParameter().GetContainerSpecGcsPath()
	jobKey := lastComponent(containerSpecGCSPath)

	for {
		now := time.Now()

		job := &pb.Job{}
		if err := r.storage.Get(ctx, fqn, job); err != nil {
			return fmt.Errorf("error getting job %q: %w", fqn, err)
		}

		if job.CurrentState == pb.JobState_JOB_STATE_UNKNOWN {
			job.CurrentState = pb.JobState_JOB_STATE_QUEUED
			job.CurrentStateTime = timestamppb.New(now)

			job.Environment = &pb.Environment{
				Dataset: "bigquery.googleapis.com/cloud_dataflow",
			}

			job.PipelineDescription = &pb.PipelineDescription{}

			if jobKey == "File_Format_Conversion" {
				job.Labels = map[string]string{
					"goog-dataflow-provided-template-name": "file_format_conversion",
					"goog-dataflow-provided-template-type": "flex",
				}
			}
			if jobKey == "Cloud_PubSub_to_GCS_Text_Flex" {
				job.Labels = map[string]string{
					"goog-dataflow-provided-template-name":    "cloud_pubsub_to_gcs_text_flex",
					"goog-dataflow-provided-template-type":    "flex",
					"goog-dataflow-provided-template-version": "2024-07-23-00_rc00",
				}
			}

			if err := r.storage.Update(ctx, fqn, job); err != nil {
				return fmt.Errorf("error updating job %q: %v", fqn, err)
			}

			time.Sleep(10 * time.Second)
		}

		if job.CurrentState == pb.JobState_JOB_STATE_QUEUED {
			job.CurrentState = pb.JobState_JOB_STATE_PENDING
			job.CurrentStateTime = timestamppb.New(now)
			job.JobMetadata = &pb.JobMetadata{}
			job.Steps = []*pb.Step{{Name: "placeholder"}}

			switch jobKey {
			case "Cloud_PubSub_to_GCS_Text_Flex":
				job.Type = pb.JobType_JOB_TYPE_STREAMING
			case "PubSub_Avro_to_BigQuery":
				job.Type = pb.JobType_JOB_TYPE_STREAMING
			case "File_Format_Conversion":
				job.Type = pb.JobType_JOB_TYPE_BATCH
			default:
				job.Type = pb.JobType_JOB_TYPE_BATCH
			}

			job.Environment.Experiments = buildExperiments(jobKey)

			if jobKey == "PubSub_Avro_to_BigQuery" {
				job.Labels = map[string]string{
					"goog-dataflow-provided-template-name": "pubsub_avro_to_bigquery",
					"goog-dataflow-provided-template-type": "flex",
				}
			}

			sdkPipelineOptions := buildSDKPipelineOptions(project, req, job, jobKey)
			sdkPipelineOptionsVal, err := structpb.NewStruct(sdkPipelineOptions)
			if err != nil {
				return fmt.Errorf("building structpb for sdkPipelineOptions: %v", err)
			}
			job.Environment.SdkPipelineOptions = sdkPipelineOptionsVal

			job.Environment.ServiceAccountEmail = replacePlaceholders("${projectNumber}-compute@developer.gserviceaccount.com", project, job)
			job.Environment.TempStoragePrefix = replacePlaceholders("storage.googleapis.com/dataflow-staging-us-central1-${projectNumber}/tmp", project, job)
			job.Environment.UserAgent = &structpb.Struct{}
			job.Environment.Version = &structpb.Struct{}
			job.Environment.WorkerPools = []*pb.WorkerPool{{Kind: "placeholder"}}

			switch jobKey {
			case "File_Format_Conversion":
				job.Environment.ShuffleMode = pb.ShuffleMode_SERVICE_BASED
			}

			if err := r.storage.Update(ctx, fqn, job); err != nil {
				return fmt.Errorf("error updating job %q: %v", fqn, err)
			}
			time.Sleep(1 * time.Second)
		}

		if job.CurrentState == pb.JobState_JOB_STATE_PENDING {
			job.CurrentState = pb.JobState_JOB_STATE_RUNNING
			job.CurrentStateTime = timestamppb.New(now)
			job.StageStates = []*pb.ExecutionStageState{{ExecutionStageName: "placeholder"}}

			if err := r.storage.Update(ctx, fqn, job); err != nil {
				return fmt.Errorf("error updating job %q: %v", fqn, err)
			}

			// We have started the job
			return nil
		}
	}
}

func buildExperiments(jobKey string) []string {
	experiments := []string{
		"enable_fnapi_multimap_side_input_bulk_read",
		"enable_compute_default_service_account_org_policy",
		"use_multi_hop_delegation",
		"enable_dataprep_new_billing",
		"disable_baggins_exp",
		"enable_remote_image_ping",
		"enable_billing_v_1_5",
		"enable_worker_memory_cloud_monitoring",
		"shuffle_mode=auto",
		"auto_runner_v2_min_sdk=2.54.0",
		"limit_resizing_by_cpu_util",
		"enable_cloud_permissions_checking",
		"limit_preemptible_worker_pct",
		"use_dataflow_service_account_in_igm",
		"enable_oom_sampler",
		"enable_worker_disk_cloud_monitoring",
		"sideinput_io_metrics",
		"enable_recommendations",
		"primeflex_slow_start_pct=5",
		"auto_google_template_runner_v2",
		"enable_async_job_creation",
		"enable_zonal_outage_aware_routing",
		"enable_worker_cloud_monitoring_exporter",
		"regional_physical_zone_separation_enabled",
		"configure_shuffle_service_addresses_in_control_plane",
		"use_templates_regional_bucket",
		"use_job_admission_controller",
		"enable_always_on_exception_sampling",
		"enable_data_sampling_telemetry",
		"override_controller_service_account",
		"enable_secure_boot",
		"shuffle_service_address_type=DIRECTPATH_WITH_CFE_FALLBACK",
		"enable_throttled_based_rescaling",
		"enable_cmek_org_policy_check",
		"use_e2_for_default_machine_type_worker_regions=africa-south1,europe-north2,europe-southwest1,europe-west10,europe-west12,europe-west8,europe-west9,me-central1,me-central2,me-west1,northamerica-south1,southamerica-west1,us-east10,us-east5,us-east7,us-south1,us-west8",
		"ek_regions=",
		"primeflex_slow_start_seconds=3600",
		"min_sdk_version_to_reject_worker_in_different_region_than_service=2.44.0",
		"delayed_launch",
		"enable_memory_sampler",
		"use_worker_zone_chooser_by_default",
		"auto_high_core_runner_v2",
		"disable_primeflex",
		"disable_runner_v2_reason=java_job_google_template",
		"enable_always_on_exception_sampling",
	}

	if jobKey == "PubSub_Avro_to_BigQuery" {
		experiments = []string{
			"autoscale_windmill_service_by_default",
			"disable_baggins_exp",
			"disable_prime_streaming_engine",
			"disable_runner_v2_reason=java_job_google_template",
			"ek_regions=",
			"enable_always_on_exception_sampling",
			"enable_always_on_exception_sampling",
			"enable_async_job_creation",
			"enable_billing_v_1_5",
			"enable_cloud_permissions_checking",
			"enable_cmek_org_policy_check",
			"enable_compute_default_service_account_org_policy",
			"enable_data_sampling_telemetry",
			"enable_fnapi_multimap_side_input_bulk_read",
			"enable_kms_on_streaming_engine",
			"enable_lightweight_streaming_update",
			"enable_lightweight_streaming_update_worker_utilization_hint",
			"enable_memory_sampler",
			"enable_oom_sampler",
			"enable_recommendations",
			"enable_remote_image_ping",
			"enable_secure_boot",
			"enable_streaming_java_vmr",
			"enable_streaming_pubsub_io_stackdriver_metrics",
			"enable_streaming_scaling",
			"enable_streaming_service_billing",
			"enable_throttled_based_rescaling",
			"enable_worker_cloud_monitoring_exporter",
			"enable_worker_disk_cloud_monitoring",
			"enable_worker_memory_cloud_monitoring",
			"enable_zonal_outage_aware_routing",
			"force_zone_isolation=zi-fail-open",
			"override_controller_service_account",
			"primeflex_slow_start_pct=5",
			"primeflex_slow_start_seconds=3600",
			"streaming_engine_isolation_mode=PER_JOB",
			"use_dataflow_service_account_in_igm",
			"use_e2_for_default_machine_type_worker_regions=africa-south1,europe-north2,europe-southwest1,europe-west10,europe-west12,europe-west8,europe-west9,me-central1,me-central2,me-west1,northamerica-south1,southamerica-west1,us-east10,us-east5,us-east7,us-south1,us-west8",
			"use_job_admission_controller",
			"use_multi_hop_delegation",
			"use_ppm_for_windmill",
			"use_templates_regional_bucket",
			"use_timer_backlog_in_streaming_autoscaling",
			"use_worker_zone_chooser_by_default",
			"use_workflow_job_manager_components_for_streaming",
			"use_workflow_job_manager_components_for_streaming_reload",
		}
	}

	if jobKey == "Cloud_PubSub_to_GCS_Text_Flex" {
		experiments = []string{
			"autoscale_windmill_service_by_default",
			"disable_baggins_exp",
			"disable_prime_streaming_engine",
			"disable_runner_v2_reason=java_job_google_template",
			"ek_regions=",
			"enable_always_on_exception_sampling",
			"enable_always_on_exception_sampling",
			"enable_async_job_creation",
			"enable_billing_v_1_5",
			"enable_cloud_permissions_checking",
			"enable_cmek_org_policy_check",
			"enable_compute_default_service_account_org_policy",
			"enable_data_sampling_telemetry",
			"enable_fnapi_multimap_side_input_bulk_read",
			"enable_kms_on_streaming_engine",
			"enable_lightweight_streaming_update",
			"enable_lightweight_streaming_update_worker_utilization_hint",
			"enable_memory_sampler",
			"enable_oom_sampler",
			"enable_recommendations",
			"enable_remote_image_ping",
			"enable_secure_boot",
			"enable_streaming_java_vmr",
			"enable_streaming_pubsub_io_stackdriver_metrics",
			"enable_streaming_scaling",
			"enable_streaming_service_billing",
			"enable_throttled_based_rescaling",
			"enable_worker_cloud_monitoring_exporter",
			"enable_worker_disk_cloud_monitoring",
			"enable_worker_memory_cloud_monitoring",
			"enable_zonal_outage_aware_routing",
			"force_zone_isolation=zi-fail-open",
			"override_controller_service_account",
			"primeflex_slow_start_pct=5",
			"primeflex_slow_start_seconds=3600",
			"streaming_engine_isolation_mode=PER_JOB",
			"use_dataflow_service_account_in_igm",
			"use_e2_for_default_machine_type_worker_regions=africa-south1,europe-north2,europe-southwest1,europe-west10,europe-west12,europe-west8,europe-west9,me-central1,me-central2,me-west1,northamerica-south1,southamerica-west1,us-east10,us-east5,us-east7,us-south1,us-west8",
			"use_job_admission_controller",
			"use_multi_hop_delegation",
			"use_ppm_for_windmill",
			"use_templates_regional_bucket",
			"use_timer_backlog_in_streaming_autoscaling",
			"use_worker_zone_chooser_by_default",
			"use_workflow_job_manager_components_for_streaming",
			"use_workflow_job_manager_components_for_streaming_reload",
		}
	}
	return experiments
}

func buildSDKPipelineOptions(project *projects.ProjectData, req *pb.LaunchFlexTemplateRequest, job *pb.Job, jobKey string) map[string]any {
	m := map[string]any{
		"display_data": buildDisplayData(job),
		"options":      buildOptions(project, req, job, jobKey),
	}

	if jobKey == "Cloud_PubSub_to_GCS_Text_Flex" {
		m["revision"] = 6
	}

	return m
}

func buildOptions(project *projects.ProjectData, req *pb.LaunchFlexTemplateRequest, job *pb.Job, jobKey string) map[string]any {
	m := map[string]any{
		"apiRootUrl":               "https://dataflow.googleapis.com/",
		"appName":                  "FileFormatConversion",
		"autoscalingAlgorithm":     nil,
		"containsHeaders":          false,
		"credentialFactoryClass":   "org.apache.beam.sdk.extensions.gcp.auth.GcpCredentialFactory",
		"csvFileEncoding":          "UTF-8",
		"csvFormat":                "Default",
		"dataflowEndpoint":         "",
		"dataflowKmsKey":           nil,
		"dataflowServiceOptions":   nil,
		"dataflowWorkerJar":        nil,
		"defaultEnvironmentConfig": nil,
		"defaultEnvironmentType":   nil,
		"delimiter":                ",",
		"diskSizeGb":               0,
		"enableCloudDebugger":      false,
		"enableStreamingEngine":    false,
		"environmentOptions":       nil,
		"experiments": []any{
			"disable_runner_v2_reason=java_job_google_template",
			"enable_always_on_exception_sampling",
		},
		"filesToStage": []any{
			"/template/file-format-conversion/file-format-conversion.jar",
		},
		"gcpTempLocation":          "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
		"gcsPerformanceMetrics":    false,
		"gcsUploadBufferSizeBytes": nil,
		"googleApiTrace":           nil,
		"hotKeyLoggingEnabled":     false,
		"inputFileFormat":          "csv",
		"inputFileSpec":            "gs://config-connector-samples/dataflowflextemplate/numbertest.csv",
		"jobName":                  "dataflowflextemplatejob-${uniqueId}",
		"labels": map[string]any{
			"goog-dataflow-provided-template-name": "file_format_conversion",
			"goog-dataflow-provided-template-type": "flex",
		},
		"logDetailedCsvConversionErrors":    false,
		"maxNumWorkers":                     0,
		"network":                           nil,
		"numShards":                         0,
		"numWorkers":                        0,
		"numberOfWorkerHarnessThreads":      0,
		"optionsId":                         0,
		"outputBucket":                      "gs://storagebucket-${uniqueId}",
		"outputFileFormat":                  "avro",
		"outputFilePrefix":                  "output",
		"overrideWindmillBinary":            nil,
		"pathValidatorClass":                "org.apache.beam.sdk.extensions.gcp.storage.GcsPathValidator",
		"pipelineUrl":                       "gs://dataflow-staging-us-central1-${projectNumber}/staging/pipeline-ehKedGonx8czD0ig9N0xE85y-KvEho1OrQ2jYCtevqg.pb",
		"project":                           "${projectId}",
		"recordJfrOnGcThrashing":            false,
		"region":                            "us-central1",
		"resourceHints":                     []any{},
		"runner":                            "org.apache.beam.runners.dataflow.DataflowRunner",
		"saveProfilesToGcs":                 nil,
		"schema":                            "gs://config-connector-samples/dataflowflextemplate/numbers.avsc",
		"sdkContainerImage":                 nil,
		"sdkHarnessContainerImageOverrides": nil,
		"serviceAccount":                    "${projectNumber}-compute@developer.gserviceaccount.com",
		"stableUniqueNames":                 "WARNING",
		"stagerClass":                       "org.apache.beam.runners.dataflow.util.GcsStager",
		"stagingLocation":                   "gs://dataflow-staging-us-central1-${projectNumber}/staging",
		"streaming":                         false,
		"subnetwork":                        nil,
		"tempLocation":                      "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
		"templateLocation":                  "gs://dataflow-staging-us-central1-${projectNumber}/staging/template_launches/${jobID}/job_object",
		"userAgent":                         "Apache_Beam_SDK_for_Java/2.41.0(JRE_11_environment)",
		"workerDiskType":                    nil,
		"workerHarnessContainerImage":       nil,
		"workerMachineType":                 nil,
		"workerRegion":                      nil,
		"zone":                              nil,
	}

	switch jobKey {
	case "PubSub_Avro_to_BigQuery":
		m = map[string]any{
			"HTTPWriteTimeout":         0,
			"apiRootUrl":               "https://dataflow.googleapis.com/",
			"appName":                  "PubsubAvroToBigQuery",
			"autoscalingAlgorithm":     nil,
			"createDisposition":        "CREATE_NEVER",
			"credentialFactoryClass":   "org.apache.beam.sdk.extensions.gcp.auth.GcpCredentialFactory",
			"dataflowEndpoint":         "",
			"dataflowKmsKey":           nil,
			"dataflowWorkerJar":        nil,
			"defaultEnvironmentConfig": nil,
			"defaultEnvironmentType":   nil,
			"diskSizeGb":               0,
			"enableCloudDebugger":      false,
			"enableStreamingEngine":    false,
			"experiments": []any{
				"enable_always_on_exception_sampling",
				"enable_streaming_java_vmr",
				"disable_runner_v2_reason=java_job_google_template",
			},
			"filesToStage": []any{
				"/usr/local/jre1.8.0-latest/lib/ext/cldrdata.jar",
				"/usr/local/jre1.8.0-latest/lib/ext/jaccess.jar",
				"/usr/local/jre1.8.0-latest/lib/ext/dnsns.jar",
				"/usr/local/jre1.8.0-latest/lib/ext/nashorn.jar",
				"/usr/local/jre1.8.0-latest/lib/ext/localedata.jar",
				"/usr/local/jre1.8.0-latest/lib/ext/jfxrt.jar",
				"/template/pubsub-avro-to-bigquery/pubsub-avro-to-bigquery.jar",
			},
			"gcpTempLocation":          "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
			"gcsPerformanceMetrics":    false,
			"gcsUploadBufferSizeBytes": nil,
			"googleApiTrace":           nil,
			"inputSubscription":        req.GetLaunchParameter().GetParameters()["inputSubscription"],
			"jobName":                  "dataflowflextemplatejob-${uniqueId}",
			"labels": map[string]any{
				"goog-dataflow-provided-template-name": "pubsub_avro_to_bigquery",
				"goog-dataflow-provided-template-type": "flex",
			},
			"maxNumWorkers":                0,
			"maxStreamingBatchSize":        65536,
			"maxStreamingRowsToBatch":      500,
			"network":                      nil,
			"numStreamingKeys":             50,
			"numWorkers":                   0,
			"numberOfWorkerHarnessThreads": 0,
			"optionsId":                    0,
			"outputTableSpec":              "${projectId}:bigquerydataset${uniqueId}.bigquerytable${uniqueId}",
			"outputTopic":                  "projects/${projectId}/topics/pubsubtopic1-${uniqueId}",
			"overrideWindmillBinary":       nil,
			"pathValidatorClass":           "org.apache.beam.sdk.extensions.gcp.storage.GcsPathValidator",
			"pipelineUrl":                  "gs://dataflow-staging-us-central1-${projectNumber}/staging/pipeline-FIMPE_iCnNdL5_-MrVLNGQ.pb",
			"project":                      "${projectId}",
			"region":                       "us-central1",
			"runner":                       "org.apache.beam.runners.dataflow.DataflowRunner",
			"saveProfilesToGcs":            nil,
			"schemaPath":                   "gs://config-connector-samples/dataflowflextemplate/numbers.avsc",
			"serviceAccount":               "${projectNumber}-compute@developer.gserviceaccount.com",
			"stableUniqueNames":            "WARNING",
			"stagerClass":                  "org.apache.beam.runners.dataflow.util.GcsStager",
			"stagingLocation":              "gs://dataflow-staging-us-central1-${projectNumber}/staging",
			"streaming":                    true,
			"subnetwork":                   nil,
			"tempLocation":                 "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
			"templateLocation":             "gs://dataflow-staging-us-central1-${projectNumber}/staging/template_launches/${jobID}/job_object",
			"useGrpcForGcs":                false,
			"userAgent":                    "Apache_Beam_SDK_for_Java/2.20.0(JRE_8_environment)",
			"workerDiskType":               nil,
			"workerHarnessContainerImage":  "gcr.io/cloud-dataflow/v1beta3/IMAGE:beam-2.20.0",
			"workerMachineType":            nil,
			"workerRegion":                 nil,
			"writeDisposition":             "WRITE_APPEND",
			"zone":                         nil,
		}
	}

	if jobKey == "Cloud_PubSub_to_GCS_Text_Flex" {
		m = map[string]any{
			"apiRootUrl":               "https://dataflow.googleapis.com/",
			"appName":                  "PubsubToText",
			"autoscalingAlgorithm":     nil,
			"credentialFactoryClass":   "org.apache.beam.sdk.extensions.gcp.auth.GcpCredentialFactory",
			"dataflowEndpoint":         "",
			"dataflowKmsKey":           nil,
			"dataflowServiceOptions":   nil,
			"dataflowWorkerJar":        nil,
			"dayPattern":               "dd",
			"defaultEnvironmentConfig": nil,
			"defaultEnvironmentType":   nil,
			"diskSizeGb":               0,
			"enableStreamingEngine":    false,
			"environmentOptions":       nil,
			"experiments": []any{
				"enable_always_on_exception_sampling",
				"enable_streaming_java_vmr",
				"disable_runner_v2_reason=java_job_google_template",
			},
			"filesToStage": []any{
				"/template/pubsub-to-text/libs/conscrypt-openjdk-uber-2.5.2.jar",
				"/template/pubsub-to-text/classpath/googlecloud-to-googlecloud-1.0-SNAPSHOT.jar",
				"/template/pubsub-to-text/resources",
			},
			"gcpOauthScopes": []any{
				"https://www.googleapis.com/auth/cloud-platform",
				"https://www.googleapis.com/auth/devstorage.full_control",
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/datastore",
				"https://www.googleapis.com/auth/bigquery",
				"https://www.googleapis.com/auth/bigquery.insertdata",
				"https://www.googleapis.com/auth/pubsub",
			},
			"gcpTempLocation":            "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
			"gcsHttpRequestReadTimeout":  nil,
			"gcsHttpRequestWriteTimeout": nil,
			"gcsPerformanceMetrics":      false,
			"gcsRewriteDataOpBatchLimit": nil,
			"gcsUploadBufferSizeBytes":   1048576,
			"googleApiTrace":             nil,
			"hotKeyLoggingEnabled":       false,
			"hourPattern":                "HH",
			"impersonateServiceAccount":  nil,
			"inputSubscription":          req.GetLaunchParameter().GetParameters()["inputSubscription"],
			"inputTopic":                 nil,
			"jobName":                    "streamingdataflowflextemplatejob2-${uniqueId}",
			"labels": map[string]any{
				"goog-dataflow-provided-template-name":    "cloud_pubsub_to_gcs_text_flex",
				"goog-dataflow-provided-template-type":    "flex",
				"goog-dataflow-provided-template-version": "2024-07-23-00_rc00",
			},
			"maxNumWorkers":                     0,
			"minutePattern":                     "mm",
			"monthPattern":                      "MM",
			"network":                           nil,
			"numShards":                         0,
			"numWorkers":                        0,
			"numberOfWorkerHarnessThreads":      0,
			"optionsId":                         0,
			"outputDirectory":                   req.GetLaunchParameter().GetParameters()["outputDirectory"],
			"outputFilenamePrefix":              "output",
			"outputFilenameSuffix":              "",
			"outputShardTemplate":               "W-P-SS-of-NN",
			"overrideWindmillBinary":            nil,
			"pathValidatorClass":                "org.apache.beam.sdk.extensions.gcp.storage.GcsPathValidator",
			"pipelineUrl":                       "gs://dataflow-staging-us-central1-${projectNumber}/staging/pipeline-kpUNy71B-lVUfqOc0ii7_TzhEanbhDEXimMNOIqU70w.pb",
			"project":                           "${projectId}",
			"recordJfrOnGcThrashing":            false,
			"region":                            "us-central1",
			"resourceHints":                     []any{},
			"runner":                            "org.apache.beam.runners.dataflow.DataflowRunner",
			"saveProfilesToGcs":                 nil,
			"sdkContainerImage":                 nil,
			"sdkHarnessContainerImageOverrides": nil,
			"serviceAccount":                    "${projectNumber}-compute@developer.gserviceaccount.com",
			"stableUniqueNames":                 "WARNING",
			"stagerClass":                       "org.apache.beam.runners.dataflow.util.GcsStager",
			"stagingLocation":                   "gs://dataflow-staging-us-central1-${projectNumber}/staging",
			"streaming":                         true,
			"subnetwork":                        nil,
			"tempLocation":                      "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
			"templateLocation":                  "gs://dataflow-staging-us-central1-${projectNumber}/staging/template_launches/${jobID}/job_object",
			"transformsToOverride":              []any{},
			"updateCompatibilityVersion":        nil,
			"usePublicIps":                      nil,
			"userAgent":                         "Apache_Beam_SDK_for_Java/2.57.0(JRE_11_environment)",
			"userTempLocation":                  nil,
			"windowDuration":                    "5m",
			"workerDiskType":                    nil,
			"workerHarnessContainerImage":       nil,
			"workerMachineType":                 nil,
			"workerRegion":                      nil,
			"yearPattern":                       "YYYY",
			"zone":                              nil,
		}
	}

	if req.GetLaunchParameter().GetUpdate() {
		m["update"] = true
	}

	for k, v := range m {
		if s, ok := v.(string); ok {
			m[k] = replacePlaceholders(s, project, job)
		}
	}
	return m
}

func replacePlaceholders(s string, project *projects.ProjectData, job *pb.Job) string {
	s = strings.ReplaceAll(s, "${projectNumber}", strconv.FormatInt(project.Number, 10))
	s = strings.ReplaceAll(s, "${jobID}", job.GetId())
	return s
}

func buildDisplayData(job *pb.Job) []any {
	values := []map[string]any{
		/*		{
					"key":       "schema",
					"namespace": "com.google.cloud.teleport.v2.transforms.ParquetConverters$ParquetOptions",
					"type":      "STRING",
					"value":     "gs://config-connector-samples/dataflowflextemplate/numbers.avsc",
				},
				{
					"key":       "templateLocation",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "gs://dataflow-staging-us-central1-${projectNumber}/staging/template_launches/${jobID}/job_object",
				},
				{
					"key":       "outputFileFormat",
					"namespace": "com.google.cloud.teleport.v2.templates.FileFormatConversion$FileFormatConversionOptions",
					"type":      "STRING",
					"value":     "avro",
				},
				{
					"key":       "schema",
					"namespace": "com.google.cloud.teleport.v2.transforms.AvroConverters$AvroOptions",
					"type":      "STRING",
					"value":     "gs://config-connector-samples/dataflowflextemplate/numbers.avsc",
				},
				{
					"key":       "appName",
					"namespace": "org.apache.beam.sdk.options.ApplicationNameOptions",
					"type":      "STRING",
					"value":     "FileFormatConversion",
				},
				{
					"key":       "tempLocation",
					"namespace": "org.apache.beam.sdk.options.PipelineOptions",
					"type":      "STRING",
					"value":     "gs://dataflow-staging-us-central1-${projectNumber}/tmp",
				},
				{
					"key":       "region",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "us-central1",
				},
				{
					"key":        "runner",
					"namespace":  "org.apache.beam.sdk.options.PipelineOptions",
					"shortValue": "DataflowRunner",
					"type":       "JAVA_CLASS",
					"value":      "org.apache.beam.runners.dataflow.DataflowRunner",
				},
				{
					"key":       "inputFileFormat",
					"namespace": "com.google.cloud.teleport.v2.templates.FileFormatConversion$FileFormatConversionOptions",
					"type":      "STRING",
					"value":     "csv",
				},
				{
					"key":       "sdkContainerImage",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineWorkerPoolOptions",
					"type":      "STRING",
					"value":     "",
				},
				{
					"key":       "project",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "${projectId}",
				},
				{
					"key":       "inputFileSpec",
					"namespace": "com.google.cloud.teleport.v2.transforms.CsvConverters$CsvPipelineOptions",
					"type":      "STRING",
					"value":     "gs://config-connector-samples/dataflowflextemplate/numbertest.csv",
				},
				{
					"key":       "labels",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "{goog-dataflow-provided-template-name=file_format_conversion, goog-dataflow-provided-template-type=flex}",
				},
				{
					"key":       "outputBucket",
					"namespace": "com.google.cloud.teleport.v2.transforms.ParquetConverters$ParquetOptions",
					"type":      "STRING",
					"value":     "gs://storagebucket-${uniqueId}",
				},
				{
					"key":       "inputFileSpec",
					"namespace": "com.google.cloud.teleport.v2.transforms.AvroConverters$AvroOptions",
					"type":      "STRING",
					"value":     "gs://config-connector-samples/dataflowflextemplate/numbertest.csv",
				},
				{
					"key":       "stagingLocation",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "gs://dataflow-staging-us-central1-${projectNumber}/staging",
				},
				{
					"key":       "pipelineUrl",
					"namespace": "org.apache.beam.runners.dataflow.options.DataflowPipelineOptions",
					"type":      "STRING",
					"value":     "gs://dataflow-staging-us-central1-${projectNumber}/staging/pipeline-ehKedGonx8czD0ig9N0xE85y-KvEho1OrQ2jYCtevqg.pb",
				},
				{
					"key":       "outputBucket",
					"namespace": "com.google.cloud.teleport.v2.transforms.AvroConverters$AvroOptions",
					"type":      "STRING",
					"value":     "gs://storagebucket-${uniqueId}",
				},
				{
					"key":       "jobName",
					"namespace": "org.apache.beam.sdk.options.PipelineOptions",
					"type":      "STRING",
					"value":     "dataflowflextemplatejob-${uniqueId}",
				},
				{
					"key":       "inputFileSpec",
					"namespace": "com.google.cloud.teleport.v2.transforms.ParquetConverters$ParquetOptions",
					"type":      "STRING",
					"value":     "gs://config-connector-samples/dataflowflextemplate/numbertest.csv",
				},
				{
					"key":       "userAgent",
					"namespace": "org.apache.beam.sdk.options.PipelineOptions",
					"type":      "STRING",
					"value":     "Apache_Beam_SDK_for_Java/2.41.0(JRE_11_environment)",
				},
				{
					"key":       "filesToStage",
					"namespace": "org.apache.beam.sdk.options.FileStagingOptions",
					"type":      "STRING",
					"value":     "[/template/file-format-conversion/file-format-conversion.jar]",
				},
		*/
	}

	return asAnySlice(values)
}

// asAnySlice converts []T => []any
// This lets us work with structpb, which only handles []any
func asAnySlice[T any](values []T) []any {
	var ret []any
	for _, v := range values {
		ret = append(ret, v)
	}
	return ret
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}
