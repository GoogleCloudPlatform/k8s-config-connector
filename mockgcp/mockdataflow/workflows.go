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
	case pb.JobState_JOB_STATE_DRAINING:
		obj.CurrentState = pb.JobState_JOB_STATE_DRAINED
		obj.CurrentStateTime = timestamppb.New(now)
		obj.RequestedState = pb.JobState_JOB_STATE_UNKNOWN
	default:
		return fmt.Errorf("unexpected state for job, got=%q, expected=JobState_JOB_STATE_CANCELLING: %v", fqn, obj.CurrentState)
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return fmt.Errorf("error updating job %q: %v", fqn, err)
	}

	return nil
}

func (r *MockService) StartJob(fqn string, project *projects.ProjectData, req *pb.LaunchFlexTemplateRequest) error {
	ctx := context.TODO()

	{
		now := time.Now()

		job := &pb.Job{}
		if err := r.storage.Get(ctx, fqn, job); err != nil {
			return fmt.Errorf("error getting job %q: %w", fqn, err)
		}

		switch job.CurrentState {
		case pb.JobState_JOB_STATE_UNKNOWN:
			job.CurrentState = pb.JobState_JOB_STATE_QUEUED
			job.CurrentStateTime = timestamppb.New(now)

			job.Environment = &pb.Environment{
				Dataset: "bigquery.googleapis.com/cloud_dataflow",
			}

			job.PipelineDescription = &pb.PipelineDescription{}

			job.Labels = map[string]string{
				"goog-dataflow-provided-template-name": "file_format_conversion",
				"goog-dataflow-provided-template-type": "flex",
			}

		default:
			return fmt.Errorf("unexpected state for job, got=%q, expected=JobState_JOB_STATE_UNKNOWN: %v", fqn, job.CurrentState)
		}

		if err := r.storage.Update(ctx, fqn, job); err != nil {
			return fmt.Errorf("error updating job %q: %v", fqn, err)
		}
	}
	time.Sleep(2 * time.Second)
	{
		now := time.Now()

		job := &pb.Job{}
		if err := r.storage.Get(ctx, fqn, job); err != nil {
			return fmt.Errorf("error getting job %q: %w", fqn, err)
		}

		switch job.CurrentState {
		case pb.JobState_JOB_STATE_QUEUED:
			job.CurrentState = pb.JobState_JOB_STATE_PENDING
			job.CurrentStateTime = timestamppb.New(now)
			job.JobMetadata = &pb.JobMetadata{}
			job.Steps = []*pb.Step{{Name: "placeholder"}}
			job.Type = pb.JobType_JOB_TYPE_BATCH

			job.Environment.Experiments = buildExperiments()

		default:
			return fmt.Errorf("unexpected state for job, got=%q, expected=JobState_JOB_STATE_QUEUED: %v", fqn, job.CurrentState)
		}

		if err := r.storage.Update(ctx, fqn, job); err != nil {
			return fmt.Errorf("error updating job %q: %v", fqn, err)
		}
	}
	// This interval is short, we don't usually catch it...
	time.Sleep(0 * time.Second)
	{
		now := time.Now()

		job := &pb.Job{}
		if err := r.storage.Get(ctx, fqn, job); err != nil {
			return fmt.Errorf("error getting job %q: %w", fqn, err)
		}

		switch job.CurrentState {
		case pb.JobState_JOB_STATE_PENDING:
			job.CurrentState = pb.JobState_JOB_STATE_RUNNING
			job.CurrentStateTime = timestamppb.New(now)
			job.StageStates = []*pb.ExecutionStageState{{ExecutionStageName: "placeholder"}}

			sdkPipelineOptions := buildSDKPipelineOptions(project, job)
			sdkPipelineOptionsVal, err := structpb.NewStruct(sdkPipelineOptions)
			if err != nil {
				return fmt.Errorf("building structpb for sdkPipelineOptions: %v", err)
			}
			job.Environment.SdkPipelineOptions = sdkPipelineOptionsVal

			job.Environment.ServiceAccountEmail = replacePlaceholders("${projectNumber}-compute@developer.gserviceaccount.com", project, job)
			job.Environment.ShuffleMode = pb.ShuffleMode_SERVICE_BASED
			job.Environment.TempStoragePrefix = replacePlaceholders("storage.googleapis.com/dataflow-staging-us-central1-${projectNumber}/tmp", project, job)
			job.Environment.UserAgent = &structpb.Struct{}
			job.Environment.Version = &structpb.Struct{}
			job.Environment.WorkerPools = []*pb.WorkerPool{{Kind: "placeholder"}}
		default:
			return fmt.Errorf("unexpected state for job, got=%q, expected=JobState_JOB_STATE_PENDING: %v", fqn, job.CurrentState)
		}

		if err := r.storage.Update(ctx, fqn, job); err != nil {
			return fmt.Errorf("error updating job %q: %v", fqn, err)
		}
	}

	return nil
}

func buildExperiments() []string {
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

	return experiments
}

func buildSDKPipelineOptions(project *projects.ProjectData, job *pb.Job) map[string]any {
	return map[string]any{
		"display_data": buildDisplayData(job),
		"options":      buildOptions(project, job),
	}
}

func buildOptions(project *projects.ProjectData, job *pb.Job) map[string]any {
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
