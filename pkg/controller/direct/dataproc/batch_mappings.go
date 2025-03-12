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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExecutionConfig_NetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_NetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_NetworkUri{
		NetworkUri: direct.ValueOf(in),
	}
	return out
}

func ExecutionConfig_SubnetworkUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.ExecutionConfig_SubnetworkUri {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig_SubnetworkUri{
		SubnetworkUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainJarFileUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainJarFileUri {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainJarFileUri{
		MainJarFileUri: direct.ValueOf(in),
	}
	return out
}

func SparkBatch_MainClass_ToProto(mapCtx *direct.MapContext, in *string) *pb.SparkBatch_MainClass {
	if in == nil {
		return nil
	}
	out := &pb.SparkBatch_MainClass{
		MainClass: direct.ValueOf(in),
	}
	return out
}
