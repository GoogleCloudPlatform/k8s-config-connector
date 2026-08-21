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

package benchmark

type Resource struct {
	Kind    string
	RawFile string
	Proto   string
}

// Ground Truth resources
// TODO: We should exclude these ground truth from the gemnimi input datapoints.
var GroundTruthResources = []Resource{
	{
		Kind:    "BigQueryConnectionConnection",
		RawFile: "bigqueryconnection/v1beta1/connection_types.go",
		Proto:   "google.cloud.bigquery.connection.v1.Connection",
	},
	{
		Kind:    "KMSAutokeyConfig",
		RawFile: "kms/v1beta1/autokeyconfig_types.go",
		Proto:   "google.cloud.kms.v1.AutokeyConfig",
	},
	{
		Kind:    "BigQueryDataTransferConfig",
		RawFile: "bigquerydatatransfer/v1beta1/transferconfig_types.go",
		Proto:   "google.cloud.bigquery.datatransfer.v1.TransferConfig",
	},
	{
		Kind:    "CloudBuildWorkerPool",
		RawFile: "cloudbuild/v1beta1/workerpool_types.go",
		Proto:   "google.devtools.cloudbuild.v1.WorkerPool",
	},
	{
		Kind:    "DataformRepository",
		RawFile: "dataform/v1beta1/repository_types.go",
		Proto:   "google.cloud.dataform.v1beta1.Repository",
	},
}
