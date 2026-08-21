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
// proto.message: google.cloud.hypercomputecluster.v1.Cluster
// api.group: hypercomputecluster.cnrm.cloud.google.com

package hypercomputecluster

import (
	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(hypercomputeclusterclusterFuzzer())
}

func hypercomputeclusterclusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		HypercomputeClusterClusterSpec_FromProto, HypercomputeClusterClusterSpec_ToProto,
		HypercomputeClusterClusterObservedState_FromProto, HypercomputeClusterClusterObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".network_resources")
	f.SpecField(".storage_resources")
	f.SpecField(".compute_resources")
	f.SpecField(".orchestrator")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".reconciling")
	f.StatusField(".orchestrator.slurm.login_nodes.instances")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.FilterSpec = func(in *pb.Cluster) {
		for _, sr := range in.StorageResources {
			if sr != nil && sr.GetConfig() != nil {
				if nb := sr.GetConfig().GetNewBucket(); nb != nil {
					if scOneof, ok := nb.GetOption().(*pb.NewBucketConfig_StorageClass_); ok && scOneof != nil {
						if scOneof.StorageClass == pb.NewBucketConfig_STORAGE_CLASS_UNSPECIFIED {
							scOneof.StorageClass = pb.NewBucketConfig_STANDARD
						}
					}
				}
			}
		}
	}

	return f
}
