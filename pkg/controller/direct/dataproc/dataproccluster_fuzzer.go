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

// +tool:fuzz-gen
// proto.message: google.cloud.dataproc.v1.Cluster
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocClusterFuzzer())
}

func dataprocClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		DataprocClusterSpec_v1beta1_FromProto, DataprocClusterSpec_v1beta1_ToProto,
		DataprocClusterStatus_v1beta1_FromProto, DataprocClusterStatus_v1beta1_ToProto,
	)

	// Top level metadata/identity fields not in KRM spec
	f.UnimplementedFields.Insert(".project_id")
	f.UnimplementedFields.Insert(".cluster_name")
	f.UnimplementedFields.Insert(".labels")

	// Kubernetes cluster config is not fully triaged or fuzzed
	f.Unimplemented_NotYetTriaged(".virtual_cluster_config.kubernetes_cluster_config")

	// Spec fields
	f.SpecFields.Insert(".config")
	f.SpecFields.Insert(".virtual_cluster_config")

	// Status fields
	f.StatusFields.Insert(".config")
	f.StatusFields.Insert(".status")
	f.StatusFields.Insert(".status_history")
	f.StatusFields.Insert(".cluster_uuid")
	f.StatusFields.Insert(".metrics")

	return f
}
